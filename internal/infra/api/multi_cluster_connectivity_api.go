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

// ClusterRemoveCredentials represents credentials for the cluster remove payload
type ClusterRemoveCredentials struct {
	LoginDomain string `json:"loginDomain,omitempty"`
	Password    string `json:"password,omitempty"`
	User        string `json:"user,omitempty"`
}

// ClusterRemovePayload represents the JSON body for the POST-based cluster remove endpoint
type ClusterRemovePayload struct {
	Credentials ClusterRemoveCredentials `json:"credentials"`
	Force       bool                     `json:"force"`
}

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

// Multi Cluster Delete API does not support query params so GetDeleteQP is not implemented for now, but keeping the logic in place in case it's needed in the future
func (c *ClusterAPI) GetDeleteQP() []string {
	return nil
}

func (c *ClusterAPI) RscName() string {
	return RscNameMultiClusterConnectivity
}
