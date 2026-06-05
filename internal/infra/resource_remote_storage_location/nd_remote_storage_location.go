package resource_remote_storage_location

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/common/utils"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func acceptHostKey(value types.Bool) bool {
	return !value.IsNull() && !value.IsUnknown() && value.ValueBool()
}

// rscCreateRemoteStorageLocation creates an nd_remote_storage_location resource.
func (r *remoteStorageLocationResource) rscCreateRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, input *RemoteStorageLocationModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}

	log.Printf("[INFO] Create nd_remote_storage_location name=%s", input.Name.ValueString())

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient)
	remoteStorageAPI.AcceptHostKey = acceptHostKey(input.AcceptHostKey)
	payload, err := json.Marshal(input.GetModelData())
	if err != nil {
		dg.AddError(
			"Error Creating ND Remote Storage Location",
			fmt.Sprintf("Could not create nd_remote_storage_location, data marshal error: %v", err),
		)
		return
	}

	res, err := remoteStorageAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Creating ND Remote Storage Location",
			fmt.Sprintf("Could not create nd_remote_storage_location, unexpected error: %v %v", err, res),
		)
		return
	}

	if found := r.rscGetRemoteStorageLocation(ctx, dg, input); !found && !dg.HasError() {
		dg.AddError(
			"Error Creating ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location %q after create", input.Name.ValueString()),
		)
	}
}

// rscGetRemoteStorageLocation retrieves nd_remote_storage_location information by name.
func (r *remoteStorageLocationResource) rscGetRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, in *RemoteStorageLocationModel) bool {
	if in == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return false
	}
	log.Printf("[INFO] Read nd_remote_storage_location name=%s", in.Name.ValueString())

	preservedPassword := in.Password
	preservedSshKey := in.SshKey
	preservedPassphrase := in.Passphrase
	preservedIgnoreHostKeyValidation := in.IgnoreHostKeyValidation
	preservedAcceptHostKey := in.AcceptHostKey

	if preservedPassword.IsUnknown() {
		preservedPassword = types.StringNull()
	}
	if preservedSshKey.IsUnknown() {
		preservedSshKey = types.StringNull()
	}
	if preservedPassphrase.IsUnknown() {
		preservedPassphrase = types.StringNull()
	}
	if preservedIgnoreHostKeyValidation.IsUnknown() {
		preservedIgnoreHostKeyValidation = types.BoolNull()
	}
	if preservedAcceptHostKey.IsUnknown() {
		preservedAcceptHostKey = types.BoolNull()
	}

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient)
	remoteStorageAPI.Name = in.Name.ValueString()

	respData, err := remoteStorageAPI.Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			log.Printf("[WARN] nd_remote_storage_location name=%s not found", in.Name.ValueString())
			return false
		}
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location, unexpected error: %v %s", err, string(respData)),
		)
		return false
	}
	if respData == nil {
		log.Printf("[WARN] nd_remote_storage_location name=%s not found", in.Name.ValueString())
		return false
	}

	var remoteStorageResp NDFCRemoteStorageLocationModel
	if err := json.Unmarshal(respData, &remoteStorageResp); err != nil {
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not unmarshal nd_remote_storage_location response, unexpected error: %v", err),
		)
		return false
	}

	in.SetModelData(&remoteStorageResp)

	if remoteStorageResp.Authentication.Password == "" {
		in.Password = preservedPassword
	}
	if remoteStorageResp.Authentication.SshKey == "" {
		in.SshKey = preservedSshKey
	}
	if remoteStorageResp.Authentication.Passphrase == "" {
		in.Passphrase = preservedPassphrase
	}
	if remoteStorageResp.Authentication.IgnoreHostKeyValidation == nil {
		in.IgnoreHostKeyValidation = preservedIgnoreHostKeyValidation
	}
	in.AcceptHostKey = preservedAcceptHostKey
	return true
}

// rscUpdateRemoteStorageLocation updates an nd_remote_storage_location resource.
func (r *remoteStorageLocationResource) rscUpdateRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, input *RemoteStorageLocationModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}

	log.Printf("[INFO] Update nd_remote_storage_location name=%s", input.Name.ValueString())

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient)
	remoteStorageAPI.Name = input.Name.ValueString()
	remoteStorageAPI.AcceptHostKey = acceptHostKey(input.AcceptHostKey)

	payload, err := json.Marshal(input.GetModelData())
	if err != nil {
		dg.AddError(
			"Error Updating ND Remote Storage Location",
			fmt.Sprintf("Could not update nd_remote_storage_location, data marshal error: %v", err),
		)
		return
	}

	res, err := remoteStorageAPI.Put(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Updating ND Remote Storage Location",
			fmt.Sprintf("Could not update nd_remote_storage_location, unexpected error: %v %v", err, res),
		)
		return
	}

	if found := r.rscGetRemoteStorageLocation(ctx, dg, input); !found && !dg.HasError() {
		dg.AddError(
			"Error Updating ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location %q after update", input.Name.ValueString()),
		)
	}
}

// rscDeleteRemoteStorageLocation deletes an nd_remote_storage_location resource by name.
func (r *remoteStorageLocationResource) rscDeleteRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, state *RemoteStorageLocationModel) {
	log.Printf("[INFO] Delete nd_remote_storage_location name=%s", state.Name.ValueString())

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient)
	remoteStorageAPI.Name = state.Name.ValueString()

	res, err := remoteStorageAPI.Delete(nil)
	if err != nil {
		dg.AddError(
			"Error Deleting ND Remote Storage Location",
			fmt.Sprintf("Could not delete nd_remote_storage_location, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting ND Remote Storage Location: error=%s", err.Error())
		return
	}
}
