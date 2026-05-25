package resource_backup

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
	_ resource.Resource                = &backupNdResource{}
	_ resource.ResourceWithConfigure   = &backupNdResource{}
	_ resource.ResourceWithImportState = &backupNdResource{}
)

// NewBackupResource is a helper function to simplify the provider implementation.
func NewBackupResource() resource.Resource {
	return &backupNdResource{}
}

// backupNdResource is the resource implementation.
type backupNdResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *backupNdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_backup"
}

// Schema defines the schema for the resource.
func (r *backupNdResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = BackupResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *backupNdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *backupNdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_backup")

	var in BackupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating ND Backup: name=%s", in.Name.ValueString())

	r.rscCreateBackup(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_backup with name=%s", in.Name.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *backupNdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_backup")

	var state BackupModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading ND Backup: name=%s", state.Name.ValueString())

	r.rscGetBackup(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_backup with name=%s", state.Name.ValueString())
}

// Update is not supported for nd_backup. All updatable attributes use
// RequiresReplace plan modifiers, so Terraform will never invoke Update.
// This method exists only to satisfy the resource.Resource interface.
func (r *backupNdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"The nd_backup resource does not support in-place updates. Changes require resource replacement.",
	)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *backupNdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_backup")
	var state BackupModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.rscDeleteBackup(ctx, &resp.Diagnostics, &state)
	log.Printf("[DEBUG] End delete of resource nd_backup with name=%s", state.Name.ValueString())
}

// ImportState imports a nd backup resource by name.
func (r *backupNdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_backup")

	var state BackupModel
	state.Name = types.StringValue(req.ID)

	r.rscGetBackup(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	// TODO: The value for the `encryption_key` attribute will not be imported when the nd_backup resource imports an already created backup from Nexus Dashboard.
	// Need to use Environment Variables or a credentials file to supply the encryption key during the import of the resource.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_backup with name=%s", state.Name.ValueString())
}
