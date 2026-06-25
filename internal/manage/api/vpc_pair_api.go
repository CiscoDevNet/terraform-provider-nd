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
	"terraform-provider-nd/internal/common/ndapi"

	"github.com/netascode/go-nd"
)

// VpcPairAPI is the API client for the vpc pair resource

const urlVpcPair = "/manage/fabrics/%s/switches/%s/vpcPair"
const urlVpcPairGet = "/manage/fabrics/%s/switches/%s/vpcPair"
const urlVpcPairRecmd = "/manage/fabrics/%s/switches/%s/vpcPairRecommendation?useVirtualPeerLink=%t"

type VpcPairAPI struct {
	ndapi.NexusDashboardAPICommon
	mutex              *sync.Mutex
	GetRecommendations bool
	FabricName         string
	SwitchID           string
	VirtualPeerLink    bool
}

func (c *VpcPairAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *VpcPairAPI) GetUrl() string {
	if c.GetRecommendations {
		return fmt.Sprintf(urlVpcPairRecmd, c.FabricName, c.SwitchID, c.VirtualPeerLink)
	} else {
		return fmt.Sprintf(urlVpcPairGet, c.FabricName, c.SwitchID)
	}
}

func (c *VpcPairAPI) PostUrl() string {
	return fmt.Sprintf(urlVpcPair, c.FabricName, c.SwitchID)
}

func (c *VpcPairAPI) PutUrl() string {
	return fmt.Sprintf(urlVpcPair, c.FabricName, c.SwitchID)
}

func (c *VpcPairAPI) DeleteUrl() string {
	return fmt.Sprintf(urlVpcPairGet, c.FabricName, c.SwitchID)
}
func (c *VpcPairAPI) GetDeleteQP() []string {
	return nil
}

func (c *VpcPairAPI) RscName() string {
	return "vpc-pair"
}

func NewVpcPairAPI(lock *sync.Mutex, client *nd.Client) *VpcPairAPI {
	papi := new(VpcPairAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NexusDashboardAPI = papi
	return papi
}
