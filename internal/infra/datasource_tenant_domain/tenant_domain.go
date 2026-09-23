// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_tenant_domain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &tenantDomainNdDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantDomainNdDataSource{}
)

// NewTenantDomainDataSource returns a tenant-domain datasource.
func NewTenantDomainDataSource() datasource.DataSource {
	return &tenantDomainNdDataSource{}
}

type tenantDomainNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the datasource type name.
func (d *tenantDomainNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_domain"
}

// Schema defines the schema for the datasource.
func (d *tenantDomainNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = TenantDomainDataSourceSchema(ctx)
}

// Configure adds the provider-configured infra client to the datasource.
func (d *tenantDomainNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
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

	d.infraClient = infraClient
}

// Read retrieves a tenant domain by name and saves it in Terraform state.
func (d *tenantDomainNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_tenant_domain")

	var data TenantDomainModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Name.IsNull() || data.Name.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant Domain Name Required",
			"The name attribute must contain a known value to read an ND tenant domain.",
		)
		return
	}

	tenantDomainName := data.Name.ValueString()
	log.Printf("[DEBUG] Reading ND Tenant Domain: name=%s", tenantDomainName)

	tenantDomainAPI := api.NewTenantDomainAPI(d.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantDomainAPI.TenantDomainName = tenantDomainName

	respData, err := tenantDomainAPI.Get()
	err = ndapi.ClassifyRequestError(http.MethodGet, tenantDomainAPI.GetUrl(), respData, err)
	if err != nil {
		if errors.Is(err, ndapi.ErrNotFound) {
			resp.Diagnostics.AddError(
				"Error Reading ND Tenant Domain",
				fmt.Sprintf("Could not read nd tenant domain with name %q: resource not found", tenantDomainName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Tenant Domain",
			fmt.Sprintf("Could not read nd tenant domain with name %q, unexpected error: %s %s", tenantDomainName, err.Error(), string(respData)),
		)
		return
	}

	var tenantDomainResp NDFCTenantDomainModel
	if err := json.Unmarshal(respData, &tenantDomainResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Tenant Domain",
			fmt.Sprintf("Could not unmarshal nd tenant domain response with name %q, unexpected error: %s", tenantDomainName, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(&tenantDomainResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_tenant_domain with name=%s", data.Name.ValueString())
}
