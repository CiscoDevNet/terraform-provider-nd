// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
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
	"github.com/netascode/go-nd"
)

// fabricGroupCRUDModel supplies group-specific policy for one Terraform operation.
type fabricGroupCRUDModel struct {
	*FabricGroupVxlanAciModel
	importing                bool
	verifyRouteServerRemoval bool
	writeAccepted            bool
}

var (
	_ resource_fabric_common.FabricModel         = (*fabricGroupCRUDModel)(nil)
	_ resource_fabric_common.FabricAPIOverride   = (*fabricGroupCRUDModel)(nil)
	_ resource_fabric_common.FabricWriteReadback = (*fabricGroupCRUDModel)(nil)
)

func (v *fabricGroupCRUDModel) WrapFabricAPI(ctx context.Context, fabricAPI *api.FabricAPI) resource_fabric_common.FabricCRUDAPI {
	name := v.GetFabricName()
	fabricAPI.FabricName = url.PathEscape(name)
	return fabricGroupAPI{ctx: ctx, fabricAPI: fabricAPI, name: name}
}

type fabricGroupAPI struct {
	ctx       context.Context
	fabricAPI *api.FabricAPI
	name      string
}

func (a fabricGroupAPI) Get() ([]byte, error) {
	return a.request(http.MethodGet, nil)
}

func (a fabricGroupAPI) Post(payload []byte, _ *ndapi.APIOptions) (nd.Res, error) {
	_, err := a.request(http.MethodPost, payload)
	return nd.Res{}, err
}

func (a fabricGroupAPI) Put(payload []byte, _ *ndapi.APIOptions) (nd.Res, error) {
	_, err := a.request(http.MethodPut, payload)
	return nd.Res{}, err
}

func (a fabricGroupAPI) Delete(_ ...[]byte) (nd.Res, error) {
	_, err := a.request(http.MethodDelete, nil)
	return nd.Res{}, err
}

func (a fabricGroupAPI) request(method string, payload []byte) ([]byte, error) {
	name := a.name
	if name == "" {
		return nil, fmt.Errorf("fabric-group name is required")
	}
	if method == http.MethodPut || method == http.MethodDelete {
		if err := verifyFabricGroup(a.ctx, a.fabricAPI, name); err != nil {
			if errors.Is(err, ndapi.ErrNotFound) {
				if method == http.MethodDelete {
					return nil, nil
				}
				return nil, fmt.Errorf("fabric group %q no longer exists. Refresh Terraform state before applying again", name)
			}
			return nil, err
		}
	}
	if method == http.MethodPost || method == http.MethodPut {
		var err error
		payload, err = marshalFabricGroupPayload(payload)
		if err != nil {
			return nil, err
		}
	}
	body, err := requestFabricGroup(a.ctx, a.fabricAPI, method, payload)
	if err != nil {
		switch method {
		case http.MethodPost, http.MethodPut:
			return nil, fmt.Errorf("Error Writing Fabric Group: %s for fabric group %q failed: %w", method, name, err)
		case http.MethodDelete:
			if errors.Is(err, ndapi.ErrNotFound) {
				return nil, nil
			}
			return nil, fmt.Errorf("could not delete fabric group %q: %w. Remove member fabrics before deleting the group", name, err)
		}
	}
	return body, err
}

// requestFabricGroup suppresses bodies for every verb because responses can
// include the BGP key. Use the same per-request lock as the default fabric API.
func requestFabricGroup(ctx context.Context, fabricAPI *api.FabricAPI, method string, payload []byte) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	guard := ndapi.Acquire(fabricAPI.FabricScope(), fabricAPI.RscName(), ndapi.LockCRUD)
	defer guard.Release()
	if err := ctx.Err(); err != nil {
		return nil, err
	}

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
		_, err = fabricAPI.Client.Post(endpoint, string(payload), mods...)
	case http.MethodPut:
		endpoint = fabricAPI.PutUrl()
		_, err = fabricAPI.Client.Put(endpoint, string(payload), mods...)
	case http.MethodDelete:
		endpoint = fabricAPI.DeleteUrl()
		_, err = fabricAPI.Client.Delete(endpoint, "", mods...)
	default:
		return nil, fmt.Errorf("unsupported fabric-group request method %q", method)
	}
	if err != nil {
		return nil, ndapi.ClassifyRequestError(method, endpoint, nil, err)
	}
	return body, nil
}

