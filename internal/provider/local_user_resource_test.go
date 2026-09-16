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

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_local_user"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	nd "github.com/netascode/go-nd"
)

const localUserDefaultTenantDomain = "all-tenants-domain"

// TestAccLocalUserResourceCRUD exercises local-user creation, mutable updates,
// drift handling, read/delete 404 behavior, and schema validation.
func TestAccLocalUserResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	luCfg := cfg.ND.LocalUser
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)

	requiredLoginID := fmt.Sprintf("%s_required_%s", luCfg.LoginID, suffix)
	fullLoginID := fmt.Sprintf("%s_full_%s", luCfg.LoginID, suffix)
	missingDestroyLoginID := fmt.Sprintf("%s_del404_%s", luCfg.LoginID, suffix)
	invalidLoginID := fmt.Sprintf("%s_invalid_%s", luCfg.LoginID, suffix)

	securityDomains := copyLocalUserSecurityDomains(luCfg.SecurityDomains)
	updatedSecurityDomains := copyLocalUserSecurityDomains(luCfg.SecurityDomains)
	for domainName, roles := range updatedSecurityDomains {
		updatedRole := "observer"
		if len(roles) == 1 && roles[0] == updatedRole {
			updatedRole = "designer"
		}
		updatedSecurityDomains[domainName] = []string{updatedRole}
		break
	}

	tenantDomain := luCfg.TenantDomain
	if tenantDomain == "" {
		tenantDomain = localUserDefaultTenantDomain
	}
	remoteIDClaim := luCfg.RemoteIDClaim
	if remoteIDClaim == "" {
		remoteIDClaim = "tf_remote_id_claim"
	}
	email := luCfg.Email
	if email == "" {
		email = "tf_local_user@example.com"
	}
	firstName := luCfg.FirstName
	if firstName == "" {
		firstName = "Terraform"
	}
	lastName := luCfg.LastName
	if lastName == "" {
		lastName = "User"
	}

	preexistingUser := new(resource_local_user.NDFCLocalUserModel)
	requiredUser := new(resource_local_user.NDFCLocalUserModel)
	fullUser := new(resource_local_user.NDFCLocalUserModel)
	invalidUser := new(resource_local_user.NDFCLocalUserModel)

	cleanupEnabled := false
	t.Cleanup(func() {
		if !cleanupEnabled {
			return
		}
		deleteLocalUserOutsideTerraform(t, requiredLoginID)
		deleteLocalUserOutsideTerraform(t, fullLoginID)
		deleteLocalUserOutsideTerraform(t, missingDestroyLoginID)
		deleteLocalUserOutsideTerraform(t, invalidLoginID)
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
	s10 := &helper.StepInfo{}
	s11 := &helper.StepInfo{}
	s12 := &helper.StepInfo{}
	s13 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(luCfg.LoginID) == "" {
				t.Fatal("local_user.login_id must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(luCfg.UserPassword) == "" {
				t.Fatal("local_user.user_password must be configured in the acceptance-test testbed")
			}
			if len(securityDomains) == 0 {
				t.Fatal("local_user.security_domains must contain at least one security domain in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject creation when the local user already exists in Nexus Dashboard")

					generateRequiredLocalUserObject(
						&preexistingUser,
						requiredLoginID,
						luCfg.UserPassword,
						securityDomains,
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						configArgs,
						[]interface{}{preexistingUser},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteLocalUserOutsideTerraform(t, requiredLoginID)
					createLocalUserOutsideTerraform(t, *preexistingUser)
				},
				ExpectError: regexp.MustCompile(`(?is)Error Creating ND Local User.*Could not create nd_local_user`),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Create a local user with required attributes")

					generateRequiredLocalUserObject(
						&requiredUser,
						requiredLoginID,
						luCfg.UserPassword,
						securityDomains,
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s2.Name,
						configArgs,
						[]interface{}{requiredUser},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					deleteLocalUserOutsideTerraform(t, requiredLoginID)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserOptionalStringsAbsentChecks("nd_local_user.required_user")...,
					)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Create a second local user with the complete supported configuration")

					helper.GenerateLocalUserObject(
						&fullUser,
						fullLoginID,
						luCfg.UserPassword,
						securityDomains,
						map[string]interface{}{
							"email":                     email,
							"first_name":                firstName,
							"last_name":                 lastName,
							"remote_id_claim":           fmt.Sprintf("%s_full_%s", remoteIDClaim, suffix),
							"remote_user_authorization": true,
							"tenant_domain":             tenantDomain,
						},
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s3.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Configure every optional attribute on the required-only local user")

					helper.ModifyLocalUserObject(&requiredUser, map[string]interface{}{
						"email":                     email,
						"first_name":                firstName,
						"last_name":                 lastName,
						"remote_id_claim":           fmt.Sprintf("%s_required_%s", remoteIDClaim, suffix),
						"remote_user_authorization": true,
						"tenant_domain":             tenantDomain,
					})
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s4.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Update the local-user security-domain roles")

					for domainName, roles := range updatedSecurityDomains {
						helper.AddSecurityDomain(&requiredUser, domainName, roles)
					}
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s5.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s5.Cfg = *tfConfig
					return s5.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s5.Index, s5.Name, s5.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Clear optional strings and restore defaulted optional attributes")

					generateRequiredLocalUserObject(
						&requiredUser,
						requiredLoginID,
						luCfg.UserPassword,
						updatedSecurityDomains,
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s6.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s6.Index, s6.Name, s6.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						append(
							localUserStateChecks("nd_local_user.required_user", *requiredUser),
							localUserOptionalStringsAbsentChecks("nd_local_user.required_user")...,
						),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan with optional local-user attributes omitted")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s7.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s8.Index = 8
					s8.Name = fmt.Sprintf("%s - %s", t.Name(), "Re-add every optional attribute for drift verification")

					helper.ModifyLocalUserObject(&requiredUser, map[string]interface{}{
						"email":                     email,
						"first_name":                firstName,
						"last_name":                 lastName,
						"remote_id_claim":           fmt.Sprintf("%s_required_%s", remoteIDClaim, suffix),
						"remote_user_authorization": true,
						"tenant_domain":             tenantDomain,
					})
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s8.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s8.Cfg = *tfConfig
					return s8.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s8.Index, s8.Name, s8.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s9.Index = 9
					s9.Name = fmt.Sprintf("%s - %s", t.Name(), "Detect an out-of-band local-user attribute change")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s9.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s9.Cfg = *tfConfig
					return s9.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s9.Index, s9.Name, s9.Cfg)
					driftedUser := *requiredUser
					driftedUser.FirstName = "Changed outside Terraform"
					updateLocalUserOutsideTerraform(t, driftedUser)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s10.Index = 10
					s10.Name = fmt.Sprintf("%s - %s", t.Name(), "Recreate the local user after an out-of-band deletion")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user,full_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s10.Name,
						configArgs,
						[]interface{}{requiredUser, fullUser},
						&tfConfig,
					)
					s10.Cfg = *tfConfig
					return s10.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s10.Index, s10.Name, s10.Cfg)
					deleteLocalUserOutsideTerraform(t, requiredLoginID)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.required_user", *requiredUser),
						localUserStateChecks("nd_local_user.full_user", *fullUser)...,
					)...,
				),
			},
			{
				Config: func() string {
					s11.Index = 11
					s11.Name = fmt.Sprintf("%s - %s", t.Name(), "Delete the fully configured local user by removing it from configuration")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s11.Name,
						configArgs,
						[]interface{}{requiredUser},
						&tfConfig,
					)
					s11.Cfg = *tfConfig
					return s11.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s11.Index, s11.Name, s11.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					localUserStateChecks("nd_local_user.required_user", *requiredUser)...,
				),
				PostApplyFunc: func() {
					assertLocalUserAbsentOutsideTerraform(t, fullLoginID)
				},
			},
			{
				Config: func() string {
					s12.Index = 12
					s12.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy the remaining local user")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "required_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s12.Name,
						configArgs,
						[]interface{}{requiredUser},
						&tfConfig,
					)
					s12.Cfg = *tfConfig
					return s12.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s12.Index, s12.Name, s12.Cfg) },
				Destroy:   true,
				PostApplyFunc: func() {
					assertLocalUserAbsentOutsideTerraform(t, requiredLoginID)
				},
			},
			{
				Config: func() string {
					s13.Index = 13
					s13.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an unsupported local-user role")

					generateRequiredLocalUserObject(
						&invalidUser,
						invalidLoginID,
						luCfg.UserPassword,
						map[string][]string{"all": {"unsupported-role"}},
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "invalid_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s13.Name,
						configArgs,
						[]interface{}{invalidUser},
						&tfConfig,
					)
					s13.Cfg = *tfConfig
					return s13.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s13.Index, s13.Name, s13.Cfg) },
				ExpectError: regexp.MustCompile(
					`(?is)Invalid Attribute Value Match.*security_domains.*roles.*value\s+must\s+be\s+one\s+of:.*got:\s+"unsupported-role"`,
				),
			},
		},
	})

	missingDestroyUser := new(resource_local_user.NDFCLocalUserModel)
	s14 := &helper.StepInfo{}
	s15 := &helper.StepInfo{}

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
					s14.Index = 14
					s14.Name = fmt.Sprintf("%s - %s", t.Name(), "Create a local user for already-missing destroy verification")

					generateRequiredLocalUserObject(
						&missingDestroyUser,
						missingDestroyLoginID,
						luCfg.UserPassword,
						securityDomains,
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "missing_destroy_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s14.Name,
						configArgs,
						[]interface{}{missingDestroyUser},
						&tfConfig,
					)
					s14.Cfg = *tfConfig
					return s14.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s14.Index, s14.Name, s14.Cfg)
					deleteLocalUserOutsideTerraform(t, missingDestroyLoginID)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						localUserStateChecks("nd_local_user.missing_destroy_user", *missingDestroyUser),
						localUserOptionalStringsAbsentChecks("nd_local_user.missing_destroy_user")...,
					)...,
				),
			},
			{
				Config: func() string {
					s15.Index = 15
					s15.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy successfully when the Terraform-managed local user is already missing remotely")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "missing_destroy_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s15.Name,
						configArgs,
						[]interface{}{missingDestroyUser},
						&tfConfig,
					)
					s15.Cfg = *tfConfig
					return s15.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s15.Index, s15.Name, s15.Cfg)
					deleteLocalUserOutsideTerraform(t, missingDestroyLoginID)
				},
				Destroy: true,
				PostApplyFunc: func() {
					assertLocalUserAbsentOutsideTerraform(t, missingDestroyLoginID)
				},
			},
		},
	})
}

