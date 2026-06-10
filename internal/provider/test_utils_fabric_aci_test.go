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
	"terraform-provider-nd/internal/manage/resource_fabric_aci"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func FabricAciModelHelperStateCheck(RscName string, c resource_fabric_aci.NDFCFabricAciModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Spec.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Spec.Hostname))
	}
	if c.Spec.Credentials.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Spec.Credentials.Username))
	}
	if c.Spec.Credentials.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Spec.Credentials.Password))
	}
	if c.Spec.Credentials.LoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("login_domain").String(), c.Spec.Credentials.LoginDomain))
	}
	if c.Spec.Aci.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.Spec.Aci.FabricName))
	}
	if c.Spec.Aci.LicenseTier != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("license_tier").String(), c.Spec.Aci.LicenseTier))
	}
	if c.Spec.Aci.Telemetry.TelemetryStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_status").String(), c.Spec.Aci.Telemetry.TelemetryStatus))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_status").String(), "disabled"))
	}
	if c.Spec.Aci.Telemetry.TelemetryNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_network").String(), c.Spec.Aci.Telemetry.TelemetryNetwork))
	}
	if c.Spec.Aci.Telemetry.TelemetryEpg != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_epg").String(), c.Spec.Aci.Telemetry.TelemetryEpg))
	}
	if c.Spec.Aci.Telemetry.TelemetryStreamingProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_streaming_protocol").String(), c.Spec.Aci.Telemetry.TelemetryStreamingProtocol))
	}
	if c.Spec.Aci.Orchestration.OrchestrationStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration_status").String(), c.Spec.Aci.Orchestration.OrchestrationStatus))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration_status").String(), "disabled"))
	}
	if c.Spec.Aci.SecurityDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), c.Spec.Aci.SecurityDomain))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), "all"))
	}
	if c.Spec.Aci.VerifyCa != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("verify_ca").String(), strconv.FormatBool(*c.Spec.Aci.VerifyCa)))
	}
	return ret
}
