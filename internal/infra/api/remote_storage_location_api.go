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

// Remote Storage Location API endpoints
const (
	UrlRemoteStorageLocation       = "/infra/remoteStorage"
	UrlRemoteStorageLocationByName = "/infra/remoteStorage/%s"
)

const RscNameRemoteStorageLocation = "remote_storage_location"

type RemoteStorageLocationAPI struct {
	ndapi.NexusDashboardAPICommon
	Name          string
	AcceptHostKey bool
}

func NewRemoteStorageLocationAPI(client *nd.Client) *RemoteStorageLocationAPI {
	papi := new(RemoteStorageLocationAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *RemoteStorageLocationAPI) GetUrl() string {
	if c.Name != "" {
		return fmt.Sprintf(UrlRemoteStorageLocationByName, c.Name)
	}
	return UrlRemoteStorageLocation
}

func (c *RemoteStorageLocationAPI) PostUrl() string {
	return c.withAcceptHostKeyQuery(UrlRemoteStorageLocation)
}

func (c *RemoteStorageLocationAPI) PutUrl() string {
	return c.withAcceptHostKeyQuery(fmt.Sprintf(UrlRemoteStorageLocationByName, c.Name))
}

func (c *RemoteStorageLocationAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlRemoteStorageLocationByName, c.Name)
}

// Remote Storage Location Delete API does not support query params.
func (c *RemoteStorageLocationAPI) GetDeleteQP() []string {
	return nil
}

func (c *RemoteStorageLocationAPI) RscName() string {
	return RscNameRemoteStorageLocation
}

func (c *RemoteStorageLocationAPI) withAcceptHostKeyQuery(url string) string {
	if c.AcceptHostKey {
		return fmt.Sprintf("%s?acceptHostKey=true", url)
	}
	return url
}