func verifyFabricGroup(ctx context.Context, fabricAPI *api.FabricAPI, name string) error {
	body, err := requestFabricGroup(ctx, fabricAPI, http.MethodGet, nil)
	if err != nil {
		return err
	}
	var identity struct {
		Name       string `json:"name"`
		Category   string `json:"category"`
		Management struct {
			Type string `json:"type"`
		} `json:"management"`
	}
	if err := json.Unmarshal(body, &identity); err != nil {
		return fmt.Errorf("could not decode the API response for fabric group %q", name)
	}
	return validateFabricGroupIdentity(name, identity.Name, identity.Category, identity.Management.Type)
}

func (v *fabricGroupCRUDModel) ReadAfterFabricWrite(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, method string) {
	v.writeAccepted = true
	pollCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	for pollCtx.Err() == nil {
		attempt := fabricGroupReadbackModel{fabricGroupCRUDModel: v, ctx: pollCtx}
		var readDiags diag.Diagnostics
		found := resource_fabric_common.RscGetFabric(pollCtx, client, &readDiags, &attempt)
		if pollCtx.Err() != nil && !attempt.matches {
			break
		}
		dg.Append(readDiags...)
		if dg.HasError() || (found && attempt.matches) {
			return
		}
		timer := time.NewTimer(2 * time.Second)
		select {
		case <-pollCtx.Done():
			timer.Stop()
		case <-timer.C:
		}
	}
	dg.AddError("Error Waiting for Fabric Group", fmt.Sprintf("ND accepted %s for fabric group %q, but its settings did not become readable and match the plan: %s", method, v.GetFabricName(), pollCtx.Err()))
}

func (v *fabricGroupCRUDModel) SetModelData(data *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	return v.setModelData(data, !v.importing)
}

type fabricGroupReadbackModel struct {
	*fabricGroupCRUDModel
	ctx     context.Context
	matches bool
}

func (v *fabricGroupReadbackModel) SetModelData(data *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var diags diag.Diagnostics
	v.matches, diags = v.readback(v.ctx, data)
	return diags
}

func (v *fabricGroupCRUDModel) readback(ctx context.Context, data *resource_fabric_common.NDFCFabricCommonModel) (bool, diag.Diagnostics) {
	// Each attempt starts from the plan, preserving the original comparison
	// values and the key even when a prior GET returned stale settings.
	actual := *v.FabricGroupVxlanAciModel
	diags := actual.setModelData(data, true)
	if diags.HasError() {
		return false, diags
	}
	if v.verifyRouteServerRemoval && len(data.Management.RouteServerCollection) != 0 {
		return false, diags
	}

	var planned, returned types.Object
	objectType := FabricGroupVxlanAciResourceSchema(ctx).Type()
	diags.Append(tfsdk.ValueFrom(ctx, v.FabricGroupVxlanAciModel, objectType, &planned)...)
	diags.Append(tfsdk.ValueFrom(ctx, &actual, objectType, &returned)...)
	if diags.HasError() {
		return false, diags
	}

	attributes := returned.Attributes()
	for name, value := range planned.Attributes() {
		switch {
		case value.IsUnknown(), name == "security_group_status", name == "multisite_inter_connect_bgp_key":
			// The operational status and the saved key do not establish convergence.
			continue
		case value.IsNull():
			// Omitted attributes are not part of the write. Apply state must
			// retain their planned nulls even if GET returns backend defaults.
			attributes[name] = value
		case !value.Equal(attributes[name]):
			return false, diags
		}
	}

	normalized, normalizationDiags := types.ObjectValue(planned.AttributeTypes(ctx), attributes)
	diags.Append(normalizationDiags...)
	if diags.HasError() {
		return false, diags
	}
	diags.Append(tfsdk.ValueAs(ctx, normalized, &actual)...)
	if diags.HasError() {
		return false, diags
	}
	*v.FabricGroupVxlanAciModel = actual
	return true, diags
}
