// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_policy

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"terraform-provider-nd/internal/common/templates"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/deployment"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

type policyAPIObject struct {
	PolicyId            string          `json:"policyId,omitempty"`
	TemplateName        string          `json:"templateName,omitempty"`
	EntityType          string          `json:"entityType,omitempty"`
	EntityName          string          `json:"entityName,omitempty"`
	Description         string          `json:"description,omitempty"`
	Source              string          `json:"source"`
	Priority            *int64          `json:"priority,omitempty"`
	SwitchId            json.RawMessage `json:"switchId,omitempty"`
	SwitchIds           []string        `json:"switchIds,omitempty"`
	TemplateInputs      map[string]any  `json:"templateInputs,omitempty"`
	SecondaryEntityName string          `json:"secondaryEntityName,omitempty"`
	SecondaryEntityType string          `json:"secondaryEntityType,omitempty"`
	GeneratedConfig     string          `json:"generatedConfig,omitempty"`
	MarkDeleted         *bool           `json:"markDeleted,omitempty"`
	Status              string          `json:"status,omitempty"`
	Message             string          `json:"message,omitempty"`
}

type inventorySwitchesResponse struct {
	Switches []inventorySwitchEntry `json:"switches,omitempty"`
}

type inventorySwitchEntry struct {
	SerialNumber       string `json:"serialNumber,omitempty"`
	FabricName         string `json:"fabricName,omitempty"`
	Fabric             string `json:"fabric,omitempty"`
	IPAddress          string `json:"ip,omitempty"`
	FabricManagementIp string `json:"fabricManagementIp,omitempty"`
}

const urlInventorySwitches = "/manage/inventory/switches"

func (s inventorySwitchEntry) resolvedFabricName() string {
	if strings.TrimSpace(s.FabricName) != "" {
		return s.FabricName
	}
	return strings.TrimSpace(s.Fabric)
}

func (s inventorySwitchEntry) matches(id string) bool {
	return s.SerialNumber == id || s.FabricManagementIp == id || s.IPAddress == id
}

func (r *policyResource) newPolicyAPI(in *NDFCPolicyModel) *api.PolicyAPI {
	policyAPI := api.NewPolicyAPI(r.manageClient.ApiClient, in.FabricName)
	policyAPI.FabricName = in.FabricName
	policyAPI.PolicyID = in.PolicyId
	policyAPI.PolicyGroup = in.IsPolicyGroup
	return policyAPI
}

func (r *policyResource) rscCreatePolicy(ctx context.Context, dg *diag.Diagnostics, input *PolicyModel) {
	if input == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return
	}

	inData := policyModelData(input)
	if !r.preparePolicyWrite(ctx, dg, inData) {
		return
	}

	payload, err := buildPolicyWritePayload(inData, true)
	if err != nil {
		dg.AddError("Error Creating Policy", fmt.Sprintf("Could not marshal policy payload: %v", err))
		return
	}

	policyAPI := r.newPolicyAPI(inData)
	tflog.Debug(ctx, "Creating policy", map[string]interface{}{
		"fabric_name":     inData.FabricName,
		"is_policy_group": inData.IsPolicyGroup,
		"url":             policyAPI.PostUrl(),
	})

	res, err := policyAPI.Post(payload, nil)
	if err != nil {
		tflog.Error(ctx, "Policy create API call failed", map[string]interface{}{
			"url":      policyAPI.PostUrl(),
			"error":    err.Error(),
			"response": res.Raw,
		})
		dg.AddError("Error Creating Policy", fmt.Sprintf("Could not create policy: %v %v", err, res))
		return
	}

	policyID, err := parseCreateResponse(res, inData.IsPolicyGroup)
	if err != nil {
		dg.AddError("Error Creating Policy", err.Error())
		return
	}
	inData.PolicyId = policyID

	if inData.Deploy {
		if err := r.deployPolicy(ctx, inData); err != nil {
			dg.AddError("Error Deploying Policy", err.Error())
			delDg := diag.Diagnostics{}
			rollback := *input
			setPolicyIDs(&rollback, policyID)
			r.rscDeletePolicy(ctx, &delDg, &rollback)
			if delDg.HasError() {
				dg.Append(delDg...)
			}
			return
		}
	}

	if !r.refreshPolicyState(ctx, dg, input, inData) {
		return
	}
}

