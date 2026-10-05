// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInfraRemoteStorageLocationDataSourceNFS(t *testing.T) {
	testAccRemoteStorageLocationDataSource(t, "nfs", false)
}

func TestAccInfraRemoteStorageLocationDataSourceSCPWithPassword(t *testing.T) {
	testAccRemoteStorageLocationDataSource(t, "scp", false)
}

func TestAccInfraRemoteStorageLocationDataSourceSFTPWithPassword(t *testing.T) {
	testAccRemoteStorageLocationDataSource(t, "sftp", false)
}

func TestAccInfraRemoteStorageLocationDataSourceSCPWithSSHKey(t *testing.T) {
	testAccRemoteStorageLocationDataSource(t, "scp", true)
}

func TestAccInfraRemoteStorageLocationDataSourceSFTPWithSSHKey(t *testing.T) {
	testAccRemoteStorageLocationDataSource(t, "sftp", true)
}

// testAccRemoteStorageLocationDataSource creates a location, then reads it through
// the datasource and compares the common and protocol-specific API-backed fields.
func testAccRemoteStorageLocationDataSource(t *testing.T, protocol string, useSSHKey bool) {
	t.Helper()

	cfg := helper.GetConfig("global")
	storageConfig := cfg.ND.RemoteStorage
	locationName := fmt.Sprintf("ds-%s-%s", protocol, acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	resourceName := "nd_remote_storage_location.storage_test"
	dataSourceName := "data.nd_remote_storage_location.storage_test"

	x := map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "storage_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	values := map[string]interface{}{
		"name":        locationName,
		"description": fmt.Sprintf("%s datasource test storage", protocol),
		"hostname":    storageConfig.Hostname,
	}
	if protocol == "nfs" {
		values["path"] = "/mnt/tank/nfsstore"
		values["nfs"] = map[string]interface{}{
			"port":            2049,
			"limit":           "10MB",
			"read_write":      false,
			"alert_threshold": 80,
		}
	} else {
		values["path"] = "/tmp"
		authentication := map[string]interface{}{
			"protocol": protocol,
			"port":     22,
			"username": storageConfig.Username,
		}
		if useSSHKey {
			authentication["ssh_key"] = storageConfig.SshKey
			authentication["passphrase"] = storageConfig.Passphrase
			authentication["ignore_host_key_validation"] = true
		} else {
			authentication["password"] = storageConfig.Password
			authentication["accept_host_key"] = true
		}
		values["scp_sftp"] = authentication
	}

	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	helper.GenerateRemoteStorageLocationObject(&rsc, values)
	dataSource := &helper.RemoteStorageLocationDataSourceTestData{
		RscName:   "storage_test",
		Name:      locationName,
		DependsOn: resourceName,
	}

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(dataSourceName, "name", locationName),
		resource.TestCheckNoResourceAttr(dataSourceName, "id"),
		resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.password"),
		resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.ssh_key"),
		resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.passphrase"),
		resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.accept_host_key"),
	}
	attributes := []string{"name", "description", "hostname", "path"}
	if protocol == "nfs" {
		attributes = append(attributes, "nfs.port", "nfs.limit", "nfs.read_write", "nfs.alert_threshold")
		matchingChecks = append(matchingChecks,
			resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.protocol"),
			resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.port"),
			resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.username"),
			resource.TestCheckNoResourceAttr(dataSourceName, "scp_sftp.ignore_host_key_validation"),
		)
	} else {
		attributes = append(attributes, "scp_sftp.protocol", "scp_sftp.port", "scp_sftp.username", "scp_sftp.ignore_host_key_validation")
		matchingChecks = append(matchingChecks,
			resource.TestCheckResourceAttr(dataSourceName, "scp_sftp.protocol", protocol),
			resource.TestCheckNoResourceAttr(dataSourceName, "nfs.port"),
			resource.TestCheckNoResourceAttr(dataSourceName, "nfs.limit"),
			resource.TestCheckNoResourceAttr(dataSourceName, "nfs.read_write"),
			resource.TestCheckNoResourceAttr(dataSourceName, "nfs.alert_threshold"),
		)
	}
	for _, attribute := range attributes {
		matchingChecks = append(matchingChecks,
			resource.TestCheckResourceAttrPair(dataSourceName, attribute, resourceName, attribute),
		)
	}

	tfConfig := new(string)
	s1 := &helper.StepInfo{
		Index: 1,
		Name:  fmt.Sprintf("%s - Create the remote storage location used by datasource lookups", t.Name()),
	}
	helper.GetTFConfigWithSingleResource(s1.Name, x, []interface{}{rsc}, &tfConfig)
	s1.Cfg = *tfConfig
	s2 := &helper.StepInfo{
		Index: 2,
		Name:  fmt.Sprintf("%s - Read the location and match datasource attributes with the resource", t.Name()),
	}
	helper.GetTFConfigWithSingleResource(s2.Name, x, []interface{}{rsc, dataSource}, &tfConfig)
	s2.Cfg = *tfConfig

	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(storageConfig.Hostname) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if protocol != "nfs" {
				if strings.TrimSpace(storageConfig.Username) == "" {
					t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
				}
				if useSSHKey {
					if strings.TrimSpace(storageConfig.SshKey) == "" {
						t.Fatal("remote_storage.ssh_key must be configured in the acceptance-test testbed")
					}
				} else if strings.TrimSpace(storageConfig.Password) == "" {
					t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
				}
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:    s1.Cfg,
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(resourceName, *rsc)...,
				),
			},
			{
				Config: s2.Cfg,
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					if protocol == "nfs" {
						t.Logf("Sleeping 90 seconds after NFS remote storage location create to let the controller settle before datasource lookup and destroy")
						time.Sleep(90 * time.Second)
					}
				},
				Check: resource.ComposeTestCheckFunc(matchingChecks...),
			},
		},
	})
}

func TestAccInfraRemoteStorageLocationDataSourceMissing(t *testing.T) {
	cfg := helper.GetConfig("global")
	missingName := fmt.Sprintf("ds-missing-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	x := map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "missing_storage",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
	dataSource := &helper.RemoteStorageLocationDataSourceTestData{
		RscName: "missing_storage",
		Name:    missingName,
	}
	tfConfig := new(string)
	s1 := &helper.StepInfo{
		Index: 1,
		Name:  fmt.Sprintf("%s - Reject datasource lookup for a missing remote storage location", t.Name()),
	}
	helper.GetTFConfigWithSingleResource(s1.Name, x, []interface{}{dataSource}, &tfConfig)
	s1.Cfg = *tfConfig

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:    s1.Cfg,
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not read nd remote storage location with name\s+%q:\s+resource\s+not\s+found`, missingName),
				),
			},
		},
	})
}
