// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_remote_storage_location

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

var (
	_ datasource.DataSource              = &remoteStorageLocationNdDataSource{}
	_ datasource.DataSourceWithConfigure = &remoteStorageLocationNdDataSource{}
)

// NewRemoteStorageLocationDataSource returns a remote storage location datasource.
func NewRemoteStorageLocationDataSource() datasource.DataSource {
	return &remoteStorageLocationNdDataSource{}
}

type remoteStorageLocationNdDataSource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the datasource type name.
func (d *remoteStorageLocationNdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_remote_storage_location"
}

// Schema defines the schema for the datasource.
func (d *remoteStorageLocationNdDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RemoteStorageLocationDataSourceSchema(ctx)
}

// Configure adds the provider configured infra client to the datasource.
func (d *remoteStorageLocationNdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read retrieves a remote storage location by name and saves it in Terraform state.
func (d *remoteStorageLocationNdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_remote_storage_location")

	var data RemoteStorageLocationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Remote Storage Location Name Required",
			"The name attribute must contain a known, non-empty value to read an ND remote storage location.",
		)
		return
	}

	name := data.Name.ValueString()
	log.Printf("[DEBUG] Reading ND Remote Storage Location: name=%s", name)

	remoteStorageAPI := api.NewRemoteStorageLocationAPI(d.infraClient.ApiClient, ndapi.DefaultFabric)
	remoteStorageAPI.Name = name

	respData, err := remoteStorageAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading ND Remote Storage Location",
				fmt.Sprintf("Could not read nd remote storage location with name %q: resource not found", name),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not read nd remote storage location with name %q, unexpected error: %s", name, err.Error()),
		)
		return
	}

	remoteStorageResp, storageType, err := decodeRemoteStorageLocationDataSourceResponse(respData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading ND Remote Storage Location",
			fmt.Sprintf("Could not unmarshal nd remote storage location response with name %q, unexpected error: %s", name, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(data.SetModelData(remoteStorageResp)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The API discriminator selects the known block, even when its fields are omitted.
	if storageType == "nfs" {
		data.Nfs.state = attr.ValueStateKnown
		data.ScpSftp = NewScpSftpValueNull()
	} else {
		data.Nfs = NewNfsValueNull()
		data.ScpSftp.state = attr.ValueStateKnown
		// ND can omit the default false value from SCP/SFTP authentication readback.
		if data.ScpSftp.IgnoreHostKeyValidation.IsNull() {
			data.ScpSftp.IgnoreHostKeyValidation = types.BoolValue(false)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_remote_storage_location with name=%s", data.Name.ValueString())
}

// remoteStorageLocationDataSourceResponse represents the flat API storage specification.
type remoteStorageLocationDataSourceResponse struct {
	Name                string                   `json:"name"`
	Description         string                   `json:"description"`
	Hostname            string                   `json:"hostname"`
	Path                string                   `json:"path"`
	StorageLocationType string                   `json:"type"`
	Port                *int64                   `json:"port"`
	Limit               string                   `json:"limit"`
	ReadWrite           *bool                    `json:"readWrite"`
	AlertThreshold      *int64                   `json:"alertThreshold"`
	Authentication      *NDFCAuthenticationValue `json:"authentication"`
}

// decodeRemoteStorageLocationDataSourceResponse expands the API specification into
// the generated datasource model. It accepts wrapped and direct storage responses.
func decodeRemoteStorageLocationDataSourceResponse(response []byte) (*NDFCRemoteStorageLocationModel, string, error) {
	var wrapped struct {
		Spec   *remoteStorageLocationDataSourceResponse `json:"spec"`
		Status struct {
			HealthState string `json:"healthState"`
			Message     string `json:"message"`
		} `json:"status"`
	}
	if err := json.Unmarshal(response, &wrapped); err != nil {
		return nil, "", err
	}

	storageResponse := wrapped.Spec
	if storageResponse == nil {
		storageResponse = new(remoteStorageLocationDataSourceResponse)
		if err := json.Unmarshal(response, storageResponse); err != nil {
			return nil, "", err
		}
	}

	model := &NDFCRemoteStorageLocationModel{
		Name:               storageResponse.Name,
		Description:        storageResponse.Description,
		Hostname:           storageResponse.Hostname,
		Path:               storageResponse.Path,
		HealthState:        wrapped.Status.HealthState,
		HealthStateMessage: wrapped.Status.Message,
	}

	switch storageResponse.StorageLocationType {
	case "nfs":
		model.Nfs = NDFCNfsValue{
			Port:           storageResponse.Port,
			Limit:          storageResponse.Limit,
			ReadWrite:      storageResponse.ReadWrite,
			AlertThreshold: storageResponse.AlertThreshold,
		}
	case "scp", "sftp":
		if storageResponse.Authentication == nil {
			return nil, "", fmt.Errorf("%s remote storage response is missing authentication", storageResponse.StorageLocationType)
		}
		model.ScpSftp = NDFCScpSftpValue{
			Protocol:       storageResponse.StorageLocationType,
			Port:           storageResponse.Port,
			Authentication: *storageResponse.Authentication,
		}
	default:
		return nil, "", fmt.Errorf("unsupported or missing remote storage type %q", storageResponse.StorageLocationType)
	}

	return model, storageResponse.StorageLocationType, nil
}