func (r *policyResource) rscReadPolicy(ctx context.Context, dg *diag.Diagnostics, model *PolicyModel) bool {
	inData := policyModelData(model)
	if err := r.ensurePolicyFabricName(ctx, inData); err != nil {
		dg.AddError("Error Reading Policy", err.Error())
		return false
	}

	outData, err := r.fetchPolicy(ctx, inData)
	if err != nil {
		if isNotFoundError(err) {
			return false
		}
		dg.AddError("Error Reading Policy", fmt.Sprintf("Could not read policy %s: %v", inData.PolicyId, err))
		return false
	}

	applyPolicyState(dg, model, outData)
	return true
}

func (r *policyResource) rscUpdatePolicy(ctx context.Context, dg *diag.Diagnostics, input *PolicyModel) {
	inData := policyModelData(input)
	if !r.preparePolicyWrite(ctx, dg, inData) {
		return
	}

	payload, err := buildPolicyWritePayload(inData, false)
	if err != nil {
		dg.AddError("Error Updating Policy", fmt.Sprintf("Could not marshal policy payload: %v", err))
		return
	}

	policyAPI := r.newPolicyAPI(inData)
	tflog.Debug(ctx, "Updating policy", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"policy_id":   inData.PolicyId,
		"url":         policyAPI.PutUrl(),
	})

	res, err := policyAPI.Put(payload, nil)
	if err != nil {
		tflog.Error(ctx, "Policy update API call failed", map[string]interface{}{
			"url":      policyAPI.PutUrl(),
			"error":    err.Error(),
			"response": res.Raw,
		})
		dg.AddError("Error Updating Policy", fmt.Sprintf("Could not update policy: %v %v", err, res))
		return
	}

	if inData.Deploy {
		if err := r.deployPolicy(ctx, inData); err != nil {
			dg.AddError("Error Deploying Policy", err.Error())
			return
		}
	}

	if !r.refreshPolicyState(ctx, dg, input, inData) {
		return
	}
}

func (r *policyResource) rscDeletePolicy(ctx context.Context, dg *diag.Diagnostics, state *PolicyModel) {
	if state == nil {
		dg.AddError("Error Deleting Policy", "The current policy state is missing.")
		return
	}

	inData := policyModelData(state)
	if strings.TrimSpace(inData.PolicyId) == "" {
		dg.AddError("Error Deleting Policy", "Policy ID is missing from state.")
		return
	}
	if err := r.ensurePolicyFabricName(ctx, inData); err != nil {
		dg.AddError("Error Deleting Policy", err.Error())
		return
	}

	serials := r.resolveSerialNumbers(ctx, dg, inData.FabricName, inData.SerialNumbers)
	if dg.HasError() {
		return
	}
	inData.SerialNumbers = serials

	policyAPI := r.newPolicyAPI(inData)
	markErr := r.markDeletePolicy(ctx, policyAPI, inData.PolicyId)
	if markErr != nil {
		tflog.Warn(ctx, "Mark-delete failed; falling back to direct delete then switch deploy", map[string]interface{}{
			"policy_id": inData.PolicyId,
			"error":     markErr.Error(),
		})
		if err := r.deletePolicyRecord(ctx, policyAPI); err != nil && !isNotFoundError(err) {
			dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not delete policy %s: %v", inData.PolicyId, err))
			return
		}
		r.cleanupSwitchDeploy(ctx, dg, inData)
		return
	}

	r.cleanupSwitchDeploy(ctx, dg, inData)
	if dg.HasError() {
		return
	}

	outData, err := r.fetchPolicy(ctx, inData)
	if err != nil {
		if isNotFoundError(err) {
			return
		}
		dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not verify policy %s after cleanup: %v", inData.PolicyId, err))
		return
	}
	if outData == nil {
		return
	}

	if err := r.deletePolicyRecord(ctx, policyAPI); err != nil && !isNotFoundError(err) {
		dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not delete remaining policy %s: %v", inData.PolicyId, err))
	}
}