// TestAccLocalUserResourceImport verifies API-backed import state, the
// write-only password difference after import, import destroy, and missing
// object handling.
func TestAccLocalUserResourceImport(t *testing.T) {
	cfg := helper.GetConfig("global")
	luCfg := cfg.ND.LocalUser
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)

	loginID := fmt.Sprintf("%s_import_%s", luCfg.LoginID, suffix)
	missingLoginID := fmt.Sprintf("%s_missing_import_%s", luCfg.LoginID, suffix)
	securityDomains := copyLocalUserSecurityDomains(luCfg.SecurityDomains)
	tenantDomain := luCfg.TenantDomain
	if tenantDomain == "" {
		tenantDomain = localUserDefaultTenantDomain
	}
	remoteIDClaim := luCfg.RemoteIDClaim
	if remoteIDClaim == "" {
		remoteIDClaim = "tf_remote_id_claim"
	}
	email := luCfg.Email
	if email == "" {
		email = "tf_local_user@example.com"
	}
	firstName := luCfg.FirstName
	if firstName == "" {
		firstName = "Terraform"
	}
	lastName := luCfg.LastName
	if lastName == "" {
		lastName = "User"
	}

	importUser := new(resource_local_user.NDFCLocalUserModel)
	missingUser := new(resource_local_user.NDFCLocalUserModel)
	cleanupEnabled := false
	t.Cleanup(func() {
		if !cleanupEnabled {
			return
		}
		deleteLocalUserOutsideTerraform(t, loginID)
		deleteLocalUserOutsideTerraform(t, missingLoginID)
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(luCfg.LoginID) == "" {
				t.Fatal("local_user.login_id must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(luCfg.UserPassword) == "" {
				t.Fatal("local_user.user_password must be configured in the acceptance-test testbed")
			}
			if len(securityDomains) == 0 {
				t.Fatal("local_user.security_domains must contain at least one security domain in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Import an existing local user without an API-returned password")

					helper.GenerateLocalUserObject(
						&importUser,
						loginID,
						luCfg.UserPassword,
						securityDomains,
						map[string]interface{}{
							"email":                     email,
							"first_name":                firstName,
							"last_name":                 lastName,
							"remote_id_claim":           fmt.Sprintf("%s_import_%s", remoteIDClaim, suffix),
							"remote_user_authorization": true,
							"tenant_domain":             tenantDomain,
						},
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "import_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s1.Name,
						configArgs,
						[]interface{}{importUser},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteLocalUserOutsideTerraform(t, loginID)
					createLocalUserOutsideTerraform(t, *importUser)
				},
				ResourceName:       "nd_local_user.import_user",
				ImportState:        true,
				ImportStateId:      loginID,
				ImportStatePersist: true,
				ImportStateCheck:   localUserImportStateCheck(*importUser),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Detect the configured password as a plan change after import")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "import_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s2.Name,
						configArgs,
						[]interface{}{importUser},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig:          func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy the imported local user")

					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "import_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s3.Name,
						configArgs,
						[]interface{}{importUser},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Destroy:   true,
				PostApplyFunc: func() {
					assertLocalUserAbsentOutsideTerraform(t, loginID)
				},
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject import of a missing local user")

					generateRequiredLocalUserObject(
						&missingUser,
						missingLoginID,
						luCfg.UserPassword,
						securityDomains,
					)
					configArgs := map[string]string{
						"RscType":  "nd_local_user",
						"RscName":  "import_user",
						"User":     cfg.ND.User,
						"Password": cfg.ND.Password,
						"Host":     cfg.ND.URL,
						"Insecure": cfg.ND.Insecure,
					}
					tfConfig := new(string)
					helper.GetTFConfigWithSingleResource(
						s4.Name,
						configArgs,
						[]interface{}{missingUser},
						&tfConfig,
					)
					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
					deleteLocalUserOutsideTerraform(t, missingLoginID)
				},
				ResourceName:  "nd_local_user.import_user",
				ImportState:   true,
				ImportStateId: missingLoginID,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(
						`(?is)Could\s+not\s+import\s+nd\s+local\s+user\s+with\s+id\s+%q:\s+resource\s+not\s+found`,
						missingLoginID,
					),
				),
			},
		},
	})
}

