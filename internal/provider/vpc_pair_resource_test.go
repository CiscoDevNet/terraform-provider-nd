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
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_vpc_pair"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type vpcPairTestContext struct {
	configMap    map[string]string
	fabricName   string
	leafSwitches []string
	spineSwitch  string
}

func loadVpcPairTestContext(t *testing.T) *vpcPairTestContext {
	t.Helper()

	cfg := helper.GetConfig("global")
	ctx := &vpcPairTestContext{
		configMap: map[string]string{
			"RscType":  "nd_vpc_pair",
			"RscName":  "vpc_pair_test",
			"User":     cfg.ND.User,
			"Password": cfg.ND.Password,
			"Host":     cfg.ND.URL,
			"Insecure": cfg.ND.Insecure,
		},
		fabricName: cfg.ND.Fabric,
	}

	for _, sw := range cfg.ND.Inventory.Switches {
		switch {
		case strings.EqualFold(sw.Role, "leaf"):
			ctx.leafSwitches = append(ctx.leafSwitches, sw.Serial)
		case strings.EqualFold(sw.Role, "spine") && ctx.spineSwitch == "":
			ctx.spineSwitch = sw.Serial
		}
	}

	return ctx
}

func requireLeafPair(t *testing.T, ctx *vpcPairTestContext) {
	t.Helper()

	if len(ctx.leafSwitches) < 2 {
		t.Skip("Need at least 2 leaf switches in nd.inventory.switches for vPC pair acceptance tests")
	}
}

func vpcPairConfigForStep(t *testing.T, testName string, stepCount *int, cfg map[string]string, model *resource_vpc_pair.NDFCVpcPairModel) string {
	t.Helper()

	*stepCount++
	tfConfig := new(string)
	stepName := fmt.Sprintf("%s_%d", testName, *stepCount)
	helper.GetTFConfigWithSingleResource(stepName, cfg, []interface{}{model}, &tfConfig)
	return *tfConfig
}

func TestAccVpcPairResourceCreateRead(t *testing.T) {
	ctx := loadVpcPairTestContext(t)
	requireLeafPair(t, ctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 verifies create with a valid pair of leaf switches.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						ctx.fabricName,
						ctx.leafSwitches[0],
						ctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2 reapplies the same config to force a refresh/read cycle and
			// confirm that state remains stable after the provider reads from ND.
			{
				Config: vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}

func TestAccVpcPairResourceUpdateDeploy(t *testing.T) {
	ctx := loadVpcPairTestContext(t)
	requireLeafPair(t, ctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 creates the vPC pair without deployment so the update path
			// can change only the operational deploy flag in the next step.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						ctx.fabricName,
						ctx.leafSwitches[0],
						ctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2 exercises Update by switching deploy from false to true and
			// then asserting the updated state is preserved after apply/read.
			{
				Config: func() string {
					helper.ModifyVpcPairObject(&vpcPairRsc, map[string]interface{}{
						"deploy": true,
					})
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}

func TestAccVpcPairResourceDelete(t *testing.T) {
	ctx := loadVpcPairTestContext(t)
	requireLeafPair(t, ctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// This step creates the vPC pair and lets the acceptance framework
			// trigger Terraform destroy during test teardown, which exercises the
			// provider Delete implementation for this resource.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						ctx.fabricName,
						ctx.leafSwitches[0],
						ctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}

func TestAccVpcPairResourceImport(t *testing.T) {
	ctx := loadVpcPairTestContext(t)
	requireLeafPair(t, ctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 creates the resource that will be imported in the following
			// step using the provider's documented fabric/switch1:switch2 ID format.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						ctx.fabricName,
						ctx.leafSwitches[0],
						ctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2 verifies that ImportState reconstructs the same resource
			// state from the ND API using the compound import identifier.
			{
				ResourceName:      "nd_vpc_pair.vpc_pair_test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId: fmt.Sprintf(
					"%s/%s:%s",
					ctx.fabricName,
					ctx.leafSwitches[0],
					ctx.leafSwitches[1],
				),
			},
		},
	})
}

func TestAccVpcPairResourceInvalidSwitchRole(t *testing.T) {
	ctx := loadVpcPairTestContext(t)
	requireLeafPair(t, ctx)

	if ctx.spineSwitch == "" {
		t.Skip("Need at least 1 spine switch in nd.inventory.switches for invalid-role vPC pair acceptance test")
	}

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// This negative case proves the testbed selection rules and ND-side
			// validation by attempting to form a vPC pair with a spine switch.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						ctx.fabricName,
						ctx.spineSwitch,
						ctx.leafSwitches[0],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, ctx.configMap, vpcPairRsc)
				}(),
				ExpectError: regexp.MustCompile(`(?i)(not vpc capable|spines is not supported)`),
			},
		},
	})
}