func (r *policyResource) preparePolicyWrite(ctx context.Context, dg *diag.Diagnostics, inData *NDFCPolicyModel) bool {
	if !r.resolveSwitchesAndFabric(ctx, dg, inData) {
		return false
	}

	if inData.IsPolicyGroup {
		if len(inData.SerialNumbers) == 0 {
			dg.AddAttributeError(
				path.Root("serial_numbers"),
				"Invalid Configuration",
				"serial_numbers must contain at least one switch when is_policy_group is true.",
			)
			return false
		}
	} else if len(inData.SerialNumbers) != 1 {
		dg.AddAttributeError(
			path.Root("serial_numbers"),
			"Invalid Configuration",
			"serial_numbers must contain exactly one switch when is_policy_group is false.",
		)
		return false
	}

	if err := r.validateTemplateInputs(ctx, inData); err != nil {
		dg.AddAttributeError(path.Root("policy_parameters"), "Invalid Template Inputs", err.Error())
		return false
	}

	return true
}

func (r *policyResource) resolveSwitchesAndFabric(ctx context.Context, dg *diag.Diagnostics, inData *NDFCPolicyModel) bool {
	if inData == nil {
		dg.AddError("Invalid Input", "The input model is nil")
		return false
	}

	inventory, err := r.loadInventorySwitches(ctx)
	if err != nil {
		dg.AddError("Error Reading Switch Inventory", err.Error())
		return false
	}

	expectedFabric := strings.TrimSpace(inData.FabricName)
	resolved := make([]string, 0, len(inData.SerialNumbers))
	seen := make(map[string]bool, len(inData.SerialNumbers))
	fabric := expectedFabric

	for _, id := range inData.SerialNumbers {
		entry, ok := findInventorySwitch(inventory, id)
		if !ok {
			dg.AddAttributeError(
				path.Root("serial_numbers"),
				"Switch Not Found",
				fmt.Sprintf("Switch %q was not found in inventory. Verify the serial or management IP.", id),
			)
			return false
		}

		serial := entry.SerialNumber
		if serial == "" {
			serial = id
		}
		switchFabric := entry.resolvedFabricName()
		if switchFabric == "" {
			dg.AddError(
				"Unable to Resolve Fabric",
				fmt.Sprintf("Switch %q does not have a fabric association in inventory.", serial),
			)
			return false
		}
		if fabric == "" {
			fabric = switchFabric
		} else if switchFabric != fabric {
			dg.AddError(
				"Invalid Configuration",
				fmt.Sprintf("Switches belong to different fabrics: %s is in %s, expected %s.", serial, switchFabric, fabric),
			)
			return false
		}
		if !seen[serial] {
			seen[serial] = true
			resolved = append(resolved, serial)
		}
	}

	if fabric == "" {
		dg.AddError(
			"Unable to Resolve Fabric",
			"Could not determine fabric_name from serial_numbers. Verify the switches are associated with a fabric in inventory.",
		)
		return false
	}

	inData.SerialNumbers = resolved
	inData.FabricName = fabric
	tflog.Debug(ctx, "Resolved policy fabric from switch inventory", map[string]interface{}{
		"fabric_name":     fabric,
		"serial_numbers":  resolved,
		"is_policy_group": inData.IsPolicyGroup,
	})
	return true
}

