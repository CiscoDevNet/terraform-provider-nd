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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type inventorySwitchDataSourceConfig struct {
	FabricName   types.String `tfsdk:"fabric_name"`
	SerialNumber types.String `tfsdk:"serial_number"`
	ManagementIP types.String `tfsdk:"management_ip"`
	Switches     types.List   `tfsdk:"switches"`
}

type inventorySwitchDataSourceState struct {
	FabricName   types.String `tfsdk:"fabric_name"`
	SerialNumber types.String `tfsdk:"serial_number"`
	ManagementIP types.String `tfsdk:"management_ip"`
	Switches     types.List   `tfsdk:"switches"`
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
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	module := client.GetModule(ModuleKey)
	if module == nil {
		resp.Diagnostics.AddError("Manage Module Not Found", "The manage module was not registered with the provider.")
		return
	}

	manageClient, ok := module.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", module),
		)
		return
	}

	d.manageClient = manageClient
}

func (d *inventorySwitchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: inventory_switch")

	var config inventorySwitchDataSourceConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.FabricName.IsNull() || config.FabricName.IsUnknown() || config.FabricName.ValueString() == "" {
		resp.Diagnostics.AddError("Fabric Name Required", "The fabric_name attribute must contain a known value to read switches.")
		return
	}
	if config.SerialNumber.IsUnknown() || config.ManagementIP.IsUnknown() {
		resp.Diagnostics.AddError("Switch Selector Unknown", "serial_number and management_ip must be known values when configured.")
		return
	}
	if d.manageClient == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The Nexus Dashboard client is not configured.")
		return
	}

	fabricName := config.FabricName.ValueString()
	serialNumber := ""
	if !config.SerialNumber.IsNull() {
		serialNumber = config.SerialNumber.ValueString()
	}
	managementIP := ""
	if !config.ManagementIP.IsNull() {
		managementIP = config.ManagementIP.ValueString()
	}

	var matchingSwitches []inventorySwitchAPIEntry
	if serialNumber != "" {
		log.Printf("[DEBUG] Reading switch with serial %s from fabric: %s", serialNumber, fabricName)
		invAPI := api.NewInventoryAPI(d.manageClient.ApiClient, ndapi.DefaultFabric).
			WithFabric(fabricName).
			WithSerial(serialNumber).
			SetOperation(api.OpGetSwitch)
		responseBody, err := invAPI.Get()
		if err != nil {
			resp.Diagnostics.AddError("Error Reading Inventory Switch", fmt.Sprintf("Could not read switch with serial_number %q from fabric %q: %s", serialNumber, fabricName, err))
			return
		}

		var entry inventorySwitchAPIEntry
		if err := json.Unmarshal(responseBody, &entry); err != nil {
			resp.Diagnostics.AddError("Error Decoding Inventory Switch", fmt.Sprintf("Could not decode switch response: %s", err))
			return
		}
		matchingSwitches = filterInventorySwitches([]inventorySwitchAPIEntry{entry}, serialNumber, managementIP)
	} else {
		log.Printf("[DEBUG] Reading switches for fabric: %s", fabricName)
		invAPI := api.NewInventoryAPI(d.manageClient.ApiClient, ndapi.DefaultFabric).
			WithFabric(fabricName).
			SetOperation(api.OpGetAllSwitches)
		responseBody, err := invAPI.Get()
		if err != nil {
			resp.Diagnostics.AddError("Error Reading Inventory Switches", fmt.Sprintf("Could not read switches for fabric %q: %s", fabricName, err))
			return
		}

		var response inventorySwitchesResponse
		if err := json.Unmarshal(responseBody, &response); err != nil {
			resp.Diagnostics.AddError("Error Decoding Inventory Switches", fmt.Sprintf("Could not decode switch response: %s", err))
			return
		}
		matchingSwitches = filterInventorySwitches(response.Switches, "", managementIP)
	}

	if (serialNumber != "" || managementIP != "") && len(matchingSwitches) == 0 {
		resp.Diagnostics.AddError(
			"Inventory Switch Not Found",
			fmt.Sprintf("No switch matched fabric %q with serial_number %q and management_ip %q.", fabricName, serialNumber, managementIP),
		)
		return
	}

	switches, diags := inventorySwitchesList(matchingSwitches)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state := inventorySwitchDataSourceState{
		FabricName:   config.FabricName,
		SerialNumber: config.SerialNumber,
		ManagementIP: config.ManagementIP,
		Switches:     switches,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Read %d switches for fabric %s after applying selectors", len(matchingSwitches), fabricName)
}

