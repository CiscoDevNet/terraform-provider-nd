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
	"terraform-provider-nd/internal/infra/resource_change_control"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func ChangeControlModelHelperStateCheck(RscName string, c resource_change_control.NDFCChangeControlModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.AdminStatus != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_status").String(), strconv.FormatBool(*c.AdminStatus)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_status").String(), "false"))
	}
	if c.Orchestration != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration").String(), strconv.FormatBool(*c.Orchestration)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration").String(), "false"))
	}
	if c.NumberOfApprovers != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("number_of_approvers").String(), strconv.Itoa(int(*c.NumberOfApprovers))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("number_of_approvers").String(), "1"))
	}
	if c.AllowSelfApproval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_self_approval").String(), strconv.FormatBool(*c.AllowSelfApproval)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_self_approval").String(), "true"))
	}
	if c.NdManagedFabrics != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nd_managed_fabrics").String(), strconv.FormatBool(*c.NdManagedFabrics)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nd_managed_fabrics").String(), "false"))
	}
	if c.BypassTelemetryChangeControl != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bypass_telemetry_change_control").String(), strconv.FormatBool(*c.BypassTelemetryChangeControl)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bypass_telemetry_change_control").String(), "false"))
	}
	if c.TicketNamePrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ticket_name_prefix").String(), c.TicketNamePrefix))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ticket_name_prefix").String(), "TICKET_"))
	}
	return ret
}
