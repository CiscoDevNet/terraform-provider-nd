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

	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccMultiClusterConnectivityResourceCRUD exercises create/read/update
// against a remote ND cluster onboarded via nd_multi_cluster_connectivity.
//
// Required testbed.yaml settings (under `nd.multi_cluster_connectivity`):
//
//	cluster_name, hostname, username, password
//
// Optional: login_domain, multi_cluster_login_domain.
func TestAccMultiClusterConnectivityResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	mccCfg := cfg.ND.MultiClusterConnectivity

	if mccCfg.Hostname == "" || mccCfg.Username == "" || mccCfg.Password == "" {
		t.Skip("nd.multi_cluster_connectivity.{hostname,username,password} required in testbed.yaml")
	}

	x := &map[string]string{
		"RscType":  "nd_multi_cluster_connectivity",
		"RscName":  "rsc_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	rscAddr := "nd_multi_cluster_connectivity.rsc_test"

	tfConfig := new(string)
	stepCount := new(int)

	rsc := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)

	// Per-step info objects captured by both the Config builder (write) and
	// PreConfig logger (read). Declared up-front so each step's closures
	// reference an independent struct.
	s1 := &stepInfo{}
	s2 := &stepInfo{}
	s4 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with required attributes only.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create", t.Name(), *stepCount)

					overrides := map[string]interface{}{}
					if mccCfg.ClusterName != "" {
						overrides["cluster_name"] = mccCfg.ClusterName
					}
					helper.GenerateMultiClusterConnectivityObject(&rsc,
						mccCfg.Hostname, mccCfg.Username, mccCfg.Password, overrides,
					)

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(
						rscAddr, *rsc, path.Empty(),
					)...,
				),
			},
			// Step 2: Update with login_domain and multi_cluster_login_domain.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_update_optional", t.Name(), *stepCount)

					updates := map[string]interface{}{}
					if mccCfg.LoginDomain != "" {
						updates["login_domain"] = mccCfg.LoginDomain
					}
					if mccCfg.MultiClusterLoginDomain != "" {
						updates["multi_cluster_login_domain"] = mccCfg.MultiClusterLoginDomain
					}
					helper.ModifyMultiClusterConnectivityObject(&rsc, updates)

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(
						rscAddr, *rsc, path.Empty(),
					)...,
				),
			},
			// Step 3: Import. cluster_name is the natural identifier and is
			// derived from prior state (it is Optional+Computed and may be
			// assigned by the API when not specified in the config).
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_import =====", t.Name())
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "cluster_name",
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[rscAddr]
					if !ok {
						return "", fmt.Errorf("resource %s not found in state", rscAddr)
					}
					name := rs.Primary.Attributes["cluster_name"]
					if name == "" {
						return "", fmt.Errorf("cluster_name not set in state for %s", rscAddr)
					}
					return name, nil
				},
				ImportStateVerifyIgnore: []string{
					"username",
					"password",
					"login_domain",
					"multi_cluster_login_domain",
				},
			},
			// Step 4: Explicit destroy. Re-uses the last rendered config
			// and tears down the resource; the framework also performs an
			// implicit destroy at the end of the case, but this step makes
			// destroy failures attributable to a specific step.
			{
				PreConfig: func() {
					s4.name = fmt.Sprintf("%s_4_destroy", t.Name())
					s4.cfg = *tfConfig
					helper.LogStep(t, 4, s4.name, s4.cfg)
				},
				Config:  *tfConfig,
				Destroy: true,
			},
		},
	})
}
