// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_common

type NDFCInterfaceCommonModel struct {
	Id         string                         `json:"-"`
	FabricName string                         `json:"-"`
	SwitchId   string                         `json:"-"`
	Deploy     bool                           `json:"-"`
	Interfaces map[string]NDFCInterfacesValue `json:"interfaces,omitempty"`
}

type NDFCInterfacesValue struct {
	InterfaceName string              `json:"interfaceName,omitempty"`
	InterfaceType string              `json:"interfaceType,omitempty"`
	SwitchId      string              `json:"switchId,omitempty"`
	ConfigData    NDFCConfigDataValue `json:"configData,omitempty"`
}

type NDFCConfigDataValue struct {
	Mode      string                       `json:"mode,omitempty"`
	NetworkOS NDFCConfigDataNetworkOSValue `json:"networkOS,omitempty"`
}

type NDFCConfigDataNetworkOSValue struct {
	NetworkOsType string                   `json:"networkOSType,omitempty"`
	Policy        NDFCNetworkOSPolicyValue `json:"policy,omitempty"`
}

type NDFCNetworkOSPolicyValue struct {
	PolicyType                    string                       `json:"policyType,omitempty"`
	TemplateName                  string                       `json:"templateName,omitempty"`
	TemplateConfig                map[string]string            `json:"templateConfig,omitempty"`
	AdminState                    *bool                        `json:"adminState,omitempty"`
	Description                   string                       `json:"description,omitempty"`
	Mtu                           string                       `json:"mtu,omitempty"`
	ExtraConfig                   string                       `json:"extraConfig,omitempty"`
	AccessVlan                    *int64                       `json:"accessVlan,omitempty"`
	AllowedVlans                  string                       `json:"allowedVlans,omitempty"`
	Bandwidth                     *int64                       `json:"bandwidth,omitempty"`
	BandwidthCapacityPercent      *int64                       `json:"bandwidthCapacityPercent,omitempty"`
	BorderRouter                  *bool                        `json:"borderRouter,omitempty"`
	BpduFilter                    string                       `json:"bpduFilter,omitempty"`
	BpduGuard                     string                       `json:"bpduGuard,omitempty"`
	Cdp                           *bool                        `json:"cdp,omitempty"`
	DebounceLinkupTimer           *int64                       `json:"debounceLinkupTimer,omitempty"`
	DebounceTimer                 *int64                       `json:"debounceTimer,omitempty"`
	DuplexMode                    string                       `json:"duplexMode,omitempty"`
	ErrorDetectionAcl             *bool                        `json:"errorDetectionAcl,omitempty"`
	Fec                           string                       `json:"fec,omitempty"`
	GuardMode                     string                       `json:"guardMode,omitempty"`
	IfName                        string                       `json:"ifName,omitempty"`
	InheritBandwidth              *int64                       `json:"inheritBandwidth,omitempty"`
	Ip                            string                       `json:"ip,omitempty"`
	IpRedirects                   *bool                        `json:"ipRedirects,omitempty"`
	Ipv6                          string                       `json:"ipv6,omitempty"`
	Ipv6Prefix                    *int64                       `json:"ipv6Prefix,omitempty"`
	LinkType                      string                       `json:"linkType,omitempty"`
	Mask                          string                       `json:"mask,omitempty"`
	Monitor                       *bool                        `json:"monitor,omitempty"`
	NativeVlan                    *int64                       `json:"nativeVlan,omitempty"`
	NegotiateAuto                 *bool                        `json:"negotiateAuto,omitempty"`
	Netflow                       *bool                        `json:"netflow,omitempty"`
	NetflowMonitor                string                       `json:"netflowMonitor,omitempty"`
	NetflowSampler                string                       `json:"netflowSampler,omitempty"`
	OrphanPort                    *bool                        `json:"orphanPort,omitempty"`
	PathCost                      *int64                       `json:"pathCost,omitempty"`
	Pfc                           *bool                        `json:"pfc,omitempty"`
	PimDrPriority                 *int64                       `json:"pimDrPriority,omitempty"`
	PimSparse                     *bool                        `json:"pimSparse,omitempty"`
	PortChannelId                 string                       `json:"portChannelId,omitempty"`
	PortChannelMode               string                       `json:"portChannelMode,omitempty"`
	PortTypeEdgeTrunk             *bool                        `json:"portTypeEdgeTrunk,omitempty"`
	Prefix                        string                       `json:"prefix,omitempty"`
	Ptp                           *bool                        `json:"ptp,omitempty"`
	PtpProfile                    string                       `json:"ptpProfile,omitempty"`
	PtpRole                       *bool                        `json:"ptpRole,omitempty"`
	PtpVlan                       *int64                       `json:"ptpVlan,omitempty"`
	PvlanAllowedVlans             string                       `json:"pvlanAllowedVlans,omitempty"`
	PvlanMode                     string                       `json:"pvlanMode,omitempty"`
	PvlanNativeVlan               *int64                       `json:"pvlanNativeVlan,omitempty"`
	Qos                           *bool                        `json:"qos,omitempty"`
	QosPolicy                     string                       `json:"qosPolicy,omitempty"`
	QueuingPolicy                 string                       `json:"queuingPolicy,omitempty"`
	RoutingTag                    string                       `json:"routingTag,omitempty"`
	SecurityGroupProxy            *bool                        `json:"securityGroupProxy,omitempty"`
	Speed                         string                       `json:"speed,omitempty"`
	StormControl                  *bool                        `json:"stormControl,omitempty"`
	StormControlAction            string                       `json:"stormControlAction,omitempty"`
	StormControlBroadcastLevel    *float64                     `json:"stormControlBroadcastLevel,omitempty"`
	StormControlBroadcastLevelPps *int64                       `json:"stormControlBroadcastLevelPps,omitempty"`
	StormControlMulticastLevel    *float64                     `json:"stormControlMulticastLevel,omitempty"`
	StormControlMulticastLevelPps *int64                       `json:"stormControlMulticastLevelPps,omitempty"`
	StormControlUnicastLevel      *float64                     `json:"stormControlUnicastLevel,omitempty"`
	StormControlUnicastLevelPps   *int64                       `json:"stormControlUnicastLevelPps,omitempty"`
	UnicastBandwidthPercent       *int64                       `json:"unicastBandwidthPercent,omitempty"`
	V6prefix                      *int64                       `json:"v6prefix,omitempty"`
	VlanMapping                   *bool                        `json:"vlanMapping,omitempty"`
	VrfInterface                  string                       `json:"vrfInterface,omitempty"`
	WanLink                       *bool                        `json:"wanLink,omitempty"`
	AssociationList               NDFCAssociationListValues    `json:"associationList,omitempty"`
	MappingList                   NDFCMappingListValues        `json:"mappingList,omitempty"`
	IgmpGroupRanges               NDFCIgmpGroupRangesValues    `json:"igmpGroupRanges,omitempty"`
	MstPathCostList               NDFCMstPathCostListValues    `json:"mstPathCostList,omitempty"`
	MstPortPrioList               NDFCMstPortPrioListValues    `json:"mstPortPrioList,omitempty"`
	VlanMappingEntries            NDFCVlanMappingEntriesValues `json:"vlanMappingEntries,omitempty"`
}

