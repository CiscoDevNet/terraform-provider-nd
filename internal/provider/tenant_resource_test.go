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
	"time"

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
	tenantFabricOne                    = "ansible-test"
	tenantFabricTwo                    = "ansible-test-2"
	tenantAssociationOrchestrationWait = 5 * time.Second
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
	tenant2DriftDescription := "Tenant2 description changed outside Terraform"
	tenant2DriftLocalName := fmt.Sprintf("local_tenant2_drift_%s", tenant2Suffix)

	allowedVlansCreate := []string{"1", "5-10"}
	allowedVlansUpdate := []string{"5-10", "11-20", "30"}
	allowedVlansDrift := []string{"100-110", "200"}

	testCases := []string{
		"Create tenant1 with required attributes only",
		"Create tenant2 with description and optional fabric association values",
		"Update tenant1 by setting the description and adding two fabric associations",
		"Update tenant1 by clearing the description, removing one fabric association, and updating the remaining association",
		"Update tenant1 by clearing optional fabric association values",
		"Update tenant1 by replacing the fabric association to change its tenant prefix",
		"Update tenant1 by removing all fabric associations",
		"Delete tenant1 by removing it from the Terraform configuration",
		"Import tenant2 by name",
		"Verify importing a missing tenant returns a not-found error",
		"Detect tenant2 drift after out-of-band changes",
		"Destroy tenant2 after an out-of-band deletion",
	}
	for _, testCase := range testCases {
		t.Logf("Tenant test case: %s", testCase)
	}
	stepName := func(step int) string {
		return fmt.Sprintf("%s - %s", t.Name(), testCases[step-1])
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

	tenant1Rsc := new(resource_tenant.NDFCTenantModel)
	tenant2Rsc := new(resource_tenant.NDFCTenantModel)

	tenant1FullAssociationCreate := resource_tenant.NDFCFabricAssociationsValue{
		TenantPrefix: tenant1Prefix,
		LocalName:    tenant1LocalName,
		AllowedVlans: allowedVlansCreate,
	}
	tenant1FullAssociationUpdate := resource_tenant.NDFCFabricAssociationsValue{
		TenantPrefix: tenant1Prefix,
		LocalName:    tenant1UpdatedLocalName,
		AllowedVlans: allowedVlansUpdate,
	}
	tenant1PrefixOnlyAssociation := resource_tenant.NDFCFabricAssociationsValue{
		TenantPrefix: tenant1Prefix,
	}
	tenant1UpdatedPrefixAssociation := resource_tenant.NDFCFabricAssociationsValue{
		TenantPrefix: tenant1UpdatedPrefix,
	}
	tenant2Association := resource_tenant.NDFCFabricAssociationsValue{
		LocalName:    tenant2LocalName,
		AllowedVlans: allowedVlansCreate,
	}
	tenant1CreateAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricOne: {},
		tenantFabricTwo: tenant1FullAssociationCreate,
	}
	tenant1UpdateAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricTwo: tenant1FullAssociationUpdate,
	}
	tenant1PrefixOnlyAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricTwo: tenant1PrefixOnlyAssociation,
	}
	tenant1UpdatedPrefixAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricTwo: tenant1UpdatedPrefixAssociation,
	}
	tenant2Associations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricOne: tenant2Association,
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
	s12 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create tenant1 with required attributes only.
			{
				Config: func() string {
					s1.Name = stepName(1)

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
						resource.TestCheckResourceAttr("nd_tenant.tenant_one", "fabric_associations.%", "0"),
					)...,
				),
			},
			// Step 2: Add tenant2 with description and fabric association
			// optional values. Tenant2 intentionally does not set tenant_prefix.
			{
				Config: func() string {
					s2.Name = stepName(2)

					helper.GenerateTenantObject(&tenant2Rsc, tenant2Name, map[string]interface{}{
						"description":         "Tenant2 acceptance test",
						"fabric_associations": tenant2Associations,
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
						tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
					)...,
				),
			},
			// Step 3: Update tenant1 description and add two fabric
			// associations: one required-only and one with optional values.
			{
				Config: func() string {
					s3.Name = stepName(3)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"description":         "Tenant1 acceptance test",
						"fabric_associations": tenant1CreateAssociations,
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
								tenant1CreateAssociations,
							),
							tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
						)...,
					)...,
				),
			},
			// Step 4: Remove tenant1 description, remove the required-only
			// association, and update allowed_vlans/local_name on the remaining
			// association.
			{
				Config: func() string {
					s4.Name = stepName(4)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"description":         "",
						"fabric_associations": tenant1UpdateAssociations,
					})
					helper.GetTFConfigWithSingleResource(s4.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, 4, s4.Name, s4.Cfg)
					waitForTenantAssociationOrchestration(t)
				},
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
								tenantAssociationStateChecks("nd_tenant.tenant_one", tenant1UpdateAssociations),
								tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
							)...,
						)...,
					)...,
				),
			},
			// Step 5: Remove local_name and allowed_vlans from the tenant1
			// association while keeping the original tenant_prefix.
			{
				Config: func() string {
					s5.Name = stepName(5)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": tenant1PrefixOnlyAssociations,
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
								tenantAssociationStateChecks("nd_tenant.tenant_one", tenant1PrefixOnlyAssociations),
								tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
							)...,
						)...,
					)...,
				),
			},
			// Step 6: Update tenant_prefix on the existing tenant1 association.
			// The provider deletes and recreates the association before applying
			// the regular update payload.
			{
				Config: func() string {
					s6.Name = stepName(6)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": tenant1UpdatedPrefixAssociations,
					})
					helper.GetTFConfigWithSingleResource(s6.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, 6, s6.Name, s6.Cfg)
					waitForTenantAssociationOrchestration(t)
				},
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
							tenantAssociationStateChecks("nd_tenant.tenant_one", tenant1UpdatedPrefixAssociations),
							tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
						)...,
					)...,
				),
			},
			// Step 7: Remove every tenant1 fabric association.
			{
				Config: func() string {
					s7.Name = stepName(7)

					helper.ModifyTenantObject(&tenant1Rsc, map[string]interface{}{
						"fabric_associations": map[string]resource_tenant.NDFCFabricAssociationsValue(nil),
					})
					helper.GetTFConfigWithSingleResource(s7.Name, *xBothTenants,
						[]interface{}{tenant1Rsc, tenant2Rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, 7, s7.Name, s7.Cfg)
					waitForTenantAssociationOrchestration(t)
				},
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
								resource.TestCheckResourceAttr("nd_tenant.tenant_one", "fabric_associations.%", "0"),
							},
							tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
						)...,
					)...,
				),
			},
			// Step 8: Remove tenant1 from Terraform config.
			{
				Config: func() string {
					s8.Name = stepName(8)

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
						tenantAssociationStateChecks("nd_tenant.tenant_two", tenant2Associations)...,
					)...,
				),
			},
			// Step 9: Import tenant2 by name. Import hydrates
			// fabric_associations from the manage-side association API because
			// the tenant GET API does not return that collection.
			{
				PreConfig: func() {
					t.Logf("===== STEP 9: %s =====", stepName(9))
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
					t.Logf("===== STEP 10: %s =====", stepName(10))
				},
				ResourceName:  "nd_tenant.tenant_two",
				ImportState:   true,
				ImportStateId: tenant1Name,
				ExpectError:   regexp.MustCompile(`(?is)tenant.*not\s+found`),
			},
			// Step 11: Change every mutable tenant attribute type covered by this
			// resource outside Terraform, then verify refresh produces a non-empty
			// plan. PlanOnly ensures Terraform does not repair the drift.
			{
				PreConfig: func() {
					s11.Name = stepName(11)
					s11.Cfg = *tfConfig
					helper.LogStep(t, 11, s11.Name, s11.Cfg)
					updateTenantConfigurationOutsideTerraform(
						t,
						tenant2Name,
						tenant2DriftDescription,
						tenantFabricOne,
						tenant2DriftLocalName,
						allowedVlansDrift,
					)
				},
				Config:             *tfConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Step 12: Clear tenant2 after an outside-Terraform delete. The
			// provider should treat the backend "tenant not found: <name>"
			// response as successful cleanup.
			{
				PreConfig: func() {
					s12.Name = stepName(12)
					s12.Cfg = *tfConfig
					helper.LogStep(t, 12, s12.Name, s12.Cfg)
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
	testCases := []string{
		"Verify tenant creation rollback after a partial fabric association failure",
		"Create the tenant update rollback baseline",
		"Verify tenant update rollback after a partial fabric association failure",
		"Verify the restored tenant configuration and destroy the tenant",
	}
	stepName := func(step int) string {
		return fmt.Sprintf("%s - %s", t.Name(), testCases[step-1])
	}

	createTenantName := fmt.Sprintf("tf_create_rollback_%s", suffix)
	updateTenantName := fmt.Sprintf("tf_update_rollback_%s", suffix)
	missingFabricName := fmt.Sprintf("tf_missing_fabric_%s", suffix)
	originalDescription := "Tenant rollback acceptance test"
	updatedDescription := "Tenant rollback description must not be applied"
	originalPrefix := fmt.Sprintf("tn_rollback_%s", suffix)
	originalLocalName := fmt.Sprintf("local_rollback_%s", suffix)
	updatedLocalName := fmt.Sprintf("local_rollback_updated_%s", suffix)
	originalVlans := []string{"1", "5-10"}
	updatedVlans := []string{"11-20", "30"}

	originalAssociationOne := resource_tenant.NDFCFabricAssociationsValue{
		LocalName:    originalLocalName,
		AllowedVlans: originalVlans,
	}
	originalAssociationTwo := resource_tenant.NDFCFabricAssociationsValue{
		TenantPrefix: originalPrefix,
	}
	updatedAssociationOne := resource_tenant.NDFCFabricAssociationsValue{
		LocalName:    updatedLocalName,
		AllowedVlans: updatedVlans,
	}
	createAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricOne:   {},
		missingFabricName: {},
	}
	originalUpdateAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricOne: originalAssociationOne,
		tenantFabricTwo: originalAssociationTwo,
	}
	failedUpdateAssociations := map[string]resource_tenant.NDFCFabricAssociationsValue{
		tenantFabricOne:   updatedAssociationOne,
		tenantFabricTwo:   originalAssociationTwo,
		missingFabricName: {},
	}

	createTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&createTenant, createTenantName, map[string]interface{}{
		"fabric_associations": createAssociations,
	})

	originalUpdateTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&originalUpdateTenant, updateTenantName, map[string]interface{}{
		"description":         originalDescription,
		"fabric_associations": originalUpdateAssociations,
	})

	failedUpdateTenant := new(resource_tenant.NDFCTenantModel)
	helper.GenerateTenantObject(&failedUpdateTenant, updateTenantName, map[string]interface{}{
		"description":         updatedDescription,
		"fabric_associations": failedUpdateAssociations,
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
		Name:  stepName(1),
	}
	s2 := &helper.StepInfo{
		Index: 2,
		Name:  stepName(2),
	}
	s3 := &helper.StepInfo{
		Index: 3,
		Name:  stepName(3),
	}
	s4 := &helper.StepInfo{
		Index: 4,
		Name:  stepName(4),
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
			// The backend can reject a missing fabric with HTTP 400 or report it as
			// a failed item in a 207 response. Rollback removes any associations the
			// API applied and deletes the tenant.
			{
				Config:      *createConfig,
				PreConfig:   func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				ExpectError: regexp.MustCompile(`(?is)stage=create.*(status="failed"|StatusCode\s+400.*fabric.*does\s+not\s+exist)`),
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
							originalUpdateAssociations,
						)...,
					)...,
				),
			},
			// The existing association update is valid, but adding the missing
			// fabric association fails with either HTTP 400 or a failed 207 item.
			// Update rollback must restore the complete old set and must not post
			// the planned description.
			{
				Config:      *failedUpdateConfig,
				PreConfig:   func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				ExpectError: regexp.MustCompile(`(?is)stage=regular_update.*(status="failed"|StatusCode\s+400.*fabric.*does\s+not\s+exist)`),
			},
			{
				PreConfig: func() {
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
					waitForTenantAssociationOrchestration(t)
					assertTenantConfigurationOutsideTerraform(
						t,
						updateTenantName,
						originalDescription,
						originalUpdateAssociations,
					)
				},
				Config:  *destroyConfig,
				Destroy: true,
			},
		},
	})
}

