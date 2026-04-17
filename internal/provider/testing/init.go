// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config represents the top-level testbed configuration
type Config struct {
	ND NDConfig `yaml:"nd"`
}

// NDConfig represents the ND controller configuration
type NDConfig struct {
	URL             string           `yaml:"url"`
	User            string           `yaml:"user"`
	Password        string           `yaml:"pwd"`
	Insecure        string           `yaml:"insecure"`
	Fabric          string           `yaml:"fabric"`
	Switches        []string         `yaml:"switches"`
	SwitchIP        []string         `yaml:"switch_ip"`
	FabricPrefix    string           `yaml:"fabric_prefix"`
	IntegrationTest IntegratedConfig `yaml:"integration_test"`
	Inventory       InventoryConfig  `yaml:"inventory"`
}

// InventoryConfig represents the inventory-specific test configuration
type InventoryConfig struct {
	Fabric   string            `yaml:"fabric"`
	User     string            `yaml:"user"`
	Password string            `yaml:"pwd"`
	Switches []InventorySwitch `yaml:"switches"`
	Deploy   bool              `yaml:"deploy"`
	Recalc   bool              `yaml:"recalculate"`
	Preserve bool              `yaml:"preserve_config"`
	Mode     string            `yaml:"mode"`
	MaxHop   int               `yaml:"max_hop"`
}

// InventorySwitch represents a switch in the inventory config
type InventorySwitch struct {
	Serial string `yaml:"serial"`
	IP     string `yaml:"ip"`
	Role   string `yaml:"role"`
}

// IntegratedConfig represents the integration test configuration
type IntegratedConfig struct {
	Fabric           string            `yaml:"fabric"`
	User             string            `yaml:"user"`
	Password         string            `yaml:"pwd"`
	Switches         []string          `yaml:"switches"`
	InventoryDevices []InventoryDevice `yaml:"inventory_devices"`
}

// InventoryDevice represents a device in the inventory
type InventoryDevice struct {
	Device string `yaml:"device"`
	Role   string `yaml:"role"`
}

var (
	globalConfig *Config
	configMu     sync.RWMutex
	isMocked     bool
	mockConfigs  map[string]*Config
)

// InitConfig loads the testbed YAML configuration and optionally sets up mock mode
func InitConfig(configPath string, mockedServer string) {
	configMu.Lock()
	defer configMu.Unlock()

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read test config file %s: %v", configPath, err))
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse test config file %s: %v", configPath, err))
	}

	globalConfig = cfg

	if mockedServer != "" {
		isMocked = true
		mockConfigs = make(map[string]*Config)
		log.Printf("Mock server mode enabled")
	}

	log.Printf("Test config loaded: url=%s fabric=%s switches=%d inventory_switches=%d",
		cfg.ND.URL, cfg.ND.Fabric, len(cfg.ND.Switches), len(cfg.ND.Inventory.Switches))
}

// GetConfig returns the configuration for a given module.
// When mocked, each module may get a unique config with mock port.
// When not mocked, all modules share the global config.
func GetConfig(module string) *Config {
	configMu.RLock()
	defer configMu.RUnlock()

	if globalConfig == nil {
		panic("Test config not initialized. Call InitConfig() first.")
	}

	if isMocked && mockConfigs != nil {
		if cfg, ok := mockConfigs[module]; ok {
			return cfg
		}
	}

	return globalConfig
}

// IsMocked returns whether mock mode is active
func IsMocked() bool {
	return isMocked
}

// StartMockServer starts a mock server for the given module (stub for future implementation)
func StartMockServer(module string) {
	log.Printf("Mock server for module %s: not yet implemented", module)
}

// StopMock stops all mock servers (stub for future implementation)
func StopMock() {
	if isMocked {
		log.Printf("Stopping mock servers")
	}
}
