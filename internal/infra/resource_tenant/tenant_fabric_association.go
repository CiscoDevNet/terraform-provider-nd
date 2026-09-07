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
	"errors"
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/common/utils"
	manageapi "terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tidwall/gjson"
)

type tenantFabricAssociationPayload struct {
	Items []tenantFabricAssociationRequestItem `json:"items"`
}

type tenantFabricAssociationListResponse struct {
	TenantFabricAssociations []tenantFabricAssociationItem `json:"tenantFabricAssociations"`
}

// tenantFabricAssociationRequestItem uses pointers for optional mutable fields
// so an omitted value and an explicit clear remain distinct during JSON
// serialization. A nil pointer is omitted on create, while a non-nil pointer to
// "" or [] is included on update to clear the backend value.
type tenantFabricAssociationRequestItem struct {
	AllowedVlans *[]string `json:"allowedVlans,omitempty"`
	Associate    bool      `json:"associate"`
	FabricName   string    `json:"fabricName"`
	LocalName    *string   `json:"localName,omitempty"`
	TenantName   string    `json:"tenantName"`
	TenantPrefix string    `json:"tenantPrefix,omitempty"`
}

type tenantFabricAssociationItem struct {
	AllowedVlans []string `json:"allowedVlans,omitempty"`
	FabricName   string   `json:"fabricName"`
	LocalName    string   `json:"localName,omitempty"`
	TenantName   string   `json:"tenantName"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
}

type tenantFabricAssociationStage string

const (
	tenantFabricAssociationStageCreate                     tenantFabricAssociationStage = "create"
	tenantFabricAssociationStageRead                       tenantFabricAssociationStage = "read"
	tenantFabricAssociationStageTenantPrefixChangeDelete   tenantFabricAssociationStage = "tenant_prefix_change_delete"
	tenantFabricAssociationStageTenantPrefixDeleteRollback tenantFabricAssociationStage = "tenant_prefix_change_delete_rollback"
	tenantFabricAssociationStageOptionalValuesClearDelete  tenantFabricAssociationStage = "optional_values_clear_delete"
	tenantFabricAssociationStageOptionalValuesRollback     tenantFabricAssociationStage = "optional_values_clear_delete_rollback"
	tenantFabricAssociationStageRegularUpdate              tenantFabricAssociationStage = "regular_update"
	tenantFabricAssociationStageRegularUpdateRollback      tenantFabricAssociationStage = "regular_update_rollback"
	tenantFabricAssociationStageTenantDelete               tenantFabricAssociationStage = "tenant_delete"

	tenantFabricAssociationDeleteRetryInterval = 5 * time.Second
	tenantFabricAssociationDeleteRetryTimeout  = time.Minute
)

func tenantFabricAssociationStageMessage(stage tenantFabricAssociationStage, message string) string {
	return fmt.Sprintf("[stage=%s] %s", stage, message)
}

// normalizeTenantFabricAssociation makes association comparison deterministic.
// The manage API and Terraform config can return VLANs in different orders, so
// this sorts a copy of allowed_vlans before diff detection or payload
// construction.
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
// payload. A nil Terraform value becomes an empty slice; the request builder
// decides whether to omit it on create or include it explicitly on update.
func allowedVlansForPayload(values []string) []string {
	if values == nil {
		return []string{}
	}

	return values
}

// newTenantFabricAssociationRequestItem maps one Terraform association into a
// manage API request item. includeEmptyOptionalValues is true for updates of an
// existing association so removing local_name or allowed_vlans sends explicit
// empty values instead of omitting those fields.
func newTenantFabricAssociationRequestItem(
	tenantName string,
	fabricName string,
	association NDFCFabricAssociationsValue,
	associate bool,
	includeEmptyOptionalValues bool,
) tenantFabricAssociationRequestItem {
	association = normalizeTenantFabricAssociation(association)
	item := tenantFabricAssociationRequestItem{
		Associate:    associate,
		FabricName:   fabricName,
		TenantName:   tenantName,
		TenantPrefix: association.TenantPrefix,
	}

	if association.AllowedVlans != nil || includeEmptyOptionalValues {
		allowedVlans := allowedVlansForPayload(association.AllowedVlans)
		item.AllowedVlans = &allowedVlans
	}
	if association.LocalName != "" || includeEmptyOptionalValues {
		localName := association.LocalName
		item.LocalName = &localName
	}

	return item
}

func tenantFabricAssociationsEqual(existing, desired NDFCFabricAssociationsValue) bool {
	existing = normalizeTenantFabricAssociation(existing)
	desired = normalizeTenantFabricAssociation(desired)

	return existing.LocalName == desired.LocalName &&
		existing.TenantPrefix == desired.TenantPrefix &&
		slices.Equal(existing.AllowedVlans, desired.AllowedVlans)
}

func sortedTenantFabricAssociationKeys(associations map[string]NDFCFabricAssociationsValue) []string {
	keys := make([]string, 0, len(associations))
	for key := range associations {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func tenantPrefixChangedAssociationKeys(
	existingAssociations map[string]NDFCFabricAssociationsValue,
	desiredAssociations map[string]NDFCFabricAssociationsValue,
) []string {
	changed := make([]string, 0, len(existingAssociations))
	for key, existingAssociation := range existingAssociations {
		desiredAssociation, ok := desiredAssociations[key]
		if ok && existingAssociation.TenantPrefix != desiredAssociation.TenantPrefix {
			changed = append(changed, key)
		}
	}
	sort.Strings(changed)
	return changed
}

// optionalValuesClearedAssociationKeys returns existing associations whose
// local_name or allowed_vlans is being cleared. The backend retains the old
// value when these fields are updated to "" or [], so the association must be
// deleted before regular reconciliation recreates it with the desired defaults.
// Tenant-prefix changes are handled by their dedicated recreate stage.
func optionalValuesClearedAssociationKeys(
	existingAssociations map[string]NDFCFabricAssociationsValue,
	desiredAssociations map[string]NDFCFabricAssociationsValue,
) []string {
	changed := make([]string, 0, len(existingAssociations))
	for key, existingAssociation := range existingAssociations {
		desiredAssociation, ok := desiredAssociations[key]
		if !ok || existingAssociation.TenantPrefix != desiredAssociation.TenantPrefix {
			continue
		}

		localNameCleared := existingAssociation.LocalName != "" && desiredAssociation.LocalName == ""
		allowedVlansCleared := len(existingAssociation.AllowedVlans) > 0 && len(desiredAssociation.AllowedVlans) == 0
		if localNameCleared || allowedVlansCleared {
			changed = append(changed, key)
		}
	}
	sort.Strings(changed)
	return changed
}

// tenantFabricAssociationReconciliationPayload builds the operations required
// to make existingAssociations match desiredAssociations. removalOverrides can
// supply freshly read backend values for delete items without changing the
// comparison baseline used to detect desired create or update items.
func tenantFabricAssociationReconciliationPayload(
	tenantName string,
	existingAssociations map[string]NDFCFabricAssociationsValue,
	desiredAssociations map[string]NDFCFabricAssociationsValue,
	removalOverrides map[string]NDFCFabricAssociationsValue,
) tenantFabricAssociationPayload {
	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationRequestItem, 0, len(existingAssociations)+len(desiredAssociations)),
	}

	for _, key := range sortedTenantFabricAssociationKeys(existingAssociations) {
		existingAssociation := existingAssociations[key]
		if _, ok := desiredAssociations[key]; ok {
			continue
		}

		removeAssociation := existingAssociation
		if override, ok := removalOverrides[key]; ok {
			removeAssociation = override
		}
		payload.Items = append(payload.Items, newTenantFabricAssociationRequestItem(tenantName, key, removeAssociation, false, false))
	}

	for _, key := range sortedTenantFabricAssociationKeys(desiredAssociations) {
		desiredAssociation := desiredAssociations[key]
		existingAssociation, ok := existingAssociations[key]
		if !ok || !tenantFabricAssociationsEqual(existingAssociation, desiredAssociation) {
			payload.Items = append(payload.Items, newTenantFabricAssociationRequestItem(tenantName, key, desiredAssociation, true, ok))
		}
	}

	return payload
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

// tenantFabricAssociationsForModel converts the Terraform nested map into the
// internal association map used by create and update helpers. The Terraform map
// key is the fabric name used by the manage API.
// It ignores null or unknown nested objects, preserves omitted allowed_vlans as
// nil, normalizes VLAN order, and rejects an empty fabric name key.
func tenantFabricAssociationsForModel(ctx context.Context, dg *diag.Diagnostics, model *TenantModel) map[string]NDFCFabricAssociationsValue {
	if model == nil || model.FabricAssociations.IsNull() || model.FabricAssociations.IsUnknown() {
		return nil
	}

	var elements map[string]FabricAssociationsValue
	dg.Append(model.FabricAssociations.ElementsAs(ctx, &elements, false)...)
	if dg.HasError() {
		return nil
	}

	associations := make(map[string]NDFCFabricAssociationsValue, len(elements))
	for fabricName, element := range elements {
		if fabricName == "" {
			dg.AddError(
				"Invalid Tenant Fabric Association",
				"The fabric_associations map key cannot be empty.",
			)
			continue
		}

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

		associations[fabricName] = normalizeTenantFabricAssociation(NDFCFabricAssociationsValue{
			LocalName:    element.LocalName.ValueString(),
			TenantPrefix: element.TenantPrefix.ValueString(),
			AllowedVlans: allowedVlans,
		})
	}

	if dg.HasError() {
		return nil
	}

	return associations
}

// rscPostTenantFabricAssociations sends a ready-made association payload and
// checks the per-item result array. Tenant deletion retries only the failed
// items when every failure is the known transient orchestration/not-found race;
// all other stages and failures return immediately.
func (r *tenantResource) rscPostTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, payload tenantFabricAssociationPayload, stage tenantFabricAssociationStage) {
	createCount, deleteCount := tenantFabricAssociationPayloadCounts(payload)
	log.Printf("[DEBUG] Start rscPostTenantFabricAssociations: stage=%s items=%d associate_true=%d associate_false=%d", stage, len(payload.Items), createCount, deleteCount)
	defer log.Printf("[DEBUG] End rscPostTenantFabricAssociations: stage=%s items=%d associate_true=%d associate_false=%d", stage, len(payload.Items), createCount, deleteCount)

	if len(payload.Items) == 0 {
		return
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	pendingPayload := payload
	var failed []string
	attempt := 0
	post := func(postCtx context.Context) (bool, error) {
		attempt++
		if err := postCtx.Err(); err != nil {
			return false, err
		}

		payloadBytes, err := json.Marshal(pendingPayload)
		if err != nil {
			return true, errors.New(tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Could not marshal tenant fabric association payload, unexpected error: %s", err.Error())))
		}

		res, err := tenantFabricAssocAPI.Post(payloadBytes, nil)
		if err != nil {
			return true, errors.New(tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Could not update tenant fabric associations, unexpected error: %s %s", err.Error(), res.String())))
		}

		results := res.Get("results")
		if !results.Exists() || !results.IsArray() {
			return true, errors.New(tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Tenant fabric association response did not include a valid results array: %s", res.String())))
		}

		failed = failed[:0]
		allFailuresRetryable := stage == tenantFabricAssociationStageTenantDelete
		var retryFabricNames map[string]struct{}
		if allFailuresRetryable {
			retryFabricNames = make(map[string]struct{})
		}
		results.ForEach(func(_, item gjson.Result) bool {
			status := item.Get("status").String()
			if !strings.EqualFold(status, "failed") {
				return true
			}

			fabricName := item.Get("fabricName").String()
			message := item.Get("message").String()
			failed = append(failed, tenantFabricAssociationStageMessage(stage, fmt.Sprintf("fabricName=%q status=%q message=%q", fabricName, status, message)))
			if !allFailuresRetryable {
				return true
			}
			if fabricName == "" || !isTransientTenantFabricAssociationDeleteFailure(message) {
				allFailuresRetryable = false
				return true
			}
			retryFabricNames[fabricName] = struct{}{}
			return true
		})

		if len(failed) == 0 || !allFailuresRetryable {
			return true, nil
		}

		retryItems := make([]tenantFabricAssociationRequestItem, 0, len(retryFabricNames))
		for _, item := range pendingPayload.Items {
			if _, ok := retryFabricNames[item.FabricName]; ok {
				retryItems = append(retryItems, item)
			}
		}
		if len(retryItems) != len(retryFabricNames) {
			return true, nil
		}

		pendingPayload.Items = retryItems
		log.Printf("[WARN] Tenant fabric association deletion is still in use by orchestration; retrying: stage=%s attempt=%d pending_items=%d", stage, attempt, len(retryItems))
		return false, nil
	}

	var pollErr error
	if stage == tenantFabricAssociationStageTenantDelete {
		pollErr = utils.PollUntil(ctx, tenantFabricAssociationDeleteRetryInterval, tenantFabricAssociationDeleteRetryTimeout, post)
	} else {
		_, pollErr = post(ctx)
	}

	if pollErr != nil {
		detail := pollErr.Error()
		switch {
		case errors.Is(pollErr, utils.ErrPollTimeout):
			detail = fmt.Sprintf("Timed out after %s retrying transient tenant fabric association deletion failures: %s", tenantFabricAssociationDeleteRetryTimeout, strings.Join(failed, "; "))
		case errors.Is(pollErr, context.Canceled), errors.Is(pollErr, context.DeadlineExceeded):
			if stage == tenantFabricAssociationStageTenantDelete {
				detail = tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Context ended while retrying tenant fabric association deletion: %s", pollErr.Error()))
			} else {
				detail = tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Context ended before the tenant fabric association request completed: %s", pollErr.Error()))
			}
		}
		dg.AddError("Error Updating Tenant Fabric Associations", detail)
		return
	}

	if len(failed) > 0 {
		dg.AddError(
			"Error Updating Tenant Fabric Associations",
			fmt.Sprintf("Tenant fabric association request failed: %s", strings.Join(failed, "; ")),
		)
	}
}

func isTransientTenantFabricAssociationDeleteFailure(message string) bool {
	normalized := strings.ToLower(strings.Join(strings.Fields(message), " "))
	return strings.Contains(normalized, "cannot be deleted because it is in use by the orchestration service") &&
		strings.Contains(normalized, "tenant ") &&
		strings.Contains(normalized, " not found")
}

// rscGetTenantFabricAssociations reads manage-side tenant fabric associations
// and filters them for this resource. The first filter always keeps only API
// objects whose tenantName equals tenantName. When configuredAssociations is
// provided, a second filter keeps only fabrics present in that map; passing nil
// skips the fabric filter and returns every association for the tenant. That
// all-association mode is used to hydrate authoritative nd_tenant state and to
// remove every association before tenant deletion.
func (r *tenantResource) rscGetTenantFabricAssociations(dg *diag.Diagnostics, tenantName string, configuredAssociations map[string]NDFCFabricAssociationsValue, stage tenantFabricAssociationStage) map[string]NDFCFabricAssociationsValue {
	filteredCount := 0
	filterEnabled := configuredAssociations != nil
	log.Printf("[DEBUG] Start rscGetTenantFabricAssociations: stage=%s tenant_name=%s fabric_filter_enabled=%t fabric_filter_count=%d", stage, tenantName, filterEnabled, len(configuredAssociations))
	defer func() {
		log.Printf("[DEBUG] End rscGetTenantFabricAssociations: stage=%s tenant_name=%s filtered_count=%d", stage, tenantName, filteredCount)
	}()

	if configuredAssociations != nil {
		if len(configuredAssociations) == 0 {
			return nil
		}
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)
	respData, err := tenantFabricAssocAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Reading Tenant Fabric Associations",
			tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Could not read tenant fabric associations, unexpected error: %v %v", err, string(respData))),
		)
		return nil
	}

	var associationResp tenantFabricAssociationListResponse
	err = json.Unmarshal(respData, &associationResp)
	if err != nil {
		dg.AddError(
			"Error Reading Tenant Fabric Associations",
			tenantFabricAssociationStageMessage(stage, fmt.Sprintf("Could not unmarshal tenant fabric association response, unexpected error: %v", err)),
		)
		return nil
	}

	filtered := make(map[string]NDFCFabricAssociationsValue, len(associationResp.TenantFabricAssociations))
	for _, association := range associationResp.TenantFabricAssociations {
		if association.TenantName != tenantName {
			continue
		}
		if configuredAssociations != nil {
			if _, ok := configuredAssociations[association.FabricName]; !ok {
				continue
			}
		}

		filtered[association.FabricName] = normalizeTenantFabricAssociation(NDFCFabricAssociationsValue{
			AllowedVlans: association.AllowedVlans,
			LocalName:    association.LocalName,
			TenantPrefix: association.TenantPrefix,
		})
	}

	filteredCount = len(filtered)
	return filtered
}

// rscReadTenantFabricAssociations injects all backend associations for the
// tenant into the tenant API model before SetModelData runs. nd_tenant is
// authoritative for the tenant's fabric_associations map, so normal read and
// import both preserve every backend association in state and let Terraform
// surface any out-of-band association as drift.
func (r *tenantResource) rscReadTenantFabricAssociations(dg *diag.Diagnostics, tenantResp *NDFCTenantModel, tenantName string) {
	log.Printf("[DEBUG] Start rscReadTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscReadTenantFabricAssociations: tenant_name=%s", tenantName)

	filtered := r.rscGetTenantFabricAssociations(dg, tenantName, nil, tenantFabricAssociationStageRead)
	if dg.HasError() {
		return
	}

	tenantResp.FabricAssociations = filtered
}

// rscDeleteTenantFabricAssociationsForRecreate deletes the selected existing
// associations before regular reconciliation recreates them. Any delete error
// can represent a partial multi-status operation, so the complete previous
// association state is restored before returning the original diagnostics.
func (r *tenantResource) rscDeleteTenantFabricAssociationsForRecreate(
	ctx context.Context,
	dg *diag.Diagnostics,
	oldState *TenantModel,
	tenantName string,
	currentAssociations map[string]NDFCFabricAssociationsValue,
	associationKeys []string,
	deleteStage tenantFabricAssociationStage,
	rollbackStage tenantFabricAssociationStage,
) bool {
	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationRequestItem, 0, len(associationKeys)),
	}
	for _, fabricName := range associationKeys {
		currentAssociation, ok := currentAssociations[fabricName]
		if !ok {
			continue
		}
		payload.Items = append(payload.Items, newTenantFabricAssociationRequestItem(tenantName, fabricName, currentAssociation, false, false))
	}

	var deleteDiags diag.Diagnostics
	r.rscPostTenantFabricAssociations(ctx, &deleteDiags, payload, deleteStage)
	if !deleteDiags.HasError() {
		return true
	}

	log.Printf("[ERROR] Tenant association recreate deletion failed for tenant id=%s stage=%s; restoring previous associations", tenantName, deleteStage)

	var rollbackDiags diag.Diagnostics
	r.rscRestoreTenantFabricAssociations(ctx, &rollbackDiags, oldState, rollbackStage)

	dg.Append(deleteDiags...)
	if rollbackDiags.HasError() {
		dg.Append(rollbackDiags...)
	} else {
		log.Printf("[INFO] Restored tenant fabric associations after recreate deletion failure: id=%s stage=%s", tenantName, deleteStage)
	}
	return false
}

// rscSyncConfiguredTenantFabricAssociations builds one mixed update payload for
// Terraform changes to fabric_associations. Deletion candidates are current
// state associations whose fabric_name is absent from the desired plan; normal
// read hydrates state with every backend association for this tenant, so
// out-of-band associations are also removed when Terraform applies the desired
// configuration. Create/update candidates are desired associations that are new
// or differ by local_name or allowed_vlans. Associations are deleted before
// regular reconciliation recreates them when tenant_prefix changes, or when
// local_name/allowed_vlans is cleared because the backend retains those values
// on an in-place empty update. A partially failed recreate-delete request is
// rolled back immediately. The backend read is filtered by tenantName and old
// fabric_name values so removals use current API values when available. The
// return value is true only when the regular update may have partially applied
// and the caller must restore the previous complete association state.
func (r *tenantResource) rscSyncConfiguredTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, oldState *TenantModel, tenantModel *TenantModel) bool {
	tenantName := tenantModel.Name.ValueString()
	log.Printf("[DEBUG] Start rscSyncConfiguredTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscSyncConfiguredTenantFabricAssociations: tenant_name=%s", tenantName)

	oldAssociations := tenantFabricAssociationsForModel(ctx, dg, oldState)
	if dg.HasError() {
		return false
	}

	desiredAssociations := tenantFabricAssociationsForModel(ctx, dg, tenantModel)
	if dg.HasError() {
		return false
	}
	prefixChangedKeys := tenantPrefixChangedAssociationKeys(oldAssociations, desiredAssociations)
	optionalValuesClearedKeys := optionalValuesClearedAssociationKeys(oldAssociations, desiredAssociations)

	currentOldAssociationsByKey := make(map[string]NDFCFabricAssociationsValue, len(oldAssociations))
	if len(oldAssociations) > 0 {
		readStage := tenantFabricAssociationStageRegularUpdate
		if len(prefixChangedKeys) > 0 {
			readStage = tenantFabricAssociationStageTenantPrefixChangeDelete
		} else if len(optionalValuesClearedKeys) > 0 {
			readStage = tenantFabricAssociationStageOptionalValuesClearDelete
		}
		currentOldAssociationsByKey = r.rscGetTenantFabricAssociations(dg, tenantName, oldAssociations, readStage)
		if dg.HasError() {
			return false
		}
	}

	if !r.rscDeleteTenantFabricAssociationsForRecreate(
		ctx,
		dg,
		oldState,
		tenantName,
		currentOldAssociationsByKey,
		prefixChangedKeys,
		tenantFabricAssociationStageTenantPrefixChangeDelete,
		tenantFabricAssociationStageTenantPrefixDeleteRollback,
	) {
		return false
	}

	if !r.rscDeleteTenantFabricAssociationsForRecreate(
		ctx,
		dg,
		oldState,
		tenantName,
		currentOldAssociationsByKey,
		optionalValuesClearedKeys,
		tenantFabricAssociationStageOptionalValuesClearDelete,
		tenantFabricAssociationStageOptionalValuesRollback,
	) {
		return false
	}

	payload := tenantFabricAssociationReconciliationPayload(
		tenantName,
		oldAssociations,
		desiredAssociations,
		currentOldAssociationsByKey,
	)

	r.rscPostTenantFabricAssociations(ctx, dg, payload, tenantFabricAssociationStageRegularUpdate)
	return dg.HasError()
}

// rscRestoreTenantFabricAssociations restores the complete association set
// from the previous Terraform state after a multi-item update partially fails.
// It reads the actual backend state first so rollback only reverses operations
// that were applied and recreates associations that were successfully removed.
func (r *tenantResource) rscRestoreTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, previousState *TenantModel, stage tenantFabricAssociationStage) {
	if previousState == nil {
		dg.AddError(
			"Error Restoring Tenant Fabric Associations",
			tenantFabricAssociationStageMessage(stage, "The previous tenant state is nil."),
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

	currentAssociations := r.rscGetTenantFabricAssociations(dg, tenantName, nil, stage)
	if dg.HasError() {
		return
	}

	payload := tenantFabricAssociationReconciliationPayload(
		tenantName,
		currentAssociations,
		desiredAssociations,
		nil,
	)

	r.rscPostTenantFabricAssociations(ctx, dg, payload, stage)
}

// rscDeleteTenantFabricAssociations removes every backend association for the
// tenant being deleted. It passes nil configuredAssociations, so
// rscGetTenantFabricAssociations filters only by tenantName and does not limit
// deletion to the Terraform-configured fabric_name list.
func (r *tenantResource) rscDeleteTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, tenantName string) {
	log.Printf("[DEBUG] Start rscDeleteTenantFabricAssociations: tenant_name=%s", tenantName)
	defer log.Printf("[DEBUG] End rscDeleteTenantFabricAssociations: tenant_name=%s", tenantName)

	associations := r.rscGetTenantFabricAssociations(dg, tenantName, nil, tenantFabricAssociationStageTenantDelete)
	if dg.HasError() {
		return
	}

	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationRequestItem, 0, len(associations)),
	}
	for _, fabricName := range sortedTenantFabricAssociationKeys(associations) {
		payload.Items = append(payload.Items, newTenantFabricAssociationRequestItem(tenantName, fabricName, associations[fabricName], false, false))
	}

	r.rscPostTenantFabricAssociations(ctx, dg, payload, tenantFabricAssociationStageTenantDelete)
}
