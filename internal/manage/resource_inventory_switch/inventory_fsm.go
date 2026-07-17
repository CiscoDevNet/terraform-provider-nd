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
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/looplab/fsm"
)

// FSM state names
const (
	StateInit               = "init"
	StateDiscovering        = "discovering"
	StateAddingSwitches     = "adding_switches"
	StateCheckingReadiness  = "checking_readiness"
	StateWaitingReady       = "waiting_ready"
	StateSavingCreds        = "saving_creds"
	StateUpdatingRoles      = "updating_roles"
	StateSavingConfig       = "saving_config"
	StateConfigSaveWait     = "config_save_wait"
	StateRediscovering      = "rediscovering"
	StateDeployingConfig    = "deploying_config"
	StateQueryingBootstrap  = "querying_bootstrap"
	StateImportingBootstrap = "importing_bootstrap"
	StateDone               = "done"
	StateFailed             = "failed"
)

// FSM event names
const (
	EventDiscover        = "discover"
	EventAddSwitches     = "add_switches"
	EventCheck           = "check"
	EventWait            = "wait"
	EventPoll            = "poll"
	EventSaveCreds       = "save_creds"
	EventUpdateRoles     = "update_roles"
	EventConfigSave      = "config_save"
	EventRediscover      = "rediscover_switches"
	EventRetryWait       = "retry_wait"
	EventConfigDeploy    = "config_deploy"
	EventFinish          = "finish"
	EventFail            = "fail"
	EventQueryBootstrap  = "query_bootstrap"
	EventImportBootstrap = "import_bootstrap"
)

// Retryable error patterns from config-save
var configSaveRetryablePatterns = []string{
	"import is not yet completed",
	"migration mode",
	"ensure switches are reachable",
}

const (
	BaseFSMTimeout       = 10 * time.Minute // Base timeout for FSM operations
	PerSwitchTimeout     = 5 * time.Minute  // Additional timeout per switch
	MaxConfigSaveRetries = 50
	ConfigSaveRetryWait  = 30 * time.Second
)

// calculateFSMTimeout returns the total timeout based on number of switches.
// Formula: BaseFSMTimeout + (PerSwitchTimeout * numSwitches)
// Examples:
//   - 1 switch:  10 + 5  = 15 minutes
//   - 2 switches: 10 + 10 = 20 minutes
//   - 3 switches: 10 + 15 = 25 minutes
//   - 5 switches: 10 + 25 = 35 minutes
func calculateFSMTimeout(numSwitches int) time.Duration {
	if numSwitches < 1 {
		numSwitches = 1
	}
	return BaseFSMTimeout + (PerSwitchTimeout * time.Duration(numSwitches))
}

// InventoryFSM manages the state machine for switch creation (discovery and bootstrap modes)
type InventoryFSM struct {
	FSM               *fsm.FSM
	ctx               context.Context
	r                 *inventorySwitchResource
	invAPI            *api.InventoryAPI
	switchesData      *NDFCInventorySwitchModel
	updateSwitches    *NDFCInventorySwitchModel
	discovery         DiscoveryStatusResponse
	bootstrapEntries  map[string]BootstrapListEntry // populated by onQueryBootstrap
	dg                *diag.Diagnostics
	deadline          time.Time
	timeoutDuration   time.Duration // Total timeout for error reporting
	configSaveRetries int
	isCreate          bool
	lastErr           error
	serialSet         map[string]bool
}

