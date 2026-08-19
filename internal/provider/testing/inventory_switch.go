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
func GenerateInventorySwitchObject(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	fabricName string,
	username string,
	password string,
	serialNumber string,
	ipAddress string,
	role string,
) {
	inv := new(resource_inventory_switch.NDFCInventorySwitchModel)

	inv.FabricName = fabricName
	inv.Mode = "discovery"
	inv.DiscoveryUsername = username
	inv.DiscoveryPassword = password
	inv.SnmpV3AuthProtocol = "md5"

	preserveConfig := false
	inv.PreserveConfig = &preserveConfig

	maxHop := int64(0)
	inv.MaxHop = &maxHop

	if role == "" {
		role = "leaf"
	}
	inv.SwitchDetail = resource_inventory_switch.NDFCSwitchDetailValue{
		SerialNumber: serialNumber,
		IpAddress:    ipAddress,
		SwitchRole:   role,
	}

	*obj = inv
}

// ModifyInventorySwitchObject modifies fields on an existing inventory switch model.
// Supported keys: "preserve_config", "mode", "max_hop", "switch_role"
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
		case "preserve_config":
			v := val.(bool)
			inv.PreserveConfig = &v
		case "mode":
			inv.Mode = val.(string)
		case "max_hop":
			v := int64(val.(int))
			inv.MaxHop = &v
		case "switch_role":
			inv.SwitchDetail.SwitchRole = val.(string)
		}
	}

	*obj = inv
}

// GenerateBootstrapSwitchObject creates a bootstrap-mode inventory switch model for testing.
func GenerateBootstrapSwitchObject(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	fabricName string,
	serialNumber string,
	ipAddress string,
	hostname string,
	role string,
	gatewayIpMask string,
	password string,
) {
	inv := new(resource_inventory_switch.NDFCInventorySwitchModel)

	inv.FabricName = fabricName
	inv.Mode = "bootstrap"
	inv.BootstrapPassword = password
	inv.SnmpV3AuthProtocol = "md5"

	preserveConfig := false
	inv.PreserveConfig = &preserveConfig

	maxHop := int64(0)
	inv.MaxHop = &maxHop

	if role == "" {
		role = "leaf"
	}
	inv.SwitchDetail = resource_inventory_switch.NDFCSwitchDetailValue{
		SerialNumber:  serialNumber,
		IpAddress:     ipAddress,
		Hostname:      hostname,
		SwitchRole:    role,
		GatewayIpMask: gatewayIpMask,
	}

	*obj = inv
}

// GenerateInventorySwitchFromConfig creates an inventory switch model from testbed config.
func GenerateInventorySwitchFromConfig(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	fabricName string,
	user string,
	password string,
	sw InventorySwitch,
) {
	mode := sw.Mode
	if mode == "" {
		mode = "discovery"
	}

	if mode == "bootstrap" {
		GenerateBootstrapSwitchObject(obj,
			fabricName,
			sw.Serial, sw.IP, sw.Hostname,
			sw.Role, sw.GatewayIpMask, sw.PoapPassword,
		)
	} else {
		GenerateInventorySwitchObject(obj,
			fabricName, user, password,
			sw.Serial, sw.IP, sw.Role,
		)
	}
}

// ModifyDiscoveryCredentials changes the discovery credentials on an existing switch model.
func ModifyDiscoveryCredentials(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	username string,
	password string,
	snmpV3AuthProtocol string,
) {
	inv := *obj
	if inv == nil {
		return
	}

	inv.DiscoveryUsername = username
	inv.DiscoveryPassword = password
	inv.SnmpV3AuthProtocol = snmpV3AuthProtocol
	*obj = inv
}

// ModifySwitchRole changes the role of the switch
func ModifySwitchRole(
	obj **resource_inventory_switch.NDFCInventorySwitchModel,
	newRole string,
) {
	inv := *obj
	if inv == nil {
		return
	}

	inv.SwitchDetail.SwitchRole = newRole
	*obj = inv
}