func policyModelData(model *PolicyModel) *NDFCPolicyModel {
	if model == nil {
		return new(NDFCPolicyModel)
	}

	inData := model.GetModelData()
	if inData == nil {
		inData = new(NDFCPolicyModel)
	}
	// Generated GetModelData omits computed fabric_name; copy it from the TF model.
	if strings.TrimSpace(inData.FabricName) == "" && !model.FabricName.IsNull() && !model.FabricName.IsUnknown() {
		inData.FabricName = model.FabricName.ValueString()
	}
	if strings.TrimSpace(inData.PolicyId) == "" && !model.PolicyId.IsNull() && !model.PolicyId.IsUnknown() {
		inData.PolicyId = model.PolicyId.ValueString()
	}
	return inData
}

func (r *policyResource) ensurePolicyFabricName(ctx context.Context, inData *NDFCPolicyModel) error {
	if inData == nil {
		return fmt.Errorf("input model is nil")
	}
	if strings.TrimSpace(inData.FabricName) != "" {
		return nil
	}

	fabric, err := r.resolvePolicyFabricName(ctx, inData.SerialNumbers)
	if err != nil {
		return err
	}
	inData.FabricName = fabric
	return nil
}

func (r *policyResource) resolvePolicyFabricName(ctx context.Context, switchIDs []string) (string, error) {
	if len(switchIDs) == 0 {
		return "", fmt.Errorf("serial_numbers is empty; cannot resolve fabric_name")
	}

	inventory, err := r.loadInventorySwitches(ctx)
	if err != nil {
		return "", err
	}

	var fabric string
	for _, id := range switchIDs {
		entry, ok := findInventorySwitch(inventory, id)
		if !ok {
			return "", fmt.Errorf("switch %s not found in inventory switches response", id)
		}
		switchFabric := entry.resolvedFabricName()
		if switchFabric == "" {
			return "", fmt.Errorf("switch %s does not have a fabric association in inventory", id)
		}
		if fabric == "" {
			fabric = switchFabric
			continue
		}
		if fabric != switchFabric {
			return "", fmt.Errorf("switches belong to different fabrics: %s=%s, %s=%s", switchIDs[0], fabric, id, switchFabric)
		}
	}

	tflog.Debug(ctx, "Resolved policy fabric from inventory switches", map[string]interface{}{
		"serial_numbers": switchIDs,
		"fabric_name":    fabric,
	})
	return fabric, nil
}

func (r *policyResource) loadInventorySwitches(ctx context.Context) ([]inventorySwitchEntry, error) {
	tflog.Debug(ctx, "Reading inventory switches to resolve policy fabric", map[string]interface{}{
		"url": urlInventorySwitches,
	})

	respData, err := r.manageClient.ApiClient.GetRawJson(urlInventorySwitches)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", urlInventorySwitches, err)
	}

	var resp inventorySwitchesResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, fmt.Errorf("decode inventory switches response: %w", err)
	}
	return resp.Switches, nil
}

func findInventorySwitch(switches []inventorySwitchEntry, id string) (inventorySwitchEntry, bool) {
	for _, sw := range switches {
		if sw.matches(id) {
			return sw, true
		}
	}
	return inventorySwitchEntry{}, false
}

func (r *policyResource) resolveSerialNumbers(ctx context.Context, dg *diag.Diagnostics, fabricName string, ids []string) []string {
	if len(ids) == 0 {
		return ids
	}

	resolved := make([]string, 0, len(ids))
	seen := make(map[string]bool, len(ids))
	for _, id := range ids {
		serial := id
		if net.ParseIP(id) != nil {
			s, ok := r.manageClient.SwitchDB.GetSerialByIP(ctx, fabricName, id)
			if !ok {
				dg.AddAttributeError(
					path.Root("serial_numbers"),
					"Switch Not Found",
					fmt.Sprintf("Could not resolve IP address %q to a serial number in fabric %q. Verify the switch is in the fabric inventory.", id, fabricName),
				)
				return nil
			}
			tflog.Debug(ctx, "Resolved switch IP to serial", map[string]interface{}{
				"ip":     id,
				"serial": s,
				"fabric": fabricName,
			})
			serial = s
		}
		if !seen[serial] {
			seen[serial] = true
			resolved = append(resolved, serial)
		}
	}
	return resolved
}

