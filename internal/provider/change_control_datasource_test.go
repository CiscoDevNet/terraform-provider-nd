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
	"testing"

	"terraform-provider-nd/internal/infra/resource_change_control"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccChangeControlDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	resourceName := "nd_change_control.change_control_datasource_test"
	dataSourceName := "data.nd_change_control.change_control_datasource_test"

	x := &map[string]string{
		"RscType":  "nd_change_control",
		"RscName":  "change_control_datasource_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	changeControl := new(resource_change_control.NDFCChangeControlModel)
	changeControlDataSource := &helper.ChangeControlDataSourceTestData{
		RscName:   "change_control_datasource_test",
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
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Configure change control for datasource reads")

					helper.GenerateChangeControlObject(&changeControl, map[string]interface{}{
						"admin_status":                    true,
						"orchestration":                   true,
						"nd_managed_fabrics":              false,
						"bypass_telemetry_change_control": true,
						"number_of_approvers":             3,
						"allow_self_approval":             false,
						"ticket_name_prefix":              "TF_DS_CC_",
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
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Read the change control singleton through the datasource")

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
					resource.TestCheckResourceAttr(dataSourceName, "ticket_name_prefix", "TF_DS_CC_"),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Match all datasource attributes with the configured change control settings")

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
