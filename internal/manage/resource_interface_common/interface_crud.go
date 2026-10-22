// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	commontypes "terraform-provider-nd/internal/common/types"
	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// truncateBody returns the first maxLen bytes of b (as a string) for
// inclusion in diagnostic/log messages, avoiding excessively large output.
func truncateBody(b []byte, maxLen int) string {
	if len(b) <= maxLen {
		return string(b)
	}
	return string(b[:maxLen]) + "...(truncated)"
}

// InterfaceModel is the contract between per-type Terraform models and the
// shared CRUD backend. See `interface_ethernet/resource_custom_codec.go` for
// the ethernet implementation.
type InterfaceModel interface {
	GetModelData() *NDFCInterfaceCommonModel
	SetModelData(*NDFCInterfaceCommonModel) diag.Diagnostics
	GetFabricName() string
	// GetInterfaceType returns the per-type discriminator (e.g. "ethernet",
	// "portChannel") used both for payload injection and handler dispatch.
	GetInterfaceType() string
}

// InterfacePreMarshal is implemented by models that need to adjust the
// common model before any POST/PUT marshalling (injecting interfaceType,
// resolving switch_id IPs, filling interfaceName from the map key, etc.).
type InterfacePreMarshal interface {
	PreMarshal(ctx context.Context, data *NDFCInterfaceCommonModel)
}

// InterfacePostUnmarshal is implemented by models that need to post-process
// a GET response before SetModelData (normalising operData, filtering drift).
type InterfacePostUnmarshal interface {
	PostUnmarshal(ctx context.Context, data *NDFCInterfaceCommonModel)
}

// RscCreateInterfaces creates all interfaces described by `model`, grouped
// by switchId, and optionally deploys them. Handler dispatch key = per-type
// GetInterfaceType().
func RscCreateInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	interfaceHandler(model.GetInterfaceType()).Create(ctx, client, dg, model)
}

// RscGetInterfaces refreshes the model from the live API. Returns false if
// *all* entries 404 (so the resource can be removed from state). A per-entry
// 404 is handled by removing the entry from the map but keeping the resource.
func RscGetInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) bool {
	return interfaceHandler(model.GetInterfaceType()).Read(ctx, client, dg, model)
}

// RscUpdateInterfaces diffs plan vs prior state by (switchId, interfaceName)
// and issues targeted create/update/delete operations. Both the prior state
// and the new plan are required so that interfaces removed from configuration
// can be detected and deleted from ND.
func RscUpdateInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, state InterfaceModel, plan InterfaceModel) {
	interfaceHandler(plan.GetInterfaceType()).Update(ctx, client, dg, state, plan)
}

// RscDeleteInterfaces removes all interfaces in the model via a single bulk
// action call.
func RscDeleteInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	interfaceHandler(model.GetInterfaceType()).Delete(ctx, client, dg, model)
}

// RscDeployInterfaces fires /interfaceActions/deploy for the full set of
// interfaces in the model. Safe to call multiple times; always acquires
// LockDeploy and uses DeployPost (per Development_Guide §Deploy).
func RscDeployInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	inData := model.GetModelData()
	if pm, ok := model.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	if inData == nil || len(inData.Interfaces) == 0 {
		return
	}
	fabricName := model.GetFabricName()

	// Build the action payload from map keys + switch_id per entry.
	payload := NDFCInterfaceActionPayload{Interfaces: make([]NDFCInterfaceAction, 0, len(inData.Interfaces))}
	for name, v := range inData.Interfaces {
		payload.Interfaces = append(payload.Interfaces, NDFCInterfaceAction{
			InterfaceName: name,
			SwitchId:      v.SwitchId,
		})
	}
	// Stable order for logs/payload diffs.
	sort.Slice(payload.Interfaces, func(i, j int) bool {
		if payload.Interfaces[i].SwitchId != payload.Interfaces[j].SwitchId {
			return payload.Interfaces[i].SwitchId < payload.Interfaces[j].SwitchId
		}
		return payload.Interfaces[i].InterfaceName < payload.Interfaces[j].InterfaceName
	})
	body, err := json.Marshal(payload)
	if err != nil {
		dg.AddError("Error Deploying Interfaces",
			fmt.Sprintf("Could not marshal deploy payload for fabric %q: %v", fabricName, err))
		return
	}

	// LockDeploy per Development_Guide §Deploy. Must use DeployPost.
	guard := ndapi.Acquire(fabricName, "", ndapi.LockDeploy)
	defer guard.Release()
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.Operation = api.OpInterfaceDeploy
	tflog.Debug(ctx, "Deploying interfaces", map[string]interface{}{
		"fabric_name": fabricName,
		"count":       len(payload.Interfaces),
	})
	res, err := ifAPI.DeployPost(body, nil)
	if err != nil {
		dg.AddError("Error Deploying Interfaces",
			fmt.Sprintf("deploy call failed for fabric %q: %v: %s", fabricName, err, res.String()))
		return
	}
	// Parse HTTP 207 response and surface per-entry errors.
	if err := bulkResponseErrs("(deploy)", res.Raw); err != nil {
		dg.AddError("Error Deploying Interfaces", err.Error())
	}
}

