package provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	// provider is built and run locally, and "test" when running acceptance
	// testing.
	version string
}

// ndProviderModel describes the provider data model.
type ndProviderModel struct {
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	URL         types.String `tfsdk:"url"`
	LoginDomain types.String `tfsdk:"login_domain"`
	IsInsecure  types.Bool   `tfsdk:"insecure"`
	ProxyUrl    types.String `tfsdk:"proxy_url"`
	ProxyCreds  types.String `tfsdk:"proxy_creds"`
	MaxRetries  types.Int64  `tfsdk:"retries"`
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
				Description: "Username for the Nexus Dashboard Account. This can also be set as the ND_USERNAME environment variable.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"password": schema.StringAttribute{
				Description: "Password for the Nexus Dashboard Account. This can also be set as the ND_PASSWORD environment variable.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"url": schema.StringAttribute{
				Description: "URL of the Cisco Nexus Dashboard web interface. This can also be set as the ND_URL environment variable.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^(https?)://[^\s/$.?#].[^\s]*$`),
						"The url must contain only alphanumeric characters",
					),
				},
			},
			"login_domain": schema.StringAttribute{
				Description: "Login domain for the Nexus Dashboard Account. This can also be set as the ND_LOGIN_DOMAIN environment variable. Defaults to `DefaultAuth`.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"insecure": schema.BoolAttribute{
				Description: "Allow insecure HTTPS client. This can also be set as the ND_INSECURE environment variable. Defaults to `true`.",
				Optional:    true,
			},
			"proxy_url": schema.StringAttribute{
				Description: "Proxy Server URL with port number. This can also be set as the ND_PROXY_URL environment variable.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^(https?)://[^\s/$.?#].[^\s]*$`),
						"The proxy_url must contain only alphanumeric characters",
					),
				},
			},
			"proxy_creds": schema.StringAttribute{
				Description: "Proxy server credentials in the form of username:password. This can also be set as the ND_PROXY_CREDS environment variable.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"retries": schema.Int64Attribute{
				Description: "Number of retries for REST API calls. This can also be set as the ND_RETRIES environment variable. Defaults to `2`.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 10),
				},
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
	isInsecure := getBoolAttribute(resp, data.IsInsecure, "ND_INSECURE", true)
	proxyUrl := getStringAttribute(data.ProxyUrl, "ND_PROXY_URL")
	url := getStringAttribute(data.URL, "ND_URL")
	loginDomain := getStringAttribute(data.LoginDomain, "ND_LOGIN_DOMAIN")
	proxyCreds := getStringAttribute(data.ProxyCreds, "ND_PROXY_CREDS")
	maxRetries := int64(getIntAttribute(resp, data.MaxRetries, "ND_RETRIES", 2))

	if username == "" {
		resp.Diagnostics.AddError(
			"Username not provided",
			"Username must be provided for the ND provider",
		)
	}

	if password == "" {
		resp.Diagnostics.AddError(
			"Authentication details not provided",
			"Password must be provided for the ND provider",
		)
	}

	if loginDomain == "" {
		loginDomain = "DefaultAuth"
	}

	ndClient := GetClient(url, username, password, proxyUrl, proxyCreds, loginDomain, isInsecure, maxRetries)

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
