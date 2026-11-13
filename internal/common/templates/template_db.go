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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
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
	DefaultValue   json.RawMessage   `json:"defaultValue,omitempty"`
	Optional       bool              `json:"optional"`
	Annotations    map[string]string `json:"annotations,omitempty"`
	MetaProperties map[string]string `json:"metaProperties,omitempty"`
}

// TemplateInfo holds the cached metadata for a single template.
type TemplateInfo struct {
	TemplateName       string
	ContentType        string
	SupportedPlatforms json.RawMessage
	Params             []TemplateParam
	// Derived indexes — built once at load time.
	paramByName map[string]*TemplateParam // camelCase name → param
	tfNames     map[string]bool           // snake_case tf_name set
}

// HasParam returns true if the template declares a parameter with the
// given camelCase model name.
func (t *TemplateInfo) HasParam(modelName string) bool {
	if t == nil || t.paramByName == nil {
		return false
	}
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
// template cannot be loaded (interface_rsc fail-open).
func (db *TemplateDB) GetTemplate(ctx context.Context, templateName string) *TemplateInfo {
	info, err := db.LoadTemplate(ctx, templateName)
	if err != nil {
		tflog.Warn(ctx, "TemplateDB: failed to load template", map[string]interface{}{
			"template": templateName,
			"error":    err.Error(),
		})
		return nil
	}
	return info
}

// LoadTemplate is GetTemplate with the fetch error. nd_policy uses this
// so a missing template fails apply instead of being skipped.
func (db *TemplateDB) LoadTemplate(ctx context.Context, templateName string) (*TemplateInfo, error) {
	if db == nil {
		return nil, fmt.Errorf("TemplateDB: not initialised")
	}

	db.mu.RLock()
	if info, ok := db.templates[templateName]; ok {
		db.mu.RUnlock()
		return info, nil
	}
	db.mu.RUnlock()

	db.mu.Lock()
	defer db.mu.Unlock()

	if info, ok := db.templates[templateName]; ok {
		return info, nil
	}

	info, err := db.fetchTemplate(ctx, templateName)
	if err != nil {
		return nil, err
	}
	db.templates[templateName] = info

	tflog.Debug(ctx, "TemplateDB: loaded template", map[string]interface{}{
		"template": templateName,
		"params":   len(info.Params),
	})
	return info, nil
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

	path := fmt.Sprintf(templateParamsURL, url.PathEscape(templateName))
	raw, err := db.client.GetRawJson(path)
	if err != nil {
		return nil, fmt.Errorf("TemplateDB: GET %s: %w", path, err)
	}

	info, err := decodeTemplate(templateName, raw)
	if err != nil {
		return nil, fmt.Errorf("TemplateDB: unmarshal %s: %w", templateName, err)
	}
	return info, nil
}

// parametersEnvelope is the wrapper object ND returns for some templates.
type parametersEnvelope struct {
	Name               string          `json:"name,omitempty"`
	TemplateName       string          `json:"templateName,omitempty"`
	ContentType        string          `json:"contentType,omitempty"`
	SupportedPlatforms json.RawMessage `json:"supportedPlatforms,omitempty"`
	Parameters         []TemplateParam `json:"parameters,omitempty"`
}

// decodeTemplate accepts either a bare parameter array (interface_rsc)
// or a {parameters, contentType, supportedPlatforms} object (nd_policy / ND 4.2).
func decodeTemplate(templateName string, raw []byte) (*TemplateInfo, error) {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 {
		return nil, fmt.Errorf("empty parameters payload")
	}

	var params []TemplateParam
	info := &TemplateInfo{TemplateName: templateName}

	if trimmed[0] == '[' {
		if err := json.Unmarshal(raw, &params); err != nil {
			return nil, err
		}
	} else {
		var env parametersEnvelope
		if err := json.Unmarshal(raw, &env); err != nil {
			return nil, err
		}
		params = env.Parameters
		info.ContentType = env.ContentType
		info.SupportedPlatforms = env.SupportedPlatforms
		if info.TemplateName == "" {
			info.TemplateName = env.TemplateName
			if info.TemplateName == "" {
				info.TemplateName = env.Name
			}
		}
	}

	info.Params = params
	info.paramByName = make(map[string]*TemplateParam, len(params))
	info.tfNames = make(map[string]bool, len(params))
	for i := range params {
		p := &params[i]
		p.Name = strings.TrimSpace(p.Name)
		if p.Name == "" {
			continue
		}
		info.paramByName[p.Name] = p
		info.tfNames[CamelToSnake(p.Name)] = true
	}
	return info, nil
}
