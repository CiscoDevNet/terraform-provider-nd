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
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const definitionsPath = "./generator/defs"

type docConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
	Subcategory string `yaml:"subcategory"`
}

func main() {
	if err := updateDocCategories(); err != nil {
		log.Fatal(err)
	}
}

func updateDocCategories() error {
	categories, err := loadDocCategories()
	if err != nil {
		return err
	}

	// Visit generated documents once, using definition metadata as a lookup.
	return filepath.WalkDir("docs", func(filename string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(filename) != ".md" {
			return nil
		}
		category, ok := categories[filename]
		if !ok {
			return nil
		}
		return updateDoc(filename, category)
	})
}

func loadDocCategories() (map[string]string, error) {
	var manifest struct {
		Files []string `yaml:"files"`
	}
	if err := readYAML(filepath.Join(definitionsPath, "defs.yaml"), &manifest); err != nil {
		return nil, err
	}

	directories := map[string]string{
		"resource":   "resources",
		"datasource": "data-sources",
		"action":     "actions",
	}
	categories := make(map[string]string)
	// Use the same active definitions as the provider generator.
	for _, filename := range manifest.Files {
		var config map[string]docConfig
		if err := readYAML(filepath.Join(definitionsPath, filename), &config); err != nil {
			return nil, err
		}

		for kind, directory := range directories {
			metadata := config[kind]
			category := strings.TrimSpace(metadata.Subcategory)
			if category == "" {
				category = strings.TrimSpace(metadata.DocCategory)
			}
			if category == "" {
				continue
			}

			name := strings.TrimSpace(metadata.Name)
			if name == "" {
				return nil, fmt.Errorf("%s: %s documentation category requires a name", filename, kind)
			}
			docPath := filepath.Join("docs", directory, name+".md")
			if previous, exists := categories[docPath]; exists && previous != category {
				return nil, fmt.Errorf("%s: conflicting categories for %s: %q and %q", filename, docPath, previous, category)
			}
			categories[docPath] = category
		}
	}
	return categories, nil
}

func readYAML(filename string, config any) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read %s: %w", filename, err)
	}
	if err := yaml.Unmarshal(content, config); err != nil {
		return fmt.Errorf("parse %s: %w", filename, err)
	}
	return nil
}

func updateDoc(filename, category string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read %s: %w", filename, err)
	}

	lines := bytes.SplitAfter(content, []byte("\n"))
	if string(bytes.TrimSpace(lines[0])) != "---" {
		return nil
	}
	headerEnd := -1
	for i := 1; i < len(lines); i++ {
		if string(bytes.TrimSpace(lines[i])) == "---" {
			headerEnd = i
			break
		}
	}
	if headerEnd < 0 {
		return fmt.Errorf("%s: missing closing front matter delimiter", filename)
	}

	var header yaml.Node
	if err := yaml.Unmarshal(bytes.Join(lines[1:headerEnd], nil), &header); err != nil {
		return fmt.Errorf("parse front matter in %s: %w", filename, err)
	}
	if len(header.Content) == 0 || header.Content[0].Kind != yaml.MappingNode {
		return fmt.Errorf("%s: front matter must be a YAML mapping", filename)
	}

	fields := header.Content[0].Content
	for i := 0; i < len(fields); i += 2 {
		key, value := fields[i], fields[i+1]
		if key.Value != "subcategory" {
			continue
		}

		value.Value = category
		value.Style = yaml.DoubleQuotedStyle

		var encoded bytes.Buffer
		encoder := yaml.NewEncoder(&encoded)
		encoder.SetIndent(2)
		if err := encoder.Encode(&header); err != nil {
			return fmt.Errorf("encode front matter in %s: %w", filename, err)
		}
		if err := encoder.Close(); err != nil {
			return fmt.Errorf("close front matter encoder for %s: %w", filename, err)
		}
		updatedHeader := encoded.Bytes()
		if bytes.HasSuffix(lines[0], []byte("\r\n")) {
			updatedHeader = bytes.ReplaceAll(updatedHeader, []byte("\n"), []byte("\r\n"))
		}

		// Keep both delimiters and the Markdown body byte-for-byte.
		var updated bytes.Buffer
		updated.Write(lines[0])
		updated.Write(updatedHeader)
		updated.Write(bytes.Join(lines[headerEnd:], nil))
		if err := writeDocAtomically(filename, updated.Bytes()); err != nil {
			return fmt.Errorf("write %s: %w", filename, err)
		}
		log.Printf("Updated %s: %s", filename, category)
		return nil
	}
	return nil
}

func writeDocAtomically(filename string, content []byte) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	// Keep the temporary file on the same filesystem for atomic replacement.
	temporary, err := os.CreateTemp(filepath.Dir(filename), ".doc-category-*")
	if err != nil {
		return err
	}
	defer os.Remove(temporary.Name())
	defer temporary.Close()

	if _, err := temporary.Write(content); err != nil {
		return err
	}
	if err := temporary.Chmod(info.Mode().Perm()); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	return os.Rename(temporary.Name(), filename)
}
