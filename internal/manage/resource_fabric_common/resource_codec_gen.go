// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_common

type NDFCFabricCommonModel struct {
	Id                         string                             `json:"-"`
	FabricName                 string                             `json:"name,omitempty"`
	LicenseTier                string                             `json:"licenseTier,omitempty"`
	TelemetryCollection        *bool                              `json:"telemetryCollection,omitempty"`
	TelemetryCollectionType    string                             `json:"telemetryCollectionType,omitempty"`
	TelemetryStreamingProtocol string                             `json:"telemetryStreamingProtocol,omitempty"`
	TelemetrySourceInterface   string                             `json:"telemetrySourceInterface,omitempty"`
	TelemetrySourceVrf         string                             `json:"telemetrySourceVrf,omitempty"`
	SecurityDomain             string                             `json:"securityDomain,omitempty"`
	Category                   string                             `json:"category,omitempty"`
	Location                   NDFCLocationValue                  `json:"location,omitempty"`
	AlertSuspend               string                             `json:"alertSuspend,omitempty"`
	Management                 NDFCManagementValue                `json:"management,omitempty"`
	ExternalStreamingSettings  NDFCExternalStreamingSettingsValue `json:"externalStreamingSettings,omitempty"`
	Meta                       NDFCMetaValue                      `json:"meta,omitempty"`
	TelemetrySettings          NDFCTelemetrySettingsValue         `json:"telemetrySettings,omitempty"`
	FeatureStatus              NDFCFeatureStatusValue             `json:"featureStatus,omitempty"`
}

type NDFCLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCManagementValue struct {
	FabricType                                 string                              `json:"type,omitempty"`
	BgpAsn                                     string                              `json:"bgpAsn,omitempty"`
	CreateBgpConfig                            *bool                               `json:"createBgpConfig,omitempty"`
	SuperSpineBgpAs                            string                              `json:"superSpineBgpAs,omitempty"`
	LeafBgpAs                                  string                              `json:"leafBgpAs,omitempty"`
	BorderBgpAs                                string                              `json:"borderBgpAs,omitempty"`
	BgpAsMode                                  string                              `json:"bgpAsMode,omitempty"`
	BgpAsnAutoAllocation                       *bool                               `json:"bgpAsnAutoAllocation,omitempty"`
	BgpAsnRange                                string                              `json:"bgpAsnRange,omitempty"`
	BgpAllowAsInNum                            *int64                              `json:"bgpAllowAsInNum,omitempty"`
	BgpMaxPath                                 *int64                              `json:"bgpMaxPath,omitempty"`
	BgpUnderlayFailureProtect                  *bool                               `json:"bgpUnderlayFailureProtect,omitempty"`
	AutoConfigureEbgpEvpnPeering               *bool                               `json:"autoConfigureEbgpEvpnPeering,omitempty"`
	AllowLeafSameAs                            *bool                               `json:"allowLeafSameAs,omitempty"`
	AssignIpv4ToLoopback0                      *bool                               `json:"assignIpv4ToLoopback0,omitempty"`
	TargetSubnetMask                           *int64                              `json:"targetSubnetMask,omitempty"`
	AnycastGatewayMac                          string                              `json:"anycastGatewayMac,omitempty"`
	PerformanceMonitoring                      *bool                               `json:"performanceMonitoring,omitempty"`
	ReplicationMode                            string                              `json:"replicationMode,omitempty"`
	MulticastGroupSubnet                       string                              `json:"multicastGroupSubnet,omitempty"`
	UnderlayMulticastGroupAddressLimit         *int64                              `json:"underlayMulticastGroupAddressLimit,omitempty"`
	TenantRoutedMulticast                      *bool                               `json:"tenantRoutedMulticast,omitempty"`
	RendezvousPointCount                       *int64                              `json:"rendezvousPointCount,omitempty"`
	RendezvousPointLoopbackId                  *int64                              `json:"rendezvousPointLoopbackId,omitempty"`
	VpcPeerLinkVlan                            string                              `json:"vpcPeerLinkVlan,omitempty"`
	VpcPeerLinkEnableNativeVlan                *bool                               `json:"vpcPeerLinkEnableNativeVlan,omitempty"`
	VpcPeerKeepAliveOption                     string                              `json:"vpcPeerKeepAliveOption,omitempty"`
	VpcAutoRecoveryTimer                       *int64                              `json:"vpcAutoRecoveryTimer,omitempty"`
	VpcDelayRestoreTimer                       *int64                              `json:"vpcDelayRestoreTimer,omitempty"`
	VpcPeerLinkPortChannelId                   string                              `json:"vpcPeerLinkPortChannelId,omitempty"`
	VpcIpv6NeighborDiscoverySync               *bool                               `json:"vpcIpv6NeighborDiscoverySync,omitempty"`
	AdvertisePhysicalIp                        *bool                               `json:"advertisePhysicalIp,omitempty"`
	VpcDomainIdRange                           string                              `json:"vpcDomainIdRange,omitempty"`
	BgpLoopbackId                              *int64                              `json:"bgpLoopbackId,omitempty"`
	AllowSameLoopbackIpOnSwitches              *bool                               `json:"allowSameLoopbackIpOnSwitches,omitempty"`
	NveLoopbackId                              *int64                              `json:"nveLoopbackId,omitempty"`
	VrfTemplate                                string                              `json:"vrfTemplate,omitempty"`
	NetworkTemplate                            string                              `json:"networkTemplate,omitempty"`
	VrfExtensionTemplate                       string                              `json:"vrfExtensionTemplate,omitempty"`
	NetworkExtensionTemplate                   string                              `json:"networkExtensionTemplate,omitempty"`
	L3VniNoVlanDefaultOption                   *bool                               `json:"l3VniNoVlanDefaultOption,omitempty"`
	SiteId                                     string                              `json:"siteId,omitempty"`
	FabricMtu                                  *int64                              `json:"fabricMtu,omitempty"`
	L2HostInterfaceMtu                         *int64                              `json:"l2HostInterfaceMtu,omitempty"`
	TenantDhcp                                 *bool                               `json:"tenantDhcp,omitempty"`
	Nxapi                                      *bool                               `json:"nxapi,omitempty"`
	NxapiHttpsPort                             *int64                              `json:"nxapiHttpsPort,omitempty"`
	NxapiHttp                                  *bool                               `json:"nxapiHttp,omitempty"`
	NxapiHttpPort                              *int64                              `json:"nxapiHttpPort,omitempty"`
	SnmpTrap                                   *bool                               `json:"snmpTrap,omitempty"`
	AnycastBorderGatewayAdvertisePhysicalIp    *bool                               `json:"anycastBorderGatewayAdvertisePhysicalIp,omitempty"`
	GreenfieldDebugFlag                        string                              `json:"greenfieldDebugFlag,omitempty"`
	TcamAllocation                             *bool                               `json:"tcamAllocation,omitempty"`
	RealTimeInterfaceStatisticsCollection      *bool                               `json:"realTimeInterfaceStatisticsCollection,omitempty"`
	InterfaceStatisticsLoadInterval            *int64                              `json:"interfaceStatisticsLoadInterval,omitempty"`
	BgpLoopbackIpRange                         string                              `json:"bgpLoopbackIpRange,omitempty"`
	NveLoopbackIpRange                         string                              `json:"nveLoopbackIpRange,omitempty"`
	AnycastRendezvousPointIpRange              string                              `json:"anycastRendezvousPointIpRange,omitempty"`
	IntraFabricSubnetRange                     string                              `json:"intraFabricSubnetRange,omitempty"`
	L2VniRange                                 string                              `json:"l2VniRange,omitempty"`
	L3VniRange                                 string                              `json:"l3VniRange,omitempty"`
	NetworkVlanRange                           string                              `json:"networkVlanRange,omitempty"`
	VrfVlanRange                               string                              `json:"vrfVlanRange,omitempty"`
	SubInterfaceDot1qRange                     string                              `json:"subInterfaceDot1qRange,omitempty"`
	VrfLiteAutoConfig                          string                              `json:"vrfLiteAutoConfig,omitempty"`
	VrfLiteSubnetRange                         string                              `json:"vrfLiteSubnetRange,omitempty"`
	VrfLiteSubnetTargetMask                    *int64                              `json:"vrfLiteSubnetTargetMask,omitempty"`
	VrfLiteIpv6SubnetRange                     string                              `json:"vrfLiteIpv6SubnetRange,omitempty"`
	VrfLiteIpv6SubnetTargetMask                *int64                              `json:"vrfLiteIpv6SubnetTargetMask,omitempty"`
	AutoUniqueVrfLiteIpPrefix                  *bool                               `json:"autoUniqueVrfLiteIpPrefix,omitempty"`
	PerVrfLoopbackAutoProvision                *bool                               `json:"perVrfLoopbackAutoProvision,omitempty"`
	PerVrfLoopbackIpRange                      string                              `json:"perVrfLoopbackIpRange,omitempty"`
	PerVrfLoopbackAutoProvisionIpv6            *bool                               `json:"perVrfLoopbackAutoProvisionIpv6,omitempty"`
	PerVrfLoopbackIpv6Range                    string                              `json:"perVrfLoopbackIpv6Range,omitempty"`
	Banner                                     string                              `json:"banner,omitempty"`
	Day0Bootstrap                              *bool                               `json:"day0Bootstrap,omitempty"`
	Day0PlugAndPlay                            *bool                               `json:"day0PlugAndPlay,omitempty"`
	InbandDay0Bootstrap                        *bool                               `json:"inbandDay0Bootstrap,omitempty"`
	LocalDhcpServer                            *bool                               `json:"localDhcpServer,omitempty"`
	DhcpProtocolVersion                        string                              `json:"dhcpProtocolVersion,omitempty"`
	DhcpStartAddress                           string                              `json:"dhcpStartAddress,omitempty"`
	DhcpEndAddress                             string                              `json:"dhcpEndAddress,omitempty"`
	DomainName                                 string                              `json:"domainName,omitempty"`
	ManagementGateway                          string                              `json:"managementGateway,omitempty"`
	ManagementIpv4Prefix                       *int64                              `json:"managementIpv4Prefix,omitempty"`
	ManagementIpv6Prefix                       *int64                              `json:"managementIpv6Prefix,omitempty"`
	BootstrapMultiSubnet                       string                              `json:"bootstrapMultiSubnet,omitempty"`
	BootstrapSubnetCollection                  NDFCBootstrapSubnetCollectionValues `json:"bootstrapSubnetCollection,omitempty"`
	ExtraConfigNxosBootstrap                   string                              `json:"extraConfigNxosBootstrap,omitempty"`
	ExtraConfigXeBootstrap                     string                              `json:"extraConfigXeBootstrap,omitempty"`
	RealTimeBackup                             *bool                               `json:"realTimeBackup,omitempty"`
	ScheduledBackup                            *bool                               `json:"scheduledBackup,omitempty"`
	ScheduledBackupTime                        string                              `json:"scheduledBackupTime,omitempty"`
	UnderlayIpv6                               *bool                               `json:"underlayIpv6,omitempty"`
	Ipv6MulticastGroupSubnet                   string                              `json:"ipv6MulticastGroupSubnet,omitempty"`
	TenantRoutedMulticastIpv6                  *bool                               `json:"tenantRoutedMulticastIpv6,omitempty"`
	MvpnVrfRouteImportId                       *bool                               `json:"mvpnVrfRouteImportId,omitempty"`
	MvpnVrfRouteImportIdRange                  string                              `json:"mvpnVrfRouteImportIdRange,omitempty"`
	VrfRouteImportIdReallocation               *bool                               `json:"vrfRouteImportIdReallocation,omitempty"`
	L3vniMulticastGroup                        string                              `json:"l3vniMulticastGroup,omitempty"`
	L3VniIpv6MulticastGroup                    string                              `json:"l3VniIpv6MulticastGroup,omitempty"`
	RendezvousPointMode                        string                              `json:"rendezvousPointMode,omitempty"`
	AutoGenerateMulticastGroupAddress          *bool                               `json:"autoGenerateMulticastGroupAddress,omitempty"`
	PhantomRendezvousPointLoopbackId1          *int64                              `json:"phantomRendezvousPointLoopbackId1,omitempty"`
	PhantomRendezvousPointLoopbackId2          *int64                              `json:"phantomRendezvousPointLoopbackId2,omitempty"`
	PhantomRendezvousPointLoopbackId3          *int64                              `json:"phantomRendezvousPointLoopbackId3,omitempty"`
	PhantomRendezvousPointLoopbackId4          *int64                              `json:"phantomRendezvousPointLoopbackId4,omitempty"`
	AdvertisePhysicalIpOnBorder                *bool                               `json:"advertisePhysicalIpOnBorder,omitempty"`
	FabricVpcDomainId                          *bool                               `json:"fabricVpcDomainId,omitempty"`
	SharedVpcDomainId                          *int64                              `json:"sharedVpcDomainId,omitempty"`
	VpcLayer3PeerRouter                        *bool                               `json:"vpcLayer3PeerRouter,omitempty"`
	FabricVpcQos                               *bool                               `json:"fabricVpcQos,omitempty"`
	FabricVpcQosPolicyName                     string                              `json:"fabricVpcQosPolicyName,omitempty"`
	EnablePeerSwitch                           *bool                               `json:"enablePeerSwitch,omitempty"`
	AnycastLoopbackId                          *int64                              `json:"anycastLoopbackId,omitempty"`
	BgpAuthentication                          *bool                               `json:"bgpAuthentication,omitempty"`
	BgpAuthenticationKeyType                   string                              `json:"bgpAuthenticationKeyType,omitempty"`
	BgpAuthenticationKey                       string                              `json:"bgpAuthenticationKey,omitempty"`
	PimHelloAuthentication                     *bool                               `json:"pimHelloAuthentication,omitempty"`
	PimHelloAuthenticationKey                  string                              `json:"pimHelloAuthenticationKey,omitempty"`
	Bfd                                        *bool                               `json:"bfd,omitempty"`
	BfdIbgp                                    *bool                               `json:"bfdIbgp,omitempty"`
	BfdAuthentication                          *bool                               `json:"bfdAuthentication,omitempty"`
	BfdAuthenticationKeyId                     *int64                              `json:"bfdAuthenticationKeyId,omitempty"`
	BfdAuthenticationKey                       string                              `json:"bfdAuthenticationKey,omitempty"`
	Macsec                                     *bool                               `json:"macsec,omitempty"`
	MacsecCipherSuite                          string                              `json:"macsecCipherSuite,omitempty"`
	MacsecKeyString                            string                              `json:"macsecKeyString,omitempty"`
	MacsecAlgorithm                            string                              `json:"macsecAlgorithm,omitempty"`
	MacsecFallbackKeyString                    string                              `json:"macsecFallbackKeyString,omitempty"`
	MacsecFallbackAlgorithm                    string                              `json:"macsecFallbackAlgorithm,omitempty"`
	MacsecReportTimer                          *int64                              `json:"macsecReportTimer,omitempty"`
	MonitoredMode                              *bool                               `json:"monitoredMode,omitempty"`
	OverlayMode                                string                              `json:"overlayMode,omitempty"`
	PrivateVlan                                *bool                               `json:"privateVlan,omitempty"`
	DefaultPrivateVlanSecondaryNetworkTemplate string                              `json:"defaultPrivateVlanSecondaryNetworkTemplate,omitempty"`
	PowerRedundancyMode                        string                              `json:"powerRedundancyMode,omitempty"`
	CoppPolicy                                 string                              `json:"coppPolicy,omitempty"`
	NveHoldDownTimer                           *int64                              `json:"nveHoldDownTimer,omitempty"`
	Cdp                                        *bool                               `json:"cdp,omitempty"`
	NextGenerationOam                          *bool                               `json:"nextGenerationOAM,omitempty"`
	NgoamSouthBoundLoopDetect                  *bool                               `json:"ngoamSouthBoundLoopDetect,omitempty"`
	NgoamSouthBoundLoopDetectProbeInterval     *int64                              `json:"ngoamSouthBoundLoopDetectProbeInterval,omitempty"`
	NgoamSouthBoundLoopDetectRecoveryInterval  *int64                              `json:"ngoamSouthBoundLoopDetectRecoveryInterval,omitempty"`
	StrictConfigComplianceMode                 *bool                               `json:"strictConfigComplianceMode,omitempty"`
	AdvancedSshOption                          *bool                               `json:"advancedSshOption,omitempty"`
	Ptp                                        *bool                               `json:"ptp,omitempty"`
	PtpLoopbackId                              *int64                              `json:"ptpLoopbackId,omitempty"`
	PtpDomainId                                *int64                              `json:"ptpDomainId,omitempty"`
	DefaultQueuingPolicy                       *bool                               `json:"defaultQueuingPolicy,omitempty"`
	DefaultQueuingPolicyCloudscale             string                              `json:"defaultQueuingPolicyCloudscale,omitempty"`
	DefaultQueuingPolicyRSeries                string                              `json:"defaultQueuingPolicyRSeries,omitempty"`
	DefaultQueuingPolicyOther                  string                              `json:"defaultQueuingPolicyOther,omitempty"`
	AimlQos                                    *bool                               `json:"aimlQos,omitempty"`
	AimlQosPolicy                              string                              `json:"aimlQosPolicy,omitempty"`
	PriorityFlowControlWatchInterval           *int64                              `json:"priorityFlowControlWatchInterval,omitempty"`
	Dlb                                        *bool                               `json:"dlb,omitempty"`
	DlbMode                                    string                              `json:"dlbMode,omitempty"`
	DlbMixedModeDefault                        string                              `json:"dlbMixedModeDefault,omitempty"`
	FlowletAging                               *int64                              `json:"flowletAging,omitempty"`
	FlowletDscp                                string                              `json:"flowletDscp,omitempty"`
	PerPacketDscp                              string                              `json:"perPacketDscp,omitempty"`
	AiLoadSharing                              *bool                               `json:"aiLoadSharing,omitempty"`
	RoceV2                                     string                              `json:"rocev2,omitempty"`
	Cnp                                        string                              `json:"cnp,omitempty"`
	WredMin                                    *int64                              `json:"wredMin,omitempty"`
	WredMax                                    *int64                              `json:"wredMax,omitempty"`
	WredDropProbability                        *int64                              `json:"wredDropProbability,omitempty"`
	WredWeight                                 *int64                              `json:"wredWeight,omitempty"`
	BandwidthRemaining                         *int64                              `json:"bandwidthRemaining,omitempty"`
	StaticUnderlayIpAllocation                 *bool                               `json:"staticUnderlayIpAllocation,omitempty"`
	BgpLoopbackIpv6Range                       string                              `json:"bgpLoopbackIpv6Range,omitempty"`
	NveLoopbackIpv6Range                       string                              `json:"nveLoopbackIpv6Range,omitempty"`
	Ipv6AnycastRendezvousPointIpRange          string                              `json:"ipv6AnycastRendezvousPointIpRange,omitempty"`
	ExtraConfigAaa                             string                              `json:"extraConfigAaa,omitempty"`
	ExtraConfigFabric                          string                              `json:"extraConfigFabric,omitempty"`
	Aaa                                        *bool                               `json:"aaa,omitempty"`
	Ipv6LinkLocal                              *bool                               `json:"ipv6LinkLocal,omitempty"`
	FabricInterfaceType                        string                              `json:"fabricInterfaceType,omitempty"`
	Ipv6SubnetTargetMask                       *int64                              `json:"ipv6SubnetTargetMask,omitempty"`
	LinkStateRoutingProtocol                   string                              `json:"linkStateRoutingProtocol,omitempty"`
	RouteReflectorCount                        *int64                              `json:"routeReflectorCount,omitempty"`
	VpcTorDelayRestoreTimer                    *int64                              `json:"vpcTorDelayRestoreTimer,omitempty"`
	LeafTorIdRange                             *bool                               `json:"leafTorIdRange,omitempty"`
	LeafTorVpcPortChannelIdRange               string                              `json:"leafTorVpcPortChannelIdRange,omitempty"`
	LinkStateRoutingTag                        string                              `json:"linkStateRoutingTag,omitempty"`
	OspfAreaId                                 string                              `json:"ospfAreaId,omitempty"`
	OspfAuthentication                         *bool                               `json:"ospfAuthentication,omitempty"`
	OspfAuthenticationKeyId                    *int64                              `json:"ospfAuthenticationKeyId,omitempty"`
	OspfAuthenticationKey                      string                              `json:"ospfAuthenticationKey,omitempty"`
	IsisLevel                                  string                              `json:"isisLevel,omitempty"`
	IsisAreaNumber                             string                              `json:"isisAreaNumber,omitempty"`
	IsisPointToPoint                           *bool                               `json:"isisPointToPoint,omitempty"`
	IsisAuthentication                         *bool                               `json:"isisAuthentication,omitempty"`
	IsisAuthenticationKeychainName             string                              `json:"isisAuthenticationKeychainName,omitempty"`
	IsisAuthenticationKeychainKeyId            *int64                              `json:"isisAuthenticationKeychainKeyId,omitempty"`
	IsisAuthenticationKey                      string                              `json:"isisAuthenticationKey,omitempty"`
	IsisOverload                               *bool                               `json:"isisOverload,omitempty"`
	IsisOverloadElapseTime                     *int64                              `json:"isisOverloadElapseTime,omitempty"`
	BfdOspf                                    *bool                               `json:"bfdOspf,omitempty"`
	BfdIsis                                    *bool                               `json:"bfdIsis,omitempty"`
	BfdPim                                     *bool                               `json:"bfdPim,omitempty"`
	AutoBgpNeighborDescription                 *bool                               `json:"autoBgpNeighborDescription,omitempty"`
	IbgpPeerTemplate                           string                              `json:"ibgpPeerTemplate,omitempty"`
	LeafibgpPeerTemplate                       string                              `json:"leafibgpPeerTemplate,omitempty"`
	SecurityGroupTag                           *bool                               `json:"securityGroupTag,omitempty"`
	SecurityGroupTagPrefix                     string                              `json:"securityGroupTagPrefix,omitempty"`
	SecurityGroupTagIdRange                    string                              `json:"securityGroupTagIdRange,omitempty"`
	SecurityGroupTagPreprovision               *bool                               `json:"securityGroupTagPreprovision,omitempty"`
	SecurityGroupTagMacSegmentation            *bool                               `json:"securityGroupTagMacSegmentation,omitempty"`
	SecurityGroupStatus                        string                              `json:"securityGroupStatus,omitempty"`
	VrfLiteMacsec                              *bool                               `json:"vrfLiteMacsec,omitempty"`
	QuantumKeyDistribution                     *bool                               `json:"quantumKeyDistribution,omitempty"`
	VrfLiteMacsecCipherSuite                   string                              `json:"vrfLiteMacsecCipherSuite,omitempty"`
	VrfLiteMacsecKeyString                     string                              `json:"vrfLiteMacsecKeyString,omitempty"`
	VrfLiteMacsecAlgorithm                     string                              `json:"vrfLiteMacsecAlgorithm,omitempty"`
	VrfLiteMacsecFallbackKeyString             string                              `json:"vrfLiteMacsecFallbackKeyString,omitempty"`
	VrfLiteMacsecFallbackAlgorithm             string                              `json:"vrfLiteMacsecFallbackAlgorithm,omitempty"`
	QuantumKeyDistributionProfileName          string                              `json:"quantumKeyDistributionProfileName,omitempty"`
	KeyManagementEntityServerIp                string                              `json:"keyManagementEntityServerIp,omitempty"`
	KeyManagementEntityServerPort              *int64                              `json:"keyManagementEntityServerPort,omitempty"`
	TrustpointLabel                            string                              `json:"trustpointLabel,omitempty"`
	SkipCertificateVerification                *bool                               `json:"skipCertificateVerification,omitempty"`
	HostInterfaceAdminState                    *bool                               `json:"hostInterfaceAdminState,omitempty"`
	BrownfieldNetworkNameFormat                string                              `json:"brownfieldNetworkNameFormat,omitempty"`
	BrownfieldSkipOverlayNetworkAttachments    *bool                               `json:"brownfieldSkipOverlayNetworkAttachments,omitempty"`
	PolicyBasedRouting                         *bool                               `json:"policyBasedRouting,omitempty"`
	PtpVlanId                                  *int64                              `json:"ptpVlanId,omitempty"`
	MplsHandoff                                *bool                               `json:"mplsHandoff,omitempty"`
	MplsLoopbackIdentifier                     *int64                              `json:"mplsLoopbackIdentifier,omitempty"`
	MplsIsisAreaNumber                         string                              `json:"mplsIsisAreaNumber,omitempty"`
	StpRootOption                              string                              `json:"stpRootOption,omitempty"`
	StpVlanRange                               string                              `json:"stpVlanRange,omitempty"`
	MstInstanceRange                           string                              `json:"mstInstanceRange,omitempty"`
	StpBridgePriority                          *int64                              `json:"stpBridgePriority,omitempty"`
	AllowVlanOnLeafTorPairing                  string                              `json:"allowVlanOnLeafTorPairing,omitempty"`
	PreInterfaceConfigLeaf                     string                              `json:"preInterfaceConfigLeaf,omitempty"`
	PreInterfaceConfigSpine                    string                              `json:"preInterfaceConfigSpine,omitempty"`
	PreInterfaceConfigTor                      string                              `json:"preInterfaceConfigTor,omitempty"`
	ExtraConfigLeaf                            string                              `json:"extraConfigLeaf,omitempty"`
	ExtraConfigSpine                           string                              `json:"extraConfigSpine,omitempty"`
	ExtraConfigTor                             string                              `json:"extraConfigTor,omitempty"`
	ExtraConfigIntraFabricLinks                string                              `json:"extraConfigIntraFabricLinks,omitempty"`
	MplsLoopbackIpRange                        string                              `json:"mplsLoopbackIpRange,omitempty"`
	Ipv6SubnetRange                            string                              `json:"ipv6SubnetRange,omitempty"`
	RouterIdRange                              string                              `json:"routerIdRange,omitempty"`
	AutoSymmetricVrfLite                       *bool                               `json:"autoSymmetricVrfLite,omitempty"`
	AutoVrfLiteDefaultVrf                      *bool                               `json:"autoVrfLiteDefaultVrf,omitempty"`
	AutoSymmetricDefaultVrf                    *bool                               `json:"autoSymmetricDefaultVrf,omitempty"`
	DefaultVrfRedistributionBgpRouteMap        string                              `json:"defaultVrfRedistributionBgpRouteMap,omitempty"`
	IpServiceLevelAgreementIdRange             string                              `json:"ipServiceLevelAgreementIdRange,omitempty"`
	ObjectTrackingNumberRange                  string                              `json:"objectTrackingNumberRange,omitempty"`
	ServiceNetworkVlanRange                    string                              `json:"serviceNetworkVlanRange,omitempty"`
	RouteMapSequenceNumberRange                string                              `json:"routeMapSequenceNumberRange,omitempty"`
	InbandManagement                           *bool                               `json:"inbandManagement,omitempty"`
	SeedSwitchCoreInterfaces                   string                              `json:"seedSwitchCoreInterfaces,omitempty"`
	SpineSwitchCoreInterfaces                  string                              `json:"spineSwitchCoreInterfaces,omitempty"`
	InbandDhcpServers                          string                              `json:"inbandDhcpServers,omitempty"`
	UnNumberedBootstrapLbId                    *int64                              `json:"unNumberedBootstrapLbId,omitempty"`
	UnNumberedDhcpStartAddress                 string                              `json:"unNumberedDhcpStartAddress,omitempty"`
	UnNumberedDhcpEndAddress                   string                              `json:"unNumberedDhcpEndAddress,omitempty"`
	HeartbeatInterval                          *int64                              `json:"heartbeatInterval,omitempty"`
	AllowSmartSwitchOnboarding                 *bool                               `json:"allowSmartSwitchOnboarding,omitempty"`
	EnableDpuPinning                           *bool                               `json:"enableDpuPinning,omitempty"`
	ConnectivityDomainName                     string                              `json:"connectivityDomainName,omitempty"`
	HypershieldConnectivityProxyServer         string                              `json:"hypershieldConnectivityProxyServer,omitempty"`
	HypershieldConnectivityProxyServerPort     *int64                              `json:"hypershieldConnectivityProxyServerPort,omitempty"`
	HypershieldConnectivitySourceIntf          string                              `json:"hypershieldConnectivitySourceIntf,omitempty"`
	DnsCollection                              []string                            `json:"dnsCollection,omitempty"`
	DnsVrfCollection                           []string                            `json:"dnsVrfCollection,omitempty"`
	NtpServerCollection                        []string                            `json:"ntpServerCollection,omitempty"`
	NtpServerVrfCollection                     []string                            `json:"ntpServerVrfCollection,omitempty"`
	SyslogServerCollection                     []string                            `json:"syslogServerCollection,omitempty"`
	SyslogSeverityCollection                   []int64                             `json:"syslogSeverityCollection,omitempty"`
	SyslogServerVrfCollection                  []string                            `json:"syslogServerVrfCollection,omitempty"`
	NetflowSettings                            NDFCManagementNetflowSettingsValue  `json:"netflowSettings,omitempty"`
}

