// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_vpc_pair

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVpcPairModel struct {
	FabricName         string `json:"-"`
	SwitchId1          string `json:"peerSwitchId,omitempty"`
	SwitchId2          string `json:"switchId,omitempty"`
	UseVirtualPeerlink *bool  `json:"useVirtualPeerLink,omitempty"`
	VpcAction          string `json:"vpcAction,omitempty"`
	Deploy             bool   `json:"-"`
}

func (v *VpcPairModel) SetModelData(jsonData *NDFCVpcPairModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.SwitchId1 != "" {
		v.SwitchId1 = types.StringValue(jsonData.SwitchId1)
	} else {
		v.SwitchId1 = types.StringNull()
	}

	if jsonData.SwitchId2 != "" {
		v.SwitchId2 = types.StringValue(jsonData.SwitchId2)
	} else {
		v.SwitchId2 = types.StringNull()
	}

	if jsonData.UseVirtualPeerlink != nil {
		v.UseVirtualPeerlink = types.BoolValue(*jsonData.UseVirtualPeerlink)

	} else {
		v.UseVirtualPeerlink = types.BoolNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)

	return err
}

func (v VpcPairModel) GetModelData() *NDFCVpcPairModel {
	var data = new(NDFCVpcPairModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.SwitchId1.IsNull() && !v.SwitchId1.IsUnknown() {
		data.SwitchId1 = v.SwitchId1.ValueString()
	} else {
		data.SwitchId1 = ""
	}

	if !v.SwitchId2.IsNull() && !v.SwitchId2.IsUnknown() {
		data.SwitchId2 = v.SwitchId2.ValueString()
	} else {
		data.SwitchId2 = ""
	}

	if !v.UseVirtualPeerlink.IsNull() && !v.UseVirtualPeerlink.IsUnknown() {
		data.UseVirtualPeerlink = new(bool)
		*data.UseVirtualPeerlink = v.UseVirtualPeerlink.ValueBool()
	} else {
		data.UseVirtualPeerlink = nil
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	return data
}
