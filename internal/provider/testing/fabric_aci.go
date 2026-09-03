// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import "terraform-provider-nd/internal/manage/resource_fabric_aci"

// GenerateFabricAciObject creates a Fabric ACI model for testing.
func GenerateFabricAciObject(
	obj **resource_fabric_aci.NDFCFabricAciModel,
	fabricName string,
	hostname string,
	username string,
	password string,
	overrides map[string]interface{},
) {
	fabric := new(resource_fabric_aci.NDFCFabricAciModel)
	fabric.Spec.Aci.FabricName = fabricName
	fabric.Spec.Hostname = hostname
	fabric.Spec.Credentials.Username = username
	fabric.Spec.Credentials.Password = password

	applyFabricAciValues(fabric, overrides)
	*obj = fabric
}

// ModifyFabricAciObject modifies fields on an existing Fabric ACI model.
func ModifyFabricAciObject(
	obj **resource_fabric_aci.NDFCFabricAciModel,
	values map[string]interface{},
) {
	fabric := *obj
	if fabric == nil {
		return
	}

	applyFabricAciValues(fabric, values)
	*obj = fabric
}

func applyFabricAciValues(fabric *resource_fabric_aci.NDFCFabricAciModel, values map[string]interface{}) {
	for key, value := range values {
		switch key {
		case "latitude":
			latitude := value.(float64)
			fabric.Spec.Location.Latitude = &latitude
		case "longitude":
			longitude := value.(float64)
			fabric.Spec.Location.Longitude = &longitude
		case "verify_ca":
			verifyCA := value.(bool)
			fabric.Spec.Aci.VerifyCa = &verifyCA
		case "security_domain":
			fabric.Spec.Aci.SecurityDomain = value.(string)
		case "license_tier":
			fabric.Spec.Aci.LicenseTier = value.(string)
		case "orchestration_status":
			fabric.Spec.Aci.Orchestration.OrchestrationStatus = value.(string)
		}
	}
}
