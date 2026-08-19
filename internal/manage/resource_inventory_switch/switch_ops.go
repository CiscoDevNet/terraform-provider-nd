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
	"terraform-provider-nd/internal/manage/api"
	"time"

	"terraform-provider-nd/internal/common/ndapi"

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
	Switches                 []NDFCSwitchDetailValue `json:"switches"`
	PlatformType             string                  `json:"platformType,omitempty"`
	PreserveConfig           bool                    `json:"preserveConfig"`
	Username                 string                  `json:"username,omitempty"`
	Password                 string                  `json:"password,omitempty"`
	SnmpV3AuthProtocol       string                  `json:"snmpV3AuthProtocol,omitempty"`
	RemoteCredentialStore    string                  `json:"remoteCredentialStore,omitempty"`
	RemoteCredentialStoreKey string                  `json:"remoteCredentialStoreKey,omitempty"`
	MaxHop                   *int64                  `json:"maxHop,omitempty"`
	UseCredentialForWrite    bool                    `json:"useCredentialForWrite,omitempty"`
}

// BootstrapSwitchEntry wraps NDFCSwitchDetailValue with bootstrap-specific credential fields
type BootstrapSwitchEntry struct {
	NDFCSwitchDetailValue
	UseNewCredentials *bool  `json:"useNewCredentials,omitempty"`
	DiscoveryUsername string `json:"discoveryUsername,omitempty"`
	DiscoveryPassword string `json:"discoveryPassword,omitempty"`
}

