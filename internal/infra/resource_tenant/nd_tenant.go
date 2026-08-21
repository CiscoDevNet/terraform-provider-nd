// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

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

// rscCreateTenant posts the tenant first, followed by its fabric associations.
// If any association fails, it removes all tenant associations and deletes the
// newly created tenant so the resource is not left partially created.
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
		Items: make([]tenantFabricAssociationRequestItem, 0, len(desiredAssociations)),
	}
	for _, fabricName := range sortedTenantFabricAssociationKeys(desiredAssociations) {
		associationPayload.Items = append(associationPayload.Items, newTenantFabricAssociationRequestItem(id, fabricName, desiredAssociations[fabricName], true, false))
	}

	var associationDiags diag.Diagnostics
	r.rscPostTenantFabricAssociations(ctx, &associationDiags, associationPayload, tenantFabricAssociationStageCreate)
	if associationDiags.HasError() {
		log.Printf("[ERROR] Fabric association creation failed for tenant id=%s; rolling back tenant creation", id)

		// Use separate diagnostics for rollback. The association diagnostics
		// already contain an error, and passing them to the delete helpers would
		// cause their HasError checks to stop cleanup before the tenant is deleted.
		var rollbackDiags diag.Diagnostics
		r.rscDeleteTenant(ctx, &rollbackDiags, input)

		dg.Append(associationDiags...)
		if rollbackDiags.HasError() {
			dg.AddError(
				"Error Rolling Back Tenant Creation",
				fmt.Sprintf("Tenant %q was created, but fabric association creation failed and the tenant could not be completely removed. Manual cleanup may be required.", id),
			)
			dg.Append(rollbackDiags...)
			return
		}

		log.Printf("[INFO] Rolled back tenant creation after fabric association failure: id=%s", id)
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

// rscPutTenant updates the tenant fields handled by the infra tenant API.
func (r *tenantResource) rscPutTenant(dg *diag.Diagnostics, tenantModel *TenantModel) {
	const errorSummary = "Error Updating Tenant"

	id := tenantModel.Id.ValueString()
	inData := tenantModel.GetModelData()
	tenantAPI := api.NewTenantAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = id

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			errorSummary,
			fmt.Sprintf("Could not update tenant, Data Marshall error: %s", err.Error()),
		)
		log.Printf("[ERROR] %s id=%s: error=%s", errorSummary, id, err.Error())
		return
	}

	res, err := tenantAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			errorSummary,
			fmt.Sprintf("Could not update tenant, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] %s id=%s: error=%s", errorSummary, id, err.Error())
	}
}

// rscUpdateTenant posts fabric association changes before updating description.
// If any association fails, it restores the complete previous association set
// and returns without changing the tenant description.
func (r *tenantResource) rscUpdateTenant(ctx context.Context, dg *diag.Diagnostics, oldState *TenantModel, tenantModel *TenantModel) {
	if tenantModel == nil {
		dg.AddError(
			"Invalid Input",
			"The tenant model is nil",
		)
		return
	}

	id := tenantModel.Id.ValueString()

	var associationDiags diag.Diagnostics
	rollbackRequired := r.rscSyncConfiguredTenantFabricAssociations(ctx, &associationDiags, oldState, tenantModel)
	if associationDiags.HasError() {
		dg.Append(associationDiags...)
		if rollbackRequired {
			log.Printf("[ERROR] Fabric association regular update failed for tenant id=%s; rolling back tenant update", id)

			// The association endpoint can partially apply a multi-item request. Read
			// the backend and reconcile it to the complete association set from the
			// previous Terraform state.
			var associationRollbackDiags diag.Diagnostics
			r.rscRestoreTenantFabricAssociations(ctx, &associationRollbackDiags, oldState, tenantFabricAssociationStageRegularUpdateRollback)

			if associationRollbackDiags.HasError() {
				dg.AddError(
					"Error Rolling Back Tenant Fabric Associations",
					tenantFabricAssociationStageMessage(tenantFabricAssociationStageRegularUpdateRollback, fmt.Sprintf("The fabric association update for tenant %q failed and the previous association configuration could not be completely restored. Manual cleanup may be required.", id)),
				)
				dg.Append(associationRollbackDiags...)
			} else {
				log.Printf("[INFO] Rolled back tenant fabric associations after regular update failure: id=%s", id)
			}
		}

		return
	}

	if oldState != nil && !oldState.Description.Equal(tenantModel.Description) {
		r.rscPutTenant(dg, tenantModel)
		if dg.HasError() {
			return
		}
	}

	if r.rscGetTenant(ctx, dg, tenantModel) {
		dg.AddError(
			"Error Updating Tenant",
			fmt.Sprintf("Tenant %q was not found after update", id),
		)
		return
	}
}

// rscDeleteTenant removes all backend fabric associations before deleting the tenant.
// If association removal fails, it returns without deleting the tenant; otherwise,
// it deletes the tenant and treats an already-missing tenant as successfully removed.
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

	r.rscDeleteTenantFabricAssociations(ctx, dg, state.Name.ValueString())
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
