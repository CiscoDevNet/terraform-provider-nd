// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

// NDFCRemoteStorageLocationTestData mirrors the schema attributes of the
// nd_remote_storage_location resource. It is used by the gotmpl renderer
// and the state-check helper in the provider test package.
//
// All optional/sensitive fields use pointers so the template can decide
// whether to emit the corresponding HCL attribute.
type NDFCRemoteStorageLocationTestData struct {
	Name                    string
	Description             string
	StorageLocationType     string
	ReadWrite               *bool
	Hostname                string
	Port                    *int64
	Path                    string
	AlertThreshold          *int64
	Limit                   string
	Username                string
	Password                string
	SshKey                  string
	Passphrase              string
	IgnoreHostKeyValidation *bool
}

// GenerateRemoteStorageLocationObject builds a fresh
// NDFCRemoteStorageLocationTestData from a values map. Only keys present in
// the map are set; everything else is left at the Go zero value (and the
// template suppresses unset attributes).
func GenerateRemoteStorageLocationObject(
	obj **NDFCRemoteStorageLocationTestData,
	values map[string]interface{},
) {
	rsl := new(NDFCRemoteStorageLocationTestData)
	applyRemoteStorageLocationValues(rsl, values)
	*obj = rsl
}

// ModifyRemoteStorageLocationObject mutates an existing model with a new
// values map. Used between create/update steps to swap configuration.
//
// Pointer-typed fields are reset to nil first so callers can drop an
// attribute by simply omitting it from the next values map.
func ModifyRemoteStorageLocationObject(
	obj **NDFCRemoteStorageLocationTestData,
	values map[string]interface{},
) {
	rsl := *obj
	if rsl == nil {
		rsl = new(NDFCRemoteStorageLocationTestData)
	}
	*rsl = NDFCRemoteStorageLocationTestData{}
	applyRemoteStorageLocationValues(rsl, values)
	*obj = rsl
}

// applyRemoteStorageLocationValues is the shared key→field mapper used by
// both Generate and Modify so attribute handling stays in sync.
func applyRemoteStorageLocationValues(
	rsl *NDFCRemoteStorageLocationTestData,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "name":
			rsl.Name = val.(string)
		case "description":
			rsl.Description = val.(string)
		case "storage_location_type":
			rsl.StorageLocationType = val.(string)
		case "read_write":
			v := val.(bool)
			rsl.ReadWrite = &v
		case "hostname":
			rsl.Hostname = val.(string)
		case "port":
			v := int64(val.(int))
			rsl.Port = &v
		case "path":
			rsl.Path = val.(string)
		case "alert_threshold":
			v := int64(val.(int))
			rsl.AlertThreshold = &v
		case "limit":
			rsl.Limit = val.(string)
		case "username":
			rsl.Username = val.(string)
		case "password":
			rsl.Password = val.(string)
		case "ssh_key":
			rsl.SshKey = val.(string)
		case "passphrase":
			rsl.Passphrase = val.(string)
		case "ignore_host_key_validation":
			v := val.(bool)
			rsl.IgnoreHostKeyValidation = &v
		}
	}
}
