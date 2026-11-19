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
	"fmt"

	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource_fabric_common.FabricPreMarshal = (*FabricGroupVxlanAciModel)(nil)

func (v *FabricGroupVxlanAciModel) GetFabricName() string {
	if !v.Id.IsNull() && !v.Id.IsUnknown() && v.Id.ValueString() != "" {
		return v.Id.ValueString()
	}
	return v.FabricName.ValueString()
}

// PreMarshal supplies hidden discriminators. The generated codec owns
// securityGroupTag conversion; no custom JSON methods are needed for that field.
func (v *FabricGroupVxlanAciModel) PreMarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	data.Category = "fabricGroup"
	data.Management.FabricType = "vxlanAci"
	data.Management.SecurityGroupStatus = ""
}

func (v *FabricGroupVxlanAciModel) marshalPayload(ctx context.Context, create bool) ([]byte, diag.Diagnostics) {
	var diags diag.Diagnostics
	// The generated collection codec cannot return conversion diagnostics and
	// panics on invalid elements. Check conversion before calling GetModelData.
	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		var servers []RouteServerCollectionValue
		diags.Append(v.RouteServerCollection.ElementsAs(ctx, &servers, false)...)
		if diags.HasError() {
			return nil, diags
		}
	}
	data := v.GetModelData()
	v.PreMarshal(ctx, data)
	body, err := json.Marshal(data.Management)
	if err != nil {
		diags.AddError("Error Preparing Fabric Group", "Could not marshal fabric-group management settings.")
		return nil, diags
	}
	var management map[string]interface{}
	if err := json.Unmarshal(body, &management); err != nil {
		diags.AddError("Error Preparing Fabric Group", "Could not decode generated fabric-group management settings.")
		return nil, diags
	}
	pruneEmptyObjects(management)
	if (!create && v.RouteServerCollection.IsNull()) || (!v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0) {
		// The generated omitempty tag drops empty lists. PUT must send [] to
		// remove route servers when the collection is emptied or omitted.
		management["routeServerCollection"] = []interface{}{}
	}
	// The common model also contains unrelated empty fabric objects. Send only
	// the fabricGroup envelope defined by this resource's API contract.
	body, err = json.Marshal(map[string]interface{}{
		"name":       data.FabricName,
		"category":   data.Category,
		"management": management,
	})
	if err != nil {
		diags.AddError("Error Preparing Fabric Group", "Could not marshal the fabric-group request.")
	}
	return body, diags
}

func pruneEmptyObjects(object map[string]interface{}) {
	for name, value := range object {
		if nested, ok := value.(map[string]interface{}); ok {
			pruneEmptyObjects(nested)
			if len(nested) == 0 {
				delete(object, name)
			}
		}
	}
}

func (v *FabricGroupVxlanAciModel) setModelData(data *resource_fabric_common.NDFCFabricCommonModel, preserveOptional bool) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Invalid Fabric Group Response", "The API returned no fabric-group model.")
		return diags
	}
	if data.Category != "fabricGroup" || data.Management.FabricType != "vxlanAci" {
		diags.AddError("Invalid Fabric Group Type", fmt.Sprintf("Object %q must have category fabricGroup and management.type vxlanAci.", v.GetFabricName()))
		return diags
	}
	if data.FabricName == "" || data.FabricName != v.GetFabricName() {
		diags.AddError("Invalid Fabric Group Response", "The returned fabric-group name does not match the requested name.")
		return diags
	}
	if data.Management.SecurityGroupTag != nil {
		mode, ok := data.Management.SecurityGroupTag.(string)
		if !ok || (mode != "loose" && mode != "strict") {
			diags.AddError("Invalid Fabric Group Response", "securityGroupTag must be the string loose or strict for a VXLAN-ACI fabric group.")
			return diags
		}
	}
	data.Id = data.FabricName
	current := *v
	diags.Append(current.SetModelData(data)...)
	if diags.HasError() {
		return diags
	}
	// GET cannot restore a configured encrypted key reliably: it may omit,
	// mask, or re-encrypt it. Import deliberately leaves this attribute null.
	current.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey
	if len(data.Management.RouteServerCollection) == 0 && !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0 {
		current.RouteServerCollection = v.RouteServerCollection
	}
	if preserveOptional {
		// These optional attributes are not computed in the generated schema.
		// Backend defaults must not turn omitted configuration into known state.
		if v.ScheduledBackup.IsNull() {
			current.ScheduledBackup = types.BoolNull()
		}
		if v.ScheduledBackupTime.IsNull() {
			current.ScheduledBackupTime = types.StringNull()
		}
	}
	*v = current
	return diags
}
