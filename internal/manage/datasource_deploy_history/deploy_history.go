// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_deploy_history

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

// These assignments make Go check the required method signatures at compile time.
// DataSource requires Metadata, Schema, and Read; DataSourceWithConfigure adds
// Configure. The blank identifier (_) discards each value; no methods are called.
var (
	_ datasource.DataSource              = &deployHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &deployHistoryDataSource{}
)

// NewDeployHistoryDataSource is the factory registered by this package's init().
// The framework calls it to create an instance, then calls the instance's methods.
// &deployHistoryDataSource{} creates a non-nil pointer to a new struct whose
// manageClient field starts as nil. Returning an interface preserves that pointer.
func NewDeployHistoryDataSource() datasource.DataSource {
	return &deployHistoryDataSource{}
}

// deployHistoryDataSource holds the client between Configure and Read calls.
// A receiver such as (source *deployHistoryDataSource) names the instance pointer
// "source" inside a method. Updating its fields changes that instance, not a copy.
type deployHistoryDataSource struct {
	manageClient *manage.NexusDashboardManage
}

// NDDeployHistoryModel is an alias (=) for the generated API-facing Go model.
// Both names refer to the same type; the alias neither converts data nor changes
// the endpoint. Its ordinary Go fields differ from the Terraform model's types.
type NDDeployHistoryModel = NDFCDeployHistoryModel

// Metadata gives Terraform the full data-source name, such as nd_deploy_history.
func (source *deployHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deploy_history"
}

// Schema supplies the generated input/output definitions, including nested types.
// Change generator/defs/deploy_history.yaml and regenerate to update this schema.
func (source *deployHistoryDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DeployHistoryDataSourceSchema(ctx)
}

// Configure receives the provider's clients from the framework; it does not fetch
// history but it saves the client and module pointer for later use. The context parameter is named _ because it is unused. req is passed
// by value, while resp is a pointer so diagnostics remain visible to the caller.
func (source *deployHistoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

	// The framework may call this before provider configuration is available.
	if req.ProviderData == nil {
		return
	}

	// This is a type assertion, not a data conversion. The comma-ok form returns
	// false instead of panicking if the supplied value does not implement GetModule.
	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	// Find the manage module in the shared provider's module registry.
	manageModule := client.GetModule("manage")
	if manageModule == nil {
		resp.Diagnostics.AddError(
			"Manage Module Not Found",
			"The manage module was not registered with the provider.",
		)
		return
	}

	// Check the module's concrete pointer type before accessing its API client.
	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule),
		)
		return
	}
	// Retain the shared module pointer for Read; this does not create another client.
	source.manageClient = manageClient
}