func (r *policyResource) validateTemplateInputs(ctx context.Context, inData *NDFCPolicyModel) error {
	info, err := templates.GetInstance().LoadTemplate(ctx, inData.TemplateName)
	if err != nil {
		return fmt.Errorf("could not read template %q parameters: %w", inData.TemplateName, err)
	}

	tflog.Debug(ctx, "Validated policy template schema", map[string]interface{}{
		"template_name":       info.TemplateName,
		"content_type":        info.ContentType,
		"supported_platforms": string(info.SupportedPlatforms),
		"parameter_count":     len(info.Params),
	})

	if unknown := info.UnknownParamNames(inData.PolicyParameters); len(unknown) > 0 {
		return fmt.Errorf("policy_parameters contains keys not defined by template %q: %s", inData.TemplateName, strings.Join(unknown, ", "))
	}
	if missing := info.MissingRequiredNames(inData.PolicyParameters); len(missing) > 0 {
		return fmt.Errorf("required template parameter %q is missing from policy_parameters", missing[0])
	}
	return nil
}

func (r *policyResource) refreshPolicyState(ctx context.Context, dg *diag.Diagnostics, model *PolicyModel, inData *NDFCPolicyModel) bool {
	outData, err := r.fetchPolicy(ctx, inData)
	if err != nil {
		dg.AddError("Error Reading Policy", fmt.Sprintf("Could not read policy %s after apply: %v", inData.PolicyId, err))
		return false
	}
	applyPolicyState(dg, model, outData)
	return true
}

func (r *policyResource) fetchPolicy(ctx context.Context, inData *NDFCPolicyModel) (*NDFCPolicyModel, error) {
	policyAPI := r.newPolicyAPI(inData)
	tflog.Debug(ctx, "Fetching policy from ND", map[string]interface{}{
		"fabric_name":     inData.FabricName,
		"policy_id":       inData.PolicyId,
		"is_policy_group": inData.IsPolicyGroup,
		"url":             policyAPI.GetUrl(),
	})

	resp, err := policyAPI.Get()
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", policyAPI.GetUrl(), err)
	}

	obj, err := decodePolicyObject(resp)
	if err != nil {
		return nil, err
	}

	return policyObjectToModel(obj, inData), nil
}

func decodePolicyObject(resp []byte) (*policyAPIObject, error) {
	var obj policyAPIObject
	if err := json.Unmarshal(resp, &obj); err != nil {
		return nil, fmt.Errorf("decode policy response: %w", err)
	}
	if obj.PolicyId != "" || obj.TemplateName != "" {
		return &obj, nil
	}

	var wrapped struct {
		Policies     []policyAPIObject `json:"policies"`
		PolicyGroups []policyAPIObject `json:"policyGroups"`
	}
	if err := json.Unmarshal(resp, &wrapped); err == nil {
		if len(wrapped.Policies) > 0 {
			return &wrapped.Policies[0], nil
		}
		if len(wrapped.PolicyGroups) > 0 {
			return &wrapped.PolicyGroups[0], nil
		}
	}

	return &obj, nil
}

func (r *policyResource) findPolicyFabric(ctx context.Context, policyID string, isGroup bool) (string, error) {
	resp, err := r.manageClient.ApiClient.GetRawJson(api.UrlFabrics)
	if err != nil {
		return "", fmt.Errorf("list fabrics: %w", err)
	}

	var names []string
	gjson.GetBytes(resp, "fabrics").ForEach(func(_, fab gjson.Result) bool {
		name := fab.Get("name").String()
		if name == "" {
			name = fab.Get("fabricName").String()
		}
		if name != "" {
			names = append(names, name)
		}
		return true
	})
	if len(names) == 0 {
		return "", fmt.Errorf("no fabrics returned by GET %s", api.UrlFabrics)
	}

	for _, fabric := range names {
		inData := &NDFCPolicyModel{
			FabricName:    fabric,
			PolicyId:      policyID,
			IsPolicyGroup: isGroup,
		}
		if _, err := r.fetchPolicy(ctx, inData); err == nil {
			return fabric, nil
		} else if !isNotFoundError(err) {
			tflog.Debug(ctx, "Policy lookup failed for fabric during import", map[string]interface{}{
				"fabric_name": fabric,
				"policy_id":   policyID,
				"error":       err.Error(),
			})
		}
	}

	return "", fmt.Errorf("policy %s was not found in any fabric", policyID)
}