// waitForDeploySync polls /interfacesSummary until all named interfaces
// show configSyncStatus == inSync (or deployed), or a timeout expires.
// On timeout, a Warning diagnostic is added so the user sees it in output.
// Interfaces are grouped by switchId to handle multi-switch resources.
const (
	deploySyncPollInterval = 3 * time.Second
	deploySyncTimeout      = 120 * time.Second
)

func waitForDeploySync(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, fabricName string, interfaces map[string]NDFCInterfacesValue) {
	if len(interfaces) == 0 {
		return
	}
	// Group interface names by switchId.
	bySwitch := make(map[string]map[string]struct{})
	for name, v := range interfaces {
		if v.SwitchId == "" {
			continue
		}
		if _, ok := bySwitch[v.SwitchId]; !ok {
			bySwitch[v.SwitchId] = make(map[string]struct{})
		}
		bySwitch[v.SwitchId][name] = struct{}{}
	}

	for switchId, nameSet := range bySwitch {
		waitForSwitchSync(ctx, client, dg, fabricName, switchId, nameSet)
	}
}

func waitForSwitchSync(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, fabricName, switchId string, nameSet map[string]struct{}) {
	deadline := time.Now().Add(deploySyncTimeout)

	for time.Now().Before(deadline) {
		statusMap := fetchDeployStatus(ctx, client, fabricName, switchId)
		if statusMap == nil {
			// Summary call failed — wait and retry.
			time.Sleep(deploySyncPollInterval)
			continue
		}
		allInSync := true
		for name := range nameSet {
			status, ok := statusMap[name]
			if !ok || !isDeployedStatus(status) {
				allInSync = false
				break
			}
		}
		if allInSync {
			tflog.Debug(ctx, "all interfaces in sync after deploy", map[string]interface{}{
				"fabric_name": fabricName,
				"switch_id":   switchId,
				"count":       len(nameSet),
			})
			return
		}
		time.Sleep(deploySyncPollInterval)
	}
	dg.AddWarning("Deploy Sync Timeout",
		fmt.Sprintf("Timed out (%s) waiting for interfaces on switch %q in fabric %q to reach inSync. "+
			"Some interfaces may not yet be deployed. The next plan will detect any drift.",
			deploySyncTimeout.String(), switchId, fabricName))
}

// ---- switch ID resolution ------------------------------------------------

// resolveSwitchIDs resolves any IP addresses in switch_id fields (both
// resource-level and per-entry) to serial numbers via SwitchDB singleton,
// per AGENTS.md §SwitchDB rules. If a value is already a serial (not an
// IP), it is left unchanged. Called by the default CRUD implementations
// after GetModelData + PreMarshal, on the working copy.
func resolveSwitchIDs(ctx context.Context, dg *diag.Diagnostics, data *NDFCInterfaceCommonModel) {
	if data == nil {
		return
	}
	db := manage.GetSwitchDBInstance()
	fabricName := data.FabricName

	// Resolve the resource-level switch_id.
	if ip := net.ParseIP(strings.TrimSpace(data.SwitchId)); ip != nil {
		serial, ok := db.GetSerialByIP(ctx, fabricName, data.SwitchId)
		if !ok {
			dg.AddError("Switch IP resolution failed",
				"Could not resolve resource-level switch_id IP "+data.SwitchId+
					" to a serial number via SwitchDB for fabric "+fabricName)
			return
		}
		tflog.Debug(ctx, "resolved resource switch_id IP to serial",
			map[string]interface{}{"ip": data.SwitchId, "serial": serial})
		data.SwitchId = serial
	}

	// Resolve per-entry switch_id overrides.
	for name, entry := range data.Interfaces {
		sid := strings.TrimSpace(entry.SwitchId)
		if sid == "" {
			continue
		}
		if ip := net.ParseIP(sid); ip != nil {
			serial, ok := db.GetSerialByIP(ctx, fabricName, sid)
			if !ok {
				dg.AddError("Switch IP resolution failed",
					"Could not resolve switch_id IP "+sid+
						" for interface "+name+" via SwitchDB in fabric "+fabricName)
				return
			}
			tflog.Debug(ctx, "resolved per-entry switch_id IP to serial",
				map[string]interface{}{"interface": name, "ip": sid, "serial": serial})
			entry.SwitchId = serial
			data.Interfaces[name] = entry
		}
	}
}

// ---- default implementations ---------------------------------------------

