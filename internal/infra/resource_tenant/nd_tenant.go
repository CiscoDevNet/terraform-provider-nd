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

// rscCreateTenant creates a tenant resource.
func (r *tenantResource) rscCreateTenant(ctx context.Context, dg *diag.Diagnostics, input *TenantModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	input.Id = input.Name
	id := input.Id.ValueString()

	desiredAssociations := tenantFabricAssociationsForModel(ctx, dg, input)
	if dg.HasError() {
		return
	}

	inData := input.GetModelData()
	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	tenantPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Could not create tenant, Data Marshall error: %s", err.Error()),
		)
		return
	}

	res, err := tenantAPI.Post(tenantPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Could not create tenant, unexpected error: %s %s", err.Error(), res.String()),
		)
		return
	}

	associationPayload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationItem, 0, len(desiredAssociations)),
	}
	for _, association := range desiredAssociations {
		associationPayload.Items = append(associationPayload.Items, newTenantFabricAssociationItem(id, association, true))
	}

	r.rscPostTenantFabricAssociations(dg, associationPayload)
	if dg.HasError() {
		return
	}

	if r.rscGetTenant(ctx, dg, input) {
		dg.AddError(
			"Error Creating Tenant",
			fmt.Sprintf("Tenant %q was not found after fabric association update", id),
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

	id := state.Id.ValueString()
	log.Printf("[INFO] Read nd_tenant id=%s", id)

	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = id

	respData, err := tenantAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading Tenant",
			fmt.Sprintf("Could not read tenant, unexpected error: %s %s", err.Error(), string(respData)),
		)
		return false
	}

	var tenantResp NDFCTenantModel
	err = json.Unmarshal(respData, &tenantResp)
	if err != nil {
		dg.AddError(
			"Error Reading Tenant",
			fmt.Sprintf("Could not unmarshal tenant response, unexpected error: %s", err.Error()),
		)
		return false
	}

	tenantResp.Id = id
	r.rscReadTenantFabricAssociations(dg, &tenantResp, id)
	if dg.HasError() {
		return false
	}

	dg.Append(state.SetModelData(&tenantResp)...)
	if dg.HasError() {
		return false
	}

	return false
}

// rscUpdateTenant updates a tenant resource.
func (r *tenantResource) rscUpdateTenant(ctx context.Context, dg *diag.Diagnostics, oldState *TenantModel, tenantModel *TenantModel) {
	if tenantModel == nil {
		dg.AddError(
			"Invalid Input",
			"The tenant model is nil",
		)
		return
	}

	id := tenantModel.Id.ValueString()

	if oldState != nil && !oldState.Description.Equal(tenantModel.Description) {
		inData := tenantModel.GetModelData()
		tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
		tenantAPI.TenantName = id

		inDataBytes, err := json.Marshal(inData)
		if err != nil {
			dg.AddError(
				"Error Updating Tenant",
				fmt.Sprintf("Could not update tenant, Data Marshall error: %s", err.Error()),
			)
			log.Printf("[ERROR] Error Updating Tenant id=%s: error=%s", id, err.Error())
			return
		}

		res, err := tenantAPI.Put(inDataBytes, nil)
		if err != nil {
			dg.AddError(
				"Error Updating Tenant",
				fmt.Sprintf("Could not update tenant, unexpected error: %s %s", err.Error(), res.String()),
			)
			log.Printf("[ERROR] Error Updating Tenant id=%s: error=%s", id, err.Error())
			return
		}
	}

	r.rscSyncConfiguredTenantFabricAssociations(ctx, dg, oldState, tenantModel)
	if dg.HasError() {
		return
	}

	if r.rscGetTenant(ctx, dg, tenantModel) {
		dg.AddError(
			"Error Updating Tenant",
			fmt.Sprintf("Tenant %q was not found after update", id),
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

	id := state.Id.ValueString()
	log.Printf("[INFO] Delete nd_tenant id=%s", id)

	r.rscDeleteTenantFabricAssociations(dg, state)
	if dg.HasError() {
		return
	}

	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = id

	res, err := tenantAPI.Delete()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting Tenant",
			fmt.Sprintf("Could not delete tenant, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] Error Deleting Tenant id=%s: error=%s", id, err.Error())
		return
	}
}
