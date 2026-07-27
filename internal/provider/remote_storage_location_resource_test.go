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
	"strconv"
	"testing"
	"time"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// importIgnoreSensitiveFields lists the sensitive attributes that the ND
// remote storage location API does not return on GET. Imports cannot
// reconstruct them, so ImportStateVerify must skip these.
var importIgnoreSensitiveFields = []string{
	"username",
	"password",
	"ssh_key",
	"passphrase",
	"ignore_host_key_validation",
	"accept_host_key",
}

// TestAccRemoteStorageLocationResourceNAS exercises the NFS remote storage
// location lifecycle: create -> update -> import -> destroy.
//
// The update step changes description, read_write, and limit, and omits
// alert_threshold so the API default is read back.
func TestAccRemoteStorageLocationResourceNAS(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	locationName := fmt.Sprintf("nas-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "nas_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.nas_test"

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create NAS remote storage location.
			{
				Config: func() string {
					s1.Name = fmt.Sprintf("%s_%d_create_nas", t.Name(), 1)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  locationName,
						"description":           "nas_storage_location",
						"storage_location_type": "nfs",
						"read_write":            true,
						"hostname":              host,
						"port":                  2049,
						"path":                  "/mnt/tank/nfsstore",
						"limit":                 "10MB",
						"alert_threshold":       70, // non-default value
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			// Step 2: In-place update of description, read_write, and limit.
			// Omitting read_write, alert_threshold should reset it to the backend default.
			{
				Config: func() string {
					s2.Name = fmt.Sprintf("%s_%d_update_nas", t.Name(), 2)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  locationName,
						"description":           "nas_storage_location_updated",
						"storage_location_type": "nfs",
						"hostname":              host,
						"port":                  2049,
						"path":                  "/mnt/tank/nfsstore",
						"limit":                 "10GB",
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, 2, s2.Name, s2.Cfg)
					t.Logf("Sleeping 90 seconds after NAS create before update to let the controller settle")
					time.Sleep(90 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "alert_threshold", "80"),
						resource.TestCheckResourceAttr(rscAddr, "read_write", "false"),
					)...,
				),
			},
			// Step 3: ImportState verification.
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_import_nas =====", t.Name())
					t.Logf("Sleeping 90 seconds before NAS import to let the controller settle")
					time.Sleep(90 * time.Second)
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore:              importIgnoreSensitiveFields,
			},
			// Step 4: Refresh-only wait. The NAS remote storage delete API
			// occasionally returns HTTP 500 "There was a problem proxying
			// the request" if the controller has not finished settling the
			// previous update. Sleep before the framework's post-test
			// destroy runs so the cleanup DELETE succeeds.
			{
				PreConfig: func() {
					t.Logf("===== STEP 4: %s_4_wait_before_destroy =====", t.Name())
					t.Logf("Sleeping 90 seconds before post-test destroy of NAS remote storage")
					time.Sleep(90 * time.Second)
				},
				RefreshState: true,
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceSCP exercises the SCP/SFTP remote
// storage location lifecycle with password authentication:
// create (scp) -> update (sftp) -> import -> destroy.
//
// Both `storage_location_type` and `username`/`password` have
// RequiresReplace plan modifiers, so the update step destroys & recreates
// the resource - this still validates the full HCL surface.
func TestAccRemoteStorageLocationResourceSCP(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	password := cfg.ND.RemoteStorage.Password
	locationName := fmt.Sprintf("scp-sftp-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "scp_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.scp_test"

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SCP location with username/password.
			{
				Config: func() string {
					s1.Name = fmt.Sprintf("%s_%d_create_scp", t.Name(), 1)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  locationName,
						"storage_location_type": "scp",
						"hostname":              host,
						"port":                  22,
						"path":                  "/tmp",
						"username":              username,
						"password":              password,
						"accept_host_key":       true,
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			// Step 2: Switch to SFTP and update path. Replaces the resource
			// because storage_location_type/username/password have
			// RequiresReplace plan modifiers.
			{
				Config: func() string {
					s2.Name = fmt.Sprintf("%s_%d_update_sftp", t.Name(), 2)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  locationName,
						"description":           "sftp_storage_location",
						"storage_location_type": "sftp",
						"hostname":              host,
						"port":                  22,
						"path":                  "/tmp",
						"username":              username,
						"password":              password,
						"accept_host_key":       true,
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			// Step 3: ImportState verification.
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_import_scp =====", t.Name())
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore:              importIgnoreSensitiveFields,
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceSCPWithSSH exercises the SCP/SFTP
// remote storage location lifecycle with SSH-key authentication:
// create (scp + ssh_key) -> update (sftp + ssh_key) -> import -> destroy.
func TestAccRemoteStorageLocationResourceSCPWithSSH(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	sshKey := cfg.ND.RemoteStorage.SshKey
	passphrase := cfg.ND.RemoteStorage.Passphrase
	locationName := fmt.Sprintf("scp-sftp-ssh-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "scp_ssh_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.scp_ssh_test"

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SCP location with SSH key + passphrase.
			{
				Config: func() string {
					s1.Name = fmt.Sprintf("%s_%d_create_scp_ssh", t.Name(), 1)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       locationName,
						"description":                "scp_storage_location_ssh",
						"storage_location_type":      "scp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   username,
						"ssh_key":                    sshKey,
						"passphrase":                 passphrase,
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			// Step 2: Switch to SFTP, keep SSH-key auth, change path.
			// Replaces because storage_location_type/ssh_key are RequiresReplace.
			{
				Config: func() string {
					s2.Name = fmt.Sprintf("%s_%d_update_sftp_ssh", t.Name(), 2)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       locationName,
						"storage_location_type":      "sftp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   username,
						"ssh_key":                    sshKey,
						"passphrase":                 passphrase,
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			// Step 3: ImportState verification.
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_import_scp_ssh =====", t.Name())
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore:              importIgnoreSensitiveFields,
			},
		},
	})
}

func remoteStorageLocationStateChecks(resourceName string, rsc helper.NDFCRemoteStorageLocationTestData) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{}

	if rsc.Name != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "name", rsc.Name))
	}
	if rsc.Description != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "description", rsc.Description))
	}
	if rsc.StorageLocationType != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "storage_location_type", rsc.StorageLocationType))
	}
	if rsc.ReadWrite != nil {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "read_write", strconv.FormatBool(*rsc.ReadWrite)))
	}
	if rsc.Hostname != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "hostname", rsc.Hostname))
	}
	if rsc.Port != nil {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "port", strconv.Itoa(int(*rsc.Port))))
	}
	if rsc.Path != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "path", rsc.Path))
	}
	if rsc.AlertThreshold != nil {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "alert_threshold", strconv.Itoa(int(*rsc.AlertThreshold))))
	}
	if rsc.Limit != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "limit", rsc.Limit))
	}
	if rsc.Username != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "username", rsc.Username))
	}
	if rsc.Password != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "password", rsc.Password))
	}
	if rsc.SshKey != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "ssh_key", rsc.SshKey))
	}
	if rsc.Passphrase != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "passphrase", rsc.Passphrase))
	}
	if rsc.IgnoreHostKeyValidation != nil {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "ignore_host_key_validation", strconv.FormatBool(*rsc.IgnoreHostKeyValidation)))
	}
	if rsc.AcceptHostKey != nil {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "accept_host_key", strconv.FormatBool(*rsc.AcceptHostKey)))
	}
	if rsc.Name != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "id", rsc.Name))
	}
	return checks
}

