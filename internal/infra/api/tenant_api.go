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

// Tenant API endpoints
const (
	UrlTenants      = "/infra/tenants"
	UrlTenantByName = "/infra/tenants/%s"
)

const RscNameTenant = "tenant"

type TenantAPI struct {
	ndapi.NexusDashboardAPICommon
	TenantName string
}

func NewTenantAPI(client *nd.Client) *TenantAPI {
	papi := new(TenantAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *TenantAPI) GetUrl() string {
	if c.TenantName != "" {
		return fmt.Sprintf(UrlTenantByName, c.TenantName)
	}
	return UrlTenants
}

func (c *TenantAPI) PostUrl() string {
	return UrlTenants
}

func (c *TenantAPI) PutUrl() string {
	return fmt.Sprintf(UrlTenantByName, c.TenantName)
}

func (c *TenantAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlTenantByName, c.TenantName)
}

func (c *TenantAPI) GetDeleteQP() []string {
	return nil
}

func (c *TenantAPI) RscName() string {
	return RscNameTenant
}
