// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package manage

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
)

// newTestSwitchDB creates a SwitchDB with pre-populated data (no API client needed).
// Mirrors the dual-indexing logic in ensureLoaded: both FabricManagementIp and
// IPAddress are indexed when they differ.
func newTestSwitchDB(fabricName string, entries []*SwitchEntry) *SwitchDB {
	db := &SwitchDB{
		byIP: make(map[string]map[string]*SwitchEntry),
	}
	fabricMap := make(map[string]*SwitchEntry, len(entries)*2)
	for _, e := range entries {
		if e.FabricManagementIp == "" && e.IPAddress == "" {
			continue
		}
		if e.FabricManagementIp != "" {
			fabricMap[e.FabricManagementIp] = e
		}
		if e.IPAddress != "" && e.IPAddress != e.FabricManagementIp {
			fabricMap[e.IPAddress] = e
		}
	}
	db.byIP[fabricName] = fabricMap
	return db
}

func TestGetSwitchByIP(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1", Hostname: "spine1", SwitchRole: "spine"},
		{SerialNumber: "SER002", IPAddress: "10.1.1.2", FabricManagementIp: "10.1.1.2", Hostname: "leaf1", SwitchRole: "leaf"},
	})
	ctx := context.Background()

	// Hit
	entry, ok := db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1")
	if !ok || entry.SerialNumber != "SER001" {
		t.Errorf("expected SER001, got %v (ok=%v)", entry, ok)
	}

	// Miss (wrong IP)
	_, ok = db.GetSwitchByIP(ctx, "fabric1", "10.1.1.99")
	if ok {
		t.Error("expected miss for unknown IP")
	}

	// Miss (wrong fabric)
	_, ok = db.GetSwitchByIP(ctx, "fabric2", "10.1.1.1")
	if ok {
		t.Error("expected miss for unknown fabric")
	}
}

func TestGetSerialByIP(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1"},
	})
	ctx := context.Background()

	serial, ok := db.GetSerialByIP(ctx, "fabric1", "10.1.1.1")
	if !ok || serial != "SER001" {
		t.Errorf("expected SER001, got %q (ok=%v)", serial, ok)
	}

	serial, ok = db.GetSerialByIP(ctx, "fabric1", "10.1.1.99")
	if ok {
		t.Errorf("expected miss, got %q", serial)
	}
}

func TestGetSwitchBySerial(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1", Hostname: "spine1"},
		{SerialNumber: "SER002", IPAddress: "10.1.1.2", FabricManagementIp: "10.1.1.2", Hostname: "leaf1"},
	})
	ctx := context.Background()

	entry, ok := db.GetSwitchBySerial(ctx, "fabric1", "SER002")
	if !ok || entry.Hostname != "leaf1" {
		t.Errorf("expected leaf1, got %v (ok=%v)", entry, ok)
	}

	_, ok = db.GetSwitchBySerial(ctx, "fabric1", "UNKNOWN")
	if ok {
		t.Error("expected miss for unknown serial")
	}
}

func TestGetIPBySerial(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.2.2.1"},
	})
	ctx := context.Background()

	// FabricManagementIp takes priority
	ip, ok := db.GetIPBySerial(ctx, "fabric1", "SER001")
	if !ok || ip != "10.2.2.1" {
		t.Errorf("expected 10.2.2.1, got %q (ok=%v)", ip, ok)
	}

	// Fallback to IPAddress when FabricManagementIp is empty
	db2 := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: ""},
	})
	ip, ok = db2.GetIPBySerial(ctx, "fabric1", "SER001")
	if !ok || ip != "10.1.1.1" {
		t.Errorf("expected 10.1.1.1 fallback, got %q (ok=%v)", ip, ok)
	}
}

func TestClearFabric(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1"},
	})
	ctx := context.Background()

	// Verify present
	if _, ok := db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1"); !ok {
		t.Fatal("expected switch before clear")
	}

	db.ClearFabric("fabric1")

	// Fabric is gone — no loader, so lookup returns false
	if _, ok := db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1"); ok {
		t.Error("expected miss after ClearFabric")
	}
}

func TestIsFabricLoaded(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{})

	if !db.IsFabricLoaded("fabric1") {
		t.Error("expected fabric1 to be loaded")
	}
	if db.IsFabricLoaded("fabric2") {
		t.Error("expected fabric2 to not be loaded")
	}
}

