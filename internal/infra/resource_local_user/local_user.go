package resource_local_user

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	var in LocalUserModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Nexus Dashboard identifies a local user by loginID; Terraform uses the
	// same value as the resource ID.
	in.Id = in.LoginId
	tflog.Debug(ctx, "Creating ND local user", map[string]interface{}{
		"login_id": in.LoginId.ValueString(),
	})

	r.rscCreateLocalUser(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created ND local user", map[string]interface{}{
		"login_id": in.LoginId.ValueString(),
	})
}

// Read refreshes the Terraform state with the latest data.
func (r *localUserNdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LocalUserModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Id = state.LoginId

	id := state.LoginId.ValueString()
	tflog.Debug(ctx, "Reading ND local user", map[string]interface{}{
		"login_id": id,
	})

	if r.rscGetLocalUser(ctx, &resp.Diagnostics, &state) {
		tflog.Warn(ctx, "ND local user no longer exists; removing it from Terraform state", map[string]interface{}{
			"login_id": id,
		})
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read ND local user", map[string]interface{}{
		"login_id": state.LoginId.ValueString(),
	})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *localUserNdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

	plan.Id = plan.LoginId
	tflog.Debug(ctx, "Updating ND local user", map[string]interface{}{
		"login_id": plan.LoginId.ValueString(),
	})

	r.rscUpdateLocalUser(ctx, &resp.Diagnostics, &plan, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated ND local user", map[string]interface{}{
		"login_id": plan.LoginId.ValueString(),
	})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *localUserNdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LocalUserModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.Id = state.LoginId

	tflog.Debug(ctx, "Deleting ND local user", map[string]interface{}{
		"login_id": state.LoginId.ValueString(),
	})

	r.rscDeleteLocalUser(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Deleted ND local user", map[string]interface{}{
		"login_id": state.LoginId.ValueString(),
	})
}

// ImportState imports a nd local user resource by login_id.
func (r *localUserNdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if strings.TrimSpace(req.ID) == "" {
		resp.Diagnostics.AddError(
			"Invalid ND Local User Import ID",
			"The import ID must be a non-empty local-user login ID.",
		)
		return
	}

	var state LocalUserModel
	state.LoginId = types.StringValue(req.ID)
	state.Id = state.LoginId
	state.UserPassword = types.StringNull()
	tflog.Debug(ctx, "Importing ND local user", map[string]interface{}{
		"login_id": req.ID,
	})

	if r.rscGetLocalUser(ctx, &resp.Diagnostics, &state) {
		resp.Diagnostics.AddError(
			"Error Importing ND Local User",
			fmt.Sprintf("Could not import nd local user with id %q: resource not found", req.ID),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	// Nexus Dashboard does not return password values. Imported state therefore
	// cannot restore user_password; configuration must supply it for later plans.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Imported ND local user", map[string]interface{}{
		"login_id": req.ID,
	})
}
