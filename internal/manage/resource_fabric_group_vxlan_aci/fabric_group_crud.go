// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_group_vxlan_aci

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	nd "github.com/netascode/go-nd"
)

// The manage_1.2.224 contract uses baseFabric for POST /fabrics and PUT
// /fabrics/{fabricName}, a direct fabric object for GET, and no DELETE body.
// POST/PUT may return 202, so completion is determined by canonical GET.
const (
	fabricGroupReadbackTimeout  = 2 * time.Minute
	fabricGroupReadbackInterval = 2 * time.Second
)

// request reuses fabric endpoint construction and the standard CRUD lock scope.
// All verbs suppress payload logging because even GET/error bodies may contain
// stored authentication material. Callers report only the error, never the body.
func (r *fabricGroupVxlanAciResource) request(ctx context.Context, name, method string, payload []byte) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("fabric-group name is required")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	fabricAPI := api.NewFabricAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	fabricAPI.FabricName = url.PathEscape(name)
	guard := ndapi.Acquire(fabricAPI.FabricScope(), fabricAPI.RscName(), ndapi.LockCRUD)
	defer guard.Release()

	mods := (&ndapi.APIOptions{DisablePayloadLog: true}).Mods(payload)
	mods = append(mods, func(req *nd.Req) {
		req.HttpReq = req.HttpReq.WithContext(ctx)
	})
	var body []byte
	var err error
	var endpoint string
	switch method {
	case http.MethodGet:
		endpoint = fabricAPI.GetUrl()
		body, err = fabricAPI.Client.GetRawJson(endpoint, mods...)
	case http.MethodPost:
		endpoint = fabricAPI.PostUrl()
		result, requestErr := fabricAPI.Client.Post(endpoint, string(payload), mods...)
		body, err = []byte(result.Raw), requestErr
	case http.MethodPut:
		endpoint = fabricAPI.PutUrl()
		result, requestErr := fabricAPI.Client.Put(endpoint, string(payload), mods...)
		body, err = []byte(result.Raw), requestErr
	case http.MethodDelete:
		endpoint = fabricAPI.DeleteUrl()
		result, requestErr := fabricAPI.Client.Delete(endpoint, "", mods...)
		body, err = []byte(result.Raw), requestErr
	default:
		return nil, fmt.Errorf("unsupported fabric-group request method %q", method)
	}
	return body, ndapi.ClassifyRequestError(method, endpoint, body, err)
}

func (r *fabricGroupVxlanAciResource) read(ctx context.Context, model *FabricGroupVxlanAciModel, preserveOptional bool, dg *diag.Diagnostics) bool {
	name := model.GetFabricName()
	tflog.Debug(ctx, "Reading VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": name})
	body, err := r.request(ctx, name, http.MethodGet, nil)
	if errors.Is(err, ndapi.ErrNotFound) {
		return false
	}
	if err != nil {
		tflog.Error(ctx, "Fabric group GET failed", map[string]interface{}{"fabric_name": name})
		dg.AddError("Error Reading Fabric Group", fmt.Sprintf("Could not read fabric group %q: %s", name, err.Error()))
		return false
	}
	var data resource_fabric_common.NDFCFabricCommonModel
	if err := json.Unmarshal(body, &data); err != nil {
		tflog.Error(ctx, "Could not decode fabric group response", map[string]interface{}{"fabric_name": name})
		// JSON type errors can contain a rejected value. Omit response and decoder
		// details because the same object may include a sensitive BGP key.
		dg.AddError("Invalid Fabric Group Response", fmt.Sprintf("Could not decode the API response for fabric group %q.", name))
		return false
	}
	dg.Append(model.setModelData(&data, preserveOptional)...)
	if dg.HasError() {
		tflog.Error(ctx, "Could not hydrate fabric group state", map[string]interface{}{"fabric_name": name})
		return false
	}
	tflog.Debug(ctx, "Read VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": name})
	return true
}

