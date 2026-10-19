// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
)

// policyAllowedAttrs maps each built-in policyType to the set of
// policy-scoped tf_name attributes that are valid for it. Generated from
// the attribute summary in spec/3-features/001-interface-common/Attribute-summary.MD.
//
// This static map is the plan-time safety net — it runs during
// ValidateConfig (terraform plan) where no API client is available.
// At apply time, validateAttrsAgainstTemplate in resource_custom_codec.go
// queries the live ND TemplateDB (internal/common/templates) for the
// authoritative parameter list, catching any drift from this static map.
//
// Identity/structural fields (mode, network_os_type, policy_type, switch_id,
// template_name) are always allowed and are NOT in these sets.
//
// The special key "userDefined" is absent because userDefined accepts any
// attributes (the user-provided template defines the shape).
var policyAllowedAttrs = map[string]map[string]bool{
	"iosXeAccess": toSet(
		"access_vlan", "admin_state", "bpdu_guard", "description", "extra_config", "mtu", "speed",
	),
	"accessHost": toSet(
		"access_vlan", "admin_state", "bandwidth", "bpdu_filter", "bpdu_guard", "cdp",
		"debounce_linkup_timer", "debounce_timer", "description", "duplex_mode",
		"error_detection_acl", "extra_config", "fec", "inherit_bandwidth", "link_type",
		"monitor", "mtu", "negotiate_auto", "netflow", "netflow_monitor", "netflow_sampler",
		"orphan_port", "pfc", "port_type_edge_trunk", "qos", "qos_policy", "queuing_policy",
		"speed", "storm_control", "storm_control_action",
		"storm_control_broadcast_level", "storm_control_broadcast_level_pps",
		"storm_control_multicast_level", "storm_control_multicast_level_pps",
		"storm_control_unicast_level", "storm_control_unicast_level_pps",
	),
	"ipfmAccessHost": toSet(
		"access_vlan", "admin_state", "bpdu_guard", "description", "extra_config", "mtu",
		"port_type_edge_trunk", "ptp", "ptp_profile", "ptp_role", "ptp_vlan", "speed",
	),
	"dot1qTunnelHost": toSet(
		"access_vlan", "admin_state", "bandwidth", "bpdu_filter", "bpdu_guard", "cdp",
		"debounce_linkup_timer", "debounce_timer", "description", "duplex_mode",
		"error_detection_acl", "extra_config", "fec", "inherit_bandwidth", "link_type",
		"mtu", "negotiate_auto", "netflow", "netflow_monitor", "netflow_sampler",
		"orphan_port", "pfc", "port_type_edge_trunk", "qos", "qos_policy", "queuing_policy",
		"speed", "storm_control", "storm_control_action",
		"storm_control_broadcast_level", "storm_control_broadcast_level_pps",
		"storm_control_multicast_level", "storm_control_multicast_level_pps",
		"storm_control_unicast_level", "storm_control_unicast_level_pps",
	),
	"iosXeStackwiseDualActive": toSet(
		"extra_config",
	),
	"iosXeStackwiseLinkMember": toSet(
		"extra_config",
	),
	"pvlanHost": toSet(
		"admin_state", "association_list", "bpdu_guard", "cdp", "description", "duplex_mode",
		"extra_config", "mapping_list", "mtu", "orphan_port", "pfc", "port_type_edge_trunk",
		"pvlan_allowed_vlans", "pvlan_mode", "pvlan_native_vlan",
		"qos", "qos_policy", "queuing_policy", "speed",
	),
	"csrMultisiteIfcMember": toSet(
		"cdp", "description", "extra_config", "if_name", "ip", "mask",
	),
	"iosXeInternalL3PoMember": toSet(
		"admin_state", "description", "extra_config", "port_channel_id", "port_channel_mode",
	),
	"iosXeNumbered": toSet(
		"admin_state", "description", "extra_config", "ip", "mtu", "prefix", "speed", "v6prefix",
	),
	"iosXeRoutedHost": toSet(
		"admin_state", "description", "extra_config", "ip", "mtu", "prefix", "speed", "vrf_interface",
	),
	"iosXrRoutedHost": toSet(
		"admin_state", "description", "extra_config", "ip", "mtu", "prefix", "vrf_interface",
	),
	"dataBrokerL3Host": toSet(
		"admin_state", "description", "ip", "prefix",
	),
	"endPointLocator": toSet(
		"admin_state", "description", "extra_config", "ip", "ipv6", "ipv6_prefix",
		"mtu", "prefix", "routing_tag", "speed",
	),
	"ipfmL3Port": toSet(
		"admin_state", "bandwidth_capacity_percent", "border_router", "description",
		"extra_config", "igmp_group_ranges", "ip", "mtu", "prefix", "ptp", "ptp_profile",
		"ptp_role", "routing_tag", "security_group_proxy", "speed",
		"unicast_bandwidth_percent", "vrf_interface", "wan_link",
	),
	"routedHost": toSet(
		"admin_state", "description", "extra_config", "fec", "ip", "ip_redirects",
		"mtu", "netflow", "netflow_monitor", "netflow_sampler", "pfc",
		"pim_dr_priority", "pim_sparse", "prefix", "qos", "qos_policy",
		"queuing_policy", "routing_tag", "speed", "vrf_interface",
	),
	"iosXeTrunkHost": toSet(
		"admin_state", "allowed_vlans", "bpdu_guard", "description", "extra_config",
		"mtu", "speed",
	),
	"classicHost": toSet(
		"admin_state", "allowed_vlans", "bpdu_guard", "cdp", "description", "duplex_mode",
		"extra_config", "guard_mode", "mst_path_cost_list", "mst_port_prio_list",
		"mtu", "netflow", "netflow_monitor", "netflow_sampler", "orphan_port",
		"path_cost", "port_type_edge_trunk", "speed",
	),
	"dataBrokerHost": toSet(
		"admin_state", "description",
	),
	"dataBrokerPoMember": toSet(
		"admin_state", "description", "port_channel_id", "port_channel_mode",
	),
	"ipfmTrunkHost": toSet(
		"admin_state", "allowed_vlans", "bpdu_guard", "description", "extra_config",
		"mtu", "port_type_edge_trunk", "ptp", "ptp_profile", "ptp_role", "ptp_vlan", "speed",
	),
	"trunkHost": toSet(
		"admin_state", "allowed_vlans", "bandwidth", "bpdu_filter", "bpdu_guard", "cdp",
		"debounce_linkup_timer", "debounce_timer", "description", "duplex_mode",
		"error_detection_acl", "extra_config", "fec", "inherit_bandwidth", "link_type",
		"monitor", "mtu", "native_vlan", "negotiate_auto", "netflow", "netflow_monitor",
		"netflow_sampler", "orphan_port", "pfc", "port_type_edge_trunk",
		"qos", "qos_policy", "queuing_policy", "speed",
		"storm_control", "storm_control_action",
		"storm_control_broadcast_level", "storm_control_broadcast_level_pps",
		"storm_control_multicast_level", "storm_control_multicast_level_pps",
		"storm_control_unicast_level", "storm_control_unicast_level_pps",
		"vlan_mapping", "vlan_mapping_entries",
	),
	"iosXeMonitor": toSet(),
	"monitor":      toSet(),
}

