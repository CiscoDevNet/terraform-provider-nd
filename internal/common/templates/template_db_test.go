// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	"testing"
)

// --- CamelToSnake -----------------------------------------------------------

func TestCamelToSnakeBasic(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"accessVlan", "access_vlan"},
		{"adminState", "admin_state"},
		{"bpduGuard", "bpdu_guard"},
		{"extraConfig", "extra_config"},
		{"mtu", "mtu"},
		{"speed", "speed"},
		{"portTypeEdgeTrunk", "port_type_edge_trunk"},
		{"stormControlBroadcastLevel", "storm_control_broadcast_level"},
		{"stormControlBroadcastLevelPps", "storm_control_broadcast_level_pps"},
		{"vrfInterface", "vrf_interface"},
		{"routingTag", "routing_tag"},
		{"pimDrPriority", "pim_dr_priority"},
		{"description", "description"},
		{"cdp", "cdp"},
	}
	for _, tc := range tests {
		got := CamelToSnake(tc.in)
		if got != tc.want {
			t.Errorf("CamelToSnake(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestCamelToSnakeAcronyms(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"ipAddress", "ip_address"},
		{"ipV6Address", "ip_v6_address"},
		{"networkOSType", "network_os_type"},
		{"pvlanNativeVlan", "pvlan_native_vlan"},
		{"qosPolicy", "qos_policy"},
		{"ipRedirects", "ip_redirects"},
		{"ipv6Prefix", "ipv6_prefix"},
	}
	for _, tc := range tests {
		got := CamelToSnake(tc.in)
		if got != tc.want {
			t.Errorf("CamelToSnake(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestCamelToSnakeEdge(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"", ""},
		{"x", "x"},
		{"X", "x"},
		{"policyType", "policy_type"},
	}
	for _, tc := range tests {
		got := CamelToSnake(tc.in)
		if got != tc.want {
			t.Errorf("CamelToSnake(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// --- PolicyToTemplate -------------------------------------------------------

func TestPolicyToTemplateCompleteness(t *testing.T) {
	// Verify that every Ethernet policy from the attribute summary has a
	// mapping. These are the 24 Ethernet policy types.
	ethernetPolicies := []string{
		"iosXeAccess", "accessHost", "ipfmAccessHost",
		"dot1qTunnelHost",
		"iosXeStackwiseDualActive", "iosXeStackwiseLinkMember",
		"pvlanHost",
		"csrMultisiteIfcMember", "iosXeInternalL3PoMember", "iosXeNumbered",
		"iosXeRoutedHost", "iosXrRoutedHost", "dataBrokerL3Host",
		"endPointLocator", "ipfmL3Port", "routedHost",
		"iosXeTrunkHost", "classicHost", "dataBrokerHost",
		"dataBrokerPoMember", "ipfmTrunkHost", "trunkHost",
		"iosXeMonitor", "monitor",
	}
	for _, p := range ethernetPolicies {
		if _, ok := PolicyToTemplate[p]; !ok {
			t.Errorf("PolicyToTemplate missing ethernet policy %q", p)
		}
	}
}

// --- TemplateInfo helpers ---------------------------------------------------

func TestTemplateInfoHasParamAndType(t *testing.T) {
	info := &TemplateInfo{
		TemplateName: "int_access_host",
		Params: []TemplateParam{
			{Name: "accessVlan", ParamType: "integer"},
			{Name: "mtu", ParamType: "string"},
			{Name: "adminState", ParamType: "boolean"},
		},
		paramByName: map[string]*TemplateParam{},
		tfNames:     map[string]bool{},
	}
	for i := range info.Params {
		p := &info.Params[i]
		info.paramByName[p.Name] = p
		info.tfNames[CamelToSnake(p.Name)] = true
	}

	if !info.HasParam("accessVlan") {
		t.Error("expected HasParam(accessVlan) = true")
	}
	if info.HasParam("bpduGuard") {
		t.Error("expected HasParam(bpduGuard) = false")
	}
	if pt := info.ParamType("mtu"); pt != "string" {
		t.Errorf("ParamType(mtu) = %q, want %q", pt, "string")
	}
	if pt := info.ParamType("missing"); pt != "" {
		t.Errorf("ParamType(missing) = %q, want empty", pt)
	}

	names := info.AllowedTFNames()
	for _, want := range []string{"access_vlan", "mtu", "admin_state"} {
		if !names[want] {
			t.Errorf("AllowedTFNames missing %q", want)
		}
	}
	if names["bpdu_guard"] {
		t.Error("AllowedTFNames should not contain bpdu_guard")
	}
}
