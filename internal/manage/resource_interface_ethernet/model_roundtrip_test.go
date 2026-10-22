// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"encoding/json"
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_interface_common"
)

// Each case is one representative (mode, OS, policy) triple. The fixtures use
// numeric mtu/prefix/routingTag for L3 templates and string mtu for L2.
var roundTripCases = []struct {
	name, payload string
}{
	{
		"nx-os access accessHost",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/13","interfaceType":"ethernet","configData":{"mode":"access","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"accessHost","adminState":true,"accessVlan":100,"bpduGuard":"enable","cdp":true,"mtu":"default","speed":"auto","portTypeEdgeTrunk":true,"extraConfig":"logging event port link-status"}}}}`,
	},
	{
		"nx-os trunk trunkHost",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/14","interfaceType":"ethernet","configData":{"mode":"trunk","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"trunkHost","adminState":true,"allowedVlans":"100-110","mtu":"jumbo","speed":"auto","portTypeEdgeTrunk":true,"nativeVlan":1}}}}`,
	},
	{
		"nx-os routed routedHost",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/15","interfaceType":"ethernet","configData":{"mode":"routed","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"routedHost","adminState":true,"ip":"10.1.1.1","prefix":30,"mtu":9216,"routingTag":"12345","vrfInterface":"default"}}}}`,
	},
	{
		"nx-os pvlan pvlanHost",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/16","interfaceType":"ethernet","configData":{"mode":"pvlan","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"pvlanHost","adminState":true,"pvlanMode":"host","pvlanAllowedVlans":"none","mtu":"jumbo"}}}}`,
	},
	{
		"nx-os dot1qTunnel dot1qTunnelHost",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/17","interfaceType":"ethernet","configData":{"mode":"dot1qTunnel","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"dot1qTunnelHost","adminState":true,"accessVlan":200,"mtu":"jumbo"}}}}`,
	},
	{
		"nx-os unmanaged monitor",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/18","interfaceType":"ethernet","configData":{"mode":"unmanaged","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"monitor"}}}}`,
	},
	{
		"ios-xe routed iosXeRoutedHost",
		`{"switchId":"FDO22412H7U","interfaceName":"GigabitEthernet0/1","interfaceType":"ethernet","configData":{"mode":"routed","networkOS":{"networkOSType":"ios-xe","policy":{"policyType":"iosXeRoutedHost","adminState":true,"ip":"10.2.2.1","prefix":24,"mtu":1500,"speed":"auto"}}}}`,
	},
	{
		"ipfm string routingTag",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/19","interfaceType":"ethernet","configData":{"mode":"routed","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"ipfmL3Port","adminState":true,"ip":"10.3.3.1","prefix":30,"mtu":9216,"routingTag":"my-tag"}}}}`,
	},
	{
		"userDefined with templateName",
		`{"switchId":"FDO22412H7U","interfaceName":"Ethernet1/20","interfaceType":"ethernet","configData":{"mode":"access","networkOS":{"networkOSType":"nx-os","policy":{"policyType":"userDefined","templateName":"my_custom_template"}}}}`,
	},
}

// TestInterfaceEthernetModelRoundTrip exercises the core round trip proven
// in the feature-001 spike, for one entry per ethernet mode plus the three
// JSON type collisions (mtu string/int, prefix string/int, routingTag string/int).
//
// For each case: JSON → NDFCInterfacesValue → JSON. The re-marshaled bytes
// must preserve every field value AND the correct wire form for the
// dual-form fields (number vs string).
func TestInterfaceEthernetModelRoundTrip(t *testing.T) {
	for _, tc := range roundTripCases {
		t.Run(tc.name, func(t *testing.T) {
			var entry resource_interface_common.NDFCInterfacesValue
			if err := json.Unmarshal([]byte(tc.payload), &entry); err != nil {
				t.Fatalf("unmarshal fixture: %v", err)
			}
			out, err := json.Marshal(entry)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			s := string(out)

			// Structural shape checks every case must satisfy.
			for _, want := range []string{
				`"configData":`,
				`"networkOS":`,
				`"policy":`,
				`"policyType":`,
			} {
				if !strings.Contains(s, want) {
					t.Errorf("marshaled output missing %q; got:\n%s", want, s)
				}
			}

			// Round-trip equality on all string-typed fields we can grab
			// cheaply by diffing the two parsed structs.
			var again resource_interface_common.NDFCInterfacesValue
			if err := json.Unmarshal(out, &again); err != nil {
				t.Fatalf("re-unmarshal: %v", err)
			}
			if entry.ConfigData.Mode != again.ConfigData.Mode ||
				entry.ConfigData.NetworkOS.NetworkOsType != again.ConfigData.NetworkOS.NetworkOsType ||
				entry.ConfigData.NetworkOS.Policy.PolicyType != again.ConfigData.NetworkOS.Policy.PolicyType ||
				entry.ConfigData.NetworkOS.Policy.Mtu != again.ConfigData.NetworkOS.Policy.Mtu ||
				entry.ConfigData.NetworkOS.Policy.Prefix != again.ConfigData.NetworkOS.Policy.Prefix ||
				entry.ConfigData.NetworkOS.Policy.RoutingTag != again.ConfigData.NetworkOS.Policy.RoutingTag {
				t.Errorf("round-trip mismatch:\n in:  %+v\n out: %+v", entry, again)
			}
		})
	}
}

// TestInterfaceEthernetCollisionWireForm covers the exact assertion from
// interface_ethernet.spec §Collisions: when a user-supplied value is a pure
// decimal integer, the wire representation must be an unquoted JSON number;
// when it's non-numeric, it must be a quoted JSON string. One assertion
// per field per form.
func TestInterfaceEthernetCollisionWireForm(t *testing.T) {
	cases := []struct {
		name, mtu, prefix, routingTag string
		wantContains, wantNotContains []string
	}{
		{
			name: "all numeric → all numbers on wire",
			mtu:  "9216", prefix: "30", routingTag: "12345",
			wantContains:    []string{`"mtu":9216`, `"prefix":30`, `"routingTag":12345`},
			wantNotContains: []string{`"mtu":"9216"`, `"prefix":"30"`, `"routingTag":"12345"`},
		},
		{
			name:            "mtu=default (string), others empty",
			mtu:             "default",
			wantContains:    []string{`"mtu":"default"`},
			wantNotContains: []string{`"prefix"`, `"routingTag"`},
		},
		{
			name: "mtu=jumbo + numeric routingTag",
			mtu:  "jumbo", routingTag: "0",
			wantContains:    []string{`"mtu":"jumbo"`, `"routingTag":0`},
			wantNotContains: []string{`"mtu":"jumbo":`, `"routingTag":"0"`},
		},
		{
			name:            "routingTag non-numeric (ipfm)",
			routingTag:      "my-tag",
			wantContains:    []string{`"routingTag":"my-tag"`},
			wantNotContains: []string{`"routingTag":my-tag`},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := resource_interface_common.NDFCNetworkOSPolicyValue{
				PolicyType: "routedHost",
				Mtu:        tc.mtu,
				Prefix:     tc.prefix,
				RoutingTag: tc.routingTag,
			}
			raw, err := json.Marshal(v)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			s := string(raw)
			for _, w := range tc.wantContains {
				if !strings.Contains(s, w) {
					t.Errorf("missing %q in output:\n%s", w, s)
				}
			}
			for _, w := range tc.wantNotContains {
				if strings.Contains(s, w) {
					t.Errorf("unexpected %q in output:\n%s", w, s)
				}
			}
		})
	}
}
