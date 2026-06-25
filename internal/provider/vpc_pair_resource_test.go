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
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_vpc_pair"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVpcPairResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")
	leafSwitches := make([]string, 0, len(cfg.ND.Inventory.Switches))
	for _, sw := range cfg.ND.Inventory.Switches {
		if strings.EqualFold(sw.Role, "leaf") {
			leafSwitches = append(leafSwitches, sw.Serial)
		}
	}

	if len(leafSwitches) < 2 {
		t.Skip("Need at least 2 leaf switches in nd.inventory.switches for vPC pair acceptance test")
	}

	x := &map[string]string{
		"RscType":  "nd_vpc_pair",
		"RscName":  "vpc_pair_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	vpcPairRsc := new(resource_vpc_pair.NDFCVpcPairModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateVpcPairObject(
						&vpcPairRsc,
						cfg.ND.Fabric,
						leafSwitches[0],
						leafSwitches[1],
						false,
						false,
					)

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{vpcPairRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyVpcPairObject(&vpcPairRsc, map[string]interface{}{
						"deploy": true,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{vpcPairRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					VpcPairModelHelperStateCheck(
						"nd_vpc_pair.vpc_pair_test",
						*vpcPairRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
