// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_vpc_pair

import (
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVpcPairModel struct {
	FabricName         string                  `json:"fabricName,omitempty"`
	SwitchId           string                  `json:"switchId,omitempty"`
	PeerSwitchId       string                  `json:"peerSwitchId,omitempty"`
	UseVirtualPeerlink *bool                   `json:"useVirtualPeerLink,omitempty"`
	VpcPairDetails     NDFCVpcPairDetailsValue `json:"vpcPairDetails,omitempty"`
}

type NDFCVpcPairDetailsValue struct {
	TemplateType               string   `json:"type,omitempty"`
	AdminState                 *bool    `json:"adminState,omitempty"`
	AllowedVlans               string   `json:"allowedVlans,omitempty"`
	DomainId                   *int64   `json:"domainId,omitempty"`
	EnableMirrorConfig         *bool    `json:"enableMirrorConfig,omitempty"`
	FabricPathSwitchId         *int64   `json:"fabricPathSwitchId,omitempty"`
	IsVpcPlus                  *bool    `json:"isVpcPlus,omitempty"`
	IsVteps                    *bool    `json:"isVteps,omitempty"`
	KeepAliveHoldTimeout       *int64   `json:"keepAliveHoldTimeout,omitempty"`
	KeepAliveVrf               string   `json:"keepAliveVrf,omitempty"`
	NveInterface               *int64   `json:"nveInterface,omitempty"`
	PeerSwitchMemberInterfaces []string `json:"peerSwitchMemberInterfaces,omitempty"`
	PeerSwitchKeepAliveLocalIp string   `json:"peerSwitchKeepAliveLocalIp,omitempty"`
	PeerSwitchNativeVlan       *int64   `json:"peerSwitchNativeVlan,omitempty"`
	PeerSwitchPoDescription    string   `json:"peerSwitchPoDescription,omitempty"`
	PeerSwitchPoId             *int64   `json:"peerSwitchPoId,omitempty"`
	PeerSwitchSourceLoopback   *int64   `json:"peerSwitchSourceLoopback,omitempty"`
	PoMode                     string   `json:"poMode,omitempty"`
	SwitchMemberInterfaces     []string `json:"switchMemberInterfaces,omitempty"`
	SwitchKeepAliveLocalIp     string   `json:"switchKeepAliveLocalIp,omitempty"`
	SwitchNativeVlan           *int64   `json:"switchNativeVlan,omitempty"`
	SwitchPoDescription        string   `json:"switchPoDescription,omitempty"`
	SwitchPoId                 *int64   `json:"switchPoId,omitempty"`
	SwitchSourceLoopback       *int64   `json:"switchSourceLoopback,omitempty"`
}

func (v *VpcPairModel) SetModelData(jsonData *NDFCVpcPairModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	if jsonData.PeerSwitchId != "" {
		v.PeerSwitchId = types.StringValue(jsonData.PeerSwitchId)
	} else {
		v.PeerSwitchId = types.StringNull()
	}

	if jsonData.UseVirtualPeerlink != nil {
		v.UseVirtualPeerlink = types.BoolValue(*jsonData.UseVirtualPeerlink)

	} else {
		v.UseVirtualPeerlink = types.BoolNull()
	}

	v.VpcPairDetails.SetValue(&jsonData.VpcPairDetails)

	return err
}

func (v *VpcPairDetailsValue) SetValue(jsonData *NDFCVpcPairDetailsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	valueStateKnown := false
	if jsonData.TemplateType != "" {
		v.TemplateType = types.StringValue(jsonData.TemplateType)
		valueStateKnown = true
	} else {
		v.TemplateType = types.StringNull()
	}

	if jsonData.AdminState != nil {
		v.AdminState = types.BoolValue(*jsonData.AdminState)
		valueStateKnown = true

	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.AllowedVlans)
		valueStateKnown = true
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.DomainId != nil {
		v.DomainId = types.Int64Value(*jsonData.DomainId)
		valueStateKnown = true
	} else {
		v.DomainId = types.Int64Null()
	}

	if jsonData.EnableMirrorConfig != nil {
		v.EnableMirrorConfig = types.BoolValue(*jsonData.EnableMirrorConfig)
		valueStateKnown = true

	} else {
		v.EnableMirrorConfig = types.BoolNull()
	}

	if jsonData.FabricPathSwitchId != nil {
		v.FabricPathSwitchId = types.Int64Value(*jsonData.FabricPathSwitchId)
		valueStateKnown = true
	} else {
		v.FabricPathSwitchId = types.Int64Null()
	}

	if jsonData.IsVpcPlus != nil {
		v.IsVpcPlus = types.BoolValue(*jsonData.IsVpcPlus)
		valueStateKnown = true

	} else {
		v.IsVpcPlus = types.BoolNull()
	}

	if jsonData.IsVteps != nil {
		v.IsVteps = types.BoolValue(*jsonData.IsVteps)
		valueStateKnown = true

	} else {
		v.IsVteps = types.BoolNull()
	}

	if jsonData.KeepAliveHoldTimeout != nil {
		v.KeepAliveHoldTimeout = types.Int64Value(*jsonData.KeepAliveHoldTimeout)
		valueStateKnown = true
	} else {
		v.KeepAliveHoldTimeout = types.Int64Null()
	}

	if jsonData.KeepAliveVrf != "" {
		v.KeepAliveVrf = types.StringValue(jsonData.KeepAliveVrf)
		valueStateKnown = true
	} else {
		v.KeepAliveVrf = types.StringNull()
	}

	if jsonData.NveInterface != nil {
		v.NveInterface = types.Int64Value(*jsonData.NveInterface)
		valueStateKnown = true
	} else {
		v.NveInterface = types.Int64Null()
	}

	if len(jsonData.PeerSwitchMemberInterfaces) == 0 {
		log.Printf("v.PeerSwitchMemberInterfaces is empty")
		v.PeerSwitchMemberInterfaces = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.PeerSwitchMemberInterfaces))
		for i, item := range jsonData.PeerSwitchMemberInterfaces {
			listData[i] = types.StringValue(item)
		}
		v.PeerSwitchMemberInterfaces, err = types.SetValue(types.StringType, listData)
		valueStateKnown = true
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if jsonData.PeerSwitchKeepAliveLocalIp != "" {
		v.PeerSwitchKeepAliveLocalIp = types.StringValue(jsonData.PeerSwitchKeepAliveLocalIp)
		valueStateKnown = true
	} else {
		v.PeerSwitchKeepAliveLocalIp = types.StringNull()
	}

	if jsonData.PeerSwitchNativeVlan != nil {
		v.PeerSwitchNativeVlan = types.Int64Value(*jsonData.PeerSwitchNativeVlan)
		valueStateKnown = true
	} else {
		v.PeerSwitchNativeVlan = types.Int64Null()
	}

	if jsonData.PeerSwitchPoDescription != "" {
		v.PeerSwitchPoDescription = types.StringValue(jsonData.PeerSwitchPoDescription)
		valueStateKnown = true
	} else {
		v.PeerSwitchPoDescription = types.StringNull()
	}

	if jsonData.PeerSwitchPoId != nil {
		v.PeerSwitchPoId = types.Int64Value(*jsonData.PeerSwitchPoId)
		valueStateKnown = true
	} else {
		v.PeerSwitchPoId = types.Int64Null()
	}

	if jsonData.PeerSwitchSourceLoopback != nil {
		v.PeerSwitchSourceLoopback = types.Int64Value(*jsonData.PeerSwitchSourceLoopback)
		valueStateKnown = true
	} else {
		v.PeerSwitchSourceLoopback = types.Int64Null()
	}

	if jsonData.PoMode != "" {
		v.PoMode = types.StringValue(jsonData.PoMode)
		valueStateKnown = true
	} else {
		v.PoMode = types.StringNull()
	}

	if len(jsonData.SwitchMemberInterfaces) == 0 {
		log.Printf("v.SwitchMemberInterfaces is empty")
		v.SwitchMemberInterfaces = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SwitchMemberInterfaces))
		for i, item := range jsonData.SwitchMemberInterfaces {
			listData[i] = types.StringValue(item)
		}
		v.SwitchMemberInterfaces, err = types.SetValue(types.StringType, listData)
		valueStateKnown = true
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if jsonData.SwitchKeepAliveLocalIp != "" {
		v.SwitchKeepAliveLocalIp = types.StringValue(jsonData.SwitchKeepAliveLocalIp)
		valueStateKnown = true
	} else {
		v.SwitchKeepAliveLocalIp = types.StringNull()
	}

	if jsonData.SwitchNativeVlan != nil {
		v.SwitchNativeVlan = types.Int64Value(*jsonData.SwitchNativeVlan)
		valueStateKnown = true
	} else {
		v.SwitchNativeVlan = types.Int64Null()
	}

	if jsonData.SwitchPoDescription != "" {
		v.SwitchPoDescription = types.StringValue(jsonData.SwitchPoDescription)
		valueStateKnown = true
	} else {
		v.SwitchPoDescription = types.StringNull()
	}

	if jsonData.SwitchPoId != nil {
		v.SwitchPoId = types.Int64Value(*jsonData.SwitchPoId)
		valueStateKnown = true
	} else {
		v.SwitchPoId = types.Int64Null()
	}

	if jsonData.SwitchSourceLoopback != nil {
		v.SwitchSourceLoopback = types.Int64Value(*jsonData.SwitchSourceLoopback)
		valueStateKnown = true
	} else {
		v.SwitchSourceLoopback = types.Int64Null()
	}

	if valueStateKnown {
		v.state = attr.ValueStateKnown
	}

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

	if !v.SwitchId.IsNull() && !v.SwitchId.IsUnknown() {
		data.SwitchId = v.SwitchId.ValueString()
	} else {
		data.SwitchId = ""
	}

	return data
}