// NewInventoryFSM creates a new FSM for switch creation (discovery and bootstrap modes).
//
// The FSM drives the full lifecycle of adding switches to a fabric. A single
// call to Run() fires the initial event based on mode; every subsequent
// transition is chained from enter_<state> callbacks.
//
// Discovery mode:
//
//	+--------+  discover   +--------------+  add_switches  +------------------+
//	|  init  |----------->| discovering  |--------------->| adding_switches  |---+
//	+--------+            +--------------+                +------------------+   |
//	                                                                          check
//	                                                                            |
//	                                                                            v
//	                                                                +---------------------+
//	                                                                | checking_readiness  |---> ...
//	                                                                +---------------------+
//
// Bootstrap mode:
//
//	+--------+  query_bootstrap  +----------------------+  import_bootstrap  +------------------------+
//	|  init  |------------------>| querying_bootstrap   |-------------------->| importing_bootstrap   |---+
//	+--------+                   +----------------------+                    +------------------------+   |
//	                                                                                                   check
//	                                                                                                     |
//	                                                                                                     v
//	                                                                                         +---------------------+
//	                                                                                         | checking_readiness  |---> ...
//	                                                                                         +---------------------+
//
// Common tail (both modes):
//
//	                      +----------------+  poll   +---------------------+
//	                      | waiting_ready  |-------->| checking_readiness  |
//	                      +-------+--------+         +-+-----+-------+----+
//	                              ^                    |     |       |
//	                              |               wait | rediscover  | save_creds
//	                              +--------------------+     |       |
//	                              |                          v       |
//	                              |  wait   +----------------+       |
//	                              +---------| rediscovering  |       |
//	                                        +----------------+       |
//	                                                                 v
//	          saving_creds -> updating_roles -> saving_config -> deploying_config -> done
//
//	Any non-terminal state can transition to "failed" via the "fail" event.
func NewInventoryFSM(
	ctx context.Context,
	r *inventorySwitchResource,
	isCreate bool,
	invAPI *api.InventoryAPI,
	switchesData *NDFCInventorySwitchModel,
	dg *diag.Diagnostics,
) *InventoryFSM {
	timeout := calculateFSMTimeout(len(switchesData.Switches))
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
		StateCheckingReadiness, StateRediscovering, StateWaitingReady, StateSavingCreds,
		StateUpdatingRoles, StateSavingConfig, StateConfigSaveWait,
		StateDeployingConfig,
		StateQueryingBootstrap, StateImportingBootstrap,
	}

	inv.FSM = fsm.NewFSM(
		StateInit,
		fsm.Events{
			// Discovery flow
			{Name: EventDiscover, Src: []string{StateInit}, Dst: StateDiscovering},
			{Name: EventAddSwitches, Src: []string{StateDiscovering}, Dst: StateAddingSwitches},
			{Name: EventCheck, Src: []string{StateAddingSwitches, StateImportingBootstrap}, Dst: StateCheckingReadiness},
			// Bootstrap flow
			{Name: EventQueryBootstrap, Src: []string{StateInit}, Dst: StateQueryingBootstrap},
			{Name: EventImportBootstrap, Src: []string{StateQueryingBootstrap}, Dst: StateImportingBootstrap},
			{Name: EventRetryWait, Src: []string{StateSavingConfig, StateConfigSaveWait, StateRediscovering}, Dst: StateConfigSaveWait},
			{Name: EventRediscover, Src: []string{StateCheckingReadiness, StateConfigSaveWait}, Dst: StateRediscovering},
			{Name: EventWait, Src: []string{StateCheckingReadiness, StateRediscovering}, Dst: StateWaitingReady},
			{Name: EventPoll, Src: []string{StateWaitingReady}, Dst: StateCheckingReadiness},
			{Name: EventSaveCreds, Src: []string{StateCheckingReadiness}, Dst: StateSavingCreds},
			{Name: EventUpdateRoles, Src: []string{StateSavingCreds}, Dst: StateUpdatingRoles},
			{Name: EventConfigSave, Src: []string{StateUpdatingRoles, StateConfigSaveWait}, Dst: StateSavingConfig},
			{Name: EventConfigDeploy, Src: []string{StateSavingConfig}, Dst: StateDeployingConfig},
			{Name: EventFinish, Src: []string{StateDeployingConfig, StateSavingConfig}, Dst: StateDone},
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
			"enter_saving_config":       func(ctx context.Context, e *fsm.Event) { inv.onConfigSave(ctx, e) },
			"enter_config_save_wait":    func(ctx context.Context, e *fsm.Event) { inv.onConfigSaveWait(ctx, e) },
			"enter_deploying_config":    func(ctx context.Context, e *fsm.Event) { inv.onConfigDeploy(ctx, e) },
			"enter_failed":              func(ctx context.Context, e *fsm.Event) { inv.onFailed(ctx, e) },
		},
	)
	return inv
}

