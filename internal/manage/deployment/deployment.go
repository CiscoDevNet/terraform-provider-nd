// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package deployment

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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	nd "github.com/netascode/go-nd"
)

// DeployOptions holds optional parameters for config-save and deploy operations.
type DeployOptions struct {
	SerialNumbers                 []string // switch serial numbers; empty or ["ALL"] = fabric-wide
	TicketId                      string   // change control ticket ID for config-save
	ForceShowRun                  bool     // fetch latest running config before deploy
	IncludeAllFabricGroupSwitches bool     // deploy to all member fabrics of a fabric group
}

// DeployResult holds the outcome of a config-save/deploy operation.
type DeployResult struct {
	Status string
}

// checkDeployResponse parses the deploy API response and returns an error
// if any switch reports a non-success status.
func checkDeployResponse(resp string, fabricName string) (string, error) {
	type switchResult struct {
		SwitchId   string `json:"switchId"`
		SwitchName string `json:"switchName"`
		Status     string `json:"status"`
		Message    string `json:"message"`
	}
	type deployResponse struct {
		SwitchIds []switchResult `json:"switchIds"`
	}

	var dr deployResponse
	if err := json.Unmarshal([]byte(resp), &dr); err != nil {
		return "", fmt.Errorf("deploy response parse error for fabric %s: %w, body: %s", fabricName, err, resp)
	}

	var failed []string
	for _, sw := range dr.SwitchIds {
		if !strings.EqualFold(sw.Status, "success") {
			failed = append(failed, fmt.Sprintf("%s(%s): %s", sw.SwitchName, sw.SwitchId, sw.Message))
		}
	}

	if len(failed) > 0 {
		return "", fmt.Errorf("deploy failed for fabric %s: %s", fabricName, strings.Join(failed, "; "))
	}

	if len(dr.SwitchIds) == 0 {
		return "Deploy completed (no pending changes)", nil
	}

	var msgs []string
	for _, sw := range dr.SwitchIds {
		msgs = append(msgs, fmt.Sprintf("%s(%s): %s", sw.SwitchName, sw.SwitchId, sw.Message))
	}
	return strings.Join(msgs, "; "), nil
}

// fetchDeploymentFailures fetches the most recent deployment history records
// and returns a human-readable summary of any failed entries, including the
// full configCommandResponses for each.
func fetchDeploymentFailures(client *nd.Client, fabricName string, maxRecords int) string {
	histAPI := api.NewConfigAPI(client, ndapi.DefaultFabric)
	histAPI.FabricName = fabricName
	histAPI.SetOperation(api.OpDeployHistory)
	histAPI.SetQueryParams(
		"sort=startTimestamp:desc",
		fmt.Sprintf("max=%d", maxRecords),
		"offset=0",
	)

	raw, err := histAPI.Get()
	if err != nil {
		log.Printf("[DEBUG] fetchDeploymentFailures: could not fetch history: %v", err)
		return ""
	}

	type cmdResponse struct {
		Command                string `json:"command"`
		Status                 string `json:"status"`
		CommandExecutionStatus string `json:"commandExecutionStatus"`
		CliResponse            string `json:"cliResponse"`
	}
	type historyRecord struct {
		SwitchId               string        `json:"switchId"`
		Hostname               string        `json:"hostname"`
		Status                 string        `json:"status"`
		StatusDescription      string        `json:"statusDescription"`
		StartTimestamp         string        `json:"startTimestamp"`
		ConfigCommandResponses []cmdResponse `json:"configCommandResponses"`
	}
	type historyResponse struct {
		DeploymentRecords []historyRecord `json:"deploymentRecords"`
	}

	var hr historyResponse
	if err := json.Unmarshal(raw, &hr); err != nil {
		log.Printf("[DEBUG] fetchDeploymentFailures: parse error: %v", err)
		return ""
	}

	var parts []string
	for _, rec := range hr.DeploymentRecords {
		if strings.EqualFold(rec.Status, "success") {
			continue
		}
		var cmds []string
		for _, cmd := range rec.ConfigCommandResponses {
			cmds = append(cmds, fmt.Sprintf("    cmd=%q status=%s cli=%q",
				cmd.Command, cmd.CommandExecutionStatus, cmd.CliResponse))
		}
		entry := fmt.Sprintf("  %s(%s) status=%s: %s\n%s",
			rec.Hostname, rec.SwitchId, rec.Status, rec.StatusDescription,
			strings.Join(cmds, "\n"))
		parts = append(parts, entry)
	}

	if len(parts) == 0 {
		return ""
	}
	return fmt.Sprintf("\nDeployment history (recent failures):\n%s", strings.Join(parts, "\n"))
}

