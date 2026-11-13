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

const (
	urlPolicies         = "/manage/fabrics/%s/policies"
	urlPolicy           = "/manage/fabrics/%s/policies/%s"
	urlPolicyMarkDelete = "/manage/fabrics/%s/policyActions/markDelete"
	urlPolicyPushConfig = "/manage/fabrics/%s/policyActions/pushConfig"
	urlPolicyRemove     = "/manage/fabrics/%s/policyActions/remove"

	urlPolicyGroups    = "/manage/fabrics/%s/policyGroups"
	urlPolicyGroup     = "/manage/fabrics/%s/policyGroups/%s"
	urlGroupMarkDelete = "/manage/fabrics/%s/policyGroups/actions/markDelete"
	urlGroupRemove     = "/manage/fabrics/%s/policyGroups/actions/remove"

	rscNamePolicy = "policy"
)

// PolicyOperation selects which Manage policy URL family to call.
type PolicyOperation int

const (
	// OpPolicyCRUD is collection POST and per-ID GET/PUT/DELETE.
	OpPolicyCRUD PolicyOperation = iota
	// OpMarkDelete is POST .../markDelete.
	OpMarkDelete
	// OpPushConfig is POST .../policyActions/pushConfig (single policy only).
	OpPushConfig
	// OpRemove is POST bulk remove (fallback only).
	OpRemove
)

// PolicyAPI is the API client for nd_policy (single policy or policy group).
type PolicyAPI struct {
	ndapi.NexusDashboardAPICommon
	FabricName  string
	PolicyID    string
	PolicyGroup bool
	Operation   PolicyOperation
}

// NewPolicyAPI creates a PolicyAPI scoped to a fabric for locking.
func NewPolicyAPI(client *nd.Client, fabric string) *PolicyAPI {
	papi := new(PolicyAPI)
	papi.Client = client
	papi.Fabric = fabric
	papi.FabricName = fabric
	papi.NexusDashboardAPI = papi
	return papi
}

func (c *PolicyAPI) collectionURL() string {
	if c.PolicyGroup {
		return fmt.Sprintf(urlPolicyGroups, c.FabricName)
	}
	return fmt.Sprintf(urlPolicies, c.FabricName)
}

func (c *PolicyAPI) objectURL() string {
	if c.PolicyGroup {
		return fmt.Sprintf(urlPolicyGroup, c.FabricName, c.PolicyID)
	}
	return fmt.Sprintf(urlPolicy, c.FabricName, c.PolicyID)
}

func (c *PolicyAPI) actionURL() string {
	switch c.Operation {
	case OpMarkDelete:
		if c.PolicyGroup {
			return fmt.Sprintf(urlGroupMarkDelete, c.FabricName)
		}
		return fmt.Sprintf(urlPolicyMarkDelete, c.FabricName)
	case OpPushConfig:
		return fmt.Sprintf(urlPolicyPushConfig, c.FabricName)
	case OpRemove:
		if c.PolicyGroup {
			return fmt.Sprintf(urlGroupRemove, c.FabricName)
		}
		return fmt.Sprintf(urlPolicyRemove, c.FabricName)
	default:
		return c.collectionURL()
	}
}

func (c *PolicyAPI) GetUrl() string {
	if strings.TrimSpace(c.PolicyID) == "" {
		return c.collectionURL()
	}
	return c.objectURL()
}

func (c *PolicyAPI) PostUrl() string {
	switch c.Operation {
	case OpMarkDelete, OpPushConfig, OpRemove:
		return c.actionURL()
	default:
		return c.collectionURL()
	}
}

func (c *PolicyAPI) PutUrl() string {
	return c.objectURL()
}

func (c *PolicyAPI) DeleteUrl() string {
	return c.objectURL()
}

func (c *PolicyAPI) GetQueryParams() []string {
	return nil
}

func (c *PolicyAPI) RscName() string {
	return rscNamePolicy
}

// SetOperation selects the URL family for the next call.
func (c *PolicyAPI) SetOperation(op PolicyOperation) *PolicyAPI {
	c.Operation = op
	return c
}
