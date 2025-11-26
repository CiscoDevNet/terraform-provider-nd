package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"time"

	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/net/html"
)

const ndAuthPayload = `{
	"userName": "%s",
	"userPasswd": "%s"
}`

// Default timeout for NGINX in ND is 90 Seconds.
// Allow the client to set a shorter or longer time depending on their
// environment
const DefaultReqTimeoutVal int = 100
const DefaultBackoffMinDelay int = 4
const DefaultBackoffMaxDelay int = 60
const DefaultBackoffDelayFactor float64 = 3

// Client is the main entry point
type Client struct {
	baseURL            *url.URL
	httpClient         *http.Client
	authToken          *Auth
	mutex              sync.Mutex
	username           string
	password           string
	insecure           bool
	proxyUrl           string
	proxyCreds         string
	domain             string
	skipLoggingPayload bool
	maxRetries         int64
	backoffMinDelay    int64
	backoffMaxDelay    int64
	backoffDelayFactor float64
}

// singleton implementation of a client
var clientImpl *Client

func initClient(clientUrl, username, password, proxyUrl, proxyCreds, loginDomain string, isInsecure bool, maxRetries int64) *Client {

	bUrl, err := url.Parse(clientUrl)
	if err != nil {
		// cannot move forward if url is undefined
		log.Fatal(err)
	}

	client := &Client{
		baseURL:    bUrl,
		username:   username,
		httpClient: http.DefaultClient,
		password:   password,
		insecure:   isInsecure,
		proxyUrl:   proxyUrl,
		proxyCreds: proxyCreds,
		domain:     loginDomain,
		maxRetries: maxRetries,
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       client.insecure,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS13,
		},
	}

	if client.proxyUrl != "" {
		transport = client.configProxy(transport)
	}

	client.httpClient = &http.Client{
		Transport: transport,
	}

	return client
}

// GetClient returns a singleton
func GetClient(clientUrl, username, password, proxyUrl, proxyCreds, loginDomain string, isInsecure bool, maxRetries int64) *Client {
	if clientImpl == nil {
		return initClient(clientUrl, username, password, proxyUrl, proxyCreds, loginDomain, isInsecure, maxRetries)
	}
	return clientImpl
}

func (c *Client) configProxy(transport *http.Transport) *http.Transport {
	log.Printf("[DEBUG]: Using Proxy Server: %s ", c.proxyUrl)
	pUrl, err := url.Parse(c.proxyUrl)
	if err != nil {
		log.Fatal(err)
	}
	transport.Proxy = http.ProxyURL(pUrl)

	if c.proxyCreds != "" {
		basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(c.proxyCreds))
		transport.ProxyConnectHeader = http.Header{}
		transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
	}
	return transport
}

func (c *Client) makeFullUrl(method string, path string) (string, error) {
	path = strings.TrimLeft(path, "/")
	path = fmt.Sprintf("/%v", path)
	url, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	if method == "PATCH" {
		validateString := url.Query()
		validateString.Set("validate", "false")
		url.RawQuery = validateString.Encode()
	}
	fURL := c.baseURL.ResolveReference(url)
	return fURL.String(), nil
}

func (c *Client) MakeRestRequest(method string, path string, body *gabs.Container, authenticated bool, skipLoggingPayload bool) (*http.Request, error) {
	fURL, err := c.makeFullUrl(method, path)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, fURL, nil)
	} else {
		req, err = http.NewRequest(method, fURL, bytes.NewBuffer((body.Bytes())))
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	log.Printf("[DEBUG] HTTP request %s %s", method, path)

	if skipLoggingPayload {
		log.Printf("HTTP request %s %s", method, path)
	} else {
		log.Printf("HTTP request %s %s %v", method, path, req)
	}

	if authenticated {
		req, err = c.InjectAuthenticationHeader(req, path)
		if err != nil {
			return req, err
		}
	}

	if !skipLoggingPayload {
		log.Printf("HTTP request after injection %s %s %v", method, path, req)
	}

	return req, nil
}

