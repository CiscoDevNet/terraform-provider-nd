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
	tflog.Debug(ctx, "Preparing vPC pair create request", map[string]interface{}{
		"fabric_name":          inData.FabricName,
		"switch_id_1":          inData.SwitchId1,
		"switch_id_2":          inData.SwitchId2,
		"use_virtual_peerlink": inData.UseVirtualPeerlink,
		"deploy":               inData.Deploy,
	})

	vpcPairAPI := api.NewVpcPairAPI(r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2
	// set action to pair for create operation and use PUT method .
	inData.VpcAction = "pair"

	payload, err := json.Marshal(inData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal vPC pair create payload", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"error":       err.Error(),
		})
		dg.AddError(
			"Error Creating vPC Pair",
			fmt.Sprintf("Could not create vPC pair, data marshal error: %v", err),
		)
		return
	}

	res, err := vpcPairAPI.Put(payload, nil)
	if err != nil {
		tflog.Error(ctx, "vPC pair create API call failed", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"url":         vpcPairAPI.PutUrl(),
			"error":       err.Error(),
			"response":    res.Raw,
		})
		dg.AddError(
			"Error Creating vPC Pair",
			fmt.Sprintf("Could not create vPC pair, unexpected error: %v %v", err, res),
		)
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Config save and deploy failed during vPC pair create", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"deploy":      inData.Deploy,
		})
		return
	}

	tflog.Debug(ctx, "vPC pair create request completed, refreshing state", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"switch_id_2": inData.SwitchId2,
	})

	time.Sleep(2 * time.Second)
	r.rscGetVpcPair(ctx, dg, input)
}

// GetVpcPair retrieves vPC pair information by fabric and switch identifier.
func (r *vpcPairResource) rscGetVpcPair(ctx context.Context, dg *diag.Diagnostics, in *VpcPairModel) {
	outData, err := r.readVpcPairState(ctx, in.GetModelData())
	if err != nil {
		tflog.Error(ctx, "Failed to read vPC pair state", map[string]interface{}{
			"fabric_name": in.GetModelData().FabricName,
			"switch_id_2": in.GetModelData().SwitchId2,
			"error":       err.Error(),
		})
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
	tflog.Debug(ctx, "Preparing vPC pair update request", map[string]interface{}{
		"fabric_name":          inData.FabricName,
		"switch_id_1":          inData.SwitchId1,
		"switch_id_2":          inData.SwitchId2,
		"use_virtual_peerlink": inData.UseVirtualPeerlink,
		"deploy":               inData.Deploy,
	})

	vpcPairAPI := api.NewVpcPairAPI(r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2
	inData.VpcAction = "pair"

	payload, err := json.Marshal(inData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal vPC pair update payload", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"error":       err.Error(),
		})
		dg.AddError(
			"Error Updating vPC Pair",
			fmt.Sprintf("Could not update vPC pair, data marshal error: %v", err),
		)
		tflog.Error(ctx, "Error Updating vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	res, err := vpcPairAPI.Put(payload, nil)
	if err != nil {
		tflog.Error(ctx, "vPC pair update API call failed", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"url":         vpcPairAPI.PutUrl(),
			"error":       err.Error(),
			"response":    res.Raw,
		})
		dg.AddError(
			"Error Updating vPC Pair",
			fmt.Sprintf("Could not update vPC pair, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Updating vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Config save and deploy failed during vPC pair update", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"deploy":      inData.Deploy,
		})
		return
	}

	r.rscGetVpcPair(ctx, dg, vpcPairModel)
	tflog.Debug(ctx, "vPC pair update completed", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"switch_id_2": inData.SwitchId2,
	})
}

func (r *vpcPairResource) readVpcPairState(ctx context.Context, inData *NDFCVpcPairModel) (*NDFCVpcPairModel, error) {
	vpcPairAPI := api.NewVpcPairAPI(r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2

	tflog.Debug(ctx, "Fetching vPC pair state from ND", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"switch_id_2": inData.SwitchId2,
		"url":         vpcPairAPI.GetUrl(),
	})

	respData, err := vpcPairAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", vpcPairAPI.GetUrl(), err)
	}

	var outData NDFCVpcPairModel
	if err := json.Unmarshal(respData, &outData); err != nil {
		tflog.Error(ctx, "Failed to decode vPC pair state response", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"error":       err.Error(),
		})
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

	tflog.Debug(ctx, "Fetching switches inventory to derive vPC pair deploy state", map[string]interface{}{
		"fabric_name": fabricName,
		"switch_id_1": switchID1,
		"switch_id_2": switchID2,
		"url":         invAPI.GetUrl(),
	})

	respData, err := invAPI.Get()
	if err != nil {
		tflog.Error(ctx, "Failed to fetch switches inventory for vPC pair deploy state", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id_1": switchID1,
			"switch_id_2": switchID2,
			"url":         invAPI.GetUrl(),
			"error":       err.Error(),
		})
		return false, fmt.Errorf("could not read switches for fabric %s: %w", fabricName, err)
	}

	deployState, err := parseVpcPairDeployState(ctx, respData, switchID1, switchID2)
	if err != nil {
		tflog.Error(ctx, "Failed to parse switches inventory for vPC pair deploy state", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id_1": switchID1,
			"switch_id_2": switchID2,
			"error":       err.Error(),
		})
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

