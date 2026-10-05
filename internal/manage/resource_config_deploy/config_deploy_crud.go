// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_config_deploy

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/deployment"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ModuleKey = "manage"

// ticketIdRegex validates ticket_id: begins with a letter, followed by up to 63
// letters, numbers, underscores, or hyphens (total length 1-64).
var ticketIdRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]{0,63}$`)

var _ resource.Resource = &configDeployResource{}
var _ resource.ResourceWithConfigure = &configDeployResource{}
var _ resource.ResourceWithValidateConfig = &configDeployResource{}

func NewConfigDeployResource() resource.Resource {
	return &configDeployResource{}
}

type configDeployResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *configDeployResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_deploy"
}

func (r *configDeployResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ConfigDeployResourceSchema(ctx)
}

func (r *configDeployResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T.", req.ProviderData),
		)
		return
	}

	module := client.GetModule(ModuleKey)
	if module == nil {
		resp.Diagnostics.AddError("Module Not Found",
			fmt.Sprintf("Could not find module '%s'.", ModuleKey))
		return
	}

	manageModule, ok := module.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T.", module))
		return
	}

	r.manageClient = manageModule
}

func (r *configDeployResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ConfigDeployModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Only validate the deploy/config_save combination when both are known.
	// Unknown (interpolated) values would otherwise raise false positives at plan time.
	if !data.Deploy.IsUnknown() && !data.ConfigSave.IsUnknown() {
		if !data.Deploy.ValueBool() && !data.ConfigSave.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("deploy"),
				"Invalid Configuration",
				"At least one of 'deploy' or 'config_save' must be true.",
			)
		}
	}

	// Validate switch_ids: if set, must not be empty and "ALL" cannot be combined with others
	if !data.SwitchIds.IsNull() && !data.SwitchIds.IsUnknown() {
		var switchIds []string
		data.SwitchIds.ElementsAs(ctx, &switchIds, false)
		if len(switchIds) == 0 {
			resp.Diagnostics.AddAttributeError(
				path.Root("switch_ids"),
				"Invalid Configuration",
				"switch_ids must contain at least one entry when specified.",
			)
		}
		for _, s := range switchIds {
			if strings.EqualFold(s, "ALL") && len(switchIds) > 1 {
				resp.Diagnostics.AddAttributeError(
					path.Root("switch_ids"),
					"Invalid Configuration",
					"switch_ids cannot combine 'ALL' with specific switch IDs.",
				)
				break
			}
		}
	}

	// Validate ticket_id format: must begin with a letter; remaining characters
	// may be letters, numbers, underscores, or hyphens; length 1-64.
	if !data.TicketId.IsNull() && !data.TicketId.IsUnknown() {
		if tid := data.TicketId.ValueString(); !ticketIdRegex.MatchString(tid) {
			resp.Diagnostics.AddAttributeError(
				path.Root("ticket_id"),
				"Invalid Configuration",
				"ticket_id must begin with a letter and contain only letters, numbers, underscores, or hyphens (1-64 characters).",
			)
		}
	}
}

func (r *configDeployResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConfigDeployModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := plan.GetModelData()

	// Check behaviour.execute_on_create
	if data.Behaviour.ExecuteOnCreate != nil && *data.Behaviour.ExecuteOnCreate {
		r.executeDeployment(ctx, &resp.Diagnostics, data)
	}

	data.Id = data.FabricName
	plan.SetModelData(data)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *configDeployResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConfigDeployModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// For always_execute: reset to false in state so Terraform sees a diff
	// when the user has always_execute = true, triggering Update on every apply.
	// Set the field directly to avoid round-tripping through SetModelData,
	// which would wipe computed fields like Status.
	if state.AlwaysExecute.ValueBool() {
		state.AlwaysExecute = types.BoolValue(false)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *configDeployResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConfigDeployModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := plan.GetModelData()

	// Check behaviour.execute_on_update
	if data.Behaviour.ExecuteOnUpdate != nil && *data.Behaviour.ExecuteOnUpdate {
		r.executeDeployment(ctx, &resp.Diagnostics, data)
	}

	data.Id = data.FabricName
	plan.SetModelData(data)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *configDeployResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConfigDeployModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := state.GetModelData()

	// Check behaviour.execute_on_destroy
	if data.Behaviour.ExecuteOnDestroy != nil && *data.Behaviour.ExecuteOnDestroy {
		tflog.Info(ctx, "Executing config-save/deploy on destroy", map[string]interface{}{
			"fabric_name": data.FabricName,
		})
		r.executeDeployment(ctx, &resp.Diagnostics, data)
	}
}

// executeDeployment performs config-save and/or deploy via deployment package.
// switch_ids that are IP addresses are resolved to serial numbers via SwitchDB
// before being sent to the deploy API.
func (r *configDeployResource) executeDeployment(ctx context.Context, dg *diag.Diagnostics, data *NDFCConfigDeployModel) {
	serials := r.resolveSwitchIds(ctx, dg, data.FabricName, data.SwitchIds)
	if dg.HasError() {
		return
	}

	opts := &deployment.DeployOptions{
		SerialNumbers:                 serials,
		TicketId:                      data.TicketId,
		ForceShowRun:                  data.ForceShowRun,
		IncludeAllFabricGroupSwitches: data.IncludeAllFabricGroupSwitches,
	}

	result := deployment.ConfigSaveAndDeployWithOpts(
		ctx, r.manageClient.ApiClient, data.FabricName,
		data.ConfigSave, data.Deploy, opts, dg,
	)

	// Preserve whatever status was produced, including a successful config-save
	// status when a subsequent deploy fails, for observability in state.
	data.Status = result.Status
}

// resolveSwitchIds converts any IP addresses in the switch_ids list to serial
// numbers using SwitchDB. Serial numbers and "ALL" are passed through unchanged.
func (r *configDeployResource) resolveSwitchIds(ctx context.Context, dg *diag.Diagnostics, fabricName string, ids []string) []string {
	if len(ids) == 0 {
		return ids
	}

	resolved := make([]string, 0, len(ids))
	for _, id := range ids {
		if strings.EqualFold(id, "ALL") {
			resolved = append(resolved, id)
			continue
		}

		if net.ParseIP(id) != nil {
			// It's an IP address — resolve to serial via SwitchDB
			serial, ok := r.manageClient.SwitchDB.GetSerialByIP(ctx, fabricName, id)
			if !ok {
				dg.AddAttributeError(
					path.Root("switch_ids"),
					"Switch Not Found",
					fmt.Sprintf("Could not resolve IP address %q to a serial number in fabric %q. Verify the switch is in the fabric inventory.", id, fabricName),
				)
				return nil
			}
			tflog.Debug(ctx, "Resolved switch IP to serial", map[string]interface{}{
				"ip":     id,
				"serial": serial,
				"fabric": fabricName,
			})
			resolved = append(resolved, serial)
		} else {
			// Assume it's a serial number
			resolved = append(resolved, id)
		}
	}
	return resolved
}
