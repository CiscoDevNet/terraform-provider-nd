// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	nd "github.com/netascode/go-nd"
)

const fabricGroupVxlanAciTestAnycastGatewayMAC = "0022.bdf8.19ff"

func newFabricGroupVxlanAciAcceptanceClient(t *testing.T) *nd.Client {
	t.Helper()

	cfg := helper.GetConfig("global")
	client, err := nd.NewClient(
		cfg.ND.URL,
		"/api/v1",
		cfg.ND.User,
		cfg.ND.Password,
		"",
		cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("failed to create ND client for VXLAN-ACI Fabric Group acceptance test: %s", err.Error())
	}

	return &client
}

func fabricGroupVxlanAciAcceptancePayload(
	t *testing.T,
	fabricName string,
	anycastGatewayMAC string,
	managementOverrides map[string]interface{},
) []byte {
	t.Helper()

	management := map[string]interface{}{"type": "vxlanAci"}
	if anycastGatewayMAC != "" {
		management["anycastGatewayMac"] = anycastGatewayMAC
	}
	for key, value := range managementOverrides {
		management[key] = value
	}

	payload, err := json.Marshal(map[string]interface{}{
		"category":   "fabricGroup",
		"name":       fabricName,
		"management": management,
	})
	if err != nil {
		t.Fatalf("failed to marshal Fabric Group %q for acceptance setup: %s", fabricName, err.Error())
	}

	return payload
}

func waitForFabricGroupVxlanAciOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	fabricName string,
	present bool,
) {
	t.Helper()

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	deadline := time.Now().Add(2 * time.Minute)
	for {
		_, err := fabricAPI.Get()
		if present && err == nil {
			return
		}
		if !present && err != nil && strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			t.Fatalf("failed while waiting for Fabric Group %q: %s", fabricName, err.Error())
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for Fabric Group %q present=%t", fabricName, present)
		}
		time.Sleep(2 * time.Second)
	}
}

func updateFabricGroupVxlanAciOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	fabricName string,
	anycastGatewayMAC string,
	managementOverrides map[string]interface{},
) {
	t.Helper()

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	// Preserve readable settings so this setup changes only the named drift
	// attributes rather than resetting unrelated fields through PUT defaults.
	response, err := fabricAPI.Get()
	if err != nil {
		t.Fatalf("failed to read Fabric Group %q before out-of-band update: %s", fabricName, err.Error())
	}
	var current struct {
		Management map[string]interface{} `json:"management"`
	}
	if err := json.Unmarshal(response, &current); err != nil {
		t.Fatalf("failed to decode Fabric Group %q before out-of-band update: %s", fabricName, err.Error())
	}
	if current.Management == nil {
		t.Fatalf("Fabric Group %q response has no management settings", fabricName)
	}
	delete(current.Management, "securityGroupStatus") // Read-only status.
	for key, value := range managementOverrides {
		current.Management[key] = value
	}
	payload := fabricGroupVxlanAciAcceptancePayload(
		t,
		fabricName,
		anycastGatewayMAC,
		current.Management,
	)
	_, err = fabricAPI.Put(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to update Fabric Group %q outside Terraform: %s", fabricName, err.Error())
	}

	deadline := time.Now().Add(2 * time.Minute)
	for {
		response, getErr := fabricAPI.Get()
		if getErr == nil {
			var body struct {
				Management map[string]interface{} `json:"management"`
			}
			if unmarshalErr := json.Unmarshal(response, &body); unmarshalErr != nil {
				t.Fatalf("failed to decode Fabric Group %q while waiting for drift update: %s", fabricName, unmarshalErr.Error())
			}
			matches := true
			for key, expected := range managementOverrides {
				if fmt.Sprint(body.Management[key]) != fmt.Sprint(expected) {
					matches = false
					break
				}
			}
			if matches {
				return
			}
		} else if !strings.Contains(getErr.Error(), "StatusCode 404") {
			t.Fatalf("failed to read Fabric Group %q after out-of-band update: %s", fabricName, getErr.Error())
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for the out-of-band Fabric Group %q update", fabricName)
		}
		time.Sleep(2 * time.Second)
	}
}

func deleteFabricGroupVxlanAciIfExistsOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	fabricName string,
) {
	t.Helper()

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	res, err := fabricAPI.Delete()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		t.Fatalf("failed to delete Fabric Group %q outside Terraform: %s %s", fabricName, err.Error(), res.String())
	}
	waitForFabricGroupVxlanAciOutsideTerraform(t, client, fabricName, false)
}

