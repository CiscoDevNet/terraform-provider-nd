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
	"math/rand"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/looplab/fsm"
)

// FSM state names
const (
	StateInit               = "init"
	StateDiscovering        = "discovering"
	StateAddingSwitches     = "adding_switches"
	StateQueryingBootstrap  = "querying_bootstrap"
	StateImportingBootstrap = "importing_bootstrap"
	StateCheckingReadiness  = "checking_readiness"
	StateWaitingReady       = "waiting_ready"
	StateRediscovering      = "rediscovering"
	StateSavingCreds        = "saving_creds"
	StateUpdatingRoles      = "updating_roles"
	StateDone               = "done"
	StateFailed             = "failed"
)

// FSM event names
const (
	EventDiscover        = "discover"
	EventAddSwitches     = "add_switches"
	EventQueryBootstrap  = "query_bootstrap"
	EventImportBootstrap = "import_bootstrap"
	EventCheck           = "check"
	EventWait            = "wait"
	EventPoll            = "poll"
	EventRediscover      = "rediscover_switches"
	EventSaveCreds       = "save_creds"
	EventUpdateRoles     = "update_roles"
	EventFinish          = "finish"
	EventFail            = "fail"
)

const (
	DefaultWaitForReady = 30 * time.Minute // Default wait_for_ready timeout
)

// jitteredInterval returns a duration randomly varied around base by ±jitterFraction.
// e.g. jitteredInterval(10s, 0.3) returns 7s–13s.
func jitteredInterval(base time.Duration, jitterFraction float64) time.Duration {
	jitter := float64(base) * jitterFraction
	min := float64(base) - jitter
	return time.Duration(min + rand.Float64()*2*jitter)
}

// parseTimeout parses a duration string (e.g. "30m", "60s") and returns the duration.
// Returns the fallback if the input is empty or unparseable.
func parseTimeout(s string, fallback time.Duration) time.Duration {
	if s == "" {
		return fallback
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return fallback
	}
	return d
}

// InventoryFSM manages the state machine for switch creation (discovery or bootstrap)
type InventoryFSM struct {
	FSM              *fsm.FSM
	ctx              context.Context
	r                *inventorySwitchResource
	invAPI           *api.InventoryAPI
	switchesData     *NDFCInventorySwitchModel
	discovery        DiscoveryStatusResponse
	bootstrapEntries map[string]NDFCSwitchDetailValue
	dg               *diag.Diagnostics
	deadline         time.Time
	timeoutDuration  time.Duration
	isCreate         bool
	lastErr          error
	serialSet        map[string]bool
}