// ConfigSaveAndDeploy performs both config-save and config-deploy operations.
// Acquires LockDeploy (Global.RLock + Fabric.WLock) — blocks all CRUD on
// the fabric and waits for in-flight CRUD to finish before proceeding.
func ConfigSaveAndDeploy(ctx context.Context, client *nd.Client, fabricName string, recalculate bool, deploy bool, dg *diag.Diagnostics) {
	ConfigSaveAndDeployWithOpts(ctx, client, fabricName, recalculate, deploy, nil, dg)
}

// ConfigSaveAndDeployWithOpts performs config-save and/or deploy with additional options.
// Acquires LockDeploy (Global.RLock + Fabric.WLock).
// Returns a DeployResult with status information from the API responses.
func ConfigSaveAndDeployWithOpts(ctx context.Context, client *nd.Client, fabricName string, recalculate bool, deploy bool, opts *DeployOptions, dg *diag.Diagnostics) *DeployResult {
	// NOTE: lock scope is intentionally DefaultFabric ("global"), not fabricName.
	// All resource CRUD in this provider currently locks on the same global scope,
	// so a deploy must take the global fabric write-lock to reliably block all
	// in-flight/concurrent CRUD. Narrowing this to per-fabric would require every
	// resource's CRUD to also adopt per-fabric scoping; until that cross-cutting
	// change is made, per-fabric locking here would silently lose mutual exclusion.
	guard := ndapi.Acquire(ndapi.DefaultFabric, "", ndapi.LockDeploy)
	defer guard.Release()

	result := &DeployResult{}

	if opts == nil {
		opts = &DeployOptions{}
	}

	// Accumulate status from each phase so a successful config-save is not lost
	// when a later deploy fails, and both phases are reflected when both run.
	var statuses []string

	if recalculate {
		status, err := doConfigSave(ctx, client, fabricName, opts)
		if status != "" {
			statuses = append(statuses, status)
			result.Status = strings.Join(statuses, "; ")
		}
		if err != nil {
			dg.AddError("Error Saving Config", fmt.Sprintf("%v", err))
			return result
		}
	}

	if deploy {
		status, err := doDeploy(ctx, client, fabricName, opts)
		if status != "" {
			statuses = append(statuses, status)
			result.Status = strings.Join(statuses, "; ")
		}
		if err != nil {
			dg.AddError("Error Deploying Config", fmt.Sprintf("%v", err))
			return result
		}
	}

	return result
}

// transientRetryMaxAttempts is the maximum number of retries for transient
// config-save/deploy failures such as "switches are reloading" after fresh discovery.
const transientRetryMaxAttempts = 12

// transientRetryInterval is the time between retries.
const transientRetryInterval = 15 * time.Second

// isTransientSwitchError checks whether an error is a transient condition that
// should be retried (e.g. switches still stabilising after discovery).
func isTransientSwitchError(body string) bool {
	lower := strings.ToLower(body)
	return strings.Contains(lower, "are reloading") ||
		strings.Contains(lower, "not reachable")
}

// doConfigSave performs config-save with optional ticketId.
// Retries automatically on transient "switches are reloading" errors.
// Must be called within a deploy context (LockDeploy already held).
func doConfigSave(ctx context.Context, client *nd.Client, fabricName string, opts *DeployOptions) (string, error) {
	for attempt := 0; attempt <= transientRetryMaxAttempts; attempt++ {
		configAPI := api.NewConfigAPI(client, ndapi.DefaultFabric)
		configAPI.FabricName = fabricName
		configAPI.SetOperation(api.OpConfigSave)

		if opts != nil && opts.TicketId != "" {
			configAPI.SetQueryParams(fmt.Sprintf("ticketId=%s", opts.TicketId))
		}

		resp, err := configAPI.DeployPost(nil, nil)
		if err != nil {
			body := resp.String()
			if isTransientSwitchError(body) && attempt < transientRetryMaxAttempts {
				tflog.Info(ctx, "Config save: switches still reloading, will retry", map[string]interface{}{
					"fabric_name": fabricName,
					"attempt":     attempt + 1,
					"max_retries": transientRetryMaxAttempts,
					"retry_in":    transientRetryInterval.String(),
				})
				select {
				case <-ctx.Done():
					return "", fmt.Errorf("config save cancelled while waiting for switches: %w", ctx.Err())
				case <-time.After(transientRetryInterval):
					continue
				}
			}
			return "", fmt.Errorf("config save failed for fabric %s: %w: %s", fabricName, err, body)
		}

		status := resp.Get("status").String()
		tflog.Info(ctx, "Fabric config saved (recalculate)", map[string]interface{}{
			"fabric_name": fabricName,
			"status":      status,
		})
		return status, nil
	}

	return "", fmt.Errorf("config save failed for fabric %s: exhausted retries waiting for switches to stop reloading", fabricName)
}