func fabricGroupVxlanAciStateChecks(resourceName string, fabricGroup resource_fabric_common.NDFCFabricCommonModel) []resource.TestCheckFunc {
	checks := FabricGroupVxlanAciModelHelperStateCheck(resourceName, fabricGroup, path.Empty())
	checks = append(checks,
		resource.TestCheckResourceAttr(resourceName, "id", fabricGroup.FabricName),
		resource.TestCheckResourceAttr(resourceName, "route_server_collection.#", fmt.Sprintf("%d", len(fabricGroup.Management.RouteServerCollection))),
	)
	for index, routeServer := range fabricGroup.Management.RouteServerCollection {
		checks = append(checks,
			resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("route_server_collection.%d.route_server_ip", index), routeServer.RouteServerIp),
			resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("route_server_collection.%d.route_server_asn", index), routeServer.RouteServerAsn),
		)
	}
	return checks
}

func fabricGroupVxlanAciConfigArgs(resourceName string) map[string]string {
	cfg := helper.GetConfig("global")
	return map[string]string{
		"RscType":  "nd_fabric_group_vxlan_aci",
		"RscName":  resourceName,
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
}

func TestAccFabricGroupVxlanAciResourceMinimal(t *testing.T) {
	const resourceAddress = "nd_fabric_group_vxlan_aci.fabric_group_test"
	fabricName := fmt.Sprintf(
		"tf_fg_vxlan_aci_%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlpha),
	)
	configArgs := fabricGroupVxlanAciConfigArgs("fabric_group_test")
	tfConfig := new(string)
	fabricGroup := new(resource_fabric_common.NDFCFabricCommonModel)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Create a VXLAN-ACI Fabric Group with only its required attributes", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(
						&fabricGroup,
						fabricName,
						fabricGroupVxlanAciTestAnycastGatewayMAC,
					)
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Verify an empty plan after creating the minimal VXLAN-ACI Fabric Group", t.Name())
					helper.GetTFConfigWithSingleResource(
						s2.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				PlanOnly:  true,
			},
			{
				PreConfig: func() {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - Import the minimal VXLAN-ACI Fabric Group and verify API readback", t.Name())
					s3.Cfg = *tfConfig
					helper.LogStep(t, s3.Index, s3.Name, s3.Cfg)
				},
				ResourceName:                         resourceAddress,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        fabricName,
				ImportStateVerifyIdentifierAttribute: "fabric_name",
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - Verify an empty plan after importing the minimal VXLAN-ACI Fabric Group", t.Name())
					helper.GetTFConfigWithSingleResource(
						s4.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				PlanOnly:  true,
			},
		},
	})
}

