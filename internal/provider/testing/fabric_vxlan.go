// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

// defaultFabricValues returns sensible defaults for a VXLAN fabric.
// Tests can override any of these via the overrides map passed to GenerateFabricVxlanObject.
func defaultFabricValues() map[string]interface{} {
	return map[string]interface{}{
		"license_tier":                  "premier",
		"category":                      "fabric",
		"security_domain":               "all",
		"anycast_gateway_mac":           "2020.0000.00aa",
		"replication_mode":              "multicast",
		"multicast_group_subnet":        "239.1.1.0/25",
		"target_subnet_mask":            30,
		"rendezvous_point_count":        2,
		"bgp_loopback_ip_range":         "10.2.0.0/22",
		"nve_loopback_ip_range":         "10.3.0.0/22",
		"intra_fabric_subnet_range":     "10.4.0.0/16",
		"bgp_loopback_id":               0,
		"nve_loopback_id":               1,
		"vrf_template":                  "Default_VRF_Universal",
		"network_template":              "Default_Network_Universal",
		"vrf_extension_template":        "Default_VRF_Extension_Universal",
		"network_extension_template":    "Default_Network_Extension_Universal",
		"vpc_peer_link_vlan":            "3600",
		"vpc_peer_keep_alive_option":    "management",
		"vpc_auto_recovery_timer":       360,
		"vpc_delay_restore_timer":       150,
		"vpc_peer_link_port_channel_id": "500",
		"vpc_domain_id_range":           "1-1000",
		"l2_vni_range":                  "30000-49000",
		"l3_vni_range":                  "50000-59000",
		"network_vlan_range":            "2300-2999",
		"vrf_vlan_range":                "2000-2299",
		"greenfield_debug_flag":         "enable",
	}
}

// GenerateFabricVxlanObject creates a VXLAN fabric model object for testing.
// fabricName, bgpAsn, and fabricType are mandatory identifiers.
// overrides lets each test supply unique values for any field so that
// multiple fabrics can coexist without conflicting on IP ranges, VNI ranges, etc.
// Any key not present in overrides gets the value from defaultFabricValues().
func GenerateFabricVxlanObject(obj **resource_fabric_common.NDFCFabricCommonModel,
	fabricName string, bgpAsn string, fabricType string,
	overrides map[string]interface{}) {

	fabric := new(resource_fabric_common.NDFCFabricCommonModel)

	fabric.FabricName = fabricName
	fabric.Management.FabricType = fabricType
	fabric.Management.BgpAsn = bgpAsn

	// Merge defaults with caller overrides (overrides win)
	merged := defaultFabricValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyFabricCommonValues(fabric, merged)

	*obj = fabric
}

// ModifyFabricVxlanObject modifies fields on an existing fabric model.
// Uses the same key set as GenerateFabricVxlanObject overrides.
func ModifyFabricVxlanObject(
	obj **resource_fabric_common.NDFCFabricCommonModel,
	values map[string]interface{},
) {
	fabric := *obj
	if fabric == nil {
		return
	}

	applyFabricCommonValues(fabric, values)

	*obj = fabric
}
