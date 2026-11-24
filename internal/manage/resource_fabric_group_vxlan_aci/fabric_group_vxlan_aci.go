// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_group_vxlan_aci

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ModuleKey = "manage"

var (
	_ resource.Resource                = &fabricGroupVxlanAciResource{}
	_ resource.ResourceWithConfigure   = &fabricGroupVxlanAciResource{}
	_ resource.ResourceWithImportState = &fabricGroupVxlanAciResource{}
)

func NewFabricGroupVxlanAciResource() resource.Resource {
	return &fabricGroupVxlanAciResource{}
}

type fabricGroupVxlanAciResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *fabricGroupVxlanAciResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_group_vxlan_aci"
}

func (r *fabricGroupVxlanAciResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricGroupVxlanAciResourceSchema(ctx)
}

func (r *fabricGroupVxlanAciResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected registry.ClientProvider, got %T.", req.ProviderData))
		return
	}
	module := client.GetModule(ModuleKey)
	manageClient, ok := module.(*manage.NexusDashboardManage)
	if !ok || manageClient == nil || manageClient.ApiClient == nil {
		resp.Diagnostics.AddError("Manage Module Not Available", "The provider must configure the manage module and its API client before using fabric groups.")
		return
	}
	r.manageClient = manageClient
	tflog.Debug(ctx, "Configured fabric_group_vxlan_aci resource")
}

func (r *fabricGroupVxlanAciResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	resp.Diagnostics.Append(plan.validateRouteServerConversion(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.Id = plan.FabricName
	tflog.Debug(ctx, "Creating VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": plan.FabricName.ValueString()})
	model := fabricGroupCRUDModel{FabricGroupVxlanAciModel: &plan}
	resource_fabric_common.RscCreateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &model)
	if !model.writeAccepted {
		return
	}
	// Return the accepted object's identity even if asynchronous readback fails.
	if plan.SecurityGroupStatus.IsUnknown() {
		plan.SecurityGroupStatus = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, "Created VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": plan.FabricName.ValueString()})
	}
}

func (r *fabricGroupVxlanAciResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	model := fabricGroupCRUDModel{FabricGroupVxlanAciModel: &state}
	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &model)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		tflog.Warn(ctx, "Fabric group no longer exists; removing it from state", map[string]interface{}{"fabric_name": state.GetFabricName()})
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *fabricGroupVxlanAciResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	if plan.FabricName.ValueString() != state.GetFabricName() {
		resp.Diagnostics.AddError("Fabric Group Rename Requires Replacement", "Changing fabric_name requires replacement of the fabric group.")
		return
	}
	resp.Diagnostics.Append(plan.validateRouteServerConversion(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.Id = plan.FabricName
	model := fabricGroupCRUDModel{
		FabricGroupVxlanAciModel: &plan,
		verifyRouteServerRemoval: plan.RouteServerCollection.IsNull() && len(state.RouteServerCollection.Elements()) != 0,
	}
	tflog.Debug(ctx, "Updating VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": plan.FabricName.ValueString()})
	resource_fabric_common.RscUpdateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &model)
	if !model.writeAccepted || resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, "Updated VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": plan.FabricName.ValueString()})
	}
}

func (r *fabricGroupVxlanAciResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	model := fabricGroupCRUDModel{FabricGroupVxlanAciModel: &state}
	resource_fabric_common.RscDeleteFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &model)
}

func (r *fabricGroupVxlanAciResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if !r.configured(&resp.Diagnostics) {
		return
	}
	if strings.TrimSpace(req.ID) == "" {
		resp.Diagnostics.AddError("Invalid Fabric Group Import ID", "Import requires the non-empty fabric-group name.")
		return
	}
	state := FabricGroupVxlanAciModel{Id: types.StringValue(req.ID), FabricName: types.StringValue(req.ID)}
	// Import hydrates optional backup settings too. The encrypted BGP key is
	// intentionally not imported; configure it before changing authentication.
	model := fabricGroupCRUDModel{FabricGroupVxlanAciModel: &state, importing: true}
	if !resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &model) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Error Importing Fabric Group", fmt.Sprintf("Fabric group %q was not found.", req.ID))
		}
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, "Imported VXLAN-ACI fabric group", map[string]interface{}{"fabric_name": req.ID})
	}
}

func (r *fabricGroupVxlanAciResource) configured(dg *diag.Diagnostics) bool {
	if r.manageClient == nil || r.manageClient.ApiClient == nil {
		dg.AddError("Manage Client Not Configured", "Configure the manage provider module before using fabric_group_vxlan_aci.")
		return false
	}
	return true
}
