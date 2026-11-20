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
	"os"
	"regexp"
	"strconv"
	"testing"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSwitchConfigDiffDataSource(test *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		test.Skip("Set TF_ACC to run acceptance tests")
	}
	config := helper.GetConfig("global")
	if config.ND.Fabric == "" || len(config.ND.Switches) == 0 || config.ND.Switches[0] == "" {
		test.Skip("Need nd.fabric and nd.switches in the testbed config")
	}
	validSwitchID := config.ND.Switches[0]
	invalidFabric := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	invalidSwitchID := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	const sourceName = "data.nd_switch_config_diff.test"
	testCases := []struct {
		name       string
		fabricName string
		switchID   string
		wantError  *regexp.Regexp
	}{
		{name: "valid fabric and switch", fabricName: config.ND.Fabric, switchID: validSwitchID},
		{name: "invalid fabric", fabricName: invalidFabric, switchID: validSwitchID, wantError: regexp.MustCompile(`Error Reading Switch Configuration Diff`)},
		{name: "invalid switch", fabricName: config.ND.Fabric, switchID: invalidSwitchID, wantError: regexp.MustCompile(`Error Reading Switch Configuration Diff`)},
		{name: "invalid fabric and switch", fabricName: invalidFabric, switchID: invalidSwitchID, wantError: regexp.MustCompile(`Error Reading Switch Configuration Diff`)},
	}
	for _, testCase := range testCases {
		testCase := testCase
		test.Run(testCase.name, func(test *testing.T) {
			step := resource.TestStep{
				Config: fmt.Sprintf(`
provider "nd" {
  username = %q
  password = %q
  url      = %q
  insecure = %t
}

data "nd_switch_config_diff" "test" {
  fabric_name = %q
  switch_id   = %q
}
				`, config.ND.User, config.ND.Password, config.ND.URL, config.ND.Insecure == "true", testCase.fabricName, testCase.switchID),
				ExpectError: testCase.wantError,
			}
			if testCase.wantError == nil {
				step.Check = resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(sourceName, "fabric_name", testCase.fabricName),
					resource.TestCheckResourceAttr(sourceName, "switch_id", testCase.switchID),
					func(state *terraform.State) error {
						saved, exists := state.RootModule().Resources[sourceName]
						if !exists || saved.Primary == nil {
							return fmt.Errorf("missing state for %s", sourceName)
						}
						countValue := saved.Primary.Attributes["diff_configs.#"]
						if countValue == "" {
							return nil
						}
						count, err := strconv.Atoi(countValue)
						if err != nil || count < 0 {
							return fmt.Errorf("invalid diff count %q", countValue)
						}
						for index := 0; index < count; index++ {
							attribute := fmt.Sprintf("diff_configs.%d.difference_type", index)
							differenceType, present := saved.Primary.Attributes[attribute]
							if !present {
								return fmt.Errorf("missing %s", attribute)
							}
							switch differenceType {
							case "equal", "insert", "delete":
							default:
								return fmt.Errorf("invalid %s: %q", attribute, differenceType)
							}
						}
						return nil
					},
				)
			}
			resource.Test(test, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(test, "global") },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    []resource.TestStep{step},
			})
		})
	}
}
