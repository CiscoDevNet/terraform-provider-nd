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
type NDFCRemoteStorageLocationTestData struct {
	Name        string
	Description string
	Hostname    string
	Path        string
	Nfs         *NDFCRemoteStorageLocationNFSTestData
	ScpSftp     *NDFCRemoteStorageLocationSCPSFTPTestData
}

type NDFCRemoteStorageLocationNFSTestData struct {
	Port           *int64
	Limit          string
	ReadWrite      *bool
	AlertThreshold *int64
}

type NDFCRemoteStorageLocationSCPSFTPTestData struct {
	Protocol                string
	Port                    *int64
	Username                string
	Password                string
	SshKey                  string
	Passphrase              string
	IgnoreHostKeyValidation *bool
	AcceptHostKey           *bool
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
// The model is reset first so callers can remove a branch or nested attribute
// by omitting it from the next values map.
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
		case "hostname":
			rsl.Hostname = val.(string)
		case "path":
			rsl.Path = val.(string)
		case "nfs":
			nfs := new(NDFCRemoteStorageLocationNFSTestData)
			applyRemoteStorageLocationNFSValues(nfs, val.(map[string]interface{}))
			rsl.Nfs = nfs
		case "scp_sftp":
			scpSftp := new(NDFCRemoteStorageLocationSCPSFTPTestData)
			applyRemoteStorageLocationSCPSFTPValues(scpSftp, val.(map[string]interface{}))
			rsl.ScpSftp = scpSftp
		}
	}
}

func applyRemoteStorageLocationNFSValues(
	nfs *NDFCRemoteStorageLocationNFSTestData,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "port":
			v := int64(val.(int))
			nfs.Port = &v
		case "limit":
			nfs.Limit = val.(string)
		case "read_write":
			v := val.(bool)
			nfs.ReadWrite = &v
		case "alert_threshold":
			v := int64(val.(int))
			nfs.AlertThreshold = &v
		}
	}
}

func applyRemoteStorageLocationSCPSFTPValues(
	scpSftp *NDFCRemoteStorageLocationSCPSFTPTestData,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "protocol":
			scpSftp.Protocol = val.(string)
		case "port":
			v := int64(val.(int))
			scpSftp.Port = &v
		case "username":
			scpSftp.Username = val.(string)
		case "password":
			scpSftp.Password = val.(string)
		case "ssh_key":
			scpSftp.SshKey = val.(string)
		case "passphrase":
			scpSftp.Passphrase = val.(string)
		case "ignore_host_key_validation":
			v := val.(bool)
			scpSftp.IgnoreHostKeyValidation = &v
		case "accept_host_key":
			v := val.(bool)
			scpSftp.AcceptHostKey = &v
		}
	}
}