type NDFCBootstrapSubnetCollectionValues []NDFCBootstrapSubnetCollectionValue

type NDFCBootstrapSubnetCollectionValue struct {
	StartIp        string `json:"startIp,omitempty"`
	EndIp          string `json:"endIp,omitempty"`
	DefaultGateway string `json:"defaultGateway,omitempty"`
	SubnetPrefix   *int64 `json:"subnetPrefix,omitempty"`
}

type NDFCManagementNetflowSettingsValue struct {
	NetflowEnable             *bool                               `json:"netflow,omitempty"`
	NetflowExporterCollection NDFCNetflowExporterCollectionValues `json:"netflowExporterCollection,omitempty"`
	NetflowRecordCollection   NDFCNetflowRecordCollectionValues   `json:"netflowRecordCollection,omitempty"`
	NetflowMonitorCollection  NDFCNetflowMonitorCollectionValues  `json:"netflowMonitorCollection,omitempty"`
	NetflowSamplerCollection  NDFCNetflowSamplerCollectionValues  `json:"netflowSamplerCollection,omitempty"`
}

type NDFCNetflowExporterCollectionValues []NDFCNetflowExporterCollectionValue

type NDFCNetflowExporterCollectionValue struct {
	ExporterName        string `json:"exporterName,omitempty"`
	ExporterIp          string `json:"exporterIp,omitempty"`
	Vrf                 string `json:"vrf,omitempty"`
	SourceInterfaceName string `json:"sourceInterfaceName,omitempty"`
	UdpPort             *int64 `json:"udpPort,omitempty"`
}