// filterInventorySwitches applies the optional selectors locally because the
// fabric switches API ignores the management IP query parameter.
func filterInventorySwitches(switches []inventorySwitchAPIEntry, serialNumber, managementIP string) []inventorySwitchAPIEntry {
	filtered := make([]inventorySwitchAPIEntry, 0, len(switches))
	for _, entry := range switches {
		if serialNumber != "" && entry.SerialNumber != serialNumber {
			continue
		}
		entryManagementIP := firstNonEmpty(entry.FabricManagementIP, entry.IPAddress)
		if managementIP != "" && entryManagementIP != managementIP {
			continue
		}
		filtered = append(filtered, entry)
	}
	return filtered
}

type inventorySwitchesResponse struct {
	Switches []inventorySwitchAPIEntry `json:"switches,omitempty"`
}

// This response model intentionally excludes additionalData.username and
// internal additionalData.meta values from Terraform state.
type inventorySwitchAPIEntry struct {
	FabricName         string                        `json:"fabricName,omitempty"`
	SerialNumber       string                        `json:"serialNumber,omitempty"`
	FabricManagementIP string                        `json:"fabricManagementIp,omitempty"`
	IPAddress          string                        `json:"ip,omitempty"`
	Hostname           string                        `json:"hostname,omitempty"`
	Model              string                        `json:"model,omitempty"`
	SoftwareVersion    string                        `json:"softwareVersion,omitempty"`
	SoftwareImage      string                        `json:"softwareImage,omitempty"`
	SwitchRole         string                        `json:"switchRole,omitempty"`
	PreserveConfig     *bool                         `json:"preserveConfig,omitempty"`
	VDCID              *int64                        `json:"vdcId,omitempty"`
	VDCMAC             string                        `json:"vdcMac,omitempty"`
	InInventory        *bool                         `json:"inInventory,omitempty"`
	SwitchID           string                        `json:"switchId,omitempty"`
	FabricType         string                        `json:"fabricType,omitempty"`
	SystemUptime       string                        `json:"systemUpTime,omitempty"`
	AdvisoryLevel      string                        `json:"advisoryLevel,omitempty"`
	AlertSuspend       string                        `json:"alertSuspend,omitempty"`
	AnomalyLevel       string                        `json:"anomalyLevel,omitempty"`
	VpcConfigured      *bool                         `json:"vpcConfigured,omitempty"`
	AdditionalData     inventorySwitchAdditionalData `json:"additionalData,omitempty"`
}

type inventorySwitchAdditionalData struct {
	Status                        string `json:"discoveryStatus,omitempty"`
	PlatformType                  string `json:"platformType,omitempty"`
	SourceInterfaceName           string `json:"sourceInterfaceName,omitempty"`
	SourceVRFName                 string `json:"sourceVrfName,omitempty"`
	ConfigSyncStatus              string `json:"configSyncStatus,omitempty"`
	DiscoveredSystemMode          string `json:"discoveredSystemMode,omitempty"`
	IntendedSystemMode            string `json:"intendedSystemMode,omitempty"`
	SystemMode                    string `json:"systemMode,omitempty"`
	DomainName                    string `json:"domainName,omitempty"`
	HypershieldConnectivityStatus string `json:"hypershieldConnectivityStatus,omitempty"`
	HypershieldTenant             string `json:"hypershieldTenant,omitempty"`
	RemoteCredentialStore         string `json:"remoteCredentialStore,omitempty"`
	ScalableUnit                  string `json:"scalableUnit,omitempty"`
	SmartSwitch                   *bool  `json:"smartSwitch,omitempty"`
	Usage                         string `json:"usage,omitempty"`
	Vendor                        string `json:"vendor,omitempty"`
}