// createDefault groups entries by switchId and POSTs one create payload per
// switch. Deploy=true triggers a follow-up deploy call for the full set.
func createDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	inData := model.GetModelData()
	if pm, ok := model.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	resolveSwitchIDs(ctx, dg, inData)
	if dg.HasError() {
		return
	}
	fabricName := inData.FabricName

	// Group by (switchId, policyType) → list of NDFCInterfacesValue.
	// The ND bulk interface create endpoint rejects mixed policy types in a
	// single request, so we must sub-group.
	type switchPolicy struct {
		switchId   string
		policyType string
	}
	groups := map[switchPolicy][]NDFCInterfacesValue{}
	for _, v := range inData.Interfaces {
		key := switchPolicy{switchId: v.SwitchId, policyType: v.ConfigData.NetworkOS.Policy.PolicyType}
		groups[key] = append(groups[key], v)
	}

	for key, entries := range groups {
		payload := NDFCInterfacesPayload{Interfaces: entries}
		body, err := json.Marshal(payload)
		if err != nil {
			dg.AddError("Error Creating Interfaces",
				fmt.Sprintf("marshal failed for switch %q policy %q: %v", key.switchId, key.policyType, err))
			return
		}
		ifAPI := api.NewInterfaceAPI(client, fabricName)
		ifAPI.FabricName = fabricName
		ifAPI.SwitchId = key.switchId
		ifAPI.Operation = api.OpInterfaceCreatePerSwitch
		tflog.Debug(ctx, "Creating interfaces", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id":   key.switchId,
			"policy_type": key.policyType,
			"count":       len(entries),
		})
		res, err := ifAPI.Post(body, nil)
		if err != nil {
			dg.AddError("Error Creating Interfaces",
				fmt.Sprintf("create failed on switch %q policy %q: %v: %s", key.switchId, key.policyType, err, res.String()))
			return
		}
		// Parse HTTP 207 GenericBulkResponse and raise per-entry errors.
		var br GenericBulkResponse
		if parseErr := json.Unmarshal([]byte(res.Raw), &br); parseErr == nil {
			for _, r := range br.Results {
				if r.Status != "" && !strings.EqualFold(r.Status, "success") {
					dg.AddError("Interface create reported non-success",
						fmt.Sprintf("switch=%q interface=%q status=%q message=%q",
							key.switchId, r.Name, r.Status, r.Message))
				}
			}
		}
	}

	// On partial failure, read back successfully-created entries into state
	// so a retry does not duplicate them. Return without deploying.
	if dg.HasError() {
		var readDg diag.Diagnostics
		RscGetInterfaces(ctx, client, &readDg, model) // best-effort; ignore read errors
		return
	}

	// Optional deploy after create.
	if inData.Deploy {
		RscDeployInterfaces(ctx, client, dg, model)
		if dg.HasError() {
			// Deploy failed — rollback by removing the successfully-created
			// interfaces so both state and infra are clean. The caller must
			// NOT write state when diagnostics contain errors.
			rollbackEntries := make([]NDFCInterfaceAction, 0, len(inData.Interfaces))
			for name, v := range inData.Interfaces {
				rollbackEntries = append(rollbackEntries, NDFCInterfaceAction{
					InterfaceName: name,
					SwitchId:      v.SwitchId,
				})
			}
			if rbErr := bulkRemove(ctx, client, fabricName, rollbackEntries); rbErr != nil {
				dg.AddError("Error Rolling Back Interfaces After Deploy Failure",
					fmt.Sprintf("rollback (bulk remove) failed for fabric %q: %v", fabricName, rbErr))
			}
			// Clear model interfaces to signal deploy-rollback to the caller.
			// The caller uses this to distinguish deploy-failure (infra clean,
			// don't write state) from partial-create-failure (write partial state).
			cleared := model.GetModelData()
			cleared.Interfaces = make(map[string]NDFCInterfacesValue)
			model.SetModelData(cleared)
			return
		}
		// Wait for deploy to propagate (configSyncStatus → inSync).
		waitForDeploySync(ctx, client, dg, fabricName, inData.Interfaces)
	}

	// Single read-back: after deploy (reflects deployed state) or after
	// create-only (refreshes API-returned fields like meta/id).
	if !RscGetInterfaces(ctx, client, dg, model) && !dg.HasError() {
		dg.AddError("Error Creating Interfaces",
			fmt.Sprintf("No interfaces found after create in fabric %q", fabricName))
	}
}

