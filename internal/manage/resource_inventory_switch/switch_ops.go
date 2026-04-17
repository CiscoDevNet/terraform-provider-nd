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
	"terraform-provider-nd/internal/manage/api"
	"time"

	. "terraform-provider-nd/internal/common/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// --- Removed structs (replaced by generated NDFCSwitchesValue / NDFCInventorySwitchModel) ---
//
// // ShallowDiscoveryRequest represents the request payload for shallow discovery
// // Replaced by: marshaling NDFCInventorySwitchModel directly
// type ShallowDiscoveryRequest struct {
// 	SeedIpCollection   []string `json:"seedIpCollection"`
// 	MaxHop             int64    `json:"maxHop,omitempty"`
// 	PlatformType       string   `json:"platformType,omitempty"`
// 	SnmpV3AuthProtocol string   `json:"snmpV3AuthProtocol,omitempty"`
// 	Username           string   `json:"username,omitempty"`
// 	Password           string   `json:"password,omitempty"`
// 	CredentialStore    string   `json:"remoteCredentialStore,omitempty"`
// 	CredentialStoreKey string   `json:"remoteCredentialStoreKey,omitempty"`
// }
//
// // SwitchDiscoveryModel represents a switch for adding to fabric
// // Replaced by: NDFCSwitchesValue in AddSwitchesRequest
// type SwitchDiscoveryModel struct {
// 	Hostname     string `json:"hostname,omitempty"`
// 	IP           string `json:"ip,omitempty"`
// 	SerialNumber string `json:"serialNumber"`
// 	Model        string `json:"model,omitempty"`
// 	SwitchRole   string `json:"switchRole,omitempty"`
// 	VdcId        *int64 `json:"vdcId,omitempty"`
// 	VdcMac       string `json:"vdcMac,omitempty"`
// }
//
// // SwitchDataResponse represents API response for a switch
// // Replaced by: NDFCSwitchesValue in waitForManageable
// type SwitchDataResponse struct {
// 	SerialNumber    string `json:"serialNumber"`
// 	Hostname        string `json:"hostname"`
// 	IP              string `json:"ip"`
// 	Model           string `json:"model"`
// 	SoftwareVersion string `json:"softwareVersion"`
// 	SwitchRole      string `json:"switchRole"`
// 	Status          string `json:"status"`
// 	StatusReason    string `json:"statusReason,omitempty"`
// 	VdcId           *int64 `json:"vdcId,omitempty"`
// 	VdcMac          string `json:"vdcMac,omitempty"`
// }
//
// // DiscoverySwitchStatus represents status of a discovered switch
// // Replaced by: NDFCSwitchesValue in DiscoveryStatusResponse
// type DiscoverySwitchStatus struct {
// 	SerialNumber    string `json:"serialNumber"`
// 	Status          string `json:"status"`
// 	StatusReason    string `json:"statusReason,omitempty"`
// 	Hostname        string `json:"hostname,omitempty"`
// 	IP              string `json:"ip,omitempty"`
// 	Model           string `json:"model,omitempty"`
// 	SoftwareVersion string `json:"softwareVersion,omitempty"`
// }
// --- End removed structs ---

// AddSwitchesRequest represents the request to add switches to a fabric
type AddSwitchesRequest struct {
	Switches                 []NDFCSwitchesValue `json:"switches"`
	PlatformType             string              `json:"platformType,omitempty"`
	PreserveConfig           bool                `json:"preserveConfig"`
	Username                 string              `json:"username,omitempty"`
	Password                 string              `json:"password,omitempty"`
	SnmpV3AuthProtocol       string              `json:"snmpV3AuthProtocol,omitempty"`
	RemoteCredentialStore    string              `json:"remoteCredentialStore,omitempty"`
	RemoteCredentialStoreKey string              `json:"remoteCredentialStoreKey,omitempty"`
	MaxHop                   *int64              `json:"maxHop,omitempty"`
	UseCredentialForWrite    bool                `json:"useCredentialForWrite,omitempty"`
}

// BootstrapSwitchRequest represents POAP bootstrap request
type BootstrapSwitchRequest struct {
	Switches []BootstrapSwitchModel `json:"switches"`
}

