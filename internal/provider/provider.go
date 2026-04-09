package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"
	"time"

	_ "terraform-provider-nd/internal/manage/resource_fabric_vxlan"
	_ "terraform-provider-nd/internal/manage/resource_inventory_switch"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	nd "github.com/netascode/go-nd"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &NexusDashboardProvider{
		version: "",
	}
)

// NexusDashboardProvider is the provider implementation.
type NexusDashboardProvider struct {
	// version is set to the provider version on release.
	version string
}

// NDClient is now defined in client.go

// New returns a function that initializes and returns a new NexusDashboardProvider.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NexusDashboardProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *NexusDashboardProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nd"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *NexusDashboardProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = NdProviderSchema(ctx)
}

// Configure prepares a Nexus Dashboard API client for data sources and resources.
func (p *NexusDashboardProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Nexus Dashboard client")

	// Retrieve provider data from configuration
	var config NdModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practising defaults for the provider configuration,
	// use environment variables to fill in the values.
	if config.Url.IsUnknown() || config.Url.IsNull() {
		url := os.Getenv("ND_URL")
		config.Url = types.StringValue(url)
	}

	if config.Username.IsUnknown() || config.Username.IsNull() {
		username := os.Getenv("ND_USERNAME")
		config.Username = types.StringValue(username)
	}

	if config.Password.IsUnknown() || config.Password.IsNull() {
		password := os.Getenv("ND_PASSWORD")
		config.Password = types.StringValue(password)
	}

	if config.Domain.IsUnknown() || config.Domain.IsNull() {
		domain := os.Getenv("ND_DOMAIN")
		config.Domain = types.StringValue(domain)
	}

	if config.Insecure.IsUnknown() || config.Insecure.IsNull() {
		insecureStr := os.Getenv("ND_INSECURE")
		insecure, err := strconv.ParseBool(insecureStr)
		if err != nil {
			// Default to false if environment variable is not set or invalid
			insecure = false
		}
		config.Insecure = types.BoolValue(insecure)
	}

	if config.Timeout.IsUnknown() || config.Timeout.IsNull() {
		timeoutStr := os.Getenv("ND_TIMEOUT")
		var timeout int64 = 60 // Default timeout
		if timeoutStr != "" {
			var err error
			timeout, err = strconv.ParseInt(timeoutStr, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError(
					"Invalid ND_TIMEOUT Environment Variable",
					fmt.Sprintf("Unable to parse ND_TIMEOUT environment variable: %v", err),
				)
				return
			}
		}
		config.Timeout = types.Int64Value(timeout)
	}

	// Validate required configuration values
	if config.Url.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Nexus Dashboard URL",
			"The provider cannot create the Nexus Dashboard API client without a URL. "+
				"Set the URL value in the configuration or use the ND_URL environment variable.",
		)
	}

	if config.Username.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Nexus Dashboard Username",
			"The provider cannot create the Nexus Dashboard API client without a username. "+
				"Set the username value in the configuration or use the ND_USERNAME environment variable.",
		)
	}

	if config.Password.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Nexus Dashboard Password",
			"The provider cannot create the Nexus Dashboard API client without a password. "+
				"Set the password value in the configuration or use the ND_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Extract values from the configuration
	url := config.Url.ValueString()
	username := config.Username.ValueString()
	password := config.Password.ValueString()
	domain := config.Domain.ValueString()
	insecure := config.Insecure.ValueBool()
	timeout := time.Duration(config.Timeout.ValueInt64()) * time.Second
	// Set up logging
	ctx = tflog.SetField(ctx, "nd_url", url)
	ctx = tflog.SetField(ctx, "nd_username", username)
	ctx = tflog.SetField(ctx, "nd_domain", domain)
	ctx = tflog.SetField(ctx, "nd_insecure", insecure)
	ctx = tflog.SetField(ctx, "nd_timeout", timeout)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "nd_password")

	tflog.Debug(ctx, "Creating Nexus Dashboard client")

	basePath := "/api/v1"

	// Create the shared API client
	client, err := nd.NewClient(url, basePath, username, password,
		domain, insecure, nd.MaxRetries(500),
		nd.RequestTimeout(time.Duration(timeout)))
	if err != nil {
		tflog.Error(ctx, "Error creating Nexus Dashboard client", map[string]interface{}{"error": err.Error()})
		resp.Diagnostics.AddError(
			"Unable to Create Nexus Dashboard API Client",
			"An unexpected error occurred when creating the Nexus Dashboard API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Nexus Dashboard Client Error: "+err.Error(),
		)
		return
	}

	ndClient := &NDClient{
		URL:       url,
		Username:  username,
		Password:  password,
		Domain:    domain,
		Insecure:  insecure,
		Timeout:   timeout,
		ApiClient: &client,
		NDModules: make(map[string]interface{}),
	}

	// Register module-specific clients (eager initialization)
	// Each team adds one line here for their module
	ndClient.NDModules[manage.ModuleKey] = manage.NewManage(&client)

	// ndClient.NDModules[onemanage.ModuleKey] = onemanage.NewClient(&client)

	// Make the Nexus Dashboard client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = ndClient
	resp.ResourceData = ndClient

	tflog.Info(ctx, "Configured Nexus Dashboard client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *NexusDashboardProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	dataSources := []func() datasource.DataSource{}

	// Get all manage data sources
	dataSources = append(dataSources, registry.GetAllDataSources()...)

	return dataSources
}

// Resources defines the resources implemented in the provider.
func (p *NexusDashboardProvider) Resources(_ context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{}

	// Get all manage resources
	resources = append(resources, registry.GetAllResources()...)

	return resources
}
