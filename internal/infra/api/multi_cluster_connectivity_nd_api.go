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
	"sync"

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

type ClusterAPI struct {
	NDInfraAPICommon
	mutex       *sync.Mutex
	ClusterName string
}

func NewClusterAPI(lock *sync.Mutex, client *nd.Client) *ClusterAPI {
	papi := new(ClusterAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDInfraAPI = papi
	return papi
}

func (c *ClusterAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *ClusterAPI) GetUrl() string {
	if c.ClusterName != "" {
		return fmt.Sprintf(UrlClusterByName, c.ClusterName)
	}
	return UrlCluster
}

func (c *ClusterAPI) PostUrl() string {
	return UrlCluster
}

func (c *ClusterAPI) PutUrl() string {
	return fmt.Sprintf(UrlClusterByName, c.ClusterName)
}

func (c *ClusterAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName)
}

func (c *ClusterAPI) GetDeleteQP() []string {
	return nil
}

func (c *ClusterAPI) PostDelete(payload []byte) (nd.Res, error) {
	lock := c.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}

	res, err := c.Client.Post(fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *ClusterAPI) RscName() string {
	return "multi_cluster_connectivity"
}
