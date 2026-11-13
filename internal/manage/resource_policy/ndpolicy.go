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

func (r *policyResource) newPolicyAPI(in *NDFCPolicyModel) *api.PolicyAPI {
	policyAPI := api.NewPolicyAPI(r.manageClient.ApiClient, in.FabricName)
	policyAPI.FabricName = in.FabricName
	policyAPI.PolicyID = in.Id
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
	inData.Id = policyID

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
		dg.AddError("Error Reading Policy", fmt.Sprintf("Could not read policy %s: %v", inData.Id, err))
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
		"policy_id":   inData.Id,
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
	if strings.TrimSpace(inData.Id) == "" {
		dg.AddError("Error Deleting Policy", "Policy ID is missing from state.")
		return
	}
	if err := r.ensurePolicyFabricName(ctx, inData); err != nil {
		dg.AddError("Error Deleting Policy", err.Error())
		return
	}

	serials := r.resolveSwitchIds(ctx, dg, inData.FabricName, inData.SwitchIds)
	if dg.HasError() {
		return
	}
	inData.SwitchIds = serials

	policyAPI := r.newPolicyAPI(inData)
	markErr := r.markDeletePolicy(ctx, policyAPI, inData.Id)
	if markErr != nil {
		tflog.Warn(ctx, "Mark-delete failed; falling back to direct delete then switch deploy", map[string]interface{}{
			"policy_id": inData.Id,
			"error":     markErr.Error(),
		})
		if err := r.deletePolicyRecord(ctx, policyAPI); err != nil && !isNotFoundError(err) {
			dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not delete policy %s: %v", inData.Id, err))
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
		dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not verify policy %s after cleanup: %v", inData.Id, err))
		return
	}
	if outData == nil {
		return
	}

	if err := r.deletePolicyRecord(ctx, policyAPI); err != nil && !isNotFoundError(err) {
		dg.AddError("Error Deleting Policy", fmt.Sprintf("Could not delete remaining policy %s: %v", inData.Id, err))
	}
}

func (r *policyResource) preparePolicyWrite(ctx context.Context, dg *diag.Diagnostics, inData *NDFCPolicyModel) bool {
	if !r.resolveSwitchesAndFabric(ctx, dg, inData) {
		return false
	}

	if inData.IsPolicyGroup {
		if len(inData.SwitchIds) == 0 {
			dg.AddAttributeError(
				path.Root("switch_ids"),
				"Invalid Configuration",
				"switch_ids must contain at least one switch when is_policy_group is true.",
			)
			return false
		}
	} else if len(inData.SwitchIds) != 1 {
		dg.AddAttributeError(
			path.Root("switch_ids"),
			"Invalid Configuration",
			"switch_ids must contain exactly one switch when is_policy_group is false.",
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

	fabric, resolved, err := r.resolvePolicySwitches(ctx, inData.SwitchIds)
	if err != nil {
		dg.AddError("Unable to Resolve Fabric", err.Error())
		return false
	}
	inData.SwitchIds = resolved
	if inData.IsPolicyGroup {
		inData.SwitchId = ""
	} else {
		inData.SwitchId = resolved[0]
	}
	inData.FabricName = fabric
	tflog.Debug(ctx, "Resolved policy fabric from switch inventory", map[string]interface{}{
		"fabric_name":     fabric,
		"switch_ids":      resolved,
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
	if strings.TrimSpace(inData.Id) == "" && !model.Id.IsNull() && !model.Id.IsUnknown() {
		inData.Id = model.Id.ValueString()
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

	fabric, err := r.resolvePolicyFabricName(ctx, inData.SwitchIds)
	if err != nil {
		return err
	}
	inData.FabricName = fabric
	return nil
}

func (r *policyResource) resolvePolicyFabricName(ctx context.Context, switchIDs []string) (string, error) {
	fabricName, _, err := r.resolvePolicySwitches(ctx, switchIDs)
	return fabricName, err
}

func (r *policyResource) resolvePolicySwitches(ctx context.Context, switchIDs []string) (string, []string, error) {
	if len(switchIDs) == 0 {
		return "", nil, fmt.Errorf("switch_ids is empty; cannot resolve fabric_name")
	}
	var fabricName string
	resolved := make([]string, 0, len(switchIDs))
	seen := make(map[string]bool, len(switchIDs))
	for _, id := range switchIDs {
		entry, ok := r.manageClient.SwitchDB.GetSwitchByID(ctx, id)
		if !ok {
			return "", nil, fmt.Errorf("switch %s not found in global inventory", id)
		}
		switchFabric := strings.TrimSpace(entry.FabricName)
		if switchFabric == "" {
			return "", nil, fmt.Errorf("switch %s has no fabric association in global inventory", id)
		}
		if fabricName == "" {
			fabricName = switchFabric
		} else if fabricName != switchFabric {
			return "", nil, fmt.Errorf("switches belong to different fabrics: expected %s, %s=%s", fabricName, id, switchFabric)
		}
		if entry.SerialNumber == "" {
			return "", nil, fmt.Errorf("global inventory entry for switch %s has no serial number", id)
		}
		if !seen[entry.SerialNumber] {
			seen[entry.SerialNumber] = true
			resolved = append(resolved, entry.SerialNumber)
		}
	}

	tflog.Debug(ctx, "Resolved policy fabric from inventory switches", map[string]interface{}{
		"switch_ids":     switchIDs,
		"serial_numbers": resolved,
		"fabric_name":    fabricName,
	})
	return fabricName, resolved, nil
}

func (r *policyResource) resolveSwitchIds(ctx context.Context, dg *diag.Diagnostics, fabricName string, ids []string) []string {
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
					path.Root("switch_ids"),
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
		dg.AddError("Error Reading Policy", fmt.Sprintf("Could not read policy %s after apply: %v", inData.Id, err))
		return false
	}
	applyPolicyState(dg, model, outData)
	return true
}

func (r *policyResource) fetchPolicy(ctx context.Context, inData *NDFCPolicyModel) (*NDFCPolicyModel, error) {
	policyAPI := r.newPolicyAPI(inData)
	tflog.Debug(ctx, "Fetching policy from ND", map[string]interface{}{
		"fabric_name":     inData.FabricName,
		"policy_id":       inData.Id,
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

func decodePolicyObject(resp []byte) (*NDFCPolicyModel, error) {
	var obj NDFCPolicyModel
	if err := json.Unmarshal(resp, &obj); err != nil {
		return nil, fmt.Errorf("decode policy response: %w", err)
	}
	if obj.Id != "" || obj.TemplateName != "" {
		return &obj, nil
	}

	var wrapped struct {
		Policies     []NDFCPolicyModel `json:"policies"`
		PolicyGroups []NDFCPolicyModel `json:"policyGroups"`
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
			Id:            policyID,
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
	payload, err := json.Marshal(map[string][]string{"policyIds": {inData.Id}})
	if err != nil {
		return fmt.Errorf("marshal pushConfig payload: %w", err)
	}

	tflog.Debug(ctx, "Pushing policy config", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"policy_id":   inData.Id,
		"url":         policyAPI.PostUrl(),
	})

	res, err := policyAPI.Post(payload, nil)
	if err != nil {
		return fmt.Errorf("pushConfig failed for %s: %w: %s", inData.Id, err, res.Raw)
	}
	if err := checkBulkStatus(res, inData.Id); err != nil {
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
	if len(inData.SwitchIds) == 0 {
		return
	}
	if err := r.switchDeploy(ctx, inData, true); err != nil {
		dg.AddError("Error Deploying Switch Cleanup", err.Error())
	}
}

func (r *policyResource) switchDeploy(ctx context.Context, inData *NDFCPolicyModel, recalculate bool) error {
	opts := &deployment.DeployOptions{
		SerialNumbers: inData.SwitchIds,
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
		"serials":     inData.SwitchIds,
		"status":      result.Status,
	})
	return nil
}

func buildPolicyWritePayload(in *NDFCPolicyModel, wrap bool) ([]byte, error) {
	if in == nil {
		return nil, fmt.Errorf("policy model is nil")
	}
	apiModel := *in
	apiModel.Id = ""
	if in.IsPolicyGroup {
		apiModel.SwitchId = ""
		apiModel.SwitchIds = append([]string(nil), in.SwitchIds...)
	} else {
		if len(in.SwitchIds) != 1 {
			return nil, fmt.Errorf("single-switch policy requires exactly one switch ID; got %d", len(in.SwitchIds))
		}
		apiModel.SwitchId = in.SwitchIds[0]
		apiModel.SwitchIds = nil
	}
	if !wrap {
		return json.Marshal(apiModel)
	}
	envelope := struct {
		Policies     []NDFCPolicyModel `json:"policies,omitempty"`
		PolicyGroups []NDFCPolicyModel `json:"policyGroups,omitempty"`
	}{}
	if in.IsPolicyGroup {
		envelope.PolicyGroups = []NDFCPolicyModel{apiModel}
	} else {
		envelope.Policies = []NDFCPolicyModel{apiModel}
	}
	return json.Marshal(envelope)
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

func policyObjectToModel(obj *NDFCPolicyModel, inData *NDFCPolicyModel) *NDFCPolicyModel {
	model := *obj
	model.IsPolicyGroup = inData.IsPolicyGroup
	model.Deploy = inData.Deploy
	model.FabricName = inData.FabricName
	out := &model
	if out.Id == "" {
		out.Id = inData.Id
	}
	out.SwitchIds = serialsFromPolicyObject(obj)
	if len(out.SwitchIds) == 0 {
		out.SwitchIds = inData.SwitchIds
	}
	splitTemplateInputs(inData.PolicyParameters, obj.PolicyParameters, out)
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

func serialsFromPolicyObject(obj *NDFCPolicyModel) []string {
	if len(obj.SwitchIds) > 0 {
		return obj.SwitchIds
	}
	if obj.SwitchId != "" {
		return []string{obj.SwitchId}
	}
	return nil
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
	setPolicyIDs(model, outData.Id)
	if !keepSource.IsNull() && !keepSource.IsUnknown() && keepSource.ValueString() == "" &&
		(model.Source.IsNull() || model.Source.ValueString() == "") {
		model.Source = types.StringValue("")
	}
}

func setPolicyIDs(model *PolicyModel, policyID string) {
	if model == nil || strings.TrimSpace(policyID) == "" {
		return
	}
	model.Id = types.StringValue(policyID)
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