type NDFCNetflowRecordCollectionValues []NDFCNetflowRecordCollectionValue

type NDFCNetflowRecordCollectionValue struct {
	RecordName     string `json:"recordName,omitempty"`
	RecordTemplate string `json:"recordTemplate,omitempty"`
	Layer2Record   string `json:"layer2Record,omitempty"`
}

type NDFCNetflowMonitorCollectionValues []NDFCNetflowMonitorCollectionValue

type NDFCNetflowMonitorCollectionValue struct {
	MonitorName       string `json:"monitorName,omitempty"`
	MonitorRecordName string `json:"recordName,omitempty"`
	Exporter1Name     string `json:"exporter1Name,omitempty"`
	Exporter2Name     string `json:"exporter2Name,omitempty"`
}

type NDFCNetflowSamplerCollectionValues []NDFCNetflowSamplerCollectionValue

type NDFCNetflowSamplerCollectionValue struct {
	SamplerName  string `json:"samplerName,omitempty"`
	NumSamples   *int64 `json:"numSamples,omitempty"`
	SamplingRate *int64 `json:"samplingRate,omitempty"`
}

type NDFCExternalStreamingSettingsValue struct {
	Email      NDFCEmailValues                          `json:"email,omitempty"`
	MessageBus NDFCMessageBusValues                     `json:"messageBus,omitempty"`
	Syslog     NDFCExternalStreamingSettingsSyslogValue `json:"syslog,omitempty"`
}

