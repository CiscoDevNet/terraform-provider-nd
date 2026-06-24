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

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

func FabricVxlanIbgpModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.LicenseTier != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("license_tier").String(), c.LicenseTier))
	}
	if c.FeatureStatus.ControllerStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("controller_status").String(), c.FeatureStatus.ControllerStatus))
	}
	if c.FeatureStatus.TelemetryStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_status").String(), c.FeatureStatus.TelemetryStatus))
	}
	if c.FeatureStatus.OrchestrationStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration_status").String(), c.FeatureStatus.OrchestrationStatus))
	}
	if c.FeatureStatus.TrapForwarderStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trap_forwarder_status").String(), c.FeatureStatus.TrapForwarderStatus))
	}
	if c.TelemetryCollection != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_collection").String(), strconv.FormatBool(*c.TelemetryCollection)))
	}
	if c.TelemetryCollectionType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_collection_type").String(), c.TelemetryCollectionType))
	}
	if c.TelemetryStreamingProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_streaming_protocol").String(), c.TelemetryStreamingProtocol))
	}
	if c.TelemetrySourceInterface != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_source_interface").String(), c.TelemetrySourceInterface))
	}
	if c.TelemetrySourceVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_source_vrf").String(), c.TelemetrySourceVrf))
	}
	if c.SecurityDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), c.SecurityDomain))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), "all"))
	}

	if c.Management.BgpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_asn").String(), c.Management.BgpAsn))
	}
	if c.Management.TargetSubnetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("target_subnet_mask").String(), strconv.Itoa(int(*c.Management.TargetSubnetMask))))
	}
	if c.Management.AnycastGatewayMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gateway_mac").String(), c.Management.AnycastGatewayMac))
	}
	if c.Management.PerformanceMonitoring != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("performance_monitoring").String(), strconv.FormatBool(*c.Management.PerformanceMonitoring)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("performance_monitoring").String(), "false"))
	}
	if c.Management.ReplicationMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("replication_mode").String(), c.Management.ReplicationMode))
	}
	if c.Management.MulticastGroupSubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multicast_group_subnet").String(), c.Management.MulticastGroupSubnet))
	}
	if c.Management.UnderlayMulticastGroupAddressLimit != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_multicast_group_address_limit").String(), strconv.Itoa(int(*c.Management.UnderlayMulticastGroupAddressLimit))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_multicast_group_address_limit").String(), "128"))
	}
	if c.Management.TenantRoutedMulticast != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_routed_multicast").String(), strconv.FormatBool(*c.Management.TenantRoutedMulticast)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_routed_multicast").String(), "false"))
	}
	if c.Management.RendezvousPointCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_count").String(), strconv.Itoa(int(*c.Management.RendezvousPointCount))))
	}
	if c.Category != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), c.Category))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), "fabric"))
	}
	if c.AlertSuspend != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("alert_suspend").String(), c.AlertSuspend))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("alert_suspend").String(), "disabled"))
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
	if c.Management.VpcPeerLinkEnableNativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_enable_native_vlan").String(), strconv.FormatBool(*c.Management.VpcPeerLinkEnableNativeVlan)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_peer_link_enable_native_vlan").String(), "false"))
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
	if c.Management.VpcIpv6NeighborDiscoverySync != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_ipv6_neighbor_discovery_sync").String(), strconv.FormatBool(*c.Management.VpcIpv6NeighborDiscoverySync)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_ipv6_neighbor_discovery_sync").String(), "true"))
	}
	if c.Management.AdvertisePhysicalIp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_physical_ip").String(), strconv.FormatBool(*c.Management.AdvertisePhysicalIp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_physical_ip").String(), "false"))
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
	if c.Management.L3VniNoVlanDefaultOption != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_no_vlan_default_option").String(), strconv.FormatBool(*c.Management.L3VniNoVlanDefaultOption)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_no_vlan_default_option").String(), "false"))
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
	if c.Management.TenantDhcp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_dhcp").String(), strconv.FormatBool(*c.Management.TenantDhcp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_dhcp").String(), "true"))
	}
	if c.Management.Nxapi != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi").String(), strconv.FormatBool(*c.Management.Nxapi)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi").String(), "false"))
	}
	if c.Management.NxapiHttpsPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), strconv.Itoa(int(*c.Management.NxapiHttpsPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_https_port").String(), "443"))
	}
	if c.Management.NxapiHttp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http").String(), strconv.FormatBool(*c.Management.NxapiHttp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http").String(), "true"))
	}
	if c.Management.NxapiHttpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), strconv.Itoa(int(*c.Management.NxapiHttpPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), "80"))
	}
	if c.Management.SnmpTrap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_trap").String(), strconv.FormatBool(*c.Management.SnmpTrap)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_trap").String(), "true"))
	}
	if c.Management.AnycastBorderGatewayAdvertisePhysicalIp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_border_gateway_advertise_physical_ip").String(), strconv.FormatBool(*c.Management.AnycastBorderGatewayAdvertisePhysicalIp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_border_gateway_advertise_physical_ip").String(), "false"))
	}
	if c.Management.GreenfieldDebugFlag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("greenfield_debug_flag").String(), c.Management.GreenfieldDebugFlag))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("greenfield_debug_flag").String(), "disable"))
	}
	if c.Management.TcamAllocation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tcam_allocation").String(), strconv.FormatBool(*c.Management.TcamAllocation)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tcam_allocation").String(), "true"))
	}
	if c.Management.RealTimeInterfaceStatisticsCollection != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_interface_statistics_collection").String(), strconv.FormatBool(*c.Management.RealTimeInterfaceStatisticsCollection)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_interface_statistics_collection").String(), "false"))
	}
	if c.Management.InterfaceStatisticsLoadInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_statistics_load_interval").String(), strconv.Itoa(int(*c.Management.InterfaceStatisticsLoadInterval))))
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
	if c.Management.AnycastRendezvousPointIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_rendezvous_point_ip_range").String(), c.Management.AnycastRendezvousPointIpRange))
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
	if c.Management.SubInterfaceDot1qRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sub_interface_dot1q_range").String(), c.Management.SubInterfaceDot1qRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sub_interface_dot1q_range").String(), "2-511"))
	}
	if c.Management.VrfLiteAutoConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_auto_config").String(), c.Management.VrfLiteAutoConfig))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_auto_config").String(), "manual"))
	}
	if c.Management.VrfLiteSubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_subnet_range").String(), c.Management.VrfLiteSubnetRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_subnet_range").String(), "10.33.0.0/16"))
	}
	if c.Management.VrfLiteSubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.VrfLiteSubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_subnet_target_mask").String(), "30"))
	}
	if c.Management.VrfLiteIpv6SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_ipv6_subnet_range").String(), c.Management.VrfLiteIpv6SubnetRange))
	}
	if c.Management.VrfLiteIpv6SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_ipv6_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.VrfLiteIpv6SubnetTargetMask))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_ipv6_subnet_target_mask").String(), "126"))
	}
	if c.Management.AutoUniqueVrfLiteIpPrefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_unique_vrf_lite_ip_prefix").String(), strconv.FormatBool(*c.Management.AutoUniqueVrfLiteIpPrefix)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_unique_vrf_lite_ip_prefix").String(), "false"))
	}
	if c.Management.PerVrfLoopbackAutoProvision != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision").String(), strconv.FormatBool(*c.Management.PerVrfLoopbackAutoProvision)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision").String(), "false"))
	}
	if c.Management.PerVrfLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_ip_range").String(), c.Management.PerVrfLoopbackIpRange))
	}
	if c.Management.PerVrfLoopbackAutoProvisionIpv6 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision_ipv6").String(), strconv.FormatBool(*c.Management.PerVrfLoopbackAutoProvisionIpv6)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_auto_provision_ipv6").String(), "false"))
	}
	if c.Management.PerVrfLoopbackIpv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_vrf_loopback_ipv6_range").String(), c.Management.PerVrfLoopbackIpv6Range))
	}
	if c.Management.Banner != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("banner").String(), c.Management.Banner))
	}
	if c.Management.Day0Bootstrap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("day0_bootstrap").String(), strconv.FormatBool(*c.Management.Day0Bootstrap)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("day0_bootstrap").String(), "false"))
	}
	if c.Management.LocalDhcpServer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local_dhcp_server").String(), strconv.FormatBool(*c.Management.LocalDhcpServer)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local_dhcp_server").String(), "false"))
	}
	if c.Management.DhcpProtocolVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_protocol_version").String(), c.Management.DhcpProtocolVersion))
	}
	if c.Management.DhcpStartAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_start_address").String(), c.Management.DhcpStartAddress))
	}
	if c.Management.DhcpEndAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_end_address").String(), c.Management.DhcpEndAddress))
	}
	if c.Management.ManagementGateway != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("management_gateway").String(), c.Management.ManagementGateway))
	}
	if c.Management.ManagementIpv4Prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("management_ipv4_prefix").String(), strconv.Itoa(int(*c.Management.ManagementIpv4Prefix))))
	}
	if c.Management.ManagementIpv6Prefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("management_ipv6_prefix").String(), strconv.Itoa(int(*c.Management.ManagementIpv6Prefix))))
	}
	if c.Management.BootstrapMultiSubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bootstrap_multi_subnet").String(), c.Management.BootstrapMultiSubnet))
	}
	if c.Management.ExtraConfigNxosBootstrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_nxos_bootstrap").String(), c.Management.ExtraConfigNxosBootstrap))
	}
	if c.Management.RealTimeBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_backup").String(), strconv.FormatBool(*c.Management.RealTimeBackup)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_backup").String(), "true"))
	}
	if c.Management.ScheduledBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup").String(), strconv.FormatBool(*c.Management.ScheduledBackup)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup").String(), "false"))
	}
	if c.Management.ScheduledBackupTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup_time").String(), c.Management.ScheduledBackupTime))
	}
	if c.Management.UnderlayIpv6 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_ipv6").String(), strconv.FormatBool(*c.Management.UnderlayIpv6)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_ipv6").String(), "false"))
	}
	if c.Management.Ipv6MulticastGroupSubnet != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_multicast_group_subnet").String(), c.Management.Ipv6MulticastGroupSubnet))
	}
	if c.Management.TenantRoutedMulticastIpv6 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_routed_multicast_ipv6").String(), strconv.FormatBool(*c.Management.TenantRoutedMulticastIpv6)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_routed_multicast_ipv6").String(), "false"))
	}
	if c.Management.MvpnVrfRouteImportId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mvpn_vrf_route_import_id").String(), strconv.FormatBool(*c.Management.MvpnVrfRouteImportId)))
	}
	if c.Management.MvpnVrfRouteImportIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mvpn_vrf_route_import_id_range").String(), c.Management.MvpnVrfRouteImportIdRange))
	}
	if c.Management.VrfRouteImportIdReallocation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_route_import_id_reallocation").String(), strconv.FormatBool(*c.Management.VrfRouteImportIdReallocation)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_route_import_id_reallocation").String(), "false"))
	}
	if c.Management.L3vniMulticastGroup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3vni_multicast_group").String(), c.Management.L3vniMulticastGroup))
	}
	if c.Management.L3VniIpv6MulticastGroup != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_ipv6_multicast_group").String(), c.Management.L3VniIpv6MulticastGroup))
	}
	if c.Management.RendezvousPointMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_mode").String(), c.Management.RendezvousPointMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("rendezvous_point_mode").String(), "asm"))
	}
	if c.Management.AutoGenerateMulticastGroupAddress != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_generate_multicast_group_address").String(), strconv.FormatBool(*c.Management.AutoGenerateMulticastGroupAddress)))
	}
	if c.Management.PhantomRendezvousPointLoopbackId1 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rendezvous_point_loopback_id1").String(), strconv.Itoa(int(*c.Management.PhantomRendezvousPointLoopbackId1))))
	}
	if c.Management.PhantomRendezvousPointLoopbackId2 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rendezvous_point_loopback_id2").String(), strconv.Itoa(int(*c.Management.PhantomRendezvousPointLoopbackId2))))
	}
	if c.Management.PhantomRendezvousPointLoopbackId3 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rendezvous_point_loopback_id3").String(), strconv.Itoa(int(*c.Management.PhantomRendezvousPointLoopbackId3))))
	}
	if c.Management.PhantomRendezvousPointLoopbackId4 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("phantom_rendezvous_point_loopback_id4").String(), strconv.Itoa(int(*c.Management.PhantomRendezvousPointLoopbackId4))))
	}
	if c.Management.AdvertisePhysicalIpOnBorder != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_physical_ip_on_border").String(), strconv.FormatBool(*c.Management.AdvertisePhysicalIpOnBorder)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advertise_physical_ip_on_border").String(), "true"))
	}
	if c.Management.FabricVpcDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_domain_id").String(), strconv.FormatBool(*c.Management.FabricVpcDomainId)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_domain_id").String(), "false"))
	}
	if c.Management.SharedVpcDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("shared_vpc_domain_id").String(), strconv.Itoa(int(*c.Management.SharedVpcDomainId))))
	}
	if c.Management.VpcLayer3PeerRouter != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_layer3_peer_router").String(), strconv.FormatBool(*c.Management.VpcLayer3PeerRouter)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_layer3_peer_router").String(), "true"))
	}
	if c.Management.FabricVpcQos != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos").String(), strconv.FormatBool(*c.Management.FabricVpcQos)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos").String(), "false"))
	}
	if c.Management.FabricVpcQosPolicyName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos_policy_name").String(), c.Management.FabricVpcQosPolicyName))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_vpc_qos_policy_name").String(), "spine_qos_for_fabric_vpc_peering"))
	}
	if c.Management.EnablePeerSwitch != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_peer_switch").String(), strconv.FormatBool(*c.Management.EnablePeerSwitch)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_peer_switch").String(), "false"))
	}
	if c.Management.AnycastLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_loopback_id").String(), strconv.Itoa(int(*c.Management.AnycastLoopbackId))))
	}
	if c.Management.BgpAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_authentication").String(), strconv.FormatBool(*c.Management.BgpAuthentication)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_authentication").String(), "false"))
	}
	if c.Management.BgpAuthenticationKeyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_authentication_key_type").String(), c.Management.BgpAuthenticationKeyType))
	}
	if c.Management.BgpAuthenticationKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_authentication_key").String(), c.Management.BgpAuthenticationKey))
	}
	if c.Management.PimHelloAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_authentication").String(), strconv.FormatBool(*c.Management.PimHelloAuthentication)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_authentication").String(), "false"))
	}
	if c.Management.PimHelloAuthenticationKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pim_hello_authentication_key").String(), c.Management.PimHelloAuthenticationKey))
	}
	if c.Management.Bfd != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd").String(), strconv.FormatBool(*c.Management.Bfd)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd").String(), "false"))
	}
	if c.Management.BfdIbgp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ibgp").String(), strconv.FormatBool(*c.Management.BfdIbgp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ibgp").String(), "false"))
	}
	if c.Management.BfdAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_authentication").String(), strconv.FormatBool(*c.Management.BfdAuthentication)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_authentication").String(), "false"))
	}
	if c.Management.BfdAuthenticationKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_authentication_key_id").String(), strconv.Itoa(int(*c.Management.BfdAuthenticationKeyId))))
	}
	if c.Management.BfdAuthenticationKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_authentication_key").String(), c.Management.BfdAuthenticationKey))
	}
	if c.Management.Macsec != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec").String(), strconv.FormatBool(*c.Management.Macsec)))
	}
	if c.Management.MacsecCipherSuite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_cipher_suite").String(), c.Management.MacsecCipherSuite))
	}
	if c.Management.MacsecKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_key_string").String(), c.Management.MacsecKeyString))
	}
	if c.Management.MacsecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_algorithm").String(), c.Management.MacsecAlgorithm))
	}
	if c.Management.MacsecFallbackKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_key_string").String(), c.Management.MacsecFallbackKeyString))
	}
	if c.Management.MacsecFallbackAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_algorithm").String(), c.Management.MacsecFallbackAlgorithm))
	}
	if c.Management.MacsecReportTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_report_timer").String(), strconv.Itoa(int(*c.Management.MacsecReportTimer))))
	}
	if c.Management.OverlayMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), c.Management.OverlayMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("overlay_mode").String(), "cli"))
	}
	if c.Management.PrivateVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("private_vlan").String(), strconv.FormatBool(*c.Management.PrivateVlan)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("private_vlan").String(), "false"))
	}
	if c.Management.DefaultPrivateVlanSecondaryNetworkTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_private_vlan_secondary_network_template").String(), c.Management.DefaultPrivateVlanSecondaryNetworkTemplate))
	}
	if c.Management.PowerRedundancyMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), c.Management.PowerRedundancyMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), "redundant"))
	}
	if c.Management.CoppPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), c.Management.CoppPolicy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), "strict"))
	}
	if c.Management.NveHoldDownTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_hold_down_timer").String(), strconv.Itoa(int(*c.Management.NveHoldDownTimer))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_hold_down_timer").String(), "180"))
	}
	if c.Management.Cdp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp").String(), strconv.FormatBool(*c.Management.Cdp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp").String(), "false"))
	}
	if c.Management.NextGenerationOam != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("next_generation_oam").String(), strconv.FormatBool(*c.Management.NextGenerationOam)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("next_generation_oam").String(), "true"))
	}
	if c.Management.NgoamSouthBoundLoopDetect != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ngoam_south_bound_loop_detect").String(), strconv.FormatBool(*c.Management.NgoamSouthBoundLoopDetect)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ngoam_south_bound_loop_detect").String(), "false"))
	}
	if c.Management.NgoamSouthBoundLoopDetectProbeInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ngoam_south_bound_loop_detect_probe_interval").String(), strconv.Itoa(int(*c.Management.NgoamSouthBoundLoopDetectProbeInterval))))
	}
	if c.Management.NgoamSouthBoundLoopDetectRecoveryInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ngoam_south_bound_loop_detect_recovery_interval").String(), strconv.Itoa(int(*c.Management.NgoamSouthBoundLoopDetectRecoveryInterval))))
	}
	if c.Management.StrictConfigComplianceMode != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("strict_config_compliance_mode").String(), strconv.FormatBool(*c.Management.StrictConfigComplianceMode)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("strict_config_compliance_mode").String(), "false"))
	}
	if c.Management.AdvancedSshOption != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advanced_ssh_option").String(), strconv.FormatBool(*c.Management.AdvancedSshOption)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advanced_ssh_option").String(), "false"))
	}
	if c.Management.Ptp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp").String(), strconv.FormatBool(*c.Management.Ptp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp").String(), "false"))
	}
	if c.Management.PtpLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_loopback_id").String(), strconv.Itoa(int(*c.Management.PtpLoopbackId))))
	}
	if c.Management.PtpDomainId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_domain_id").String(), strconv.Itoa(int(*c.Management.PtpDomainId))))
	}
	if c.Management.DefaultQueuingPolicy != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy").String(), strconv.FormatBool(*c.Management.DefaultQueuingPolicy)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy").String(), "false"))
	}
	if c.Management.DefaultQueuingPolicyCloudscale != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_cloudscale").String(), c.Management.DefaultQueuingPolicyCloudscale))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_cloudscale").String(), "queuing_policy_default_8q_cloudscale"))
	}
	if c.Management.DefaultQueuingPolicyRSeries != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_r_series").String(), c.Management.DefaultQueuingPolicyRSeries))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_r_series").String(), "queuing_policy_default_r_series"))
	}
	if c.Management.DefaultQueuingPolicyOther != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_other").String(), c.Management.DefaultQueuingPolicyOther))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_queuing_policy_other").String(), "queuing_policy_default_other"))
	}
	if c.Management.AimlQos != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aiml_qos").String(), strconv.FormatBool(*c.Management.AimlQos)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aiml_qos").String(), "false"))
	}
	if c.Management.AimlQosPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aiml_qos_policy").String(), c.Management.AimlQosPolicy))
	}
	if c.Management.PriorityFlowControlWatchInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("priority_flow_control_watch_interval").String(), strconv.Itoa(int(*c.Management.PriorityFlowControlWatchInterval))))
	}
	if c.Management.Dlb != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dlb").String(), strconv.FormatBool(*c.Management.Dlb)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dlb").String(), "false"))
	}
	if c.Management.DlbMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dlb_mode").String(), c.Management.DlbMode))
	}
	if c.Management.DlbMixedModeDefault != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dlb_mixed_mode_default").String(), c.Management.DlbMixedModeDefault))
	}
	if c.Management.FlowletAging != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("flowlet_aging").String(), strconv.Itoa(int(*c.Management.FlowletAging))))
	}
	if c.Management.FlowletDscp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("flowlet_dscp").String(), c.Management.FlowletDscp))
	}
	if c.Management.PerPacketDscp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("per_packet_dscp").String(), c.Management.PerPacketDscp))
	}
	if c.Management.AiLoadSharing != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ai_load_sharing").String(), strconv.FormatBool(*c.Management.AiLoadSharing)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ai_load_sharing").String(), "false"))
	}
	if c.Management.RoceV2 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("roce_v2").String(), c.Management.RoceV2))
	}
	if c.Management.Cnp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cnp").String(), c.Management.Cnp))
	}
	if c.Management.WredMin != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wred_min").String(), strconv.Itoa(int(*c.Management.WredMin))))
	}
	if c.Management.WredMax != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wred_max").String(), strconv.Itoa(int(*c.Management.WredMax))))
	}
	if c.Management.WredDropProbability != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wred_drop_probability").String(), strconv.Itoa(int(*c.Management.WredDropProbability))))
	}
	if c.Management.WredWeight != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("wred_weight").String(), strconv.Itoa(int(*c.Management.WredWeight))))
	}
	if c.Management.BandwidthRemaining != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bandwidth_remaining").String(), strconv.Itoa(int(*c.Management.BandwidthRemaining))))
	}
	if c.Management.StaticUnderlayIpAllocation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_allocation").String(), strconv.FormatBool(*c.Management.StaticUnderlayIpAllocation)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("static_underlay_ip_allocation").String(), "false"))
	}
	if c.Management.BgpLoopbackIpv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_loopback_ipv6_range").String(), c.Management.BgpLoopbackIpv6Range))
	}
	if c.Management.NveLoopbackIpv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nve_loopback_ipv6_range").String(), c.Management.NveLoopbackIpv6Range))
	}
	if c.Management.Ipv6AnycastRendezvousPointIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_anycast_rendezvous_point_ip_range").String(), c.Management.Ipv6AnycastRendezvousPointIpRange))
	}
	if c.Management.ExtraConfigAaa != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_aaa").String(), c.Management.ExtraConfigAaa))
	}
	if c.Management.Aaa != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa").String(), strconv.FormatBool(*c.Management.Aaa)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa").String(), "false"))
	}
	if c.Management.Ipv6LinkLocal != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_link_local").String(), strconv.FormatBool(*c.Management.Ipv6LinkLocal)))
	}
	if c.Management.FabricInterfaceType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_interface_type").String(), c.Management.FabricInterfaceType))
	}
	if c.Management.Ipv6SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.Ipv6SubnetTargetMask))))
	}
	if c.Management.LinkStateRoutingProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_protocol").String(), c.Management.LinkStateRoutingProtocol))
	}
	if c.Management.RouteReflectorCount != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_reflector_count").String(), strconv.Itoa(int(*c.Management.RouteReflectorCount))))
	}
	if c.Management.VpcTorDelayRestoreTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_tor_delay_restore_timer").String(), strconv.Itoa(int(*c.Management.VpcTorDelayRestoreTimer))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vpc_tor_delay_restore_timer").String(), "30"))
	}
	if c.Management.LeafTorIdRange != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("leaf_tor_id_range").String(), strconv.FormatBool(*c.Management.LeafTorIdRange)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("leaf_tor_id_range").String(), "false"))
	}
	if c.Management.LeafTorVpcPortChannelIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("leaf_tor_vpc_port_channel_id_range").String(), c.Management.LeafTorVpcPortChannelIdRange))
	}
	if c.Management.LinkStateRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_state_routing_tag").String(), c.Management.LinkStateRoutingTag))
	}
	if c.Management.OspfAreaId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_area_id").String(), c.Management.OspfAreaId))
	}
	if c.Management.OspfAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_authentication").String(), strconv.FormatBool(*c.Management.OspfAuthentication)))
	}
	if c.Management.OspfAuthenticationKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_authentication_key_id").String(), strconv.Itoa(int(*c.Management.OspfAuthenticationKeyId))))
	}
	if c.Management.OspfAuthenticationKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ospf_authentication_key").String(), c.Management.OspfAuthenticationKey))
	}
	if c.Management.IsisLevel != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_level").String(), c.Management.IsisLevel))
	}
	if c.Management.IsisAreaNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_area_number").String(), c.Management.IsisAreaNumber))
	}
	if c.Management.IsisPointToPoint != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_point_to_point").String(), strconv.FormatBool(*c.Management.IsisPointToPoint)))
	}
	if c.Management.IsisAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_authentication").String(), strconv.FormatBool(*c.Management.IsisAuthentication)))
	}
	if c.Management.IsisAuthenticationKeychainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_authentication_keychain_name").String(), c.Management.IsisAuthenticationKeychainName))
	}
	if c.Management.IsisAuthenticationKeychainKeyId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_authentication_keychain_key_id").String(), strconv.Itoa(int(*c.Management.IsisAuthenticationKeychainKeyId))))
	}
	if c.Management.IsisAuthenticationKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_authentication_key").String(), c.Management.IsisAuthenticationKey))
	}
	if c.Management.IsisOverload != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_overload").String(), strconv.FormatBool(*c.Management.IsisOverload)))
	}
	if c.Management.IsisOverloadElapseTime != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("isis_overload_elapse_time").String(), strconv.Itoa(int(*c.Management.IsisOverloadElapseTime))))
	}
	if c.Management.BfdOspf != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_ospf").String(), strconv.FormatBool(*c.Management.BfdOspf)))
	}
	if c.Management.BfdIsis != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_isis").String(), strconv.FormatBool(*c.Management.BfdIsis)))
	}
	if c.Management.BfdPim != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_pim").String(), strconv.FormatBool(*c.Management.BfdPim)))
	}
	if c.Management.AutoBgpNeighborDescription != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_bgp_neighbor_description").String(), strconv.FormatBool(*c.Management.AutoBgpNeighborDescription)))
	}
	if c.Management.IbgpPeerTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ibgp_peer_template").String(), c.Management.IbgpPeerTemplate))
	}
	if c.Management.LeafibgpPeerTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("leafibgp_peer_template").String(), c.Management.LeafibgpPeerTemplate))
	}
	if c.Management.SecurityGroupTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag").String(), strconv.FormatBool(*c.Management.SecurityGroupTag)))
	}
	if c.Management.SecurityGroupTagPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_prefix").String(), c.Management.SecurityGroupTagPrefix))
	}
	if c.Management.SecurityGroupTagIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_id_range").String(), c.Management.SecurityGroupTagIdRange))
	}
	if c.Management.SecurityGroupTagPreprovision != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_preprovision").String(), strconv.FormatBool(*c.Management.SecurityGroupTagPreprovision)))
	}
	if c.Management.SecurityGroupTagMacSegmentation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_mac_segmentation").String(), strconv.FormatBool(*c.Management.SecurityGroupTagMacSegmentation)))
	}
	if c.Management.SecurityGroupStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_status").String(), c.Management.SecurityGroupStatus))
	}
	if c.Management.VrfLiteMacsec != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec").String(), strconv.FormatBool(*c.Management.VrfLiteMacsec)))
	}
	if c.Management.QuantumKeyDistribution != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("quantum_key_distribution").String(), strconv.FormatBool(*c.Management.QuantumKeyDistribution)))
	}
	if c.Management.VrfLiteMacsecCipherSuite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec_cipher_suite").String(), c.Management.VrfLiteMacsecCipherSuite))
	}
	if c.Management.VrfLiteMacsecKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec_key_string").String(), c.Management.VrfLiteMacsecKeyString))
	}
	if c.Management.VrfLiteMacsecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec_algorithm").String(), c.Management.VrfLiteMacsecAlgorithm))
	}
	if c.Management.VrfLiteMacsecFallbackKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec_fallback_key_string").String(), c.Management.VrfLiteMacsecFallbackKeyString))
	}
	if c.Management.VrfLiteMacsecFallbackAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_lite_macsec_fallback_algorithm").String(), c.Management.VrfLiteMacsecFallbackAlgorithm))
	}
	if c.Management.QuantumKeyDistributionProfileName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("quantum_key_distribution_profile_name").String(), c.Management.QuantumKeyDistributionProfileName))
	}
	if c.Management.KeyManagementEntityServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("key_management_entity_server_ip").String(), c.Management.KeyManagementEntityServerIp))
	}
	if c.Management.KeyManagementEntityServerPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("key_management_entity_server_port").String(), strconv.Itoa(int(*c.Management.KeyManagementEntityServerPort))))
	}
	if c.Management.TrustpointLabel != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trustpoint_label").String(), c.Management.TrustpointLabel))
	}
	if c.Management.SkipCertificateVerification != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("skip_certificate_verification").String(), strconv.FormatBool(*c.Management.SkipCertificateVerification)))
	}
	if c.Management.HostInterfaceAdminState != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("host_interface_admin_state").String(), strconv.FormatBool(*c.Management.HostInterfaceAdminState)))
	}
	if c.Management.BrownfieldNetworkNameFormat != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_network_name_format").String(), c.Management.BrownfieldNetworkNameFormat))
	}
	if c.Management.BrownfieldSkipOverlayNetworkAttachments != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("brownfield_skip_overlay_network_attachments").String(), strconv.FormatBool(*c.Management.BrownfieldSkipOverlayNetworkAttachments)))
	}
	if c.Management.PolicyBasedRouting != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy_based_routing").String(), strconv.FormatBool(*c.Management.PolicyBasedRouting)))
	}
	if c.Management.PtpVlanId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ptp_vlan_id").String(), strconv.Itoa(int(*c.Management.PtpVlanId))))
	}
	if c.Management.MplsHandoff != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), strconv.FormatBool(*c.Management.MplsHandoff)))
	}
	if c.Management.MplsLoopbackIdentifier != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_identifier").String(), strconv.Itoa(int(*c.Management.MplsLoopbackIdentifier))))
	}
	if c.Management.MplsIsisAreaNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_isis_area_number").String(), c.Management.MplsIsisAreaNumber))
	}
	if c.Management.StpRootOption != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_root_option").String(), c.Management.StpRootOption))
	}
	if c.Management.StpVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_vlan_range").String(), c.Management.StpVlanRange))
	}
	if c.Management.MstInstanceRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mst_instance_range").String(), c.Management.MstInstanceRange))
	}
	if c.Management.StpBridgePriority != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("stp_bridge_priority").String(), strconv.Itoa(int(*c.Management.StpBridgePriority))))
	}
	if c.Management.AllowVlanOnLeafTorPairing != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_vlan_on_leaf_tor_pairing").String(), c.Management.AllowVlanOnLeafTorPairing))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_vlan_on_leaf_tor_pairing").String(), "none"))
	}
	if c.Management.PreInterfaceConfigLeaf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pre_interface_config_leaf").String(), c.Management.PreInterfaceConfigLeaf))
	}
	if c.Management.PreInterfaceConfigSpine != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pre_interface_config_spine").String(), c.Management.PreInterfaceConfigSpine))
	}
	if c.Management.PreInterfaceConfigTor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("pre_interface_config_tor").String(), c.Management.PreInterfaceConfigTor))
	}
	if c.Management.ExtraConfigLeaf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_leaf").String(), c.Management.ExtraConfigLeaf))
	}
	if c.Management.ExtraConfigSpine != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_spine").String(), c.Management.ExtraConfigSpine))
	}
	if c.Management.ExtraConfigTor != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_tor").String(), c.Management.ExtraConfigTor))
	}
	if c.Management.ExtraConfigIntraFabricLinks != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_intra_fabric_links").String(), c.Management.ExtraConfigIntraFabricLinks))
	}
	if c.Management.MplsLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_ip_range").String(), c.Management.MplsLoopbackIpRange))
	}
	if c.Management.Ipv6SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_subnet_range").String(), c.Management.Ipv6SubnetRange))
	}
	if c.Management.RouterIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("router_id_range").String(), c.Management.RouterIdRange))
	}
	if c.Management.AutoSymmetricVrfLite != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_vrf_lite").String(), strconv.FormatBool(*c.Management.AutoSymmetricVrfLite)))
	}
	if c.Management.AutoVrfLiteDefaultVrf != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_vrf_lite_default_vrf").String(), strconv.FormatBool(*c.Management.AutoVrfLiteDefaultVrf)))
	}
	if c.Management.AutoSymmetricDefaultVrf != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_symmetric_default_vrf").String(), strconv.FormatBool(*c.Management.AutoSymmetricDefaultVrf)))
	}
	if c.Management.DefaultVrfRedistributionBgpRouteMap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf_redistribution_bgp_route_map").String(), c.Management.DefaultVrfRedistributionBgpRouteMap))
	}
	if c.Management.IpServiceLevelAgreementIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip_service_level_agreement_id_range").String(), c.Management.IpServiceLevelAgreementIdRange))
	}
	if c.Management.ObjectTrackingNumberRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("object_tracking_number_range").String(), c.Management.ObjectTrackingNumberRange))
	}
	if c.Management.ServiceNetworkVlanRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("service_network_vlan_range").String(), c.Management.ServiceNetworkVlanRange))
	}
	if c.Management.RouteMapSequenceNumberRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_map_sequence_number_range").String(), c.Management.RouteMapSequenceNumberRange))
	}
	if c.Management.InbandManagement != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_management").String(), strconv.FormatBool(*c.Management.InbandManagement)))
	}
	if c.Management.SeedSwitchCoreInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("seed_switch_core_interfaces").String(), c.Management.SeedSwitchCoreInterfaces))
	}
	if c.Management.SpineSwitchCoreInterfaces != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("spine_switch_core_interfaces").String(), c.Management.SpineSwitchCoreInterfaces))
	}
	if c.Management.InbandDhcpServers != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_dhcp_servers").String(), c.Management.InbandDhcpServers))
	}
	if c.Management.UnNumberedBootstrapLbId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("un_numbered_bootstrap_lb_id").String(), strconv.Itoa(int(*c.Management.UnNumberedBootstrapLbId))))
	}
	if c.Management.UnNumberedDhcpStartAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("un_numbered_dhcp_start_address").String(), c.Management.UnNumberedDhcpStartAddress))
	}
	if c.Management.UnNumberedDhcpEndAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("un_numbered_dhcp_end_address").String(), c.Management.UnNumberedDhcpEndAddress))
	}
	if c.Management.HeartbeatInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("heartbeat_interval").String(), strconv.Itoa(int(*c.Management.HeartbeatInterval))))
	}
	if c.Management.AllowSmartSwitchOnboarding != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_smart_switch_onboarding").String(), strconv.FormatBool(*c.Management.AllowSmartSwitchOnboarding)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_smart_switch_onboarding").String(), "false"))
	}
	if c.Management.ConnectivityDomainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("connectivity_domain_name").String(), c.Management.ConnectivityDomainName))
	}
	if c.Management.HypershieldConnectivityProxyServer != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hypershield_connectivity_proxy_server").String(), c.Management.HypershieldConnectivityProxyServer))
	}
	if c.Management.HypershieldConnectivityProxyServerPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hypershield_connectivity_proxy_server_port").String(), strconv.Itoa(int(*c.Management.HypershieldConnectivityProxyServerPort))))
	}
	if c.Management.HypershieldConnectivitySourceIntf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hypershield_connectivity_source_intf").String(), c.Management.HypershieldConnectivitySourceIntf))
	}
	if c.Management.NetflowSettings.NetflowEnable != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_enable").String(), strconv.FormatBool(*c.Management.NetflowSettings.NetflowEnable)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_enable").String(), "false"))
	}

	if c.TelemetrySettings.FlowCollection.TrafficAnalytics != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("traffic_analytics").String(), c.TelemetrySettings.FlowCollection.TrafficAnalytics))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("net_flow").String(), strconv.FormatBool(*c.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow)))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("s_flow").String(), strconv.FormatBool(*c.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow)))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("flow_telemetry").String(), strconv.FormatBool(*c.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry)))
	}
	if c.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("traffic_analytics_rules_enabled").String(), strconv.FormatBool(*c.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled)))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("traffic_analytics_mode").String(), c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("udp_categorization").String(), c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization))
	}
	if c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("traffic_analytics_filter_rules").String(), c.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules))
	}
	if c.TelemetrySettings.FlowCollection.OperatingMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("operating_mode").String(), c.TelemetrySettings.FlowCollection.OperatingMode))
	}
	if c.TelemetrySettings.FlowCollection.UdpCategorizationSupport != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("udp_categorization_support").String(), c.TelemetrySettings.FlowCollection.UdpCategorizationSupport))
	}
	if c.TelemetrySettings.Microburst.Microburst != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("microburst").String(), strconv.FormatBool(*c.TelemetrySettings.Microburst.Microburst)))
	}
	if c.TelemetrySettings.Microburst.Sensitivity != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sensitivity").String(), c.TelemetrySettings.Microburst.Sensitivity))
	}
	if c.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("analysis_settings_is_enabled").String(), strconv.FormatBool(*c.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled)))
	}
	if c.TelemetrySettings.Nas.Server != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("server").String(), c.TelemetrySettings.Nas.Server))
	}
	if c.TelemetrySettings.Nas.ExportSettings.ExportType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("export_type").String(), c.TelemetrySettings.Nas.ExportSettings.ExportType))
	}
	if c.TelemetrySettings.Nas.ExportSettings.ExportFormat != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("export_format").String(), c.TelemetrySettings.Nas.ExportSettings.ExportFormat))
	}
	if c.ExternalStreamingSettings.Syslog.SyslogFacility != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("syslog_facility").String(), c.ExternalStreamingSettings.Syslog.SyslogFacility))
	}
	return ret
}

