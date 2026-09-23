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
	"testing"

	"terraform-provider-nd/internal/infra/resource_tenant"
	"terraform-provider-nd/internal/infra/resource_tenant_domain"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTenantDomainDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	domainName := fmt.Sprintf("tf_domain_ds_%s", suffix)
	missingDomainName := domainName + "_missing"
	description := "tenant domain datasource acceptance test"
	tenantNames := []string{
		fmt.Sprintf("TDA_DS_%s", suffix),
		fmt.Sprintf("TDB_DS_%s", suffix),
	}

	tenantAResourceName := "nd_tenant.tenant_a"
	tenantBResourceName := "nd_tenant.tenant_b"
	resourceName := "nd_tenant_domain.domain_test"
	dataSourceName := "data.nd_tenant_domain.domain_test"

	x := &map[string]string{
		"RscType":  "nd_tenant_domain",
		"RscName":  "tenant_a,tenant_b,domain_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	tenantA := new(resource_tenant.NDFCTenantModel)
	tenantB := new(resource_tenant.NDFCTenantModel)
	tenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	tenantDomainResource := &helper.TenantDomainTestResource{
		Model:     tenantDomain,
		DependsOn: tenantAResourceName + ", " + tenantBResourceName,
	}
	tenantDomainDataSource := &helper.TenantDomainDataSourceTestData{
		RscName:   "domain_test",
		Name:      domainName,
		DependsOn: resourceName,
	}
	missingTenantDomainDataSource := &helper.TenantDomainDataSourceTestData{
		RscName: "missing_domain",
		Name:    missingDomainName,
	}

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "tenant_names.#", resourceName, "tenant_names.#"),
	}
	for _, tenantName := range tenantNames {
		matchingChecks = append(
			matchingChecks,
			resource.TestCheckTypeSetElemAttr(dataSourceName, "tenant_names.*", tenantName),
			resource.TestCheckTypeSetElemAttr(resourceName, "tenant_names.*", tenantName),
		)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the prerequisite tenants and tenant domain used by datasource lookups")

					helper.GenerateTenantObject(&tenantA, tenantNames[0], nil)
					helper.GenerateTenantObject(&tenantB, tenantNames[1], nil)
					helper.GenerateTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name":         domainName,
						"description":  description,
						"tenant_names": tenantNames,
					})
					tenantDomainResource.Model = tenantDomain
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						*x,
						[]interface{}{tenantA, tenantB, tenantDomainResource},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							append(
								TenantModelHelperStateCheck(tenantAResourceName, *tenantA, path.Empty()),
								TenantModelHelperStateCheck(tenantBResourceName, *tenantB, path.Empty())...,
							),
							TenantDomainModelHelperStateCheck(resourceName, *tenantDomain, path.Empty())...,
						),
						resource.TestCheckResourceAttr(resourceName, "tenant_names.#", strconv.Itoa(len(tenantNames))),
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Read the existing tenant domain through the datasource")

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{tenantA, tenantB, tenantDomainResource, tenantDomainDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", domainName),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Match datasource attributes with the configured tenant domain")

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						*x,
						[]interface{}{tenantA, tenantB, tenantDomainResource, tenantDomainDataSource},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check:     resource.ComposeTestCheckFunc(matchingChecks...),
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
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject datasource lookup for a missing tenant domain")

					helper.GetTFConfigWithSingleResource(
						s4.Name,
						*x,
						[]interface{}{missingTenantDomainDataSource},
						&missingTFConfig,
					)
					s4.Cfg = *missingTFConfig
					return *missingTFConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not read nd tenant domain with name\s+%q:\s+resource\s+not\s+found`, missingDomainName),
				),
			},
		},
	})
}
