// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_inventory_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCInventorySwitchModel struct {
	Id                       string                `json:"id,omitempty"`
	FabricName               string                `json:"fabricName,omitempty"`
	Mode                     string                `json:"mode,omitempty"`
	SeedIpCollection         []string              `json:"seedIpCollection,omitempty"`
	PreserveConfig           *bool                 `json:"preserveConfig,omitempty"`
	WaitForBootstrap         string                `json:"-"`
	WaitForDiscover          string                `json:"-"`
	WaitForReady             string                `json:"-"`
	SwitchDetail             NDFCSwitchDetailValue `json:"-"`
	DiscoveryUsername        string                `json:"username,omitempty"`
	DiscoveryPassword        string                `json:"password,omitempty"`
	BootstrapPassword        string                `json:"-"`
	UseNewCredentials        bool                  `json:"-"`
	DiscoveryCredForLan      bool                  `json:"-"`
	PlatformType             string                `json:"platformType,omitempty"`
	SnmpV3AuthProtocol       string                `json:"snmpV3AuthProtocol,omitempty"`
	RemoteCredentialStore    string                `json:"remoteCredentialStore,omitempty"`
	RemoteCredentialStoreKey string                `json:"remoteCredentialStoreKey,omitempty"`
	SourceInterfaceName      string                `json:"-"`
	SourceVrfName            string                `json:"-"`
	MaxHop                   *int64                `json:"maxHop,omitempty"`
}

type NDFCSwitchDetailValue struct {
	Hostname              string `json:"hostname,omitempty"`
	SerialNumber          string `json:"serialNumber,omitempty"`
	IpAddress             string `json:"ip,omitempty"`
	FabricManagementIp    string `json:"fabricManagementIp,omitempty"`
	Model                 string `json:"model,omitempty"`
	SoftwareVersion       string `json:"softwareVersion,omitempty"`
	SoftwareImage         string `json:"softwareImage,omitempty"`
	SwitchRole            string `json:"switchRole,omitempty"`
	PreserveConfig        *bool  `json:"preserveConfig,omitempty"`
	Status                string `json:"status,omitempty"`
	StatusReason          string `json:"statusReason,omitempty"`
	GatewayIpMask         string `json:"gatewayIpMask,omitempty"`
	DiscoveryAuthProtocol string `json:"discoveryAuthProtocol,omitempty"`
	VdcId                 *int64 `json:"vdcId,omitempty"`
	VdcMac                string `json:"vdcMac,omitempty"`
	SwitchPassword        string `json:"password,omitempty"`
	IpAddrBootstrap       string `json:"ipAddress,omitempty"`
	PublicKey             string `json:"publicKey,omitempty"`
	Fingerprint           string `json:"fingerPrint,omitempty"`
	InInventory           *bool  `json:"inInventory,omitempty"`
}

