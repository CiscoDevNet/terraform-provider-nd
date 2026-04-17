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
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCInventorySwitchModel struct {
	FabricName               string                       `json:"fabricName,omitempty"`
	Mode                     string                       `json:"mode,omitempty"`
	SeedIpCollection         []string                     `json:"seedIpCollection,omitempty"`
	PreserveConfig           *bool                        `json:"preserveConfig,omitempty"`
	Switches                 map[string]NDFCSwitchesValue `json:"-"`
	Username                 string                       `json:"username,omitempty"`
	Password                 string                       `json:"password,omitempty"`
	PlatformType             string                       `json:"platformType,omitempty"`
	SnmpV3AuthProtocol       string                       `json:"snmpV3AuthProtocol,omitempty"`
	RemoteCredentialStore    string                       `json:"remoteCredentialStore,omitempty"`
	RemoteCredentialStoreKey string                       `json:"remoteCredentialStoreKey,omitempty"`
	MaxHop                   *int64                       `json:"maxHop,omitempty"`
	Recalculate              bool                         `json:"-"`
	Deploy                   bool                         `json:"-"`
}

type NDFCSwitchesValue struct {
	Hostname              string `json:"hostname,omitempty"`
	SerialNumber          string `json:"serialNumber,omitempty"`
	IpAddress             string `json:"ip,omitempty"`
	FabricManagementIp    string `json:"fabricManagementIp,omitempty"`
	Model                 string `json:"model,omitempty"`
	SoftwareVersion       string `json:"softwareVersion,omitempty"`
	SwitchRole            string `json:"switchRole,omitempty"`
	PreserveConfig        *bool  `json:"preserveConfig,omitempty"`
	Status                string `json:"status,omitempty"`
	StatusReason          string `json:"statusReason,omitempty"`
	GatewayIpMask         string `json:"gatewayIpMask,omitempty"`
	DiscoveryAuthProtocol string `json:"discoveryAuthProtocol,omitempty"`
	PoapPassword          string `json:"poapPassword,omitempty"`
	VdcId                 *int64 `json:"vdcId,omitempty"`
	VdcMac                string `json:"vdcMac,omitempty"`
}

func (v *InventorySwitchModel) SetModelData(jsonData *NDFCInventorySwitchModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

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

	if len(jsonData.Switches) == 0 {
		log.Printf("v.Switches is empty")
		v.Switches = types.MapNull(SwitchesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]SwitchesValue)
		for key, item := range jsonData.Switches {
			data := new(SwitchesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in SwitchesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Switches, err = types.MapValueFrom(context.Background(), SwitchesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]SwitchesValue to  Map")

		}
	}
	if jsonData.Username != "" {
		v.Username = types.StringValue(jsonData.Username)
	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Password != "" {
		v.Password = types.StringValue(jsonData.Password)
	} else {
		v.Password = types.StringNull()
	}

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

	if jsonData.MaxHop != nil {
		v.MaxHop = types.Int64Value(*jsonData.MaxHop)

	} else {
		v.MaxHop = types.Int64Null()
	}

	v.Recalculate = types.BoolValue(jsonData.Recalculate)

	v.Deploy = types.BoolValue(jsonData.Deploy)

	return err
}

func (v *SwitchesValue) SetValue(jsonData *NDFCSwitchesValue) diag.Diagnostics {

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

	if jsonData.PoapPassword != "" {
		v.PoapPassword = types.StringValue(jsonData.PoapPassword)
	} else {
		v.PoapPassword = types.StringNull()
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

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Username = v.Username.ValueString()
	} else {
		data.Username = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Password = v.Password.ValueString()
	} else {
		data.Password = ""
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

	if !v.MaxHop.IsNull() && !v.MaxHop.IsUnknown() {
		data.MaxHop = new(int64)
		*data.MaxHop = v.MaxHop.ValueInt64()

	} else {
		data.MaxHop = nil
	}

	if !v.Recalculate.IsNull() && !v.Recalculate.IsUnknown() {
		data.Recalculate = v.Recalculate.ValueBool()
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.Switches.IsNull() && !v.Switches.IsUnknown() {
		elements1 := make(map[string]SwitchesValue, len(v.Switches.Elements()))

		data.Switches = make(map[string]NDFCSwitchesValue)

		diag := v.Switches.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(NDFCSwitchesValue)

			// hostname | String| []| false
			if !ele1.Hostname.IsNull() && !ele1.Hostname.IsUnknown() {

				data1.Hostname = ele1.Hostname.ValueString()
			} else {
				data1.Hostname = ""
			}

			// serial_number | String| []| false
			if !ele1.SerialNumber.IsNull() && !ele1.SerialNumber.IsUnknown() {

				data1.SerialNumber = ele1.SerialNumber.ValueString()
			} else {
				data1.SerialNumber = ""
			}

			// ip_address | String| []| false
			if !ele1.IpAddress.IsNull() && !ele1.IpAddress.IsUnknown() {

				data1.IpAddress = ele1.IpAddress.ValueString()
			} else {
				data1.IpAddress = ""
			}

			// model | String| []| false
			if !ele1.Model.IsNull() && !ele1.Model.IsUnknown() {

				data1.Model = ele1.Model.ValueString()
			} else {
				data1.Model = ""
			}

			// software_version | String| []| false
			if !ele1.SoftwareVersion.IsNull() && !ele1.SoftwareVersion.IsUnknown() {

				data1.SoftwareVersion = ele1.SoftwareVersion.ValueString()
			} else {
				data1.SoftwareVersion = ""
			}

			// switch_role | String| []| false
			if !ele1.SwitchRole.IsNull() && !ele1.SwitchRole.IsUnknown() {

				data1.SwitchRole = ele1.SwitchRole.ValueString()
			} else {
				data1.SwitchRole = ""
			}

			// preserve_config | Bool| []| true
			// status | String| []| false
			// status_reason | String| []| false
			// gateway_ip_mask | String| []| false
			if !ele1.GatewayIpMask.IsNull() && !ele1.GatewayIpMask.IsUnknown() {

				data1.GatewayIpMask = ele1.GatewayIpMask.ValueString()
			} else {
				data1.GatewayIpMask = ""
			}

			// discovery_auth_protocol | String| []| false
			if !ele1.DiscoveryAuthProtocol.IsNull() && !ele1.DiscoveryAuthProtocol.IsUnknown() {

				data1.DiscoveryAuthProtocol = ele1.DiscoveryAuthProtocol.ValueString()
			} else {
				data1.DiscoveryAuthProtocol = ""
			}

			// poap_password | String| []| false
			if !ele1.PoapPassword.IsNull() && !ele1.PoapPassword.IsUnknown() {

				data1.PoapPassword = ele1.PoapPassword.ValueString()
			} else {
				data1.PoapPassword = ""
			}

			// vdc_id | Int64| []| false
			if !ele1.VdcId.IsNull() && !ele1.VdcId.IsUnknown() {

				data1.VdcId = new(int64)
				*data1.VdcId = ele1.VdcId.ValueInt64()

			} else {
				data1.VdcId = nil
			}

			// vdc_mac | String| []| false
			if !ele1.VdcMac.IsNull() && !ele1.VdcMac.IsUnknown() {

				data1.VdcMac = ele1.VdcMac.ValueString()
			} else {
				data1.VdcMac = ""
			}

			data.Switches[k1] = *data1

		}
	}

	return data
}
