package resource_local_user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// rscCreateLocalUser creates a nd local user resource
func (r *localUserNdResource) rscCreateLocalUser(ctx context.Context, dg *diag.Diagnostics, input *LocalUserModel) {
	id := input.LoginId.ValueString()
	input.Id = input.LoginId
	inData := input.GetModelData()

	// Create nd local user API client
	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient)
	localUserAPI.LoginId = id

	// Convert model data to JSON
	localUserPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating ND Local User",
			fmt.Sprintf("Could not marshal nd_local_user create payload: %s", err.Error()),
		)
		return
	}

	// Call the API to create the nd local user
	_, err = localUserAPI.Post(localUserPayload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		tflog.Error(ctx, "ND local user create request failed", map[string]interface{}{
			"login_id": id,
			"error":    err.Error(),
		})
		dg.AddError(
			"Error Creating ND Local User",
			fmt.Sprintf("Could not create nd_local_user with login_id %q: %s", id, err.Error()),
		)
		return
	}

	notFound := r.rscGetLocalUser(ctx, dg, input)
	if dg.HasError() {
		return
	}
	if notFound {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not read nd_local_user with login_id %q after create: resource not found", id),
		)
	}
}

// rscGetLocalUser retrieves nd local user information by login_id
func (r *localUserNdResource) rscGetLocalUser(ctx context.Context, dg *diag.Diagnostics, in *LocalUserModel) bool {
	id := in.LoginId.ValueString()
	// Nexus Dashboard deliberately omits the password from GET responses.
	preservedUserPassword := in.UserPassword

	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient)
	localUserAPI.LoginId = in.LoginId.ValueString()
	respData, err := localUserAPI.Get()

	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		tflog.Error(ctx, "ND local user read request failed", map[string]interface{}{
			"login_id": id,
			"error":    err.Error(),
		})
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not read nd_local_user with login_id %q: %s. Response: %s", id, err.Error(), string(respData)),
		)
		return false
	}

	var localUserResp NDFCLocalUserModel
	err = json.Unmarshal(respData, &localUserResp)
	if err != nil {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not unmarshal nd_local_user response for login_id %q: %s", id, err.Error()),
		)
		return false
	}
	if localUserResp.LoginId == "" {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("The nd_local_user response for login_id %q did not contain loginID.", id),
		)
		return false
	}

	dg.Append(in.SetModelData(&localUserResp)...)
	if dg.HasError() {
		return false
	}
	// The resource ID is Terraform-only and is derived from the API login ID.
	in.Id = in.LoginId

	// Restore the configured value after generated hydration clears the
	// API-omitted password field.
	in.UserPassword = preservedUserPassword
	return false
}

// rscUpdateLocalUser updates a nd local user with the provided payload.
// If the plan's user_password (sourced from the configuration file) matches
// the prior state's user_password, the password field is omitted from the
// payload to avoid resending an unchanged password to the API.
func (r *localUserNdResource) rscUpdateLocalUser(ctx context.Context, dg *diag.Diagnostics, plan *LocalUserModel, state *LocalUserModel) {
	id := plan.LoginId.ValueString()
	plan.Id = plan.LoginId
	inData := plan.GetModelData()

	// The plan's user_password reflects the value from the configuration file.
	// If it matches the prior state's user_password (i.e. the user did not change
	// the password in the config), omit the password from the update payload to
	// avoid resending an unchanged password to the API.
	planPasswordSet := !plan.UserPassword.IsNull() && !plan.UserPassword.IsUnknown()
	statePasswordSet := !state.UserPassword.IsNull() && !state.UserPassword.IsUnknown()
	passwordUnchanged := statePasswordSet && plan.UserPassword.ValueString() == state.UserPassword.ValueString()
	if planPasswordSet && passwordUnchanged {
		tflog.Debug(ctx, "Omitting unchanged password from ND local user update payload", map[string]interface{}{
			"login_id": id,
		})
		inData.UserPassword = ""
	}

	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient)
	localUserAPI.LoginId = plan.LoginId.ValueString()

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating ND Local User",
			fmt.Sprintf("Could not marshal nd_local_user update payload: %s", err.Error()),
		)
		return
	}
	_, err = localUserAPI.Put(inDataBytes, &ndapi.APIOptions{DisablePayloadLog: true})

	if err != nil {
		tflog.Error(ctx, "ND local user update request failed", map[string]interface{}{
			"login_id": id,
			"error":    err.Error(),
		})
		dg.AddError(
			"Error Updating ND Local User",
			fmt.Sprintf("Could not update nd_local_user with login_id %q: %s", id, err.Error()),
		)
		return
	}
	// Read the updated nd local user
	notFound := r.rscGetLocalUser(ctx, dg, plan)
	if dg.HasError() {
		return
	}
	if notFound {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not read nd_local_user with login_id %q after update: resource not found", id),
		)
	}
}

// rscDeleteLocalUser deletes a nd local user by login_id
func (r *localUserNdResource) rscDeleteLocalUser(ctx context.Context, dg *diag.Diagnostics, state *LocalUserModel) {
	id := state.LoginId.ValueString()
	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient)
	localUserAPI.LoginId = id

	res, err := localUserAPI.Delete(nil)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			tflog.Debug(ctx, "ND local user is already absent during delete", map[string]interface{}{
				"login_id": id,
			})
			return
		}
		tflog.Error(ctx, "ND local user delete request failed", map[string]interface{}{
			"login_id": id,
			"error":    err.Error(),
		})
		dg.AddError(
			"Error Deleting ND Local User",
			fmt.Sprintf("Could not delete nd_local_user with login_id %q: %s. Response: %s", id, err.Error(), res.String()),
		)
		return
	}
}
