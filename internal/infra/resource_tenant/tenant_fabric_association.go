// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_tenant

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	manageapi "terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tidwall/gjson"
)

type tenantFabricAssociationPayload struct {
	Items []tenantFabricAssociationItem `json:"items"`
}

type tenantFabricAssociationListResponse struct {
	TenantFabricAssociations []tenantFabricAssociationItem `json:"tenantFabricAssociations"`
}

type tenantFabricAssociationItem struct {
	AllowedVlans []string `json:"allowedVlans,omitempty"`
	Associate    bool     `json:"associate"`
	FabricName   string   `json:"fabricName"`
	LocalName    string   `json:"localName,omitempty"`
	TenantName   string   `json:"tenantName"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
}

// normalizeTenantFabricAssociation makes association comparison deterministic.
// The manage API and Terraform config can return VLANs in different orders, so
// this sorts a copy of allowed_vlans before map lookup, diff detection, or
// payload construction.
func normalizeTenantFabricAssociation(association NDFCFabricAssociationsValue) NDFCFabricAssociationsValue {
	association.AllowedVlans = sortedStringCopy(association.AllowedVlans)
	return association
}

// sortedStringCopy returns a sorted copy of values and keeps nil as nil.
// Preserving nil lets Terraform state keep the difference between an omitted
// allowed_vlans attribute and an explicitly configured empty set.
func sortedStringCopy(values []string) []string {
	if values == nil {
		return nil
	}

	out := append([]string{}, values...)
	sort.Strings(out)
	return out
}

// allowedVlansForPayload prepares configured VLAN values for the manage API
// payload. A nil Terraform value becomes an empty slice, which is omitted from
// JSON by the request item tag so the backend can apply its default [] value.
func allowedVlansForPayload(values []string) []string {
	if values == nil {
		return []string{}
	}

	return values
}

// newTenantFabricAssociationItem maps one Terraform association into the
// manage API request item. The caller decides the operation with associate:
// true creates or updates an association, false removes it.
func newTenantFabricAssociationItem(tenantName string, association NDFCFabricAssociationsValue, associate bool) tenantFabricAssociationItem {
	association = normalizeTenantFabricAssociation(association)
	return tenantFabricAssociationItem{
		AllowedVlans: allowedVlansForPayload(association.AllowedVlans),
		Associate:    associate,
		FabricName:   association.FabricName,
		LocalName:    association.LocalName,
		TenantName:   tenantName,
		TenantPrefix: association.TenantPrefix,
	}
}

func tenantFabricAssociationsEqual(existing, desired NDFCFabricAssociationsValue) bool {
	existing = normalizeTenantFabricAssociation(existing)
	desired = normalizeTenantFabricAssociation(desired)

	return existing.FabricName == desired.FabricName &&
		existing.LocalName == desired.LocalName &&
		existing.TenantPrefix == desired.TenantPrefix &&
		slices.Equal(existing.AllowedVlans, desired.AllowedVlans)
}

func tenantFabricAssociationPayloadCounts(payload tenantFabricAssociationPayload) (int, int) {
	associateTrue := 0
	associateFalse := 0
	for _, item := range payload.Items {
		if item.Associate {
			associateTrue++
			continue
		}
		associateFalse++
	}
	return associateTrue, associateFalse
}

// tenantFabricAssociationsByFabricName indexes configured associations by
// fabric_name because a single tenant can have at most one configured
// association per fabric in Terraform. It rejects empty fabric_name values and
// duplicate fabric_name values so update/delete logic has a stable key.
func tenantFabricAssociationsByFabricName(dg *diag.Diagnostics, associations []NDFCFabricAssociationsValue) map[string]NDFCFabricAssociationsValue {
	result := make(map[string]NDFCFabricAssociationsValue, len(associations))
	for _, association := range associations {
		key := association.FabricName
		if key == "" {
			dg.AddError(
				"Invalid Tenant Fabric Association",
				"fabric_name cannot be empty.",
			)
			continue
		}

		if _, ok := result[key]; ok {
			dg.AddError(
				"Invalid Tenant Fabric Association",
				fmt.Sprintf("fabric_name %q is duplicated. Each tenant fabric association must use a unique fabric_name.", key),
			)
			continue
		}

		result[key] = normalizeTenantFabricAssociation(association)
	}
	return result
}

// tenantFabricAssociationsForModel converts the Terraform nested set into the
// internal association slice used by create and update helpers.
// It ignores null or unknown nested objects, preserves omitted allowed_vlans as
// nil, normalizes VLAN order, and validates uniqueness by fabric_name.
func tenantFabricAssociationsForModel(ctx context.Context, dg *diag.Diagnostics, model *TenantModel) []NDFCFabricAssociationsValue {
	if model == nil || model.FabricAssociations.IsNull() || model.FabricAssociations.IsUnknown() {
		return nil
	}

	var elements []FabricAssociationsValue
	dg.Append(model.FabricAssociations.ElementsAs(ctx, &elements, false)...)
	if dg.HasError() {
		return nil
	}

	associations := make([]NDFCFabricAssociationsValue, 0, len(elements))
	for _, element := range elements {
		if element.IsNull() || element.IsUnknown() {
			continue
		}

		var allowedVlans []string
		if !element.AllowedVlans.IsNull() && !element.AllowedVlans.IsUnknown() {
			dg.Append(element.AllowedVlans.ElementsAs(ctx, &allowedVlans, false)...)
			if dg.HasError() {
				return nil
			}
		}

		associations = append(associations, normalizeTenantFabricAssociation(NDFCFabricAssociationsValue{
			FabricName:   element.FabricName.ValueString(),
			LocalName:    element.LocalName.ValueString(),
			TenantPrefix: element.TenantPrefix.ValueString(),
			AllowedVlans: allowedVlans,
		}))
	}

	tenantFabricAssociationsByFabricName(dg, associations)
	if dg.HasError() {
		return nil
	}

	return associations
}

// rscPostTenantFabricAssociations sends a ready-made association payload and
// checks the per-item result array. Payload selection is intentionally done by
// create, update, or delete helpers so this function stays only responsible for
// POST execution and response failure handling.
func (r *tenantResource) rscPostTenantFabricAssociations(dg *diag.Diagnostics, payload tenantFabricAssociationPayload) {
	createCount, deleteCount := tenantFabricAssociationPayloadCounts(payload)
	log.Printf("[DEBUG] Start rscPostTenantFabricAssociations: items=%d associate_true=%d associate_false=%d", len(payload.Items), createCount, deleteCount)
	defer log.Printf("[DEBUG] End rscPostTenantFabricAssociations: items=%d associate_true=%d associate_false=%d", len(payload.Items), createCount, deleteCount)

	if len(payload.Items) == 0 {
		return
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant Fabric Associations",
			fmt.Sprintf("Could not marshal tenant fabric association payload, unexpected error: %v", err),
		)
		return
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	res, err := tenantFabricAssocAPI.Post(payloadBytes, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Tenant Fabric Associations",
			fmt.Sprintf("Could not update tenant fabric associations, unexpected error: %v %v", err, res),
		)
		return
	}

	results := res.Get("results")
	if !results.Exists() || !results.IsArray() {
		dg.AddError(
			"Error Updating Tenant Fabric Associations",
			fmt.Sprintf("Tenant fabric association response did not include a valid results array: %s", res.String()),
		)
		return
	}

	var failed []string
	results.ForEach(func(_, item gjson.Result) bool {
		if strings.EqualFold(item.Get("status").String(), "failed") {
			failed = append(failed, fmt.Sprintf("fabricName=%q message=%q", item.Get("fabricName").String(), item.Get("message").String()))
		}
		return true
	})

	if len(failed) > 0 {
		dg.AddError(
			"Error Updating Tenant Fabric Associations",
			fmt.Sprintf("Tenant fabric association request failed: %s", strings.Join(failed, "; ")),
		)
	}
}

// rscGetTenantFabricAssociations reads manage-side tenant fabric associations
// and filters them for this resource. The first filter always keeps only API
// objects whose tenantName equals tenantName. When configuredAssociations is
// provided, a second filter keeps only fabrics present in that set; passing nil
// skips the fabric filter and returns every association for the tenant. That
// all-association mode is used to hydrate authoritative nd_tenant state and to
// remove every association before tenant deletion.
func (r *tenantResource) rscGetTenantFabricAssociations(dg *diag.Diagnostics, tenantName string, configuredAssociations []NDFCFabricAssociationsValue) []NDFCFabricAssociationsValue {
	filteredCount := 0
	filterEnabled := configuredAssociations != nil
	log.Printf("[DEBUG] Start rscGetTenantFabricAssociations: tenant_name=%s fabric_filter_enabled=%t fabric_filter_count=%d", tenantName, filterEnabled, len(configuredAssociations))
	defer func() {
		log.Printf("[DEBUG] End rscGetTenantFabricAssociations: tenant_name=%s filtered_count=%d", tenantName, filteredCount)
	}()

	configuredByKey := map[string]NDFCFabricAssociationsValue(nil)
	if configuredAssociations != nil {
		configuredByKey = tenantFabricAssociationsByFabricName(dg, configuredAssociations)
		if dg.HasError() || len(configuredByKey) == 0 {
			return nil
		}
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	respData, err := tenantFabricAssocAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Reading Tenant Fabric Associations",
			fmt.Sprintf("Could not read tenant fabric associations, unexpected error: %v %v", err, string(respData)),
		)
		return nil
	}

	var associationResp tenantFabricAssociationListResponse
	err = json.Unmarshal(respData, &associationResp)
	if err != nil {
		dg.AddError(
			"Error Reading Tenant Fabric Associations",
			fmt.Sprintf("Could not unmarshal tenant fabric association response, unexpected error: %v", err),
		)
		return nil
	}

	filtered := make([]NDFCFabricAssociationsValue, 0, len(associationResp.TenantFabricAssociations))
	for _, association := range associationResp.TenantFabricAssociations {
		if association.TenantName != tenantName {
			continue
		}
		if configuredByKey != nil {
			if _, ok := configuredByKey[association.FabricName]; !ok {
				continue
			}
		}

		filtered = append(filtered, normalizeTenantFabricAssociation(NDFCFabricAssociationsValue{
			AllowedVlans: association.AllowedVlans,
			FabricName:   association.FabricName,
			LocalName:    association.LocalName,
			TenantPrefix: association.TenantPrefix,
		}))
	}

	filteredCount = len(filtered)
	return filtered
}

// rscReadTenantFabricAssociations injects all backend associations for the
// tenant into the tenant API model before SetModelData runs. nd_tenant is
// authoritative for the tenant's fabric_associations set, so normal read and
// import both preserve every backend association in state and let Terraform
// surface any out-of-band association as drift.
func (r *tenantResource) rscReadTenantFabricAssociations(dg *diag.Diagnostics, tenantResp *NDFCTenantModel, tenantName string) {
	log.Printf("[DEBUG] Start rscReadTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscReadTenantFabricAssociations: tenant_name=%s", tenantName)

	filtered := r.rscGetTenantFabricAssociations(dg, tenantName, nil)
	if dg.HasError() {
		return
	}

	tenantResp.FabricAssociations = filtered
}

// rscSyncConfiguredTenantFabricAssociations builds one mixed update payload for
// Terraform changes to fabric_associations. Deletion candidates are current
// state associations whose fabric_name is absent from the desired plan; normal
// read hydrates state with every backend association for this tenant, so
// out-of-band associations are also removed when Terraform applies the desired
// configuration. Create/update candidates are desired associations that are new
// or differ by local_name or allowed_vlans. tenant_prefix is immutable in the
// Tenant Fabric Associations API; a tenant_prefix difference for an existing
// fabric association is still detected and sent so the API returns the
// immutability error instead of the provider silently ignoring the requested
// change. To change tenant_prefix, remove the association first and then create
// it again. The backend read is filtered by tenantName and old fabric_name
// values so removals use current API values when available.
func (r *tenantResource) rscSyncConfiguredTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, oldState *TenantModel, tenantModel *TenantModel) {
	tenantName := tenantModel.Name.ValueString()
	log.Printf("[DEBUG] Start rscSyncConfiguredTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscSyncConfiguredTenantFabricAssociations: tenant_name=%s", tenantName)

	oldAssociations := tenantFabricAssociationsForModel(ctx, dg, oldState)
	if dg.HasError() {
		return
	}

	desiredAssociations := tenantFabricAssociationsForModel(ctx, dg, tenantModel)
	if dg.HasError() {
		return
	}

	oldByKey := tenantFabricAssociationsByFabricName(dg, oldAssociations)
	desiredByKey := tenantFabricAssociationsByFabricName(dg, desiredAssociations)
	if dg.HasError() {
		return
	}

	currentOldAssociationsByKey := make(map[string]NDFCFabricAssociationsValue, len(oldByKey))
	if len(oldByKey) > 0 {
		currentOldAssociations := r.rscGetTenantFabricAssociations(dg, tenantModel.Name.ValueString(), oldAssociations)
		if dg.HasError() {
			return
		}
		for _, association := range currentOldAssociations {
			currentOldAssociationsByKey[association.FabricName] = association
		}
	}

	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationItem, 0, len(oldAssociations)+len(desiredAssociations)),
	}

	for _, oldAssociation := range oldAssociations {
		oldAssociation = normalizeTenantFabricAssociation(oldAssociation)
		key := oldAssociation.FabricName
		if _, ok := desiredByKey[key]; !ok {
			removeAssociation, ok := currentOldAssociationsByKey[key]
			if !ok {
				payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, oldAssociation, false))
				continue
			}
			payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, removeAssociation, false))
		}
	}

	for _, desiredAssociation := range desiredAssociations {
		desiredAssociation = normalizeTenantFabricAssociation(desiredAssociation)
		key := desiredAssociation.FabricName
		oldAssociation, ok := oldByKey[key]
		if !ok || !tenantFabricAssociationsEqual(oldAssociation, desiredAssociation) {
			payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, desiredAssociation, true))
		}
	}

	r.rscPostTenantFabricAssociations(dg, payload)
}

// rscRestoreTenantFabricAssociations restores the complete association set
// from the previous Terraform state after a multi-item update partially fails.
// It reads the actual backend state first so rollback only reverses operations
// that were applied and recreates associations that were successfully removed.
func (r *tenantResource) rscRestoreTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, previousState *TenantModel) {
	if previousState == nil {
		dg.AddError(
			"Error Restoring Tenant Fabric Associations",
			"The previous tenant state is nil.",
		)
		return
	}

	tenantName := previousState.Name.ValueString()
	log.Printf("[DEBUG] Start rscRestoreTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscRestoreTenantFabricAssociations: tenant_name=%s", tenantName)

	desiredAssociations := tenantFabricAssociationsForModel(ctx, dg, previousState)
	if dg.HasError() {
		return
	}

	currentAssociations := r.rscGetTenantFabricAssociations(dg, tenantName, nil)
	if dg.HasError() {
		return
	}

	desiredByKey := tenantFabricAssociationsByFabricName(dg, desiredAssociations)
	currentByKey := tenantFabricAssociationsByFabricName(dg, currentAssociations)
	if dg.HasError() {
		return
	}

	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationItem, 0, len(currentAssociations)+len(desiredAssociations)),
	}

	for _, currentAssociation := range currentAssociations {
		if _, ok := desiredByKey[currentAssociation.FabricName]; !ok {
			payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, currentAssociation, false))
		}
	}

	for _, desiredAssociation := range desiredAssociations {
		currentAssociation, ok := currentByKey[desiredAssociation.FabricName]
		if !ok || !tenantFabricAssociationsEqual(currentAssociation, desiredAssociation) {
			payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, desiredAssociation, true))
		}
	}

	r.rscPostTenantFabricAssociations(dg, payload)
}

// rscDeleteTenantFabricAssociations removes every backend association for the
// tenant being deleted. It passes nil configuredAssociations, so
// rscGetTenantFabricAssociations filters only by tenantName and does not limit
// deletion to the Terraform-configured fabric_name list.
func (r *tenantResource) rscDeleteTenantFabricAssociations(dg *diag.Diagnostics, state *TenantModel) {
	tenantName := state.Name.ValueString()
	log.Printf("[DEBUG] Start rscDeleteTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscDeleteTenantFabricAssociations: tenant_name=%s", tenantName)

	associations := r.rscGetTenantFabricAssociations(dg, tenantName, nil)
	if dg.HasError() {
		return
	}

	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationItem, 0, len(associations)),
	}
	for _, association := range associations {
		payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, association, false))
	}

	r.rscPostTenantFabricAssociations(dg, payload)
}
