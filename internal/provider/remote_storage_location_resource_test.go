// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/common/utils"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_remote_storage_location"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	nd "github.com/netascode/go-nd"
)

const (
	remoteStorageLocationTestDeletePollInterval = 10 * time.Second
	remoteStorageLocationTestDeletePollTimeout  = 5 * time.Minute
)

var passwordImportIgnoreFields = []string{
	"scp_sftp.password",
	"scp_sftp.accept_host_key",
}

var sshKeyImportIgnoreFields = []string{
	"scp_sftp.ssh_key",
	"scp_sftp.passphrase",
}

// TestAccRemoteStorageLocationResourceNAS exercises NFS-specific defaults,
// optional attributes, in-place updates, import, branch replacement, and
// destroy.
func TestAccRemoteStorageLocationResourceNAS(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	password := cfg.ND.RemoteStorage.Password
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
	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(password) == "" {
				t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create an NFS location with only required attributes")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/mnt/tank/nfsstore",
						"nfs": map[string]interface{}{
							"limit": "10MB",
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						remoteStorageLocationNFSDefaultChecks(rscAddr)...,
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Configure every NFS-specific optional attribute")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "nas_storage_location",
						"hostname":    host,
						"path":        "/mnt/tank/nfsstore",
						"nfs": map[string]interface{}{
							"read_write":      true,
							"port":            2049,
							"limit":           "10MB",
							"alert_threshold": 70,
						},
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					t.Logf("Sleeping 90 seconds after NAS create before update to let the controller settle")
					time.Sleep(90 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Remove NFS optional attributes and restore backend defaults")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/mnt/tank/nfsstore",
						"nfs": map[string]interface{}{
							"limit": "10MB",
						},
					})

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s3.Index, s3.Name, s3.Cfg)
					t.Logf("Sleeping 90 seconds before resetting NFS optional attributes")
					time.Sleep(90 * time.Second)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						remoteStorageLocationNFSDefaultChecks(rscAddr)...,
					)...,
				),
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan with NFS optional attributes omitted")

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				PlanOnly:  true,
			},
			{
				PreConfig: func() {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Import the NFS location and verify its API-backed state")
					helper.LogStep(t, s5.Index, s5.Name, "")
					t.Logf("Sleeping 90 seconds before NFS import to let the controller settle")
					time.Sleep(90 * time.Second)
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Replace the NFS location with an SCP location")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s6.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s6.Index, s6.Name, s6.Cfg)
					t.Logf("Sleeping 90 seconds before replacing the managed NFS location")
					time.Sleep(90 * time.Second)
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(rscAddr, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.ignore_host_key_validation", "false"),
					)...,
				),
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy the replacement SCP location")

					helper.GetTFConfigWithSingleResource(s7.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				Destroy:   true,
				PostApplyFunc: func() {
					assertRemoteStorageLocationAbsentOutsideTerraform(t, locationName)
				},
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceSCPWithPassword exercises password
// authentication, in-place updates, SCP-to-SFTP replacement, import, and
// destroy.
func TestAccRemoteStorageLocationResourceSCPWithPassword(t *testing.T) {
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
	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(password) == "" {
				t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create an SCP location with password authentication and a default port")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.ignore_host_key_validation", "false"),
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Update a mutable SCP attribute without replacing the location")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "scp_storage_location_updated",
						"hostname":    host,
						"path":        "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.ignore_host_key_validation", "false"),
					)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan after refreshing password-authenticated state")

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Replace the SCP location with SFTP while retaining password authentication")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "sftp_storage_location",
						"hostname":    host,
						"path":        "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "sftp",
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.ignore_host_key_validation", "false"),
					)...,
				),
			},
			{
				PreConfig: func() {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Import the password-authenticated SFTP location")
					helper.LogStep(t, s5.Index, s5.Name, "")
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore:              passwordImportIgnoreFields,
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify the managed password-authenticated state remains stable after import verification")

					helper.GetTFConfigWithSingleResource(s6.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s6.Index, s6.Name, s6.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy the managed password-authenticated location")

					helper.GetTFConfigWithSingleResource(s7.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				Destroy:   true,
				PostApplyFunc: func() {
					assertRemoteStorageLocationAbsentOutsideTerraform(t, locationName)
				},
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceSCPWithSSHKey exercises SSH-key
// authentication, in-place updates, SCP-to-SFTP replacement, import, and
// destroy.
func TestAccRemoteStorageLocationResourceSCPWithSSHKey(t *testing.T) {
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
	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(sshKey) == "" {
				t.Fatal("remote_storage.ssh_key must be configured in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create an SCP location with SSH-key authentication and a default port")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "scp_storage_location_ssh",
						"hostname":    host,
						"path":        "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "scp",
							"username":                   username,
							"ssh_key":                    sshKey,
							"passphrase":                 passphrase,
							"ignore_host_key_validation": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.accept_host_key", "false"),
					)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Update a mutable SCP attribute while retaining SSH-key authentication")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "scp_storage_location_ssh_updated",
						"hostname":    host,
						"path":        "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "scp",
							"username":                   username,
							"ssh_key":                    sshKey,
							"passphrase":                 passphrase,
							"ignore_host_key_validation": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.accept_host_key", "false"),
					)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify an empty plan after refreshing SSH-key authentication state")

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Replace the SCP location with SFTP while retaining SSH-key authentication")

					helper.ModifyRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "sftp",
							"username":                   username,
							"ssh_key":                    sshKey,
							"passphrase":                 passphrase,
							"ignore_host_key_validation": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				Check: resource.ComposeTestCheckFunc(
					append(
						remoteStorageLocationStateChecks(rscAddr, *rsc),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.port", "22"),
						resource.TestCheckResourceAttr(rscAddr, "scp_sftp.accept_host_key", "false"),
					)...,
				),
			},
			{
				PreConfig: func() {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Import the SSH-key-authenticated SFTP location")
					helper.LogStep(t, s5.Index, s5.Name, "")
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        locationName,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore:              sshKeyImportIgnoreFields,
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Verify the managed SSH-key-authenticated state remains stable after import verification")

					helper.GetTFConfigWithSingleResource(s6.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s6.Index, s6.Name, s6.Cfg) },
				PlanOnly:  true,
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy the managed SSH-key-authenticated location")

					helper.GetTFConfigWithSingleResource(s7.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				Destroy:   true,
				PostApplyFunc: func() {
					assertRemoteStorageLocationAbsentOutsideTerraform(t, locationName)
				},
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
	if rsc.Hostname != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "hostname", rsc.Hostname))
	}
	if rsc.Path != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "path", rsc.Path))
	}

	switch {
	case rsc.Nfs != nil:
		if rsc.Nfs.Port != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "nfs.port", strconv.Itoa(int(*rsc.Nfs.Port))))
		}
		if rsc.Nfs.Limit != "" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "nfs.limit", rsc.Nfs.Limit))
		}
		if rsc.Nfs.ReadWrite != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "nfs.read_write", strconv.FormatBool(*rsc.Nfs.ReadWrite)))
		}
		if rsc.Nfs.AlertThreshold != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "nfs.alert_threshold", strconv.Itoa(int(*rsc.Nfs.AlertThreshold))))
		}
		checks = append(checks, resource.TestCheckNoResourceAttr(resourceName, "scp_sftp.protocol"))
	case rsc.ScpSftp != nil:
		checks = append(checks,
			resource.TestCheckResourceAttr(resourceName, "scp_sftp.protocol", rsc.ScpSftp.Protocol),
			resource.TestCheckNoResourceAttr(resourceName, "nfs.limit"),
		)
		if rsc.ScpSftp.Port != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.port", strconv.Itoa(int(*rsc.ScpSftp.Port))))
		}
		if rsc.ScpSftp.Username != "" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.username", rsc.ScpSftp.Username))
		}
		if rsc.ScpSftp.Password != "" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.password", rsc.ScpSftp.Password))
		}
		if rsc.ScpSftp.SshKey != "" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.ssh_key", rsc.ScpSftp.SshKey))
		}
		if rsc.ScpSftp.Passphrase != "" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.passphrase", rsc.ScpSftp.Passphrase))
		}
		if rsc.ScpSftp.IgnoreHostKeyValidation != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.ignore_host_key_validation", strconv.FormatBool(*rsc.ScpSftp.IgnoreHostKeyValidation)))
		}
		if rsc.ScpSftp.AcceptHostKey != nil {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "scp_sftp.accept_host_key", strconv.FormatBool(*rsc.ScpSftp.AcceptHostKey)))
		}
	}
	if rsc.Name != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "id", rsc.Name))
	}
	return checks
}

