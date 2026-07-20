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

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// RemoteStorageLocationModelHelperStateCheck returns a slice of TestCheckFunc
// that validates an nd_remote_storage_location resource state matches the
// values held in the supplied test data model.
//
// Sensitive fields (`username`, `password`, `ssh_key`, `passphrase`,
// `ignore_host_key_validation`) are not returned by the API but the provider
// preserves them in state from the prior plan/state, so we still verify them
// here. Callers using ImportState should add these to ImportStateVerifyIgnore
// because the imported state will not have any prior plan to preserve from.
func RemoteStorageLocationModelHelperStateCheck(
	rscName string,
	c helper.NDFCRemoteStorageLocationTestData,
	attrPath path.Path,
) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.Description != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("description").String(), c.Description))
	}
	if c.StorageLocationType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("storage_location_type").String(), c.StorageLocationType))
	}
	// `read_write` is only meaningful for NFS storage. The provider requires
	// it to be set for NFS and forces it to null for SCP/SFTP because the
	// API does not honor or echo it for non-NFS storage.
	switch c.StorageLocationType {
	case "nfs":
		if c.ReadWrite == nil {
			panic("test data: read_write must be set for nfs storage_location_type")
		}
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("read_write").String(), strconv.FormatBool(*c.ReadWrite)))
	default:
		ret = append(ret, resource.TestCheckNoResourceAttr(rscName, attrPath.AtName("read_write").String()))
	}
	if c.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("hostname").String(), c.Hostname))
	}
	if c.Port != nil {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("port").String(), strconv.FormatInt(*c.Port, 10)))
	}
	if c.Path != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("path").String(), c.Path))
	}
	if c.AlertThreshold != nil {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("alert_threshold").String(), strconv.FormatInt(*c.AlertThreshold, 10)))
	}
	if c.Limit != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("limit").String(), c.Limit))
	}
	if c.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("username").String(), c.Username))
	}
	if c.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("password").String(), c.Password))
	}
	if c.SshKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("ssh_key").String(), c.SshKey))
	}
	if c.Passphrase != "" {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("passphrase").String(), c.Passphrase))
	}
	if c.IgnoreHostKeyValidation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(rscName, attrPath.AtName("ignore_host_key_validation").String(), strconv.FormatBool(*c.IgnoreHostKeyValidation)))
	}
	return ret
}
