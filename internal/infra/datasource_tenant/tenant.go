// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_tenant

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra"
	infraapi "terraform-provider-nd/internal/infra/api"
	manageapi "terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &tenantNdDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantNdDataSource{}
)

// NewTenantDataSource returns a tenant datasource.
func NewTenantDataSource() datasource.DataSource {
	return &tenantNdDataSource{}
}

type tenantNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

type tenantFabricAssociationListResponse struct {
	TenantFabricAssociations []tenantFabricAssociationItem `json:"tenantFabricAssociations"`
}

type tenantFabricAssociationItem struct {
	AllowedVlans []string `json:"allowedVlans,omitempty"`
	FabricName   string   `json:"fabricName"`
	LocalName    string   `json:"localName,omitempty"`
	TenantName   string   `json:"tenantName"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
}

// Metadata returns the datasource type name.
func (d *tenantNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

// Schema defines the schema for the datasource.
func (d *tenantNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = TenantDataSourceSchema(ctx)
}

// Configure adds the provider configured infra client to the datasource.
func (d *tenantNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves a tenant by name, including all its fabric associations, and
// saves it in Terraform state.
func (d *tenantNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_tenant")

	var data TenantModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Name.IsNull() || data.Name.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant Name Required",
			"The name attribute must contain a known value to read an ND tenant.",
		)
		return
	}

	tenantName := data.Name.ValueString()
	log.Printf("[DEBUG] Reading ND Tenant: name=%s", tenantName)

	tenantAPI := infraapi.NewTenantAPI(d.infraClient.ApiClient, ndapi.DefaultFabric)
	tenantAPI.TenantName = tenantName

	respData, err := tenantAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading ND Tenant",
				fmt.Sprintf("Could not read nd tenant with name %q: resource not found", tenantName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Tenant",
			fmt.Sprintf("Could not read nd tenant with name %q, unexpected error: %s %s", tenantName, err.Error(), string(respData)),
		)
		return
	}

	var tenantResp NDFCTenantModel
	if err := json.Unmarshal(respData, &tenantResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Tenant",
			fmt.Sprintf("Could not unmarshal nd tenant response with name %q, unexpected error: %s", tenantName, err.Error()),
		)
		return
	}

	tenantFabricAssociationAPI := manageapi.NewTenantFabricAssociationAPI(d.infraClient.ApiClient, ndapi.DefaultFabric)
	associationRespData, err := tenantFabricAssociationAPI.Get()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tenant Fabric Associations",
			fmt.Sprintf("Could not read fabric associations for nd tenant with name %q, unexpected error: %s %s", tenantName, err.Error(), string(associationRespData)),
		)
		return
	}

	var associationResp tenantFabricAssociationListResponse
	if err := json.Unmarshal(associationRespData, &associationResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tenant Fabric Associations",
			fmt.Sprintf("Could not unmarshal fabric associations for nd tenant with name %q, unexpected error: %s", tenantName, err.Error()),
		)
		return
	}

	tenantResp.FabricAssociations = make(map[string]NDFCFabricAssociationsValue)
	for _, association := range associationResp.TenantFabricAssociations {
		if association.TenantName != tenantName {
			continue
		}

		tenantResp.FabricAssociations[association.FabricName] = NDFCFabricAssociationsValue{
			AllowedVlans: association.AllowedVlans,
			LocalName:    association.LocalName,
			TenantPrefix: association.TenantPrefix,
		}
	}

	resp.Diagnostics.Append(data.SetModelData(&tenantResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_tenant with name=%s", data.Name.ValueString())
}
