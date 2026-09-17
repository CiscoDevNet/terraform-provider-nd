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

	"terraform-provider-nd/internal/manage/resource_config_deploy"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccConfigDeploySwitchScoped tests config-save + switch-scoped deploy
// after inventory discovery, then updates to fabric-wide deploy.
func TestAccConfigDeploySwitchScoped(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

	x := &map[string]string{
		"RscType":   "nd_config_deploy",
		"RscName":   "fabric_test,switch_1,deploy_test",
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
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
	deployRsc := new(resource_config_deploy.NDFCConfigDeployModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + switch + switch-scoped deploy
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
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

					helper.GenerateConfigDeployObject(&deployRsc,
						invCfg.Fabric, true, true,
						[]string{switches[0].Serial},
						nil,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						ConfigDeployModelHelperStateCheck(
							"nd_config_deploy.deploy_test",
							*deployRsc,
							path.Empty(),
						),
						resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "fabric_name", invCfg.Fabric),
						resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "true"),
						resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "config_save", "true"),
						resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
						resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "switch_ids.#", "1"),
					)...,
				),
			},
			// Step 2: Update to fabric-wide deploy (remove switch_ids)
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyConfigDeployObject(&deployRsc, map[string]interface{}{
						"switch_ids": []string{},
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "fabric_name", invCfg.Fabric),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "true"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "config_save", "true"),
				),
			},
		},
	})
}

// TestAccConfigDeployAlwaysDeploy tests the always_deploy mechanism.
// Step 1: Create with always_deploy=false.
// Step 2: Set always_deploy=true — should trigger an update.
// Step 3: Re-apply same config — always_deploy should force another update.
func TestAccConfigDeployAlwaysDeploy(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

	x := &map[string]string{
		"RscType":   "nd_config_deploy",
		"RscName":   "fabric_test,switch_1,deploy_test",
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
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
	deployRsc := new(resource_config_deploy.NDFCConfigDeployModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with always_deploy=false
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
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

					helper.GenerateConfigDeployObject(&deployRsc,
						invCfg.Fabric, true, true, nil, nil,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "always_deploy", "false"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "true"),
					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
				),
			},
			// Step 2: Enable always_deploy
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyConfigDeployObject(&deployRsc, map[string]interface{}{
						"always_deploy": true,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "always_deploy", "true"),
					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
				),
			},
			// Step 3: Re-apply same config — always_deploy forces update
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "always_deploy", "true"),
					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
				),
			},
		},
	})
}

// TestAccConfigDeployConfigSaveOnly tests config-save without deploy.
func TestAccConfigDeployConfigSaveOnly(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

	x := &map[string]string{
		"RscType":   "nd_config_deploy",
		"RscName":   "fabric_test,switch_1,deploy_test",
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
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
	deployRsc := new(resource_config_deploy.NDFCConfigDeployModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create fabric + switch + config-save only
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
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

					helper.GenerateConfigDeployObject(&deployRsc,
						invCfg.Fabric, true, false, nil, nil,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "config_save", "true"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "false"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "status", "Config save is completed"),
				),
			},
			// Step 2: Switch to deploy only
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyConfigDeployObject(&deployRsc, map[string]interface{}{
						"config_save": false,
						"deploy":      true,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "config_save", "false"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "true"),
					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
				),
			},
		},
	})
}

// TestAccConfigDeployBehaviourOnDestroy tests deploy execution on resource destruction.
func TestAccConfigDeployBehaviourOnDestroy(t *testing.T) {
	cfg := helper.GetConfig("global")
	invCfg := cfg.ND.Inventory

	switches := invCfg.GetSwitchesByMode("discovery")
	if len(switches) < 1 {
		t.Skip("Need at least 1 discovery switch in testbed config")
	}

	onDestroyBehaviour := &resource_config_deploy.NDFCBehaviourValue{
		DeployOnCreate:  helper.BoolPtr(true),
		DeployOnUpdate:  helper.BoolPtr(true),
		DeployOnDestroy: helper.BoolPtr(true),
	}

	x := &map[string]string{
		"RscType":   "nd_config_deploy",
		"RscName":   "fabric_test,switch_1,deploy_test",
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
	switchRsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
	deployRsc := new(resource_config_deploy.NDFCConfigDeployModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with deploy_on_destroy = true
			{
				PreConfig: func() {
					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
					helper.LogStep(t, *stepCount+1, name, "")
				},
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

					helper.GenerateConfigDeployObject(&deployRsc,
						invCfg.Fabric, true, true, nil, onDestroyBehaviour,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.VxlanResource(fabricRsc), switchRsc, deployRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "behaviour.deploy_on_create", "true"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "behaviour.deploy_on_update", "true"),
					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "behaviour.deploy_on_destroy", "true"),
					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
				),
			},
			// The resource.Test framework will automatically destroy all resources
			// after the last step. With deploy_on_destroy=true, the deploy should execute
			// during destruction without errors.
		},
	})
}