// write returns whether ND accepted the write, including when later readback
// fails. It updates model only when all requested settings have converged.
func (r *fabricGroupVxlanAciResource) write(ctx context.Context, model *FabricGroupVxlanAciModel, method string, dg *diag.Diagnostics) bool {
	payload, diags := model.marshalPayload(ctx, method == http.MethodPost)
	dg.Append(diags...)
	if dg.HasError() {
		return false
	}
	name := model.GetFabricName()
	if _, err := r.request(ctx, name, method, payload); err != nil {
		tflog.Error(ctx, "Fabric group write failed", map[string]interface{}{"fabric_name": name, "method": method})
		dg.AddError("Error Writing Fabric Group", fmt.Sprintf("%s for fabric group %q failed: %s", method, name, err.Error()))
		return false
	}

	pollCtx, cancel := context.WithTimeout(ctx, fabricGroupReadbackTimeout)
	defer cancel()
	for {
		if err := pollCtx.Err(); err != nil {
			tflog.Error(ctx, "Fabric group readback did not converge", map[string]interface{}{"fabric_name": name, "method": method})
			dg.AddError("Error Waiting for Fabric Group", fmt.Sprintf("ND accepted %s for fabric group %q, but its settings did not become readable and match the plan: %s", method, name, err.Error()))
			return true
		}
		// Always start from the plan. An intermediate response must not replace
		// values needed for secret preservation or the convergence comparison.
		actual := *model
		found := r.read(pollCtx, &actual, true, dg)
		if dg.HasError() {
			return true
		}
		if found {
			matches, diags := fabricGroupReadbackMatches(pollCtx, model, &actual)
			dg.Append(diags...)
			if dg.HasError() {
				return true
			}
			if matches {
				*model = actual
				return true
			}
		}
		tflog.Debug(ctx, "Waiting for fabric group settings to converge", map[string]interface{}{"fabric_name": name})
		timer := time.NewTimer(fabricGroupReadbackInterval)
		select {
		case <-pollCtx.Done():
			timer.Stop()
		case <-timer.C:
		}
	}
}

func (r *fabricGroupVxlanAciResource) delete(ctx context.Context, model *FabricGroupVxlanAciModel, dg *diag.Diagnostics) {
	name := model.GetFabricName()
	tflog.Debug(ctx, "Deleting VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": name})
	// Refuse to delete a different fabric kind that now occupies the same name.
	existing := *model
	if !r.read(ctx, &existing, true, dg) {
		return // An explicit GET 404 already satisfies deletion.
	}
	_, err := r.request(ctx, name, http.MethodDelete, nil)
	if err != nil && !errors.Is(err, ndapi.ErrNotFound) {
		tflog.Error(ctx, "Fabric group DELETE failed", map[string]interface{}{"fabric_name": name})
		dg.AddError("Error Deleting Fabric Group", fmt.Sprintf("Could not delete fabric group %q: %s. Remove member fabrics before deleting the group.", name, err.Error()))
		return
	}
	// DELETE returns 204 synchronously in the API contract.
	tflog.Debug(ctx, "Deleted VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": name})
}

func fabricGroupReadbackMatches(ctx context.Context, expected, actual *FabricGroupVxlanAciModel) (bool, diag.Diagnostics) {
	var wanted, returned types.Object
	var diags diag.Diagnostics
	objectType := FabricGroupVxlanAciResourceSchema(ctx).Type()
	diags.Append(tfsdk.ValueFrom(ctx, expected, objectType, &wanted)...)
	diags.Append(tfsdk.ValueFrom(ctx, actual, objectType, &returned)...)
	if diags.HasError() {
		return false, diags
	}
	attributes := returned.Attributes()
	for name, value := range wanted.Attributes() {
		// Operational security-group status can change independently of settings.
		if name == "security_group_status" || value.IsUnknown() {
			continue
		}
		if !value.Equal(attributes[name]) {
			return false, diags
		}
	}
	return true, diags
}
