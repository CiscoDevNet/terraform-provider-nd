// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_aci

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                   = &fabricAciResource{}
	_ resource.ResourceWithConfigure      = &fabricAciResource{}
	_ resource.ResourceWithImportState    = &fabricAciResource{}
	_ resource.ResourceWithValidateConfig = &fabricAciResource{}
)

// NewFabricAciResource is a helper function to simplify the provider implementation.
func NewFabricAciResource() resource.Resource {
	return &fabricAciResource{}
}

// fabricAciResource is the resource implementation.
type fabricAciResource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the resource type name.
func (r *fabricAciResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_aci"
}

// Schema defines the schema for the resource.
func (r *fabricAciResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricAciResourceSchema(ctx)
}

// ValidateConfig enforces conditional validation that generated schema
// validators cannot express.
func (r *fabricAciResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config FabricAciModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.Telemetry.IsNull() && !config.Telemetry.IsUnknown() {
		if config.Telemetry.Network.ValueString() == "outband" {
			if !config.Telemetry.Epg.IsNull() && !config.Telemetry.Epg.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("telemetry").AtName("epg"),
					"Invalid attribute for telemetry network",
					"Attribute `telemetry.epg` can only be set when `telemetry.network` is `inband`. Remove `telemetry.epg` when `telemetry.network` is `outband`.",
				)
			}
		}
	}
}

// Configure adds the provider configured client to the resource.
func (r *fabricAciResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule),
		)
		return
	}

	r.manageClient = manageClient
}

func fabricAciCredentialAvailable(value types.String) bool {
	return !value.IsNull() && !value.IsUnknown() && value.ValueString() != ""
}

func fabricAciEnvironmentVariableName(fabricName string, variableName string) string {
	prefix := strings.ToUpper(strings.ReplaceAll(fabricName, "-", "_"))
	return prefix + "_" + variableName
}

// Create creates the resource and sets the initial Terraform state.
func (r *fabricAciResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_fabric_aci")

	var in FabricAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	in.Id = in.FabricName
	log.Printf("[DEBUG] Creating Fabric ACI: id=%s", in.Id.ValueString())

	r.rscCreateFabricAci(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_fabric_aci with id '%s'", in.Id.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *fabricAciResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_fabric_aci")

	var state FabricAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.Id.ValueString()
	log.Printf("[DEBUG] Reading Fabric ACI: id=%s", id)

	if r.rscGetFabricAci(ctx, &resp.Diagnostics, &state) {
		log.Printf("[DEBUG] Fabric ACI %q not found, removing from state", id)
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_fabric_aci with id '%s'", state.Id.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *fabricAciResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_fabric_aci")

	var plan FabricAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = plan.FabricName
	log.Printf("[DEBUG] Updating Fabric ACI: id=%s", plan.Id.ValueString())

	r.rscUpdateFabricAci(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] End update of resource nd_fabric_aci with id '%s'", plan.Id.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *fabricAciResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_fabric_aci")

	var state FabricAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !fabricAciCredentialAvailable(state.Username) || !fabricAciCredentialAvailable(state.Password) {
		resp.Diagnostics.AddError(
			"Missing Fabric ACI Delete Credentials",
			fmt.Sprintf(
				"Cannot delete imported nd_fabric_aci %q because its username and password are not stored in Terraform state. Keep the resource configuration present and run terraform apply so Terraform performs the regular update and stores the configured credentials in state, then run terraform destroy.",
				state.Id.ValueString(),
			),
		)
		return
	}

	id := state.Id.ValueString()
	force := false
	forceEnvironmentVariable := fabricAciEnvironmentVariableName(id, "FORCE")
	if forceValue, ok := os.LookupEnv(forceEnvironmentVariable); ok {
		var err error
		force, err = strconv.ParseBool(forceValue)
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid Fabric ACI Force Environment Variable",
				fmt.Sprintf("Environment variable %s must contain a valid Boolean value: %v", forceEnvironmentVariable, err),
			)
			return
		}
	}

	r.rscDeleteFabricAci(ctx, &resp.Diagnostics, &state, force)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End delete of resource nd_fabric_aci with id '%s'", id)
}

// ImportState imports a fabric ACI resource by id.
func (r *fabricAciResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_fabric_aci")

	var state FabricAciModel
	state.Id = types.StringValue(req.ID)
	state.FabricName = state.Id

	// Setting default value false to VerifyCa because the API does not return it.
	state.VerifyCa = types.BoolValue(false)

	if r.rscGetFabricAci(ctx, &resp.Diagnostics, &state) {
		resp.Diagnostics.AddError(
			"Error Importing Fabric ACI",
			fmt.Sprintf("Could not import nd_fabric_aci with id %q: resource not found", state.Id.ValueString()),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// The API does not return credentials. Import each value only when its
	// fabric-scoped environment variable is present.
	if username, ok := os.LookupEnv(fabricAciEnvironmentVariableName(req.ID, "USERNAME")); ok {
		state.Username = types.StringValue(username)
	}
	if password, ok := os.LookupEnv(fabricAciEnvironmentVariableName(req.ID, "PASSWORD")); ok {
		state.Password = types.StringValue(password)
	}
	if loginDomain, ok := os.LookupEnv(fabricAciEnvironmentVariableName(req.ID, "LOGIN_DOMAIN")); ok {
		state.LoginDomain = types.StringValue(loginDomain)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_fabric_aci with id '%s'", state.Id.ValueString())
}
