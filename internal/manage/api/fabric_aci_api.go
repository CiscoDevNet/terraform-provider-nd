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

// Fabric ACI API endpoints.
const (
	UrlAciCluster             = "/infra/clusters"
	UrlAciClusterByName       = "/infra/clusters/%s"
	UrlAciClusterRemoveByName = "/infra/clusters/%s/remove"
	UrlAciFabricByName        = "/manage/fabrics/%s"
)

const RscNameFabricAci = "fabric_aci"

type FabricAciAPI struct {
	ndapi.NexusDashboardAPICommon
	ClusterName string
	Delete      bool
}

func NewFabricAciAPI(client *nd.Client) *FabricAciAPI {
	papi := new(FabricAciAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *FabricAciAPI) GetUrl() string {
	if c.ClusterName != "" {
		return fmt.Sprintf(UrlAciClusterByName, c.ClusterName)
	}
	return UrlAciCluster
}

func (c *FabricAciAPI) PostUrl() string {
	if c.Delete {
		return c.DeleteUrl()
	}
	return UrlAciCluster
}

func (c *FabricAciAPI) PutUrl() string {
	return fmt.Sprintf(UrlAciFabricByName, c.ClusterName)
}

func (c *FabricAciAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlAciClusterRemoveByName, c.ClusterName)
}

func (c *FabricAciAPI) GetDeleteQP() []string {
	return nil
}

func (c *FabricAciAPI) RscName() string {
	return RscNameFabricAci
}
