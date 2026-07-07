// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_external

import (
	"context"
	"log"
	"strconv"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v *FabricExternalModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.LicenseTier != "" {
		v.LicenseTier = types.StringValue(jsonData.LicenseTier)
	} else {
		v.LicenseTier = types.StringNull()
	}

	if jsonData.FeatureStatus.ControllerStatus != "" {
		v.ControllerStatus = types.StringValue(jsonData.FeatureStatus.ControllerStatus)

	} else {
		v.ControllerStatus = types.StringNull()
	}

	if jsonData.FeatureStatus.TelemetryStatus != "" {
		v.TelemetryStatus = types.StringValue(jsonData.FeatureStatus.TelemetryStatus)

	} else {
		v.TelemetryStatus = types.StringNull()
	}

	if jsonData.FeatureStatus.OrchestrationStatus != "" {
		v.OrchestrationStatus = types.StringValue(jsonData.FeatureStatus.OrchestrationStatus)

	} else {
		v.OrchestrationStatus = types.StringNull()
	}

	if jsonData.FeatureStatus.TrapForwarderStatus != "" {
		v.TrapForwarderStatus = types.StringValue(jsonData.FeatureStatus.TrapForwarderStatus)

	} else {
		v.TrapForwarderStatus = types.StringNull()
	}

	if jsonData.TelemetryCollection != nil {
		v.TelemetryCollection = types.BoolValue(*jsonData.TelemetryCollection)

	} else {
		v.TelemetryCollection = types.BoolNull()
	}

	if jsonData.TelemetryCollectionType != "" {
		v.TelemetryCollectionType = types.StringValue(jsonData.TelemetryCollectionType)
	} else {
		v.TelemetryCollectionType = types.StringNull()
	}

	if jsonData.TelemetryStreamingProtocol != "" {
		v.TelemetryStreamingProtocol = types.StringValue(jsonData.TelemetryStreamingProtocol)
	} else {
		v.TelemetryStreamingProtocol = types.StringNull()
	}

	if jsonData.TelemetrySourceInterface != "" {
		v.TelemetrySourceInterface = types.StringValue(jsonData.TelemetrySourceInterface)
	} else {
		v.TelemetrySourceInterface = types.StringNull()
	}

	if jsonData.TelemetrySourceVrf != "" {
		v.TelemetrySourceVrf = types.StringValue(jsonData.TelemetrySourceVrf)
	} else {
		v.TelemetrySourceVrf = types.StringNull()
	}

	if jsonData.SecurityDomain != "" {
		v.SecurityDomain = types.StringValue(jsonData.SecurityDomain)
	} else {
		v.SecurityDomain = types.StringNull()
	}

	if len(jsonData.Meta.AllowedActions) == 0 {
		log.Printf("v.AllowedActions is empty")
		v.AllowedActions = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Meta.AllowedActions))
		for i, item := range jsonData.Meta.AllowedActions {
			listData[i] = types.StringValue(item)
		}
		v.AllowedActions, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.Management.BgpAsn != "" {
		v.BgpAsn = types.StringValue(jsonData.Management.BgpAsn)

	} else {
		v.BgpAsn = types.StringNull()
	}

	if jsonData.Category != "" {
		v.Category = types.StringValue(jsonData.Category)
	} else {
		v.Category = types.StringNull()
	}

	v.Location.SetValue(&jsonData.Location)
	v.Location.state = attr.ValueStateKnown

	if jsonData.AlertSuspend != "" {
		v.AlertSuspend = types.StringValue(jsonData.AlertSuspend)
	} else {
		v.AlertSuspend = types.StringNull()
	}

	if jsonData.Management.CreateBgpConfig != nil {
		v.CreateBgpConfig = types.BoolValue(*jsonData.Management.CreateBgpConfig)

	} else {
		v.CreateBgpConfig = types.BoolNull()
	}

	if jsonData.Management.Aaa != nil {
		v.Aaa = types.BoolValue(*jsonData.Management.Aaa)

	} else {
		v.Aaa = types.BoolNull()
	}

	if jsonData.Management.AdvancedSshOption != nil {
		v.AdvancedSshOption = types.BoolValue(*jsonData.Management.AdvancedSshOption)

	} else {
		v.AdvancedSshOption = types.BoolNull()
	}

	if jsonData.Management.AllowSameLoopbackIpOnSwitches != nil {
		v.AllowSameLoopbackIpOnSwitches = types.BoolValue(*jsonData.Management.AllowSameLoopbackIpOnSwitches)

	} else {
		v.AllowSameLoopbackIpOnSwitches = types.BoolNull()
	}

	if jsonData.Management.AllowSmartSwitchOnboarding != nil {
		v.AllowSmartSwitchOnboarding = types.BoolValue(*jsonData.Management.AllowSmartSwitchOnboarding)

	} else {
		v.AllowSmartSwitchOnboarding = types.BoolNull()
	}

	if jsonData.Management.ConnectivityDomainName != "" {
		v.ConnectivityDomainName = types.StringValue(jsonData.Management.ConnectivityDomainName)

	} else {
		v.ConnectivityDomainName = types.StringNull()
	}

	if jsonData.Management.HypershieldConnectivityProxyServer != "" {
		v.HypershieldConnectivityProxyServer = types.StringValue(jsonData.Management.HypershieldConnectivityProxyServer)

	} else {
		v.HypershieldConnectivityProxyServer = types.StringNull()
	}

	if jsonData.Management.HypershieldConnectivityProxyServerPort != nil {
		v.HypershieldConnectivityProxyServerPort = types.Int64Value(*jsonData.Management.HypershieldConnectivityProxyServerPort)

	} else {
		v.HypershieldConnectivityProxyServerPort = types.Int64Null()
	}

	if jsonData.Management.HypershieldConnectivitySourceIntf != "" {
		v.HypershieldConnectivitySourceIntf = types.StringValue(jsonData.Management.HypershieldConnectivitySourceIntf)

	} else {
		v.HypershieldConnectivitySourceIntf = types.StringNull()
	}

	if jsonData.Management.Day0Bootstrap != nil {
		v.Day0Bootstrap = types.BoolValue(*jsonData.Management.Day0Bootstrap)

	} else {
		v.Day0Bootstrap = types.BoolNull()
	}

	if jsonData.Management.Day0PlugAndPlay != nil {
		v.Day0PlugAndPlay = types.BoolValue(*jsonData.Management.Day0PlugAndPlay)

	} else {
		v.Day0PlugAndPlay = types.BoolNull()
	}

	if jsonData.Management.InbandDay0Bootstrap != nil {
		v.InbandDay0Bootstrap = types.BoolValue(*jsonData.Management.InbandDay0Bootstrap)

	} else {
		v.InbandDay0Bootstrap = types.BoolNull()
	}

	if len(jsonData.Management.BootstrapSubnetCollection) == 0 {
		log.Printf("v.BootstrapSubnetCollection is empty")
		v.BootstrapSubnetCollection = types.ListNull(BootstrapSubnetCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]BootstrapSubnetCollectionValue, len(jsonData.Management.BootstrapSubnetCollection))
		for i, item := range jsonData.Management.BootstrapSubnetCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.BootstrapSubnetCollection, err = types.ListValueFrom(context.Background(), BootstrapSubnetCollectionValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if jsonData.Management.Cdp != nil {
		v.Cdp = types.BoolValue(*jsonData.Management.Cdp)

	} else {
		v.Cdp = types.BoolNull()
	}

	if jsonData.Management.CoppPolicy != "" {
		v.CoppPolicy = types.StringValue(jsonData.Management.CoppPolicy)

	} else {
		v.CoppPolicy = types.StringNull()
	}

	if jsonData.Management.DhcpEndAddress != "" {
		v.DhcpEndAddress = types.StringValue(jsonData.Management.DhcpEndAddress)

	} else {
		v.DhcpEndAddress = types.StringNull()
	}

	if jsonData.Management.DhcpProtocolVersion != "" {
		v.DhcpProtocolVersion = types.StringValue(jsonData.Management.DhcpProtocolVersion)

	} else {
		v.DhcpProtocolVersion = types.StringNull()
	}

	if jsonData.Management.DhcpStartAddress != "" {
		v.DhcpStartAddress = types.StringValue(jsonData.Management.DhcpStartAddress)

	} else {
		v.DhcpStartAddress = types.StringNull()
	}

	if jsonData.Management.LocalDhcpServer != nil {
		v.LocalDhcpServer = types.BoolValue(*jsonData.Management.LocalDhcpServer)

	} else {
		v.LocalDhcpServer = types.BoolNull()
	}

	if len(jsonData.Management.DnsCollection) == 0 {
		log.Printf("v.DnsCollection is empty")
		v.DnsCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.DnsCollection))
		for i, item := range jsonData.Management.DnsCollection {
			listData[i] = types.StringValue(item)
		}
		v.DnsCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.Management.DnsVrfCollection) == 0 {
		log.Printf("v.DnsVrfCollection is empty")
		v.DnsVrfCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.DnsVrfCollection))
		for i, item := range jsonData.Management.DnsVrfCollection {
			listData[i] = types.StringValue(item)
		}
		v.DnsVrfCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.Management.DomainName != "" {
		v.DomainName = types.StringValue(jsonData.Management.DomainName)

	} else {
		v.DomainName = types.StringNull()
	}

	if jsonData.Management.EnableDpuPinning != nil {
		v.EnableDpuPinning = types.BoolValue(*jsonData.Management.EnableDpuPinning)

	} else {
		v.EnableDpuPinning = types.BoolNull()
	}

	if jsonData.Management.ExtraConfigAaa != "" {
		v.ExtraConfigAaa = types.StringValue(jsonData.Management.ExtraConfigAaa)

	} else {
		v.ExtraConfigAaa = types.StringNull()
	}

	if jsonData.Management.ExtraConfigFabric != "" {
		v.ExtraConfigFabric = types.StringValue(jsonData.Management.ExtraConfigFabric)

	} else {
		v.ExtraConfigFabric = types.StringNull()
	}

	if jsonData.Management.ExtraConfigNxosBootstrap != "" {
		v.ExtraConfigNxosBootstrap = types.StringValue(jsonData.Management.ExtraConfigNxosBootstrap)

	} else {
		v.ExtraConfigNxosBootstrap = types.StringNull()
	}

	if jsonData.Management.ExtraConfigXeBootstrap != "" {
		v.ExtraConfigXeBootstrap = types.StringValue(jsonData.Management.ExtraConfigXeBootstrap)

	} else {
		v.ExtraConfigXeBootstrap = types.StringNull()
	}

	if jsonData.Management.InbandManagement != nil {
		v.InbandManagement = types.BoolValue(*jsonData.Management.InbandManagement)

	} else {
		v.InbandManagement = types.BoolNull()
	}

	if jsonData.Management.RealTimeInterfaceStatisticsCollection != nil {
		v.RealTimeInterfaceStatisticsCollection = types.BoolValue(*jsonData.Management.RealTimeInterfaceStatisticsCollection)

	} else {
		v.RealTimeInterfaceStatisticsCollection = types.BoolNull()
	}

	if jsonData.Management.InterfaceStatisticsLoadInterval != nil {
		v.InterfaceStatisticsLoadInterval = types.Int64Value(*jsonData.Management.InterfaceStatisticsLoadInterval)

	} else {
		v.InterfaceStatisticsLoadInterval = types.Int64Null()
	}

	if jsonData.Management.ManagementGateway != "" {
		v.ManagementGateway = types.StringValue(jsonData.Management.ManagementGateway)

	} else {
		v.ManagementGateway = types.StringNull()
	}

	if jsonData.Management.ManagementIpv4Prefix != nil {
		v.ManagementIpv4Prefix = types.Int64Value(*jsonData.Management.ManagementIpv4Prefix)

	} else {
		v.ManagementIpv4Prefix = types.Int64Null()
	}

	if jsonData.Management.ManagementIpv6Prefix != nil {
		v.ManagementIpv6Prefix = types.Int64Value(*jsonData.Management.ManagementIpv6Prefix)

	} else {
		v.ManagementIpv6Prefix = types.Int64Null()
	}

	if jsonData.Management.MonitoredMode != nil {
		v.MonitoredMode = types.BoolValue(*jsonData.Management.MonitoredMode)

	} else {
		v.MonitoredMode = types.BoolNull()
	}

	if jsonData.Management.MplsHandoff != nil {
		v.MplsHandoff = types.BoolValue(*jsonData.Management.MplsHandoff)

	} else {
		v.MplsHandoff = types.BoolNull()
	}

	if jsonData.Management.MplsLoopbackIdentifier != nil {
		v.MplsLoopbackIdentifier = types.Int64Value(*jsonData.Management.MplsLoopbackIdentifier)

	} else {
		v.MplsLoopbackIdentifier = types.Int64Null()
	}

	if jsonData.Management.MplsLoopbackIpRange != "" {
		v.MplsLoopbackIpRange = types.StringValue(jsonData.Management.MplsLoopbackIpRange)

	} else {
		v.MplsLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.NetflowSettings.NetflowEnable != nil {
		v.NetflowEnable = types.BoolValue(*jsonData.Management.NetflowSettings.NetflowEnable)

	} else {
		v.NetflowEnable = types.BoolNull()
	}

	if len(jsonData.Management.NetflowSettings.NetflowExporterCollection) == 0 {
		log.Printf("v.NetflowExporterCollection is empty")
		v.NetflowExporterCollection = types.ListNull(NetflowExporterCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]NetflowExporterCollectionValue, len(jsonData.Management.NetflowSettings.NetflowExporterCollection))
		for i, item := range jsonData.Management.NetflowSettings.NetflowExporterCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.NetflowExporterCollection, err = types.ListValueFrom(context.Background(), NetflowExporterCollectionValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if len(jsonData.Management.NetflowSettings.NetflowRecordCollection) == 0 {
		log.Printf("v.NetflowRecordCollection is empty")
		v.NetflowRecordCollection = types.ListNull(NetflowRecordCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]NetflowRecordCollectionValue, len(jsonData.Management.NetflowSettings.NetflowRecordCollection))
		for i, item := range jsonData.Management.NetflowSettings.NetflowRecordCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.NetflowRecordCollection, err = types.ListValueFrom(context.Background(), NetflowRecordCollectionValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if len(jsonData.Management.NetflowSettings.NetflowMonitorCollection) == 0 {
		log.Printf("v.NetflowMonitorCollection is empty")
		v.NetflowMonitorCollection = types.ListNull(NetflowMonitorCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]NetflowMonitorCollectionValue, len(jsonData.Management.NetflowSettings.NetflowMonitorCollection))
		for i, item := range jsonData.Management.NetflowSettings.NetflowMonitorCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.NetflowMonitorCollection, err = types.ListValueFrom(context.Background(), NetflowMonitorCollectionValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if jsonData.Management.Nxapi != nil {
		v.Nxapi = types.BoolValue(*jsonData.Management.Nxapi)

	} else {
		v.Nxapi = types.BoolNull()
	}

	if jsonData.Management.NxapiHttpsPort != nil {
		v.NxapiHttpsPort = types.Int64Value(*jsonData.Management.NxapiHttpsPort)

	} else {
		v.NxapiHttpsPort = types.Int64Null()
	}

	if jsonData.Management.NxapiHttp != nil {
		v.NxapiHttp = types.BoolValue(*jsonData.Management.NxapiHttp)

	} else {
		v.NxapiHttp = types.BoolNull()
	}

	if jsonData.Management.NxapiHttpPort != nil {
		v.NxapiHttpPort = types.Int64Value(*jsonData.Management.NxapiHttpPort)

	} else {
		v.NxapiHttpPort = types.Int64Null()
	}

	if jsonData.Management.PerformanceMonitoring != nil {
		v.PerformanceMonitoring = types.BoolValue(*jsonData.Management.PerformanceMonitoring)

	} else {
		v.PerformanceMonitoring = types.BoolNull()
	}

	if jsonData.Management.PowerRedundancyMode != "" {
		v.PowerRedundancyMode = types.StringValue(jsonData.Management.PowerRedundancyMode)

	} else {
		v.PowerRedundancyMode = types.StringNull()
	}

	if jsonData.Management.Ptp != nil {
		v.Ptp = types.BoolValue(*jsonData.Management.Ptp)

	} else {
		v.Ptp = types.BoolNull()
	}

	if jsonData.Management.PtpLoopbackId != nil {
		v.PtpLoopbackId = types.Int64Value(*jsonData.Management.PtpLoopbackId)

	} else {
		v.PtpLoopbackId = types.Int64Null()
	}

	if jsonData.Management.PtpDomainId != nil {
		v.PtpDomainId = types.Int64Value(*jsonData.Management.PtpDomainId)

	} else {
		v.PtpDomainId = types.Int64Null()
	}

	if jsonData.Management.RealTimeBackup != nil {
		v.RealTimeBackup = types.BoolValue(*jsonData.Management.RealTimeBackup)

	} else {
		v.RealTimeBackup = types.BoolNull()
	}

	if jsonData.Management.ScheduledBackup != nil {
		v.ScheduledBackup = types.BoolValue(*jsonData.Management.ScheduledBackup)

	} else {
		v.ScheduledBackup = types.BoolNull()
	}

	if jsonData.Management.ScheduledBackupTime != "" {
		v.ScheduledBackupTime = types.StringValue(jsonData.Management.ScheduledBackupTime)

	} else {
		v.ScheduledBackupTime = types.StringNull()
	}

	if jsonData.Management.SnmpTrap != nil {
		v.SnmpTrap = types.BoolValue(*jsonData.Management.SnmpTrap)

	} else {
		v.SnmpTrap = types.BoolNull()
	}

	if jsonData.Management.SubInterfaceDot1qRange != "" {
		v.SubInterfaceDot1qRange = types.StringValue(jsonData.Management.SubInterfaceDot1qRange)

	} else {
		v.SubInterfaceDot1qRange = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.TrafficAnalytics != "" {
		v.TrafficAnalytics = types.StringValue(jsonData.TelemetrySettings.FlowCollection.TrafficAnalytics)

	} else {
		v.TrafficAnalytics = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow != nil {
		v.NetFlow = types.BoolValue(*jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow)

	} else {
		v.NetFlow = types.BoolNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow != nil {
		v.SFlow = types.BoolValue(*jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow)

	} else {
		v.SFlow = types.BoolNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry != nil {
		v.FlowTelemetry = types.BoolValue(*jsonData.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry)

	} else {
		v.FlowTelemetry = types.BoolNull()
	}

	if len(jsonData.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules) == 0 {
		log.Printf("v.VrfFlowRules is empty")
		v.VrfFlowRules = types.ListNull(VrfFlowRulesValue{}.Type(context.Background()))
	} else {
		listData := make([]VrfFlowRulesValue, len(jsonData.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules))
		for i, item := range jsonData.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.VrfFlowRules, err = types.ListValueFrom(context.Background(), VrfFlowRulesValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if len(jsonData.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules) == 0 {
		log.Printf("v.InterfaceFlowRules is empty")
		v.InterfaceFlowRules = types.ListNull(InterfaceFlowRulesValue{}.Type(context.Background()))
	} else {
		listData := make([]InterfaceFlowRulesValue, len(jsonData.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules))
		for i, item := range jsonData.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.InterfaceFlowRules, err = types.ListValueFrom(context.Background(), InterfaceFlowRulesValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if len(jsonData.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules) == 0 {
		log.Printf("v.L3OutFlowRules is empty")
		v.L3OutFlowRules = types.ListNull(L3OutFlowRulesValue{}.Type(context.Background()))
	} else {
		listData := make([]L3OutFlowRulesValue, len(jsonData.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules))
		for i, item := range jsonData.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.L3OutFlowRules, err = types.ListValueFrom(context.Background(), L3OutFlowRulesValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if jsonData.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled != nil {
		v.TrafficAnalyticsRulesEnabled = types.BoolValue(*jsonData.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled)

	} else {
		v.TrafficAnalyticsRulesEnabled = types.BoolNull()
	}

	if len(jsonData.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules) == 0 {
		log.Printf("v.InterfaceRules is empty")
		v.InterfaceRules = types.ListNull(InterfaceRulesValue{}.Type(context.Background()))
	} else {
		listData := make([]InterfaceRulesValue, len(jsonData.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules))
		for i, item := range jsonData.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.InterfaceRules, err = types.ListValueFrom(context.Background(), InterfaceRulesValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}
	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode != "" {
		v.TrafficAnalyticsMode = types.StringValue(jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode)

	} else {
		v.TrafficAnalyticsMode = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization != "" {
		v.UdpCategorization = types.StringValue(jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization)

	} else {
		v.UdpCategorization = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules != "" {
		v.TrafficAnalyticsFilterRules = types.StringValue(jsonData.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules)

	} else {
		v.TrafficAnalyticsFilterRules = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.OperatingMode != "" {
		v.OperatingMode = types.StringValue(jsonData.TelemetrySettings.FlowCollection.OperatingMode)

	} else {
		v.OperatingMode = types.StringNull()
	}

	if jsonData.TelemetrySettings.FlowCollection.UdpCategorizationSupport != "" {
		v.UdpCategorizationSupport = types.StringValue(jsonData.TelemetrySettings.FlowCollection.UdpCategorizationSupport)

	} else {
		v.UdpCategorizationSupport = types.StringNull()
	}

	if jsonData.TelemetrySettings.Microburst.Microburst != nil {
		v.Microburst = types.BoolValue(*jsonData.TelemetrySettings.Microburst.Microburst)

	} else {
		v.Microburst = types.BoolNull()
	}

	if jsonData.TelemetrySettings.Microburst.Sensitivity != "" {
		v.Sensitivity = types.StringValue(jsonData.TelemetrySettings.Microburst.Sensitivity)

	} else {
		v.Sensitivity = types.StringNull()
	}

	if jsonData.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled != nil {
		v.AnalysisSettingsIsEnabled = types.BoolValue(*jsonData.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled)

	} else {
		v.AnalysisSettingsIsEnabled = types.BoolNull()
	}

	if jsonData.TelemetrySettings.Nas.Server != "" {
		v.Server = types.StringValue(jsonData.TelemetrySettings.Nas.Server)

	} else {
		v.Server = types.StringNull()
	}

	if jsonData.TelemetrySettings.Nas.ExportSettings.ExportType != "" {
		v.ExportType = types.StringValue(jsonData.TelemetrySettings.Nas.ExportSettings.ExportType)

	} else {
		v.ExportType = types.StringNull()
	}

	if jsonData.TelemetrySettings.Nas.ExportSettings.ExportFormat != "" {
		v.ExportFormat = types.StringValue(jsonData.TelemetrySettings.Nas.ExportSettings.ExportFormat)

	} else {
		v.ExportFormat = types.StringNull()
	}

	if jsonData.TelemetrySettings.EnergyManagement.Cost != nil {
		v.Cost = types.Float64Value(float64(*jsonData.TelemetrySettings.EnergyManagement.Cost))

	} else {
		v.Cost = types.Float64Null()
	}

	if len(jsonData.ExternalStreamingSettings.Email) == 0 {
		log.Printf("v.Email is empty")
		v.Email = types.ListNull(EmailValue{}.Type(context.Background()))
	} else {
		listData := make([]EmailValue, len(jsonData.ExternalStreamingSettings.Email))
		for i, item := range jsonData.ExternalStreamingSettings.Email {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.Email, err = types.ListValueFrom(context.Background(), EmailValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}

	if len(jsonData.ExternalStreamingSettings.Syslog.SyslogServers) == 0 {
		log.Printf("v.SyslogServers is empty")
		v.SyslogServers = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.ExternalStreamingSettings.Syslog.SyslogServers))
		for i, item := range jsonData.ExternalStreamingSettings.Syslog.SyslogServers {
			listData[i] = types.StringValue(item)
		}
		v.SyslogServers, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.ExternalStreamingSettings.Syslog.SyslogFacility != "" {
		v.SyslogFacility = types.StringValue(jsonData.ExternalStreamingSettings.Syslog.SyslogFacility)

	} else {
		v.SyslogFacility = types.StringNull()
	}

	if len(jsonData.ExternalStreamingSettings.Syslog.CollectionSettings.SyslogAnomalies) == 0 {
		log.Printf("v.SyslogAnomalies is empty")
		v.SyslogAnomalies = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.ExternalStreamingSettings.Syslog.CollectionSettings.SyslogAnomalies))
		for i, item := range jsonData.ExternalStreamingSettings.Syslog.CollectionSettings.SyslogAnomalies {
			listData[i] = types.StringValue(item)
		}
		v.SyslogAnomalies, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.ExternalStreamingSettings.MessageBus) == 0 {
		log.Printf("v.MessageBus is empty")
		v.MessageBus = types.ListNull(MessageBusValue{}.Type(context.Background()))
	} else {
		listData := make([]MessageBusValue, len(jsonData.ExternalStreamingSettings.MessageBus))
		for i, item := range jsonData.ExternalStreamingSettings.MessageBus {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.MessageBus, err = types.ListValueFrom(context.Background(), MessageBusValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
	}

	return err
}

func (v *LocationValue) SetValue(jsonData *resource_fabric_common.NDFCLocationValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Latitude != nil {
		v.Latitude = types.Float64Value(float64(*jsonData.Latitude))
	} else {
		v.Latitude = types.Float64Null()
	}

	if jsonData.Longitude != nil {
		v.Longitude = types.Float64Value(float64(*jsonData.Longitude))
	} else {
		v.Longitude = types.Float64Null()
	}

	return err
}

func (v *BootstrapSubnetCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCBootstrapSubnetCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.StartIp != "" {
		v.StartIp = types.StringValue(jsonData.StartIp)
	} else {
		v.StartIp = types.StringNull()
	}

	if jsonData.EndIp != "" {
		v.EndIp = types.StringValue(jsonData.EndIp)
	} else {
		v.EndIp = types.StringNull()
	}

	if jsonData.DefaultGateway != "" {
		v.DefaultGateway = types.StringValue(jsonData.DefaultGateway)
	} else {
		v.DefaultGateway = types.StringNull()
	}

	if jsonData.SubnetPrefix != nil {
		v.SubnetPrefix = types.Int64Value(*jsonData.SubnetPrefix)

	} else {
		v.SubnetPrefix = types.Int64Null()
	}

	return err
}

func (v *NetflowExporterCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCNetflowExporterCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.ExporterName != "" {
		v.ExporterName = types.StringValue(jsonData.ExporterName)
	} else {
		v.ExporterName = types.StringNull()
	}

	if jsonData.ExporterIp != "" {
		v.ExporterIp = types.StringValue(jsonData.ExporterIp)
	} else {
		v.ExporterIp = types.StringNull()
	}

	if jsonData.Vrf != "" {
		v.Vrf = types.StringValue(jsonData.Vrf)
	} else {
		v.Vrf = types.StringNull()
	}

	if jsonData.SourceInterfaceName != "" {
		v.SourceInterfaceName = types.StringValue(jsonData.SourceInterfaceName)
	} else {
		v.SourceInterfaceName = types.StringNull()
	}

	if jsonData.UdpPort != nil {
		v.UdpPort = types.Int64Value(*jsonData.UdpPort)

	} else {
		v.UdpPort = types.Int64Null()
	}

	return err
}

func (v *NetflowRecordCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCNetflowRecordCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.RecordName != "" {
		v.RecordName = types.StringValue(jsonData.RecordName)
	} else {
		v.RecordName = types.StringNull()
	}

	if jsonData.RecordTemplate != "" {
		v.RecordTemplate = types.StringValue(jsonData.RecordTemplate)
	} else {
		v.RecordTemplate = types.StringNull()
	}

	if jsonData.Layer2Record != "" {
		x, _ := strconv.ParseBool(jsonData.Layer2Record)
		v.Layer2Record = types.BoolValue(x)
	} else {
		v.Layer2Record = types.BoolNull()
	}

	return err
}

func (v *NetflowMonitorCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCNetflowMonitorCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.MonitorName != "" {
		v.MonitorName = types.StringValue(jsonData.MonitorName)
	} else {
		v.MonitorName = types.StringNull()
	}

	if jsonData.MonitorRecordName != "" {
		v.MonitorRecordName = types.StringValue(jsonData.MonitorRecordName)
	} else {
		v.MonitorRecordName = types.StringNull()
	}

	if jsonData.Exporter1Name != "" {
		v.Exporter1Name = types.StringValue(jsonData.Exporter1Name)
	} else {
		v.Exporter1Name = types.StringNull()
	}

	if jsonData.Exporter2Name != "" {
		v.Exporter2Name = types.StringValue(jsonData.Exporter2Name)
	} else {
		v.Exporter2Name = types.StringNull()
	}

	return err
}

func (v *VrfFlowRulesValue) SetValue(jsonData *resource_fabric_common.NDFCVrfFlowRulesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.VrfFlowRuleName != "" {
		v.VrfFlowRuleName = types.StringValue(jsonData.VrfFlowRuleName)
	} else {
		v.VrfFlowRuleName = types.StringNull()
	}

	if jsonData.VrfFlowRuleUuid != "" {
		v.VrfFlowRuleUuid = types.StringValue(jsonData.VrfFlowRuleUuid)
	} else {
		v.VrfFlowRuleUuid = types.StringNull()
	}

	if jsonData.VrfFlowRuleTenant != "" {
		v.VrfFlowRuleTenant = types.StringValue(jsonData.VrfFlowRuleTenant)
	} else {
		v.VrfFlowRuleTenant = types.StringNull()
	}

	if jsonData.VrfFlowRuleVrf != "" {
		v.VrfFlowRuleVrf = types.StringValue(jsonData.VrfFlowRuleVrf)
	} else {
		v.VrfFlowRuleVrf = types.StringNull()
	}

	if len(jsonData.VrfFlowRuleSubnets) == 0 {
		log.Printf("v.VrfFlowRuleSubnets is empty")
		v.VrfFlowRuleSubnets = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.VrfFlowRuleSubnets))
		for i, item := range jsonData.VrfFlowRuleSubnets {
			listData[i] = types.StringValue(item)
		}
		v.VrfFlowRuleSubnets, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if len(jsonData.VrfFlowRuleAttributes) == 0 {
		log.Printf("v.VrfFlowRuleAttributes is empty")
		v.VrfFlowRuleAttributes = types.ListNull(VrfFlowRuleAttributesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.VrfFlowRuleAttributes contains %d elements", len(jsonData.VrfFlowRuleAttributes))
		listData := make([]VrfFlowRuleAttributesValue, 0)
		for _, item := range jsonData.VrfFlowRuleAttributes {
			data := new(VrfFlowRuleAttributesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in VrfFlowRuleAttributesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.VrfFlowRuleAttributes, err = types.ListValueFrom(context.Background(), VrfFlowRuleAttributesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []VrfFlowRuleAttributesValue to  List")
			return err
		}
	}

	return err
}

func (v *VrfFlowRuleAttributesValue) SetValue(jsonData *resource_fabric_common.NDFCVrfFlowRuleAttributesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.VrfFlowRuleBidirectional != nil {
		v.VrfFlowRuleBidirectional = types.BoolValue(*jsonData.VrfFlowRuleBidirectional)

	} else {
		v.VrfFlowRuleBidirectional = types.BoolNull()
	}

	if jsonData.VrfFlowRuleDstIp != "" {
		v.VrfFlowRuleDstIp = types.StringValue(jsonData.VrfFlowRuleDstIp)
	} else {
		v.VrfFlowRuleDstIp = types.StringNull()
	}

	if jsonData.VrfFlowRuleSrcIp != "" {
		v.VrfFlowRuleSrcIp = types.StringValue(jsonData.VrfFlowRuleSrcIp)
	} else {
		v.VrfFlowRuleSrcIp = types.StringNull()
	}

	if jsonData.VrfFlowRuleDstPort != "" {
		v.VrfFlowRuleDstPort = types.StringValue(jsonData.VrfFlowRuleDstPort)
	} else {
		v.VrfFlowRuleDstPort = types.StringNull()
	}

	if jsonData.VrfFlowRuleSrcPort != "" {
		v.VrfFlowRuleSrcPort = types.StringValue(jsonData.VrfFlowRuleSrcPort)
	} else {
		v.VrfFlowRuleSrcPort = types.StringNull()
	}

	if jsonData.VrfFlowRuleProtocol != "" {
		v.VrfFlowRuleProtocol = types.StringValue(jsonData.VrfFlowRuleProtocol)
	} else {
		v.VrfFlowRuleProtocol = types.StringNull()
	}

	if jsonData.VrfFlowRuleAttributeId != "" {
		v.VrfFlowRuleAttributeId = types.StringValue(jsonData.VrfFlowRuleAttributeId)
	} else {
		v.VrfFlowRuleAttributeId = types.StringNull()
	}

	return err
}

func (v *InterfaceFlowRulesValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceFlowRulesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceFlowRuleName != "" {
		v.InterfaceFlowRuleName = types.StringValue(jsonData.InterfaceFlowRuleName)
	} else {
		v.InterfaceFlowRuleName = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleUuid != "" {
		v.InterfaceFlowRuleUuid = types.StringValue(jsonData.InterfaceFlowRuleUuid)
	} else {
		v.InterfaceFlowRuleUuid = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleType != "" {
		v.InterfaceFlowRuleType = types.StringValue(jsonData.InterfaceFlowRuleType)
	} else {
		v.InterfaceFlowRuleType = types.StringNull()
	}

	if len(jsonData.InterfaceFlowRuleInterfaceCollection) == 0 {
		log.Printf("v.InterfaceFlowRuleInterfaceCollection is empty")
		v.InterfaceFlowRuleInterfaceCollection = types.ListNull(InterfaceFlowRuleInterfaceCollectionValue{}.Type(context.Background()))
	} else {
		log.Printf("v.InterfaceFlowRuleInterfaceCollection contains %d elements", len(jsonData.InterfaceFlowRuleInterfaceCollection))
		listData := make([]InterfaceFlowRuleInterfaceCollectionValue, 0)
		for _, item := range jsonData.InterfaceFlowRuleInterfaceCollection {
			data := new(InterfaceFlowRuleInterfaceCollectionValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfaceFlowRuleInterfaceCollectionValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.InterfaceFlowRuleInterfaceCollection, err = types.ListValueFrom(context.Background(), InterfaceFlowRuleInterfaceCollectionValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []InterfaceFlowRuleInterfaceCollectionValue to  List")
			return err
		}
	}

	if len(jsonData.InterfaceFlowRuleSubnets) == 0 {
		log.Printf("v.InterfaceFlowRuleSubnets is empty")
		v.InterfaceFlowRuleSubnets = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.InterfaceFlowRuleSubnets))
		for i, item := range jsonData.InterfaceFlowRuleSubnets {
			listData[i] = types.StringValue(item)
		}
		v.InterfaceFlowRuleSubnets, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}
	if len(jsonData.InterfaceFlowRuleAttributes) == 0 {
		log.Printf("v.InterfaceFlowRuleAttributes is empty")
		v.InterfaceFlowRuleAttributes = types.ListNull(InterfaceFlowRuleAttributesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.InterfaceFlowRuleAttributes contains %d elements", len(jsonData.InterfaceFlowRuleAttributes))
		listData := make([]InterfaceFlowRuleAttributesValue, 0)
		for _, item := range jsonData.InterfaceFlowRuleAttributes {
			data := new(InterfaceFlowRuleAttributesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfaceFlowRuleAttributesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.InterfaceFlowRuleAttributes, err = types.ListValueFrom(context.Background(), InterfaceFlowRuleAttributesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []InterfaceFlowRuleAttributesValue to  List")
			return err
		}
	}

	return err
}

func (v *InterfaceFlowRuleInterfaceCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceFlowRuleInterfaceCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceFlowRuleSwitchId != "" {
		v.InterfaceFlowRuleSwitchId = types.StringValue(jsonData.InterfaceFlowRuleSwitchId)
	} else {
		v.InterfaceFlowRuleSwitchId = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleSwitchName != "" {
		v.InterfaceFlowRuleSwitchName = types.StringValue(jsonData.InterfaceFlowRuleSwitchName)
	} else {
		v.InterfaceFlowRuleSwitchName = types.StringNull()
	}

	if len(jsonData.InterfaceFlowRuleInterfaces) == 0 {
		log.Printf("v.InterfaceFlowRuleInterfaces is empty")
		v.InterfaceFlowRuleInterfaces = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.InterfaceFlowRuleInterfaces))
		for i, item := range jsonData.InterfaceFlowRuleInterfaces {
			listData[i] = types.StringValue(item)
		}
		v.InterfaceFlowRuleInterfaces, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v *InterfaceFlowRuleAttributesValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceFlowRuleAttributesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceFlowRuleBidirectional != nil {
		v.InterfaceFlowRuleBidirectional = types.BoolValue(*jsonData.InterfaceFlowRuleBidirectional)

	} else {
		v.InterfaceFlowRuleBidirectional = types.BoolNull()
	}

	if jsonData.InterfaceFlowRuleDstIp != "" {
		v.InterfaceFlowRuleDstIp = types.StringValue(jsonData.InterfaceFlowRuleDstIp)
	} else {
		v.InterfaceFlowRuleDstIp = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleSrcIp != "" {
		v.InterfaceFlowRuleSrcIp = types.StringValue(jsonData.InterfaceFlowRuleSrcIp)
	} else {
		v.InterfaceFlowRuleSrcIp = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleDstPort != "" {
		v.InterfaceFlowRuleDstPort = types.StringValue(jsonData.InterfaceFlowRuleDstPort)
	} else {
		v.InterfaceFlowRuleDstPort = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleSrcPort != "" {
		v.InterfaceFlowRuleSrcPort = types.StringValue(jsonData.InterfaceFlowRuleSrcPort)
	} else {
		v.InterfaceFlowRuleSrcPort = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleProtocol != "" {
		v.InterfaceFlowRuleProtocol = types.StringValue(jsonData.InterfaceFlowRuleProtocol)
	} else {
		v.InterfaceFlowRuleProtocol = types.StringNull()
	}

	if jsonData.InterfaceFlowRuleAttributeId != "" {
		v.InterfaceFlowRuleAttributeId = types.StringValue(jsonData.InterfaceFlowRuleAttributeId)
	} else {
		v.InterfaceFlowRuleAttributeId = types.StringNull()
	}

	return err
}

func (v *L3OutFlowRulesValue) SetValue(jsonData *resource_fabric_common.NDFCL3OutFlowRulesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.L3OutFlowRuleName != "" {
		v.L3OutFlowRuleName = types.StringValue(jsonData.L3OutFlowRuleName)
	} else {
		v.L3OutFlowRuleName = types.StringNull()
	}

	if jsonData.L3OutFlowRuleUuid != "" {
		v.L3OutFlowRuleUuid = types.StringValue(jsonData.L3OutFlowRuleUuid)
	} else {
		v.L3OutFlowRuleUuid = types.StringNull()
	}

	if jsonData.L3OutFlowRuleType != "" {
		v.L3OutFlowRuleType = types.StringValue(jsonData.L3OutFlowRuleType)
	} else {
		v.L3OutFlowRuleType = types.StringNull()
	}

	if len(jsonData.L3OutFlowRuleInterfaceCollection) == 0 {
		log.Printf("v.L3OutFlowRuleInterfaceCollection is empty")
		v.L3OutFlowRuleInterfaceCollection = types.ListNull(L3OutFlowRuleInterfaceCollectionValue{}.Type(context.Background()))
	} else {
		log.Printf("v.L3OutFlowRuleInterfaceCollection contains %d elements", len(jsonData.L3OutFlowRuleInterfaceCollection))
		listData := make([]L3OutFlowRuleInterfaceCollectionValue, 0)
		for _, item := range jsonData.L3OutFlowRuleInterfaceCollection {
			data := new(L3OutFlowRuleInterfaceCollectionValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in L3OutFlowRuleInterfaceCollectionValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.L3OutFlowRuleInterfaceCollection, err = types.ListValueFrom(context.Background(), L3OutFlowRuleInterfaceCollectionValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []L3OutFlowRuleInterfaceCollectionValue to  List")
			return err
		}
	}

	if len(jsonData.L3OutFlowRuleSubnets) == 0 {
		log.Printf("v.L3OutFlowRuleSubnets is empty")
		v.L3OutFlowRuleSubnets = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.L3OutFlowRuleSubnets))
		for i, item := range jsonData.L3OutFlowRuleSubnets {
			listData[i] = types.StringValue(item)
		}
		v.L3OutFlowRuleSubnets, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v *L3OutFlowRuleInterfaceCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCL3OutFlowRuleInterfaceCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.L3OutFlowRuleTenant != "" {
		v.L3OutFlowRuleTenant = types.StringValue(jsonData.L3OutFlowRuleTenant)
	} else {
		v.L3OutFlowRuleTenant = types.StringNull()
	}

	if jsonData.L3OutFlowRuleL3Out != "" {
		v.L3OutFlowRuleL3Out = types.StringValue(jsonData.L3OutFlowRuleL3Out)
	} else {
		v.L3OutFlowRuleL3Out = types.StringNull()
	}

	if jsonData.L3OutFlowRuleEncap != "" {
		v.L3OutFlowRuleEncap = types.StringValue(jsonData.L3OutFlowRuleEncap)
	} else {
		v.L3OutFlowRuleEncap = types.StringNull()
	}

	if jsonData.L3OutFlowRuleSwitchName != "" {
		v.L3OutFlowRuleSwitchName = types.StringValue(jsonData.L3OutFlowRuleSwitchName)
	} else {
		v.L3OutFlowRuleSwitchName = types.StringNull()
	}

	if jsonData.L3OutFlowRuleSwitchId != "" {
		v.L3OutFlowRuleSwitchId = types.StringValue(jsonData.L3OutFlowRuleSwitchId)
	} else {
		v.L3OutFlowRuleSwitchId = types.StringNull()
	}

	if len(jsonData.L3OutFlowRuleInterfaces) == 0 {
		log.Printf("v.L3OutFlowRuleInterfaces is empty")
		v.L3OutFlowRuleInterfaces = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.L3OutFlowRuleInterfaces))
		for i, item := range jsonData.L3OutFlowRuleInterfaces {
			listData[i] = types.StringValue(item)
		}
		v.L3OutFlowRuleInterfaces, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v *InterfaceRulesValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceRulesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceRuleName != "" {
		v.InterfaceRuleName = types.StringValue(jsonData.InterfaceRuleName)
	} else {
		v.InterfaceRuleName = types.StringNull()
	}

	if len(jsonData.InterfaceRuleInterfaceCollection) == 0 {
		log.Printf("v.InterfaceRuleInterfaceCollection is empty")
		v.InterfaceRuleInterfaceCollection = types.ListNull(InterfaceRuleInterfaceCollectionValue{}.Type(context.Background()))
	} else {
		log.Printf("v.InterfaceRuleInterfaceCollection contains %d elements", len(jsonData.InterfaceRuleInterfaceCollection))
		listData := make([]InterfaceRuleInterfaceCollectionValue, 0)
		for _, item := range jsonData.InterfaceRuleInterfaceCollection {
			data := new(InterfaceRuleInterfaceCollectionValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfaceRuleInterfaceCollectionValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.InterfaceRuleInterfaceCollection, err = types.ListValueFrom(context.Background(), InterfaceRuleInterfaceCollectionValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []InterfaceRuleInterfaceCollectionValue to  List")
			return err
		}
	}

	if jsonData.InterfaceRuleEnabled != nil {
		v.InterfaceRuleEnabled = types.BoolValue(*jsonData.InterfaceRuleEnabled)

	} else {
		v.InterfaceRuleEnabled = types.BoolNull()
	}

	if jsonData.InterfaceRuleEnableFabricInterconnect != nil {
		v.InterfaceRuleEnableFabricInterconnect = types.BoolValue(*jsonData.InterfaceRuleEnableFabricInterconnect)

	} else {
		v.InterfaceRuleEnableFabricInterconnect = types.BoolNull()
	}

	if jsonData.InterfaceRuleEnableL3Out != nil {
		v.InterfaceRuleEnableL3Out = types.BoolValue(*jsonData.InterfaceRuleEnableL3Out)

	} else {
		v.InterfaceRuleEnableL3Out = types.BoolNull()
	}

	if jsonData.InterfaceRuleUuid != "" {
		v.InterfaceRuleUuid = types.StringValue(jsonData.InterfaceRuleUuid)
	} else {
		v.InterfaceRuleUuid = types.StringNull()
	}

	if len(jsonData.InterfaceRuleSubnets) == 0 {
		log.Printf("v.InterfaceRuleSubnets is empty")
		v.InterfaceRuleSubnets = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.InterfaceRuleSubnets))
		for i, item := range jsonData.InterfaceRuleSubnets {
			listData[i] = types.StringValue(item)
		}
		v.InterfaceRuleSubnets, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v *InterfaceRuleInterfaceCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceRuleInterfaceCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceRuleSwitchId != "" {
		v.InterfaceRuleSwitchId = types.StringValue(jsonData.InterfaceRuleSwitchId)
	} else {
		v.InterfaceRuleSwitchId = types.StringNull()
	}

	if jsonData.InterfaceRuleSwitchName != "" {
		v.InterfaceRuleSwitchName = types.StringValue(jsonData.InterfaceRuleSwitchName)
	} else {
		v.InterfaceRuleSwitchName = types.StringNull()
	}

	if jsonData.InterfaceRuleVrfName != "" {
		v.InterfaceRuleVrfName = types.StringValue(jsonData.InterfaceRuleVrfName)
	} else {
		v.InterfaceRuleVrfName = types.StringNull()
	}

	if len(jsonData.InterfaceRuleInterfaces) == 0 {
		log.Printf("v.InterfaceRuleInterfaces is empty")
		v.InterfaceRuleInterfaces = types.ListNull(InterfaceRuleInterfacesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.InterfaceRuleInterfaces contains %d elements", len(jsonData.InterfaceRuleInterfaces))
		listData := make([]InterfaceRuleInterfacesValue, 0)
		for _, item := range jsonData.InterfaceRuleInterfaces {
			data := new(InterfaceRuleInterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfaceRuleInterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.InterfaceRuleInterfaces, err = types.ListValueFrom(context.Background(), InterfaceRuleInterfacesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []InterfaceRuleInterfacesValue to  List")
			return err
		}
	}
	if jsonData.InterfaceRuleTenant != "" {
		v.InterfaceRuleTenant = types.StringValue(jsonData.InterfaceRuleTenant)
	} else {
		v.InterfaceRuleTenant = types.StringNull()
	}

	if jsonData.InterfaceRuleL3Out != "" {
		v.InterfaceRuleL3Out = types.StringValue(jsonData.InterfaceRuleL3Out)
	} else {
		v.InterfaceRuleL3Out = types.StringNull()
	}

	return err
}

func (v *InterfaceRuleInterfacesValue) SetValue(jsonData *resource_fabric_common.NDFCInterfaceRuleInterfacesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.InterfaceRuleInterfaceName != "" {
		v.InterfaceRuleInterfaceName = types.StringValue(jsonData.InterfaceRuleInterfaceName)
	} else {
		v.InterfaceRuleInterfaceName = types.StringNull()
	}

	if jsonData.InterfaceRuleInterfaceType != "" {
		v.InterfaceRuleInterfaceType = types.StringValue(jsonData.InterfaceRuleInterfaceType)
	} else {
		v.InterfaceRuleInterfaceType = types.StringNull()
	}

	if jsonData.InterfaceRuleInterfaceEncap != "" {
		v.InterfaceRuleInterfaceEncap = types.StringValue(jsonData.InterfaceRuleInterfaceEncap)
	} else {
		v.InterfaceRuleInterfaceEncap = types.StringNull()
	}

	if jsonData.InterfaceRuleInterfaceVrfName != "" {
		v.InterfaceRuleInterfaceVrfName = types.StringValue(jsonData.InterfaceRuleInterfaceVrfName)
	} else {
		v.InterfaceRuleInterfaceVrfName = types.StringNull()
	}

	return err
}

func (v *EmailValue) SetValue(jsonData *resource_fabric_common.NDFCEmailValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Name != "" {
		v.Name = types.StringValue(jsonData.Name)
	} else {
		v.Name = types.StringNull()
	}

	if jsonData.ReceiverEmail != "" {
		v.ReceiverEmail = types.StringValue(jsonData.ReceiverEmail)
	} else {
		v.ReceiverEmail = types.StringNull()
	}

	if jsonData.Format != "" {
		v.Format = types.StringValue(jsonData.Format)
	} else {
		v.Format = types.StringNull()
	}

	if jsonData.StartDate != "" {
		v.StartDate = types.StringValue(jsonData.StartDate)
	} else {
		v.StartDate = types.StringNull()
	}

	if jsonData.CollectionFrequencyInDays != nil {
		v.CollectionFrequencyInDays = types.Int64Value(*jsonData.CollectionFrequencyInDays)

	} else {
		v.CollectionFrequencyInDays = types.Int64Null()
	}

	if jsonData.CollectionSettings.CollectionType != "" {
		v.CollectionType = types.StringValue(jsonData.CollectionSettings.CollectionType)

	} else {
		v.CollectionType = types.StringNull()
	}

	if len(jsonData.CollectionSettings.Anomalies) == 0 {
		log.Printf("v.Anomalies is empty")
		v.Anomalies = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Anomalies))
		for i, item := range jsonData.CollectionSettings.Anomalies {
			listData[i] = types.StringValue(item)
		}
		v.Anomalies, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.Advisories) == 0 {
		log.Printf("v.Advisories is empty")
		v.Advisories = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Advisories))
		for i, item := range jsonData.CollectionSettings.Advisories {
			listData[i] = types.StringValue(item)
		}
		v.Advisories, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.RiskAndConformanceReports) == 0 {
		log.Printf("v.RiskAndConformanceReports is empty")
		v.RiskAndConformanceReports = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.RiskAndConformanceReports))
		for i, item := range jsonData.CollectionSettings.RiskAndConformanceReports {
			listData[i] = types.StringValue(item)
		}
		v.RiskAndConformanceReports, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.OnlyIncludeActiveAlerts != nil {
		v.OnlyIncludeActiveAlerts = types.BoolValue(*jsonData.OnlyIncludeActiveAlerts)

	} else {
		v.OnlyIncludeActiveAlerts = types.BoolNull()
	}

	return err
}

func (v *MessageBusValue) SetValue(jsonData *resource_fabric_common.NDFCMessageBusValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Server != "" {
		v.Server = types.StringValue(jsonData.Server)
	} else {
		v.Server = types.StringNull()
	}

	if jsonData.CollectionType != "" {
		v.CollectionType = types.StringValue(jsonData.CollectionType)
	} else {
		v.CollectionType = types.StringNull()
	}

	if jsonData.CollectionSettings.CollectionSettingsCollectionType != "" {
		v.CollectionSettingsCollectionType = types.StringValue(jsonData.CollectionSettings.CollectionSettingsCollectionType)

	} else {
		v.CollectionSettingsCollectionType = types.StringNull()
	}

	if len(jsonData.CollectionSettings.Anomalies) == 0 {
		log.Printf("v.Anomalies is empty")
		v.Anomalies = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Anomalies))
		for i, item := range jsonData.CollectionSettings.Anomalies {
			listData[i] = types.StringValue(item)
		}
		v.Anomalies, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.Advisories) == 0 {
		log.Printf("v.Advisories is empty")
		v.Advisories = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Advisories))
		for i, item := range jsonData.CollectionSettings.Advisories {
			listData[i] = types.StringValue(item)
		}
		v.Advisories, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.Statistics) == 0 {
		log.Printf("v.Statistics is empty")
		v.Statistics = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Statistics))
		for i, item := range jsonData.CollectionSettings.Statistics {
			listData[i] = types.StringValue(item)
		}
		v.Statistics, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.Faults) == 0 {
		log.Printf("v.Faults is empty")
		v.Faults = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.Faults))
		for i, item := range jsonData.CollectionSettings.Faults {
			listData[i] = types.StringValue(item)
		}
		v.Faults, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.CollectionSettings.AuditLogs) == 0 {
		log.Printf("v.AuditLogs is empty")
		v.AuditLogs = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.CollectionSettings.AuditLogs))
		for i, item := range jsonData.CollectionSettings.AuditLogs {
			listData[i] = types.StringValue(item)
		}
		v.AuditLogs, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v FabricExternalModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.LicenseTier.IsNull() && !v.LicenseTier.IsUnknown() {
		data.LicenseTier = v.LicenseTier.ValueString()
	} else {
		data.LicenseTier = ""
	}

	if !v.ControllerStatus.IsNull() && !v.ControllerStatus.IsUnknown() {
		data.FeatureStatus.ControllerStatus = v.ControllerStatus.ValueString()
	} else {
		data.FeatureStatus.ControllerStatus = ""
	}

	if !v.TelemetryStatus.IsNull() && !v.TelemetryStatus.IsUnknown() {
		data.FeatureStatus.TelemetryStatus = v.TelemetryStatus.ValueString()
	} else {
		data.FeatureStatus.TelemetryStatus = ""
	}

	if !v.OrchestrationStatus.IsNull() && !v.OrchestrationStatus.IsUnknown() {
		data.FeatureStatus.OrchestrationStatus = v.OrchestrationStatus.ValueString()
	} else {
		data.FeatureStatus.OrchestrationStatus = ""
	}

	if !v.TrapForwarderStatus.IsNull() && !v.TrapForwarderStatus.IsUnknown() {
		data.FeatureStatus.TrapForwarderStatus = v.TrapForwarderStatus.ValueString()
	} else {
		data.FeatureStatus.TrapForwarderStatus = ""
	}

	if !v.TelemetryCollection.IsNull() && !v.TelemetryCollection.IsUnknown() {
		data.TelemetryCollection = new(bool)
		*data.TelemetryCollection = v.TelemetryCollection.ValueBool()
	} else {
		data.TelemetryCollection = nil
	}

	if !v.TelemetryCollectionType.IsNull() && !v.TelemetryCollectionType.IsUnknown() {
		data.TelemetryCollectionType = v.TelemetryCollectionType.ValueString()
	} else {
		data.TelemetryCollectionType = ""
	}

	if !v.TelemetryStreamingProtocol.IsNull() && !v.TelemetryStreamingProtocol.IsUnknown() {
		data.TelemetryStreamingProtocol = v.TelemetryStreamingProtocol.ValueString()
	} else {
		data.TelemetryStreamingProtocol = ""
	}

	if !v.TelemetrySourceInterface.IsNull() && !v.TelemetrySourceInterface.IsUnknown() {
		data.TelemetrySourceInterface = v.TelemetrySourceInterface.ValueString()
	} else {
		data.TelemetrySourceInterface = ""
	}

	if !v.TelemetrySourceVrf.IsNull() && !v.TelemetrySourceVrf.IsUnknown() {
		data.TelemetrySourceVrf = v.TelemetrySourceVrf.ValueString()
	} else {
		data.TelemetrySourceVrf = ""
	}

	if !v.SecurityDomain.IsNull() && !v.SecurityDomain.IsUnknown() {
		data.SecurityDomain = v.SecurityDomain.ValueString()
	} else {
		data.SecurityDomain = ""
	}

	if !v.AllowedActions.IsNull() && !v.AllowedActions.IsUnknown() {
		listStringData := make([]string, len(v.AllowedActions.Elements()))
		dg := v.AllowedActions.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Meta.AllowedActions = make([]string, len(listStringData))

		copy(data.Meta.AllowedActions, listStringData)
	}

	if !v.BgpAsn.IsNull() && !v.BgpAsn.IsUnknown() {
		data.Management.BgpAsn = v.BgpAsn.ValueString()
	} else {
		data.Management.BgpAsn = ""
	}

	if !v.Category.IsNull() && !v.Category.IsUnknown() {
		data.Category = v.Category.ValueString()
	} else {
		data.Category = ""
	}

	if !v.AlertSuspend.IsNull() && !v.AlertSuspend.IsUnknown() {
		data.AlertSuspend = v.AlertSuspend.ValueString()
	} else {
		data.AlertSuspend = ""
	}

	if !v.CreateBgpConfig.IsNull() && !v.CreateBgpConfig.IsUnknown() {
		data.Management.CreateBgpConfig = new(bool)
		*data.Management.CreateBgpConfig = v.CreateBgpConfig.ValueBool()
	} else {
		data.Management.CreateBgpConfig = nil
	}

	if !v.Aaa.IsNull() && !v.Aaa.IsUnknown() {
		data.Management.Aaa = new(bool)
		*data.Management.Aaa = v.Aaa.ValueBool()
	} else {
		data.Management.Aaa = nil
	}

	if !v.AdvancedSshOption.IsNull() && !v.AdvancedSshOption.IsUnknown() {
		data.Management.AdvancedSshOption = new(bool)
		*data.Management.AdvancedSshOption = v.AdvancedSshOption.ValueBool()
	} else {
		data.Management.AdvancedSshOption = nil
	}

	if !v.AllowSameLoopbackIpOnSwitches.IsNull() && !v.AllowSameLoopbackIpOnSwitches.IsUnknown() {
		data.Management.AllowSameLoopbackIpOnSwitches = new(bool)
		*data.Management.AllowSameLoopbackIpOnSwitches = v.AllowSameLoopbackIpOnSwitches.ValueBool()
	} else {
		data.Management.AllowSameLoopbackIpOnSwitches = nil
	}

	if !v.AllowSmartSwitchOnboarding.IsNull() && !v.AllowSmartSwitchOnboarding.IsUnknown() {
		data.Management.AllowSmartSwitchOnboarding = new(bool)
		*data.Management.AllowSmartSwitchOnboarding = v.AllowSmartSwitchOnboarding.ValueBool()
	} else {
		data.Management.AllowSmartSwitchOnboarding = nil
	}

	if !v.ConnectivityDomainName.IsNull() && !v.ConnectivityDomainName.IsUnknown() {
		data.Management.ConnectivityDomainName = v.ConnectivityDomainName.ValueString()
	} else {
		data.Management.ConnectivityDomainName = ""
	}

	if !v.HypershieldConnectivityProxyServer.IsNull() && !v.HypershieldConnectivityProxyServer.IsUnknown() {
		data.Management.HypershieldConnectivityProxyServer = v.HypershieldConnectivityProxyServer.ValueString()
	} else {
		data.Management.HypershieldConnectivityProxyServer = ""
	}

	if !v.HypershieldConnectivityProxyServerPort.IsNull() && !v.HypershieldConnectivityProxyServerPort.IsUnknown() {
		data.Management.HypershieldConnectivityProxyServerPort = new(int64)
		*data.Management.HypershieldConnectivityProxyServerPort = v.HypershieldConnectivityProxyServerPort.ValueInt64()

	} else {
		data.Management.HypershieldConnectivityProxyServerPort = nil
	}

	if !v.HypershieldConnectivitySourceIntf.IsNull() && !v.HypershieldConnectivitySourceIntf.IsUnknown() {
		data.Management.HypershieldConnectivitySourceIntf = v.HypershieldConnectivitySourceIntf.ValueString()
	} else {
		data.Management.HypershieldConnectivitySourceIntf = ""
	}

	if !v.Day0Bootstrap.IsNull() && !v.Day0Bootstrap.IsUnknown() {
		data.Management.Day0Bootstrap = new(bool)
		*data.Management.Day0Bootstrap = v.Day0Bootstrap.ValueBool()
	} else {
		data.Management.Day0Bootstrap = nil
	}

	if !v.Day0PlugAndPlay.IsNull() && !v.Day0PlugAndPlay.IsUnknown() {
		data.Management.Day0PlugAndPlay = new(bool)
		*data.Management.Day0PlugAndPlay = v.Day0PlugAndPlay.ValueBool()
	} else {
		data.Management.Day0PlugAndPlay = nil
	}

	if !v.InbandDay0Bootstrap.IsNull() && !v.InbandDay0Bootstrap.IsUnknown() {
		data.Management.InbandDay0Bootstrap = new(bool)
		*data.Management.InbandDay0Bootstrap = v.InbandDay0Bootstrap.ValueBool()
	} else {
		data.Management.InbandDay0Bootstrap = nil
	}

	if !v.Cdp.IsNull() && !v.Cdp.IsUnknown() {
		data.Management.Cdp = new(bool)
		*data.Management.Cdp = v.Cdp.ValueBool()
	} else {
		data.Management.Cdp = nil
	}

	if !v.CoppPolicy.IsNull() && !v.CoppPolicy.IsUnknown() {
		data.Management.CoppPolicy = v.CoppPolicy.ValueString()
	} else {
		data.Management.CoppPolicy = ""
	}

	if !v.DhcpEndAddress.IsNull() && !v.DhcpEndAddress.IsUnknown() {
		data.Management.DhcpEndAddress = v.DhcpEndAddress.ValueString()
	} else {
		data.Management.DhcpEndAddress = ""
	}

	if !v.DhcpProtocolVersion.IsNull() && !v.DhcpProtocolVersion.IsUnknown() {
		data.Management.DhcpProtocolVersion = v.DhcpProtocolVersion.ValueString()
	} else {
		data.Management.DhcpProtocolVersion = ""
	}

	if !v.DhcpStartAddress.IsNull() && !v.DhcpStartAddress.IsUnknown() {
		data.Management.DhcpStartAddress = v.DhcpStartAddress.ValueString()
	} else {
		data.Management.DhcpStartAddress = ""
	}

	if !v.LocalDhcpServer.IsNull() && !v.LocalDhcpServer.IsUnknown() {
		data.Management.LocalDhcpServer = new(bool)
		*data.Management.LocalDhcpServer = v.LocalDhcpServer.ValueBool()
	} else {
		data.Management.LocalDhcpServer = nil
	}

	if !v.DnsCollection.IsNull() && !v.DnsCollection.IsUnknown() {
		listStringData := make([]string, len(v.DnsCollection.Elements()))
		dg := v.DnsCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.DnsCollection = make([]string, len(listStringData))

		copy(data.Management.DnsCollection, listStringData)
	}

	if !v.DnsVrfCollection.IsNull() && !v.DnsVrfCollection.IsUnknown() {
		listStringData := make([]string, len(v.DnsVrfCollection.Elements()))
		dg := v.DnsVrfCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.DnsVrfCollection = make([]string, len(listStringData))

		copy(data.Management.DnsVrfCollection, listStringData)
	}

	if !v.DomainName.IsNull() && !v.DomainName.IsUnknown() {
		data.Management.DomainName = v.DomainName.ValueString()
	} else {
		data.Management.DomainName = ""
	}

	if !v.EnableDpuPinning.IsNull() && !v.EnableDpuPinning.IsUnknown() {
		data.Management.EnableDpuPinning = new(bool)
		*data.Management.EnableDpuPinning = v.EnableDpuPinning.ValueBool()
	} else {
		data.Management.EnableDpuPinning = nil
	}

	if !v.ExtraConfigAaa.IsNull() && !v.ExtraConfigAaa.IsUnknown() {
		data.Management.ExtraConfigAaa = v.ExtraConfigAaa.ValueString()
	} else {
		data.Management.ExtraConfigAaa = ""
	}

	if !v.ExtraConfigFabric.IsNull() && !v.ExtraConfigFabric.IsUnknown() {
		data.Management.ExtraConfigFabric = v.ExtraConfigFabric.ValueString()
	} else {
		data.Management.ExtraConfigFabric = ""
	}

	if !v.ExtraConfigNxosBootstrap.IsNull() && !v.ExtraConfigNxosBootstrap.IsUnknown() {
		data.Management.ExtraConfigNxosBootstrap = v.ExtraConfigNxosBootstrap.ValueString()
	} else {
		data.Management.ExtraConfigNxosBootstrap = ""
	}

	if !v.ExtraConfigXeBootstrap.IsNull() && !v.ExtraConfigXeBootstrap.IsUnknown() {
		data.Management.ExtraConfigXeBootstrap = v.ExtraConfigXeBootstrap.ValueString()
	} else {
		data.Management.ExtraConfigXeBootstrap = ""
	}

	if !v.InbandManagement.IsNull() && !v.InbandManagement.IsUnknown() {
		data.Management.InbandManagement = new(bool)
		*data.Management.InbandManagement = v.InbandManagement.ValueBool()
	} else {
		data.Management.InbandManagement = nil
	}

	if !v.RealTimeInterfaceStatisticsCollection.IsNull() && !v.RealTimeInterfaceStatisticsCollection.IsUnknown() {
		data.Management.RealTimeInterfaceStatisticsCollection = new(bool)
		*data.Management.RealTimeInterfaceStatisticsCollection = v.RealTimeInterfaceStatisticsCollection.ValueBool()
	} else {
		data.Management.RealTimeInterfaceStatisticsCollection = nil
	}

	if !v.InterfaceStatisticsLoadInterval.IsNull() && !v.InterfaceStatisticsLoadInterval.IsUnknown() {
		data.Management.InterfaceStatisticsLoadInterval = new(int64)
		*data.Management.InterfaceStatisticsLoadInterval = v.InterfaceStatisticsLoadInterval.ValueInt64()

	} else {
		data.Management.InterfaceStatisticsLoadInterval = nil
	}

	if !v.ManagementGateway.IsNull() && !v.ManagementGateway.IsUnknown() {
		data.Management.ManagementGateway = v.ManagementGateway.ValueString()
	} else {
		data.Management.ManagementGateway = ""
	}

	if !v.ManagementIpv4Prefix.IsNull() && !v.ManagementIpv4Prefix.IsUnknown() {
		data.Management.ManagementIpv4Prefix = new(int64)
		*data.Management.ManagementIpv4Prefix = v.ManagementIpv4Prefix.ValueInt64()

	} else {
		data.Management.ManagementIpv4Prefix = nil
	}

	if !v.ManagementIpv6Prefix.IsNull() && !v.ManagementIpv6Prefix.IsUnknown() {
		data.Management.ManagementIpv6Prefix = new(int64)
		*data.Management.ManagementIpv6Prefix = v.ManagementIpv6Prefix.ValueInt64()

	} else {
		data.Management.ManagementIpv6Prefix = nil
	}

	if !v.MonitoredMode.IsNull() && !v.MonitoredMode.IsUnknown() {
		data.Management.MonitoredMode = new(bool)
		*data.Management.MonitoredMode = v.MonitoredMode.ValueBool()
	} else {
		data.Management.MonitoredMode = nil
	}

	if !v.MplsHandoff.IsNull() && !v.MplsHandoff.IsUnknown() {
		data.Management.MplsHandoff = new(bool)
		*data.Management.MplsHandoff = v.MplsHandoff.ValueBool()
	} else {
		data.Management.MplsHandoff = nil
	}

	if !v.MplsLoopbackIdentifier.IsNull() && !v.MplsLoopbackIdentifier.IsUnknown() {
		data.Management.MplsLoopbackIdentifier = new(int64)
		*data.Management.MplsLoopbackIdentifier = v.MplsLoopbackIdentifier.ValueInt64()

	} else {
		data.Management.MplsLoopbackIdentifier = nil
	}

	if !v.MplsLoopbackIpRange.IsNull() && !v.MplsLoopbackIpRange.IsUnknown() {
		data.Management.MplsLoopbackIpRange = v.MplsLoopbackIpRange.ValueString()
	} else {
		data.Management.MplsLoopbackIpRange = ""
	}

	if !v.NetflowEnable.IsNull() && !v.NetflowEnable.IsUnknown() {
		data.Management.NetflowSettings.NetflowEnable = new(bool)
		*data.Management.NetflowSettings.NetflowEnable = v.NetflowEnable.ValueBool()
	} else {
		data.Management.NetflowSettings.NetflowEnable = nil
	}

	if !v.Nxapi.IsNull() && !v.Nxapi.IsUnknown() {
		data.Management.Nxapi = new(bool)
		*data.Management.Nxapi = v.Nxapi.ValueBool()
	} else {
		data.Management.Nxapi = nil
	}

	if !v.NxapiHttpsPort.IsNull() && !v.NxapiHttpsPort.IsUnknown() {
		data.Management.NxapiHttpsPort = new(int64)
		*data.Management.NxapiHttpsPort = v.NxapiHttpsPort.ValueInt64()

	} else {
		data.Management.NxapiHttpsPort = nil
	}

	if !v.NxapiHttp.IsNull() && !v.NxapiHttp.IsUnknown() {
		data.Management.NxapiHttp = new(bool)
		*data.Management.NxapiHttp = v.NxapiHttp.ValueBool()
	} else {
		data.Management.NxapiHttp = nil
	}

	if !v.NxapiHttpPort.IsNull() && !v.NxapiHttpPort.IsUnknown() {
		data.Management.NxapiHttpPort = new(int64)
		*data.Management.NxapiHttpPort = v.NxapiHttpPort.ValueInt64()

	} else {
		data.Management.NxapiHttpPort = nil
	}

	if !v.PerformanceMonitoring.IsNull() && !v.PerformanceMonitoring.IsUnknown() {
		data.Management.PerformanceMonitoring = new(bool)
		*data.Management.PerformanceMonitoring = v.PerformanceMonitoring.ValueBool()
	} else {
		data.Management.PerformanceMonitoring = nil
	}

	if !v.PowerRedundancyMode.IsNull() && !v.PowerRedundancyMode.IsUnknown() {
		data.Management.PowerRedundancyMode = v.PowerRedundancyMode.ValueString()
	} else {
		data.Management.PowerRedundancyMode = ""
	}

	if !v.Ptp.IsNull() && !v.Ptp.IsUnknown() {
		data.Management.Ptp = new(bool)
		*data.Management.Ptp = v.Ptp.ValueBool()
	} else {
		data.Management.Ptp = nil
	}

	if !v.PtpLoopbackId.IsNull() && !v.PtpLoopbackId.IsUnknown() {
		data.Management.PtpLoopbackId = new(int64)
		*data.Management.PtpLoopbackId = v.PtpLoopbackId.ValueInt64()

	} else {
		data.Management.PtpLoopbackId = nil
	}

	if !v.PtpDomainId.IsNull() && !v.PtpDomainId.IsUnknown() {
		data.Management.PtpDomainId = new(int64)
		*data.Management.PtpDomainId = v.PtpDomainId.ValueInt64()

	} else {
		data.Management.PtpDomainId = nil
	}

	if !v.RealTimeBackup.IsNull() && !v.RealTimeBackup.IsUnknown() {
		data.Management.RealTimeBackup = new(bool)
		*data.Management.RealTimeBackup = v.RealTimeBackup.ValueBool()
	} else {
		data.Management.RealTimeBackup = nil
	}

	if !v.ScheduledBackup.IsNull() && !v.ScheduledBackup.IsUnknown() {
		data.Management.ScheduledBackup = new(bool)
		*data.Management.ScheduledBackup = v.ScheduledBackup.ValueBool()
	} else {
		data.Management.ScheduledBackup = nil
	}

	if !v.ScheduledBackupTime.IsNull() && !v.ScheduledBackupTime.IsUnknown() {
		data.Management.ScheduledBackupTime = v.ScheduledBackupTime.ValueString()
	} else {
		data.Management.ScheduledBackupTime = ""
	}

	if !v.SnmpTrap.IsNull() && !v.SnmpTrap.IsUnknown() {
		data.Management.SnmpTrap = new(bool)
		*data.Management.SnmpTrap = v.SnmpTrap.ValueBool()
	} else {
		data.Management.SnmpTrap = nil
	}

	if !v.SubInterfaceDot1qRange.IsNull() && !v.SubInterfaceDot1qRange.IsUnknown() {
		data.Management.SubInterfaceDot1qRange = v.SubInterfaceDot1qRange.ValueString()
	} else {
		data.Management.SubInterfaceDot1qRange = ""
	}

	if !v.TrafficAnalytics.IsNull() && !v.TrafficAnalytics.IsUnknown() {
		data.TelemetrySettings.FlowCollection.TrafficAnalytics = v.TrafficAnalytics.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.TrafficAnalytics = ""
	}

	if !v.NetFlow.IsNull() && !v.NetFlow.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow = new(bool)
		*data.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow = v.NetFlow.ValueBool()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.NetFlow = nil
	}

	if !v.SFlow.IsNull() && !v.SFlow.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow = new(bool)
		*data.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow = v.SFlow.ValueBool()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.SFlow = nil
	}

	if !v.FlowTelemetry.IsNull() && !v.FlowTelemetry.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry = new(bool)
		*data.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry = v.FlowTelemetry.ValueBool()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionModes.FlowTelemetry = nil
	}

	if !v.TrafficAnalyticsRulesEnabled.IsNull() && !v.TrafficAnalyticsRulesEnabled.IsUnknown() {
		data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled = new(bool)
		*data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled = v.TrafficAnalyticsRulesEnabled.ValueBool()
	} else {
		data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.TrafficAnalyticsRulesEnabled = nil
	}

	if !v.TrafficAnalyticsMode.IsNull() && !v.TrafficAnalyticsMode.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode = v.TrafficAnalyticsMode.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsMode = ""
	}

	if !v.UdpCategorization.IsNull() && !v.UdpCategorization.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization = v.UdpCategorization.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.UdpCategorization = ""
	}

	if !v.TrafficAnalyticsFilterRules.IsNull() && !v.TrafficAnalyticsFilterRules.IsUnknown() {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules = v.TrafficAnalyticsFilterRules.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.FlowCollectionCapabilities.TrafficAnalyticsFilterRules = ""
	}

	if !v.OperatingMode.IsNull() && !v.OperatingMode.IsUnknown() {
		data.TelemetrySettings.FlowCollection.OperatingMode = v.OperatingMode.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.OperatingMode = ""
	}

	if !v.UdpCategorizationSupport.IsNull() && !v.UdpCategorizationSupport.IsUnknown() {
		data.TelemetrySettings.FlowCollection.UdpCategorizationSupport = v.UdpCategorizationSupport.ValueString()
	} else {
		data.TelemetrySettings.FlowCollection.UdpCategorizationSupport = ""
	}

	if !v.Microburst.IsNull() && !v.Microburst.IsUnknown() {
		data.TelemetrySettings.Microburst.Microburst = new(bool)
		*data.TelemetrySettings.Microburst.Microburst = v.Microburst.ValueBool()
	} else {
		data.TelemetrySettings.Microburst.Microburst = nil
	}

	if !v.Sensitivity.IsNull() && !v.Sensitivity.IsUnknown() {
		data.TelemetrySettings.Microburst.Sensitivity = v.Sensitivity.ValueString()
	} else {
		data.TelemetrySettings.Microburst.Sensitivity = ""
	}

	if !v.AnalysisSettingsIsEnabled.IsNull() && !v.AnalysisSettingsIsEnabled.IsUnknown() {
		data.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled = new(bool)
		*data.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled = v.AnalysisSettingsIsEnabled.ValueBool()
	} else {
		data.TelemetrySettings.AnalysisSettings.AnalysisSettingsIsEnabled = nil
	}

	if !v.Server.IsNull() && !v.Server.IsUnknown() {
		data.TelemetrySettings.Nas.Server = v.Server.ValueString()
	} else {
		data.TelemetrySettings.Nas.Server = ""
	}

	if !v.ExportType.IsNull() && !v.ExportType.IsUnknown() {
		data.TelemetrySettings.Nas.ExportSettings.ExportType = v.ExportType.ValueString()
	} else {
		data.TelemetrySettings.Nas.ExportSettings.ExportType = ""
	}

	if !v.ExportFormat.IsNull() && !v.ExportFormat.IsUnknown() {
		data.TelemetrySettings.Nas.ExportSettings.ExportFormat = v.ExportFormat.ValueString()
	} else {
		data.TelemetrySettings.Nas.ExportSettings.ExportFormat = ""
	}

	if !v.Cost.IsNull() && !v.Cost.IsUnknown() {
		data.TelemetrySettings.EnergyManagement.Cost = new(float64)
		*data.TelemetrySettings.EnergyManagement.Cost = v.Cost.ValueFloat64()
	} else {
		data.TelemetrySettings.EnergyManagement.Cost = nil
	}

	if !v.SyslogServers.IsNull() && !v.SyslogServers.IsUnknown() {
		listStringData := make([]string, len(v.SyslogServers.Elements()))
		dg := v.SyslogServers.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.ExternalStreamingSettings.Syslog.SyslogServers = make([]string, len(listStringData))

		copy(data.ExternalStreamingSettings.Syslog.SyslogServers, listStringData)
	}

	if !v.SyslogFacility.IsNull() && !v.SyslogFacility.IsUnknown() {
		data.ExternalStreamingSettings.Syslog.SyslogFacility = v.SyslogFacility.ValueString()
	} else {
		data.ExternalStreamingSettings.Syslog.SyslogFacility = ""
	}

	if !v.SyslogAnomalies.IsNull() && !v.SyslogAnomalies.IsUnknown() {
		listStringData := make([]string, len(v.SyslogAnomalies.Elements()))
		dg := v.SyslogAnomalies.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.ExternalStreamingSettings.Syslog.CollectionSettings.SyslogAnomalies = make([]string, len(listStringData))

		copy(data.ExternalStreamingSettings.Syslog.CollectionSettings.SyslogAnomalies, listStringData)
	}

	//MARSHAL_BODY

	// Nested types Location # latitude
	if !v.Location.Latitude.IsNull() && !v.Location.Latitude.IsUnknown() {
		data.Location.Latitude = new(float64)
		*data.Location.Latitude = v.Location.Latitude.ValueFloat64()
	} else {
		data.Location.Latitude = nil
	}

	// Nested types Location # longitude
	if !v.Location.Longitude.IsNull() && !v.Location.Longitude.IsUnknown() {
		data.Location.Longitude = new(float64)
		*data.Location.Longitude = v.Location.Longitude.ValueFloat64()
	} else {
		data.Location.Longitude = nil
	}

	//MARSHALL_LIST  BEGIN BootstrapSubnetCollection[i1]

	if !v.BootstrapSubnetCollection.IsNull() && !v.BootstrapSubnetCollection.IsUnknown() {
		elements := make([]BootstrapSubnetCollectionValue, len(v.BootstrapSubnetCollection.Elements()))
		// Not augmenting

		data.Management.BootstrapSubnetCollection = make([]resource_fabric_common.NDFCBootstrapSubnetCollectionValue, len(v.BootstrapSubnetCollection.Elements()))

		// -- Set here 1 |.Management.BootstrapSubnetCollection[i1]|BootstrapSubnetCollection[i1]| --

		diag := v.BootstrapSubnetCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.BootstrapSubnetCollection[i1] BootstrapSubnetCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.StartIp.IsNull() && !ele1.StartIp.IsUnknown() {

				data.Management.BootstrapSubnetCollection[i1].StartIp = ele1.StartIp.ValueString()
			} else {
				data.Management.BootstrapSubnetCollection[i1].StartIp = ""
			}

			if !ele1.EndIp.IsNull() && !ele1.EndIp.IsUnknown() {

				data.Management.BootstrapSubnetCollection[i1].EndIp = ele1.EndIp.ValueString()
			} else {
				data.Management.BootstrapSubnetCollection[i1].EndIp = ""
			}

			if !ele1.DefaultGateway.IsNull() && !ele1.DefaultGateway.IsUnknown() {

				data.Management.BootstrapSubnetCollection[i1].DefaultGateway = ele1.DefaultGateway.ValueString()
			} else {
				data.Management.BootstrapSubnetCollection[i1].DefaultGateway = ""
			}

			if !ele1.SubnetPrefix.IsNull() && !ele1.SubnetPrefix.IsUnknown() {

				data.Management.BootstrapSubnetCollection[i1].SubnetPrefix = new(int64)
				*data.Management.BootstrapSubnetCollection[i1].SubnetPrefix = ele1.SubnetPrefix.ValueInt64()

			} else {
				data.Management.BootstrapSubnetCollection[i1].SubnetPrefix = nil
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN NetflowExporterCollection[i1]

	if !v.NetflowExporterCollection.IsNull() && !v.NetflowExporterCollection.IsUnknown() {
		elements := make([]NetflowExporterCollectionValue, len(v.NetflowExporterCollection.Elements()))
		// Not augmenting

		data.Management.NetflowSettings.NetflowExporterCollection = make([]resource_fabric_common.NDFCNetflowExporterCollectionValue, len(v.NetflowExporterCollection.Elements()))

		// -- Set here 1 |.Management.NetflowSettings.NetflowExporterCollection[i1]|NetflowExporterCollection[i1]| --

		diag := v.NetflowExporterCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.NetflowSettings.NetflowExporterCollection[i1] NetflowExporterCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.ExporterName.IsNull() && !ele1.ExporterName.IsUnknown() {

				data.Management.NetflowSettings.NetflowExporterCollection[i1].ExporterName = ele1.ExporterName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowExporterCollection[i1].ExporterName = ""
			}

			if !ele1.ExporterIp.IsNull() && !ele1.ExporterIp.IsUnknown() {

				data.Management.NetflowSettings.NetflowExporterCollection[i1].ExporterIp = ele1.ExporterIp.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowExporterCollection[i1].ExporterIp = ""
			}

			if !ele1.Vrf.IsNull() && !ele1.Vrf.IsUnknown() {

				data.Management.NetflowSettings.NetflowExporterCollection[i1].Vrf = ele1.Vrf.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowExporterCollection[i1].Vrf = ""
			}

			if !ele1.SourceInterfaceName.IsNull() && !ele1.SourceInterfaceName.IsUnknown() {

				data.Management.NetflowSettings.NetflowExporterCollection[i1].SourceInterfaceName = ele1.SourceInterfaceName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowExporterCollection[i1].SourceInterfaceName = ""
			}

			if !ele1.UdpPort.IsNull() && !ele1.UdpPort.IsUnknown() {

				data.Management.NetflowSettings.NetflowExporterCollection[i1].UdpPort = new(int64)
				*data.Management.NetflowSettings.NetflowExporterCollection[i1].UdpPort = ele1.UdpPort.ValueInt64()

			} else {
				data.Management.NetflowSettings.NetflowExporterCollection[i1].UdpPort = nil
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN NetflowRecordCollection[i1]

	if !v.NetflowRecordCollection.IsNull() && !v.NetflowRecordCollection.IsUnknown() {
		elements := make([]NetflowRecordCollectionValue, len(v.NetflowRecordCollection.Elements()))
		// Not augmenting

		data.Management.NetflowSettings.NetflowRecordCollection = make([]resource_fabric_common.NDFCNetflowRecordCollectionValue, len(v.NetflowRecordCollection.Elements()))

		// -- Set here 1 |.Management.NetflowSettings.NetflowRecordCollection[i1]|NetflowRecordCollection[i1]| --

		diag := v.NetflowRecordCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.NetflowSettings.NetflowRecordCollection[i1] NetflowRecordCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.RecordName.IsNull() && !ele1.RecordName.IsUnknown() {

				data.Management.NetflowSettings.NetflowRecordCollection[i1].RecordName = ele1.RecordName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowRecordCollection[i1].RecordName = ""
			}

			if !ele1.RecordTemplate.IsNull() && !ele1.RecordTemplate.IsUnknown() {

				data.Management.NetflowSettings.NetflowRecordCollection[i1].RecordTemplate = ele1.RecordTemplate.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowRecordCollection[i1].RecordTemplate = ""
			}

			if !ele1.Layer2Record.IsNull() && !ele1.Layer2Record.IsUnknown() {

				data.Management.NetflowSettings.NetflowRecordCollection[i1].Layer2Record = strconv.FormatBool(ele1.Layer2Record.ValueBool())
			} else {
				data.Management.NetflowSettings.NetflowRecordCollection[i1].Layer2Record = ""
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN NetflowMonitorCollection[i1]

	if !v.NetflowMonitorCollection.IsNull() && !v.NetflowMonitorCollection.IsUnknown() {
		elements := make([]NetflowMonitorCollectionValue, len(v.NetflowMonitorCollection.Elements()))
		// Not augmenting

		data.Management.NetflowSettings.NetflowMonitorCollection = make([]resource_fabric_common.NDFCNetflowMonitorCollectionValue, len(v.NetflowMonitorCollection.Elements()))

		// -- Set here 1 |.Management.NetflowSettings.NetflowMonitorCollection[i1]|NetflowMonitorCollection[i1]| --

		diag := v.NetflowMonitorCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.NetflowSettings.NetflowMonitorCollection[i1] NetflowMonitorCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.MonitorName.IsNull() && !ele1.MonitorName.IsUnknown() {

				data.Management.NetflowSettings.NetflowMonitorCollection[i1].MonitorName = ele1.MonitorName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowMonitorCollection[i1].MonitorName = ""
			}

			if !ele1.MonitorRecordName.IsNull() && !ele1.MonitorRecordName.IsUnknown() {

				data.Management.NetflowSettings.NetflowMonitorCollection[i1].MonitorRecordName = ele1.MonitorRecordName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowMonitorCollection[i1].MonitorRecordName = ""
			}

			if !ele1.Exporter1Name.IsNull() && !ele1.Exporter1Name.IsUnknown() {

				data.Management.NetflowSettings.NetflowMonitorCollection[i1].Exporter1Name = ele1.Exporter1Name.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowMonitorCollection[i1].Exporter1Name = ""
			}

			if !ele1.Exporter2Name.IsNull() && !ele1.Exporter2Name.IsUnknown() {

				data.Management.NetflowSettings.NetflowMonitorCollection[i1].Exporter2Name = ele1.Exporter2Name.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowMonitorCollection[i1].Exporter2Name = ""
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN VrfFlowRules[i1]

	if !v.VrfFlowRules.IsNull() && !v.VrfFlowRules.IsUnknown() {
		elements := make([]VrfFlowRulesValue, len(v.VrfFlowRules.Elements()))
		// Not augmenting

		data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules = make([]resource_fabric_common.NDFCVrfFlowRulesValue, len(v.VrfFlowRules.Elements()))

		// -- Set here 1 |.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1]|VrfFlowRules[i1]| --

		diag := v.VrfFlowRules.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1] VrfFlowRules[i1]
		for i1, ele1 := range elements {
			if !ele1.VrfFlowRuleName.IsNull() && !ele1.VrfFlowRuleName.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleName = ele1.VrfFlowRuleName.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleName = ""
			}

			if !ele1.VrfFlowRuleUuid.IsNull() && !ele1.VrfFlowRuleUuid.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleUuid = ele1.VrfFlowRuleUuid.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleUuid = ""
			}

			if !ele1.VrfFlowRuleTenant.IsNull() && !ele1.VrfFlowRuleTenant.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleTenant = ele1.VrfFlowRuleTenant.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleTenant = ""
			}

			if !ele1.VrfFlowRuleVrf.IsNull() && !ele1.VrfFlowRuleVrf.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleVrf = ele1.VrfFlowRuleVrf.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleVrf = ""
			}

			if !ele1.VrfFlowRuleSubnets.IsNull() && !ele1.VrfFlowRuleSubnets.IsUnknown() {

				// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1]|[]|vrf_flow_rule_subnets|
				listStringData := make([]string, len(ele1.VrfFlowRuleSubnets.Elements()))
				dg := ele1.VrfFlowRuleSubnets.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleSubnets = make([]string, len(listStringData))
				copy(data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleSubnets, listStringData)
			}

			// here 507 |.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1]|TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1]|

			//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2]

			if !ele1.VrfFlowRuleAttributes.IsNull() && !ele1.VrfFlowRuleAttributes.IsUnknown() {
				elements := make([]VrfFlowRuleAttributesValue, len(ele1.VrfFlowRuleAttributes.Elements()))
				// Not augmenting

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes|TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes = make([]resource_fabric_common.NDFCVrfFlowRuleAttributesValue, len(ele1.VrfFlowRuleAttributes.Elements()))

				diag := ele1.VrfFlowRuleAttributes.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				// .TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2] TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2]
				for i2, ele2 := range elements {
					if !ele2.VrfFlowRuleBidirectional.IsNull() && !ele2.VrfFlowRuleBidirectional.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleBidirectional = new(bool)
						*data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleBidirectional = ele2.VrfFlowRuleBidirectional.ValueBool()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleBidirectional = nil
					}

					if !ele2.VrfFlowRuleDstIp.IsNull() && !ele2.VrfFlowRuleDstIp.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleDstIp = ele2.VrfFlowRuleDstIp.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleDstIp = ""
					}

					if !ele2.VrfFlowRuleSrcIp.IsNull() && !ele2.VrfFlowRuleSrcIp.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleSrcIp = ele2.VrfFlowRuleSrcIp.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleSrcIp = ""
					}

					if !ele2.VrfFlowRuleDstPort.IsNull() && !ele2.VrfFlowRuleDstPort.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleDstPort = ele2.VrfFlowRuleDstPort.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleDstPort = ""
					}

					if !ele2.VrfFlowRuleSrcPort.IsNull() && !ele2.VrfFlowRuleSrcPort.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleSrcPort = ele2.VrfFlowRuleSrcPort.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleSrcPort = ""
					}

					if !ele2.VrfFlowRuleProtocol.IsNull() && !ele2.VrfFlowRuleProtocol.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleProtocol = ele2.VrfFlowRuleProtocol.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleProtocol = ""
					}

					if !ele2.VrfFlowRuleAttributeId.IsNull() && !ele2.VrfFlowRuleAttributeId.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleAttributeId = ele2.VrfFlowRuleAttributeId.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2].VrfFlowRuleAttributeId = ""
					}

				} /* for loop */
			} /* isNull if check */

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN InterfaceFlowRules[i1]

	if !v.InterfaceFlowRules.IsNull() && !v.InterfaceFlowRules.IsUnknown() {
		elements := make([]InterfaceFlowRulesValue, len(v.InterfaceFlowRules.Elements()))
		// Not augmenting

		data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules = make([]resource_fabric_common.NDFCInterfaceFlowRulesValue, len(v.InterfaceFlowRules.Elements()))

		// -- Set here 1 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|InterfaceFlowRules[i1]| --

		diag := v.InterfaceFlowRules.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1] InterfaceFlowRules[i1]
		for i1, ele1 := range elements {
			if !ele1.InterfaceFlowRuleName.IsNull() && !ele1.InterfaceFlowRuleName.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleName = ele1.InterfaceFlowRuleName.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleName = ""
			}

			if !ele1.InterfaceFlowRuleUuid.IsNull() && !ele1.InterfaceFlowRuleUuid.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleUuid = ele1.InterfaceFlowRuleUuid.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleUuid = ""
			}

			if !ele1.InterfaceFlowRuleType.IsNull() && !ele1.InterfaceFlowRuleType.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleType = ele1.InterfaceFlowRuleType.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleType = ""
			}

			// here 507 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|

			//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]

			if !ele1.InterfaceFlowRuleInterfaceCollection.IsNull() && !ele1.InterfaceFlowRuleInterfaceCollection.IsUnknown() {
				elements := make([]InterfaceFlowRuleInterfaceCollectionValue, len(ele1.InterfaceFlowRuleInterfaceCollection.Elements()))
				// Not augmenting

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection = make([]resource_fabric_common.NDFCInterfaceFlowRuleInterfaceCollectionValue, len(ele1.InterfaceFlowRuleInterfaceCollection.Elements()))

				diag := ele1.InterfaceFlowRuleInterfaceCollection.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				// .TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2] TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]
				for i2, ele2 := range elements {
					if !ele2.InterfaceFlowRuleSwitchId.IsNull() && !ele2.InterfaceFlowRuleSwitchId.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleSwitchId = ele2.InterfaceFlowRuleSwitchId.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleSwitchId = ""
					}

					if !ele2.InterfaceFlowRuleSwitchName.IsNull() && !ele2.InterfaceFlowRuleSwitchName.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleSwitchName = ele2.InterfaceFlowRuleSwitchName.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleSwitchName = ""
					}

					if !ele2.InterfaceFlowRuleInterfaces.IsNull() && !ele2.InterfaceFlowRuleInterfaces.IsUnknown() {

						// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]|[]|interface_flow_rule_interfaces|
						listStringData := make([]string, len(ele2.InterfaceFlowRuleInterfaces.Elements()))
						dg := ele2.InterfaceFlowRuleInterfaces.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleInterfaces = make([]string, len(listStringData))
						copy(data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2].InterfaceFlowRuleInterfaces, listStringData)
					}

				} /* for loop */
			} /* isNull if check */

			if !ele1.InterfaceFlowRuleSubnets.IsNull() && !ele1.InterfaceFlowRuleSubnets.IsUnknown() {

				// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|[]|interface_flow_rule_subnets|
				listStringData := make([]string, len(ele1.InterfaceFlowRuleSubnets.Elements()))
				dg := ele1.InterfaceFlowRuleSubnets.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleSubnets = make([]string, len(listStringData))
				copy(data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleSubnets, listStringData)
			}

			// here 507 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1]|

			//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2]

			if !ele1.InterfaceFlowRuleAttributes.IsNull() && !ele1.InterfaceFlowRuleAttributes.IsUnknown() {
				elements := make([]InterfaceFlowRuleAttributesValue, len(ele1.InterfaceFlowRuleAttributes.Elements()))
				// Not augmenting

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes = make([]resource_fabric_common.NDFCInterfaceFlowRuleAttributesValue, len(ele1.InterfaceFlowRuleAttributes.Elements()))

				diag := ele1.InterfaceFlowRuleAttributes.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				// .TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2] TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2]
				for i2, ele2 := range elements {
					if !ele2.InterfaceFlowRuleBidirectional.IsNull() && !ele2.InterfaceFlowRuleBidirectional.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleBidirectional = new(bool)
						*data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleBidirectional = ele2.InterfaceFlowRuleBidirectional.ValueBool()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleBidirectional = nil
					}

					if !ele2.InterfaceFlowRuleDstIp.IsNull() && !ele2.InterfaceFlowRuleDstIp.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleDstIp = ele2.InterfaceFlowRuleDstIp.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleDstIp = ""
					}

					if !ele2.InterfaceFlowRuleSrcIp.IsNull() && !ele2.InterfaceFlowRuleSrcIp.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleSrcIp = ele2.InterfaceFlowRuleSrcIp.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleSrcIp = ""
					}

					if !ele2.InterfaceFlowRuleDstPort.IsNull() && !ele2.InterfaceFlowRuleDstPort.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleDstPort = ele2.InterfaceFlowRuleDstPort.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleDstPort = ""
					}

					if !ele2.InterfaceFlowRuleSrcPort.IsNull() && !ele2.InterfaceFlowRuleSrcPort.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleSrcPort = ele2.InterfaceFlowRuleSrcPort.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleSrcPort = ""
					}

					if !ele2.InterfaceFlowRuleProtocol.IsNull() && !ele2.InterfaceFlowRuleProtocol.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleProtocol = ele2.InterfaceFlowRuleProtocol.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleProtocol = ""
					}

					if !ele2.InterfaceFlowRuleAttributeId.IsNull() && !ele2.InterfaceFlowRuleAttributeId.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleAttributeId = ele2.InterfaceFlowRuleAttributeId.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2].InterfaceFlowRuleAttributeId = ""
					}

				} /* for loop */
			} /* isNull if check */

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN L3OutFlowRules[i1]

	if !v.L3OutFlowRules.IsNull() && !v.L3OutFlowRules.IsUnknown() {
		elements := make([]L3OutFlowRulesValue, len(v.L3OutFlowRules.Elements()))
		// Not augmenting

		data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules = make([]resource_fabric_common.NDFCL3OutFlowRulesValue, len(v.L3OutFlowRules.Elements()))

		// -- Set here 1 |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1]|L3OutFlowRules[i1]| --

		diag := v.L3OutFlowRules.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1] L3OutFlowRules[i1]
		for i1, ele1 := range elements {
			if !ele1.L3OutFlowRuleName.IsNull() && !ele1.L3OutFlowRuleName.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleName = ele1.L3OutFlowRuleName.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleName = ""
			}

			if !ele1.L3OutFlowRuleUuid.IsNull() && !ele1.L3OutFlowRuleUuid.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleUuid = ele1.L3OutFlowRuleUuid.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleUuid = ""
			}

			if !ele1.L3OutFlowRuleType.IsNull() && !ele1.L3OutFlowRuleType.IsUnknown() {

				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleType = ele1.L3OutFlowRuleType.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleType = ""
			}

			// here 507 |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1]|TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1]|

			//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]

			if !ele1.L3OutFlowRuleInterfaceCollection.IsNull() && !ele1.L3OutFlowRuleInterfaceCollection.IsUnknown() {
				elements := make([]L3OutFlowRuleInterfaceCollectionValue, len(ele1.L3OutFlowRuleInterfaceCollection.Elements()))
				// Not augmenting

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection|TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection = make([]resource_fabric_common.NDFCL3OutFlowRuleInterfaceCollectionValue, len(ele1.L3OutFlowRuleInterfaceCollection.Elements()))

				diag := ele1.L3OutFlowRuleInterfaceCollection.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				// .TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2] TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]
				for i2, ele2 := range elements {
					if !ele2.L3OutFlowRuleTenant.IsNull() && !ele2.L3OutFlowRuleTenant.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleTenant = ele2.L3OutFlowRuleTenant.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleTenant = ""
					}

					if !ele2.L3OutFlowRuleL3Out.IsNull() && !ele2.L3OutFlowRuleL3Out.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleL3Out = ele2.L3OutFlowRuleL3Out.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleL3Out = ""
					}

					if !ele2.L3OutFlowRuleEncap.IsNull() && !ele2.L3OutFlowRuleEncap.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleEncap = ele2.L3OutFlowRuleEncap.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleEncap = ""
					}

					if !ele2.L3OutFlowRuleSwitchName.IsNull() && !ele2.L3OutFlowRuleSwitchName.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleSwitchName = ele2.L3OutFlowRuleSwitchName.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleSwitchName = ""
					}

					if !ele2.L3OutFlowRuleSwitchId.IsNull() && !ele2.L3OutFlowRuleSwitchId.IsUnknown() {

						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleSwitchId = ele2.L3OutFlowRuleSwitchId.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleSwitchId = ""
					}

					if !ele2.L3OutFlowRuleInterfaces.IsNull() && !ele2.L3OutFlowRuleInterfaces.IsUnknown() {

						// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]|[]|l3_out_flow_rule_interfaces|
						listStringData := make([]string, len(ele2.L3OutFlowRuleInterfaces.Elements()))
						dg := ele2.L3OutFlowRuleInterfaces.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleInterfaces = make([]string, len(listStringData))
						copy(data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2].L3OutFlowRuleInterfaces, listStringData)
					}

				} /* for loop */
			} /* isNull if check */

			if !ele1.L3OutFlowRuleSubnets.IsNull() && !ele1.L3OutFlowRuleSubnets.IsUnknown() {

				// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1]|[]|l3_out_flow_rule_subnets|
				listStringData := make([]string, len(ele1.L3OutFlowRuleSubnets.Elements()))
				dg := ele1.L3OutFlowRuleSubnets.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleSubnets = make([]string, len(listStringData))
				copy(data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleSubnets, listStringData)
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN InterfaceRules[i1]

	if !v.InterfaceRules.IsNull() && !v.InterfaceRules.IsUnknown() {
		elements := make([]InterfaceRulesValue, len(v.InterfaceRules.Elements()))
		// Not augmenting

		data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules = make([]resource_fabric_common.NDFCInterfaceRulesValue, len(v.InterfaceRules.Elements()))

		// -- Set here 1 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1]|InterfaceRules[i1]| --

		diag := v.InterfaceRules.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1] InterfaceRules[i1]
		for i1, ele1 := range elements {
			if !ele1.InterfaceRuleName.IsNull() && !ele1.InterfaceRuleName.IsUnknown() {

				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleName = ele1.InterfaceRuleName.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleName = ""
			}

			// here 507 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1]|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1]|

			//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]

			if !ele1.InterfaceRuleInterfaceCollection.IsNull() && !ele1.InterfaceRuleInterfaceCollection.IsUnknown() {
				elements := make([]InterfaceRuleInterfaceCollectionValue, len(ele1.InterfaceRuleInterfaceCollection.Elements()))
				// Not augmenting

				// -- Set here 2 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection = make([]resource_fabric_common.NDFCInterfaceRuleInterfaceCollectionValue, len(ele1.InterfaceRuleInterfaceCollection.Elements()))

				diag := ele1.InterfaceRuleInterfaceCollection.ElementsAs(context.Background(), &elements, false)
				if diag != nil {
					panic(diag)
				}
				// .TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2] TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]
				for i2, ele2 := range elements {
					if !ele2.InterfaceRuleSwitchId.IsNull() && !ele2.InterfaceRuleSwitchId.IsUnknown() {

						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleSwitchId = ele2.InterfaceRuleSwitchId.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleSwitchId = ""
					}

					if !ele2.InterfaceRuleSwitchName.IsNull() && !ele2.InterfaceRuleSwitchName.IsUnknown() {

						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleSwitchName = ele2.InterfaceRuleSwitchName.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleSwitchName = ""
					}

					if !ele2.InterfaceRuleVrfName.IsNull() && !ele2.InterfaceRuleVrfName.IsUnknown() {

						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleVrfName = ele2.InterfaceRuleVrfName.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleVrfName = ""
					}

					// here 507 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]|

					//MARSHALL_LIST  BEGIN TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3]

					if !ele2.InterfaceRuleInterfaces.IsNull() && !ele2.InterfaceRuleInterfaces.IsUnknown() {
						elements := make([]InterfaceRuleInterfacesValue, len(ele2.InterfaceRuleInterfaces.Elements()))
						// Not augmenting

						// -- Set here 2 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3]| --
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces = make([]resource_fabric_common.NDFCInterfaceRuleInterfacesValue, len(ele2.InterfaceRuleInterfaces.Elements()))

						diag := ele2.InterfaceRuleInterfaces.ElementsAs(context.Background(), &elements, false)
						if diag != nil {
							panic(diag)
						}
						// .TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3] TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3]
						for i3, ele3 := range elements {
							if !ele3.InterfaceRuleInterfaceName.IsNull() && !ele3.InterfaceRuleInterfaceName.IsUnknown() {

								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceName = ele3.InterfaceRuleInterfaceName.ValueString()
							} else {
								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceName = ""
							}

							if !ele3.InterfaceRuleInterfaceType.IsNull() && !ele3.InterfaceRuleInterfaceType.IsUnknown() {

								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceType = ele3.InterfaceRuleInterfaceType.ValueString()
							} else {
								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceType = ""
							}

							if !ele3.InterfaceRuleInterfaceEncap.IsNull() && !ele3.InterfaceRuleInterfaceEncap.IsUnknown() {

								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceEncap = ele3.InterfaceRuleInterfaceEncap.ValueString()
							} else {
								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceEncap = ""
							}

							if !ele3.InterfaceRuleInterfaceVrfName.IsNull() && !ele3.InterfaceRuleInterfaceVrfName.IsUnknown() {

								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceVrfName = ele3.InterfaceRuleInterfaceVrfName.ValueString()
							} else {
								data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3].InterfaceRuleInterfaceVrfName = ""
							}

						} /* for loop */
					} /* isNull if check */

					if !ele2.InterfaceRuleTenant.IsNull() && !ele2.InterfaceRuleTenant.IsUnknown() {

						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleTenant = ele2.InterfaceRuleTenant.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleTenant = ""
					}

					if !ele2.InterfaceRuleL3Out.IsNull() && !ele2.InterfaceRuleL3Out.IsUnknown() {

						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleL3Out = ele2.InterfaceRuleL3Out.ValueString()
					} else {
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleL3Out = ""
					}

				} /* for loop */
			} /* isNull if check */

			if !ele1.InterfaceRuleEnabled.IsNull() && !ele1.InterfaceRuleEnabled.IsUnknown() {

				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnabled = new(bool)
				*data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnabled = ele1.InterfaceRuleEnabled.ValueBool()
			} else {
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnabled = nil
			}

			if !ele1.InterfaceRuleEnableFabricInterconnect.IsNull() && !ele1.InterfaceRuleEnableFabricInterconnect.IsUnknown() {

				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableFabricInterconnect = new(bool)
				*data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableFabricInterconnect = ele1.InterfaceRuleEnableFabricInterconnect.ValueBool()
			} else {
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableFabricInterconnect = nil
			}

			if !ele1.InterfaceRuleEnableL3Out.IsNull() && !ele1.InterfaceRuleEnableL3Out.IsUnknown() {

				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableL3Out = new(bool)
				*data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableL3Out = ele1.InterfaceRuleEnableL3Out.ValueBool()
			} else {
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleEnableL3Out = nil
			}

			if !ele1.InterfaceRuleUuid.IsNull() && !ele1.InterfaceRuleUuid.IsUnknown() {

				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleUuid = ele1.InterfaceRuleUuid.ValueString()
			} else {
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleUuid = ""
			}

			if !ele1.InterfaceRuleSubnets.IsNull() && !ele1.InterfaceRuleSubnets.IsUnknown() {

				// Nested List:String inside a list - which is not having NDFCNested |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1]|[]|interface_rule_subnets|
				listStringData := make([]string, len(ele1.InterfaceRuleSubnets.Elements()))
				dg := ele1.InterfaceRuleSubnets.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleSubnets = make([]string, len(listStringData))
				copy(data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleSubnets, listStringData)
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN Email[i1]

	if !v.Email.IsNull() && !v.Email.IsUnknown() {
		elements := make([]EmailValue, len(v.Email.Elements()))
		// Not augmenting

		data.ExternalStreamingSettings.Email = make([]resource_fabric_common.NDFCEmailValue, len(v.Email.Elements()))

		// -- Set here 1 |.ExternalStreamingSettings.Email[i1]|Email[i1]| --

		diag := v.Email.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .ExternalStreamingSettings.Email[i1] Email[i1]
		for i1, ele1 := range elements {
			if !ele1.Name.IsNull() && !ele1.Name.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].Name = ele1.Name.ValueString()
			} else {
				data.ExternalStreamingSettings.Email[i1].Name = ""
			}

			if !ele1.ReceiverEmail.IsNull() && !ele1.ReceiverEmail.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].ReceiverEmail = ele1.ReceiverEmail.ValueString()
			} else {
				data.ExternalStreamingSettings.Email[i1].ReceiverEmail = ""
			}

			if !ele1.Format.IsNull() && !ele1.Format.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].Format = ele1.Format.ValueString()
			} else {
				data.ExternalStreamingSettings.Email[i1].Format = ""
			}

			if !ele1.StartDate.IsNull() && !ele1.StartDate.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].StartDate = ele1.StartDate.ValueString()
			} else {
				data.ExternalStreamingSettings.Email[i1].StartDate = ""
			}

			if !ele1.CollectionFrequencyInDays.IsNull() && !ele1.CollectionFrequencyInDays.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].CollectionFrequencyInDays = new(int64)
				*data.ExternalStreamingSettings.Email[i1].CollectionFrequencyInDays = ele1.CollectionFrequencyInDays.ValueInt64()

			} else {
				data.ExternalStreamingSettings.Email[i1].CollectionFrequencyInDays = nil
			}

			//-----inline nesting Start---- .ExternalStreamingSettings.Email[i1]

			if !ele1.CollectionType.IsNull() && !ele1.CollectionType.IsUnknown() {
				data.ExternalStreamingSettings.Email[i1].CollectionSettings.CollectionType = ele1.CollectionType.ValueString()
			} else {
				data.ExternalStreamingSettings.Email[i1].CollectionSettings.CollectionType = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.Email[i1]

			if !ele1.Anomalies.IsNull() && !ele1.Anomalies.IsUnknown() {
				listStringData := make([]string, len(ele1.Anomalies.Elements()))
				dg := ele1.Anomalies.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.Email[i1].CollectionSettings.Anomalies = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.Email[i1].CollectionSettings.Anomalies, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.Email[i1]

			if !ele1.Advisories.IsNull() && !ele1.Advisories.IsUnknown() {
				listStringData := make([]string, len(ele1.Advisories.Elements()))
				dg := ele1.Advisories.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.Email[i1].CollectionSettings.Advisories = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.Email[i1].CollectionSettings.Advisories, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.Email[i1]

			if !ele1.RiskAndConformanceReports.IsNull() && !ele1.RiskAndConformanceReports.IsUnknown() {
				listStringData := make([]string, len(ele1.RiskAndConformanceReports.Elements()))
				dg := ele1.RiskAndConformanceReports.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.Email[i1].CollectionSettings.RiskAndConformanceReports = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.Email[i1].CollectionSettings.RiskAndConformanceReports, listStringData)
			}

			//-----inline nesting end----

			if !ele1.OnlyIncludeActiveAlerts.IsNull() && !ele1.OnlyIncludeActiveAlerts.IsUnknown() {

				data.ExternalStreamingSettings.Email[i1].OnlyIncludeActiveAlerts = new(bool)
				*data.ExternalStreamingSettings.Email[i1].OnlyIncludeActiveAlerts = ele1.OnlyIncludeActiveAlerts.ValueBool()
			} else {
				data.ExternalStreamingSettings.Email[i1].OnlyIncludeActiveAlerts = nil
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN MessageBus[i1]

	if !v.MessageBus.IsNull() && !v.MessageBus.IsUnknown() {
		elements := make([]MessageBusValue, len(v.MessageBus.Elements()))
		// Not augmenting

		data.ExternalStreamingSettings.MessageBus = make([]resource_fabric_common.NDFCMessageBusValue, len(v.MessageBus.Elements()))

		// -- Set here 1 |.ExternalStreamingSettings.MessageBus[i1]|MessageBus[i1]| --

		diag := v.MessageBus.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .ExternalStreamingSettings.MessageBus[i1] MessageBus[i1]
		for i1, ele1 := range elements {
			if !ele1.Server.IsNull() && !ele1.Server.IsUnknown() {

				data.ExternalStreamingSettings.MessageBus[i1].Server = ele1.Server.ValueString()
			} else {
				data.ExternalStreamingSettings.MessageBus[i1].Server = ""
			}

			if !ele1.CollectionType.IsNull() && !ele1.CollectionType.IsUnknown() {

				data.ExternalStreamingSettings.MessageBus[i1].CollectionType = ele1.CollectionType.ValueString()
			} else {
				data.ExternalStreamingSettings.MessageBus[i1].CollectionType = ""
			}

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.CollectionSettingsCollectionType.IsNull() && !ele1.CollectionSettingsCollectionType.IsUnknown() {
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.CollectionSettingsCollectionType = ele1.CollectionSettingsCollectionType.ValueString()
			} else {
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.CollectionSettingsCollectionType = ""
			}
			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.Anomalies.IsNull() && !ele1.Anomalies.IsUnknown() {
				listStringData := make([]string, len(ele1.Anomalies.Elements()))
				dg := ele1.Anomalies.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Anomalies = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Anomalies, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.Advisories.IsNull() && !ele1.Advisories.IsUnknown() {
				listStringData := make([]string, len(ele1.Advisories.Elements()))
				dg := ele1.Advisories.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Advisories = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Advisories, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.Statistics.IsNull() && !ele1.Statistics.IsUnknown() {
				listStringData := make([]string, len(ele1.Statistics.Elements()))
				dg := ele1.Statistics.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Statistics = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Statistics, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.Faults.IsNull() && !ele1.Faults.IsUnknown() {
				listStringData := make([]string, len(ele1.Faults.Elements()))
				dg := ele1.Faults.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Faults = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.Faults, listStringData)
			}

			//-----inline nesting end----

			//-----inline nesting Start---- .ExternalStreamingSettings.MessageBus[i1]

			if !ele1.AuditLogs.IsNull() && !ele1.AuditLogs.IsUnknown() {
				listStringData := make([]string, len(ele1.AuditLogs.Elements()))
				dg := ele1.AuditLogs.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.AuditLogs = make([]string, len(listStringData))

				copy(data.ExternalStreamingSettings.MessageBus[i1].CollectionSettings.AuditLogs, listStringData)
			}

			//-----inline nesting end----

		} /* for loop */
	} /* isNull if check */

	return data
}
