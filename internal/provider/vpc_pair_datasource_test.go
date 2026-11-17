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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVpcPairDataSource(t *testing.T) {
	vpcPairTestctx := loadVpcPairTestctx(t)
	requireLeafPair(t, vpcPairTestctx)

	terraformConfig := fmt.Sprintf(`
provider "nd" {
  username = %q
  password = %q
  url      = %q
  insecure = %t
}

data "nd_vpc_pair" "vpc_pair_test" {
  fabric_name = %q
  switch_id   = %q
}
`, vpcPairTestctx.configMap["User"], vpcPairTestctx.configMap["Password"],
		vpcPairTestctx.configMap["Host"], vpcPairTestctx.configMap["Insecure"] == "true",
		vpcPairTestctx.fabricName, vpcPairTestctx.leafSwitches[1])

	const dataSourceName = "data.nd_vpc_pair.vpc_pair_test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: terraformConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "fabric_name", vpcPairTestctx.fabricName),
					resource.TestCheckResourceAttr(dataSourceName, "switch_id", vpcPairTestctx.leafSwitches[1]),
					resource.TestCheckResourceAttr(dataSourceName, "peer_switch_id", vpcPairTestctx.leafSwitches[0]),
				),
			},
		},
	})
}
