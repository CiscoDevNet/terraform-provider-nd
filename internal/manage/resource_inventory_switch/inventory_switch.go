// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
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

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

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
	StatusNotReachable   = "notReachable"

	MaxDiscoveryWaitTime    = 10 * time.Minute
	PollInterval            = 10 * time.Second
	ObservationWindow       = 3 * time.Minute
	ObservationPollInterval = 30 * time.Second
)

// rscCreateInventory creates a single inventory switch
func (r *inventorySwitchResource) rscCreateInventory(ctx context.Context, dg *diag.Diagnostics, input *InventorySwitchModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}
	switchData := input.GetModelData()

	// Fill in a mandatory param thats not part of config (tf_hide)
	switchData.MaxHop = new(int64)
	*switchData.MaxHop = 0

	invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	invAPI.FabricName = switchData.FabricName

	log.Printf("Creating inventory for fabric %s with mode %s", switchData.FabricName, switchData.Mode)

	// Both discovery and bootstrap use the FSM
	fsm := NewInventoryFSM(ctx, r, true, invAPI, switchData, dg)
	fsm.Run()
	if dg.HasError() {
		return
	}

	// Read back the created state
	if !r.rscGetInventory(ctx, dg, input) && !dg.HasError() {
		dg.AddError("Error Reading Inventory", "Switch not found in fabric after create")
	}
}

// rscGetInventory reads the current state of the single inventory switch
func (r *inventorySwitchResource) rscGetInventory(ctx context.Context, dg *diag.Diagnostics, input *InventorySwitchModel) bool {
	DumpInventorySwitchModel("rscGetInventory INPUT", input)
	fabricName := input.FabricName.ValueString()

	switchData := input.GetModelData()
	lookupSerial := switchData.SwitchDetail.SerialNumber
	lookupIP := switchData.SwitchDetail.IpAddress

	// Fetch all switches in the fabric
	resp, err := r.getAllSwitchesByFabric(ctx, fabricName, false)
	if err != nil {
		dg.AddError("Error Reading Inventory", err.Error())
		return false
	}

	// Find our switch by serial or IP
	outData := *switchData
	found := false
	var matchedEntry *FabricSwitchEntry
	for i := range resp.Switches {
		sw := &resp.Switches[i]
		if lookupSerial != "" && sw.SerialNumber == lookupSerial {
			log.Printf("Found switch by serial %s in fabric %s", lookupSerial, fabricName)
			matchedEntry = sw
			found = true
			break
		}
		if lookupIP != "" && (sw.FabricManagementIp == lookupIP || sw.IpAddress == lookupIP) {
			log.Printf("Found switch by IP %s in fabric %s", lookupIP, fabricName)
			matchedEntry = sw
			found = true
			break
		}
	}

	if found {
		outData.SwitchDetail = matchedEntry.NDFCSwitchDetailValue

		// Set id from serial number
		outData.Id = matchedEntry.SerialNumber

		// Populate fields from additionalData
		if matchedEntry.AdditionalData.PlatformType != "" {
			outData.PlatformType = matchedEntry.AdditionalData.PlatformType
		}

		// source_interface_name and source_vrf_name: only fill from API
		// if the config had them set; otherwise reset to empty (null)
		if switchData.SourceInterfaceName != "" {
			if matchedEntry.AdditionalData.SourceInterfaceName != "" {
				outData.SourceInterfaceName = matchedEntry.AdditionalData.SourceInterfaceName
			}
		} else {
			outData.SourceInterfaceName = ""
		}
		if switchData.SourceVrfName != "" {
			if matchedEntry.AdditionalData.SourceVrfName != "" {
				outData.SourceVrfName = matchedEntry.AdditionalData.SourceVrfName
			}
		} else {
			outData.SourceVrfName = ""
		}

		// status comes from additionalData.discoveryStatus in GET /switches
		if matchedEntry.AdditionalData.DiscoveryStatus != "" {
			outData.SwitchDetail.Status = matchedEntry.AdditionalData.DiscoveryStatus
		}
		// statusReason is not returned by GET /switches
		outData.SwitchDetail.StatusReason = ""

		// Preserve config-only fields from the input that the API does not return
		outData.SwitchDetail.GatewayIpMask = switchData.SwitchDetail.GatewayIpMask
		outData.SwitchDetail.DiscoveryAuthProtocol = switchData.SwitchDetail.DiscoveryAuthProtocol

		input.SetModelData(&outData)
		DumpInventorySwitchModel("rscGetInventory OUTPUT", input)
		log.Printf("Read inventory for fabric %s, found=%v", fabricName, found)
	} else {
		log.Printf("Switch not found in fabric %s (serial=%s, ip=%s)", fabricName, lookupSerial, lookupIP)
	}
	return found
}