func inventorySwitchesList(apiSwitches []inventorySwitchAPIEntry) (types.List, diag.Diagnostics) {
	attributeTypes := inventorySwitchAttributeTypes()
	objectType := types.ObjectType{AttrTypes: attributeTypes}
	values := make([]attr.Value, 0, len(apiSwitches))

	for _, entry := range apiSwitches {
		attributes := map[string]attr.Value{
			"fabric_name":                     stringValue(entry.FabricName),
			"serial_number":                   stringValue(entry.SerialNumber),
			"ip_address":                      stringValue(firstNonEmpty(entry.FabricManagementIP, entry.IPAddress)),
			"hostname":                        stringValue(entry.Hostname),
			"model":                           stringValue(entry.Model),
			"software_version":                stringValue(entry.SoftwareVersion),
			"software_image":                  stringValue(entry.SoftwareImage),
			"switch_role":                     stringValue(entry.SwitchRole),
			"status":                          stringValue(entry.AdditionalData.Status),
			"platform_type":                   stringValue(entry.AdditionalData.PlatformType),
			"preserve_config":                 boolValue(entry.PreserveConfig),
			"vdc_id":                          int64Value(entry.VDCID),
			"vdc_mac":                         stringValue(entry.VDCMAC),
			"source_interface_name":           stringValue(entry.AdditionalData.SourceInterfaceName),
			"source_vrf_name":                 stringValue(entry.AdditionalData.SourceVRFName),
			"in_inventory":                    boolValue(entry.InInventory),
			"switch_id":                       stringValue(entry.SwitchID),
			"fabric_type":                     stringValue(entry.FabricType),
			"system_uptime":                   stringValue(entry.SystemUptime),
			"advisory_level":                  stringValue(entry.AdvisoryLevel),
			"alert_suspend":                   stringValue(entry.AlertSuspend),
			"anomaly_level":                   stringValue(entry.AnomalyLevel),
			"vpc_configured":                  boolValue(entry.VpcConfigured),
			"config_sync_status":              stringValue(entry.AdditionalData.ConfigSyncStatus),
			"discovered_system_mode":          stringValue(entry.AdditionalData.DiscoveredSystemMode),
			"intended_system_mode":            stringValue(entry.AdditionalData.IntendedSystemMode),
			"system_mode":                     stringValue(entry.AdditionalData.SystemMode),
			"domain_name":                     stringValue(entry.AdditionalData.DomainName),
			"hypershield_connectivity_status": stringValue(entry.AdditionalData.HypershieldConnectivityStatus),
			"hypershield_tenant":              stringValue(entry.AdditionalData.HypershieldTenant),
			"remote_credential_store":         stringValue(entry.AdditionalData.RemoteCredentialStore),
			"scalable_unit":                   stringValue(entry.AdditionalData.ScalableUnit),
			"smart_switch":                    boolValue(entry.AdditionalData.SmartSwitch),
			"usage":                           stringValue(entry.AdditionalData.Usage),
			"vendor":                          stringValue(entry.AdditionalData.Vendor),
		}

		value, objectDiags := types.ObjectValue(attributeTypes, attributes)
		if objectDiags.HasError() {
			return types.ListNull(objectType), objectDiags
		}
		values = append(values, value)
	}

	list, listDiags := types.ListValue(objectType, values)
	return list, listDiags
}

func inventorySwitchAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fabric_name":                     types.StringType,
		"serial_number":                   types.StringType,
		"ip_address":                      types.StringType,
		"hostname":                        types.StringType,
		"model":                           types.StringType,
		"software_version":                types.StringType,
		"software_image":                  types.StringType,
		"switch_role":                     types.StringType,
		"status":                          types.StringType,
		"platform_type":                   types.StringType,
		"preserve_config":                 types.BoolType,
		"vdc_id":                          types.Int64Type,
		"vdc_mac":                         types.StringType,
		"source_interface_name":           types.StringType,
		"source_vrf_name":                 types.StringType,
		"in_inventory":                    types.BoolType,
		"switch_id":                       types.StringType,
		"fabric_type":                     types.StringType,
		"system_uptime":                   types.StringType,
		"advisory_level":                  types.StringType,
		"alert_suspend":                   types.StringType,
		"anomaly_level":                   types.StringType,
		"vpc_configured":                  types.BoolType,
		"config_sync_status":              types.StringType,
		"discovered_system_mode":          types.StringType,
		"intended_system_mode":            types.StringType,
		"system_mode":                     types.StringType,
		"domain_name":                     types.StringType,
		"hypershield_connectivity_status": types.StringType,
		"hypershield_tenant":              types.StringType,
		"remote_credential_store":         types.StringType,
		"scalable_unit":                   types.StringType,
		"smart_switch":                    types.BoolType,
		"usage":                           types.StringType,
		"vendor":                          types.StringType,
	}
}

func stringValue(value string) attr.Value {
	if value == "" {
		return types.StringNull()
	}
	return types.StringValue(value)
}

func boolValue(value *bool) attr.Value {
	if value == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*value)
}

func int64Value(value *int64) attr.Value {
	if value == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*value)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