func (r *policyResource) deployPolicy(ctx context.Context, inData *NDFCPolicyModel) error {
	if inData.IsPolicyGroup {
		return r.switchDeploy(ctx, inData, true)
	}

	policyAPI := r.newPolicyAPI(inData)
	policyAPI.SetOperation(api.OpPushConfig)
	payload, err := json.Marshal(map[string][]string{"policyIds": {inData.PolicyId}})
	if err != nil {
		return fmt.Errorf("marshal pushConfig payload: %w", err)
	}

	tflog.Debug(ctx, "Pushing policy config", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"policy_id":   inData.PolicyId,
		"url":         policyAPI.PostUrl(),
	})

	res, err := policyAPI.Post(payload, nil)
	if err != nil {
		return fmt.Errorf("pushConfig failed for %s: %w: %s", inData.PolicyId, err, res.Raw)
	}
	if err := checkBulkStatus(res, inData.PolicyId); err != nil {
		return err
	}
	return nil
}

func (r *policyResource) markDeletePolicy(ctx context.Context, policyAPI *api.PolicyAPI, policyID string) error {
	policyAPI.SetOperation(api.OpMarkDelete)
	payload, err := json.Marshal(map[string][]string{"policyIds": {policyID}})
	if err != nil {
		return fmt.Errorf("marshal markDelete payload: %w", err)
	}

	tflog.Debug(ctx, "Mark-deleting policy", map[string]interface{}{
		"policy_id": policyID,
		"url":       policyAPI.PostUrl(),
	})

	res, err := policyAPI.Post(payload, nil)
	if err != nil {
		return fmt.Errorf("markDelete failed for %s: %w: %s", policyID, err, res.Raw)
	}
	return checkBulkStatus(res, policyID)
}

func (r *policyResource) deletePolicyRecord(ctx context.Context, policyAPI *api.PolicyAPI) error {
	policyAPI.SetOperation(api.OpPolicyCRUD)
	tflog.Debug(ctx, "Deleting policy record", map[string]interface{}{
		"url": policyAPI.DeleteUrl(),
	})
	res, err := policyAPI.Delete()
	if err != nil {
		return fmt.Errorf("DELETE %s: %w: %s", policyAPI.DeleteUrl(), err, res.Raw)
	}
	return nil
}

func (r *policyResource) cleanupSwitchDeploy(ctx context.Context, dg *diag.Diagnostics, inData *NDFCPolicyModel) {
	if len(inData.SerialNumbers) == 0 {
		return
	}
	if err := r.switchDeploy(ctx, inData, true); err != nil {
		dg.AddError("Error Deploying Switch Cleanup", err.Error())
	}
}

func (r *policyResource) switchDeploy(ctx context.Context, inData *NDFCPolicyModel, recalculate bool) error {
	opts := &deployment.DeployOptions{
		SerialNumbers: inData.SerialNumbers,
	}
	localDg := diag.Diagnostics{}
	result := deployment.ConfigSaveAndDeployWithOpts(
		ctx, r.manageClient.ApiClient, inData.FabricName,
		recalculate, true, opts, &localDg,
	)
	if localDg.HasError() {
		var parts []string
		for _, d := range localDg.Errors() {
			parts = append(parts, d.Detail())
		}
		return fmt.Errorf("switch deploy failed for fabric %s: %s", inData.FabricName, strings.Join(parts, "; "))
	}
	tflog.Debug(ctx, "Switch deploy completed", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"serials":     inData.SerialNumbers,
		"status":      result.Status,
	})
	return nil
}

