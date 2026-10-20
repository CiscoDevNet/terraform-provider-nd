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

	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMultiClusterConnectivityDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	mccCfg := cfg.ND.MultiClusterConnectivity

	resourceName := "nd_multi_cluster_connectivity.cluster_test"
	dataSourceName := "data.nd_multi_cluster_connectivity.cluster_test"

	x := &map[string]string{
		"RscType":  "nd_multi_cluster_connectivity",
		"RscName":  "cluster_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	cluster := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)
	helper.GenerateMultiClusterConnectivityObject(
		&cluster,
		mccCfg.Hostname,
		cfg.ND.User,
		cfg.ND.Password,
		map[string]interface{}{
			"login_domain":             mccCfg.LoginDomain,
			"multi_cluster_login_domain": mccCfg.MultiClusterLoginDomain,
		},
	)

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}

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
						"Create the cluster used by the datasource",
					)

					helper.GetTFConfigWithSingleResource(
						s1.Name,
						*x,
						[]interface{}{cluster},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteMultiClusterConnectivityIfExistsOutsideTerraform(t, cluster)
				},
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(resourceName, *cluster, path.Empty())...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Read the existing cluster through the datasource",
					)

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{cluster},
						&tfConfig,
					)
					s2Config := *tfConfig
					s2Config += fmt.Sprintf(
						"\ndata \"nd_multi_cluster_connectivity\" \"cluster_test\" {\n"+
							"  cluster_name = %s.cluster_name\n"+
							"}\n",
						resourceName,
					)
					s2.Cfg = s2Config
					*tfConfig = s2Config
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_name", resourceName, "cluster_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hostname", resourceName, "hostname"),
				),
			},
		},
	})
}

func TestAccMultiClusterConnectivityDataSourceNotFound(t *testing.T) {
	cfg := helper.GetConfig("global")
	missingClusterName := fmt.Sprintf(
		"tf_multi_cluster_connectivity_missing_%s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlpha),
	)

	s1 := &helper.StepInfo{}

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
						"Reject a datasource lookup for a missing cluster",
					)
					s1.Cfg = fmt.Sprintf(
						"provider \"nd\" {\n"+
							"  username = %q\n"+
							"  password = %q\n"+
							"  url      = %q\n"+
							"  insecure = %s\n"+
							"}\n\n"+
							"data \"nd_multi_cluster_connectivity\" \"missing\" {\n"+
							"  cluster_name = %q\n"+
							"}\n",
						cfg.ND.User,
						cfg.ND.Password,
						cfg.ND.URL,
						cfg.ND.Insecure,
						missingClusterName,
					)
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(
						`Could not read nd multi-cluster connectivity with cluster_name\s+%q:\s+resource\s+not\s+found`,
						missingClusterName,
					),
				),
			},
		},
	})
}
