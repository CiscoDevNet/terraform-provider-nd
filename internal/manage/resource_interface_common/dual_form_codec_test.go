// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"encoding/json"
	"strings"
	"testing"
)

// policyWithFields is a lightweight helper to build a NDFCNetworkOSPolicyValue
// with just the fields each sub-test exercises.
func policyWithFields(policy, mtu, prefix, routingTag string) NDFCNetworkOSPolicyValue {
	return NDFCNetworkOSPolicyValue{
		PolicyType: policy,
		Mtu:        mtu,
		Prefix:     prefix,
		RoutingTag: routingTag,
	}
}

// TestDualFormMarshalNumericForms verifies pure-numeric collision strings are
// emitted as JSON numbers (unquoted) by MarshalJSON on NDFCNetworkOSPolicyValue.
func TestDualFormMarshalNumericForms(t *testing.T) {
	v := policyWithFields("routedHost", "9216", "30", "12345")
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	s := string(raw)
	// Unquoted number tokens for each of the three collision fields.
	for _, want := range []string{`"mtu":9216`, `"prefix":30`, `"routingTag":12345`} {
		if !strings.Contains(s, want) {
			t.Errorf("marshal output missing numeric form %q; got:\n%s", want, s)
		}
	}
	// Must not emit the quoted-string variant for these three fields.
	for _, bad := range []string{`"mtu":"9216"`, `"prefix":"30"`, `"routingTag":"12345"`} {
		if strings.Contains(s, bad) {
			t.Errorf("marshal output has quoted form it should not: %q\n%s", bad, s)
		}
	}
}

// TestDualFormMarshalStringForms verifies non-numeric collision strings are
// emitted as JSON strings (quoted) — e.g. mtu=default|jumbo for nx-os L2.
func TestDualFormMarshalStringForms(t *testing.T) {
	v := policyWithFields("accessHost", "default", "", "")
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	s := string(raw)
	if !strings.Contains(s, `"mtu":"default"`) {
		t.Errorf("expected quoted mtu=default, got:\n%s", s)
	}
	// omitempty: prefix/routingTag empty → not present at all.
	if strings.Contains(s, `"prefix"`) || strings.Contains(s, `"routingTag"`) {
		t.Errorf("empty collision fields leaked into output:\n%s", s)
	}
}

// TestDualFormMarshalIPFMRoutingTagIsString verifies that for ipfmL3Port,
// a numeric routingTag is emitted as a JSON string (not a number) because
// the IPFM template defines routingTag as string-typed.
func TestDualFormMarshalIPFMRoutingTagIsString(t *testing.T) {
	v := policyWithFields("ipfmL3Port", "9216", "30", "12345")
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	s := string(raw)
	// mtu and prefix should be numbers (ipfmL3Port declares them as number).
	if !strings.Contains(s, `"mtu":9216`) {
		t.Errorf("expected numeric mtu for ipfmL3Port, got:\n%s", s)
	}
	if !strings.Contains(s, `"prefix":30`) {
		t.Errorf("expected numeric prefix for ipfmL3Port, got:\n%s", s)
	}
	// routingTag should be a quoted string (ipfmL3Port declares it as string).
	if !strings.Contains(s, `"routingTag":"12345"`) {
		t.Errorf("expected quoted routingTag for ipfmL3Port, got:\n%s", s)
	}
	// Must NOT have bare numeric routingTag.
	if strings.Contains(s, `"routingTag":12345`) && !strings.Contains(s, `"routingTag":"12345"`) {
		t.Errorf("routingTag should be string for ipfmL3Port, got:\n%s", s)
	}
}

// TestDualFormMarshalL2MtuIsString verifies that for nx-os L2 policies,
// a numeric mtu like "100" is emitted as a JSON string, not a number.
func TestDualFormMarshalL2MtuIsString(t *testing.T) {
	v := policyWithFields("accessHost", "100", "", "")
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	s := string(raw)
	// mtu should be a quoted string for accessHost even when numeric.
	if !strings.Contains(s, `"mtu":"100"`) {
		t.Errorf("expected quoted mtu for accessHost, got:\n%s", s)
	}
}

