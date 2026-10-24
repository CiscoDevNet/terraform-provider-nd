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
	"strconv"
	"strings"
	"testing"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccInventorySwitchDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	// Read the existing switches directly. Discovery resources are intentionally
	// not created here because they require the switches to be reachable.
	const fabricName = "fabric1_vxlan"
	switches := []struct {
		serialNumber string
		managementIP string
	}{
		{serialNumber: "9R518K2AT3R", managementIP: "192.168.10.204"},
		{serialNumber: "9N2SFZYSDIJ", managementIP: "192.168.10.201"},
		{serialNumber: "9NXSLRWEM7J", managementIP: "192.168.10.203"},
	}
	const dataSourceName = "data.nd_inventory_switch.lookup"

	steps := []struct {
		name         string
		serialNumber string
		managementIP string
	}{
		// Fetch every switch in the existing fabric.
		{name: "fabric_only"},
		// The selector cases target the first existing switch.
		{name: "serial_only", serialNumber: switches[0].serialNumber},
		{name: "management_ip_only", managementIP: switches[0].managementIP},
		{name: "serial_and_management_ip", serialNumber: switches[0].serialNumber, managementIP: switches[0].managementIP},
	}

	resourceSteps := make([]resource.TestStep, 0, len(steps))
	for _, step := range steps {
		step := step
		tfConfig := new(string)
		helper.GetTFConfigWithSingleResource(
			"inventory_switch_datasource_"+step.name,
			map[string]string{
				"User":     cfg.ND.User,
				"Password": cfg.ND.Password,
				"Host":     cfg.ND.URL,
				"Insecure": cfg.ND.Insecure,
			},
			nil,
			&tfConfig,
		)
		*tfConfig += inventorySwitchDataSourceBlock("lookup", fabricName, step.serialNumber, step.managementIP, nil)

		checks := []resource.TestCheckFunc{}
		if step.serialNumber != "" || step.managementIP != "" {
			checks = append(
				checks,
				resource.TestCheckResourceAttr(dataSourceName, "switches.#", "1"),
				checkInventorySwitchDataSourceContains(dataSourceName, switches[0].serialNumber, switches[0].managementIP),
			)
		} else {
			checks = append(
				checks,
				resource.TestCheckResourceAttr(dataSourceName, "switches.#", "3"),
				checkInventorySwitchDataSourceContains(dataSourceName, switches[0].serialNumber, switches[0].managementIP),
				checkInventorySwitchDataSourceContains(dataSourceName, switches[1].serialNumber, switches[1].managementIP),
				checkInventorySwitchDataSourceContains(dataSourceName, switches[2].serialNumber, switches[2].managementIP),
			)
		}

		resourceSteps = append(resourceSteps, resource.TestStep{
			Config: *tfConfig,
			Check:  resource.ComposeTestCheckFunc(checks...),
		})
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    resourceSteps,
	})
}

func inventorySwitchDataSourceBlock(dataSourceName, fabricName, serialNumber, managementIP string, dependsOn []string) string {
	block := fmt.Sprintf(`
data "nd_inventory_switch" %q {
  fabric_name = %q
`, dataSourceName, fabricName)
	if len(dependsOn) > 0 {
		block += fmt.Sprintf("  depends_on = [%s]\n", strings.Join(dependsOn, ", "))
	}
	if serialNumber != "" {
		block += fmt.Sprintf("  serial_number = %q\n", serialNumber)
	}
	if managementIP != "" {
		block += fmt.Sprintf("  management_ip = %q\n", managementIP)
	}
	return block + "}\n"
}

func checkInventorySwitchDataSourceContains(dataSourceName, serialNumber, managementIP string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		dataSource, ok := state.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("data source %q not found in state", dataSourceName)
		}
		if dataSource.Primary == nil {
			return fmt.Errorf("data source %q has no primary instance", dataSourceName)
		}

		attributes := dataSource.Primary.Attributes
		count, err := strconv.Atoi(attributes["switches.#"])
		if err != nil {
			return fmt.Errorf("could not read switches count from data source %q: %w", dataSourceName, err)
		}
		for index := 0; index < count; index++ {
			prefix := fmt.Sprintf("switches.%d.", index)
			if attributes[prefix+"serial_number"] == serialNumber && attributes[prefix+"ip_address"] == managementIP {
				return nil
			}
		}
		return fmt.Errorf("data source %q did not return switch with serial %q and management IP %q", dataSourceName, serialNumber, managementIP)
	}
}
