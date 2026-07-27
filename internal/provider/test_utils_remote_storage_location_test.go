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
	if c.StorageLocationType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storage_location_type").String(), c.StorageLocationType))
	}
	if c.ReadWrite != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("read_write").String(), strconv.FormatBool(*c.ReadWrite)))
	}
	if c.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Hostname))
	}
	if c.Port != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port").String(), strconv.Itoa(int(*c.Port))))
	}
	if c.Path != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("path").String(), c.Path))
	}
	if c.AlertThreshold != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("alert_threshold").String(), strconv.Itoa(int(*c.AlertThreshold))))
	}
	if c.Limit != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("limit").String(), c.Limit))
	}
	if c.Authentication.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Authentication.Username))
	}
	if c.Authentication.AuthenticationType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("authentication_type").String(), c.Authentication.AuthenticationType))
	}
	if c.Authentication.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Authentication.Password))
	}
	if c.Authentication.SshKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ssh_key").String(), c.Authentication.SshKey))
	}
	if c.Authentication.Passphrase != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("passphrase").String(), c.Authentication.Passphrase))
	}
	if c.Authentication.IgnoreHostKeyValidation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ignore_host_key_validation").String(), strconv.FormatBool(*c.Authentication.IgnoreHostKeyValidation)))
	}
	if c.AcceptHostKey {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("accept_host_key").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("accept_host_key").String(), "false"))
	}
	if c.HealthState != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("health_state").String(), c.HealthState))
	}
	if c.HealthStateMessage != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("health_state_message").String(), c.HealthStateMessage))
	}
	return ret
}