type NDFCAssociationListValues []NDFCAssociationListValue

type NDFCAssociationListValue struct {
	PVlan *int64 `json:"pVlan,omitempty"`
	SVlan *int64 `json:"sVlan,omitempty"`
}

type NDFCMappingListValues []NDFCMappingListValue

type NDFCMappingListValue struct {
	PVlan *int64 `json:"pVlan,omitempty"`
	SVlan *int64 `json:"sVlan,omitempty"`
}

type NDFCIgmpGroupRangesValues []NDFCIgmpGroupRangesValue

type NDFCIgmpGroupRangesValue struct {
	GroupAddress string `json:"groupAddress,omitempty"`
	Prefix       *int64 `json:"prefix,omitempty"`
}

type NDFCMstPathCostListValues []NDFCMstPathCostListValue

type NDFCMstPathCostListValue struct {
	MstInstanceList []string `json:"mstInstanceList,omitempty"`
	PathCost        *int64   `json:"pathCost,omitempty"`
}

type NDFCMstPortPrioListValues []NDFCMstPortPrioListValue

type NDFCMstPortPrioListValue struct {
	MstInstanceList []string `json:"mstInstanceList,omitempty"`
	PortPriority    *int64   `json:"portPriority,omitempty"`
}

type NDFCVlanMappingEntriesValues []NDFCVlanMappingEntriesValue

type NDFCVlanMappingEntriesValue struct {
	CustomerVlanId      []string `json:"customerVlanId,omitempty"`
	Dot1qTunnel         *bool    `json:"dot1qTunnel,omitempty"`
	CustomerInnerVlanId *int64   `json:"customerInnerVlanId,omitempty"`
	ProviderVlanId      *int64   `json:"providerVlanId,omitempty"`
}
