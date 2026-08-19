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
	"time"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ModuleKey = "manage"

var _ resource.Resource = &inventorySwitchResource{}
var _ resource.ResourceWithConfigure = &inventorySwitchResource{}
var _ resource.ResourceWithImportState = &inventorySwitchResource{}
var _ resource.ResourceWithValidateConfig = &inventorySwitchResource{}

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

func (r *inventorySwitchResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config InventorySwitchModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	mode := config.Mode.ValueString()
	if mode == "" || config.Mode.IsUnknown() {
		mode = "discovery" // default
	}

	// Validate mode is a known value
	if mode != "discovery" && mode != "bootstrap" {
		resp.Diagnostics.AddAttributeError(
			path.Root("mode"),
			"Invalid Attribute Value",
			fmt.Sprintf("mode must be \"discovery\" or \"bootstrap\", got %q", mode),
		)
		return
	}

	credStore := config.RemoteCredentialStore.ValueString()
	if credStore == "" || config.RemoteCredentialStore.IsUnknown() {
		credStore = "local"
	}

	if mode == "discovery" {
		// Username/password only required when using local credential store
		if credStore == "local" {
			if config.DiscoveryUsername.IsNull() || config.DiscoveryUsername.ValueString() == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("discovery_username"),
					"Missing Required Attribute",
					"discovery_username is required when mode is \"discovery\" and remote_credential_store is \"local\"",
				)
			}
			if config.DiscoveryPassword.IsNull() || config.DiscoveryPassword.ValueString() == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("discovery_password"),
					"Missing Required Attribute",
					"discovery_password is required when mode is \"discovery\" and remote_credential_store is \"local\"",
				)
			}
		}
		if !config.BootstrapPassword.IsNull() && config.BootstrapPassword.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("bootstrap_password"),
				"Invalid Attribute",
				"bootstrap_password is not applicable when mode is \"discovery\"",
			)
		}
		if !config.UseNewCredentials.IsNull() && config.UseNewCredentials.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("use_new_credentials"),
				"Invalid Attribute",
				"use_new_credentials is only applicable when mode is \"bootstrap\"",
			)
		}
	}

	if mode == "bootstrap" {
		if config.BootstrapPassword.IsNull() || config.BootstrapPassword.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("bootstrap_password"),
				"Missing Required Attribute",
				"bootstrap_password is required when mode is \"bootstrap\"",
			)
		}

		if config.SwitchDetail.SerialNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("switch_detail").AtName("serial_number"),
				"Missing Required Attribute",
				"switch_detail.serial_number is required when mode is \"bootstrap\"",
			)
		}

		if !config.UseNewCredentials.IsNull() && config.UseNewCredentials.ValueBool() && credStore == "local" {
			if config.DiscoveryUsername.IsNull() || config.DiscoveryUsername.ValueString() == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("discovery_username"),
					"Missing Required Attribute",
					"discovery_username is required when use_new_credentials is true and remote_credential_store is \"local\"",
				)
			}
			if config.DiscoveryPassword.IsNull() || config.DiscoveryPassword.ValueString() == "" {
				resp.Diagnostics.AddAttributeError(
					path.Root("discovery_password"),
					"Missing Required Attribute",
					"discovery_password is required when use_new_credentials is true and remote_credential_store is \"local\"",
				)
			}
		}
		if !config.DiscoveryCredForLan.IsNull() && config.DiscoveryCredForLan.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root("discovery_cred_for_lan"),
				"Invalid Attribute",
				"discovery_cred_for_lan is only applicable when mode is \"discovery\"",
			)
		}
		if !config.SourceInterfaceName.IsNull() && config.SourceInterfaceName.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("source_interface_name"),
				"Invalid Attribute",
				"source_interface_name is only applicable when mode is \"discovery\"",
			)
		}
		if !config.SourceVrfName.IsNull() && config.SourceVrfName.ValueString() != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("source_vrf_name"),
				"Invalid Attribute",
				"source_vrf_name is only applicable when mode is \"discovery\"",
			)
		}
	}

	validateDuration := func(attr string, val string) {
		if _, err := time.ParseDuration(val); err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root(attr),
				"Invalid Duration",
				fmt.Sprintf("%s must be a valid duration (e.g. \"30m\", \"10m30s\"): %v", attr, err),
			)
		}
	}

	if !config.WaitForReady.IsNull() && !config.WaitForReady.IsUnknown() {
		validateDuration("wait_for_ready", config.WaitForReady.ValueString())
	}
	if !config.WaitForBootstrap.IsNull() && !config.WaitForBootstrap.IsUnknown() {
		validateDuration("wait_for_bootstrap", config.WaitForBootstrap.ValueString())
	}
	if !config.WaitForDiscover.IsNull() && !config.WaitForDiscover.IsUnknown() {
		validateDuration("wait_for_discover", config.WaitForDiscover.ValueString())
	}
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

	found := r.rscGetInventory(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	if !found {
		tflog.Warn(ctx, "Switch not found in fabric, removing from state", map[string]interface{}{
			"fabric_name": state.FabricName.ValueString(),
		})
		resp.State.RemoveResource(ctx)
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
