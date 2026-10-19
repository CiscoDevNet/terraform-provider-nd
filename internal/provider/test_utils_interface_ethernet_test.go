// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"terraform-provider-nd/internal/manage/resource_interface_common"
)

func InterfaceEthernetModelHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfaceCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.SwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_id").String(), c.SwitchId))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	for key, value := range c.Interfaces {
		attrNewPath := attrPath.AtName("interfaces").AtName(key)
		ret = append(ret, InterfaceEthernetInterfacesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func InterfaceEthernetAssociationListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCAssociationListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.PVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("p_vlan").String(), strconv.Itoa(int(*c.PVlan))))
	}
	if c.SVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("s_vlan").String(), strconv.Itoa(int(*c.SVlan))))
	}
	return ret
}

func InterfaceEthernetMappingListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCMappingListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.PVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("p_vlan").String(), strconv.Itoa(int(*c.PVlan))))
	}
	if c.SVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("s_vlan").String(), strconv.Itoa(int(*c.SVlan))))
	}
	return ret
}

func InterfaceEthernetIgmpGroupRangesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCIgmpGroupRangesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.GroupAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("group_address").String(), c.GroupAddress))
	}
	if c.Prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("prefix").String(), strconv.Itoa(int(*c.Prefix))))
	}
	return ret
}

func InterfaceEthernetMstPathCostListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCMstPathCostListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.PathCost != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("path_cost").String(), strconv.Itoa(int(*c.PathCost))))
	}
	return ret
}

func InterfaceEthernetMstPortPrioListValueHelperStateCheck(RscName string, c resource_interface_common.NDFCMstPortPrioListValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.PortPriority != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_priority").String(), strconv.Itoa(int(*c.PortPriority))))
	}
	return ret
}

func InterfaceEthernetVlanMappingEntriesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCVlanMappingEntriesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Dot1qTunnel != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dot1q_tunnel").String(), strconv.FormatBool(*c.Dot1qTunnel)))
	}
	if c.CustomerInnerVlanId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("customer_inner_vlan_id").String(), strconv.Itoa(int(*c.CustomerInnerVlanId))))
	}
	if c.ProviderVlanId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("provider_vlan_id").String(), strconv.Itoa(int(*c.ProviderVlanId))))
	}
	return ret
}

