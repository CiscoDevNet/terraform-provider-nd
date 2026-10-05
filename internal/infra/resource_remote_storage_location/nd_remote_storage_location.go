package resource_remote_storage_location

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/common/utils"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	remoteStorageLocationDeletePollInterval = 10 * time.Second
	remoteStorageLocationDeletePollTimeout  = 5 * time.Minute
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

	id := input.Id.ValueString()
	log.Printf("[INFO] Create nd_remote_storage_location id=%s", id)

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	remoteStorageAPI.AcceptHostKey = acceptHostKey(input.ScpSftp.AcceptHostKey)
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

	if r.rscGetRemoteStorageLocation(ctx, dg, input) && !dg.HasError() {
		dg.AddError(
			"Error Creating ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location %q after create: resource not found", id),
		)
	}
}

// rscGetRemoteStorageLocation retrieves nd_remote_storage_location information by id.
// It returns true when the remote object was not found.
func (r *remoteStorageLocationResource) rscGetRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, in *RemoteStorageLocationModel) bool {
	if in == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return false
	}
	id := in.Id.ValueString()
	log.Printf("[INFO] Read nd_remote_storage_location id=%s", id)

	preserveScpSftpState := !in.ScpSftp.IsNull() && !in.ScpSftp.IsUnknown()
	preservedPassword := in.ScpSftp.Password
	preservedSshKey := in.ScpSftp.SshKey
	preservedPassphrase := in.ScpSftp.Passphrase
	preservedIgnoreHostKeyValidation := in.ScpSftp.IgnoreHostKeyValidation
	preservedAcceptHostKey := in.ScpSftp.AcceptHostKey

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

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	remoteStorageAPI.Name = id

	respData, err := remoteStorageAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location, unexpected error: %v %s", err, string(respData)),
		)
		return false
	}
	if respData == nil {
		log.Printf("[WARN] nd_remote_storage_location id=%s not found: empty response", id)
		return true
	}

	var remoteStorageResp NDFCRemoteStorageLocationModel
	if err := json.Unmarshal(respData, &remoteStorageResp); err != nil {
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not unmarshal nd_remote_storage_location response, unexpected error: %v", err),
		)
		return false
	}

	nfsSelected := hasNfsConfiguration(remoteStorageResp.Nfs)
	scpSftpSelected := hasScpSftpConfiguration(remoteStorageResp.ScpSftp)

	switch {
	case nfsSelected && scpSftpSelected:
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("API response for nd_remote_storage_location %q contains both NFS and SCP/SFTP configurations", id),
		)
		return false
	case nfsSelected, scpSftpSelected:
	default:
		dg.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not determine the storage type for nd_remote_storage_location %q from the API response", id),
		)
		return false
	}

	dg.Append(in.SetModelData(&remoteStorageResp)...)
	if dg.HasError() {
		return false
	}

	if nfsSelected {
		dg.Append(in.Nfs.SetValue(&remoteStorageResp.Nfs)...)
		if dg.HasError() {
			return false
		}
		in.Nfs.state = attr.ValueStateKnown
		in.ScpSftp = NewScpSftpValueNull()
	}

	if scpSftpSelected {
		dg.Append(in.ScpSftp.SetValue(&remoteStorageResp.ScpSftp)...)
		if dg.HasError() {
			return false
		}
		in.Nfs = NewNfsValueNull()
		in.ScpSftp.state = attr.ValueStateKnown

		if preserveScpSftpState && remoteStorageResp.ScpSftp.Authentication.Password == "" {
			in.ScpSftp.Password = preservedPassword
		}
		if preserveScpSftpState && remoteStorageResp.ScpSftp.Authentication.SshKey == "" {
			in.ScpSftp.SshKey = preservedSshKey
		}
		if preserveScpSftpState && remoteStorageResp.ScpSftp.Authentication.Passphrase == "" {
			in.ScpSftp.Passphrase = preservedPassphrase
		}
		if remoteStorageResp.ScpSftp.Authentication.IgnoreHostKeyValidation == nil {
			if preserveScpSftpState {
				in.ScpSftp.IgnoreHostKeyValidation = preservedIgnoreHostKeyValidation
			} else {
				in.ScpSftp.IgnoreHostKeyValidation = types.BoolValue(false)
			}
		}
		if preserveScpSftpState {
			in.ScpSftp.AcceptHostKey = preservedAcceptHostKey
		}
	}
	in.Id = in.Name
	return false
}

// rscUpdateRemoteStorageLocation updates an nd_remote_storage_location resource.
func (r *remoteStorageLocationResource) rscUpdateRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, input *RemoteStorageLocationModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}

	id := input.Id.ValueString()
	log.Printf("[INFO] Update nd_remote_storage_location id=%s", id)

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	remoteStorageAPI.Name = id
	remoteStorageAPI.AcceptHostKey = acceptHostKey(input.ScpSftp.AcceptHostKey)

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

	if r.rscGetRemoteStorageLocation(ctx, dg, input) && !dg.HasError() {
		dg.AddError(
			"Error Updating ND Remote Storage Location",
			fmt.Sprintf("Could not read nd_remote_storage_location %q after update: resource not found", id),
		)
	}
}

// rscDeleteRemoteStorageLocation deletes an nd_remote_storage_location resource by id.
func (r *remoteStorageLocationResource) rscDeleteRemoteStorageLocation(ctx context.Context, dg *diag.Diagnostics, state *RemoteStorageLocationModel) {
	id := state.Id.ValueString()
	log.Printf("[INFO] Delete nd_remote_storage_location id=%s", id)

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	remoteStorageAPI.Name = id

	res, err := remoteStorageAPI.Delete(nil)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting ND Remote Storage Location",
			fmt.Sprintf("Could not delete nd_remote_storage_location, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting ND Remote Storage Location: error=%s", err.Error())
		return
	}

	var lastResponse []byte
	pollErr := utils.PollUntil(ctx, remoteStorageLocationDeletePollInterval, remoteStorageLocationDeletePollTimeout, func(context.Context) (bool, error) {
		respData, getErr := remoteStorageAPI.Get()
		lastResponse = respData
		if getErr == nil {
			if respData == nil {
				return true, nil
			}
			log.Printf("[INFO] Waiting for nd_remote_storage_location id=%s deletion to complete", id)
			return false, nil
		}
		if strings.Contains(getErr.Error(), "StatusCode 404") {
			return true, nil
		}
		return false, fmt.Errorf("failed to verify deletion of nd_remote_storage_location %q: %w %s", id, getErr, string(respData))
	})
	if pollErr == nil {
		return
	}

	switch {
	case errors.Is(pollErr, utils.ErrPollTimeout):
		dg.AddError(
			"Error Deleting ND Remote Storage Location",
			fmt.Sprintf("Timed out after %s waiting for nd_remote_storage_location %q deletion to complete. Last response: %s", remoteStorageLocationDeletePollTimeout, id, string(lastResponse)),
		)
	case errors.Is(pollErr, context.Canceled), errors.Is(pollErr, context.DeadlineExceeded):
		dg.AddError(
			"Error Deleting ND Remote Storage Location",
			fmt.Sprintf("Context ended while waiting for nd_remote_storage_location %q deletion to complete: %v", id, pollErr),
		)
	default:
		dg.AddError("Error Deleting ND Remote Storage Location", pollErr.Error())
	}
}
