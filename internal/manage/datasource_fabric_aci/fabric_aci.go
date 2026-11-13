// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_fabric_aci

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

var (
	_ datasource.DataSource              = &fabricAciDataSource{}
	_ datasource.DataSourceWithConfigure = &fabricAciDataSource{}
)

// NewFabricAciDataSource returns a Fabric ACI datasource.
func NewFabricAciDataSource() datasource.DataSource {
	return &fabricAciDataSource{}
}

type fabricAciDataSource struct {
	manageClient *manage.NexusDashboardManage
}

// Metadata returns the datasource type name.
func (d *fabricAciDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_aci"
}

// Schema defines the schema for the datasource.
func (d *fabricAciDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = FabricAciDataSourceSchema(ctx)
}

// Configure adds the provider configured manage client to the datasource.
func (d *fabricAciDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves a Fabric ACI cluster by name and stores its returned fields in state.
func (d *fabricAciDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_fabric_aci")

	var data FabricAciModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Fabric ACI Name Required",
			"The fabric_name attribute must contain a known value to read a Fabric ACI cluster.",
		)
		return
	}

	fabricName := data.FabricName.ValueString()
	log.Printf("[DEBUG] Reading ND Fabric ACI: fabric_name=%s", fabricName)

	clusterAPI := api.NewFabricAciAPI(d.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = fabricName
	respData, err := clusterAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading ND Fabric ACI",
				fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q: resource not found", fabricName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q, unexpected error: %s %s", fabricName, err.Error(), string(respData)),
		)
		return
	}
	if len(respData) == 0 {
		resp.Diagnostics.AddError(
			"Error Reading ND Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q: empty response", fabricName),
		)
		return
	}

	var fabricAciResp fabricAciDataSourceResponse
	if err := json.Unmarshal(respData, &fabricAciResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Fabric ACI",
			fmt.Sprintf("Could not unmarshal nd_fabric_aci response with fabric_name %q: %s", fabricName, err.Error()),
		)
		return
	}

	modelData := fabricAciResp.modelData(fabricName)
	// The cluster API can return the manage-style camel-cased telemetry values,
	// while the Terraform schema exposes the canonical lowercase values.
	switch modelData.Spec.Aci.Telemetry.Network {
	case "inBand":
		modelData.Spec.Aci.Telemetry.Network = "inband"
	case "outOfBand":
		modelData.Spec.Aci.Telemetry.Network = "outband"
	}
	resp.Diagnostics.Append(data.SetModelData(&modelData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_fabric_aci with fabric_name=%s", data.FabricName.ValueString())
}

// fabricAciDataSourceResponse models the API response, where the cluster name
// can be returned as spec.name instead of spec.aci.name.
type fabricAciDataSourceResponse struct {
	Spec struct {
		NDFCSpecValue
		FabricName string `json:"name,omitempty"`
	} `json:"spec,omitempty"`
	Status struct {
		State      string `json:"state,omitempty"`
		LastUpdate struct {
			Message string `json:"message,omitempty"`
		} `json:"lastUpdate,omitempty"`
	} `json:"status,omitempty"`
}

func (r fabricAciDataSourceResponse) modelData(requestedFabricName string) NDFCFabricAciModel {
	fabricName := r.Spec.FabricName
	if fabricName == "" {
		fabricName = requestedFabricName
	}

	spec := r.Spec.NDFCSpecValue
	if spec.Aci.FabricName == "" {
		spec.Aci.FabricName = fabricName
	}

	return NDFCFabricAciModel{
		Spec: spec,
		Status: NDFCStatusValue{
			State: r.Status.State,
			LastUpdate: NDFCStatusLastUpdateValue{
				LastUpdateMessage: r.Status.LastUpdate.Message,
			},
		},
	}
}
