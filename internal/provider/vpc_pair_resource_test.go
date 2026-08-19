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

	if cfg.ND.VpcPair == nil {
		return vpcPairTestctx
	}

	inventoryBySerial := make(map[string]string, len(cfg.ND.Inventory.Switches))
	for _, sw := range cfg.ND.Inventory.Switches {
		inventoryBySerial[sw.Serial] = sw.Role
		if strings.EqualFold(sw.Role, "spine") && vpcPairTestctx.spineSwitch == "" {
			vpcPairTestctx.spineSwitch = sw.Serial
		}
	}

	for _, serial := range cfg.ND.VpcPair.Switches {
		if strings.EqualFold(inventoryBySerial[serial], "leaf") {
			vpcPairTestctx.leafSwitches = append(vpcPairTestctx.leafSwitches, serial)
		}
	}

	return vpcPairTestctx
}

func loadExternalVpcPairTestctx(t *testing.T) *vpcPairTestctx {
	t.Helper()

	vpcPairTestctx := loadVpcPairTestctx(t)
	cfg := helper.GetConfig("global")
	if cfg.ND.VpcPair == nil || cfg.ND.VpcPair.External == nil {
		return vpcPairTestctx
	}

	external := cfg.ND.VpcPair.External
	if external.Fabric != "" {
		vpcPairTestctx.fabricName = external.Fabric
	}

	switch {
	case external.PeerSwitchID != "" && external.SwitchID != "":
		vpcPairTestctx.leafSwitches = []string{external.PeerSwitchID, external.SwitchID}
	case len(external.Switches) >= 2:
		vpcPairTestctx.leafSwitches = external.Switches[:2]
	}

	return vpcPairTestctx
}

func requireLeafPair(t *testing.T, vpcPairTestctx *vpcPairTestctx) {
	t.Helper()

	if len(vpcPairTestctx.leafSwitches) < 2 {
		t.Skip("Need nd.vpc_pair.switches with at least 2 inventory leaf switches for vPC pair acceptance tests")
	}
}

