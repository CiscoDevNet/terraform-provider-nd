// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_change_control

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &changeControlNdDataSource{}
	_ datasource.DataSourceWithConfigure = &changeControlNdDataSource{}
)

// NewChangeControlDataSource returns a change control datasource.
func NewChangeControlDataSource() datasource.DataSource {
	return &changeControlNdDataSource{}
}

type changeControlNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the datasource type name.
func (d *changeControlNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_change_control"
}

// Schema defines the schema for the datasource.
func (d *changeControlNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ChangeControlDataSourceSchema(ctx)
}

// Configure adds the provider configured infra client to the datasource.
func (d *changeControlNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves the singleton change control settings and saves them in Terraform state.
func (d *changeControlNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_change_control")

	var data ChangeControlModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	changeControlAPI := api.NewChangeControlAPI(d.infraClient.ApiClient, ndapi.DefaultFabric)

	respData, err := changeControlAPI.Get()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Change Control",
			fmt.Sprintf("Could not read change control settings, unexpected error: %s", err.Error()),
		)
		return
	}

	var changeControlResp NDFCChangeControlModel
	if err := json.Unmarshal(respData, &changeControlResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Change Control",
			fmt.Sprintf("Could not unmarshal change control response, unexpected error: %s", err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(&changeControlResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource: nd_change_control")
}
