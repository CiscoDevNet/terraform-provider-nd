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

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInventorySwitchDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	datasourceConfig := cfg.ND.InventorySwitchDataSource
	if datasourceConfig.Fabric == "" {
		t.Skip("Need an inventory switch datasource fabric in the testbed config")
	}
	if datasourceConfig.SerialNumber == "" {
		t.Skip("Need an inventory switch datasource serial number in the testbed config")
	}

	const (
		allDataSource    = "data.nd_inventory_switch.all"
		serialDataSource = "data.nd_inventory_switch.by_serial"
	)

	terraformConfig := fmt.Sprintf(`
provider "nd" {
  username = %q
  password = %q
  url      = %q
  insecure = %t
}

data "nd_inventory_switch" "all" {
  fabric_name = %q
}

data "nd_inventory_switch" "by_serial" {
  fabric_name   = %q
  serial_number = %q
}
`, cfg.ND.User, cfg.ND.Password, cfg.ND.URL, cfg.ND.Insecure == "true", datasourceConfig.Fabric, datasourceConfig.Fabric, datasourceConfig.SerialNumber)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: terraformConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(allDataSource, "fabric_name", datasourceConfig.Fabric),
					resource.TestCheckResourceAttr(allDataSource, "switches.#", "3"),
					resource.TestCheckResourceAttr(serialDataSource, "fabric_name", datasourceConfig.Fabric),
					resource.TestCheckResourceAttr(serialDataSource, "serial_number", datasourceConfig.SerialNumber),
					resource.TestCheckResourceAttr(serialDataSource, "switches.#", "1"),
					resource.TestCheckResourceAttr(serialDataSource, "switches.0.serial_number", datasourceConfig.SerialNumber),
				),
			},
		},
	})
}
