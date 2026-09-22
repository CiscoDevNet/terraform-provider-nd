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
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_tenant_domain"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/netascode/go-nd"
)

type tenantAcceptancePayload struct {
	Name string `json:"name"`
}

type tenantDomainAcceptancePayload struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	TenantNames []string `json:"tenantNames"`
}

func newTenantDomainAcceptanceClient(t *testing.T) *nd.Client {
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
		t.Fatalf("failed to create ND client for tenant-domain acceptance test: %s", err.Error())
	}

	return &client
}

func createTenantOutsideTerraform(t *testing.T, client *nd.Client, name string) {
	t.Helper()

	payload, err := json.Marshal(tenantAcceptancePayload{Name: name})
	if err != nil {
		t.Fatalf("failed to marshal tenant %q for test setup: %s", name, err.Error())
	}

	res, err := client.Post("/infra/tenants", string(payload))
	if err != nil {
		t.Fatalf("failed to create tenant %q for test setup: %s %s", name, err.Error(), res.String())
	}
}

func deleteTenantIfExistsOutsideTerraform(t *testing.T, client *nd.Client, name string) {
	t.Helper()

	res, err := client.Delete(fmt.Sprintf("/infra/tenants/%s", url.PathEscape(name)), "")
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to delete tenant %q during test cleanup: %s %s", name, err.Error(), res.String())
	}
}

func createTenantDomainOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	name string,
	description string,
	tenantNames []string,
) {
	t.Helper()

	payload, err := json.Marshal(tenantDomainAcceptancePayload{
		Name:        name,
		Description: description,
		TenantNames: tenantNames,
	})
	if err != nil {
		t.Fatalf("failed to marshal tenant domain %q for out-of-band create: %s", name, err.Error())
	}

	tenantDomainAPI := api.NewTenantDomainAPI(client, ndapi.DefaultFabric)
	res, err := tenantDomainAPI.Post(payload, nil)
	if err != nil {
		t.Fatalf("failed to create tenant domain %q outside Terraform: %s %s", name, err.Error(), res.String())
	}
}

func updateTenantDomainOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	name string,
	description string,
	tenantNames []string,
) {
	t.Helper()

	payload, err := json.Marshal(tenantDomainAcceptancePayload{
		Description: description,
		TenantNames: tenantNames,
	})
	if err != nil {
		t.Fatalf("failed to marshal tenant domain %q for out-of-band update: %s", name, err.Error())
	}

	tenantDomainAPI := api.NewTenantDomainAPI(client, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = name
	res, err := tenantDomainAPI.Put(payload, nil)
	if err != nil {
		t.Fatalf("failed to update tenant domain %q outside Terraform: %s %s", name, err.Error(), res.String())
	}
}

func deleteTenantDomainIfExistsOutsideTerraform(t *testing.T, client *nd.Client, name string) {
	t.Helper()

	tenantDomainAPI := api.NewTenantDomainAPI(client, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = name
	res, err := tenantDomainAPI.Delete()
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to delete tenant domain %q outside Terraform: %s %s", name, err.Error(), res.String())
	}
}

func tenantDomainStateChecks(
	resourceName string,
	tenantDomain resource_tenant_domain.NDFCTenantDomainModel,
) []resource.TestCheckFunc {
	checks := TenantDomainModelHelperStateCheck(resourceName, tenantDomain, path.Empty())
	checks = append(
		checks,
		resource.TestCheckResourceAttr(resourceName, "tenant_names.#", strconv.Itoa(len(tenantDomain.TenantNames))),
	)
	for _, tenantName := range tenantDomain.TenantNames {
		checks = append(checks, resource.TestCheckTypeSetElemAttr(resourceName, "tenant_names.*", tenantName))
	}
	if tenantDomain.Description == "" {
		checks = append(checks, resource.TestCheckNoResourceAttr(resourceName, "description"))
	}

	return checks
}

func TestAccTenantDomainResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	client := newTenantDomainAcceptanceClient(t)
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	tenantNames := []string{
		fmt.Sprintf("TDA_%s", suffix),
		fmt.Sprintf("TDB_%s", suffix),
		fmt.Sprintf("TDC_%s", suffix),
	}
	existingDomainName := fmt.Sprintf("tf_domain_existing_%s", suffix)
	optionalDomainName := fmt.Sprintf("tf_domain_optional_%s", suffix)
	domainName := fmt.Sprintf("tf_domain_%s", suffix)
	importDomainName := fmt.Sprintf("tf_domain_import_%s", suffix)
	missingDomainName := fmt.Sprintf("tf_domain_missing_%s", suffix)
	description := "tenant domain managed by Terraform"
	driftDescription := "tenant domain changed outside Terraform"
	resourceName := "nd_tenant_domain.domain_test"
	x := &map[string]string{
		"RscType":  "nd_tenant_domain",
		"RscName":  "domain_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	tenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	importTFConfig := new(string)
	importTenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	missingTFConfig := new(string)
	missingTenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}
	s8 := &helper.StepInfo{}
	s9 := &helper.StepInfo{}
	s10 := &helper.StepInfo{}
	s11 := &helper.StepInfo{}
	s12 := &helper.StepInfo{}

	cleanupEnabled := false
	createdTenants := make([]string, 0, len(tenantNames))
	defer func() {
		if cleanupEnabled {
			deleteTenantDomainIfExistsOutsideTerraform(t, client, existingDomainName)
			deleteTenantDomainIfExistsOutsideTerraform(t, client, importDomainName)
			deleteTenantDomainIfExistsOutsideTerraform(t, client, domainName)
			deleteTenantDomainIfExistsOutsideTerraform(t, client, optionalDomainName)
			for i := len(createdTenants) - 1; i >= 0; i-- {
				deleteTenantIfExistsOutsideTerraform(t, client, createdTenants[i])
			}
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			cleanupEnabled = true
			for _, tenantName := range tenantNames {
				createTenantOutsideTerraform(t, client, tenantName)
				createdTenants = append(createdTenants, tenantName)
			}
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Report a conflict when the tenant domain already exists remotely")

					helper.GenerateTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name": existingDomainName,
					})
					helper.GetTFConfigWithSingleResource(s1.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					createTenantDomainOutsideTerraform(t, client, existingDomainName, "", []string{})
				},
				ExpectError: regexp.MustCompile(`StatusCode 409`),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Create a tenant domain with only the required name attribute")

					helper.GenerateTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name": optionalDomainName,
					})
					helper.GetTFConfigWithSingleResource(s2.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					deleteTenantDomainIfExistsOutsideTerraform(t, client, existingDomainName)
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
				},
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *tenantDomain)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the tenant domain with the first tenant")

					helper.ModifyTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name":         domainName,
						"tenant_names": tenantNames[:1],
					})
					helper.GetTFConfigWithSingleResource(s3.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *tenantDomain)...,
				),
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Add a description and replace the tenant set")

					helper.ModifyTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name":         domainName,
						"description":  description,
						"tenant_names": tenantNames[1:],
					})
					helper.GetTFConfigWithSingleResource(s4.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s4.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *tenantDomain)...,
				),
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Remove the description and tenant set")

					helper.ModifyTenantDomainObject(&tenantDomain, map[string]interface{}{
						"name": domainName,
					})
					helper.GetTFConfigWithSingleResource(s5.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s5.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s5.Index, s5.Name, s5.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *tenantDomain)...,
				),
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Detect an out-of-band description and tenant set change")

					helper.GetTFConfigWithSingleResource(s6.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s6.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s6.Index, s6.Name, s6.Cfg)
					updateTenantDomainOutsideTerraform(t, client, domainName, driftDescription, tenantNames[2:])
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Recreate the tenant domain after an out-of-band deletion")

					helper.GetTFConfigWithSingleResource(s7.Name, *x, []interface{}{tenantDomain}, &tfConfig)
					s7.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s7.Index, s7.Name, s7.Cfg)
					deleteTenantDomainIfExistsOutsideTerraform(t, client, domainName)
				},
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *tenantDomain)...,
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan: resource.PlanOptions{NoRefresh: true},
		},
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s8.Index = 8
					s8.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the tenant domain for import")

					helper.GenerateTenantDomainObject(&importTenantDomain, map[string]interface{}{
						"name": importDomainName,
					})
					helper.GetTFConfigWithSingleResource(s8.Name, *x, []interface{}{importTenantDomain}, &importTFConfig)
					s8.Cfg = *importTFConfig
					return *importTFConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s8.Index, s8.Name, s8.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					tenantDomainStateChecks(resourceName, *importTenantDomain)...,
				),
			},
			{
				PreConfig: func() {
					s9.Index = 9
					s9.Name = fmt.Sprintf("%s - %s", t.Name(), "Import the tenant domain and verify API readback")
					s9.Cfg = *importTFConfig
					helper.LogStep(t, s9.Index, s9.Name, s9.Cfg)
				},
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        importDomainName,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			{
				Config: func() string {
					s10.Index = 10
					s10.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan after import")

					helper.GetTFConfigWithSingleResource(s10.Name, *x, []interface{}{importTenantDomain}, &importTFConfig)
					s10.Cfg = *importTFConfig
					return *importTFConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s10.Index, s10.Name, s10.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s11.Index = 11
					s11.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy successfully after the imported tenant domain is deleted out of band")

					helper.GetTFConfigWithSingleResource(s11.Name, *x, []interface{}{importTenantDomain}, &importTFConfig)
					s11.Cfg = *importTFConfig
					return *importTFConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s11.Index, s11.Name, s11.Cfg)
					deleteTenantDomainIfExistsOutsideTerraform(t, client, importDomainName)
				},
				Destroy: true,
			},
			{
				PreConfig: func() {
					s12.Index = 12
					s12.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject import of a missing tenant domain")
					helper.GenerateTenantDomainObject(&missingTenantDomain, map[string]interface{}{
						"name": missingDomainName,
					})
					helper.GetTFConfigWithSingleResource(s12.Name, *x, []interface{}{missingTenantDomain}, &missingTFConfig)
					s12.Cfg = *missingTFConfig
					helper.LogStep(t, s12.Index, s12.Name, s12.Cfg)
				},
				ResourceName:  resourceName,
				ImportState:   true,
				ImportStateId: missingDomainName,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Tenant domain\s+%q\s+was\s+not\s+found`, missingDomainName),
				),
			},
		},
	})
}
