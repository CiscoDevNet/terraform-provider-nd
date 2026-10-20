// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_multi_cluster_connectivity

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &multiClusterConnectivityNdDataSource{}
	_ datasource.DataSourceWithConfigure = &multiClusterConnectivityNdDataSource{}
)

// NewMultiClusterConnectivityDataSource returns a multi-cluster connectivity
// datasource.
func NewMultiClusterConnectivityDataSource() datasource.DataSource {
	return &multiClusterConnectivityNdDataSource{}
}

type multiClusterConnectivityNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the datasource type name.
func (d *multiClusterConnectivityNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity"
}

// Schema defines the schema for the datasource.
func (d *multiClusterConnectivityNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = MultiClusterConnectivityDataSourceSchema(ctx)
}

// Configure adds the provider-configured infra client to the datasource.
func (d *multiClusterConnectivityNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves an ND cluster by name and saves it in Terraform state.
func (d *multiClusterConnectivityNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_multi_cluster_connectivity")

	var data MultiClusterConnectivityModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ClusterName.IsNull() || data.ClusterName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Cluster Name Required",
			"The cluster_name attribute must contain a known value to read an ND cluster.",
		)
		return
	}

	clusterName := data.ClusterName.ValueString()
	log.Printf("[DEBUG] Reading ND cluster: cluster_name=%s", clusterName)

	clusterAPI := api.NewClusterAPI(d.infraClient.ApiClient)
	clusterAPI.ClusterName = clusterName

	respData, err := clusterAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading ND Multi-Cluster Connectivity",
				fmt.Sprintf("Could not read nd multi-cluster connectivity with cluster_name %q: resource not found", clusterName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Multi-Cluster Connectivity",
			fmt.Sprintf("Could not read nd multi-cluster connectivity with cluster_name %q, unexpected error: %s %s", clusterName, err.Error(), string(respData)),
		)
		return
	}

	var clusterResp NDFCMultiClusterConnectivityModel
	if err := json.Unmarshal(respData, &clusterResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Multi-Cluster Connectivity",
			fmt.Sprintf("Could not unmarshal nd multi-cluster connectivity response with cluster_name %q, unexpected error: %s", clusterName, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(&clusterResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_multi_cluster_connectivity with cluster_name=%s", clusterName)
}
