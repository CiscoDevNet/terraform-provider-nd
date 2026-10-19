// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_ethernet

import (
	"context"
	"log"
	"terraform-provider-nd/internal/manage/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v *InterfaceEthernetModel) SetModelData(jsonData *resource_interface_common.NDFCInterfaceCommonModel) diag.Diagnostics {
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

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)
	if len(jsonData.Interfaces) == 0 {
		log.Printf("v.Interfaces is empty")
		v.Interfaces = types.MapNull(InterfacesValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]InterfacesValue)
		for key, item := range jsonData.Interfaces {
			data := new(InterfacesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in InterfacesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.Interfaces, err = types.MapValueFrom(context.Background(), InterfacesValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]InterfacesValue to  Map")

		}
	}

	return err
}

func (v *InterfacesValue) SetValue(jsonData *resource_interface_common.NDFCInterfacesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	if jsonData.ConfigData.Mode != "" {
		v.Mode = types.StringValue(jsonData.ConfigData.Mode)
	} else {
		v.Mode = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.NetworkOsType != "" {
		v.NetworkOsType = types.StringValue(jsonData.ConfigData.NetworkOS.NetworkOsType)
	} else {
		v.NetworkOsType = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PolicyType != "" {
		v.PolicyType = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PolicyType)
	} else {
		v.PolicyType = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.TemplateName != "" {
		v.TemplateName = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.TemplateName)
	} else {
		v.TemplateName = types.StringNull()
	}

	if len(jsonData.ConfigData.NetworkOS.Policy.TemplateConfig) == 0 {
		log.Printf("v.TemplateConfig is empty")
		v.TemplateConfig = types.MapNull(types.StringType)
	} else {
		mapData := make(map[string]attr.Value, len(jsonData.ConfigData.NetworkOS.Policy.TemplateConfig))
		for key, item := range jsonData.ConfigData.NetworkOS.Policy.TemplateConfig {
			mapData[key] = types.StringValue(item)
		}
		v.TemplateConfig, err = types.MapValue(types.StringType, mapData)
		if err != nil {
			log.Printf("Error in converting map[string]string to Map")
			return err
		}
	}

	if jsonData.ConfigData.NetworkOS.Policy.AdminState != nil {
		v.AdminState = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.AdminState)
	} else {
		v.AdminState = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Description != "" {
		v.Description = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Description)
	} else {
		v.Description = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Mtu != "" {
		v.Mtu = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Mtu)
	} else {
		v.Mtu = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.ExtraConfig != "" {
		v.ExtraConfig = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.ExtraConfig)
	} else {
		v.ExtraConfig = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		v.AccessVlan = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.AccessVlan)
	} else {
		v.AccessVlan = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.AllowedVlans != "" {
		v.AllowedVlans = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.AllowedVlans)
	} else {
		v.AllowedVlans = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		v.Bandwidth = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.Bandwidth)
	} else {
		v.Bandwidth = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		v.BandwidthCapacityPercent = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
	} else {
		v.BandwidthCapacityPercent = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		v.BorderRouter = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.BorderRouter)
	} else {
		v.BorderRouter = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.BpduFilter != "" {
		v.BpduFilter = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.BpduFilter)
	} else {
		v.BpduFilter = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.BpduGuard != "" {
		v.BpduGuard = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.BpduGuard)
	} else {
		v.BpduGuard = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Cdp != nil {
		v.Cdp = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Cdp)
	} else {
		v.Cdp = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		v.DebounceLinkupTimer = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
	} else {
		v.DebounceLinkupTimer = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		v.DebounceTimer = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.DebounceTimer)
	} else {
		v.DebounceTimer = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.DuplexMode != "" {
		v.DuplexMode = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.DuplexMode)
	} else {
		v.DuplexMode = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		v.ErrorDetectionAcl = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
	} else {
		v.ErrorDetectionAcl = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Fec != "" {
		v.Fec = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Fec)
	} else {
		v.Fec = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.GuardMode != "" {
		v.GuardMode = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.GuardMode)
	} else {
		v.GuardMode = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.IfName != "" {
		v.IfName = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.IfName)
	} else {
		v.IfName = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		v.InheritBandwidth = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.InheritBandwidth)
	} else {
		v.InheritBandwidth = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Ip != "" {
		v.Ip = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Ip)
	} else {
		v.Ip = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		v.IpRedirects = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.IpRedirects)
	} else {
		v.IpRedirects = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Ipv6 != "" {
		v.Ipv6 = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Ipv6)
	} else {
		v.Ipv6 = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		v.Ipv6Prefix = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.Ipv6Prefix)
	} else {
		v.Ipv6Prefix = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.LinkType != "" {
		v.LinkType = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.LinkType)
	} else {
		v.LinkType = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Mask != "" {
		v.Mask = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Mask)
	} else {
		v.Mask = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Monitor != nil {
		v.Monitor = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Monitor)
	} else {
		v.Monitor = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		v.NativeVlan = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.NativeVlan)
	} else {
		v.NativeVlan = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		v.NegotiateAuto = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.NegotiateAuto)
	} else {
		v.NegotiateAuto = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Netflow != nil {
		v.Netflow = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Netflow)
	} else {
		v.Netflow = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.NetflowMonitor != "" {
		v.NetflowMonitor = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.NetflowMonitor)
	} else {
		v.NetflowMonitor = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.NetflowSampler != "" {
		v.NetflowSampler = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.NetflowSampler)
	} else {
		v.NetflowSampler = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		v.OrphanPort = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.OrphanPort)
	} else {
		v.OrphanPort = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PathCost != nil {
		v.PathCost = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.PathCost)
	} else {
		v.PathCost = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Pfc != nil {
		v.Pfc = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Pfc)
	} else {
		v.Pfc = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		v.PimDrPriority = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.PimDrPriority)
	} else {
		v.PimDrPriority = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PimSparse != nil {
		v.PimSparse = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.PimSparse)
	} else {
		v.PimSparse = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PortChannelId != "" {
		v.PortChannelId = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PortChannelId)
	} else {
		v.PortChannelId = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PortChannelMode != "" {
		v.PortChannelMode = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PortChannelMode)
	} else {
		v.PortChannelMode = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		v.PortTypeEdgeTrunk = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
	} else {
		v.PortTypeEdgeTrunk = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Prefix != "" {
		v.Prefix = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Prefix)
	} else {
		v.Prefix = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Ptp != nil {
		v.Ptp = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Ptp)
	} else {
		v.Ptp = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PtpProfile != "" {
		v.PtpProfile = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PtpProfile)
	} else {
		v.PtpProfile = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PtpRole != nil {
		v.PtpRole = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.PtpRole)
	} else {
		v.PtpRole = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		v.PtpVlan = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.PtpVlan)
	} else {
		v.PtpVlan = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PvlanAllowedVlans != "" {
		v.PvlanAllowedVlans = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PvlanAllowedVlans)
	} else {
		v.PvlanAllowedVlans = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PvlanMode != "" {
		v.PvlanMode = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.PvlanMode)
	} else {
		v.PvlanMode = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		v.PvlanNativeVlan = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
	} else {
		v.PvlanNativeVlan = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Qos != nil {
		v.Qos = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.Qos)
	} else {
		v.Qos = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.QosPolicy != "" {
		v.QosPolicy = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.QosPolicy)
	} else {
		v.QosPolicy = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.QueuingPolicy != "" {
		v.QueuingPolicy = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.QueuingPolicy)
	} else {
		v.QueuingPolicy = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.RoutingTag != "" {
		v.RoutingTag = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.RoutingTag)
	} else {
		v.RoutingTag = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		v.SecurityGroupProxy = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
	} else {
		v.SecurityGroupProxy = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.Speed != "" {
		v.Speed = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.Speed)
	} else {
		v.Speed = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControl != nil {
		v.StormControl = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.StormControl)
	} else {
		v.StormControl = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlAction != "" {
		v.StormControlAction = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.StormControlAction)
	} else {
		v.StormControlAction = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
		v.StormControlBroadcastLevel = types.Float64Value(float64(*jsonData.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel))
	} else {
		v.StormControlBroadcastLevel = types.Float64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		v.StormControlBroadcastLevelPps = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
	} else {
		v.StormControlBroadcastLevelPps = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
		v.StormControlMulticastLevel = types.Float64Value(float64(*jsonData.ConfigData.NetworkOS.Policy.StormControlMulticastLevel))
	} else {
		v.StormControlMulticastLevel = types.Float64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		v.StormControlMulticastLevelPps = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
	} else {
		v.StormControlMulticastLevelPps = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
		v.StormControlUnicastLevel = types.Float64Value(float64(*jsonData.ConfigData.NetworkOS.Policy.StormControlUnicastLevel))
	} else {
		v.StormControlUnicastLevel = types.Float64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		v.StormControlUnicastLevelPps = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
	} else {
		v.StormControlUnicastLevelPps = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		v.UnicastBandwidthPercent = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
	} else {
		v.UnicastBandwidthPercent = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.V6prefix != nil {
		v.V6prefix = types.Int64Value(*jsonData.ConfigData.NetworkOS.Policy.V6prefix)
	} else {
		v.V6prefix = types.Int64Null()
	}

	if jsonData.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		v.VlanMapping = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.VlanMapping)
	} else {
		v.VlanMapping = types.BoolNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.VrfInterface != "" {
		v.VrfInterface = types.StringValue(jsonData.ConfigData.NetworkOS.Policy.VrfInterface)
	} else {
		v.VrfInterface = types.StringNull()
	}

	if jsonData.ConfigData.NetworkOS.Policy.WanLink != nil {
		v.WanLink = types.BoolValue(*jsonData.ConfigData.NetworkOS.Policy.WanLink)
	} else {
		v.WanLink = types.BoolNull()
	}

	if len(jsonData.ConfigData.NetworkOS.Policy.AssociationList) == 0 {
		log.Printf("v.AssociationList is empty")
		v.AssociationList = types.ListNull(AssociationListValue{}.Type(context.Background()))
	} else {
		listData := make([]AssociationListValue, len(jsonData.ConfigData.NetworkOS.Policy.AssociationList))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.AssociationList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.AssociationList, err = types.ListValueFrom(context.Background(), AssociationListValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if len(jsonData.ConfigData.NetworkOS.Policy.MappingList) == 0 {
		log.Printf("v.MappingList is empty")
		v.MappingList = types.ListNull(MappingListValue{}.Type(context.Background()))
	} else {
		listData := make([]MappingListValue, len(jsonData.ConfigData.NetworkOS.Policy.MappingList))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.MappingList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.MappingList, err = types.ListValueFrom(context.Background(), MappingListValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if len(jsonData.ConfigData.NetworkOS.Policy.IgmpGroupRanges) == 0 {
		log.Printf("v.IgmpGroupRanges is empty")
		v.IgmpGroupRanges = types.ListNull(IgmpGroupRangesValue{}.Type(context.Background()))
	} else {
		listData := make([]IgmpGroupRangesValue, len(jsonData.ConfigData.NetworkOS.Policy.IgmpGroupRanges))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.IgmpGroupRanges {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.IgmpGroupRanges, err = types.ListValueFrom(context.Background(), IgmpGroupRangesValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if len(jsonData.ConfigData.NetworkOS.Policy.MstPathCostList) == 0 {
		log.Printf("v.MstPathCostList is empty")
		v.MstPathCostList = types.ListNull(MstPathCostListValue{}.Type(context.Background()))
	} else {
		listData := make([]MstPathCostListValue, len(jsonData.ConfigData.NetworkOS.Policy.MstPathCostList))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.MstPathCostList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.MstPathCostList, err = types.ListValueFrom(context.Background(), MstPathCostListValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if len(jsonData.ConfigData.NetworkOS.Policy.MstPortPrioList) == 0 {
		log.Printf("v.MstPortPrioList is empty")
		v.MstPortPrioList = types.ListNull(MstPortPrioListValue{}.Type(context.Background()))
	} else {
		listData := make([]MstPortPrioListValue, len(jsonData.ConfigData.NetworkOS.Policy.MstPortPrioList))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.MstPortPrioList {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.MstPortPrioList, err = types.ListValueFrom(context.Background(), MstPortPrioListValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if len(jsonData.ConfigData.NetworkOS.Policy.VlanMappingEntries) == 0 {
		log.Printf("v.VlanMappingEntries is empty")
		v.VlanMappingEntries = types.ListNull(VlanMappingEntriesValue{}.Type(context.Background()))
	} else {
		listData := make([]VlanMappingEntriesValue, len(jsonData.ConfigData.NetworkOS.Policy.VlanMappingEntries))
		for i, item := range jsonData.ConfigData.NetworkOS.Policy.VlanMappingEntries {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.VlanMappingEntries, err = types.ListValueFrom(context.Background(), VlanMappingEntriesValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}

	return err
}

func (v *AssociationListValue) SetValue(jsonData *resource_interface_common.NDFCAssociationListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.PVlan != nil {
		v.PVlan = types.Int64Value(*jsonData.PVlan)
	} else {
		v.PVlan = types.Int64Null()
	}

	if jsonData.SVlan != nil {
		v.SVlan = types.Int64Value(*jsonData.SVlan)
	} else {
		v.SVlan = types.Int64Null()
	}

	return err
}

func (v *MappingListValue) SetValue(jsonData *resource_interface_common.NDFCMappingListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.PVlan != nil {
		v.PVlan = types.Int64Value(*jsonData.PVlan)
	} else {
		v.PVlan = types.Int64Null()
	}

	if jsonData.SVlan != nil {
		v.SVlan = types.Int64Value(*jsonData.SVlan)
	} else {
		v.SVlan = types.Int64Null()
	}

	return err
}

func (v *IgmpGroupRangesValue) SetValue(jsonData *resource_interface_common.NDFCIgmpGroupRangesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.GroupAddress != "" {
		v.GroupAddress = types.StringValue(jsonData.GroupAddress)
	} else {
		v.GroupAddress = types.StringNull()
	}

	if jsonData.Prefix != nil {
		v.Prefix = types.Int64Value(*jsonData.Prefix)
	} else {
		v.Prefix = types.Int64Null()
	}

	return err
}

func (v *MstPathCostListValue) SetValue(jsonData *resource_interface_common.NDFCMstPathCostListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if len(jsonData.MstInstanceList) == 0 {
		log.Printf("v.MstInstanceList is empty")
		v.MstInstanceList = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.MstInstanceList))
		for i, item := range jsonData.MstInstanceList {
			listData[i] = types.StringValue(item)
		}
		v.MstInstanceList, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.PathCost != nil {
		v.PathCost = types.Int64Value(*jsonData.PathCost)
	} else {
		v.PathCost = types.Int64Null()
	}

	return err
}

func (v *MstPortPrioListValue) SetValue(jsonData *resource_interface_common.NDFCMstPortPrioListValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if len(jsonData.MstInstanceList) == 0 {
		log.Printf("v.MstInstanceList is empty")
		v.MstInstanceList = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.MstInstanceList))
		for i, item := range jsonData.MstInstanceList {
			listData[i] = types.StringValue(item)
		}
		v.MstInstanceList, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.PortPriority != nil {
		v.PortPriority = types.Int64Value(*jsonData.PortPriority)
	} else {
		v.PortPriority = types.Int64Null()
	}

	return err
}

func (v *VlanMappingEntriesValue) SetValue(jsonData *resource_interface_common.NDFCVlanMappingEntriesValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if len(jsonData.CustomerVlanId) == 0 {
		log.Printf("v.CustomerVlanId is empty")
		v.CustomerVlanId = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.CustomerVlanId))
		for i, item := range jsonData.CustomerVlanId {
			listData[i] = types.StringValue(item)
		}
		v.CustomerVlanId, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	if jsonData.Dot1qTunnel != nil {
		v.Dot1qTunnel = types.BoolValue(*jsonData.Dot1qTunnel)

	} else {
		v.Dot1qTunnel = types.BoolNull()
	}

	if jsonData.CustomerInnerVlanId != nil {
		v.CustomerInnerVlanId = types.Int64Value(*jsonData.CustomerInnerVlanId)
	} else {
		v.CustomerInnerVlanId = types.Int64Null()
	}

	if jsonData.ProviderVlanId != nil {
		v.ProviderVlanId = types.Int64Value(*jsonData.ProviderVlanId)
	} else {
		v.ProviderVlanId = types.Int64Null()
	}

	return err
}

func (v InterfaceEthernetModel) GetModelData() *resource_interface_common.NDFCInterfaceCommonModel {
	var data = new(resource_interface_common.NDFCInterfaceCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.SwitchId.IsNull() && !v.SwitchId.IsUnknown() {
		data.SwitchId = v.SwitchId.ValueString()
	} else {
		data.SwitchId = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.Interfaces.IsNull() && !v.Interfaces.IsUnknown() {
		elements1 := make(map[string]InterfacesValue, len(v.Interfaces.Elements()))

		data.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
		diag := v.Interfaces.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(resource_interface_common.NDFCInterfacesValue)

			// interface_name | String| []| true

			// interface_type | String| []| true

			// switch_id | String| []| false
			if !ele1.SwitchId.IsNull() && !ele1.SwitchId.IsUnknown() {

				data1.SwitchId = ele1.SwitchId.ValueString()
			} else {
				data1.SwitchId = ""
			}

			// mode | String| [configData]| false
			if !ele1.Mode.IsNull() && !ele1.Mode.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.Mode = ele1.Mode.ValueString()
			} else {
				data1.ConfigData.Mode = ""
			}

			// network_os_type | String| [configData networkOS]| false
			if !ele1.NetworkOsType.IsNull() && !ele1.NetworkOsType.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.NetworkOsType = ele1.NetworkOsType.ValueString()
			} else {
				data1.ConfigData.NetworkOS.NetworkOsType = ""
			}

			// policy_type | String| [configData networkOS policy]| false
			if !ele1.PolicyType.IsNull() && !ele1.PolicyType.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PolicyType = ele1.PolicyType.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PolicyType = ""
			}

			// template_name | String| [configData networkOS policy]| false
			if !ele1.TemplateName.IsNull() && !ele1.TemplateName.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.TemplateName = ele1.TemplateName.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.TemplateName = ""
			}

			// template_config | Map:String| [configData networkOS policy]| false
			if !ele1.TemplateConfig.IsNull() && !ele1.TemplateConfig.IsUnknown() {
				//-----inline nested----
				mapStringData := make(map[string]string, len(ele1.TemplateConfig.Elements()))
				dg := ele1.TemplateConfig.ElementsAs(context.Background(), &mapStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.ConfigData.NetworkOS.Policy.TemplateConfig = mapStringData
			}

			// admin_state | Bool| [configData networkOS policy]| false
			if !ele1.AdminState.IsNull() && !ele1.AdminState.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.AdminState = new(bool)
				*data1.ConfigData.NetworkOS.Policy.AdminState = ele1.AdminState.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.AdminState = nil
			}

			// description | String| [configData networkOS policy]| false
			if !ele1.Description.IsNull() && !ele1.Description.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Description = ele1.Description.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Description = ""
			}

			// mtu | String| [configData networkOS policy]| false
			if !ele1.Mtu.IsNull() && !ele1.Mtu.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Mtu = ele1.Mtu.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Mtu = ""
			}

			// extra_config | String| [configData networkOS policy]| false
			if !ele1.ExtraConfig.IsNull() && !ele1.ExtraConfig.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.ExtraConfig = ele1.ExtraConfig.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.ExtraConfig = ""
			}

			// access_vlan | Int64| [configData networkOS policy]| false
			if !ele1.AccessVlan.IsNull() && !ele1.AccessVlan.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.AccessVlan = new(int64)
				*data1.ConfigData.NetworkOS.Policy.AccessVlan = ele1.AccessVlan.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.AccessVlan = nil
			}

			// allowed_vlans | String| [configData networkOS policy]| false
			if !ele1.AllowedVlans.IsNull() && !ele1.AllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.AllowedVlans = ele1.AllowedVlans.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.AllowedVlans = ""
			}

			// bandwidth | Int64| [configData networkOS policy]| false
			if !ele1.Bandwidth.IsNull() && !ele1.Bandwidth.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Bandwidth = new(int64)
				*data1.ConfigData.NetworkOS.Policy.Bandwidth = ele1.Bandwidth.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.Bandwidth = nil
			}

			// bandwidth_capacity_percent | Int64| [configData networkOS policy]| false
			if !ele1.BandwidthCapacityPercent.IsNull() && !ele1.BandwidthCapacityPercent.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent = new(int64)
				*data1.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent = ele1.BandwidthCapacityPercent.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent = nil
			}

			// border_router | Bool| [configData networkOS policy]| false
			if !ele1.BorderRouter.IsNull() && !ele1.BorderRouter.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.BorderRouter = new(bool)
				*data1.ConfigData.NetworkOS.Policy.BorderRouter = ele1.BorderRouter.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.BorderRouter = nil
			}

			// bpdu_filter | String| [configData networkOS policy]| false
			if !ele1.BpduFilter.IsNull() && !ele1.BpduFilter.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.BpduFilter = ele1.BpduFilter.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.BpduFilter = ""
			}

			// bpdu_guard | String| [configData networkOS policy]| false
			if !ele1.BpduGuard.IsNull() && !ele1.BpduGuard.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.BpduGuard = ele1.BpduGuard.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.BpduGuard = ""
			}

			// cdp | Bool| [configData networkOS policy]| false
			if !ele1.Cdp.IsNull() && !ele1.Cdp.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Cdp = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Cdp = ele1.Cdp.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Cdp = nil
			}

			// debounce_linkup_timer | Int64| [configData networkOS policy]| false
			if !ele1.DebounceLinkupTimer.IsNull() && !ele1.DebounceLinkupTimer.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.DebounceLinkupTimer = new(int64)
				*data1.ConfigData.NetworkOS.Policy.DebounceLinkupTimer = ele1.DebounceLinkupTimer.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.DebounceLinkupTimer = nil
			}

			// debounce_timer | Int64| [configData networkOS policy]| false
			if !ele1.DebounceTimer.IsNull() && !ele1.DebounceTimer.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.DebounceTimer = new(int64)
				*data1.ConfigData.NetworkOS.Policy.DebounceTimer = ele1.DebounceTimer.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.DebounceTimer = nil
			}

			// duplex_mode | String| [configData networkOS policy]| false
			if !ele1.DuplexMode.IsNull() && !ele1.DuplexMode.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.DuplexMode = ele1.DuplexMode.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.DuplexMode = ""
			}

			// error_detection_acl | Bool| [configData networkOS policy]| false
			if !ele1.ErrorDetectionAcl.IsNull() && !ele1.ErrorDetectionAcl.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.ErrorDetectionAcl = new(bool)
				*data1.ConfigData.NetworkOS.Policy.ErrorDetectionAcl = ele1.ErrorDetectionAcl.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.ErrorDetectionAcl = nil
			}

			// fec | String| [configData networkOS policy]| false
			if !ele1.Fec.IsNull() && !ele1.Fec.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Fec = ele1.Fec.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Fec = ""
			}

			// guard_mode | String| [configData networkOS policy]| false
			if !ele1.GuardMode.IsNull() && !ele1.GuardMode.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.GuardMode = ele1.GuardMode.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.GuardMode = ""
			}

			// if_name | String| [configData networkOS policy]| false
			if !ele1.IfName.IsNull() && !ele1.IfName.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.IfName = ele1.IfName.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.IfName = ""
			}

			// inherit_bandwidth | Int64| [configData networkOS policy]| false
			if !ele1.InheritBandwidth.IsNull() && !ele1.InheritBandwidth.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.InheritBandwidth = new(int64)
				*data1.ConfigData.NetworkOS.Policy.InheritBandwidth = ele1.InheritBandwidth.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.InheritBandwidth = nil
			}

			// ip | String| [configData networkOS policy]| false
			if !ele1.Ip.IsNull() && !ele1.Ip.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Ip = ele1.Ip.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Ip = ""
			}

			// ip_redirects | Bool| [configData networkOS policy]| false
			if !ele1.IpRedirects.IsNull() && !ele1.IpRedirects.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.IpRedirects = new(bool)
				*data1.ConfigData.NetworkOS.Policy.IpRedirects = ele1.IpRedirects.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.IpRedirects = nil
			}

			// ipv6 | String| [configData networkOS policy]| false
			if !ele1.Ipv6.IsNull() && !ele1.Ipv6.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Ipv6 = ele1.Ipv6.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Ipv6 = ""
			}

			// ipv6_prefix | Int64| [configData networkOS policy]| false
			if !ele1.Ipv6Prefix.IsNull() && !ele1.Ipv6Prefix.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Ipv6Prefix = new(int64)
				*data1.ConfigData.NetworkOS.Policy.Ipv6Prefix = ele1.Ipv6Prefix.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.Ipv6Prefix = nil
			}

			// link_type | String| [configData networkOS policy]| false
			if !ele1.LinkType.IsNull() && !ele1.LinkType.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.LinkType = ele1.LinkType.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.LinkType = ""
			}

			// mask | String| [configData networkOS policy]| false
			if !ele1.Mask.IsNull() && !ele1.Mask.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Mask = ele1.Mask.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Mask = ""
			}

			// monitor | Bool| [configData networkOS policy]| false
			if !ele1.Monitor.IsNull() && !ele1.Monitor.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Monitor = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Monitor = ele1.Monitor.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Monitor = nil
			}

			// native_vlan | Int64| [configData networkOS policy]| false
			if !ele1.NativeVlan.IsNull() && !ele1.NativeVlan.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.NativeVlan = new(int64)
				*data1.ConfigData.NetworkOS.Policy.NativeVlan = ele1.NativeVlan.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.NativeVlan = nil
			}

			// negotiate_auto | Bool| [configData networkOS policy]| false
			if !ele1.NegotiateAuto.IsNull() && !ele1.NegotiateAuto.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.NegotiateAuto = new(bool)
				*data1.ConfigData.NetworkOS.Policy.NegotiateAuto = ele1.NegotiateAuto.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.NegotiateAuto = nil
			}

			// netflow | Bool| [configData networkOS policy]| false
			if !ele1.Netflow.IsNull() && !ele1.Netflow.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Netflow = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Netflow = ele1.Netflow.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Netflow = nil
			}

			// netflow_monitor | String| [configData networkOS policy]| false
			if !ele1.NetflowMonitor.IsNull() && !ele1.NetflowMonitor.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.NetflowMonitor = ele1.NetflowMonitor.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.NetflowMonitor = ""
			}

			// netflow_sampler | String| [configData networkOS policy]| false
			if !ele1.NetflowSampler.IsNull() && !ele1.NetflowSampler.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.NetflowSampler = ele1.NetflowSampler.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.NetflowSampler = ""
			}

			// orphan_port | Bool| [configData networkOS policy]| false
			if !ele1.OrphanPort.IsNull() && !ele1.OrphanPort.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.OrphanPort = new(bool)
				*data1.ConfigData.NetworkOS.Policy.OrphanPort = ele1.OrphanPort.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.OrphanPort = nil
			}

			// path_cost | Int64| [configData networkOS policy]| false
			if !ele1.PathCost.IsNull() && !ele1.PathCost.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PathCost = new(int64)
				*data1.ConfigData.NetworkOS.Policy.PathCost = ele1.PathCost.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.PathCost = nil
			}

			// pfc | Bool| [configData networkOS policy]| false
			if !ele1.Pfc.IsNull() && !ele1.Pfc.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Pfc = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Pfc = ele1.Pfc.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Pfc = nil
			}

			// pim_dr_priority | Int64| [configData networkOS policy]| false
			if !ele1.PimDrPriority.IsNull() && !ele1.PimDrPriority.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PimDrPriority = new(int64)
				*data1.ConfigData.NetworkOS.Policy.PimDrPriority = ele1.PimDrPriority.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.PimDrPriority = nil
			}

			// pim_sparse | Bool| [configData networkOS policy]| false
			if !ele1.PimSparse.IsNull() && !ele1.PimSparse.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PimSparse = new(bool)
				*data1.ConfigData.NetworkOS.Policy.PimSparse = ele1.PimSparse.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.PimSparse = nil
			}

			// port_channel_id | String| [configData networkOS policy]| false
			if !ele1.PortChannelId.IsNull() && !ele1.PortChannelId.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PortChannelId = ele1.PortChannelId.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PortChannelId = ""
			}

			// port_channel_mode | String| [configData networkOS policy]| false
			if !ele1.PortChannelMode.IsNull() && !ele1.PortChannelMode.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PortChannelMode = ele1.PortChannelMode.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PortChannelMode = ""
			}

			// port_type_edge_trunk | Bool| [configData networkOS policy]| false
			if !ele1.PortTypeEdgeTrunk.IsNull() && !ele1.PortTypeEdgeTrunk.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk = new(bool)
				*data1.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk = ele1.PortTypeEdgeTrunk.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk = nil
			}

			// prefix | String| [configData networkOS policy]| false
			if !ele1.Prefix.IsNull() && !ele1.Prefix.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Prefix = ele1.Prefix.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Prefix = ""
			}

			// ptp | Bool| [configData networkOS policy]| false
			if !ele1.Ptp.IsNull() && !ele1.Ptp.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Ptp = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Ptp = ele1.Ptp.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Ptp = nil
			}

			// ptp_profile | String| [configData networkOS policy]| false
			if !ele1.PtpProfile.IsNull() && !ele1.PtpProfile.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PtpProfile = ele1.PtpProfile.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PtpProfile = ""
			}

			// ptp_role | Bool| [configData networkOS policy]| false
			if !ele1.PtpRole.IsNull() && !ele1.PtpRole.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PtpRole = new(bool)
				*data1.ConfigData.NetworkOS.Policy.PtpRole = ele1.PtpRole.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.PtpRole = nil
			}

			// ptp_vlan | Int64| [configData networkOS policy]| false
			if !ele1.PtpVlan.IsNull() && !ele1.PtpVlan.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PtpVlan = new(int64)
				*data1.ConfigData.NetworkOS.Policy.PtpVlan = ele1.PtpVlan.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.PtpVlan = nil
			}

			// pvlan_allowed_vlans | String| [configData networkOS policy]| false
			if !ele1.PvlanAllowedVlans.IsNull() && !ele1.PvlanAllowedVlans.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PvlanAllowedVlans = ele1.PvlanAllowedVlans.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PvlanAllowedVlans = ""
			}

			// pvlan_mode | String| [configData networkOS policy]| false
			if !ele1.PvlanMode.IsNull() && !ele1.PvlanMode.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PvlanMode = ele1.PvlanMode.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.PvlanMode = ""
			}

			// pvlan_native_vlan | Int64| [configData networkOS policy]| false
			if !ele1.PvlanNativeVlan.IsNull() && !ele1.PvlanNativeVlan.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.PvlanNativeVlan = new(int64)
				*data1.ConfigData.NetworkOS.Policy.PvlanNativeVlan = ele1.PvlanNativeVlan.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.PvlanNativeVlan = nil
			}

			// qos | Bool| [configData networkOS policy]| false
			if !ele1.Qos.IsNull() && !ele1.Qos.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Qos = new(bool)
				*data1.ConfigData.NetworkOS.Policy.Qos = ele1.Qos.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.Qos = nil
			}

			// qos_policy | String| [configData networkOS policy]| false
			if !ele1.QosPolicy.IsNull() && !ele1.QosPolicy.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.QosPolicy = ele1.QosPolicy.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.QosPolicy = ""
			}

			// queuing_policy | String| [configData networkOS policy]| false
			if !ele1.QueuingPolicy.IsNull() && !ele1.QueuingPolicy.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.QueuingPolicy = ele1.QueuingPolicy.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.QueuingPolicy = ""
			}

			// routing_tag | String| [configData networkOS policy]| false
			if !ele1.RoutingTag.IsNull() && !ele1.RoutingTag.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.RoutingTag = ele1.RoutingTag.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.RoutingTag = ""
			}

			// security_group_proxy | Bool| [configData networkOS policy]| false
			if !ele1.SecurityGroupProxy.IsNull() && !ele1.SecurityGroupProxy.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.SecurityGroupProxy = new(bool)
				*data1.ConfigData.NetworkOS.Policy.SecurityGroupProxy = ele1.SecurityGroupProxy.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.SecurityGroupProxy = nil
			}

			// speed | String| [configData networkOS policy]| false
			if !ele1.Speed.IsNull() && !ele1.Speed.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.Speed = ele1.Speed.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.Speed = ""
			}

			// storm_control | Bool| [configData networkOS policy]| false
			if !ele1.StormControl.IsNull() && !ele1.StormControl.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControl = new(bool)
				*data1.ConfigData.NetworkOS.Policy.StormControl = ele1.StormControl.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.StormControl = nil
			}

			// storm_control_action | String| [configData networkOS policy]| false
			if !ele1.StormControlAction.IsNull() && !ele1.StormControlAction.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlAction = ele1.StormControlAction.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlAction = ""
			}

			// storm_control_broadcast_level | Float64| [configData networkOS policy]| false
			if !ele1.StormControlBroadcastLevel.IsNull() && !ele1.StormControlBroadcastLevel.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel = new(float64)
				*data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel = ele1.StormControlBroadcastLevel.ValueFloat64()
			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel = nil
			}

			// storm_control_broadcast_level_pps | Int64| [configData networkOS policy]| false
			if !ele1.StormControlBroadcastLevelPps.IsNull() && !ele1.StormControlBroadcastLevelPps.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps = new(int64)
				*data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps = ele1.StormControlBroadcastLevelPps.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps = nil
			}

			// storm_control_multicast_level | Float64| [configData networkOS policy]| false
			if !ele1.StormControlMulticastLevel.IsNull() && !ele1.StormControlMulticastLevel.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevel = new(float64)
				*data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevel = ele1.StormControlMulticastLevel.ValueFloat64()
			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevel = nil
			}

			// storm_control_multicast_level_pps | Int64| [configData networkOS policy]| false
			if !ele1.StormControlMulticastLevelPps.IsNull() && !ele1.StormControlMulticastLevelPps.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps = new(int64)
				*data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps = ele1.StormControlMulticastLevelPps.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps = nil
			}

			// storm_control_unicast_level | Float64| [configData networkOS policy]| false
			if !ele1.StormControlUnicastLevel.IsNull() && !ele1.StormControlUnicastLevel.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevel = new(float64)
				*data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevel = ele1.StormControlUnicastLevel.ValueFloat64()
			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevel = nil
			}

			// storm_control_unicast_level_pps | Int64| [configData networkOS policy]| false
			if !ele1.StormControlUnicastLevelPps.IsNull() && !ele1.StormControlUnicastLevelPps.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps = new(int64)
				*data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps = ele1.StormControlUnicastLevelPps.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps = nil
			}

			// unicast_bandwidth_percent | Int64| [configData networkOS policy]| false
			if !ele1.UnicastBandwidthPercent.IsNull() && !ele1.UnicastBandwidthPercent.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent = new(int64)
				*data1.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent = ele1.UnicastBandwidthPercent.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent = nil
			}

			// v6prefix | Int64| [configData networkOS policy]| false
			if !ele1.V6prefix.IsNull() && !ele1.V6prefix.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.V6prefix = new(int64)
				*data1.ConfigData.NetworkOS.Policy.V6prefix = ele1.V6prefix.ValueInt64()

			} else {
				data1.ConfigData.NetworkOS.Policy.V6prefix = nil
			}

			// vlan_mapping | Bool| [configData networkOS policy]| false
			if !ele1.VlanMapping.IsNull() && !ele1.VlanMapping.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.VlanMapping = new(bool)
				*data1.ConfigData.NetworkOS.Policy.VlanMapping = ele1.VlanMapping.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.VlanMapping = nil
			}

			// vrf_interface | String| [configData networkOS policy]| false
			if !ele1.VrfInterface.IsNull() && !ele1.VrfInterface.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.VrfInterface = ele1.VrfInterface.ValueString()
			} else {
				data1.ConfigData.NetworkOS.Policy.VrfInterface = ""
			}

			// wan_link | Bool| [configData networkOS policy]| false
			if !ele1.WanLink.IsNull() && !ele1.WanLink.IsUnknown() {
				//-----inline nested----
				data1.ConfigData.NetworkOS.Policy.WanLink = new(bool)
				*data1.ConfigData.NetworkOS.Policy.WanLink = ele1.WanLink.ValueBool()
			} else {
				data1.ConfigData.NetworkOS.Policy.WanLink = nil
			}

			// association_list | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.AssociationList.IsNull() && !ele1.AssociationList.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.AssociationList = make([]resource_interface_common.NDFCAssociationListValue, len(ele1.AssociationList.Elements()))
				elements2 := make([]AssociationListValue, len(ele1.AssociationList.Elements()))
				diag := ele1.AssociationList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCAssociationListValue)
					if !ele2.PVlan.IsNull() && !ele2.PVlan.IsUnknown() {
						data2.PVlan = new(int64)
						*data2.PVlan = ele2.PVlan.ValueInt64()

					} else {
						data2.PVlan = nil
					}
					data1.ConfigData.NetworkOS.Policy.AssociationList[i2] = *data2
					if !ele2.SVlan.IsNull() && !ele2.SVlan.IsUnknown() {
						data2.SVlan = new(int64)
						*data2.SVlan = ele2.SVlan.ValueInt64()

					} else {
						data2.SVlan = nil
					}
					data1.ConfigData.NetworkOS.Policy.AssociationList[i2] = *data2
				}
			}

			// mapping_list | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.MappingList.IsNull() && !ele1.MappingList.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.MappingList = make([]resource_interface_common.NDFCMappingListValue, len(ele1.MappingList.Elements()))
				elements2 := make([]MappingListValue, len(ele1.MappingList.Elements()))
				diag := ele1.MappingList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCMappingListValue)
					if !ele2.PVlan.IsNull() && !ele2.PVlan.IsUnknown() {
						data2.PVlan = new(int64)
						*data2.PVlan = ele2.PVlan.ValueInt64()

					} else {
						data2.PVlan = nil
					}
					data1.ConfigData.NetworkOS.Policy.MappingList[i2] = *data2
					if !ele2.SVlan.IsNull() && !ele2.SVlan.IsUnknown() {
						data2.SVlan = new(int64)
						*data2.SVlan = ele2.SVlan.ValueInt64()

					} else {
						data2.SVlan = nil
					}
					data1.ConfigData.NetworkOS.Policy.MappingList[i2] = *data2
				}
			}

			// igmp_group_ranges | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.IgmpGroupRanges.IsNull() && !ele1.IgmpGroupRanges.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.IgmpGroupRanges = make([]resource_interface_common.NDFCIgmpGroupRangesValue, len(ele1.IgmpGroupRanges.Elements()))
				elements2 := make([]IgmpGroupRangesValue, len(ele1.IgmpGroupRanges.Elements()))
				diag := ele1.IgmpGroupRanges.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCIgmpGroupRangesValue)
					if !ele2.GroupAddress.IsNull() && !ele2.GroupAddress.IsUnknown() {
						data2.GroupAddress = ele2.GroupAddress.ValueString()
					} else {
						data2.GroupAddress = ""
					}
					data1.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i2] = *data2
					if !ele2.Prefix.IsNull() && !ele2.Prefix.IsUnknown() {
						data2.Prefix = new(int64)
						*data2.Prefix = ele2.Prefix.ValueInt64()

					} else {
						data2.Prefix = nil
					}
					data1.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i2] = *data2
				}
			}

			// mst_path_cost_list | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.MstPathCostList.IsNull() && !ele1.MstPathCostList.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.MstPathCostList = make([]resource_interface_common.NDFCMstPathCostListValue, len(ele1.MstPathCostList.Elements()))
				elements2 := make([]MstPathCostListValue, len(ele1.MstPathCostList.Elements()))
				diag := ele1.MstPathCostList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCMstPathCostListValue)
					if !ele2.MstInstanceList.IsNull() && !ele2.MstInstanceList.IsUnknown() {
						listStringData := make([]string, len(ele2.MstInstanceList.Elements()))
						dg := ele2.MstInstanceList.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data2.MstInstanceList = make([]string, len(listStringData))
						copy(data2.MstInstanceList, listStringData)
					}
					data1.ConfigData.NetworkOS.Policy.MstPathCostList[i2] = *data2
					if !ele2.PathCost.IsNull() && !ele2.PathCost.IsUnknown() {
						data2.PathCost = new(int64)
						*data2.PathCost = ele2.PathCost.ValueInt64()

					} else {
						data2.PathCost = nil
					}
					data1.ConfigData.NetworkOS.Policy.MstPathCostList[i2] = *data2
				}
			}

			// mst_port_prio_list | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.MstPortPrioList.IsNull() && !ele1.MstPortPrioList.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.MstPortPrioList = make([]resource_interface_common.NDFCMstPortPrioListValue, len(ele1.MstPortPrioList.Elements()))
				elements2 := make([]MstPortPrioListValue, len(ele1.MstPortPrioList.Elements()))
				diag := ele1.MstPortPrioList.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCMstPortPrioListValue)
					if !ele2.MstInstanceList.IsNull() && !ele2.MstInstanceList.IsUnknown() {
						listStringData := make([]string, len(ele2.MstInstanceList.Elements()))
						dg := ele2.MstInstanceList.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data2.MstInstanceList = make([]string, len(listStringData))
						copy(data2.MstInstanceList, listStringData)
					}
					data1.ConfigData.NetworkOS.Policy.MstPortPrioList[i2] = *data2
					if !ele2.PortPriority.IsNull() && !ele2.PortPriority.IsUnknown() {
						data2.PortPriority = new(int64)
						*data2.PortPriority = ele2.PortPriority.ValueInt64()

					} else {
						data2.PortPriority = nil
					}
					data1.ConfigData.NetworkOS.Policy.MstPortPrioList[i2] = *data2
				}
			}

			// vlan_mapping_entries | List| [configData networkOS policy]| false//-----inline nested----

			if !ele1.VlanMappingEntries.IsNull() && !ele1.VlanMappingEntries.IsUnknown() {

				data1.ConfigData.NetworkOS.Policy.VlanMappingEntries = make([]resource_interface_common.NDFCVlanMappingEntriesValue, len(ele1.VlanMappingEntries.Elements()))
				elements2 := make([]VlanMappingEntriesValue, len(ele1.VlanMappingEntries.Elements()))
				diag := ele1.VlanMappingEntries.ElementsAs(context.Background(), &elements2, false)
				if diag.HasError() {
					panic(diag.Errors())
				}
				for i2, ele2 := range elements2 {
					data2 := new(resource_interface_common.NDFCVlanMappingEntriesValue)
					if !ele2.CustomerVlanId.IsNull() && !ele2.CustomerVlanId.IsUnknown() {
						listStringData := make([]string, len(ele2.CustomerVlanId.Elements()))
						dg := ele2.CustomerVlanId.ElementsAs(context.Background(), &listStringData, false)
						if dg.HasError() {
							panic(dg.Errors())
						}
						data2.CustomerVlanId = make([]string, len(listStringData))
						copy(data2.CustomerVlanId, listStringData)
					}
					data1.ConfigData.NetworkOS.Policy.VlanMappingEntries[i2] = *data2
					if !ele2.Dot1qTunnel.IsNull() && !ele2.Dot1qTunnel.IsUnknown() {
						data2.Dot1qTunnel = new(bool)
						*data2.Dot1qTunnel = ele2.Dot1qTunnel.ValueBool()
					} else {
						data2.Dot1qTunnel = nil
					}
					data1.ConfigData.NetworkOS.Policy.VlanMappingEntries[i2] = *data2
					if !ele2.CustomerInnerVlanId.IsNull() && !ele2.CustomerInnerVlanId.IsUnknown() {
						data2.CustomerInnerVlanId = new(int64)
						*data2.CustomerInnerVlanId = ele2.CustomerInnerVlanId.ValueInt64()

					} else {
						data2.CustomerInnerVlanId = nil
					}
					data1.ConfigData.NetworkOS.Policy.VlanMappingEntries[i2] = *data2
					if !ele2.ProviderVlanId.IsNull() && !ele2.ProviderVlanId.IsUnknown() {
						data2.ProviderVlanId = new(int64)
						*data2.ProviderVlanId = ele2.ProviderVlanId.ValueInt64()

					} else {
						data2.ProviderVlanId = nil
					}
					data1.ConfigData.NetworkOS.Policy.VlanMappingEntries[i2] = *data2
				}
			}

			data.Interfaces[k1] = *data1

		}
	}

	return data
}
