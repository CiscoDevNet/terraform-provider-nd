package resource_remote_storage_location

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ resource.Resource                   = &remoteStorageLocationResource{}
	_ resource.ResourceWithConfigure      = &remoteStorageLocationResource{}
	_ resource.ResourceWithImportState    = &remoteStorageLocationResource{}
	_ resource.ResourceWithValidateConfig = &remoteStorageLocationResource{}
	_ resource.ResourceWithModifyPlan     = &remoteStorageLocationResource{}
)

// NewRemoteStorageLocationResource is a helper function to simplify the provider implementation.
func NewRemoteStorageLocationResource() resource.Resource {
	return &remoteStorageLocationResource{}
}

// remoteStorageLocationResource is the resource implementation.
type remoteStorageLocationResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *remoteStorageLocationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_remote_storage_location"
}

// Schema defines the schema for the resource.
func (r *remoteStorageLocationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = RemoteStorageLocationResourceSchema(ctx)
}

// ValidateConfig enforces authentication field conflicts.
func (r *remoteStorageLocationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var storageType types.String
	var readWrite types.Bool
	var password types.String
	var sshKey types.String
	var passphrase types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("storage_location_type"), &storageType)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("read_write"), &readWrite)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &password)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ssh_key"), &sshKey)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("passphrase"), &passphrase)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !storageType.IsNull() && !storageType.IsUnknown() &&
		storageType.ValueString() == "nfs" && readWrite.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("read_write"),
			"Missing required attribute",
			"Attribute `read_write` must be set when `storage_location_type` is `nfs`.",
		)
	}

	isConfigured := func(v types.String) bool {
		return !v.IsNull() && !v.IsUnknown()
	}

	if isConfigured(password) && isConfigured(sshKey) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Invalid authentication configuration",
			"Attributes `password` and `ssh_key` cannot both be set. Configure only one authentication method.",
		)

		resp.Diagnostics.AddAttributeError(
			path.Root("ssh_key"),
			"Invalid authentication configuration",
			"Attributes `ssh_key` and `password` cannot both be set. Configure only one authentication method.",
		)
	}

	if isConfigured(password) && isConfigured(passphrase) {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Invalid authentication configuration",
			"Attributes `password` and `passphrase` cannot both be set.",
		)

		resp.Diagnostics.AddAttributeError(
			path.Root("passphrase"),
			"Invalid authentication configuration",
			"Attributes `passphrase` and `password` cannot both be set.",
		)
	}
}

// ModifyPlan nulls out attributes that are only meaningful for certain
// storage_location_type values, so plan and post-apply state agree.
//
// `read_write` is only honored by the NFS backend - the API does not echo
// it back for SCP/SFTP, so we force it to null in the plan to match what
// the resource will store after a successful apply.
func (r *remoteStorageLocationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan RemoteStorageLocationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.StorageLocationType.IsNull() && !plan.StorageLocationType.IsUnknown() &&
		plan.StorageLocationType.ValueString() != "nfs" && !plan.ReadWrite.IsNull() {
		plan.ReadWrite = types.BoolNull()
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

// Configure adds the provider configured client to the resource.
func (r *remoteStorageLocationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *remoteStorageLocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_remote_storage_location")

	var in RemoteStorageLocationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating ND Remote Storage Location: name=%s", in.Name.ValueString())

	r.rscCreateRemoteStorageLocation(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_remote_storage_location with name=%s", in.Name.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *remoteStorageLocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading ND Remote Storage Location: name=%s", state.Name.ValueString())

	found := r.rscGetRemoteStorageLocation(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		log.Printf("[DEBUG] Removed nd_remote_storage_location name=%s from state because it was not found", state.Name.ValueString())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_remote_storage_location with name=%s", state.Name.ValueString())
}

// Update updates the resource and saves the latest Terraform state.
func (r *remoteStorageLocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_remote_storage_location")

	var in RemoteStorageLocationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Updating ND Remote Storage Location: name=%s", in.Name.ValueString())

	r.rscUpdateRemoteStorageLocation(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End update of resource nd_remote_storage_location with name=%s", in.Name.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *remoteStorageLocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.rscDeleteRemoteStorageLocation(ctx, &resp.Diagnostics, &state)
	log.Printf("[DEBUG] End delete of resource nd_remote_storage_location with name=%s", state.Name.ValueString())
}

// ImportState imports an nd_remote_storage_location resource by name.
func (r *remoteStorageLocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	state.Name = types.StringValue(req.ID)

	found := r.rscGetRemoteStorageLocation(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError(
			"Error Importing ND Remote Storage Location",
			fmt.Sprintf("Could not find nd_remote_storage_location %q", state.Name.ValueString()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_remote_storage_location with name=%s", state.Name.ValueString())
}