// Run kicks off the FSM by firing the initial event based on mode.
// Discovery mode fires EventDiscover; bootstrap mode fires EventQueryBootstrap.
func (inv *InventoryFSM) Run() {
	tflog.Info(inv.ctx, "Starting inventory FSM", map[string]interface{}{
		"mode": inv.switchesData.Mode,
	})

	var initialEvent string
	switch inv.switchesData.Mode {
	case ModeBootstrap:
		initialEvent = EventQueryBootstrap
	default:
		initialEvent = EventDiscover
	}

	if err := inv.FSM.Event(inv.ctx, initialEvent); err != nil {
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

// --- State callbacks ---

func (inv *InventoryFSM) onDiscover(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	discovery, err := inv.r.shallowDiscover(ctx, inv.invAPI, inv.switchesData)
	if err != nil {
		inv.triggerFail(ctx, e, "%v", err)
		return
	}
	inv.discovery = *discovery

	// On create, all discovered switches must be available (not already managed)
	if inv.isCreate {
		var alreadyManaged []string
		for _, sw := range inv.discovery.Switches {
			if sw.Status == StatusAlreadyManaged {
				alreadyManaged = append(alreadyManaged, fmt.Sprintf("%s (%s)", sw.SerialNumber, sw.IpAddress))
			}
		}
		if len(alreadyManaged) > 0 {
			inv.triggerFail(ctx, e,
				"not all switches are available to be managed in this fabric; "+
					"already managed: %s", strings.Join(alreadyManaged, ", "))
			return
		}
	}

	e.FSM.Event(ctx, EventAddSwitches)
}

func (inv *InventoryFSM) onAddSwitches(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	// In update flow, filter out switches that are already managed
	discoveryForAdd := inv.discovery
	if !inv.isCreate {
		filtered := make([]NDFCSwitchesValue, 0, len(inv.discovery.Switches))
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
		e.FSM.Event(ctx, EventCheck)
		return
	}

	addReq := inv.r.buildAddSwitchesRequest(&discoveryForAdd, inv.switchesData)

	payload, err := json.Marshal(addReq)
	if err != nil {
		inv.triggerFail(ctx, e, "could not marshal add switches request: %v", err)
		return
	}

	inv.invAPI.SetOperation(api.OpAddSwitches)
	_, err = inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true}) // disable payload logging (contains credentials)
	if err != nil {
		inv.triggerFail(ctx, e, "add switches failed: %v", err)
		return
	}

	tflog.Info(ctx, "Switches added to fabric", map[string]interface{}{
		"fabric_name": inv.invAPI.FabricName,
		"count":       len(addReq.Switches),
	})

	e.FSM.Event(ctx, EventCheck)
}

// --- Bootstrap mode callbacks ---

// onQueryBootstrap queries the bootstrap API to find switches in the POAP loop.
func (inv *InventoryFSM) onQueryBootstrap(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	entries, err := inv.r.queryBootstrapList(ctx, inv.invAPI)
	if err != nil {
		inv.triggerFail(ctx, e, "bootstrap list query failed: %v", err)
		return
	}

	inv.bootstrapEntries = entries
	e.FSM.Event(ctx, EventImportBootstrap)
}

