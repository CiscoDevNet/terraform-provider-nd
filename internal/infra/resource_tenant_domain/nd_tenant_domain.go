// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_tenant_domain

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// CRUD helpers are called only by tenantDomainResource lifecycle methods with
// pointers to concrete local models.

// rscCreateTenantDomain creates a tenant domain resource.
func (r *tenantDomainResource) rscCreateTenantDomain(dg *diag.Diagnostics, tenantDomainModel *TenantDomainModel) {
	tenantDomainModel.Id = tenantDomainModel.Name
	id := tenantDomainModel.Id.ValueString()

	inData := tenantDomainModel.GetModelData()
	tenantDomainAPI := api.NewTenantDomainAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	if inData.TenantNames == nil {
		inData.TenantNames = []string{}
	}

	tenantDomainPayload, err := json.Marshal(inData)

	if err != nil {
		dg.AddError(
			"Error Creating Tenant Domain",
			fmt.Sprintf("Could not create tenant domain, Data Marshall error: %s", err.Error()),
		)
		return
	}

	res, err := tenantDomainAPI.Post(tenantDomainPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating Tenant Domain",
			fmt.Sprintf("Could not create tenant domain, unexpected error: %s %s", err.Error(), res.String()),
		)
		return
	}

	if r.rscGetTenantDomain(dg, tenantDomainModel) {
		dg.AddError(
			"Error Creating Tenant Domain",
			fmt.Sprintf("Tenant domain %q was not found after create", id),
		)
		return
	}
}

// rscGetTenantDomain retrieves tenant domain information by name.
func (r *tenantDomainResource) rscGetTenantDomain(dg *diag.Diagnostics, tenantDomainModel *TenantDomainModel) bool {
	id := tenantDomainModel.Id.ValueString()
	log.Printf("[INFO] Read nd_tenant_domain id=%s", id)

	tenantDomainAPI := api.NewTenantDomainAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = id

	respData, err := tenantDomainAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading Tenant Domain",
			fmt.Sprintf("Could not read tenant domain, unexpected error: %s %s", err.Error(), string(respData)),
		)
		return false
	}

	var tenantDomainResp NDFCTenantDomainModel
	err = json.Unmarshal(respData, &tenantDomainResp)
	if err != nil {
		dg.AddError(
			"Error Reading Tenant Domain",
			fmt.Sprintf("Could not unmarshal tenant domain response, unexpected error: %s", err.Error()),
		)
		return false
	}

	tenantDomainResp.Id = id
	dg.Append(tenantDomainModel.SetModelData(&tenantDomainResp)...)
	if dg.HasError() {
		return false
	}

	return false
}

// rscUpdateTenantDomain updates a tenant domain resource.
func (r *tenantDomainResource) rscUpdateTenantDomain(dg *diag.Diagnostics, tenantDomainModel *TenantDomainModel) {
	id := tenantDomainModel.Id.ValueString()
	inData := tenantDomainModel.GetModelData()
	tenantDomainAPI := api.NewTenantDomainAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = id

	if inData.TenantNames == nil {
		inData.TenantNames = []string{}
	}

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant Domain",
			fmt.Sprintf("Could not update tenant domain, Data Marshall error: %s", err.Error()),
		)
		log.Printf("[ERROR] Error Updating Tenant Domain id=%s: error=%s", id, err.Error())
		return
	}

	res, err := tenantDomainAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant Domain",
			fmt.Sprintf("Could not update tenant domain, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] Error Updating Tenant Domain id=%s: error=%s", id, err.Error())
		return
	}

	if r.rscGetTenantDomain(dg, tenantDomainModel) {
		dg.AddError(
			"Error Updating Tenant Domain",
			fmt.Sprintf("Tenant domain %q was not found after update", id),
		)
		return
	}
}

// rscDeleteTenantDomain deletes a tenant domain resource by name.
func (r *tenantDomainResource) rscDeleteTenantDomain(dg *diag.Diagnostics, tenantDomainModel *TenantDomainModel) {
	id := tenantDomainModel.Id.ValueString()
	log.Printf("[INFO] Delete nd_tenant_domain id=%s", id)

	tenantDomainAPI := api.NewTenantDomainAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = id

	res, err := tenantDomainAPI.Delete()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting Tenant Domain",
			fmt.Sprintf("Could not delete tenant domain, unexpected error: %s %s", err.Error(), res.String()),
		)
		log.Printf("[ERROR] Error Deleting Tenant Domain id=%s: error=%s", id, err.Error())
		return
	}
}
