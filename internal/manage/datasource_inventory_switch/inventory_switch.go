// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_inventory_switch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

var (
	_ datasource.DataSource              = &inventorySwitchDataSource{}
	_ datasource.DataSourceWithConfigure = &inventorySwitchDataSource{}
)

// NewInventorySwitchDataSource creates the inventory switch data source.
func NewInventorySwitchDataSource() datasource.DataSource {
	return &inventorySwitchDataSource{}
}

type inventorySwitchDataSource struct {
	manageClient *manage.NexusDashboardManage
}

func (d *inventorySwitchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_switch"
}

func (d *inventorySwitchDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = InventorySwitchDataSourceSchema(ctx)
}

func (d *inventorySwitchDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.manageClient = manageClient
}

func (d *inventorySwitchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: inventory_switch")

	var data InventorySwitchModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Fabric Name Required",
			"The fabric_name attribute must contain a known value to read switch inventory.",
		)
		return
	}
	if d.manageClient == nil || d.manageClient.ApiClient == nil {
		resp.Diagnostics.AddError(
			"Manage Client Not Configured",
			"The manage API client is not available to read switch inventory.",
		)
		return
	}

	fabricName := data.FabricName.ValueString()
	log.Printf("[DEBUG] Reading switch inventory for fabric %s", fabricName)

	inventoryAPI := api.NewInventoryAPI(d.manageClient.ApiClient, ndapi.DefaultFabric)
	inventoryAPI.FabricName = fabricName
	serialFilter := ""
	if !data.SerialNumber.IsNull() && !data.SerialNumber.IsUnknown() {
		serialFilter = data.SerialNumber.ValueString()
	}

	var inventory NDFCInventorySwitchModel
	if serialFilter != "" {
		inventoryAPI.SerialNumber = serialFilter
		inventoryAPI.SetOperation(api.OpGetSwitch)
	} else {
		inventoryAPI.SetOperation(api.OpGetAllSwitches)
	}

	responseData, err := inventoryAPI.Get()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Switch Inventory",
			fmt.Sprintf("Could not read switches for fabric %q: %s", fabricName, err),
		)
		return
	}

	if serialFilter != "" {
		var switchValue NDFCSwitchesValue
		if err := json.Unmarshal(responseData, &switchValue); err != nil {
			resp.Diagnostics.AddError(
				"Error Decoding Switch Inventory",
				fmt.Sprintf("Could not decode switch %q in fabric %q: %s", serialFilter, fabricName, err),
			)
			return
		}
		inventory.Switches = append(inventory.Switches, switchValue)
	} else {
		if err := json.Unmarshal(responseData, &inventory); err != nil {
			resp.Diagnostics.AddError(
				"Error Decoding Switch Inventory",
				fmt.Sprintf("Could not decode switches for fabric %q: %s", fabricName, err),
			)
			return
		}
	}

	// The API response may omit these top-level query attributes. Keep them in
	// the Terraform state and let the generated mapper populate the response.
	inventory.FabricName = fabricName
	inventory.SerialNumber = serialFilter
	resp.Diagnostics.Append(data.SetModelData(&inventory)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource inventory_switch for fabric %s", fabricName)
}
