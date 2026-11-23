// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"regexp"
	"testing"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccManageSwitchConfigDiffRead(test *testing.T) {
	cfg := switchConfigDiffAcceptanceConfig(test)
	if cfg.ND.Fabric == "" || len(cfg.ND.Switches) == 0 || cfg.ND.Switches[0] == "" {
		test.Skip("Need nd.fabric and nd.switches in the testbed config")
	}
	x := switchConfigDiffProviderConfig(cfg)
	tfConfig := new(string)

	resource.Test(test, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(test, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GetTFConfigWithSingleResource(test.Name(), *x, []interface{}{
						helper.SwitchConfigDiffDataSource("test", cfg.ND.Fabric, cfg.ND.Switches[0]),
					}, &tfConfig)
					return *tfConfig
				}(),
			},
		},
	})
}

func TestAccManageSwitchConfigDiffInvalidFabric(test *testing.T) {
	cfg := switchConfigDiffAcceptanceConfig(test)
	if cfg.ND.Fabric == "" || len(cfg.ND.Switches) == 0 || cfg.ND.Switches[0] == "" {
		test.Skip("Need nd.fabric and nd.switches in the testbed config")
	}
	fabricName := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	x := switchConfigDiffProviderConfig(cfg)
	tfConfig := new(string)

	resource.Test(test, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(test, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GetTFConfigWithSingleResource(test.Name(), *x, []interface{}{
						helper.SwitchConfigDiffDataSource("test", fabricName, cfg.ND.Switches[0]),
					}, &tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Error Reading Switch Configuration Diff`),
			},
		},
	})
}

func TestAccManageSwitchConfigDiffInvalidSwitch(test *testing.T) {
	cfg := switchConfigDiffAcceptanceConfig(test)
	if cfg.ND.Fabric == "" || len(cfg.ND.Switches) == 0 || cfg.ND.Switches[0] == "" {
		test.Skip("Need nd.fabric and nd.switches in the testbed config")
	}
	switchID := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	x := switchConfigDiffProviderConfig(cfg)
	tfConfig := new(string)

	resource.Test(test, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(test, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GetTFConfigWithSingleResource(test.Name(), *x, []interface{}{
						helper.SwitchConfigDiffDataSource("test", cfg.ND.Fabric, switchID),
					}, &tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Error Reading Switch Configuration Diff`),
			},
		},
	})
}

func TestAccManageSwitchConfigDiffInvalidFabricAndSwitch(test *testing.T) {
	cfg := switchConfigDiffAcceptanceConfig(test)
	if cfg.ND.Fabric == "" {
		test.Skip("Need nd.fabric in the testbed config")
	}
	fabricName := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	switchID := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	x := switchConfigDiffProviderConfig(cfg)
	tfConfig := new(string)

	resource.Test(test, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(test, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GetTFConfigWithSingleResource(test.Name(), *x, []interface{}{
						helper.SwitchConfigDiffDataSource("test", fabricName, switchID),
					}, &tfConfig)
					return *tfConfig
				}(),
				ExpectError: regexp.MustCompile(`Error Reading Switch Configuration Diff`),
			},
		},
	})
}

func switchConfigDiffAcceptanceConfig(test *testing.T) *helper.Config {
	test.Helper()
	if os.Getenv("TF_ACC") == "" {
		test.Skip("Set TF_ACC to run acceptance tests")
	}
	return helper.GetConfig("global")
}

func switchConfigDiffProviderConfig(config *helper.Config) *map[string]string {
	return &map[string]string{
		"RscType":  "nd_switch_config_diff",
		"RscName":  "test",
		"User":     config.ND.User,
		"Password": config.ND.Password,
		"Host":     config.ND.URL,
		"Insecure": config.ND.Insecure,
	}
}
