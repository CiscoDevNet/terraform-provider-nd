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

	"terraform-provider-nd/internal/manage/resource_fabric_vxlan"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFabricVxlanResourceCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":  "nd_fabric_vxlan",
		"RscName":  "fabric_test",
		"User":     helper.GetConfig("global").ND.User,
		"Password": helper.GetConfig("global").ND.Password,
		"Host":     helper.GetConfig("global").ND.URL,
		"Insecure": helper.GetConfig("global").ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_vxlan.NDFCFabricVxlanModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create a basic VXLAN fabric
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						helper.GetConfig("global").ND.Fabric,
						"55000",
						"vxlanIbgp",
						nil,
					)

					(*x)["RscName"] = "fabric_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanModelHelperStateCheck(
						"nd_fabric_vxlan.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Modify fabric parameters (MTU, overlay mode)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyFabricVxlanObject(&fabricRsc, map[string]interface{}{
						"fabric_mtu":            9000,
						"l2_host_interface_mtu": 9000,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanModelHelperStateCheck(
						"nd_fabric_vxlan.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 3: Modify more parameters (VPC, VNI ranges)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyFabricVxlanObject(&fabricRsc, map[string]interface{}{
						"vpc_peer_link_vlan": "3601",
						"l2_vni_range":       "30000-48000",
						"l3_vni_range":       "50000-58000",
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanModelHelperStateCheck(
						"nd_fabric_vxlan.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