// readDefault refreshes each interface entry via getInterfaceDetails.
// Returns false when *every* entry 404s (resource gone from server).
// This function does NOT touch the deploy flag — use
// RscDeriveDeployStatus separately (from Read only, not Create/Update).
//
// configOnly deviation: the feature spec requires ?configOnly=true on GET to
// suppress operational fields and avoid churn. However, adding configOnly=true
// caused deploy-status drift in acceptance tests: ND returns a reduced
// response that omits fields the provider maps into state, leading to
// false null diffs on every refresh. Since the provider's schema only maps
// configuration-level attributes (not operData.*), the operational fields
// in the full response are harmlessly ignored by the codec. The full GET is
// therefore the safer default until ND's configOnly response is verified to
// include all mapped fields.
func readDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) bool {
	inData := model.GetModelData()
	// PreMarshal populates per-entry SwitchId from the resource-level field,
	// which is needed for the per-interface GET URLs.
	if pm, ok := model.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	resolveSwitchIDs(ctx, dg, inData)
	if dg.HasError() {
		return false
	}
	fabricName := inData.FabricName
	if len(inData.Interfaces) == 0 {
		return false
	}

	fresh := make(map[string]NDFCInterfacesValue, len(inData.Interfaces))
	foundAny := false
	for name, v := range inData.Interfaces {
		ifAPI := api.NewInterfaceAPI(client, fabricName)
		ifAPI.FabricName = fabricName
		ifAPI.SwitchId = v.SwitchId
		ifAPI.InterfaceName = name
		ifAPI.Operation = api.OpInterfaceGetOne
		resp, err := ifAPI.Get()
		if err != nil {
			if strings.Contains(err.Error(), "StatusCode 404") {
				tflog.Warn(ctx, "interface not found on remote; dropping entry", map[string]interface{}{
					"fabric_name":    fabricName,
					"switch_id":      v.SwitchId,
					"interface_name": name,
				})
				continue
			}
			dg.AddError("Error reading Interface",
				fmt.Sprintf("GET %s/%s/%s failed: %v: %s", fabricName, v.SwitchId, name, err, string(resp)))
			return false
		}
		var out NDFCInterfacesValue
		if err := json.Unmarshal(resp, &out); err != nil {
			dg.AddError("Error reading Interface",
				fmt.Sprintf("unmarshal failed for %s/%s/%s: %v\nresponse: %s",
					fabricName, v.SwitchId, name, err, truncateBody(resp, 512)))
			return false
		}
		fresh[name] = out
		foundAny = true
	}
	if !foundAny {
		return false
	}
	inData.Interfaces = fresh
	if pu, ok := model.(InterfacePostUnmarshal); ok {
		pu.PostUnmarshal(ctx, inData)
	}
	model.SetModelData(inData)
	return true
}

// RscDeriveDeployStatus checks /interfacesSummary configSyncStatus for
// each interface in the model. It can only DOWNGRADE deploy from true to
// false (when interfaces are out of sync) — never upgrade from false to
// true. This detects deploy failures/drift when the user intended deploy,
// without forcing a deploy the user explicitly skipped.
//
// Call this from the resource-level Read function ONLY — not from
// Create/Update read-backs (Terraform requires those to return state
// matching the plan).
func RscDeriveDeployStatus(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	inData := model.GetModelData()
	// Only check when state says deploy=true — if false, nothing to downgrade.
	if !inData.Deploy {
		return
	}
	if pm, ok := model.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	fabricName := inData.FabricName
	if len(inData.Interfaces) == 0 {
		return
	}

	// Group interfaces by switchId so we check all switches.
	bySwitch := make(map[string][]string)
	for name, v := range inData.Interfaces {
		if v.SwitchId != "" {
			bySwitch[v.SwitchId] = append(bySwitch[v.SwitchId], name)
		}
	}
	if len(bySwitch) == 0 {
		return
	}

	for switchId, names := range bySwitch {
		statusMap := fetchDeployStatus(ctx, client, fabricName, switchId)
		if statusMap == nil {
			// Summary call failed — keep existing deploy value from state
			// rather than flipping it on a transient error. Surface a
			// warning so the user knows deploy status could not be verified.
			dg.AddWarning("Deploy Status Check Failed",
				fmt.Sprintf("Could not fetch interface summary for switch %q in fabric %q. "+
					"The deploy flag was not verified and may not reflect the actual deployment state.",
					switchId, fabricName))
			continue
		}
		for _, name := range names {
			status, ok := statusMap[name]
			if !ok {
				// Interface not in summary — treat as not in sync.
				tflog.Debug(ctx, "interface not in summary; treating as not deployed", map[string]interface{}{
					"fabric_name":    fabricName,
					"switch_id":      switchId,
					"interface_name": name,
				})
				inData.Deploy = false
				model.SetModelData(inData)
				return
			}
			if !isDeployedStatus(status) {
				tflog.Debug(ctx, "interface not in sync; downgrading deploy flag", map[string]interface{}{
					"fabric_name":        fabricName,
					"switch_id":          switchId,
					"interface_name":     name,
					"config_sync_status": status,
				})
				inData.Deploy = false
				model.SetModelData(inData)
				return
			}
		}
	}
	// All in sync on all switches — deploy stays true.
}

