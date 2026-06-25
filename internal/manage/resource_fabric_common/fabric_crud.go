// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// FabricModel is the interface that all fabric type models must implement.
// It bridges the per-type Terraform models and the shared CRUD backend.
type FabricModel interface {
	GetModelData() *NDFCFabricCommonModel
	SetModelData(*NDFCFabricCommonModel) diag.Diagnostics
	GetFabricName() string
	GetFabricType() string
}

// FabricPreMarshal is an optional interface. Implement it on a fabric model
// to adjust the common model data before it is marshalled for POST/PUT.
// Use cases: injecting fabric type, removing irrelevant fields, setting defaults.
type FabricPreMarshal interface {
	PreMarshal(ctx context.Context, data *NDFCFabricCommonModel)
}

// FabricPostUnmarshal is an optional interface. Implement it on a fabric model
// to adjust the common model data after GET unmarshal, before SetModelData.
// Use cases: normalizing API response quirks, filtering read-only noise.
type FabricPostUnmarshal interface {
	PostUnmarshal(ctx context.Context, data *NDFCFabricCommonModel)
}

// RscCreateFabric creates a fabric resource via the NDFC API.
func RscCreateFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Create(ctx, client, dg, model)
}

// createDefault is the built-in create logic.
func createDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	inData := model.GetModelData()
	if pm, ok := model.(FabricPreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}

	log.Printf("Creating fabric %s with category %s", inData.FabricName, inData.Category)

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)

	fabricPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, Data Marshall error: %v", err),
		)
		return
	}

	res, err := fabricAPI.Post(fabricPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, unexpected error: %v: %s", err, res.String()),
		)
		return
	}
	// ND BUG - license_tier is not set in the response. Try delaying the read
	time.Sleep(2 * time.Second)
	if !RscGetFabric(ctx, client, dg, model) {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Fabric %q was not found after creation", model.GetFabricName()),
		)
	}
}

// RscGetFabric retrieves fabric information by name and populates the model.
// Returns true if the resource was found, false if it no longer exists (404).
func RscGetFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	return fabricHandler(model.GetFabricType()).Read(ctx, client, dg, model)
}

// readDefault is the built-in read logic.
// Returns true if the resource was found and read successfully, false if the
// resource no longer exists on the remote infrastructure (HTTP 404). When false
// is returned, no error diagnostic is added — the caller should remove the
// resource from state so Terraform can plan a re-create.
func readDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	fabricName := model.GetFabricName()
	tflog.Debug(ctx, "Read fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	respData, err := fabricAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			tflog.Warn(ctx, "Fabric not found on remote, marking for re-creation", map[string]interface{}{
				"fabric_name": fabricName,
			})
			return false
		}
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v: %s", err, string(respData)),
		)
		return false
	}
	var outData NDFCFabricCommonModel

	err = json.Unmarshal(respData, &outData)
	if err != nil {
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v %v", err, respData),
		)
		return false
	}
	if pu, ok := model.(FabricPostUnmarshal); ok {
		pu.PostUnmarshal(ctx, &outData)
	}
	model.SetModelData(&outData)
	return true
}

// RscUpdateFabric updates a fabric resource via the NDFC API.
func RscUpdateFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Update(ctx, client, dg, model)
}

// updateDefault is the built-in update logic.
func updateDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	inData := model.GetModelData()
	if pm, ok := model.(FabricPreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	tflog.Debug(ctx, "Update fabric", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"category":    inData.Category,
	})

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = inData.FabricName

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	res, err := fabricAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	// Read the updated fabric
	if !RscGetFabric(ctx, client, dg, model) {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Fabric %q was not found after update", inData.FabricName),
		)
		return
	}
	log.Printf("Updated fabric %s with category %s", inData.FabricName, inData.Category)
}

// RscDeleteFabric deletes a fabric resource via the NDFC API.
func RscDeleteFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Delete(ctx, client, dg, model)
}

// deleteDefault is the built-in delete logic.
func deleteDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricName := model.GetFabricName()
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	tflog.Debug(ctx, "Delete fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	res, err := fabricAPI.Delete()
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric",
			fmt.Sprintf("Could not delete fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Deleting Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
}