func FabricVxlanIbgpBootstrapSubnetCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCBootstrapSubnetCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.StartIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("start_ip").String(), c.StartIp))
	}
	if c.EndIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("end_ip").String(), c.EndIp))
	}
	if c.DefaultGateway != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_gateway").String(), c.DefaultGateway))
	}
	if c.SubnetPrefix != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("subnet_prefix").String(), strconv.Itoa(int(*c.SubnetPrefix))))
	}
	return ret
}

func FabricVxlanIbgpNetflowExporterCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowExporterCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.ExporterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("exporter_name").String(), c.ExporterName))
	}
	if c.ExporterIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("exporter_ip").String(), c.ExporterIp))
	}
	if c.Vrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf").String(), c.Vrf))
	}
	if c.SourceInterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("source_interface_name").String(), c.SourceInterfaceName))
	}
	if c.UdpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("udp_port").String(), strconv.Itoa(int(*c.UdpPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("udp_port").String(), "1"))
	}
	return ret
}

func FabricVxlanIbgpNetflowRecordCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowRecordCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.RecordName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("record_name").String(), c.RecordName))
	}
	if c.RecordTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("record_template").String(), c.RecordTemplate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("record_template").String(), "netflowIpv4Record"))
	}
	if c.Layer2Record != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("layer2_record").String(), c.Layer2Record))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("layer2_record").String(), "false"))
	}
	return ret
}

func FabricVxlanIbgpNetflowMonitorCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowMonitorCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.MonitorName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("monitor_name").String(), c.MonitorName))
	}
	if c.MonitorRecordName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("monitor_record_name").String(), c.MonitorRecordName))
	}
	if c.Exporter1Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("exporter1_name").String(), c.Exporter1Name))
	}
	if c.Exporter2Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("exporter2_name").String(), c.Exporter2Name))
	}
	return ret
}

func FabricVxlanIbgpNetflowSamplerCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowSamplerCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.SamplerName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sampler_name").String(), c.SamplerName))
	}
	if c.NumSamples != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("num_samples").String(), strconv.Itoa(int(*c.NumSamples))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("num_samples").String(), "1"))
	}
	if c.SamplingRate != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sampling_rate").String(), strconv.Itoa(int(*c.SamplingRate))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sampling_rate").String(), "1"))
	}
	return ret
}

func FabricVxlanIbgpVrfFlowRuleAttributesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCVrfFlowRuleAttributesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.VrfFlowRuleBidirectional != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_bidirectional").String(), strconv.FormatBool(*c.VrfFlowRuleBidirectional)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_bidirectional").String(), "false"))
	}
	if c.VrfFlowRuleDstIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_dst_ip").String(), c.VrfFlowRuleDstIp))
	}
	if c.VrfFlowRuleSrcIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_src_ip").String(), c.VrfFlowRuleSrcIp))
	}
	if c.VrfFlowRuleDstPort != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_dst_port").String(), c.VrfFlowRuleDstPort))
	}
	if c.VrfFlowRuleSrcPort != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_src_port").String(), c.VrfFlowRuleSrcPort))
	}
	if c.VrfFlowRuleProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_protocol").String(), c.VrfFlowRuleProtocol))
	}
	if c.VrfFlowRuleAttributeId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_attribute_id").String(), c.VrfFlowRuleAttributeId))
	}
	return ret
}

func FabricVxlanIbgpVrfFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCVrfFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.VrfFlowRuleName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_name").String(), c.VrfFlowRuleName))
	}
	if c.VrfFlowRuleUuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_uuid").String(), c.VrfFlowRuleUuid))
	}
	if c.VrfFlowRuleTenant != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_tenant").String(), c.VrfFlowRuleTenant))
	}
	if c.VrfFlowRuleVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_flow_rule_vrf").String(), c.VrfFlowRuleVrf))
	}
	return ret
}

func FabricVxlanIbgpInterfaceFlowRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceFlowRuleSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_switch_id").String(), c.InterfaceFlowRuleSwitchId))
	}
	if c.InterfaceFlowRuleSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_switch_name").String(), c.InterfaceFlowRuleSwitchName))
	}
	return ret
}

func FabricVxlanIbgpInterfaceFlowRuleAttributesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRuleAttributesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceFlowRuleBidirectional != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_bidirectional").String(), strconv.FormatBool(*c.InterfaceFlowRuleBidirectional)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_bidirectional").String(), "false"))
	}
	if c.InterfaceFlowRuleDstIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_dst_ip").String(), c.InterfaceFlowRuleDstIp))
	}
	if c.InterfaceFlowRuleSrcIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_src_ip").String(), c.InterfaceFlowRuleSrcIp))
	}
	if c.InterfaceFlowRuleDstPort != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_dst_port").String(), c.InterfaceFlowRuleDstPort))
	}
	if c.InterfaceFlowRuleSrcPort != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_src_port").String(), c.InterfaceFlowRuleSrcPort))
	}
	if c.InterfaceFlowRuleProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_protocol").String(), c.InterfaceFlowRuleProtocol))
	}
	if c.InterfaceFlowRuleAttributeId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_attribute_id").String(), c.InterfaceFlowRuleAttributeId))
	}
	return ret
}

func FabricVxlanIbgpInterfaceFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceFlowRuleName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_name").String(), c.InterfaceFlowRuleName))
	}
	if c.InterfaceFlowRuleUuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_uuid").String(), c.InterfaceFlowRuleUuid))
	}
	if c.InterfaceFlowRuleType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_type").String(), c.InterfaceFlowRuleType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_type").String(), "physical"))
	}
	return ret
}

func FabricVxlanIbgpL3OutFlowRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCL3OutFlowRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.L3OutFlowRuleTenant != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_tenant").String(), c.L3OutFlowRuleTenant))
	}
	if c.L3OutFlowRuleL3Out != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_l3_out").String(), c.L3OutFlowRuleL3Out))
	}
	if c.L3OutFlowRuleEncap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_encap").String(), c.L3OutFlowRuleEncap))
	}
	if c.L3OutFlowRuleSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_switch_name").String(), c.L3OutFlowRuleSwitchName))
	}
	if c.L3OutFlowRuleSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_switch_id").String(), c.L3OutFlowRuleSwitchId))
	}
	return ret
}

func FabricVxlanIbgpL3OutFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCL3OutFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.L3OutFlowRuleName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_name").String(), c.L3OutFlowRuleName))
	}
	if c.L3OutFlowRuleUuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_uuid").String(), c.L3OutFlowRuleUuid))
	}
	if c.L3OutFlowRuleType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_type").String(), c.L3OutFlowRuleType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_out_flow_rule_type").String(), "subInterface"))
	}
	return ret
}

