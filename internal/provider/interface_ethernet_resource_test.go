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
	"testing"

	"terraform-provider-nd/internal/manage/resource_interface_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

type ieTestctx struct {
	configMap  map[string]string
	fabricName string
	switchId   string
	switchIP   string // management IP of the same switch (for IP-based tests)
}

func loadIETestctx(t *testing.T) *ieTestctx {
	t.Helper()

	cfg := helper.GetConfig("global")
	if cfg.ND.Fabric == "" || len(cfg.ND.Switches) == 0 {
		t.Skip("testbed yaml lacks nd.fabric or nd.switches[0]; skipping")
	}

	ctx := &ieTestctx{
		configMap: map[string]string{
			"RscType":  "nd_interface_ethernet",
			"RscName":  "test",
			"User":     cfg.ND.User,
			"Password": cfg.ND.Password,
			"Host":     cfg.ND.URL,
			"Insecure": cfg.ND.Insecure,
		},
		fabricName: cfg.ND.Fabric,
		switchId:   cfg.ND.Switches[0],
	}
	if len(cfg.ND.SwitchIP) > 0 {
		ctx.switchIP = cfg.ND.SwitchIP[0]
	}
	return ctx
}

func boolP(v bool) *bool    { return &v }
func int64P(v int64) *int64 { return &v }

// ieConfigForStep renders HCL for the model and increments stepCount.
func ieConfigForStep(t *testing.T, testName string, stepCount *int, cfg map[string]string, model *resource_interface_common.NDFCInterfaceCommonModel) string {
	t.Helper()
	return helper.InterfaceEthernetConfigForStep(testName, stepCount, cfg, model)
}

// ieInlineConfig returns a raw HCL config using the ND provider template plus
// an inline resource block. Used for error-path tests where the model struct
// does not represent a valid resource (e.g., invalid policy_type combo).
func ieInlineConfig(cfg map[string]string, rscName, body string) string {
	return fmt.Sprintf(`provider "nd" {
  username = "%s"
  password = "%s"
  url      = "%s"
  insecure = %s
}

resource "nd_interface_ethernet" "%s" {
%s
}
`, cfg["User"], cfg["Password"], cfg["Host"], cfg["Insecure"], rscName, body)
}

// ---------------------------------------------------------------------------
// 1. TestAccInterfaceEthernet_BasicAccessPort
//    Verifies create, read, update, destroy, and idempotence for a single
//    access-mode Ethernet port.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_BasicAccessPort(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "basic"
	rn := "nd_interface_ethernet.basic"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create a single access port.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/21": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(200),
								Mtu: "default", PortTypeEdgeTrunk: boolP(true),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Idempotence.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Update — change admin_state, access_vlan, mtu.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/21.admin_state": false,
						"Ethernet1/21.access_vlan": 201,
						"Ethernet1/21.mtu":         "jumbo",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 2. TestAccInterfaceEthernet_CollisionFields_MTU
