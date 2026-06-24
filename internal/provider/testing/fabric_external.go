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

// defaultFabricExternalValues returns sensible defaults for an External Connectivity fabric.
func defaultFabricExternalValues() map[string]interface{} {
	return map[string]interface{}{
		"license_tier":              "premier",
		"category":                  "fabric",
		"security_domain":           "all",
		"sub_interface_dot1q_range": "2-511",
	}
}

// GenerateFabricExternalObject creates an External Connectivity fabric model object for testing.
// fabricName and bgpAsn are mandatory identifiers.
// overrides lets each test supply unique values for any field so that
// multiple fabrics can coexist without conflicting.
// Any key not present in overrides gets the value from defaultFabricExternalValues().
func GenerateFabricExternalObject(obj **resource_fabric_common.NDFCFabricCommonModel,
	fabricName string, bgpAsn string,
	overrides map[string]interface{}) {

	fabric := new(resource_fabric_common.NDFCFabricCommonModel)

	fabric.FabricName = fabricName
	fabric.Management.FabricType = "externalConnectivity"
	fabric.Management.BgpAsn = bgpAsn

	// Merge defaults with caller overrides (overrides win)
	merged := defaultFabricExternalValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyFabricCommonValues(fabric, merged)

	*obj = fabric
}

// ModifyFabricExternalObject modifies fields on an existing External Connectivity fabric model.
func ModifyFabricExternalObject(
	obj **resource_fabric_common.NDFCFabricCommonModel,
	values map[string]interface{},
) {
	fabric := *obj
	if fabric == nil {
		return
	}

	applyFabricCommonValues(fabric, values)

	*obj = fabric
}
