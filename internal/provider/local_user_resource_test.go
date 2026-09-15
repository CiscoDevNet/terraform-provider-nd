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

	"terraform-provider-nd/internal/infra/resource_local_user"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccLocalUserMultiResource creates two independent nd_local_user
// resources in the same apply and exercises distinct lifecycles per user:
//
//	user_one: Create (no remote_id_claim, no tenant_domain) → Update (add
//	          remote_id_claim and tenant_domain) → ImportState → Destroy.
//	user_two: Create → Destroy (covers the minimal lifecycle path).
func TestAccLocalUserMultiResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	luCfg := cfg.ND.LocalUser

	x := &map[string]string{
		"RscType":  "nd_local_user",
		"RscName":  "user_one,user_two",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	missingTFConfig := new(string)

	user1 := new(resource_local_user.NDFCLocalUserModel)
	user2 := new(resource_local_user.NDFCLocalUserModel)
	missingUser := new(resource_local_user.NDFCLocalUserModel)

	domains := luCfg.SecurityDomains
	if len(domains) == 0 {
		domains = map[string][]string{"all": {"approver", "designer"}}
	}

	loginIDA := luCfg.LoginID + "_a"
	loginIDB := luCfg.LoginID + "_b"
	missingLoginID := loginIDA + "_missing"

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
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create both local users")

					helper.GenerateLocalUserObject(&user1,
						loginIDA,
						luCfg.UserPassword,
						domains,
						nil,
					)
					helper.GenerateLocalUserObject(&user2,
						loginIDB,
						luCfg.UserPassword,
						domains,
						nil,
					)

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{user1, user2}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_one",
							*user1,
							path.Empty(),
						),
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_two",
							*user2,
							path.Empty(),
						)...,
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Add a remote ID claim and tenant domain to the first local user")

					helper.ModifyLocalUserObject(&user1, map[string]interface{}{
						"remote_id_claim": "tf_remote_id_claim_a",
						"tenant_domain":   "all-tenants-domain",
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{user1, user2}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_one",
							*user1,
							path.Empty(),
						),
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_two",
							*user2,
							path.Empty(),
						)...,
					)...,
				),
			},
			{
				PreConfig: func() {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Import the first local user and verify API readback")
					s3.Cfg = *tfConfig
					helper.LogStep(t, s3.Index, s3.Name, s3.Cfg)
				},
				ResourceName:                         "nd_local_user.user_one",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        loginIDA,
				ImportStateVerifyIdentifierAttribute: "login_id",
				ImportStateVerifyIgnore: []string{
					"user_password",
				},
			},
			{
				PreConfig: func() {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject import of a missing local user")
					helper.GenerateLocalUserObject(&missingUser, missingLoginID, luCfg.UserPassword, domains, nil)
					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{missingUser}, &missingTFConfig)
					s4.Cfg = *missingTFConfig
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
				},
				ResourceName:  "nd_local_user.user_one",
				ImportState:   true,
				ImportStateId: missingLoginID,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not import nd local user with id %q:\s+resource not found`, missingLoginID),
				),
			},
		},
	})
}

// TestAccLocalUserResourceCRUD exercises required-only create, optional
// attribute configuration, clearing omitted optional strings, an empty plan
// with those attributes omitted, and re-adding the optional values.
func TestAccLocalUserResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	luCfg := cfg.ND.LocalUser

	x := &map[string]string{
		"RscType":  "nd_local_user",
		"RscName":  "user_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)

	userRsc := new(resource_local_user.NDFCLocalUserModel)

	loginID := luCfg.LoginID
	if loginID == "" {
		loginID = "tf_local_user_acc"
	}
	userPassword := luCfg.UserPassword
	if userPassword == "" {
		userPassword = "Str0ng_P@ssw0rd_123!"
	}

	initialDomains := luCfg.SecurityDomains
	if len(initialDomains) == 0 {
		initialDomains = map[string][]string{
			"all": {"approver", "designer"},
		}
	}
	optionalValues := map[string]interface{}{
		"email":                     "tf_local_user_test@example.com",
		"first_name":                "Test",
		"last_name":                 "User",
		"remote_id_claim":           "tf_remote_id_claim_crud",
		"remote_user_authorization": true,
		"tenant_domain":             "all-tenants-domain",
	}
	optionalStringsAbsentChecks := func() []resource.TestCheckFunc {
		return []resource.TestCheckFunc{
			resource.TestCheckNoResourceAttr("nd_local_user.user_test", "email"),
			resource.TestCheckNoResourceAttr("nd_local_user.user_test", "first_name"),
			resource.TestCheckNoResourceAttr("nd_local_user.user_test", "last_name"),
			resource.TestCheckNoResourceAttr("nd_local_user.user_test", "remote_id_claim"),
		}
	}

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create a local user with required attributes")

					helper.GenerateLocalUserObject(&userRsc, loginID, userPassword, initialDomains, nil)
					userRsc.Email = ""
					userRsc.FirstName = ""
					userRsc.LastName = ""
					userRsc.RemoteUserAuthorization = nil

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_test",
							*userRsc,
							path.Empty(),
						),
						optionalStringsAbsentChecks()...,
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Configure optional local-user attributes")

					helper.ModifyLocalUserObject(&userRsc, optionalValues)
					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(
						"nd_local_user.user_test",
						*userRsc,
						path.Empty(),
					)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Remove optional local-user attributes")

					helper.GenerateLocalUserObject(&userRsc, loginID, userPassword, initialDomains, nil)
					userRsc.Email = ""
					userRsc.FirstName = ""
					userRsc.LastName = ""
					userRsc.RemoteUserAuthorization = nil

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						LocalUserModelHelperStateCheck(
							"nd_local_user.user_test",
							*userRsc,
							path.Empty(),
						),
						optionalStringsAbsentChecks()...,
					)...,
				),
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan with optional local-user attributes omitted")

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Re-add optional local-user attributes and verify the object")

					helper.ModifyLocalUserObject(&userRsc, optionalValues)
					helper.GetTFConfigWithSingleResource(s5.Name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s5.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s5.Index, s5.Name, s5.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(
						"nd_local_user.user_test",
						*userRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
