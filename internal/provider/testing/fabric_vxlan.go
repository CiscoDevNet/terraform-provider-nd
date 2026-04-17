// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_fabric_vxlan"
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
func GenerateFabricVxlanObject(obj **resource_fabric_vxlan.NDFCFabricVxlanModel,
	fabricName string, bgpAsn string, fabricType string,
	overrides map[string]interface{}) {

	fabric := new(resource_fabric_vxlan.NDFCFabricVxlanModel)

	fabric.FabricName = fabricName
	fabric.Management.FabricType = fabricType
	fabric.Management.BgpAsn = bgpAsn

	// Merge defaults with caller overrides (overrides win)
	merged := defaultFabricValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyFabricValues(fabric, merged)

	*obj = fabric
}

// ModifyFabricVxlanObject modifies fields on an existing fabric model.
// Uses the same key set as GenerateFabricVxlanObject overrides.
func ModifyFabricVxlanObject(
	obj **resource_fabric_vxlan.NDFCFabricVxlanModel,
	values map[string]interface{},
) {
	fabric := *obj
	if fabric == nil {
		return
	}

	applyFabricValues(fabric, values)

	*obj = fabric
}

// applyFabricValues is the shared engine that sets fields on a fabric model
// from a key-value map. Used by both Generate and Modify.
func applyFabricValues(fabric *resource_fabric_vxlan.NDFCFabricVxlanModel, values map[string]interface{}) {
	for key, val := range values {
		switch key {
		// Top-level fields
		case "license_tier":
			fabric.LicenseTier = val.(string)
		case "category":
			fabric.Category = val.(string)
		case "security_domain":
			fabric.SecurityDomain = val.(string)

		// Management string fields
		case "anycast_gateway_mac":
			fabric.Management.AnycastGatewayMac = val.(string)
		case "replication_mode":
			fabric.Management.ReplicationMode = val.(string)
		case "multicast_group_subnet":
			fabric.Management.MulticastGroupSubnet = val.(string)
		case "bgp_loopback_ip_range":
			fabric.Management.BgpLoopbackIpRange = val.(string)
		case "nve_loopback_ip_range":
			fabric.Management.NveLoopbackIpRange = val.(string)
		case "intra_fabric_subnet_range":
			fabric.Management.IntraFabricSubnetRange = val.(string)
		case "overlay_mode":
			fabric.Management.OverlayMode = val.(string)
		case "site_id":
			fabric.Management.SiteId = val.(string)
		case "vpc_peer_link_vlan":
			fabric.Management.VpcPeerLinkVlan = val.(string)
		case "vpc_peer_keep_alive_option":
			fabric.Management.VpcPeerKeepAliveOption = val.(string)
		case "vpc_peer_link_port_channel_id":
			fabric.Management.VpcPeerLinkPortChannelId = val.(string)
		case "vpc_domain_id_range":
			fabric.Management.VpcDomainIdRange = val.(string)
		case "vrf_template":
			fabric.Management.VrfTemplate = val.(string)
		case "network_template":
			fabric.Management.NetworkTemplate = val.(string)
		case "vrf_extension_template":
			fabric.Management.VrfExtensionTemplate = val.(string)
		case "network_extension_template":
			fabric.Management.NetworkExtensionTemplate = val.(string)
		case "l2_vni_range":
			fabric.Management.L2VniRange = val.(string)
		case "l3_vni_range":
			fabric.Management.L3VniRange = val.(string)
		case "network_vlan_range":
			fabric.Management.NetworkVlanRange = val.(string)
		case "vrf_vlan_range":
			fabric.Management.VrfVlanRange = val.(string)
		case "bgp_as_mode":
			fabric.Management.BgpAsMode = val.(string)
		case "super_spine_bgp_as":
			fabric.Management.SuperSpineBgpAs = val.(string)
		case "leaf_bgp_as":
			fabric.Management.LeafBgpAs = val.(string)
		case "border_bgp_as":
			fabric.Management.BorderBgpAs = val.(string)
		case "anycast_rendezvous_point_ip_range":
			fabric.Management.AnycastRendezvousPointIpRange = val.(string)
		case "sub_interface_dot1q_range":
			fabric.Management.SubInterfaceDot1qRange = val.(string)

		// Management *int64 fields
		case "target_subnet_mask":
			v := int64(val.(int))
			fabric.Management.TargetSubnetMask = &v
		case "rendezvous_point_count":
			v := int64(val.(int))
			fabric.Management.RendezvousPointCount = &v
		case "bgp_loopback_id":
			v := int64(val.(int))
			fabric.Management.BgpLoopbackId = &v
		case "nve_loopback_id":
			v := int64(val.(int))
			fabric.Management.NveLoopbackId = &v
		case "vpc_auto_recovery_timer":
			v := int64(val.(int))
			fabric.Management.VpcAutoRecoveryTimer = &v
		case "vpc_delay_restore_timer":
			v := int64(val.(int))
			fabric.Management.VpcDelayRestoreTimer = &v
		case "fabric_mtu":
			v := int64(val.(int))
			fabric.Management.FabricMtu = &v
		case "l2_host_interface_mtu":
			v := int64(val.(int))
			fabric.Management.L2HostInterfaceMtu = &v
		}
	}
}
