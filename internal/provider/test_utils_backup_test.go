// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"strconv"

	"terraform-provider-nd/internal/infra/resource_backup"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// BackupModelHelperStateCheck returns a slice of TestCheckFunc that validates
// an nd_backup resource state matches the values held in the supplied
// NDFCBackupModel. The `encryption_key` attribute is intentionally not
// validated here because it is not returned by the API and is reconstructed
// from prior state by the provider.
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
	ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("destination").String(), c.Destination))
	if c.TelemetryData != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_data").String(), strconv.FormatBool(*c.TelemetryData)))
	}
	return ret
}
