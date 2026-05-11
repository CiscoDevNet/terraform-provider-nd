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

// Local User API endpoints
const (
	UrlLocalUser          = "/infra/aaa/localUsers"
	UrlLocalUserByLoginId = "/infra/aaa/localUsers/%s"
)

const RscNameLocalUser = "local_user"

type LocalUserAPI struct {
	ndapi.NexusDashboardAPICommon
	LoginId string
}

func NewLocalUserAPI(client *nd.Client, fabric string) *LocalUserAPI {
	papi := new(LocalUserAPI)
	papi.Client = client
	papi.Fabric = fabric
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *LocalUserAPI) GetUrl() string {
	if c.LoginId != "" {
		return fmt.Sprintf(UrlLocalUserByLoginId, c.LoginId)
	}
	return UrlLocalUser
}

func (c *LocalUserAPI) PostUrl() string {
	return UrlLocalUser
}

func (c *LocalUserAPI) PutUrl() string {
	return fmt.Sprintf(UrlLocalUserByLoginId, c.LoginId)
}

func (c *LocalUserAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlLocalUserByLoginId, c.LoginId)
}

// Local User Delete API does not support query params.
func (c *LocalUserAPI) GetDeleteQP() []string {
	return nil
}

func (c *LocalUserAPI) RscName() string {
	return RscNameLocalUser
}