// onImportBootstrap builds the importBootstrap payload by merging user config
// with bootstrap API data, posts it, and transitions to checking_readiness.
//
// Idempotency: if some serials are missing from the POAP bootstrap list,
// they may already be in the fabric inventory (from a previous partial run).
// Those serials are added to serialSet so downstream steps still process them,
// but only the POAP-listed switches are sent to importBootstrap.
func (inv *InventoryFSM) onImportBootstrap(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	buildResult := inv.r.buildBootstrapRequestFromAPI(inv.switchesData, inv.bootstrapEntries)

	// Check if missing serials are already in the fabric inventory
	if len(buildResult.MissingFromBootstrap) > 0 {
		resp, err := inv.r.getAllSwitchesByFabric(ctx, inv.switchesData.FabricName, true)
		if err != nil {
			inv.triggerFail(ctx, e, "could not check fabric inventory for missing bootstrap switches: %v", err)
			return
		}

		var genuinelyMissing []string
		var alreadyInFabric []string
		for _, serial := range buildResult.MissingFromBootstrap {
			if _, found := resp.SwitchesBySerial[serial]; found {
				alreadyInFabric = append(alreadyInFabric, serial)
			} else {
				genuinelyMissing = append(genuinelyMissing, serial)
			}
		}

		if len(genuinelyMissing) > 0 {
			inv.triggerFail(ctx, e, "switches not found in POAP bootstrap list or fabric inventory: %v", genuinelyMissing)
			return
		}

		if len(alreadyInFabric) > 0 {
			tflog.Info(ctx, "Switches already in fabric inventory, skipping import for these", map[string]interface{}{
				"serials": alreadyInFabric,
			})
		}
	}

	// Build serial set for readiness checking — includes ALL config serials
	inv.serialSet = make(map[string]bool, len(inv.switchesData.Switches))
	for serial := range inv.switchesData.Switches {
		inv.serialSet[serial] = true
	}

	// Only POST importBootstrap if there are switches to import
	if len(buildResult.Request.Switches) > 0 {
		payload, err := json.Marshal(buildResult.Request)
		if err != nil {
			inv.triggerFail(ctx, e, "could not marshal bootstrap request: %v", err)
			return
		}

		inv.invAPI.SetOperation(api.OpImportBootstrap)
		res, err := inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
		if err != nil {
			inv.triggerFail(ctx, e, "importBootstrap failed: %v: %s", err, res.String())
			return
		}

		tflog.Info(ctx, "Bootstrap import initiated", map[string]interface{}{
			"fabric_name": inv.invAPI.FabricName,
			"count":       len(buildResult.Request.Switches),
		})
	} else {
		tflog.Info(ctx, "All switches already in fabric, skipping importBootstrap")
	}

	e.FSM.Event(ctx, EventCheck)
}

// onCheckReadiness performs a single readiness check for all target switches.
// If all ready, transitions to save_creds. Otherwise, transitions to waiting_ready.
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
		tflog.Info(ctx, "All switches ready", map[string]interface{}{
			"count": len(inv.serialSet),
		})
		e.FSM.Event(ctx, EventSaveCreds)
		return
	}

	// Switches that need rediscovery — transition to rediscovering state
	if len(result.NeedRediscover) > 0 {
		e.FSM.Event(ctx, EventRediscover, result.NeedRediscover)
		return
	}

	e.FSM.Event(ctx, EventWait)
}

// onRediscover triggers rediscovery for switches that need it.
// Returns to waiting_ready (normal flow) or config_save_wait (config save retry flow) based on source state.
func (inv *InventoryFSM) onRediscover(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	// Extract serials passed via event args
	if len(e.Args) > 0 {
		if serials, ok := e.Args[0].([]string); ok && len(serials) > 0 {
			inv.r.triggerRediscovery(ctx, inv.invAPI, serials)
		}
	}

	// Return to the appropriate wait state based on where we came from
	if e.Src == StateConfigSaveWait {
		e.FSM.Event(ctx, EventRetryWait)
	} else {
		e.FSM.Event(ctx, EventWait)
	}
}

// onWaitReady sleeps for PollInterval then fires a poll event to re-check readiness.
func (inv *InventoryFSM) onWaitReady(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	select {
	case <-time.After(PollInterval):
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

	if inv.switchesData.Username == "" || inv.switchesData.Password == "" {
		e.FSM.Event(ctx, EventUpdateRoles)
		return
	}

	// Collect serial numbers from serialSet (works for both discovery and bootstrap modes)
	serials := make([]string, 0, len(inv.serialSet))
	for s := range inv.serialSet {
		serials = append(serials, s)
	}

	credReq := SwitchCredentialsRequest{
		SwitchIds:      serials,
		SwitchUsername: inv.switchesData.Username,
		SwitchPassword: inv.switchesData.Password,
	}

	payload, err := json.Marshal(credReq)
	if err != nil {
		inv.triggerFail(ctx, e, "could not marshal credentials request: %v", err)
		return
	}

	inv.invAPI.SetOperation(api.OpCreateCredentials)
	_, err = inv.invAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true}) // disable payload logging (contains credentials)
	if err != nil {
		inv.triggerFail(ctx, e, "save credentials failed: %v", err)
		return
	}

	tflog.Info(ctx, "Credentials saved for switches")
	e.FSM.Event(ctx, EventUpdateRoles)
}

