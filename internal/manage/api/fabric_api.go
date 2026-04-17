// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
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

// FabricAPI is the API client for fabric resources
const (
	// Standard Fabric API endpoints
	UrlFabrics      = "/manage/fabrics"
	UrlFabricByName = "/manage/fabrics/%s"

	// Legacy endpoints - keeping for compatibility
	UrlFabricGetAll         = "/lan-fabric/rest/control/fabrics"
	urlPerFabric            = UrlFabricGetAll + "/%s"
	urlFabricTemplate       = urlPerFabric + "/%s"
	UrlSwitchesByFabric     = "/lan-fabric/rest/control/fabrics/%s/inventory/switchesByFabric"
	urlFabricNameFromSerial = "/lan-fabric/rest/control/switches/%s/fabric-name"
)

type FabricAPI struct {
	NDManageAPICommon
	mutex               *sync.Mutex
	FabricName          string
	GetSwitchesInFabric bool
	Serialnumber        string
	FabricType          string
}

func NewFabricAPI(lock *sync.Mutex, client *nd.Client) *FabricAPI {
	papi := new(FabricAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDManageAPI = papi
	return papi
}

func (c *FabricAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *FabricAPI) GetUrl() string {
	// Use the modern API endpoint first if fabric name is specified
	if c.FabricName != "" {
		return fmt.Sprintf(UrlFabricByName, c.FabricName)
	}

	// Fall back to legacy endpoints for special cases
	if c.Serialnumber != "" {
		return fmt.Sprintf(urlFabricNameFromSerial, c.Serialnumber)
	}
	// Use modern API endpoint for listing all fabrics
	return UrlFabrics
}

func (c *FabricAPI) PostUrl() string {
	// Use modern API endpoint for creating fabrics
	return UrlFabrics
}

func (c *FabricAPI) PutUrl() string {
	// Use modern API endpoint for updating fabric
	if c.FabricName != "" {
		return fmt.Sprintf(UrlFabricByName, c.FabricName)
	}
	// Fall back to legacy endpoint if needed
	return fmt.Sprintf(urlFabricTemplate, c.FabricName, c.FabricType)
}

func (c *FabricAPI) DeleteUrl() string {
	// Use modern API endpoint for deleting fabric
	if c.FabricName != "" {
		return fmt.Sprintf(UrlFabricByName, c.FabricName)
	}
	// Fall back to legacy endpoint if needed
	return fmt.Sprintf(urlPerFabric, c.FabricName)
}

func (c *FabricAPI) GetDeleteQP() []string {
	return nil
}

func (c *FabricAPI) RscName() string {
	return "fabric"
}