// NewInventoryFSM creates a new FSM for switch creation.
//
// Discovery mode:
//
//	init → discovering → adding_switches → saving_creds → updating_roles
//	  → checking_readiness → [waiting_ready ↔ checking_readiness] → done
//
// Bootstrap mode:
//
//	init → querying_bootstrap → importing_bootstrap → saving_creds → updating_roles
//	  → checking_readiness → [waiting_ready ↔ checking_readiness] → done
//
// Any state can transition to "failed" via the "fail" event.
func NewInventoryFSM(ctx context.Context,
	r *inventorySwitchResource,
	isCreate bool,
	invAPI *api.InventoryAPI,
	switchesData *NDFCInventorySwitchModel,
	dg *diag.Diagnostics) *InventoryFSM {

	timeout := parseTimeout(switchesData.WaitForReady, DefaultWaitForReady)
	inv := &InventoryFSM{
		ctx:             ctx,
		r:               r,
		isCreate:        isCreate,
		invAPI:          invAPI,
		switchesData:    switchesData,
		dg:              dg,
		deadline:        time.Now().Add(timeout),
		timeoutDuration: timeout,
	}

	allStates := []string{
		StateInit, StateDiscovering, StateAddingSwitches,
		StateQueryingBootstrap, StateImportingBootstrap,
		StateCheckingReadiness, StateRediscovering, StateWaitingReady,
		StateSavingCreds, StateUpdatingRoles,
	}

	inv.FSM = fsm.NewFSM(
		StateInit,
		fsm.Events{
			// Discovery path
			{Name: EventDiscover, Src: []string{StateInit}, Dst: StateDiscovering},
			{Name: EventAddSwitches, Src: []string{StateDiscovering}, Dst: StateAddingSwitches},
			// Bootstrap path
			{Name: EventQueryBootstrap, Src: []string{StateInit}, Dst: StateQueryingBootstrap},
			{Name: EventImportBootstrap, Src: []string{StateQueryingBootstrap}, Dst: StateImportingBootstrap},
			// Common path: save creds + roles first, then wait for readiness
			{Name: EventSaveCreds, Src: []string{StateAddingSwitches, StateImportingBootstrap}, Dst: StateSavingCreds},
			{Name: EventUpdateRoles, Src: []string{StateSavingCreds}, Dst: StateUpdatingRoles},
			{Name: EventCheck, Src: []string{StateUpdatingRoles}, Dst: StateCheckingReadiness},
			{Name: EventRediscover, Src: []string{StateCheckingReadiness}, Dst: StateRediscovering},
			{Name: EventWait, Src: []string{StateCheckingReadiness, StateRediscovering}, Dst: StateWaitingReady},
			{Name: EventPoll, Src: []string{StateWaitingReady}, Dst: StateCheckingReadiness},
			{Name: EventFinish, Src: []string{StateCheckingReadiness}, Dst: StateDone},
			{Name: EventFail, Src: allStates, Dst: StateFailed},
		},
		fsm.Callbacks{
			"enter_discovering":         func(ctx context.Context, e *fsm.Event) { inv.onDiscover(ctx, e) },
			"enter_adding_switches":     func(ctx context.Context, e *fsm.Event) { inv.onAddSwitches(ctx, e) },
			"enter_querying_bootstrap":  func(ctx context.Context, e *fsm.Event) { inv.onQueryBootstrap(ctx, e) },
			"enter_importing_bootstrap": func(ctx context.Context, e *fsm.Event) { inv.onImportBootstrap(ctx, e) },
			"enter_checking_readiness":  func(ctx context.Context, e *fsm.Event) { inv.onCheckReadiness(ctx, e) },
			"enter_rediscovering":       func(ctx context.Context, e *fsm.Event) { inv.onRediscover(ctx, e) },
			"enter_waiting_ready":       func(ctx context.Context, e *fsm.Event) { inv.onWaitReady(ctx, e) },
			"enter_saving_creds":        func(ctx context.Context, e *fsm.Event) { inv.onSaveCreds(ctx, e) },
			"enter_updating_roles":      func(ctx context.Context, e *fsm.Event) { inv.onUpdateRoles(ctx, e) },
			"enter_failed":              func(ctx context.Context, e *fsm.Event) { inv.onFailed(ctx, e) },
		},
	)
	return inv
}

// Run kicks off the FSM by firing the initial event based on mode
func (inv *InventoryFSM) Run() {
	// Tag all logs from this FSM with switch identity
	inv.ctx = tflog.SetField(inv.ctx, "switch_ip", inv.switchesData.SwitchDetail.IpAddress)
	inv.ctx = tflog.SetField(inv.ctx, "switch_serial", inv.switchesData.SwitchDetail.SerialNumber)

	tflog.Info(inv.ctx, "Starting inventory FSM", map[string]interface{}{
		"mode": inv.switchesData.Mode,
	})

	var startEvent string
	if inv.switchesData.Mode == ModeBootstrap {
		startEvent = EventQueryBootstrap
	} else {
		startEvent = EventDiscover
	}

	if err := inv.FSM.Event(inv.ctx, startEvent); err != nil {
		if !inv.dg.HasError() {
			inv.dg.AddError("Error Creating Inventory", fmt.Sprintf("FSM error: %v", err))
		}
	}

	tflog.Info(inv.ctx, "Inventory FSM completed", map[string]interface{}{
		"final_state": inv.FSM.Current(),
	})
}

// triggerFail records an error and transitions to the failed state
func (inv *InventoryFSM) triggerFail(ctx context.Context, e *fsm.Event, format string, args ...interface{}) {
	inv.lastErr = fmt.Errorf(format, args...)
	if err := e.FSM.Event(ctx, EventFail); err != nil {
		inv.dg.AddError("Error Creating Inventory", inv.lastErr.Error())
	}
}