// BootstrapSwitchModel represents a switch for POAP bootstrap
type BootstrapSwitchModel struct {
	GatewayIpMask         string `json:"gatewayIpMask"`
	Model                 string `json:"model,omitempty"`
	SoftwareVersion       string `json:"softwareVersion,omitempty"`
	Password              string `json:"password,omitempty"`
	DiscoveryAuthProtocol string `json:"discoveryAuthProtocol,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	IP                    string `json:"ip,omitempty"`
	SerialNumber          string `json:"serialNumber"`
	InInventory           bool   `json:"inInventory,omitempty"`
	PublicKey             string `json:"publicKey,omitempty"`
	FingerPrint           string `json:"fingerPrint,omitempty"`
}

// SwitchCredentialsRequest represents credentials save request
type SwitchCredentialsRequest struct {
	SwitchIds      []string `json:"switchIds"`
	SwitchUsername string   `json:"switchUsername"`
	SwitchPassword string   `json:"switchPassword"`
}

// RemoveSwitchesRequest represents the request to remove switches
type RemoveSwitchesRequest struct {
	SwitchIds []string `json:"switchIds"`
}

type SwitchRole struct {
	SwitchId string `json:"switchId"`
	Role     string `json:"role"`
}

// SwitchRoleUpdateRequest represents the request to update switch role
type SwitchRoleUpdateRequest struct {
	SwitchRoles []SwitchRole `json:"switchRoles"`
}

// DiscoveryStatusResponse represents discovery status response
type DiscoveryStatusResponse struct {
	Switches []NDFCSwitchesValue `json:"switches,omitempty"`
}

// FabricSwitchAdditionalData represents the additionalData nested object in GET /switches response
type FabricSwitchAdditionalData struct {
	DiscoveryStatus      string `json:"discoveryStatus,omitempty"`
	DiscoveredSystemMode string `json:"discoveredSystemMode,omitempty"`
	SystemMode           string `json:"systemMode,omitempty"`
}

// FabricSwitchEntry represents a single switch in the GET /switches response
type FabricSwitchEntry struct {
	NDFCSwitchesValue
	AdditionalData FabricSwitchAdditionalData `json:"additionalData,omitempty"`
}

// FabricSwitchesResponse represents the GET /fabrics/{fabric}/switches response
type FabricSwitchesResponse struct {
	Switches         []FabricSwitchEntry           `json:"switches,omitempty"`
	SwitchesBySerial map[string]*FabricSwitchEntry `json:"-"`
	SwitchesByIP     map[string]*FabricSwitchEntry `json:"-"`
}

// Helper functions

// shallowDiscover performs a shallow discovery for the given switch data.
// It collects seed IPs, posts the discovery request, and returns the parsed response.
func (r *inventorySwitchResource) shallowDiscover(ctx context.Context, invAPI *api.InventoryAPI, switchesData *NDFCInventorySwitchModel) (*DiscoveryStatusResponse, error) {
	seedIPs := r.collectSeedIPs(switchesData)
	if len(seedIPs) == 0 {
		return nil, fmt.Errorf("no seed IPs found in switch configuration")
	}

	switchesData.SeedIpCollection = seedIPs

	payload, err := json.Marshal(switchesData)
	if err != nil {
		return nil, fmt.Errorf("could not marshal discovery request: %v", err)
	}

	invAPI.SetOperation(api.OpShallowDiscovery)
	respData, err := invAPI.Post(payload)
	if err != nil {
		return nil, fmt.Errorf("shallow discovery failed: %v", err)
	}

	var discovery DiscoveryStatusResponse
	if err := json.Unmarshal([]byte(respData.Raw), &discovery); err != nil {
		return nil, fmt.Errorf("could not parse discovery response: %v", err)
	}

	tflog.Info(ctx, "Shallow discovery completed", map[string]interface{}{
		"seed_ips": seedIPs,
		"switches": len(discovery.Switches),
	})

	return &discovery, nil
}

// SwitchReadinessResult holds the result of a switchesReady check.
type SwitchReadinessResult struct {
	Ready          bool
	NeedRediscover []string
	Found          int
	Expected       int
}

// switchesReady checks whether all switches identified by serialSet are present
// and ready in the fabric. It returns a SwitchReadinessResult summarising the state.
func (r *inventorySwitchResource) switchesReady(ctx context.Context, fabricName string, serialSet map[string]bool) (*SwitchReadinessResult, error) {
	resp, err := r.getAllSwitchesByFabric(ctx, fabricName, false)
	if err != nil {
		return nil, err
	}

	result := &SwitchReadinessResult{
		Ready:    true,
		Expected: len(serialSet),
	}

	for _, sw := range resp.Switches {
		if !serialSet[sw.SerialNumber] {
			continue
		}
		result.Found++

		// Check migration mode — both discoveredSystemMode and systemMode must be "normal"
		if sw.AdditionalData.DiscoveredSystemMode != "normal" || sw.AdditionalData.SystemMode != "normal" {
			tflog.Debug(ctx, "Switch in migration mode, waiting", map[string]interface{}{
				"serial":               sw.SerialNumber,
				"discoveredSystemMode": sw.AdditionalData.DiscoveredSystemMode,
				"systemMode":           sw.AdditionalData.SystemMode,
			})
			result.NeedRediscover = append(result.NeedRediscover, sw.SerialNumber)
			result.Ready = false
			continue
		}

		// Check discovery status — trigger rediscover if not ok
		if sw.AdditionalData.DiscoveryStatus != "ok" {
			tflog.Debug(ctx, "Switch discovery not ok, will rediscover", map[string]interface{}{
				"serial": sw.SerialNumber,
				"status": sw.AdditionalData.DiscoveryStatus,
			})
			result.NeedRediscover = append(result.NeedRediscover, sw.SerialNumber)
			result.Ready = false
			continue
		}
	}

	// Not all switches present yet
	if result.Found < result.Expected {
		tflog.Debug(ctx, "Not all switches in fabric yet", map[string]interface{}{
			"found":    result.Found,
			"expected": result.Expected,
		})
		result.Ready = false
	}

	return result, nil
}

func (r *inventorySwitchResource) getAllSwitchesByFabric(ctx context.Context, fabricName string, fillLookupMaps bool) (*FabricSwitchesResponse, error) {
	invAPI := api.NewInventoryAPI(nil, r.manageClient.ApiClient)
	invAPI.FabricName = fabricName
	invAPI.SetOperation(api.OpGetAllSwitches)

	respData, err := invAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("could not read switches for fabric %s: %w", fabricName, err)
	}

	var resp FabricSwitchesResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal switches response: %w", err)
	}

	// Fill lookup maps if requested
	if fillLookupMaps {
		resp.SwitchesByIP = make(map[string]*FabricSwitchEntry, len(resp.Switches))
		resp.SwitchesBySerial = make(map[string]*FabricSwitchEntry, len(resp.Switches))

		for i := range resp.Switches {
			entry := &resp.Switches[i]
			if entry.SerialNumber != "" {
				resp.SwitchesBySerial[entry.SerialNumber] = entry
			}
			ip := entry.FabricManagementIp
			if ip == "" {
				ip = entry.IpAddress
			}
			if ip != "" {
				resp.SwitchesByIP[ip] = entry
			}
		}
	}
	return &resp, nil
}

func (r *inventorySwitchResource) collectSeedIPs(data *NDFCInventorySwitchModel) []string {
	var ips []string
	for _, sw := range data.Switches {
		if sw.IpAddress != "" {
			ips = append(ips, sw.IpAddress)
		}
	}
	return ips
}

func (r *inventorySwitchResource) getSerialNumbers(data *DiscoveryStatusResponse) []string {
	var serials []string
	for _, sw := range data.Switches {
		serials = append(serials, sw.SerialNumber)
	}
	return serials
}

func (r *inventorySwitchResource) getModelSerialNumbers(data *NDFCInventorySwitchModel) []string {
	var serials []string
	for _, sw := range data.Switches {
		if sw.SerialNumber != "" {
			serials = append(serials, sw.SerialNumber)
		} else if sw.IpAddress != "" {
			serials = append(serials, sw.IpAddress)
		}
	}
	return serials
}

func (r *inventorySwitchResource) buildAddSwitchesRequest(data *DiscoveryStatusResponse, input *NDFCInventorySwitchModel) AddSwitchesRequest {

	req := AddSwitchesRequest{
		Switches:                 make([]NDFCSwitchesValue, 0),
		PreserveConfig:           *input.PreserveConfig,
		Username:                 input.Username,
		Password:                 input.Password,
		SnmpV3AuthProtocol:       input.SnmpV3AuthProtocol,
		RemoteCredentialStore:    input.RemoteCredentialStore,
		RemoteCredentialStoreKey: input.RemoteCredentialStoreKey,
		MaxHop:                   input.MaxHop,
		PlatformType:             "nx-os",
		UseCredentialForWrite:    true,
	}

	for _, sw := range data.Switches {
		req.Switches = append(req.Switches, NDFCSwitchesValue{
			SerialNumber:    sw.SerialNumber,
			Hostname:        sw.Hostname,
			IpAddress:       sw.IpAddress,
			Model:           sw.Model,
			SwitchRole:      sw.SwitchRole,
			VdcId:           sw.VdcId,
			VdcMac:          sw.VdcMac,
			SoftwareVersion: sw.SoftwareVersion,
			PreserveConfig:  input.PreserveConfig,
		})
	}

	return req
}

func (r *inventorySwitchResource) buildBootstrapRequest(data *NDFCInventorySwitchModel) BootstrapSwitchRequest {
	req := BootstrapSwitchRequest{}

	for serial, sw := range data.Switches {
		req.Switches = append(req.Switches, BootstrapSwitchModel{
			SerialNumber:          serial,
			Hostname:              sw.Hostname,
			IP:                    sw.IpAddress,
			Model:                 sw.Model,
			SoftwareVersion:       sw.SoftwareVersion,
			GatewayIpMask:         sw.GatewayIpMask,
			Password:              sw.PoapPassword,
			DiscoveryAuthProtocol: sw.DiscoveryAuthProtocol,
		})
	}

	return req
}

// triggerRediscovery sends a rediscovery request for the given switch serials.
func (r *inventorySwitchResource) triggerRediscovery(ctx context.Context, invAPI *api.InventoryAPI, serials []string) {
	tflog.Info(ctx, "Triggering rediscovery", map[string]interface{}{
		"serials": serials,
	})
	rediscoverReq := RemoveSwitchesRequest{SwitchIds: serials}
	payload, err := json.Marshal(rediscoverReq)
	if err == nil {
		invAPI.SetOperation(api.OpRediscover)
		_, _ = invAPI.Post(payload)
	}
}

// waitForManageable polls until all specified switches are ready in the fabric.
// Uses switchesReady for the readiness check and triggerRediscovery for switches that need it.
func (r *inventorySwitchResource) waitForManageable(ctx context.Context, invAPI *api.InventoryAPI, serials []string) error {
	deadline := time.Now().Add(MaxDiscoveryWaitTime)

	serialSet := make(map[string]bool, len(serials))
	for _, s := range serials {
		serialSet[s] = true
	}

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled waiting for switches: %v", ctx.Err())
		default:
		}

		result, err := r.switchesReady(ctx, invAPI.FabricName, serialSet)
		if err != nil {
			select {
			case <-time.After(PollInterval):
			case <-ctx.Done():
				return fmt.Errorf("context cancelled waiting for switches: %v", ctx.Err())
			}
			continue
		}

		if len(result.NeedRediscover) > 0 {
			r.triggerRediscovery(ctx, invAPI, result.NeedRediscover)
		}

		if result.Ready {
			tflog.Info(ctx, "All switches manageable", map[string]interface{}{
				"count": len(serials),
			})
			return nil
		}

		select {
		case <-time.After(PollInterval):
		case <-ctx.Done():
			return fmt.Errorf("context cancelled waiting for switches: %v", ctx.Err())
		}
	}

	return fmt.Errorf("timeout waiting for switches to become manageable")
}

func (r *inventorySwitchResource) updateSwitchRoles(ctx context.Context, dg *diag.Diagnostics, invAPI *api.InventoryAPI, discovery *DiscoveryStatusResponse, config *NDFCInventorySwitchModel) {
	// Build IP -> desired role lookup from config
	ipToRole := make(map[string]string)
	for _, sw := range config.Switches {
		log.Printf("Switch: %v:%v", sw.IpAddress, sw.SwitchRole)
		if sw.IpAddress != "" && sw.SwitchRole != "" {
			ipToRole[sw.IpAddress] = sw.SwitchRole
		}
	}

	roleReq := SwitchRoleUpdateRequest{}
	roleReq.SwitchRoles = []SwitchRole{}

	for _, sw := range discovery.Switches {
		role, found := ipToRole[sw.IpAddress]
		if !found || role == "" {
			continue
		}
		tflog.Info(ctx, "Updated switch role", map[string]interface{}{
			"serial": sw.SerialNumber,
			"role":   role,
		})
		roleReq.SwitchRoles = append(roleReq.SwitchRoles, SwitchRole{
			SwitchId: sw.SerialNumber,
			Role:     role,
		})
	}

	payload, err := json.Marshal(roleReq)
	if err != nil {
		dg.AddError("Error Updating Role", fmt.Sprintf("Could not marshal role update:%v", err))
		return
	}

	invAPI.FabricName = config.FabricName
	invAPI.SetOperation(api.OpUpdateSwitchRole)
	resp, err := invAPI.Post(payload)
	if err != nil {
		dg.AddError("Error Updating Role", fmt.Sprintf("Could not update role for %v:%v: %s", config.FabricName, err, resp.String()))
		return
	}

	tflog.Info(ctx, "Updated switch roles", map[string]interface{}{
		"fabric_name": config.FabricName,
		"count":       len(roleReq.SwitchRoles),
	})
}

func (r *inventorySwitchResource) removeSwitches(ctx context.Context, dg *diag.Diagnostics, invAPI *api.InventoryAPI, serials []string) {
	if len(serials) == 0 {
		return
	}

	removeReq := RemoveSwitchesRequest{
		SwitchIds: serials,
	}

	payload, err := json.Marshal(removeReq)
	if err != nil {
		dg.AddError("Error Removing Switches", fmt.Sprintf("Could not marshal remove request: %v", err))
		return
	}

	invAPI.SetOperation(api.OpRemoveSwitches)
	_, err = invAPI.Post(payload)
	if err != nil {
		dg.AddError("Error Removing Switches", fmt.Sprintf("Could not remove switches: %v", err))
		return
	}

	tflog.Info(ctx, "Removed switches from fabric", map[string]interface{}{
		"fabric_name": invAPI.FabricName,
		"serials":     serials,
	})
}

func (r *inventorySwitchResource) diffSwitches(plan, state *NDFCInventorySwitchModel) map[string][]NDFCSwitchesValue {
	action := make(map[string][]NDFCSwitchesValue)

	for serial, planSw := range plan.Switches {
		if _, exists := state.Switches[serial]; !exists {
			log.Printf("New switch in plan: %s", serial)
			action["add"] = append(action["add"], planSw)
		} else {
			cf := false
			swState := state.Switches[serial]
			updateAction := planSw.CreatePlan(swState, &cf)
			if updateAction == RequiresUpdate {
				log.Printf("Update switch config in plan: %s", serial)
				action["add"] = append(action["add"], planSw)
			} else if updateAction == RequiresReplace {
				// delete and create
				log.Printf("Update switch config in plan requires replacement: %s", serial)
				action["del"] = append(action["del"], swState)
				action["add"] = append(action["add"], planSw)
			} else {
				log.Printf("No changes to switch %s", serial)
			}
		}
	}

	// Find switches to remove (in state but not in plan)
	for serial := range state.Switches {
		if _, exists := plan.Switches[serial]; !exists {
			log.Printf("Switch removed in plan: %s", serial)
			action["del"] = append(action["del"], state.Switches[serial])
		}
	}

	return action
}

// createBootstrapSwitches handles POAP bootstrap switch addition
func (r *inventorySwitchResource) createBootstrapSwitches(ctx context.Context, dg *diag.Diagnostics, invAPI *api.InventoryAPI, switchesData *NDFCInventorySwitchModel) {

	bootstrapReq := r.buildBootstrapRequest(switchesData)
	payload, err := json.Marshal(bootstrapReq)
	if err != nil {
		dg.AddError("Error Creating Inventory", fmt.Sprintf("Could not marshal bootstrap request: %v", err))
		return
	}

	invAPI.SetOperation(api.OpBootstrap)
	res, err := invAPI.Post(payload)
	if err != nil {
		dg.AddError("Error Creating Inventory", fmt.Sprintf("Bootstrap failed: %v: %s", err, res.String()))
		return
	}

	tflog.Info(ctx, "Bootstrap initiated for switches", map[string]interface{}{
		"fabric_name": invAPI.FabricName,
		"count":       len(bootstrapReq.Switches),
	})

	// Wait for switches to become manageable
	err = r.waitForManageable(ctx, invAPI, r.getModelSerialNumbers(switchesData))
	if err != nil {
		dg.AddError("Error Creating Inventory", fmt.Sprintf("Wait for manageable failed: %v", err))
		return
	}
}