func buildPolicyWritePayload(in *NDFCPolicyModel, wrap bool) ([]byte, error) {
	obj := map[string]any{
		"templateName": in.TemplateName,
		"entityType":   in.EntityType,
		"entityName":   in.EntityName,
		"source":       in.Source,
	}
	if in.IsPolicyGroup {
		obj["switchIds"] = in.SerialNumbers
	} else {
		obj["switchId"] = in.SerialNumbers[0]
	}
	if in.Description != "" {
		obj["description"] = in.Description
	}
	if in.Priority != nil {
		obj["priority"] = *in.Priority
	}
	if len(in.PolicyParameters) > 0 {
		obj["templateInputs"] = in.PolicyParameters
	}
	if in.SecondaryEntityName != "" {
		obj["secondaryEntityName"] = in.SecondaryEntityName
	}
	if in.SecondaryEntityType != "" {
		obj["secondaryEntityType"] = in.SecondaryEntityType
	}

	if !wrap {
		return json.Marshal(obj)
	}
	if in.IsPolicyGroup {
		return json.Marshal(map[string]any{"policyGroups": []any{obj}})
	}
	return json.Marshal(map[string]any{"policies": []any{obj}})
}

func parseCreateResponse(res gjson.Result, isGroup bool) (string, error) {
	key := "policies"
	if isGroup {
		key = "policyGroups"
	}

	items := res.Get(key)
	if !items.IsArray() || len(items.Array()) == 0 {
		policyID := res.Get("policyId").String()
		status := res.Get("status").String()
		if status != "" && !strings.EqualFold(status, "success") {
			return "", fmt.Errorf("create %s failed: status=%s message=%s", key, status, res.Get("message").String())
		}
		if policyID == "" {
			return "", fmt.Errorf("create %s returned no policyId: %s", key, res.Raw)
		}
		return policyID, nil
	}

	item := items.Array()[0]
	status := item.Get("status").String()
	if status != "" && !strings.EqualFold(status, "success") {
		return "", fmt.Errorf("create %s failed: status=%s message=%s", key, status, item.Get("message").String())
	}
	policyID := item.Get("policyId").String()
	if policyID == "" {
		return "", fmt.Errorf("create %s returned an empty policyId: %s", key, res.Raw)
	}
	return policyID, nil
}

func checkBulkStatus(res gjson.Result, policyID string) error {
	keys := []string{"policies", "policyGroups"}
	found := false
	for _, key := range keys {
		items := res.Get(key)
		if !items.IsArray() {
			continue
		}
		found = true
		for _, item := range items.Array() {
			status := item.Get("status").String()
			if status != "" && !strings.EqualFold(status, "success") {
				id := item.Get("policyId").String()
				if id == "" {
					id = policyID
				}
				return fmt.Errorf("policy action failed for %s: status=%s message=%s", id, status, item.Get("message").String())
			}
		}
	}
	if found {
		return nil
	}
	if res.IsArray() {
		for _, item := range res.Array() {
			status := item.Get("status").String()
			if status != "" && !strings.EqualFold(status, "success") {
				return fmt.Errorf("policy action failed for %s: status=%s message=%s", policyID, status, item.Get("message").String())
			}
		}
		return nil
	}
	status := res.Get("status").String()
	if status != "" && !strings.EqualFold(status, "success") {
		return fmt.Errorf("policy action failed for %s: status=%s message=%s", policyID, status, res.Get("message").String())
	}
	return nil
}

