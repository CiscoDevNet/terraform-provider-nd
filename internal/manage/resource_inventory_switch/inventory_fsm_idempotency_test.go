// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
// SPDX-License-Identifier: MPL-2.0

package resource_inventory_switch

import "testing"

func TestAlreadyPresentInFabricErrorIsScoped(t *testing.T) {
	message := `HTTP 400: Switch SERIAL1 already present in fabric vxlan_akabhask. Please remove the switch(es) from existing fabric(s)`
	if !isAlreadyPresentInFabricError(message, "vxlan_akabhask") {
		t.Fatal("expected exact same-fabric response to be accepted")
	}
	if isAlreadyPresentInFabricError(message, "another_fabric") {
		t.Fatal("response for a different fabric must not be accepted")
	}
}

func TestAllDiscoveredSwitchesPresent(t *testing.T) {
	existing := &FabricSwitchesResponse{
		SwitchesBySerial: map[string]*FabricSwitchEntry{"SERIAL1": {}},
		SwitchesByIP:     map[string]*FabricSwitchEntry{"192.0.2.2": {}},
	}
	if !allDiscoveredSwitchesPresent(&DiscoveryStatusResponse{Switches: []NDFCSwitchDetailValue{
		{SerialNumber: "SERIAL1"},
		{IpAddress: "192.0.2.2"},
	}}, existing) {
		t.Fatal("expected all discovered switches to be found")
	}
	if allDiscoveredSwitchesPresent(&DiscoveryStatusResponse{Switches: []NDFCSwitchDetailValue{
		{SerialNumber: "MISSING"},
	}}, existing) {
		t.Fatal("missing serial must not be treated as converged")
	}
}
