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
	"testing"
	"time"

	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
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
}

// TestAccRemoteStorageLocationResourceNAS exercises the NFS remote storage
// location lifecycle: create -> update -> import -> destroy.
//
// The update step changes description, read_write, and limit. `port` is
// required by the schema, so it is kept at 2049 in both steps.
func TestAccRemoteStorageLocationResourceNAS(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "nas_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.nas_test"

	s1 := &stepInfo{}
	s2 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create NAS remote storage location.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_nas", t.Name(), *stepCount)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  "nas",
						"description":           "nas_storage_location",
						"storage_location_type": "nfs",
						"read_write":            true,
						"hostname":              host,
						"port":                  2049,
						"path":                  "/export/path/",
						"limit":                 "10MB",
						"alert_threshold":       70,
					})

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
				),
			},
			// Step 2: In-place update of description, read_write, and limit.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_update_nas", t.Name(), *stepCount)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                  "nas",
						"description":           "nas_storage_location_updated",
						"storage_location_type": "nfs",
						"read_write":            false,
						"hostname":              host,
						"port":                  2049,
						"path":                  "/export/path/",
						"limit":                 "10GB",
						"alert_threshold":       70,
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, 2, s2.name, s2.cfg)
					t.Logf("Sleeping 90 seconds after NAS create before update to let the controller settle")
					time.Sleep(90 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
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
				ImportStateId:                        "nas",
				ImportStateVerifyIdentifierAttribute: "name",
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

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "scp_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.scp_test"

	s1 := &stepInfo{}
	s2 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SCP location with username/password.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_scp", t.Name(), *stepCount)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "scp-sftp",
						"storage_location_type":      "scp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   username,
						"password":                   password,
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
				),
			},
			// Step 2: Switch to SFTP and update path. Replaces the resource
			// because storage_location_type/username/password have
			// RequiresReplace plan modifiers.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_update_sftp", t.Name(), *stepCount)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "scp-sftp",
						"description":                "sftp_storage_location",
						"storage_location_type":      "sftp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/home/cisco/tmp1",
						"username":                   username,
						"password":                   password,
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
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
				ImportStateId:                        "scp-sftp",
				ImportStateVerifyIdentifierAttribute: "name",
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

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "scp_ssh_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	rscAddr := "nd_remote_storage_location.scp_ssh_test"

	s1 := &stepInfo{}
	s2 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create SCP location with SSH key + passphrase.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_scp_ssh", t.Name(), *stepCount)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "scp-sftp-ssh",
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

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
				),
			},
			// Step 2: Switch to SFTP, keep SSH-key auth, change path.
			// Replaces because storage_location_type/ssh_key are RequiresReplace.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_update_sftp_ssh", t.Name(), *stepCount)

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "scp-sftp-ssh",
						"storage_location_type":      "sftp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/home/cisco/tmp1",
						"username":                   username,
						"ssh_key":                    sshKey,
						"passphrase":                 passphrase,
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
				Check: resource.ComposeTestCheckFunc(
					RemoteStorageLocationModelHelperStateCheck(rscAddr, *rsc, path.Empty())...,
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
				ImportStateId:                        "scp-sftp-ssh",
				ImportStateVerifyIdentifierAttribute: "name",
				ImportStateVerifyIgnore:              importIgnoreSensitiveFields,
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceAuthConflicts verifies plan-time
// validation for mutually exclusive SCP/SFTP authentication fields.
func TestAccRemoteStorageLocationResourceAuthConflicts(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "auth_conflict_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	rsc := new(helper.NDFCRemoteStorageLocationTestData)

	s1 := &stepInfo{}
	s2 := &stepInfo{}
	authConflictErr := regexp.MustCompile("Invalid authentication configuration")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: password and ssh_key cannot be configured together.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_password_ssh_key_conflict", t.Name(), *stepCount)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "auth-conflict-password-ssh-key",
						"storage_location_type":      "scp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   "testuser",
						"password":                   "test-password",
						"ssh_key":                    "test-private-key",
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
			// Step 2: password and passphrase cannot be configured together.
			{
				Config: func() string {
					*stepCount++
					s2.name = fmt.Sprintf("%s_%d_password_passphrase_conflict", t.Name(), *stepCount)

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":                       "auth-conflict-password-passphrase",
						"storage_location_type":      "sftp",
						"hostname":                   host,
						"port":                       22,
						"path":                       "/tmp",
						"username":                   "testuser",
						"password":                   "test-password",
						"passphrase":                 "test-passphrase",
						"ignore_host_key_validation": true,
					})

					helper.GetTFConfigWithSingleResource(s2.name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, 2, s2.name, s2.cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
		},
	})
}
