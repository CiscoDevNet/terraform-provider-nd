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
	"testing"
	"time"

	"terraform-provider-nd/internal/infra/resource_backup"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccBackupResource exercises the create + import lifecycle of a single
// nd_backup resource. All nd_backup attributes use the RequiresReplace plan
// modifier and the resource does not support in-place updates, so
// create + import is the full meaningful coverage path.
//
// Notes:
//   - `encryption_key` is not returned by the API, so it is excluded from the
//     ImportStateVerify diff comparison.
//   - `telemetry_data` is set explicitly in the config because the provider
//     preserves whatever value was in the plan; leaving it unset surfaces as
//     "unknown value for telemetry_data after apply".
//   - A standalone CheckDestroy is intentionally NOT used: the controller
//     reports HTTP 403 "Backup cannot be deleted while in progress" if the
//     backup job has not finished, which would make the post-test destroy
//     flaky. The framework still attempts destroy for cleanup at end of test.
func TestAccBackupResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	bkCfg := cfg.ND.Backup

	x := &map[string]string{
		"RscType":  "nd_backup",
		"RscName":  "backup_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	backupRsc := new(resource_backup.NDFCBackupModel)

	name := bkCfg.Name
	if name == "" {
		name = "tf-backup-acc"
	}
	encryptionKey := bkCfg.EncryptionKey
	if encryptionKey == "" {
		encryptionKey = "backupKey123"
	}
	backupType := bkCfg.Type
	if backupType == "" {
		backupType = "configOnly"
	}

	s1 := &stepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the backup using testbed values.
			{
				Config: func() string {
					*stepCount++
					s1.name = fmt.Sprintf("%s_%d_create_backup", t.Name(), *stepCount)

					helper.GenerateBackupObject(&backupRsc,
						name,
						encryptionKey,
						map[string]interface{}{
							"type":           backupType,
							"destination":    bkCfg.Destination,
							"telemetry_data": bkCfg.TelemetryData,
						},
					)

					helper.GetTFConfigWithSingleResource(s1.name, *x,
						[]interface{}{backupRsc}, &tfConfig)

					s1.cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, s1.name, s1.cfg) },
				Check: resource.ComposeTestCheckFunc(
					BackupModelHelperStateCheck(
						"nd_backup.backup_test",
						*backupRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: ImportState verification. `encryption_key` is not
			// returned by the API and is excluded from the diff comparison.
			{
				PreConfig: func() {
					t.Logf("===== STEP 2: %s_2_import_backup =====", t.Name())
				},
				ResourceName:                         "nd_backup.backup_test",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        name,
				ImportStateVerifyIdentifierAttribute: "name",
				ImportStateVerifyIgnore: []string{
					"encryption_key",
					"telemetry_data",
				},
			},
			// Step 3: Refresh-only wait step. The controller returns HTTP
			// 403 "Backup cannot be deleted while in progress" until the
			// backup job finishes. Sleep before the framework's post-test
			// destroy runs so the cleanup DELETE succeeds.
			{
				PreConfig: func() {
					t.Logf("===== STEP 3: %s_3_wait_for_backup_completion =====", t.Name())
					t.Logf("Sleeping 5 minutes to allow backup job to finish before destroy")
					time.Sleep(5 * time.Minute)
				},
				// Config:       s1.cfg,
				RefreshState: true,
			},
		},
	})
}