func InterfaceEthernetInterfacesValueHelperStateCheck(RscName string, c resource_interface_common.NDFCInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_id").String(), c.SwitchId))
	}
	if c.ConfigData.Mode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), c.ConfigData.Mode))
	}
	if c.ConfigData.NetworkOS.NetworkOsType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_os_type").String(), c.ConfigData.NetworkOS.NetworkOsType))
	}
	if c.ConfigData.NetworkOS.Policy.PolicyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy_type").String(), c.ConfigData.NetworkOS.Policy.PolicyType))
	}
	if c.ConfigData.NetworkOS.Policy.TemplateName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_name").String(), c.ConfigData.NetworkOS.Policy.TemplateName))
	}
	if c.ConfigData.NetworkOS.Policy.AdminState != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("admin_state").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.AdminState)))
	}
	if c.ConfigData.NetworkOS.Policy.Description != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("description").String(), c.ConfigData.NetworkOS.Policy.Description))
	}
	if c.ConfigData.NetworkOS.Policy.Mtu != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), c.ConfigData.NetworkOS.Policy.Mtu))
	}
	if c.ConfigData.NetworkOS.Policy.ExtraConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config").String(), c.ConfigData.NetworkOS.Policy.ExtraConfig))
	}
	if c.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("access_vlan").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.AccessVlan))))
	}
	if c.ConfigData.NetworkOS.Policy.AllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allowed_vlans").String(), c.ConfigData.NetworkOS.Policy.AllowedVlans))
	}
	if c.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bandwidth").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.Bandwidth))))
	}
	if c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bandwidth_capacity_percent").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent))))
	}
	if c.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_router").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.BorderRouter)))
	}
	if c.ConfigData.NetworkOS.Policy.BpduFilter != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bpdu_filter").String(), c.ConfigData.NetworkOS.Policy.BpduFilter))
	}
	if c.ConfigData.NetworkOS.Policy.BpduGuard != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bpdu_guard").String(), c.ConfigData.NetworkOS.Policy.BpduGuard))
	}
	if c.ConfigData.NetworkOS.Policy.Cdp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Cdp)))
	}
	if c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("debounce_linkup_timer").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer))))
	}
	if c.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("debounce_timer").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.DebounceTimer))))
	}
	if c.ConfigData.NetworkOS.Policy.DuplexMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("duplex_mode").String(), c.ConfigData.NetworkOS.Policy.DuplexMode))
	}
	if c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("error_detection_acl").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)))
	}
	if c.ConfigData.NetworkOS.Policy.Fec != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fec").String(), c.ConfigData.NetworkOS.Policy.Fec))
	}
	if c.ConfigData.NetworkOS.Policy.GuardMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("guard_mode").String(), c.ConfigData.NetworkOS.Policy.GuardMode))
	}
	if c.ConfigData.NetworkOS.Policy.IfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("if_name").String(), c.ConfigData.NetworkOS.Policy.IfName))
	}
	if c.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inherit_bandwidth").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.InheritBandwidth))))
	}
	if c.ConfigData.NetworkOS.Policy.Ip != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip").String(), c.ConfigData.NetworkOS.Policy.Ip))
	}
	if c.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip_redirects").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.IpRedirects)))
	}
	if c.ConfigData.NetworkOS.Policy.Ipv6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6").String(), c.ConfigData.NetworkOS.Policy.Ipv6))
	}
	if c.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_prefix").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.Ipv6Prefix))))
	}
	if c.ConfigData.NetworkOS.Policy.LinkType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_type").String(), c.ConfigData.NetworkOS.Policy.LinkType))
	}
	if c.ConfigData.NetworkOS.Policy.Mask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mask").String(), c.ConfigData.NetworkOS.Policy.Mask))
	}
	if c.ConfigData.NetworkOS.Policy.Monitor != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("monitor").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Monitor)))
	}
	if c.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("native_vlan").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.NativeVlan))))
	}
	if c.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("negotiate_auto").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.NegotiateAuto)))
	}
	if c.ConfigData.NetworkOS.Policy.Netflow != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Netflow)))
	}
	if c.ConfigData.NetworkOS.Policy.NetflowMonitor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_monitor").String(), c.ConfigData.NetworkOS.Policy.NetflowMonitor))
	}
	if c.ConfigData.NetworkOS.Policy.NetflowSampler != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_sampler").String(), c.ConfigData.NetworkOS.Policy.NetflowSampler))
	}
	if c.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orphan_port").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.OrphanPort)))
	}
	if c.ConfigData.NetworkOS.Policy.PathCost != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("path_cost").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.PathCost))))
	}
	if c.ConfigData.NetworkOS.Policy.Pfc != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pfc").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Pfc)))
	}
	if c.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_dr_priority").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.PimDrPriority))))
	}
	if c.ConfigData.NetworkOS.Policy.PimSparse != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_sparse").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.PimSparse)))
	}
	if c.ConfigData.NetworkOS.Policy.PortChannelId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_channel_id").String(), c.ConfigData.NetworkOS.Policy.PortChannelId))
	}
	if c.ConfigData.NetworkOS.Policy.PortChannelMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_channel_mode").String(), c.ConfigData.NetworkOS.Policy.PortChannelMode))
	}
	if c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_type_edge_trunk").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)))
	}
	if c.ConfigData.NetworkOS.Policy.Prefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("prefix").String(), c.ConfigData.NetworkOS.Policy.Prefix))
	}
	if c.ConfigData.NetworkOS.Policy.Ptp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Ptp)))
	}
	if c.ConfigData.NetworkOS.Policy.PtpProfile != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_profile").String(), c.ConfigData.NetworkOS.Policy.PtpProfile))
	}
	if c.ConfigData.NetworkOS.Policy.PtpRole != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_role").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.PtpRole)))
	}
	if c.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_vlan").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.PtpVlan))))
	}
	if c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_allowed_vlans").String(), c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans))
	}
	if c.ConfigData.NetworkOS.Policy.PvlanMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_mode").String(), c.ConfigData.NetworkOS.Policy.PvlanMode))
	}
	if c.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pvlan_native_vlan").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.PvlanNativeVlan))))
	}
	if c.ConfigData.NetworkOS.Policy.Qos != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("qos").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.Qos)))
	}
	if c.ConfigData.NetworkOS.Policy.QosPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("qos_policy").String(), c.ConfigData.NetworkOS.Policy.QosPolicy))
	}
	if c.ConfigData.NetworkOS.Policy.QueuingPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("queuing_policy").String(), c.ConfigData.NetworkOS.Policy.QueuingPolicy))
	}
	if c.ConfigData.NetworkOS.Policy.RoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_tag").String(), c.ConfigData.NetworkOS.Policy.RoutingTag))
	}
	if c.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_proxy").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.SecurityGroupProxy)))
	}
	if c.ConfigData.NetworkOS.Policy.Speed != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("speed").String(), c.ConfigData.NetworkOS.Policy.Speed))
	}
	if c.ConfigData.NetworkOS.Policy.StormControl != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storm_control").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.StormControl)))
	}
	if c.ConfigData.NetworkOS.Policy.StormControlAction != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storm_control_action").String(), c.ConfigData.NetworkOS.Policy.StormControlAction))
	}
	if c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storm_control_broadcast_level_pps").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps))))
	}
	if c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storm_control_multicast_level_pps").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps))))
	}
	if c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("storm_control_unicast_level_pps").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps))))
	}
	if c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("unicast_bandwidth_percent").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent))))
	}
	if c.ConfigData.NetworkOS.Policy.V6prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("v6prefix").String(), strconv.Itoa(int(*c.ConfigData.NetworkOS.Policy.V6prefix))))
	}
	if c.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vlan_mapping").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.VlanMapping)))
	}
	if c.ConfigData.NetworkOS.Policy.VrfInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_interface").String(), c.ConfigData.NetworkOS.Policy.VrfInterface))
	}
	if c.ConfigData.NetworkOS.Policy.WanLink != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wan_link").String(), strconv.FormatBool(*c.ConfigData.NetworkOS.Policy.WanLink)))
	}
	return ret
}