// checkDeadline returns false and triggers fail if the context is cancelled or the deadline has passed
func (inv *InventoryFSM) checkDeadline(ctx context.Context, e *fsm.Event) bool {
	select {
	case <-ctx.Done():
		inv.triggerFail(ctx, e, "context cancelled in state %s: %v", e.Dst, ctx.Err())
		return false
	default:
	}
	if time.Now().After(inv.deadline) {
		inv.triggerFail(ctx, e, "timeout in state %s after %v", e.Dst, inv.timeoutDuration)
		return false
	}
	return true
}

// --- Discovery path callbacks ---

func (inv *InventoryFSM) onDiscover(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	discoverTimeout := parseTimeout(inv.switchesData.WaitForDiscover, 0)
	discoverDeadline := time.Now().Add(discoverTimeout)

	for {
		discovery, err := inv.r.shallowDiscover(ctx, inv.invAPI, inv.switchesData)
		if err != nil {
			inv.triggerFail(ctx, e, "%v", err)
			return
		}
		inv.discovery = *discovery

		// Check if any discovered switch is not reachable
		notReachable := false
		for _, sw := range inv.discovery.Switches {
			if sw.Status == StatusNotReachable {
				notReachable = true
				tflog.Info(ctx, "Switch not reachable", map[string]interface{}{
					"ip":     sw.IpAddress,
					"serial": sw.SerialNumber,
					"status": sw.Status,
				})
			}
		}

		if !notReachable {
			break
		}

		// No wait_for_discover configured — fail immediately
		if discoverTimeout == 0 {
			inv.triggerFail(ctx, e, "switch is not reachable during discovery")
			return
		}

		// Timeout expired
		if time.Now().After(discoverDeadline) {
			inv.triggerFail(ctx, e, "switch not reachable within wait_for_discover timeout (%v)", discoverTimeout)
			return
		}

		tflog.Info(ctx, "Waiting for switch to become reachable, retrying discovery...", map[string]interface{}{
			"remaining": time.Until(discoverDeadline).Round(time.Second).String(),
		})
		time.Sleep(jitteredInterval(PollInterval, 0.2))
	}

	// On create, the discovered switch must be available (not already managed)
	if inv.isCreate {
		for _, sw := range inv.discovery.Switches {
			if sw.Status == StatusAlreadyManaged {
				inv.triggerFail(ctx, e,
					"switch is already managed in this fabric: %s (%s)",
					sw.SerialNumber, sw.IpAddress)
				return
			}
		}
	}

	e.FSM.Event(ctx, EventAddSwitches)
}

func (inv *InventoryFSM) onAddSwitches(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	// In update flow, filter out switch if already managed
	discoveryForAdd := inv.discovery
	if !inv.isCreate {
		filtered := make([]NDFCSwitchDetailValue, 0, len(inv.discovery.Switches))
		for _, sw := range inv.discovery.Switches {
			if sw.Status == StatusAlreadyManaged {
				tflog.Info(ctx, "Skipping already-managed switch in update flow", map[string]interface{}{
					"serial": sw.SerialNumber,
					"ip":     sw.IpAddress,
				})
				continue
			}
			filtered = append(filtered, sw)
		}
		discoveryForAdd.Switches = filtered
	}

	if len(discoveryForAdd.Switches) == 0 {
		tflog.Info(ctx, "No switches to add")
		e.FSM.Event(ctx, EventSaveCreds)
		return
	}

	addReq := inv.r.buildAddSwitchesRequest(&discoveryForAdd, inv.switchesData)

	payload, err := json.Marshal(addReq)
	if err != nil {
		inv.triggerFail(ctx, e, "could not marshal add switches request: %v", err)
		return
	}

	inv.invAPI.SetOperation(api.OpAddSwitches)
	res, err := inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		inv.triggerFail(ctx, e, "add switches failed: %v: %s", err, res.String())
		return
	}

	tflog.Info(ctx, "Switch added to fabric", map[string]interface{}{
		"fabric_name": inv.invAPI.FabricName,
	})

	e.FSM.Event(ctx, EventSaveCreds)
}

// --- Bootstrap path callbacks ---

