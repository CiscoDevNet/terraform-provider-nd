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
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// switchDetailStateChecks returns TestCheckFuncs for the switch_detail nested block.
func switchDetailStateChecks(rscName string, sd resource_inventory_switch.NDFCSwitchDetailValue) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{}
	if sd.IpAddress != "" {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "switch_detail.ip_address", sd.IpAddress))
	}
	if sd.SerialNumber != "" {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "switch_detail.serial_number", sd.SerialNumber))
	}
	if sd.SwitchRole != "" {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "switch_detail.switch_role", sd.SwitchRole))
	}
	if sd.Hostname != "" {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "switch_detail.hostname", sd.Hostname))
	}
	if sd.GatewayIpMask != "" {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "switch_detail.gateway_ip_mask", sd.GatewayIpMask))
	}
	// Computed fields should be non-empty after apply
	checks = append(checks, resource.TestCheckResourceAttrSet(rscName, "switch_detail.model"))
	checks = append(checks, resource.TestCheckResourceAttrSet(rscName, "switch_detail.software_version"))
	return checks
}

// TestAccInventorySwitchMultiResource tests creating two separate
// nd_inventory_switch resources on the same fabric (one switch each).
func TestAccInventorySwitchMultiResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 2 {
		t.Skip("Need at least 2 discovery switches in testbed config for multi-resource test")
	}

	x := &map[string]string{
		"RscType":   "nd_inventory_switch",
		"RscName":   "fabric_test,switch_1,switch_2",
		"User":      cfg.ND.User,
		"Password":  cfg.ND.Password,
		"Host":      cfg.ND.URL,
		"Insecure":  cfg.ND.Insecure,
		"DependsOn": "nd_fabric_vxlan.fabric_test;nd_inventory_switch.switch_1",
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)
	switchRsc1 := new(resource_inventory_switch.NDFCInventorySwitchModel)
	switchRsc2 := new(resource_inventory_switch.NDFCInventorySwitchModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + both switches in one apply
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						cfg.ND.Fabric, "55000", "vxlanIbgp", nil,
					)

					helper.GenerateInventorySwitchFromConfig(&switchRsc1,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						switches[0],
					)

					helper.GenerateInventorySwitchFromConfig(&switchRsc2,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						switches[1],
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc1, switchRsc2}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							InventorySwitchModelHelperStateCheck(
								"nd_inventory_switch.switch_1",
								*switchRsc1,
								path.Empty(),
							),
							switchDetailStateChecks("nd_inventory_switch.switch_1", switchRsc1.SwitchDetail)...,
						),
						append(
							InventorySwitchModelHelperStateCheck(
								"nd_inventory_switch.switch_2",
								*switchRsc2,
								path.Empty(),
							),
							switchDetailStateChecks("nd_inventory_switch.switch_2", switchRsc2.SwitchDetail)...,
						)...,
					)...,
				),
			},
		},
	})
}

// TestAccInventorySwitchBootstrapMultiResource tests creating two separate
// nd_inventory_switch resources in bootstrap mode on the same fabric.
func TestAccInventorySwitchBootstrapMultiResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("bootstrap")
	if len(switches) < 2 {
		t.Skip("Need at least 2 bootstrap switches in testbed config for multi-resource test")
	}

	x := &map[string]string{
		"RscType":   "nd_inventory_switch",
		"RscName":   "fabric_test,switch_1,switch_2",
		"User":      cfg.ND.User,
		"Password":  cfg.ND.Password,
		"Host":      cfg.ND.URL,
		"Insecure":  cfg.ND.Insecure,
		"DependsOn": "nd_fabric_vxlan.fabric_test;nd_inventory_switch.switch_1",
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)
	switchRsc1 := new(resource_inventory_switch.NDFCInventorySwitchModel)
	switchRsc2 := new(resource_inventory_switch.NDFCInventorySwitchModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + both bootstrap switches in one apply
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						cfg.ND.Fabric, "55000", "vxlanIbgp", nil,
					)

					helper.GenerateBootstrapSwitchObject(&switchRsc1,
						invCfg.Fabric,
						switches[0].Serial, switches[0].IP, switches[0].Hostname,
						switches[0].Role, switches[0].GatewayIpMask, switches[0].PoapPassword,
					)

					helper.GenerateBootstrapSwitchObject(&switchRsc2,
						invCfg.Fabric,
						switches[1].Serial, switches[1].IP, switches[1].Hostname,
						switches[1].Role, switches[1].GatewayIpMask, switches[1].PoapPassword,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc1, switchRsc2}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							InventorySwitchModelHelperStateCheck(
								"nd_inventory_switch.switch_1",
								*switchRsc1,
								path.Empty(),
							),
							switchDetailStateChecks("nd_inventory_switch.switch_1", switchRsc1.SwitchDetail)...,
						),
						append(
							InventorySwitchModelHelperStateCheck(
								"nd_inventory_switch.switch_2",
								*switchRsc2,
								path.Empty(),
							),
							switchDetailStateChecks("nd_inventory_switch.switch_2", switchRsc2.SwitchDetail)...,
						)...,
					)...,
				),
			},
		},
	})
}

// TestAccInventorySwitchUpdateCredentials tests the update flow for
// changing discovery credentials (username, password, snmp) on a discovery-mode switch.
func TestAccInventorySwitchUpdateCredentials(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

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

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + discover switch with initial credentials
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricVxlanObject(&fabricRsc,
						cfg.ND.Fabric, "55000", "vxlanIbgp", nil,
					)

					helper.GenerateInventorySwitchFromConfig(&switchRsc,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						switches[0],
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
			// Step 2: Update SNMP auth protocol to SHA
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyDiscoveryCredentials(&switchRsc,
						invCfg.User, invCfg.Password, "sha",
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
			// Step 3: Revert SNMP auth protocol back to MD5
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyDiscoveryCredentials(&switchRsc,
						invCfg.User, invCfg.Password, "md5",
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
		},
	})
}

func TestAccInventorySwitchResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

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

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + discover switch
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

					helper.GenerateInventorySwitchFromConfig(&switchRsc,
						invCfg.Fabric, invCfg.User, invCfg.Password,
						switches[0],
					)

					(*x)["RscName"] = "fabric_test,switch_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
			// Step 2: Update switch role to spine
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifySwitchRole(&switchRsc, "spine")

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
			// Step 3: Revert switch role back to leaf
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifySwitchRole(&switchRsc, "leaf")

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InventorySwitchModelHelperStateCheck(
							"nd_inventory_switch.switch_test",
							*switchRsc,
							path.Empty(),
						),
						switchDetailStateChecks("nd_inventory_switch.switch_test", switchRsc.SwitchDetail)...,
					)...,
				),
			},
		},
	})
}
