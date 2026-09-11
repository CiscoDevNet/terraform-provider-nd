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

// Change Control API endpoints
const (
	UrlChangeControl = "/infra/settings/fabricManagement/changeControl"
)

const RscNameChangeControl = "change_control"

// ChangeControlAPI is the API client for change control system settings.
type ChangeControlAPI struct {
	ndapi.NexusDashboardAPICommon
}

func NewChangeControlAPI(client *nd.Client, fabric string) *ChangeControlAPI {
	papi := new(ChangeControlAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	papi.Fabric = fabric
	return papi
}

func (c *ChangeControlAPI) GetUrl() string {
	return UrlChangeControl
}

func (c *ChangeControlAPI) PutUrl() string {
	return UrlChangeControl
}

func (c *ChangeControlAPI) GetDeleteQP() []string {
	return nil
}

func (c *ChangeControlAPI) RscName() string {
	return RscNameChangeControl
}
