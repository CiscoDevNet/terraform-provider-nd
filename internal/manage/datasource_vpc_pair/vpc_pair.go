// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_vpc_pair

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

const ModuleKey = "manage"

var (
	_ datasource.DataSource              = &vpcPairDataSource{}
	_ datasource.DataSourceWithConfigure = &vpcPairDataSource{}
)

// NewVpcPairDataSource creates a data source for reading a vPC pair.
func NewVpcPairDataSource() datasource.DataSource {
	return &vpcPairDataSource{}
}

type vpcPairDataSource struct {
	manageClient *manage.NexusDashboardManage
}

func (d *vpcPairDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_pair"
}

func (d *vpcPairDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpcPairDataSourceSchema(ctx)
}

func (d *vpcPairDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vpcPairDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpcPairModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if d.manageClient == nil || d.manageClient.ApiClient == nil {
		resp.Diagnostics.AddError(
			"Manage Client Not Configured",
			"The manage API client is not available to read a vPC pair.",
		)
		return
	}

	fabricName, ok := requiredDataSourceString(data.FabricName, "fabric_name", &resp.Diagnostics)
	if !ok {
		return
	}
	switchID, ok := requiredDataSourceString(data.SwitchId, "switch_id", &resp.Diagnostics)
	if !ok {
		return
	}

	pairAPI := api.NewVpcPairAPI(d.manageClient.ApiClient, fabricName)
	pairAPI.FabricName = fabricName
	pairAPI.SwitchID = switchID

	responseData, err := pairAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"vPC Pair Not Found",
				fmt.Sprintf("No vPC pair was found for switch %q in fabric %q.", switchID, fabricName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("Could not read vPC pair for switch %q in fabric %q: %s", switchID, fabricName, err),
		)
		return
	}

	if len(responseData) == 0 {
		resp.Diagnostics.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("The vPC pair response for switch %q in fabric %q was empty.", switchID, fabricName),
		)
		return
	}

	var response NDFCVpcPairModel
	if err := json.Unmarshal(responseData, &response); err != nil {
		resp.Diagnostics.AddError(
			"Error Decoding vPC Pair",
			fmt.Sprintf("Could not decode the vPC pair response for switch %q in fabric %q: %s", switchID, fabricName, err),
		)
		return
	}

	response.FabricName = fabricName
	if response.SwitchId == "" {
		response.SwitchId = switchID
	}

	if strings.TrimSpace(response.PeerSwitchId) == "" {
		resp.Diagnostics.AddError(
			"Error Decoding vPC Pair",
			fmt.Sprintf("The vPC pair response for switch %q in fabric %q did not include peerSwitchId.", switchID, fabricName),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(&response)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func requiredDataSourceString(value interface {
	IsNull() bool
	IsUnknown() bool
	ValueString() string
}, attributeName string, diagnostics *diag.Diagnostics) (string, bool) {
	if value.IsNull() || value.IsUnknown() || strings.TrimSpace(value.ValueString()) == "" {
		diagnostics.AddError(
			"Required Attribute Missing",
			fmt.Sprintf("The %q attribute must contain a known, non-empty value to read a vPC pair.", attributeName),
		)
		return "", false
	}

	return value.ValueString(), true
}
