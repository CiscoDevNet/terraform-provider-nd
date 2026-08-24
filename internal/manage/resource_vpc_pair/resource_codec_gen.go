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
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCVpcPairModel struct {
	FabricName         string                  `json:"-"`
	SwitchId1          string                  `json:"peerSwitchId,omitempty"`
	SwitchId2          string                  `json:"switchId,omitempty"`
	UseVirtualPeerlink *bool                   `json:"useVirtualPeerLink,omitempty"`
	VpcPairDetails     NDFCVpcPairDetailsValue `json:"vpcPairDetails,omitempty"`
	VpcAction          string                  `json:"vpcAction,omitempty"`
	Deploy             bool                    `json:"-"`
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
	LoopbackSecondaryIp        string   `json:"loopbackSecondaryIp,omitempty"`
	NveInterface               *int64   `json:"nveInterface,omitempty"`
	PeerSwitchKeepAliveLocalIp string   `json:"peerSwitchKeepAliveLocalIp,omitempty"`
	PeerSwitchDomainConfig     string   `json:"peerSwitchDomainConfig,omitempty"`
	PeerSwitchMemberInterfaces []string `json:"peerSwitchMemberInterfaces,omitempty"`
	PeerSwitchNativeVlan       *int64   `json:"peerSwitchNativeVlan,omitempty"`
	PeerSwitchPoDescription    string   `json:"peerSwitchPoDescription,omitempty"`
	PeerSwitchPoConfig         string   `json:"peerSwitchPoConfig,omitempty"`
	PeerSwitchPoId             *int64   `json:"peerSwitchPoId,omitempty"`
	PeerSwitchPrimaryIp        string   `json:"peerSwitchPrimaryIp,omitempty"`
	PeerSwitchSourceLoopback   *int64   `json:"peerSwitchSourceLoopback,omitempty"`
	PoMode                     string   `json:"poMode,omitempty"`
	SwitchKeepAliveLocalIp     string   `json:"switchKeepAliveLocalIp,omitempty"`
	SwitchDomainConfig         string   `json:"switchDomainConfig,omitempty"`
	SwitchMemberInterfaces     []string `json:"switchMemberInterfaces,omitempty"`
	SwitchNativeVlan           *int64   `json:"switchNativeVlan,omitempty"`
	SwitchPoDescription        string   `json:"switchPoDescription,omitempty"`
	SwitchPoConfig             string   `json:"switchPoConfig,omitempty"`
	SwitchPoId                 *int64   `json:"switchPoId,omitempty"`
	SwitchPrimaryIp            string   `json:"switchPrimaryIp,omitempty"`
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

	v.VpcPairDetails.SetValue(&jsonData.VpcPairDetails)
	v.VpcPairDetails.state = attr.ValueStateKnown

	v.Deploy = types.BoolValue(jsonData.Deploy)

	return err
}

