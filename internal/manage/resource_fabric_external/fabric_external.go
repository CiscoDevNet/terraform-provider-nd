// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_external

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource              = &fabricExternalResource{}
	_ resource.ResourceWithConfigure = &fabricExternalResource{}
)

// NewFabricExternalResource is a helper function to simplify the provider implementation.
func NewFabricExternalResource() resource.Resource {
	return &fabricExternalResource{}
}

// fabricExternalResource is the resource implementation.
type fabricExternalResource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the resource type name.
func (r *fabricExternalResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_external"
}

// Schema defines the schema for the resource.
func (r *fabricExternalResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricExternalResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *fabricExternalResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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

	manageModule := client.GetModule(ModuleKey)
	if manageModule == nil {
		resp.Diagnostics.AddError(
			"Manage Module Not Found",
			"The manage module was not registered with the provider.",
		)
		return
	}

	r.manageClient = manageModule.(*manage.NexusDashboardManage)
}

// Create creates the resource and sets the initial Terraform state.
func (r *fabricExternalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in FabricExternalModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Fabric External", map[string]interface{}{
		"fabric_name": in.FabricName.ValueString(),
	})

	resource_fabric_common.RscCreateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *fabricExternalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricExternalModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Fabric External", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *fabricExternalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FabricExternalModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating Fabric External", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
	})

	resource_fabric_common.RscUpdateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *fabricExternalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricExternalModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting Fabric External", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	resource_fabric_common.RscDeleteFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
}