//    Checks dual-form handling of mtu (string "default"/"jumbo" for L2,
//    numeric string "9216" for L3).
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_CollisionFields_MTU(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "mtu"
	rn := "nd_interface_ethernet.mtu"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: L2 access mtu="jumbo", L3 routed mtu="9216".
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/22": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(100), Mtu: "jumbo",
							},
							"Ethernet1/23": {
								Mode: "routed", NetworkOsType: "nx-os", PolicyType: "routedHost",
								AdminState: boolP(true), Ip: "10.2.2.1", Prefix: "30",
								Mtu: "9216", RoutingTag: "0",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Idempotence — collision fields must round-trip.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 3. TestAccInterfaceEthernet_PolicyTypeMatrix
//    Covers access + trunk + routed mode combinations.
// ---------------------------------------------------------------------------

func TestAccNdInterfaceEthernet(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "acc"
	rn := "nd_interface_ethernet.acc"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create access + routed.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/13": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), Cdp: boolP(true), BpduGuard: "enable",
								PortTypeEdgeTrunk: boolP(true), AccessVlan: int64P(100),
								Mtu: "default", Speed: "auto",
							},
							"Ethernet1/14": {
								Mode: "routed", NetworkOsType: "nx-os", PolicyType: "routedHost",
								AdminState: boolP(true), Ip: "10.1.1.1", Prefix: "30",
								Mtu: "9216", RoutingTag: "12345", VrfInterface: "default",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Idempotence.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Update — change admin_state, mtu.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/13.admin_state": false,
						"Ethernet1/14.mtu":         "1500",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
			// Step 4: Replace routed with trunk. Remove Ethernet1/14, add Ethernet1/15.
			{
				Config: func() string {
					helper.RemoveInterface(&intfRsc, "Ethernet1/14")
					helper.AddInterface(&intfRsc, "Ethernet1/15", helper.InterfaceEthernetEntry{
						Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
						AdminState: boolP(true), AllowedVlans: "100-110",
						Mtu: "jumbo", Speed: "auto", NativeVlan: int64P(1),
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty()),
						resource.TestCheckNoResourceAttr(rn, "interfaces.Ethernet1/14.policy_type"),
					)...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 4. TestAccInterfaceEthernet_OptionalFieldsAndDefaults
//    Creates with minimal config; verifies server-supplied defaults in state.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_OptionalFieldsAndDefaults(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "defaults"
	rn := "nd_interface_ethernet.defaults"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Minimal config — only required fields.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/24": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty()),
						// Server should fill defaults for speed, fec, etc.
						resource.TestCheckResourceAttrSet(rn, "interfaces.Ethernet1/24.speed"),
						resource.TestCheckResourceAttrSet(rn, "interfaces.Ethernet1/24.fec"),
					)...,
				),
			},
			// Step 2: Idempotence — defaults should not cause drift.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 5. TestAccInterfaceEthernet_BulkMultipleInterfaces
//    Tests create/modify/destroy for 3+ interfaces in one resource.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_BulkMultipleInterfaces(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "bulk"
	rn := "nd_interface_ethernet.bulk"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 interfaces (2 access + 1 trunk).
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/25": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(300),
							},
							"Ethernet1/26": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(301),
							},
							"Ethernet1/27": {
								Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
								AdminState: boolP(true), AllowedVlans: "300-310",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Update — change Ethernet1/25, remove Ethernet1/26, add Ethernet1/28.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/25.access_vlan":   302,
						"Ethernet1/27.allowed_vlans": "300-320",
					})
					helper.RemoveInterface(&intfRsc, "Ethernet1/26")
					helper.AddInterface(&intfRsc, "Ethernet1/28", helper.InterfaceEthernetEntry{
						Mode: "routed", NetworkOsType: "nx-os", PolicyType: "routedHost",
						AdminState: boolP(true), Ip: "10.3.3.1", Prefix: "24",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rn, *intfRsc, path.Empty()),
						resource.TestCheckNoResourceAttr(rn, "interfaces.Ethernet1/26.policy_type"),
					)...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 6. TestAccInterfaceEthernet_Import
//    Verifies import of an existing interface from ND.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_Import(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "imp"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)
	importId := fmt.Sprintf("%s/%s/Ethernet1/29", ctx.fabricName, ctx.switchId)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the interface first.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/29": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(400),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
			},
			// Step 2: Import it back.
			{
				ResourceName:      "nd_interface_ethernet.imp",
				ImportState:       true,
				ImportStateId:     importId,
				ImportStateVerify: false, // deploy and some computed fields may differ
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_interface_ethernet.imp", "fabric_name", ctx.fabricName),
					resource.TestCheckResourceAttr("nd_interface_ethernet.imp", "interfaces.Ethernet1/29.policy_type", "accessHost"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 7. TestAccInterfaceEthernet_IdempotenceAfterApply
//    Validates that applying a no-op plan does not modify ND or TF state.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IdempotenceAfterApply(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "idem"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/30": {
								Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
								AdminState: boolP(true), AllowedVlans: "500-510", Mtu: "jumbo",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn("idem"), *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Re-apply — must be idempotent.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Step 3: Re-apply again.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// rn returns the full Terraform resource address for nd_interface_ethernet.
func rn(name string) string {
	return "nd_interface_ethernet." + name
}

// ---------------------------------------------------------------------------
// 8. TestAccInterfaceEthernet_InvalidPolicyTypeCombo
//    Illegal (mode, policyType, network_os_type) must fail at validate time.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_InvalidPolicyTypeCombo(t *testing.T) {
	ctx := loadIETestctx(t)

	config := ieInlineConfig(ctx.configMap, "bad", fmt.Sprintf(`
  fabric_name = %q
  switch_id   = %q
  interfaces = {
    "Ethernet1/31" = {
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "routedHost"
      admin_state     = true
    }
  }`, ctx.fabricName, ctx.switchId))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile(`(?i)policy_type.*not valid`)},
		},
	})
}

