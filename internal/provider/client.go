package provider

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"

	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ndAuthPayload = `{
	"userName": "%s",
	"userPasswd": "%s"
}`

// Client is the main entry point
type Client struct {
	BaseURL            *url.URL
	httpClient         *http.Client
	AuthToken          *Auth
	Mutex              sync.Mutex
	username           string
	password           string
	insecure           bool
	proxyUrl           string
	proxyCreds         string
	domain             string
	version            string
	skipLoggingPayload bool
}

// singleton implementation of a client
var clientImpl *Client

type Option func(*Client)

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}

func Password(password string) Option {
	return func(client *Client) {
		client.password = password
	}
}

func ProxyUrl(pUrl string) Option {
	return func(client *Client) {
		client.proxyUrl = pUrl
	}
}

func ProxyCreds(pcreds string) Option {
	return func(client *Client) {
		client.proxyCreds = pcreds
	}
}

func Domain(domain string) Option {
	return func(client *Client) {
		client.domain = domain
	}
}

func Version(version string) Option {
	return func(client *Client) {
		client.version = version
	}
}

func SkipLoggingPayload(skipLoggingPayload bool) Option {
	return func(client *Client) {
		client.skipLoggingPayload = skipLoggingPayload
	}
}

func initClient(clientUrl, username string, options ...Option) *Client {
	var transport *http.Transport
	bUrl, err := url.Parse(clientUrl)
	if err != nil {
		// cannot move forward if url is undefined
		log.Fatal(err)
	}
	client := &Client{
		BaseURL:    bUrl,
		username:   username,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	transport = client.useInsecureHTTPClient(client.insecure)
	if client.proxyUrl != "" {
		transport = client.configProxy(transport)
	}

	client.httpClient = &http.Client{
		Transport: transport,
	}

	return client
}

// GetClient returns a singleton
func GetClient(clientUrl, username string, options ...Option) *Client {
	if clientImpl == nil {
		return initClient(clientUrl, username, options...)
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

func (c *Client) useInsecureHTTPClient(insecure bool) *http.Transport {
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
			InsecureSkipVerify:       insecure,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS13,
		},
	}

	return transport
}

func (c *Client) MakeRestRequest(method string, path string, body *gabs.Container, authenticated bool, skipLoggingPayload bool) (*http.Request, error) {
	if path != "/login" {
		if strings.HasPrefix(path, "/") {
			path = path[1:]
		}
		path = fmt.Sprintf("/%v", path)
	}
	url, err := url.Parse(path)

	if err != nil {
		return nil, err
	}
	if method == "PATCH" {
		validateString := url.Query()
		validateString.Set("validate", "false")
		url.RawQuery = validateString.Encode()
	}
	fURL := c.BaseURL.ResolveReference(url)

	var req *http.Request
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, fURL.String(), nil)
	} else {
		req, err = http.NewRequest(method, fURL.String(), bytes.NewBuffer((body.Bytes())))
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

// Authenticate is used to
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

	if c.AuthToken == nil {
		c.AuthToken = &Auth{}
	}

	c.AuthToken.Token = token
	c.AuthToken.CalculateExpiry(1200) //refreshTime=1200 Sec

	return nil
}

func (c *Client) Do(req *http.Request, skipLoggingPayload bool) (*gabs.Container, *http.Response, error) {
	log.Printf("[DEBUG] Beginning DO method %s", req.URL.String())
	log.Printf("[TRACE] HTTP Request Method and URL: %s %s", req.Method, req.URL.String())

	if !skipLoggingPayload {
		log.Printf("[TRACE] HTTP Request Body: %v", req.Body)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
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

	if req.Method != "DELETE" && resp.StatusCode != 204 {
		obj, err := gabs.ParseJSON(bodyBytes)
		if err != nil {
			log.Printf("Error occurred while json parsing %+v", err)
			return nil, resp, err
		}
		log.Printf("[DEBUG] Exit from do method")
		return obj, resp, err
	} else if req.Method == "DELETE" && resp.StatusCode == 204 {
		return nil, resp, nil
	} else if resp.StatusCode == 204 {
		return nil, nil, nil
	} else {
		return nil, resp, err
	}
}

func DoRestRequest(ctx context.Context, diags *diag.Diagnostics, client *Client, path, method string, payload *gabs.Container) *gabs.Container {
	if !strings.HasPrefix("/", path) {
		path = fmt.Sprintf("/%s", path)
	}
	var restRequest *http.Request
	var err error

	restRequest, err = client.MakeRestRequest(method, path, payload, true, client.skipLoggingPayload)
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := client.Do(restRequest, client.skipLoggingPayload)

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
