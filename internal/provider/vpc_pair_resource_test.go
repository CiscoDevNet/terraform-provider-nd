// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"
	"terraform-provider-nd/internal/manage/resource_vpc_pair"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	nd "github.com/netascode/go-nd"
)

type vpcPairTestctx struct {
	configMap    map[string]string
	fabricName   string
	leafSwitches []string
	spineSwitch  string
}

func loadVpcPairTestctx(t *testing.T) *vpcPairTestctx {
	t.Helper()

	cfg := helper.GetConfig("global")
	vpcPairTestctx := &vpcPairTestctx{
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
			vpcPairTestctx.leafSwitches = append(vpcPairTestctx.leafSwitches, sw.Serial)
		case strings.EqualFold(sw.Role, "spine") && vpcPairTestctx.spineSwitch == "":
			vpcPairTestctx.spineSwitch = sw.Serial
		}
	}

	return vpcPairTestctx
}

func requireLeafPair(t *testing.T, vpcPairTestctx *vpcPairTestctx) {
	t.Helper()

	if len(vpcPairTestctx.leafSwitches) < 2 {
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

func newVpcPairTestClient(t *testing.T) *nd.Client {
	t.Helper()

	cfg := helper.GetConfig("global")
	insecure, err := strconv.ParseBool(cfg.ND.Insecure)
	if err != nil {
		t.Fatalf("failed to parse nd.insecure test setting %q: %v", cfg.ND.Insecure, err)
	}

	client, err := nd.NewClient(cfg.ND.URL, "/api/v1", cfg.ND.User, cfg.ND.Password, "", insecure)
	if err != nil {
		t.Fatalf("failed to create ND client for vPC pair acceptance helper: %v", err)
	}

	return &client
}

func deleteVpcPairOutsideTerraform(t *testing.T, model *resource_vpc_pair.NDFCVpcPairModel) {
	t.Helper()

	client := newVpcPairTestClient(t)
	vpcPairAPI := api.NewVpcPairAPI(nil, client)
	vpcPairAPI.FabricName = model.FabricName
	vpcPairAPI.SwitchID = model.SwitchId2

	deleteRequest := *model
	deleteRequest.VpcAction = "unPair"

	payload, err := json.Marshal(deleteRequest)
	if err != nil {
		t.Fatalf("failed to marshal out-of-band vPC pair delete payload: %v", err)
	}

	res, err := vpcPairAPI.Put(payload, nil)
	if err != nil {
		t.Fatalf("failed to delete vPC pair outside Terraform: %v %v", err, res.Raw)
	}

	var diagnostics diag.Diagnostics
	deployment.ConfigSaveAndDeploy(context.Background(), client, model.FabricName, true, model.Deploy, &diagnostics)
	if diagnostics.HasError() {
		t.Fatalf("failed to save/deploy config after out-of-band vPC pair delete: %v", diagnostics)
	}

	time.Sleep(2 * time.Second)
}

func TestAccVpcPairResourceCreateRead(t *testing.T) {
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

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
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
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
				Config: vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc),
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
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

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
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
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
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
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

func TestAccVpcPairResourceRecreateAfterOutOfBandDelete(t *testing.T) {
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 creates a valid leaf-to-leaf vPC pair so the next step can
			// simulate a deletion that happened outside Terraform.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2 deletes the pair directly in ND before Terraform refreshes
			// state, proving the provider removes stale state on Read and plans
			// the same configuration as a recreation instead of failing.
			{
				PreConfig: func() {
					deleteVpcPairOutsideTerraform(t, vpcPairRsc)
				},
				Config: vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc),
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
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

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
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
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
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

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
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
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
					vpcPairTestctx.fabricName,
					vpcPairTestctx.leafSwitches[0],
					vpcPairTestctx.leafSwitches[1],
				),
			},
		},
	})
}

func TestAccVpcPairResourceInvalidSwitchRole(t *testing.T) {
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	if vpcPairTestctx.spineSwitch == "" {
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
						vpcPairTestctx.fabricName,
						vpcPairTestctx.spineSwitch,
						vpcPairTestctx.leafSwitches[0],
						false,
						false,
					)
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
				}(),
				ExpectError: regexp.MustCompile(`(?i)(not vpc capable|spines is not supported)`),
			},
		},
	})
}
