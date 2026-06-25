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
	"strings"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type fabricSwitchesResponse struct {
	Switches []fabricSwitchEntry `json:"switches,omitempty"`
}

type fabricSwitchEntry struct {
	SerialNumber   string                     `json:"serialNumber,omitempty"`
	VpcConfigured  bool                       `json:"vpcConfigured,omitempty"`
	AdditionalData fabricSwitchAdditionalData `json:"additionalData,omitempty"`
	VpcData        fabricSwitchVpcData        `json:"vpcData,omitempty"`
}

type fabricSwitchAdditionalData struct {
	ConfigSyncStatus string `json:"configSyncStatus,omitempty"`
}

type fabricSwitchVpcData struct {
	PeerSwitchID string `json:"peerSwitchId,omitempty"`
}

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

	res, err := vpcPairAPI.Put(payload, nil)
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
	outData, err := r.readVpcPairState(ctx, in.GetModelData())
	if err != nil {
		dg.AddError(
			"Error Reading vPC Pair",
			fmt.Sprintf("Could not read vPC pair, unexpected error: %v", err),
		)
		return
	}
	if outData.UseVirtualPeerlink == nil {
		outData.UseVirtualPeerlink = in.GetModelData().UseVirtualPeerlink
	}
	in.SetModelData(outData)
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
	inData.VpcAction = "pair"

	payload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating vPC Pair",
			fmt.Sprintf("Could not update vPC pair, data marshal error: %v", err),
		)
		tflog.Error(ctx, "Error Updating vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	res, err := vpcPairAPI.Put(payload, nil)
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

func (r *vpcPairResource) readVpcPairState(ctx context.Context, inData *NDFCVpcPairModel) (*NDFCVpcPairModel, error) {
	vpcPairAPI := api.NewVpcPairAPI(nil, r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2

	respData, err := vpcPairAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", vpcPairAPI.GetUrl(), err)
	}

	var outData NDFCVpcPairModel
	if err := json.Unmarshal(respData, &outData); err != nil {
		return nil, fmt.Errorf("decode vpc pair response: %w", err)
	}

	if outData.FabricName == "" {
		outData.FabricName = inData.FabricName
	}
	if outData.SwitchId1 == "" {
		outData.SwitchId1 = inData.SwitchId1
	}
	if outData.SwitchId2 == "" {
		outData.SwitchId2 = inData.SwitchId2
	}

	// Deploy is an operation flag in this resource, not durable remote state.
	// Preserve the caller's value to avoid apply/read drift after update.
	outData.Deploy = inData.Deploy

	return &outData, nil
}

func (r *vpcPairResource) getVpcPairDeployState(ctx context.Context, fabricName, switchID1, switchID2 string) (bool, error) {
	invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, fabricName)
	invAPI.FabricName = fabricName
	invAPI.SetOperation(api.OpGetAllSwitches)

	respData, err := invAPI.Get()
	if err != nil {
		return false, fmt.Errorf("could not read switches for fabric %s: %w", fabricName, err)
	}

	deployState, err := parseVpcPairDeployState(respData, switchID1, switchID2)
	if err != nil {
		return false, err
	}

	tflog.Debug(ctx, "Derived vPC pair deploy state from switches API", map[string]interface{}{
		"fabric_name": fabricName,
		"switch_id_1": switchID1,
		"switch_id_2": switchID2,
		"deploy":      deployState,
	})

	return deployState, nil
}

func parseVpcPairDeployState(respData []byte, switchID1, switchID2 string) (bool, error) {
	var resp fabricSwitchesResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return false, fmt.Errorf("could not decode switches response: %w", err)
	}

	switchesBySerial := make(map[string]fabricSwitchEntry, len(resp.Switches))
	for _, sw := range resp.Switches {
		switchesBySerial[sw.SerialNumber] = sw
	}

	sw1, ok := switchesBySerial[switchID1]
	if !ok {
		return false, fmt.Errorf("switch %s not found in switches response", switchID1)
	}

	sw2, ok := switchesBySerial[switchID2]
	if !ok {
		return false, fmt.Errorf("switch %s not found in switches response", switchID2)
	}

	if !sw1.VpcConfigured || !sw2.VpcConfigured {
		return false, nil
	}

	if !strings.EqualFold(sw1.AdditionalData.ConfigSyncStatus, "inSync") || !strings.EqualFold(sw2.AdditionalData.ConfigSyncStatus, "inSync") {
		return false, nil
	}

	if sw1.VpcData.PeerSwitchID != switchID2 || sw2.VpcData.PeerSwitchID != switchID1 {
		return false, nil
	}

	return true, nil
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

	res, err := vpcPairAPI.Put(payload, nil)
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

	if model.FabricName.IsNull() || model.FabricName.IsUnknown() ||
		model.SwitchId1.IsNull() || model.SwitchId1.IsUnknown() ||
		model.SwitchId2.IsNull() || model.SwitchId2.IsUnknown() {
		return
	}

	model.Id = types.StringValue(fmt.Sprintf(
		"%s/%s:%s",
		model.FabricName.ValueString(),
		model.SwitchId1.ValueString(),
		model.SwitchId2.ValueString(),
	))
}
