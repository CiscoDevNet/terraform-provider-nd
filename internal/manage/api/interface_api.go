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

	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

// Interfaces API endpoints (relative to /manage path exposed by go-nd).
const (
	UrlFabricSwitchInterfaces     = "/manage/fabrics/%s/switches/%s/interfaces"
	UrlFabricSwitchInterface      = "/manage/fabrics/%s/switches/%s/interfaces/%s"
	UrlFabricInterfaceActDeploy   = "/manage/fabrics/%s/interfaceActions/deploy"
	UrlFabricInterfaceActRemove   = "/manage/fabrics/%s/interfaceActions/remove"
	UrlFabricInterfaceActPreview  = "/manage/fabrics/%s/interfaceActions/preview"
	UrlFabricInterfaceActSummary  = "/manage/fabrics/%s/interfacesSummary"
	UrlFabricInterfaceActOverview = "/manage/fabrics/%s/interfacesOverview"
)

// RscNameInterface is the per-fabric resource mutex key for interface CRUD.
// All interface resources (ethernet, port_channel, vpc, ...) share this one
// name so that they serialise within a fabric.
const RscNameInterface = "interface"

// InterfaceOperation discriminates the URL/method surface the InterfaceAPI
// client should hit. See the dispatch functions below.
type InterfaceOperation int

const (
	OpInterfaceGetOne InterfaceOperation = iota
	OpInterfaceListPerSwitch
	OpInterfaceCreatePerSwitch // POST .../switches/{sw}/interfaces (bulk)
	OpInterfaceUpdateOne       // PUT  .../switches/{sw}/interfaces/{name}
	OpInterfaceDeleteOne       // DELETE .../switches/{sw}/interfaces/{name}
	OpInterfaceRemoveBulk      // POST .../interfaceActions/remove
	OpInterfaceDeploy          // POST .../interfaceActions/deploy
	OpInterfaceSummary         // GET  .../interfacesSummary?switchId={sw}
)

// InterfaceAPI is the API client for interface resources.
// Fabric is set on NexusDashboardAPICommon so the per-fabric lock scope
// isolates CRUD on fabric A from a deploy on fabric B.
type InterfaceAPI struct {
	ndapi.NexusDashboardAPICommon
	FabricName    string
	SwitchId      string
	InterfaceName string
	// ConfigOnly, if true, adds ?configOnly=true to GET URLs. We use this
	// on Read so operational churn (speed/MTU reconciliation) doesn't drive
	// state diffs.
	ConfigOnly bool
	Operation  InterfaceOperation
}

// NewInterfaceAPI constructs an InterfaceAPI scoped to a fabric. The caller
// must set FabricName (used in URLs) and Operation before dispatch.
func NewInterfaceAPI(client *nd.Client, fabric string) *InterfaceAPI {
	papi := new(InterfaceAPI)
	papi.Client = client
	papi.Fabric = fabric
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *InterfaceAPI) GetUrl() string {
	switch c.Operation {
	case OpInterfaceGetOne:
		return fmt.Sprintf(UrlFabricSwitchInterface, c.FabricName, c.SwitchId, c.InterfaceName)
	case OpInterfaceListPerSwitch:
		return fmt.Sprintf(UrlFabricSwitchInterfaces, c.FabricName, c.SwitchId)
	case OpInterfaceSummary:
		return fmt.Sprintf(UrlFabricInterfaceActSummary+"?switchId=%s", c.FabricName, c.SwitchId)
	default:
		return fmt.Sprintf(UrlFabricSwitchInterface, c.FabricName, c.SwitchId, c.InterfaceName)
	}
}

func (c *InterfaceAPI) PostUrl() string {
	switch c.Operation {
	case OpInterfaceCreatePerSwitch:
		return fmt.Sprintf(UrlFabricSwitchInterfaces, c.FabricName, c.SwitchId)
	case OpInterfaceRemoveBulk:
		return fmt.Sprintf(UrlFabricInterfaceActRemove, c.FabricName)
	case OpInterfaceDeploy:
		return fmt.Sprintf(UrlFabricInterfaceActDeploy, c.FabricName)
	default:
		return fmt.Sprintf(UrlFabricSwitchInterfaces, c.FabricName, c.SwitchId)
	}
}

func (c *InterfaceAPI) PutUrl() string {
	return fmt.Sprintf(UrlFabricSwitchInterface, c.FabricName, c.SwitchId, c.InterfaceName)
}

func (c *InterfaceAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlFabricSwitchInterface, c.FabricName, c.SwitchId, c.InterfaceName)
}

func (c *InterfaceAPI) GetQueryParams() []string {
	if c.ConfigOnly {
		return []string{"configOnly=true"}
	}
	return nil
}

func (c *InterfaceAPI) RscName() string {
	return RscNameInterface
}
