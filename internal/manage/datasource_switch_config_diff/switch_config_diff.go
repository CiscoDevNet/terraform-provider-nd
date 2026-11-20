package datasource_switch_config_diff

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

const ModuleKey = "manage"

// Compile-time checks ensure this type implements the provider framework's
// datasource interfaces, including the optional provider-configuration hook.
var (
	_ datasource.DataSource              = &switchConfigDiffDataSource{}
	_ datasource.DataSourceWithConfigure = &switchConfigDiffDataSource{}
)

// The provider's manage client is retained here so Read can call the ND API.
type switchConfigDiffDataSource struct {
	manageClient *manage.NexusDashboardManage
}

// The API response model is generated alongside the Terraform-facing model.
// Type alias for the API response model used by this data source.
type NDSwitchConfigDiffModel = NDFCSwitchConfigDiffModel

// The framework calls this when constructing the data source.
func NewSwitchConfigDiffDataSource() datasource.DataSource {
	return &switchConfigDiffDataSource{}
}

// Metadata supplies the Terraform type name used in configuration and state.
func (source *switchConfigDiffDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_config_diff"
}

// Schema publishes the generated input and computed-output attributes.
func (source *switchConfigDiffDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SwitchConfigDiffDataSourceSchema(ctx)
}

// Configure obtains the shared manage module from provider data and retains its
// client for subsequent Read calls. Diagnostics stop configuration on mismatch.
func (source *switchConfigDiffDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected DataSource Configure Type", fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	manageModule := client.GetModule(ModuleKey)
	if manageModule == nil {
		resp.Diagnostics.AddError("Manage Module Not Found", "The manage module was not registered with the provider.")
		return
	}
	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Manage Module Type", fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule))
		return
	}
	source.manageClient = manageClient
}

// Read validates the configured fabric and switch, requests their configuration
// diff from ND, converts the response to Terraform values, and writes state.
func (source *switchConfigDiffDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Decode the user's Terraform arguments into the GO model first.
	var data SwitchConfigDiffModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Require concrete non-blank identifiers before constructing an API request.
	if data.FabricName.IsNull() || data.FabricName.IsUnknown() || strings.TrimSpace(data.FabricName.ValueString()) == "" {
		resp.Diagnostics.AddError("Fabric Name Required", "The fabric_name attribute must contain a known, non-empty value to read switch configuration differences.")
		return
	}
	if data.SwitchId.IsNull() || data.SwitchId.IsUnknown() || strings.TrimSpace(data.SwitchId.ValueString()) == "" {
		resp.Diagnostics.AddError("Switch ID Required", "The switch_id attribute must contain a known, non-empty value to read switch configuration differences.")
		return
	}

	// Configure is responsible for storing a usable provider client.
	if source.manageClient == nil || source.manageClient.ApiClient == nil {
		resp.Diagnostics.AddError("Manage Client Not Configured", "The manage client must be configured before reading switch configuration differences.")
		return
	}

	// Select the switch-diff endpoint and query option using the input identifiers.
	fabricName := data.FabricName.ValueString()
	switchID := data.SwitchId.ValueString()
	diffAPI := api.NewConfigAPI(source.manageClient.ApiClient, fabricName)
	diffAPI.FabricName = fabricName
	diffAPI.SwitchId = switchID
	diffAPI.SetOperation(api.OpSwitchConfigDiff)
	diffAPI.SetQueryParams("forceShowRun=false")

	// Fetch the diff; API errors become Terraform diagnostics and leave state unset.
	responseData, err := diffAPI.Get()
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Switch Configuration Diff", fmt.Sprintf("Could not read configuration differences for switch %q in fabric %q: %v", switchID, fabricName, err))
		return
	}

	// Decode ND's JSON payload into its API-shaped model. Malformed payloads fail
	// here rather than producing partial or misleading Terraform state.
	var diff NDSwitchConfigDiffModel
	if err := json.Unmarshal(responseData, &diff); err != nil {
		resp.Diagnostics.AddError("Error Decoding Switch Configuration Diff", fmt.Sprintf("Could not decode configuration differences for switch %q in fabric %q: %v", switchID, fabricName, err))
		return
	}

	// Keep identifiers from configuration authoritative, then map all API fields
	// into Terraform framework values before committing the completed state.
	diff.FabricName = fabricName
	diff.SwitchId = switchID
	resp.Diagnostics.Append(data.SetModelData(&diff)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
