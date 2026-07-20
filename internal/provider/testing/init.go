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
	URL                      string                         `yaml:"url"`
	User                     string                         `yaml:"user"`
	Password                 string                         `yaml:"pwd"`
	Insecure                 string                         `yaml:"insecure"`
	Fabric                   string                         `yaml:"fabric"`
	Switches                 []string                       `yaml:"switches"`
	SwitchIP                 []string                       `yaml:"switch_ip"`
	FabricPrefix             string                         `yaml:"fabric_prefix"`
	IntegrationTest          IntegratedConfig               `yaml:"integration_test"`
	Inventory                InventoryConfig                `yaml:"inventory"`
	VpcPair                  *VpcPairConfig                 `yaml:"vpc_pair"`
	LocalUser                LocalUserConfig                `yaml:"local_user"`
	MultiClusterConnectivity MultiClusterConnectivityConfig `yaml:"multi_cluster_connectivity"`
	RemoteStorage            RemoteStorageConfig            `yaml:"remote_storage"`
}

// MultiClusterConnectivityConfig represents the multi-cluster connectivity test configuration
type MultiClusterConnectivityConfig struct {
	ClusterName             string `yaml:"cluster_name"`
	Hostname                string `yaml:"hostname"`
	LoginDomain             string `yaml:"login_domain"`
	MultiClusterLoginDomain string `yaml:"multi_cluster_login_domain"`
}

// LocalUserConfig represents the local-user-specific test configuration
type LocalUserConfig struct {
	LoginID                 string              `yaml:"login_id"`
	UserPassword            string              `yaml:"user_password"`
	Email                   string              `yaml:"email"`
	FirstName               string              `yaml:"first_name"`
	LastName                string              `yaml:"last_name"`
	RemoteIDClaim           string              `yaml:"remote_id_claim"`
	RemoteUserAuthorization bool                `yaml:"remote_user_authorization"`
	TenantDomain            string              `yaml:"tenant_domain"`
	SecurityDomains         map[string][]string `yaml:"security_domains"`
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

// VpcPairConfig represents the acceptance-test switch selection for vPC pair tests.
type VpcPairConfig struct {
	Switches []string `yaml:"switches"`
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

// RemoteStorageConfig represents the nd_remote_storage_location-specific
// test configuration. Tests use the Hostname/Username/Password/SshKey
// fields when set; otherwise sensible literal defaults from the user's
// testbed are applied inside the test itself.
type RemoteStorageConfig struct {
	Hostname   string `yaml:"hostname"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	SshKey     string `yaml:"ssh_key"`
	Passphrase string `yaml:"passphrase"`
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