// TODO: Enable once IP-to-serial resolution is implemented in the CRUD layer.
// TestAccConfigDeployMultiSwitchByIP tests config-save + deploy using switch IP
// addresses in switch_ids, with 2 inventory switches. Verifies that both
// switches are deployed successfully.
// func TestAccConfigDeployMultiSwitchByIP(t *testing.T) {
// 	cfg := helper.GetConfig("global")
// 	invCfg := cfg.ND.Inventory
//
// 	switches := invCfg.GetSwitchesByMode("discovery")
// 	if len(switches) < 2 {
// 		t.Skip("Need at least 2 discovery switches in testbed config")
// 	}
//
// 	x := &map[string]string{
// 		"RscType":  "nd_config_deploy",
// 		"RscName":  "fabric_test,switch_1,switch_2,deploy_test",
// 		"User":     cfg.ND.User,
// 		"Password": cfg.ND.Password,
// 		"Host":     cfg.ND.URL,
// 		"Insecure": cfg.ND.Insecure,
// 		"DependsOn": "nd_fabric_vxlan.fabric_test;" +
// 			"nd_fabric_vxlan.fabric_test;" +
// 			"nd_inventory_switch.switch_1, nd_inventory_switch.switch_2",
// 	}
//
// 	tfConfig := new(string)
// 	stepCount := new(int)
// 	*stepCount = 0
//
// 	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)
// 	switch1Rsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
// 	switch2Rsc := new(resource_inventory_switch.NDFCInventorySwitchModel)
// 	deployRsc := new(resource_config_deploy.NDFCConfigDeployModel)
//
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t, "global") },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Step 1: Create fabric + 2 switches + deploy scoped to both by IP
// 			{
// 				PreConfig: func() {
// 					name := fmt.Sprintf("%s_%d", t.Name(), *stepCount+1)
// 					helper.LogStep(t, *stepCount+1, name, "")
// 				},
// 				Config: func() string {
// 					*stepCount++
// 					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)
//
// 					helper.GenerateFabricVxlanObject(&fabricRsc,
// 						cfg.ND.Fabric, "55000", "vxlanIbgp", nil,
// 					)
//
// 					helper.GenerateInventorySwitchFromConfig(&switch1Rsc,
// 						invCfg.Fabric, invCfg.User, invCfg.Password,
// 						switches[0],
// 					)
//
// 					helper.GenerateInventorySwitchFromConfig(&switch2Rsc,
// 						invCfg.Fabric, invCfg.User, invCfg.Password,
// 						switches[1],
// 					)
//
// 					helper.GenerateConfigDeployObject(&deployRsc,
// 						invCfg.Fabric, true, true,
// 						[]string{switches[0].IP, switches[1].IP},
// 						nil,
// 					)
//
// 					helper.GetTFConfigWithSingleResource(tName, *x,
// 						[]interface{}{helper.VxlanResource(fabricRsc), switch1Rsc, switch2Rsc, deployRsc}, &tfConfig)
//
// 					return *tfConfig
// 				}(),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "fabric_name", invCfg.Fabric),
// 					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "deploy", "true"),
// 					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "config_save", "true"),
// 					resource.TestCheckResourceAttr("nd_config_deploy.deploy_test", "switch_ids.#", "2"),
// 					resource.TestCheckResourceAttrSet("nd_config_deploy.deploy_test", "status"),
// 				),
// 			},
// 		},
// 	})
// }
