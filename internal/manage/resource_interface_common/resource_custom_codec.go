// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// ---- Payload wrappers used by the bulk interface endpoints. --------------

// NDFCInterfacesPayload wraps a list of interface objects for the bulk
// POST .../switches/{sw}/interfaces create endpoint.
type NDFCInterfacesPayload struct {
	Interfaces []NDFCInterfacesValue `json:"interfaces"`
}

// NDFCInterfaceAction is the {interfaceName, switchId} selector used by the
// bulk actions endpoints (deploy, remove, preview, rediscover).
type NDFCInterfaceAction struct {
	InterfaceName string `json:"interfaceName"`
	SwitchId      string `json:"switchId"`
}

// NDFCInterfaceActionPayload wraps a list of interface selectors for the
// bulk actions endpoints.
type NDFCInterfaceActionPayload struct {
	Interfaces []NDFCInterfaceAction `json:"interfaces"`
}

// GenericBulkResult is one row of the HTTP 207 createInterface response.
type GenericBulkResult struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// GenericBulkResponse is the HTTP 207 response body of createInterface and
// the bulk action endpoints.
type GenericBulkResponse struct {
	Results []GenericBulkResult `json:"results,omitempty"`
}

// ---- Dual-form codec for the three JSON type collisions. -----------------
//
// Per spec `interface_ethernet.spec` §Collisions, the keys `mtu`, `prefix`
// and `routingTag` are declared as `integer` in some templates and `string`
// in others (e.g. nx-os L2 `mtu` is one of `default`/`jumbo` strings, while
// L3/ios-xe `mtu` is a bounded integer). We expose each as a Terraform
// `String` and write a hand-written MarshalJSON / UnmarshalJSON on
// NDFCNetworkOSPolicyValue that:
//
//   - on Marshal: uses the active policyType to determine the expected JSON
//     type for each collision field. Policies that define the field as an
//     integer get a JSON number; string-typed variants get a JSON string.
//     When the policy is unknown, falls back to value-sniffing (pure decimal
//     integer → JSON number, anything else → JSON string).
//   - on Unmarshal: accepts either a JSON number or a JSON string, storing
//     the canonical textual form in the Go struct.
//
// The implementation uses a type alias to delegate to encoding/json's
// generated marshal for every other field.

// collisionFieldJSONType maps (policyType, fieldName) to the expected JSON
// wire type: "number" for JSON integers, "string" for JSON strings.
// Absent entries fall back to the legacy value-sniffing heuristic.
var collisionFieldJSONType = map[string]map[string]string{
	// nx-os L2 templates: mtu is a string enum (default|jumbo).
	"accessHost":      {"mtu": "string"},
	"trunkHost":       {"mtu": "string"},
	"classicHost":     {"mtu": "string"},
	"dot1qTunnelHost": {"mtu": "string"},
	"pvlanHost":       {"mtu": "string"},
	"dataBrokerHost":  {"mtu": "string"},
	"ipfmAccessHost":  {"mtu": "string"},
	"ipfmTrunkHost":   {"mtu": "string"},
	// nx-os L3 and ios-xe templates: mtu, prefix, routingTag are integers.
	"routedHost":       {"mtu": "number", "prefix": "number", "routingTag": "number"},
	"endPointLocator":  {"mtu": "number", "prefix": "number", "routingTag": "number"},
	"dataBrokerL3Host": {"mtu": "number", "prefix": "number", "routingTag": "number"},
	"iosXeRoutedHost":  {"mtu": "number", "prefix": "number"},
	"iosXeNumbered":    {"mtu": "number", "prefix": "string"},
	"iosXeAccess":      {"mtu": "number"},
	"iosXeTrunkHost":   {"mtu": "number"},
	// IPFM L3: routingTag is a string in IPFM templates.
	"ipfmL3Port": {"mtu": "number", "prefix": "number", "routingTag": "string"},
}

type policyRaw struct {
	Mtu        json.RawMessage `json:"mtu,omitempty"`
	Prefix     json.RawMessage `json:"prefix,omitempty"`
	RoutingTag json.RawMessage `json:"routingTag,omitempty"`
}

// policyShadow is a type alias to reuse the default (field-reflection) JSON
// encoding for every field except mtu/prefix/routingTag. We rename those
// three on the alias so the default encoding does not emit them, then our
// MarshalJSON merges in the dual-form bytes manually.
type policyShadow NDFCNetworkOSPolicyValue

