// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	"strings"
	"unicode"
)

// PolicyToTemplate maps each built-in policyType to its canonical
// templateName as documented in the Attribute Summary. This is the only
// hardcoded knowledge; the per-template attribute list is fetched live
// from the ND Template Library API.
var PolicyToTemplate = map[string]string{
	// Ethernet — access
	"iosXeAccess":    "ios_xe_int_access_host",
	"accessHost":     "int_access_host",
	"ipfmAccessHost": "int_ipfm_access_host",

	// Ethernet — dot1qTunnel
	"dot1qTunnelHost": "int_dot1q_tunnel_host",

	// Ethernet — managed
	"iosXeStackwiseDualActive": "ios_xe_int_stackwise_dual_active",
	"iosXeStackwiseLinkMember": "ios_xe_int_stackwise_link",

	// Ethernet — pvlan
	"pvlanHost": "int_pvlan_host",

	// Ethernet — routed
	"csrMultisiteIfcMember":   "csr_multisite_ifc_underlay",
	"iosXeInternalL3PoMember": "ios_xe_int_l3_port_channel_member_internal",
	"iosXeNumbered":           "ios_xe_int_fabric_num",
	"iosXeRoutedHost":         "ios_xe_int_routed_host",
	"iosXrRoutedHost":         "ios_xr_int_routed_host",
	"dataBrokerL3Host":        "int_ndb_routed_host",
	"endPointLocator":         "epl_routed_intf",
	"ipfmL3Port":              "int_ipfm_l3_port",
	"routedHost":              "int_routed_host",

	// Ethernet — trunk
	"iosXeTrunkHost":     "ios_xe_int_trunk_host",
	"classicHost":        "int_trunk_classic",
	"dataBrokerHost":     "int_ndb_trunk_host",
	"dataBrokerPoMember": "int_ndb_port_channel_member_trunk_host",
	"ipfmTrunkHost":      "int_ipfm_trunk_host",
	"trunkHost":          "int_trunk_host",

	// Ethernet — unmanaged
	"iosXeMonitor": "ios_xe_int_monitor_ethernet",
	"monitor":      "int_monitor_ethernet",

	// Port-channel — access
	"iosXeAccessPoHost": "ios_xe_int_port_channel_access_host",
	"accessPoHost":      "int_port_channel_access_host",
	"ipfmAccessPoHost":  "int_ipfm_port_channel_access_host",

	// Port-channel — dot1qTunnel
	"dot1qTunnelPoHost": "int_port_channel_dot1q_tunnel_host",

	// Port-channel — fex
	"fexPo": "int_port_channel_fex",

	// Port-channel — pvlan
	"pvlanPoHost": "int_port_channel_pvlan_host",

	// Port-channel — routed
	"iosXeInternalPoMember": "ios_xe_int_l3_port_channel_internal",
	"iosXeL3PortChannel":    "ios_xe_int_l3_port_channel",
	"l3Po":                  "int_l3_port_channel",

	// Port-channel — trunk
	"iosXeTrunkPoHost": "ios_xe_int_port_channel_trunk_host",
	"dataBrokerPo":     "int_ndb_port_channel_trunk_host",
	"ipfmTrunkPoHost":  "int_ipfm_port_channel_trunk_host",
	"trunkPoHost":      "int_port_channel_trunk_host",

	// Port-channel — unmanaged
	"monitorPo": "int_monitor_port_channel",
}

// CamelToSnake converts a camelCase or PascalCase identifier to
// snake_case, following the same conventions as the Terraform code
// generator. Consecutive uppercase letters are treated as an acronym
// (e.g. "ipV6Address" → "ip_v6_address", "bpduGuard" → "bpdu_guard").
func CamelToSnake(s string) string {
	if s == "" {
		return s
	}
	var b strings.Builder
	b.Grow(len(s) + 4) // small extra for underscores
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				prev := runes[i-1]
				// Insert underscore before an uppercase letter when the
				// previous character is lowercase/digit, OR when it's
				// the start of a new word after an acronym run
				// (e.g. "...OS|T..." → "...os_t...").
				if unicode.IsLower(prev) || unicode.IsDigit(prev) {
					b.WriteRune('_')
				} else if unicode.IsUpper(prev) && i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
					// Transition from acronym to new word: "OSType" → "os_type"
					b.WriteRune('_')
				}
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
