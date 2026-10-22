// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_inventory_switch

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestInventorySwitchPayload(t *testing.T) {
	payload, err := os.ReadFile("testdata/inventory_switches.json")
	if err != nil {
		t.Fatalf("read inventory switch fixture: %v", err)
	}

	var response inventorySwitchesResponse
	if err := json.Unmarshal(payload, &response); err != nil {
		t.Fatalf("decode inventory switch fixture: %v", err)
	}
	if got, want := len(response.Switches), 2; got != want {
		t.Fatalf("decoded switch count = %d, want %d", got, want)
	}

	const (
		leafSerial  = "9R518K2AT3R"
		leafIP      = "192.168.10.204"
		spineSerial = "9N2SFZYSDIJ"
		spineIP     = "192.168.10.201"
	)

	tests := []struct {
		name         string
		serialNumber string
		managementIP string
		wantSerials  []string
	}{
		{name: "fabric_only", wantSerials: []string{leafSerial, spineSerial}},
		{name: "serial_only", serialNumber: leafSerial, wantSerials: []string{leafSerial}},
		{name: "management_ip_only", managementIP: spineIP, wantSerials: []string{spineSerial}},
		{name: "serial_and_management_ip", serialNumber: leafSerial, managementIP: leafIP, wantSerials: []string{leafSerial}},
		{name: "serial_and_mismatched_management_ip", serialNumber: leafSerial, managementIP: spineIP},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filtered := filterInventorySwitches(response.Switches, test.serialNumber, test.managementIP)
			if got, want := len(filtered), len(test.wantSerials); got != want {
				t.Fatalf("filtered switch count = %d, want %d", got, want)
			}

			list, diags := inventorySwitchesList(filtered)
			if diags.HasError() {
				t.Fatalf("convert switches to Terraform list: %v", diags)
			}
			values := list.Elements()
			if got, want := len(values), len(test.wantSerials); got != want {
				t.Fatalf("Terraform list element count = %d, want %d", got, want)
			}

			for index, value := range values {
				object, ok := value.(types.Object)
				if !ok {
					t.Fatalf("list element %d has type %T, want types.Object", index, value)
				}
				attributes := object.Attributes()
				assertStringAttribute(t, attributes, "serial_number", test.wantSerials[index])
				if test.wantSerials[index] == leafSerial {
					assertStringAttribute(t, attributes, "ip_address", leafIP)
					assertStringAttribute(t, attributes, "hostname", "leaf3")
					assertStringAttribute(t, attributes, "status", "ok")
				} else {
					assertStringAttribute(t, attributes, "ip_address", spineIP)
					assertStringAttribute(t, attributes, "hostname", "spine2")
				}
				if _, exists := attributes["username"]; exists {
					t.Fatal("sensitive additionalData.username must not be exposed in Terraform state")
				}
			}
		})
	}
}

func assertStringAttribute(t *testing.T, attributes map[string]attr.Value, name, want string) {
	t.Helper()

	value, ok := attributes[name].(types.String)
	if !ok {
		t.Fatalf("attribute %q has type %T, want types.String", name, attributes[name])
	}
	if got := value.ValueString(); got != want {
		t.Fatalf("attribute %q = %q, want %q", name, got, want)
	}
}
