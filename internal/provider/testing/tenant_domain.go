// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import "terraform-provider-nd/internal/infra/resource_tenant_domain"

// TenantDomainTestResource pairs a tenant-domain model with test-only
// dependency metadata used while rendering its Terraform configuration.
type TenantDomainTestResource struct {
	Model     *resource_tenant_domain.NDFCTenantDomainModel
	DependsOn string
}

// TenantDomainDataSourceTestData contains the values rendered by the
// tenant-domain datasource acceptance-test template.
type TenantDomainDataSourceTestData struct {
	RscName   string
	Name      string
	DependsOn string
}

// GenerateTenantDomainObject builds a fresh tenant-domain model from the
// supplied Terraform attribute values.
func GenerateTenantDomainObject(
	obj **resource_tenant_domain.NDFCTenantDomainModel,
	values map[string]interface{},
) {
	tenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	applyTenantDomainValues(tenantDomain, values)
	*obj = tenantDomain
}

// ModifyTenantDomainObject replaces the configurable tenant-domain values.
// Attributes omitted from values are cleared from the rendered configuration.
func ModifyTenantDomainObject(
	obj **resource_tenant_domain.NDFCTenantDomainModel,
	values map[string]interface{},
) {
	tenantDomain := new(resource_tenant_domain.NDFCTenantDomainModel)
	applyTenantDomainValues(tenantDomain, values)
	*obj = tenantDomain
}

func applyTenantDomainValues(
	tenantDomain *resource_tenant_domain.NDFCTenantDomainModel,
	values map[string]interface{},
) {
	for key, value := range values {
		switch key {
		case "name":
			tenantDomain.Name = value.(string)
			tenantDomain.Id = tenantDomain.Name
		case "description":
			tenantDomain.Description = value.(string)
		case "tenant_names":
			tenantNames := value.([]string)
			tenantDomain.TenantNames = append([]string(nil), tenantNames...)
		}
	}
}
