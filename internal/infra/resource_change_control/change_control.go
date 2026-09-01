// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_change_control

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	ModuleKey             = "infra"
	changeControlImportID = "changeControl"
)

var (
	_ resource.Resource                   = &changeControlResource{}
	_ resource.ResourceWithConfigure      = &changeControlResource{}
	_ resource.ResourceWithImportState    = &changeControlResource{}
	_ resource.ResourceWithValidateConfig = &changeControlResource{}
)

func NewChangeControlResource() resource.Resource {
	return &changeControlResource{}
}

type changeControlResource struct {
	infraClient *infra.NexusDashboardInfra
}

func (r *changeControlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_change_control"
}

func (r *changeControlResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ChangeControlResourceSchema(ctx)
}

// ValidateConfig validates the relationship between the change control
// feature flags before Terraform creates a plan.
func (r *changeControlResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config ChangeControlModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ValidateConfig receives configuration, so omitted attributes can still be
	// null even though their schema defaults will be present in the plan.
	adminStatus := false
	if !config.AdminStatus.IsNull() {
		if config.AdminStatus.IsUnknown() {
			return
		}
		adminStatus = config.AdminStatus.ValueBool()
	}

	orchestration := false
	if !config.Orchestration.IsNull() {
		if config.Orchestration.IsUnknown() {
			return
		}
		orchestration = config.Orchestration.ValueBool()
	}

	ndManagedFabrics := false
	if !config.NdManagedFabrics.IsNull() {
		if config.NdManagedFabrics.IsUnknown() {
			return
		}
		ndManagedFabrics = config.NdManagedFabrics.ValueBool()
	}

	if adminStatus && !orchestration && !ndManagedFabrics {
		resp.Diagnostics.AddAttributeError(
			path.Root("admin_status"),
			"Invalid Change Control Configuration",
			"admin_status can be enabled only when orchestration or nd_managed_fabrics is enabled.",
		)
	}
}

func (r *changeControlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	infraModule, ok := module.(*infra.NexusDashboardInfra)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Module Type",
			fmt.Sprintf("Expected *infra.NexusDashboardInfra, got: %T. Please report this issue to the provider developers.", module),
		)
		return
	}

	r.infraClient = infraModule
	tflog.Debug(ctx, "Configured change_control resource", map[string]interface{}{
		"module": ModuleKey,
	})
}

func (r *changeControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ChangeControlModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The API exposes no POST because change control is a singleton. Use its
	// supported PUT operation to apply the initial Terraform configuration.
	tflog.Debug(ctx, "Creating change_control resource with PUT", map[string]interface{}{
		"resource": changeControlImportID,
	})

	r.rscPutChangeControl(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *changeControlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ChangeControlModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading change_control resource", map[string]interface{}{
		"resource": changeControlImportID,
	})

	r.rscGetChangeControl(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *changeControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ChangeControlModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating change_control resource", map[string]interface{}{
		"resource": changeControlImportID,
	})

	r.rscPutChangeControl(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *changeControlResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	defaults := ChangeControlModel{
		Id:                           types.StringValue(changeControlImportID),
		AdminStatus:                  types.BoolValue(false),
		Orchestration:                types.BoolValue(false),
		NumberOfApprovers:            types.Int64Value(1),
		AllowSelfApproval:            types.BoolValue(true),
		NdManagedFabrics:             types.BoolValue(false),
		BypassTelemetryChangeControl: types.BoolValue(false),
		TicketNamePrefix:             types.StringValue("TICKET_"),
	}

	// The API does not support DELETE because change control is a built-in
	// singleton. Restore its default settings with PUT before Terraform removes
	// the resource from state.
	tflog.Debug(ctx, "Resetting change_control settings to defaults before removing Terraform state", map[string]interface{}{
		"resource": changeControlImportID,
	})

	r.rscPutChangeControl(ctx, &resp.Diagnostics, &defaults)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *changeControlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != changeControlImportID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected import ID %q for nd_change_control, got %q.", changeControlImportID, req.ID),
		)
		return
	}

	var state ChangeControlModel

	tflog.Debug(ctx, "Importing change_control resource", map[string]interface{}{
		"resource": req.ID,
	})

	// All change_control settings are returned by the canonical GET endpoint;
	// there are no sensitive or write-only attributes to restore separately.
	r.rscGetChangeControl(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
