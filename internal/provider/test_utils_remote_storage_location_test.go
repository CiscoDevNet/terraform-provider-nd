// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"terraform-provider-nd/internal/infra/resource_remote_storage_location"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func RemoteStorageLocationModelHelperStateCheck(RscName string, c resource_remote_storage_location.NDFCRemoteStorageLocationModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.Description != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("description").String(), c.Description))
	}
	if c.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Hostname))
	}
	if c.Path != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("path").String(), c.Path))
	}
	if c.HealthState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("health_state").String(), c.HealthState))
	}
	if c.HealthStateMessage != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("health_state_message").String(), c.HealthStateMessage))
	}
	return ret
}
