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
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccInventorySwitchMultiResource tests splitting switches across two
// separate nd_inventory_switch resources on the same fabric.
// Resource 1: first 2 switches, Resource 2: third switch.
func TestAccInventorySwitchMultiResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	if len(invCfg.Switches) < 3 {
		t.Skip("Need at least 3 switches in testbed config for multi-resource test")
	}

	// Config map: 3 resource names (fabric + 2 switch groups)
	// DependsOn is semicolon-separated: first switch group depends on fabric,
	// second switch group depends on first switch group.
	x := &map[string]string{
		"RscType":   "nd_inventory_switch",
		"RscName":   "fabric_test,switch_group_1,switch_group_2",
		"User":      cfg.ND.User,
		"Password":  cfg.ND.Password,
		"Host":      cfg.ND.URL,
		"Insecure":  cfg.ND.Insecure,
		"DependsOn": "nd_fabric_vxlan.fabric_test;nd_inventory_switch.switch_group_1",
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_vxlan.NDFCFabricVxlanModel)
	switchRsc1 := new(resource_inventory_switch.NDFCInventorySwitchModel)
	switchRsc2 := new(resource_inventory_switch.NDFCInventorySwitchModel)

	// Group 1: first 2 switches
	group1Switches := map[string]string{
		invCfg.Switches[0].Serial: invCfg.Switches[0].IP,
		invCfg.Switches[1].Serial: invCfg.Switches[1].IP,
	}
	group1Roles := map[string]string{
		invCfg.Switches[0].Serial: invCfg.Switches[0].Role,
		invCfg.Switches[1].Serial: invCfg.Switches[1].Role,
	}

	// Group 2: third switch
	group2Switches := map[string]string{
		invCfg.Switches[2].Serial: invCfg.Switches[2].IP,
	}
	group2Roles := map[string]string{
		invCfg.Switches[2].Serial: invCfg.Switches[2].Role,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + both switch groups in one apply
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						cfg.ND.Fabric, "55000", "vxlanIbgp", nil,
					)

					helper.GenerateInventorySwitchObject(&switchRsc1,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						group1Switches, group1Roles, invCfg.Deploy, invCfg.Recalc,
					)

					helper.GenerateInventorySwitchObject(&switchRsc2,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						group2Switches, group2Roles, invCfg.Deploy, invCfg.Recalc,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc, switchRsc1, switchRsc2}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_group_1",
							*switchRsc1,
							path.Empty(),
						),
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_group_2",
							*switchRsc2,
							path.Empty(),
						)...,
					)...,
				),
			},
		},
	})
}

func TestAccInventorySwitchResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	x := &map[string]string{
		"RscType":   "nd_inventory_switch",
		"RscName":   "fabric_test,switch_test",
		"User":      cfg.ND.User,
		"Password":  cfg.ND.Password,
		"Host":      cfg.ND.URL,
		"Insecure":  cfg.ND.Insecure,
		"DependsOn": "nd_fabric_vxlan.fabric_test",
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_vxlan.NDFCFabricVxlanModel)
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)

	// Build initial switch map from testbed config (first 2 switches)
	initialSwitches := map[string]string{}
	initialRoles := map[string]string{}
	for i := 0; i < 2 && i < len(invCfg.Switches); i++ {
		initialSwitches[invCfg.Switches[i].Serial] = invCfg.Switches[i].IP
		initialRoles[invCfg.Switches[i].Serial] = invCfg.Switches[i].Role
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + discover 2 switches
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						cfg.ND.Fabric,
						"55000",
						"vxlanIbgp",
						nil,
					)

					helper.GenerateInventorySwitchObject(&switchRsc,
						invCfg.Fabric,
						invCfg.User,
						invCfg.Password,
						initialSwitches,
						initialRoles,
						false,
						false,
					)

					(*x)["RscName"] = "fabric_test,switch_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc, switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					InventorySwitchModelHelperStateCheck(
						"nd_inventory_switch.switch_test",
						*switchRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Add a third switch
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					if len(invCfg.Switches) >= 3 {
						helper.AddSwitches(&switchRsc,
							map[string]string{invCfg.Switches[2].Serial: invCfg.Switches[2].IP},
							map[string]string{invCfg.Switches[2].Serial: invCfg.Switches[2].Role},
						)
					}

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc, switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					InventorySwitchModelHelperStateCheck(
						"nd_inventory_switch.switch_test",
						*switchRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 3: Enable deploy and recalculate
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyInventorySwitchObject(&switchRsc, map[string]interface{}{
						"deploy":      true,
						"recalculate": true,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc, switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					InventorySwitchModelHelperStateCheck(
						"nd_inventory_switch.switch_test",
						*switchRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 4: Remove the third switch
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					if len(invCfg.Switches) >= 3 {
						helper.DeleteSwitches(&switchRsc, []string{invCfg.Switches[2].Serial})
					}

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{fabricRsc, switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					InventorySwitchModelHelperStateCheck(
						"nd_inventory_switch.switch_test",
						*switchRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
