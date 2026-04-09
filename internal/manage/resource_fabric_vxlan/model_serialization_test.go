// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_vxlan

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNDFCFabricVxlanModelSerialization(t *testing.T) {
	// Find the root directory of the project
	rootDir := findProjectRoot(t)

	// Path to the input JSON file
	inputFile := filepath.Join(rootDir, "fabric1.json")

	// Read the input JSON file
	inputData, err := os.ReadFile(inputFile)
	require.NoError(t, err, "Failed to read input file: %v", err)

	// Unmarshal the JSON into the model
	var model NDFCFabricVxlanModel
	err = json.Unmarshal(inputData, &model)
	require.NoError(t, err, "Failed to unmarshal JSON into NDFCFabricVxlanModel: %v", err)

	// Marshal the model back to JSON
	outputData, err := json.Marshal(&model)
	require.NoError(t, err, "Failed to marshal NDFCFabricVxlanModel back to JSON: %v", err)

	// Create a temporary output file
	outputFile := filepath.Join(os.TempDir(), "fabric1_output.json")
	err = os.WriteFile(outputFile, outputData, 0644)
	require.NoError(t, err, "Failed to write output file: %v", err)

	// For better debugging in case of failures, also write a pretty-printed version
	var prettyInput, prettyOutput map[string]interface{}
	err = json.Unmarshal(inputData, &prettyInput)
	require.NoError(t, err, "Failed to unmarshal input for pretty printing")

	err = json.Unmarshal(outputData, &prettyOutput)
	require.NoError(t, err, "Failed to unmarshal output for pretty printing")

	prettyInputData, err := json.MarshalIndent(prettyInput, "", "  ")
	require.NoError(t, err, "Failed to marshal pretty input")

	prettyOutputData, err := json.MarshalIndent(prettyOutput, "", "  ")
	require.NoError(t, err, "Failed to marshal pretty output")

	prettyInputFile := filepath.Join(os.TempDir(), "fabric1_input_pretty.json")
	prettyOutputFile := filepath.Join(os.TempDir(), "fabric1_output_pretty.json")

	err = os.WriteFile(prettyInputFile, prettyInputData, 0644)
	require.NoError(t, err, "Failed to write pretty input file")

	err = os.WriteFile(prettyOutputFile, prettyOutputData, 0644)
	require.NoError(t, err, "Failed to write pretty output file")

	t.Logf("Wrote pretty-printed input to: %s", prettyInputFile)
	t.Logf("Wrote pretty-printed output to: %s", prettyOutputFile)

	// Compare the input and output JSON (ignoring formatting differences)
	var inputJson, outputJson interface{}
	err = json.Unmarshal(inputData, &inputJson)
	require.NoError(t, err, "Failed to parse input JSON for comparison")

	err = json.Unmarshal(outputData, &outputJson)
	require.NoError(t, err, "Failed to parse output JSON for comparison")

	// Compare the JSON objects
	// We use assert instead of require to see all failures
	assert.Equal(t, inputJson, outputJson, "The input and output JSON should be equivalent")

	// Report any keys that exist in input but not in output
	inputMap, ok := inputJson.(map[string]interface{})
	require.True(t, ok, "Input JSON should be a map")

	outputMap, ok := outputJson.(map[string]interface{})
	require.True(t, ok, "Output JSON should be a map")

	checkMissingKeys(t, inputMap, outputMap, "")
}

// findProjectRoot attempts to find the project root by looking for typical markers
func findProjectRoot(t *testing.T) string {
	// Start with the current working directory
	dir, err := os.Getwd()
	require.NoError(t, err, "Failed to get current working directory")

	// Go up the directory tree looking for the project root
	for {
		// Check if we've reached the root of the filesystem
		if dir == "/" || dir == "." {
			t.Fatal("Could not find project root")
		}

		// Check for typical markers of a project root
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		// Go up one directory
		dir = filepath.Dir(dir)
	}
}

// checkMissingKeys recursively compares two maps to find keys that exist in the input but not in the output
func checkMissingKeys(t *testing.T, input, output map[string]interface{}, path string) {
	for key, inputValue := range input {
		currentPath := path
		if currentPath == "" {
			currentPath = key
		} else {
			currentPath = path + "." + key
		}

		outputValue, exists := output[key]
		if !exists {
			t.Errorf("Key missing in output: %s", currentPath)
			continue
		}

		// If this is a nested map, recurse
		if inputMap, ok := inputValue.(map[string]interface{}); ok {
			if outputMap, ok := outputValue.(map[string]interface{}); ok {
				checkMissingKeys(t, inputMap, outputMap, currentPath)
			} else {
				t.Errorf("Expected map at %s but got %T", currentPath, outputValue)
			}
		} else if inputSlice, ok := inputValue.([]interface{}); ok {
			if outputSlice, ok := outputValue.([]interface{}); ok {
				// For arrays, just check length
				if len(inputSlice) > len(outputSlice) {
					t.Errorf("Array at %s has fewer elements in output (%d) than in input (%d)",
						currentPath, len(outputSlice), len(inputSlice))
				}
			} else {
				t.Errorf("Expected array at %s but got %T", currentPath, outputValue)
			}
		}
	}
}
