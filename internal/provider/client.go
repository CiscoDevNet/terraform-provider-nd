// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"time"

	"terraform-provider-nd/internal/registry"

	nd "github.com/netascode/go-nd"
)

var _ registry.ClientProvider = (*NDClient)(nil)

type NDClient struct {
	URL       string
	Username  string
	Password  string
	Domain    string
	Insecure  bool
	Timeout   time.Duration
	ApiClient *nd.Client
	NDModules map[string]interface{}
}

func (c *NDClient) GetModule(name string) interface{} {
	return c.NDModules[name]
}
