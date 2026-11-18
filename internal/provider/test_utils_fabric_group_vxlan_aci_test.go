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
	"fmt"
	"strconv"
	. "terraform-provider-nd/internal/common/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

func FabricGroupVxlanAciModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}

	if c.Management.L2VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_vni_range").String(), c.Management.L2VniRange))
	}
	if c.Management.L3VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_range").String(), c.Management.L3VniRange))
	}
	if c.Management.AnycastGatewayMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gateway_mac").String(), c.Management.AnycastGatewayMac))
	}
	if c.Management.MultisiteOverlayInterConnectType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_overlay_inter_connect_type").String(), c.Management.MultisiteOverlayInterConnectType))
	}
	if c.Management.RouteServerRedistributeDirectRouteMap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_redistribute_direct_route_map").String(), strconv.FormatBool(*c.Management.RouteServerRedistributeDirectRouteMap)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_redistribute_direct_route_map").String(), "false"))
	}
	if c.Management.RouteServerRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_routing_tag").String(), strconv.Itoa(int(*c.Management.RouteServerRoutingTag))))
	}
	if c.Management.EnableMsOverlayIfcBgpDesc != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_ms_overlay_ifc_bgp_desc").String(), strconv.FormatBool(*c.Management.EnableMsOverlayIfcBgpDesc)))
	}
	if c.Management.AutoMultisiteUnderlayInterConnect != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_multisite_underlay_inter_connect").String(), strconv.FormatBool(*c.Management.AutoMultisiteUnderlayInterConnect)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_multisite_underlay_inter_connect").String(), "false"))
	}
	if c.Management.MultisiteDelayRestore != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_delay_restore").String(), strconv.Itoa(int(*c.Management.MultisiteDelayRestore))))
	}
	if c.Management.MultisiteInterConnectBgpAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_authentication").String(), strconv.FormatBool(*c.Management.MultisiteInterConnectBgpAuthentication)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_authentication").String(), "false"))
	}
	if c.Management.MultisiteInterConnectBgpAuthKeyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_auth_key_type").String(), c.Management.MultisiteInterConnectBgpAuthKeyType))
	}
	if c.Management.MultisiteInterConnectBgpKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_key").String(), c.Management.MultisiteInterConnectBgpKey))
	}
	if c.Management.MultisiteLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_loopback_id").String(), strconv.Itoa(int(*c.Management.MultisiteLoopbackId))))
	}
	if c.Management.BorderGatewayRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_gateway_routing_tag").String(), strconv.Itoa(int(*c.Management.BorderGatewayRoutingTag))))
	}
	if c.Management.SecurityGroupTag != nil {
		value, conversionErr := AnyToString(c.Management.SecurityGroupTag)
		if conversionErr != nil {
			ret = append(ret, resource.TestCheckResourceAttrWith(RscName, attrPath.AtName("security_group_tag").String(), func(string) error { return conversionErr }))
		} else {
			ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag").String(), fmt.Sprint(value)))
		}

	}
	if c.Management.SecurityGroupTagPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_prefix").String(), c.Management.SecurityGroupTagPrefix))
	}
	if c.Management.SecurityGroupTagMacSegmentation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_mac_segmentation").String(), strconv.FormatBool(*c.Management.SecurityGroupTagMacSegmentation)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_mac_segmentation").String(), "false"))
	}
	if c.Management.SecurityGroupTagIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_id_range").String(), c.Management.SecurityGroupTagIdRange))
	}
	if c.Management.SecurityGroupTagPreprovision != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_preprovision").String(), strconv.FormatBool(*c.Management.SecurityGroupTagPreprovision)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_preprovision").String(), "false"))
	}
	if c.Management.SecurityGroupStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_status").String(), c.Management.SecurityGroupStatus))
	}
	if c.Management.MultisiteLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_loopback_ip_range").String(), c.Management.MultisiteLoopbackIpRange))
	}
	if c.Management.MultisiteUnderlaySubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_subnet_range").String(), c.Management.MultisiteUnderlaySubnetRange))
	}
	if c.Management.MultisiteUnderlaySubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.MultisiteUnderlaySubnetTargetMask))))
	}
	if c.Management.ScheduledBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup").String(), strconv.FormatBool(*c.Management.ScheduledBackup)))
	}
	if c.Management.ScheduledBackupTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup_time").String(), c.Management.ScheduledBackupTime))
	}
	return ret
}

func FabricGroupVxlanAciRouteServerCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCRouteServerCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.RouteServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_ip").String(), c.RouteServerIp))
	}
	if c.RouteServerAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_asn").String(), c.RouteServerAsn))
	}
	return ret
}
