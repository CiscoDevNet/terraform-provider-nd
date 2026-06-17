// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

const UrlTenantFabricAssociations = "/manage/tenantFabricAssociations"

const RscNameTenantFabricAssociation = "tenant_fabric_association"

type TenantFabricAssociationAPI struct {
	ndapi.NexusDashboardAPICommon
}

func NewTenantFabricAssociationAPI(client *nd.Client) *TenantFabricAssociationAPI {
	papi := new(TenantFabricAssociationAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *TenantFabricAssociationAPI) GetUrl() string {
	return UrlTenantFabricAssociations
}

func (c *TenantFabricAssociationAPI) PostUrl() string {
	return UrlTenantFabricAssociations
}

func (c *TenantFabricAssociationAPI) PutUrl() string {
	return UrlTenantFabricAssociations
}

func (c *TenantFabricAssociationAPI) DeleteUrl() string {
	return UrlTenantFabricAssociations
}

func (c *TenantFabricAssociationAPI) GetDeleteQP() []string {
	return nil
}

func (c *TenantFabricAssociationAPI) RscName() string {
	return RscNameTenantFabricAssociation
}
