// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/infra/resource_backup"
)

// defaultBackupValues returns sensible defaults for an nd_backup resource.
// Tests can override any of these via the overrides map passed to
// GenerateBackupObject.
//
// `encryption_key` is mandatory and must be supplied by the caller (it has
// no useful default). `telemetry_data` is intentionally omitted from the
// defaults: it is only valid when type=full and destination is a NAS remote
// location, so tests should set it explicitly when needed.
func defaultBackupValues() map[string]interface{} {
	return map[string]interface{}{
		"type":        "configOnly",
		"destination": "",
	}
}

// GenerateBackupObject creates a backup model object for testing.
//   - name and encryptionKey are mandatory.
//   - overrides lets each test supply unique values for any field listed in
//     defaultBackupValues() (plus `telemetry_data`). Any key not present in
//     overrides gets the value from defaultBackupValues().
func GenerateBackupObject(
	obj **resource_backup.NDFCBackupModel,
	name string,
	encryptionKey string,
	overrides map[string]interface{},
) {
	backup := new(resource_backup.NDFCBackupModel)

	backup.Name = name
	backup.EncryptionKey = encryptionKey

	merged := defaultBackupValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyBackupValues(backup, merged)

	*obj = backup
}

// applyBackupValues sets fields on a backup model from a key-value map.
func applyBackupValues(backup *resource_backup.NDFCBackupModel, values map[string]interface{}) {
	for key, val := range values {
		switch key {
		case "name":
			backup.Name = val.(string)
		case "type":
			backup.Type = val.(string)
		case "destination":
			backup.Destination = val.(string)
		case "encryption_key":
			backup.EncryptionKey = val.(string)
		case "telemetry_data":
			v := val.(bool)
			backup.TelemetryData = &v
		}
	}
}
