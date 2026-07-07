// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vpc_pair

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource                = &vpcPairResource{}
	_ resource.ResourceWithConfigure   = &vpcPairResource{}
	_ resource.ResourceWithImportState = &vpcPairResource{}
)

// NewVpcPairResource is a helper function to simplify the provider implementation.
func NewVpcPairResource() resource.Resource {
	return &vpcPairResource{}
}

// vpcPairResource is the resource implementation.
type vpcPairResource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the resource type name.
func (r *vpcPairResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_pair"
}

// Schema defines the schema for the resource.
func (r *vpcPairResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = VpcPairResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *vpcPairResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *vpcPairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in VpcPairModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode vPC pair create plan")
		return
	}

	tflog.Debug(ctx, "Creating VPC Pair", map[string]interface{}{
		"switch_1_serial_number": in.SwitchId1.ValueString(),
		"switch_2_serial_number": in.SwitchId2.ValueString(),
	})

	r.rscCreateVpcPair(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *vpcPairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VpcPairModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode vPC pair read state")
		return
	}

	tflog.Debug(ctx, "Reading VPC Pair", map[string]interface{}{
		"fabric_name":            state.FabricName.ValueString(),
		"switch_1_serial_number": state.SwitchId1.ValueString(),
		"switch_2_serial_number": state.SwitchId2.ValueString(),
	})

	outData, err := r.readVpcPairState(ctx, state.GetModelData())
	if err != nil {
		if isVpcPairNotFoundError(err) {
			tflog.Warn(ctx, "vPC pair no longer exists in ND, removing from Terraform state", map[string]interface{}{
				"fabric_name":            state.FabricName.ValueString(),
				"switch_1_serial_number": state.SwitchId1.ValueString(),
				"switch_2_serial_number": state.SwitchId2.ValueString(),
				"error":                  err.Error(),
			})
			resp.State.RemoveResource(ctx)
			return
		}

		tflog.Error(ctx, "Failed to read vPC pair state", map[string]interface{}{
			"fabric_name":            state.FabricName.ValueString(),
			"switch_1_serial_number": state.SwitchId1.ValueString(),
			"switch_2_serial_number": state.SwitchId2.ValueString(),
			"error":                  err.Error(),
		})
		resp.Diagnostics.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("Could not read vPC pair, unexpected error: %v", err),
		)
		return
	}

	if outData.UseVirtualPeerlink == nil {
		outData.UseVirtualPeerlink = state.GetModelData().UseVirtualPeerlink
	}
	state.SetModelData(outData)
	setVpcPairID(&state)

	tflog.Debug(ctx, "Read vPC pair", map[string]interface{}{
		"fabric_name": outData.FabricName,
		"switch_id":   outData.SwitchId2,
	})

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func isVpcPairNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	return strings.Contains(msg, "StatusCode 404")
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *vpcPairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VpcPairModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode vPC pair update plan")
		return
	}

	tflog.Debug(ctx, "Updating VPC Pair", map[string]interface{}{
		"fabric_name":            plan.FabricName.ValueString(),
		"switch_1_serial_number": plan.SwitchId1.ValueString(),
		"switch_2_serial_number": plan.SwitchId2.ValueString(),
	})

	r.rscUpdateVpcPair(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vpcPairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VpcPairModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to decode vPC pair delete state")
		return
	}

	tflog.Debug(ctx, "Deleting VPC Pair", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
		"switch_id":   state.SwitchId2.ValueString(),
	})

	r.rscDeleteVpcPair(ctx, &resp.Diagnostics, &state)
}

