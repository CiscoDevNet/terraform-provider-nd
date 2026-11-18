// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_fabric_aci

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCFabricAciModel struct {
	Status NDFCStatusValue `json:"status,omitempty"`
	Spec   NDFCSpecValue   `json:"spec,omitempty"`
}

type NDFCStatusValue struct {
	State      string                    `json:"state,omitempty"`
	LastUpdate NDFCStatusLastUpdateValue `json:"lastUpdate,omitempty"`
}

type NDFCStatusLastUpdateValue struct {
	LastUpdateMessage string `json:"message,omitempty"`
}

type NDFCSpecValue struct {
	Hostname string                `json:"onboardUrl,omitempty"`
	Location NDFCSpecLocationValue `json:"location,omitempty"`
	Aci      NDFCAciValue          `json:"aci,omitempty"`
}

type NDFCSpecLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCAciValue struct {
	FabricName     string                    `json:"name,omitempty"`
	LicenseTier    string                    `json:"licenseTier,omitempty"`
	Telemetry      NDFCTelemetryValue        `json:"telemetry,omitempty"`
	SecurityDomain string                    `json:"securityDomain,omitempty"`
	Orchestration  NDFCAciOrchestrationValue `json:"orchestration,omitempty"`
}

type NDFCTelemetryValue struct {
	Status            string `json:"status,omitempty"`
	Network           string `json:"network,omitempty"`
	Epg               string `json:"epg,omitempty"`
	StreamingProtocol string `json:"streamingProtocol,omitempty"`
}

type NDFCAciOrchestrationValue struct {
	OrchestrationStatus string `json:"status,omitempty"`
}

func (v *FabricAciModel) SetModelData(jsonData *NDFCFabricAciModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Spec.Aci.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.Spec.Aci.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

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

	if jsonData.Spec.Aci.LicenseTier != "" {
		v.LicenseTier = types.StringValue(jsonData.Spec.Aci.LicenseTier)
	} else {
		v.LicenseTier = types.StringNull()
	}

	v.Telemetry.SetValue(&jsonData.Spec.Aci.Telemetry)
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

	if jsonData.Status.State != "" {
		v.State = types.StringValue(jsonData.Status.State)
	} else {
		v.State = types.StringNull()
	}

	if jsonData.Status.LastUpdate.LastUpdateMessage != "" {
		v.LastUpdateMessage = types.StringValue(jsonData.Status.LastUpdate.LastUpdateMessage)
	} else {
		v.LastUpdateMessage = types.StringNull()
	}

	return err
}

func (v *TelemetryValue) SetValue(jsonData *NDFCTelemetryValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	valueStateKnown := false
	if jsonData.Status != "" {
		v.Status = types.StringValue(jsonData.Status)
		valueStateKnown = true
	} else {
		v.Status = types.StringNull()
	}

	if jsonData.Network != "" {
		v.Network = types.StringValue(jsonData.Network)
		valueStateKnown = true
	} else {
		v.Network = types.StringNull()
	}

	if jsonData.Epg != "" {
		v.Epg = types.StringValue(jsonData.Epg)
		valueStateKnown = true
	} else {
		v.Epg = types.StringNull()
	}

	if jsonData.StreamingProtocol != "" {
		v.StreamingProtocol = types.StringValue(jsonData.StreamingProtocol)
		valueStateKnown = true
	} else {
		v.StreamingProtocol = types.StringNull()
	}

	if valueStateKnown {
		v.state = attr.ValueStateKnown
	}

	return err
}

func (v FabricAciModel) GetModelData() *NDFCFabricAciModel {
	var data = new(NDFCFabricAciModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.Spec.Aci.FabricName = v.FabricName.ValueString()
	} else {
		data.Spec.Aci.FabricName = ""
	}

	return data
}