func (v *InventorySwitchModel) SetModelData(jsonData *NDFCInventorySwitchModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Mode != "" {
		v.Mode = types.StringValue(jsonData.Mode)
	} else {
		v.Mode = types.StringNull()
	}

	if jsonData.PreserveConfig != nil {
		v.PreserveConfig = types.BoolValue(*jsonData.PreserveConfig)

	} else {
		v.PreserveConfig = types.BoolNull()
	}

	if jsonData.WaitForBootstrap != "" {
		v.WaitForBootstrap = types.StringValue(jsonData.WaitForBootstrap)
	} else {
		v.WaitForBootstrap = types.StringNull()
	}

	if jsonData.WaitForDiscover != "" {
		v.WaitForDiscover = types.StringValue(jsonData.WaitForDiscover)
	} else {
		v.WaitForDiscover = types.StringNull()
	}

	if jsonData.WaitForReady != "" {
		v.WaitForReady = types.StringValue(jsonData.WaitForReady)
	} else {
		v.WaitForReady = types.StringNull()
	}

	v.SwitchDetail.SetValue(&jsonData.SwitchDetail)
	v.SwitchDetail.state = attr.ValueStateKnown

	if jsonData.DiscoveryUsername != "" {
		v.DiscoveryUsername = types.StringValue(jsonData.DiscoveryUsername)
	} else {
		v.DiscoveryUsername = types.StringNull()
	}

	if jsonData.DiscoveryPassword != "" {
		v.DiscoveryPassword = types.StringValue(jsonData.DiscoveryPassword)
	} else {
		v.DiscoveryPassword = types.StringNull()
	}

	if jsonData.BootstrapPassword != "" {
		v.BootstrapPassword = types.StringValue(jsonData.BootstrapPassword)
	} else {
		v.BootstrapPassword = types.StringNull()
	}

	v.UseNewCredentials = types.BoolValue(jsonData.UseNewCredentials)

	v.DiscoveryCredForLan = types.BoolValue(jsonData.DiscoveryCredForLan)
	if jsonData.PlatformType != "" {
		v.PlatformType = types.StringValue(jsonData.PlatformType)
	} else {
		v.PlatformType = types.StringNull()
	}

	if jsonData.SnmpV3AuthProtocol != "" {
		v.SnmpV3AuthProtocol = types.StringValue(jsonData.SnmpV3AuthProtocol)
	} else {
		v.SnmpV3AuthProtocol = types.StringNull()
	}

	if jsonData.RemoteCredentialStore != "" {
		v.RemoteCredentialStore = types.StringValue(jsonData.RemoteCredentialStore)
	} else {
		v.RemoteCredentialStore = types.StringNull()
	}

	if jsonData.RemoteCredentialStoreKey != "" {
		v.RemoteCredentialStoreKey = types.StringValue(jsonData.RemoteCredentialStoreKey)
	} else {
		v.RemoteCredentialStoreKey = types.StringNull()
	}

	if jsonData.SourceInterfaceName != "" {
		v.SourceInterfaceName = types.StringValue(jsonData.SourceInterfaceName)
	} else {
		v.SourceInterfaceName = types.StringNull()
	}

	if jsonData.SourceVrfName != "" {
		v.SourceVrfName = types.StringValue(jsonData.SourceVrfName)
	} else {
		v.SourceVrfName = types.StringNull()
	}

	return err
}

func (v *SwitchDetailValue) SetValue(jsonData *NDFCSwitchDetailValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.IpAddress != "" {
		v.IpAddress = types.StringValue(jsonData.IpAddress)
	} else if jsonData.FabricManagementIp != "" {
		v.IpAddress = types.StringValue(jsonData.FabricManagementIp)
	} else {
		v.IpAddress = types.StringNull()
	}

	if jsonData.Model != "" {
		v.Model = types.StringValue(jsonData.Model)
	} else {
		v.Model = types.StringNull()
	}

	if jsonData.SoftwareVersion != "" {
		v.SoftwareVersion = types.StringValue(jsonData.SoftwareVersion)
	} else {
		v.SoftwareVersion = types.StringNull()
	}

	if jsonData.SoftwareImage != "" {
		v.SoftwareImage = types.StringValue(jsonData.SoftwareImage)
	} else {
		v.SoftwareImage = types.StringNull()
	}

	if jsonData.SwitchRole != "" {
		v.SwitchRole = types.StringValue(jsonData.SwitchRole)
	} else {
		v.SwitchRole = types.StringNull()
	}

	if jsonData.Status != "" {
		v.Status = types.StringValue(jsonData.Status)
	} else {
		v.Status = types.StringNull()
	}

	if jsonData.StatusReason != "" {
		v.StatusReason = types.StringValue(jsonData.StatusReason)
	} else {
		v.StatusReason = types.StringNull()
	}

	if jsonData.GatewayIpMask != "" {
		v.GatewayIpMask = types.StringValue(jsonData.GatewayIpMask)
	} else {
		v.GatewayIpMask = types.StringNull()
	}

	if jsonData.DiscoveryAuthProtocol != "" {
		v.DiscoveryAuthProtocol = types.StringValue(jsonData.DiscoveryAuthProtocol)
	} else {
		v.DiscoveryAuthProtocol = types.StringNull()
	}

	if jsonData.VdcId != nil {
		v.VdcId = types.Int64Value(*jsonData.VdcId)

	} else {
		v.VdcId = types.Int64Null()
	}

	if jsonData.VdcMac != "" {
		v.VdcMac = types.StringValue(jsonData.VdcMac)
	} else {
		v.VdcMac = types.StringNull()
	}

	return err
}