// rscUpdateInventory updates a single inventory switch via direct API calls
func (r *inventorySwitchResource) rscUpdateInventory(ctx context.Context, dg *diag.Diagnostics, plan *InventorySwitchModel, state *InventorySwitchModel) {
	DumpInventorySwitchModel("rscUpdateInventory PLAN", plan)
	DumpInventorySwitchModel("rscUpdateInventory STATE", state)
	fabricName := plan.FabricName.ValueString()
	invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	invAPI.FabricName = fabricName

	planData := plan.GetModelData()
	stateData := state.GetModelData()

	// Compare plan vs state for the single switch
	planSw := planData.SwitchDetail
	stateSw := stateData.SwitchDetail

	// Update switch role if changed
	if planSw.SwitchRole != stateSw.SwitchRole && planSw.SwitchRole != "" {
		serial := stateSw.SerialNumber
		if serial == "" {
			serial = planSw.SerialNumber
		}
		if serial != "" {
			roleReq := SwitchRoleUpdateRequest{
				SwitchRoles: []SwitchRole{
					{SwitchId: serial, Role: planSw.SwitchRole},
				},
			}
			payload, err := json.Marshal(roleReq)
			if err != nil {
				dg.AddError("Error Updating Role", fmt.Sprintf("Could not marshal role update: %v", err))
				return
			}
			invAPI.SetOperation(api.OpUpdateSwitchRole)
			resp, err := invAPI.Post(payload, nil)
			if err != nil {
				dg.AddError("Error Updating Role", fmt.Sprintf("Could not update role: %v: %s", err, resp.String()))
				return
			}
			tflog.Info(ctx, "Updated switch role", map[string]interface{}{
				"serial": serial,
				"role":   planSw.SwitchRole,
			})
		}
	}

	if planData.Mode == ModeDiscovery {
		serial := stateSw.SerialNumber
		if serial == "" {
			serial = planSw.SerialNumber
		}

		// Handle discovery credentials change (username, password, snmp, credential store)
		usernameChanged := planData.DiscoveryUsername != stateData.DiscoveryUsername
		passwordChanged := planData.DiscoveryPassword != stateData.DiscoveryPassword
		snmpChanged := planData.SnmpV3AuthProtocol != stateData.SnmpV3AuthProtocol
		credStoreChanged := planData.RemoteCredentialStore != stateData.RemoteCredentialStore
		credStoreKeyChanged := planData.RemoteCredentialStoreKey != stateData.RemoteCredentialStoreKey

		credsChanged := usernameChanged || passwordChanged || snmpChanged || credStoreChanged || credStoreKeyChanged

		if credsChanged && serial != "" {
			credReq := ChangeDiscoveryCredentialRequest{
				SwitchIds:                []string{serial},
				SnmpV3AuthProtocol:       planData.SnmpV3AuthProtocol,
				Username:                 planData.DiscoveryUsername,
				Password:                 planData.DiscoveryPassword,
				RemoteCredentialStore:    planData.RemoteCredentialStore,
				RemoteCredentialStoreKey: planData.RemoteCredentialStoreKey,
			}
			payload, err := json.Marshal(credReq)
			if err != nil {
				dg.AddError("Error Updating Credentials", fmt.Sprintf("Could not marshal credentials change: %v", err))
				return
			}
			invAPI.SetOperation(api.OpChangeDiscoveryCredentials)
			resp, err := invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
			if err != nil {
				dg.AddError("Error Updating Credentials", fmt.Sprintf("Could not change discovery credentials: %v: %s", err, resp.String()))
				return
			}
			tflog.Info(ctx, "Updated discovery credentials", map[string]interface{}{
				"serial": serial,
			})
		}

		// Handle discovery IP change
		if planSw.IpAddress != stateSw.IpAddress && planSw.IpAddress != "" && serial != "" {
			ipReq := []IpSwitchIdPair{
				{SwitchId: serial, Ip: planSw.IpAddress},
			}
			payload, err := json.Marshal(ipReq)
			if err != nil {
				dg.AddError("Error Updating IP", fmt.Sprintf("Could not marshal IP change: %v", err))
				return
			}
			invAPI.SetOperation(api.OpChangeIpCollection)
			resp, err := invAPI.Post(payload, nil)
			if err != nil {
				dg.AddError("Error Updating IP", fmt.Sprintf("Could not change switch IP: %v: %s", err, resp.String()))
				return
			}
			tflog.Info(ctx, "Updated switch IP", map[string]interface{}{
				"serial": serial,
				"ip":     planSw.IpAddress,
			})
		}

		// Handle discovery source interface or VRF change
		// Skip if plan values are empty (user unset them)
		ifChanged := planData.SourceInterfaceName != stateData.SourceInterfaceName
		vrfChanged := planData.SourceVrfName != stateData.SourceVrfName
		hasValues := planData.SourceInterfaceName != "" || planData.SourceVrfName != ""
		if (ifChanged || vrfChanged) && hasValues && serial != "" {
			ifVrfReq := ChangeDiscoveryInterfaceOrVrfRequest{
				SwitchIds:     []string{serial},
				VrfName:       planData.SourceVrfName,
				InterfaceName: planData.SourceInterfaceName,
			}
			payload, err := json.Marshal(ifVrfReq)
			if err != nil {
				dg.AddError("Error Updating Discovery Interface/VRF", fmt.Sprintf("Could not marshal request: %v", err))
				return
			}
			invAPI.SetOperation(api.OpChangeDiscoveryInterfaceOrVrf)
			resp, err := invAPI.Post(payload, nil)
			if err != nil {
				dg.AddError("Error Updating Discovery Interface/VRF", fmt.Sprintf("Could not change discovery interface/VRF: %v: %s", err, resp.String()))
				return
			}

			// Parse item-level results to detect failures
			var actionResp SwitchActionResponse
			if err := json.Unmarshal([]byte(resp.Raw), &actionResp); err == nil {
				for _, item := range actionResp.Items {
					if item.Status != "success" {
						dg.AddError("Error Updating Discovery Interface/VRF",
							fmt.Sprintf("Switch %s: %s: %s", item.SwitchId, item.Status, item.Message))
						return
					}
				}
			}

			tflog.Info(ctx, "Updated discovery interface/VRF", map[string]interface{}{
				"serial":    serial,
				"interface": planData.SourceInterfaceName,
				"vrf":       planData.SourceVrfName,
			})
		}
	}
	// Read back the updated state
	if !r.rscGetInventory(ctx, dg, plan) && !dg.HasError() {
		dg.AddError("Error Reading Inventory", "Switch not found in fabric after update")
	}
}