func TestAccFabricGroupVxlanAciResourceCRUD(t *testing.T) {
	const resourceAddress = "nd_fabric_group_vxlan_aci.fabric_group_test"
	var client *nd.Client
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	existingName := fmt.Sprintf("tf_fg_existing_vxlan_aci_%s", suffix)
	fabricName := fmt.Sprintf("tf_fg_full_vxlan_aci_%s", suffix)
	replacementName := fmt.Sprintf("tf_fg_replacement_vxlan_aci_%s", suffix)
	missingName := fmt.Sprintf("tf_fg_missing_vxlan_aci_%s", suffix)
	configArgs := fabricGroupVxlanAciConfigArgs("fabric_group_test")
	tfConfig := new(string)
	fabricGroup := new(resource_fabric_common.NDFCFabricCommonModel)
	cleanupEnabled := false
	t.Cleanup(func() {
		if !cleanupEnabled {
			return
		}
		deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, existingName)
		deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, fabricName)
		deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, replacementName)
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}
	s8 := &helper.StepInfo{}
	s9 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			client = newFabricGroupVxlanAciAcceptanceClient(t)
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Report a conflict when the VXLAN-ACI Fabric Group already exists remotely", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(
						&fabricGroup,
						existingName,
						fabricGroupVxlanAciTestAnycastGatewayMAC,
					)
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, existingName)
					fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
					payload := fabricGroupVxlanAciAcceptancePayload(t, existingName, fabricGroupVxlanAciTestAnycastGatewayMAC, nil)
					if _, err := fabricAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true}); err != nil {
						t.Fatalf("failed to create VXLAN-ACI Fabric Group %q outside Terraform: %s", existingName, err.Error())
					}
					waitForFabricGroupVxlanAciOutsideTerraform(t, client, existingName, true)
				},
				ExpectError: regexp.MustCompile(`(?is)Error Writing Fabric Group.*StatusCode\s+(403|409)\b`),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Create the VXLAN-ACI Fabric Group with its supported shared settings", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(
						&fabricGroup,
						fabricName,
						fabricGroupVxlanAciTestAnycastGatewayMAC,
					)
					values := map[string]interface{}{
						"l2_vni_range":                               "30000-49000",
						"l3_vni_range":                               "50000-59000",
						"scheduled_backup":                           false,
						"security_group_tag":                         "strict",
						"security_group_tag_prefix":                  "TF_",
						"security_group_tag_id_range":                "10000-14000",
						"security_group_tag_preprovision":            false,
						"security_group_tag_mac_segmentation":        false,
						"multisite_overlay_inter_connect_type":       "directPeering",
						"route_server_redistribute_direct_route_map": false,
						"route_server_routing_tag":                   int64(54321),
						"auto_multisite_underlay_inter_connect":      false,
						"multisite_delay_restore":                    int64(300),
						"multisite_inter_connect_bgp_authentication": false,
						"multisite_loopback_id":                      int64(100),
						"border_gateway_routing_tag":                 int64(54321),
						"multisite_loopback_ip_range":                "10.10.0.0/24",
						"multisite_underlay_subnet_range":            "10.10.1.0/24",
						"multisite_underlay_subnet_target_mask":      int64(30),
						"enable_ms_overlay_ifc_bgp_desc":             true,
					}

					helper.ModifyFabricGroupVxlanAciObject(&fabricGroup, values)
					helper.GetTFConfigWithSingleResource(
						s2.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, existingName)
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
				},
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - Configure route-server connectivity on the VXLAN-ACI Fabric Group", t.Name())
					helper.ModifyFabricGroupVxlanAciObject(&fabricGroup, map[string]interface{}{
						"multisite_overlay_inter_connect_type": "routeServer",
						"route_server_collection": []resource_fabric_common.NDFCRouteServerCollectionValue{
							{RouteServerIp: "192.0.2.10", RouteServerAsn: "65001"},
						},
						"route_server_redistribute_direct_route_map": true,
					})
					helper.GetTFConfigWithSingleResource(
						s3.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - Clear the route-server collection on the VXLAN-ACI Fabric Group", t.Name())
					helper.ModifyFabricGroupVxlanAciObject(&fabricGroup, map[string]interface{}{
						"multisite_overlay_inter_connect_type":       "directPeering",
						"route_server_collection":                    []resource_fabric_common.NDFCRouteServerCollectionValue{},
						"route_server_redistribute_direct_route_map": false,
					})
					helper.GetTFConfigWithSingleResource(
						s4.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResourceWithEmptyRouteServers(fabricGroup)},
						&tfConfig,
					)
					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - Detect an out-of-band VNI range change on the VXLAN-ACI Fabric Group", t.Name())
					helper.GetTFConfigWithSingleResource(
						s5.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResourceWithEmptyRouteServers(fabricGroup)},
						&tfConfig,
					)
					s5.Cfg = *tfConfig
					return s5.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s5.Index, s5.Name, s5.Cfg)
					updateFabricGroupVxlanAciOutsideTerraform(
						t,
						client,
						fabricName,
						fabricGroupVxlanAciTestAnycastGatewayMAC,
						map[string]interface{}{
							"l2VniRange":       "31000-49000",
							"l3VniRange":       "50000-59000",
							"securityGroupTag": "strict",
						},
					)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - Recreate the VXLAN-ACI Fabric Group after an out-of-band deletion", t.Name())
					helper.GetTFConfigWithSingleResource(
						s6.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResourceWithEmptyRouteServers(fabricGroup)},
						&tfConfig,
					)
					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s6.Index, s6.Name, s6.Cfg)
					deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, fabricName)
				},
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - Replace the VXLAN-ACI Fabric Group when its name changes", t.Name())
					helper.ModifyFabricGroupVxlanAciObject(&fabricGroup, map[string]interface{}{
						"fabric_name": replacementName,
					})
					helper.GetTFConfigWithSingleResource(
						s7.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResourceWithEmptyRouteServers(fabricGroup)},
						&tfConfig,
					)
					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				PreConfig: func() {
					s8.Index = 8
					s8.Name = fmt.Sprintf("%s - Reject import of a missing VXLAN-ACI Fabric Group", t.Name())
					s8.Cfg = *tfConfig
					helper.LogStep(t, s8.Index, s8.Name, s8.Cfg)
				},
				ResourceName:  resourceAddress,
				ImportState:   true,
				ImportStateId: missingName,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`(?s)Error Importing Fabric Group.*Fabric\s+group\s+"%s"\s+was\s+not\s+found\.`, regexp.QuoteMeta(missingName)),
				),
			},
			{
				PreConfig: func() {
					s9.Index = 9
					s9.Name = fmt.Sprintf("%s - Reject an empty import ID for the VXLAN-ACI Fabric Group", t.Name())
					s9.Cfg = *tfConfig
					helper.LogStep(t, s9.Index, s9.Name, s9.Cfg)
				},
				ResourceName:  resourceAddress,
				ImportState:   true,
				ImportStateId: " ",
				ExpectError:   regexp.MustCompile(`Invalid Fabric Group Import ID`),
			},
		},
	})
}

