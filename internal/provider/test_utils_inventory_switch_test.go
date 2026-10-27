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
	"terraform-provider-nd/internal/manage/resource_inventory_switch"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InventorySwitchModelHelperStateCheck(RscName string, c resource_inventory_switch.NDFCInventorySwitchModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.Mode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), c.Mode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), "discovery"))
	}

	if c.PreserveConfig != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), strconv.FormatBool(*c.PreserveConfig)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), "false"))
	}
	if c.WaitForBootstrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wait_for_bootstrap").String(), c.WaitForBootstrap))
	}
	if c.WaitForDiscover != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wait_for_discover").String(), c.WaitForDiscover))
	}
	if c.WaitForReady != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wait_for_ready").String(), c.WaitForReady))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wait_for_ready").String(), "30m"))
	}
	if c.DiscoveryUsername != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_username").String(), c.DiscoveryUsername))
	}
	if c.DiscoveryPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_password").String(), c.DiscoveryPassword))
	}
	if c.BootstrapPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_password").String(), c.BootstrapPassword))
	}
	if c.UseNewCredentials {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_new_credentials").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("use_new_credentials").String(), "false"))
	}
	if c.DiscoveryCredForLan {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_cred_for_lan").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_cred_for_lan").String(), "false"))
	}
	if c.PlatformType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("platform_type").String(), c.PlatformType))
	}
	if c.SnmpV3AuthProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_v3_auth_protocol").String(), c.SnmpV3AuthProtocol))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_v3_auth_protocol").String(), "md5"))
	}
	if c.RemoteCredentialStore != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store").String(), c.RemoteCredentialStore))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store").String(), "local"))
	}
	if c.RemoteCredentialStoreKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store_key").String(), c.RemoteCredentialStoreKey))
	}
	if c.SourceInterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_interface_name").String(), c.SourceInterfaceName))
	}
	if c.SourceVrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_vrf_name").String(), c.SourceVrfName))
	}

	return ret
}