func toSet(names ...string) map[string]bool {
	s := make(map[string]bool, len(names))
	for _, n := range names {
		s[n] = true
	}
	return s
}

// configuredPolicyAttrs returns the set of policy-scoped tf_name attributes
// that the user explicitly configured (non-null and non-unknown) on an
// interface entry. Identity/structural attributes are excluded.
func configuredPolicyAttrs(e InterfacesValue) map[string]bool {
	set := make(map[string]bool)
	type namedAttr struct {
		name string
		val  attr.Value
	}
	// List every policy-scoped attribute. Identity fields (mode,
	// network_os_type, policy_type, switch_id, template_name) are excluded.
	attrs := []namedAttr{
		{"access_vlan", e.AccessVlan},
		{"admin_state", e.AdminState},
		{"allowed_vlans", e.AllowedVlans},
		{"association_list", e.AssociationList},
		{"bandwidth", e.Bandwidth},
		{"bandwidth_capacity_percent", e.BandwidthCapacityPercent},
		{"border_router", e.BorderRouter},
		{"bpdu_filter", e.BpduFilter},
		{"bpdu_guard", e.BpduGuard},
		{"cdp", e.Cdp},
		{"debounce_linkup_timer", e.DebounceLinkupTimer},
		{"debounce_timer", e.DebounceTimer},
		{"description", e.Description},
		{"duplex_mode", e.DuplexMode},
		{"error_detection_acl", e.ErrorDetectionAcl},
		{"extra_config", e.ExtraConfig},
		{"fec", e.Fec},
		{"guard_mode", e.GuardMode},
		{"if_name", e.IfName},
		{"igmp_group_ranges", e.IgmpGroupRanges},
		{"inherit_bandwidth", e.InheritBandwidth},
		{"ip", e.Ip},
		{"ip_redirects", e.IpRedirects},
		{"ipv6", e.Ipv6},
		{"ipv6_prefix", e.Ipv6Prefix},
		{"link_type", e.LinkType},
		{"mapping_list", e.MappingList},
		{"mask", e.Mask},
		{"monitor", e.Monitor},
		{"mst_path_cost_list", e.MstPathCostList},
		{"mst_port_prio_list", e.MstPortPrioList},
		{"mtu", e.Mtu},
		{"native_vlan", e.NativeVlan},
		{"negotiate_auto", e.NegotiateAuto},
		{"netflow", e.Netflow},
		{"netflow_monitor", e.NetflowMonitor},
		{"netflow_sampler", e.NetflowSampler},
		{"orphan_port", e.OrphanPort},
		{"path_cost", e.PathCost},
		{"pfc", e.Pfc},
		{"pim_dr_priority", e.PimDrPriority},
		{"pim_sparse", e.PimSparse},
		{"port_channel_id", e.PortChannelId},
		{"port_channel_mode", e.PortChannelMode},
		{"port_type_edge_trunk", e.PortTypeEdgeTrunk},
		{"prefix", e.Prefix},
		{"ptp", e.Ptp},
		{"ptp_profile", e.PtpProfile},
		{"ptp_role", e.PtpRole},
		{"ptp_vlan", e.PtpVlan},
		{"pvlan_allowed_vlans", e.PvlanAllowedVlans},
		{"pvlan_mode", e.PvlanMode},
		{"pvlan_native_vlan", e.PvlanNativeVlan},
		{"qos", e.Qos},
		{"qos_policy", e.QosPolicy},
		{"queuing_policy", e.QueuingPolicy},
		{"routing_tag", e.RoutingTag},
		{"security_group_proxy", e.SecurityGroupProxy},
		{"speed", e.Speed},
		{"storm_control", e.StormControl},
		{"storm_control_action", e.StormControlAction},
		{"storm_control_broadcast_level", e.StormControlBroadcastLevel},
		{"storm_control_broadcast_level_pps", e.StormControlBroadcastLevelPps},
		{"storm_control_multicast_level", e.StormControlMulticastLevel},
		{"storm_control_multicast_level_pps", e.StormControlMulticastLevelPps},
		{"storm_control_unicast_level", e.StormControlUnicastLevel},
		{"storm_control_unicast_level_pps", e.StormControlUnicastLevelPps},
		{"unicast_bandwidth_percent", e.UnicastBandwidthPercent},
		{"v6prefix", e.V6prefix},
		{"vlan_mapping", e.VlanMapping},
		{"vlan_mapping_entries", e.VlanMappingEntries},
		{"vrf_interface", e.VrfInterface},
		{"wan_link", e.WanLink},
	}
	for _, a := range attrs {
		if !a.val.IsNull() && !a.val.IsUnknown() {
			set[a.name] = true
		}
	}
	return set
}
