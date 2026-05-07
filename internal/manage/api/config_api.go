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

// Config API endpoints
const (
	UrlFabricConfigSave   = "/manage/fabrics/%s/actions/configSave"
	UrlFabricConfigDeploy = "/manage/fabrics/%s/actions/configDeploy"
)

const RscNameConfig = "config"

// ConfigAPI is the API client for fabric config operations (save/deploy)
type ConfigAPI struct {
	ndapi.NexusDashboardAPICommon
	FabricName string
	Operation  ConfigOperation
}

// ConfigOperation defines the type of config operation
type ConfigOperation int

const (
	OpConfigSave ConfigOperation = iota
	OpConfigDeploy
)

// NewConfigAPI creates a new ConfigAPI instance
func NewConfigAPI(client *nd.Client, fabric string) *ConfigAPI {
	papi := new(ConfigAPI)
	papi.Client = client
	papi.Fabric = fabric
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *ConfigAPI) GetUrl() string {
	switch c.Operation {
	case OpConfigSave:
		return fmt.Sprintf(UrlFabricConfigSave, c.FabricName)
	case OpConfigDeploy:
		return fmt.Sprintf(UrlFabricConfigDeploy, c.FabricName)
	default:
		return fmt.Sprintf(UrlFabricConfigSave, c.FabricName)
	}
}

func (c *ConfigAPI) PostUrl() string {
	return c.GetUrl()
}

func (c *ConfigAPI) PutUrl() string {
	return c.GetUrl()
}

func (c *ConfigAPI) DeleteUrl() string {
	return c.GetUrl()
}

func (c *ConfigAPI) GetDeleteQP() []string {
	return nil
}

func (c *ConfigAPI) RscName() string {
	return RscNameConfig
}

// SetOperation sets the current operation
func (c *ConfigAPI) SetOperation(op ConfigOperation) *ConfigAPI {
	c.Operation = op
	return c
}

// WithFabric sets the fabric name
func (c *ConfigAPI) WithFabric(fabricName string) *ConfigAPI {
	c.FabricName = fabricName
	return c
}
