// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_common

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/netascode/go-nd"
)

// FabricHandler is the strategy interface for fabric CRUD operations.
// Embed DefaultFabricHandler and override only the methods you need.
type FabricHandler interface {
	Create(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel)
	Read(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool
	Update(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel)
	Delete(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel)
}

// DefaultFabricHandler implements FabricHandler using the common CRUD flow.
// Embed it in custom handlers so you only override methods you need.
type DefaultFabricHandler struct{}

func (DefaultFabricHandler) Create(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	createDefault(ctx, client, dg, model)
}
func (DefaultFabricHandler) Read(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	return readDefault(ctx, client, dg, model)
}
func (DefaultFabricHandler) Update(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	updateDefault(ctx, client, dg, model)
}
func (DefaultFabricHandler) Delete(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	deleteDefault(ctx, client, dg, model)
}

// handler registry — keyed by fabric type string (e.g. "vxlanEbgp")
var (
	handlersMu sync.RWMutex
	handlers   = map[string]FabricHandler{}
)

// RegisterFabricHandler registers a custom handler for a fabric type.
// Call this from an init() in the fabric-specific package.
func RegisterFabricHandler(fabricType string, h FabricHandler) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	handlers[fabricType] = h
}

// fabricHandler returns the registered handler for the fabric type, or
// DefaultFabricHandler (common CRUD flow) if none is registered.
func fabricHandler(fabricType string) FabricHandler {
	handlersMu.RLock()
	defer handlersMu.RUnlock()
	if h, ok := handlers[fabricType]; ok {
		return h
	}
	return DefaultFabricHandler{}
}
