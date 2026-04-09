// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"strconv"

	"terraform-provider-nd/internal/manage/resource_fabric_vxlan"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func FabricVxlanModelHelperStateCheck(RscName string, c resource_fabric_vxlan.NDFCFabricVxlanModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.LicenseTier != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("license_tier").String(), c.LicenseTier))
	}
	if c.SecurityDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), c.SecurityDomain))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), "all"))
	}
	if c.Category != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), c.Category))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), "fabric"))
	}
	if c.Management.FabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_type").String(), c.Management.FabricType))
	}
	if c.Management.BgpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_asn").String(), c.Management.BgpAsn))
	}
	if c.Management.SuperSpineBgpAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("super_spine_bgp_as").String(), c.Management.SuperSpineBgpAs))
	}
	if c.Management.LeafBgpAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("leaf_bgp_as").String(), c.Management.LeafBgpAs))
	}
	if c.Management.BorderBgpAs != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_bgp_as").String(), c.Management.BorderBgpAs))
	}
	if c.Management.BgpAsMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_as_mode").String(), c.Management.BgpAsMode))
	}
	if c.Management.TargetSubnetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("target_subnet_mask").String(), strconv.Itoa(int(*c.Management.TargetSubnetMask))))
	}
	if c.Management.AnycastGatewayMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gateway_mac").String(), c.Management.AnycastGatewayMac))
	}
	if c.Management.ReplicationMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("replication_mode").String(), c.Management.ReplicationMode))
	}
	if c.Management.MulticastGroupSubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multicast_group_subnet").String(), c.Management.MulticastGroupSubnet))
	}
	if c.Management.RendezvousPointCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_count").String(), strconv.Itoa(int(*c.Management.RendezvousPointCount))))
	}
	if c.Management.RendezvousPointLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_loopback_id").String(), strconv.Itoa(int(*c.Management.RendezvousPointLoopbackId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_loopback_id").String(), "254"))
	}
	if c.Management.VpcPeerLinkVlan != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_vlan").String(), c.Management.VpcPeerLinkVlan))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_vlan").String(), "3600"))
	}
	if c.Management.VpcPeerKeepAliveOption != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_keep_alive_option").String(), c.Management.VpcPeerKeepAliveOption))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_keep_alive_option").String(), "management"))
	}
	if c.Management.VpcAutoRecoveryTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_auto_recovery_timer").String(), strconv.Itoa(int(*c.Management.VpcAutoRecoveryTimer))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_auto_recovery_timer").String(), "360"))
	}
	if c.Management.VpcDelayRestoreTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_delay_restore_timer").String(), strconv.Itoa(int(*c.Management.VpcDelayRestoreTimer))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_delay_restore_timer").String(), "150"))
	}
	if c.Management.VpcPeerLinkPortChannelId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_port_channel_id").String(), c.Management.VpcPeerLinkPortChannelId))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_port_channel_id").String(), "500"))
	}
	if c.Management.VpcDomainIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_domain_id_range").String(), c.Management.VpcDomainIdRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_domain_id_range").String(), "1-1000"))
	}
	if c.Management.BgpLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_loopback_id").String(), strconv.Itoa(int(*c.Management.BgpLoopbackId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_loopback_id").String(), "0"))
	}
	if c.Management.NveLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_loopback_id").String(), strconv.Itoa(int(*c.Management.NveLoopbackId))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_loopback_id").String(), "1"))
	}
	if c.Management.VrfTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_template").String(), c.Management.VrfTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_template").String(), "Default_VRF_Universal"))
	}
	if c.Management.NetworkTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_template").String(), c.Management.NetworkTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_template").String(), "Default_Network_Universal"))
	}
	if c.Management.VrfExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), c.Management.VrfExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), "Default_VRF_Extension_Universal"))
	}
	if c.Management.NetworkExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), c.Management.NetworkExtensionTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), "Default_Network_Extension_Universal"))
	}
	if c.Management.SiteId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("site_id").String(), c.Management.SiteId))
	}
	if c.Management.FabricMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu").String(), strconv.Itoa(int(*c.Management.FabricMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_mtu").String(), "9216"))
	}
	if c.Management.L2HostInterfaceMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_interface_mtu").String(), strconv.Itoa(int(*c.Management.L2HostInterfaceMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_host_interface_mtu").String(), "9216"))
	}
	if c.Management.BgpLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_loopback_ip_range").String(), c.Management.BgpLoopbackIpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_loopback_ip_range").String(), "10.2.0.0/22"))
	}
	if c.Management.NveLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_loopback_ip_range").String(), c.Management.NveLoopbackIpRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_loopback_ip_range").String(), "10.3.0.0/22"))
	}
	if c.Management.IntraFabricSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("intra_fabric_subnet_range").String(), c.Management.IntraFabricSubnetRange))
	}
	if c.Management.L2VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_vni_range").String(), c.Management.L2VniRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_vni_range").String(), "30000-49000"))
	}
	if c.Management.L3VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_range").String(), c.Management.L3VniRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_range").String(), "50000-59000"))
	}
	if c.Management.NetworkVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_vlan_range").String(), c.Management.NetworkVlanRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_vlan_range").String(), "2300-2999"))
	}
	if c.Management.VrfVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_vlan_range").String(), c.Management.VrfVlanRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_vlan_range").String(), "2000-2299"))
	}
	if c.Management.OverlayMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), c.Management.OverlayMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), "cli"))
	}

	return ret
}