func generateRequiredLocalUserObject(
	obj **resource_local_user.NDFCLocalUserModel,
	loginID string,
	userPassword string,
	securityDomains map[string][]string,
) {
	helper.GenerateLocalUserObject(obj, loginID, userPassword, securityDomains, nil)
	(*obj).Email = ""
	(*obj).FirstName = ""
	(*obj).LastName = ""
	(*obj).RemoteUserAuthorization = nil
}

func copyLocalUserSecurityDomains(in map[string][]string) map[string][]string {
	out := make(map[string][]string, len(in))
	for domainName, roles := range in {
		out[domainName] = append([]string(nil), roles...)
	}
	return out
}

func localUserStateChecks(resourceName string, model resource_local_user.NDFCLocalUserModel) []resource.TestCheckFunc {
	checks := LocalUserModelHelperStateCheck(resourceName, model, path.Empty())
	checks = append(checks,
		resource.TestCheckResourceAttr(resourceName, "id", model.LoginId),
		resource.TestCheckResourceAttr(resourceName, "security_domains.%", fmt.Sprintf("%d", len(model.Rbac.SecurityDomains))),
	)

	for domainName, securityDomain := range model.Rbac.SecurityDomains {
		rolesPath := fmt.Sprintf("security_domains.%s.roles", domainName)
		checks = append(checks,
			resource.TestCheckResourceAttr(resourceName, rolesPath+".#", fmt.Sprintf("%d", len(securityDomain.Roles))),
		)
		for _, role := range securityDomain.Roles {
			checks = append(checks,
				resource.TestCheckTypeSetElemAttr(resourceName, rolesPath+".*", role),
			)
		}
	}

	return checks
}

