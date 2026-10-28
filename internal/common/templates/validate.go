// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	"sort"
	"strings"
)

// UnknownParamNames is the interface_rsc allow-list check using native
// template parameter names (featureName, TRACK_ID). Interface ethernet
// uses AllowedTFNames() for the snake_case variant of the same idea.
func (t *TemplateInfo) UnknownParamNames(provided map[string]string) []string {
	if t == nil {
		return nil
	}
	var unknown []string
	for key := range provided {
		if !t.HasParam(key) {
			unknown = append(unknown, key)
		}
	}
	sort.Strings(unknown)
	return unknown
}

// UnknownTFNames is the interface_rsc validateAttrsAgainstTemplate check:
// configured snake_case TF attributes that the live template does not declare.
func (t *TemplateInfo) UnknownTFNames(configured map[string]bool) []string {
	if t == nil {
		return nil
	}
	allowed := t.AllowedTFNames()
	var unknown []string
	for name := range configured {
		if !allowed[name] {
			unknown = append(unknown, name)
		}
	}
	sort.Strings(unknown)
	return unknown
}

// MissingRequiredNames returns required parameters that are absent from
// provided and have no default. nd_policy fail-closed uses this; interface
// does not enforce required params.
func (t *TemplateInfo) MissingRequiredNames(provided map[string]string) []string {
	if t == nil {
		return nil
	}
	var missing []string
	for i := range t.Params {
		p := &t.Params[i]
		if p.Name == "" || p.IsOptional() {
			continue
		}
		if _, ok := provided[p.Name]; ok {
			continue
		}
		if p.HasDefault() {
			continue
		}
		missing = append(missing, p.Name)
	}
	sort.Strings(missing)
	return missing
}

// IsOptional treats optional=true, IsMandatory=false, or a conditional
// IsMandatory expression (e.g. PROTOCOL==tcp) as optional.
func (p TemplateParam) IsOptional() bool {
	if p.Optional {
		return true
	}
	if p.Annotations == nil {
		return false
	}
	mandatory := strings.Trim(strings.TrimSpace(p.Annotations["IsMandatory"]), `"`)
	if mandatory == "" {
		return false
	}
	if strings.EqualFold(mandatory, "true") {
		return false
	}
	return true
}

// HasDefault is true when defaultValue or metaProperties.defaultValue is set.
func (p TemplateParam) HasDefault() bool {
	if len(p.DefaultValue) > 0 && string(p.DefaultValue) != "null" && string(p.DefaultValue) != `""` {
		return true
	}
	if p.MetaProperties == nil {
		return false
	}
	dv := strings.TrimSpace(p.MetaProperties["defaultValue"])
	return dv != "" && dv != "null"
}