// rscDeleteInventory deletes a single inventory switch
func (r *inventorySwitchResource) rscDeleteInventory(ctx context.Context, dg *diag.Diagnostics, state *InventorySwitchModel) {
	fabricName := state.FabricName.ValueString()
	invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	invAPI.FabricName = fabricName

	stateData := state.GetModelData()
	sw := stateData.SwitchDetail

	// Resolve the serial number for the switch
	var serial string
	if sw.SerialNumber != "" {
		serial = sw.SerialNumber
	} else if sw.IpAddress != "" {
		// Look up serial by IP from the fabric
		switches, err := r.getAllSwitchesByFabric(ctx, fabricName, true)
		if err != nil {
			dg.AddError("Error Reading Inventory", err.Error())
			return
		}
		if swByIP, ok := switches.SwitchesByIP[sw.IpAddress]; ok {
			serial = swByIP.SerialNumber
		} else {
			tflog.Warn(ctx, "Switch not found in inventory by IP", map[string]interface{}{
				"ip": sw.IpAddress,
			})
			return
		}
	} else {
		tflog.Warn(ctx, "State has no serial or IP - likely corrupt state", map[string]interface{}{})
		return
	}

	swDelList := []string{serial}

	tflog.Info(ctx, "Switch to be removed from fabric", map[string]interface{}{
		"fabric": fabricName,
		"serial": serial,
	})

	// Remove credentials first
	credRemoveReq := RemoveSwitchesRequest{SwitchIds: swDelList}
	payload, err := json.Marshal(credRemoveReq)
	if err != nil {
		dg.AddError("Error Deleting Inventory", fmt.Sprintf("Could not marshal credentials remove request: %v", err))
		return
	}

	invAPI.SetOperation(api.OpRemoveCredentials)
	res, err := invAPI.Post(payload, nil)
	if err != nil {
		tflog.Warn(ctx, "Could not remove credentials (best-effort), continuing with switch removal", map[string]interface{}{
			"serial": serial,
			"error":  fmt.Sprintf("%v: %s", err, res.String()),
		})
	}

	// Remove switch
	r.removeSwitches(ctx, dg, invAPI, swDelList)
}

// rscImportInventory imports an existing inventory switch
func (r *inventorySwitchResource) rscImportInventory(ctx context.Context, dg *diag.Diagnostics, id string, resp *resource.ImportStateResponse) {
	// ID format: fabric_name/serial_number
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 {
		dg.AddError("Invalid Import ID", "Expected format: fabric_name/serial_number")
		return
	}

	fabricName := parts[0]
	serial := strings.TrimSpace(parts[1])

	if fabricName == "" || serial == "" {
		dg.AddError("Invalid Import ID", "Fabric name and serial number are required")
		return
	}

	// Create a model to populate
	input := &InventorySwitchModel{}
	input.FabricName = types.StringValue(fabricName)
	input.Mode = types.StringValue(ModeDiscovery)

	// Initialize switch detail with serial
	switchData := &NDFCInventorySwitchModel{
		FabricName: fabricName,
		SwitchDetail: NDFCSwitchDetailValue{
			SerialNumber: serial,
		},
	}
	input.SetModelData(switchData)

	// Read the actual state
	r.rscGetInventory(ctx, dg, input)
	if dg.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, input)...)
}