type NDFCEmailValues []NDFCEmailValue

type NDFCEmailValue struct {
	Name                      string                      `json:"name,omitempty"`
	ReceiverEmail             string                      `json:"receiverEmail,omitempty"`
	Format                    string                      `json:"format,omitempty"`
	StartDate                 string                      `json:"startDate,omitempty"`
	CollectionFrequencyInDays *int64                      `json:"collectionFrequencyInDays,omitempty"`
	OnlyIncludeActiveAlerts   *bool                       `json:"onlyIncludeActiveAlerts,omitempty"`
	CollectionSettings        NDFCCollectionSettingsValue `json:"collectionSettings,omitempty"`
}

type NDFCCollectionSettingsValue struct {
	CollectionType            string   `json:"collectionType,omitempty"`
	Anomalies                 []string `json:"anomalies,omitempty"`
	Advisories                []string `json:"advisories,omitempty"`
	RiskAndConformanceReports []string `json:"riskAndConformanceReports,omitempty"`
}

type NDFCMessageBusValues []NDFCMessageBusValue

type NDFCMessageBusValue struct {
	Server             string                                `json:"server,omitempty"`
	CollectionType     string                                `json:"collectionType,omitempty"`
	CollectionSettings NDFCMessageBusCollectionSettingsValue `json:"collectionSettings,omitempty"`
}

