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
	"strings"

	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

// Config API endpoints
const (
	UrlFabricConfigSave    = "/manage/fabrics/%s/actions/configSave"
	UrlFabricDeploy        = "/manage/fabrics/%s/actions/deploy"
	UrlSwitchDeploy        = "/manage/fabrics/%s/switchActions/deploy"
	UrlFabricPreview       = "/manage/fabrics/%s/actions/preview"
	UrlSwitchPreview       = "/manage/fabrics/%s/switchActions/preview"
	UrlFabricDeployHistory = "/manage/fabrics/%s/deploymentHistory"
)

const RscNameConfig = "config"

// ConfigAPI is the API client for fabric config operations (save/deploy)
type ConfigAPI struct {
	ndapi.NexusDashboardAPICommon
	FabricName  string
	Operation   ConfigOperation
	QueryParams []string
}

// ConfigOperation defines the type of config operation
type ConfigOperation int

const (
	OpConfigSave ConfigOperation = iota
	OpFabricDeploy
	OpSwitchDeploy
	OpFabricPreview
	OpSwitchPreview
	OpDeployHistory
)

// NewConfigAPI creates a new ConfigAPI instance
func NewConfigAPI(client *nd.Client, fabric string) *ConfigAPI {
	papi := new(ConfigAPI)
	papi.Client = client
	papi.Fabric = fabric
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *ConfigAPI) urlForOp() string {
	switch c.Operation {
	case OpConfigSave:
		return fmt.Sprintf(UrlFabricConfigSave, c.FabricName)
	case OpFabricDeploy:
		return fmt.Sprintf(UrlFabricDeploy, c.FabricName)
	case OpSwitchDeploy:
		return fmt.Sprintf(UrlSwitchDeploy, c.FabricName)
	case OpFabricPreview:
		return fmt.Sprintf(UrlFabricPreview, c.FabricName)
	case OpSwitchPreview:
		return fmt.Sprintf(UrlSwitchPreview, c.FabricName)
	case OpDeployHistory:
		return fmt.Sprintf(UrlFabricDeployHistory, c.FabricName)
	default:
		return fmt.Sprintf(UrlFabricConfigSave, c.FabricName)
	}
}

func (c *ConfigAPI) appendQueryParams(url string) string {
	if len(c.QueryParams) > 0 {
		return url + "?" + strings.Join(c.QueryParams, "&")
	}
	return url
}

func (c *ConfigAPI) GetUrl() string {
	return c.appendQueryParams(c.urlForOp())
}

func (c *ConfigAPI) PostUrl() string {
	return c.appendQueryParams(c.urlForOp())
}

func (c *ConfigAPI) PutUrl() string {
	return c.urlForOp()
}

func (c *ConfigAPI) DeleteUrl() string {
	return c.urlForOp()
}

func (c *ConfigAPI) GetDeleteQP() []string {
	return nil
}

// SetQueryParams sets query parameters for the API call
func (c *ConfigAPI) SetQueryParams(params ...string) *ConfigAPI {
	c.QueryParams = params
	return c
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
