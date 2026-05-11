package resource_local_user

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource                = &localUserNdResource{}
	_ resource.ResourceWithConfigure   = &localUserNdResource{}
	_ resource.ResourceWithImportState = &localUserNdResource{}
)

// NewLocalUserResource is a helper function to simplify the provider implementation.
func NewLocalUserResource() resource.Resource {
	return &localUserNdResource{}
}

// localUserNdResource is the resource implementation.
type localUserNdResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *localUserNdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_local_user"
}

// Schema defines the schema for the resource.
func (r *localUserNdResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = LocalUserResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *localUserNdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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

	infraModule := client.GetModule(ModuleKey)
	if infraModule == nil {
		resp.Diagnostics.AddError(
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	r.infraClient = infraModule.(*infra.NexusDashboardInfra)
}

// Create creates the resource and sets the initial Terraform state.
func (r *localUserNdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_local_user")

	var in LocalUserModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating ND Local User: login_id=%s", in.LoginId.ValueString())

	r.rscCreateLocalUser(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_local_user with id '%s'", in.LoginId.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *localUserNdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_local_user")

	var state LocalUserModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading ND Local User: login_id=%s", state.LoginId.ValueString())

	r.rscGetLocalUser(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_local_user with id '%s'", state.LoginId.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *localUserNdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_local_user")
	var plan, state LocalUserModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Updating ND Local User: login_id=%s", plan.LoginId.ValueString())

	r.rscUpdateLocalUser(ctx, &resp.Diagnostics, &plan, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] End update of resource nd_local_user with id '%s'", plan.LoginId.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *localUserNdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_local_user")
	var state LocalUserModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.rscDeleteLocalUser(ctx, &resp.Diagnostics, &state)
	log.Printf("[DEBUG] End delete of resource nd_local_user with id '%s'", state.LoginId.ValueString())
}

// ImportState imports a nd local user resource by login_id.
func (r *localUserNdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_local_user")

	var state LocalUserModel
	state.LoginId = types.StringValue(req.ID)

	r.rscGetLocalUser(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	// TODO: The value for the `user_password` attribute will not be imported when the nd_local_user resource imports an already created local user from Nexus Dashboard.
	// Need to use Environment Variables or a credentials file to supply the password during the import of the resource.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_local_user")
}