type NDFCMessageBusCollectionSettingsValue struct {
	CollectionSettingsCollectionType string   `json:"collectionType,omitempty"`
	Anomalies                        []string `json:"anomalies,omitempty"`
	Advisories                       []string `json:"advisories,omitempty"`
	Statistics                       []string `json:"statistics,omitempty"`
	Faults                           []string `json:"faults,omitempty"`
	AuditLogs                        []string `json:"auditLogs,omitempty"`
}

type NDFCExternalStreamingSettingsSyslogValue struct {
	SyslogServers      []string                          `json:"servers,omitempty"`
	SyslogFacility     string                            `json:"facility,omitempty"`
	CollectionSettings NDFCSyslogCollectionSettingsValue `json:"collectionSettings,omitempty"`
}

type NDFCSyslogCollectionSettingsValue struct {
	SyslogAnomalies []string `json:"anomalies,omitempty"`
}

type NDFCMetaValue struct {
	AllowedActions []string `json:"allowedActions,omitempty"`
}

type NDFCTelemetrySettingsValue struct {
	DummyField       string                                     `json:"-"`
	FlowCollection   NDFCTelemetrySettingsFlowCollectionValue   `json:"flowCollection,omitempty"`
	Microburst       NDFCTelemetrySettingsMicroburstValue       `json:"microburst,omitempty"`
	AnalysisSettings NDFCTelemetrySettingsAnalysisSettingsValue `json:"analysisSettings,omitempty"`
	Nas              NDFCTelemetrySettingsNasValue              `json:"nas,omitempty"`
	EnergyManagement NDFCTelemetrySettingsEnergyManagementValue `json:"energyManagement,omitempty"`
}

