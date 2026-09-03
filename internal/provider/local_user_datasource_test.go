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

	"terraform-provider-nd/internal/infra/resource_local_user"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLocalUserDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	localUserConfig := cfg.ND.LocalUser
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	loginID := fmt.Sprintf("%s_ds_%s", localUserConfig.LoginID, suffix)
	missingID := loginID + "_missing"
	defaultTenantDomain := "all-tenants-domain"
	resourceName := "nd_local_user.user_test"
	dataSourceName := "data.nd_local_user.user_test"

	securityDomains := localUserConfig.SecurityDomains
	if len(securityDomains) == 0 {
		securityDomains = map[string][]string{
			"all": {"approver", "designer"},
		}
	}

	x := &map[string]string{
		"RscType":  "nd_local_user",
		"RscName":  "user_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	localUser := new(resource_local_user.NDFCLocalUserModel)
	localUserDataSource := &helper.LocalUserDataSourceTestData{
		RscName:   "user_test",
		LoginId:   loginID,
		DependsOn: resourceName,
	}
	missingLocalUserDataSource := &helper.LocalUserDataSourceTestData{
		RscName: "missing_user",
		LoginId: missingID,
	}
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
		resource.TestCheckResourceAttrPair(dataSourceName, "login_id", resourceName, "login_id"),
		resource.TestCheckResourceAttrPair(dataSourceName, "email", resourceName, "email"),
		resource.TestCheckResourceAttrPair(dataSourceName, "first_name", resourceName, "first_name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "last_name", resourceName, "last_name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "remote_id_claim", resourceName, "remote_id_claim"),
		resource.TestCheckResourceAttrPair(dataSourceName, "remote_user_authorization", resourceName, "remote_user_authorization"),
		resource.TestCheckResourceAttr(dataSourceName, "tenant_domain", defaultTenantDomain),
		resource.TestCheckResourceAttrPair(dataSourceName, "tenant_domain", resourceName, "tenant_domain"),
		resource.TestCheckResourceAttr(dataSourceName, "security_domains.%", strconv.Itoa(len(securityDomains))),
		// The local-user GET API does not return the write-only password.
		resource.TestCheckNoResourceAttr(dataSourceName, "user_password"),
	}
	for domainName, roles := range securityDomains {
		rolesPath := fmt.Sprintf("security_domains.%s.roles", domainName)
		matchingChecks = append(
			matchingChecks,
			resource.TestCheckResourceAttrPair(dataSourceName, rolesPath+".#", resourceName, rolesPath+".#"),
		)
		for _, role := range roles {
			matchingChecks = append(
				matchingChecks,
				resource.TestCheckTypeSetElemAttr(dataSourceName, rolesPath+".*", role),
				resource.TestCheckTypeSetElemAttr(resourceName, rolesPath+".*", role),
			)
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the local user used by datasource lookups")

					helper.GenerateLocalUserObject(
						&localUser,
						loginID,
						localUserConfig.UserPassword,
						securityDomains,
						map[string]interface{}{
							"email":                     fmt.Sprintf("tf_local_user_ds_%s@mail.com", suffix),
							"first_name":                "tf_datasource_first",
							"last_name":                 "tf_datasource_last",
							"remote_id_claim":           fmt.Sprintf("tf_remote_id_claim_ds_%s", suffix),
							"remote_user_authorization": false,
						},
					)
					helper.GetTFConfigWithSingleResource(s1.Name, *x, []interface{}{localUser}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					LocalUserModelHelperStateCheck(resourceName, *localUser, path.Empty())...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Read the existing local user through the datasource")

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{localUser, localUserDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "id", loginID),
					resource.TestCheckResourceAttr(dataSourceName, "login_id", loginID),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Match datasource attributes with the configured local user")

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						*x,
						[]interface{}{localUser, localUserDataSource},
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
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject datasource lookup for a missing local user")

					helper.GetTFConfigWithSingleResource(
						s4.Name,
						*x,
						[]interface{}{missingLocalUserDataSource},
						&missingTFConfig,
					)
					s4.Cfg = *missingTFConfig
					return *missingTFConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not read nd local user with id\s+%q:\s+resource not found`, missingID),
				),
			},
		},
	})
}
