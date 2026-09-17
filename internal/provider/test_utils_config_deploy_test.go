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
	"terraform-provider-nd/internal/manage/resource_config_deploy"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func ConfigDeployModelHelperStateCheck(RscName string, c resource_config_deploy.NDFCConfigDeployModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	if c.ConfigSave {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("config_save").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("config_save").String(), "false"))
	}
	if c.ForceShowRun {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("force_show_run").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("force_show_run").String(), "false"))
	}
	if c.IncludeAllFabricGroupSwitches {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("include_all_fabric_group_switches").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("include_all_fabric_group_switches").String(), "false"))
	}
	if c.AlwaysDeploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("always_deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("always_deploy").String(), "false"))
	}
	if c.Status != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("status").String(), c.Status))
	}
	if c.TicketId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ticket_id").String(), c.TicketId))
	}
	return ret
}