func localUserOptionalStringsAbsentChecks(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckNoResourceAttr(resourceName, "email"),
		resource.TestCheckNoResourceAttr(resourceName, "first_name"),
		resource.TestCheckNoResourceAttr(resourceName, "last_name"),
		resource.TestCheckNoResourceAttr(resourceName, "remote_id_claim"),
	}
}

func localUserImportStateCheck(expected resource_local_user.NDFCLocalUserModel) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		if len(states) != 1 {
			return fmt.Errorf("expected one imported local-user state, got %d", len(states))
		}

		state := states[0]
		if state.ID != expected.LoginId {
			return fmt.Errorf("imported local-user ID: expected %q, got %q", expected.LoginId, state.ID)
		}

		expectedRemoteAuthorization := false
		if expected.RemoteUserAuthorization != nil {
			expectedRemoteAuthorization = *expected.RemoteUserAuthorization
		}
		expectedTenantDomain := expected.Rbac.TenantDomain
		if expectedTenantDomain == "" {
			expectedTenantDomain = localUserDefaultTenantDomain
		}
		expectedAttributes := map[string]string{
			"id":                        expected.LoginId,
			"login_id":                  expected.LoginId,
			"email":                     expected.Email,
			"first_name":                expected.FirstName,
			"last_name":                 expected.LastName,
			"remote_id_claim":           expected.RemoteIdClaim,
			"remote_user_authorization": fmt.Sprintf("%t", expectedRemoteAuthorization),
			"tenant_domain":             expectedTenantDomain,
			"security_domains.%":        fmt.Sprintf("%d", len(expected.Rbac.SecurityDomains)),
		}
		for attributeName, expectedValue := range expectedAttributes {
			if actualValue := state.Attributes[attributeName]; actualValue != expectedValue {
				return fmt.Errorf("imported local-user attribute %q: expected %q, got %q", attributeName, expectedValue, actualValue)
			}
		}
		if password := state.Attributes["user_password"]; password != "" {
			return fmt.Errorf("imported local-user password must be absent because the API does not return it")
		}

		for domainName, securityDomain := range expected.Rbac.SecurityDomains {
			rolesPrefix := fmt.Sprintf("security_domains.%s.roles.", domainName)
			rolesCountAttribute := rolesPrefix + "#"
			if actualCount := state.Attributes[rolesCountAttribute]; actualCount != fmt.Sprintf("%d", len(securityDomain.Roles)) {
				return fmt.Errorf("imported local-user security domain %q role count: expected %d, got %q", domainName, len(securityDomain.Roles), actualCount)
			}

			for _, expectedRole := range securityDomain.Roles {
				roleFound := false
				for attributeName, actualRole := range state.Attributes {
					if attributeName != rolesCountAttribute && strings.HasPrefix(attributeName, rolesPrefix) && actualRole == expectedRole {
						roleFound = true
						break
					}
				}
				if !roleFound {
					return fmt.Errorf("imported local-user security domain %q is missing role %q", domainName, expectedRole)
				}
			}
		}

		return nil
	}
}

