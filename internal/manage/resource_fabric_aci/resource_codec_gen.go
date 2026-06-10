// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_aci

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCFabricAciModel struct {
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type NDFCSpecValue struct {
	ClusterType string                   `json:"clusterType,omitempty"`
	Hostname    string                   `json:"onboardUrl,omitempty"`
	Location    NDFCSpecLocationValue    `json:"location,omitempty"`
	Credentials NDFCSpecCredentialsValue `json:"credentials,omitempty"`
	Aci         NDFCSpecAciValue         `json:"aci,omitempty"`
}

type NDFCSpecLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCSpecCredentialsValue struct {
	Username    string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	LoginDomain string `json:"loginDomain,omitempty"`
}

type NDFCSpecAciValue struct {
	FabricName     string                    `json:"name,omitempty"`
	LicenseTier    string                    `json:"licenseTier,omitempty"`
	SecurityDomain string                    `json:"securityDomain,omitempty"`
	VerifyCa       *bool                     `json:"verifyCA,omitempty"`
	Telemetry      NDFCAciTelemetryValue     `json:"telemetry,omitempty"`
	Orchestration  NDFCAciOrchestrationValue `json:"orchestration,omitempty"`
}

type NDFCAciTelemetryValue struct {
	TelemetryStatus            string `json:"status,omitempty"`
	TelemetryNetwork           string `json:"network,omitempty"`
	TelemetryEpg               string `json:"epg,omitempty"`
	TelemetryStreamingProtocol string `json:"streamingProtocol,omitempty"`
}

type NDFCAciOrchestrationValue struct {
	OrchestrationStatus string `json:"status,omitempty"`
}

func (v *FabricAciModel) SetModelData(jsonData *NDFCFabricAciModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Spec.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Spec.Hostname)

	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.Spec.Location.Latitude != nil {
		v.Latitude = types.Float64Value(float64(*jsonData.Spec.Location.Latitude))

	} else {
		v.Latitude = types.Float64Null()
	}

	if jsonData.Spec.Location.Longitude != nil {
		v.Longitude = types.Float64Value(float64(*jsonData.Spec.Location.Longitude))

	} else {
		v.Longitude = types.Float64Null()
	}

	if jsonData.Spec.Credentials.Username != "" {
		v.Username = types.StringValue(jsonData.Spec.Credentials.Username)

	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Spec.Credentials.Password != "" {
		v.Password = types.StringValue(jsonData.Spec.Credentials.Password)

	} else {
		v.Password = types.StringNull()
	}

	if jsonData.Spec.Credentials.LoginDomain != "" {
		v.LoginDomain = types.StringValue(jsonData.Spec.Credentials.LoginDomain)

	} else {
		v.LoginDomain = types.StringNull()
	}

	if jsonData.Spec.Aci.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.Spec.Aci.FabricName)

	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Spec.Aci.LicenseTier != "" {
		v.LicenseTier = types.StringValue(jsonData.Spec.Aci.LicenseTier)

	} else {
		v.LicenseTier = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryStatus != "" {
		v.TelemetryStatus = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryStatus)

	} else {
		v.TelemetryStatus = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryNetwork != "" {
		v.TelemetryNetwork = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryNetwork)

	} else {
		v.TelemetryNetwork = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryEpg != "" {
		v.TelemetryEpg = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryEpg)

	} else {
		v.TelemetryEpg = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryStreamingProtocol != "" {
		v.TelemetryStreamingProtocol = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryStreamingProtocol)

	} else {
		v.TelemetryStreamingProtocol = types.StringNull()
	}

	if jsonData.Spec.Aci.Orchestration.OrchestrationStatus != "" {
		v.OrchestrationStatus = types.StringValue(jsonData.Spec.Aci.Orchestration.OrchestrationStatus)

	} else {
		v.OrchestrationStatus = types.StringNull()
	}

	if jsonData.Spec.Aci.SecurityDomain != "" {
		v.SecurityDomain = types.StringValue(jsonData.Spec.Aci.SecurityDomain)

	} else {
		v.SecurityDomain = types.StringNull()
	}

	if jsonData.Spec.Aci.VerifyCa != nil {
		v.VerifyCa = types.BoolValue(*jsonData.Spec.Aci.VerifyCa)

	} else {
		v.VerifyCa = types.BoolNull()
	}

	return err
}

func (v FabricAciModel) GetModelData() *NDFCFabricAciModel {
	var data = new(NDFCFabricAciModel)

	//MARSHAL_BODY

	if !v.Hostname.IsNull() && !v.Hostname.IsUnknown() {
		data.Spec.Hostname = v.Hostname.ValueString()
	} else {
		data.Spec.Hostname = ""
	}

	if !v.Latitude.IsNull() && !v.Latitude.IsUnknown() {
		data.Spec.Location.Latitude = new(float64)
		*data.Spec.Location.Latitude = v.Latitude.ValueFloat64()
	} else {
		data.Spec.Location.Latitude = nil
	}

	if !v.Longitude.IsNull() && !v.Longitude.IsUnknown() {
		data.Spec.Location.Longitude = new(float64)
		*data.Spec.Location.Longitude = v.Longitude.ValueFloat64()
	} else {
		data.Spec.Location.Longitude = nil
	}

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Spec.Credentials.Username = v.Username.ValueString()
	} else {
		data.Spec.Credentials.Username = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Spec.Credentials.Password = v.Password.ValueString()
	} else {
		data.Spec.Credentials.Password = ""
	}

	if !v.LoginDomain.IsNull() && !v.LoginDomain.IsUnknown() {
		data.Spec.Credentials.LoginDomain = v.LoginDomain.ValueString()
	} else {
		data.Spec.Credentials.LoginDomain = ""
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.Spec.Aci.FabricName = v.FabricName.ValueString()
	} else {
		data.Spec.Aci.FabricName = ""
	}

	if !v.LicenseTier.IsNull() && !v.LicenseTier.IsUnknown() {
		data.Spec.Aci.LicenseTier = v.LicenseTier.ValueString()
	} else {
		data.Spec.Aci.LicenseTier = ""
	}

	if !v.TelemetryStatus.IsNull() && !v.TelemetryStatus.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryStatus = v.TelemetryStatus.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryStatus = ""
	}

	if !v.TelemetryNetwork.IsNull() && !v.TelemetryNetwork.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryNetwork = v.TelemetryNetwork.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryNetwork = ""
	}

	if !v.TelemetryEpg.IsNull() && !v.TelemetryEpg.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryEpg = v.TelemetryEpg.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryEpg = ""
	}

	if !v.TelemetryStreamingProtocol.IsNull() && !v.TelemetryStreamingProtocol.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryStreamingProtocol = v.TelemetryStreamingProtocol.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryStreamingProtocol = ""
	}

	if !v.OrchestrationStatus.IsNull() && !v.OrchestrationStatus.IsUnknown() {
		data.Spec.Aci.Orchestration.OrchestrationStatus = v.OrchestrationStatus.ValueString()
	} else {
		data.Spec.Aci.Orchestration.OrchestrationStatus = ""
	}

	if !v.SecurityDomain.IsNull() && !v.SecurityDomain.IsUnknown() {
		data.Spec.Aci.SecurityDomain = v.SecurityDomain.ValueString()
	} else {
		data.Spec.Aci.SecurityDomain = ""
	}

	if !v.VerifyCa.IsNull() && !v.VerifyCa.IsUnknown() {
		data.Spec.Aci.VerifyCa = new(bool)
		*data.Spec.Aci.VerifyCa = v.VerifyCa.ValueBool()
	} else {
		data.Spec.Aci.VerifyCa = nil
	}

	return data
}