func remoteStorageLocationNFSDefaultChecks(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "nfs.port", "2049"),
		resource.TestCheckResourceAttr(resourceName, "nfs.read_write", "false"),
		resource.TestCheckResourceAttr(resourceName, "nfs.alert_threshold", "80"),
		resource.TestCheckNoResourceAttr(resourceName, "description"),
		resource.TestCheckNoResourceAttr(resourceName, "scp_sftp.protocol"),
	}
}

func remoteStorageLocationSCPSFTPDefaultChecks(resourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "scp_sftp.port", "22"),
		resource.TestCheckResourceAttr(resourceName, "scp_sftp.ignore_host_key_validation", "false"),
		resource.TestCheckResourceAttr(resourceName, "scp_sftp.accept_host_key", "false"),
	}
}

// TestAccRemoteStorageLocationResourceCommonLifecycle exercises behavior that
// is shared by NFS, SCP, and SFTP locations: duplicate creation, drift,
// Read-404 recreation, and deletion by removing the resource from config. It
// uses SCP password authentication to avoid NFS backend synchronization delays.
func TestAccRemoteStorageLocationResourceCommonLifecycle(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	password := cfg.ND.RemoteStorage.Password
	locationName := fmt.Sprintf("lifecycle-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	rscAddr := "nd_remote_storage_location.lifecycle_test"

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "lifecycle_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(password) == "" {
				t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject creation when the remote storage location already exists")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":        locationName,
						"description": "terraform_managed_location",
						"hostname":    host,
						"path":        "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"port":            22,
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
					createRemoteStorageLocationOutsideTerraform(t, *rsc)
				},
				ExpectError: regexp.MustCompile(`(?is)Error Creating ND Remote Storage Location.*Could not create nd_remote_storage_location`),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the Terraform-managed remote storage location")

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
				},
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Detect an out-of-band description change")

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return s3.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s3.Index, s3.Name, s3.Cfg)
					drifted := *rsc
					drifted.Description = "changed_outside_terraform"
					updateRemoteStorageLocationOutsideTerraform(t, drifted)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Recreate the remote storage location after an out-of-band deletion")

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
				},
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Delete the remote storage location by removing it from configuration")

					helper.GetTFConfigWithSingleResource(s5.Name, *x, nil, &tfConfig)

					s5.Cfg = *tfConfig
					return s5.Cfg
				}(),
				PreConfig: func() { helper.LogStep(t, s5.Index, s5.Name, s5.Cfg) },
				PostApplyFunc: func() {
					assertRemoteStorageLocationAbsentOutsideTerraform(t, locationName)
				},
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceDeleteMissing verifies that Delete is
// idempotent when the Terraform-managed location is already absent remotely.
func TestAccRemoteStorageLocationResourceDeleteMissing(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	password := cfg.ND.RemoteStorage.Password
	locationName := fmt.Sprintf("delete-missing-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	rscAddr := "nd_remote_storage_location.delete_missing_test"

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "delete_missing_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	cleanupEnabled := false
	t.Cleanup(func() {
		if cleanupEnabled {
			deleteRemoteStorageLocationOutsideTerraform(t, locationName)
		}
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(password) == "" {
				t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
			}
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		AdditionalCLIOptions: &resource.AdditionalCLIOptions{
			Plan: resource.PlanOptions{NoRefresh: true},
		},
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create an SCP location for already-missing destroy verification")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"port":            22,
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
				},
				Check: resource.ComposeTestCheckFunc(
					remoteStorageLocationStateChecks(rscAddr, *rsc)...,
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Destroy successfully when the remote storage location is already missing")

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return s2.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
				},
				Destroy: true,
				PostApplyFunc: func() {
					assertRemoteStorageLocationAbsentOutsideTerraform(t, locationName)
				},
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceImportMissing verifies that import
// reports an explicit error for a location that does not exist.
func TestAccRemoteStorageLocationResourceImportMissing(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	username := cfg.ND.RemoteStorage.Username
	password := cfg.ND.RemoteStorage.Password
	locationName := fmt.Sprintf("missing-import-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	rscAddr := "nd_remote_storage_location.import_missing_test"

	x := &map[string]string{
		"RscType":  "nd_remote_storage_location",
		"RscName":  "import_missing_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	rsc := new(helper.NDFCRemoteStorageLocationTestData)
	s1 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			if strings.TrimSpace(host) == "" {
				t.Fatal("remote_storage.hostname must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(username) == "" {
				t.Fatal("remote_storage.username must be configured in the acceptance-test testbed")
			}
			if strings.TrimSpace(password) == "" {
				t.Fatal("remote_storage.password must be configured in the acceptance-test testbed")
			}
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject import of a missing remote storage location")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     locationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":        "scp",
							"port":            22,
							"username":        username,
							"password":        password,
							"accept_host_key": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return s1.Cfg
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteRemoteStorageLocationOutsideTerraform(t, locationName)
				},
				ResourceName:  rscAddr,
				ImportState:   true,
				ImportStateId: locationName,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(
						`(?is)Could\s+not\s+import\s+nd_remote_storage_location\s+with\s+id\s+%q:\s+resource\s+not\s+found`,
						locationName,
					),
				),
			},
		},
	})
}

