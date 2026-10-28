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
	"strconv"
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_policy"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const policyResourceName = "nd_policy.policy_test"

type policyTestctx struct {
	configMap    map[string]string
	fabricName   string
	serials      []string
	templateName string
	parameters   map[string]string
}

func loadPolicyTestctx(t *testing.T) *policyTestctx {
	t.Helper()

	cfg := helper.GetConfig("global")
	ctx := &policyTestctx{
		configMap: map[string]string{
			"RscType":  "nd_policy",
			"RscName":  "policy_test",
			"User":     cfg.ND.User,
			"Password": cfg.ND.Password,
			"Host":     cfg.ND.URL,
			"Insecure": cfg.ND.Insecure,
		},
		fabricName:   cfg.ND.Fabric,
		templateName: "feature_enable",
		parameters: map[string]string{
			"featureName": "lacp",
		},
	}

	for _, sw := range cfg.ND.Inventory.Switches {
		if strings.EqualFold(sw.Role, "leaf") && sw.Serial != "" {
			ctx.serials = append(ctx.serials, sw.Serial)
		}
	}
	if len(ctx.serials) == 0 {
		ctx.serials = append(ctx.serials, cfg.ND.Switches...)
	}

	return ctx
}

func requirePolicyFabric(t *testing.T, ctx *policyTestctx) {
	t.Helper()
	if ctx.fabricName == "" {
		t.Skip("Need nd.fabric in the testbed for policy acceptance tests")
	}
}

func requirePolicySerials(t *testing.T, ctx *policyTestctx, count int) {
	t.Helper()
	requirePolicyFabric(t, ctx)
	if len(ctx.serials) < count {
		t.Skipf("Need at least %d switches in the testbed for this policy acceptance test", count)
	}
}

func policyConfigForStep(t *testing.T, testName string, stepCount *int, cfg map[string]string, model *resource_policy.NDFCPolicyModel) string {
	t.Helper()
	return policyConfigForResources(t, testName, stepCount, cfg, []string{cfg["RscName"]}, []*resource_policy.NDFCPolicyModel{model}, "")
}

func policyConfigForResources(
	t *testing.T,
	testName string,
	stepCount *int,
	cfg map[string]string,
	names []string,
	models []*resource_policy.NDFCPolicyModel,
	dependsOn string,
) string {
	t.Helper()

	*stepCount++
	stepCfg := make(map[string]string, len(cfg)+2)
	for k, v := range cfg {
		stepCfg[k] = v
	}
	stepCfg["RscName"] = strings.Join(names, ",")
	if dependsOn != "" {
		stepCfg["DependsOn"] = dependsOn
	}

	rscs := make([]interface{}, 0, len(models))
	for _, model := range models {
		rscs = append(rscs, model)
	}

	tfConfig := new(string)
	stepName := fmt.Sprintf("%s_%d", testName, *stepCount)
	helper.GetTFConfigWithSingleResource(stepName, stepCfg, rscs, &tfConfig)
	return *tfConfig
}

func policyStateChecks(rscName string, model resource_policy.NDFCPolicyModel) []resource.TestCheckFunc {
	checks := PolicyModelHelperStateCheck(rscName, model, path.Empty())
	checks = append(checks,
		resource.TestCheckResourceAttrSet(rscName, "policy_id"),
		resource.TestCheckResourceAttrSet(rscName, "fabric_name"),
		resource.TestCheckResourceAttr(rscName, "serial_numbers.#", strconv.Itoa(len(model.SerialNumbers))),
	)
	for _, serial := range model.SerialNumbers {
		checks = append(checks, resource.TestCheckTypeSetElemAttr(rscName, "serial_numbers.*", serial))
	}
	for key, value := range model.PolicyParameters {
		checks = append(checks, resource.TestCheckResourceAttr(rscName, "policy_parameters."+key, value))
	}
	return checks
}

