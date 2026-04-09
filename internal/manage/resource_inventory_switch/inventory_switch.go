// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_inventory_switch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	ModeDiscovery = "discovery"
	ModeBootstrap = "bootstrap"

	StatusManageable     = "Manageable"
	StatusUnreachable    = "Unreachable"
	StatusPending        = "Pending"
	StatusFailed         = "Failed"
	StatusMigrationMode  = "Migration"
	StatusAlreadyManaged = "alreadyManaged"

	MaxDiscoveryWaitTime = 5 * time.Minute
	PollInterval         = 10 * time.Second
)

// rscCreateInventory creates inventory switches
func (r *inventorySwitchResource) rscCreateInventory(ctx context.Context, dg *diag.Diagnostics, input *InventorySwitchModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}
	// Get switch data from model
	switchesData := input.GetModelData()

	invAPI := api.NewInventoryAPI(nil, r.manageClient.ApiClient)
	invAPI.FabricName = switchesData.FabricName

	log.Printf("Creating inventory for fabric %s with mode %s", switchesData.FabricName, switchesData.Mode)

	if switchesData.Mode == ModeBootstrap {
		r.createBootstrapSwitches(ctx, dg, invAPI, switchesData)
		if dg.HasError() {
			return
		}
		// Bootstrap uses the old linear flow for config save/deploy
		deployment.ConfigSaveAndDeploy(ctx, r.manageClient.ApiClient, switchesData.FabricName, switchesData.Recalculate, switchesData.Deploy, dg)
		if dg.HasError() {
			return
		}
	} else {
		// Discovery mode uses the FSM which handles the full lifecycle
		// including retryable config-save errors
		fsm := NewInventoryFSM(ctx, r, true, invAPI, switchesData, dg)
		fsm.Run()
		if dg.HasError() {
			return
		}
	}

	// Read back the created state
	r.rscGetInventory(ctx, dg, input)
}

// rscGetInventory reads the current state of inventory switches
func (r *inventorySwitchResource) rscGetInventory(ctx context.Context, dg *diag.Diagnostics, input *InventorySwitchModel) {
	DumpInventorySwitchModel("rscGetInventory INPUT", input)
	fabricName := input.FabricName.ValueString()

	// Get data model — input may only have IPs, not serial numbers yet
	switchesData := input.GetModelData()

	// Build lookup maps from input: serial -> map key, IP -> map key
	serialToKey := make(map[string]string)
	ipToKey := make(map[string]string)
	for key, sw := range switchesData.Switches {
		log.Printf("Processing switch %s", sw.IpAddress)
		if sw.SerialNumber != "" {
			log.Printf("Processing switch via serial key %s:%s", key, sw.SerialNumber)
			serialToKey[sw.SerialNumber] = key
		}
		if sw.IpAddress != "" {
			log.Printf("Processing switch via IP key %s:%s", key, sw.IpAddress)
			ipToKey[sw.IpAddress] = key
		}
	}

	// Fetch all switches in the fabric - no need to fill lookup maps
	resp, err := r.getAllSwitchesByFabric(ctx, fabricName, false)
	if err != nil {
		dg.AddError("Error Reading Inventory", err.Error())
		return
	}
	// Response data to populate
	outData := *switchesData
	outData.Switches = make(map[string]NDFCSwitchesValue)

	// Filter to only switches that match our input by serial or IP
	for _, sw := range resp.Switches {
		var mapKey string
		if key, ok := serialToKey[sw.SerialNumber]; ok {
			log.Printf("Found switch %s in fabric %s key %s", sw.SerialNumber, fabricName, key)
			mapKey = key
		} else if key, ok := ipToKey[sw.FabricManagementIp]; ok {
			log.Printf("Found switch %s in fabric %s key %s", sw.FabricManagementIp, fabricName, key)
			mapKey = key
		} else {
			log.Printf("Switch %s found in fabric %s - not part of resource", sw.SerialNumber, fabricName)
			continue
		}
		outData.Switches[mapKey] = sw.NDFCSwitchesValue
	}

	input.SetModelData(&outData)
	DumpInventorySwitchModel("rscGetInventory OUTPUT", input)
	log.Printf("Read inventory for fabric %s, found %d switches", fabricName, len(outData.Switches))
}