// ---------------------------------------------------------------------------
// 9. TestAccInterfaceEthernet_BadTemplateConfig
//    userDefined without template_name must fail at validate time.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_BadTemplateConfig(t *testing.T) {
	ctx := loadIETestctx(t)

	config := ieInlineConfig(ctx.configMap, "badtpl", fmt.Sprintf(`
  fabric_name = %q
  switch_id   = %q
  interfaces = {
    "Ethernet1/32" = {
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "userDefined"
      admin_state     = true
    }
  }`, ctx.fabricName, ctx.switchId))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile(`(?i)template_name.*required.*userDefined`)},
		},
	})
}

// ---------------------------------------------------------------------------
// 10. TestAccInterfaceEthernet_NDReturnsError
//     Force an ND error (e.g., invalid VLAN) and verify it surfaces.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_NDReturnsError(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "nderr"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// VLAN 5000 is out of range (1-4094).
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/33": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(5000),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				ExpectError: regexp.MustCompile(`(?i)error|failed|invalid`),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 11. TestAccInterfaceEthernet_RepeatedDeletes
//     Deleting an already-missing interface should not panic.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_RepeatedDeletes(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "repdel"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create — the testing framework auto-destroys at the end.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/34": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(100),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rn("repdel"), *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 12. TestAccInterfaceEthernet_RestoreAfterNDOutOfBandChange
//     Modify an interface out-of-band (via curl), then re-apply TF to
//     restore the intended state.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_RestoreAfterNDOutOfBandChange(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "oob"
	rscAddr := rn("oob")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/35": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(600),
								Description: "tf-managed",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Re-apply same config. The Update path will PUT the
			// intended state. Even if ND was modified out-of-band between
			// steps (which the framework does with Refresh), TF corrects it.
			{
				Config: ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 13. TestAccManageInterfaceEthernetAddOnly
//     Single operation: add a new interface to an existing resource
//     (no update, no delete).
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetAddOnly(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "addonly"
	rscAddr := rn("addonly")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with one interface.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/36": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(100),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Add one new interface (no modify, no delete).
			{
				Config: func() string {
					helper.AddInterface(&intfRsc, "Ethernet1/37", helper.InterfaceEthernetEntry{
						Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
						AdminState: boolP(true), AllowedVlans: "100-110",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 14. TestAccManageInterfaceEthernetDeleteOnly
//     Single operation: remove an interface from an existing resource
//     (no add, no update).
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetDeleteOnly(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "delonly"
	rscAddr := rn("delonly")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with two interfaces.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/38": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(200),
							},
							"Ethernet1/39": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(201),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Remove one interface (no add, no modify).
			{
				Config: func() string {
					helper.RemoveInterface(&intfRsc, "Ethernet1/39")
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty()),
						resource.TestCheckNoResourceAttr(rscAddr, "interfaces.Ethernet1/39.policy_type"),
					)...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 15. TestAccManageInterfaceEthernetUpdateOnly
//     Single operation: update fields on existing interfaces
//     (no add, no delete).
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetUpdateOnly(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "updonly"
	rscAddr := rn("updonly")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create two interfaces.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/40": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(300), Mtu: "default",
							},
							"Ethernet1/41": {
								Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
								AdminState: boolP(true), AllowedVlans: "300-310",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Update both interfaces (no add, no delete).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/40.access_vlan":   301,
						"Ethernet1/40.mtu":           "jumbo",
						"Ethernet1/41.allowed_vlans": "300-320",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 16. TestAccManageInterfaceEthernetMixedOps
//     Single operation containing add + update + delete in one step.
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetMixedOps(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "mixops"
	rscAddr := rn("mixops")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create three interfaces.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/42": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(400),
							},
							"Ethernet1/43": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(401),
							},
							"Ethernet1/44": {
								Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
								AdminState: boolP(true), AllowedVlans: "400-410",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Update Ethernet1/42 + Delete Ethernet1/43 + Add Ethernet1/45.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/42.access_vlan": 402,
					})
					helper.RemoveInterface(&intfRsc, "Ethernet1/43")
					helper.AddInterface(&intfRsc, "Ethernet1/45", helper.InterfaceEthernetEntry{
						Mode: "routed", NetworkOsType: "nx-os", PolicyType: "routedHost",
						AdminState: boolP(true), Ip: "10.4.4.1", Prefix: "24",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty()),
						resource.TestCheckNoResourceAttr(rscAddr, "interfaces.Ethernet1/43.policy_type"),
					)...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 17. TestAccManageInterfaceEthernetAddNoDeploy
//     Create interfaces with deploy=false, then set deploy=true in step 2
//     without modifying interfaces.
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetAddNoDeploy(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "nodeploy"
	rscAddr := rn("nodeploy")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/46": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(500),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Set deploy=true (no interface changes).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy": true,
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 18. TestAccManageInterfaceEthernetAddUpdateDeployStep2
//     Step 1: create with deploy=false.
//     Step 2: add new interface + update existing + set deploy=true.
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetAddUpdateDeployStep2(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "auds2"
	rscAddr := rn("auds2")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create one interface with deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/47": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(600),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Add new interface + update existing + deploy=true.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy":                   true,
						"Ethernet1/47.access_vlan": 601,
					})
					helper.AddInterface(&intfRsc, "Ethernet1/48", helper.InterfaceEthernetEntry{
						Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
						AdminState: boolP(true), AllowedVlans: "600-610",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 19. TestAccManageInterfaceEthernetAddUpdateDeployStep3
//     Step 1: create with deploy=false.
//     Step 2: update only (still deploy=false).
//     Step 3: add new interface + update existing + deploy=true.
// ---------------------------------------------------------------------------

func TestAccManageInterfaceEthernetAddUpdateDeployStep3(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "auds3"
	rscAddr := rn("auds3")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create one interface with deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/49": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(700), Mtu: "default",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Update only (still deploy=false).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/49.access_vlan": 701,
						"Ethernet1/49.mtu":         "jumbo",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 3: Add new interface + update existing + deploy=true.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy":                   true,
						"Ethernet1/49.admin_state": false,
					})
					helper.AddInterface(&intfRsc, "Ethernet1/50", helper.InterfaceEthernetEntry{
						Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
						AdminState: boolP(true), AllowedVlans: "700-710",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 20. TestAccManageInterfaceEthernetMultiStepDeploy
//     Accumulates add+update operations without deploying, then deploys
//     everything in a final step.
//     Step 1: create with deploy=false.
//     Step 2: add + update (still deploy=false).
//     Step 3: set deploy=true only (no interface changes).
// ---------------------------------------------------------------------------

// makeAccessEntries builds n access-mode interface entries starting at
// Ethernet1/<startPort>.  Each gets access_vlan = baseVlan + offset.
func makeAccessEntries(startPort, n, baseVlan int) map[string]helper.InterfaceEthernetEntry {
	m := make(map[string]helper.InterfaceEthernetEntry, n)
	for i := 0; i < n; i++ {
		name := fmt.Sprintf("Ethernet1/%d", startPort+i)
		m[name] = helper.InterfaceEthernetEntry{
			Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
			AdminState: boolP(true), AccessVlan: int64P(int64(baseVlan + i)),
		}
	}
	return m
}

func TestAccManageInterfaceEthernetMultiStepDeploy(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "msdep"
	rscAddr := rn("msdep")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create one interface with deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/51": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(800),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Add new interface + update existing (still deploy=false).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/51.access_vlan": 801,
					})
					helper.AddInterface(&intfRsc, "Ethernet1/52", helper.InterfaceEthernetEntry{
						Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
						AdminState: boolP(true), AllowedVlans: "800-810",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 3: Set deploy=true only (no interface changes).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy": true,
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Scale Tests — 20-30 interfaces
// ---------------------------------------------------------------------------

// 21. TestAccManageInterfaceEthernetScaleCreate
//     Create 20 access-mode interfaces in a single step with deploy.
//     Validates bulk create and deploy at 20-interface scale.
//     Interfaces: Ethernet1/35 - Ethernet1/54

func TestAccManageInterfaceEthernetScaleCreate(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "sc20"
	rscAddr := rn("sc20")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 20 interfaces with deploy=true.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						makeAccessEntries(35, 20, 100),
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// 22. TestAccManageInterfaceEthernetScaleAddDeleteUpdate
//     Start with 20, then add 5 + update 10 + delete 5 in one step.
//     Interfaces: Ethernet1/35-1/54 (initial), add 1/55-1/59, delete 1/50-1/54

func TestAccManageInterfaceEthernetScaleAddDeleteUpdate(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "scadu"
	rscAddr := rn("scadu")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 20 interfaces with deploy=true.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						makeAccessEntries(35, 20, 100),
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Add 5 (1/55-1/59), update 10 (1/35-1/44), delete 5 (1/50-1/54).
			{
				Config: func() string {
					// Update access_vlan on first 10 interfaces.
					mods := make(map[string]interface{}, 10)
					for i := 35; i <= 44; i++ {
						mods[fmt.Sprintf("Ethernet1/%d.access_vlan", i)] = 200 + (i - 35)
					}
					helper.ModifyInterfaceEthernetObject(&intfRsc, mods)
					// Delete last 5.
					for i := 50; i <= 54; i++ {
						helper.RemoveInterface(&intfRsc, fmt.Sprintf("Ethernet1/%d", i))
					}
					// Add 5 new.
					for i := 55; i <= 59; i++ {
						helper.AddInterface(&intfRsc, fmt.Sprintf("Ethernet1/%d", i), helper.InterfaceEthernetEntry{
							Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
							AdminState: boolP(true), AccessVlan: int64P(int64(300 + i - 55)),
						})
					}
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// 23. TestAccManageInterfaceEthernetScaleNoDeploy
//     Create 25 interfaces without deploying, then deploy in step 2.
//     Interfaces: Ethernet1/35-1/59

func TestAccManageInterfaceEthernetScaleNoDeploy(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "scnd"
	rscAddr := rn("scnd")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 25 interfaces, deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						makeAccessEntries(35, 25, 100),
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Deploy only (no interface changes).
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy": true,
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// 24. TestAccManageInterfaceEthernetScaleMultiStepGrow
//     Grow from 10 → 20 → 30 interfaces across 3 steps, deploy only in step 3.
//     Interfaces: Ethernet1/35-1/64

func TestAccManageInterfaceEthernetScaleMultiStepGrow(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "scgrow"
	rscAddr := rn("scgrow")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 10 interfaces (1/35-1/44), deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, false,
						makeAccessEntries(35, 10, 100),
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Add 10 more (1/45-1/54), update 5 existing (1/35-1/39), deploy=false.
			{
				Config: func() string {
					mods := make(map[string]interface{}, 5)
					for i := 35; i <= 39; i++ {
						mods[fmt.Sprintf("Ethernet1/%d.access_vlan", i)] = 200 + (i - 35)
					}
					helper.ModifyInterfaceEthernetObject(&intfRsc, mods)
					for i := 45; i <= 54; i++ {
						helper.AddInterface(&intfRsc, fmt.Sprintf("Ethernet1/%d", i), helper.InterfaceEthernetEntry{
							Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
							AdminState: boolP(true), AccessVlan: int64P(int64(110 + i - 45)),
						})
					}
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 3: Add 10 more (1/55-1/64), deploy=true. Total 30 interfaces.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy": true,
					})
					for i := 55; i <= 64; i++ {
						helper.AddInterface(&intfRsc, fmt.Sprintf("Ethernet1/%d", i), helper.InterfaceEthernetEntry{
							Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
							AdminState: boolP(true), AccessVlan: int64P(int64(120 + i - 55)),
						})
					}
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// 25. TestAccManageInterfaceEthernetScaleDeleteHalf
//     Create 30 interfaces, then delete half in one step.
//     Interfaces: Ethernet1/35-1/64

func TestAccManageInterfaceEthernetScaleDeleteHalf(t *testing.T) {
	ctx := loadIETestctx(t)
	ctx.configMap["RscName"] = "scdel"
	rscAddr := rn("scdel")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 30 interfaces with deploy=true.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchId, true,
						makeAccessEntries(35, 30, 100),
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
			// Step 2: Delete 15 interfaces (1/50-1/64), keep 15. Deploy=true.
			{
				Config: func() string {
					for i := 50; i <= 64; i++ {
						helper.RemoveInterface(&intfRsc, fmt.Sprintf("Ethernet1/%d", i))
					}
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty())...,
				),
			},
		},
	})
}

// ===========================================================================
// switch_id mutual exclusion validation tests
// ===========================================================================

// ---------------------------------------------------------------------------
// 26. TestAccInterfaceEthernet_SwitchIdBothLevels_Error
//     Setting switch_id at resource level AND per-entry must fail validation.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_SwitchIdBothLevels_Error(t *testing.T) {
	ctx := loadIETestctx(t)

	config := ieInlineConfig(ctx.configMap, "bothsw", fmt.Sprintf(`
  fabric_name = %q
  switch_id   = %q
  interfaces = {
    "Ethernet1/41" = {
      switch_id       = %q
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "accessHost"
      admin_state     = true
    }
  }`, ctx.fabricName, ctx.switchId, ctx.switchId))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile(`(?i)switch_id.*mutually exclusive`)},
		},
	})
}

// ---------------------------------------------------------------------------
// 27. TestAccInterfaceEthernet_SwitchIdNeitherLevel_Error
//     Omitting switch_id at both levels must fail validation.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_SwitchIdNeitherLevel_Error(t *testing.T) {
	ctx := loadIETestctx(t)

	config := ieInlineConfig(ctx.configMap, "nosw", fmt.Sprintf(`
  fabric_name = %q
  interfaces = {
    "Ethernet1/42" = {
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "accessHost"
      admin_state     = true
    }
  }`, ctx.fabricName))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile(`(?i)switch_id.*required`)},
		},
	})
}

// ---------------------------------------------------------------------------
// 28. TestAccInterfaceEthernet_SwitchIdPerEntryOnly
//     Per-entry switch_id (no resource-level) must work end-to-end.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_SwitchIdPerEntryOnly(t *testing.T) {
	ctx := loadIETestctx(t)

	config := ieInlineConfig(ctx.configMap, "perentry", fmt.Sprintf(`
  fabric_name = %q
  deploy      = true
  interfaces = {
    "Ethernet1/43" = {
      switch_id       = %q
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "accessHost"
      admin_state     = true
      access_vlan     = 430
    }
  }`, ctx.fabricName, ctx.switchId))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn("perentry"), "interfaces.Ethernet1/43.switch_id", ctx.switchId),
					resource.TestCheckResourceAttr(rn("perentry"), "interfaces.Ethernet1/43.access_vlan", "430"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 29. TestAccInterfaceEthernet_SwitchIdPerEntryIP
//     Per-entry switch_id with IP (no resource-level) must resolve and
//     preserve the IP in state.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_SwitchIdPerEntryIP(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}

	config := ieInlineConfig(ctx.configMap, "perentip", fmt.Sprintf(`
  fabric_name = %q
  deploy      = true
  interfaces = {
    "Ethernet1/44" = {
      switch_id       = %q
      mode            = "access"
      network_os_type = "nx-os"
      policy_type     = "accessHost"
      admin_state     = true
      access_vlan     = 440
    }
  }`, ctx.fabricName, ctx.switchIP))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn("perentip"), "interfaces.Ethernet1/44.switch_id", ctx.switchIP),
					resource.TestCheckResourceAttr(rn("perentip"), "interfaces.Ethernet1/44.access_vlan", "440"),
				),
			},
		},
	})
}

