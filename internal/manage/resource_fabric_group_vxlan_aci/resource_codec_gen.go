// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_group_vxlan_aci

import (
	"context"
	"log"
	. "terraform-provider-nd/internal/common/types"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

func (v *FabricGroupVxlanAciModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
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

	if jsonData.Management.AnycastGatewayMac != "" {
		v.AnycastGatewayMac = types.StringValue(jsonData.Management.AnycastGatewayMac)
	} else {
		v.AnycastGatewayMac = types.StringNull()
	}

	if jsonData.Management.MultisiteOverlayInterConnectType != "" {
		v.MultisiteOverlayInterConnectType = types.StringValue(jsonData.Management.MultisiteOverlayInterConnectType)
	} else {
		v.MultisiteOverlayInterConnectType = types.StringNull()
	}

	if len(jsonData.Management.RouteServerCollection) == 0 {
		log.Printf("v.RouteServerCollection is empty")
		v.RouteServerCollection = types.ListNull(RouteServerCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]RouteServerCollectionValue, len(jsonData.Management.RouteServerCollection))
		for i, item := range jsonData.Management.RouteServerCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.RouteServerCollection, err = types.ListValueFrom(context.Background(), RouteServerCollectionValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if jsonData.Management.RouteServerRedistributeDirectRouteMap != nil {
		v.RouteServerRedistributeDirectRouteMap = types.BoolValue(*jsonData.Management.RouteServerRedistributeDirectRouteMap)
	} else {
		v.RouteServerRedistributeDirectRouteMap = types.BoolNull()
	}

	if jsonData.Management.RouteServerRoutingTag != nil {
		v.RouteServerRoutingTag = types.Int64Value(*jsonData.Management.RouteServerRoutingTag)
	} else {
		v.RouteServerRoutingTag = types.Int64Null()
	}

	if jsonData.Management.EnableMsOverlayIfcBgpDesc != nil {
		v.EnableMsOverlayIfcBgpDesc = types.BoolValue(*jsonData.Management.EnableMsOverlayIfcBgpDesc)
	} else {
		v.EnableMsOverlayIfcBgpDesc = types.BoolNull()
	}

	if jsonData.Management.AutoMultisiteUnderlayInterConnect != nil {
		v.AutoMultisiteUnderlayInterConnect = types.BoolValue(*jsonData.Management.AutoMultisiteUnderlayInterConnect)
	} else {
		v.AutoMultisiteUnderlayInterConnect = types.BoolNull()
	}

	if jsonData.Management.MultisiteDelayRestore != nil {
		v.MultisiteDelayRestore = types.Int64Value(*jsonData.Management.MultisiteDelayRestore)
	} else {
		v.MultisiteDelayRestore = types.Int64Null()
	}

	if jsonData.Management.MultisiteInterConnectBgpAuthentication != nil {
		v.MultisiteInterConnectBgpAuthentication = types.BoolValue(*jsonData.Management.MultisiteInterConnectBgpAuthentication)
	} else {
		v.MultisiteInterConnectBgpAuthentication = types.BoolNull()
	}

	if jsonData.Management.MultisiteInterConnectBgpAuthKeyType != "" {
		v.MultisiteInterConnectBgpAuthKeyType = types.StringValue(jsonData.Management.MultisiteInterConnectBgpAuthKeyType)
	} else {
		v.MultisiteInterConnectBgpAuthKeyType = types.StringNull()
	}

	if jsonData.Management.MultisiteInterConnectBgpKey != "" {
		v.MultisiteInterConnectBgpKey = types.StringValue(jsonData.Management.MultisiteInterConnectBgpKey)
	} else {
		v.MultisiteInterConnectBgpKey = types.StringNull()
	}

	if jsonData.Management.MultisiteLoopbackId != nil {
		v.MultisiteLoopbackId = types.Int64Value(*jsonData.Management.MultisiteLoopbackId)
	} else {
		v.MultisiteLoopbackId = types.Int64Null()
	}

	if jsonData.Management.BorderGatewayRoutingTag != nil {
		v.BorderGatewayRoutingTag = types.Int64Value(*jsonData.Management.BorderGatewayRoutingTag)
	} else {
		v.BorderGatewayRoutingTag = types.Int64Null()
	}

	{
		var input any = jsonData.Management.SecurityGroupTag

		if input == nil {
			v.SecurityGroupTag = types.StringNull()
		} else {
			converted, conversionErr := AnyToString(input)
			if conversionErr != nil {
				err.AddError("Invalid API response", "Unable to convert API value for Terraform attribute security_group_tag: "+conversionErr.Error())
				return err
			}
			v.SecurityGroupTag = types.StringValue(converted)

		}
	}

	if jsonData.Management.SecurityGroupTagPrefix != "" {
		v.SecurityGroupTagPrefix = types.StringValue(jsonData.Management.SecurityGroupTagPrefix)
	} else {
		v.SecurityGroupTagPrefix = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagMacSegmentation != nil {
		v.SecurityGroupTagMacSegmentation = types.BoolValue(*jsonData.Management.SecurityGroupTagMacSegmentation)
	} else {
		v.SecurityGroupTagMacSegmentation = types.BoolNull()
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

	if jsonData.Management.MultisiteLoopbackIpRange != "" {
		v.MultisiteLoopbackIpRange = types.StringValue(jsonData.Management.MultisiteLoopbackIpRange)
	} else {
		v.MultisiteLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlaySubnetRange != "" {
		v.MultisiteUnderlaySubnetRange = types.StringValue(jsonData.Management.MultisiteUnderlaySubnetRange)
	} else {
		v.MultisiteUnderlaySubnetRange = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlaySubnetTargetMask != nil {
		v.MultisiteUnderlaySubnetTargetMask = types.Int64Value(*jsonData.Management.MultisiteUnderlaySubnetTargetMask)
	} else {
		v.MultisiteUnderlaySubnetTargetMask = types.Int64Null()
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

	return err
}

func (v *RouteServerCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCRouteServerCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.RouteServerIp != "" {
		v.RouteServerIp = types.StringValue(jsonData.RouteServerIp)
	} else {
		v.RouteServerIp = types.StringNull()
	}

	if jsonData.RouteServerAsn != "" {
		v.RouteServerAsn = types.StringValue(jsonData.RouteServerAsn)
	} else {
		v.RouteServerAsn = types.StringNull()
	}

	return err
}

func (v FabricGroupVxlanAciModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
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

	if !v.AnycastGatewayMac.IsNull() && !v.AnycastGatewayMac.IsUnknown() {
		data.Management.AnycastGatewayMac = v.AnycastGatewayMac.ValueString()
	} else {
		data.Management.AnycastGatewayMac = ""
	}

	if !v.MultisiteOverlayInterConnectType.IsNull() && !v.MultisiteOverlayInterConnectType.IsUnknown() {
		data.Management.MultisiteOverlayInterConnectType = v.MultisiteOverlayInterConnectType.ValueString()
	} else {
		data.Management.MultisiteOverlayInterConnectType = ""
	}

	if !v.RouteServerRedistributeDirectRouteMap.IsNull() && !v.RouteServerRedistributeDirectRouteMap.IsUnknown() {
		data.Management.RouteServerRedistributeDirectRouteMap = new(bool)
		*data.Management.RouteServerRedistributeDirectRouteMap = v.RouteServerRedistributeDirectRouteMap.ValueBool()
	} else {
		data.Management.RouteServerRedistributeDirectRouteMap = nil
	}

	if !v.RouteServerRoutingTag.IsNull() && !v.RouteServerRoutingTag.IsUnknown() {
		data.Management.RouteServerRoutingTag = new(int64)
		*data.Management.RouteServerRoutingTag = v.RouteServerRoutingTag.ValueInt64()

	} else {
		data.Management.RouteServerRoutingTag = nil
	}

	if !v.EnableMsOverlayIfcBgpDesc.IsNull() && !v.EnableMsOverlayIfcBgpDesc.IsUnknown() {
		data.Management.EnableMsOverlayIfcBgpDesc = new(bool)
		*data.Management.EnableMsOverlayIfcBgpDesc = v.EnableMsOverlayIfcBgpDesc.ValueBool()
	} else {
		data.Management.EnableMsOverlayIfcBgpDesc = nil
	}

	if !v.AutoMultisiteUnderlayInterConnect.IsNull() && !v.AutoMultisiteUnderlayInterConnect.IsUnknown() {
		data.Management.AutoMultisiteUnderlayInterConnect = new(bool)
		*data.Management.AutoMultisiteUnderlayInterConnect = v.AutoMultisiteUnderlayInterConnect.ValueBool()
	} else {
		data.Management.AutoMultisiteUnderlayInterConnect = nil
	}

	if !v.MultisiteDelayRestore.IsNull() && !v.MultisiteDelayRestore.IsUnknown() {
		data.Management.MultisiteDelayRestore = new(int64)
		*data.Management.MultisiteDelayRestore = v.MultisiteDelayRestore.ValueInt64()

	} else {
		data.Management.MultisiteDelayRestore = nil
	}

	if !v.MultisiteInterConnectBgpAuthentication.IsNull() && !v.MultisiteInterConnectBgpAuthentication.IsUnknown() {
		data.Management.MultisiteInterConnectBgpAuthentication = new(bool)
		*data.Management.MultisiteInterConnectBgpAuthentication = v.MultisiteInterConnectBgpAuthentication.ValueBool()
	} else {
		data.Management.MultisiteInterConnectBgpAuthentication = nil
	}

	if !v.MultisiteInterConnectBgpAuthKeyType.IsNull() && !v.MultisiteInterConnectBgpAuthKeyType.IsUnknown() {
		data.Management.MultisiteInterConnectBgpAuthKeyType = v.MultisiteInterConnectBgpAuthKeyType.ValueString()
	} else {
		data.Management.MultisiteInterConnectBgpAuthKeyType = ""
	}

	if !v.MultisiteInterConnectBgpKey.IsNull() && !v.MultisiteInterConnectBgpKey.IsUnknown() {
		data.Management.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey.ValueString()
	} else {
		data.Management.MultisiteInterConnectBgpKey = ""
	}

	if !v.MultisiteLoopbackId.IsNull() && !v.MultisiteLoopbackId.IsUnknown() {
		data.Management.MultisiteLoopbackId = new(int64)
		*data.Management.MultisiteLoopbackId = v.MultisiteLoopbackId.ValueInt64()

	} else {
		data.Management.MultisiteLoopbackId = nil
	}

	if !v.BorderGatewayRoutingTag.IsNull() && !v.BorderGatewayRoutingTag.IsUnknown() {
		data.Management.BorderGatewayRoutingTag = new(int64)
		*data.Management.BorderGatewayRoutingTag = v.BorderGatewayRoutingTag.ValueInt64()

	} else {
		data.Management.BorderGatewayRoutingTag = nil
	}

	if !v.SecurityGroupTag.IsNull() && !v.SecurityGroupTag.IsUnknown() {
		data.Management.SecurityGroupTag = v.SecurityGroupTag.ValueString()
	} else {
		data.Management.SecurityGroupTag = nil
	}

	if !v.SecurityGroupTagPrefix.IsNull() && !v.SecurityGroupTagPrefix.IsUnknown() {
		data.Management.SecurityGroupTagPrefix = v.SecurityGroupTagPrefix.ValueString()
	} else {
		data.Management.SecurityGroupTagPrefix = ""
	}

	if !v.SecurityGroupTagMacSegmentation.IsNull() && !v.SecurityGroupTagMacSegmentation.IsUnknown() {
		data.Management.SecurityGroupTagMacSegmentation = new(bool)
		*data.Management.SecurityGroupTagMacSegmentation = v.SecurityGroupTagMacSegmentation.ValueBool()
	} else {
		data.Management.SecurityGroupTagMacSegmentation = nil
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

	if !v.MultisiteLoopbackIpRange.IsNull() && !v.MultisiteLoopbackIpRange.IsUnknown() {
		data.Management.MultisiteLoopbackIpRange = v.MultisiteLoopbackIpRange.ValueString()
	} else {
		data.Management.MultisiteLoopbackIpRange = ""
	}

	if !v.MultisiteUnderlaySubnetRange.IsNull() && !v.MultisiteUnderlaySubnetRange.IsUnknown() {
		data.Management.MultisiteUnderlaySubnetRange = v.MultisiteUnderlaySubnetRange.ValueString()
	} else {
		data.Management.MultisiteUnderlaySubnetRange = ""
	}

	if !v.MultisiteUnderlaySubnetTargetMask.IsNull() && !v.MultisiteUnderlaySubnetTargetMask.IsUnknown() {
		data.Management.MultisiteUnderlaySubnetTargetMask = new(int64)
		*data.Management.MultisiteUnderlaySubnetTargetMask = v.MultisiteUnderlaySubnetTargetMask.ValueInt64()

	} else {
		data.Management.MultisiteUnderlaySubnetTargetMask = nil
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

	//MARSHALL_LIST  BEGIN RouteServerCollection[i1]

	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		elements := make([]RouteServerCollectionValue, len(v.RouteServerCollection.Elements()))
		// Not augmenting

		data.Management.RouteServerCollection = make([]resource_fabric_common.NDFCRouteServerCollectionValue, len(v.RouteServerCollection.Elements()))

		// -- Set here 1 |.Management.RouteServerCollection[i1]|RouteServerCollection[i1]| --

		diag := v.RouteServerCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.RouteServerCollection[i1] RouteServerCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.RouteServerIp.IsNull() && !ele1.RouteServerIp.IsUnknown() {

				data.Management.RouteServerCollection[i1].RouteServerIp = ele1.RouteServerIp.ValueString()
			} else {
				data.Management.RouteServerCollection[i1].RouteServerIp = ""
			}

			if !ele1.RouteServerAsn.IsNull() && !ele1.RouteServerAsn.IsUnknown() {

				data.Management.RouteServerCollection[i1].RouteServerAsn = ele1.RouteServerAsn.ValueString()
			} else {
				data.Management.RouteServerCollection[i1].RouteServerAsn = ""
			}

		} /* for loop */
	} /* isNull if check */

	return data
}
