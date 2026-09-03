// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_aci

import (
	"terraform-provider-nd/internal/common/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fabricAciManageUpdatePayload is the payload shape for
// PUT /manage/fabrics/{fabricName}. The generated NDFCFabricAciModel
// is kept for the infra /clusters create/read/delete payload.
type fabricAciManageUpdatePayload struct {
	Category       string `json:"category,omitempty"`
	FabricName     string `json:"name,omitempty"`
	LicenseTier    string `json:"licenseTier,omitempty"`
	SecurityDomain string `json:"securityDomain,omitempty"`
	// TelemetryCollection is derived from whether the Terraform telemetry object
	// is configured, rather than from its computed status attribute.
	TelemetryCollection        *bool                      `json:"telemetryCollection,omitempty"`
	TelemetryCollectionType    string                     `json:"telemetryCollectionType,omitempty"`
	TelemetryStreamingProtocol string                     `json:"telemetryStreamingProtocol,omitempty"`
	Location                   *NDFCSpecLocationValue     `json:"location,omitempty"`
	Management                 *fabricAciManageManagement `json:"management,omitempty"`
}

type fabricAciManageManagement struct {
	Type          string `json:"type,omitempty"`
	Orchestration *bool  `json:"orchestration,omitempty"`
	Epg           string `json:"epg,omitempty"`
}

func (v FabricAciModel) manageUpdatePayload() *fabricAciManageUpdatePayload {
	data := &fabricAciManageUpdatePayload{
		Category: "fabric",
		Management: &fabricAciManageManagement{
			Type: "aci",
		},
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	}
	if !v.LicenseTier.IsNull() && !v.LicenseTier.IsUnknown() {
		data.LicenseTier = v.LicenseTier.ValueString()
	}

	if !v.SecurityDomain.IsNull() && !v.SecurityDomain.IsUnknown() {
		data.SecurityDomain = v.SecurityDomain.ValueString()
	}

	telemetryCollection := fabricAciTelemetryEnabled(v.Telemetry)
	data.TelemetryCollection = &telemetryCollection
	if telemetryCollection {
		if !v.Telemetry.Network.IsNull() && !v.Telemetry.Network.IsUnknown() {
			data.TelemetryCollectionType = telemetryNetworkToCollectionType(v.Telemetry.Network.ValueString())
		}
		if !v.Telemetry.StreamingProtocol.IsNull() && !v.Telemetry.StreamingProtocol.IsUnknown() {
			data.TelemetryStreamingProtocol = v.Telemetry.StreamingProtocol.ValueString()
		}
		if !v.Telemetry.Epg.IsNull() && !v.Telemetry.Epg.IsUnknown() {
			data.Management.Epg = v.Telemetry.Epg.ValueString()
		}
	}
	if !v.Latitude.IsNull() && !v.Latitude.IsUnknown() {
		if data.Location == nil {
			data.Location = &NDFCSpecLocationValue{}
		}
		data.Location.Latitude = new(float64)
		*data.Location.Latitude = v.Latitude.ValueFloat64()
	}
	if !v.Longitude.IsNull() && !v.Longitude.IsUnknown() {
		if data.Location == nil {
			data.Location = &NDFCSpecLocationValue{}
		}
		data.Location.Longitude = new(float64)
		*data.Location.Longitude = v.Longitude.ValueFloat64()
	}
	if !v.OrchestrationStatus.IsNull() && !v.OrchestrationStatus.IsUnknown() {
		orchestration := utils.EnabledDisabledToBool(v.OrchestrationStatus.ValueString())
		data.Management.Orchestration = &orchestration
	}
	return data
}

func fabricAciTelemetryEnabled(telemetry TelemetryValue) bool {
	return !telemetry.IsNull() && !telemetry.IsUnknown()
}

func fabricAciTelemetryStatus(telemetry TelemetryValue) string {
	return utils.BoolToEnabledDisabled(fabricAciTelemetryEnabled(telemetry))
}

func (v *FabricAciModel) NormalizeTelemetryNetworkState() {
	if v == nil || v.Telemetry.IsNull() || v.Telemetry.IsUnknown() || v.Telemetry.Network.IsNull() || v.Telemetry.Network.IsUnknown() {
		return
	}

	v.Telemetry.Network = types.StringValue(normalizeTelemetryNetworkState(v.Telemetry.Network.ValueString()))
}

func telemetryNetworkToCollectionType(network string) string {
	switch normalizeTelemetryNetworkState(network) {
	case "inband":
		return "inBand"
	case "outband":
		return "outOfBand"
	}
	return network
}

func telemetryNetworkToCreatePayload(network string) string {
	switch normalizeTelemetryNetworkState(network) {
	case "inband":
		return "inband"
	case "outband":
		return "outOfBand"
	}
	return network
}

func normalizeTelemetryNetworkState(network string) string {
	switch network {
	case "inband", "inBand":
		return "inband"
	case "outband", "outOfBand":
		return "outband"
	}
	return network
}
