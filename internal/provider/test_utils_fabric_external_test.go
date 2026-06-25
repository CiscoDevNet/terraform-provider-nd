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

func FabricExternalModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
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
	if c.Management.CreateBgpConfig != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("create_bgp_config").String(), strconv.FormatBool(*c.Management.CreateBgpConfig)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("create_bgp_config").String(), "true"))
	}
	if c.Management.Aaa != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("aaa").String(), strconv.FormatBool(*c.Management.Aaa)))
	}
	if c.Management.AdvancedSshOption != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advanced_ssh_option").String(), strconv.FormatBool(*c.Management.AdvancedSshOption)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("advanced_ssh_option").String(), "false"))
	}
	if c.Management.AllowSameLoopbackIpOnSwitches != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_same_loopback_ip_on_switches").String(), strconv.FormatBool(*c.Management.AllowSameLoopbackIpOnSwitches)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("allow_same_loopback_ip_on_switches").String(), "false"))
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
	if c.Management.Day0Bootstrap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("day0_bootstrap").String(), strconv.FormatBool(*c.Management.Day0Bootstrap)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("day0_bootstrap").String(), "false"))
	}
	if c.Management.Day0PlugAndPlay != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("day0_plug_and_play").String(), strconv.FormatBool(*c.Management.Day0PlugAndPlay)))
	}
	if c.Management.InbandDay0Bootstrap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_day0_bootstrap").String(), strconv.FormatBool(*c.Management.InbandDay0Bootstrap)))
	}
	if c.Management.Cdp != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp").String(), strconv.FormatBool(*c.Management.Cdp)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cdp").String(), "false"))
	}
	if c.Management.CoppPolicy != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), c.Management.CoppPolicy))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("copp_policy").String(), "manual"))
	}
	if c.Management.DhcpEndAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_end_address").String(), c.Management.DhcpEndAddress))
	}
	if c.Management.DhcpProtocolVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_protocol_version").String(), c.Management.DhcpProtocolVersion))
	}
	if c.Management.DhcpStartAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_start_address").String(), c.Management.DhcpStartAddress))
	}
	if c.Management.LocalDhcpServer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local_dhcp_server").String(), strconv.FormatBool(*c.Management.LocalDhcpServer)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local_dhcp_server").String(), "false"))
	}
	if c.Management.DomainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("domain_name").String(), c.Management.DomainName))
	}
	if c.Management.EnableDpuPinning != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_dpu_pinning").String(), strconv.FormatBool(*c.Management.EnableDpuPinning)))
	}
	if c.Management.ExtraConfigAaa != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_aaa").String(), c.Management.ExtraConfigAaa))
	}
	if c.Management.ExtraConfigFabric != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_fabric").String(), c.Management.ExtraConfigFabric))
	}
	if c.Management.ExtraConfigNxosBootstrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_nxos_bootstrap").String(), c.Management.ExtraConfigNxosBootstrap))
	}
	if c.Management.ExtraConfigXeBootstrap != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("extra_config_xe_bootstrap").String(), c.Management.ExtraConfigXeBootstrap))
	}
	if c.Management.InbandManagement != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_management").String(), strconv.FormatBool(*c.Management.InbandManagement)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inband_management").String(), "false"))
	}
	if c.Management.RealTimeInterfaceStatisticsCollection != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_interface_statistics_collection").String(), strconv.FormatBool(*c.Management.RealTimeInterfaceStatisticsCollection)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_interface_statistics_collection").String(), "false"))
	}
	if c.Management.InterfaceStatisticsLoadInterval != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_statistics_load_interval").String(), strconv.Itoa(int(*c.Management.InterfaceStatisticsLoadInterval))))
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
	if c.Management.MonitoredMode != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("monitored_mode").String(), strconv.FormatBool(*c.Management.MonitoredMode)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("monitored_mode").String(), "false"))
	}
	if c.Management.MplsHandoff != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), strconv.FormatBool(*c.Management.MplsHandoff)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_handoff").String(), "false"))
	}
	if c.Management.MplsLoopbackIdentifier != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_identifier").String(), strconv.Itoa(int(*c.Management.MplsLoopbackIdentifier))))
	}
	if c.Management.MplsLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mpls_loopback_ip_range").String(), c.Management.MplsLoopbackIpRange))
	}
	if c.Management.NetflowSettings.NetflowEnable != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_enable").String(), strconv.FormatBool(*c.Management.NetflowSettings.NetflowEnable)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_enable").String(), "false"))
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
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http").String(), "false"))
	}
	if c.Management.NxapiHttpPort != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), strconv.Itoa(int(*c.Management.NxapiHttpPort))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("nxapi_http_port").String(), "80"))
	}
	if c.Management.PerformanceMonitoring != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("performance_monitoring").String(), strconv.FormatBool(*c.Management.PerformanceMonitoring)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("performance_monitoring").String(), "false"))
	}
	if c.Management.PowerRedundancyMode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), c.Management.PowerRedundancyMode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("power_redundancy_mode").String(), "redundant"))
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
	if c.Management.RealTimeBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("real_time_backup").String(), strconv.FormatBool(*c.Management.RealTimeBackup)))
	}
	if c.Management.ScheduledBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup").String(), strconv.FormatBool(*c.Management.ScheduledBackup)))
	}
	if c.Management.ScheduledBackupTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup_time").String(), c.Management.ScheduledBackupTime))
	}
	if c.Management.SnmpTrap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_trap").String(), strconv.FormatBool(*c.Management.SnmpTrap)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_trap").String(), "true"))
	}
	if c.Management.SubInterfaceDot1qRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sub_interface_dot1q_range").String(), c.Management.SubInterfaceDot1qRange))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("sub_interface_dot1q_range").String(), "2-511"))
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

func FabricExternalBootstrapSubnetCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCBootstrapSubnetCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalNetflowExporterCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowExporterCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalNetflowRecordCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowRecordCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalNetflowMonitorCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCNetflowMonitorCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalVrfFlowRuleAttributesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCVrfFlowRuleAttributesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalVrfFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCVrfFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalInterfaceFlowRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.InterfaceFlowRuleSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_switch_id").String(), c.InterfaceFlowRuleSwitchId))
	}
	if c.InterfaceFlowRuleSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_flow_rule_switch_name").String(), c.InterfaceFlowRuleSwitchName))
	}
	return ret
}

func FabricExternalInterfaceFlowRuleAttributesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRuleAttributesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalInterfaceFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalL3OutFlowRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCL3OutFlowRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalL3OutFlowRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCL3OutFlowRulesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalInterfaceRuleInterfacesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRuleInterfacesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalInterfaceRuleInterfaceCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRuleInterfaceCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalInterfaceRulesValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCInterfaceRulesValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalEmailValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCEmailValue, attrPath path.Path) []resource.TestCheckFunc {
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

func FabricExternalMessageBusValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCMessageBusValue, attrPath path.Path) []resource.TestCheckFunc {
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
