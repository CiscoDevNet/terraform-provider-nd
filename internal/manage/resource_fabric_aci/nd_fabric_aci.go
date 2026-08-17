// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_aci

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aciClusterResponse struct {
	Spec   aciClusterResponseSpec   `json:"spec,omitempty"`
	Status aciClusterResponseStatus `json:"status,omitempty"`
}

type aciClusterResponseSpec struct {
	NDFCSpecValue
	FabricName string `json:"name,omitempty"`
}

type aciClusterResponseStatus struct {
	State      string                       `json:"state,omitempty"`
	LastUpdate aciClusterResponseLastUpdate `json:"lastUpdate,omitempty"`
}

type aciClusterResponseLastUpdate struct {
	Message string `json:"message,omitempty"`
}

type aciClusterCreatePayload struct {
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type aciClusterRemovePayload struct {
	Credentials NDFCSpecCredentialsValue `json:"credentials"`
	Force       bool                     `json:"force"`
}

func (r aciClusterResponse) modelData() NDFCFabricAciModel {
	data := NDFCFabricAciModel{
		Spec: r.Spec.NDFCSpecValue,
		Status: NDFCStatusValue{
			State: r.Status.State,
			LastUpdate: NDFCStatusLastUpdateValue{
				LastUpdateMessage: r.Status.LastUpdate.Message,
			},
		},
	}
	if data.Spec.Aci.FabricName == "" {
		data.Spec.Aci.FabricName = r.Spec.FabricName
	}
	return data
}

// rscCreateFabricAci creates a fabric ACI resource.
func (r *fabricAciResource) rscCreateFabricAci(ctx context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Create nd_fabric_aci id=%s", id)

	inData := fabricAciModel.GetModelData()
	inData.Spec.ClusterType = "APIC"
	inData.Spec.Aci.Telemetry.TelemetryNetwork = telemetryNetworkToCreatePayload(inData.Spec.Aci.Telemetry.TelemetryNetwork)
	createPayload := aciClusterCreatePayload{Spec: inData.Spec}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)

	clusterPayload, err := json.Marshal(createPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(clusterPayload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, unexpected error: %v %v", err, res),
		)
		return
	}

	if r.rscGetFabricAci(ctx, dg, fabricAciModel) && !dg.HasError() {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci %q after create: resource not found", id),
		)
	}
}

// rscGetFabricAci retrieves fabric ACI information by id.
// It returns true when the remote object was not found.
func (r *fabricAciResource) rscGetFabricAci(ctx context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel) bool {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Read nd_fabric_aci id=%s", id)

	preservedUsername := fabricAciModel.Username
	preservedPassword := fabricAciModel.Password
	preservedLoginDomain := fabricAciModel.LoginDomain
	if preservedLoginDomain.IsUnknown() {
		preservedLoginDomain = types.StringNull()
	}

	preservedVerifyCa := fabricAciModel.VerifyCa
	if preservedVerifyCa.IsUnknown() {
		preservedVerifyCa = types.BoolNull()
	}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	respData, err := clusterAPI.Get()

	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read fabric ACI, unexpected error: %v %v", err, respData),
		)
		return false
	}
	if respData == nil {
		log.Printf("[WARN] nd_fabric_aci id=%s not found: empty response", id)
		return true
	}

	var clusterResp aciClusterResponse
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not unmarshal fabric ACI response, unexpected error: %v", err),
		)
		return false
	}

	modelData := clusterResp.modelData()
	dg.Append(fabricAciModel.SetModelData(&modelData)...)
	if dg.HasError() {
		return false
	}
	fabricAciModel.NormalizeTelemetryNetworkState()

	fabricAciModel.Username = preservedUsername
	fabricAciModel.Password = preservedPassword
	fabricAciModel.LoginDomain = preservedLoginDomain
	fabricAciModel.VerifyCa = preservedVerifyCa
	fabricAciModel.Id = types.StringValue(id)
	return false
}

// rscUpdateFabricAci updates a fabric ACI resource.
func (r *fabricAciResource) rscUpdateFabricAci(ctx context.Context, dg *diag.Diagnostics, plan *FabricAciModel) {
	id := plan.Id.ValueString()
	log.Printf("[INFO] Update nd_fabric_aci id=%s", id)

	inData := plan.manageUpdatePayload()

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric ACI",
			fmt.Sprintf("Could not update fabric ACI, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}

	res, err := clusterAPI.Put(inDataBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric ACI",
			fmt.Sprintf("Could not update fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}

	if r.rscGetFabricAci(ctx, dg, plan) && !dg.HasError() {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci %q after update: resource not found", id),
		)
	}
}

// rscDeleteFabricAci deletes a fabric ACI resource by id.
func (r *fabricAciResource) rscDeleteFabricAci(_ context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel, force bool) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Delete nd_fabric_aci id=%s", id)

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	clusterAPI.Delete = true

	modelData := fabricAciModel.GetModelData()
	removePayload := aciClusterRemovePayload{
		Credentials: modelData.Spec.Credentials,
		Force:       force,
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}
}
