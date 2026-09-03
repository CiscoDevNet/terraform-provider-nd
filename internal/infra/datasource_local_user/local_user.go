// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_local_user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_local_user"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &localUserNdDataSource{}
	_ datasource.DataSourceWithConfigure = &localUserNdDataSource{}
)

// NewLocalUserDataSource returns a local user datasource.
func NewLocalUserDataSource() datasource.DataSource {
	return &localUserNdDataSource{}
}

type localUserNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the datasource type name.
func (d *localUserNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_local_user"
}

// Schema defines the schema for the datasource.
func (d *localUserNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LocalUserDataSourceSchema(ctx)
}

// Configure adds the provider configured infra client to the datasource.
func (d *localUserNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves a local user by login ID and saves it in Terraform state.
func (d *localUserNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_local_user")

	var data resource_local_user.LocalUserModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.LoginId.IsNull() || data.LoginId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Local User Login ID Required",
			"The login_id attribute must contain a known value to read an ND local user.",
		)
		return
	}

	id := data.LoginId.ValueString()
	log.Printf("[DEBUG] Reading ND Local User: id=%s", id)

	localUserAPI := api.NewLocalUserAPI(d.infraClient.ApiClient)
	localUserAPI.LoginId = id

	respData, err := localUserAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading ND Local User",
				fmt.Sprintf("Could not read nd local user with id %q: resource not found", id),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not read nd local user with id %q, unexpected error: %s %s", id, err.Error(), string(respData)),
		)
		return
	}

	var localUserResp resource_local_user.NDFCLocalUserModel
	if err := json.Unmarshal(respData, &localUserResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Local User",
			fmt.Sprintf("Could not unmarshal nd local user response with id %q, unexpected error: %s", id, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(&localUserResp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// The data source ID is Terraform-only and is derived from the API login ID.
	data.Id = data.LoginId

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_local_user with id=%s", data.Id.ValueString())
}