// updateDefault diffs plan vs prior state and the live server, then applies
// the changes in the order prescribed by feature 001 plan.md: delete, deploy,
// add, update, deploy.
//
// Partial-failure / rollback strategy:
//
// ND operations (delete, add, update, deploy) are applied sequentially. If a
// later operation fails, earlier successful operations are NOT rolled back on
// ND. Instead, the Update returns an error WITHOUT calling resp.State.Set(),
// which preserves the prior Terraform state unchanged. On the next
// `terraform plan`, the Read function refreshes state from ND (which now
// reflects the partial mutations), and Terraform shows the delta as planned
// changes. The user then runs `terraform apply` to converge ND to the desired
// state.
//
// Why we do NOT reverse successful ND operations:
//  1. ND's interface APIs are not transactional — there is no atomic
//     multi-operation rollback. Compensating calls (re-create deleted
//     interfaces, undo PUTs) would need the exact prior payloads, which
//     would introduce a second failure point and could leave ND in a worse
//     state than the partial-apply.
//  2. The Terraform model handles partial drift natively: Read detects what
//     ND actually has, plan shows the remaining delta, and the next apply
//     converges. This is the idiomatic Terraform recovery path.
//  3. Deploy is async and not reversible — once ND pushes config to a
//     switch, undoing it requires a second deploy of the old config, which
//     may itself fail or conflict with other pending changes.
//
// It uses the generated NDFCInterfacesValue.CreatePlan (resource_compare_gen.go)
// to decide which in-both entries actually changed. CreatePlan also copies
// server-computed defaults into the plan for fields the user did not set,
// preventing false update diffs.
//
// Both the prior state model and the new plan model are required so that:
//   - Interfaces removed from configuration (in state but absent from plan)
//     are detected and deleted from ND.
//   - Interfaces whose identity fields (switch_id, mode, network_os_type)
//     changed are handled as delete + create rather than in-place update,
//     because CreatePlan returns RequiresReplace for those fields.
func updateDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, stateModel InterfaceModel, planModel InterfaceModel) {
	plan := planModel.GetModelData()
	if pm, ok := planModel.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, plan)
	}
	resolveSwitchIDs(ctx, dg, plan)
	if dg.HasError() {
		return
	}

	stateData := stateModel.GetModelData()
	if pm, ok := stateModel.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, stateData)
	}

	fabricName := plan.FabricName

	// Snapshot the live server state for every interface that's in the plan.
	//
	// Design rationale: adds/updates are classified by comparing the plan
	// against the live server (not prior TF state) for two reasons:
	//   1. CreatePlan backfills server-computed defaults into the plan so the
	//      PUT payload includes them; comparing against TF state would miss
	//      defaults that were never in state.
	//   2. If an interface was deleted out-of-band on ND, the snapshot reports
	//      it as absent and the provider re-creates it — correct recovery when
	//      the Terraform config still declares it.
	// Deletions use prior-state vs plan (line 490) and are not affected.
	serverSnap := refreshServerSnapshot(ctx, client, dg, plan, fabricName)
	if dg.HasError() {
		return
	}

	// Classify entries into delete / add / update / replace sets.
	toDelete := make([]NDFCInterfaceAction, 0)
	toAdd := make(map[string][]NDFCInterfacesValue) // keyed by switchId
	toUpdate := make([]NDFCInterfacesValue, 0)

	// 1. Detect deletions: entries in prior state but absent from plan.
	for name, v := range stateData.Interfaces {
		if _, ok := plan.Interfaces[name]; !ok {
			toDelete = append(toDelete, NDFCInterfaceAction{InterfaceName: name, SwitchId: v.SwitchId})
		}
	}

	// 2. Classify plan entries as add / update / replace.
	for name, pv := range plan.Interfaces {
		sv, inServer := serverSnap[name]
		if !inServer {
			entry := pv
			entry.InterfaceName = name
			toAdd[pv.SwitchId] = append(toAdd[pv.SwitchId], entry)
			continue
		}
		// Both present — compare. CreatePlan returns:
		//   ActionNone      → nothing changed
		//   RequiresUpdate  → in-place PUT
		//   RequiresReplace → identity field changed → delete old + create new
		pvCopy := pv
		cf := false
		action := pvCopy.CreatePlan(sv, &cf)
		switch {
		case action == commontypes.ActionNone:
			// No change.
		case action == commontypes.RequiresReplace:
			// Identity field changed (mode, network_os_type, or switch_id).
			// Delete the old entry and create the new one.
			toDelete = append(toDelete, NDFCInterfaceAction{InterfaceName: name, SwitchId: sv.SwitchId})
			entry := pvCopy
			entry.InterfaceName = name
			toAdd[pvCopy.SwitchId] = append(toAdd[pvCopy.SwitchId], entry)
		default:
			// RequiresUpdate, ControlFlagUpdate, etc. — in-place PUT.
			entry := pvCopy // use pvCopy — CreatePlan may have filled computed fields
			entry.InterfaceName = name
			toUpdate = append(toUpdate, entry)
		}
	}

	tflog.Debug(ctx, "update: classified diff", map[string]interface{}{
		"fabric_name": fabricName,
		"to_delete":   len(toDelete),
		"to_add":      len(toAdd),
		"to_update":   len(toUpdate),
	})

	// Order: delete → deploy → add → update → deploy. The two deploys are
	// gated on plan.Deploy to match create-time behaviour.
	if len(toDelete) > 0 {
		if err := bulkRemove(ctx, client, fabricName, toDelete); err != nil {
			dg.AddError("Error Updating Interfaces", err.Error())
			return
		}
		if plan.Deploy {
			deployInterfaces(ctx, client, dg, fabricName, toDelete)
			if dg.HasError() {
				return
			}
			// Wait for delete-deploy to propagate before proceeding.
			delIfaces := make(map[string]NDFCInterfacesValue, len(toDelete))
			for _, e := range toDelete {
				delIfaces[e.InterfaceName] = NDFCInterfacesValue{SwitchId: e.SwitchId}
			}
			waitForDeploySync(ctx, client, dg, fabricName, delIfaces)
		}
	}

	if len(toAdd) > 0 {
		for switchId, entries := range toAdd {
			if err := bulkCreatePerSwitch(ctx, client, fabricName, switchId, entries); err != nil {
				dg.AddError("Error Updating Interfaces", err.Error())
				return
			}
		}
	}

	for _, entry := range toUpdate {
		if err := singleUpdate(ctx, client, fabricName, entry); err != nil {
			dg.AddError("Error Updating Interface", err.Error())
			return
		}
	}

	// Deploy when the plan says deploy=true AND either:
	//   - there were interface mutations (adds/updates), OR
	//   - deploy flag itself changed (false→true), meaning the user wants
	//     to push previously-saved config to the switch.
	deployFlagChanged := plan.Deploy && !stateData.Deploy
	needsDeploy := plan.Deploy && (len(toAdd) > 0 || len(toUpdate) > 0 || deployFlagChanged)
	if needsDeploy {
		RscDeployInterfaces(ctx, client, dg, planModel)
		if dg.HasError() {
			return
		}
		// Wait for deploy to propagate.
		waitForDeploySync(ctx, client, dg, fabricName, plan.Interfaces)
	}

	// Single read-back: after deploy (reflects deployed state) or after
	// mutations-only (refreshes API-returned fields). One GET instead of two.
	if !RscGetInterfaces(ctx, client, dg, planModel) && !dg.HasError() {
		dg.AddError("Error Updating Interfaces",
			fmt.Sprintf("No interfaces found after update in fabric %q", fabricName))
	}
}

