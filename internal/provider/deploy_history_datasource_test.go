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

func TestAccDeployHistoryDataSource(test *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		test.Skip("Set TF_ACC to run acceptance tests")
	}
	config := helper.GetConfig("global")
	fabricName := config.ND.DeployHistoryDataSource.Fabric
	if fabricName == "" {
		fabricName = config.ND.Fabric
	}
	if fabricName == "" {
		test.Skip("Need nd.deploy_history_datasource.fabric or nd.fabric in the testbed config")
	}
	const sourceName = "data.nd_deploy_history.test"
	validLimit := 5
	zeroLimit := 0
	overLimit := 51
	missingFabric := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	testCases := []struct {
		name        string
		fabricName  string
		recordLimit *int
		wantError   *regexp.Regexp
	}{
		{name: "valid fabric", fabricName: fabricName},
		{name: "invalid fabric", fabricName: missingFabric, wantError: regexp.MustCompile(`Error Reading Deployment History`)},
		{name: "zero record count", fabricName: fabricName, recordLimit: &zeroLimit, wantError: regexp.MustCompile(`(?s)Invalid Record Count.*at least 1`)},
		{name: "valid record count", fabricName: fabricName, recordLimit: &validLimit},
		{name: "record count over 50", fabricName: fabricName, recordLimit: &overLimit, wantError: regexp.MustCompile(`(?is)(no_of_records|invalid record count).*50`)},
	}
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			recordCountConfig := ""
			if testCase.recordLimit != nil {
				recordCountConfig = fmt.Sprintf("  no_of_records = %d\n", *testCase.recordLimit)
			}
			step := resource.TestStep{
				Config: fmt.Sprintf(`
provider "nd" {
  username = %q
  password = %q
  url      = %q
  insecure = %t
}

data "nd_deploy_history" "test" {
  fabric_name = %q
%s}
`, config.ND.User, config.ND.Password, config.ND.URL, config.ND.Insecure == "true", testCase.fabricName, recordCountConfig),
				ExpectError: testCase.wantError,
			}
			if testCase.wantError == nil {
				checks := []resource.TestCheckFunc{
					resource.TestCheckResourceAttr(sourceName, "fabric_name", testCase.fabricName),
				}
				if testCase.recordLimit != nil {
					recordLimit := *testCase.recordLimit
					checks = append(checks,
						resource.TestCheckResourceAttr(sourceName, "no_of_records", strconv.Itoa(recordLimit)),
						func(state *terraform.State) error {
							saved, exists := state.RootModule().Resources[sourceName]
							if !exists || saved.Primary == nil {
								return fmt.Errorf("missing state for %s", sourceName)
							}
							countValue := saved.Primary.Attributes["deployment_records.#"]
							if countValue == "" {
								return nil
							}
							count, err := strconv.Atoi(countValue)
							if err != nil || count < 0 || count > recordLimit {
								return fmt.Errorf("expected at most %d records, got %q", recordLimit, countValue)
							}
							return nil
						},
					)
				}
				step.Check = resource.ComposeTestCheckFunc(checks...)
			}
			resource.Test(test, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(test, "global") },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    []resource.TestStep{step},
			})
		})
	}
}