// TestDualFormUnmarshalAccepts both JSON number and JSON string inputs for
// each of the three collision fields and canonicalises them to a Go string.
func TestDualFormUnmarshalAccepts(t *testing.T) {
	cases := []struct {
		name, in, wantMtu, wantPrefix, wantRoutingTag string
	}{
		{
			"nx-os L3 numeric form",
			`{"policyType":"routedHost","mtu":9216,"prefix":30,"routingTag":12345}`,
			"9216", "30", "12345",
		},
		{
			"nx-os L2 string form (mtu)",
			`{"policyType":"accessHost","mtu":"default"}`,
			"default", "", "",
		},
		{
			"ipfm string form (routingTag)",
			`{"policyType":"ipfmL3Port","routingTag":"my-tag"}`,
			"", "", "my-tag",
		},
		{
			"mixed forms",
			`{"policyType":"routedHost","mtu":"jumbo","prefix":31,"routingTag":"0"}`,
			"jumbo", "31", "0",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var v NDFCNetworkOSPolicyValue
			if err := json.Unmarshal([]byte(tc.in), &v); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if v.Mtu != tc.wantMtu {
				t.Errorf("Mtu: got %q want %q", v.Mtu, tc.wantMtu)
			}
			if v.Prefix != tc.wantPrefix {
				t.Errorf("Prefix: got %q want %q", v.Prefix, tc.wantPrefix)
			}
			if v.RoutingTag != tc.wantRoutingTag {
				t.Errorf("RoutingTag: got %q want %q", v.RoutingTag, tc.wantRoutingTag)
			}
		})
	}
}

// TestDualFormRoundTrip ensures Marshal → Unmarshal is identity for the
// collision fields across the number/string split.
func TestDualFormRoundTrip(t *testing.T) {
	cases := []NDFCNetworkOSPolicyValue{
		policyWithFields("routedHost", "9216", "30", "12345"),
		policyWithFields("accessHost", "default", "", ""),
		policyWithFields("ipfmL3Port", "9216", "30", "boundary-tag"),
	}
	for i, v := range cases {
		raw, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("case %d marshal: %v", i, err)
		}
		var back NDFCNetworkOSPolicyValue
		if err := json.Unmarshal(raw, &back); err != nil {
			t.Fatalf("case %d unmarshal: %v (payload=%s)", i, err, raw)
		}
		if back.Mtu != v.Mtu || back.Prefix != v.Prefix || back.RoutingTag != v.RoutingTag {
			t.Errorf("case %d round-trip mismatch: got (%q,%q,%q) want (%q,%q,%q)",
				i, back.Mtu, back.Prefix, back.RoutingTag, v.Mtu, v.Prefix, v.RoutingTag)
		}
	}
}

// TestInterfacesValueNestedPayloadShape asserts that marshalling an
// InterfacesValue produces the exact wire shape required by the ND
// createInterface API — `configData.networkOS.policy.<attr>` with
// interfaceType at the root and the collision dual-form respected.
func TestInterfacesValueNestedPayloadShape(t *testing.T) {
	entry := NDFCInterfacesValue{
		InterfaceName: "Ethernet1/14",
		InterfaceType: "ethernet",
		SwitchId:      "FDO22412H7U",
		ConfigData: NDFCConfigDataValue{
			Mode: "routed",
			NetworkOS: NDFCConfigDataNetworkOSValue{
				NetworkOsType: "nx-os",
				Policy:        policyWithFields("routedHost", "9216", "30", "12345"),
			},
		},
	}
	raw, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(raw)
	// Shape probes.
	for _, want := range []string{
		`"interfaceName":"Ethernet1/14"`,
		`"interfaceType":"ethernet"`,
		`"switchId":"FDO22412H7U"`,
		`"configData":`,
		`"networkOS":`,
		`"networkOSType":"nx-os"`,
		`"policy":`,
		`"policyType":"routedHost"`,
		// Collision fields all numeric here.
		`"mtu":9216`,
		`"prefix":30`,
		`"routingTag":12345`,
	} {
		if !strings.Contains(s, want) {
			t.Errorf("marshaled payload missing %q; got:\n%s", want, s)
		}
	}

	// Round-trip: parse back and ensure structural equality on the probed fields.
	var back NDFCInterfacesValue
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if back.InterfaceName != "Ethernet1/14" ||
		back.ConfigData.Mode != "routed" ||
		back.ConfigData.NetworkOS.NetworkOsType != "nx-os" ||
		back.ConfigData.NetworkOS.Policy.PolicyType != "routedHost" ||
		back.ConfigData.NetworkOS.Policy.Mtu != "9216" ||
		back.ConfigData.NetworkOS.Policy.Prefix != "30" ||
		back.ConfigData.NetworkOS.Policy.RoutingTag != "12345" {
		t.Errorf("round-trip lost field data: %+v", back)
	}
}