type NDFCTelemetrySettingsFlowCollectionValue struct {
	TrafficAnalytics           string                                            `json:"trafficAnalytics,omitempty"`
	OperatingMode              string                                            `json:"operatingMode,omitempty"`
	UdpCategorizationSupport   string                                            `json:"udpCategorization,omitempty"`
	FlowCollectionModes        NDFCFlowCollectionFlowCollectionModesValue        `json:"flowCollectionModes,omitempty"`
	FlowRules                  NDFCFlowCollectionFlowRulesValue                  `json:"flowRules,omitempty"`
	TrafficAnalyticsRules      NDFCFlowCollectionTrafficAnalyticsRulesValue      `json:"trafficAnalyticsRules,omitempty"`
	FlowCollectionCapabilities NDFCFlowCollectionFlowCollectionCapabilitiesValue `json:"flowCollectionCapabilities,omitempty"`
}

type NDFCFlowCollectionFlowCollectionModesValue struct {
	NetFlow       *bool `json:"netFlow,omitempty"`
	SFlow         *bool `json:"sFlow,omitempty"`
	FlowTelemetry *bool `json:"flowTelemetry,omitempty"`
}

type NDFCFlowCollectionFlowRulesValue struct {
	VrfFlowRules       NDFCVrfFlowRulesValues       `json:"vrfFlowRules,omitempty"`
	InterfaceFlowRules NDFCInterfaceFlowRulesValues `json:"interfaceFlowRules,omitempty"`
	L3OutFlowRules     NDFCL3OutFlowRulesValues     `json:"l3OutFlowRules,omitempty"`
}

type NDFCVrfFlowRulesValues []NDFCVrfFlowRulesValue

type NDFCVrfFlowRulesValue struct {
	VrfFlowRuleName       string                          `json:"name,omitempty"`
	VrfFlowRuleUuid       string                          `json:"uuid,omitempty"`
	VrfFlowRuleTenant     string                          `json:"tenant,omitempty"`
	VrfFlowRuleVrf        string                          `json:"vrf,omitempty"`
	VrfFlowRuleSubnets    []string                        `json:"subnets,omitempty"`
	VrfFlowRuleAttributes NDFCVrfFlowRuleAttributesValues `json:"attributes,omitempty"`
}

type NDFCVrfFlowRuleAttributesValues []NDFCVrfFlowRuleAttributesValue

type NDFCVrfFlowRuleAttributesValue struct {
	VrfFlowRuleBidirectional *bool  `json:"bidirectional,omitempty"`
	VrfFlowRuleDstIp         string `json:"dstIp,omitempty"`
	VrfFlowRuleSrcIp         string `json:"srcIp,omitempty"`
	VrfFlowRuleDstPort       string `json:"dstPort,omitempty"`
	VrfFlowRuleSrcPort       string `json:"srcPort,omitempty"`
	VrfFlowRuleProtocol      string `json:"protocol,omitempty"`
	VrfFlowRuleAttributeId   string `json:"attributeId,omitempty"`
}

type NDFCInterfaceFlowRulesValues []NDFCInterfaceFlowRulesValue

type NDFCInterfaceFlowRulesValue struct {
	InterfaceFlowRuleName                string                                         `json:"name,omitempty"`
	InterfaceFlowRuleUuid                string                                         `json:"uuid,omitempty"`
	InterfaceFlowRuleType                string                                         `json:"type,omitempty"`
	InterfaceFlowRuleInterfaceCollection NDFCInterfaceFlowRuleInterfaceCollectionValues `json:"interfaceCollection,omitempty"`
	InterfaceFlowRuleSubnets             []string                                       `json:"subnets,omitempty"`
	InterfaceFlowRuleAttributes          NDFCInterfaceFlowRuleAttributesValues          `json:"attributes,omitempty"`
}

type NDFCInterfaceFlowRuleInterfaceCollectionValues []NDFCInterfaceFlowRuleInterfaceCollectionValue

type NDFCInterfaceFlowRuleInterfaceCollectionValue struct {
	InterfaceFlowRuleSwitchId   string   `json:"switchId,omitempty"`
	InterfaceFlowRuleSwitchName string   `json:"switchName,omitempty"`
	InterfaceFlowRuleInterfaces []string `json:"interfaces,omitempty"`
}

type NDFCInterfaceFlowRuleAttributesValues []NDFCInterfaceFlowRuleAttributesValue