// Read retrieves history without changing ND configuration. Its data flow is:
// Terraform config -> Terraform model -> GET -> JSON -> API model -> Terraform model -> state response.
func (source *deployHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: deploy_history")

	// This model uses framework values (types.String, types.Int64, types.List),
	// which represent known, null, and unknown Terraform values.
	var data DeployHistoryModel

	// Config.Get fills data through its address (&data), matching tfsdk field tags.
	// It returns a diagnostics slice; ... expands it into arguments for Append.
	// HasError stops processing on errors, but allows warnings to remain in the response.
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Null means absent; unknown means not yet resolved by Terraform. Reject both
	// and blank names. TrimSpace checks validity without modifying the supplied name.
	if data.FabricName.IsNull() || data.FabricName.IsUnknown() || strings.TrimSpace(data.FabricName.ValueString()) == "" {
		resp.Diagnostics.AddError(
			"Fabric Name Required",
			"The fabric_name attribute must contain a known, non-empty value to read deployment history.",
		)
		return
	}

	// An omitted count (null) can use the default, but an unresolved count cannot.
	if data.NoOfRecords.IsUnknown() {
		resp.Diagnostics.AddError(
			"Record Count Unknown",
			"The no_of_records attribute must contain a known value when reading deployment history.",
		)
		return
	}

	// Use 10 for the API limit when omitted. ValueInt64 extracts the ordinary Go
	// value from a known Terraform number; reject nonpositive limits before slicing.
	recordCount := int64(10)
	if !data.NoOfRecords.IsNull() {
		recordCount = data.NoOfRecords.ValueInt64()
	}
	if recordCount < 1 {
		resp.Diagnostics.AddError("Invalid Record Count", "The no_of_records attribute must be at least 1.")
		return
	}

	// Configure must have supplied a usable client. || short-circuits, so the
	// ApiClient field is accessed only when manageClient itself is non-nil.
	if source.manageClient == nil || source.manageClient.ApiClient == nil {
		resp.Diagnostics.AddError("Manage Client Not Configured", "The manage client must be configured before reading deployment history.")
		return
	}

	// Extract the Go string and configure the existing API wrapper. No request is
	// sent yet. DefaultFabric sets lock scope; FabricName selects the fabric in the URL.
	fabricName := data.FabricName.ValueString()
	historyAPI := api.NewConfigAPI(source.manageClient.ApiClient, ndapi.DefaultFabric)
	historyAPI.FabricName = fabricName
	historyAPI.SetOperation(api.OpDeployHistory)

	// Ask ND for newest completion times first, at most recordCount entries, and
	// the first page. These are server query parameters, not local sorting logic;
	// their support depends on ND. fmt.Sprintf formats the integer as max=<count>.
	historyAPI.SetQueryParams("sort=completeTimestamp:desc", fmt.Sprintf("max=%d", recordCount), "offset=0")

	// Get builds the URL and sends an authenticated GET through the shared client.
	// The response is JSON bytes, not yet a populated model. Report request errors.
	responseData, err := historyAPI.Get()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Deployment History",
			fmt.Sprintf("Could not read deployment history for fabric %q: %v", fabricName, err),
		)
		return
	}

	// Decode JSON into the API model using json field tags. &history lets Unmarshal
	// populate its fields. Undeclared response fields are ignored. err is scoped to
	// this if statement; malformed JSON or incompatible field types stop the read.
	var history NDDeployHistoryModel
	if err := json.Unmarshal(responseData, &history); err != nil {
		resp.Diagnostics.AddError(
			"Error Decoding Deployment History",
			fmt.Sprintf("Could not decode deployment history for fabric %q: %v", fabricName, err),
		)
		return
	}

	// Local sorting is disabled below; the active reader preserves ND response order.
	// If enabled, this block compares parsed times and retains response order for ties.
	// sort.SliceStable(history.DeploymentRecords, func(firstIndex, secondIndex int) bool {
	// 	firstTime, firstErr := time.Parse(time.RFC3339Nano, history.DeploymentRecords[firstIndex].CompleteTimestamp)
	// 	if firstErr != nil {
	// 		return false
	// 	}
	// 	secondTime, secondErr := time.Parse(time.RFC3339Nano, history.DeploymentRecords[secondIndex].CompleteTimestamp)
	// 	if secondErr != nil {
	// 		return true
	// 	}
	// 	return firstTime.After(secondTime)
	// })

	// Enforce the cap even if ND ignores max. len returns int, so convert it to
	// int64 for comparison. [:recordCount] keeps the first N records, excluding
	// index N; it does not trim nested command responses or fetch additional pages.
	if int64(len(history.DeploymentRecords)) > recordCount {
		history.DeploymentRecords = history.DeploymentRecords[:recordCount]
	}

	// History responses may omit the input fabric name. Preserve the requested name
	// and keep an omitted no_of_records null in state, despite using 10 in the GET.
	history.FabricName = fabricName
	history.NoOfRecords = nil
	if !data.NoOfRecords.IsNull() {

		// &recordCount supplies the *int64 expected by the API model for a known count.
		history.NoOfRecords = &recordCount
	}

	// Convert API Go fields and nested slices into Terraform values. This generated
	// method updates data and returns diagnostics; it does not write state or call ND.
	resp.Diagnostics.Append(data.SetModelData(&history)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate the framework's state response from data, including computed records.
	// State.Set is not json.Marshal and does not directly write the state file.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource deploy_history with fabric_name=%s", fabricName)
}
