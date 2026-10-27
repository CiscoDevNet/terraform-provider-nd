// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_inventory_switch

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCInventorySwitchModel struct {
	FabricName   string             `json:"fabricName,omitempty"`
	SerialNumber string             `json:"serialNumber,omitempty"`
	Switches     NDFCSwitchesValues `json:"switches,omitempty"`
}

type NDFCSwitchesValues []NDFCSwitchesValue

type NDFCSwitchesValue struct {
	FabricName      string                  `json:"fabricName,omitempty"`
	SerialNumber    string                  `json:"serialNumber,omitempty"`
	IpAddress       string                  `json:"fabricManagementIp,omitempty"`
	Hostname        string                  `json:"hostname,omitempty"`
	Model           string                  `json:"model,omitempty"`
	SoftwareVersion string                  `json:"softwareVersion,omitempty"`
	SoftwareImage   string                  `json:"softwareImage,omitempty"`
	SwitchRole      string                  `json:"switchRole,omitempty"`
	PreserveConfig  *bool                   `json:"preserveConfig,omitempty"`
	VdcId           *int64                  `json:"vdcId,omitempty"`
	VdcMac          string                  `json:"vdcMac,omitempty"`
	InInventory     *bool                   `json:"inInventory,omitempty"`
	SwitchId        string                  `json:"switchId,omitempty"`
	FabricType      string                  `json:"fabricType,omitempty"`
	SystemUptime    string                  `json:"systemUpTime,omitempty"`
	AdvisoryLevel   string                  `json:"advisoryLevel,omitempty"`
	AlertSuspend    string                  `json:"alertSuspend,omitempty"`
	AnomalyLevel    string                  `json:"anomalyLevel,omitempty"`
	VpcConfigured   *bool                   `json:"vpcConfigured,omitempty"`
	AdditionalData  NDFCAdditionalDataValue `json:"additionalData,omitempty"`
}

type NDFCAdditionalDataValue struct {
	Status                        string `json:"discoveryStatus,omitempty"`
	PlatformType                  string `json:"platformType,omitempty"`
	SourceInterfaceName           string `json:"sourceInterfaceName,omitempty"`
	SourceVrfName                 string `json:"sourceVrfName,omitempty"`
	ConfigSyncStatus              string `json:"configSyncStatus,omitempty"`
	DiscoveredSystemMode          string `json:"discoveredSystemMode,omitempty"`
	IntendedSystemMode            string `json:"intendedSystemMode,omitempty"`
	SystemMode                    string `json:"systemMode,omitempty"`
	DomainName                    string `json:"domainName,omitempty"`
	HypershieldConnectivityStatus string `json:"hypershieldConnectivityStatus,omitempty"`
	HypershieldTenant             string `json:"hypershieldTenant,omitempty"`
	RemoteCredentialStore         string `json:"remoteCredentialStore,omitempty"`
	ScalableUnit                  string `json:"scalableUnit,omitempty"`
	SmartSwitch                   *bool  `json:"smartSwitch,omitempty"`
	Usage                         string `json:"usage,omitempty"`
	Vendor                        string `json:"vendor,omitempty"`
	Username                      string `json:"username,omitempty"`
}