func boolPtr(v bool) *bool {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func buildDefaultVpcPairDetailsPayload(fabricName string) resource_vpc_pair.NDFCVpcPairDetailsValue {
	return resource_vpc_pair.NDFCVpcPairDetailsValue{
		TemplateType:               "default",
		AdminState:                 boolPtr(true),
		AllowedVlans:               "all",
		DomainId:                   int64Ptr(5),
		EnableMirrorConfig:         boolPtr(false),
		FabricPathSwitchId:         int64Ptr(100),
		IsVpcPlus:                  boolPtr(false),
		IsVteps:                    boolPtr(true),
		KeepAliveHoldTimeout:       int64Ptr(3),
		KeepAliveVrf:               "management",
		LoopbackSecondaryIp:        "10.3.0.2",
		NveInterface:               int64Ptr(1),
		PeerSwitchKeepAliveLocalIp: "192.168.10.102",
		PeerSwitchMemberInterfaces: []string{"e1/2"},
		PeerSwitchNativeVlan:       int64Ptr(3600),
		PeerSwitchPoDescription:    "vpc-peer-link leaf1--leaf2",
		PeerSwitchPoId:             int64Ptr(500),
		PeerSwitchPrimaryIp:        "10.3.0.4",
		PeerSwitchSourceLoopback:   int64Ptr(1),
		PoMode:                     "active",
		SwitchKeepAliveLocalIp:     "192.168.10.101",
		SwitchMemberInterfaces:     []string{"e1/2"},
		SwitchNativeVlan:           int64Ptr(3600),
		SwitchPoDescription:        "vpc-peer-link leaf1--leaf2",
		SwitchPoId:                 int64Ptr(500),
		SwitchPrimaryIp:            "10.3.0.3",
		SwitchSourceLoopback:       int64Ptr(1),
	}
}

func buildRequiredVpcPairDetailsPayload() resource_vpc_pair.NDFCVpcPairDetailsValue {
	return resource_vpc_pair.NDFCVpcPairDetailsValue{
		TemplateType:               "default",
		DomainId:                   int64Ptr(1),
		KeepAliveVrf:               "default",
		PeerSwitchKeepAliveLocalIp: "192.168.10.102",
		SwitchKeepAliveLocalIp:     "192.168.10.101",
	}
}

func buildBoundaryVpcPairDetailsPayload() resource_vpc_pair.NDFCVpcPairDetailsValue {
	return resource_vpc_pair.NDFCVpcPairDetailsValue{
		TemplateType:               "default",
		AdminState:                 boolPtr(true),
		AllowedVlans:               "all",
		DomainId:                   int64Ptr(1000),
		EnableMirrorConfig:         boolPtr(false),
		FabricPathSwitchId:         int64Ptr(4096),
		IsVpcPlus:                  boolPtr(false),
		IsVteps:                    boolPtr(false),
		KeepAliveHoldTimeout:       int64Ptr(10),
		KeepAliveVrf:               "management",
		LoopbackSecondaryIp:        "10.3.0.12",
		NveInterface:               int64Ptr(4),
		PeerSwitchKeepAliveLocalIp: "192.168.10.112",
		PeerSwitchMemberInterfaces: []string{"e1/2"},
		PeerSwitchNativeVlan:       int64Ptr(4094),
		PeerSwitchPoDescription:    "vpc-peer-link boundary leaf1--leaf2",
		PeerSwitchPoId:             int64Ptr(4096),
		PeerSwitchPrimaryIp:        "10.3.0.14",
		PeerSwitchSourceLoopback:   int64Ptr(4),
		PoMode:                     "passive",
		SwitchKeepAliveLocalIp:     "192.168.10.111",
		SwitchMemberInterfaces:     []string{"e1/2"},
		SwitchNativeVlan:           int64Ptr(4094),
		SwitchPoDescription:        "vpc-peer-link boundary leaf1--leaf2",
		SwitchPoId:                 int64Ptr(4096),
		SwitchPrimaryIp:            "10.3.0.13",
		SwitchSourceLoopback:       int64Ptr(4),
	}
}

func buildUpdatedVpcPairDetailsPayload(fabricName string) resource_vpc_pair.NDFCVpcPairDetailsValue {
	updated := buildDefaultVpcPairDetailsPayload(fabricName)
	updated.PeerSwitchPoDescription = "vpc-peer-link leaf1--leaf2 updated"
	updated.SwitchPoDescription = "vpc-peer-link leaf1--leaf2 updated"
	updated.KeepAliveHoldTimeout = int64Ptr(4)
	return updated
}

func vpcPairDetailsStateCheck(rscName string, details resource_vpc_pair.NDFCVpcPairDetailsValue) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.template_type", details.TemplateType),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.admin_state", strconv.FormatBool(*details.AdminState)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.allowed_vlans", details.AllowedVlans),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.domain_id", strconv.FormatInt(*details.DomainId, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.enable_mirror_config", strconv.FormatBool(*details.EnableMirrorConfig)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.fabric_path_switch_id", strconv.FormatInt(*details.FabricPathSwitchId, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.is_vpc_plus", strconv.FormatBool(*details.IsVpcPlus)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.is_vteps", strconv.FormatBool(*details.IsVteps)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.keep_alive_hold_timeout", strconv.FormatInt(*details.KeepAliveHoldTimeout, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.keep_alive_vrf", details.KeepAliveVrf),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.loopback_secondary_ip", details.LoopbackSecondaryIp),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.nve_interface", strconv.FormatInt(*details.NveInterface, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_keep_alive_local_ip", details.PeerSwitchKeepAliveLocalIp),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_member_interfaces.0", details.PeerSwitchMemberInterfaces[0]),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_native_vlan", strconv.FormatInt(*details.PeerSwitchNativeVlan, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_po_description", details.PeerSwitchPoDescription),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_po_id", strconv.FormatInt(*details.PeerSwitchPoId, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_primary_ip", details.PeerSwitchPrimaryIp),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_source_loopback", strconv.FormatInt(*details.PeerSwitchSourceLoopback, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.po_mode", details.PoMode),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_keep_alive_local_ip", details.SwitchKeepAliveLocalIp),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_member_interfaces.0", details.SwitchMemberInterfaces[0]),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_native_vlan", strconv.FormatInt(*details.SwitchNativeVlan, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_po_description", details.SwitchPoDescription),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_po_id", strconv.FormatInt(*details.SwitchPoId, 10)),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_primary_ip", details.SwitchPrimaryIp),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_source_loopback", strconv.FormatInt(*details.SwitchSourceLoopback, 10)),
		resource.TestCheckResourceAttrSet(rscName, "id"),
		resource.TestCheckResourceAttrSet(rscName, "fabric_name"),
	}
}

func vpcPairDetailsDefaultStateCheck(rscName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.template_type", "default"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.domain_id", "1"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.keep_alive_vrf", "default"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.peer_switch_keep_alive_local_ip", "192.168.10.102"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.switch_keep_alive_local_ip", "192.168.10.101"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.admin_state", "true"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.allowed_vlans", "all"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.enable_mirror_config", "false"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.is_vpc_plus", "false"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.is_vteps", "false"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.keep_alive_hold_timeout", "3"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.nve_interface", "1"),
		resource.TestCheckResourceAttr(rscName, "vpc_pair_details.po_mode", "active"),
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