// refreshServerSnapshot GETs the current state of every interface named in
// `plan` so we can diff plan vs server without clobbering the model. 404
// per-entry is treated as "not in server" (returns absence rather than error).
func refreshServerSnapshot(ctx context.Context, client *nd.Client, dg *diag.Diagnostics,
	plan *NDFCInterfaceCommonModel, fabricName string) map[string]NDFCInterfacesValue {

	snap := make(map[string]NDFCInterfacesValue, len(plan.Interfaces))
	for name, v := range plan.Interfaces {
		ifAPI := api.NewInterfaceAPI(client, fabricName)
		ifAPI.FabricName = fabricName
		ifAPI.SwitchId = v.SwitchId
		ifAPI.InterfaceName = name
		ifAPI.Operation = api.OpInterfaceGetOne
		resp, err := ifAPI.Get()
		if err != nil {
			if strings.Contains(err.Error(), "StatusCode 404") {
				continue // not in server
			}
			dg.AddError("Error snapshotting server state",
				fmt.Sprintf("GET %s/%s/%s failed: %v: %s", fabricName, v.SwitchId, name, err, string(resp)))
			return nil
		}
		var out NDFCInterfacesValue
		if err := json.Unmarshal(resp, &out); err != nil {
			dg.AddError("Error snapshotting server state",
				fmt.Sprintf("unmarshal failed for %s/%s/%s: %v\nresponse: %s",
					fabricName, v.SwitchId, name, err, truncateBody(resp, 512)))
			return nil
		}
		snap[name] = out
	}
	return snap
}

