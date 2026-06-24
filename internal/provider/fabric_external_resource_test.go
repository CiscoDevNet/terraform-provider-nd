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

	"terraform-provider-nd/internal/manage/resource_fabric_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFabricExternalResourceCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":  "nd_fabric_external",
		"RscName":  "fabric_test",
		"User":     helper.GetConfig("global").ND.User,
		"Password": helper.GetConfig("global").ND.Password,
		"Host":     helper.GetConfig("global").ND.URL,
		"Insecure": helper.GetConfig("global").ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create a basic External Connectivity fabric
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricExternalObject(&fabricRsc,
						helper.GetConfig("global").ND.Fabric+"_external",
						"65001",
						nil,
					)

					(*x)["RscName"] = "fabric_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.ExternalResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricExternalModelHelperStateCheck(
						"nd_fabric_external.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Modify fabric parameters (sub_interface_dot1q_range)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyFabricExternalObject(&fabricRsc, map[string]interface{}{
						"sub_interface_dot1q_range": "2-1000",
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.ExternalResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricExternalModelHelperStateCheck(
						"nd_fabric_external.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
