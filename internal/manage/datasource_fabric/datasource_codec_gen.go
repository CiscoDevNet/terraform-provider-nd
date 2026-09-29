// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_fabric

import (
	"context"
	"log"
	"strconv"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v *FabricModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
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

	if jsonData.Management.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.Management.FabricType)

	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.Management.BgpAsn != "" {
		v.BgpAsn = types.StringValue(jsonData.Management.BgpAsn)

	} else {
		v.BgpAsn = types.StringNull()
	}

	if jsonData.Management.CreateBgpConfig != nil {
		v.CreateBgpConfig = types.BoolValue(*jsonData.Management.CreateBgpConfig)

	} else {
		v.CreateBgpConfig = types.BoolNull()
	}

	if jsonData.Management.SuperSpineBgpAs != "" {
		v.SuperSpineBgpAs = types.StringValue(jsonData.Management.SuperSpineBgpAs)

	} else {
		v.SuperSpineBgpAs = types.StringNull()
	}

	if jsonData.Management.LeafBgpAs != "" {
		v.LeafBgpAs = types.StringValue(jsonData.Management.LeafBgpAs)

	} else {
		v.LeafBgpAs = types.StringNull()
	}

	if jsonData.Management.BorderBgpAs != "" {
		v.BorderBgpAs = types.StringValue(jsonData.Management.BorderBgpAs)

	} else {
		v.BorderBgpAs = types.StringNull()
	}

	if jsonData.Management.BgpAsMode != "" {
		v.BgpAsMode = types.StringValue(jsonData.Management.BgpAsMode)

	} else {
		v.BgpAsMode = types.StringNull()
	}

	if jsonData.Management.BgpAsnAutoAllocation != nil {
		v.BgpAsnAutoAllocation = types.BoolValue(*jsonData.Management.BgpAsnAutoAllocation)

	} else {
		v.BgpAsnAutoAllocation = types.BoolNull()
	}

	if jsonData.Management.BgpAsnRange != "" {
		v.BgpAsnRange = types.StringValue(jsonData.Management.BgpAsnRange)

	} else {
		v.BgpAsnRange = types.StringNull()
	}

	if jsonData.Management.BgpAllowAsInNum != nil {
		v.BgpAllowAsInNum = types.Int64Value(*jsonData.Management.BgpAllowAsInNum)

	} else {
		v.BgpAllowAsInNum = types.Int64Null()
	}

	if jsonData.Management.BgpMaxPath != nil {
		v.BgpMaxPath = types.Int64Value(*jsonData.Management.BgpMaxPath)

	} else {
		v.BgpMaxPath = types.Int64Null()
	}

	if jsonData.Management.BgpUnderlayFailureProtect != nil {
		v.BgpUnderlayFailureProtect = types.BoolValue(*jsonData.Management.BgpUnderlayFailureProtect)

	} else {
		v.BgpUnderlayFailureProtect = types.BoolNull()
	}

	if jsonData.Management.AutoConfigureEbgpEvpnPeering != nil {
		v.AutoConfigureEbgpEvpnPeering = types.BoolValue(*jsonData.Management.AutoConfigureEbgpEvpnPeering)

	} else {
		v.AutoConfigureEbgpEvpnPeering = types.BoolNull()
	}

	if jsonData.Management.AllowLeafSameAs != nil {
		v.AllowLeafSameAs = types.BoolValue(*jsonData.Management.AllowLeafSameAs)

	} else {
		v.AllowLeafSameAs = types.BoolNull()
	}

	if jsonData.Management.AssignIpv4ToLoopback0 != nil {
		v.AssignIpv4ToLoopback0 = types.BoolValue(*jsonData.Management.AssignIpv4ToLoopback0)

	} else {
		v.AssignIpv4ToLoopback0 = types.BoolNull()
	}

	if jsonData.Management.TargetSubnetMask != nil {
		v.TargetSubnetMask = types.Int64Value(*jsonData.Management.TargetSubnetMask)

	} else {
		v.TargetSubnetMask = types.Int64Null()
	}

	if jsonData.Management.AnycastGatewayMac != "" {
		v.AnycastGatewayMac = types.StringValue(jsonData.Management.AnycastGatewayMac)

	} else {
		v.AnycastGatewayMac = types.StringNull()
	}

	if jsonData.Management.PerformanceMonitoring != nil {
		v.PerformanceMonitoring = types.BoolValue(*jsonData.Management.PerformanceMonitoring)

	} else {
		v.PerformanceMonitoring = types.BoolNull()
	}

	if jsonData.Management.ReplicationMode != "" {
		v.ReplicationMode = types.StringValue(jsonData.Management.ReplicationMode)

	} else {
		v.ReplicationMode = types.StringNull()
	}

	if jsonData.Management.MulticastGroupSubnet != "" {
		v.MulticastGroupSubnet = types.StringValue(jsonData.Management.MulticastGroupSubnet)

	} else {
		v.MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.Management.UnderlayMulticastGroupAddressLimit != nil {
		v.UnderlayMulticastGroupAddressLimit = types.Int64Value(*jsonData.Management.UnderlayMulticastGroupAddressLimit)

	} else {
		v.UnderlayMulticastGroupAddressLimit = types.Int64Null()
	}

	if jsonData.Management.TenantRoutedMulticast != nil {
		v.TenantRoutedMulticast = types.BoolValue(*jsonData.Management.TenantRoutedMulticast)

	} else {
		v.TenantRoutedMulticast = types.BoolNull()
	}

	if jsonData.Management.RendezvousPointCount != nil {
		v.RendezvousPointCount = types.Int64Value(*jsonData.Management.RendezvousPointCount)

	} else {
		v.RendezvousPointCount = types.Int64Null()
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

	if jsonData.Management.RendezvousPointLoopbackId != nil {
		v.RendezvousPointLoopbackId = types.Int64Value(*jsonData.Management.RendezvousPointLoopbackId)

	} else {
		v.RendezvousPointLoopbackId = types.Int64Null()
	}

	if jsonData.Management.VpcPeerLinkVlan != "" {
		v.VpcPeerLinkVlan = types.StringValue(jsonData.Management.VpcPeerLinkVlan)

	} else {
		v.VpcPeerLinkVlan = types.StringNull()
	}

	if jsonData.Management.VpcPeerLinkEnableNativeVlan != nil {
		v.VpcPeerLinkEnableNativeVlan = types.BoolValue(*jsonData.Management.VpcPeerLinkEnableNativeVlan)

	} else {
		v.VpcPeerLinkEnableNativeVlan = types.BoolNull()
	}

	if jsonData.Management.VpcPeerKeepAliveOption != "" {
		v.VpcPeerKeepAliveOption = types.StringValue(jsonData.Management.VpcPeerKeepAliveOption)

	} else {
		v.VpcPeerKeepAliveOption = types.StringNull()
	}

	if jsonData.Management.VpcAutoRecoveryTimer != nil {
		v.VpcAutoRecoveryTimer = types.Int64Value(*jsonData.Management.VpcAutoRecoveryTimer)

	} else {
		v.VpcAutoRecoveryTimer = types.Int64Null()
	}

	if jsonData.Management.VpcDelayRestoreTimer != nil {
		v.VpcDelayRestoreTimer = types.Int64Value(*jsonData.Management.VpcDelayRestoreTimer)

	} else {
		v.VpcDelayRestoreTimer = types.Int64Null()
	}

	if jsonData.Management.VpcPeerLinkPortChannelId != "" {
		v.VpcPeerLinkPortChannelId = types.StringValue(jsonData.Management.VpcPeerLinkPortChannelId)

	} else {
		v.VpcPeerLinkPortChannelId = types.StringNull()
	}

	if jsonData.Management.VpcIpv6NeighborDiscoverySync != nil {
		v.VpcIpv6NeighborDiscoverySync = types.BoolValue(*jsonData.Management.VpcIpv6NeighborDiscoverySync)

	} else {
		v.VpcIpv6NeighborDiscoverySync = types.BoolNull()
	}

	if jsonData.Management.AdvertisePhysicalIp != nil {
		v.AdvertisePhysicalIp = types.BoolValue(*jsonData.Management.AdvertisePhysicalIp)

	} else {
		v.AdvertisePhysicalIp = types.BoolNull()
	}

	if jsonData.Management.VpcDomainIdRange != "" {
		v.VpcDomainIdRange = types.StringValue(jsonData.Management.VpcDomainIdRange)

	} else {
		v.VpcDomainIdRange = types.StringNull()
	}

	if jsonData.Management.BgpLoopbackId != nil {
		v.BgpLoopbackId = types.Int64Value(*jsonData.Management.BgpLoopbackId)

	} else {
		v.BgpLoopbackId = types.Int64Null()
	}

	if jsonData.Management.AllowSameLoopbackIpOnSwitches != nil {
		v.AllowSameLoopbackIpOnSwitches = types.BoolValue(*jsonData.Management.AllowSameLoopbackIpOnSwitches)

	} else {
		v.AllowSameLoopbackIpOnSwitches = types.BoolNull()
	}

	if jsonData.Management.NveLoopbackId != nil {
		v.NveLoopbackId = types.Int64Value(*jsonData.Management.NveLoopbackId)

	} else {
		v.NveLoopbackId = types.Int64Null()
	}

	if jsonData.Management.VrfTemplate != "" {
		v.VrfTemplate = types.StringValue(jsonData.Management.VrfTemplate)

	} else {
		v.VrfTemplate = types.StringNull()
	}

	if jsonData.Management.NetworkTemplate != "" {
		v.NetworkTemplate = types.StringValue(jsonData.Management.NetworkTemplate)

	} else {
		v.NetworkTemplate = types.StringNull()
	}

	if jsonData.Management.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.Management.VrfExtensionTemplate)

	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	if jsonData.Management.NetworkExtensionTemplate != "" {
		v.NetworkExtensionTemplate = types.StringValue(jsonData.Management.NetworkExtensionTemplate)

	} else {
		v.NetworkExtensionTemplate = types.StringNull()
	}

	if jsonData.Management.L3VniNoVlanDefaultOption != nil {
		v.L3VniNoVlanDefaultOption = types.BoolValue(*jsonData.Management.L3VniNoVlanDefaultOption)

	} else {
		v.L3VniNoVlanDefaultOption = types.BoolNull()
	}

	if jsonData.Management.SiteId != "" {
		v.SiteId = types.StringValue(jsonData.Management.SiteId)

	} else {
		v.SiteId = types.StringNull()
	}

	if jsonData.Management.FabricMtu != nil {
		v.FabricMtu = types.Int64Value(*jsonData.Management.FabricMtu)

	} else {
		v.FabricMtu = types.Int64Null()
	}

	if jsonData.Management.L2HostInterfaceMtu != nil {
		v.L2HostInterfaceMtu = types.Int64Value(*jsonData.Management.L2HostInterfaceMtu)

	} else {
		v.L2HostInterfaceMtu = types.Int64Null()
	}

	if jsonData.Management.TenantDhcp != nil {
		v.TenantDhcp = types.BoolValue(*jsonData.Management.TenantDhcp)

	} else {
		v.TenantDhcp = types.BoolNull()
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

	if jsonData.Management.SnmpTrap != nil {
		v.SnmpTrap = types.BoolValue(*jsonData.Management.SnmpTrap)

	} else {
		v.SnmpTrap = types.BoolNull()
	}

	if jsonData.Management.AnycastBorderGatewayAdvertisePhysicalIp != nil {
		v.AnycastBorderGatewayAdvertisePhysicalIp = types.BoolValue(*jsonData.Management.AnycastBorderGatewayAdvertisePhysicalIp)

	} else {
		v.AnycastBorderGatewayAdvertisePhysicalIp = types.BoolNull()
	}

	if jsonData.Management.GreenfieldDebugFlag != "" {
		v.GreenfieldDebugFlag = types.StringValue(jsonData.Management.GreenfieldDebugFlag)

	} else {
		v.GreenfieldDebugFlag = types.StringNull()
	}

	if jsonData.Management.TcamAllocation != nil {
		v.TcamAllocation = types.BoolValue(*jsonData.Management.TcamAllocation)

	} else {
		v.TcamAllocation = types.BoolNull()
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

	if jsonData.Management.BgpLoopbackIpRange != "" {
		v.BgpLoopbackIpRange = types.StringValue(jsonData.Management.BgpLoopbackIpRange)

	} else {
		v.BgpLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.NveLoopbackIpRange != "" {
		v.NveLoopbackIpRange = types.StringValue(jsonData.Management.NveLoopbackIpRange)

	} else {
		v.NveLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.AnycastRendezvousPointIpRange != "" {
		v.AnycastRendezvousPointIpRange = types.StringValue(jsonData.Management.AnycastRendezvousPointIpRange)

	} else {
		v.AnycastRendezvousPointIpRange = types.StringNull()
	}

	if jsonData.Management.IntraFabricSubnetRange != "" {
		v.IntraFabricSubnetRange = types.StringValue(jsonData.Management.IntraFabricSubnetRange)

	} else {
		v.IntraFabricSubnetRange = types.StringNull()
	}

	if jsonData.Management.L2VniRange != "" {
		v.L2VniRange = types.StringValue(jsonData.Management.L2VniRange)

	} else {
		v.L2VniRange = types.StringNull()
	}

	if jsonData.Management.L3VniRange != "" {
		v.L3VniRange = types.StringValue(jsonData.Management.L3VniRange)

	} else {
		v.L3VniRange = types.StringNull()
	}

	if jsonData.Management.NetworkVlanRange != "" {
		v.NetworkVlanRange = types.StringValue(jsonData.Management.NetworkVlanRange)

	} else {
		v.NetworkVlanRange = types.StringNull()
	}

	if jsonData.Management.VrfVlanRange != "" {
		v.VrfVlanRange = types.StringValue(jsonData.Management.VrfVlanRange)

	} else {
		v.VrfVlanRange = types.StringNull()
	}

	if jsonData.Management.SubInterfaceDot1qRange != "" {
		v.SubInterfaceDot1qRange = types.StringValue(jsonData.Management.SubInterfaceDot1qRange)

	} else {
		v.SubInterfaceDot1qRange = types.StringNull()
	}

	if jsonData.Management.VrfLiteAutoConfig != "" {
		v.VrfLiteAutoConfig = types.StringValue(jsonData.Management.VrfLiteAutoConfig)

	} else {
		v.VrfLiteAutoConfig = types.StringNull()
	}

	if jsonData.Management.VrfLiteSubnetRange != "" {
		v.VrfLiteSubnetRange = types.StringValue(jsonData.Management.VrfLiteSubnetRange)

	} else {
		v.VrfLiteSubnetRange = types.StringNull()
	}

	if jsonData.Management.VrfLiteSubnetTargetMask != nil {
		v.VrfLiteSubnetTargetMask = types.Int64Value(*jsonData.Management.VrfLiteSubnetTargetMask)

	} else {
		v.VrfLiteSubnetTargetMask = types.Int64Null()
	}

	if jsonData.Management.VrfLiteIpv6SubnetRange != "" {
		v.VrfLiteIpv6SubnetRange = types.StringValue(jsonData.Management.VrfLiteIpv6SubnetRange)

	} else {
		v.VrfLiteIpv6SubnetRange = types.StringNull()
	}

	if jsonData.Management.VrfLiteIpv6SubnetTargetMask != nil {
		v.VrfLiteIpv6SubnetTargetMask = types.Int64Value(*jsonData.Management.VrfLiteIpv6SubnetTargetMask)

	} else {
		v.VrfLiteIpv6SubnetTargetMask = types.Int64Null()
	}

	if jsonData.Management.AutoUniqueVrfLiteIpPrefix != nil {
		v.AutoUniqueVrfLiteIpPrefix = types.BoolValue(*jsonData.Management.AutoUniqueVrfLiteIpPrefix)

	} else {
		v.AutoUniqueVrfLiteIpPrefix = types.BoolNull()
	}

	if jsonData.Management.PerVrfLoopbackAutoProvision != nil {
		v.PerVrfLoopbackAutoProvision = types.BoolValue(*jsonData.Management.PerVrfLoopbackAutoProvision)

	} else {
		v.PerVrfLoopbackAutoProvision = types.BoolNull()
	}

	if jsonData.Management.PerVrfLoopbackIpRange != "" {
		v.PerVrfLoopbackIpRange = types.StringValue(jsonData.Management.PerVrfLoopbackIpRange)

	} else {
		v.PerVrfLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.PerVrfLoopbackAutoProvisionIpv6 != nil {
		v.PerVrfLoopbackAutoProvisionIpv6 = types.BoolValue(*jsonData.Management.PerVrfLoopbackAutoProvisionIpv6)

	} else {
		v.PerVrfLoopbackAutoProvisionIpv6 = types.BoolNull()
	}

	if jsonData.Management.PerVrfLoopbackIpv6Range != "" {
		v.PerVrfLoopbackIpv6Range = types.StringValue(jsonData.Management.PerVrfLoopbackIpv6Range)

	} else {
		v.PerVrfLoopbackIpv6Range = types.StringNull()
	}

	if jsonData.Management.Banner != "" {
		v.Banner = types.StringValue(jsonData.Management.Banner)

	} else {
		v.Banner = types.StringNull()
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

	if jsonData.Management.LocalDhcpServer != nil {
		v.LocalDhcpServer = types.BoolValue(*jsonData.Management.LocalDhcpServer)

	} else {
		v.LocalDhcpServer = types.BoolNull()
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

	if jsonData.Management.DhcpEndAddress != "" {
		v.DhcpEndAddress = types.StringValue(jsonData.Management.DhcpEndAddress)

	} else {
		v.DhcpEndAddress = types.StringNull()
	}

	if jsonData.Management.DomainName != "" {
		v.DomainName = types.StringValue(jsonData.Management.DomainName)

	} else {
		v.DomainName = types.StringNull()
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

	if jsonData.Management.BootstrapMultiSubnet != "" {
		v.BootstrapMultiSubnet = types.StringValue(jsonData.Management.BootstrapMultiSubnet)

	} else {
		v.BootstrapMultiSubnet = types.StringNull()
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

	if jsonData.Management.UnderlayIpv6 != nil {
		v.UnderlayIpv6 = types.BoolValue(*jsonData.Management.UnderlayIpv6)

	} else {
		v.UnderlayIpv6 = types.BoolNull()
	}

	if jsonData.Management.Ipv6MulticastGroupSubnet != "" {
		v.Ipv6MulticastGroupSubnet = types.StringValue(jsonData.Management.Ipv6MulticastGroupSubnet)

	} else {
		v.Ipv6MulticastGroupSubnet = types.StringNull()
	}

	if jsonData.Management.TenantRoutedMulticastIpv6 != nil {
		v.TenantRoutedMulticastIpv6 = types.BoolValue(*jsonData.Management.TenantRoutedMulticastIpv6)

	} else {
		v.TenantRoutedMulticastIpv6 = types.BoolNull()
	}

	if jsonData.Management.MvpnVrfRouteImportId != nil {
		v.MvpnVrfRouteImportId = types.BoolValue(*jsonData.Management.MvpnVrfRouteImportId)

	} else {
		v.MvpnVrfRouteImportId = types.BoolNull()
	}

	if jsonData.Management.MvpnVrfRouteImportIdRange != "" {
		v.MvpnVrfRouteImportIdRange = types.StringValue(jsonData.Management.MvpnVrfRouteImportIdRange)

	} else {
		v.MvpnVrfRouteImportIdRange = types.StringNull()
	}

	if jsonData.Management.VrfRouteImportIdReallocation != nil {
		v.VrfRouteImportIdReallocation = types.BoolValue(*jsonData.Management.VrfRouteImportIdReallocation)

	} else {
		v.VrfRouteImportIdReallocation = types.BoolNull()
	}

	if jsonData.Management.L3vniMulticastGroup != "" {
		v.L3vniMulticastGroup = types.StringValue(jsonData.Management.L3vniMulticastGroup)

	} else {
		v.L3vniMulticastGroup = types.StringNull()
	}

	if jsonData.Management.L3VniIpv6MulticastGroup != "" {
		v.L3VniIpv6MulticastGroup = types.StringValue(jsonData.Management.L3VniIpv6MulticastGroup)

	} else {
		v.L3VniIpv6MulticastGroup = types.StringNull()
	}

	if jsonData.Management.RendezvousPointMode != "" {
		v.RendezvousPointMode = types.StringValue(jsonData.Management.RendezvousPointMode)

	} else {
		v.RendezvousPointMode = types.StringNull()
	}

	if jsonData.Management.AutoGenerateMulticastGroupAddress != nil {
		v.AutoGenerateMulticastGroupAddress = types.BoolValue(*jsonData.Management.AutoGenerateMulticastGroupAddress)

	} else {
		v.AutoGenerateMulticastGroupAddress = types.BoolNull()
	}

	if jsonData.Management.PhantomRendezvousPointLoopbackId1 != nil {
		v.PhantomRendezvousPointLoopbackId1 = types.Int64Value(*jsonData.Management.PhantomRendezvousPointLoopbackId1)

	} else {
		v.PhantomRendezvousPointLoopbackId1 = types.Int64Null()
	}

	if jsonData.Management.PhantomRendezvousPointLoopbackId2 != nil {
		v.PhantomRendezvousPointLoopbackId2 = types.Int64Value(*jsonData.Management.PhantomRendezvousPointLoopbackId2)

	} else {
		v.PhantomRendezvousPointLoopbackId2 = types.Int64Null()
	}

	if jsonData.Management.PhantomRendezvousPointLoopbackId3 != nil {
		v.PhantomRendezvousPointLoopbackId3 = types.Int64Value(*jsonData.Management.PhantomRendezvousPointLoopbackId3)

	} else {
		v.PhantomRendezvousPointLoopbackId3 = types.Int64Null()
	}

	if jsonData.Management.PhantomRendezvousPointLoopbackId4 != nil {
		v.PhantomRendezvousPointLoopbackId4 = types.Int64Value(*jsonData.Management.PhantomRendezvousPointLoopbackId4)

	} else {
		v.PhantomRendezvousPointLoopbackId4 = types.Int64Null()
	}

	if jsonData.Management.AdvertisePhysicalIpOnBorder != nil {
		v.AdvertisePhysicalIpOnBorder = types.BoolValue(*jsonData.Management.AdvertisePhysicalIpOnBorder)

	} else {
		v.AdvertisePhysicalIpOnBorder = types.BoolNull()
	}

	if jsonData.Management.FabricVpcDomainId != nil {
		v.FabricVpcDomainId = types.BoolValue(*jsonData.Management.FabricVpcDomainId)

	} else {
		v.FabricVpcDomainId = types.BoolNull()
	}

	if jsonData.Management.SharedVpcDomainId != nil {
		v.SharedVpcDomainId = types.Int64Value(*jsonData.Management.SharedVpcDomainId)

	} else {
		v.SharedVpcDomainId = types.Int64Null()
	}

	if jsonData.Management.VpcLayer3PeerRouter != nil {
		v.VpcLayer3PeerRouter = types.BoolValue(*jsonData.Management.VpcLayer3PeerRouter)

	} else {
		v.VpcLayer3PeerRouter = types.BoolNull()
	}

	if jsonData.Management.FabricVpcQos != nil {
		v.FabricVpcQos = types.BoolValue(*jsonData.Management.FabricVpcQos)

	} else {
		v.FabricVpcQos = types.BoolNull()
	}

	if jsonData.Management.FabricVpcQosPolicyName != "" {
		v.FabricVpcQosPolicyName = types.StringValue(jsonData.Management.FabricVpcQosPolicyName)

	} else {
		v.FabricVpcQosPolicyName = types.StringNull()
	}

	if jsonData.Management.EnablePeerSwitch != nil {
		v.EnablePeerSwitch = types.BoolValue(*jsonData.Management.EnablePeerSwitch)

	} else {
		v.EnablePeerSwitch = types.BoolNull()
	}

	if jsonData.Management.AnycastLoopbackId != nil {
		v.AnycastLoopbackId = types.Int64Value(*jsonData.Management.AnycastLoopbackId)

	} else {
		v.AnycastLoopbackId = types.Int64Null()
	}

	if jsonData.Management.BgpAuthentication != nil {
		v.BgpAuthentication = types.BoolValue(*jsonData.Management.BgpAuthentication)

	} else {
		v.BgpAuthentication = types.BoolNull()
	}

	if jsonData.Management.BgpAuthenticationKeyType != "" {
		v.BgpAuthenticationKeyType = types.StringValue(jsonData.Management.BgpAuthenticationKeyType)

	} else {
		v.BgpAuthenticationKeyType = types.StringNull()
	}

	if jsonData.Management.BgpAuthenticationKey != "" {
		v.BgpAuthenticationKey = types.StringValue(jsonData.Management.BgpAuthenticationKey)

	} else {
		v.BgpAuthenticationKey = types.StringNull()
	}

	if jsonData.Management.PimHelloAuthentication != nil {
		v.PimHelloAuthentication = types.BoolValue(*jsonData.Management.PimHelloAuthentication)

	} else {
		v.PimHelloAuthentication = types.BoolNull()
	}

	if jsonData.Management.PimHelloAuthenticationKey != "" {
		v.PimHelloAuthenticationKey = types.StringValue(jsonData.Management.PimHelloAuthenticationKey)

	} else {
		v.PimHelloAuthenticationKey = types.StringNull()
	}

	if jsonData.Management.Bfd != nil {
		v.Bfd = types.BoolValue(*jsonData.Management.Bfd)

	} else {
		v.Bfd = types.BoolNull()
	}

	if jsonData.Management.BfdIbgp != nil {
		v.BfdIbgp = types.BoolValue(*jsonData.Management.BfdIbgp)

	} else {
		v.BfdIbgp = types.BoolNull()
	}

	if jsonData.Management.BfdAuthentication != nil {
		v.BfdAuthentication = types.BoolValue(*jsonData.Management.BfdAuthentication)

	} else {
		v.BfdAuthentication = types.BoolNull()
	}

	if jsonData.Management.BfdAuthenticationKeyId != nil {
		v.BfdAuthenticationKeyId = types.Int64Value(*jsonData.Management.BfdAuthenticationKeyId)

	} else {
		v.BfdAuthenticationKeyId = types.Int64Null()
	}

	if jsonData.Management.BfdAuthenticationKey != "" {
		v.BfdAuthenticationKey = types.StringValue(jsonData.Management.BfdAuthenticationKey)

	} else {
		v.BfdAuthenticationKey = types.StringNull()
	}

	if jsonData.Management.Macsec != nil {
		v.Macsec = types.BoolValue(*jsonData.Management.Macsec)

	} else {
		v.Macsec = types.BoolNull()
	}

	if jsonData.Management.MacsecCipherSuite != "" {
		v.MacsecCipherSuite = types.StringValue(jsonData.Management.MacsecCipherSuite)

	} else {
		v.MacsecCipherSuite = types.StringNull()
	}

	if jsonData.Management.MacsecKeyString != "" {
		v.MacsecKeyString = types.StringValue(jsonData.Management.MacsecKeyString)

	} else {
		v.MacsecKeyString = types.StringNull()
	}

	if jsonData.Management.MacsecAlgorithm != "" {
		v.MacsecAlgorithm = types.StringValue(jsonData.Management.MacsecAlgorithm)

	} else {
		v.MacsecAlgorithm = types.StringNull()
	}

	if jsonData.Management.MacsecFallbackKeyString != "" {
		v.MacsecFallbackKeyString = types.StringValue(jsonData.Management.MacsecFallbackKeyString)

	} else {
		v.MacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.Management.MacsecFallbackAlgorithm != "" {
		v.MacsecFallbackAlgorithm = types.StringValue(jsonData.Management.MacsecFallbackAlgorithm)

	} else {
		v.MacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.Management.MacsecReportTimer != nil {
		v.MacsecReportTimer = types.Int64Value(*jsonData.Management.MacsecReportTimer)

	} else {
		v.MacsecReportTimer = types.Int64Null()
	}

	if jsonData.Management.MonitoredMode != nil {
		v.MonitoredMode = types.BoolValue(*jsonData.Management.MonitoredMode)

	} else {
		v.MonitoredMode = types.BoolNull()
	}

	if jsonData.Management.OverlayMode != "" {
		v.OverlayMode = types.StringValue(jsonData.Management.OverlayMode)

	} else {
		v.OverlayMode = types.StringNull()
	}

	if jsonData.Management.PrivateVlan != nil {
		v.PrivateVlan = types.BoolValue(*jsonData.Management.PrivateVlan)

	} else {
		v.PrivateVlan = types.BoolNull()
	}

	if jsonData.Management.DefaultPrivateVlanSecondaryNetworkTemplate != "" {
		v.DefaultPrivateVlanSecondaryNetworkTemplate = types.StringValue(jsonData.Management.DefaultPrivateVlanSecondaryNetworkTemplate)

	} else {
		v.DefaultPrivateVlanSecondaryNetworkTemplate = types.StringNull()
	}

	if jsonData.Management.PowerRedundancyMode != "" {
		v.PowerRedundancyMode = types.StringValue(jsonData.Management.PowerRedundancyMode)

	} else {
		v.PowerRedundancyMode = types.StringNull()
	}

	if jsonData.Management.CoppPolicy != "" {
		v.CoppPolicy = types.StringValue(jsonData.Management.CoppPolicy)

	} else {
		v.CoppPolicy = types.StringNull()
	}

	if jsonData.Management.NveHoldDownTimer != nil {
		v.NveHoldDownTimer = types.Int64Value(*jsonData.Management.NveHoldDownTimer)

	} else {
		v.NveHoldDownTimer = types.Int64Null()
	}

	if jsonData.Management.Cdp != nil {
		v.Cdp = types.BoolValue(*jsonData.Management.Cdp)

	} else {
		v.Cdp = types.BoolNull()
	}

	if jsonData.Management.NextGenerationOam != nil {
		v.NextGenerationOam = types.BoolValue(*jsonData.Management.NextGenerationOam)

	} else {
		v.NextGenerationOam = types.BoolNull()
	}

	if jsonData.Management.NgoamSouthBoundLoopDetect != nil {
		v.NgoamSouthBoundLoopDetect = types.BoolValue(*jsonData.Management.NgoamSouthBoundLoopDetect)

	} else {
		v.NgoamSouthBoundLoopDetect = types.BoolNull()
	}

	if jsonData.Management.NgoamSouthBoundLoopDetectProbeInterval != nil {
		v.NgoamSouthBoundLoopDetectProbeInterval = types.Int64Value(*jsonData.Management.NgoamSouthBoundLoopDetectProbeInterval)

	} else {
		v.NgoamSouthBoundLoopDetectProbeInterval = types.Int64Null()
	}

	if jsonData.Management.NgoamSouthBoundLoopDetectRecoveryInterval != nil {
		v.NgoamSouthBoundLoopDetectRecoveryInterval = types.Int64Value(*jsonData.Management.NgoamSouthBoundLoopDetectRecoveryInterval)

	} else {
		v.NgoamSouthBoundLoopDetectRecoveryInterval = types.Int64Null()
	}

	if jsonData.Management.StrictConfigComplianceMode != nil {
		v.StrictConfigComplianceMode = types.BoolValue(*jsonData.Management.StrictConfigComplianceMode)

	} else {
		v.StrictConfigComplianceMode = types.BoolNull()
	}

	if jsonData.Management.AdvancedSshOption != nil {
		v.AdvancedSshOption = types.BoolValue(*jsonData.Management.AdvancedSshOption)

	} else {
		v.AdvancedSshOption = types.BoolNull()
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

	if jsonData.Management.DefaultQueuingPolicy != nil {
		v.DefaultQueuingPolicy = types.BoolValue(*jsonData.Management.DefaultQueuingPolicy)

	} else {
		v.DefaultQueuingPolicy = types.BoolNull()
	}

	if jsonData.Management.DefaultQueuingPolicyCloudscale != "" {
		v.DefaultQueuingPolicyCloudscale = types.StringValue(jsonData.Management.DefaultQueuingPolicyCloudscale)

	} else {
		v.DefaultQueuingPolicyCloudscale = types.StringNull()
	}

	if jsonData.Management.DefaultQueuingPolicyRSeries != "" {
		v.DefaultQueuingPolicyRSeries = types.StringValue(jsonData.Management.DefaultQueuingPolicyRSeries)

	} else {
		v.DefaultQueuingPolicyRSeries = types.StringNull()
	}

	if jsonData.Management.DefaultQueuingPolicyOther != "" {
		v.DefaultQueuingPolicyOther = types.StringValue(jsonData.Management.DefaultQueuingPolicyOther)

	} else {
		v.DefaultQueuingPolicyOther = types.StringNull()
	}

	if jsonData.Management.AimlQos != nil {
		v.AimlQos = types.BoolValue(*jsonData.Management.AimlQos)

	} else {
		v.AimlQos = types.BoolNull()
	}

	if jsonData.Management.AimlQosPolicy != "" {
		v.AimlQosPolicy = types.StringValue(jsonData.Management.AimlQosPolicy)

	} else {
		v.AimlQosPolicy = types.StringNull()
	}

	if jsonData.Management.PriorityFlowControlWatchInterval != nil {
		v.PriorityFlowControlWatchInterval = types.Int64Value(*jsonData.Management.PriorityFlowControlWatchInterval)

	} else {
		v.PriorityFlowControlWatchInterval = types.Int64Null()
	}

	if jsonData.Management.Dlb != nil {
		v.Dlb = types.BoolValue(*jsonData.Management.Dlb)

	} else {
		v.Dlb = types.BoolNull()
	}

	if jsonData.Management.DlbMode != "" {
		v.DlbMode = types.StringValue(jsonData.Management.DlbMode)

	} else {
		v.DlbMode = types.StringNull()
	}

	if jsonData.Management.DlbMixedModeDefault != "" {
		v.DlbMixedModeDefault = types.StringValue(jsonData.Management.DlbMixedModeDefault)

	} else {
		v.DlbMixedModeDefault = types.StringNull()
	}

	if jsonData.Management.FlowletAging != nil {
		v.FlowletAging = types.Int64Value(*jsonData.Management.FlowletAging)

	} else {
		v.FlowletAging = types.Int64Null()
	}

	if jsonData.Management.FlowletDscp != "" {
		v.FlowletDscp = types.StringValue(jsonData.Management.FlowletDscp)

	} else {
		v.FlowletDscp = types.StringNull()
	}

	if jsonData.Management.PerPacketDscp != "" {
		v.PerPacketDscp = types.StringValue(jsonData.Management.PerPacketDscp)

	} else {
		v.PerPacketDscp = types.StringNull()
	}

	if jsonData.Management.AiLoadSharing != nil {
		v.AiLoadSharing = types.BoolValue(*jsonData.Management.AiLoadSharing)

	} else {
		v.AiLoadSharing = types.BoolNull()
	}

	if jsonData.Management.RoceV2 != "" {
		v.RoceV2 = types.StringValue(jsonData.Management.RoceV2)

	} else {
		v.RoceV2 = types.StringNull()
	}

	if jsonData.Management.Cnp != "" {
		v.Cnp = types.StringValue(jsonData.Management.Cnp)

	} else {
		v.Cnp = types.StringNull()
	}

	if jsonData.Management.WredMin != nil {
		v.WredMin = types.Int64Value(*jsonData.Management.WredMin)

	} else {
		v.WredMin = types.Int64Null()
	}

	if jsonData.Management.WredMax != nil {
		v.WredMax = types.Int64Value(*jsonData.Management.WredMax)

	} else {
		v.WredMax = types.Int64Null()
	}

	if jsonData.Management.WredDropProbability != nil {
		v.WredDropProbability = types.Int64Value(*jsonData.Management.WredDropProbability)

	} else {
		v.WredDropProbability = types.Int64Null()
	}

	if jsonData.Management.WredWeight != nil {
		v.WredWeight = types.Int64Value(*jsonData.Management.WredWeight)

	} else {
		v.WredWeight = types.Int64Null()
	}

	if jsonData.Management.BandwidthRemaining != nil {
		v.BandwidthRemaining = types.Int64Value(*jsonData.Management.BandwidthRemaining)

	} else {
		v.BandwidthRemaining = types.Int64Null()
	}

	if jsonData.Management.StaticUnderlayIpAllocation != nil {
		v.StaticUnderlayIpAllocation = types.BoolValue(*jsonData.Management.StaticUnderlayIpAllocation)

	} else {
		v.StaticUnderlayIpAllocation = types.BoolNull()
	}

	if jsonData.Management.BgpLoopbackIpv6Range != "" {
		v.BgpLoopbackIpv6Range = types.StringValue(jsonData.Management.BgpLoopbackIpv6Range)

	} else {
		v.BgpLoopbackIpv6Range = types.StringNull()
	}

	if jsonData.Management.NveLoopbackIpv6Range != "" {
		v.NveLoopbackIpv6Range = types.StringValue(jsonData.Management.NveLoopbackIpv6Range)

	} else {
		v.NveLoopbackIpv6Range = types.StringNull()
	}

	if jsonData.Management.Ipv6AnycastRendezvousPointIpRange != "" {
		v.Ipv6AnycastRendezvousPointIpRange = types.StringValue(jsonData.Management.Ipv6AnycastRendezvousPointIpRange)

	} else {
		v.Ipv6AnycastRendezvousPointIpRange = types.StringNull()
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

	if jsonData.Management.Aaa != nil {
		v.Aaa = types.BoolValue(*jsonData.Management.Aaa)

	} else {
		v.Aaa = types.BoolNull()
	}

	if jsonData.Management.Ipv6LinkLocal != nil {
		v.Ipv6LinkLocal = types.BoolValue(*jsonData.Management.Ipv6LinkLocal)

	} else {
		v.Ipv6LinkLocal = types.BoolNull()
	}

	if jsonData.Management.FabricInterfaceType != "" {
		v.FabricInterfaceType = types.StringValue(jsonData.Management.FabricInterfaceType)

	} else {
		v.FabricInterfaceType = types.StringNull()
	}

	if jsonData.Management.Ipv6SubnetTargetMask != nil {
		v.Ipv6SubnetTargetMask = types.Int64Value(*jsonData.Management.Ipv6SubnetTargetMask)

	} else {
		v.Ipv6SubnetTargetMask = types.Int64Null()
	}

	if jsonData.Management.LinkStateRoutingProtocol != "" {
		v.LinkStateRoutingProtocol = types.StringValue(jsonData.Management.LinkStateRoutingProtocol)

	} else {
		v.LinkStateRoutingProtocol = types.StringNull()
	}

	if jsonData.Management.RouteReflectorCount != nil {
		v.RouteReflectorCount = types.Int64Value(*jsonData.Management.RouteReflectorCount)

	} else {
		v.RouteReflectorCount = types.Int64Null()
	}

	if jsonData.Management.VpcTorDelayRestoreTimer != nil {
		v.VpcTorDelayRestoreTimer = types.Int64Value(*jsonData.Management.VpcTorDelayRestoreTimer)

	} else {
		v.VpcTorDelayRestoreTimer = types.Int64Null()
	}

	if jsonData.Management.LeafTorIdRange != nil {
		v.LeafTorIdRange = types.BoolValue(*jsonData.Management.LeafTorIdRange)

	} else {
		v.LeafTorIdRange = types.BoolNull()
	}

	if jsonData.Management.LeafTorVpcPortChannelIdRange != "" {
		v.LeafTorVpcPortChannelIdRange = types.StringValue(jsonData.Management.LeafTorVpcPortChannelIdRange)

	} else {
		v.LeafTorVpcPortChannelIdRange = types.StringNull()
	}

	if jsonData.Management.LinkStateRoutingTag != "" {
		v.LinkStateRoutingTag = types.StringValue(jsonData.Management.LinkStateRoutingTag)

	} else {
		v.LinkStateRoutingTag = types.StringNull()
	}

	if jsonData.Management.OspfAreaId != "" {
		v.OspfAreaId = types.StringValue(jsonData.Management.OspfAreaId)

	} else {
		v.OspfAreaId = types.StringNull()
	}

	if jsonData.Management.OspfAuthentication != nil {
		v.OspfAuthentication = types.BoolValue(*jsonData.Management.OspfAuthentication)

	} else {
		v.OspfAuthentication = types.BoolNull()
	}

	if jsonData.Management.OspfAuthenticationKeyId != nil {
		v.OspfAuthenticationKeyId = types.Int64Value(*jsonData.Management.OspfAuthenticationKeyId)

	} else {
		v.OspfAuthenticationKeyId = types.Int64Null()
	}

	if jsonData.Management.OspfAuthenticationKey != "" {
		v.OspfAuthenticationKey = types.StringValue(jsonData.Management.OspfAuthenticationKey)

	} else {
		v.OspfAuthenticationKey = types.StringNull()
	}

	if jsonData.Management.IsisLevel != "" {
		v.IsisLevel = types.StringValue(jsonData.Management.IsisLevel)

	} else {
		v.IsisLevel = types.StringNull()
	}

	if jsonData.Management.IsisAreaNumber != "" {
		v.IsisAreaNumber = types.StringValue(jsonData.Management.IsisAreaNumber)

	} else {
		v.IsisAreaNumber = types.StringNull()
	}

	if jsonData.Management.IsisPointToPoint != nil {
		v.IsisPointToPoint = types.BoolValue(*jsonData.Management.IsisPointToPoint)

	} else {
		v.IsisPointToPoint = types.BoolNull()
	}

	if jsonData.Management.IsisAuthentication != nil {
		v.IsisAuthentication = types.BoolValue(*jsonData.Management.IsisAuthentication)

	} else {
		v.IsisAuthentication = types.BoolNull()
	}

	if jsonData.Management.IsisAuthenticationKeychainName != "" {
		v.IsisAuthenticationKeychainName = types.StringValue(jsonData.Management.IsisAuthenticationKeychainName)

	} else {
		v.IsisAuthenticationKeychainName = types.StringNull()
	}

	if jsonData.Management.IsisAuthenticationKeychainKeyId != nil {
		v.IsisAuthenticationKeychainKeyId = types.Int64Value(*jsonData.Management.IsisAuthenticationKeychainKeyId)

	} else {
		v.IsisAuthenticationKeychainKeyId = types.Int64Null()
	}

	if jsonData.Management.IsisAuthenticationKey != "" {
		v.IsisAuthenticationKey = types.StringValue(jsonData.Management.IsisAuthenticationKey)

	} else {
		v.IsisAuthenticationKey = types.StringNull()
	}

	if jsonData.Management.IsisOverload != nil {
		v.IsisOverload = types.BoolValue(*jsonData.Management.IsisOverload)

	} else {
		v.IsisOverload = types.BoolNull()
	}

	if jsonData.Management.IsisOverloadElapseTime != nil {
		v.IsisOverloadElapseTime = types.Int64Value(*jsonData.Management.IsisOverloadElapseTime)

	} else {
		v.IsisOverloadElapseTime = types.Int64Null()
	}

	if jsonData.Management.BfdOspf != nil {
		v.BfdOspf = types.BoolValue(*jsonData.Management.BfdOspf)

	} else {
		v.BfdOspf = types.BoolNull()
	}

	if jsonData.Management.BfdIsis != nil {
		v.BfdIsis = types.BoolValue(*jsonData.Management.BfdIsis)

	} else {
		v.BfdIsis = types.BoolNull()
	}

	if jsonData.Management.BfdPim != nil {
		v.BfdPim = types.BoolValue(*jsonData.Management.BfdPim)

	} else {
		v.BfdPim = types.BoolNull()
	}

	if jsonData.Management.AutoBgpNeighborDescription != nil {
		v.AutoBgpNeighborDescription = types.BoolValue(*jsonData.Management.AutoBgpNeighborDescription)

	} else {
		v.AutoBgpNeighborDescription = types.BoolNull()
	}

	if jsonData.Management.IbgpPeerTemplate != "" {
		v.IbgpPeerTemplate = types.StringValue(jsonData.Management.IbgpPeerTemplate)

	} else {
		v.IbgpPeerTemplate = types.StringNull()
	}

	if jsonData.Management.LeafibgpPeerTemplate != "" {
		v.LeafibgpPeerTemplate = types.StringValue(jsonData.Management.LeafibgpPeerTemplate)

	} else {
		v.LeafibgpPeerTemplate = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTag != nil {
		v.SecurityGroupTag = types.BoolValue(*jsonData.Management.SecurityGroupTag)

	} else {
		v.SecurityGroupTag = types.BoolNull()
	}

	if jsonData.Management.SecurityGroupTagPrefix != "" {
		v.SecurityGroupTagPrefix = types.StringValue(jsonData.Management.SecurityGroupTagPrefix)

	} else {
		v.SecurityGroupTagPrefix = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagIdRange != "" {
		v.SecurityGroupTagIdRange = types.StringValue(jsonData.Management.SecurityGroupTagIdRange)

	} else {
		v.SecurityGroupTagIdRange = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagPreprovision != nil {
		v.SecurityGroupTagPreprovision = types.BoolValue(*jsonData.Management.SecurityGroupTagPreprovision)

	} else {
		v.SecurityGroupTagPreprovision = types.BoolNull()
	}

	if jsonData.Management.SecurityGroupTagMacSegmentation != nil {
		v.SecurityGroupTagMacSegmentation = types.BoolValue(*jsonData.Management.SecurityGroupTagMacSegmentation)

	} else {
		v.SecurityGroupTagMacSegmentation = types.BoolNull()
	}

	if jsonData.Management.SecurityGroupStatus != "" {
		v.SecurityGroupStatus = types.StringValue(jsonData.Management.SecurityGroupStatus)

	} else {
		v.SecurityGroupStatus = types.StringNull()
	}

	if jsonData.Management.VrfLiteMacsec != nil {
		v.VrfLiteMacsec = types.BoolValue(*jsonData.Management.VrfLiteMacsec)

	} else {
		v.VrfLiteMacsec = types.BoolNull()
	}

	if jsonData.Management.QuantumKeyDistribution != nil {
		v.QuantumKeyDistribution = types.BoolValue(*jsonData.Management.QuantumKeyDistribution)

	} else {
		v.QuantumKeyDistribution = types.BoolNull()
	}

	if jsonData.Management.VrfLiteMacsecCipherSuite != "" {
		v.VrfLiteMacsecCipherSuite = types.StringValue(jsonData.Management.VrfLiteMacsecCipherSuite)

	} else {
		v.VrfLiteMacsecCipherSuite = types.StringNull()
	}

	if jsonData.Management.VrfLiteMacsecKeyString != "" {
		v.VrfLiteMacsecKeyString = types.StringValue(jsonData.Management.VrfLiteMacsecKeyString)

	} else {
		v.VrfLiteMacsecKeyString = types.StringNull()
	}

	if jsonData.Management.VrfLiteMacsecAlgorithm != "" {
		v.VrfLiteMacsecAlgorithm = types.StringValue(jsonData.Management.VrfLiteMacsecAlgorithm)

	} else {
		v.VrfLiteMacsecAlgorithm = types.StringNull()
	}

	if jsonData.Management.VrfLiteMacsecFallbackKeyString != "" {
		v.VrfLiteMacsecFallbackKeyString = types.StringValue(jsonData.Management.VrfLiteMacsecFallbackKeyString)

	} else {
		v.VrfLiteMacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.Management.VrfLiteMacsecFallbackAlgorithm != "" {
		v.VrfLiteMacsecFallbackAlgorithm = types.StringValue(jsonData.Management.VrfLiteMacsecFallbackAlgorithm)

	} else {
		v.VrfLiteMacsecFallbackAlgorithm = types.StringNull()
	}

	if jsonData.Management.QuantumKeyDistributionProfileName != "" {
		v.QuantumKeyDistributionProfileName = types.StringValue(jsonData.Management.QuantumKeyDistributionProfileName)

	} else {
		v.QuantumKeyDistributionProfileName = types.StringNull()
	}

	if jsonData.Management.KeyManagementEntityServerIp != "" {
		v.KeyManagementEntityServerIp = types.StringValue(jsonData.Management.KeyManagementEntityServerIp)

	} else {
		v.KeyManagementEntityServerIp = types.StringNull()
	}

	if jsonData.Management.KeyManagementEntityServerPort != nil {
		v.KeyManagementEntityServerPort = types.Int64Value(*jsonData.Management.KeyManagementEntityServerPort)

	} else {
		v.KeyManagementEntityServerPort = types.Int64Null()
	}

	if jsonData.Management.TrustpointLabel != "" {
		v.TrustpointLabel = types.StringValue(jsonData.Management.TrustpointLabel)

	} else {
		v.TrustpointLabel = types.StringNull()
	}

	if jsonData.Management.SkipCertificateVerification != nil {
		v.SkipCertificateVerification = types.BoolValue(*jsonData.Management.SkipCertificateVerification)

	} else {
		v.SkipCertificateVerification = types.BoolNull()
	}

	if jsonData.Management.HostInterfaceAdminState != nil {
		v.HostInterfaceAdminState = types.BoolValue(*jsonData.Management.HostInterfaceAdminState)

	} else {
		v.HostInterfaceAdminState = types.BoolNull()
	}

	if jsonData.Management.BrownfieldNetworkNameFormat != "" {
		v.BrownfieldNetworkNameFormat = types.StringValue(jsonData.Management.BrownfieldNetworkNameFormat)

	} else {
		v.BrownfieldNetworkNameFormat = types.StringNull()
	}

	if jsonData.Management.BrownfieldSkipOverlayNetworkAttachments != nil {
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolValue(*jsonData.Management.BrownfieldSkipOverlayNetworkAttachments)

	} else {
		v.BrownfieldSkipOverlayNetworkAttachments = types.BoolNull()
	}

	if jsonData.Management.PolicyBasedRouting != nil {
		v.PolicyBasedRouting = types.BoolValue(*jsonData.Management.PolicyBasedRouting)

	} else {
		v.PolicyBasedRouting = types.BoolNull()
	}

	if jsonData.Management.PtpVlanId != nil {
		v.PtpVlanId = types.Int64Value(*jsonData.Management.PtpVlanId)

	} else {
		v.PtpVlanId = types.Int64Null()
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

	if jsonData.Management.MplsIsisAreaNumber != "" {
		v.MplsIsisAreaNumber = types.StringValue(jsonData.Management.MplsIsisAreaNumber)

	} else {
		v.MplsIsisAreaNumber = types.StringNull()
	}

	if jsonData.Management.StpRootOption != "" {
		v.StpRootOption = types.StringValue(jsonData.Management.StpRootOption)

	} else {
		v.StpRootOption = types.StringNull()
	}

	if jsonData.Management.StpVlanRange != "" {
		v.StpVlanRange = types.StringValue(jsonData.Management.StpVlanRange)

	} else {
		v.StpVlanRange = types.StringNull()
	}

	if jsonData.Management.MstInstanceRange != "" {
		v.MstInstanceRange = types.StringValue(jsonData.Management.MstInstanceRange)

	} else {
		v.MstInstanceRange = types.StringNull()
	}

	if jsonData.Management.StpBridgePriority != nil {
		v.StpBridgePriority = types.Int64Value(*jsonData.Management.StpBridgePriority)

	} else {
		v.StpBridgePriority = types.Int64Null()
	}

	if jsonData.Management.AllowVlanOnLeafTorPairing != "" {
		v.AllowVlanOnLeafTorPairing = types.StringValue(jsonData.Management.AllowVlanOnLeafTorPairing)

	} else {
		v.AllowVlanOnLeafTorPairing = types.StringNull()
	}

	if jsonData.Management.PreInterfaceConfigLeaf != "" {
		v.PreInterfaceConfigLeaf = types.StringValue(jsonData.Management.PreInterfaceConfigLeaf)

	} else {
		v.PreInterfaceConfigLeaf = types.StringNull()
	}

	if jsonData.Management.PreInterfaceConfigSpine != "" {
		v.PreInterfaceConfigSpine = types.StringValue(jsonData.Management.PreInterfaceConfigSpine)

	} else {
		v.PreInterfaceConfigSpine = types.StringNull()
	}

	if jsonData.Management.PreInterfaceConfigTor != "" {
		v.PreInterfaceConfigTor = types.StringValue(jsonData.Management.PreInterfaceConfigTor)

	} else {
		v.PreInterfaceConfigTor = types.StringNull()
	}

	if jsonData.Management.ExtraConfigLeaf != "" {
		v.ExtraConfigLeaf = types.StringValue(jsonData.Management.ExtraConfigLeaf)

	} else {
		v.ExtraConfigLeaf = types.StringNull()
	}

	if jsonData.Management.ExtraConfigSpine != "" {
		v.ExtraConfigSpine = types.StringValue(jsonData.Management.ExtraConfigSpine)

	} else {
		v.ExtraConfigSpine = types.StringNull()
	}

	if jsonData.Management.ExtraConfigTor != "" {
		v.ExtraConfigTor = types.StringValue(jsonData.Management.ExtraConfigTor)

	} else {
		v.ExtraConfigTor = types.StringNull()
	}

	if jsonData.Management.ExtraConfigIntraFabricLinks != "" {
		v.ExtraConfigIntraFabricLinks = types.StringValue(jsonData.Management.ExtraConfigIntraFabricLinks)

	} else {
		v.ExtraConfigIntraFabricLinks = types.StringNull()
	}

	if jsonData.Management.MplsLoopbackIpRange != "" {
		v.MplsLoopbackIpRange = types.StringValue(jsonData.Management.MplsLoopbackIpRange)

	} else {
		v.MplsLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.Ipv6SubnetRange != "" {
		v.Ipv6SubnetRange = types.StringValue(jsonData.Management.Ipv6SubnetRange)

	} else {
		v.Ipv6SubnetRange = types.StringNull()
	}

	if jsonData.Management.RouterIdRange != "" {
		v.RouterIdRange = types.StringValue(jsonData.Management.RouterIdRange)

	} else {
		v.RouterIdRange = types.StringNull()
	}

	if jsonData.Management.AutoSymmetricVrfLite != nil {
		v.AutoSymmetricVrfLite = types.BoolValue(*jsonData.Management.AutoSymmetricVrfLite)

	} else {
		v.AutoSymmetricVrfLite = types.BoolNull()
	}

	if jsonData.Management.AutoVrfLiteDefaultVrf != nil {
		v.AutoVrfLiteDefaultVrf = types.BoolValue(*jsonData.Management.AutoVrfLiteDefaultVrf)

	} else {
		v.AutoVrfLiteDefaultVrf = types.BoolNull()
	}

	if jsonData.Management.AutoSymmetricDefaultVrf != nil {
		v.AutoSymmetricDefaultVrf = types.BoolValue(*jsonData.Management.AutoSymmetricDefaultVrf)

	} else {
		v.AutoSymmetricDefaultVrf = types.BoolNull()
	}

	if jsonData.Management.DefaultVrfRedistributionBgpRouteMap != "" {
		v.DefaultVrfRedistributionBgpRouteMap = types.StringValue(jsonData.Management.DefaultVrfRedistributionBgpRouteMap)

	} else {
		v.DefaultVrfRedistributionBgpRouteMap = types.StringNull()
	}

	if jsonData.Management.IpServiceLevelAgreementIdRange != "" {
		v.IpServiceLevelAgreementIdRange = types.StringValue(jsonData.Management.IpServiceLevelAgreementIdRange)

	} else {
		v.IpServiceLevelAgreementIdRange = types.StringNull()
	}

	if jsonData.Management.ObjectTrackingNumberRange != "" {
		v.ObjectTrackingNumberRange = types.StringValue(jsonData.Management.ObjectTrackingNumberRange)

	} else {
		v.ObjectTrackingNumberRange = types.StringNull()
	}

	if jsonData.Management.ServiceNetworkVlanRange != "" {
		v.ServiceNetworkVlanRange = types.StringValue(jsonData.Management.ServiceNetworkVlanRange)

	} else {
		v.ServiceNetworkVlanRange = types.StringNull()
	}

	if jsonData.Management.RouteMapSequenceNumberRange != "" {
		v.RouteMapSequenceNumberRange = types.StringValue(jsonData.Management.RouteMapSequenceNumberRange)

	} else {
		v.RouteMapSequenceNumberRange = types.StringNull()
	}

	if jsonData.Management.InbandManagement != nil {
		v.InbandManagement = types.BoolValue(*jsonData.Management.InbandManagement)

	} else {
		v.InbandManagement = types.BoolNull()
	}

	if jsonData.Management.SeedSwitchCoreInterfaces != "" {
		v.SeedSwitchCoreInterfaces = types.StringValue(jsonData.Management.SeedSwitchCoreInterfaces)

	} else {
		v.SeedSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.Management.SpineSwitchCoreInterfaces != "" {
		v.SpineSwitchCoreInterfaces = types.StringValue(jsonData.Management.SpineSwitchCoreInterfaces)

	} else {
		v.SpineSwitchCoreInterfaces = types.StringNull()
	}

	if jsonData.Management.InbandDhcpServers != "" {
		v.InbandDhcpServers = types.StringValue(jsonData.Management.InbandDhcpServers)

	} else {
		v.InbandDhcpServers = types.StringNull()
	}

	if jsonData.Management.UnNumberedBootstrapLbId != nil {
		v.UnNumberedBootstrapLbId = types.Int64Value(*jsonData.Management.UnNumberedBootstrapLbId)

	} else {
		v.UnNumberedBootstrapLbId = types.Int64Null()
	}

	if jsonData.Management.UnNumberedDhcpStartAddress != "" {
		v.UnNumberedDhcpStartAddress = types.StringValue(jsonData.Management.UnNumberedDhcpStartAddress)

	} else {
		v.UnNumberedDhcpStartAddress = types.StringNull()
	}

	if jsonData.Management.UnNumberedDhcpEndAddress != "" {
		v.UnNumberedDhcpEndAddress = types.StringValue(jsonData.Management.UnNumberedDhcpEndAddress)

	} else {
		v.UnNumberedDhcpEndAddress = types.StringNull()
	}

	if jsonData.Management.HeartbeatInterval != nil {
		v.HeartbeatInterval = types.Int64Value(*jsonData.Management.HeartbeatInterval)

	} else {
		v.HeartbeatInterval = types.Int64Null()
	}

	if jsonData.Management.AllowSmartSwitchOnboarding != nil {
		v.AllowSmartSwitchOnboarding = types.BoolValue(*jsonData.Management.AllowSmartSwitchOnboarding)

	} else {
		v.AllowSmartSwitchOnboarding = types.BoolNull()
	}

	if jsonData.Management.EnableDpuPinning != nil {
		v.EnableDpuPinning = types.BoolValue(*jsonData.Management.EnableDpuPinning)

	} else {
		v.EnableDpuPinning = types.BoolNull()
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

	if len(jsonData.Management.NtpServerCollection) == 0 {
		log.Printf("v.NtpServerCollection is empty")
		v.NtpServerCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.NtpServerCollection))
		for i, item := range jsonData.Management.NtpServerCollection {
			listData[i] = types.StringValue(item)
		}
		v.NtpServerCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.Management.NtpServerVrfCollection) == 0 {
		log.Printf("v.NtpServerVrfCollection is empty")
		v.NtpServerVrfCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.NtpServerVrfCollection))
		for i, item := range jsonData.Management.NtpServerVrfCollection {
			listData[i] = types.StringValue(item)
		}
		v.NtpServerVrfCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.Management.SyslogServerCollection) == 0 {
		log.Printf("v.SyslogServerCollection is empty")
		v.SyslogServerCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.SyslogServerCollection))
		for i, item := range jsonData.Management.SyslogServerCollection {
			listData[i] = types.StringValue(item)
		}
		v.SyslogServerCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if len(jsonData.Management.SyslogServerVrfCollection) == 0 {
		log.Printf("v.SyslogServerVrfCollection is empty")
		v.SyslogServerVrfCollection = types.SetNull(types.StringType)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.SyslogServerVrfCollection))
		for i, item := range jsonData.Management.SyslogServerVrfCollection {
			listData[i] = types.StringValue(item)
		}
		v.SyslogServerVrfCollection, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
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
	if len(jsonData.Management.NetflowSettings.NetflowSamplerCollection) == 0 {
		log.Printf("v.NetflowSamplerCollection is empty")
		v.NetflowSamplerCollection = types.ListNull(NetflowSamplerCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]NetflowSamplerCollectionValue, len(jsonData.Management.NetflowSettings.NetflowSamplerCollection))
		for i, item := range jsonData.Management.NetflowSettings.NetflowSamplerCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.NetflowSamplerCollection, err = types.ListValueFrom(context.Background(), NetflowSamplerCollectionValue{}.Type(context.Background()), listData)

		if err != nil {
			return err
		}
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

func (v *NetflowSamplerCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCNetflowSamplerCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SamplerName != "" {
		v.SamplerName = types.StringValue(jsonData.SamplerName)
	} else {
		v.SamplerName = types.StringNull()
	}

	if jsonData.NumSamples != nil {
		v.NumSamples = types.Int64Value(*jsonData.NumSamples)

	} else {
		v.NumSamples = types.Int64Null()
	}

	if jsonData.SamplingRate != nil {
		v.SamplingRate = types.Int64Value(*jsonData.SamplingRate)

	} else {
		v.SamplingRate = types.Int64Null()
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

func (v FabricModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	return data
}
