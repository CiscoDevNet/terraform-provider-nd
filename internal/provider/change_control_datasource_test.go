// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"terraform-provider-nd/internal/infra/resource_change_control"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInfraChangeControlDataSource(t *testing.T) {
	// Test each target independently because an ND deployment can support one
	// change control target while rejecting the other.
	t.Run("Orchestration", func(t *testing.T) {
		testAccInfraChangeControlDataSourceTarget(
			t,
			"orchestration",
			true,
			false,
			"TF_DS_ORCH_CC_",
			regexp.MustCompile(
				`(?is)(?:Unable\s+to\s+enable\s+change\s+control\s+for\s+Orchestration\.\s+Confirm\s+the\s+Orchestration\s+feature\s+is\s+enabled\.|orchestration.{0,120}(?:not supported|unsupported|not available|not allowed)|(?:not supported|unsupported|not available|not allowed).{0,120}orchestration)`,
			),
		)
	})

	t.Run("NDManagedFabrics", func(t *testing.T) {
		testAccInfraChangeControlDataSourceTarget(
			t,
			"nd_managed_fabrics",
			false,
			true,
			"TF_DS_FABRIC_CC_",
			regexp.MustCompile(
				`(?is)Enabling the change control feature is\s+not allowed if tenant associations exist for some managed LAN fabrics|(?:ND[- ]managed fabrics?.{0,120}(?:not supported|unsupported|not available|not allowed)|(?:not supported|unsupported|not available|not allowed).{0,120}ND[- ]managed fabrics?)`,
			),
		)
	})
}

func testAccInfraChangeControlDataSourceTarget(
	t *testing.T,
	capability string,
	orchestration bool,
	ndManagedFabrics bool,
	ticketNamePrefix string,
	rejection *regexp.Regexp,
) {
	t.Helper()

	cfg := helper.GetConfig("global")
	testName := fmt.Sprintf("change_control_datasource_%s_test", capability)
	resourceName := "nd_change_control." + testName
	dataSourceName := "data.nd_change_control." + testName

	x := &map[string]string{
		"RscType":  "nd_change_control",
		"RscName":  testName,
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	changeControl := new(resource_change_control.NDFCChangeControlModel)
	changeControlDataSource := &helper.ChangeControlDataSourceTestData{
		RscName:   testName,
		DependsOn: resourceName,
	}
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "admin_status", resourceName, "admin_status"),
		resource.TestCheckResourceAttrPair(dataSourceName, "allow_self_approval", resourceName, "allow_self_approval"),
		resource.TestCheckResourceAttrPair(dataSourceName, "bypass_telemetry_change_control", resourceName, "bypass_telemetry_change_control"),
		resource.TestCheckResourceAttrPair(dataSourceName, "nd_managed_fabrics", resourceName, "nd_managed_fabrics"),
		resource.TestCheckResourceAttrPair(dataSourceName, "number_of_approvers", resourceName, "number_of_approvers"),
		resource.TestCheckResourceAttrPair(dataSourceName, "orchestration", resourceName, "orchestration"),
		resource.TestCheckResourceAttrPair(dataSourceName, "ticket_name_prefix", resourceName, "ticket_name_prefix"),
	}

	// A missing-object lookup is not applicable because change control is a
	// built-in GET/PUT singleton with no datasource lookup attribute.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ErrorCheck:               skipUnsupportedChangeControlTargetError(t, capability, rejection),
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Configure change control for %s datasource reads", t.Name(), capability)

					helper.GenerateChangeControlObject(&changeControl, map[string]interface{}{
						"admin_status":                    true,
						"orchestration":                   orchestration,
						"nd_managed_fabrics":              ndManagedFabrics,
						"bypass_telemetry_change_control": true,
						"number_of_approvers":             3,
						"allow_self_approval":             false,
						"ticket_name_prefix":              ticketNamePrefix,
					})
					helper.GetTFConfigWithSingleResource(s1.Name, *x, []interface{}{changeControl}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					ChangeControlModelHelperStateCheck(resourceName, *changeControl, path.Empty())...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Read the %s change control configuration through the datasource", t.Name(), capability)

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{changeControl, changeControlDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "ticket_name_prefix", ticketNamePrefix),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - Match all datasource attributes with the configured %s settings", t.Name(), capability)

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						*x,
						[]interface{}{changeControl, changeControlDataSource},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check:     resource.ComposeTestCheckFunc(matchingChecks...),
			},
		},
	})
}
