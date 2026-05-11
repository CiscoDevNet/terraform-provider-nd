package resource_local_user

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rscCreateLocalUser creates a nd local user resource
func (r *localUserNdResource) rscCreateLocalUser(ctx context.Context, dg *diag.Diagnostics, input *LocalUserModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	log.Printf("[INFO] Create nd_local_user login_id=%s", input.LoginId.ValueString())
	inData := input.GetModelData()

	// Create nd local user API client
	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	// Convert model data to JSON
	localUserPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating ND Local User",
			fmt.Sprintf("Could not create nd local user, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the nd local user
	res, err := localUserAPI.Post(localUserPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating ND Local User",
			fmt.Sprintf("Could not create nd local user, unexpected error: %v %v", err, res),
		)
		return
	}
	r.rscGetLocalUser(ctx, dg, input)
}

// rscGetLocalUser retrieves nd local user information by login_id
func (r *localUserNdResource) rscGetLocalUser(ctx context.Context, dg *diag.Diagnostics, in *LocalUserModel) {
	log.Printf("[INFO] Read nd_local_user login_id=%s", in.LoginId.ValueString())
	// Preserve sensitive fields that are not returned by the API
	preservedUserPassword := in.UserPassword

	// Preserve reuse_limitation and time_interval_limitation as well since they are not returned by the API
	preservedReuseLimitation := in.ReuseLimitation
	preservedTimeIntervalLimitation := in.TimeIntervalLimitation

	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	localUserAPI.LoginId = in.LoginId.ValueString()
	respData, err := localUserAPI.Get()

	if err != nil {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not read nd local user, unexpected error: %v %v", err, respData),
		)
		return
	}

	var localUserResp NDFCLocalUserModel
	err = json.Unmarshal(respData, &localUserResp)
	if err != nil {
		dg.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not unmarshal nd local user response, unexpected error: %v", err),
		)
		return
	}

	in.SetModelData(&localUserResp)

	// Restore sensitive fields after SetModelData (API does not return them)
	in.UserPassword = preservedUserPassword

	// Restore reuse_limitation and time_interval_limitation only when they were
	// known (i.e. user-supplied) prior to the API call, since the API does not
	// return these fields. If the value was unknown (user did not specify) and
	// the API did not return it, leave it as Null so Terraform does not see an
	// unknown value after apply.
	if !preservedReuseLimitation.IsNull() && !preservedReuseLimitation.IsUnknown() {
		in.ReuseLimitation = preservedReuseLimitation
	} else {
		in.ReuseLimitation = types.Int64Null()
	}
	if !preservedTimeIntervalLimitation.IsNull() && !preservedTimeIntervalLimitation.IsUnknown() {
		in.TimeIntervalLimitation = preservedTimeIntervalLimitation
	} else {
		in.TimeIntervalLimitation = types.Int64Null()
	}
}

// rscUpdateLocalUser updates a nd local user with the provided payload.
// If the plan's user_password (sourced from the configuration file) matches
// the prior state's user_password, the password field is omitted from the
// payload to avoid resending an unchanged password to the API.
func (r *localUserNdResource) rscUpdateLocalUser(ctx context.Context, dg *diag.Diagnostics, plan *LocalUserModel, state *LocalUserModel) {
	log.Printf("[INFO] Update nd_local_user login_id=%s", plan.LoginId.ValueString())
	inData := plan.GetModelData()

	// The plan's user_password reflects the value from the configuration file.
	// If it matches the prior state's user_password (i.e. the user did not change
	// the password in the config), omit the password from the update payload to
	// avoid resending an unchanged password to the API.
	planPasswordSet := !plan.UserPassword.IsNull() && !plan.UserPassword.IsUnknown()
	statePasswordSet := state != nil && !state.UserPassword.IsNull() && !state.UserPassword.IsUnknown()
	passwordUnchanged := statePasswordSet && plan.UserPassword.ValueString() == state.UserPassword.ValueString()
	if planPasswordSet && passwordUnchanged {
		log.Printf("[DEBUG] Skipping user_password in update payload for login_id=%s (unchanged)", plan.LoginId.ValueString())
		inData.UserPassword = ""
	}

	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	localUserAPI.LoginId = plan.LoginId.ValueString()

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating ND Local User",
			fmt.Sprintf("Could not update nd local user, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating ND Local User: error=%s", err.Error())
		return
	}
	res, err := localUserAPI.Put(inDataBytes, nil)

	if err != nil {
		dg.AddError(
			"Error Updating ND Local User",
			fmt.Sprintf("Could not update nd local user, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating ND Local User: error=%s", err.Error())
		return
	}
	// Read the updated nd local user
	r.rscGetLocalUser(ctx, dg, plan)
}

// rscDeleteLocalUser deletes a nd local user by login_id
func (r *localUserNdResource) rscDeleteLocalUser(ctx context.Context, dg *diag.Diagnostics, state *LocalUserModel) {
	log.Printf("[INFO] Delete nd_local_user login_id=%s", state.LoginId.ValueString())
	localUserAPI := api.NewLocalUserAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	localUserAPI.LoginId = state.LoginId.ValueString()

	res, err := localUserAPI.Delete(nil)
	if err != nil {
		dg.AddError(
			"Error Deleting ND Local User",
			fmt.Sprintf("Could not delete nd local user, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting ND Local User: error=%s", err.Error())
		return
	}
}
