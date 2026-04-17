// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	"terraform-provider-nd/internal/manage/resource_fabric_vxlan"
	"terraform-provider-nd/internal/manage/resource_inventory_switch"
)

var testOutputDir string

func init() {
	ts := time.Now().Format("2006_01_02_15-04-05")
	testOutputDir = filepath.Join(os.TempDir(), fmt.Sprintf("tftest_%s", ts))
}

// GetTFConfigWithSingleResource generates a complete Terraform HCL config string
// from Go model objects using .gotmpl templates.
//
// Parameters:
//   - tt: test step name (used for snapshot filename)
//   - cfg: provider config map with keys: User, Password, Host, Insecure, RscType, RscName
//   - rscs: slice of resource model pointers
//   - out: output pointer that receives the generated HCL string
func GetTFConfigWithSingleResource(tt string, cfg map[string]string, rscs []interface{}, out **string) {
	// Find the directory containing templates (same directory as this file)
	_, filename, _, _ := runtime.Caller(0)
	tmplDir := filepath.Dir(filename)

	// Load all .gotmpl files
	tmplFiles, err := filepath.Glob(filepath.Join(tmplDir, "*.gotmpl"))
	if err != nil || len(tmplFiles) == 0 {
		panic(fmt.Sprintf("Failed to find .gotmpl files in %s: %v", tmplDir, err))
	}

	// Custom template functions
	functions := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"deref_bool": func(a *bool) bool {
			if a == nil {
				return false
			}
			return *a
		},
		"deref_int64": func(a *int64) int64 {
			if a == nil {
				return 0
			}
			return *a
		},
		"trimSuffix": func(s, suffix string) string {
			return strings.TrimSuffix(s, suffix)
		},
	}

	// Parse all templates
	t, err := template.New("").Funcs(functions).ParseFiles(tmplFiles...)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse templates: %v", err))
	}

	var output bytes.Buffer

	// Note: We do NOT render the HEADER (required_providers) block because
	// ProtoV6ProviderFactories injects the provider directly. Including
	// required_providers causes "Inconsistent dependency lock file" errors.

	// Execute ND provider template
	providerArgs := map[string]string{
		"User":     cfg["User"],
		"Password": cfg["Password"],
		"Host":     cfg["Host"],
		"Insecure": cfg["Insecure"],
	}
	err = t.ExecuteTemplate(&output, "ND", providerArgs)
	if err != nil {
		panic(fmt.Sprintf("Failed to execute ND template: %v", err))
	}

	// Split resource names by comma
	rscNames := strings.Split(cfg["RscName"], ",")

	// Split DependsOn by semicolon to support per-resource depends_on.
	// If a single value is provided, all inventory switch resources share it.
	// If multiple values are provided (semicolon-separated), each inventory
	// switch resource gets the corresponding entry by index.
	var dependsOnList []string
	if dep, ok := cfg["DependsOn"]; ok && dep != "" {
		dependsOnList = strings.Split(dep, ";")
		for i := range dependsOnList {
			dependsOnList[i] = strings.TrimSpace(dependsOnList[i])
		}
	}

	// Execute resource-specific templates
	invSwitchIdx := 0
	for i, rsc := range rscs {
		args := make(map[string]interface{})
		rscName := ""
		if i < len(rscNames) {
			rscName = strings.TrimSpace(rscNames[i])
		}

		switch v := rsc.(type) {
		case *resource_fabric_vxlan.NDFCFabricVxlanModel:
			args["FabricVxlan"] = v
			args["RscName"] = rscName
			err = t.ExecuteTemplate(&output, "ND_FABRIC_VXLAN_RSC", args)
			if err != nil {
				panic(fmt.Sprintf("Failed to execute ND_FABRIC_VXLAN_RSC template: %v", err))
			}

		case *resource_inventory_switch.NDFCInventorySwitchModel:
			args["InventorySwitch"] = v
			args["RscName"] = rscName
			// Assign per-resource depends_on
			if len(dependsOnList) > 0 {
				idx := invSwitchIdx
				if idx >= len(dependsOnList) {
					idx = len(dependsOnList) - 1
				}
				if dependsOnList[idx] != "" {
					args["DependsOn"] = dependsOnList[idx]
				}
			}
			invSwitchIdx++
			err = t.ExecuteTemplate(&output, "ND_INVENTORY_SWITCH_RSC", args)
			if err != nil {
				panic(fmt.Sprintf("Failed to execute ND_INVENTORY_SWITCH_RSC template: %v", err))
			}

		default:
			panic(fmt.Sprintf("Unknown resource type: %T", rsc))
		}
	}

	result := output.String()
	*out = &result

	// Write snapshot to temp directory for debugging
	writeSnapshot(tt, result)
}

// writeSnapshot writes the generated HCL to a temp file for post-test debugging
func writeSnapshot(testName, content string) {
	if err := os.MkdirAll(testOutputDir, 0755); err != nil {
		log.Printf("Warning: failed to create snapshot dir %s: %v", testOutputDir, err)
		return
	}

	// Sanitize test name for filename
	safeName := strings.ReplaceAll(testName, "/", "_")
	safeName = strings.ReplaceAll(safeName, " ", "_")
	filePath := filepath.Join(testOutputDir, safeName+".tf")

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		log.Printf("Warning: failed to write snapshot %s: %v", filePath, err)
		return
	}

	log.Printf("Snapshot written to: %s", filePath)
}
