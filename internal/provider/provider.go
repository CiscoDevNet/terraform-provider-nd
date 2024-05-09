package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ciscoecosystem/mso-go-client/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &ndProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ndProvider{
			version: version,
		}
	}
}

// ndProvider is the provider implementation.
type ndProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ndProviderModel describes the provider data model.
type ndProviderModel struct {
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	IsInsecure types.String `tfsdk:"insecure"`
	ProxyUrl   types.String `tfsdk:"proxy_url"`
	URL        types.String `tfsdk:"url"`
	Domain     types.String `tfsdk:"domain"`
	Platform   types.String `tfsdk:"platform"`
	PrivateKey types.String `tfsdk:"private_key"`
	Certname   types.String `tfsdk:"cert_name"`
	ProxyCreds types.String `tfsdk:"proxy_creds"`
	MaxRetries types.String `tfsdk:"retries"`
}

// Metadata returns the provider type name.
func (p *ndProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nd"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *ndProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description: "Username for the ND Account. This can also be set as the ND_USERNAME environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for the ND Account. This can also be set as the ND_PASSWORD environment variable.",
				Optional:    true,
			},
			// temporary used schema attributes until muxing is removed and the correct schema attributes can be used which are commented out below
			"insecure": schema.StringAttribute{
				Description: "Allow insecure HTTPS client. This can also be set as the ND_INSECURE environment variable. Defaults to `true`.",
				Optional:    true,
			},
			"proxy_url": schema.StringAttribute{
				Description: "Proxy Server URL with port number. This can also be set as the ND_PROXY_URL environment variable.",
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Description: "URL of the Cisco ND web interface. This can also be set as the ND_URL environment variable.",
				Optional:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "Private key path for signature calculation. This can also be set as the ND_PRIVATE_KEY environment variable.",
				Optional:    true,
			},
			"domain": schema.StringAttribute{
				Description: "URL of the Cisco ND web interface. This can also be set as the ND_DOMAIN environment variable.",
				Optional:    true,
			},
			"platform": schema.StringAttribute{
				Description: "Platform of the Cisco ND web interface. This can also be set as the ND_PLATFORM environment variable.",
				Optional:    true,
			},
			"cert_name": schema.StringAttribute{
				Description: "Certificate name for the User in Cisco ND. This can also be set as the ND_CERT_NAME environment variable.",
				Optional:    true,
			},
			"proxy_creds": schema.StringAttribute{
				Description: "Proxy server credentials in the form of username:password. This can also be set as the ND_PROXY_CREDS environment variable.",
				Optional:    true,
			},
			"retries": schema.StringAttribute{
				Description: "Number of retries for REST API calls. This can also be set as the ND_RETRIES environment variable. Defaults to `2`.",
				Optional:    true,
			},
		},
	}
}

// Configure prepares a ND API client for data sources and resources.
func (p *ndProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data ndProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	username := getStringAttribute(data.Username, "ND_USERNAME")
	password := getStringAttribute(data.Password, "ND_PASSWORD")
	isInsecure := stringToBool(resp, "insecure", getStringAttribute(data.IsInsecure, "ND_INSECURE"), true)
	proxyUrl := getStringAttribute(data.ProxyUrl, "ND_PROXY_URL")
	url := getStringAttribute(data.URL, "ND_URL")
	domain := getStringAttribute(data.Domain, "ND_DOMAIN")
	platform := getStringAttribute(data.Platform, "ND_PLATFORM")
	privateKey := getStringAttribute(data.PrivateKey, "ND_PRIVATE_KEY")
	certName := getStringAttribute(data.Certname, "ND_CERT_NAME")
	proxyCreds := getStringAttribute(data.ProxyCreds, "ND_PROXY_CREDS")
	maxRetries := stringToInt(resp, "retries", getStringAttribute(data.MaxRetries, "ND_RETRIES"), 2)

	if username == "" {
		resp.Diagnostics.AddError(
			"Username not provided",
			"Username must be provided for the ND provider",
		)
	}

	if password == "" && (privateKey == "" || certName == "") {
		resp.Diagnostics.AddError(
			"Authentication details not provided",
			"Either 'password' OR 'private_key' and 'cert_name' must be provided for the ND provider",
		)
	}

	if url == "" {
		resp.Diagnostics.AddError(
			"Url not provided",
			"Url must be provided for the ND provider",
		)
	} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		resp.Diagnostics.AddError(
			"Incorrect url prefix",
			fmt.Sprintf("Url '%s' must start with 'http://' or 'https://'", url),
		)
	}

	// temporary conditional until muxing is removed and the correct retries schema attribute is used
	if maxRetries < 0 || maxRetries > 9 {
		resp.Diagnostics.AddError(
			"Incorrect retry amount",
			fmt.Sprintf("Retries must be between 0 and 9 inclusive, got: %d", maxRetries),
		)
	}

	var ndClient *client.Client
	if password != "" {
		ndClient = client.GetClient(url, username, client.Password(password), client.Insecure(isInsecure), client.ProxyUrl(proxyUrl), client.ProxyCreds(proxyCreds), client.Domain(domain), client.Platform(platform))
	} else {
		ndClient = nil
	}

	resp.DataSourceData = ndClient
	resp.ResourceData = ndClient
}

func (p *ndProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewVersionDataSource,
		NewSiteDataSource,
	}
}

func (p *ndProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSiteResource,
	}
}

func getStringAttribute(attribute basetypes.StringValue, envKey string) string {

	if attribute.IsNull() {
		return os.Getenv(envKey)
	}
	return attribute.ValueString()

}

// Placeholder for future use when correct type can be used since muxing has been removed
func getBoolAttribute(resp *provider.ConfigureResponse, attribute basetypes.BoolValue, envKey string, defaultValue bool) bool {

	if attribute.IsNull() {
		envValue := os.Getenv(envKey)
		if envValue == "" {
			return defaultValue
		}
		boolValue, err := strconv.ParseBool(envValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", envValue),
				fmt.Sprintf("A boolean value must be provided for %s", envKey),
			)
		}
		return boolValue
	}
	return attribute.ValueBool()

}

// Placeholder for future use when correct type can be used since muxing has been removed
func getIntAttribute(resp *provider.ConfigureResponse, attribute basetypes.Int64Value, envKey string, defaultValue int) int {

	if attribute.IsNull() {
		envValue := os.Getenv(envKey)
		if envValue == "" {
			return defaultValue
		}
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", envValue),
				fmt.Sprintf("A integer value must be provided for %s", envKey),
			)
		}
		return intValue
	}
	return int(attribute.ValueInt64())
}

// Temporary function so correct type can be used untill since muxing has been removed
func stringToBool(resp *provider.ConfigureResponse, attributeName, stringValue string, boolValue bool) bool {
	var err error
	if stringValue != "" {
		boolValue, err = strconv.ParseBool(stringValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", stringValue),
				fmt.Sprintf("A boolean value must be provided for the %s attribute", attributeName),
			)
		}
	}
	return boolValue
}

// Temporary function so correct type can be used untill since muxing has been removed
func stringToInt(resp *provider.ConfigureResponse, attributeName, stringValue string, intValue int) int {
	var err error
	if stringValue != "" {
		intValue, err = strconv.Atoi(stringValue)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Invalid input '%s'", stringValue),
				fmt.Sprintf("A integer must be provided for the %s attribute", attributeName),
			)
		}
	}
	return intValue
}
