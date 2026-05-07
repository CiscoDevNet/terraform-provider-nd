// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_vpc_pair

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RscCreateVpcPair creates a vPC pair resource using the vPC pair model.
func (r *vpcPairResource) rscCreateVpcPair(ctx context.Context, dg *diag.Diagnostics, input *VpcPairModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()
	log.Printf("Creating vPC pair for fabric %s switch %s", inData.FabricName, inData.SwitchId2)

	vpcPairAPI := api.NewVpcPairAPI(nil, r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2
	// set action to pair for create operation and use PUT method .
	inData.VpcAction = "pair"

	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating vPC Pair",
			fmt.Sprintf("Could not create vPC pair, data marshal error: %v", err),
		)
		return
	}

	res, err := vpcPairAPI.Put(payload)
	if err != nil {
		dg.AddError(
			"Error Creating vPC Pair",
			fmt.Sprintf("Could not create vPC pair, unexpected error: %v %v", err, res),
		)
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Error Save and deploy while Creating vPC Pair")
		return
	}

	time.Sleep(2 * time.Second)
	r.rscGetVpcPair(ctx, dg, input)
}

// GetVpcPair retrieves vPC pair information by fabric and switch identifier.
func (r *vpcPairResource) rscGetVpcPair(ctx context.Context, dg *diag.Diagnostics, in *VpcPairModel) {
	vpcPairAPI := api.NewVpcPairAPI(nil, r.manageClient.ApiClient)
	vpcPairAPI.FabricName = in.FabricName.ValueString()
	vpcPairAPI.SwitchID = in.SwitchId2.ValueString()

	respData, err := vpcPairAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("Could not read vPC pair, unexpected error: %v %v", err, respData),
		)
		return
	}

	var outData NDFCVpcPairModel
	if err := json.Unmarshal(respData, &outData); err != nil {
		dg.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("Could not decode vPC pair response: %v %s", err, string(respData)),
		)
		return
	}

	if outData.FabricName == "" {
		outData.FabricName = in.GetModelData().FabricName
	}
	if outData.SwitchId1 == "" {
		outData.SwitchId1 = in.GetModelData().SwitchId1
	}
	if outData.SwitchId2 == "" {
		outData.SwitchId2 = in.GetModelData().SwitchId2
	}
	outData.Deploy = in.GetModelData().Deploy
	if outData.UseVirtualPeerlink == nil {
		outData.UseVirtualPeerlink = in.GetModelData().UseVirtualPeerlink
	}
	in.SetModelData(&outData)
	setVpcPairID(in)

	tflog.Debug(ctx, "Read vPC pair", map[string]interface{}{
		"fabric_name": outData.FabricName,
		"switch_id":   outData.SwitchId2,
	})
}

// UpdateVpcPair updates a vPC pair with the provided payload.
func (r *vpcPairResource) rscUpdateVpcPair(ctx context.Context, dg *diag.Diagnostics, vpcPairModel *VpcPairModel) {
	inData := vpcPairModel.GetModelData()
	log.Printf("Updating vPC pair for fabric %s switch %s", inData.FabricName, inData.SwitchId2)

	vpcPairAPI := api.NewVpcPairAPI(nil, r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2

	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating vPC Pair",
			fmt.Sprintf("Could not update vPC pair, data marshal error: %v", err),
		)
		tflog.Error(ctx, "Error Updating vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	res, err := vpcPairAPI.Put(payload)
	if err != nil {
		dg.AddError(
			"Error Updating vPC Pair",
			fmt.Sprintf("Could not update vPC pair, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Updating vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Error Save and deploy while Updating vPC Pair")
		return
	}

	r.rscGetVpcPair(ctx, dg, vpcPairModel)
	log.Printf("Updated vPC pair for fabric %s switch %s", inData.FabricName, inData.SwitchId2)
}

// DeleteVpcPair deletes a vPC pair by fabric and switch identifier.
func (r *vpcPairResource) rscDeleteVpcPair(ctx context.Context, dg *diag.Diagnostics, state *VpcPairModel) {
	if state == nil {
		dg.AddError(
			"Error Deleting vPC Pair",
			"The current vPC pair state is missing.",
		)
		return
	}

	inData := state.GetModelData()
	log.Printf("Deleting vPC pair for fabric %s switch %s", inData.FabricName, inData.SwitchId2)

	vpcPairAPI := api.NewVpcPairAPI(nil, r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2

	// The ND API removes a pair through the same PUT endpoint with the unPair action.
	inData.VpcAction = "unPair"

	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Deleting vPC Pair",
			fmt.Sprintf("Could not delete vPC pair, data marshal error: %v", err),
		)
		tflog.Error(ctx, "Error Deleting vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	res, err := vpcPairAPI.Put(payload)
	if err != nil {
		dg.AddError(
			"Error Deleting vPC Pair",
			fmt.Sprintf("Could not delete vPC pair, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Deleting vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Error Save and deploy while Deleting vPC Pair")
	}
}

func setVpcPairID(model *VpcPairModel) {
	if model == nil {
		return
	}

	if model.FabricName.IsNull() || model.FabricName.IsUnknown() || model.SwitchId2.IsNull() || model.SwitchId2.IsUnknown() {
		return
	}

	model.Id = types.StringValue(fmt.Sprintf("%s/%s", model.FabricName.ValueString(), model.SwitchId2.ValueString()))
}