// TestAccRemoteStorageLocationResourceAuthConflicts verifies profile-specific
// schema validation and mutually exclusive NFS/SCP/SFTP attributes.
func TestAccRemoteStorageLocationResourceAuthConflicts(t *testing.T) {
	cfg := helper.GetConfig("global")
	host := cfg.ND.RemoteStorage.Hostname
	passwordSSHKeyConflictName := fmt.Sprintf("auth-conflict-password-ssh-key-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	passwordPassphraseConflictName := fmt.Sprintf("auth-conflict-password-passphrase-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	missingAuthenticationName := fmt.Sprintf("missing-authentication-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	passphraseWithoutSSHKeyName := fmt.Sprintf("passphrase-without-ssh-key-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	hostKeyConflictName := fmt.Sprintf("host-key-conflict-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	invalidPortName := fmt.Sprintf("invalid-port-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	invalidThresholdName := fmt.Sprintf("invalid-threshold-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	invalidLimitName := fmt.Sprintf("invalid-limit-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	invalidTypeName := fmt.Sprintf("invalid-type-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	invalidName := fmt.Sprintf("Invalid_Name_%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	bothBranchesName := fmt.Sprintf("both-branches-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	missingNFSLimitName := fmt.Sprintf("missing-nfs-limit-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	missingBranchName := fmt.Sprintf("missing-branch-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	missingSCPProtocolName := fmt.Sprintf("missing-scp-protocol-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
	missingSCPUsernameName := fmt.Sprintf("missing-scp-username-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))

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
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}
	s8 := &helper.StepInfo{}
	s9 := &helper.StepInfo{}
	s10 := &helper.StepInfo{}
	s11 := &helper.StepInfo{}
	s12 := &helper.StepInfo{}
	s13 := &helper.StepInfo{}
	s14 := &helper.StepInfo{}
	s15 := &helper.StepInfo{}
	authConflictErr := regexp.MustCompile("Invalid Attribute Combination")
	missingAuthenticationErr := regexp.MustCompile("Missing SCP/SFTP authentication")
	passphraseRequiresSSHKeyErr := regexp.MustCompile("Invalid SCP/SFTP passphrase configuration")
	hostKeyConflictErr := regexp.MustCompile("Invalid host-key configuration")
	branchSelectionErr := regexp.MustCompile("Configure exactly one of `nfs` or `scp_sftp`")
	nfsLimitRequiredErr := regexp.MustCompile("(?i)(missing configuration for required attribute|required.*limit|limit.*required)")
	scpSftpRequiredAttributeErr := regexp.MustCompile("(?is)(missing configuration for required attribute|required.*(protocol|username)|(protocol|username).*required)")
	invalidAttributeErr := regexp.MustCompile("Invalid Attribute Value")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject password and SSH key configured together")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     passwordSSHKeyConflictName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "scp",
							"port":                       22,
							"username":                   "testuser",
							"password":                   "test-password",
							"ssh_key":                    "test-private-key",
							"ignore_host_key_validation": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject password and SSH-key passphrase configured together")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     passwordPassphraseConflictName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "sftp",
							"port":                       22,
							"username":                   "testuser",
							"password":                   "test-password",
							"passphrase":                 "test-passphrase",
							"ignore_host_key_validation": true,
						},
					})

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				PlanOnly:    true,
				ExpectError: authConflictErr,
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an SCP location without an authentication method")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     missingAuthenticationName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol": "scp",
							"port":     22,
							"username": "testuser",
						},
					})

					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig:   func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				PlanOnly:    true,
				ExpectError: missingAuthenticationErr,
			},
			{
				Config: func() string {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject a passphrase without SSH-key authentication")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     passphraseWithoutSSHKeyName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":   "sftp",
							"port":       22,
							"username":   "testuser",
							"passphrase": "test-passphrase",
						},
					})

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return s4.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s4.Index, s4.Name, s4.Cfg) },
				PlanOnly:    true,
				ExpectError: passphraseRequiresSSHKeyErr,
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject accept_host_key and ignore_host_key_validation configured together")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     hostKeyConflictName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol":                   "sftp",
							"port":                       22,
							"username":                   "testuser",
							"ssh_key":                    "test-private-key",
							"ignore_host_key_validation": true,
							"accept_host_key":            true,
						},
					})

					helper.GetTFConfigWithSingleResource(s5.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s5.Cfg = *tfConfig
					return s5.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s5.Index, s5.Name, s5.Cfg) },
				PlanOnly:    true,
				ExpectError: hostKeyConflictErr,
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject a remote storage port outside the valid range")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     invalidPortName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol": "scp",
							"port":     0,
							"username": "testuser",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s6.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return s6.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s6.Index, s6.Name, s6.Cfg) },
				PlanOnly:    true,
				ExpectError: invalidAttributeErr,
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an NFS alert threshold outside the valid range")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     invalidThresholdName,
						"hostname": host,
						"path":     "/mnt/tank/nfsstore",
						"nfs": map[string]interface{}{
							"limit":           "10MB",
							"alert_threshold": 101,
						},
					})

					helper.GetTFConfigWithSingleResource(s7.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return s7.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				PlanOnly:    true,
				ExpectError: invalidAttributeErr,
			},
			{
				Config: func() string {
					s8.Index = 8
					s8.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an invalid NFS storage limit")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     invalidLimitName,
						"hostname": host,
						"path":     "/mnt/tank/nfsstore",
						"nfs": map[string]interface{}{
							"limit": "10TB",
						},
					})

					helper.GetTFConfigWithSingleResource(s8.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s8.Cfg = *tfConfig
					return s8.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s8.Index, s8.Name, s8.Cfg) },
				PlanOnly:    true,
				ExpectError: invalidAttributeErr,
			},
			{
				Config: func() string {
					s9.Index = 9
					s9.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an unsupported remote storage location type")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     invalidTypeName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol": "ftp",
							"username": "testuser",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s9.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s9.Cfg = *tfConfig
					return s9.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s9.Index, s9.Name, s9.Cfg) },
				PlanOnly:    true,
				ExpectError: invalidAttributeErr,
			},
			{
				Config: func() string {
					s10.Index = 10
					s10.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an invalid remote storage location name")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     invalidName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol": "scp",
							"username": "testuser",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s10.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s10.Cfg = *tfConfig
					return s10.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s10.Index, s10.Name, s10.Cfg) },
				PlanOnly:    true,
				ExpectError: invalidAttributeErr,
			},
			{
				Config: func() string {
					s11.Index = 11
					s11.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject NFS and SCP/SFTP branches configured together")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     bothBranchesName,
						"hostname": host,
						"path":     "/tmp",
						"nfs": map[string]interface{}{
							"limit": "10GB",
						},
						"scp_sftp": map[string]interface{}{
							"protocol": "scp",
							"username": "testuser",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s11.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s11.Cfg = *tfConfig
					return s11.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s11.Index, s11.Name, s11.Cfg) },
				PlanOnly:    true,
				ExpectError: branchSelectionErr,
			},
			{
				Config: func() string {
					s12.Index = 12
					s12.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an NFS location without a storage limit")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     missingNFSLimitName,
						"hostname": host,
						"path":     "/mnt/tank/nfsstore",
						"nfs":      map[string]interface{}{},
					})

					helper.GetTFConfigWithSingleResource(s12.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s12.Cfg = *tfConfig
					return s12.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s12.Index, s12.Name, s12.Cfg) },
				PlanOnly:    true,
				ExpectError: nfsLimitRequiredErr,
			},
			{
				Config: func() string {
					s13.Index = 13
					s13.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject a remote storage location without a storage branch")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     missingBranchName,
						"hostname": host,
						"path":     "/tmp",
					})

					helper.GetTFConfigWithSingleResource(s13.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s13.Cfg = *tfConfig
					return s13.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s13.Index, s13.Name, s13.Cfg) },
				PlanOnly:    true,
				ExpectError: branchSelectionErr,
			},
			{
				Config: func() string {
					s14.Index = 14
					s14.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an SCP/SFTP branch without a protocol")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     missingSCPProtocolName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"username": "testuser",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s14.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s14.Cfg = *tfConfig
					return s14.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s14.Index, s14.Name, s14.Cfg) },
				PlanOnly:    true,
				ExpectError: scpSftpRequiredAttributeErr,
			},
			{
				Config: func() string {
					s15.Index = 15
					s15.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject an SCP/SFTP branch without a username")

					helper.GenerateRemoteStorageLocationObject(&rsc, map[string]interface{}{
						"name":     missingSCPUsernameName,
						"hostname": host,
						"path":     "/tmp",
						"scp_sftp": map[string]interface{}{
							"protocol": "scp",
							"password": "test-password",
						},
					})

					helper.GetTFConfigWithSingleResource(s15.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s15.Cfg = *tfConfig
					return s15.Cfg
				}(),
				PreConfig:   func() { helper.LogStep(t, s15.Index, s15.Name, s15.Cfg) },
				PlanOnly:    true,
				ExpectError: scpSftpRequiredAttributeErr,
			},
		},
	})
}