func TestLoadFromJSON(t *testing.T) {
	switches := []switchInventoryItem{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.2.2.1", Hostname: "spine1", SwitchRole: "spine", Model: "N9K"},
		{SerialNumber: "SER002", IPAddress: "", FabricManagementIp: "10.2.2.2", Hostname: "leaf1", SwitchRole: "leaf", Model: "N9K"},
		{SerialNumber: "SER003", IPAddress: "", FabricManagementIp: "", Hostname: "orphan", SwitchRole: "leaf", Model: "N9K"}, // no IP, should be skipped
	}

	resp := switchInventoryResponse{Switches: switches}
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}

	db := &SwitchDB{
		byIP: make(map[string]map[string]*SwitchEntry),
	}
	ctx := context.Background()

	// Simulate what ensureLoaded does: unmarshal and populate
	var parsed switchInventoryResponse
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatal(err)
	}

	fabricMap := make(map[string]*SwitchEntry, len(parsed.Switches)*2)
	for _, sw := range parsed.Switches {
		if sw.FabricManagementIp == "" && sw.IPAddress == "" {
			continue
		}
		entry := &SwitchEntry{
			SerialNumber:       sw.SerialNumber,
			IPAddress:          sw.IPAddress,
			FabricManagementIp: sw.FabricManagementIp,
			Hostname:           sw.Hostname,
			SwitchRole:         sw.SwitchRole,
			Model:              sw.Model,
		}
		if sw.FabricManagementIp != "" {
			fabricMap[sw.FabricManagementIp] = entry
		}
		if sw.IPAddress != "" && sw.IPAddress != sw.FabricManagementIp {
			fabricMap[sw.IPAddress] = entry
		}
	}
	db.byIP["testfabric"] = fabricMap

	// SER001 indexed by FabricManagementIp
	entry, ok := db.GetSwitchByIP(ctx, "testfabric", "10.2.2.1")
	if !ok || entry.SerialNumber != "SER001" {
		t.Errorf("expected SER001 at 10.2.2.1, got %v", entry)
	}

	// SER001 also reachable by IPAddress
	entry, ok = db.GetSwitchByIP(ctx, "testfabric", "10.1.1.1")
	if !ok || entry.SerialNumber != "SER001" {
		t.Errorf("expected SER001 at 10.1.1.1, got %v", entry)
	}

	// SER002 indexed by FabricManagementIp (IPAddress is empty)
	entry, ok = db.GetSwitchByIP(ctx, "testfabric", "10.2.2.2")
	if !ok || entry.SerialNumber != "SER002" {
		t.Errorf("expected SER002 at 10.2.2.2, got %v", entry)
	}

	// SER003 should be skipped (no IP)
	// SER001 has 2 entries (both IPs), SER002 has 1 entry → total 3
	if len(db.byIP["testfabric"]) != 3 {
		t.Errorf("expected 3 entries (SER001 x2 + SER002 x1), got %d", len(db.byIP["testfabric"]))
	}
}

func TestConcurrentReads(t *testing.T) {
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1"},
		{SerialNumber: "SER002", IPAddress: "10.1.1.2", FabricManagementIp: "10.1.1.2"},
	})
	ctx := context.Background()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1")
		}()
		go func() {
			defer wg.Done()
			db.GetSwitchBySerial(ctx, "fabric1", "SER002")
		}()
	}
	wg.Wait()
}

func TestDualIPIndexing(t *testing.T) {
	// When both FabricManagementIp and IPAddress are set and differ,
	// the entry is indexed under both addresses.
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.2.2.1"},
	})
	ctx := context.Background()

	// Lookup by FabricManagementIp
	entry, ok := db.GetSwitchByIP(ctx, "fabric1", "10.2.2.1")
	if !ok || entry.SerialNumber != "SER001" {
		t.Error("expected hit by FabricManagementIp")
	}

	// Lookup by IPAddress — also works, same entry
	entry, ok = db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1")
	if !ok || entry.SerialNumber != "SER001" {
		t.Error("expected hit by IPAddress")
	}

	// Both point to the same SwitchEntry
	e1, _ := db.GetSwitchByIP(ctx, "fabric1", "10.2.2.1")
	e2, _ := db.GetSwitchByIP(ctx, "fabric1", "10.1.1.1")
	if e1 != e2 {
		t.Error("expected both IPs to point to the same SwitchEntry")
	}
}

func TestSameIPNotDuplicated(t *testing.T) {
	// When FabricManagementIp == IPAddress, only one map entry is created.
	db := newTestSwitchDB("fabric1", []*SwitchEntry{
		{SerialNumber: "SER001", IPAddress: "10.1.1.1", FabricManagementIp: "10.1.1.1"},
	})

	if len(db.byIP["fabric1"]) != 1 {
		t.Errorf("expected 1 map entry when IPs are identical, got %d", len(db.byIP["fabric1"]))
	}
}