func (inv *InventoryFSM) onQueryBootstrap(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	serial := inv.switchesData.SwitchDetail.SerialNumber
	bootstrapTimeout := parseTimeout(inv.switchesData.WaitForBootstrap, 0)
	bootstrapDeadline := time.Now().Add(bootstrapTimeout)

	for {
		entries, err := inv.r.queryBootstrapList(ctx, inv.invAPI)
		if err != nil {
			inv.triggerFail(ctx, e, "%v", err)
			return
		}
		inv.bootstrapEntries = entries

		if _, ok := entries[serial]; ok {
			e.FSM.Event(ctx, EventImportBootstrap)
			return
		}

		// No wait_for_bootstrap configured — fail immediately
		if bootstrapTimeout == 0 {
			inv.triggerFail(ctx, e, "switch serial %s not found in bootstrap list", serial)
			return
		}

		// Wait timeout expired
		if time.Now().After(bootstrapDeadline) {
			inv.triggerFail(ctx, e, "switch serial %s not found in bootstrap list after %v", serial, bootstrapTimeout)
			return
		}

		tflog.Debug(ctx, "Switch not yet in bootstrap list, waiting", map[string]interface{}{
			"serial":            serial,
			"bootstrap_timeout": bootstrapTimeout.String(),
		})

		select {
		case <-time.After(jitteredInterval(PollInterval, 0.3)):
		case <-ctx.Done():
			inv.triggerFail(ctx, e, "context cancelled waiting for bootstrap: %v", ctx.Err())
			return
		}
	}
}

func (inv *InventoryFSM) onImportBootstrap(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	serial := inv.switchesData.SwitchDetail.SerialNumber

	// Pre-check: skip import if the switch already exists in the fabric
	fabricSwitches, err := inv.r.getAllSwitchesByFabric(ctx, inv.invAPI.FabricName, true)
	if err != nil {
		inv.triggerFail(ctx, e, "could not check existing switches: %v", err)
		return
	}
	if _, exists := fabricSwitches.SwitchesBySerial[serial]; exists {
		tflog.Info(ctx, "Switch already exists in fabric, skipping bootstrap import", map[string]interface{}{
			"serial": serial,
		})
		inv.discovery = DiscoveryStatusResponse{
			Switches: []NDFCSwitchDetailValue{
				{SerialNumber: serial, IpAddress: inv.switchesData.SwitchDetail.IpAddress},
			},
		}
		e.FSM.Event(ctx, EventSaveCreds)
		return
	}

	buildResult := inv.r.buildBootstrapRequestFromAPI(inv.switchesData, inv.bootstrapEntries)
	if len(buildResult.MissingFromBootstrap) > 0 {
		inv.triggerFail(ctx, e, "switch not found in bootstrap list: %s",
			strings.Join(buildResult.MissingFromBootstrap, ", "))
		return
	}

	payload, err := json.Marshal(buildResult.Request)
	if err != nil {
		inv.triggerFail(ctx, e, "could not marshal bootstrap request: %v", err)
		return
	}

	inv.invAPI.SetOperation(api.OpImportBootstrap)
	res, err := inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		inv.triggerFail(ctx, e, "bootstrap import failed: %v: %s", err, res.String())
		return
	}

	// Populate discovery with bootstrap serials for readiness tracking
	inv.discovery = DiscoveryStatusResponse{
		Switches: []NDFCSwitchDetailValue{
			{SerialNumber: serial, IpAddress: inv.switchesData.SwitchDetail.IpAddress},
		},
	}

	tflog.Info(ctx, "Bootstrap import initiated", map[string]interface{}{
		"fabric_name": inv.invAPI.FabricName,
		"serial":      serial,
	})

	e.FSM.Event(ctx, EventSaveCreds)
}

// --- Common path callbacks ---