func (inv *InventoryFSM) onUpdateRoles(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	ipToRole := make(map[string]string)
	for _, sw := range inv.switchesData.Switches {
		if sw.IpAddress != "" && sw.SwitchRole != "" {
			ipToRole[sw.IpAddress] = sw.SwitchRole
		}
	}

	roleReq := SwitchRoleUpdateRequest{
		SwitchRoles: []SwitchRole{},
	}

	for _, sw := range inv.discovery.Switches {
		role, found := ipToRole[sw.IpAddress]
		if !found || role == "" {
			continue
		}
		roleReq.SwitchRoles = append(roleReq.SwitchRoles, SwitchRole{
			SwitchId: sw.SerialNumber,
			Role:     role,
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

		tflog.Info(ctx, "Updated switch roles", map[string]interface{}{
			"fabric_name": inv.switchesData.FabricName,
			"count":       len(roleReq.SwitchRoles),
		})
	}

	e.FSM.Event(ctx, EventConfigSave)
}

func (inv *InventoryFSM) onConfigSave(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	if !inv.switchesData.Recalculate {
		tflog.Info(ctx, "onConfigSave: recalculate setting is off, not doing config save, move to deployment", map[string]interface{}{
			"fabric_name": inv.switchesData.FabricName,
		})
		e.FSM.Event(ctx, EventConfigDeploy)
		return
	}
	errDiag := diag.Diagnostics{}
	deployment.ConfigSaveAndDeploy(ctx, inv.r.manageClient.ApiClient, inv.switchesData.FabricName, true, false, &errDiag) //deployment.ConfigSave(ctx, inv.r.manageClient.ApiClient, inv.switchesData.FabricName, nil)
	if errDiag.HasError() {
		// Check for retryable errors (import not completed, migration mode)
		if inv.isRetryableConfigSaveError(errDiag[0].Summary()) || inv.isRetryableConfigSaveError(errDiag[0].Detail()) {
			inv.configSaveRetries++
			if inv.configSaveRetries > MaxConfigSaveRetries {
				inv.triggerFail(ctx, e, "config save failed after %d retries: %s", MaxConfigSaveRetries, errDiag[0].Detail())
				return
			}

			tflog.Info(ctx, "Config save retryable error, re-checking switches", map[string]interface{}{
				"retry":   inv.configSaveRetries,
				"message": errDiag[0].Detail(),
			})
			e.FSM.Event(ctx, EventRetryWait)
			return
		}

		inv.triggerFail(ctx, e, "config save failed: %s: %s", errDiag[0].Summary(), errDiag[0].Detail())
		return
	}

	e.FSM.Event(ctx, EventConfigDeploy)
}

// onConfigSaveWait sleeps for ConfigSaveRetryWait then retries config-save directly.
func (inv *InventoryFSM) onConfigSaveWait(ctx context.Context, e *fsm.Event) {
	tflog.Info(ctx, "Waiting for config save ready", map[string]interface{}{
		"fabric_name": inv.switchesData.FabricName,
	})

	if !inv.checkDeadline(ctx, e) {
		return
	}
	select {
	case <-time.After(ConfigSaveRetryWait):
	case <-ctx.Done():
		inv.triggerFail(ctx, e, "context cancelled during config save wait: %v", ctx.Err())
		return
	}

	// Check if all switches are still ready before retrying config save
	result, err := inv.r.switchesReady(ctx, inv.switchesData.FabricName, inv.serialSet)
	if err != nil {
		tflog.Debug(ctx, "GET switches failed during config save wait, will retry", map[string]interface{}{"error": err.Error()})
		e.FSM.Event(ctx, EventRetryWait)
		return
	}

	if len(result.NeedRediscover) > 0 {
		tflog.Info(ctx, "Switches need rediscovery during config save wait", map[string]interface{}{
			"serials": result.NeedRediscover,
		})
		e.FSM.Event(ctx, EventRediscover, result.NeedRediscover)
		return
	}

	if !result.Ready {
		tflog.Debug(ctx, "Switches not ready yet, will retry config save wait", map[string]interface{}{
			"found":    result.Found,
			"expected": result.Expected,
		})
		e.FSM.Event(ctx, EventRetryWait)
		return
	}

	e.FSM.Event(ctx, EventConfigSave)
}

func (inv *InventoryFSM) onConfigDeploy(ctx context.Context, e *fsm.Event) {
	if !inv.checkDeadline(ctx, e) {
		return
	}

	if !inv.switchesData.Deploy {
		tflog.Info(ctx, "onConfigDeploy: deploy setting is off, not doing config deploy, move to finish", map[string]interface{}{
			"fabric_name": inv.switchesData.FabricName,
		})
		e.FSM.Event(ctx, EventFinish)
		return
	}

	errDiag := diag.Diagnostics{}
	deployment.ConfigSaveAndDeploy(ctx, inv.r.manageClient.ApiClient, inv.switchesData.FabricName, false, true, &errDiag)
	if errDiag.HasError() {
		inv.triggerFail(ctx, e, "config deploy failed: %s: %s", errDiag[0].Summary(), errDiag[0].Detail())
		return
	}

	tflog.Info(ctx, "Config deployed", map[string]interface{}{
		"fabric_name": inv.switchesData.FabricName,
	})

	e.FSM.Event(ctx, EventFinish)
}

func (inv *InventoryFSM) onFailed(_ context.Context, e *fsm.Event) {
	if inv.lastErr != nil {
		inv.dg.AddError("Error Creating Inventory", inv.lastErr.Error())
	}

	// Bootstrap cleanup: if switches were already imported into the fabric,
	// remove them so they return to the POAP loop for the next run.
	if inv.switchesData.Mode == ModeBootstrap && inv.serialSet != nil && len(inv.serialSet) > 0 {
		// Only clean up if we got past the import step
		postImportStates := map[string]bool{
			StateCheckingReadiness: true,
			StateWaitingReady:      true,
			StateRediscovering:     true,
			StateSavingCreds:       true,
			StateUpdatingRoles:     true,
			StateSavingConfig:      true,
			StateConfigSaveWait:    true,
			StateDeployingConfig:   true,
		}
		if postImportStates[e.Src] {
			serials := make([]string, 0, len(inv.serialSet))
			for s := range inv.serialSet {
				serials = append(serials, s)
			}
			tflog.Warn(inv.ctx, "Bootstrap failed after import, removing switches to restore POAP state", map[string]interface{}{
				"serials": serials,
			})
			cleanupDg := diag.Diagnostics{}
			inv.r.removeSwitches(inv.ctx, &cleanupDg, inv.invAPI, serials)
			if cleanupDg.HasError() {
				tflog.Error(inv.ctx, "Bootstrap cleanup failed: could not remove imported switches", map[string]interface{}{
					"error": cleanupDg[0].Detail(),
				})
				inv.dg.AddWarning("Bootstrap Cleanup Failed",
					fmt.Sprintf("Could not remove imported switches after failure. "+
						"Manually remove switches %v from the fabric before retrying.", serials))
			} else {
				inv.dg.AddWarning("Bootstrap Rolled Back",
					fmt.Sprintf("Switches %v were removed from the fabric after failure. "+
						"They should return to the POAP loop for retry.", serials))
			}
		}
	}
}

// isRetryableConfigSaveError checks if the error message indicates a retryable condition
func (inv *InventoryFSM) isRetryableConfigSaveError(msg string) bool {
	lower := strings.ToLower(msg)
	for _, pattern := range configSaveRetryablePatterns {
		if strings.Contains(lower, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}
