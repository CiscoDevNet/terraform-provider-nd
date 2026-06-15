// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"

	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

// Tenant domain API endpoints
const (
	UrlTenantDomains      = "/infra/tenantDomains"
	UrlTenantDomainByName = "/infra/tenantDomains/%s"
)

const RscNameTenantDomain = "tenant_domain"

type TenantDomainAPI struct {
	ndapi.NexusDashboardAPICommon
	TenantDomainName string
}

func NewTenantDomainAPI(client *nd.Client, fabric string) *TenantDomainAPI {
	papi := new(TenantDomainAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	papi.Fabric = fabric
	return papi
}

func (c *TenantDomainAPI) GetUrl() string {
	if c.TenantDomainName != "" {
		return fmt.Sprintf(UrlTenantDomainByName, c.TenantDomainName)
	}
	return UrlTenantDomains
}

func (c *TenantDomainAPI) PostUrl() string {
	return UrlTenantDomains
}

func (c *TenantDomainAPI) PutUrl() string {
	return fmt.Sprintf(UrlTenantDomainByName, c.TenantDomainName)
}

func (c *TenantDomainAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlTenantDomainByName, c.TenantDomainName)
}

func (c *TenantDomainAPI) GetDeleteQP() []string {
	return nil
}

func (c *TenantDomainAPI) RscName() string {
	return RscNameTenantDomain
}
