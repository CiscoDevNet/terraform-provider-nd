// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_inventory_switch

type NDFCInventorySwitchModel struct {
	FabricName   string             `json:"fabricName,omitempty"`
	SerialNumber string             `json:"serialNumber,omitempty"`
	ManagementIp string             `json:"managementIp,omitempty"`
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
}