func policyObjectToModel(obj *policyAPIObject, inData *NDFCPolicyModel) *NDFCPolicyModel {
	out := &NDFCPolicyModel{
		PolicyId:            obj.PolicyId,
		FabricName:          inData.FabricName,
		IsPolicyGroup:       inData.IsPolicyGroup,
		Deploy:              inData.Deploy,
		EntityName:          obj.EntityName,
		EntityType:          obj.EntityType,
		Description:         obj.Description,
		TemplateName:        obj.TemplateName,
		Source:              obj.Source,
		Priority:            obj.Priority,
		SecondaryEntityName: obj.SecondaryEntityName,
		SecondaryEntityType: obj.SecondaryEntityType,
		GeneratedConfig:     obj.GeneratedConfig,
		Deleted:             obj.MarkDeleted,
	}
	if out.PolicyId == "" {
		out.PolicyId = inData.PolicyId
	}
	out.SerialNumbers = serialsFromPolicyObject(obj)
	if len(out.SerialNumbers) == 0 {
		out.SerialNumbers = inData.SerialNumbers
	}
	splitTemplateInputs(inData.PolicyParameters, stringifyTemplateInputs(obj.TemplateInputs), out)
	if out.Source == "" {
		out.Source = inData.Source
	}
	if out.Priority == nil {
		out.Priority = inData.Priority
	}
	if out.Description == "" {
		out.Description = inData.Description
	}
	if out.SecondaryEntityName == "" {
		out.SecondaryEntityName = inData.SecondaryEntityName
	}
	if out.SecondaryEntityType == "" {
		out.SecondaryEntityType = inData.SecondaryEntityType
	}
	return out
}

func serialsFromPolicyObject(obj *policyAPIObject) []string {
	if len(obj.SwitchIds) > 0 {
		return obj.SwitchIds
	}
	if len(obj.SwitchId) == 0 || string(obj.SwitchId) == "null" {
		return nil
	}
	var serial string
	if err := json.Unmarshal(obj.SwitchId, &serial); err == nil && serial != "" {
		return []string{serial}
	}
	var serials []string
	if err := json.Unmarshal(obj.SwitchId, &serials); err == nil {
		return serials
	}
	return nil
}

func stringifyTemplateInputs(in map[string]any) map[string]string {
	if len(in) == 0 {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = anyToString(v)
	}
	return out
}

func anyToString(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return t
	case bool:
		return strconv.FormatBool(t)
	case float64:
		if t == float64(int64(t)) {
			return strconv.FormatInt(int64(t), 10)
		}
		return strconv.FormatFloat(t, 'f', -1, 64)
	case json.Number:
		return t.String()
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return fmt.Sprint(t)
		}
		return string(b)
	}
}

var controllerTemplateInputKeys = map[string]bool{
	"FABRIC_NAME":   true,
	"SERIAL_NUMBER": true,
	"POLICY_ID":     true,
}

func splitTemplateInputs(planned, got map[string]string, out *NDFCPolicyModel) {
	out.PolicyParameters = make(map[string]string)

	if len(planned) == 0 {
		for k, v := range got {
			if controllerTemplateInputKeys[k] {
				continue
			}
			out.PolicyParameters[k] = v
		}
		return
	}

	for k, v := range got {
		if _, ok := planned[k]; ok {
			out.PolicyParameters[k] = v
		}
	}
	for k, v := range planned {
		if _, ok := out.PolicyParameters[k]; !ok {
			out.PolicyParameters[k] = v
		}
	}
}

func applyPolicyState(dg *diag.Diagnostics, model *PolicyModel, outData *NDFCPolicyModel) {
	keepSource := model.Source
	if modelDiags := model.SetModelData(outData); modelDiags.HasError() {
		dg.Append(modelDiags...)
		return
	}
	setPolicyIDs(model, outData.PolicyId)
	if !keepSource.IsNull() && !keepSource.IsUnknown() && keepSource.ValueString() == "" &&
		(model.Source.IsNull() || model.Source.ValueString() == "") {
		model.Source = types.StringValue("")
	}
}

func setPolicyIDs(model *PolicyModel, policyID string) {
	if model == nil || strings.TrimSpace(policyID) == "" {
		return
	}
	model.PolicyId = types.StringValue(policyID)
}

func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "statuscode 404") ||
		strings.Contains(msg, "does not exist") ||
		strings.Contains(msg, "not found")
}
