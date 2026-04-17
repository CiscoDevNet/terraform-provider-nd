// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_inventory_switch

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ModuleKey = "manage"

var _ resource.Resource = &inventorySwitchResource{}
var _ resource.ResourceWithConfigure = &inventorySwitchResource{}
var _ resource.ResourceWithImportState = &inventorySwitchResource{}

func NewInventorySwitchResource() resource.Resource {
	return &inventorySwitchResource{}
}

type inventorySwitchResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *inventorySwitchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_switch"
}

func (r *inventorySwitchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = InventorySwitchResourceSchema(ctx)
}

func (r *inventorySwitchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	module := client.GetModule(ModuleKey)
	if module == nil {
		resp.Diagnostics.AddError(
			"Module Not Found",
			fmt.Sprintf("Could not find module '%s'. Please report this issue to the provider developers.", ModuleKey),
		)
		return
	}

	manageModule, ok := module.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", module),
		)
		return
	}

	r.manageClient = manageModule
	tflog.Debug(ctx, "Configured inventory_switch resource", map[string]interface{}{
		"module": ModuleKey,
	})
}

func (r *inventorySwitchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InventorySwitchModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating inventory_switch resource", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
	})

	r.rscCreateInventory(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *inventorySwitchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InventorySwitchModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading inventory_switch resource", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	r.rscGetInventory(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *inventorySwitchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InventorySwitchModel
	var state InventorySwitchModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating inventory_switch resource", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
	})

	r.rscUpdateInventory(ctx, &resp.Diagnostics, &plan, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *inventorySwitchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InventorySwitchModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting inventory_switch resource", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	r.rscDeleteInventory(ctx, &resp.Diagnostics, &state)
}

func (r *inventorySwitchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing inventory_switch resource", map[string]interface{}{
		"id": req.ID,
	})

	r.rscImportInventory(ctx, &resp.Diagnostics, req.ID, resp)
}
