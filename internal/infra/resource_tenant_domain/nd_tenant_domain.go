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
	"errors"
	"fmt"
	"log"
	"net/http"

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

	modelDiags, err := r.rscGetTenantDomain(tenantDomainModel)
	dg.Append(modelDiags...)
	if errors.Is(err, ndapi.ErrNotFound) {
		dg.AddError(
			"Error Creating Tenant Domain",
			fmt.Sprintf("Tenant domain %q was not found after create", id),
		)
		return
	}
	if err != nil {
		dg.AddError(
			"Error Creating Tenant Domain",
			fmt.Sprintf("Could not read tenant domain %q after create: %s", id, err.Error()),
		)
		return
	}
}

// rscGetTenantDomain retrieves tenant domain information by name. API and
// response-decoding failures are returned as errors, while model-conversion
// diagnostics are returned without flattening their Terraform details.
func (r *tenantDomainResource) rscGetTenantDomain(tenantDomainModel *TenantDomainModel) (diag.Diagnostics, error) {
	id := tenantDomainModel.Id.ValueString()
	log.Printf("[INFO] Read nd_tenant_domain id=%s", id)

	tenantDomainAPI := api.NewTenantDomainAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = id

	respData, err := tenantDomainAPI.Get()
	err = ndapi.ClassifyRequestError(http.MethodGet, tenantDomainAPI.GetUrl(), respData, err)
	if err != nil {
		responseBody := ""
		var requestErr *ndapi.RequestError
		if errors.As(err, &requestErr) {
			responseBody = string(requestErr.Response)
		}
		return nil, fmt.Errorf("GET tenant domain %q: %w; response: %s", id, err, responseBody)
	}

	var tenantDomainResp NDFCTenantDomainModel
	err = json.Unmarshal(respData, &tenantDomainResp)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal tenant domain %q response: %w", id, err)
	}

	tenantDomainResp.Id = id
	return tenantDomainModel.SetModelData(&tenantDomainResp), nil
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

	modelDiags, err := r.rscGetTenantDomain(tenantDomainModel)
	dg.Append(modelDiags...)
	if errors.Is(err, ndapi.ErrNotFound) {
		dg.AddError(
			"Error Updating Tenant Domain",
			fmt.Sprintf("Tenant domain %q was not found after update", id),
		)
		return
	}
	if err != nil {
		dg.AddError(
			"Error Updating Tenant Domain",
			fmt.Sprintf("Could not read tenant domain %q after update: %s", id, err.Error()),
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
	err = ndapi.ClassifyRequestError(http.MethodDelete, tenantDomainAPI.DeleteUrl(), []byte(res.String()), err)
	if err != nil {
		if errors.Is(err, ndapi.ErrNotFound) {
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