func (c *Client) Authenticate() error {
	body, err := gabs.ParseJSON([]byte(fmt.Sprintf(ndAuthPayload, c.username, c.password)))
	if err != nil {
		return err
	}

	if c.domain != "" {
		body.Set(c.domain, "domain")
	}

	req, err := c.MakeRestRequest("POST", "/login", body, false, c.skipLoggingPayload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	obj, _, err := c.Do(req, c.skipLoggingPayload)

	if err != nil {
		return err
	}

	if obj == nil {
		return errors.New("Empty response")
	}

	token := obj.S("token").String()

	if token == "" || token == "{}" {
		return errors.New("Invalid Username or Password")
	}

	if c.authToken == nil {
		c.authToken = &Auth{}
	}

	c.authToken.Token = token
	c.authToken.CalculateExpiry(1200) //refreshTime=1200 Sec

	return nil
}

func (c *Client) Do(req *http.Request, skipLoggingPayload bool) (*gabs.Container, *http.Response, error) {
	log.Printf("[DEBUG] Beginning DO method %s", req.URL.String())
	log.Printf("[TRACE] HTTP Request Method and URL: %s %s", req.Method, req.URL.String())

	var body []byte
	if req.Body != nil && c.maxRetries != 0 {
		body, _ = io.ReadAll(req.Body)
	}

	for attempts := int64(0); ; attempts++ {
		if c.maxRetries != 0 {
			req.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		if !skipLoggingPayload {
			log.Printf("[TRACE] HTTP Request Body: %v", req.Body)
		}

		resp, err := c.httpClient.Do(req)

		if err != nil {
			if ok := c.backoff(attempts); !ok {
				log.Printf("[ERROR] HTTP Connection error occured: %+v", err)
				log.Printf("[DEBUG] Exit from Do method")
				return nil, nil, errors.New(fmt.Sprintf("Failed to connect to ND. Verify that you are connecting to an ND.\nError message: %+v", err))
			} else {
				log.Printf("[ERROR] HTTP Connection failed: %s, retries: %v", err, attempts)
				continue
			}
		}

		if !skipLoggingPayload {
			log.Printf("[TRACE] HTTP Response: %d %s %v", resp.StatusCode, resp.Status, resp)
		} else {
			log.Printf("[TRACE] HTTP Response: %d %s", resp.StatusCode, resp.Status)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, err
		}

		bodyStr := string(bodyBytes)
		err = resp.Body.Close()
		if err != nil {
			return nil, nil, err
		}

		if !skipLoggingPayload {
			log.Printf("[DEBUG] HTTP response unique string %s %s %s", req.Method, req.URL.String(), bodyStr)
		}

		if req.Method == "POST" && resp.StatusCode == 200 && bodyStr == "" {
			log.Printf("[DEBUG] Exit from do method")
			return nil, resp, nil
		} else if req.Method != "DELETE" && resp.StatusCode != 204 {
			obj, err := gabs.ParseJSON(bodyBytes)
			if err != nil {
				log.Printf("Error occurred while json parsing %+v", err)
				return nil, resp, err
			}
			log.Printf("[DEBUG] Exit from do method")
			return obj, resp, err
		} else if req.Method == "DELETE" && resp.StatusCode == 204 {
			log.Printf("[DEBUG] Exit from do method")
			return nil, resp, nil
		} else if resp.StatusCode == 204 {
			log.Printf("[DEBUG] Exit from do method")
			return nil, nil, nil
		} else {
			if ok := c.backoff(attempts); !ok {
				obj, err := gabs.ParseJSON(bodyBytes)
				if err != nil {
					log.Printf("[ERROR] Error occured while json parsing: %+v with HTTP StatusCode 405, 500-504", err)

					// If nginx is too busy or the page is not found, ND's nginx will response with an HTML doc instead of a JSON Response.
					// In those cases, parse the HTML response for the message and return that to the user
					htmlErr := c.checkHtmlResp(bodyStr)
					log.Printf("[ERROR] Error occured while json parsing: %s", htmlErr.Error())
					log.Printf("[DEBUG] Exit from Do method")
					return nil, resp, errors.New(fmt.Sprintf("Failed to parse JSON response from: %s. Verify that you are connecting to an ND.\nHTTP response status: %s\nMessage: %s", req.URL.String(), resp.Status, htmlErr))
				}
				log.Printf("[DEBUG] Exit from Do method")
				return obj, resp, nil
			} else {
				log.Printf("[ERROR] HTTP Request failed: StatusCode %v, Retries: %v", resp.StatusCode, attempts)
				continue
			}
		}
	}
}

// func (c *Client) DoRestRequest(ctx context.Context, diags *diag.Diagnostics, client *Client, path, method string, payload *gabs.Container) *gabs.Container {
func (c *Client) DoRestRequest(ctx context.Context, diags *diag.Diagnostics, path, method string, payload *gabs.Container) *gabs.Container {
	if !strings.HasPrefix("/", path) {
		path = fmt.Sprintf("/%s", path)
	}
	var restRequest *http.Request
	var err error

	restRequest, err = c.MakeRestRequest(method, path, payload, true, c.skipLoggingPayload)
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := c.Do(restRequest, c.skipLoggingPayload)

	// Return nil when the object is not found and ignore 404 not found error
	// The resource ID will be set it to nil and the state file content will be deleted when the object is not found
	if restResponse.StatusCode == 404 {
		return nil
	}

	if restResponse != nil && cont.Data() != nil && (restResponse.StatusCode != 200 && restResponse.StatusCode != 201) {
		diags.AddError(
			fmt.Sprintf("The %s %s rest request failed.", method, path),
			fmt.Sprintf("Code: %d Response: %s, err: %s. Please report this issue to the provider developers.", restResponse.StatusCode, cont.Data().(map[string]interface{})["errors"], err),
		)
		tflog.Debug(ctx, fmt.Sprintf("%v", cont.Search("errors")))
		return nil
	} else if err != nil {
		diags.AddError(
			fmt.Sprintf("The %s %s rest request failed.", method, path),
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return cont
}

func (c *Client) backoff(attempts int64) bool {
	log.Printf("[DEBUG] Begining backoff method: attempts %v on %v", attempts, c.maxRetries)
	if attempts >= c.maxRetries {
		log.Printf("[DEBUG] Exit from backoff method with return value false")
		return false
	}

	minDelay := time.Duration(DefaultBackoffMinDelay) * time.Second
	if c.backoffMinDelay != 0 {
		minDelay = time.Duration(c.backoffMinDelay) * time.Second
	}

	maxDelay := time.Duration(DefaultBackoffMaxDelay) * time.Second
	if c.backoffMaxDelay != 0 {
		maxDelay = time.Duration(c.backoffMaxDelay) * time.Second
	}

	factor := DefaultBackoffDelayFactor
	if c.backoffDelayFactor != 0 {
		factor = c.backoffDelayFactor
	}

	min := float64(minDelay)
	backoff := min * math.Pow(factor, float64(attempts))
	if backoff > float64(maxDelay) {
		backoff = float64(maxDelay)
	}
	backoff = (rand.Float64()/2+0.5)*(backoff-min) + min
	backoffDuration := time.Duration(backoff)
	log.Printf("[TRACE] Starting sleeping for %v", backoffDuration.Round(time.Second))
	time.Sleep(backoffDuration)
	log.Printf("[DEBUG] Exit from backoff method with return value true")
	return true
}

// If nginx is too busy or the page is not found, ND's nginx will response with an HTML doc instead of a JSON Response.
// In those cases, parse the HTML response for the message and return that to the user
//
// Sample Response Body: https://github.com/nginx/nginx-releases/blob/master/html/50x.html
// <!DOCTYPE html>
// <html>
// <head>
// <title>Error</title>
// <style>
//
//	body {
//	    width: 35em;
//	    margin: 0 auto;
//	    font-family: Tahoma, Verdana, Arial, sans-serif;
//	}
//
// </style>
// </head>
// <body>
// <h1>An error occurred.</h1>
// <p>Sorry, the page you are looking for is currently unavailable.<br/>
// Please try again later.</p>
// <p>If you are the system administrator of this resource then you should check
// the <a href="http://nginx.org/r/error_log">error log</a> for details.</p>
// <p><em>Faithfully yours, nginx.</em></p>
// </body>
// </html>
//
// Sample return error:
// An error occurred. Sorry, the page you are looking for is currently unavailable. If you are the system administrator of this
// resource then you should check the error log for details. Faithfully yours, nginx.
func (c *Client) checkHtmlResp(body string) error {
	reader := strings.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	errStr := ""
	prevTag := ""
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			break
		}
		tag, _ := tokenizer.TagName()
		token := tokenizer.Token()

		if prevTag == "a" || prevTag == "p" || prevTag == "body" {
			data := strings.TrimSpace(token.Data)
			if data == "" {
				continue
			}
			if errStr == "" {
				errStr = data
			} else {
				errStr = errStr + " " + data
			}
		}
		prevTag = string(tag)
	}
	if errStr == "" {
		errStr = "Empty ND HTML Response"
	}
	log.Printf("[DEBUG] HTML Error Parsing Result: %s", errStr)
	return fmt.Errorf(errStr)
}