// bulkCreatePerSwitch POSTs a group of interfaces for a single switch,
// sub-grouped by policyType (ND rejects mixed policy types in one request),
// and surfaces any per-entry non-success rows in the HTTP 207 response.
func bulkCreatePerSwitch(ctx context.Context, client *nd.Client, fabricName, switchId string,
	entries []NDFCInterfacesValue) error {
	// Sub-group by policyType.
	byPolicy := map[string][]NDFCInterfacesValue{}
	for _, e := range entries {
		pt := e.ConfigData.NetworkOS.Policy.PolicyType
		byPolicy[pt] = append(byPolicy[pt], e)
	}
	for pt, group := range byPolicy {
		payload := NDFCInterfacesPayload{Interfaces: group}
		body, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal failed for switch %q policy %q: %w", switchId, pt, err)
		}
		ifAPI := api.NewInterfaceAPI(client, fabricName)
		ifAPI.FabricName = fabricName
		ifAPI.SwitchId = switchId
		ifAPI.Operation = api.OpInterfaceCreatePerSwitch
		tflog.Debug(ctx, "update: adding interfaces", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id":   switchId,
			"policy_type": pt,
			"count":       len(group),
		})
		res, err := ifAPI.Post(body, nil)
		if err != nil {
			return fmt.Errorf("create failed on switch %q policy %q: %v: %s", switchId, pt, err, res.String())
		}
		if err := bulkResponseErrs(switchId, res.Raw); err != nil {
			return err
		}
	}
	return nil
}

// singleUpdate issues one PUT per changed interface. A 404 is treated as
// "the server dropped it between snapshot and PUT" and we re-create it.
func singleUpdate(ctx context.Context, client *nd.Client, fabricName string,
	entry NDFCInterfacesValue) error {
	body, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshal failed for %s/%s/%s: %w", fabricName, entry.SwitchId, entry.InterfaceName, err)
	}
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.SwitchId = entry.SwitchId
	ifAPI.InterfaceName = entry.InterfaceName
	ifAPI.Operation = api.OpInterfaceUpdateOne
	tflog.Debug(ctx, "update: PUT interface", map[string]interface{}{
		"fabric_name":    fabricName,
		"switch_id":      entry.SwitchId,
		"interface_name": entry.InterfaceName,
	})
	res, err := ifAPI.Put(body, nil)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return bulkCreatePerSwitch(ctx, client, fabricName, entry.SwitchId, []NDFCInterfacesValue{entry})
		}
		return fmt.Errorf("PUT failed for %s/%s/%s: %v: %s", fabricName, entry.SwitchId, entry.InterfaceName, err, res.String())
	}
	return nil
}

// bulkRemove calls /interfaceActions/remove with the given selectors.
func bulkRemove(ctx context.Context, client *nd.Client, fabricName string,
	entries []NDFCInterfaceAction) error {
	payload := NDFCInterfaceActionPayload{Interfaces: entries}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal failed for fabric %q: %w", fabricName, err)
	}
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.Operation = api.OpInterfaceRemoveBulk
	tflog.Debug(ctx, "update: bulk remove", map[string]interface{}{
		"fabric_name": fabricName,
		"count":       len(entries),
	})
	res, err := ifAPI.Post(body, nil)
	if err != nil {
		return fmt.Errorf("bulk remove failed for fabric %q: %v: %s", fabricName, err, res.String())
	}
	// Parse HTTP 207 response and surface per-entry errors.
	if err := bulkResponseErrs("(remove)", res.Raw); err != nil {
		return err
	}
	return nil
}

// deployInterfaces fires /interfaceActions/deploy for the given entries.
// Separate from RscDeployInterfaces so we can deploy a *subset* (e.g. the
// just-deleted entries) during a staged update.
func deployInterfaces(ctx context.Context, client *nd.Client, dg *diag.Diagnostics,
	fabricName string, entries []NDFCInterfaceAction) {
	if len(entries) == 0 {
		return
	}
	body, err := json.Marshal(NDFCInterfaceActionPayload{Interfaces: entries})
	if err != nil {
		dg.AddError("Error Deploying Interfaces",
			fmt.Sprintf("marshal failed for fabric %q: %v", fabricName, err))
		return
	}
	guard := ndapi.Acquire(fabricName, "", ndapi.LockDeploy)
	defer guard.Release()
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.Operation = api.OpInterfaceDeploy
	tflog.Debug(ctx, "update: deploying subset", map[string]interface{}{
		"fabric_name": fabricName,
		"count":       len(entries),
	})
	res, err := ifAPI.DeployPost(body, nil)
	if err != nil {
		dg.AddError("Error Deploying Interfaces",
			fmt.Sprintf("deploy failed for fabric %q: %v: %s", fabricName, err, res.String()))
		return
	}
	// Parse HTTP 207 response and surface per-entry errors.
	if err := bulkResponseErrs("(deploy)", res.Raw); err != nil {
		dg.AddError("Error Deploying Interfaces", err.Error())
	}
}

