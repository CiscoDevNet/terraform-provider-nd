// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

// FabricGroupVxlanAciTestResource contains the model and rendering details needed by
// the VXLAN-ACI Fabric Group acceptance tests.
type FabricGroupVxlanAciTestResource struct {
	Model                            *resource_fabric_common.NDFCFabricCommonModel
	RenderEmptyRouteServerCollection bool
	AdditionalAttributes             string // Raw HCL for unsupported-attribute tests.
}

// GenerateFabricGroupVxlanAciObject builds a minimal fabric-group model for
// acceptance tests. Optional and computed values remain unset so tests can
// verify API default hydration and import convergence.
func GenerateFabricGroupVxlanAciObject(
	obj **resource_fabric_common.NDFCFabricCommonModel,
	fabricName string,
	anycastGatewayMAC string,
) {
	fabricGroup := new(resource_fabric_common.NDFCFabricCommonModel)
	fabricGroup.FabricName = fabricName
	fabricGroup.Category = "fabricGroup"
	fabricGroup.Management.FabricType = "vxlanAci"
	fabricGroup.Management.AnycastGatewayMac = anycastGatewayMAC

	*obj = fabricGroup
}

// FabricGroupVxlanAciResource wraps a fabric-group model for Terraform template
// dispatch through the shared fabric test renderer.
func FabricGroupVxlanAciResource(
	model *resource_fabric_common.NDFCFabricCommonModel,
) *FabricGroupVxlanAciTestResource {
	return &FabricGroupVxlanAciTestResource{Model: model}
}

// FabricGroupVxlanAciResourceWithEmptyRouteServers renders an explicit empty list so
// acceptance coverage can distinguish [] from an omitted collection.
func FabricGroupVxlanAciResourceWithEmptyRouteServers(
	model *resource_fabric_common.NDFCFabricCommonModel,
) *FabricGroupVxlanAciTestResource {
	return &FabricGroupVxlanAciTestResource{
		Model:                            model,
		RenderEmptyRouteServerCollection: true,
	}
}

// ModifyFabricGroupVxlanAciObject applies configurable values to an existing model.
func ModifyFabricGroupVxlanAciObject(
	obj **resource_fabric_common.NDFCFabricCommonModel,
	values map[string]interface{},
) {
	fabricGroup := *obj
	if fabricGroup == nil {
		return
	}

	for key, value := range values {
		switch key {
		case "fabric_name":
			fabricGroup.FabricName = value.(string)
		case "anycast_gateway_mac":
			fabricGroup.Management.AnycastGatewayMac = value.(string)
		case "l2_vni_range":
			fabricGroup.Management.L2VniRange = value.(string)
		case "l3_vni_range":
			fabricGroup.Management.L3VniRange = value.(string)
		case "scheduled_backup":
			v := value.(bool)
			fabricGroup.Management.ScheduledBackup = &v
		case "scheduled_backup_time":
			fabricGroup.Management.ScheduledBackupTime = value.(string)
		case "security_group_tag":
			fabricGroup.Management.SecurityGroupTag = value.(string)
		case "security_group_tag_prefix":
			fabricGroup.Management.SecurityGroupTagPrefix = value.(string)
		case "security_group_tag_id_range":
			fabricGroup.Management.SecurityGroupTagIdRange = value.(string)
		case "security_group_tag_preprovision":
			v := value.(bool)
			fabricGroup.Management.SecurityGroupTagPreprovision = &v
		case "security_group_tag_mac_segmentation":
			v := value.(bool)
			fabricGroup.Management.SecurityGroupTagMacSegmentation = &v
		case "multisite_overlay_inter_connect_type":
			fabricGroup.Management.MultisiteOverlayInterConnectType = value.(string)
		case "route_server_collection":
			items := value.([]resource_fabric_common.NDFCRouteServerCollectionValue)
			fabricGroup.Management.RouteServerCollection = append(
				resource_fabric_common.NDFCRouteServerCollectionValues(nil),
				items...,
			)
		case "route_server_redistribute_direct_route_map":
			v := value.(bool)
			fabricGroup.Management.RouteServerRedistributeDirectRouteMap = &v
		case "route_server_routing_tag":
			v := value.(int64)
			fabricGroup.Management.RouteServerRoutingTag = &v
		case "enable_ms_overlay_ifc_bgp_desc":
			v := value.(bool)
			fabricGroup.Management.EnableMsOverlayIfcBgpDesc = &v
		case "auto_multisite_underlay_inter_connect":
			v := value.(bool)
			fabricGroup.Management.AutoMultisiteUnderlayInterConnect = &v
		case "multisite_delay_restore":
			v := value.(int64)
			fabricGroup.Management.MultisiteDelayRestore = &v
		case "multisite_inter_connect_bgp_authentication":
			v := value.(bool)
			fabricGroup.Management.MultisiteInterConnectBgpAuthentication = &v
		case "multisite_inter_connect_bgp_auth_key_type":
			fabricGroup.Management.MultisiteInterConnectBgpAuthKeyType = value.(string)
		case "multisite_inter_connect_bgp_key":
			fabricGroup.Management.MultisiteInterConnectBgpKey = value.(string)
		case "multisite_loopback_id":
			v := value.(int64)
			fabricGroup.Management.MultisiteLoopbackId = &v
		case "border_gateway_routing_tag":
			v := value.(int64)
			fabricGroup.Management.BorderGatewayRoutingTag = &v
		case "multisite_loopback_ip_range":
			fabricGroup.Management.MultisiteLoopbackIpRange = value.(string)
		case "multisite_underlay_subnet_range":
			fabricGroup.Management.MultisiteUnderlaySubnetRange = value.(string)
		case "multisite_underlay_subnet_target_mask":
			v := value.(int64)
			fabricGroup.Management.MultisiteUnderlaySubnetTargetMask = &v
		}
	}

	*obj = fabricGroup
}
