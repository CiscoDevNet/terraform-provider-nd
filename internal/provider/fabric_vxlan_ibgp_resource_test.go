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
	"log"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/netascode/go-nd"
)

func TestAccFabricVxlanIbgpResourceCRUD(t *testing.T) {
	x := &map[string]string{
		"RscType":  "nd_fabric_vxlan_ibgp",
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
			// Step 1: Create a basic iBGP VXLAN fabric
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricIbgpObject(&fabricRsc,
						helper.GetConfig("global").ND.Fabric,
						"55000",
						nil,
					)

					(*x)["RscName"] = "fabric_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.IbgpResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanIbgpModelHelperStateCheck(
						"nd_fabric_vxlan_ibgp.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Modify fabric parameters (MTU)
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyFabricIbgpObject(&fabricRsc, map[string]interface{}{
						"fabric_mtu":            9000,
						"l2_host_interface_mtu": 9000,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.IbgpResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanIbgpModelHelperStateCheck(
						"nd_fabric_vxlan_ibgp.fabric_test",
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

					helper.ModifyFabricIbgpObject(&fabricRsc, map[string]interface{}{
						"vpc_peer_link_vlan": "3601",
						"l2_vni_range":       "30000-48000",
						"l3_vni_range":       "50000-58000",
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.IbgpResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					FabricVxlanIbgpModelHelperStateCheck(
						"nd_fabric_vxlan_ibgp.fabric_test",
						*fabricRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}

// TestAccFabricVxlanIbgpResourceDrift verifies that when a fabric is deleted
// out-of-band, the provider detects the drift (Read returns 404 → state is
// removed) and Terraform plans a re-create on the next apply.
func TestAccFabricVxlanIbgpResourceDrift(t *testing.T) {
	cfg := helper.GetConfig("global")
	fabricName := cfg.ND.Fabric + "_ibgp_drift"

	x := &map[string]string{
		"RscType":  "nd_fabric_vxlan_ibgp",
		"RscName":  "fabric_drift",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_common.NDFCFabricCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the fabric
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateFabricIbgpObject(&fabricRsc,
						fabricName, "55000", nil,
					)

					(*x)["RscName"] = "fabric_drift"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{helper.IbgpResource(fabricRsc)}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nd_fabric_vxlan_ibgp.fabric_drift",
						"fabric_name", fabricName,
					),
				),
			},
			// Step 2: Delete fabric out-of-band, then refresh.
			// Read should detect 404, remove from state → non-empty plan.
			{
				PreConfig: func() {
					client, err := nd.NewClient(
						cfg.ND.URL, "/api/v1",
						cfg.ND.User, cfg.ND.Password,
						"", cfg.ND.Insecure == "true",
						nd.MaxRetries(3),
					)
					if err != nil {
						t.Fatalf("Failed to create ND client for out-of-band delete: %v", err)
					}
					fabricAPI := api.NewFabricAPI(&client, ndapi.DefaultFabric)
					fabricAPI.FabricName = fabricName
					res, err := fabricAPI.Delete()
					if err != nil {
						t.Fatalf("Failed to delete fabric %q out-of-band: %v: %s", fabricName, err, res.String())
					}
					log.Printf("Out-of-band delete of fabric %q succeeded", fabricName)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Step 3: Re-apply the same config — Terraform should recreate the fabric
			{
				Config: *tfConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"nd_fabric_vxlan_ibgp.fabric_drift",
						"fabric_name", fabricName,
					),
				),
			},
		},
	})
}
