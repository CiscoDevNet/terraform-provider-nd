// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/infra/resource_tenant"
)

func defaultTenantValues() map[string]interface{} {
	return map[string]interface{}{}
}

// GenerateTenantObject creates a tenant model object for testing.
// name is the Terraform and ND tenant identifier. overrides lets each test
// supply optional fields such as description and fabric_associations.
func GenerateTenantObject(
	obj **resource_tenant.NDFCTenantModel,
	name string,
	overrides map[string]interface{},
) {
	tenant := new(resource_tenant.NDFCTenantModel)

	tenant.Id = name
	tenant.Name = name

	merged := defaultTenantValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyTenantValues(tenant, merged)

	*obj = tenant
}

// ModifyTenantObject modifies fields on an existing tenant model.
// Uses the same key set as GenerateTenantObject overrides.
func ModifyTenantObject(
	obj **resource_tenant.NDFCTenantModel,
	values map[string]interface{},
) {
	tenant := *obj
	if tenant == nil {
		return
	}

	applyTenantValues(tenant, values)

	*obj = tenant
}

func applyTenantValues(tenant *resource_tenant.NDFCTenantModel, values map[string]interface{}) {
	for key, val := range values {
		switch key {
		case "description":
			tenant.Description = val.(string)
		case "fabric_associations":
			tenant.FabricAssociations = val.(map[string]resource_tenant.NDFCFabricAssociationsValue)
		}
	}
}