func TestAccFabricGroupVxlanAciResourceMissingDelete(t *testing.T) {
	const resourceAddress = "nd_fabric_group_vxlan_aci.fabric_group_test"
	var client *nd.Client
	fabricName := fmt.Sprintf(
		"tf_fg_delete_vxlan_aci_%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlpha),
	)
	configArgs := fabricGroupVxlanAciConfigArgs("fabric_group_test")
	tfConfig := new(string)
	fabricGroup := new(resource_fabric_common.NDFCFabricCommonModel)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			client = newFabricGroupVxlanAciAcceptanceClient(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan: resource.PlanOptions{NoRefresh: true},
		},
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Create the VXLAN-ACI Fabric Group for stale-state deletion", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(
						&fabricGroup,
						fabricName,
						fabricGroupVxlanAciTestAnycastGatewayMAC,
					)
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						configArgs,
						[]interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					fabricGroupVxlanAciStateChecks(resourceAddress, *fabricGroup)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Destroy the VXLAN-ACI Fabric Group successfully after it is deleted remotely", t.Name())
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					deleteFabricGroupVxlanAciIfExistsOutsideTerraform(t, client, fabricName)
				},
				Destroy: true,
			},
		},
	})
}

func TestAccFabricGroupVxlanAciResourceValidation(t *testing.T) {
	configArgs := fabricGroupVxlanAciConfigArgs("fabric_group_test")
	tfConfig := new(string)
	fabricGroup := new(resource_fabric_common.NDFCFabricCommonModel)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Reject an invalid anycast gateway MAC address", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(&fabricGroup, "tf_fg_invalid_mac", "invalid-mac")
					helper.GetTFConfigWithSingleResource(s1.Name, configArgs, []interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)}, &tfConfig)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("must be a MAC address in xxxx.xxxx.xxxx format")),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Reject SGT off mode for a VXLAN-ACI Fabric Group", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(&fabricGroup, "tf_fg_invalid_sgt", fabricGroupVxlanAciTestAnycastGatewayMAC)
					helper.ModifyFabricGroupVxlanAciObject(&fabricGroup, map[string]interface{}{"security_group_tag": "off"})
					helper.GetTFConfigWithSingleResource(s2.Name, configArgs, []interface{}{helper.FabricGroupVxlanAciResource(fabricGroup)}, &tfConfig)
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`(?s)security_group_tag.*must be one of`),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - Reject a VXLAN-only attribute on a VXLAN-ACI Fabric Group", t.Name())
					helper.GenerateFabricGroupVxlanAciObject(&fabricGroup, "tf_fg_invalid_attribute", fabricGroupVxlanAciTestAnycastGatewayMAC)
					// This field is absent from the dedicated resource schema and
					// added only to this negative case through the shared renderer.
					invalidResource := helper.FabricGroupVxlanAciResource(fabricGroup)
					invalidResource.AdditionalAttributes = `vrf_template = "Default_VRF_Universal"`
					helper.GetTFConfigWithSingleResource(s3.Name, configArgs, []interface{}{invalidResource}, &tfConfig)
					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`(?s)Unsupported argument.*vrf_template`),
			},
		},
	})
}