func newRemoteStorageLocationTestClient(t *testing.T) *nd.Client {
	t.Helper()

	cfg := helper.GetConfig("global")
	client, err := nd.NewClient(
		cfg.ND.URL,
		"/api/v1",
		cfg.ND.User,
		cfg.ND.Password,
		"",
		cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("failed to create ND client for remote-storage acceptance helper: %v", err)
	}

	return &client
}

func remoteStorageLocationAPIModel(rsc helper.NDFCRemoteStorageLocationTestData) resource_remote_storage_location.NDFCRemoteStorageLocationModel {
	model := resource_remote_storage_location.NDFCRemoteStorageLocationModel{
		Name:        rsc.Name,
		Description: rsc.Description,
		Hostname:    rsc.Hostname,
		Path:        rsc.Path,
	}

	if rsc.Nfs != nil {
		model.Nfs = resource_remote_storage_location.NDFCNfsValue{
			Port:           rsc.Nfs.Port,
			Limit:          rsc.Nfs.Limit,
			ReadWrite:      rsc.Nfs.ReadWrite,
			AlertThreshold: rsc.Nfs.AlertThreshold,
		}
		return model
	}

	model.ScpSftp = resource_remote_storage_location.NDFCScpSftpValue{
		Protocol:      rsc.ScpSftp.Protocol,
		Port:          rsc.ScpSftp.Port,
		AcceptHostKey: remoteStorageLocationAcceptHostKey(rsc),
		Authentication: resource_remote_storage_location.NDFCAuthenticationValue{
			Username:                rsc.ScpSftp.Username,
			Password:                rsc.ScpSftp.Password,
			SshKey:                  rsc.ScpSftp.SshKey,
			Passphrase:              rsc.ScpSftp.Passphrase,
			IgnoreHostKeyValidation: rsc.ScpSftp.IgnoreHostKeyValidation,
		},
	}
	return model
}