func (r *vpcPairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing VPC Pair", map[string]interface{}{
		"id": req.ID,
	})

	fabricName := ""
	switchTuple := req.ID
	if parts := strings.SplitN(req.ID, "/", 2); len(parts) == 2 {
		fabricName = strings.TrimSpace(parts[0])
		switchTuple = strings.TrimSpace(parts[1])
	}

	switchParts := strings.SplitN(switchTuple, ":", 2)
	if len(switchParts) != 2 || strings.TrimSpace(switchParts[0]) == "" || strings.TrimSpace(switchParts[1]) == "" {
		tflog.Error(ctx, "Invalid vPC pair import switch tuple", map[string]interface{}{
			"id": req.ID,
		})
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			"Expected format: switch_1_serial_number:switch_2_serial_number or fabric_name/switch_1_serial_number:switch_2_serial_number",
		)
		return
	}

	fabricValue := types.StringNull()
	if fabricName != "" {
		fabricValue = types.StringValue(fabricName)
	}

	state := VpcPairModel{
		FabricName:         fabricValue,
		SwitchId1:          types.StringValue(strings.TrimSpace(switchParts[0])),
		SwitchId2:          types.StringValue(strings.TrimSpace(switchParts[1])),
		Deploy:             types.BoolValue(false),
		UseVirtualPeerlink: types.BoolNull(),
	}

	outData, err := r.readVpcPairState(ctx, state.GetModelData())
	if err != nil {
		tflog.Error(ctx, "Failed to import vPC pair using provided switch order", map[string]interface{}{
			"fabric_name":            state.FabricName.ValueString(),
			"switch_1_serial_number": state.SwitchId1.ValueString(),
			"switch_2_serial_number": state.SwitchId2.ValueString(),
			"error":                  err.Error(),
		})
		resp.Diagnostics.AddError(
			"Error Importing vPC Pair",
			fmt.Sprintf("Could not read imported vPC pair: %v", err),
		)
		return
	}

	if !matchesImportedVpcPair(state.GetModelData(), outData) {
		tflog.Debug(ctx, "Imported vPC pair did not match provided switch order, retrying with swapped switch IDs", map[string]interface{}{
			"fabric_name":            state.FabricName.ValueString(),
			"switch_1_serial_number": state.SwitchId1.ValueString(),
			"switch_2_serial_number": state.SwitchId2.ValueString(),
		})
		swappedInData := &NDFCVpcPairModel{
			FabricName:         state.FabricName.ValueString(),
			SwitchId1:          state.SwitchId2.ValueString(),
			SwitchId2:          state.SwitchId1.ValueString(),
			UseVirtualPeerlink: nil,
			Deploy:             false,
		}

		outData, err = r.readVpcPairState(ctx, swappedInData)
		if err != nil {
			tflog.Error(ctx, "Failed to import vPC pair in either switch order", map[string]interface{}{
				"fabric_name":            state.FabricName.ValueString(),
				"switch_1_serial_number": state.SwitchId1.ValueString(),
				"switch_2_serial_number": state.SwitchId2.ValueString(),
				"error":                  err.Error(),
			})
			resp.Diagnostics.AddError(
				"Error Importing vPC Pair",
				fmt.Sprintf("Could not read imported vPC pair in either switch order: %v", err),
			)
			return
		}
	}

	deployState, err := r.getVpcPairDeployState(ctx, outData.FabricName, outData.SwitchId1, outData.SwitchId2)
	if err != nil {
		tflog.Error(ctx, "Failed to resolve vPC pair deploy state during import", map[string]interface{}{
			"fabric_name":            outData.FabricName,
			"switch_1_serial_number": outData.SwitchId1,
			"switch_2_serial_number": outData.SwitchId2,
			"error":                  err.Error(),
		})
		resp.Diagnostics.AddError(
			"Error Importing vPC Pair",
			fmt.Sprintf("Could not determine imported vPC pair deploy state from fabric switches: %v", err),
		)
		return
	}
	outData.Deploy = deployState

	if diag := state.SetModelData(outData); diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	setVpcPairID(&state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func matchesImportedVpcPair(inData *NDFCVpcPairModel, outData *NDFCVpcPairModel) bool {
	if inData == nil || outData == nil {
		return false
	}

	if inData.FabricName != "" && inData.FabricName != outData.FabricName {
		return false
	}

	return inData.SwitchId1 == outData.SwitchId1 &&
		inData.SwitchId2 == outData.SwitchId2
}
