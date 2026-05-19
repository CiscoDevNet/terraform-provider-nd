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
	stepCount := new(int)
	*stepCount = 0

	user1 := new(resource_local_user.NDFCLocalUserModel)
	user2 := new(resource_local_user.NDFCLocalUserModel)

	domains := luCfg.SecurityDomains
	if len(domains) == 0 {
		domains = map[string][]string{"all": {"approver", "designer"}}
	}

	loginIDA := luCfg.LoginID + "_a"
	loginIDB := luCfg.LoginID + "_b"

	// Per-step info objects captured by both the Config builder (write) and
	// PreConfig logger (read). Declared up-front so each step's closures
	// reference an independent struct.
	s1 := &stepInfo{}
	s2 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create both users. Neither carries remote_id_claim or
			// tenant_domain at this point.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_both_users", t.Name(), *stepCount)

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

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{user1, user2}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
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
			// Step 2: Update user_one to add remote_id_claim and
			// tenant_domain. user_two is re-rendered unchanged.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_update_user_one_add_remote_id_claim_and_tenant_domain", t.Name(), *stepCount)

					helper.ModifyLocalUserObject(&user1, map[string]interface{}{
						"remote_id_claim": "tf_remote_id_claim_a",
						"tenant_domain":   "all-tenants-domain",
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{user1, user2}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
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
			// Step 3: ImportState verification for user_one. The API GET
			// does not echo write-only / non-returned fields, so they are
			// excluded from the ImportStateVerify diff comparison:
			//   - user_password (write-only / sensitive)
			//   - tenant_domain (not returned)
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_import_user_one =====", t.Name())
				},
				ResourceName:                         "nd_local_user.user_one",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        loginIDA,
				ImportStateVerifyIdentifierAttribute: "login_id",
				ImportStateVerifyIgnore: []string{
					"user_password",
					"tenant_domain",
				},
			},
		},
	})
}

// TestAccLocalUserResourceCRUD exercises create, modify (multiple field
// updates), security-domain add/remove, and final delete of a local user.
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
	stepCount := new(int)
	*stepCount = 0

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

	s1 := &stepInfo{}
	s2 := &stepInfo{}
	s3 := &stepInfo{}
	s4 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create local user with defaults from helper
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_local_user", t.Name(), *stepCount)

					helper.GenerateLocalUserObject(&userRsc,
						loginID,
						userPassword,
						initialDomains,
						map[string]interface{}{
							"remote_id_claim": "tf_remote_id_claim_crud",
							"tenant_domain":   "all-tenants-domain",
						},
					)

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(
						"nd_local_user.user_test",
						*userRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Modify scalar fields (names, email, limits)
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_modify_scalars", t.Name(), *stepCount)

					helper.ModifyLocalUserObject(&userRsc, map[string]interface{}{
						"first_name": "updated_first",
						"last_name":  "updated_last",
						"email":      "updated_user@mail.com",
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(
						"nd_local_user.user_test",
						*userRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 3: Toggle remote_user_authorization and extend the roles
			// of the existing security domain.
			{
				Config: func() string {
					*stepCount++
					s3.name = fmt.Sprintf("%s_%d_toggle_xlaunch_extend_roles", t.Name(), *stepCount)

					helper.ModifyLocalUserObject(&userRsc, map[string]interface{}{
						"remote_user_authorization": true,
					})
					helper.AddSecurityDomain(&userRsc, "all",
						[]string{"approver", "designer", "observer"},
					)

					helper.GetTFConfigWithSingleResource(s3.name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s3.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 3, s3.name, s3.cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(
						"nd_local_user.user_test",
						*userRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 4: Shrink roles back to the original set on the same
			// security domain.
			{
				Config: func() string {
					*stepCount++
					s4.name = fmt.Sprintf("%s_%d_shrink_roles", t.Name(), *stepCount)

					helper.AddSecurityDomain(&userRsc, "all",
						[]string{"approver", "designer"},
					)

					helper.GetTFConfigWithSingleResource(s4.name, *x,
						[]interface{}{userRsc}, &tfConfig)

					s4.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 4, s4.name, s4.cfg) },
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