func (v NDFCNetworkOSPolicyValue) MarshalJSON() ([]byte, error) {
	// userDefined policies must only emit policyType, templateName, and
	// templateConfig — all built-in policy fields are stripped to prevent
	// leaking into the ND request.
	if v.PolicyType == "userDefined" {
		obj := map[string]interface{}{
			"policyType": v.PolicyType,
		}
		if v.TemplateName != "" {
			obj["templateName"] = v.TemplateName
		}
		if len(v.TemplateConfig) > 0 {
			obj["templateConfig"] = v.TemplateConfig
		}
		return json.Marshal(obj)
	}

	// Marshal everything except the three collision fields using the default
	// encoder (via the type alias to short-circuit MarshalJSON recursion).
	stripped := policyShadow(v)
	stripped.Mtu = ""
	stripped.Prefix = ""
	stripped.RoutingTag = ""
	base, err := json.Marshal(stripped)
	if err != nil {
		return nil, err
	}

	// Marshal the three collision fields using the policy-aware encoder.
	pt := v.PolicyType
	extra := map[string]json.RawMessage{}
	if b, ok := encodeDualForm(v.Mtu, pt, "mtu"); ok {
		extra["mtu"] = b
	}
	if b, ok := encodeDualForm(v.Prefix, pt, "prefix"); ok {
		extra["prefix"] = b
	}
	if b, ok := encodeDualForm(v.RoutingTag, pt, "routingTag"); ok {
		extra["routingTag"] = b
	}
	if len(extra) == 0 {
		return base, nil
	}

	// Merge: strip trailing `}` from base, append `,"mtu":...`.
	out := bytes.TrimRight(base, "}")
	if !bytes.Equal(out, []byte("{")) {
		out = append(out, ',')
	}
	first := true
	for _, k := range []string{"mtu", "prefix", "routingTag"} {
		b, ok := extra[k]
		if !ok {
			continue
		}
		if !first {
			out = append(out, ',')
		}
		first = false
		out = append(out, '"')
		out = append(out, k...)
		out = append(out, '"', ':')
		out = append(out, b...)
	}
	out = append(out, '}')
	return out, nil
}

func (v *NDFCNetworkOSPolicyValue) UnmarshalJSON(data []byte) error {
	// First parse everything except the three fields via the shadow alias.
	// We don't try to filter them out of the data since struct unmarshal
	// into a *string will fail for a number — so instead we unmarshal into
	// policyShadow which has the fields as plain strings but tolerate the
	// number case below by stripping them first.
	stripped, err := stripCollisionFields(data)
	if err != nil {
		return err
	}
	var sh policyShadow
	if err := json.Unmarshal(stripped, &sh); err != nil {
		return fmt.Errorf("decode interface policy: %w", err)
	}
	*v = NDFCNetworkOSPolicyValue(sh)

	// Now extract the three collision fields in dual-form.
	var raw policyRaw
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("decode interface policy (collision fields): %w", err)
	}
	if s, ok := decodeDualForm(raw.Mtu); ok {
		v.Mtu = s
	}
	if s, ok := decodeDualForm(raw.Prefix); ok {
		v.Prefix = s
	}
	if s, ok := decodeDualForm(raw.RoutingTag); ok {
		v.RoutingTag = s
	}
	return nil
}

// encodeDualForm returns the JSON bytes to emit for a collision-field value.
// Empty string => emit nothing (ok=false).
//
// Resolution priority for determining the expected JSON wire type:
//  1. Static collisionFieldJSONType map (fastest, no network).
//  2. TemplateDB live lookup (cached after first call per template).
//  3. Value-sniffing (last resort for unknown policies/fields).
//
// "number" always emits a bare JSON number, "string" always emits a
// quoted JSON string — even if the value is a pure digit string.
func encodeDualForm(s, policyType, fieldName string) (json.RawMessage, bool) {
	if s == "" {
		return nil, false
	}
	wantType := resolveCollisionType(policyType, fieldName)
	if wantType == "string" {
		b, _ := json.Marshal(s)
		return b, true
	}
	if wantType == "number" {
		if _, err := strconv.ParseInt(s, 10, 64); err == nil {
			return json.RawMessage(s), true
		}
		// Not a valid integer for a number-typed field — emit as string.
		b, _ := json.Marshal(s)
		return b, true
	}
	// wantType == "" — unknown; value-sniff.
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return json.RawMessage(s), true
	}
	b, _ := json.Marshal(s)
	return b, true
}

// resolveCollisionType determines the expected JSON type ("number" or
// "string") for a collision field. Returns "" when undetermined.
func resolveCollisionType(policyType, fieldName string) string {
	// 1. Static map — covers all known built-in policies.
	if fields, ok := collisionFieldJSONType[policyType]; ok {
		if wantType, ok := fields[fieldName]; ok {
			return wantType
		}
	}
	// 2. TemplateDB dynamic lookup (cached, thread-safe).
	return resolveCollisionTypeFromTemplateDB(policyType, fieldName)
}

// decodeDualForm reads an incoming JSON value that could be either a number
// or a string, and returns the canonical textual form. Null / empty raw =>
// returns ("", false).
func decodeDualForm(raw json.RawMessage) (string, bool) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || bytes.Equal(raw, []byte("null")) {
		return "", false
	}
	if raw[0] == '"' {
		var s string
		if err := json.Unmarshal(raw, &s); err == nil {
			return s, true
		}
		return "", false
	}
	// Treat as number — keep the exact token.
	return string(raw), true
}

// stripCollisionFields removes mtu/prefix/routingTag from a top-level JSON
// object so the shadow unmarshal (which types them as `string`) does not
// error out when they arrive as JSON numbers. The removed fields are
// re-decoded via the dual-form path.
func stripCollisionFields(data []byte) ([]byte, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return data, nil // not an object — let caller surface the error
	}
	delete(m, "mtu")
	delete(m, "prefix")
	delete(m, "routingTag")
	return json.Marshal(m)
}