func policyImportIDFromState(rscName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[rscName]
		if !ok {
			return "", fmt.Errorf("resource %s not found in state", rscName)
		}
		policyID := rs.Primary.Attributes["policy_id"]
		if policyID == "" {
			return "", fmt.Errorf("policy_id missing from resource %s", rscName)
		}
		return policyID, nil
	}
}

func TestAccPolicyResourceCreateRead(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						ctx.templateName,
						ctx.parameters,
						false,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				Config: policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc),
				Check:  resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
		},
	})
}

func TestAccPolicyResourceUpdate(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						ctx.templateName,
						ctx.parameters,
						false,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				Config: func() string {
					helper.ModifyPolicyObject(&policyRsc, map[string]interface{}{
						"description": "Updated LACP policy",
						"priority":    50,
					})
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
		},
	})
}

func TestAccPolicyResourceGroupCreateRead(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 2)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0], ctx.serials[1]},
						ctx.templateName,
						ctx.parameters,
						true,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				Config: func() string {
					helper.ModifyPolicyObject(&policyRsc, map[string]interface{}{
						"serial_numbers": []string{ctx.serials[0]},
						"description":    "Updated policy group membership",
					})
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
		},
	})
}

func TestAccPolicyResourceImport(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						ctx.templateName,
						ctx.parameters,
						false,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				ResourceName:                         policyResourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateIdFunc:                    policyImportIDFromState(policyResourceName),
				ImportStateVerifyIdentifierAttribute: "policy_id",
				ImportStateVerifyIgnore:              []string{"generated_config"},
			},
		},
	})
}

func TestAccPolicyResourceImportDeploy(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						ctx.templateName,
						ctx.parameters,
						false,
						true,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				ResourceName:                         policyResourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateIdFunc:                    policyImportIDFromState(policyResourceName),
				ImportStateVerifyIdentifierAttribute: "policy_id",
				ImportStateVerifyIgnore:              []string{"generated_config", "deploy"},
			},
		},
	})
}

func TestAccPolicyResourceInvalidSerialCount(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicyFabric(t, ctx)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{"SWITCHONE", "SWITCHTWO"},
						ctx.templateName,
						ctx.parameters,
						false,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				ExpectError: regexp.MustCompile(`serial_numbers must contain exactly one switch`),
			},
		},
	})
}

func TestAccPolicyResourceInvalidEntityName(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicyFabric(t, ctx)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{"SWITCHONE"},
						ctx.templateName,
						ctx.parameters,
						false,
						false,
					)
					helper.ModifyPolicyObject(&policyRsc, map[string]interface{}{
						"entity_name": "leaf1",
					})
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				ExpectError: regexp.MustCompile(`entity_name must be SWITCH when entity_type is switch`),
			},
		},
	})
}