// BootstrapSwitchRequest represents POAP bootstrap request
type BootstrapSwitchRequest struct {
	Switches []BootstrapSwitchEntry `json:"switches"`
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

// ChangeDiscoveryCredentialRequest represents the request to change discovery credentials
type ChangeDiscoveryCredentialRequest struct {
	SwitchIds                []string `json:"switchIds"`
	SnmpV3AuthProtocol       string   `json:"snmpV3AuthProtocol"`
	Username                 string   `json:"username,omitempty"`
	Password                 string   `json:"password,omitempty"`
	RemoteCredentialStore    string   `json:"remoteCredentialStore,omitempty"`
	RemoteCredentialStoreKey string   `json:"remoteCredentialStoreKey,omitempty"`
}

// IpSwitchIdPair represents a single entry in the changeIpCollection request
type IpSwitchIdPair struct {
	SwitchId string `json:"switchId"`
	Ip       string `json:"ip"`
}

// DiscoveryStatusResponse represents discovery status response
type DiscoveryStatusResponse struct {
	Switches []NDFCSwitchDetailValue `json:"switches,omitempty"`
}

// FabricSwitchAdditionalData represents the additionalData nested object in GET /switches response
type FabricSwitchAdditionalData struct {
	DiscoveryStatus      string `json:"discoveryStatus,omitempty"`
	DiscoveredSystemMode string `json:"discoveredSystemMode,omitempty"`
	SystemMode           string `json:"systemMode,omitempty"`
	PlatformType         string `json:"platformType,omitempty"`
	SourceInterfaceName  string `json:"sourceInterfaceName,omitempty"`
	SourceVrfName        string `json:"sourceVrfName,omitempty"`
}

// ChangeDiscoveryInterfaceOrVrfRequest is the payload for POST changeDiscoveryInterfaceOrVrf
type ChangeDiscoveryInterfaceOrVrfRequest struct {
	SwitchIds     []string `json:"switchIds"`
	VrfName       string   `json:"vrfName"`
	InterfaceName string   `json:"interfaceName,omitempty"`
}

// SwitchActionResponseItem represents a single item in a switch action response
type SwitchActionResponseItem struct {
	SwitchId string `json:"switchId"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

// SwitchActionResponse represents the response from switch action endpoints that return item-level results
type SwitchActionResponse struct {
	Items []SwitchActionResponseItem `json:"items"`
}

// FabricSwitchEntry represents a single switch in the GET /switches response
type FabricSwitchEntry struct {
	NDFCSwitchDetailValue
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

	tflog.Debug(ctx, "Shallow discovery request", map[string]interface{}{
		"seed_ips":        seedIPs,
		"max_hop":         switchesData.MaxHop,
		"mode":            switchesData.Mode,
		"preserve_config": switchesData.PreserveConfig,
	})

	invAPI.SetOperation(api.OpShallowDiscovery)
	// Disable payload logging, to avoid printing sensitive fields
	respData, err := invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		return nil, fmt.Errorf("shallow discovery failed: %v: %s", err, respData.String())
	}

	var discovery DiscoveryStatusResponse
	if err := json.Unmarshal([]byte(respData.Raw), &discovery); err != nil {
		return nil, fmt.Errorf("could not parse discovery response: %v", err)
	}

	tflog.Info(ctx, "Shallow discovery completed", map[string]interface{}{
		"seed_ips": seedIPs,
		"switches": len(discovery.Switches),
	})

	for i, sw := range discovery.Switches {
		tflog.Debug(ctx, "Discovered switch", map[string]interface{}{
			"index":        i,
			"serial":       sw.SerialNumber,
			"ip":           sw.IpAddress,
			"model":        sw.Model,
			"hostname":     sw.Hostname,
			"status":       sw.Status,
			"switch_role":  sw.SwitchRole,
			"software_ver": sw.SoftwareVersion,
		})
	}

	return &discovery, nil
}

// SwitchStatus holds the current state of a single switch from the API.
type SwitchStatus struct {
	DiscoveryStatus   string
	SystemMode        string
	DiscoveredSysMode string
}

// SwitchReadinessResult holds the result of a switchesReady check.
type SwitchReadinessResult struct {
	Ready          bool
	NeedRediscover []string
	Found          int
	Expected       int
	SwitchStates   map[string]SwitchStatus
}

// switchesReady checks whether all switches identified by serialSet are present
// and ready in the fabric. It returns a SwitchReadinessResult summarising the state.
func (r *inventorySwitchResource) switchesReady(ctx context.Context, fabricName string, serialSet map[string]bool) (*SwitchReadinessResult, error) {
	resp, err := r.getAllSwitchesByFabric(ctx, fabricName, false)
	if err != nil {
		return nil, err
	}

	result := &SwitchReadinessResult{
		Ready:        true,
		Expected:     len(serialSet),
		SwitchStates: make(map[string]SwitchStatus),
	}

	for _, sw := range resp.Switches {
		if !serialSet[sw.SerialNumber] {
			continue
		}
		result.Found++

		ss := SwitchStatus{
			DiscoveryStatus:   sw.AdditionalData.DiscoveryStatus,
			SystemMode:        sw.AdditionalData.SystemMode,
			DiscoveredSysMode: sw.AdditionalData.DiscoveredSystemMode,
		}
		result.SwitchStates[sw.SerialNumber] = ss

		// Check system mode — both discoveredSystemMode and systemMode must be "normal"
		if sw.AdditionalData.DiscoveredSystemMode != "normal" || sw.AdditionalData.SystemMode != "normal" {
			reason := "system mode not normal"
			if sw.AdditionalData.DiscoveredSystemMode == "notApplicable" {
				reason = "awaiting system mode detection (switch may be rebooting)"
			} else if sw.AdditionalData.SystemMode == "migration" || sw.AdditionalData.DiscoveredSystemMode == "migration" {
				reason = "switch in migration mode"
			}
			tflog.Debug(ctx, "Switch not ready: "+reason, map[string]interface{}{
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
			tflog.Debug(ctx, "Switch not ready: discoveryStatus="+sw.AdditionalData.DiscoveryStatus, map[string]interface{}{
				"serial": sw.SerialNumber,
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
	invAPI := api.NewInventoryAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	invAPI.FabricName = fabricName
	invAPI.SetOperation(api.OpGetAllSwitches)

	respData, err := invAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("could not read switches for fabric %s: %w: %s", fabricName, err, string(respData))
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
	if data.SwitchDetail.IpAddress != "" {
		return []string{data.SwitchDetail.IpAddress}
	}
	return nil
}

func (r *inventorySwitchResource) getSerialNumbers(data *DiscoveryStatusResponse) []string {
	var serials []string
	for _, sw := range data.Switches {
		serials = append(serials, sw.SerialNumber)
	}
	return serials
}

func (r *inventorySwitchResource) getModelSerialNumbers(data *NDFCInventorySwitchModel) []string {
	if data.SwitchDetail.SerialNumber != "" {
		return []string{data.SwitchDetail.SerialNumber}
	}
	if data.SwitchDetail.IpAddress != "" {
		return []string{data.SwitchDetail.IpAddress}
	}
	return nil
}

func (r *inventorySwitchResource) buildAddSwitchesRequest(data *DiscoveryStatusResponse, input *NDFCInventorySwitchModel) AddSwitchesRequest {

	req := AddSwitchesRequest{
		Switches:                 make([]NDFCSwitchDetailValue, 0),
		PreserveConfig:           *input.PreserveConfig,
		Username:                 input.DiscoveryUsername,
		Password:                 input.DiscoveryPassword,
		SnmpV3AuthProtocol:       input.SnmpV3AuthProtocol,
		RemoteCredentialStore:    input.RemoteCredentialStore,
		RemoteCredentialStoreKey: input.RemoteCredentialStoreKey,
		MaxHop:                   input.MaxHop,
		PlatformType:             "nx-os",
		UseCredentialForWrite:    input.DiscoveryCredForLan,
	}

	for _, sw := range data.Switches {
		req.Switches = append(req.Switches, NDFCSwitchDetailValue{
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
	sw := data.SwitchDetail
	entry := BootstrapSwitchEntry{
		NDFCSwitchDetailValue: NDFCSwitchDetailValue{
			SerialNumber:          sw.SerialNumber,
			Hostname:              sw.Hostname,
			IpAddress:             sw.IpAddress,
			Model:                 sw.Model,
			SoftwareVersion:       sw.SoftwareVersion,
			GatewayIpMask:         sw.GatewayIpMask,
			SwitchPassword:        data.BootstrapPassword,
			DiscoveryAuthProtocol: sw.DiscoveryAuthProtocol,
		},
		UseNewCredentials: &data.UseNewCredentials,
	}
	if data.UseNewCredentials {
		entry.DiscoveryUsername = data.DiscoveryUsername
		entry.DiscoveryPassword = data.DiscoveryPassword
	}
	return BootstrapSwitchRequest{Switches: []BootstrapSwitchEntry{entry}}
}

// triggerRediscovery sends a rediscovery request for the given switch serials.
func (r *inventorySwitchResource) triggerRediscovery(ctx context.Context, invAPI *api.InventoryAPI, serials []string) {
	tflog.Debug(ctx, "Triggering rediscovery", map[string]interface{}{
		"serials": serials,
	})
	rediscoverReq := RemoveSwitchesRequest{SwitchIds: serials}
	payload, err := json.Marshal(rediscoverReq)
	if err == nil {
		invAPI.SetOperation(api.OpRediscover)
		_, _ = invAPI.Post(payload, nil)
	}
}

// waitForManageable polls until all specified switches are ready in the fabric.
// It uses a state-transition tracking approach:
//   - For the first ObservationWindow (3 min), it watches for instability (unreachable, migration).
//   - If a switch goes unstable then recovers to ok during the window, it's considered ready.
//   - After the observation window, any poll where all switches are ok is accepted immediately.
//   - Overall timeout is MaxDiscoveryWaitTime (10 min).
func (r *inventorySwitchResource) waitForManageable(ctx context.Context, invAPI *api.InventoryAPI, serials []string) error {
	deadline := time.Now().Add(MaxDiscoveryWaitTime)
	observationEnd := time.Now().Add(ObservationWindow)

	serialSet := make(map[string]bool, len(serials))
	for _, s := range serials {
		serialSet[s] = true
	}

	// Per-switch tracking: did we see this switch go unstable during observation?
	sawUnstable := make(map[string]bool, len(serials))
	// Per-switch tracking: did the switch recover after going unstable?
	recoveredAfterUnstable := make(map[string]bool, len(serials))

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

		inObservationWindow := time.Now().Before(observationEnd)

		phase := "post-observation"
		if inObservationWindow {
			phase = "observation"
		}

		// Track state transitions for each switch
		for serial, ss := range result.SwitchStates {
			isUnstable := ss.DiscoveryStatus != "ok" || ss.SystemMode != "normal" || ss.DiscoveredSysMode != "normal"

			if isUnstable {
				firstTime := !sawUnstable[serial]
				sawUnstable[serial] = true
				recoveredAfterUnstable[serial] = false
				if firstTime {
					tflog.Info(ctx, "Switch became unstable, tracking for recovery", map[string]interface{}{
						"serial":          serial,
						"discoveryStatus": ss.DiscoveryStatus,
						"systemMode":      ss.SystemMode,
						"phase":           phase,
					})
				}
			} else if sawUnstable[serial] && !recoveredAfterUnstable[serial] {
				recoveredAfterUnstable[serial] = true
				tflog.Info(ctx, "Switch recovered after instability", map[string]interface{}{
					"serial": serial,
					"phase":  phase,
				})
			}
		}

		if result.Ready {
			if !inObservationWindow {
				// Observation window passed; all switches currently ok — accept it
				tflog.Info(ctx, "All switches ready (observation window elapsed)", map[string]interface{}{
					"count": len(serials),
				})
				return nil
			}

			// Still in observation window — check if all switches that went unstable have recovered
			allRecovered := true
			for _, s := range serials {
				if sawUnstable[s] && !recoveredAfterUnstable[s] {
					allRecovered = false
					break
				}
			}

			// If at least one switch went unstable and all have recovered, we're done
			anySawUnstable := false
			for _, s := range serials {
				if sawUnstable[s] {
					anySawUnstable = true
					break
				}
			}

			if anySawUnstable && allRecovered {
				tflog.Info(ctx, "All switches recovered after instability — ready", map[string]interface{}{
					"count": len(serials),
				})
				return nil
			}

			// Otherwise keep observing
			tflog.Debug(ctx, "Switches ok, observing for potential reboot", map[string]interface{}{
				"remainingObservation": time.Until(observationEnd).Round(time.Second).String(),
			})
		}

		// Use slower poll during observation, faster after
		pollWait := PollInterval
		if inObservationWindow {
			pollWait = ObservationPollInterval
		}
		select {
		case <-time.After(pollWait):
		case <-ctx.Done():
			return fmt.Errorf("context cancelled waiting for switches: %v", ctx.Err())
		}
	}

	return fmt.Errorf("timeout waiting for switches to become manageable")
}

func (r *inventorySwitchResource) updateSwitchRoles(ctx context.Context, dg *diag.Diagnostics, invAPI *api.InventoryAPI, discovery *DiscoveryStatusResponse, config *NDFCInventorySwitchModel) {
	desiredRole := config.SwitchDetail.SwitchRole
	desiredIP := config.SwitchDetail.IpAddress

	roleReq := SwitchRoleUpdateRequest{}
	roleReq.SwitchRoles = []SwitchRole{}

	for _, sw := range discovery.Switches {
		if sw.IpAddress != desiredIP || desiredRole == "" {
			continue
		}
		tflog.Info(ctx, "Updated switch role", map[string]interface{}{
			"serial": sw.SerialNumber,
			"role":   desiredRole,
		})
		roleReq.SwitchRoles = append(roleReq.SwitchRoles, SwitchRole{
			SwitchId: sw.SerialNumber,
			Role:     desiredRole,
		})
	}

	payload, err := json.Marshal(roleReq)
	if err != nil {
		dg.AddError("Error Updating Role", fmt.Sprintf("Could not marshal role update:%v", err))
		return
	}

	invAPI.FabricName = config.FabricName
	invAPI.SetOperation(api.OpUpdateSwitchRole)
	resp, err := invAPI.Post(payload, nil)
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
	res, err := invAPI.Post(payload, nil)
	if err != nil {
		tflog.Warn(ctx, "Could not remove switches (best-effort)", map[string]interface{}{
			"serials": serials,
			"error":   fmt.Sprintf("%v: %s", err, res.String()),
		})
		return
	}

	tflog.Info(ctx, "Removed switches from fabric", map[string]interface{}{
		"fabric_name": invAPI.FabricName,
		"serials":     serials,
	})
}

// BootstrapBuildResult holds the result of building a bootstrap request from API data.
type BootstrapBuildResult struct {
	Request              BootstrapSwitchRequest
	MissingFromBootstrap []string
}

// BootstrapListResponse represents the response from the bootstrap list API.
type BootstrapListResponse struct {
	Switches []NDFCSwitchDetailValue `json:"switches,omitempty"`
}

// queryBootstrapList queries the bootstrap API and returns a map of serial -> entry.
func (r *inventorySwitchResource) queryBootstrapList(ctx context.Context, invAPI *api.InventoryAPI) (map[string]NDFCSwitchDetailValue, error) {
	invAPI.SetOperation(api.OpGetBootstrapList)
	respData, err := invAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("could not query bootstrap list: %w: %s", err, string(respData))
	}

	tflog.Debug(ctx, "Bootstrap list raw response", map[string]interface{}{
		"body": string(respData),
	})

	// Try parsing as a direct array first, then as a wrapped object
	var switchList []NDFCSwitchDetailValue
	if err := json.Unmarshal(respData, &switchList); err != nil {
		var resp BootstrapListResponse
		if err2 := json.Unmarshal(respData, &resp); err2 != nil {
			return nil, fmt.Errorf("could not parse bootstrap list response: %w (also tried array: %v)", err2, err)
		}
		switchList = resp.Switches
	}

	entries := make(map[string]NDFCSwitchDetailValue, len(switchList))
	for _, sw := range switchList {
		entries[sw.SerialNumber] = sw
	}

	tflog.Info(ctx, "Bootstrap list queried", map[string]interface{}{
		"count": len(entries),
	})
	return entries, nil
}

// buildBootstrapRequestFromAPI builds a bootstrap request by merging user config with bootstrap API data.
func (r *inventorySwitchResource) buildBootstrapRequestFromAPI(data *NDFCInventorySwitchModel, bootstrapEntries map[string]NDFCSwitchDetailValue) BootstrapBuildResult {
	result := BootstrapBuildResult{}
	serial := data.SwitchDetail.SerialNumber

	entry, ok := bootstrapEntries[serial]
	if !ok {
		result.MissingFromBootstrap = append(result.MissingFromBootstrap, serial)
		return result
	}

	discoveryAuth := data.SwitchDetail.DiscoveryAuthProtocol
	if discoveryAuth == "" {
		discoveryAuth = entry.DiscoveryAuthProtocol
	}
	if discoveryAuth == "" {
		discoveryAuth = "md5"
	}

	softwareImage := data.SwitchDetail.SoftwareImage
	if softwareImage == "" {
		softwareImage = entry.SoftwareImage
	}

	swEntry := BootstrapSwitchEntry{
		NDFCSwitchDetailValue: NDFCSwitchDetailValue{
			SerialNumber:          serial,
			Hostname:              data.SwitchDetail.Hostname,
			IpAddress:             data.SwitchDetail.IpAddress,
			Model:                 entry.Model,
			SoftwareVersion:       entry.SoftwareVersion,
			SoftwareImage:         softwareImage,
			GatewayIpMask:         data.SwitchDetail.GatewayIpMask,
			SwitchPassword:        data.BootstrapPassword,
			DiscoveryAuthProtocol: discoveryAuth,
			PublicKey:             entry.PublicKey,
			Fingerprint:           entry.Fingerprint,
		},
		UseNewCredentials: &data.UseNewCredentials,
	}
	if data.UseNewCredentials {
		swEntry.DiscoveryUsername = data.DiscoveryUsername
		swEntry.DiscoveryPassword = data.DiscoveryPassword
	}
	result.Request.Switches = append(result.Request.Switches, swEntry)
	return result
}
