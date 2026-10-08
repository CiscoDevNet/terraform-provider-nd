// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_fabric

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

// Ensure the implementation satisfies the expected interfaces
var (
	_ datasource.DataSource              = &fabricDataSource{}
	_ datasource.DataSourceWithConfigure = &fabricDataSource{}
	_ resource_fabric_common.FabricModel = (*FabricModel)(nil)
)

// NewFabricDataSource is a helper function to simplify the provider implementation.
func NewFabricDataSource() datasource.DataSource {
	return &fabricDataSource{}
}

// fabricDataSource is the datasource implementation.
type fabricDataSource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the datasource type name.
func (r *fabricDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric"
}

// Schema defines the schema for the datasource.
func (r *fabricDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = FabricDataSourceSchema(ctx)
}

// GetFabricName returns the fabric name used to look up the data source.
func (m *FabricModel) GetFabricName() string {
	if m.FabricName.IsNull() || m.FabricName.IsUnknown() {
		return ""
	}
	return m.FabricName.ValueString()
}

// GetFabricType returns the hardcoded default handler identifier for the datasource.
func (m *FabricModel) GetFabricType() string {
	return ""
}

// Configure adds the provider configured client to the datasource.
func (d *fabricDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	manageModule := client.GetModule("manage")
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
	d.manageClient = manageClient
}

// Read retrieves an fabric by name and saves it in Terraform state.
func (d *fabricDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: fabric")

	var data FabricModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Fabric Name Required",
			"The fabric_name attribute must contain a known value to read a fabric.",
		)
		return
	}

	fabricName := data.GetFabricName()
	log.Printf("[DEBUG] Reading fabric: name=%s", fabricName)

	found := resource_fabric_common.RscGetFabric(ctx, d.manageClient.ApiClient, &resp.Diagnostics, &data)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError(
			"Error Reading fabric",
			fmt.Sprintf("Could not read fabric with name %q: resource not found", fabricName),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource fabric with name=%s", data.FabricName.ValueString())
}
