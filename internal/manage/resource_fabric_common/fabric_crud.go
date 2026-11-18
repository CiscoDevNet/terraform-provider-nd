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
	"errors"
	"fmt"
	"log"
	"net/http"
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

// FabricCRUDAPI is the request surface used by the shared fabric handler.
type FabricCRUDAPI interface {
	Get() ([]byte, error)
	Post([]byte, *ndapi.APIOptions) (nd.Res, error)
	Put([]byte, *ndapi.APIOptions) (nd.Res, error)
	Delete(...[]byte) (nd.Res, error)
}

// FabricAPIOverride supplies resource-specific request handling for shared CRUD.
// Implementations must omit sensitive response bodies from request errors.
type FabricAPIOverride interface {
	WrapFabricAPI(context.Context, *api.FabricAPI) FabricCRUDAPI
}

// FabricWriteReadback handles readback after a successful POST or PUT.
type FabricWriteReadback interface {
	ReadAfterFabricWrite(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, method string)
}

func fabricAPIForModel(ctx context.Context, model FabricModel, fabricAPI *api.FabricAPI) FabricCRUDAPI {
	if override, ok := model.(FabricAPIOverride); ok {
		return override.WrapFabricAPI(ctx, fabricAPI)
	}
	return fabricAPI
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

	res, err := fabricAPIForModel(ctx, model, fabricAPI).Post(fabricPayload, nil)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, unexpected error: %v: %s", err, res.String()),
		)
		return
	}
	if readback, ok := model.(FabricWriteReadback); ok {
		readback.ReadAfterFabricWrite(ctx, client, dg, http.MethodPost)
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

// RscValidateFabricType checks that a retrieved fabric belongs to one of the
// resource types supported by an import operation.
func RscValidateFabricType(dg *diag.Diagnostics, model FabricModel, allowedTypes ...string) bool {
	actualType := model.GetModelData().Management.FabricType
	for _, allowedType := range allowedTypes {
		if actualType == allowedType {
			return true
		}
	}

	dg.AddError(
		"Error Importing Fabric",
		fmt.Sprintf("Fabric %q has type %q; expected one of %v", model.GetFabricName(), actualType, allowedTypes),
	)
	return false
}

// readDefault is the built-in read logic.
// Returns true if the resource was found and read successfully. Callers must
// check error diagnostics before treating false as a missing resource (HTTP 404)
// and removing it from state.
func readDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	fabricName := model.GetFabricName()
	tflog.Debug(ctx, "Read fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	_, overridden := model.(FabricAPIOverride)
	respData, err := fabricAPIForModel(ctx, model, fabricAPI).Get()
	if err != nil {
		if overridden && errors.Is(err, ndapi.ErrNotFound) {
			return false
		}
		if !overridden && strings.Contains(err.Error(), "StatusCode 404") {
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
		if overridden {
			dg.AddError("Error reading Fabric", fmt.Sprintf("Could not decode the API response for fabric %q.", fabricName))
			return false
		}
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v %v", err, respData),
		)
		return false
	}
	outData.Id = fabricName
	if pu, ok := model.(FabricPostUnmarshal); ok {
		pu.PostUnmarshal(ctx, &outData)
	}
	dg.Append(model.SetModelData(&outData)...)
	return !dg.HasError()
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
	res, err := fabricAPIForModel(ctx, model, fabricAPI).Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	if readback, ok := model.(FabricWriteReadback); ok {
		readback.ReadAfterFabricWrite(ctx, client, dg, http.MethodPut)
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
	res, err := fabricAPIForModel(ctx, model, fabricAPI).Delete()
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric",
			fmt.Sprintf("Could not delete fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Deleting Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
}
