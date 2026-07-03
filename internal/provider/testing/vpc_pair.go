// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_vpc_pair"
)

func GenerateVpcPairObject(
	obj **resource_vpc_pair.NDFCVpcPairModel,
	fabricName string,
	switchID1 string,
	switchID2 string,
	useVirtualPeerLink bool,
	deploy bool,
) {
	vpcPair := new(resource_vpc_pair.NDFCVpcPairModel)

	vpcPair.FabricName = fabricName
	vpcPair.Switch1SerialNumber = switchID1
	vpcPair.Switch2SerialNumber = switchID2
	vpcPair.UseVirtualPeerlink = &useVirtualPeerLink
	vpcPair.Deploy = deploy

	*obj = vpcPair
}

func ModifyVpcPairObject(
	obj **resource_vpc_pair.NDFCVpcPairModel,
	values map[string]interface{},
) {
	vpcPair := *obj
	if vpcPair == nil {
		return
	}

	for key, val := range values {
		switch key {
		case "switch_1_serial_number":
			vpcPair.Switch1SerialNumber = val.(string)
		case "switch_2_serial_number":
			vpcPair.Switch2SerialNumber = val.(string)
		case "use_virtual_peerlink":
			v := val.(bool)
			vpcPair.UseVirtualPeerlink = &v
		case "deploy":
			vpcPair.Deploy = val.(bool)
		}
	}

	*obj = vpcPair
}
