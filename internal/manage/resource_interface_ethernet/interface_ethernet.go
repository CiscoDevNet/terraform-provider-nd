// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/resource_interface_common"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

var (
	_ resource.Resource                   = &interfaceEthernetResource{}
	_ resource.ResourceWithConfigure      = &interfaceEthernetResource{}
	_ resource.ResourceWithValidateConfig = &interfaceEthernetResource{}
	_ resource.ResourceWithImportState    = &interfaceEthernetResource{}
)

// ValidateConfig applies the ethernet discriminator matrix and the
// userDefined/template_name contract. See validators.go.
func (r *interfaceEthernetResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	ethernetConfigValidator{}.ValidateResource(ctx, req, resp)
}

func NewInterfaceEthernetResource() resource.Resource {
	return &interfaceEthernetResource{}
}

type interfaceEthernetResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *interfaceEthernetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet"
}

func (r *interfaceEthernetResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = InterfaceEthernetResourceSchema(ctx)
}

func (r *interfaceEthernetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	m := client.GetModule(ModuleKey)
	if m == nil {
		resp.Diagnostics.AddError(
			"Manage Module Not Found",
			"The manage module was not registered with the provider.",
		)
		return
	}
	r.manageClient = m.(*manage.NexusDashboardManage)
}

func (r *interfaceEthernetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in InterfaceEthernetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating nd_interface_ethernet", map[string]interface{}{
		"fabric_name": in.FabricName.ValueString(),
	})
	// Apply-time validation: cross-check attributes against the live ND
	// template parameter list.
	validateAttrsAgainstTemplate(ctx, &resp.Diagnostics, &in)
	resource_interface_common.RscCreateInterfaces(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		// Distinguish deploy-failure+rollback (model cleared, infra clean)
		// from partial-create-failure (model has successfully-created entries).
		modelData := in.GetModelData()
		if len(modelData.Interfaces) == 0 {
			// Deploy failed and rollback removed everything — skip state
			// so both TF state and infra are empty for a clean retry.
			return
		}
		// Partial create failure: write partial state so retry doesn't
		// duplicate the successfully-created entries.
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
}

func (r *interfaceEthernetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Reading nd_interface_ethernet", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})
	found := resource_interface_common.RscGetInterfaces(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}
	// Derive deploy flag from actual configSyncStatus on the switch.
	// This ensures drift (e.g. failed deploy, out-of-band change) is
	// surfaced as deploy=false, triggering a plan diff if the user's
	// config says deploy=true.
	resource_interface_common.RscDeriveDeployStatus(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *interfaceEthernetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceEthernetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state InterfaceEthernetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updating nd_interface_ethernet", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
	})
	// Apply-time validation: cross-check attributes against the live ND
	// template parameter list.
	validateAttrsAgainstTemplate(ctx, &resp.Diagnostics, &plan)
	resource_interface_common.RscUpdateInterfaces(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *interfaceEthernetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing nd_interface_ethernet", map[string]interface{}{
		"id": req.ID,
	})

	// ID format: fabric_name/switch_id/intf1,intf2,...
	parts := strings.SplitN(req.ID, "/", 3)
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected format: fabric_name/switch_id/intf1,intf2,... — got %q", req.ID),
		)
		return
	}
	fabricName := parts[0]
	switchId := parts[1]
	ifaceNames := strings.Split(parts[2], ",")

	// Build a skeleton TF model so the common Read can iterate the entries.
	// The mandatory fields (mode, network_os_type, policy_type) are set to
	// unknown — Read will overwrite them from ND.
	elemType := InterfacesValue{}.Type(ctx)
	elems := make(map[string]attr.Value, len(ifaceNames))
	for _, name := range ifaceNames {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		elems[name] = NewInterfacesValueUnknown()
	}
	ifaceMap, diags := types.MapValue(elemType, elems)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state := InterfaceEthernetModel{
		Id:         types.StringValue(req.ID),
		FabricName: types.StringValue(fabricName),
		SwitchId:   types.StringValue(switchId),
		Deploy:     types.BoolValue(false),
		Interfaces: ifaceMap,
	}

	// Read from ND to populate all fields.
	// resolveSwitchIDs runs inside readDefault after PreMarshal.
	found := resource_interface_common.RscGetInterfaces(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError(
			"Import Failed",
			fmt.Sprintf("No interfaces found in fabric %q for switch %q matching %v", fabricName, switchId, ifaceNames),
		)
		return
	}
	// Derive deploy flag from actual configSyncStatus on the switch.
	resource_interface_common.RscDeriveDeployStatus(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *interfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Deleting nd_interface_ethernet", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})
	resource_interface_common.RscDeleteInterfaces(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
}
