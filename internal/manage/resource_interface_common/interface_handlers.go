// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/netascode/go-nd"
)

// InterfaceHandler is the strategy interface for per-type CRUD overrides on
// nd_interface_* resources. Embed DefaultInterfaceHandler and override only
// the methods you need (see interface_ethernet.spec §Breakout for the
// ethernet owner-mismatch Read override).
//
// Update receives both the prior state and the new plan so that the diff can
// detect removed interfaces (present in state but absent from plan).
type InterfaceHandler interface {
	Create(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel)
	Read(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) bool
	Update(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, state InterfaceModel, plan InterfaceModel)
	Delete(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel)
}

// DefaultInterfaceHandler implements InterfaceHandler by delegating every op
// to the common CRUD flow. Embed it in a custom handler if you only need to
// override one method.
type DefaultInterfaceHandler struct{}

func (DefaultInterfaceHandler) Create(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	createDefault(ctx, client, dg, model)
}
func (DefaultInterfaceHandler) Read(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) bool {
	return readDefault(ctx, client, dg, model)
}
func (DefaultInterfaceHandler) Update(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, state InterfaceModel, plan InterfaceModel) {
	updateDefault(ctx, client, dg, state, plan)
}
func (DefaultInterfaceHandler) Delete(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	deleteDefault(ctx, client, dg, model)
}

// handler registry — keyed by interface type string (e.g. "ethernet", "vpc").
var (
	handlersMu sync.RWMutex
	handlers   = map[string]InterfaceHandler{}
)

// RegisterInterfaceHandler registers a custom handler for an interface type.
// Call this from an init() in the per-type package.
func RegisterInterfaceHandler(interfaceType string, h InterfaceHandler) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	handlers[interfaceType] = h
}

// interfaceHandler returns the registered handler, or DefaultInterfaceHandler
// (common CRUD flow) when none is registered for the given type.
func interfaceHandler(interfaceType string) InterfaceHandler {
	handlersMu.RLock()
	defer handlersMu.RUnlock()
	if h, ok := handlers[interfaceType]; ok {
		return h
	}
	return DefaultInterfaceHandler{}
}