// TestAccRemoteStorageLocationResourceAuthConflicts verifies plan-time
// validation for mutually exclusive SCP/SFTP authentication fields.
func TestAccRemoteStorageLocationResourceAuthConflicts(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	passwordSSHKeyConflictName := fmt.Sprintf("auth-conflict-password-ssh-key-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	passwordPassphraseConflictName := fmt.Sprintf("auth-conflict-password-passphrase-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	readWriteSCPConflictName := fmt.Sprintf("read-write-scp-conflict-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "auth_conflict_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	authConflictErr := regexp.MustCompile("Invalid Attribute Combination")
	readWriteTypeErr := regexp.MustCompile("Attribute `read_write` can only be set when `storage_location_type` is `nfs`")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: password and ssh_key cannot be configured together.
			{
				Config: func() string {
					s1.Name = fmt.Sprintf("%s_%d_password_ssh_key_conflict", t.Name(), 1)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       passwordSSHKeyConflictName,
						"storage_location_type":      "scp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   "testuser",
						"password":                   "test-password",
						"ssh_key":                    "test-private-key",
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, 1, s1.Name, s1.Cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
			// Step 2: password and passphrase cannot be configured together.
			{
				Config: func() string {
					s2.Name = fmt.Sprintf("%s_%d_password_passphrase_conflict", t.Name(), 2)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       passwordPassphraseConflictName,
						"storage_location_type":      "sftp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   "testuser",
						"password":                   "test-password",
						"passphrase":                 "test-passphrase",
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, 2, s2.Name, s2.Cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
			// Step 3: read_write is only valid for NFS remote storage locations.
			{
				Config: func() string {
					s3.Name = fmt.Sprintf("%s_%d_read_write_scp_conflict", t.Name(), 3)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  readWriteSCPConflictName,
						"storage_location_type": "scp",
						"hostname":              host,
						"port":                  22,
						"path":                  "/tmp",
						"read_write":            false,
						"username":              "testuser",
						"password":              "test-password",
					})

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, 3, s3.Name, s3.Cfg) },
				PlanOnly:    true,
				ExpectError: readWriteTypeErr,
			},
		},
	})
}
