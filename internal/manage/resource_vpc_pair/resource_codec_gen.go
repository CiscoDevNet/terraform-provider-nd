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
	FabricName          string `json:"-"`
	Switch1SerialNumber string `json:"peerSwitchId,omitempty"`
	Switch2SerialNumber string `json:"switchId,omitempty"`
	UseVirtualPeerlink  *bool  `json:"useVirtualPeerLink,omitempty"`
	VpcAction           string `json:"vpcAction,omitempty"`
	Deploy              bool   `json:"-"`
}

func (v *VpcPairModel) SetModelData(jsonData *NDFCVpcPairModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Switch1SerialNumber != "" {
		v.Switch1SerialNumber = types.StringValue(jsonData.Switch1SerialNumber)
	} else {
		v.Switch1SerialNumber = types.StringNull()
	}

	if jsonData.Switch2SerialNumber != "" {
		v.Switch2SerialNumber = types.StringValue(jsonData.Switch2SerialNumber)
	} else {
		v.Switch2SerialNumber = types.StringNull()
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

	if !v.Switch1SerialNumber.IsNull() && !v.Switch1SerialNumber.IsUnknown() {
		data.Switch1SerialNumber = v.Switch1SerialNumber.ValueString()
	} else {
		data.Switch1SerialNumber = ""
	}

	if !v.Switch2SerialNumber.IsNull() && !v.Switch2SerialNumber.IsUnknown() {
		data.Switch2SerialNumber = v.Switch2SerialNumber.ValueString()
	} else {
		data.Switch2SerialNumber = ""
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