func (v *InventorySwitchModel) SetModelData(jsonData *NDFCInventorySwitchModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if len(jsonData.Switches) == 0 {
		log.Printf("v.Switches is empty")
		v.Switches = types.ListNull(SwitchesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.Switches contains %d elements", len(jsonData.Switches))
		listData := make([]SwitchesValue, 0)
		for _, item := range jsonData.Switches {
			data := new(SwitchesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in SwitchesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.Switches, err = types.ListValueFrom(context.Background(), SwitchesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []SwitchesValue to  List")
			return err
		}
	}

	return err
}

func (v *SwitchesValue) SetValue(jsonData *NDFCSwitchesValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.SerialNumber != "" {
		v.SerialNumber = types.StringValue(jsonData.SerialNumber)
	} else {
		v.SerialNumber = types.StringNull()
	}

	if jsonData.IpAddress != "" {
		v.IpAddress = types.StringValue(jsonData.IpAddress)
	} else {
		v.IpAddress = types.StringNull()
	}

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
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

	if jsonData.AdditionalData.Status != "" {
		v.Status = types.StringValue(jsonData.AdditionalData.Status)
	} else {
		v.Status = types.StringNull()
	}

	if jsonData.AdditionalData.PlatformType != "" {
		v.PlatformType = types.StringValue(jsonData.AdditionalData.PlatformType)
	} else {
		v.PlatformType = types.StringNull()
	}

	if jsonData.PreserveConfig != nil {
		v.PreserveConfig = types.BoolValue(*jsonData.PreserveConfig)

	} else {
		v.PreserveConfig = types.BoolNull()
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

	if jsonData.AdditionalData.SourceInterfaceName != "" {
		v.SourceInterfaceName = types.StringValue(jsonData.AdditionalData.SourceInterfaceName)
	} else {
		v.SourceInterfaceName = types.StringNull()
	}

	if jsonData.AdditionalData.SourceVrfName != "" {
		v.SourceVrfName = types.StringValue(jsonData.AdditionalData.SourceVrfName)
	} else {
		v.SourceVrfName = types.StringNull()
	}

	if jsonData.InInventory != nil {
		v.InInventory = types.BoolValue(*jsonData.InInventory)

	} else {
		v.InInventory = types.BoolNull()
	}

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	if jsonData.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.SystemUptime != "" {
		v.SystemUptime = types.StringValue(jsonData.SystemUptime)
	} else {
		v.SystemUptime = types.StringNull()
	}

	if jsonData.AdvisoryLevel != "" {
		v.AdvisoryLevel = types.StringValue(jsonData.AdvisoryLevel)
	} else {
		v.AdvisoryLevel = types.StringNull()
	}

	if jsonData.AlertSuspend != "" {
		v.AlertSuspend = types.StringValue(jsonData.AlertSuspend)
	} else {
		v.AlertSuspend = types.StringNull()
	}

	if jsonData.AnomalyLevel != "" {
		v.AnomalyLevel = types.StringValue(jsonData.AnomalyLevel)
	} else {
		v.AnomalyLevel = types.StringNull()
	}

	if jsonData.VpcConfigured != nil {
		v.VpcConfigured = types.BoolValue(*jsonData.VpcConfigured)

	} else {
		v.VpcConfigured = types.BoolNull()
	}

	if jsonData.AdditionalData.ConfigSyncStatus != "" {
		v.ConfigSyncStatus = types.StringValue(jsonData.AdditionalData.ConfigSyncStatus)
	} else {
		v.ConfigSyncStatus = types.StringNull()
	}

	if jsonData.AdditionalData.DiscoveredSystemMode != "" {
		v.DiscoveredSystemMode = types.StringValue(jsonData.AdditionalData.DiscoveredSystemMode)
	} else {
		v.DiscoveredSystemMode = types.StringNull()
	}

	if jsonData.AdditionalData.IntendedSystemMode != "" {
		v.IntendedSystemMode = types.StringValue(jsonData.AdditionalData.IntendedSystemMode)
	} else {
		v.IntendedSystemMode = types.StringNull()
	}

	if jsonData.AdditionalData.SystemMode != "" {
		v.SystemMode = types.StringValue(jsonData.AdditionalData.SystemMode)
	} else {
		v.SystemMode = types.StringNull()
	}

	if jsonData.AdditionalData.DomainName != "" {
		v.DomainName = types.StringValue(jsonData.AdditionalData.DomainName)
	} else {
		v.DomainName = types.StringNull()
	}

	if jsonData.AdditionalData.HypershieldConnectivityStatus != "" {
		v.HypershieldConnectivityStatus = types.StringValue(jsonData.AdditionalData.HypershieldConnectivityStatus)
	} else {
		v.HypershieldConnectivityStatus = types.StringNull()
	}

	if jsonData.AdditionalData.HypershieldTenant != "" {
		v.HypershieldTenant = types.StringValue(jsonData.AdditionalData.HypershieldTenant)
	} else {
		v.HypershieldTenant = types.StringNull()
	}

	if jsonData.AdditionalData.RemoteCredentialStore != "" {
		v.RemoteCredentialStore = types.StringValue(jsonData.AdditionalData.RemoteCredentialStore)
	} else {
		v.RemoteCredentialStore = types.StringNull()
	}

	if jsonData.AdditionalData.ScalableUnit != "" {
		v.ScalableUnit = types.StringValue(jsonData.AdditionalData.ScalableUnit)
	} else {
		v.ScalableUnit = types.StringNull()
	}

	if jsonData.AdditionalData.SmartSwitch != nil {
		v.SmartSwitch = types.BoolValue(*jsonData.AdditionalData.SmartSwitch)
	} else {
		v.SmartSwitch = types.BoolNull()
	}

	if jsonData.AdditionalData.Usage != "" {
		v.Usage = types.StringValue(jsonData.AdditionalData.Usage)
	} else {
		v.Usage = types.StringNull()
	}

	if jsonData.AdditionalData.Vendor != "" {
		v.Vendor = types.StringValue(jsonData.AdditionalData.Vendor)
	} else {
		v.Vendor = types.StringNull()
	}

	if jsonData.AdditionalData.Username != "" {
		v.Username = types.StringValue(jsonData.AdditionalData.Username)
	} else {
		v.Username = types.StringNull()
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

	if !v.SerialNumber.IsNull() && !v.SerialNumber.IsUnknown() {
		data.SerialNumber = v.SerialNumber.ValueString()
	} else {
		data.SerialNumber = ""
	}

	return data
}
