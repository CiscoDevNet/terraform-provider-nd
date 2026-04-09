// Code generated;  DO NOT EDIT.

package resource_fabric_vxlan

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCFabricVxlanModel struct {
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
	FeatureStatus              NDFCFeatureStatusValue             `json:"featureStatus,omitempty"`
	Meta                       NDFCMetaValue                      `json:"meta,omitempty"`
	ExternalStreamingSettings  NDFCExternalStreamingSettingsValue `json:"externalStreamingSettings,omitempty"`
	Management                 NDFCManagementValue                `json:"management,omitempty"`
	TelemetrySettings          NDFCTelemetrySettingsValue         `json:"telemetrySettings,omitempty"`
}

type NDFCLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCFeatureStatusValue struct {
	ControllerStatus    string `json:"controller,omitempty"`
	TelemetryStatus     string `json:"telemetry,omitempty"`
	OrchestrationStatus string `json:"orchestration,omitempty"`
	TrapForwarderStatus string `json:"trapForwarder,omitempty"`
}

type NDFCMetaValue struct {
	AllowedActions []string `json:"allowedActions,omitempty"`
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

type NDFCManagementValue struct {
	FabricType                                 string                             `json:"type,omitempty"`
	BgpAsn                                     string                             `json:"bgpAsn,omitempty"`
	SuperSpineBgpAs                            string                             `json:"superSpineBgpAs,omitempty"`
	LeafBgpAs                                  string                             `json:"leafBgpAs,omitempty"`
	BorderBgpAs                                string                             `json:"borderBgpAs,omitempty"`
	BgpAsMode                                  string                             `json:"bgpAsMode,omitempty"`
	TargetSubnetMask                           *int64                             `json:"targetSubnetMask,omitempty"`
	AnycastGatewayMac                          string                             `json:"anycastGatewayMac,omitempty"`
	PerformanceMonitoring                      *bool                              `json:"performanceMonitoring,omitempty"`
	ReplicationMode                            string                             `json:"replicationMode,omitempty"`
	MulticastGroupSubnet                       string                             `json:"multicastGroupSubnet,omitempty"`
	TenantRoutedMulticast                      *bool                              `json:"tenantRoutedMulticast,omitempty"`
	RendezvousPointCount                       *int64                             `json:"rendezvousPointCount,omitempty"`
	RendezvousPointLoopbackId                  *int64                             `json:"rendezvousPointLoopbackId,omitempty"`
	VpcPeerLinkVlan                            string                             `json:"vpcPeerLinkVlan,omitempty"`
	VpcPeerLinkEnableNativeVlan                *bool                              `json:"vpcPeerLinkEnableNativeVlan,omitempty"`
	VpcPeerKeepAliveOption                     string                             `json:"vpcPeerKeepAliveOption,omitempty"`
	VpcAutoRecoveryTimer                       *int64                             `json:"vpcAutoRecoveryTimer,omitempty"`
	VpcDelayRestoreTimer                       *int64                             `json:"vpcDelayRestoreTimer,omitempty"`
	VpcPeerLinkPortChannelId                   string                             `json:"vpcPeerLinkPortChannelId,omitempty"`
	VpcIpv6NeighborDiscoverySync               *bool                              `json:"vpcIpv6NeighborDiscoverySync,omitempty"`
	AdvertisePhysicalIp                        *bool                              `json:"advertisePhysicalIp,omitempty"`
	VpcDomainIdRange                           string                             `json:"vpcDomainIdRange,omitempty"`
	BgpLoopbackId                              *int64                             `json:"bgpLoopbackId,omitempty"`
	NveLoopbackId                              *int64                             `json:"nveLoopbackId,omitempty"`
	VrfTemplate                                string                             `json:"vrfTemplate,omitempty"`
	NetworkTemplate                            string                             `json:"networkTemplate,omitempty"`
	VrfExtensionTemplate                       string                             `json:"vrfExtensionTemplate,omitempty"`
	NetworkExtensionTemplate                   string                             `json:"networkExtensionTemplate,omitempty"`
	L3VniNoVlanDefaultOption                   *bool                              `json:"l3VniNoVlanDefaultOption,omitempty"`
	SiteId                                     string                             `json:"siteId,omitempty"`
	FabricMtu                                  *int64                             `json:"fabricMtu,omitempty"`
	L2HostInterfaceMtu                         *int64                             `json:"l2HostInterfaceMtu,omitempty"`
	TenantDhcp                                 *bool                              `json:"tenantDhcp,omitempty"`
	Nxapi                                      *bool                              `json:"nxapi,omitempty"`
	NxapiHttpsPort                             *int64                             `json:"nxapiHttpsPort,omitempty"`
	NxapiHttp                                  *bool                              `json:"nxapiHttp,omitempty"`
	NxapiHttpPort                              *int64                             `json:"nxapiHttpPort,omitempty"`
	SnmpTrap                                   *bool                              `json:"snmpTrap,omitempty"`
	AnycastBorderGatewayAdvertisePhysicalIp    *bool                              `json:"anycastBorderGatewayAdvertisePhysicalIp,omitempty"`
	GreenfieldDebugFlag                        string                             `json:"greenfieldDebugFlag,omitempty"`
	TcamAllocation                             *bool                              `json:"tcamAllocation,omitempty"`
	RealTimeInterfaceStatisticsCollection      *bool                              `json:"realTimeInterfaceStatisticsCollection,omitempty"`
	InterfaceStatisticsLoadInterval            *int64                             `json:"interfaceStatisticsLoadInterval,omitempty"`
	BgpLoopbackIpRange                         string                             `json:"bgpLoopbackIpRange,omitempty"`
	NveLoopbackIpRange                         string                             `json:"nveLoopbackIpRange,omitempty"`
	AnycastRendezvousPointIpRange              string                             `json:"anycastRendezvousPointIpRange,omitempty"`
	IntraFabricSubnetRange                     string                             `json:"intraFabricSubnetRange,omitempty"`
	L2VniRange                                 string                             `json:"l2VniRange,omitempty"`
	L3VniRange                                 string                             `json:"l3VniRange,omitempty"`
	NetworkVlanRange                           string                             `json:"networkVlanRange,omitempty"`
	VrfVlanRange                               string                             `json:"vrfVlanRange,omitempty"`
	SubInterfaceDot1qRange                     string                             `json:"subInterfaceDot1qRange,omitempty"`
	VrfLiteAutoConfig                          string                             `json:"vrfLiteAutoConfig,omitempty"`
	VrfLiteSubnetRange                         string                             `json:"vrfLiteSubnetRange,omitempty"`
	VrfLiteSubnetTargetMask                    *int64                             `json:"vrfLiteSubnetTargetMask,omitempty"`
	VrfLiteIpv6SubnetRange                     string                             `json:"vrfLiteIpv6SubnetRange,omitempty"`
	VrfLiteIpv6SubnetTargetMask                *int64                             `json:"vrfLiteIpv6SubnetTargetMask,omitempty"`
	AutoUniqueVrfLiteIpPrefix                  *bool                              `json:"autoUniqueVrfLiteIpPrefix,omitempty"`
	PerVrfLoopbackAutoProvision                *bool                              `json:"perVrfLoopbackAutoProvision,omitempty"`
	PerVrfLoopbackIpRange                      string                             `json:"perVrfLoopbackIpRange,omitempty"`
	PerVrfLoopbackAutoProvisionIpv6            *bool                              `json:"perVrfLoopbackAutoProvisionIpv6,omitempty"`
	PerVrfLoopbackIpv6Range                    string                             `json:"perVrfLoopbackIpv6Range,omitempty"`
	Banner                                     string                             `json:"banner,omitempty"`
	Day0Bootstrap                              *bool                              `json:"day0Bootstrap,omitempty"`
	LocalDhcpServer                            *bool                              `json:"localDhcpServer,omitempty"`
	DhcpProtocolVersion                        string                             `json:"dhcpProtocolVersion,omitempty"`
	DhcpStartAddress                           string                             `json:"dhcpStartAddress,omitempty"`
	DhcpEndAddress                             string                             `json:"dhcpEndAddress,omitempty"`
	ManagementGateway                          string                             `json:"managementGateway,omitempty"`
	ManagementIpv4Prefix                       *int64                             `json:"managementIpv4Prefix,omitempty"`
	ManagementIpv6Prefix                       *int64                             `json:"managementIpv6Prefix,omitempty"`
	BootstrapMultiSubnet                       string                             `json:"bootstrapMultiSubnet,omitempty"`
	ExtraConfigNxosBootstrap                   string                             `json:"extraConfigNxosBootstrap,omitempty"`
	RealTimeBackup                             *bool                              `json:"realTimeBackup,omitempty"`
	ScheduledBackup                            *bool                              `json:"scheduledBackup,omitempty"`
	ScheduledBackupTime                        string                             `json:"scheduledBackupTime,omitempty"`
	UnderlayIpv6                               *bool                              `json:"underlayIpv6,omitempty"`
	Ipv6MulticastGroupSubnet                   string                             `json:"ipv6MulticastGroupSubnet,omitempty"`
	TenantRoutedMulticastIpv6                  *bool                              `json:"tenantRoutedMulticastIpv6,omitempty"`
	MvpnVrfRouteImportId                       *bool                              `json:"mvpnVrfRouteImportId,omitempty"`
	MvpnVrfRouteImportIdRange                  string                             `json:"mvpnVrfRouteImportIdRange,omitempty"`
	VrfRouteImportIdReallocation               *bool                              `json:"vrfRouteImportIdReallocation,omitempty"`
	L3vniMulticastGroup                        string                             `json:"l3vniMulticastGroup,omitempty"`
	L3VniIpv6MulticastGroup                    string                             `json:"l3VniIpv6MulticastGroup,omitempty"`
	RendezvousPointMode                        string                             `json:"rendezvousPointMode,omitempty"`
	AutoGenerateMulticastGroupAddress          *bool                              `json:"autoGenerateMulticastGroupAddress,omitempty"`
	PhantomRendezvousPointLoopbackId1          *int64                             `json:"phantomRendezvousPointLoopbackId1,omitempty"`
	PhantomRendezvousPointLoopbackId2          *int64                             `json:"phantomRendezvousPointLoopbackId2,omitempty"`
	PhantomRendezvousPointLoopbackId3          *int64                             `json:"phantomRendezvousPointLoopbackId3,omitempty"`
	PhantomRendezvousPointLoopbackId4          *int64                             `json:"phantomRendezvousPointLoopbackId4,omitempty"`
	AdvertisePhysicalIpOnBorder                *bool                              `json:"advertisePhysicalIpOnBorder,omitempty"`
	FabricVpcDomainId                          *bool                              `json:"fabricVpcDomainId,omitempty"`
	SharedVpcDomainId                          *int64                             `json:"sharedVpcDomainId,omitempty"`
	VpcLayer3PeerRouter                        *bool                              `json:"vpcLayer3PeerRouter,omitempty"`
	FabricVpcQos                               *bool                              `json:"fabricVpcQos,omitempty"`
	FabricVpcQosPolicyName                     string                             `json:"fabricVpcQosPolicyName,omitempty"`
	AnycastLoopbackId                          *int64                             `json:"anycastLoopbackId,omitempty"`
	BgpAuthentication                          *bool                              `json:"bgpAuthentication,omitempty"`
	BgpAuthenticationKeyType                   string                             `json:"bgpAuthenticationKeyType,omitempty"`
	BgpAuthenticationKey                       string                             `json:"bgpAuthenticationKey,omitempty"`
	PimHelloAuthentication                     *bool                              `json:"pimHelloAuthentication,omitempty"`
	PimHelloAuthenticationKey                  string                             `json:"pimHelloAuthenticationKey,omitempty"`
	Bfd                                        *bool                              `json:"bfd,omitempty"`
	BfdIbgp                                    *bool                              `json:"bfdIbgp,omitempty"`
	BfdAuthentication                          *bool                              `json:"bfdAuthentication,omitempty"`
	BfdAuthenticationKeyId                     *int64                             `json:"bfdAuthenticationKeyId,omitempty"`
	BfdAuthenticationKey                       string                             `json:"bfdAuthenticationKey,omitempty"`
	Macsec                                     *bool                              `json:"macsec,omitempty"`
	MacsecCipherSuite                          string                             `json:"macsecCipherSuite,omitempty"`
	MacsecKeyString                            string                             `json:"macsecKeyString,omitempty"`
	MacsecAlgorithm                            string                             `json:"macsecAlgorithm,omitempty"`
	MacsecFallbackKeyString                    string                             `json:"macsecFallbackKeyString,omitempty"`
	MacsecFallbackAlgorithm                    string                             `json:"macsecFallbackAlgorithm,omitempty"`
	MacsecReportTimer                          *int64                             `json:"macsecReportTimer,omitempty"`
	OverlayMode                                string                             `json:"overlayMode,omitempty"`
	PrivateVlan                                *bool                              `json:"privateVlan,omitempty"`
	DefaultPrivateVlanSecondaryNetworkTemplate string                             `json:"defaultPrivateVlanSecondaryNetworkTemplate,omitempty"`
	PowerRedundancyMode                        string                             `json:"powerRedundancyMode,omitempty"`
	CoppPolicy                                 string                             `json:"coppPolicy,omitempty"`
	NveHoldDownTimer                           *int64                             `json:"nveHoldDownTimer,omitempty"`
	Cdp                                        *bool                              `json:"cdp,omitempty"`
	NextGenerationOam                          *bool                              `json:"nextGenerationOAM,omitempty"`
	NgoamSouthBoundLoopDetect                  *bool                              `json:"ngoamSouthBoundLoopDetect,omitempty"`
	NgoamSouthBoundLoopDetectProbeInterval     *int64                             `json:"ngoamSouthBoundLoopDetectProbeInterval,omitempty"`
	NgoamSouthBoundLoopDetectRecoveryInterval  *int64                             `json:"ngoamSouthBoundLoopDetectRecoveryInterval,omitempty"`
	StrictConfigComplianceMode                 *bool                              `json:"strictConfigComplianceMode,omitempty"`
	AdvancedSshOption                          *bool                              `json:"advancedSshOption,omitempty"`
	Ptp                                        *bool                              `json:"ptp,omitempty"`
	PtpLoopbackId                              *int64                             `json:"ptpLoopbackId,omitempty"`
	PtpDomainId                                *int64                             `json:"ptpDomainId,omitempty"`
	DefaultQueuingPolicy                       *bool                              `json:"defaultQueuingPolicy,omitempty"`
	DefaultQueuingPolicyCloudscale             string                             `json:"defaultQueuingPolicyCloudscale,omitempty"`
	DefaultQueuingPolicyRSeries                string                             `json:"defaultQueuingPolicyRSeries,omitempty"`
	DefaultQueuingPolicyOther                  string                             `json:"defaultQueuingPolicyOther,omitempty"`
	AimlQos                                    *bool                              `json:"aimlQos,omitempty"`
	AimlQosPolicy                              string                             `json:"aimlQosPolicy,omitempty"`
	PriorityFlowControlWatchInterval           *int64                             `json:"priorityFlowControlWatchInterval,omitempty"`
	StaticUnderlayIpAllocation                 *bool                              `json:"staticUnderlayIpAllocation,omitempty"`
	BgpLoopbackIpv6Range                       string                             `json:"bgpLoopbackIpv6Range,omitempty"`
	NveLoopbackIpv6Range                       string                             `json:"nveLoopbackIpv6Range,omitempty"`
	Ipv6AnycastRendezvousPointIpRange          string                             `json:"ipv6AnycastRendezvousPointIpRange,omitempty"`
	ExtraConfigAaa                             string                             `json:"extraConfigAaa,omitempty"`
	Aaa                                        *bool                              `json:"aaa,omitempty"`
	Ipv6LinkLocal                              *bool                              `json:"ipv6LinkLocal,omitempty"`
	FabricInterfaceType                        string                             `json:"fabricInterfaceType,omitempty"`
	Ipv6SubnetTargetMask                       *int64                             `json:"ipv6SubnetTargetMask,omitempty"`
	LinkStateRoutingProtocol                   string                             `json:"linkStateRoutingProtocol,omitempty"`
	RouteReflectorCount                        *int64                             `json:"routeReflectorCount,omitempty"`
	VpcTorDelayRestoreTimer                    *int64                             `json:"vpcTorDelayRestoreTimer,omitempty"`
	LeafToRIdRange                             *bool                              `json:"leafTorIdRange,omitempty"`
	LeafTorVpcPortChannelIdRange               string                             `json:"leafTorVpcPortChannelIdRange,omitempty"`
	LinkStateRoutingTag                        string                             `json:"linkStateRoutingTag,omitempty"`
	OspfAreaId                                 string                             `json:"ospfAreaId,omitempty"`
	OspfAuthentication                         *bool                              `json:"ospfAuthentication,omitempty"`
	OspfAuthenticationKeyId                    *int64                             `json:"ospfAuthenticationKeyId,omitempty"`
	OspfAuthenticationKey                      string                             `json:"ospfAuthenticationKey,omitempty"`
	IsisLevel                                  string                             `json:"isisLevel,omitempty"`
	IsisAreaNumber                             string                             `json:"isisAreaNumber,omitempty"`
	IsisPointToPoint                           *bool                              `json:"isisPointToPoint,omitempty"`
	IsisAuthentication                         *bool                              `json:"isisAuthentication,omitempty"`
	IsisAuthenticationKeychainName             string                             `json:"isisAuthenticationKeychainName,omitempty"`
	IsisAuthenticationKeychainKeyId            *int64                             `json:"isisAuthenticationKeychainKeyId,omitempty"`
	IsisAuthenticationKey                      string                             `json:"isisAuthenticationKey,omitempty"`
	IsisOverload                               *bool                              `json:"isisOverload,omitempty"`
	IsisOverloadElapseTime                     *int64                             `json:"isisOverloadElapseTime,omitempty"`
	BfdOspf                                    *bool                              `json:"bfdOspf,omitempty"`
	BfdIsis                                    *bool                              `json:"bfdIsis,omitempty"`
	BfdPim                                     *bool                              `json:"bfdPim,omitempty"`
	AutoBgpNeighborDescription                 *bool                              `json:"autoBgpNeighborDescription,omitempty"`
	IbgpPeerTemplate                           string                             `json:"ibgpPeerTemplate,omitempty"`
	LeafibgpPeerTemplate                       string                             `json:"leafibgpPeerTemplate,omitempty"`
	SecurityGroupTag                           *bool                              `json:"securityGroupTag,omitempty"`
	SecurityGroupTagPrefix                     string                             `json:"securityGroupTagPrefix,omitempty"`
	SecurityGroupTagIdRange                    string                             `json:"securityGroupTagIdRange,omitempty"`
	SecurityGroupTagPreprovision               *bool                              `json:"securityGroupTagPreprovision,omitempty"`
	SecurityGroupStatus                        string                             `json:"securityGroupStatus,omitempty"`
	VrfLiteMacsec                              *bool                              `json:"vrfLiteMacsec,omitempty"`
	QuantumKeyDistribution                     *bool                              `json:"quantumKeyDistribution,omitempty"`
	VrfLiteMacsecCipherSuite                   string                             `json:"vrfLiteMacsecCipherSuite,omitempty"`
	VrfLiteMacsecKeyString                     string                             `json:"vrfLiteMacsecKeyString,omitempty"`
	VrfLiteMacsecAlgorithm                     string                             `json:"vrfLiteMacsecAlgorithm,omitempty"`
	VrfLiteMacsecFallbackKeyString             string                             `json:"vrfLiteMacsecFallbackKeyString,omitempty"`
	VrfLiteMacsecFallbackAlgorithm             string                             `json:"vrfLiteMacsecFallbackAlgorithm,omitempty"`
	QuantumKeyDistributionProfileName          string                             `json:"quantumKeyDistributionProfileName,omitempty"`
	KeyManagementEntityServerIp                string                             `json:"keyManagementEntityServerIp,omitempty"`
	KeyManagementEntityServerPort              *int64                             `json:"keyManagementEntityServerPort,omitempty"`
	TrustpointLabel                            string                             `json:"trustpointLabel,omitempty"`
	SkipCertificateVerification                *bool                              `json:"skipCertificateVerification,omitempty"`
	HostInterfaceAdminState                    *bool                              `json:"hostInterfaceAdminState,omitempty"`
	BrownfieldNetworkNameFormat                string                             `json:"brownfieldNetworkNameFormat,omitempty"`
	BrownfieldSkipOverlayNetworkAttachments    *bool                              `json:"brownfieldSkipOverlayNetworkAttachments,omitempty"`
	PolicyBasedRouting                         *bool                              `json:"policyBasedRouting,omitempty"`
	PtpVlanId                                  *int64                             `json:"ptpVlanId,omitempty"`
	MplsHandoff                                *bool                              `json:"mplsHandoff,omitempty"`
	MplsLoopbackIdentifier                     *int64                             `json:"mplsLoopbackIdentifier,omitempty"`
	MplsIsisAreaNumber                         string                             `json:"mplsIsisAreaNumber,omitempty"`
	StpRootOption                              string                             `json:"stpRootOption,omitempty"`
	StpVlanRange                               string                             `json:"stpVlanRange,omitempty"`
	MstInstanceRange                           string                             `json:"mstInstanceRange,omitempty"`
	StpBridgePriority                          *int64                             `json:"stpBridgePriority,omitempty"`
	AllowVlanOnLeafTorPairing                  string                             `json:"allowVlanOnLeafTorPairing,omitempty"`
	PreInterfaceConfigLeaf                     string                             `json:"preInterfaceConfigLeaf,omitempty"`
	PreInterfaceConfigSpine                    string                             `json:"preInterfaceConfigSpine,omitempty"`
	PreInterfaceConfigTor                      string                             `json:"preInterfaceConfigTor,omitempty"`
	ExtraConfigLeaf                            string                             `json:"extraConfigLeaf,omitempty"`
	ExtraConfigSpine                           string                             `json:"extraConfigSpine,omitempty"`
	ExtraConfigTor                             string                             `json:"extraConfigTor,omitempty"`
	ExtraConfigIntraFabricLinks                string                             `json:"extraConfigIntraFabricLinks,omitempty"`
	MplsLoopbackIpRange                        string                             `json:"mplsLoopbackIpRange,omitempty"`
	Ipv6SubnetRange                            string                             `json:"ipv6SubnetRange,omitempty"`
	RouterIdRange                              string                             `json:"routerIdRange,omitempty"`
	AutoSymmetricVrfLite                       *bool                              `json:"autoSymmetricVrfLite,omitempty"`
	AutoVrfLiteDefaultVrf                      *bool                              `json:"autoVrfLiteDefaultVrf,omitempty"`
	AutoSymmetricDefaultVrf                    *bool                              `json:"autoSymmetricDefaultVrf,omitempty"`
	DefaultVrfRedistributionBgpRouteMap        string                             `json:"defaultVrfRedistributionBgpRouteMap,omitempty"`
	IpServiceLevelAgreementIdRange             string                             `json:"ipServiceLevelAgreementIdRange,omitempty"`
	ObjectTrackingNumberRange                  string                             `json:"objectTrackingNumberRange,omitempty"`
	ServiceNetworkVlanRange                    string                             `json:"serviceNetworkVlanRange,omitempty"`
	RouteMapSequenceNumberRange                string                             `json:"routeMapSequenceNumberRange,omitempty"`
	InbandManagement                           *bool                              `json:"inbandManagement,omitempty"`
	SeedSwitchCoreInterfaces                   string                             `json:"seedSwitchCoreInterfaces,omitempty"`
	SpineSwitchCoreInterfaces                  string                             `json:"spineSwitchCoreInterfaces,omitempty"`
	InbandDhcpServers                          string                             `json:"inbandDhcpServers,omitempty"`
	UnNumberedBootstrapLbId                    *int64                             `json:"unNumberedBootstrapLbId,omitempty"`
	UnNumberedDhcpStartAddress                 string                             `json:"unNumberedDhcpStartAddress,omitempty"`
	UnNumberedDhcpEndAddress                   string                             `json:"unNumberedDhcpEndAddress,omitempty"`
	HeartbeatInterval                          *int64                             `json:"heartbeatInterval,omitempty"`
	DnsCollection                              []string                           `json:"dnsCollection,omitempty"`
	DnsVrfCollection                           []string                           `json:"dnsVrfCollection,omitempty"`
	NtpServerCollection                        []string                           `json:"ntpServerCollection,omitempty"`
	NtpServerVrfCollection                     []string                           `json:"ntpServerVrfCollection,omitempty"`
	SyslogServerCollection                     []string                           `json:"syslogServerCollection,omitempty"`
	SyslogSeverityCollection                   []int64                            `json:"syslogSeverityCollection,omitempty"`
	SyslogServerVrfCollection                  []string                           `json:"syslogServerVrfCollection,omitempty"`
	NetflowSettings                            NDFCManagementNetflowSettingsValue `json:"netflowSettings,omitempty"`
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

func (v *FabricVxlanModel) SetModelData(jsonData *NDFCFabricVxlanModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

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

	if jsonData.Management.ExtraConfigNxosBootstrap != "" {
		v.ExtraConfigNxosBootstrap = types.StringValue(jsonData.Management.ExtraConfigNxosBootstrap)

	} else {
		v.ExtraConfigNxosBootstrap = types.StringNull()
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

	if jsonData.Management.LeafToRIdRange != nil {
		v.LeafToRIdRange = types.BoolValue(*jsonData.Management.LeafToRIdRange)

	} else {
		v.LeafToRIdRange = types.BoolNull()
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

	if len(jsonData.Management.SyslogSeverityCollection) == 0 {
		log.Printf("v.SyslogSeverityCollection is empty")
		v.SyslogSeverityCollection = types.SetNull(types.Int64Type)
	} else {
		listData := make([]attr.Value, len(jsonData.Management.SyslogSeverityCollection))
		for i, item := range jsonData.Management.SyslogSeverityCollection {
			listData[i] = types.Int64Value(item)
		}
		v.SyslogSeverityCollection, err = types.SetValue(types.Int64Type, listData)
		if err != nil {
			log.Printf("Error in converting []int64 to  List")
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

func (v *LocationValue) SetValue(jsonData *NDFCLocationValue) diag.Diagnostics {

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

func (v *NetflowExporterCollectionValue) SetValue(jsonData *NDFCNetflowExporterCollectionValue) diag.Diagnostics {

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

func (v *NetflowRecordCollectionValue) SetValue(jsonData *NDFCNetflowRecordCollectionValue) diag.Diagnostics {

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

func (v *NetflowMonitorCollectionValue) SetValue(jsonData *NDFCNetflowMonitorCollectionValue) diag.Diagnostics {

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

func (v *NetflowSamplerCollectionValue) SetValue(jsonData *NDFCNetflowSamplerCollectionValue) diag.Diagnostics {

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

func (v *VrfFlowRulesValue) SetValue(jsonData *NDFCVrfFlowRulesValue) diag.Diagnostics {

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

func (v *VrfFlowRuleAttributesValue) SetValue(jsonData *NDFCVrfFlowRuleAttributesValue) diag.Diagnostics {

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

func (v *InterfaceFlowRulesValue) SetValue(jsonData *NDFCInterfaceFlowRulesValue) diag.Diagnostics {

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

func (v *InterfaceFlowRuleInterfaceCollectionValue) SetValue(jsonData *NDFCInterfaceFlowRuleInterfaceCollectionValue) diag.Diagnostics {

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

func (v *InterfaceFlowRuleAttributesValue) SetValue(jsonData *NDFCInterfaceFlowRuleAttributesValue) diag.Diagnostics {

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

func (v *L3OutFlowRulesValue) SetValue(jsonData *NDFCL3OutFlowRulesValue) diag.Diagnostics {

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

func (v *L3OutFlowRuleInterfaceCollectionValue) SetValue(jsonData *NDFCL3OutFlowRuleInterfaceCollectionValue) diag.Diagnostics {

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

func (v *InterfaceRulesValue) SetValue(jsonData *NDFCInterfaceRulesValue) diag.Diagnostics {

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

func (v *InterfaceRuleInterfaceCollectionValue) SetValue(jsonData *NDFCInterfaceRuleInterfaceCollectionValue) diag.Diagnostics {

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

func (v *InterfaceRuleInterfacesValue) SetValue(jsonData *NDFCInterfaceRuleInterfacesValue) diag.Diagnostics {

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

func (v *EmailValue) SetValue(jsonData *NDFCEmailValue) diag.Diagnostics {

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

func (v *MessageBusValue) SetValue(jsonData *NDFCMessageBusValue) diag.Diagnostics {

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

func (v FabricVxlanModel) GetModelData() *NDFCFabricVxlanModel {
	var data = new(NDFCFabricVxlanModel)

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

	if !v.FabricType.IsNull() && !v.FabricType.IsUnknown() {
		data.Management.FabricType = v.FabricType.ValueString()
	} else {
		data.Management.FabricType = ""
	}

	if !v.BgpAsn.IsNull() && !v.BgpAsn.IsUnknown() {
		data.Management.BgpAsn = v.BgpAsn.ValueString()
	} else {
		data.Management.BgpAsn = ""
	}

	if !v.SuperSpineBgpAs.IsNull() && !v.SuperSpineBgpAs.IsUnknown() {
		data.Management.SuperSpineBgpAs = v.SuperSpineBgpAs.ValueString()
	} else {
		data.Management.SuperSpineBgpAs = ""
	}

	if !v.LeafBgpAs.IsNull() && !v.LeafBgpAs.IsUnknown() {
		data.Management.LeafBgpAs = v.LeafBgpAs.ValueString()
	} else {
		data.Management.LeafBgpAs = ""
	}

	if !v.BorderBgpAs.IsNull() && !v.BorderBgpAs.IsUnknown() {
		data.Management.BorderBgpAs = v.BorderBgpAs.ValueString()
	} else {
		data.Management.BorderBgpAs = ""
	}

	if !v.BgpAsMode.IsNull() && !v.BgpAsMode.IsUnknown() {
		data.Management.BgpAsMode = v.BgpAsMode.ValueString()
	} else {
		data.Management.BgpAsMode = ""
	}

	if !v.TargetSubnetMask.IsNull() && !v.TargetSubnetMask.IsUnknown() {
		data.Management.TargetSubnetMask = new(int64)
		*data.Management.TargetSubnetMask = v.TargetSubnetMask.ValueInt64()

	} else {
		data.Management.TargetSubnetMask = nil
	}

	if !v.AnycastGatewayMac.IsNull() && !v.AnycastGatewayMac.IsUnknown() {
		data.Management.AnycastGatewayMac = v.AnycastGatewayMac.ValueString()
	} else {
		data.Management.AnycastGatewayMac = ""
	}

	if !v.PerformanceMonitoring.IsNull() && !v.PerformanceMonitoring.IsUnknown() {
		data.Management.PerformanceMonitoring = new(bool)
		*data.Management.PerformanceMonitoring = v.PerformanceMonitoring.ValueBool()
	} else {
		data.Management.PerformanceMonitoring = nil
	}

	if !v.ReplicationMode.IsNull() && !v.ReplicationMode.IsUnknown() {
		data.Management.ReplicationMode = v.ReplicationMode.ValueString()
	} else {
		data.Management.ReplicationMode = ""
	}

	if !v.MulticastGroupSubnet.IsNull() && !v.MulticastGroupSubnet.IsUnknown() {
		data.Management.MulticastGroupSubnet = v.MulticastGroupSubnet.ValueString()
	} else {
		data.Management.MulticastGroupSubnet = ""
	}

	if !v.TenantRoutedMulticast.IsNull() && !v.TenantRoutedMulticast.IsUnknown() {
		data.Management.TenantRoutedMulticast = new(bool)
		*data.Management.TenantRoutedMulticast = v.TenantRoutedMulticast.ValueBool()
	} else {
		data.Management.TenantRoutedMulticast = nil
	}

	if !v.RendezvousPointCount.IsNull() && !v.RendezvousPointCount.IsUnknown() {
		data.Management.RendezvousPointCount = new(int64)
		*data.Management.RendezvousPointCount = v.RendezvousPointCount.ValueInt64()

	} else {
		data.Management.RendezvousPointCount = nil
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

	if !v.RendezvousPointLoopbackId.IsNull() && !v.RendezvousPointLoopbackId.IsUnknown() {
		data.Management.RendezvousPointLoopbackId = new(int64)
		*data.Management.RendezvousPointLoopbackId = v.RendezvousPointLoopbackId.ValueInt64()

	} else {
		data.Management.RendezvousPointLoopbackId = nil
	}

	if !v.VpcPeerLinkVlan.IsNull() && !v.VpcPeerLinkVlan.IsUnknown() {
		data.Management.VpcPeerLinkVlan = v.VpcPeerLinkVlan.ValueString()
	} else {
		data.Management.VpcPeerLinkVlan = ""
	}

	if !v.VpcPeerLinkEnableNativeVlan.IsNull() && !v.VpcPeerLinkEnableNativeVlan.IsUnknown() {
		data.Management.VpcPeerLinkEnableNativeVlan = new(bool)
		*data.Management.VpcPeerLinkEnableNativeVlan = v.VpcPeerLinkEnableNativeVlan.ValueBool()
	} else {
		data.Management.VpcPeerLinkEnableNativeVlan = nil
	}

	if !v.VpcPeerKeepAliveOption.IsNull() && !v.VpcPeerKeepAliveOption.IsUnknown() {
		data.Management.VpcPeerKeepAliveOption = v.VpcPeerKeepAliveOption.ValueString()
	} else {
		data.Management.VpcPeerKeepAliveOption = ""
	}

	if !v.VpcAutoRecoveryTimer.IsNull() && !v.VpcAutoRecoveryTimer.IsUnknown() {
		data.Management.VpcAutoRecoveryTimer = new(int64)
		*data.Management.VpcAutoRecoveryTimer = v.VpcAutoRecoveryTimer.ValueInt64()

	} else {
		data.Management.VpcAutoRecoveryTimer = nil
	}

	if !v.VpcDelayRestoreTimer.IsNull() && !v.VpcDelayRestoreTimer.IsUnknown() {
		data.Management.VpcDelayRestoreTimer = new(int64)
		*data.Management.VpcDelayRestoreTimer = v.VpcDelayRestoreTimer.ValueInt64()

	} else {
		data.Management.VpcDelayRestoreTimer = nil
	}

	if !v.VpcPeerLinkPortChannelId.IsNull() && !v.VpcPeerLinkPortChannelId.IsUnknown() {
		data.Management.VpcPeerLinkPortChannelId = v.VpcPeerLinkPortChannelId.ValueString()
	} else {
		data.Management.VpcPeerLinkPortChannelId = ""
	}

	if !v.VpcIpv6NeighborDiscoverySync.IsNull() && !v.VpcIpv6NeighborDiscoverySync.IsUnknown() {
		data.Management.VpcIpv6NeighborDiscoverySync = new(bool)
		*data.Management.VpcIpv6NeighborDiscoverySync = v.VpcIpv6NeighborDiscoverySync.ValueBool()
	} else {
		data.Management.VpcIpv6NeighborDiscoverySync = nil
	}

	if !v.AdvertisePhysicalIp.IsNull() && !v.AdvertisePhysicalIp.IsUnknown() {
		data.Management.AdvertisePhysicalIp = new(bool)
		*data.Management.AdvertisePhysicalIp = v.AdvertisePhysicalIp.ValueBool()
	} else {
		data.Management.AdvertisePhysicalIp = nil
	}

	if !v.VpcDomainIdRange.IsNull() && !v.VpcDomainIdRange.IsUnknown() {
		data.Management.VpcDomainIdRange = v.VpcDomainIdRange.ValueString()
	} else {
		data.Management.VpcDomainIdRange = ""
	}

	if !v.BgpLoopbackId.IsNull() && !v.BgpLoopbackId.IsUnknown() {
		data.Management.BgpLoopbackId = new(int64)
		*data.Management.BgpLoopbackId = v.BgpLoopbackId.ValueInt64()

	} else {
		data.Management.BgpLoopbackId = nil
	}

	if !v.NveLoopbackId.IsNull() && !v.NveLoopbackId.IsUnknown() {
		data.Management.NveLoopbackId = new(int64)
		*data.Management.NveLoopbackId = v.NveLoopbackId.ValueInt64()

	} else {
		data.Management.NveLoopbackId = nil
	}

	if !v.VrfTemplate.IsNull() && !v.VrfTemplate.IsUnknown() {
		data.Management.VrfTemplate = v.VrfTemplate.ValueString()
	} else {
		data.Management.VrfTemplate = ""
	}

	if !v.NetworkTemplate.IsNull() && !v.NetworkTemplate.IsUnknown() {
		data.Management.NetworkTemplate = v.NetworkTemplate.ValueString()
	} else {
		data.Management.NetworkTemplate = ""
	}

	if !v.VrfExtensionTemplate.IsNull() && !v.VrfExtensionTemplate.IsUnknown() {
		data.Management.VrfExtensionTemplate = v.VrfExtensionTemplate.ValueString()
	} else {
		data.Management.VrfExtensionTemplate = ""
	}

	if !v.NetworkExtensionTemplate.IsNull() && !v.NetworkExtensionTemplate.IsUnknown() {
		data.Management.NetworkExtensionTemplate = v.NetworkExtensionTemplate.ValueString()
	} else {
		data.Management.NetworkExtensionTemplate = ""
	}

	if !v.L3VniNoVlanDefaultOption.IsNull() && !v.L3VniNoVlanDefaultOption.IsUnknown() {
		data.Management.L3VniNoVlanDefaultOption = new(bool)
		*data.Management.L3VniNoVlanDefaultOption = v.L3VniNoVlanDefaultOption.ValueBool()
	} else {
		data.Management.L3VniNoVlanDefaultOption = nil
	}

	if !v.SiteId.IsNull() && !v.SiteId.IsUnknown() {
		data.Management.SiteId = v.SiteId.ValueString()
	} else {
		data.Management.SiteId = ""
	}

	if !v.FabricMtu.IsNull() && !v.FabricMtu.IsUnknown() {
		data.Management.FabricMtu = new(int64)
		*data.Management.FabricMtu = v.FabricMtu.ValueInt64()

	} else {
		data.Management.FabricMtu = nil
	}

	if !v.L2HostInterfaceMtu.IsNull() && !v.L2HostInterfaceMtu.IsUnknown() {
		data.Management.L2HostInterfaceMtu = new(int64)
		*data.Management.L2HostInterfaceMtu = v.L2HostInterfaceMtu.ValueInt64()

	} else {
		data.Management.L2HostInterfaceMtu = nil
	}

	if !v.TenantDhcp.IsNull() && !v.TenantDhcp.IsUnknown() {
		data.Management.TenantDhcp = new(bool)
		*data.Management.TenantDhcp = v.TenantDhcp.ValueBool()
	} else {
		data.Management.TenantDhcp = nil
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

	if !v.SnmpTrap.IsNull() && !v.SnmpTrap.IsUnknown() {
		data.Management.SnmpTrap = new(bool)
		*data.Management.SnmpTrap = v.SnmpTrap.ValueBool()
	} else {
		data.Management.SnmpTrap = nil
	}

	if !v.AnycastBorderGatewayAdvertisePhysicalIp.IsNull() && !v.AnycastBorderGatewayAdvertisePhysicalIp.IsUnknown() {
		data.Management.AnycastBorderGatewayAdvertisePhysicalIp = new(bool)
		*data.Management.AnycastBorderGatewayAdvertisePhysicalIp = v.AnycastBorderGatewayAdvertisePhysicalIp.ValueBool()
	} else {
		data.Management.AnycastBorderGatewayAdvertisePhysicalIp = nil
	}

	if !v.GreenfieldDebugFlag.IsNull() && !v.GreenfieldDebugFlag.IsUnknown() {
		data.Management.GreenfieldDebugFlag = v.GreenfieldDebugFlag.ValueString()
	} else {
		data.Management.GreenfieldDebugFlag = ""
	}

	if !v.TcamAllocation.IsNull() && !v.TcamAllocation.IsUnknown() {
		data.Management.TcamAllocation = new(bool)
		*data.Management.TcamAllocation = v.TcamAllocation.ValueBool()
	} else {
		data.Management.TcamAllocation = nil
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

	if !v.BgpLoopbackIpRange.IsNull() && !v.BgpLoopbackIpRange.IsUnknown() {
		data.Management.BgpLoopbackIpRange = v.BgpLoopbackIpRange.ValueString()
	} else {
		data.Management.BgpLoopbackIpRange = ""
	}

	if !v.NveLoopbackIpRange.IsNull() && !v.NveLoopbackIpRange.IsUnknown() {
		data.Management.NveLoopbackIpRange = v.NveLoopbackIpRange.ValueString()
	} else {
		data.Management.NveLoopbackIpRange = ""
	}

	if !v.AnycastRendezvousPointIpRange.IsNull() && !v.AnycastRendezvousPointIpRange.IsUnknown() {
		data.Management.AnycastRendezvousPointIpRange = v.AnycastRendezvousPointIpRange.ValueString()
	} else {
		data.Management.AnycastRendezvousPointIpRange = ""
	}

	if !v.IntraFabricSubnetRange.IsNull() && !v.IntraFabricSubnetRange.IsUnknown() {
		data.Management.IntraFabricSubnetRange = v.IntraFabricSubnetRange.ValueString()
	} else {
		data.Management.IntraFabricSubnetRange = ""
	}

	if !v.L2VniRange.IsNull() && !v.L2VniRange.IsUnknown() {
		data.Management.L2VniRange = v.L2VniRange.ValueString()
	} else {
		data.Management.L2VniRange = ""
	}

	if !v.L3VniRange.IsNull() && !v.L3VniRange.IsUnknown() {
		data.Management.L3VniRange = v.L3VniRange.ValueString()
	} else {
		data.Management.L3VniRange = ""
	}

	if !v.NetworkVlanRange.IsNull() && !v.NetworkVlanRange.IsUnknown() {
		data.Management.NetworkVlanRange = v.NetworkVlanRange.ValueString()
	} else {
		data.Management.NetworkVlanRange = ""
	}

	if !v.VrfVlanRange.IsNull() && !v.VrfVlanRange.IsUnknown() {
		data.Management.VrfVlanRange = v.VrfVlanRange.ValueString()
	} else {
		data.Management.VrfVlanRange = ""
	}

	if !v.SubInterfaceDot1qRange.IsNull() && !v.SubInterfaceDot1qRange.IsUnknown() {
		data.Management.SubInterfaceDot1qRange = v.SubInterfaceDot1qRange.ValueString()
	} else {
		data.Management.SubInterfaceDot1qRange = ""
	}

	if !v.VrfLiteAutoConfig.IsNull() && !v.VrfLiteAutoConfig.IsUnknown() {
		data.Management.VrfLiteAutoConfig = v.VrfLiteAutoConfig.ValueString()
	} else {
		data.Management.VrfLiteAutoConfig = ""
	}

	if !v.VrfLiteSubnetRange.IsNull() && !v.VrfLiteSubnetRange.IsUnknown() {
		data.Management.VrfLiteSubnetRange = v.VrfLiteSubnetRange.ValueString()
	} else {
		data.Management.VrfLiteSubnetRange = ""
	}

	if !v.VrfLiteSubnetTargetMask.IsNull() && !v.VrfLiteSubnetTargetMask.IsUnknown() {
		data.Management.VrfLiteSubnetTargetMask = new(int64)
		*data.Management.VrfLiteSubnetTargetMask = v.VrfLiteSubnetTargetMask.ValueInt64()

	} else {
		data.Management.VrfLiteSubnetTargetMask = nil
	}

	if !v.VrfLiteIpv6SubnetRange.IsNull() && !v.VrfLiteIpv6SubnetRange.IsUnknown() {
		data.Management.VrfLiteIpv6SubnetRange = v.VrfLiteIpv6SubnetRange.ValueString()
	} else {
		data.Management.VrfLiteIpv6SubnetRange = ""
	}

	if !v.VrfLiteIpv6SubnetTargetMask.IsNull() && !v.VrfLiteIpv6SubnetTargetMask.IsUnknown() {
		data.Management.VrfLiteIpv6SubnetTargetMask = new(int64)
		*data.Management.VrfLiteIpv6SubnetTargetMask = v.VrfLiteIpv6SubnetTargetMask.ValueInt64()

	} else {
		data.Management.VrfLiteIpv6SubnetTargetMask = nil
	}

	if !v.AutoUniqueVrfLiteIpPrefix.IsNull() && !v.AutoUniqueVrfLiteIpPrefix.IsUnknown() {
		data.Management.AutoUniqueVrfLiteIpPrefix = new(bool)
		*data.Management.AutoUniqueVrfLiteIpPrefix = v.AutoUniqueVrfLiteIpPrefix.ValueBool()
	} else {
		data.Management.AutoUniqueVrfLiteIpPrefix = nil
	}

	if !v.PerVrfLoopbackAutoProvision.IsNull() && !v.PerVrfLoopbackAutoProvision.IsUnknown() {
		data.Management.PerVrfLoopbackAutoProvision = new(bool)
		*data.Management.PerVrfLoopbackAutoProvision = v.PerVrfLoopbackAutoProvision.ValueBool()
	} else {
		data.Management.PerVrfLoopbackAutoProvision = nil
	}

	if !v.PerVrfLoopbackIpRange.IsNull() && !v.PerVrfLoopbackIpRange.IsUnknown() {
		data.Management.PerVrfLoopbackIpRange = v.PerVrfLoopbackIpRange.ValueString()
	} else {
		data.Management.PerVrfLoopbackIpRange = ""
	}

	if !v.PerVrfLoopbackAutoProvisionIpv6.IsNull() && !v.PerVrfLoopbackAutoProvisionIpv6.IsUnknown() {
		data.Management.PerVrfLoopbackAutoProvisionIpv6 = new(bool)
		*data.Management.PerVrfLoopbackAutoProvisionIpv6 = v.PerVrfLoopbackAutoProvisionIpv6.ValueBool()
	} else {
		data.Management.PerVrfLoopbackAutoProvisionIpv6 = nil
	}

	if !v.PerVrfLoopbackIpv6Range.IsNull() && !v.PerVrfLoopbackIpv6Range.IsUnknown() {
		data.Management.PerVrfLoopbackIpv6Range = v.PerVrfLoopbackIpv6Range.ValueString()
	} else {
		data.Management.PerVrfLoopbackIpv6Range = ""
	}

	if !v.Banner.IsNull() && !v.Banner.IsUnknown() {
		data.Management.Banner = v.Banner.ValueString()
	} else {
		data.Management.Banner = ""
	}

	if !v.Day0Bootstrap.IsNull() && !v.Day0Bootstrap.IsUnknown() {
		data.Management.Day0Bootstrap = new(bool)
		*data.Management.Day0Bootstrap = v.Day0Bootstrap.ValueBool()
	} else {
		data.Management.Day0Bootstrap = nil
	}

	if !v.LocalDhcpServer.IsNull() && !v.LocalDhcpServer.IsUnknown() {
		data.Management.LocalDhcpServer = new(bool)
		*data.Management.LocalDhcpServer = v.LocalDhcpServer.ValueBool()
	} else {
		data.Management.LocalDhcpServer = nil
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

	if !v.DhcpEndAddress.IsNull() && !v.DhcpEndAddress.IsUnknown() {
		data.Management.DhcpEndAddress = v.DhcpEndAddress.ValueString()
	} else {
		data.Management.DhcpEndAddress = ""
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

	if !v.BootstrapMultiSubnet.IsNull() && !v.BootstrapMultiSubnet.IsUnknown() {
		data.Management.BootstrapMultiSubnet = v.BootstrapMultiSubnet.ValueString()
	} else {
		data.Management.BootstrapMultiSubnet = ""
	}

	if !v.ExtraConfigNxosBootstrap.IsNull() && !v.ExtraConfigNxosBootstrap.IsUnknown() {
		data.Management.ExtraConfigNxosBootstrap = v.ExtraConfigNxosBootstrap.ValueString()
	} else {
		data.Management.ExtraConfigNxosBootstrap = ""
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

	if !v.UnderlayIpv6.IsNull() && !v.UnderlayIpv6.IsUnknown() {
		data.Management.UnderlayIpv6 = new(bool)
		*data.Management.UnderlayIpv6 = v.UnderlayIpv6.ValueBool()
	} else {
		data.Management.UnderlayIpv6 = nil
	}

	if !v.Ipv6MulticastGroupSubnet.IsNull() && !v.Ipv6MulticastGroupSubnet.IsUnknown() {
		data.Management.Ipv6MulticastGroupSubnet = v.Ipv6MulticastGroupSubnet.ValueString()
	} else {
		data.Management.Ipv6MulticastGroupSubnet = ""
	}

	if !v.TenantRoutedMulticastIpv6.IsNull() && !v.TenantRoutedMulticastIpv6.IsUnknown() {
		data.Management.TenantRoutedMulticastIpv6 = new(bool)
		*data.Management.TenantRoutedMulticastIpv6 = v.TenantRoutedMulticastIpv6.ValueBool()
	} else {
		data.Management.TenantRoutedMulticastIpv6 = nil
	}

	if !v.MvpnVrfRouteImportId.IsNull() && !v.MvpnVrfRouteImportId.IsUnknown() {
		data.Management.MvpnVrfRouteImportId = new(bool)
		*data.Management.MvpnVrfRouteImportId = v.MvpnVrfRouteImportId.ValueBool()
	} else {
		data.Management.MvpnVrfRouteImportId = nil
	}

	if !v.MvpnVrfRouteImportIdRange.IsNull() && !v.MvpnVrfRouteImportIdRange.IsUnknown() {
		data.Management.MvpnVrfRouteImportIdRange = v.MvpnVrfRouteImportIdRange.ValueString()
	} else {
		data.Management.MvpnVrfRouteImportIdRange = ""
	}

	if !v.VrfRouteImportIdReallocation.IsNull() && !v.VrfRouteImportIdReallocation.IsUnknown() {
		data.Management.VrfRouteImportIdReallocation = new(bool)
		*data.Management.VrfRouteImportIdReallocation = v.VrfRouteImportIdReallocation.ValueBool()
	} else {
		data.Management.VrfRouteImportIdReallocation = nil
	}

	if !v.L3vniMulticastGroup.IsNull() && !v.L3vniMulticastGroup.IsUnknown() {
		data.Management.L3vniMulticastGroup = v.L3vniMulticastGroup.ValueString()
	} else {
		data.Management.L3vniMulticastGroup = ""
	}

	if !v.L3VniIpv6MulticastGroup.IsNull() && !v.L3VniIpv6MulticastGroup.IsUnknown() {
		data.Management.L3VniIpv6MulticastGroup = v.L3VniIpv6MulticastGroup.ValueString()
	} else {
		data.Management.L3VniIpv6MulticastGroup = ""
	}

	if !v.RendezvousPointMode.IsNull() && !v.RendezvousPointMode.IsUnknown() {
		data.Management.RendezvousPointMode = v.RendezvousPointMode.ValueString()
	} else {
		data.Management.RendezvousPointMode = ""
	}

	if !v.AutoGenerateMulticastGroupAddress.IsNull() && !v.AutoGenerateMulticastGroupAddress.IsUnknown() {
		data.Management.AutoGenerateMulticastGroupAddress = new(bool)
		*data.Management.AutoGenerateMulticastGroupAddress = v.AutoGenerateMulticastGroupAddress.ValueBool()
	} else {
		data.Management.AutoGenerateMulticastGroupAddress = nil
	}

	if !v.PhantomRendezvousPointLoopbackId1.IsNull() && !v.PhantomRendezvousPointLoopbackId1.IsUnknown() {
		data.Management.PhantomRendezvousPointLoopbackId1 = new(int64)
		*data.Management.PhantomRendezvousPointLoopbackId1 = v.PhantomRendezvousPointLoopbackId1.ValueInt64()

	} else {
		data.Management.PhantomRendezvousPointLoopbackId1 = nil
	}

	if !v.PhantomRendezvousPointLoopbackId2.IsNull() && !v.PhantomRendezvousPointLoopbackId2.IsUnknown() {
		data.Management.PhantomRendezvousPointLoopbackId2 = new(int64)
		*data.Management.PhantomRendezvousPointLoopbackId2 = v.PhantomRendezvousPointLoopbackId2.ValueInt64()

	} else {
		data.Management.PhantomRendezvousPointLoopbackId2 = nil
	}

	if !v.PhantomRendezvousPointLoopbackId3.IsNull() && !v.PhantomRendezvousPointLoopbackId3.IsUnknown() {
		data.Management.PhantomRendezvousPointLoopbackId3 = new(int64)
		*data.Management.PhantomRendezvousPointLoopbackId3 = v.PhantomRendezvousPointLoopbackId3.ValueInt64()

	} else {
		data.Management.PhantomRendezvousPointLoopbackId3 = nil
	}

	if !v.PhantomRendezvousPointLoopbackId4.IsNull() && !v.PhantomRendezvousPointLoopbackId4.IsUnknown() {
		data.Management.PhantomRendezvousPointLoopbackId4 = new(int64)
		*data.Management.PhantomRendezvousPointLoopbackId4 = v.PhantomRendezvousPointLoopbackId4.ValueInt64()

	} else {
		data.Management.PhantomRendezvousPointLoopbackId4 = nil
	}

	if !v.AdvertisePhysicalIpOnBorder.IsNull() && !v.AdvertisePhysicalIpOnBorder.IsUnknown() {
		data.Management.AdvertisePhysicalIpOnBorder = new(bool)
		*data.Management.AdvertisePhysicalIpOnBorder = v.AdvertisePhysicalIpOnBorder.ValueBool()
	} else {
		data.Management.AdvertisePhysicalIpOnBorder = nil
	}

	if !v.FabricVpcDomainId.IsNull() && !v.FabricVpcDomainId.IsUnknown() {
		data.Management.FabricVpcDomainId = new(bool)
		*data.Management.FabricVpcDomainId = v.FabricVpcDomainId.ValueBool()
	} else {
		data.Management.FabricVpcDomainId = nil
	}

	if !v.SharedVpcDomainId.IsNull() && !v.SharedVpcDomainId.IsUnknown() {
		data.Management.SharedVpcDomainId = new(int64)
		*data.Management.SharedVpcDomainId = v.SharedVpcDomainId.ValueInt64()

	} else {
		data.Management.SharedVpcDomainId = nil
	}

	if !v.VpcLayer3PeerRouter.IsNull() && !v.VpcLayer3PeerRouter.IsUnknown() {
		data.Management.VpcLayer3PeerRouter = new(bool)
		*data.Management.VpcLayer3PeerRouter = v.VpcLayer3PeerRouter.ValueBool()
	} else {
		data.Management.VpcLayer3PeerRouter = nil
	}

	if !v.FabricVpcQos.IsNull() && !v.FabricVpcQos.IsUnknown() {
		data.Management.FabricVpcQos = new(bool)
		*data.Management.FabricVpcQos = v.FabricVpcQos.ValueBool()
	} else {
		data.Management.FabricVpcQos = nil
	}

	if !v.FabricVpcQosPolicyName.IsNull() && !v.FabricVpcQosPolicyName.IsUnknown() {
		data.Management.FabricVpcQosPolicyName = v.FabricVpcQosPolicyName.ValueString()
	} else {
		data.Management.FabricVpcQosPolicyName = ""
	}

	if !v.AnycastLoopbackId.IsNull() && !v.AnycastLoopbackId.IsUnknown() {
		data.Management.AnycastLoopbackId = new(int64)
		*data.Management.AnycastLoopbackId = v.AnycastLoopbackId.ValueInt64()

	} else {
		data.Management.AnycastLoopbackId = nil
	}

	if !v.BgpAuthentication.IsNull() && !v.BgpAuthentication.IsUnknown() {
		data.Management.BgpAuthentication = new(bool)
		*data.Management.BgpAuthentication = v.BgpAuthentication.ValueBool()
	} else {
		data.Management.BgpAuthentication = nil
	}

	if !v.BgpAuthenticationKeyType.IsNull() && !v.BgpAuthenticationKeyType.IsUnknown() {
		data.Management.BgpAuthenticationKeyType = v.BgpAuthenticationKeyType.ValueString()
	} else {
		data.Management.BgpAuthenticationKeyType = ""
	}

	if !v.BgpAuthenticationKey.IsNull() && !v.BgpAuthenticationKey.IsUnknown() {
		data.Management.BgpAuthenticationKey = v.BgpAuthenticationKey.ValueString()
	} else {
		data.Management.BgpAuthenticationKey = ""
	}

	if !v.PimHelloAuthentication.IsNull() && !v.PimHelloAuthentication.IsUnknown() {
		data.Management.PimHelloAuthentication = new(bool)
		*data.Management.PimHelloAuthentication = v.PimHelloAuthentication.ValueBool()
	} else {
		data.Management.PimHelloAuthentication = nil
	}

	if !v.PimHelloAuthenticationKey.IsNull() && !v.PimHelloAuthenticationKey.IsUnknown() {
		data.Management.PimHelloAuthenticationKey = v.PimHelloAuthenticationKey.ValueString()
	} else {
		data.Management.PimHelloAuthenticationKey = ""
	}

	if !v.Bfd.IsNull() && !v.Bfd.IsUnknown() {
		data.Management.Bfd = new(bool)
		*data.Management.Bfd = v.Bfd.ValueBool()
	} else {
		data.Management.Bfd = nil
	}

	if !v.BfdIbgp.IsNull() && !v.BfdIbgp.IsUnknown() {
		data.Management.BfdIbgp = new(bool)
		*data.Management.BfdIbgp = v.BfdIbgp.ValueBool()
	} else {
		data.Management.BfdIbgp = nil
	}

	if !v.BfdAuthentication.IsNull() && !v.BfdAuthentication.IsUnknown() {
		data.Management.BfdAuthentication = new(bool)
		*data.Management.BfdAuthentication = v.BfdAuthentication.ValueBool()
	} else {
		data.Management.BfdAuthentication = nil
	}

	if !v.BfdAuthenticationKeyId.IsNull() && !v.BfdAuthenticationKeyId.IsUnknown() {
		data.Management.BfdAuthenticationKeyId = new(int64)
		*data.Management.BfdAuthenticationKeyId = v.BfdAuthenticationKeyId.ValueInt64()

	} else {
		data.Management.BfdAuthenticationKeyId = nil
	}

	if !v.BfdAuthenticationKey.IsNull() && !v.BfdAuthenticationKey.IsUnknown() {
		data.Management.BfdAuthenticationKey = v.BfdAuthenticationKey.ValueString()
	} else {
		data.Management.BfdAuthenticationKey = ""
	}

	if !v.Macsec.IsNull() && !v.Macsec.IsUnknown() {
		data.Management.Macsec = new(bool)
		*data.Management.Macsec = v.Macsec.ValueBool()
	} else {
		data.Management.Macsec = nil
	}

	if !v.MacsecCipherSuite.IsNull() && !v.MacsecCipherSuite.IsUnknown() {
		data.Management.MacsecCipherSuite = v.MacsecCipherSuite.ValueString()
	} else {
		data.Management.MacsecCipherSuite = ""
	}

	if !v.MacsecKeyString.IsNull() && !v.MacsecKeyString.IsUnknown() {
		data.Management.MacsecKeyString = v.MacsecKeyString.ValueString()
	} else {
		data.Management.MacsecKeyString = ""
	}

	if !v.MacsecAlgorithm.IsNull() && !v.MacsecAlgorithm.IsUnknown() {
		data.Management.MacsecAlgorithm = v.MacsecAlgorithm.ValueString()
	} else {
		data.Management.MacsecAlgorithm = ""
	}

	if !v.MacsecFallbackKeyString.IsNull() && !v.MacsecFallbackKeyString.IsUnknown() {
		data.Management.MacsecFallbackKeyString = v.MacsecFallbackKeyString.ValueString()
	} else {
		data.Management.MacsecFallbackKeyString = ""
	}

	if !v.MacsecFallbackAlgorithm.IsNull() && !v.MacsecFallbackAlgorithm.IsUnknown() {
		data.Management.MacsecFallbackAlgorithm = v.MacsecFallbackAlgorithm.ValueString()
	} else {
		data.Management.MacsecFallbackAlgorithm = ""
	}

	if !v.MacsecReportTimer.IsNull() && !v.MacsecReportTimer.IsUnknown() {
		data.Management.MacsecReportTimer = new(int64)
		*data.Management.MacsecReportTimer = v.MacsecReportTimer.ValueInt64()

	} else {
		data.Management.MacsecReportTimer = nil
	}

	if !v.OverlayMode.IsNull() && !v.OverlayMode.IsUnknown() {
		data.Management.OverlayMode = v.OverlayMode.ValueString()
	} else {
		data.Management.OverlayMode = ""
	}

	if !v.PrivateVlan.IsNull() && !v.PrivateVlan.IsUnknown() {
		data.Management.PrivateVlan = new(bool)
		*data.Management.PrivateVlan = v.PrivateVlan.ValueBool()
	} else {
		data.Management.PrivateVlan = nil
	}

	if !v.DefaultPrivateVlanSecondaryNetworkTemplate.IsNull() && !v.DefaultPrivateVlanSecondaryNetworkTemplate.IsUnknown() {
		data.Management.DefaultPrivateVlanSecondaryNetworkTemplate = v.DefaultPrivateVlanSecondaryNetworkTemplate.ValueString()
	} else {
		data.Management.DefaultPrivateVlanSecondaryNetworkTemplate = ""
	}

	if !v.PowerRedundancyMode.IsNull() && !v.PowerRedundancyMode.IsUnknown() {
		data.Management.PowerRedundancyMode = v.PowerRedundancyMode.ValueString()
	} else {
		data.Management.PowerRedundancyMode = ""
	}

	if !v.CoppPolicy.IsNull() && !v.CoppPolicy.IsUnknown() {
		data.Management.CoppPolicy = v.CoppPolicy.ValueString()
	} else {
		data.Management.CoppPolicy = ""
	}

	if !v.NveHoldDownTimer.IsNull() && !v.NveHoldDownTimer.IsUnknown() {
		data.Management.NveHoldDownTimer = new(int64)
		*data.Management.NveHoldDownTimer = v.NveHoldDownTimer.ValueInt64()

	} else {
		data.Management.NveHoldDownTimer = nil
	}

	if !v.Cdp.IsNull() && !v.Cdp.IsUnknown() {
		data.Management.Cdp = new(bool)
		*data.Management.Cdp = v.Cdp.ValueBool()
	} else {
		data.Management.Cdp = nil
	}

	if !v.NextGenerationOam.IsNull() && !v.NextGenerationOam.IsUnknown() {
		data.Management.NextGenerationOam = new(bool)
		*data.Management.NextGenerationOam = v.NextGenerationOam.ValueBool()
	} else {
		data.Management.NextGenerationOam = nil
	}

	if !v.NgoamSouthBoundLoopDetect.IsNull() && !v.NgoamSouthBoundLoopDetect.IsUnknown() {
		data.Management.NgoamSouthBoundLoopDetect = new(bool)
		*data.Management.NgoamSouthBoundLoopDetect = v.NgoamSouthBoundLoopDetect.ValueBool()
	} else {
		data.Management.NgoamSouthBoundLoopDetect = nil
	}

	if !v.NgoamSouthBoundLoopDetectProbeInterval.IsNull() && !v.NgoamSouthBoundLoopDetectProbeInterval.IsUnknown() {
		data.Management.NgoamSouthBoundLoopDetectProbeInterval = new(int64)
		*data.Management.NgoamSouthBoundLoopDetectProbeInterval = v.NgoamSouthBoundLoopDetectProbeInterval.ValueInt64()

	} else {
		data.Management.NgoamSouthBoundLoopDetectProbeInterval = nil
	}

	if !v.NgoamSouthBoundLoopDetectRecoveryInterval.IsNull() && !v.NgoamSouthBoundLoopDetectRecoveryInterval.IsUnknown() {
		data.Management.NgoamSouthBoundLoopDetectRecoveryInterval = new(int64)
		*data.Management.NgoamSouthBoundLoopDetectRecoveryInterval = v.NgoamSouthBoundLoopDetectRecoveryInterval.ValueInt64()

	} else {
		data.Management.NgoamSouthBoundLoopDetectRecoveryInterval = nil
	}

	if !v.StrictConfigComplianceMode.IsNull() && !v.StrictConfigComplianceMode.IsUnknown() {
		data.Management.StrictConfigComplianceMode = new(bool)
		*data.Management.StrictConfigComplianceMode = v.StrictConfigComplianceMode.ValueBool()
	} else {
		data.Management.StrictConfigComplianceMode = nil
	}

	if !v.AdvancedSshOption.IsNull() && !v.AdvancedSshOption.IsUnknown() {
		data.Management.AdvancedSshOption = new(bool)
		*data.Management.AdvancedSshOption = v.AdvancedSshOption.ValueBool()
	} else {
		data.Management.AdvancedSshOption = nil
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

	if !v.DefaultQueuingPolicy.IsNull() && !v.DefaultQueuingPolicy.IsUnknown() {
		data.Management.DefaultQueuingPolicy = new(bool)
		*data.Management.DefaultQueuingPolicy = v.DefaultQueuingPolicy.ValueBool()
	} else {
		data.Management.DefaultQueuingPolicy = nil
	}

	if !v.DefaultQueuingPolicyCloudscale.IsNull() && !v.DefaultQueuingPolicyCloudscale.IsUnknown() {
		data.Management.DefaultQueuingPolicyCloudscale = v.DefaultQueuingPolicyCloudscale.ValueString()
	} else {
		data.Management.DefaultQueuingPolicyCloudscale = ""
	}

	if !v.DefaultQueuingPolicyRSeries.IsNull() && !v.DefaultQueuingPolicyRSeries.IsUnknown() {
		data.Management.DefaultQueuingPolicyRSeries = v.DefaultQueuingPolicyRSeries.ValueString()
	} else {
		data.Management.DefaultQueuingPolicyRSeries = ""
	}

	if !v.DefaultQueuingPolicyOther.IsNull() && !v.DefaultQueuingPolicyOther.IsUnknown() {
		data.Management.DefaultQueuingPolicyOther = v.DefaultQueuingPolicyOther.ValueString()
	} else {
		data.Management.DefaultQueuingPolicyOther = ""
	}

	if !v.AimlQos.IsNull() && !v.AimlQos.IsUnknown() {
		data.Management.AimlQos = new(bool)
		*data.Management.AimlQos = v.AimlQos.ValueBool()
	} else {
		data.Management.AimlQos = nil
	}

	if !v.AimlQosPolicy.IsNull() && !v.AimlQosPolicy.IsUnknown() {
		data.Management.AimlQosPolicy = v.AimlQosPolicy.ValueString()
	} else {
		data.Management.AimlQosPolicy = ""
	}

	if !v.PriorityFlowControlWatchInterval.IsNull() && !v.PriorityFlowControlWatchInterval.IsUnknown() {
		data.Management.PriorityFlowControlWatchInterval = new(int64)
		*data.Management.PriorityFlowControlWatchInterval = v.PriorityFlowControlWatchInterval.ValueInt64()

	} else {
		data.Management.PriorityFlowControlWatchInterval = nil
	}

	if !v.StaticUnderlayIpAllocation.IsNull() && !v.StaticUnderlayIpAllocation.IsUnknown() {
		data.Management.StaticUnderlayIpAllocation = new(bool)
		*data.Management.StaticUnderlayIpAllocation = v.StaticUnderlayIpAllocation.ValueBool()
	} else {
		data.Management.StaticUnderlayIpAllocation = nil
	}

	if !v.BgpLoopbackIpv6Range.IsNull() && !v.BgpLoopbackIpv6Range.IsUnknown() {
		data.Management.BgpLoopbackIpv6Range = v.BgpLoopbackIpv6Range.ValueString()
	} else {
		data.Management.BgpLoopbackIpv6Range = ""
	}

	if !v.NveLoopbackIpv6Range.IsNull() && !v.NveLoopbackIpv6Range.IsUnknown() {
		data.Management.NveLoopbackIpv6Range = v.NveLoopbackIpv6Range.ValueString()
	} else {
		data.Management.NveLoopbackIpv6Range = ""
	}

	if !v.ExtraConfigAaa.IsNull() && !v.ExtraConfigAaa.IsUnknown() {
		data.Management.ExtraConfigAaa = v.ExtraConfigAaa.ValueString()
	} else {
		data.Management.ExtraConfigAaa = ""
	}

	if !v.Aaa.IsNull() && !v.Aaa.IsUnknown() {
		data.Management.Aaa = new(bool)
		*data.Management.Aaa = v.Aaa.ValueBool()
	} else {
		data.Management.Aaa = nil
	}

	if !v.Ipv6LinkLocal.IsNull() && !v.Ipv6LinkLocal.IsUnknown() {
		data.Management.Ipv6LinkLocal = new(bool)
		*data.Management.Ipv6LinkLocal = v.Ipv6LinkLocal.ValueBool()
	} else {
		data.Management.Ipv6LinkLocal = nil
	}

	if !v.FabricInterfaceType.IsNull() && !v.FabricInterfaceType.IsUnknown() {
		data.Management.FabricInterfaceType = v.FabricInterfaceType.ValueString()
	} else {
		data.Management.FabricInterfaceType = ""
	}

	if !v.Ipv6SubnetTargetMask.IsNull() && !v.Ipv6SubnetTargetMask.IsUnknown() {
		data.Management.Ipv6SubnetTargetMask = new(int64)
		*data.Management.Ipv6SubnetTargetMask = v.Ipv6SubnetTargetMask.ValueInt64()

	} else {
		data.Management.Ipv6SubnetTargetMask = nil
	}

	if !v.LinkStateRoutingProtocol.IsNull() && !v.LinkStateRoutingProtocol.IsUnknown() {
		data.Management.LinkStateRoutingProtocol = v.LinkStateRoutingProtocol.ValueString()
	} else {
		data.Management.LinkStateRoutingProtocol = ""
	}

	if !v.RouteReflectorCount.IsNull() && !v.RouteReflectorCount.IsUnknown() {
		data.Management.RouteReflectorCount = new(int64)
		*data.Management.RouteReflectorCount = v.RouteReflectorCount.ValueInt64()

	} else {
		data.Management.RouteReflectorCount = nil
	}

	if !v.VpcTorDelayRestoreTimer.IsNull() && !v.VpcTorDelayRestoreTimer.IsUnknown() {
		data.Management.VpcTorDelayRestoreTimer = new(int64)
		*data.Management.VpcTorDelayRestoreTimer = v.VpcTorDelayRestoreTimer.ValueInt64()

	} else {
		data.Management.VpcTorDelayRestoreTimer = nil
	}

	if !v.LeafToRIdRange.IsNull() && !v.LeafToRIdRange.IsUnknown() {
		data.Management.LeafToRIdRange = new(bool)
		*data.Management.LeafToRIdRange = v.LeafToRIdRange.ValueBool()
	} else {
		data.Management.LeafToRIdRange = nil
	}

	if !v.LeafTorVpcPortChannelIdRange.IsNull() && !v.LeafTorVpcPortChannelIdRange.IsUnknown() {
		data.Management.LeafTorVpcPortChannelIdRange = v.LeafTorVpcPortChannelIdRange.ValueString()
	} else {
		data.Management.LeafTorVpcPortChannelIdRange = ""
	}

	if !v.LinkStateRoutingTag.IsNull() && !v.LinkStateRoutingTag.IsUnknown() {
		data.Management.LinkStateRoutingTag = v.LinkStateRoutingTag.ValueString()
	} else {
		data.Management.LinkStateRoutingTag = ""
	}

	if !v.OspfAreaId.IsNull() && !v.OspfAreaId.IsUnknown() {
		data.Management.OspfAreaId = v.OspfAreaId.ValueString()
	} else {
		data.Management.OspfAreaId = ""
	}

	if !v.OspfAuthentication.IsNull() && !v.OspfAuthentication.IsUnknown() {
		data.Management.OspfAuthentication = new(bool)
		*data.Management.OspfAuthentication = v.OspfAuthentication.ValueBool()
	} else {
		data.Management.OspfAuthentication = nil
	}

	if !v.OspfAuthenticationKeyId.IsNull() && !v.OspfAuthenticationKeyId.IsUnknown() {
		data.Management.OspfAuthenticationKeyId = new(int64)
		*data.Management.OspfAuthenticationKeyId = v.OspfAuthenticationKeyId.ValueInt64()

	} else {
		data.Management.OspfAuthenticationKeyId = nil
	}

	if !v.OspfAuthenticationKey.IsNull() && !v.OspfAuthenticationKey.IsUnknown() {
		data.Management.OspfAuthenticationKey = v.OspfAuthenticationKey.ValueString()
	} else {
		data.Management.OspfAuthenticationKey = ""
	}

	if !v.IsisLevel.IsNull() && !v.IsisLevel.IsUnknown() {
		data.Management.IsisLevel = v.IsisLevel.ValueString()
	} else {
		data.Management.IsisLevel = ""
	}

	if !v.IsisAreaNumber.IsNull() && !v.IsisAreaNumber.IsUnknown() {
		data.Management.IsisAreaNumber = v.IsisAreaNumber.ValueString()
	} else {
		data.Management.IsisAreaNumber = ""
	}

	if !v.IsisPointToPoint.IsNull() && !v.IsisPointToPoint.IsUnknown() {
		data.Management.IsisPointToPoint = new(bool)
		*data.Management.IsisPointToPoint = v.IsisPointToPoint.ValueBool()
	} else {
		data.Management.IsisPointToPoint = nil
	}

	if !v.IsisAuthentication.IsNull() && !v.IsisAuthentication.IsUnknown() {
		data.Management.IsisAuthentication = new(bool)
		*data.Management.IsisAuthentication = v.IsisAuthentication.ValueBool()
	} else {
		data.Management.IsisAuthentication = nil
	}

	if !v.IsisAuthenticationKeychainName.IsNull() && !v.IsisAuthenticationKeychainName.IsUnknown() {
		data.Management.IsisAuthenticationKeychainName = v.IsisAuthenticationKeychainName.ValueString()
	} else {
		data.Management.IsisAuthenticationKeychainName = ""
	}

	if !v.IsisAuthenticationKeychainKeyId.IsNull() && !v.IsisAuthenticationKeychainKeyId.IsUnknown() {
		data.Management.IsisAuthenticationKeychainKeyId = new(int64)
		*data.Management.IsisAuthenticationKeychainKeyId = v.IsisAuthenticationKeychainKeyId.ValueInt64()

	} else {
		data.Management.IsisAuthenticationKeychainKeyId = nil
	}

	if !v.IsisAuthenticationKey.IsNull() && !v.IsisAuthenticationKey.IsUnknown() {
		data.Management.IsisAuthenticationKey = v.IsisAuthenticationKey.ValueString()
	} else {
		data.Management.IsisAuthenticationKey = ""
	}

	if !v.IsisOverload.IsNull() && !v.IsisOverload.IsUnknown() {
		data.Management.IsisOverload = new(bool)
		*data.Management.IsisOverload = v.IsisOverload.ValueBool()
	} else {
		data.Management.IsisOverload = nil
	}

	if !v.IsisOverloadElapseTime.IsNull() && !v.IsisOverloadElapseTime.IsUnknown() {
		data.Management.IsisOverloadElapseTime = new(int64)
		*data.Management.IsisOverloadElapseTime = v.IsisOverloadElapseTime.ValueInt64()

	} else {
		data.Management.IsisOverloadElapseTime = nil
	}

	if !v.BfdOspf.IsNull() && !v.BfdOspf.IsUnknown() {
		data.Management.BfdOspf = new(bool)
		*data.Management.BfdOspf = v.BfdOspf.ValueBool()
	} else {
		data.Management.BfdOspf = nil
	}

	if !v.BfdIsis.IsNull() && !v.BfdIsis.IsUnknown() {
		data.Management.BfdIsis = new(bool)
		*data.Management.BfdIsis = v.BfdIsis.ValueBool()
	} else {
		data.Management.BfdIsis = nil
	}

	if !v.BfdPim.IsNull() && !v.BfdPim.IsUnknown() {
		data.Management.BfdPim = new(bool)
		*data.Management.BfdPim = v.BfdPim.ValueBool()
	} else {
		data.Management.BfdPim = nil
	}

	if !v.AutoBgpNeighborDescription.IsNull() && !v.AutoBgpNeighborDescription.IsUnknown() {
		data.Management.AutoBgpNeighborDescription = new(bool)
		*data.Management.AutoBgpNeighborDescription = v.AutoBgpNeighborDescription.ValueBool()
	} else {
		data.Management.AutoBgpNeighborDescription = nil
	}

	if !v.IbgpPeerTemplate.IsNull() && !v.IbgpPeerTemplate.IsUnknown() {
		data.Management.IbgpPeerTemplate = v.IbgpPeerTemplate.ValueString()
	} else {
		data.Management.IbgpPeerTemplate = ""
	}

	if !v.LeafibgpPeerTemplate.IsNull() && !v.LeafibgpPeerTemplate.IsUnknown() {
		data.Management.LeafibgpPeerTemplate = v.LeafibgpPeerTemplate.ValueString()
	} else {
		data.Management.LeafibgpPeerTemplate = ""
	}

	if !v.SecurityGroupTag.IsNull() && !v.SecurityGroupTag.IsUnknown() {
		data.Management.SecurityGroupTag = new(bool)
		*data.Management.SecurityGroupTag = v.SecurityGroupTag.ValueBool()
	} else {
		data.Management.SecurityGroupTag = nil
	}

	if !v.SecurityGroupTagPrefix.IsNull() && !v.SecurityGroupTagPrefix.IsUnknown() {
		data.Management.SecurityGroupTagPrefix = v.SecurityGroupTagPrefix.ValueString()
	} else {
		data.Management.SecurityGroupTagPrefix = ""
	}

	if !v.SecurityGroupTagIdRange.IsNull() && !v.SecurityGroupTagIdRange.IsUnknown() {
		data.Management.SecurityGroupTagIdRange = v.SecurityGroupTagIdRange.ValueString()
	} else {
		data.Management.SecurityGroupTagIdRange = ""
	}

	if !v.SecurityGroupTagPreprovision.IsNull() && !v.SecurityGroupTagPreprovision.IsUnknown() {
		data.Management.SecurityGroupTagPreprovision = new(bool)
		*data.Management.SecurityGroupTagPreprovision = v.SecurityGroupTagPreprovision.ValueBool()
	} else {
		data.Management.SecurityGroupTagPreprovision = nil
	}

	if !v.SecurityGroupStatus.IsNull() && !v.SecurityGroupStatus.IsUnknown() {
		data.Management.SecurityGroupStatus = v.SecurityGroupStatus.ValueString()
	} else {
		data.Management.SecurityGroupStatus = ""
	}

	if !v.VrfLiteMacsec.IsNull() && !v.VrfLiteMacsec.IsUnknown() {
		data.Management.VrfLiteMacsec = new(bool)
		*data.Management.VrfLiteMacsec = v.VrfLiteMacsec.ValueBool()
	} else {
		data.Management.VrfLiteMacsec = nil
	}

	if !v.QuantumKeyDistribution.IsNull() && !v.QuantumKeyDistribution.IsUnknown() {
		data.Management.QuantumKeyDistribution = new(bool)
		*data.Management.QuantumKeyDistribution = v.QuantumKeyDistribution.ValueBool()
	} else {
		data.Management.QuantumKeyDistribution = nil
	}

	if !v.VrfLiteMacsecCipherSuite.IsNull() && !v.VrfLiteMacsecCipherSuite.IsUnknown() {
		data.Management.VrfLiteMacsecCipherSuite = v.VrfLiteMacsecCipherSuite.ValueString()
	} else {
		data.Management.VrfLiteMacsecCipherSuite = ""
	}

	if !v.VrfLiteMacsecKeyString.IsNull() && !v.VrfLiteMacsecKeyString.IsUnknown() {
		data.Management.VrfLiteMacsecKeyString = v.VrfLiteMacsecKeyString.ValueString()
	} else {
		data.Management.VrfLiteMacsecKeyString = ""
	}

	if !v.VrfLiteMacsecAlgorithm.IsNull() && !v.VrfLiteMacsecAlgorithm.IsUnknown() {
		data.Management.VrfLiteMacsecAlgorithm = v.VrfLiteMacsecAlgorithm.ValueString()
	} else {
		data.Management.VrfLiteMacsecAlgorithm = ""
	}

	if !v.VrfLiteMacsecFallbackKeyString.IsNull() && !v.VrfLiteMacsecFallbackKeyString.IsUnknown() {
		data.Management.VrfLiteMacsecFallbackKeyString = v.VrfLiteMacsecFallbackKeyString.ValueString()
	} else {
		data.Management.VrfLiteMacsecFallbackKeyString = ""
	}

	if !v.VrfLiteMacsecFallbackAlgorithm.IsNull() && !v.VrfLiteMacsecFallbackAlgorithm.IsUnknown() {
		data.Management.VrfLiteMacsecFallbackAlgorithm = v.VrfLiteMacsecFallbackAlgorithm.ValueString()
	} else {
		data.Management.VrfLiteMacsecFallbackAlgorithm = ""
	}

	if !v.QuantumKeyDistributionProfileName.IsNull() && !v.QuantumKeyDistributionProfileName.IsUnknown() {
		data.Management.QuantumKeyDistributionProfileName = v.QuantumKeyDistributionProfileName.ValueString()
	} else {
		data.Management.QuantumKeyDistributionProfileName = ""
	}

	if !v.KeyManagementEntityServerIp.IsNull() && !v.KeyManagementEntityServerIp.IsUnknown() {
		data.Management.KeyManagementEntityServerIp = v.KeyManagementEntityServerIp.ValueString()
	} else {
		data.Management.KeyManagementEntityServerIp = ""
	}

	if !v.KeyManagementEntityServerPort.IsNull() && !v.KeyManagementEntityServerPort.IsUnknown() {
		data.Management.KeyManagementEntityServerPort = new(int64)
		*data.Management.KeyManagementEntityServerPort = v.KeyManagementEntityServerPort.ValueInt64()

	} else {
		data.Management.KeyManagementEntityServerPort = nil
	}

	if !v.TrustpointLabel.IsNull() && !v.TrustpointLabel.IsUnknown() {
		data.Management.TrustpointLabel = v.TrustpointLabel.ValueString()
	} else {
		data.Management.TrustpointLabel = ""
	}

	if !v.SkipCertificateVerification.IsNull() && !v.SkipCertificateVerification.IsUnknown() {
		data.Management.SkipCertificateVerification = new(bool)
		*data.Management.SkipCertificateVerification = v.SkipCertificateVerification.ValueBool()
	} else {
		data.Management.SkipCertificateVerification = nil
	}

	if !v.HostInterfaceAdminState.IsNull() && !v.HostInterfaceAdminState.IsUnknown() {
		data.Management.HostInterfaceAdminState = new(bool)
		*data.Management.HostInterfaceAdminState = v.HostInterfaceAdminState.ValueBool()
	} else {
		data.Management.HostInterfaceAdminState = nil
	}

	if !v.BrownfieldNetworkNameFormat.IsNull() && !v.BrownfieldNetworkNameFormat.IsUnknown() {
		data.Management.BrownfieldNetworkNameFormat = v.BrownfieldNetworkNameFormat.ValueString()
	} else {
		data.Management.BrownfieldNetworkNameFormat = ""
	}

	if !v.BrownfieldSkipOverlayNetworkAttachments.IsNull() && !v.BrownfieldSkipOverlayNetworkAttachments.IsUnknown() {
		data.Management.BrownfieldSkipOverlayNetworkAttachments = new(bool)
		*data.Management.BrownfieldSkipOverlayNetworkAttachments = v.BrownfieldSkipOverlayNetworkAttachments.ValueBool()
	} else {
		data.Management.BrownfieldSkipOverlayNetworkAttachments = nil
	}

	if !v.PolicyBasedRouting.IsNull() && !v.PolicyBasedRouting.IsUnknown() {
		data.Management.PolicyBasedRouting = new(bool)
		*data.Management.PolicyBasedRouting = v.PolicyBasedRouting.ValueBool()
	} else {
		data.Management.PolicyBasedRouting = nil
	}

	if !v.PtpVlanId.IsNull() && !v.PtpVlanId.IsUnknown() {
		data.Management.PtpVlanId = new(int64)
		*data.Management.PtpVlanId = v.PtpVlanId.ValueInt64()

	} else {
		data.Management.PtpVlanId = nil
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

	if !v.MplsIsisAreaNumber.IsNull() && !v.MplsIsisAreaNumber.IsUnknown() {
		data.Management.MplsIsisAreaNumber = v.MplsIsisAreaNumber.ValueString()
	} else {
		data.Management.MplsIsisAreaNumber = ""
	}

	if !v.StpRootOption.IsNull() && !v.StpRootOption.IsUnknown() {
		data.Management.StpRootOption = v.StpRootOption.ValueString()
	} else {
		data.Management.StpRootOption = ""
	}

	if !v.StpVlanRange.IsNull() && !v.StpVlanRange.IsUnknown() {
		data.Management.StpVlanRange = v.StpVlanRange.ValueString()
	} else {
		data.Management.StpVlanRange = ""
	}

	if !v.MstInstanceRange.IsNull() && !v.MstInstanceRange.IsUnknown() {
		data.Management.MstInstanceRange = v.MstInstanceRange.ValueString()
	} else {
		data.Management.MstInstanceRange = ""
	}

	if !v.StpBridgePriority.IsNull() && !v.StpBridgePriority.IsUnknown() {
		data.Management.StpBridgePriority = new(int64)
		*data.Management.StpBridgePriority = v.StpBridgePriority.ValueInt64()

	} else {
		data.Management.StpBridgePriority = nil
	}

	if !v.AllowVlanOnLeafTorPairing.IsNull() && !v.AllowVlanOnLeafTorPairing.IsUnknown() {
		data.Management.AllowVlanOnLeafTorPairing = v.AllowVlanOnLeafTorPairing.ValueString()
	} else {
		data.Management.AllowVlanOnLeafTorPairing = ""
	}

	if !v.PreInterfaceConfigLeaf.IsNull() && !v.PreInterfaceConfigLeaf.IsUnknown() {
		data.Management.PreInterfaceConfigLeaf = v.PreInterfaceConfigLeaf.ValueString()
	} else {
		data.Management.PreInterfaceConfigLeaf = ""
	}

	if !v.PreInterfaceConfigSpine.IsNull() && !v.PreInterfaceConfigSpine.IsUnknown() {
		data.Management.PreInterfaceConfigSpine = v.PreInterfaceConfigSpine.ValueString()
	} else {
		data.Management.PreInterfaceConfigSpine = ""
	}

	if !v.PreInterfaceConfigTor.IsNull() && !v.PreInterfaceConfigTor.IsUnknown() {
		data.Management.PreInterfaceConfigTor = v.PreInterfaceConfigTor.ValueString()
	} else {
		data.Management.PreInterfaceConfigTor = ""
	}

	if !v.ExtraConfigLeaf.IsNull() && !v.ExtraConfigLeaf.IsUnknown() {
		data.Management.ExtraConfigLeaf = v.ExtraConfigLeaf.ValueString()
	} else {
		data.Management.ExtraConfigLeaf = ""
	}

	if !v.ExtraConfigSpine.IsNull() && !v.ExtraConfigSpine.IsUnknown() {
		data.Management.ExtraConfigSpine = v.ExtraConfigSpine.ValueString()
	} else {
		data.Management.ExtraConfigSpine = ""
	}

	if !v.ExtraConfigTor.IsNull() && !v.ExtraConfigTor.IsUnknown() {
		data.Management.ExtraConfigTor = v.ExtraConfigTor.ValueString()
	} else {
		data.Management.ExtraConfigTor = ""
	}

	if !v.ExtraConfigIntraFabricLinks.IsNull() && !v.ExtraConfigIntraFabricLinks.IsUnknown() {
		data.Management.ExtraConfigIntraFabricLinks = v.ExtraConfigIntraFabricLinks.ValueString()
	} else {
		data.Management.ExtraConfigIntraFabricLinks = ""
	}

	if !v.MplsLoopbackIpRange.IsNull() && !v.MplsLoopbackIpRange.IsUnknown() {
		data.Management.MplsLoopbackIpRange = v.MplsLoopbackIpRange.ValueString()
	} else {
		data.Management.MplsLoopbackIpRange = ""
	}

	if !v.Ipv6SubnetRange.IsNull() && !v.Ipv6SubnetRange.IsUnknown() {
		data.Management.Ipv6SubnetRange = v.Ipv6SubnetRange.ValueString()
	} else {
		data.Management.Ipv6SubnetRange = ""
	}

	if !v.RouterIdRange.IsNull() && !v.RouterIdRange.IsUnknown() {
		data.Management.RouterIdRange = v.RouterIdRange.ValueString()
	} else {
		data.Management.RouterIdRange = ""
	}

	if !v.AutoSymmetricVrfLite.IsNull() && !v.AutoSymmetricVrfLite.IsUnknown() {
		data.Management.AutoSymmetricVrfLite = new(bool)
		*data.Management.AutoSymmetricVrfLite = v.AutoSymmetricVrfLite.ValueBool()
	} else {
		data.Management.AutoSymmetricVrfLite = nil
	}

	if !v.AutoVrfLiteDefaultVrf.IsNull() && !v.AutoVrfLiteDefaultVrf.IsUnknown() {
		data.Management.AutoVrfLiteDefaultVrf = new(bool)
		*data.Management.AutoVrfLiteDefaultVrf = v.AutoVrfLiteDefaultVrf.ValueBool()
	} else {
		data.Management.AutoVrfLiteDefaultVrf = nil
	}

	if !v.AutoSymmetricDefaultVrf.IsNull() && !v.AutoSymmetricDefaultVrf.IsUnknown() {
		data.Management.AutoSymmetricDefaultVrf = new(bool)
		*data.Management.AutoSymmetricDefaultVrf = v.AutoSymmetricDefaultVrf.ValueBool()
	} else {
		data.Management.AutoSymmetricDefaultVrf = nil
	}

	if !v.DefaultVrfRedistributionBgpRouteMap.IsNull() && !v.DefaultVrfRedistributionBgpRouteMap.IsUnknown() {
		data.Management.DefaultVrfRedistributionBgpRouteMap = v.DefaultVrfRedistributionBgpRouteMap.ValueString()
	} else {
		data.Management.DefaultVrfRedistributionBgpRouteMap = ""
	}

	if !v.IpServiceLevelAgreementIdRange.IsNull() && !v.IpServiceLevelAgreementIdRange.IsUnknown() {
		data.Management.IpServiceLevelAgreementIdRange = v.IpServiceLevelAgreementIdRange.ValueString()
	} else {
		data.Management.IpServiceLevelAgreementIdRange = ""
	}

	if !v.ObjectTrackingNumberRange.IsNull() && !v.ObjectTrackingNumberRange.IsUnknown() {
		data.Management.ObjectTrackingNumberRange = v.ObjectTrackingNumberRange.ValueString()
	} else {
		data.Management.ObjectTrackingNumberRange = ""
	}

	if !v.ServiceNetworkVlanRange.IsNull() && !v.ServiceNetworkVlanRange.IsUnknown() {
		data.Management.ServiceNetworkVlanRange = v.ServiceNetworkVlanRange.ValueString()
	} else {
		data.Management.ServiceNetworkVlanRange = ""
	}

	if !v.RouteMapSequenceNumberRange.IsNull() && !v.RouteMapSequenceNumberRange.IsUnknown() {
		data.Management.RouteMapSequenceNumberRange = v.RouteMapSequenceNumberRange.ValueString()
	} else {
		data.Management.RouteMapSequenceNumberRange = ""
	}

	if !v.InbandManagement.IsNull() && !v.InbandManagement.IsUnknown() {
		data.Management.InbandManagement = new(bool)
		*data.Management.InbandManagement = v.InbandManagement.ValueBool()
	} else {
		data.Management.InbandManagement = nil
	}

	if !v.SeedSwitchCoreInterfaces.IsNull() && !v.SeedSwitchCoreInterfaces.IsUnknown() {
		data.Management.SeedSwitchCoreInterfaces = v.SeedSwitchCoreInterfaces.ValueString()
	} else {
		data.Management.SeedSwitchCoreInterfaces = ""
	}

	if !v.SpineSwitchCoreInterfaces.IsNull() && !v.SpineSwitchCoreInterfaces.IsUnknown() {
		data.Management.SpineSwitchCoreInterfaces = v.SpineSwitchCoreInterfaces.ValueString()
	} else {
		data.Management.SpineSwitchCoreInterfaces = ""
	}

	if !v.InbandDhcpServers.IsNull() && !v.InbandDhcpServers.IsUnknown() {
		data.Management.InbandDhcpServers = v.InbandDhcpServers.ValueString()
	} else {
		data.Management.InbandDhcpServers = ""
	}

	if !v.UnNumberedBootstrapLbId.IsNull() && !v.UnNumberedBootstrapLbId.IsUnknown() {
		data.Management.UnNumberedBootstrapLbId = new(int64)
		*data.Management.UnNumberedBootstrapLbId = v.UnNumberedBootstrapLbId.ValueInt64()

	} else {
		data.Management.UnNumberedBootstrapLbId = nil
	}

	if !v.UnNumberedDhcpStartAddress.IsNull() && !v.UnNumberedDhcpStartAddress.IsUnknown() {
		data.Management.UnNumberedDhcpStartAddress = v.UnNumberedDhcpStartAddress.ValueString()
	} else {
		data.Management.UnNumberedDhcpStartAddress = ""
	}

	if !v.UnNumberedDhcpEndAddress.IsNull() && !v.UnNumberedDhcpEndAddress.IsUnknown() {
		data.Management.UnNumberedDhcpEndAddress = v.UnNumberedDhcpEndAddress.ValueString()
	} else {
		data.Management.UnNumberedDhcpEndAddress = ""
	}

	if !v.HeartbeatInterval.IsNull() && !v.HeartbeatInterval.IsUnknown() {
		data.Management.HeartbeatInterval = new(int64)
		*data.Management.HeartbeatInterval = v.HeartbeatInterval.ValueInt64()

	} else {
		data.Management.HeartbeatInterval = nil
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

	if !v.NtpServerCollection.IsNull() && !v.NtpServerCollection.IsUnknown() {
		listStringData := make([]string, len(v.NtpServerCollection.Elements()))
		dg := v.NtpServerCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.NtpServerCollection = make([]string, len(listStringData))

		copy(data.Management.NtpServerCollection, listStringData)
	}

	if !v.NtpServerVrfCollection.IsNull() && !v.NtpServerVrfCollection.IsUnknown() {
		listStringData := make([]string, len(v.NtpServerVrfCollection.Elements()))
		dg := v.NtpServerVrfCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.NtpServerVrfCollection = make([]string, len(listStringData))

		copy(data.Management.NtpServerVrfCollection, listStringData)
	}

	if !v.SyslogServerCollection.IsNull() && !v.SyslogServerCollection.IsUnknown() {
		listStringData := make([]string, len(v.SyslogServerCollection.Elements()))
		dg := v.SyslogServerCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.SyslogServerCollection = make([]string, len(listStringData))

		copy(data.Management.SyslogServerCollection, listStringData)
	}

	if !v.SyslogSeverityCollection.IsNull() && !v.SyslogSeverityCollection.IsUnknown() {
		listInt64Data := make([]int64, len(v.SyslogSeverityCollection.Elements()))
		dg := v.SyslogSeverityCollection.ElementsAs(context.Background(), &listInt64Data, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.SyslogSeverityCollection = make([]int64, len(listInt64Data))
		copy(data.Management.SyslogSeverityCollection, listInt64Data)
	}

	if !v.SyslogServerVrfCollection.IsNull() && !v.SyslogServerVrfCollection.IsUnknown() {
		listStringData := make([]string, len(v.SyslogServerVrfCollection.Elements()))
		dg := v.SyslogServerVrfCollection.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.Management.SyslogServerVrfCollection = make([]string, len(listStringData))

		copy(data.Management.SyslogServerVrfCollection, listStringData)
	}

	if !v.NetflowEnable.IsNull() && !v.NetflowEnable.IsUnknown() {
		data.Management.NetflowSettings.NetflowEnable = new(bool)
		*data.Management.NetflowSettings.NetflowEnable = v.NetflowEnable.ValueBool()
	} else {
		data.Management.NetflowSettings.NetflowEnable = nil
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

	//MARSHALL_LIST  BEGIN NetflowExporterCollection[i1]

	if !v.NetflowExporterCollection.IsNull() && !v.NetflowExporterCollection.IsUnknown() {
		elements := make([]NetflowExporterCollectionValue, len(v.NetflowExporterCollection.Elements()))
		// Not augmenting

		data.Management.NetflowSettings.NetflowExporterCollection = make([]NDFCNetflowExporterCollectionValue, len(v.NetflowExporterCollection.Elements()))

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

		data.Management.NetflowSettings.NetflowRecordCollection = make([]NDFCNetflowRecordCollectionValue, len(v.NetflowRecordCollection.Elements()))

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

		data.Management.NetflowSettings.NetflowMonitorCollection = make([]NDFCNetflowMonitorCollectionValue, len(v.NetflowMonitorCollection.Elements()))

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

	//MARSHALL_LIST  BEGIN NetflowSamplerCollection[i1]

	if !v.NetflowSamplerCollection.IsNull() && !v.NetflowSamplerCollection.IsUnknown() {
		elements := make([]NetflowSamplerCollectionValue, len(v.NetflowSamplerCollection.Elements()))
		// Not augmenting

		data.Management.NetflowSettings.NetflowSamplerCollection = make([]NDFCNetflowSamplerCollectionValue, len(v.NetflowSamplerCollection.Elements()))

		// -- Set here 1 |.Management.NetflowSettings.NetflowSamplerCollection[i1]|NetflowSamplerCollection[i1]| --

		diag := v.NetflowSamplerCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.NetflowSettings.NetflowSamplerCollection[i1] NetflowSamplerCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.SamplerName.IsNull() && !ele1.SamplerName.IsUnknown() {

				data.Management.NetflowSettings.NetflowSamplerCollection[i1].SamplerName = ele1.SamplerName.ValueString()
			} else {
				data.Management.NetflowSettings.NetflowSamplerCollection[i1].SamplerName = ""
			}

			if !ele1.NumSamples.IsNull() && !ele1.NumSamples.IsUnknown() {

				data.Management.NetflowSettings.NetflowSamplerCollection[i1].NumSamples = new(int64)
				*data.Management.NetflowSettings.NetflowSamplerCollection[i1].NumSamples = ele1.NumSamples.ValueInt64()

			} else {
				data.Management.NetflowSettings.NetflowSamplerCollection[i1].NumSamples = nil
			}

			if !ele1.SamplingRate.IsNull() && !ele1.SamplingRate.IsUnknown() {

				data.Management.NetflowSettings.NetflowSamplerCollection[i1].SamplingRate = new(int64)
				*data.Management.NetflowSettings.NetflowSamplerCollection[i1].SamplingRate = ele1.SamplingRate.ValueInt64()

			} else {
				data.Management.NetflowSettings.NetflowSamplerCollection[i1].SamplingRate = nil
			}

		} /* for loop */
	} /* isNull if check */

	//MARSHALL_LIST  BEGIN VrfFlowRules[i1]

	if !v.VrfFlowRules.IsNull() && !v.VrfFlowRules.IsUnknown() {
		elements := make([]VrfFlowRulesValue, len(v.VrfFlowRules.Elements()))
		// Not augmenting

		data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules = make([]NDFCVrfFlowRulesValue, len(v.VrfFlowRules.Elements()))

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

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2]|TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.VrfFlowRules[i1].VrfFlowRuleAttributes = make([]NDFCVrfFlowRuleAttributesValue, len(ele1.VrfFlowRuleAttributes.Elements()))

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

		data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules = make([]NDFCInterfaceFlowRulesValue, len(v.InterfaceFlowRules.Elements()))

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

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleInterfaceCollection = make([]NDFCInterfaceFlowRuleInterfaceCollectionValue, len(ele1.InterfaceFlowRuleInterfaceCollection.Elements()))

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

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2]|TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.InterfaceFlowRules[i1].InterfaceFlowRuleAttributes = make([]NDFCInterfaceFlowRuleAttributesValue, len(ele1.InterfaceFlowRuleAttributes.Elements()))

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

		data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules = make([]NDFCL3OutFlowRulesValue, len(v.L3OutFlowRules.Elements()))

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

				// -- Set here 2 |.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]|TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.FlowRules.L3OutFlowRules[i1].L3OutFlowRuleInterfaceCollection = make([]NDFCL3OutFlowRuleInterfaceCollectionValue, len(ele1.L3OutFlowRuleInterfaceCollection.Elements()))

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

		data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules = make([]NDFCInterfaceRulesValue, len(v.InterfaceRules.Elements()))

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

				// -- Set here 2 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2]| --
				data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection = make([]NDFCInterfaceRuleInterfaceCollectionValue, len(ele1.InterfaceRuleInterfaceCollection.Elements()))

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

						// -- Set here 2 |.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3]|TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces[i3]| --
						data.TelemetrySettings.FlowCollection.TrafficAnalyticsRules.InterfaceRules[i1].InterfaceRuleInterfaceCollection[i2].InterfaceRuleInterfaces = make([]NDFCInterfaceRuleInterfacesValue, len(ele2.InterfaceRuleInterfaces.Elements()))

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

		data.ExternalStreamingSettings.Email = make([]NDFCEmailValue, len(v.Email.Elements()))

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

		data.ExternalStreamingSettings.MessageBus = make([]NDFCMessageBusValue, len(v.MessageBus.Elements()))

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
