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
	_ resource.ResourceWithModifyPlan     = &remoteStorageLocationResource{}
	_ resource.ResourceWithValidateConfig = &remoteStorageLocationResource{}
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

// ValidateConfig enforces conditional validation that generated schema
// validators cannot express.
func (r *remoteStorageLocationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config RemoteStorageLocationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nfsKnown := !config.Nfs.IsNull() && !config.Nfs.IsUnknown()
	scpSftpKnown := !config.ScpSftp.IsNull() && !config.ScpSftp.IsUnknown()

	if !config.Nfs.IsUnknown() && !config.ScpSftp.IsUnknown() && nfsKnown == scpSftpKnown {
		resp.Diagnostics.AddAttributeError(
			path.Root("nfs"),
			"Invalid remote storage configuration",
			"Configure exactly one of `nfs` or `scp_sftp`.",
		)
	}

	if !scpSftpKnown {
		return
	}

	passwordConfigured := !config.ScpSftp.Password.IsNull() &&
		!config.ScpSftp.Password.IsUnknown() &&
		config.ScpSftp.Password.ValueString() != ""
	sshKeyConfigured := !config.ScpSftp.SshKey.IsNull() &&
		!config.ScpSftp.SshKey.IsUnknown() &&
		config.ScpSftp.SshKey.ValueString() != ""

	if !config.ScpSftp.Password.IsUnknown() && !config.ScpSftp.SshKey.IsUnknown() &&
		!passwordConfigured && !sshKeyConfigured {
		resp.Diagnostics.AddAttributeError(
			path.Root("scp_sftp").AtName("password"),
			"Missing SCP/SFTP authentication",
			"Configure exactly one of `scp_sftp.password` or `scp_sftp.ssh_key`.",
		)
	}

	if !config.ScpSftp.Passphrase.IsNull() && !config.ScpSftp.Passphrase.IsUnknown() &&
		config.ScpSftp.Passphrase.ValueString() != "" &&
		!config.ScpSftp.SshKey.IsUnknown() && !sshKeyConfigured {
		resp.Diagnostics.AddAttributeError(
			path.Root("scp_sftp").AtName("passphrase"),
			"Invalid SCP/SFTP passphrase configuration",
			"Attribute `scp_sftp.passphrase` requires `scp_sftp.ssh_key`.",
		)
	}

	if !config.ScpSftp.AcceptHostKey.IsNull() && !config.ScpSftp.AcceptHostKey.IsUnknown() &&
		config.ScpSftp.AcceptHostKey.ValueBool() &&
		!config.ScpSftp.IgnoreHostKeyValidation.IsNull() && !config.ScpSftp.IgnoreHostKeyValidation.IsUnknown() &&
		config.ScpSftp.IgnoreHostKeyValidation.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("scp_sftp").AtName("accept_host_key"),
			"Invalid host-key configuration",
			"Attributes `scp_sftp.accept_host_key` and `scp_sftp.ignore_host_key_validation` cannot both be true.",
		)
	}
}

// ModifyPlan rejects changes to the immutable path of an existing NFS location
// and requires replacement when the selected remote storage protocol family
// changes between NFS and SCP/SFTP.
func (r *remoteStorageLocationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var state RemoteStorageLocationModel
	var plan RemoteStorageLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	stateNFSKnown := !state.Nfs.IsNull() && !state.Nfs.IsUnknown()
	planNFSKnown := !plan.Nfs.IsNull() && !plan.Nfs.IsUnknown()
	if stateNFSKnown && planNFSKnown &&
		!state.Path.IsNull() && !state.Path.IsUnknown() &&
		!plan.Path.IsNull() && !plan.Path.IsUnknown() &&
		!state.Path.Equal(plan.Path) {
		resp.Diagnostics.AddAttributeError(
			path.Root("path"),
			"NFS path cannot be changed",
			"The path of an existing NFS remote storage location cannot be updated. Delete the resource and create it again with the new path.",
		)
		return
	}

	stateBranch, stateBranchKnown := remoteStorageLocationBranch(state)
	planBranch, planBranchKnown := remoteStorageLocationBranch(plan)
	if !stateBranchKnown || !planBranchKnown {
		return
	}

	if stateBranch != planBranch {
		switch planBranch {
		case "nfs":
			resp.RequiresReplace = append(resp.RequiresReplace, path.Root("nfs"))
		case "scp_sftp":
			resp.RequiresReplace = append(resp.RequiresReplace, path.Root("scp_sftp"))
		}
		return
	}
}

func remoteStorageLocationBranch(model RemoteStorageLocationModel) (string, bool) {
	if model.Nfs.IsUnknown() || model.ScpSftp.IsUnknown() {
		return "", false
	}

	switch {
	case !model.Nfs.IsNull() && model.ScpSftp.IsNull():
		return "nfs", true
	case model.Nfs.IsNull() && !model.ScpSftp.IsNull():
		return "scp_sftp", true
	default:
		return "", false
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

	in.Id = in.Name
	log.Printf("[DEBUG] Creating ND Remote Storage Location: id=%s", in.Id.ValueString())

	r.rscCreateRemoteStorageLocation(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_remote_storage_location with id=%s", in.Id.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *remoteStorageLocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.Id.IsNull() || state.Id.IsUnknown() || state.Id.ValueString() == "" {
		state.Id = state.Name
	}
	log.Printf("[DEBUG] Reading ND Remote Storage Location: id=%s", state.Id.ValueString())

	if r.rscGetRemoteStorageLocation(ctx, &resp.Diagnostics, &state) {
		resp.State.RemoveResource(ctx)
		log.Printf("[DEBUG] Removed nd_remote_storage_location id=%s from state because it was not found", state.Id.ValueString())
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_remote_storage_location with id=%s", state.Id.ValueString())
}

// Update updates the resource and saves the latest Terraform state.
func (r *remoteStorageLocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_remote_storage_location")

	var in RemoteStorageLocationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	in.Id = in.Name
	log.Printf("[DEBUG] Updating ND Remote Storage Location: id=%s", in.Id.ValueString())

	r.rscUpdateRemoteStorageLocation(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End update of resource nd_remote_storage_location with id=%s", in.Id.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *remoteStorageLocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.Id.IsNull() || state.Id.IsUnknown() || state.Id.ValueString() == "" {
		state.Id = state.Name
	}
	r.rscDeleteRemoteStorageLocation(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("[DEBUG] End delete of resource nd_remote_storage_location with id=%s", state.Id.ValueString())
}

// ImportState imports an nd_remote_storage_location resource by id.
func (r *remoteStorageLocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_remote_storage_location")

	var state RemoteStorageLocationModel
	state.Id = types.StringValue(req.ID)
	state.Name = state.Id

	if r.rscGetRemoteStorageLocation(ctx, &resp.Diagnostics, &state) {
		resp.Diagnostics.AddError(
			"Error Importing ND Remote Storage Location",
			fmt.Sprintf("Could not import nd_remote_storage_location with id %q: resource not found", state.Id.ValueString()),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_remote_storage_location with id=%s", state.Id.ValueString())
}
