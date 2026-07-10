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
	"strconv"
	"terraform-provider-nd/internal/infra/resource_backup"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func BackupModelHelperStateCheck(RscName string, c resource_backup.NDFCBackupModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.Type != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("type").String(), c.Type))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("type").String(), "configOnly"))
	}
	if c.Destination != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination").String(), c.Destination))
	}
	if c.EncryptionKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("encryption_key").String(), c.EncryptionKey))
	}
	if c.TelemetryData != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_data").String(), strconv.FormatBool(*c.TelemetryData)))
	}
	return ret
}
