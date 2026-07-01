// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"encoding/json"
	"fmt"

	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

// Backup API endpoints
const (
	UrlBackup        = "/infra/backups"
	UrlBackupById    = "/infra/backups/%s"
	UrlBackupStatus  = "/infra/backups/status"
	UrlBackupImport  = "/infra/backups/actions/import"
	UrlBackupRestore = "/infra/backups/actions/restore"
)

const RscNameBackup = "backup"

type BackupAPI struct {
	ndapi.NexusDashboardAPICommon
	Name string
}

type BackupRestorePayload struct {
	IgnorePersistentIPs             bool   `json:"ignorePersistentIPs"`
	Type                            string `json:"type"`
	IncludeTelemetryOperationalData bool   `json:"includeTelemetryOperationalData"`
}

type BackupImportPayload struct {
	EncryptionKey string `json:"encryptionKey"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	Source        string `json:"source"`
}

type BackupStatus struct {
	Raw       json.RawMessage     `json:"-"`
	Id        string              `json:"id"`
	Error     string              `json:"error"`
	Operation string              `json:"operation"`
	State     string              `json:"state"`
	Details   BackupStatusDetails `json:"details"`
}

type BackupStatusDetails struct {
	Progress *int `json:"progress"`
}

func NewBackupAPI(client *nd.Client) *BackupAPI {
	papi := new(BackupAPI)
	papi.Client = client
	papi.NexusDashboardAPI = papi
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

func (c *BackupAPI) ClearImportedBackup() (string, error) {
	res, err := c.Client.Delete(UrlBackupImport, "")
	return fmt.Sprint(res), err
}

func (c *BackupAPI) ImportBackup(payload []byte) (string, error) {
	res, err := c.Client.Post(UrlBackupImport, string(payload), nd.NoLogPayload)
	return fmt.Sprint(res), err
}

func (c *BackupAPI) Restore(payload []byte) (string, error) {
	res, err := c.Client.Post(UrlBackupRestore, string(payload))
	return fmt.Sprint(res), err
}

func (c *BackupAPI) Status() (BackupStatus, error) {
	res, err := c.Client.GetRawJson(UrlBackupStatus)
	if err != nil {
		return BackupStatus{}, err
	}

	status := BackupStatus{Raw: append(json.RawMessage(nil), res...)}
	if len(res) == 0 {
		return status, nil
	}

	err = json.Unmarshal(res, &status)
	return status, err
}

// Backup Delete API does not support query params.
func (c *BackupAPI) GetDeleteQP() []string {
	return nil
}

func (c *BackupAPI) RscName() string {
	return RscNameBackup
}