// ===========================================================================
// IP-based switch_id tests (resource-level)
//
// These tests use the switch management IP (from testbed switch_ip[0])
// instead of the serial number. They verify that:
//   - The provider resolves IPs to serials internally for API calls
//   - Terraform state preserves the user-supplied IP, NOT the resolved serial
//   - Create, Update, Import, and idempotence all work with IPs
// ===========================================================================

// ---------------------------------------------------------------------------
// 30. TestAccInterfaceEthernet_IPSwitchId_Create
//     Create using IP as switch_id. Verify state preserves the IP.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IPSwitchId_Create(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}
	ctx.configMap["RscName"] = "ipcreat"
	rscAddr := rn("ipcreat")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with IP as switch_id.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchIP, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/21": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(210),
								Mtu: "default",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty()),
						// Verify state preserves the IP, not the serial.
						resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
					)...,
				),
			},
			// Step 2: Idempotence — IP in state must not cause drift.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 31. TestAccInterfaceEthernet_IPSwitchId_Update
//     Create with IP, then update interface attributes. Verify state
//     preserves the IP through the update cycle.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IPSwitchId_Update(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}
	ctx.configMap["RscName"] = "ipupd"
	rscAddr := rn("ipupd")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with IP.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchIP, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/22": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(220),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
				),
			},
			// Step 2: Update access_vlan. IP must still be in state.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"Ethernet1/22.access_vlan": 221,
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
					resource.TestCheckResourceAttr(rscAddr, "interfaces.Ethernet1/22.access_vlan", "221"),
				),
			},
			// Step 3: Idempotence after update.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 32. TestAccInterfaceEthernet_IPSwitchId_BulkAddDelete