func (v *VpcPairDetailsValue) SetValue(jsonData *NDFCVpcPairDetailsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.TemplateType != "" {
		v.TemplateType = types.StringValue(jsonData.TemplateType)
	} else {
		v.TemplateType = types.StringNull()
	}

	if jsonData.AdminState != nil {
		v.AdminState = types.BoolValue(*jsonData.AdminState)

	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.DomainId != nil {
		v.DomainId = types.Int64Value(*jsonData.DomainId)

	} else {
		v.DomainId = types.Int64Null()
	}

	if jsonData.EnableMirrorConfig != nil {
		v.EnableMirrorConfig = types.BoolValue(*jsonData.EnableMirrorConfig)

	} else {
		v.EnableMirrorConfig = types.BoolNull()
	}

	if jsonData.FabricPathSwitchId != nil {
		v.FabricPathSwitchId = types.Int64Value(*jsonData.FabricPathSwitchId)

	} else {
		v.FabricPathSwitchId = types.Int64Null()
	}

	if jsonData.IsVpcPlus != nil {
		v.IsVpcPlus = types.BoolValue(*jsonData.IsVpcPlus)

	} else {
		v.IsVpcPlus = types.BoolNull()
	}

	if jsonData.IsVteps != nil {
		v.IsVteps = types.BoolValue(*jsonData.IsVteps)

	} else {
		v.IsVteps = types.BoolNull()
	}

	if jsonData.KeepAliveHoldTimeout != nil {
		v.KeepAliveHoldTimeout = types.Int64Value(*jsonData.KeepAliveHoldTimeout)

	} else {
		v.KeepAliveHoldTimeout = types.Int64Null()
	}

	if jsonData.KeepAliveVrf != "" {
		v.KeepAliveVrf = types.StringValue(jsonData.KeepAliveVrf)
	} else {
		v.KeepAliveVrf = types.StringNull()
	}

	if jsonData.LoopbackSecondaryIp != "" {
		v.LoopbackSecondaryIp = types.StringValue(jsonData.LoopbackSecondaryIp)
	} else {
		v.LoopbackSecondaryIp = types.StringNull()
	}

	if jsonData.NveInterface != nil {
		v.NveInterface = types.Int64Value(*jsonData.NveInterface)

	} else {
		v.NveInterface = types.Int64Null()
	}

	if jsonData.PeerSwitchKeepAliveLocalIp != "" {
		v.PeerSwitchKeepAliveLocalIp = types.StringValue(jsonData.PeerSwitchKeepAliveLocalIp)
	} else {
		v.PeerSwitchKeepAliveLocalIp = types.StringNull()
	}

	if jsonData.PeerSwitchDomainConfig != "" {
		v.PeerSwitchDomainConfig = types.StringValue(jsonData.PeerSwitchDomainConfig)
	} else {
		v.PeerSwitchDomainConfig = types.StringNull()
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
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.PeerSwitchNativeVlan != nil {
		v.PeerSwitchNativeVlan = types.Int64Value(*jsonData.PeerSwitchNativeVlan)

	} else {
		v.PeerSwitchNativeVlan = types.Int64Null()
	}

	if jsonData.PeerSwitchPoDescription != "" {
		v.PeerSwitchPoDescription = types.StringValue(jsonData.PeerSwitchPoDescription)
	} else {
		v.PeerSwitchPoDescription = types.StringNull()
	}

	if jsonData.PeerSwitchPoConfig != "" {
		v.PeerSwitchPoConfig = types.StringValue(jsonData.PeerSwitchPoConfig)
	} else {
		v.PeerSwitchPoConfig = types.StringNull()
	}

	if jsonData.PeerSwitchPoId != nil {
		v.PeerSwitchPoId = types.Int64Value(*jsonData.PeerSwitchPoId)

	} else {
		v.PeerSwitchPoId = types.Int64Null()
	}

	if jsonData.PeerSwitchPrimaryIp != "" {
		v.PeerSwitchPrimaryIp = types.StringValue(jsonData.PeerSwitchPrimaryIp)
	} else {
		v.PeerSwitchPrimaryIp = types.StringNull()
	}

	if jsonData.PeerSwitchSourceLoopback != nil {
		v.PeerSwitchSourceLoopback = types.Int64Value(*jsonData.PeerSwitchSourceLoopback)

	} else {
		v.PeerSwitchSourceLoopback = types.Int64Null()
	}

	if jsonData.PoMode != "" {
		v.PoMode = types.StringValue(jsonData.PoMode)
	} else {
		v.PoMode = types.StringNull()
	}

	if jsonData.SwitchKeepAliveLocalIp != "" {
		v.SwitchKeepAliveLocalIp = types.StringValue(jsonData.SwitchKeepAliveLocalIp)
	} else {
		v.SwitchKeepAliveLocalIp = types.StringNull()
	}

	if jsonData.SwitchDomainConfig != "" {
		v.SwitchDomainConfig = types.StringValue(jsonData.SwitchDomainConfig)
	} else {
		v.SwitchDomainConfig = types.StringNull()
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
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.SwitchNativeVlan != nil {
		v.SwitchNativeVlan = types.Int64Value(*jsonData.SwitchNativeVlan)

	} else {
		v.SwitchNativeVlan = types.Int64Null()
	}

	if jsonData.SwitchPoDescription != "" {
		v.SwitchPoDescription = types.StringValue(jsonData.SwitchPoDescription)
	} else {
		v.SwitchPoDescription = types.StringNull()
	}

	if jsonData.SwitchPoConfig != "" {
		v.SwitchPoConfig = types.StringValue(jsonData.SwitchPoConfig)
	} else {
		v.SwitchPoConfig = types.StringNull()
	}

	if jsonData.SwitchPoId != nil {
		v.SwitchPoId = types.Int64Value(*jsonData.SwitchPoId)

	} else {
		v.SwitchPoId = types.Int64Null()
	}

	if jsonData.SwitchPrimaryIp != "" {
		v.SwitchPrimaryIp = types.StringValue(jsonData.SwitchPrimaryIp)
	} else {
		v.SwitchPrimaryIp = types.StringNull()
	}

	if jsonData.SwitchSourceLoopback != nil {
		v.SwitchSourceLoopback = types.Int64Value(*jsonData.SwitchSourceLoopback)

	} else {
		v.SwitchSourceLoopback = types.Int64Null()
	}

	return err
}

func (v VpcPairModel) GetModelData() *NDFCVpcPairModel {
	var data = new(NDFCVpcPairModel)

	//MARSHAL_BODY

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

	//MARSHAL_BODY

	// Nested types VpcPairDetails # template_type
	if !v.VpcPairDetails.TemplateType.IsNull() && !v.VpcPairDetails.TemplateType.IsUnknown() {
		data.VpcPairDetails.TemplateType = v.VpcPairDetails.TemplateType.ValueString()
	} else {
		data.VpcPairDetails.TemplateType = ""
	}

	// Nested types VpcPairDetails # admin_state
	if !v.VpcPairDetails.AdminState.IsNull() && !v.VpcPairDetails.AdminState.IsUnknown() {
		data.VpcPairDetails.AdminState = new(bool)
		*data.VpcPairDetails.AdminState = v.VpcPairDetails.AdminState.ValueBool()
	} else {
		data.VpcPairDetails.AdminState = nil
	}

	// Nested types VpcPairDetails # allowed_vlans
	if !v.VpcPairDetails.AllowedVlans.IsNull() && !v.VpcPairDetails.AllowedVlans.IsUnknown() {
		data.VpcPairDetails.AllowedVlans = v.VpcPairDetails.AllowedVlans.ValueString()
	} else {
		data.VpcPairDetails.AllowedVlans = ""
	}

	// Nested types VpcPairDetails # domain_id
	if !v.VpcPairDetails.DomainId.IsNull() && !v.VpcPairDetails.DomainId.IsUnknown() {
		data.VpcPairDetails.DomainId = new(int64)
		*data.VpcPairDetails.DomainId = v.VpcPairDetails.DomainId.ValueInt64()

	} else {
		data.VpcPairDetails.DomainId = nil
	}

	// Nested types VpcPairDetails # enable_mirror_config
	if !v.VpcPairDetails.EnableMirrorConfig.IsNull() && !v.VpcPairDetails.EnableMirrorConfig.IsUnknown() {
		data.VpcPairDetails.EnableMirrorConfig = new(bool)
		*data.VpcPairDetails.EnableMirrorConfig = v.VpcPairDetails.EnableMirrorConfig.ValueBool()
	} else {
		data.VpcPairDetails.EnableMirrorConfig = nil
	}

	// Nested types VpcPairDetails # fabric_path_switch_id
	if !v.VpcPairDetails.FabricPathSwitchId.IsNull() && !v.VpcPairDetails.FabricPathSwitchId.IsUnknown() {
		data.VpcPairDetails.FabricPathSwitchId = new(int64)
		*data.VpcPairDetails.FabricPathSwitchId = v.VpcPairDetails.FabricPathSwitchId.ValueInt64()

	} else {
		data.VpcPairDetails.FabricPathSwitchId = nil
	}

	// Nested types VpcPairDetails # is_vpc_plus
	if !v.VpcPairDetails.IsVpcPlus.IsNull() && !v.VpcPairDetails.IsVpcPlus.IsUnknown() {
		data.VpcPairDetails.IsVpcPlus = new(bool)
		*data.VpcPairDetails.IsVpcPlus = v.VpcPairDetails.IsVpcPlus.ValueBool()
	} else {
		data.VpcPairDetails.IsVpcPlus = nil
	}

	// Nested types VpcPairDetails # is_vteps
	if !v.VpcPairDetails.IsVteps.IsNull() && !v.VpcPairDetails.IsVteps.IsUnknown() {
		data.VpcPairDetails.IsVteps = new(bool)
		*data.VpcPairDetails.IsVteps = v.VpcPairDetails.IsVteps.ValueBool()
	} else {
		data.VpcPairDetails.IsVteps = nil
	}

	// Nested types VpcPairDetails # keep_alive_hold_timeout
	if !v.VpcPairDetails.KeepAliveHoldTimeout.IsNull() && !v.VpcPairDetails.KeepAliveHoldTimeout.IsUnknown() {
		data.VpcPairDetails.KeepAliveHoldTimeout = new(int64)
		*data.VpcPairDetails.KeepAliveHoldTimeout = v.VpcPairDetails.KeepAliveHoldTimeout.ValueInt64()

	} else {
		data.VpcPairDetails.KeepAliveHoldTimeout = nil
	}

	// Nested types VpcPairDetails # keep_alive_vrf
	if !v.VpcPairDetails.KeepAliveVrf.IsNull() && !v.VpcPairDetails.KeepAliveVrf.IsUnknown() {
		data.VpcPairDetails.KeepAliveVrf = v.VpcPairDetails.KeepAliveVrf.ValueString()
	} else {
		data.VpcPairDetails.KeepAliveVrf = ""
	}

	// Nested types VpcPairDetails # loopback_secondary_ip
	if !v.VpcPairDetails.LoopbackSecondaryIp.IsNull() && !v.VpcPairDetails.LoopbackSecondaryIp.IsUnknown() {
		data.VpcPairDetails.LoopbackSecondaryIp = v.VpcPairDetails.LoopbackSecondaryIp.ValueString()
	} else {
		data.VpcPairDetails.LoopbackSecondaryIp = ""
	}

	// Nested types VpcPairDetails # nve_interface
	if !v.VpcPairDetails.NveInterface.IsNull() && !v.VpcPairDetails.NveInterface.IsUnknown() {
		data.VpcPairDetails.NveInterface = new(int64)
		*data.VpcPairDetails.NveInterface = v.VpcPairDetails.NveInterface.ValueInt64()

	} else {
		data.VpcPairDetails.NveInterface = nil
	}

	// Nested types VpcPairDetails # peer_switch_keep_alive_local_ip
	if !v.VpcPairDetails.PeerSwitchKeepAliveLocalIp.IsNull() && !v.VpcPairDetails.PeerSwitchKeepAliveLocalIp.IsUnknown() {
		data.VpcPairDetails.PeerSwitchKeepAliveLocalIp = v.VpcPairDetails.PeerSwitchKeepAliveLocalIp.ValueString()
	} else {
		data.VpcPairDetails.PeerSwitchKeepAliveLocalIp = ""
	}

	// Nested types VpcPairDetails # peer_switch_domain_config
	if !v.VpcPairDetails.PeerSwitchDomainConfig.IsNull() && !v.VpcPairDetails.PeerSwitchDomainConfig.IsUnknown() {
		data.VpcPairDetails.PeerSwitchDomainConfig = v.VpcPairDetails.PeerSwitchDomainConfig.ValueString()
	} else {
		data.VpcPairDetails.PeerSwitchDomainConfig = ""
	}

	// Nested types VpcPairDetails # peer_switch_member_interfaces
	if !v.VpcPairDetails.PeerSwitchMemberInterfaces.IsNull() && !v.VpcPairDetails.PeerSwitchMemberInterfaces.IsUnknown() {
		listStringData := make([]string, len(v.VpcPairDetails.PeerSwitchMemberInterfaces.Elements()))
		dg := v.VpcPairDetails.PeerSwitchMemberInterfaces.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.VpcPairDetails.PeerSwitchMemberInterfaces = make([]string, len(listStringData))
		copy(data.VpcPairDetails.PeerSwitchMemberInterfaces, listStringData)
	}

	// Nested types VpcPairDetails # peer_switch_native_vlan
	if !v.VpcPairDetails.PeerSwitchNativeVlan.IsNull() && !v.VpcPairDetails.PeerSwitchNativeVlan.IsUnknown() {
		data.VpcPairDetails.PeerSwitchNativeVlan = new(int64)
		*data.VpcPairDetails.PeerSwitchNativeVlan = v.VpcPairDetails.PeerSwitchNativeVlan.ValueInt64()

	} else {
		data.VpcPairDetails.PeerSwitchNativeVlan = nil
	}

	// Nested types VpcPairDetails # peer_switch_po_description
	if !v.VpcPairDetails.PeerSwitchPoDescription.IsNull() && !v.VpcPairDetails.PeerSwitchPoDescription.IsUnknown() {
		data.VpcPairDetails.PeerSwitchPoDescription = v.VpcPairDetails.PeerSwitchPoDescription.ValueString()
	} else {
		data.VpcPairDetails.PeerSwitchPoDescription = ""
	}

	// Nested types VpcPairDetails # peer_switch_po_config
	if !v.VpcPairDetails.PeerSwitchPoConfig.IsNull() && !v.VpcPairDetails.PeerSwitchPoConfig.IsUnknown() {
		data.VpcPairDetails.PeerSwitchPoConfig = v.VpcPairDetails.PeerSwitchPoConfig.ValueString()
	} else {
		data.VpcPairDetails.PeerSwitchPoConfig = ""
	}

	// Nested types VpcPairDetails # peer_switch_po_id
	if !v.VpcPairDetails.PeerSwitchPoId.IsNull() && !v.VpcPairDetails.PeerSwitchPoId.IsUnknown() {
		data.VpcPairDetails.PeerSwitchPoId = new(int64)
		*data.VpcPairDetails.PeerSwitchPoId = v.VpcPairDetails.PeerSwitchPoId.ValueInt64()

	} else {
		data.VpcPairDetails.PeerSwitchPoId = nil
	}

	// Nested types VpcPairDetails # peer_switch_primary_ip
	if !v.VpcPairDetails.PeerSwitchPrimaryIp.IsNull() && !v.VpcPairDetails.PeerSwitchPrimaryIp.IsUnknown() {
		data.VpcPairDetails.PeerSwitchPrimaryIp = v.VpcPairDetails.PeerSwitchPrimaryIp.ValueString()
	} else {
		data.VpcPairDetails.PeerSwitchPrimaryIp = ""
	}

	// Nested types VpcPairDetails # peer_switch_source_loopback
	if !v.VpcPairDetails.PeerSwitchSourceLoopback.IsNull() && !v.VpcPairDetails.PeerSwitchSourceLoopback.IsUnknown() {
		data.VpcPairDetails.PeerSwitchSourceLoopback = new(int64)
		*data.VpcPairDetails.PeerSwitchSourceLoopback = v.VpcPairDetails.PeerSwitchSourceLoopback.ValueInt64()

	} else {
		data.VpcPairDetails.PeerSwitchSourceLoopback = nil
	}

	// Nested types VpcPairDetails # po_mode
	if !v.VpcPairDetails.PoMode.IsNull() && !v.VpcPairDetails.PoMode.IsUnknown() {
		data.VpcPairDetails.PoMode = v.VpcPairDetails.PoMode.ValueString()
	} else {
		data.VpcPairDetails.PoMode = ""
	}

	// Nested types VpcPairDetails # switch_keep_alive_local_ip
	if !v.VpcPairDetails.SwitchKeepAliveLocalIp.IsNull() && !v.VpcPairDetails.SwitchKeepAliveLocalIp.IsUnknown() {
		data.VpcPairDetails.SwitchKeepAliveLocalIp = v.VpcPairDetails.SwitchKeepAliveLocalIp.ValueString()
	} else {
		data.VpcPairDetails.SwitchKeepAliveLocalIp = ""
	}

	// Nested types VpcPairDetails # switch_domain_config
	if !v.VpcPairDetails.SwitchDomainConfig.IsNull() && !v.VpcPairDetails.SwitchDomainConfig.IsUnknown() {
		data.VpcPairDetails.SwitchDomainConfig = v.VpcPairDetails.SwitchDomainConfig.ValueString()
	} else {
		data.VpcPairDetails.SwitchDomainConfig = ""
	}

	// Nested types VpcPairDetails # switch_member_interfaces
	if !v.VpcPairDetails.SwitchMemberInterfaces.IsNull() && !v.VpcPairDetails.SwitchMemberInterfaces.IsUnknown() {
		listStringData := make([]string, len(v.VpcPairDetails.SwitchMemberInterfaces.Elements()))
		dg := v.VpcPairDetails.SwitchMemberInterfaces.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.VpcPairDetails.SwitchMemberInterfaces = make([]string, len(listStringData))
		copy(data.VpcPairDetails.SwitchMemberInterfaces, listStringData)
	}

	// Nested types VpcPairDetails # switch_native_vlan
	if !v.VpcPairDetails.SwitchNativeVlan.IsNull() && !v.VpcPairDetails.SwitchNativeVlan.IsUnknown() {
		data.VpcPairDetails.SwitchNativeVlan = new(int64)
		*data.VpcPairDetails.SwitchNativeVlan = v.VpcPairDetails.SwitchNativeVlan.ValueInt64()

	} else {
		data.VpcPairDetails.SwitchNativeVlan = nil
	}

	// Nested types VpcPairDetails # switch_po_description
	if !v.VpcPairDetails.SwitchPoDescription.IsNull() && !v.VpcPairDetails.SwitchPoDescription.IsUnknown() {
		data.VpcPairDetails.SwitchPoDescription = v.VpcPairDetails.SwitchPoDescription.ValueString()
	} else {
		data.VpcPairDetails.SwitchPoDescription = ""
	}

	// Nested types VpcPairDetails # switch_po_config
	if !v.VpcPairDetails.SwitchPoConfig.IsNull() && !v.VpcPairDetails.SwitchPoConfig.IsUnknown() {
		data.VpcPairDetails.SwitchPoConfig = v.VpcPairDetails.SwitchPoConfig.ValueString()
	} else {
		data.VpcPairDetails.SwitchPoConfig = ""
	}

	// Nested types VpcPairDetails # switch_po_id
	if !v.VpcPairDetails.SwitchPoId.IsNull() && !v.VpcPairDetails.SwitchPoId.IsUnknown() {
		data.VpcPairDetails.SwitchPoId = new(int64)
		*data.VpcPairDetails.SwitchPoId = v.VpcPairDetails.SwitchPoId.ValueInt64()

	} else {
		data.VpcPairDetails.SwitchPoId = nil
	}

	// Nested types VpcPairDetails # switch_primary_ip
	if !v.VpcPairDetails.SwitchPrimaryIp.IsNull() && !v.VpcPairDetails.SwitchPrimaryIp.IsUnknown() {
		data.VpcPairDetails.SwitchPrimaryIp = v.VpcPairDetails.SwitchPrimaryIp.ValueString()
	} else {
		data.VpcPairDetails.SwitchPrimaryIp = ""
	}

	// Nested types VpcPairDetails # switch_source_loopback
	if !v.VpcPairDetails.SwitchSourceLoopback.IsNull() && !v.VpcPairDetails.SwitchSourceLoopback.IsUnknown() {
		data.VpcPairDetails.SwitchSourceLoopback = new(int64)
		*data.VpcPairDetails.SwitchSourceLoopback = v.VpcPairDetails.SwitchSourceLoopback.ValueInt64()

	} else {
		data.VpcPairDetails.SwitchSourceLoopback = nil
	}

	return data
}
