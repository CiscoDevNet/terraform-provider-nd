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
	"slices"
	"strings"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_tenant"
	manageapi "terraform-provider-nd/internal/manage/api"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	nd "github.com/netascode/go-nd"
	"github.com/tidwall/gjson"
)

const (
	tenantFabricOne = "tf_apic1"
	// tenantFabricOne = "ansible_test"
	tenantFabricTwo = "ansible_test_2"
)

type tenantFabricAssociationTestPayload struct {
	Items []tenantFabricAssociationTestItem `json:"items"`
}

type tenantFabricAssociationTestListResponse struct {
	TenantFabricAssociations []tenantFabricAssociationTestItem `json:"tenantFabricAssociations"`
}

type tenantFabricAssociationTestItem struct {
	AllowedVlans []string `json:"allowedVlans,omitempty"`
	Associate    bool     `json:"associate"`
	FabricName   string   `json:"fabricName"`
	LocalName    string   `json:"localName,omitempty"`
	TenantName   string   `json:"tenantName"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
}

// TestAccTenantResourceCRUD exercises tenant create/update/import cleanup and
// the manage-side fabric association workflow.
func TestAccTenantResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")

	tenant1Suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	tenant2Suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)

	tenant1Name := fmt.Sprintf("tf_tenant1_%s", tenant1Suffix)
	tenant2Name := fmt.Sprintf("tf_tenant2_%s", tenant2Suffix)
	tenant1LocalName := fmt.Sprintf("local_tenant1_%s", tenant1Suffix)
	tenant1UpdatedLocalName := fmt.Sprintf("local_tenant1_updated_%s", tenant1Suffix)
	tenant2LocalName := fmt.Sprintf("local_tenant2_%s", tenant2Suffix)
	tenant1Prefix := fmt.Sprintf("tn_tenant1_%s", tenant1Suffix)
	tenant1UpdatedPrefix := fmt.Sprintf("tn_t1upd_%s", tenant1Suffix)

	allowedVlansCreate := []string{"1", "5-10"}
	allowedVlansUpdate := []string{"5-10", "11-20", "30"}

	testCases := []map[string]string{
		{"name": "create_tenant1_required", "purpose": "create tenant1 with only required attributes"},
		{"name": "create_tenant2_optional", "purpose": "create tenant2 with description and fabric association optional values"},
		{"name": "update_tenant1_add_associations", "purpose": "add tenant1 description and two fabric associations"},
		{"name": "update_tenant1_remove_and_update_associations", "purpose": "remove tenant1 description, remove one association, and update the second association"},
		{"name": "update_tenant1_remove_association_optional_values", "purpose": "remove tenant1 association local_name and allowed_vlans while keeping tenant_prefix"},
		{"name": "update_tenant1_change_tenant_prefix_error", "purpose": "verify the API rejects tenant_prefix mutation without delete/recreate"},
		{"name": "update_tenant1_remove_all_associations", "purpose": "delete every tenant1 fabric association"},
		{"name": "delete_tenant1", "purpose": "delete tenant1 by removing it from Terraform config"},
		{"name": "import_tenant2", "purpose": "import tenant2 by name and verify API-returned values"},
		{"name": "import_missing_tenant1", "purpose": "verify import reports not found for the deleted tenant1"},
		{"name": "clear_tenant2_after_out_of_band_delete", "purpose": "destroy tenant2 after an outside-Terraform delete so the resource handles tenant not found"},
	}
	for _, tc := range testCases {
		t.Logf("tenant test case %s: %s", tc["name"], tc["purpose"])
	}

	xTenant1 := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_one",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
	xTenant2 := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_two",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
	xBothTenants := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_one,tenant_two",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	tenant1Rsc := new(resource_tenant.NDFCTenantModel)
	tenant2Rsc := new(resource_tenant.NDFCTenantModel)

	tenant1RequiredAssociation := resource_tenant.NDFCFabricAssociationsValue{
		FabricName: tenantFabricOne,
	}
	tenant1FullAssociationCreate := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: tenant1Prefix,
		LocalName:    tenant1LocalName,
		AllowedVlans: allowedVlansCreate,
	}
	tenant1FullAssociationUpdate := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: tenant1Prefix,
		LocalName:    tenant1UpdatedLocalName,
		AllowedVlans: allowedVlansUpdate,
	}
	tenant1PrefixOnlyAssociation := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: tenant1Prefix,
	}
	tenant1InvalidPrefixAssociation := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: tenant1UpdatedPrefix,
	}
	tenant2Association := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricOne,
		LocalName:    tenant2LocalName,
		AllowedVlans: allowedVlansCreate,
	}

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}
	s8 := &helper.StepInfo{}
	s11 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create tenant1 with required attributes only.
			{
				Config: func() string {
					*stepCount++
					s1.Name = fmt.Sprintf("%s_%d_create_tenant1_required", t.Name(), *stepCount)

					helper.GenerateTenantObject(&tenant1Rsc, tenant1Name, nil)
					helper.GetTFConfigWithSingleResource(s1.Name, *xTenant1,
						[]interface{}{tenant1Rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						TenantModelHelperStateCheck(
							"nd_tenant.tenant_one",
							*tenant1Rsc,
							path.Empty(),
						),
						resource.TestCheckNoResourceAttr("nd_tenant.tenant_one", "description"),
						resource.TestCheckResourceAttr("nd_tenant.tenant_one", "fabric_associations.#", "0"),
					)...,
				),
			},
			// Step 2: Add tenant2 with description and fabric association
			// optional values. Tenant2 intentionally does not set tenant_prefix.
			{
				Config: func() string {
					*stepCount++
					s2.Name = fmt.Sprintf("%s_%d_create_tenant2_optional", t.Name(), *stepCount)

					helper.GenerateTenantObject(&tenant2Rsc, tenant2Name, map[string]interface{}{
						"description": "Tenant2 acceptance test",
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
							tenant2Association,
						},
					})
					helper.GetTFConfigWithSingleResource(s2.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_one",
								*tenant1Rsc,
								path.Empty(),
							),
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_two",
								*tenant2Rsc,
								path.Empty(),
							)...,
						),
						tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
					)...,
				),
			},
			// Step 3: Update tenant1 description and add two fabric
			// associations: one required-only and one with optional values.
			{
				Config: func() string {
					*stepCount++
					s3.Name = fmt.Sprintf("%s_%d_update_tenant1_add_associations", t.Name(), *stepCount)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"description": "Tenant1 acceptance test",
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
							tenant1RequiredAssociation,
							tenant1FullAssociationCreate,
						},
					})
					helper.GetTFConfigWithSingleResource(s3.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 3, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_one",
								*tenant1Rsc,
								path.Empty(),
							),
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_two",
								*tenant2Rsc,
								path.Empty(),
							)...,
						),
						append(
							tenantAssociationStateChecks(
								"nd_tenant.tenant_one",
								tenant1RequiredAssociation,
								tenant1FullAssociationCreate,
							),
							tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
						)...,
					)...,
				),
			},
			// Step 4: Remove tenant1 description, remove the required-only
			// association, and update allowed_vlans/local_name on the remaining
			// association.
			{
				Config: func() string {
					*stepCount++
					s4.Name = fmt.Sprintf("%s_%d_update_tenant1_remove_and_update_associations", t.Name(), *stepCount)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"description": "",
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
							tenant1FullAssociationUpdate,
						},
					})
					helper.GetTFConfigWithSingleResource(s4.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 4, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_one",
								*tenant1Rsc,
								path.Empty(),
							),
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_two",
								*tenant2Rsc,
								path.Empty(),
							)...,
						),
						append(
							[]resource.TestCheckFunc{
								resource.TestCheckNoResourceAttr("nd_tenant.tenant_one", "description"),
							},
							append(
								tenantAssociationStateChecks("nd_tenant.tenant_one", tenant1FullAssociationUpdate),
								tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
							)...,
						)...,
					)...,
				),
			},
			// Step 5: Remove local_name and allowed_vlans from the tenant1
			// association while keeping the original tenant_prefix.
			{
				Config: func() string {
					*stepCount++
					s5.Name = fmt.Sprintf("%s_%d_update_tenant1_remove_association_optional_values", t.Name(), *stepCount)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
							tenant1PrefixOnlyAssociation,
						},
					})
					helper.GetTFConfigWithSingleResource(s5.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s5.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 5, s5.Name, s5.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_one",
								*tenant1Rsc,
								path.Empty(),
							),
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_two",
								*tenant2Rsc,
								path.Empty(),
							)...,
						),
						append(
							[]resource.TestCheckFunc{
								resource.TestCheckNoResourceAttr("nd_tenant.tenant_one", "description"),
								tenantAssociationOptionalAttrsAbsent("nd_tenant.tenant_one", tenantFabricTwo),
							},
							append(
								tenantAssociationStateChecks("nd_tenant.tenant_one", tenant1PrefixOnlyAssociation),
								tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
							)...,
						)...,
					)...,
				),
			},
			// Step 6: Try to mutate tenant_prefix on the existing tenant1
			// association. The backend rejects this unless the association is
			// deleted and re-created.
			{
				Config: func() string {
					*stepCount++
					s6.Name = fmt.Sprintf("%s_%d_update_tenant1_change_tenant_prefix_error", t.Name(), *stepCount)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
							tenant1InvalidPrefixAssociation,
						},
					})
					helper.GetTFConfigWithSingleResource(s6.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 6, s6.Name, s6.Cfg) },
				ExpectError: regexp.MustCompile(
					`(?is)tenant\s+prefix\s+cannot\s+be\s+changed\s+unless\s+this\s+association\s+is\s+deleted\s+and\s+re-created`,
				),
			},
			// Step 7: Return to a valid tenant1 config and remove every tenant1
			// fabric association.
			{
				Config: func() string {
					*stepCount++
					s7.Name = fmt.Sprintf("%s_%d_update_tenant1_remove_all_associations", t.Name(), *stepCount)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue(nil),
					})
					helper.GetTFConfigWithSingleResource(s7.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 7, s7.Name, s7.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_one",
								*tenant1Rsc,
								path.Empty(),
							),
							TenantModelHelperStateCheck(
								"nd_tenant.tenant_two",
								*tenant2Rsc,
								path.Empty(),
							)...,
						),
						append(
							[]resource.TestCheckFunc{
								resource.TestCheckNoResourceAttr("nd_tenant.tenant_one", "description"),
								resource.TestCheckResourceAttr("nd_tenant.tenant_one", "fabric_associations.#", "0"),
							},
							tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
						)...,
					)...,
				),
			},
			// Step 8: Remove tenant1 from Terraform config.
			{
				Config: func() string {
					*stepCount++
					s8.Name = fmt.Sprintf("%s_%d_delete_tenant1", t.Name(), *stepCount)

					helper.GetTFConfigWithSingleResource(s8.Name, *xTenant2,
						[]interface{}{tenant2Rsc}, &tfConfig)

					s8.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 8, s8.Name, s8.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						TenantModelHelperStateCheck(
							"nd_tenant.tenant_two",
							*tenant2Rsc,
							path.Empty(),
						),
						tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Association)...,
					)...,
				),
			},
			// Step 9: Import tenant2 by name. Import hydrates
			// fabric_associations from the manage-side association API because
			// the tenant GET API does not return that collection.
			{
				PreConfig: func() {
					t.Logf("===== STEP 9: %s_9_import_tenant2 =====", t.Name())
				},
				ResourceName:                         "nd_tenant.tenant_two",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        tenant2Name,
				ImportStateVerifyIdentifierAttribute: "name",
			},
			// Step 10: ImportState with the deleted tenant1 name. This
			// verifies the import path reports a tenant-not-found error for
			// non-existing tenants.
			{
				PreConfig: func() {
					t.Logf("===== STEP 10: %s_10_import_missing_tenant1 =====", t.Name())
				},
				ResourceName:  "nd_tenant.tenant_two",
				ImportState:   true,
				ImportStateId: tenant1Name,
				ExpectError:   regexp.MustCompile(`(?is)tenant.*not\s+found`),
			},
			// Step 11: Clear tenant2 after an outside-Terraform delete. The
			// provider should treat the backend "tenant not found: <name>"
			// response as successful cleanup.
			{
				PreConfig: func() {
					s11.Name = fmt.Sprintf("%s_11_clear_tenant2", t.Name())
					s11.Cfg = *tfConfig
					helper.LogStep(t, 11, s11.Name, s11.Cfg)
					deleteTenantOutsideTerraform(t, tenant2Name)
				},
				Config:  *tfConfig,
				Destroy: true,
			},
		},
	})
}

// TestAccTenantResourceAssociationRollback verifies that partial multi-status
// association failures do not leave a partially created or updated tenant.
func TestAccTenantResourceAssociationRollback(t *testing.T) {
	cfg := helper.GetConfig("global")
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)

	createTenantName := fmt.Sprintf("tf_create_rollback_%s", suffix)
	updateTenantName := fmt.Sprintf("tf_update_rollback_%s", suffix)
	missingFabricName := fmt.Sprintf("tf_missing_fabric_%s", suffix)
	originalDescription := "Tenant rollback acceptance test"
	updatedDescription := "Tenant rollback description must not be applied"
	originalPrefix := fmt.Sprintf("tn_rollback_%s", suffix)
	invalidPrefix := fmt.Sprintf("tn_changed_%s", suffix)
	originalLocalName := fmt.Sprintf("local_rollback_%s", suffix)
	updatedLocalName := fmt.Sprintf("local_rollback_updated_%s", suffix)
	originalVlans := []string{"1", "5-10"}
	updatedVlans := []string{"11-20", "30"}

	createValidAssociation := resource_tenant.NDFCFabricAssociationsValue{
		FabricName: tenantFabricOne,
	}
	createInvalidAssociation := resource_tenant.NDFCFabricAssociationsValue{
		FabricName: missingFabricName,
	}
	originalAssociationOne := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricOne,
		LocalName:    originalLocalName,
		AllowedVlans: originalVlans,
	}
	originalAssociationTwo := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: originalPrefix,
	}
	updatedAssociationOne := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricOne,
		LocalName:    updatedLocalName,
		AllowedVlans: updatedVlans,
	}
	invalidAssociationTwo := resource_tenant.NDFCFabricAssociationsValue{
		FabricName:   tenantFabricTwo,
		TenantPrefix: invalidPrefix,
	}

	createTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&createTenant, createTenantName, map[string]interface{}{
		"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
			createValidAssociation,
			createInvalidAssociation,
		},
	})

	originalUpdateTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&originalUpdateTenant, updateTenantName, map[string]interface{}{
		"description": originalDescription,
		"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
			originalAssociationOne,
			originalAssociationTwo,
		},
	})

	failedUpdateTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&failedUpdateTenant, updateTenantName, map[string]interface{}{
		"description": updatedDescription,
		"fabric_associations": []resource_tenant.NDFCFabricAssociationsValue{
			updatedAssociationOne,
			invalidAssociationTwo,
		},
	})

	createConfigArgs := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_create_rollback",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
	updateConfigArgs := &map[string]string{
		"RscType":  "nd_tenant",
		"RscName":  "tenant_update_rollback",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	createConfig := new(string)
	originalUpdateConfig := new(string)
	failedUpdateConfig := new(string)
	destroyConfig := new(string)
	s1 := &helper.StepInfo{
		Index: 1,
		Name:  fmt.Sprintf("%s_1_create_partial_association_failure", t.Name()),
	}
	s2 := &helper.StepInfo{
		Index: 2,
		Name:  fmt.Sprintf("%s_2_create_update_rollback_baseline", t.Name()),
	}
	s3 := &helper.StepInfo{
		Index: 3,
		Name:  fmt.Sprintf("%s_3_update_partial_association_failure", t.Name()),
	}
	s4 := &helper.StepInfo{
		Index: 4,
		Name:  fmt.Sprintf("%s_4_verify_update_rollback_and_destroy", t.Name()),
	}
	helper.GetTFConfigWithSingleResource(
		s1.Name,
		*createConfigArgs,
		[]interface{}{createTenant},
		&createConfig,
	)
	s1.Cfg = *createConfig
	helper.GetTFConfigWithSingleResource(
		s2.Name,
		*updateConfigArgs,
		[]interface{}{originalUpdateTenant},
		&originalUpdateConfig,
	)
	s2.Cfg = *originalUpdateConfig
	helper.GetTFConfigWithSingleResource(
		s3.Name,
		*updateConfigArgs,
		[]interface{}{failedUpdateTenant},
		&failedUpdateConfig,
	)
	s3.Cfg = *failedUpdateConfig
	helper.GetTFConfigWithSingleResource(
		s4.Name,
		*updateConfigArgs,
		[]interface{}{originalUpdateTenant},
		&destroyConfig,
	)
	s4.Cfg = *destroyConfig

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// An association API failure after tenant creation must trigger rollback.
			// Rollback removes any associations the API applied and deletes the tenant.
			{
				Config:      *createConfig,
				PreConfig:   func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				ExpectError: regexp.MustCompile(`(?is)(tenant fabric association request failed|fabric.*not\s+found|fabric.*does\s+not\s+exist)`),
			},
			// Assert rollback directly against the APIs before Terraform can
			// perform another operation, then create the update-test baseline.
			{
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					assertTenantAndAssociationsAbsentOutsideTerraform(t, createTenantName)
				},
				Config: *originalUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					append(
						TenantModelHelperStateCheck(
							"nd_tenant.tenant_update_rollback",
							*originalUpdateTenant,
							path.Empty(),
						),
						tenantAssociationStateChecks(
							"nd_tenant.tenant_update_rollback",
							originalAssociationOne,
							originalAssociationTwo,
						)...,
					)...,
				),
			},
			// The first association update is valid and the tenant_prefix change
			// is rejected. Update rollback must restore the complete old set and
			// must not post the planned description.
			{
				Config:    *failedUpdateConfig,
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				ExpectError: regexp.MustCompile(
					`(?is)tenant\s+prefix\s+cannot\s+be\s+changed\s+unless\s+this\s+association\s+is\s+deleted\s+and\s+re-created`,
				),
			},
			{
				PreConfig: func() {
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
					assertTenantConfigurationOutsideTerraform(
						t,
						updateTenantName,
						originalDescription,
						originalAssociationOne,
						originalAssociationTwo,
					)
				},
				Config:  *destroyConfig,
				Destroy: true,
			},
		},
	})
}

func tenantAssociationStateChecks(resourceName string, associations ...resource_tenant.NDFCFabricAssociationsValue) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "fabric_associations.#", fmt.Sprintf("%d", len(associations))),
	}

	for _, association := range associations {
		nestedAttrs := map[string]string{
			"fabric_name": association.FabricName,
		}
		if association.TenantPrefix != "" {
			nestedAttrs["tenant_prefix"] = association.TenantPrefix
		}
		if association.LocalName != "" {
			nestedAttrs["local_name"] = association.LocalName
		}
		if len(association.AllowedVlans) > 0 {
			nestedAttrs["allowed_vlans.#"] = fmt.Sprintf("%d", len(association.AllowedVlans))
		}

		checks = append(checks,
			resource.TestCheckTypeSetElemNestedAttrs(resourceName, "fabric_associations.*", nestedAttrs),
		)
		for _, vlan := range association.AllowedVlans {
			checks = append(checks,
				resource.TestCheckTypeSetElemAttr(resourceName, "fabric_associations.*.allowed_vlans.*", vlan),
			)
		}
	}

	return checks
}

func tenantAssociationOptionalAttrsAbsent(resourceName string, fabricName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceName)
		}
		if rs.Primary == nil {
			return fmt.Errorf("resource %q has no primary instance", resourceName)
		}

		attrs := rs.Primary.Attributes
		associationPrefix := ""
		for key, value := range attrs {
			if strings.HasPrefix(key, "fabric_associations.") &&
				strings.HasSuffix(key, ".fabric_name") &&
				value == fabricName {
				associationPrefix = strings.TrimSuffix(key, ".fabric_name")
				break
			}
		}
		if associationPrefix == "" {
			return fmt.Errorf("resource %q has no fabric association for fabric %q", resourceName, fabricName)
		}

		if value, ok := attrs[associationPrefix+".local_name"]; ok && value != "" {
			return fmt.Errorf("resource %q association %q local_name expected absent, got %q", resourceName, fabricName, value)
		}
		if value, ok := attrs[associationPrefix+".allowed_vlans.#"]; ok && value != "0" {
			return fmt.Errorf("resource %q association %q allowed_vlans expected absent or empty, got count %q", resourceName, fabricName, value)
		}
		for key, value := range attrs {
			if strings.HasPrefix(key, associationPrefix+".allowed_vlans.") &&
				!strings.HasSuffix(key, ".#") &&
				value != "" {
				return fmt.Errorf("resource %q association %q allowed_vlans expected absent, found %s=%q", resourceName, fabricName, key, value)
			}
		}

		return nil
	}
}

func newTenantTestClient(t *testing.T) nd.Client {
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
		t.Fatalf("failed to create ND client for tenant API test helper: %v", err)
	}

	return client
}

func readTenantFabricAssociationsOutsideTerraform(t *testing.T, client *nd.Client, name string) []tenantFabricAssociationTestItem {
	t.Helper()

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(client, ndapi.DefaultFabric)
	respData, err := tenantFabricAssocAPI.Get()
	if err != nil {
		t.Fatalf("failed to read tenant fabric associations for tenant %q: %v %s", name, err, string(respData))
	}

	var associationResp tenantFabricAssociationTestListResponse
	err = json.Unmarshal(respData, &associationResp)
	if err != nil {
		t.Fatalf("failed to unmarshal tenant fabric associations for tenant %q: %v", name, err)
	}

	associations := make([]tenantFabricAssociationTestItem, 0)
	for _, association := range associationResp.TenantFabricAssociations {
		if association.TenantName == name {
			associations = append(associations, association)
		}
	}

	return associations
}

func assertTenantAndAssociationsAbsentOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	client := newTenantTestClient(t)
	tenantAPI := api.NewTenantAPI(&client, ndapi.DefaultFabric)
	tenantAPI.TenantName = name

	respData, err := tenantAPI.Get()
	if err == nil {
		t.Fatalf("tenant %q still exists after create rollback: %s", name, string(respData))
	}
	if !isTenantNotFoundAPIError(err, string(respData), name) {
		t.Fatalf("failed to verify tenant %q was removed after create rollback: %v %s", name, err, string(respData))
	}

	associations := readTenantFabricAssociationsOutsideTerraform(t, &client, name)
	if len(associations) != 0 {
		t.Fatalf("tenant %q still has fabric associations after create rollback: %+v", name, associations)
	}
}

func assertTenantConfigurationOutsideTerraform(
	t *testing.T,
	name string,
	expectedDescription string,
	expectedAssociations ...resource_tenant.NDFCFabricAssociationsValue,
) {
	t.Helper()

	client := newTenantTestClient(t)
	tenantAPI := api.NewTenantAPI(&client, ndapi.DefaultFabric)
	tenantAPI.TenantName = name

	respData, err := tenantAPI.Get()
	if err != nil {
		t.Fatalf("failed to read tenant %q after update rollback: %v %s", name, err, string(respData))
	}

	var tenantResp resource_tenant.NDFCTenantModel
	if err := json.Unmarshal(respData, &tenantResp); err != nil {
		t.Fatalf("failed to unmarshal tenant %q after update rollback: %v", name, err)
	}
	if tenantResp.Description != expectedDescription {
		t.Fatalf("tenant %q description after update rollback: expected %q, got %q", name, expectedDescription, tenantResp.Description)
	}

	actualAssociations := readTenantFabricAssociationsOutsideTerraform(t, &client, name)
	actualByFabric := make(map[string]tenantFabricAssociationTestItem, len(actualAssociations))
	for _, association := range actualAssociations {
		if _, ok := actualByFabric[association.FabricName]; ok {
			t.Fatalf("tenant %q has duplicate backend associations for fabric %q", name, association.FabricName)
		}
		actualByFabric[association.FabricName] = association
	}

	if len(actualByFabric) != len(expectedAssociations) {
		t.Fatalf("tenant %q association count after update rollback: expected %d, got %d (%+v)", name, len(expectedAssociations), len(actualByFabric), actualAssociations)
	}

	for _, expected := range expectedAssociations {
		actual, ok := actualByFabric[expected.FabricName]
		if !ok {
			t.Fatalf("tenant %q is missing association for fabric %q after update rollback", name, expected.FabricName)
		}

		actualVlans := append([]string(nil), actual.AllowedVlans...)
		expectedVlans := append([]string(nil), expected.AllowedVlans...)
		slices.Sort(actualVlans)
		slices.Sort(expectedVlans)
		if actual.LocalName != expected.LocalName ||
			actual.TenantPrefix != expected.TenantPrefix ||
			!slices.Equal(actualVlans, expectedVlans) {
			t.Fatalf("tenant %q association for fabric %q after update rollback: expected %+v, got %+v", name, expected.FabricName, expected, actual)
		}
	}
}

func deleteTenantOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	res, err := deleteTenantOutsideTerraformRaw(t, name)
	if err != nil && !isTenantNotFoundAPIError(err, res, name) {
		t.Fatalf("failed to delete tenant %q outside Terraform: %v %s", name, err, res)
	}
}

func deleteTenantOutsideTerraformRaw(t *testing.T, name string) (string, error) {
	t.Helper()

	client := newTenantTestClient(t)
	deleteTenantFabricAssociationsOutsideTerraform(t, &client, name)

	tenantAPI := api.NewTenantAPI(&client, ndapi.DefaultFabric)
	tenantAPI.TenantName = name
	res, err := tenantAPI.Delete()
	return res.String(), err
}

func deleteTenantFabricAssociationsOutsideTerraform(t *testing.T, client *nd.Client, name string) {
	t.Helper()

	associations := readTenantFabricAssociationsOutsideTerraform(t, client, name)
	payload := tenantFabricAssociationTestPayload{}
	for _, association := range associations {
		association.Associate = false
		payload.Items = append(payload.Items, association)
	}
	if len(payload.Items) == 0 {
		return
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal tenant fabric association delete payload for tenant %q: %v", name, err)
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(client, ndapi.DefaultFabric)
	res, err := tenantFabricAssocAPI.Post(payloadBytes, nil)
	if err != nil {
		t.Fatalf("failed to delete tenant fabric associations for tenant %q outside Terraform: %v %s", name, err, res.String())
	}

	var failed []string
	res.Get("results").ForEach(func(_, item gjson.Result) bool {
		if strings.EqualFold(item.Get("status").String(), "failed") {
			failed = append(failed, fmt.Sprintf("fabricName=%q message=%q", item.Get("fabricName").String(), item.Get("message").String()))
		}
		return true
	})
	if len(failed) > 0 {
		t.Fatalf("failed to delete tenant fabric associations for tenant %q outside Terraform: %s", name, strings.Join(failed, "; "))
	}
}

func isTenantNotFoundAPIError(err error, response string, name string) bool {
	if err == nil {
		return false
	}

	notFoundMessage := fmt.Sprintf("tenant not found: %s", name)
	return strings.Contains(err.Error(), "StatusCode 404") ||
		strings.Contains(err.Error(), notFoundMessage) ||
		strings.Contains(response, notFoundMessage)
}
