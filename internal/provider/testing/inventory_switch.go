// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
)

// GenerateInventorySwitchObject creates an inventory switch model object for testing.
// switches is a map of serial_number -> ip_address
func GenerateInventorySwitchObject(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	fabricName string,
	username string,
	password string,
	switches map[string]string,
	roles map[string]string,
	deploy bool,
	recalculate bool,
) {
	inv := new(resource_inventory_switch.NDFCInventorySwitchModel)

	inv.FabricName = fabricName
	inv.Mode = "discovery"
	inv.Username = username
	inv.Password = password
	inv.SnmpV3AuthProtocol = "MD5"

	preserveConfig := false
	inv.PreserveConfig = &preserveConfig

	maxHop := int64(0)
	inv.MaxHop = &maxHop

	inv.Deploy = deploy
	inv.Recalculate = recalculate

	inv.Switches = make(map[string]resource_inventory_switch.NDFCSwitchesValue)
	for serial, ip := range switches {
		role := "leaf"
		if r, ok := roles[serial]; ok {
			role = r
		}
		inv.Switches[serial] = resource_inventory_switch.NDFCSwitchesValue{
			IpAddress:  ip,
			SwitchRole: role,
		}
	}

	*obj = inv
}

// ModifyInventorySwitchObject modifies fields on an existing inventory switch model.
// Supported keys: "deploy", "recalculate", "preserve_config", "mode", "max_hop"
func ModifyInventorySwitchObject(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	values map[string]interface{},
) {
	inv := *obj
	if inv == nil {
		return
	}

	for key, val := range values {
		switch key {
		case "deploy":
			inv.Deploy = val.(bool)
		case "recalculate":
			inv.Recalculate = val.(bool)
		case "preserve_config":
			v := val.(bool)
			inv.PreserveConfig = &v
		case "mode":
			inv.Mode = val.(string)
		case "max_hop":
			v := int64(val.(int))
			inv.MaxHop = &v
		}
	}

	*obj = inv
}

// AddSwitches adds switches to an existing inventory switch model
func AddSwitches(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	switches map[string]string,
	roles map[string]string,
) {
	inv := *obj
	if inv == nil {
		return
	}

	if inv.Switches == nil {
		inv.Switches = make(map[string]resource_inventory_switch.NDFCSwitchesValue)
	}

	for serial, ip := range switches {
		role := "leaf"
		if r, ok := roles[serial]; ok {
			role = r
		}
		inv.Switches[serial] = resource_inventory_switch.NDFCSwitchesValue{
			IpAddress:  ip,
			SwitchRole: role,
		}
	}

	*obj = inv
}

// DeleteSwitches removes switches by serial number from an existing inventory switch model
func DeleteSwitches(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	serials []string,
) {
	inv := *obj
	if inv == nil || inv.Switches == nil {
		return
	}

	for _, serial := range serials {
		delete(inv.Switches, serial)
	}

	*obj = inv
}

// ModifySwitchRole changes the role of a specific switch
func ModifySwitchRole(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	serial string,
	newRole string,
) {
	inv := *obj
	if inv == nil || inv.Switches == nil {
		return
	}

	if sw, ok := inv.Switches[serial]; ok {
		sw.SwitchRole = newRole
		inv.Switches[serial] = sw
	}

	*obj = inv
}