// onCheckReadiness performs a single readiness check.
// If ready, transitions to done. Otherwise, transitions to waiting_ready.
func (inv *InventoryFSM) onCheckReadiness(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	// Build serial set once on first entry
	if inv.serialSet == nil {
		serials := inv.r.getSerialNumbers(&inv.discovery)
		inv.serialSet = make(map[string]bool, len(serials))
		for _, s := range serials {
			inv.serialSet[s] = true
		}
	}

	result, err := inv.r.switchesReady(ctx, inv.switchesData.FabricName, inv.serialSet)
	if err != nil {
		tflog.Debug(ctx, "GET switches failed, will retry", map[string]interface{}{"error": err.Error()})
		e.FSM.Event(ctx, EventWait)
		return
	}

	if result.Ready {
		tflog.Info(ctx, "Switch ready", map[string]interface{}{
			"count": len(inv.serialSet),
		})
		e.FSM.Event(ctx, EventFinish)
		return
	}

	// Switches that need rediscovery
	if len(result.NeedRediscover) > 0 {
		e.FSM.Event(ctx, EventRediscover, result.NeedRediscover)
		return
	}

	e.FSM.Event(ctx, EventWait)
}

// onRediscover triggers rediscovery for switches that need it, then returns to waiting.
func (inv *InventoryFSM) onRediscover(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	if len(e.Args) > 0 {
		if serials, ok := e.Args[0].([]string); ok && len(serials) > 0 {
			inv.r.triggerRediscovery(ctx, inv.invAPI, serials)
		}
	}

	e.FSM.Event(ctx, EventWait)
}

// onWaitReady sleeps for PollInterval then fires a poll event to re-check readiness.
func (inv *InventoryFSM) onWaitReady(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	select {
	case <-time.After(jitteredInterval(PollInterval, 0.3)):
	case <-ctx.Done():
		inv.triggerFail(ctx, e, "context cancelled during wait: %v", ctx.Err())
		return
	}
	e.FSM.Event(ctx, EventPoll)
}

func (inv *InventoryFSM) onSaveCreds(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	if inv.switchesData.DiscoveryUsername == "" || inv.switchesData.DiscoveryPassword == "" {
		e.FSM.Event(ctx, EventUpdateRoles)
		return
	}

	credReq := SwitchCredentialsRequest{
		SwitchIds:      inv.r.getSerialNumbers(&inv.discovery),
		SwitchUsername: inv.switchesData.DiscoveryUsername,
		SwitchPassword: inv.switchesData.DiscoveryPassword,
	}

	payload, err := json.Marshal(credReq)
	if err != nil {
		inv.triggerFail(ctx, e, "could not marshal credentials request: %v", err)
		return
	}

	inv.invAPI.SetOperation(api.OpCreateCredentials)
	res, err := inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		inv.triggerFail(ctx, e, "save credentials failed: %v: %s", err, res.String())
		return
	}

	tflog.Info(ctx, "Credentials saved for switch")
	e.FSM.Event(ctx, EventUpdateRoles)
}

func (inv *InventoryFSM) onUpdateRoles(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	desiredRole := inv.switchesData.SwitchDetail.SwitchRole
	desiredIP := inv.switchesData.SwitchDetail.IpAddress

	roleReq := SwitchRoleUpdateRequest{
		SwitchRoles: []SwitchRole{},
	}

	for _, sw := range inv.discovery.Switches {
		if sw.IpAddress != desiredIP || desiredRole == "" {
			continue
		}
		roleReq.SwitchRoles = append(roleReq.SwitchRoles, SwitchRole{
			SwitchId: sw.SerialNumber,
			Role:     desiredRole,
		})
	}

	if len(roleReq.SwitchRoles) > 0 {
		payload, err := json.Marshal(roleReq)
		if err != nil {
			inv.triggerFail(ctx, e, "could not marshal role update: %v", err)
			return
		}

		inv.invAPI.FabricName = inv.switchesData.FabricName
		inv.invAPI.SetOperation(api.OpUpdateSwitchRole)
		resp, err := inv.invAPI.Post(payload, nil)
		if err != nil {
			inv.triggerFail(ctx, e, "could not update roles: %v: %s", err, resp.String())
			return
		}

		tflog.Info(ctx, "Updated switch role", map[string]interface{}{
			"fabric_name": inv.switchesData.FabricName,
			"role":        desiredRole,
		})
	}

	e.FSM.Event(ctx, EventCheck)
}

func (inv *InventoryFSM) onFailed(_ context.Context, _ *fsm.Event) {
	if inv.lastErr != nil {
		inv.dg.AddError("Error Creating Inventory", inv.lastErr.Error())
	}
}