func remoteStorageLocationAcceptHostKey(rsc helper.NDFCRemoteStorageLocationTestData) bool {
	return rsc.ScpSftp != nil && rsc.ScpSftp.AcceptHostKey != nil && *rsc.ScpSftp.AcceptHostKey
}

func createRemoteStorageLocationOutsideTerraform(t *testing.T, rsc helper.NDFCRemoteStorageLocationTestData) {
	t.Helper()

	payload, err := json.Marshal(remoteStorageLocationAPIModel(rsc))
	if err != nil {
		t.Fatalf("failed to marshal remote-storage create payload for %q: %v", rsc.Name, err)
	}

	client := newRemoteStorageLocationTestClient(t)
	remoteStorageAPI := api.NewRemoteStorageLocationAPI(client, ndapi.DefaultFabric)
	remoteStorageAPI.AcceptHostKey = remoteStorageLocationAcceptHostKey(rsc)
	res, err := remoteStorageAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to create remote storage location %q outside Terraform: %v %s", rsc.Name, err, res.String())
	}
}

func updateRemoteStorageLocationOutsideTerraform(t *testing.T, rsc helper.NDFCRemoteStorageLocationTestData) {
	t.Helper()

	payload, err := json.Marshal(remoteStorageLocationAPIModel(rsc))
	if err != nil {
		t.Fatalf("failed to marshal remote-storage update payload for %q: %v", rsc.Name, err)
	}

	client := newRemoteStorageLocationTestClient(t)
	remoteStorageAPI := api.NewRemoteStorageLocationAPI(client, ndapi.DefaultFabric)
	remoteStorageAPI.Name = rsc.Name
	remoteStorageAPI.AcceptHostKey = remoteStorageLocationAcceptHostKey(rsc)
	res, err := remoteStorageAPI.Put(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to update remote storage location %q outside Terraform: %v %s", rsc.Name, err, res.String())
	}
}

func deleteRemoteStorageLocationOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	client := newRemoteStorageLocationTestClient(t)
	remoteStorageAPI := api.NewRemoteStorageLocationAPI(client, ndapi.DefaultFabric)
	remoteStorageAPI.Name = name
	res, err := remoteStorageAPI.Delete(nil)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		t.Fatalf("failed to delete remote storage location %q outside Terraform: %v %s", name, err, res.String())
	}

	var lastResponse []byte
	pollErr := utils.PollUntil(context.Background(), remoteStorageLocationTestDeletePollInterval, remoteStorageLocationTestDeletePollTimeout, func(context.Context) (bool, error) {
		respData, getErr := remoteStorageAPI.Get()
		lastResponse = respData
		if getErr == nil {
			if respData == nil {
				return true, nil
			}
			return false, nil
		}
		if strings.Contains(getErr.Error(), "StatusCode 404") {
			return true, nil
		}
		return false, fmt.Errorf("failed to verify deletion of remote storage location %q: %w %s", name, getErr, string(respData))
	})
	if pollErr == nil {
		return
	}
	if errors.Is(pollErr, utils.ErrPollTimeout) {
		t.Fatalf("timed out after %s waiting for remote storage location %q deletion to complete; last response: %s", remoteStorageLocationTestDeletePollTimeout, name, string(lastResponse))
	}
	t.Fatalf("failed while waiting for remote storage location %q deletion to complete: %v", name, pollErr)
}

func assertRemoteStorageLocationAbsentOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	client := newRemoteStorageLocationTestClient(t)
	remoteStorageAPI := api.NewRemoteStorageLocationAPI(client, ndapi.DefaultFabric)
	remoteStorageAPI.Name = name
	respData, err := remoteStorageAPI.Get()
	if err == nil {
		t.Fatalf("remote storage location %q still exists in Nexus Dashboard: %s", name, string(respData))
	}
	if !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to verify remote storage location %q is absent from Nexus Dashboard: %v %s", name, err, string(respData))
	}
}
