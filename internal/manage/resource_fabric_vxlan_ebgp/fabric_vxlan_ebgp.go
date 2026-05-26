// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_vxlan_ebgp

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
	_ resource.Resource              = &fabricVxlanEbgpResource{}
	_ resource.ResourceWithConfigure = &fabricVxlanEbgpResource{}
)

// NewFabricVxlanEbgpResource is a helper function to simplify the provider implementation.
func NewFabricVxlanEbgpResource() resource.Resource {
	return &fabricVxlanEbgpResource{}
}

// fabricVxlanEbgpResource is the resource implementation.
type fabricVxlanEbgpResource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the resource type name.
func (r *fabricVxlanEbgpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_vxlan_ebgp"
}

// Schema defines the schema for the resource.
func (r *fabricVxlanEbgpResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricVxlanEbgpResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *fabricVxlanEbgpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *fabricVxlanEbgpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in FabricVxlanEbgpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Fabric VXLAN eBGP", map[string]interface{}{
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
func (r *fabricVxlanEbgpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricVxlanEbgpModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Fabric VXLAN eBGP", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *fabricVxlanEbgpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FabricVxlanEbgpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating Fabric VXLAN eBGP", map[string]interface{}{
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
func (r *fabricVxlanEbgpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricVxlanEbgpModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting Fabric VXLAN eBGP", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	resource_fabric_common.RscDeleteFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
}
