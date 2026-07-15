// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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

// Cluster API endpoints
const (
	UrlCluster             = "/infra/clusters"
	UrlClusterByName       = "/infra/clusters/%s"
	UrlClusterRemoveByName = "/infra/clusters/%s/remove"
)

const RscNameMultiClusterConnectivity = "multi_cluster_connectivity"

type ClusterAPI struct {
	ndapi.NexusDashboardAPICommon
	ClusterName string
	Delete      bool
}

func NewClusterAPI(client *nd.Client) *ClusterAPI {
	papi := new(ClusterAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *ClusterAPI) GetUrl() string {
	if c.ClusterName != "" {
		return fmt.Sprintf(UrlClusterByName, c.ClusterName)
	}
	return UrlCluster
}

func (c *ClusterAPI) PostUrl() string {
	if c.Delete {
		return c.DeleteUrl()
	}
	return UrlCluster
}

func (c *ClusterAPI) PutUrl() string {
	return fmt.Sprintf(UrlClusterByName, c.ClusterName)
}

func (c *ClusterAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName)
}

// GetDeleteQP satisfies the shared NexusDashboardAPI interface. This resource
// does not use query parameters for delete-style operations.
func (c *ClusterAPI) GetDeleteQP() []string {
	return nil
}

func (c *ClusterAPI) RscName() string {
	return RscNameMultiClusterConnectivity
}
