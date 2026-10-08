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
	"fmt"
	"sync"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	nd "github.com/netascode/go-nd"
)

// SwitchEntry represents essential switch information cached from the inventory API.
type SwitchEntry struct {
	SerialNumber       string
	IPAddress          string
	FabricManagementIp string
	Hostname           string
	SwitchRole         string
	Model              string
}

// switchInventoryResponse matches the GET /manage/fabrics/{fabric}/switches response.
type switchInventoryResponse struct {
	Switches []switchInventoryItem `json:"switches,omitempty"`
}

// switchInventoryItem represents a single switch in the inventory API response.
type switchInventoryItem struct {
	SerialNumber       string `json:"serialNumber,omitempty"`
	IPAddress          string `json:"ip,omitempty"`
	FabricManagementIp string `json:"fabricManagementIp,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	SwitchRole         string `json:"switchRole,omitempty"`
	Model              string `json:"model,omitempty"`
}

// SwitchDB is a thread-safe, lazily-populated cache of switch information
// indexed by fabric name and IP address. Serial lookups scan the fabric map.
//
// The cache is populated on first access per fabric via the inventory API
// (GET /manage/fabrics/{fabric}/switches). The write lock is held during
// the API call to prevent duplicate concurrent loads for the same fabric.
//
// Invalidation: callers must call ClearFabric after inventory mutations
// (switch add/remove/update) so subsequent lookups see fresh data.
type SwitchDB struct {
	mu     sync.RWMutex
	byIP   map[string]map[string]*SwitchEntry // fabric → IP → entry
	client *nd.Client
}

// NewSwitchDB creates a new SwitchDB backed by the given API client.
func NewSwitchDB(client *nd.Client) *SwitchDB {
	return &SwitchDB{
		byIP:   make(map[string]map[string]*SwitchEntry),
		client: client,
	}
}

// GetSerialByIP returns the serial number for a switch identified by IP address.
// Auto-loads the fabric's switch data on first access.
func (db *SwitchDB) GetSerialByIP(ctx context.Context, fabricName, ipAddress string) (string, bool) {
	entry, ok := db.GetSwitchByIP(ctx, fabricName, ipAddress)
	if !ok {
		return "", false
	}
	return entry.SerialNumber, true
}

// GetSwitchByIP returns the full SwitchEntry for a switch identified by IP address.
// Auto-loads the fabric's switch data on first access.
func (db *SwitchDB) GetSwitchByIP(ctx context.Context, fabricName, ipAddress string) (*SwitchEntry, bool) {
	return db.lookup(ctx, fabricName, func(fabricMap map[string]*SwitchEntry) (*SwitchEntry, bool) {
		entry, ok := fabricMap[ipAddress]
		return entry, ok
	})
}

// GetIPBySerial returns the IP address for a switch identified by serial number.
// Auto-loads the fabric's switch data on first access. O(n) scan over the fabric's switches.
func (db *SwitchDB) GetIPBySerial(ctx context.Context, fabricName, serialNumber string) (string, bool) {
	entry, ok := db.GetSwitchBySerial(ctx, fabricName, serialNumber)
	if !ok {
		return "", false
	}
	ip := entry.FabricManagementIp
	if ip == "" {
		ip = entry.IPAddress
	}
	return ip, ip != ""
}

// GetSwitchBySerial returns the full SwitchEntry for a switch identified by serial number.
// Auto-loads the fabric's switch data on first access. O(n) scan over the fabric's switches.
func (db *SwitchDB) GetSwitchBySerial(ctx context.Context, fabricName, serialNumber string) (*SwitchEntry, bool) {
	return db.lookup(ctx, fabricName, func(fabricMap map[string]*SwitchEntry) (*SwitchEntry, bool) {
		for _, entry := range fabricMap {
			if entry.SerialNumber == serialNumber {
				return entry, true
			}
		}
		return nil, false
	})
}

// lookup searches a fabric's switch map using the provided match function.
// If the fabric is not yet loaded, it triggers a lazy load and retries once.
func (db *SwitchDB) lookup(ctx context.Context, fabricName string, match func(map[string]*SwitchEntry) (*SwitchEntry, bool)) (*SwitchEntry, bool) {
	// Fast path: read lock
	db.mu.RLock()
	if fabricMap, ok := db.byIP[fabricName]; ok {
		entry, found := match(fabricMap)
		db.mu.RUnlock()
		return entry, found
	}
	db.mu.RUnlock()

	// Fabric not loaded — load under write lock
	if err := db.ensureLoaded(ctx, fabricName); err != nil {
		tflog.Warn(ctx, "SwitchDB: failed to load fabric", map[string]interface{}{
			"fabric": fabricName,
			"error":  err.Error(),
		})
		return nil, false
	}

	// Retry after load
	db.mu.RLock()
	defer db.mu.RUnlock()
	if fabricMap, ok := db.byIP[fabricName]; ok {
		return match(fabricMap)
	}
	return nil, false
}

// ClearFabric removes cached switch data for a fabric, forcing a reload on next access.
// Call this after inventory mutations (switch add, remove, or update).
func (db *SwitchDB) ClearFabric(fabricName string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.byIP, fabricName)
}

// IsFabricLoaded returns true if the fabric's switch data has been loaded.
func (db *SwitchDB) IsFabricLoaded(fabricName string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()
	_, ok := db.byIP[fabricName]
	return ok
}

// ensureLoaded fetches the fabric's switch inventory if not already cached.
// Holds the write lock during the API call to prevent duplicate concurrent loads.
func (db *SwitchDB) ensureLoaded(ctx context.Context, fabricName string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Double-check under write lock (another goroutine may have loaded while we waited)
	if _, ok := db.byIP[fabricName]; ok {
		return nil
	}

	if db.client == nil {
		return fmt.Errorf("SwitchDB: no API client configured")
	}

	tflog.Debug(ctx, "SwitchDB: loading fabric inventory", map[string]interface{}{
		"fabric": fabricName,
	})

	jsonData, err := db.fetchSwitches(ctx, fabricName)
	if err != nil {
		return fmt.Errorf("SwitchDB: failed to fetch switches for fabric %s: %w", fabricName, err)
	}

	var resp switchInventoryResponse
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return fmt.Errorf("SwitchDB: failed to parse switches response for fabric %s: %w", fabricName, err)
	}

	fabricMap := make(map[string]*SwitchEntry, len(resp.Switches)*2)
	for _, sw := range resp.Switches {
		if sw.FabricManagementIp == "" && sw.IPAddress == "" {
			continue // skip entries without any IP
		}

		entry := &SwitchEntry{
			SerialNumber:       sw.SerialNumber,
			IPAddress:          sw.IPAddress,
			FabricManagementIp: sw.FabricManagementIp,
			Hostname:           sw.Hostname,
			SwitchRole:         sw.SwitchRole,
			Model:              sw.Model,
		}

		// Index under both IP addresses so lookups work regardless of
		// which address the user specifies. Both point to the same entry.
		if sw.FabricManagementIp != "" {
			fabricMap[sw.FabricManagementIp] = entry
		}
		if sw.IPAddress != "" && sw.IPAddress != sw.FabricManagementIp {
			fabricMap[sw.IPAddress] = entry
		}
	}

	db.byIP[fabricName] = fabricMap

	tflog.Debug(ctx, "SwitchDB: loaded fabric inventory", map[string]interface{}{
		"fabric":   fabricName,
		"switches": len(fabricMap),
	})

	return nil
}

// fetchSwitches calls the inventory API to get all switches for a fabric.
// Uses CRUD lock scoping to participate in the provider's lock hierarchy.
func (db *SwitchDB) fetchSwitches(ctx context.Context, fabricName string) ([]byte, error) {
	invAPI := api.NewInventoryAPI(db.client, ndapi.DefaultFabric)
	invAPI.FabricName = fabricName
	invAPI.SetOperation(api.OpGetAllSwitches)

	respData, err := invAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("inventory API error: %w", err)
	}

	return respData, nil
}