func parseVpcPairDeployState(ctx context.Context, respData []byte, switchID1, switchID2 string) (bool, error) {
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
		tflog.Debug(ctx, "vPC pair deploy state is false because switches are not yet marked vPC-configured", map[string]interface{}{
			"switch_id_1":         switchID1,
			"switch_id_2":         switchID2,
			"switch_1_configured": sw1.VpcConfigured,
			"switch_2_configured": sw2.VpcConfigured,
		})
		return false, nil
	}

	if !strings.EqualFold(sw1.AdditionalData.ConfigSyncStatus, "inSync") || !strings.EqualFold(sw2.AdditionalData.ConfigSyncStatus, "inSync") {
		tflog.Debug(ctx, "vPC pair deploy state is false because config sync is not inSync on both switches", map[string]interface{}{
			"switch_id_1":          switchID1,
			"switch_id_2":          switchID2,
			"switch_1_sync_status": sw1.AdditionalData.ConfigSyncStatus,
			"switch_2_sync_status": sw2.AdditionalData.ConfigSyncStatus,
		})
		return false, nil
	}

	if sw1.VpcData.PeerSwitchID != switchID2 || sw2.VpcData.PeerSwitchID != switchID1 {
		tflog.Debug(ctx, "vPC pair deploy state is false because peer switch IDs do not match the requested pair", map[string]interface{}{
			"switch_id_1":             switchID1,
			"switch_id_2":             switchID2,
			"switch_1_peer_switch_id": sw1.VpcData.PeerSwitchID,
			"switch_2_peer_switch_id": sw2.VpcData.PeerSwitchID,
		})
		return false, nil
	}

	tflog.Debug(ctx, "vPC pair deploy state matched requested switches", map[string]interface{}{
		"switch_id_1": switchID1,
		"switch_id_2": switchID2,
	})
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
	tflog.Debug(ctx, "Preparing vPC pair delete request", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"switch_id_1": inData.SwitchId1,
		"switch_id_2": inData.SwitchId2,
		"deploy":      inData.Deploy,
	})

	vpcPairAPI := api.NewVpcPairAPI(r.manageClient.ApiClient)
	vpcPairAPI.FabricName = inData.FabricName
	vpcPairAPI.SwitchID = inData.SwitchId2

	// The ND API removes a pair through the same PUT endpoint with the unPair action.
	inData.VpcAction = "unPair"

	payload, err := json.Marshal(inData)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal vPC pair delete payload", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"error":       err.Error(),
		})
		dg.AddError(
			"Error Deleting vPC Pair",
			fmt.Sprintf("Could not delete vPC pair, data marshal error: %v", err),
		)
		tflog.Error(ctx, "Error Deleting vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	res, err := vpcPairAPI.Put(payload, nil)
	if err != nil {
		tflog.Error(ctx, "vPC pair delete API call failed", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"url":         vpcPairAPI.PutUrl(),
			"error":       err.Error(),
			"response":    res.Raw,
		})
		dg.AddError(
			"Error Deleting vPC Pair",
			fmt.Sprintf("Could not delete vPC pair, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Deleting vPC Pair", map[string]interface{}{"error": err.Error()})
		return
	}

	deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, inData.FabricName, true, inData.Deploy, dg)

	if dg.HasError() {
		tflog.Error(ctx, "Config save and deploy failed during vPC pair delete", map[string]interface{}{
			"fabric_name": inData.FabricName,
			"switch_id_2": inData.SwitchId2,
			"deploy":      inData.Deploy,
		})
		return
	}

	tflog.Debug(ctx, "vPC pair delete completed", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"switch_id_2": inData.SwitchId2,
	})
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