// deployPostWithRetry runs a deploy POST and retries on transient
// "switches reloading / not reachable" errors, mirroring doConfigSave.
// The post closure returns the response body and error. Must be called within
// a deploy context (LockDeploy already held).
func deployPostWithRetry(ctx context.Context, fabricName string, post func() (string, error)) (string, error) {
	var body string
	var err error
	for attempt := 0; attempt <= transientRetryMaxAttempts; attempt++ {
		body, err = post()
		if err == nil {
			return body, nil
		}
		if isTransientSwitchError(body) && attempt < transientRetryMaxAttempts {
			tflog.Info(ctx, "Deploy: switches still reloading, will retry", map[string]interface{}{
				"fabric_name": fabricName,
				"attempt":     attempt + 1,
				"max_retries": transientRetryMaxAttempts,
				"retry_in":    transientRetryInterval.String(),
			})
			select {
			case <-ctx.Done():
				return body, fmt.Errorf("deploy cancelled while waiting for switches: %w", ctx.Err())
			case <-time.After(transientRetryInterval):
				continue
			}
		}
		return body, err
	}
	return body, err
}

// doDeploy performs fabric-wide or switch-specific deploy.
// Must be called within a deploy context (LockDeploy already held).
func doDeploy(ctx context.Context, client *nd.Client, fabricName string, opts *DeployOptions) (string, error) {
	configAPI := api.NewConfigAPI(client, ndapi.DefaultFabric)
	configAPI.FabricName = fabricName

	var qp []string
	if opts != nil && opts.ForceShowRun {
		qp = append(qp, "forceShowRun=true")
	}

	isSwitchDeploy := opts != nil && len(opts.SerialNumbers) > 0 &&
		!strings.EqualFold(opts.SerialNumbers[0], "ALL")

	if isSwitchDeploy {
		configAPI.SetOperation(api.OpSwitchDeploy)
		if len(qp) > 0 {
			configAPI.SetQueryParams(qp...)
		}

		payload := map[string][]string{"switchIds": opts.SerialNumbers}
		body, err := json.Marshal(payload)
		if err != nil {
			return "", fmt.Errorf("could not marshal switch deploy payload: %w", err)
		}

		respStr, err := deployPostWithRetry(ctx, fabricName, func() (string, error) {
			resp, err := configAPI.DeployPost(body, nil)
			return resp.String(), err
		})
		if err != nil {
			return "", fmt.Errorf("switch deploy failed for fabric %s: %w: %s", fabricName, err, respStr)
		}

		status, chkErr := checkDeployResponse(respStr, fabricName)
		if chkErr != nil {
			history := fetchDeploymentFailures(client, fabricName, 5)
			return "", fmt.Errorf("%w%s", chkErr, history)
		}
		tflog.Info(ctx, "Switch deploy completed", map[string]interface{}{
			"fabric_name":    fabricName,
			"serial_numbers": opts.SerialNumbers,
			"status":         status,
		})
		return status, nil
	}

	// Fabric-wide deploy
	configAPI.SetOperation(api.OpFabricDeploy)
	if opts != nil && opts.IncludeAllFabricGroupSwitches {
		qp = append(qp, "inclAllFabricGroupsSwitches=true")
	}
	if len(qp) > 0 {
		configAPI.SetQueryParams(qp...)
	}

	respStr, err := deployPostWithRetry(ctx, fabricName, func() (string, error) {
		resp, err := configAPI.DeployPost(nil, nil)
		return resp.String(), err
	})
	if err != nil {
		return "", fmt.Errorf("fabric deploy failed for fabric %s: %w: %s", fabricName, err, respStr)
	}

	status, chkErr := checkDeployResponse(respStr, fabricName)
	if chkErr != nil {
		history := fetchDeploymentFailures(client, fabricName, 5)
		return "", fmt.Errorf("%w%s", chkErr, history)
	}
	tflog.Info(ctx, "Fabric config deployed", map[string]interface{}{
		"fabric_name": fabricName,
		"status":      status,
	})
	return status, nil
}