func newLocalUserTestClient(t *testing.T) *nd.Client {
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
		t.Fatalf("failed to create ND client for local-user acceptance helper: %v", err)
	}

	return &client
}

func createLocalUserOutsideTerraform(t *testing.T, model resource_local_user.NDFCLocalUserModel) {
	t.Helper()

	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("failed to marshal local-user create payload for %q: %v", model.LoginId, err)
	}

	client := newLocalUserTestClient(t)
	localUserAPI := api.NewLocalUserAPI(client)
	localUserAPI.LoginId = model.LoginId
	res, err := localUserAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to create local user %q outside Terraform: %v %s", model.LoginId, err, res.String())
	}
}

func updateLocalUserOutsideTerraform(t *testing.T, model resource_local_user.NDFCLocalUserModel) {
	t.Helper()

	model.UserPassword = ""
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("failed to marshal local-user update payload for %q: %v", model.LoginId, err)
	}

	client := newLocalUserTestClient(t)
	localUserAPI := api.NewLocalUserAPI(client)
	localUserAPI.LoginId = model.LoginId
	res, err := localUserAPI.Put(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to update local user %q outside Terraform: %v %s", model.LoginId, err, res.String())
	}
}

func deleteLocalUserOutsideTerraform(t *testing.T, loginID string) {
	t.Helper()

	client := newLocalUserTestClient(t)
	localUserAPI := api.NewLocalUserAPI(client)
	localUserAPI.LoginId = loginID
	res, err := localUserAPI.Delete(nil)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to delete local user %q outside Terraform: %v %s", loginID, err, res.String())
	}
}

func assertLocalUserAbsentOutsideTerraform(t *testing.T, loginID string) {
	t.Helper()

	client := newLocalUserTestClient(t)
	localUserAPI := api.NewLocalUserAPI(client)
	localUserAPI.LoginId = loginID
	respData, err := localUserAPI.Get()
	if err == nil {
		t.Fatalf("local user %q still exists in Nexus Dashboard: %s", loginID, string(respData))
	}
	if !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to verify local user %q is absent from Nexus Dashboard: %v %s", loginID, err, string(respData))
	}
}
