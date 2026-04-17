// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package manage

import (
	nd "github.com/netascode/go-nd"
)

const ModuleKey = "manage"

type NexusDashboardManage struct {
	ApiClient *nd.Client
}

var manageInstance *NexusDashboardManage

func NewManage(client *nd.Client) *NexusDashboardManage {
	if manageInstance == nil {
		manageInstance = &NexusDashboardManage{
			ApiClient: client,
		}
	}
	return manageInstance
}
