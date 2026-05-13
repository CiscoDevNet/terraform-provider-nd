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

// applyFabricCommonValues sets fields on an NDFCFabricCommonModel from a key-value map.
// Shared by all fabric type test helpers (ibgp, ebgp, etc.).
func applyFabricCommonValues(fabric *resource_fabric_common.NDFCFabricCommonModel, values map[string]interface{}) {
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
		case "bgp_asn_range":
			fabric.Management.BgpAsnRange = val.(string)

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
		case "bgp_allow_as_in_num":
			v := int64(val.(int))
			fabric.Management.BgpAllowAsInNum = &v
		case "bgp_max_path":
			v := int64(val.(int))
			fabric.Management.BgpMaxPath = &v

		// Management *bool fields
		case "bgp_asn_auto_allocation":
			v := val.(bool)
			fabric.Management.BgpAsnAutoAllocation = &v
		case "bgp_underlay_failure_protect":
			v := val.(bool)
			fabric.Management.BgpUnderlayFailureProtect = &v
		case "auto_configure_ebgp_evpn_peering":
			v := val.(bool)
			fabric.Management.AutoConfigureEbgpEvpnPeering = &v
		case "allow_leaf_same_as":
			v := val.(bool)
			fabric.Management.AllowLeafSameAs = &v
		case "assign_ipv4_to_loopback0":
			v := val.(bool)
			fabric.Management.AssignIpv4ToLoopback0 = &v
		}
	}
}
