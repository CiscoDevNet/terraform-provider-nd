// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import "testing"

func TestInterfaceAPIPathEscapesInterfaceName(t *testing.T) {
	interfaceAPI := &InterfaceAPI{
		FabricName:    "fabric-1",
		SwitchId:      "FDO22412H7U",
		InterfaceName: "Ethernet1/10",
		Operation:     OpInterfaceGetOne,
	}
	want := "/manage/fabrics/fabric-1/switches/FDO22412H7U/interfaces/Ethernet1%2F10"

	if got := interfaceAPI.GetUrl(); got != want {
		t.Fatalf("GetUrl() = %q, want %q", got, want)
	}
	if got := interfaceAPI.PutUrl(); got != want {
		t.Fatalf("PutUrl() = %q, want %q", got, want)
	}
	if got := interfaceAPI.DeleteUrl(); got != want {
		t.Fatalf("DeleteUrl() = %q, want %q", got, want)
	}
}