type NDFCInterfaceFlowRuleAttributesValue struct {
	InterfaceFlowRuleBidirectional *bool  `json:"bidirectional,omitempty"`
	InterfaceFlowRuleDstIp         string `json:"dstIp,omitempty"`
	InterfaceFlowRuleSrcIp         string `json:"srcIp,omitempty"`
	InterfaceFlowRuleDstPort       string `json:"dstPort,omitempty"`
	InterfaceFlowRuleSrcPort       string `json:"srcPort,omitempty"`
	InterfaceFlowRuleProtocol      string `json:"protocol,omitempty"`
	InterfaceFlowRuleAttributeId   string `json:"attributeId,omitempty"`
}

type NDFCL3OutFlowRulesValues []NDFCL3OutFlowRulesValue

type NDFCL3OutFlowRulesValue struct {
	L3OutFlowRuleName                string                                     `json:"name,omitempty"`
	L3OutFlowRuleUuid                string                                     `json:"uuid,omitempty"`
	L3OutFlowRuleType                string                                     `json:"type,omitempty"`
	L3OutFlowRuleInterfaceCollection NDFCL3OutFlowRuleInterfaceCollectionValues `json:"interfaceCollection,omitempty"`
	L3OutFlowRuleSubnets             []string                                   `json:"subnets,omitempty"`
}

type NDFCL3OutFlowRuleInterfaceCollectionValues []NDFCL3OutFlowRuleInterfaceCollectionValue

type NDFCL3OutFlowRuleInterfaceCollectionValue struct {
	L3OutFlowRuleTenant     string   `json:"tenant,omitempty"`
	L3OutFlowRuleL3Out      string   `json:"l3Out,omitempty"`
	L3OutFlowRuleEncap      string   `json:"encap,omitempty"`
	L3OutFlowRuleSwitchName string   `json:"switchName,omitempty"`
	L3OutFlowRuleSwitchId   string   `json:"switchId,omitempty"`
	L3OutFlowRuleInterfaces []string `json:"interfaces,omitempty"`
}

type NDFCFlowCollectionTrafficAnalyticsRulesValue struct {
	TrafficAnalyticsRulesEnabled *bool                    `json:"enabled,omitempty"`
	InterfaceRules               NDFCInterfaceRulesValues `json:"interfaceRules,omitempty"`
}

type NDFCInterfaceRulesValues []NDFCInterfaceRulesValue

type NDFCInterfaceRulesValue struct {
	InterfaceRuleName                     string                                     `json:"name,omitempty"`
	InterfaceRuleInterfaceCollection      NDFCInterfaceRuleInterfaceCollectionValues `json:"interfaceCollection,omitempty"`
	InterfaceRuleEnabled                  *bool                                      `json:"enabled,omitempty"`
	InterfaceRuleEnableFabricInterconnect *bool                                      `json:"enableFabricInterconnect,omitempty"`
	InterfaceRuleEnableL3Out              *bool                                      `json:"enableL3Out,omitempty"`
	InterfaceRuleUuid                     string                                     `json:"uuid,omitempty"`
	InterfaceRuleSubnets                  []string                                   `json:"subnets,omitempty"`
}

type NDFCInterfaceRuleInterfaceCollectionValues []NDFCInterfaceRuleInterfaceCollectionValue

type NDFCInterfaceRuleInterfaceCollectionValue struct {
	InterfaceRuleSwitchId   string                            `json:"switchId,omitempty"`
	InterfaceRuleSwitchName string                            `json:"switchName,omitempty"`
	InterfaceRuleVrfName    string                            `json:"vrfName,omitempty"`
	InterfaceRuleInterfaces NDFCInterfaceRuleInterfacesValues `json:"interfaces,omitempty"`
	InterfaceRuleTenant     string                            `json:"tenant,omitempty"`
	InterfaceRuleL3Out      string                            `json:"l3Out,omitempty"`
}

type NDFCInterfaceRuleInterfacesValues []NDFCInterfaceRuleInterfacesValue

type NDFCInterfaceRuleInterfacesValue struct {
	InterfaceRuleInterfaceName    string `json:"name,omitempty"`
	InterfaceRuleInterfaceType    string `json:"type,omitempty"`
	InterfaceRuleInterfaceEncap   string `json:"encap,omitempty"`
	InterfaceRuleInterfaceVrfName string `json:"vrfName,omitempty"`
}

type NDFCFlowCollectionFlowCollectionCapabilitiesValue struct {
	TrafficAnalyticsMode        string `json:"trafficAnalyticsMode,omitempty"`
	UdpCategorization           string `json:"udpCategorization,omitempty"`
	TrafficAnalyticsFilterRules string `json:"trafficAnalyticsFilterRules,omitempty"`
}

type NDFCTelemetrySettingsMicroburstValue struct {
	Microburst  *bool  `json:"microburst,omitempty"`
	Sensitivity string `json:"sensitivity,omitempty"`
}

type NDFCTelemetrySettingsAnalysisSettingsValue struct {
	AnalysisSettingsIsEnabled *bool `json:"isEnabled,omitempty"`
}

type NDFCTelemetrySettingsNasValue struct {
	Server         string                     `json:"server,omitempty"`
	ExportSettings NDFCNasExportSettingsValue `json:"exportSettings,omitempty"`
}

type NDFCNasExportSettingsValue struct {
	ExportType   string `json:"exportType,omitempty"`
	ExportFormat string `json:"exportFormat,omitempty"`
}

type NDFCTelemetrySettingsEnergyManagementValue struct {
	Cost *float64 `json:"cost,omitempty"`
}

type NDFCFeatureStatusValue struct {
	ControllerStatus    string `json:"controller,omitempty"`
	TelemetryStatus     string `json:"telemetry,omitempty"`
	OrchestrationStatus string `json:"orchestration,omitempty"`
	TrapForwarderStatus string `json:"trapForwarder,omitempty"`
}
