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

	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_backup"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	nd "github.com/netascode/go-nd"
)

type backupCreateCase struct {
	stepName             string
	name                 string
	backupType           string
	encryptionKey        string
	destination          string
	includeDestination   bool
	telemetryData        bool
	includeTelemetryData bool
}

func deleteBackupOutsideTerraform(t *testing.T, name string) {
	t.Helper()

	cfg := helper.GetConfig("global")
	client, err := nd.NewClient(
		cfg.ND.URL, "/api/v1",
		cfg.ND.User, cfg.ND.Password,
		"", cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("failed to create ND client for out-of-band backup delete: %v", err)
	}

	backupAPI := api.NewBackupAPI(&client)
	backupAPI.Name = name

	res, err := backupAPI.Delete(nil)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		t.Fatalf("failed to delete backup %q outside Terraform: %v %v", name, err, res)
	}
}

// TestAccBackupResource exercises the supported create/destroy lifecycle for
// the main nd_backup combinations. Import is intentionally not covered because
// nd_backup import is not supported.
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

	name := bkCfg.Name
	if name == "" {
		name = "tf-backup-acc"
	}
	encryptionKey := bkCfg.EncryptionKey
	if encryptionKey == "" {
		encryptionKey = "backupKey123"
	}

	scpConfigOnlyBackup := backupCreateCase{
		stepName:           "create_scp_config_only_backup",
		name:               fmt.Sprintf("%s-scp", name),
		backupType:         "configOnly",
		encryptionKey:      encryptionKey,
		destination:        "ubuntu",
		includeDestination: true,
	}

	testCases := []backupCreateCase{
		{
			stepName:      "create_nd_local_config_only_backup",
			name:          name,
			backupType:    "configOnly",
			encryptionKey: encryptionKey,
		},
		{
			stepName:             "create_nas_full_telemetry_backup",
			name:                 fmt.Sprintf("%s-nas", name),
			backupType:           "full",
			encryptionKey:        encryptionKey,
			destination:          "nas",
			includeDestination:   true,
			telemetryData:        true,
			includeTelemetryData: true,
		},
		scpConfigOnlyBackup,
	}

	steps := make([]resource.TestStep, 0, len(testCases)+1)
	for i, tc := range testCases {
		tc := tc
		stepNumber := i + 1
		stepInfo := &helper.StepInfo{}
		backupRsc := new(resource_backup.NDFCBackupModel)

		steps = append(steps, resource.TestStep{
			Config: func() string {
				stepInfo.Name = fmt.Sprintf("%s_%d_%s", t.Name(), stepNumber, tc.stepName)

				overrides := map[string]interface{}{
					"type": tc.backupType,
				}
				if tc.includeDestination {
					overrides["destination"] = tc.destination
				}
				if tc.includeTelemetryData {
					overrides["telemetry_data"] = tc.telemetryData
				}

				helper.GenerateBackupObject(&backupRsc, tc.name, tc.encryptionKey, overrides)
				helper.GetTFConfigWithSingleResource(stepInfo.Name, *x,
					[]interface{}{backupRsc}, &tfConfig)

				stepInfo.Cfg = *tfConfig
				return *tfConfig
			}(),
			PreConfig: func() { helper.LogStep(t, stepNumber, stepInfo.Name, stepInfo.Cfg) },
			Check: resource.ComposeTestCheckFunc(
				append(
					BackupModelHelperStateCheck(
						"nd_backup.backup_test",
						*backupRsc,
						path.Empty(),
					),
					backupAdditionalStateChecks("nd_backup.backup_test", tc)...,
				)...,
			),
		})
	}
	steps = append(steps, resource.TestStep{
		PreConfig: func() {
			t.Logf("===== STEP %d: %s_%d_delete_backup_outside_terraform =====",
				len(testCases)+1, t.Name(), len(testCases)+1)
			deleteBackupOutsideTerraform(t, scpConfigOnlyBackup.name)
		},
		RefreshState:       true,
		ExpectNonEmptyPlan: true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func TestAccBackupResourceInvalidTelemetryData(t *testing.T) {
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
	backupRsc := new(resource_backup.NDFCBackupModel)

	name := bkCfg.Name
	if name == "" {
		name = "tf-backup-acc"
	}
	encryptionKey := bkCfg.EncryptionKey
	if encryptionKey == "" {
		encryptionKey = "backupKey123"
	}

	tc := backupCreateCase{
		stepName:             "create_invalid_scp_config_only_telemetry_backup",
		name:                 fmt.Sprintf("%s-invalid", name),
		backupType:           "configOnly",
		encryptionKey:        encryptionKey,
		destination:          "ubuntu",
		includeDestination:   true,
		telemetryData:        true,
		includeTelemetryData: true,
	}
	stepInfo := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					stepInfo.Name = fmt.Sprintf("%s_1_%s", t.Name(), tc.stepName)

					helper.GenerateBackupObject(&backupRsc,
						tc.name,
						tc.encryptionKey,
						map[string]interface{}{
							"type":           tc.backupType,
							"destination":    tc.destination,
							"telemetry_data": tc.telemetryData,
						},
					)
					helper.GetTFConfigWithSingleResource(stepInfo.Name, *x,
						[]interface{}{backupRsc}, &tfConfig)

					stepInfo.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, 1, stepInfo.Name, stepInfo.Cfg) },
				ExpectError: regexp.MustCompile(
					`Telemetry data collection is allowed only for full\s+backups and NAS destinations`,
				),
			},
		},
	})
}

func backupAdditionalStateChecks(resourceName string, tc backupCreateCase) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "id", tc.name),
	}

	if !tc.includeDestination {
		checks = append(checks, resource.TestCheckNoResourceAttr(resourceName, "destination"))
	}

	if !tc.includeTelemetryData {
		checks = append(checks, resource.TestCheckNoResourceAttr(resourceName, "telemetry_data"))
	}

	return checks
}
