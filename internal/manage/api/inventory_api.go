// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"
	"sync"

	"github.com/netascode/go-nd"
)

// Inventory API endpoints
const (
	// Switch management endpoints
	UrlFabricSwitches     = "/manage/fabrics/%s/switches"
	UrlFabricSwitch       = "/manage/fabrics/%s/switches/%s"
	UrlFabricSwitchRole   = "/manage/fabrics/%s/switchActions/changeRoles"
	UrlFabricSwitchRemove = "/manage/fabrics/%s/switchActions/remove"

	// Discovery endpoints
	UrlShallowDiscovery = "/manage/fabrics/%s/actions/shallowDiscovery"
	UrlDiscoveryStatus  = "/manage/fabrics/%s/inventory/discover"

	// Rediscovery endpoints
	UrlSwitchActionsRediscover = "/manage/fabrics/%s/switchActions/rediscover"

	// Credentials endpoints
	UrlCredentialsSwitches       = "/manage/credentials/switches"
	UrlCredentialsSwitchesRemove = "/manage/credentials/switches/actions/remove"
	UrlCredentialsSwitchValidate = "/manage/credentials/switches/%s/actions/validate"

	// Bootstrap/POAP endpoints
	UrlBootstrapSwitches = "/manage/fabrics/%s/inventory/poap"
)

// InventoryAPI is the API client for inventory/switch resources
type InventoryAPI struct {
	NDManageAPICommon
	mutex        *sync.Mutex
	FabricName   string
	SerialNumber string
	Operation    InventoryOperation
}

// InventoryOperation defines the type of inventory operation
type InventoryOperation int

const (
	OpGetSwitch InventoryOperation = iota
	OpGetAllSwitches
	OpAddSwitches
	OpRemoveSwitches
	OpUpdateSwitchRole
	OpShallowDiscovery
	OpGetDiscoveryStatus
	OpRediscover
	OpBootstrap
	OpCreateCredentials
	OpRemoveCredentials
	OpValidateCredentials
)

func NewInventoryAPI(lock *sync.Mutex, client *nd.Client) *InventoryAPI {
	papi := new(InventoryAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDManageAPI = papi
	return papi
}

func (c *InventoryAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *InventoryAPI) GetUrl() string {
	switch c.Operation {
	case OpGetSwitch:
		return fmt.Sprintf(UrlFabricSwitch, c.FabricName, c.SerialNumber)
	case OpGetAllSwitches:
		return fmt.Sprintf(UrlFabricSwitches, c.FabricName)
	case OpGetDiscoveryStatus:
		return fmt.Sprintf(UrlDiscoveryStatus, c.FabricName)
	case OpValidateCredentials:
		return fmt.Sprintf(UrlCredentialsSwitchValidate, c.SerialNumber)
	default:
		if c.SerialNumber != "" {
			return fmt.Sprintf(UrlFabricSwitch, c.FabricName, c.SerialNumber)
		}
		return fmt.Sprintf(UrlFabricSwitches, c.FabricName)
	}
}

func (c *InventoryAPI) PostUrl() string {
	switch c.Operation {
	case OpAddSwitches:
		return fmt.Sprintf(UrlFabricSwitches, c.FabricName)
	case OpRemoveSwitches:
		return fmt.Sprintf(UrlFabricSwitchRemove, c.FabricName)
	case OpShallowDiscovery:
		return fmt.Sprintf(UrlShallowDiscovery, c.FabricName)
	case OpRediscover:
		return fmt.Sprintf(UrlSwitchActionsRediscover, c.FabricName)
	case OpBootstrap:
		return fmt.Sprintf(UrlBootstrapSwitches, c.FabricName)
	case OpCreateCredentials:
		return UrlCredentialsSwitches
	case OpRemoveCredentials:
		return UrlCredentialsSwitchesRemove
	case OpValidateCredentials:
		return fmt.Sprintf(UrlCredentialsSwitchValidate, c.SerialNumber)
	case OpUpdateSwitchRole:
		return fmt.Sprintf(UrlFabricSwitchRole, c.FabricName)
	default:
		return fmt.Sprintf(UrlFabricSwitches, c.FabricName)
	}
}

func (c *InventoryAPI) PutUrl() string {
	switch c.Operation {
	case OpUpdateSwitchRole:
		return fmt.Sprintf(UrlFabricSwitchRole, c.FabricName)
	default:
		return fmt.Sprintf(UrlFabricSwitch, c.FabricName, c.SerialNumber)
	}
}

func (c *InventoryAPI) DeleteUrl() string {
	if c.SerialNumber != "" {
		return fmt.Sprintf(UrlFabricSwitch, c.FabricName, c.SerialNumber)
	}
	return fmt.Sprintf(UrlFabricSwitches, c.FabricName)
}

func (c *InventoryAPI) GetDeleteQP() []string {
	return nil
}

func (c *InventoryAPI) RscName() string {
	return "inventory_switch"
}

// SetOperation sets the current operation for URL generation
func (c *InventoryAPI) SetOperation(op InventoryOperation) *InventoryAPI {
	c.Operation = op
	return c
}

// WithFabric sets the fabric name
func (c *InventoryAPI) WithFabric(fabricName string) *InventoryAPI {
	c.FabricName = fabricName
	return c
}

// WithSerial sets the serial number
func (c *InventoryAPI) WithSerial(serialNumber string) *InventoryAPI {
	c.SerialNumber = serialNumber
	return c
}