func TestAccPolicyResourceIPSlaComplex(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	slaSender := new(resource_policy.NDFCPolicyModel)
	ipslaTrack := new(resource_policy.NDFCPolicyModel)
	createParams := map[string]string{
		"TRACK_ID":       "421",
		"SLA_ID":         "422",
		"PROTOCOL":       "icmp",
		"DESTINATION_IP": "192.0.2.1",
		"THRESHOLD":      "5000",
		"FREQUENCY":      "60",
		"DELAY_DOWN":     "5",
		"DELAY_UP":       "5",
		"TIMEOUT":        "5000",
	}
	updateParams := map[string]string{
		"TRACK_ID":       "421",
		"SLA_ID":         "422",
		"PROTOCOL":       "icmp",
		"DESTINATION_IP": "192.0.2.1",
		"THRESHOLD":      "8000",
		"FREQUENCY":      "90",
		"DELAY_DOWN":     "10",
		"DELAY_UP":       "10",
		"TIMEOUT":        "8000",
	}
	const (
		slaSenderName = "sla_sender"
		ipslaName     = "ipsla_track"
	)
	resourceNames := []string{slaSenderName, ipslaName}
	dependsOn := ";nd_policy." + slaSenderName

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&slaSender,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						"feature_sla_sender",
						nil,
						false,
						true,
					)
					helper.ModifyPolicyObject(&slaSender, map[string]interface{}{
						"description": "Enable IPSLA sender",
					})
					helper.GeneratePolicyObject(
						&ipslaTrack,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						"service_track",
						createParams,
						false,
						true,
					)
					helper.ModifyPolicyObject(&ipslaTrack, map[string]interface{}{
						"description": "Complex IPSLA ICMP track",
					})
					return policyConfigForResources(
						t, t.Name(), &stepCount, ctx.configMap,
						resourceNames,
						[]*resource_policy.NDFCPolicyModel{slaSender, ipslaTrack},
						dependsOn,
					)
				}(),
				Check: resource.ComposeTestCheckFunc(append(
					policyStateChecks("nd_policy."+slaSenderName, *slaSender),
					append(
						policyStateChecks("nd_policy."+ipslaName, *ipslaTrack),
						resource.TestCheckResourceAttrSet("nd_policy."+ipslaName, "generated_config"),
					)...,
				)...),
			},
			{
				Config: func() string {
					helper.ModifyPolicyObject(&ipslaTrack, map[string]interface{}{
						"description":       "Updated complex IPSLA ICMP track",
						"policy_parameters": updateParams,
						"deploy":            false,
					})
					return policyConfigForResources(
						t, t.Name(), &stepCount, ctx.configMap,
						resourceNames,
						[]*resource_policy.NDFCPolicyModel{slaSender, ipslaTrack},
						dependsOn,
					)
				}(),
				Check: resource.ComposeTestCheckFunc(append(
					policyStateChecks("nd_policy."+slaSenderName, *slaSender),
					append(
						policyStateChecks("nd_policy."+ipslaName, *ipslaTrack),
						resource.TestCheckResourceAttrSet("nd_policy."+ipslaName, "generated_config"),
					)...,
				)...),
			},
		},
	})
}

func TestAccPolicyResourceServiceTrack(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)
	createParams := map[string]string{
		"TRACK_ID":       "401",
		"SLA_ID":         "402",
		"PROTOCOL":       "icmp",
		"DESTINATION_IP": "192.0.2.1",
	}
	updateParams := map[string]string{
		"TRACK_ID":       "401",
		"SLA_ID":         "402",
		"PROTOCOL":       "icmp",
		"DESTINATION_IP": "192.0.2.2",
		"THRESHOLD":      "5000",
		"FREQUENCY":      "60",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						"service_track",
						createParams,
						false,
						false,
					)
					helper.ModifyPolicyObject(&policyRsc, map[string]interface{}{
						"description": "IPSLA ICMP track",
					})
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
			{
				Config: func() string {
					helper.ModifyPolicyObject(&policyRsc, map[string]interface{}{
						"description":       "Updated IPSLA ICMP track",
						"policy_parameters": updateParams,
					})
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				Check: resource.ComposeTestCheckFunc(policyStateChecks(policyResourceName, *policyRsc)...),
			},
		},
	})
}

func TestAccPolicyResourceInvalidTemplateParameter(t *testing.T) {
	ctx := loadPolicyTestctx(t)
	requirePolicySerials(t, ctx, 1)

	stepCount := 0
	policyRsc := new(resource_policy.NDFCPolicyModel)
	badParameters := map[string]string{
		"featureName":    "lacp",
		"NOT_A_TEMPLATE": "x",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GeneratePolicyObject(
						&policyRsc,
						ctx.fabricName,
						[]string{ctx.serials[0]},
						ctx.templateName,
						badParameters,
						false,
						false,
					)
					return policyConfigForStep(t, t.Name(), &stepCount, ctx.configMap, policyRsc)
				}(),
				ExpectError: regexp.MustCompile(`policy_parameters contains keys not defined by template`),
			},
		},
	})
}