// waitForTenantAssociationOrchestration gives asynchronous association changes
// time to settle before a test removes tenant fabric associations.
func waitForTenantAssociationOrchestration(t *testing.T) {
	t.Helper()
	t.Logf("Waiting %s for tenant fabric association orchestration to settle", tenantAssociationOrchestrationWait)
	time.Sleep(tenantAssociationOrchestrationWait)
}

func tenantAssociationStateChecks(
	resourceName string,
	associations map[string]resource_tenant.NDFCFabricAssociationsValue,
) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "fabric_associations.%", fmt.Sprintf("%d", len(associations))),
	}

	for fabricName, association := range associations {
		associationPrefix := fmt.Sprintf("fabric_associations.%s", fabricName)
		if len(association.AllowedVlans) > 0 {
			checks = append(checks,
				resource.TestCheckResourceAttr(resourceName, associationPrefix+".allowed_vlans.#", fmt.Sprintf("%d", len(association.AllowedVlans))),
			)
		}

		for _, vlan := range association.AllowedVlans {
			checks = append(checks,
				resource.TestCheckTypeSetElemAttr(resourceName, associationPrefix+".allowed_vlans.*", vlan),
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
		associationPrefix := fmt.Sprintf("fabric_associations.%s", fabricName)
		associationFound := false
		for key := range attrs {
			if strings.HasPrefix(key, associationPrefix+".") {
				associationFound = true
				break
			}
		}
		if !associationFound {
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
	expectedAssociations map[string]resource_tenant.NDFCFabricAssociationsValue,
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

	for fabricName, expected := range expectedAssociations {
		actual, ok := actualByFabric[fabricName]
		if !ok {
			t.Fatalf("tenant %q is missing association for fabric %q after update rollback", name, fabricName)
		}

		actualVlans := append([]string(nil), actual.AllowedVlans...)
		expectedVlans := append([]string(nil), expected.AllowedVlans...)
		slices.Sort(actualVlans)
		slices.Sort(expectedVlans)
		if actual.LocalName != expected.LocalName ||
			actual.TenantPrefix != expected.TenantPrefix ||
			!slices.Equal(actualVlans, expectedVlans) {
			t.Fatalf("tenant %q association for fabric %q after update rollback: expected %+v, got %+v", name, fabricName, expected, actual)
		}
	}
}

func updateTenantConfigurationOutsideTerraform(
	t *testing.T,
	name string,
	description string,
	fabricName string,
	localName string,
	allowedVlans []string,
) {
	t.Helper()

	client := newTenantTestClient(t)
	tenantAPI := api.NewTenantAPI(&client, ndapi.DefaultFabric)
	tenantAPI.TenantName = name

	tenantPayload, err := json.Marshal(resource_tenant.NDFCTenantModel{
		Name:        name,
		Description: description,
	})
	if err != nil {
		t.Fatalf("failed to marshal tenant update payload for tenant %q outside Terraform: %v", name, err)
	}

	res, err := tenantAPI.Put(tenantPayload, nil)
	if err != nil {
		t.Fatalf("failed to update description for tenant %q outside Terraform: %v %s", name, err, res.String())
	}

	associations := readTenantFabricAssociationsOutsideTerraform(t, &client, name)
	var associationToUpdate *tenantFabricAssociationTestItem
	for i := range associations {
		if associations[i].FabricName != fabricName {
			continue
		}
		if associationToUpdate != nil {
			t.Fatalf("tenant %q has duplicate backend associations for fabric %q", name, fabricName)
		}
		associationToUpdate = &associations[i]
	}
	if associationToUpdate == nil {
		t.Fatalf("tenant %q has no backend association for fabric %q", name, fabricName)
	}

	associationToUpdate.Associate = true
	associationToUpdate.LocalName = localName
	associationToUpdate.AllowedVlans = append([]string(nil), allowedVlans...)
	postTenantFabricAssociationsOutsideTerraform(
		t,
		&client,
		name,
		"update",
		tenantFabricAssociationTestPayload{Items: []tenantFabricAssociationTestItem{*associationToUpdate}},
	)
	waitForTenantAssociationOrchestration(t)
}

func deleteTenantOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	client := newTenantTestClient(t)
	deleteTenantFabricAssociationsOutsideTerraform(t, &client, name)

	tenantAPI := api.NewTenantAPI(&client, ndapi.DefaultFabric)
	tenantAPI.TenantName = name
	res, err := tenantAPI.Delete()
	if err != nil && !isTenantNotFoundAPIError(err, res.String(), name) {
		t.Fatalf("failed to delete tenant %q outside Terraform: %v %s", name, err, res.String())
	}
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

	waitForTenantAssociationOrchestration(t)
	postTenantFabricAssociationsOutsideTerraform(t, client, name, "delete", payload)
}

func postTenantFabricAssociationsOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	tenantName string,
	operation string,
	payload tenantFabricAssociationTestPayload,
) {
	t.Helper()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal tenant fabric association %s payload for tenant %q: %v", operation, tenantName, err)
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(client, ndapi.DefaultFabric)
	res, err := tenantFabricAssocAPI.Post(payloadBytes, nil)
	if err != nil {
		t.Fatalf("failed to %s tenant fabric associations for tenant %q outside Terraform: %v %s", operation, tenantName, err, res.String())
	}

	results := res.Get("results")
	if !results.Exists() || !results.IsArray() {
		t.Fatalf("tenant fabric association %s response for tenant %q did not include a valid results array: %s", operation, tenantName, res.String())
	}

	var failed []string
	results.ForEach(func(_, item gjson.Result) bool {
		if strings.EqualFold(item.Get("status").String(), "failed") {
			failed = append(failed, fmt.Sprintf("fabricName=%q message=%q", item.Get("fabricName").String(), item.Get("message").String()))
		}
		return true
	})
	if len(failed) > 0 {
		t.Fatalf("failed to %s tenant fabric associations for tenant %q outside Terraform: %s", operation, tenantName, strings.Join(failed, "; "))
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
