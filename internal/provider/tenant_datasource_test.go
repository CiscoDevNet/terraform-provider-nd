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

	"terraform-provider-nd/internal/infra/resource_tenant"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTenantDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	tenantName := fmt.Sprintf("tf_tenant_ds_%s", suffix)
	missingName := tenantName + "_missing"
	fabricName := "ansible-test"
	resourceName := "nd_tenant.tenant_test"
	dataSourceName := "data.nd_tenant.tenant_test"

	associations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		fabricName: {
			TenantPrefix: fmt.Sprintf("tn_ds_%s", suffix),
			LocalName:    fmt.Sprintf("local_tenant_ds_%s", suffix),
			AllowedVlans: []string{"1", "5-10"},
		},
	}

	x := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	tenant := new(resource_tenant.NDFCTenantModel)
	tenantDataSource := &helper.TenantDataSourceTestData{
		RscName:   "tenant_test",
		Name:      tenantName,
		DependsOn: resourceName,
	}
	missingTenantDataSource := &helper.TenantDataSourceTestData{
		RscName: "missing_tenant",
		Name:    missingName,
	}

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "fabric_associations.%", resourceName, "fabric_associations.%"),
	}
	for associationFabricName, association := range associations {
		associationPath := fmt.Sprintf("fabric_associations.%s", associationFabricName)
		matchingChecks = append(
			matchingChecks,
			resource.TestCheckResourceAttrPair(dataSourceName, associationPath+".tenant_prefix", resourceName, associationPath+".tenant_prefix"),
			resource.TestCheckResourceAttrPair(dataSourceName, associationPath+".local_name", resourceName, associationPath+".local_name"),
			resource.TestCheckResourceAttrPair(dataSourceName, associationPath+".allowed_vlans.#", resourceName, associationPath+".allowed_vlans.#"),
		)
		for _, vlan := range association.AllowedVlans {
			matchingChecks = append(
				matchingChecks,
				resource.TestCheckTypeSetElemAttr(dataSourceName, associationPath+".allowed_vlans.*", vlan),
				resource.TestCheckTypeSetElemAttr(resourceName, associationPath+".allowed_vlans.*", vlan),
			)
		}
	}

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
					s1.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Create the tenant used by datasource lookups",
					)

					helper.GenerateTenantObject(&tenant, tenantName, map[string]interface{}{
						"description":         "Tenant datasource acceptance test",
						"fabric_associations": associations,
					})
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						*x,
						[]interface{}{tenant},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						TenantModelHelperStateCheck(resourceName, *tenant, path.Empty()),
						tenantAssociationStateChecks(resourceName, associations)...,
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Read the existing tenant through the datasource",
					)

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{tenant, tenantDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", tenantName),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Match datasource attributes with the configured tenant",
					)

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						*x,
						[]interface{}{tenant, tenantDataSource},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					matchingChecks...,
				),
			},
		},
	})

	missingTFConfig := new(string)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Reject datasource lookup for a missing tenant",
					)

					helper.GetTFConfigWithSingleResource(
						s4.Name,
						*x,
						[]interface{}{missingTenantDataSource},
						&missingTFConfig,
					)
					s4.Cfg = *missingTFConfig
					return *missingTFConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not read nd tenant with name\s+%q:\s+resource\s+not\s+found`, missingName),
				),
			},
		},
	})
}
