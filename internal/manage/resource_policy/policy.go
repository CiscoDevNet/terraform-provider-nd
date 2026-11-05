// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_policy

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

const (
	policyIDPrefix      = "POLICY-"
	policyGroupIDPrefix = "POLICY-GROUP-"
)

var (
	_ resource.Resource                   = &policyResource{}
	_ resource.ResourceWithConfigure      = &policyResource{}
	_ resource.ResourceWithImportState    = &policyResource{}
	_ resource.ResourceWithValidateConfig = &policyResource{}
)

// NewPolicyResource is a helper function to simplify the provider implementation.
func NewPolicyResource() resource.Resource {
	return &policyResource{}
}

type policyResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *policyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (r *policyResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = PolicyResourceSchema(ctx)
}

func (r *policyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *policyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config PolicyModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.SerialNumbers.IsNull() || config.SerialNumbers.IsUnknown() ||
		config.IsPolicyGroup.IsUnknown() {
		return
	}

	var serials []string
	resp.Diagnostics.Append(config.SerialNumbers.ElementsAs(ctx, &serials, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.IsPolicyGroup.ValueBool() {
		if len(serials) == 0 {
			resp.Diagnostics.AddAttributeError(
				path.Root("serial_numbers"),
				"Invalid Configuration",
				"serial_numbers must contain at least one switch when is_policy_group is true.",
			)
		}
	} else if len(serials) != 1 {
		resp.Diagnostics.AddAttributeError(
			path.Root("serial_numbers"),
			"Invalid Configuration",
			"serial_numbers must contain exactly one switch when is_policy_group is false.",
		)
	}

	if !config.EntityType.IsNull() && !config.EntityType.IsUnknown() &&
		!config.EntityName.IsNull() && !config.EntityName.IsUnknown() &&
		config.EntityType.ValueString() == "switch" &&
		config.EntityName.ValueString() != "SWITCH" {
		resp.Diagnostics.AddAttributeError(
			path.Root("entity_name"),
			"Invalid Configuration",
			"entity_name must be SWITCH when entity_type is switch.",
		)
	}
}

func (r *policyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode policy create plan")
		return
	}

	tflog.Debug(ctx, "Creating policy", map[string]interface{}{
		"fabric_name":     plan.FabricName.ValueString(),
		"template_name":   plan.TemplateName.ValueString(),
		"is_policy_group": plan.IsPolicyGroup.ValueBool(),
	})

	r.rscCreatePolicy(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *policyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode policy read state")
		return
	}

	tflog.Debug(ctx, "Reading policy", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
		"policy_id":   state.PolicyId.ValueString(),
	})

	found := r.rscReadPolicy(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		tflog.Warn(ctx, "Policy no longer exists in ND, removing from Terraform state", map[string]interface{}{
			"fabric_name": state.FabricName.ValueString(),
			"policy_id":   state.PolicyId.ValueString(),
		})
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *policyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PolicyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode policy update plan")
		return
	}

	tflog.Debug(ctx, "Updating policy", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
		"policy_id":   plan.PolicyId.ValueString(),
	})

	r.rscUpdatePolicy(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *policyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode policy delete state")
		return
	}

	tflog.Debug(ctx, "Deleting policy", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
		"policy_id":   state.PolicyId.ValueString(),
	})

	r.rscDeletePolicy(ctx, &resp.Diagnostics, &state)
}

func (r *policyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing policy", map[string]interface{}{
		"policy_id": req.ID,
	})

	policyID, err := parsePolicyImportID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Import ID", err.Error())
		return
	}

	isGroup := strings.HasPrefix(policyID, policyGroupIDPrefix)
	resolvedFabric, err := r.findPolicyFabric(ctx, policyID, isGroup)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Importing Policy",
			fmt.Sprintf("Could not locate policy %s: %v", policyID, err),
		)
		return
	}

	state := PolicyModel{
		FabricName:    types.StringValue(resolvedFabric),
		PolicyId:      types.StringValue(policyID),
		IsPolicyGroup: types.BoolValue(isGroup),
		Deploy:        types.BoolValue(false),
	}

	found := r.rscReadPolicy(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError(
			"Error Importing Policy",
			fmt.Sprintf("Policy %s was not found in fabric %s", policyID, resolvedFabric),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func parsePolicyImportID(id string) (string, error) {
	policyID := strings.TrimSpace(id)
	if !strings.HasPrefix(policyID, policyIDPrefix) {
		return "", fmt.Errorf("expected POLICY-<n> or POLICY-GROUP-<n>, got %q", policyID)
	}
	return policyID, nil
}
