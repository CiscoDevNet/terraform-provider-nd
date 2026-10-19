// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"testing"
)

// unit-only check: the discriminator matrix encodes every (mode, OS, policy)
// triple called out in interface_ethernet.spec. Lightweight guard so an
// accidental deletion is caught in CI without needing a full plan/apply.
func TestEthernetMatrixCoverage(t *testing.T) {
	// Each row: mode, OS, policy that must be legal.
	cases := []struct{ mode, os, policy string }{
		{"access", "ios-xe", "iosXeAccess"},
		{"access", "nx-os", "accessHost"},
		{"access", "nx-os", "ipfmAccessHost"},
		{"dot1qTunnel", "nx-os", "dot1qTunnelHost"},
		{"managed", "ios-xe", "iosXeStackwiseDualActive"},
		{"managed", "ios-xe", "iosXeStackwiseLinkMember"},
		{"pvlan", "nx-os", "pvlanHost"},
		{"routed", "ios-xe", "csrMultisiteIfcMember"},
		{"routed", "ios-xe", "iosXeInternalL3PoMember"},
		{"routed", "ios-xe", "iosXeNumbered"},
		{"routed", "ios-xe", "iosXeRoutedHost"},
		{"routed", "ios-xr", "iosXrRoutedHost"},
		{"routed", "nx-os", "dataBrokerL3Host"},
		{"routed", "nx-os", "endPointLocator"},
		{"routed", "nx-os", "ipfmL3Port"},
		{"routed", "nx-os", "routedHost"},
		{"trunk", "ios-xe", "iosXeTrunkHost"},
		{"trunk", "nx-os", "classicHost"},
		{"trunk", "nx-os", "dataBrokerHost"},
		{"trunk", "nx-os", "dataBrokerPoMember"},
		{"trunk", "nx-os", "ipfmTrunkHost"},
		{"trunk", "nx-os", "trunkHost"},
		{"unmanaged", "ios-xe", "iosXeMonitor"},
		{"unmanaged", "nx-os", "monitor"},
	}
	for _, tc := range cases {
		byOS, ok := ethernetDiscriminatorMatrix[tc.mode]
		if !ok {
			t.Errorf("mode %q missing from matrix", tc.mode)
			continue
		}
		policies, ok := byOS[tc.os]
		if !ok {
			t.Errorf("(mode=%q, os=%q) missing from matrix", tc.mode, tc.os)
			continue
		}
		if !contains(policies, tc.policy) {
			t.Errorf("policy %q not in matrix for (mode=%q, os=%q); got %v",
				tc.policy, tc.mode, tc.os, policies)
		}
	}
}

// TestEthernetMatrixRejectsUnknown ensures invalid tuples are caught — this
// is the inverse of coverage: we want negative assertions too.
func TestEthernetMatrixRejectsUnknown(t *testing.T) {
	bad := []struct{ mode, os, policy string }{
		// Right mode/OS pair, wrong policy for it.
		{"access", "nx-os", "trunkHost"},
		{"trunk", "nx-os", "accessHost"},
		{"routed", "nx-os", "iosXeRoutedHost"},
		{"routed", "ios-xe", "routedHost"},
		// Mode exists but OS doesn't support ethernet on it.
		{"access", "ios-xr", "iosXeAccess"},
		{"pvlan", "ios-xe", "pvlanHost"},
		{"dot1qTunnel", "ios-xe", "dot1qTunnelHost"},
	}
	for _, tc := range bad {
		byOS, modeOk := ethernetDiscriminatorMatrix[tc.mode]
		if !modeOk {
			// unknown mode — matrix check will emit "Unsupported mode"
			continue
		}
		policies, osOk := byOS[tc.os]
		if !osOk {
			// unknown OS for mode — handled by the (mode, OS) error.
			continue
		}
		if contains(policies, tc.policy) {
			t.Errorf("matrix unexpectedly accepts (%q, %q, %q)",
				tc.mode, tc.os, tc.policy)
		}
	}
}
