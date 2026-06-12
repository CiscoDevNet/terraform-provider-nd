package resource_tenant

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func isTenantNotFound(respData []byte, err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())
	respMsg := strings.ToLower(string(respData))
	if strings.Contains(errMsg, "statuscode 404") ||
		strings.Contains(errMsg, "status code 404") ||
		strings.Contains(errMsg, "404 not found") {
		return true
	}

	return strings.Contains(respMsg, "404") && strings.Contains(respMsg, "not found")
}

// rscCreateTenant creates a tenant resource.
func (r *tenantResource) rscCreateTenant(ctx context.Context, dg *diag.Diagnostics, input *TenantModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()
	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	tenantPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Could not create tenant, Data Marshall error: %v", err),
		)
		return
	}

	res, err := tenantAPI.Post(tenantPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Could not create tenant, unexpected error: %v %v", err, res),
		)
		return
	}

	if r.rscGetTenant(ctx, dg, input) {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Tenant %q was not found after create", input.Name.ValueString()),
		)
		return
	}
}

// rscGetTenant retrieves tenant information by name.
func (r *tenantResource) rscGetTenant(ctx context.Context, dg *diag.Diagnostics, state *TenantModel) bool {
	if state == nil {
		dg.AddError(
			"Invalid Input",
			"The state model is nil",
		)
		return false
	}

	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = state.Name.ValueString()

	respData, err := tenantAPI.Get()
	if err != nil {
		if isTenantNotFound(respData, err) {
			return true
		}
		dg.AddError(
			"Error Reading Tenant",
			fmt.Sprintf("Could not read tenant, unexpected error: %v %v", err, string(respData)),
		)
		return false
	}

	var tenantResp NDFCTenantModel
	err = json.Unmarshal(respData, &tenantResp)
	if err != nil {
		dg.AddError(
			"Error Reading Tenant",
			fmt.Sprintf("Could not unmarshal tenant response, unexpected error: %v", err),
		)
		return false
	}

	tenantResp.Id = state.Name.ValueString()
	state.SetModelData(&tenantResp)
	return false
}

// rscUpdateTenant updates a tenant resource.
func (r *tenantResource) rscUpdateTenant(ctx context.Context, dg *diag.Diagnostics, tenantModel *TenantModel) {
	if tenantModel == nil {
		dg.AddError(
			"Invalid Input",
			"The tenant model is nil",
		)
		return
	}

	inData := tenantModel.GetModelData()
	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = tenantModel.Name.ValueString()

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant",
			fmt.Sprintf("Could not update tenant, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Tenant: error=%s", err.Error())
		return
	}

	res, err := tenantAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant",
			fmt.Sprintf("Could not update tenant, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating Tenant: error=%s", err.Error())
		return
	}

	if r.rscGetTenant(ctx, dg, tenantModel) {
		dg.AddError(
			"Error Updating Tenant",
			fmt.Sprintf("Tenant %q was not found after update", tenantModel.Name.ValueString()),
		)
		return
	}
}

// rscDeleteTenant deletes a tenant resource by name.
func (r *tenantResource) rscDeleteTenant(ctx context.Context, dg *diag.Diagnostics, state *TenantModel) {
	if state == nil {
		dg.AddError(
			"Invalid Input",
			"The state model is nil",
		)
		return
	}

	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = state.Name.ValueString()

	res, err := tenantAPI.Delete()
	if err != nil {
		if isTenantNotFound(nil, err) {
			return
		}
		dg.AddError(
			"Error Deleting Tenant",
			fmt.Sprintf("Could not delete tenant, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Tenant: error=%s", err.Error())
		return
	}
}