func (v InventorySwitchModel) GetModelData() *NDFCInventorySwitchModel {
	var data = new(NDFCInventorySwitchModel)

	//MARSHAL_BODY

	if !v.Id.IsNull() && !v.Id.IsUnknown() {
		data.Id = v.Id.ValueString()
	} else {
		data.Id = ""
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.Mode.IsNull() && !v.Mode.IsUnknown() {
		data.Mode = v.Mode.ValueString()
	} else {
		data.Mode = ""
	}

	if !v.PreserveConfig.IsNull() && !v.PreserveConfig.IsUnknown() {
		data.PreserveConfig = new(bool)
		*data.PreserveConfig = v.PreserveConfig.ValueBool()
	} else {
		data.PreserveConfig = nil
	}

	if !v.WaitForBootstrap.IsNull() && !v.WaitForBootstrap.IsUnknown() {
		data.WaitForBootstrap = v.WaitForBootstrap.ValueString()
	} else {
		data.WaitForBootstrap = ""
	}

	if !v.WaitForDiscover.IsNull() && !v.WaitForDiscover.IsUnknown() {
		data.WaitForDiscover = v.WaitForDiscover.ValueString()
	} else {
		data.WaitForDiscover = ""
	}

	if !v.WaitForReady.IsNull() && !v.WaitForReady.IsUnknown() {
		data.WaitForReady = v.WaitForReady.ValueString()
	} else {
		data.WaitForReady = ""
	}

	if !v.DiscoveryUsername.IsNull() && !v.DiscoveryUsername.IsUnknown() {
		data.DiscoveryUsername = v.DiscoveryUsername.ValueString()
	} else {
		data.DiscoveryUsername = ""
	}

	if !v.DiscoveryPassword.IsNull() && !v.DiscoveryPassword.IsUnknown() {
		data.DiscoveryPassword = v.DiscoveryPassword.ValueString()
	} else {
		data.DiscoveryPassword = ""
	}

	if !v.BootstrapPassword.IsNull() && !v.BootstrapPassword.IsUnknown() {
		data.BootstrapPassword = v.BootstrapPassword.ValueString()
	} else {
		data.BootstrapPassword = ""
	}

	if !v.UseNewCredentials.IsNull() && !v.UseNewCredentials.IsUnknown() {
		data.UseNewCredentials = v.UseNewCredentials.ValueBool()
	}

	if !v.DiscoveryCredForLan.IsNull() && !v.DiscoveryCredForLan.IsUnknown() {
		data.DiscoveryCredForLan = v.DiscoveryCredForLan.ValueBool()
	}

	if !v.SnmpV3AuthProtocol.IsNull() && !v.SnmpV3AuthProtocol.IsUnknown() {
		data.SnmpV3AuthProtocol = v.SnmpV3AuthProtocol.ValueString()
	} else {
		data.SnmpV3AuthProtocol = ""
	}

	if !v.RemoteCredentialStore.IsNull() && !v.RemoteCredentialStore.IsUnknown() {
		data.RemoteCredentialStore = v.RemoteCredentialStore.ValueString()
	} else {
		data.RemoteCredentialStore = ""
	}

	if !v.RemoteCredentialStoreKey.IsNull() && !v.RemoteCredentialStoreKey.IsUnknown() {
		data.RemoteCredentialStoreKey = v.RemoteCredentialStoreKey.ValueString()
	} else {
		data.RemoteCredentialStoreKey = ""
	}

	if !v.SourceInterfaceName.IsNull() && !v.SourceInterfaceName.IsUnknown() {
		data.SourceInterfaceName = v.SourceInterfaceName.ValueString()
	} else {
		data.SourceInterfaceName = ""
	}

	if !v.SourceVrfName.IsNull() && !v.SourceVrfName.IsUnknown() {
		data.SourceVrfName = v.SourceVrfName.ValueString()
	} else {
		data.SourceVrfName = ""
	}

	//MARSHAL_BODY

	// Nested types SwitchDetail # hostname
	if !v.SwitchDetail.Hostname.IsNull() && !v.SwitchDetail.Hostname.IsUnknown() {
		data.SwitchDetail.Hostname = v.SwitchDetail.Hostname.ValueString()
	} else {
		data.SwitchDetail.Hostname = ""
	}

	// Nested types SwitchDetail # serial_number
	if !v.SwitchDetail.SerialNumber.IsNull() && !v.SwitchDetail.SerialNumber.IsUnknown() {
		data.SwitchDetail.SerialNumber = v.SwitchDetail.SerialNumber.ValueString()
	} else {
		data.SwitchDetail.SerialNumber = ""
	}

	// Nested types SwitchDetail # ip_address
	if !v.SwitchDetail.IpAddress.IsNull() && !v.SwitchDetail.IpAddress.IsUnknown() {
		data.SwitchDetail.IpAddress = v.SwitchDetail.IpAddress.ValueString()
	} else {
		data.SwitchDetail.IpAddress = ""
	}

	// Nested types SwitchDetail # model
	if !v.SwitchDetail.Model.IsNull() && !v.SwitchDetail.Model.IsUnknown() {
		data.SwitchDetail.Model = v.SwitchDetail.Model.ValueString()
	} else {
		data.SwitchDetail.Model = ""
	}

	// Nested types SwitchDetail # software_version
	if !v.SwitchDetail.SoftwareVersion.IsNull() && !v.SwitchDetail.SoftwareVersion.IsUnknown() {
		data.SwitchDetail.SoftwareVersion = v.SwitchDetail.SoftwareVersion.ValueString()
	} else {
		data.SwitchDetail.SoftwareVersion = ""
	}

	// Nested types SwitchDetail # software_image
	if !v.SwitchDetail.SoftwareImage.IsNull() && !v.SwitchDetail.SoftwareImage.IsUnknown() {
		data.SwitchDetail.SoftwareImage = v.SwitchDetail.SoftwareImage.ValueString()
	} else {
		data.SwitchDetail.SoftwareImage = ""
	}

	// Nested types SwitchDetail # switch_role
	if !v.SwitchDetail.SwitchRole.IsNull() && !v.SwitchDetail.SwitchRole.IsUnknown() {
		data.SwitchDetail.SwitchRole = v.SwitchDetail.SwitchRole.ValueString()
	} else {
		data.SwitchDetail.SwitchRole = ""
	}

	// Nested types SwitchDetail # gateway_ip_mask
	if !v.SwitchDetail.GatewayIpMask.IsNull() && !v.SwitchDetail.GatewayIpMask.IsUnknown() {
		data.SwitchDetail.GatewayIpMask = v.SwitchDetail.GatewayIpMask.ValueString()
	} else {
		data.SwitchDetail.GatewayIpMask = ""
	}

	// Nested types SwitchDetail # discovery_auth_protocol
	if !v.SwitchDetail.DiscoveryAuthProtocol.IsNull() && !v.SwitchDetail.DiscoveryAuthProtocol.IsUnknown() {
		data.SwitchDetail.DiscoveryAuthProtocol = v.SwitchDetail.DiscoveryAuthProtocol.ValueString()
	} else {
		data.SwitchDetail.DiscoveryAuthProtocol = ""
	}

	// Nested types SwitchDetail # vdc_id
	if !v.SwitchDetail.VdcId.IsNull() && !v.SwitchDetail.VdcId.IsUnknown() {
		data.SwitchDetail.VdcId = new(int64)
		*data.SwitchDetail.VdcId = v.SwitchDetail.VdcId.ValueInt64()

	} else {
		data.SwitchDetail.VdcId = nil
	}

	// Nested types SwitchDetail # vdc_mac
	if !v.SwitchDetail.VdcMac.IsNull() && !v.SwitchDetail.VdcMac.IsUnknown() {
		data.SwitchDetail.VdcMac = v.SwitchDetail.VdcMac.ValueString()
	} else {
		data.SwitchDetail.VdcMac = ""
	}

	return data
}