func FabricVxlanIbgpInterfaceRuleInterfacesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRuleInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceRuleInterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_interface_name").String(), c.InterfaceRuleInterfaceName))
	}
	if c.InterfaceRuleInterfaceType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_interface_type").String(), c.InterfaceRuleInterfaceType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_interface_type").String(), "routed"))
	}
	if c.InterfaceRuleInterfaceEncap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_interface_encap").String(), c.InterfaceRuleInterfaceEncap))
	}
	if c.InterfaceRuleInterfaceVrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_interface_vrf_name").String(), c.InterfaceRuleInterfaceVrfName))
	}
	return ret
}

func FabricVxlanIbgpInterfaceRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceRuleSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_switch_id").String(), c.InterfaceRuleSwitchId))
	}
	if c.InterfaceRuleSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_switch_name").String(), c.InterfaceRuleSwitchName))
	}
	if c.InterfaceRuleVrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_vrf_name").String(), c.InterfaceRuleVrfName))
	}
	if c.InterfaceRuleTenant != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_tenant").String(), c.InterfaceRuleTenant))
	}
	if c.InterfaceRuleL3Out != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_l3_out").String(), c.InterfaceRuleL3Out))
	}
	return ret
}

func FabricVxlanIbgpInterfaceRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRulesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceRuleName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_name").String(), c.InterfaceRuleName))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_name").String(), "TAInterfaceRule1"))
	}
	if c.InterfaceRuleEnabled != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enabled").String(), strconv.FormatBool(*c.InterfaceRuleEnabled)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enabled").String(), "true"))
	}
	if c.InterfaceRuleEnableFabricInterconnect != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enable_fabric_interconnect").String(), strconv.FormatBool(*c.InterfaceRuleEnableFabricInterconnect)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enable_fabric_interconnect").String(), "true"))
	}
	if c.InterfaceRuleEnableL3Out != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enable_l3_out").String(), strconv.FormatBool(*c.InterfaceRuleEnableL3Out)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_enable_l3_out").String(), "true"))
	}
	if c.InterfaceRuleUuid != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_rule_uuid").String(), c.InterfaceRuleUuid))
	}
	return ret
}

func FabricVxlanIbgpEmailValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCEmailValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.ReceiverEmail != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("receiver_email").String(), c.ReceiverEmail))
	}
	if c.Format != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("format").String(), c.Format))
	}
	if c.StartDate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("start_date").String(), c.StartDate))
	}
	if c.CollectionFrequencyInDays != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("collection_frequency_in_days").String(), strconv.Itoa(int(*c.CollectionFrequencyInDays))))
	}
	if c.CollectionSettings.CollectionType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("collection_type").String(), c.CollectionSettings.CollectionType))
	}
	if c.OnlyIncludeActiveAlerts != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("only_include_active_alerts").String(), strconv.FormatBool(*c.OnlyIncludeActiveAlerts)))
	}
	return ret
}

func FabricVxlanIbgpMessageBusValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCMessageBusValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Server != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("server").String(), c.Server))
	}
	if c.CollectionType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("collection_type").String(), c.CollectionType))
	}
	if c.CollectionSettings.CollectionSettingsCollectionType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("collection_settings_collection_type").String(), c.CollectionSettings.CollectionSettingsCollectionType))
	}
	return ret
}
