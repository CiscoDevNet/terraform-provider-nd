// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_vxlan_ibgp

import (
	"context"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

// GetFabricName returns the fabric name from the Terraform model.
func (v *FabricVxlanIbgpModel) GetFabricName() string {
	return v.FabricName.ValueString()
}

// PreMarshal injects the fabric type before the payload is sent to the API.
// FabricType is tf_hide in the schema, so it must be set here.
func (v *FabricVxlanIbgpModel) PreMarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	data.Management.FabricType = "vxlanIbgp"
}
