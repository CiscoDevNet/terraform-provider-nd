// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

// Run from the repository root after tfplugindocs generate:
// go run ./generator/doc_category.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const definitionsPath = "./generator/defs"

type docConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
}

func main() {
	docCategories, ok := loadDocCategories()
	if !ok {
		return
	}

	content, _ := os.ReadFile(filepath.Join(definitionsPath, "defs.yaml"))
	var manifest struct {
		Files []string `yaml:"files"`
	}
	_ = yaml.Unmarshal(content, &manifest)

	docDirectories := map[string]string{
		"resource":   "resources",
		"datasource": "data-sources",
	}
	for _, filename := range manifest.Files {
		content, err := os.ReadFile(filepath.Join(definitionsPath, filename))
		if err != nil {
			continue
		}
		var definitions map[string]docConfig
		if yaml.Unmarshal(content, &definitions) != nil {
			continue
		}

		for kind, directory := range docDirectories {
			definition := definitions[kind]
			if definition.Name == "" || definition.DocCategory == "" {
				continue
			}
			if _, ok := docCategories[definition.DocCategory]; !ok {
				continue
			}
			updateDoc(filepath.Join("docs", directory, definition.Name+".md"), definition.DocCategory)
		}
	}
}

func loadDocCategories() (map[string]struct{}, bool) {
	content, err := os.ReadFile(filepath.Join(definitionsPath, "doc_categories.yaml"))
	if err != nil {
		return nil, false
	}

	var categories []string
	if yaml.Unmarshal(content, &categories) != nil {
		return nil, false
	}

	docCategories := make(map[string]struct{}, len(categories))
	for _, category := range categories {
		docCategories[category] = struct{}{}
	}
	return docCategories, true
}

func updateDoc(filename, subcategory string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	lines := bytes.SplitAfter(content, []byte("\n"))
	if !bytes.Equal(bytes.TrimRight(lines[0], " \t\r\n"), []byte("---")) {
		return
	}
	subcategoryLine := []byte(fmt.Sprintf("subcategory: %q", subcategory))
	for i := 1; i < len(lines); i++ {
		newline := []byte("\n")
		if bytes.HasSuffix(lines[i], []byte("\r\n")) {
			newline = []byte("\r\n")
		}

		if bytes.HasPrefix(lines[i], []byte("subcategory:")) {
			lines[i] = append(subcategoryLine, newline...)
			_ = os.WriteFile(filename, bytes.Join(lines, nil), 0o644)
			return
		}
		if bytes.Equal(bytes.TrimRight(lines[i], " \t\r\n"), []byte("---")) {
			lines = append(lines, nil)
			copy(lines[i+1:], lines[i:])
			lines[i] = append(subcategoryLine, newline...)
			_ = os.WriteFile(filename, bytes.Join(lines, nil), 0o644)
			return
		}
	}
}
