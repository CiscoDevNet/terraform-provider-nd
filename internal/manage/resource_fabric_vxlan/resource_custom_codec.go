// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_vxlan

import (
	"context"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

// GetFabricName returns the fabric name from the Terraform model.
func (v *FabricVxlanModel) GetFabricName() string {
	return v.FabricName.ValueString()
}

// GetFabricType returns the fabric type identifier for handler dispatch.
func (v *FabricVxlanModel) GetFabricType() string {
	return "vxlan"
}

// PostUnmarshal preserves flow-collection settings that ND requires on create
// but omits from some fabric GET responses. Without this normalization the
// provider changes configured values to null immediately after apply.
func (v *FabricVxlanModel) PostUnmarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	configured := v.GetModelData().TelemetrySettings.FlowCollection
	returned := &data.TelemetrySettings.FlowCollection

	if returned.TrafficAnalytics == "" {
		returned.TrafficAnalytics = configured.TrafficAnalytics
	}
	if returned.FlowCollectionModes.NetFlow == nil {
		returned.FlowCollectionModes.NetFlow = configured.FlowCollectionModes.NetFlow
	}
	if returned.FlowCollectionModes.SFlow == nil {
		returned.FlowCollectionModes.SFlow = configured.FlowCollectionModes.SFlow
	}
	if returned.FlowCollectionModes.FlowTelemetry == nil {
		returned.FlowCollectionModes.FlowTelemetry = configured.FlowCollectionModes.FlowTelemetry
	}
}
