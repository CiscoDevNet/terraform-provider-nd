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
	"terraform-provider-nd/internal/infra/resource_tenant"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TenantModelHelperStateCheck(RscName string, c resource_tenant.NDFCTenantModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.Name != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("name").String(), c.Name))
	}
	if c.Description != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("description").String(), c.Description))
	}
	for key, value := range c.FabricAssociations {
		attrNewPath := attrPath.AtName("fabric_associations").AtName(key)
		ret = append(ret, FabricAssociationsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func FabricAssociationsValueHelperStateCheck(RscName string, c resource_tenant.NDFCFabricAssociationsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.TenantPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_prefix").String(), c.TenantPrefix))
	}
	if c.LocalName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local_name").String(), c.LocalName))
	}
	return ret
}
