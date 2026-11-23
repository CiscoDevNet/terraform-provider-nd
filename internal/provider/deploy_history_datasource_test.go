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

func TestAccManageDeployHistoryRead(test *testing.T) {
	config := deployHistoryAcceptanceConfig(test)
	runDeployHistoryAcceptanceTest(test, config, config.ND.Fabric, nil, nil)
}

func TestAccManageDeployHistoryInvalidFabric(test *testing.T) {
	config := deployHistoryAcceptanceConfig(test)
	fabricName := "tf_missing_" + acctest.RandStringFromCharSet(12, acctest.CharSetAlphaNum)
	runDeployHistoryAcceptanceTest(test, config, fabricName, nil, regexp.MustCompile(`Error Reading Deployment History`))
}

func TestAccManageDeployHistoryZeroRecordCount(test *testing.T) {
	config := deployHistoryAcceptanceConfig(test)
	zeroLimit := 0
	runDeployHistoryAcceptanceTest(test, config, config.ND.Fabric, &zeroLimit, regexp.MustCompile(`(?is)(no_of_records.*(between|at least|greater than or equal to).*1|Invalid Record Count.*at least 1)`))
}

func TestAccManageDeployHistoryValidRecordCount(test *testing.T) {
	config := deployHistoryAcceptanceConfig(test)
	validLimit := 5
	runDeployHistoryAcceptanceTest(test, config, config.ND.Fabric, &validLimit, nil)
}

func TestAccManageDeployHistoryRecordCountOverLimit(test *testing.T) {
	config := deployHistoryAcceptanceConfig(test)
	overLimit := 51
	runDeployHistoryAcceptanceTest(test, config, config.ND.Fabric, &overLimit, regexp.MustCompile(`(?is)(no_of_records|invalid record count).*50`))
}

func deployHistoryAcceptanceConfig(test *testing.T) *helper.Config {
	test.Helper()
	if os.Getenv("TF_ACC") == "" {
		test.Skip("Set TF_ACC to run acceptance tests")
	}
	config := helper.GetConfig("global")
	if config.ND.Fabric == "" {
		test.Skip("Need nd.fabric in the testbed config")
	}
	return config
}

func runDeployHistoryAcceptanceTest(test *testing.T, config *helper.Config, fabricName string, recordLimit *int, expectedError *regexp.Regexp) {
	test.Helper()
	x := &map[string]string{
		"RscType":  "nd_deploy_history",
		"RscName":  "test",
		"User":     config.ND.User,
		"Password": config.ND.Password,
		"Host":     config.ND.URL,
		"Insecure": config.ND.Insecure,
	}
	tfConfig := new(string)

	resource.Test(test, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(test, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					helper.GetTFConfigWithSingleResource(test.Name(), *x, []interface{}{
						helper.DeployHistoryDataSource("test", fabricName, recordLimit),
					}, &tfConfig)
					return *tfConfig
				}(),
				ExpectError: expectedError,
			},
		},
	})
}
