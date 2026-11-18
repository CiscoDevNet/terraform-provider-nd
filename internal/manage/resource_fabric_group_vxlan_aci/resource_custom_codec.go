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

const (
	fabricGroupCategory = "fabricGroup"
	fabricGroupType     = "vxlanAci"
)

func (v *FabricGroupVxlanAciModel) GetFabricName() string {
	if !v.Id.IsNull() && !v.Id.IsUnknown() && v.Id.ValueString() != "" {
		return v.Id.ValueString()
	}
	return v.FabricName.ValueString()
}

func (v *FabricGroupVxlanAciModel) GetFabricType() string {
	return fabricGroupType
}

// PreMarshal supplies hidden discriminators. The generated codec owns
// securityGroupTag conversion; no custom JSON methods are needed for that field.
func (v *FabricGroupVxlanAciModel) PreMarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	data.Category = fabricGroupCategory
	data.Management.FabricType = fabricGroupType
	data.Management.SecurityGroupStatus = ""
}

func (v *FabricGroupVxlanAciModel) validateRouteServerConversion(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	// The generated collection codec cannot return conversion diagnostics and
	// panics on invalid elements. Check conversion before calling GetModelData.
	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		var servers []RouteServerCollectionValue
		diags.Append(v.RouteServerCollection.ElementsAs(ctx, &servers, false)...)
	}
	return diags
}

func marshalFabricGroupPayload(payload []byte) ([]byte, error) {
	var group struct {
		Name       string                     `json:"name"`
		Category   string                     `json:"category"`
		Management map[string]json.RawMessage `json:"management"`
	}
	if err := json.Unmarshal(payload, &group); err != nil {
		return nil, fmt.Errorf("could not decode the generated fabric-group request")
	}
	// The group body uses only group fields from the shared model. Unset
	// struct-valued common fields can still be serialized as empty objects.
	for name, value := range group.Management {
		if string(value) == "{}" {
			delete(group.Management, name)
		}
	}
	return json.Marshal(group)
}

func validateFabricGroupIdentity(expectedName, name, category, fabricType string) error {
	if category != fabricGroupCategory || fabricType != fabricGroupType {
		return fmt.Errorf("object %q must have category %s and management.type %s", expectedName, fabricGroupCategory, fabricGroupType)
	}
	if name == "" || name != expectedName {
		return fmt.Errorf("the returned fabric-group name does not match the requested name %q", expectedName)
	}
	return nil
}

func (v *FabricGroupVxlanAciModel) setModelData(data *resource_fabric_common.NDFCFabricCommonModel, preserveOptional bool) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Invalid Fabric Group Response", "The API returned no fabric-group model.")
		return diags
	}
	if err := validateFabricGroupIdentity(v.GetFabricName(), data.FabricName, data.Category, data.Management.FabricType); err != nil {
		diags.AddError("Invalid Fabric Group Response", err.Error())
		return diags
	}
	if data.Management.SecurityGroupTag != nil {
		mode, ok := data.Management.SecurityGroupTag.(string)
		if !ok || (mode != "loose" && mode != "strict") {
			diags.AddError("Invalid Fabric Group Response", "securityGroupTag must be the string loose or strict for a VXLAN-ACI fabric group.")
			return diags
		}
	}
	current := *v
	diags.Append(current.SetModelData(data)...)
	if diags.HasError() {
		return diags
	}
	current.Id = current.FabricName
	// GET cannot restore a configured encrypted key reliably: it may omit,
	// mask, or re-encrypt it. Import deliberately leaves this attribute null.
	current.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey
	if len(data.Management.RouteServerCollection) == 0 && !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0 {
		current.RouteServerCollection = v.RouteServerCollection
	}
	if preserveOptional && !current.ScheduledBackup.ValueBool() {
		// These optional attributes are not computed in the generated schema.
		// Keep omitted values null when GET echoes disabled backup settings.
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