func vpcPairConfigWithExplicitEmptyDetails(cfg map[string]string, rscName, switchID1, switchID2 string, useVirtualPeerLink, deploy bool) string {
	return fmt.Sprintf(`provider "nd" {
  username = "%s"
  password = "%s"
  url      = "%s"
  insecure = %s
}

resource "nd_vpc_pair" "%s" {
  switch_id_1          = "%s"
  switch_id_2          = "%s"
  use_virtual_peerlink = %t
  vpc_pair_details     = {}
  deploy               = %t
}
`, cfg["User"], cfg["Password"], cfg["Host"], cfg["Insecure"], rscName, switchID1, switchID2, useVirtualPeerLink, deploy)
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
	vpcPairAPI := api.NewVpcPairAPI(client, model.FabricName)
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
						true,
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

func TestAccVpcPairResourceCreateReadWithVpcPairDetails(t *testing.T) {
	vpcPairTestctx := loadExternalVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 proves the nested block now enforces the Cisco-required
			// fields before any request is sent to ND.
			{
				Config: func() string {
					stepCount++
					return vpcPairConfigWithExplicitEmptyDetails(
						vpcPairTestctx.configMap,
						"vpc_pair_test",
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						true,
					)
				}(),
				ExpectError: regexp.MustCompile(`(?s)attributes "domain_id".*"keep_alive_vrf".*"peer_switch_keep_alive_local_ip".*"switch_keep_alive_local_ip".*"template_type" are required`),
			},
			// Step 2 creates the pair with only the Cisco-required fields plus
			// minimum boundary values. The check confirms the provider/ND path
			// populates documented defaults for omitted optional attributes.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						true,
					)
					vpcPairRsc.VpcPairDetails = buildRequiredVpcPairDetailsPayload()
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						VpcPairModelHelperStateCheck(
							"nd_vpc_pair.vpc_pair_test",
							*vpcPairRsc,
							path.Empty(),
						),
						vpcPairDetailsDefaultStateCheck("nd_vpc_pair.vpc_pair_test")...,
					)...,
				),
			},
			// Step 3 reapplies unchanged config to prove the discovered state
			// stays stable after a read/refresh cycle with required/defaulted
			// nested payload fields.
			{
				Config: vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc),
				Check: resource.ComposeTestCheckFunc(
					append(
						VpcPairModelHelperStateCheck(
							"nd_vpc_pair.vpc_pair_test",
							*vpcPairRsc,
							path.Empty(),
						),
						vpcPairDetailsDefaultStateCheck("nd_vpc_pair.vpc_pair_test")...,
					)...,
				),
			},
		},
	})
}

func TestAccVpcPairResourceUpdateVpcPairDetails(t *testing.T) {
	vpcPairTestctx := loadExternalVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 creates the external-fabric pair with the minimum required
			// nested payload. This establishes a baseline resource before
			// exercising updates at the upper boundary values.
			{
				Config: func() string {
					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						vpcPairTestctx.fabricName,
						vpcPairTestctx.leafSwitches[0],
						vpcPairTestctx.leafSwitches[1],
						false,
						true,
					)
					vpcPairRsc.VpcPairDetails = buildRequiredVpcPairDetailsPayload()
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						VpcPairModelHelperStateCheck(
							"nd_vpc_pair.vpc_pair_test",
							*vpcPairRsc,
							path.Empty(),
						),
						vpcPairDetailsDefaultStateCheck("nd_vpc_pair.vpc_pair_test")...,
					)...,
				),
			},
			// Step 2 updates the nested payload to maximum boundary values for
			// the fields Cisco constrains, proving the provider can round-trip
			// the high end of the supported ranges as well.
			{
				Config: func() string {
					vpcPairRsc.VpcPairDetails = buildBoundaryVpcPairDetailsPayload()
					return vpcPairConfigForStep(t, t.Name(), &stepCount, vpcPairTestctx.configMap, vpcPairRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						VpcPairModelHelperStateCheck(
							"nd_vpc_pair.vpc_pair_test",
							*vpcPairRsc,
							path.Empty(),
						),
						vpcPairDetailsStateCheck(
							"nd_vpc_pair.vpc_pair_test",
							vpcPairRsc.VpcPairDetails,
						)...,
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
					"%s:%s",
					vpcPairTestctx.leafSwitches[0],
					vpcPairTestctx.leafSwitches[1],
				),
			},
		},
	})
}

func TestAccVpcPairResourceImportThenUpdateDeploy(t *testing.T) {
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	stepCount := 0
	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1 creates the vPC pair without deployment so the imported
			// state starts with deploy=false, matching the manual workflow.
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
			// Step 2 imports the same pair back into empty state using the
			// compound identifier, reproducing the manual import flow.
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
			// Step 3 reapplies the unchanged config and expects no behavioral
			// drift after import, matching the "terraform plan" no-op check.
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
			// Step 4 changes only deploy from false to true and verifies the
			// provider performs an in-place update after import.
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
				ExpectError: regexp.MustCompile(`(?i)(not vpc capable|spines is not supported|different roles)`),
			},
		},
	})
}