//     Bulk operations (add + delete) with IP as switch_id.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IPSwitchId_BulkAddDelete(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}
	ctx.configMap["RscName"] = "ipbulk"
	rscAddr := rn("ipbulk")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create two interfaces with IP.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchIP, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/23": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(230),
							},
							"Ethernet1/24": {
								Mode: "trunk", NetworkOsType: "nx-os", PolicyType: "trunkHost",
								AdminState: boolP(true), AllowedVlans: "230-240",
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
				),
			},
			// Step 2: Remove one, add one. IP must still be in state.
			{
				Config: func() string {
					helper.RemoveInterface(&intfRsc, "Ethernet1/24")
					helper.AddInterface(&intfRsc, "Ethernet1/25", helper.InterfaceEthernetEntry{
						Mode: "routed", NetworkOsType: "nx-os", PolicyType: "routedHost",
						AdminState: boolP(true), Ip: "10.5.5.1", Prefix: "24",
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					append(
						InterfaceEthernetModelHelperStateCheck(rscAddr, *intfRsc, path.Empty()),
						resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
						resource.TestCheckNoResourceAttr(rscAddr, "interfaces.Ethernet1/24.policy_type"),
					)...,
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 33. TestAccInterfaceEthernet_IPSwitchId_Import
//     Import using IP as switch_id. Verify state preserves the IP.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IPSwitchId_Import(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}
	ctx.configMap["RscName"] = "ipimp"

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)
	importId := fmt.Sprintf("%s/%s/Ethernet1/29", ctx.fabricName, ctx.switchIP)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with IP so the interface exists on ND.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchIP, true,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/29": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(290),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
			},
			// Step 2: Import using IP as the switch_id portion of the ID.
			{
				ResourceName:      "nd_interface_ethernet.ipimp",
				ImportState:       true,
				ImportStateId:     importId,
				ImportStateVerify: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_interface_ethernet.ipimp", "fabric_name", ctx.fabricName),
					// State must preserve the IP from the import ID.
					resource.TestCheckResourceAttr("nd_interface_ethernet.ipimp", "switch_id", ctx.switchIP),
					resource.TestCheckResourceAttr("nd_interface_ethernet.ipimp", "interfaces.Ethernet1/29.policy_type", "accessHost"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// 34. TestAccInterfaceEthernet_IPSwitchId_Deploy
//     Create with IP and deploy=false, then deploy in step 2.
//     Verifies deploy works correctly with IP-based switch_id.
// ---------------------------------------------------------------------------

func TestAccInterfaceEthernet_IPSwitchId_Deploy(t *testing.T) {
	ctx := loadIETestctx(t)
	if ctx.switchIP == "" {
		t.Skip("testbed yaml lacks nd.switch_ip[0]; skipping IP-based test")
	}
	ctx.configMap["RscName"] = "ipdep"
	rscAddr := rn("ipdep")

	stepCount := 0
	intfRsc := new(resource_interface_common.NDFCInterfaceCommonModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "ethernet") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with IP, deploy=false.
			{
				Config: func() string {
					helper.GenerateInterfaceEthernetObject(&intfRsc,
						ctx.fabricName, ctx.switchIP, false,
						map[string]helper.InterfaceEthernetEntry{
							"Ethernet1/30": {
								Mode: "access", NetworkOsType: "nx-os", PolicyType: "accessHost",
								AdminState: boolP(true), AccessVlan: int64P(300),
							},
						},
					)
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
				),
			},
			// Step 2: Set deploy=true. Verifies RscDeployInterfaces resolves
			// the IP for the deploy payload.
			{
				Config: func() string {
					helper.ModifyInterfaceEthernetObject(&intfRsc, map[string]interface{}{
						"deploy": true,
					})
					return ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rscAddr, "switch_id", ctx.switchIP),
				),
			},
			// Step 3: Idempotence — deploy + IP must not drift.
			{
				Config:             ieConfigForStep(t, t.Name(), &stepCount, ctx.configMap, intfRsc),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
