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

// Backup API endpoints
const (
	UrlBackup     = "/infra/backups"
	UrlBackupById = "/infra/backups/%s"
)

const RscNameBackup = "backup"

type BackupAPI struct {
	ndapi.NexusDashboardAPICommon
	Name string
}

func NewBackupAPI(client *nd.Client, fabric string) *BackupAPI {
	papi := new(BackupAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
	papi.Fabric = fabric
	return papi
}

func (c *BackupAPI) GetUrl() string {
	if c.Name != "" {
		return fmt.Sprintf(UrlBackupById, c.Name)
	}
	return UrlBackup
}

func (c *BackupAPI) PostUrl() string {
	return UrlBackup
}

func (c *BackupAPI) PutUrl() string {
	return fmt.Sprintf(UrlBackupById, c.Name)
}

func (c *BackupAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlBackupById, c.Name)
}

// Backup Delete API does not support query params.
func (c *BackupAPI) GetDeleteQP() []string {
	return nil
}

func (c *BackupAPI) RscName() string {
	return RscNameBackup
}
