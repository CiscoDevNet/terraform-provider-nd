package resource_fabric_aci

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &fabricAciResource{}
	_ resource.ResourceWithConfigure   = &fabricAciResource{}
	_ resource.ResourceWithImportState = &fabricAciResource{}
)

// NewFabricAciResource is a helper function to simplify the provider implementation.
func NewFabricAciResource() resource.Resource {
	return &fabricAciResource{}
}

// fabricAciResource is the resource implementation.
type fabricAciResource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the resource type name.
func (r *fabricAciResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_aci"
}

// Schema defines the schema for the resource.
func (r *fabricAciResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricAciResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *fabricAciResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	manageModule := client.GetModule(ModuleKey)
	if manageModule == nil {
		resp.Diagnostics.AddError(
			"Manage Module Not Found",
			"The manage module was not registered with the provider.",
		)
		return
	}

	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule),
		)
		return
	}

	r.manageClient = manageClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *fabricAciResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_fabric_aci")

	var in FabricAciModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating Fabric ACI: fabric_name=%s", in.FabricName.ValueString())

	r.rscCreateFabricAci(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_fabric_aci with fabric_name '%s'", in.FabricName.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *fabricAciResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_fabric_aci")

	var state FabricAciModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading Fabric ACI: fabric_name=%s", state.FabricName.ValueString())

	notFound := r.rscGetFabricAci(ctx, &resp.Diagnostics, &state)
	if notFound {
		resp.State.RemoveResource(ctx)
		log.Printf("[DEBUG] Fabric ACI not found, removing from state: fabric_name=%s", state.FabricName.ValueString())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_fabric_aci with fabric_name '%s'", state.FabricName.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *fabricAciResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_fabric_aci")

	var plan FabricAciModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Updating Fabric ACI: fabric_name=%s", plan.FabricName.ValueString())

	r.rscUpdateFabricAci(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] End update of resource nd_fabric_aci with fabric_name '%s'", plan.FabricName.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *fabricAciResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_fabric_aci")

	var state FabricAciModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.rscDeleteFabricAci(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End delete of resource nd_fabric_aci with fabric_name '%s'", state.FabricName.ValueString())
}

// ImportState imports a fabric ACI resource by fabric name.
func (r *fabricAciResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_fabric_aci")

	var state FabricAciModel
	state.FabricName = types.StringValue(req.ID)

	notFound := r.rscGetFabricAci(ctx, &resp.Diagnostics, &state)
	if notFound {
		resp.Diagnostics.AddError(
			"Error Importing Fabric ACI",
			fmt.Sprintf("Could not find fabric ACI with fabric_name %q.", req.ID),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: The values for `username`, `password`, and `login_domain` attributes
	// will not be imported when the resource imports an already registered APIC cluster.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_fabric_aci with fabric_name '%s'", state.FabricName.ValueString())
}
