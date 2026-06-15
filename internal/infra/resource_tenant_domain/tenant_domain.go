// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_tenant_domain

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &tenantDomainResource{}
	_ resource.ResourceWithConfigure   = &tenantDomainResource{}
	_ resource.ResourceWithImportState = &tenantDomainResource{}
)

// NewTenantDomainResource is a helper function to simplify the provider implementation.
func NewTenantDomainResource() resource.Resource {
	return &tenantDomainResource{}
}

// tenantDomainResource is the resource implementation.
type tenantDomainResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *tenantDomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_domain"
}

// Schema defines the schema for the resource.
func (r *tenantDomainResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = TenantDomainResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *tenantDomainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	infraModule := client.GetModule(ModuleKey)
	if infraModule == nil {
		resp.Diagnostics.AddError(
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	infraClient, ok := infraModule.(*infra.NexusDashboardInfra)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Infra Module Type",
			fmt.Sprintf("Expected *infra.NexusDashboardInfra, got: %T. Please report this issue to the provider developers.", infraModule),
		)
		return
	}

	r.infraClient = infraClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *tenantDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_tenant_domain")

	var in TenantDomainModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating Tenant Domain: name=%s", in.Name.ValueString())

	r.rscCreateTenantDomain(&resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_tenant_domain with id '%s'", in.Id.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *tenantDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_tenant_domain")

	var state TenantDomainModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.Id.ValueString()
	log.Printf("[DEBUG] Reading Tenant Domain: id=%s", id)

	if r.rscGetTenantDomain(&resp.Diagnostics, &state) {
		log.Printf("[DEBUG] Tenant domain %q not found, removing from state", id)
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_tenant_domain with id '%s'", state.Id.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *tenantDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_tenant_domain")

	var plan TenantDomainModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TenantDomainModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id
	log.Printf("[DEBUG] Updating Tenant Domain: id=%s", plan.Id.ValueString())

	r.rscUpdateTenantDomain(&resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] End update of resource nd_tenant_domain with id '%s'", plan.Id.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *tenantDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_tenant_domain")

	var state TenantDomainModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.Id.ValueString()
	r.rscDeleteTenantDomain(&resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End delete of resource nd_tenant_domain with id '%s'", id)
}

// ImportState imports a tenant domain resource by name.
func (r *tenantDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_tenant_domain")

	var state TenantDomainModel
	state.Name = types.StringValue(req.ID)
	state.Id = state.Name

	if r.rscGetTenantDomain(&resp.Diagnostics, &state) {
		resp.Diagnostics.AddError(
			"Error Importing Tenant Domain",
			fmt.Sprintf("Tenant domain %q was not found", req.ID),
		)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource nd_tenant_domain with id '%s'", state.Id.ValueString())
}