// rscUpdateInventory updates inventory switches
func (r *inventorySwitchResource) rscUpdateInventory(ctx context.Context, dg *diag.Diagnostics, plan *InventorySwitchModel, state *InventorySwitchModel) {
	DumpInventorySwitchModel("rscUpdateInventory PLAN", plan)
	DumpInventorySwitchModel("rscUpdateInventory STATE", state)
	fabricName := plan.FabricName.ValueString()
	invAPI := api.NewInventoryAPI(nil, r.manageClient.ApiClient)
	invAPI.FabricName = fabricName

	planData := plan.GetModelData()
	stateData := state.GetModelData()

	// Identify switches to add, remove, and update
	actions := r.diffSwitches(planData, stateData)

	addSwitches := actions["add"]
	removeSwitches := actions["del"]

	// Remove switches that are no longer in plan
	if removeSwitches != nil {
		removeSerials := make([]string, 0)
		for _, delSw := range removeSwitches {
			removeSerials = append(removeSerials, delSw.SerialNumber)
		}
		if len(removeSerials) > 0 {
			r.removeSwitches(ctx, dg, invAPI, removeSerials)
			if dg.HasError() {
				return
			}
		}
	}

	// Add/update switches via FSM (handles discovery, readiness, creds, roles, config save, deploy)
	if addSwitches != nil {
		addData := &NDFCInventorySwitchModel{
			FabricName:               fabricName,
			Mode:                     planData.Mode,
			Username:                 planData.Username,
			Password:                 planData.Password,
			PreserveConfig:           planData.PreserveConfig,
			PlatformType:             planData.PlatformType,
			SnmpV3AuthProtocol:       planData.SnmpV3AuthProtocol,
			RemoteCredentialStore:    planData.RemoteCredentialStore,
			RemoteCredentialStoreKey: planData.RemoteCredentialStoreKey,
			MaxHop:                   planData.MaxHop,
			Recalculate:              planData.Recalculate,
			Deploy:                   planData.Deploy,
			Switches:                 make(map[string]NDFCSwitchesValue),
		}
		for _, sw := range addSwitches {
			addData.Switches[sw.IpAddress] = sw
		}
		// add, update, and deploy
		fsm := NewInventoryFSM(ctx, r, false, invAPI, addData, dg)
		fsm.Run()
		if dg.HasError() {
			return
		}
	}
	// Read back the updated state
	r.rscGetInventory(ctx, dg, plan)
}

// rscDeleteInventory deletes inventory switches
func (r *inventorySwitchResource) rscDeleteInventory(ctx context.Context, dg *diag.Diagnostics, state *InventorySwitchModel) {
	fabricName := state.FabricName.ValueString()
	invAPI := api.NewInventoryAPI(nil, r.manageClient.ApiClient)
	invAPI.FabricName = fabricName

	stateData := state.GetModelData()

	swDelList := make([]string, 0)

	switches, err := r.getAllSwitchesByFabric(ctx, fabricName, true)
	if err != nil {
		dg.AddError("Error Reading Inventory", err.Error())
		return
	}

	// iterate the state data and check if if the switch exists in the retrieved data

	for _, sw := range stateData.Switches {
		if sw.SerialNumber != "" {
			if _, ok := switches.SwitchesBySerial[sw.SerialNumber]; !ok {
				tflog.Warn(ctx, "Switch not found in inventory", map[string]interface{}{
					"serial": sw.SerialNumber,
				})
				continue
			} else {
				swDelList = append(swDelList, sw.SerialNumber)
			}
		} else if sw.IpAddress != "" {
			if swByIP, ok := switches.SwitchesByIP[sw.IpAddress]; ok {
				swDelList = append(swDelList, swByIP.SerialNumber)
			} else {
				tflog.Warn(ctx, "Switch not found in inventory", map[string]interface{}{
					"ip": sw.IpAddress,
				})
			}
		} else {
			tflog.Warn(ctx, "State entry has no serial or IP - likely corrupt state", map[string]interface{}{
				"serial": sw.SerialNumber,
				"ip":     sw.IpAddress,
			})
		}

	}

	// Remove credentials first
	if len(swDelList) > 0 {
		tflog.Info(ctx, "Switches to be removed from fabric", map[string]interface{}{
			"fabric":   fabricName,
			"switches": swDelList,
		})
		credRemoveReq := RemoveSwitchesRequest{
			SwitchIds: swDelList,
		}

		payload, err := json.Marshal(credRemoveReq)
		if err != nil {
			dg.AddError("Error Deleting Inventory", fmt.Sprintf("Could not marshal credentials remove request: %v", err))
			return
		}

		tflog.Info(ctx, "Removing credentials", map[string]interface{}{
			"fabric":   fabricName,
			"switches": swDelList,
		})
		invAPI.SetOperation(api.OpRemoveCredentials)
		_, err = invAPI.Post(payload)
		if err != nil {
			tflog.Warn(ctx, "Could not remove credentials, continuing with switch removal", map[string]interface{}{
				"error": err.Error(),
			})
		}

		tflog.Info(ctx, "Removing switches", map[string]interface{}{
			"fabric":   fabricName,
			"switches": swDelList,
		})
		// Remove switches
		r.removeSwitches(ctx, dg, invAPI, swDelList)
	}

}

// rscImportInventory imports existing inventory switches
func (r *inventorySwitchResource) rscImportInventory(ctx context.Context, dg *diag.Diagnostics, id string, resp *resource.ImportStateResponse) {
	// ID format: fabric_name/serial1,serial2,serial3
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 {
		dg.AddError("Invalid Import ID", "Expected format: fabric_name/serial1,serial2,serial3")
		return
	}

	fabricName := parts[0]
	serialsStr := parts[1]
	serials := strings.Split(serialsStr, ",")

	if fabricName == "" || len(serials) == 0 {
		dg.AddError("Invalid Import ID", "Fabric name and at least one serial number are required")
		return
	}

	// Create a model to populate
	input := &InventorySwitchModel{}
	input.FabricName = types.StringValue(fabricName)
	input.Mode = types.StringValue(ModeDiscovery)

	// Initialize switches map with serials
	switchesData := &NDFCInventorySwitchModel{
		FabricName: fabricName,
		Switches:   make(map[string]NDFCSwitchesValue),
	}
	for _, serial := range serials {
		serial = strings.TrimSpace(serial)
		if serial != "" {
			switchesData.Switches[serial] = NDFCSwitchesValue{}
		}
	}
	input.SetModelData(switchesData)

	// Read the actual state
	r.rscGetInventory(ctx, dg, input)
	if dg.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, input)...)
}
