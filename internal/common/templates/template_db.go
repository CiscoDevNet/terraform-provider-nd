// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Package templates provides a singleton, lazily-populated cache of ND
// configuration-template metadata fetched from the Template Library API
// (GET /manage/configTemplates/{name}/parameters). Any manage resource
// can call templates.GetInstance() to share a single, lifetime-scoped
// cache across the provider process.
package templates

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	nd "github.com/netascode/go-nd"
)

// TemplateParam mirrors one entry in the parameters array returned by
// GET /configTemplates/{name}/parameters.
type TemplateParam struct {
	Name           string            `json:"name"`
	ParamType      string            `json:"parameterType"`
	Description    string            `json:"description,omitempty"`
	DefaultValue   string            `json:"defaultValue,omitempty"`
	Optional       bool              `json:"optional"`
	Annotations    map[string]string `json:"annotations,omitempty"`
	MetaProperties map[string]string `json:"metaProperties,omitempty"`
}

// TemplateInfo holds the cached metadata for a single template.
type TemplateInfo struct {
	TemplateName string
	Params       []TemplateParam
	// Derived indexes — built once at load time.
	paramByName map[string]*TemplateParam // camelCase name → param
	tfNames     map[string]bool           // snake_case tf_name set
}

// HasParam returns true if the template declares a parameter with the
// given camelCase model name.
func (t *TemplateInfo) HasParam(modelName string) bool {
	_, ok := t.paramByName[modelName]
	return ok
}

// ParamType returns the parameterType string (e.g. "string", "integer",
// "boolean") for the named parameter, or "" if not found.
func (t *TemplateInfo) ParamType(modelName string) string {
	if p, ok := t.paramByName[modelName]; ok {
		return p.ParamType
	}
	return ""
}

// AllowedTFNames returns the set of snake_case Terraform attribute names
// derived from this template's parameter list. The set is stable and
// safe for concurrent reads.
func (t *TemplateInfo) AllowedTFNames() map[string]bool {
	return t.tfNames
}

// TemplateDB is a thread-safe, lazily-populated cache of ND template
// parameter metadata. Templates are fetched on first access by name and
// cached for the provider process lifetime. The singleton is obtained
// via GetInstance().
type TemplateDB struct {
	mu        sync.RWMutex
	templates map[string]*TemplateInfo // templateName → info
	client    *nd.Client
}

var (
	instance *TemplateDB
	once     sync.Once
)

// InitInstance initialises the singleton TemplateDB with the given API
// client. It is safe to call multiple times; only the first call takes
// effect (the client is bound once).
func InitInstance(client *nd.Client) {
	once.Do(func() {
		instance = &TemplateDB{
			templates: make(map[string]*TemplateInfo),
			client:    client,
		}
	})
}

// GetInstance returns the singleton TemplateDB. Panics if InitInstance
// has not been called.
func GetInstance() *TemplateDB {
	if instance == nil {
		panic("templates.GetInstance called before InitInstance")
	}
	return instance
}

// ResetForTest tears down the singleton so unit tests can re-initialise
// with a fresh client. Production code must not call this.
func ResetForTest() {
	instance = nil
	once = sync.Once{}
}

// GetTemplate returns the cached TemplateInfo for the given template
// name, fetching it from ND on first access. Returns nil if the
// template cannot be loaded.
func (db *TemplateDB) GetTemplate(ctx context.Context, templateName string) *TemplateInfo {
	// Fast path: read lock
	db.mu.RLock()
	if info, ok := db.templates[templateName]; ok {
		db.mu.RUnlock()
		return info
	}
	db.mu.RUnlock()

	// Slow path: write lock — fetch and cache
	db.mu.Lock()
	defer db.mu.Unlock()

	// Double-check under write lock
	if info, ok := db.templates[templateName]; ok {
		return info
	}

	info, err := db.fetchTemplate(ctx, templateName)
	if err != nil {
		tflog.Warn(ctx, "TemplateDB: failed to load template", map[string]interface{}{
			"template": templateName,
			"error":    err.Error(),
		})
		return nil
	}
	db.templates[templateName] = info

	tflog.Debug(ctx, "TemplateDB: loaded template", map[string]interface{}{
		"template": templateName,
		"params":   len(info.Params),
	})
	return info
}

// GetTemplateForPolicy resolves a policyType to its templateName, then
// returns the template info. Returns nil when the policyType is unknown
// or the template cannot be loaded.
func (db *TemplateDB) GetTemplateForPolicy(ctx context.Context, policyType string) *TemplateInfo {
	tmpl, ok := PolicyToTemplate[policyType]
	if !ok {
		return nil
	}
	return db.GetTemplate(ctx, tmpl)
}

// GetParamTypeForPolicy is a convenience method that resolves a
// policyType → templateName, loads the template, and returns the
// parameterType for the given camelCase field name.
func (db *TemplateDB) GetParamTypeForPolicy(ctx context.Context, policyType, fieldName string) string {
	info := db.GetTemplateForPolicy(ctx, policyType)
	if info == nil {
		return ""
	}
	return info.ParamType(fieldName)
}

// AllowedTFNamesForPolicy resolves a policyType to its template and
// returns the set of allowed Terraform attribute names (snake_case).
// Returns nil when the policy is unknown or the template cannot be
// loaded.
func (db *TemplateDB) AllowedTFNamesForPolicy(ctx context.Context, policyType string) map[string]bool {
	info := db.GetTemplateForPolicy(ctx, policyType)
	if info == nil {
		return nil
	}
	return info.AllowedTFNames()
}

// --- internal ---------------------------------------------------------------

// templateParamsURL is the manage API path for template parameters.
const templateParamsURL = "/manage/configTemplates/%s/parameters"

func (db *TemplateDB) fetchTemplate(ctx context.Context, templateName string) (*TemplateInfo, error) {
	if db.client == nil {
		return nil, fmt.Errorf("TemplateDB: no API client configured")
	}

	url := fmt.Sprintf(templateParamsURL, templateName)
	raw, err := db.client.GetRawJson(url)
	if err != nil {
		return nil, fmt.Errorf("TemplateDB: GET %s: %w", url, err)
	}

	var params []TemplateParam
	if err := json.Unmarshal(raw, &params); err != nil {
		return nil, fmt.Errorf("TemplateDB: unmarshal %s: %w", templateName, err)
	}

	info := &TemplateInfo{
		TemplateName: templateName,
		Params:       params,
		paramByName:  make(map[string]*TemplateParam, len(params)),
		tfNames:      make(map[string]bool, len(params)),
	}
	for i := range params {
		p := &params[i]
		info.paramByName[p.Name] = p
		info.tfNames[CamelToSnake(p.Name)] = true
	}
	return info, nil
}