// bulkResponseErrs parses an HTTP 207 GenericBulkResponse body and returns
// an error summarising any non-success results. Entries whose message
// indicates the interface is already at its default/null state are treated
// as successful (ND cannot "delete" a physical Ethernet port — it can only
// reset it to default, after which further deletes return this message).
func bulkResponseErrs(switchId, raw string) error {
	var br GenericBulkResponse
	if err := json.Unmarshal([]byte(raw), &br); err != nil {
		return nil // not an HTTP 207 — treat as plain success
	}
	msgs := make([]string, 0)
	for _, r := range br.Results {
		if r.Status != "" && !strings.EqualFold(r.Status, "success") {
			if strings.Contains(r.Message, "null Interface cannot be deleted") {
				continue // interface already at default — not an error
			}
			msgs = append(msgs, fmt.Sprintf("%s: %s (%s)", r.Name, r.Message, r.Status))
		}
	}
	if len(msgs) == 0 {
		return nil
	}
	return fmt.Errorf("switch %q had non-success entries: %s", switchId, strings.Join(msgs, "; "))
}

// interfaceSummaryEntry holds the fields we care about from the
// /interfacesSummary response.
type interfaceSummaryEntry struct {
	InterfaceName    string `json:"interfaceName"`
	ConfigSyncStatus string `json:"configSyncStatus"`
}

type interfaceSummaryResponse struct {
	Interfaces []interfaceSummaryEntry `json:"interfaces"`
}

// fetchDeployStatus calls GET /interfacesSummary?switchId=X for the given
// fabric/switch and returns a map of interfaceName → configSyncStatus.
// The configSyncStatus enum (from the OpenAPI spec) includes:
//
//	"inSync", "deployed" → interface is deployed and in sync
//	"outOfSync", "pending", "failed", etc. → not in sync
//
// On error, returns nil (caller should treat as unknown / skip the check).
func fetchDeployStatus(ctx context.Context, client *nd.Client, fabricName, switchId string) map[string]string {
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.SwitchId = switchId
	ifAPI.Operation = api.OpInterfaceSummary
	resp, err := ifAPI.Get()
	if err != nil {
		tflog.Warn(ctx, "could not fetch interface summary for deploy status; skipping", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id":   switchId,
			"error":       err.Error(),
		})
		return nil
	}
	var summary interfaceSummaryResponse
	if err := json.Unmarshal(resp, &summary); err != nil {
		tflog.Warn(ctx, "could not parse interface summary; skipping deploy status", map[string]interface{}{
			"fabric_name": fabricName,
			"switch_id":   switchId,
			"error":       err.Error(),
			"response":    truncateBody(resp, 512),
		})
		return nil
	}
	result := make(map[string]string, len(summary.Interfaces))
	for _, entry := range summary.Interfaces {
		result[entry.InterfaceName] = entry.ConfigSyncStatus
	}
	return result
}

// isDeployedStatus returns true if the configSyncStatus value indicates
// the interface config has been successfully deployed to the switch.
func isDeployedStatus(status string) bool {
	s := strings.ToLower(status)
	return s == "insync" || s == "deployed"
}

// deleteDefault removes all interfaces in the model with a single bulk
// /interfaceActions/remove call. Interfaces not in the model are untouched.
func deleteDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model InterfaceModel) {
	inData := model.GetModelData()
	if pm, ok := model.(InterfacePreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	resolveSwitchIDs(ctx, dg, inData)
	if dg.HasError() {
		return
	}
	fabricName := inData.FabricName
	if len(inData.Interfaces) == 0 {
		return
	}

	payload := NDFCInterfaceActionPayload{Interfaces: make([]NDFCInterfaceAction, 0, len(inData.Interfaces))}
	for name, v := range inData.Interfaces {
		payload.Interfaces = append(payload.Interfaces, NDFCInterfaceAction{
			InterfaceName: name,
			SwitchId:      v.SwitchId,
		})
	}
	body, err := json.Marshal(payload)
	if err != nil {
		dg.AddError("Error Deleting Interfaces",
			fmt.Sprintf("marshal failed for fabric %q: %v", fabricName, err))
		return
	}
	ifAPI := api.NewInterfaceAPI(client, fabricName)
	ifAPI.FabricName = fabricName
	ifAPI.Operation = api.OpInterfaceRemoveBulk
	tflog.Debug(ctx, "Deleting interfaces (bulk)", map[string]interface{}{
		"fabric_name": fabricName,
		"count":       len(payload.Interfaces),
	})
	res, err := ifAPI.Post(body, nil)
	if err != nil {
		dg.AddError("Error Deleting Interfaces",
			fmt.Sprintf("bulk remove failed for fabric %q: %v: %s", fabricName, err, res.String()))
		return
	}
	// Parse HTTP 207 response and surface per-entry errors.
	if err := bulkResponseErrs("(delete)", res.Raw); err != nil {
		dg.AddError("Partial Delete Failure", err.Error())
	}
}
