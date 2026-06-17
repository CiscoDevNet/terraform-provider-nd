package resource_tenant

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"

	manageapi "terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

type tenantFabricAssociationPayload struct {
	Items []tenantFabricAssociationItem `json:"items"`
}

type tenantFabricAssociationListResponse struct {
	TenantFabricAssociations []tenantFabricAssociationItem `json:"tenantFabricAssociations"`
}

type tenantFabricAssociationItem struct {
	AllowedVlans []string `json:"allowedVlans"`
	Associate    bool     `json:"associate"`
	FabricName   string   `json:"fabricName"`
	LocalName    string   `json:"localName,omitempty"`
	TenantName   string   `json:"tenantName"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
}

func normalizeTenantFabricAssociation(association NDFCFabricAssociationsValue) NDFCFabricAssociationsValue {
	association.AllowedVlans = sortedStringCopy(association.AllowedVlans)
	return association
}

func sortedStringCopy(values []string) []string {
	if values == nil {
		return nil
	}

	out := append([]string{}, values...)
	sort.Strings(out)
	return out
}

func allowedVlansForPayload(values []string) []string {
	if values == nil {
		return []string{}
	}

	return values
}

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

func (r *tenantResource) rscPostTenantFabricAssociations(dg *diag.Diagnostics, payload tenantFabricAssociationPayload) {
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

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient)
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

func (r *tenantResource) rscGetTenantFabricAssociations(dg *diag.Diagnostics, tenantName string, configuredAssociations []NDFCFabricAssociationsValue) []NDFCFabricAssociationsValue {
	configuredByKey := map[string]NDFCFabricAssociationsValue(nil)
	if configuredAssociations != nil {
		configuredByKey = tenantFabricAssociationsByFabricName(dg, configuredAssociations)
		if dg.HasError() || len(configuredByKey) == 0 {
			return nil
		}
	}

	tenantFabricAssocAPI := manageapi.NewTenantFabricAssociationAPI(r.infraClient.ApiClient)
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

	return filtered
}

func (r *tenantResource) rscReadConfiguredTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, state *TenantModel, configuredAssociations []NDFCFabricAssociationsValue) {
	if configuredAssociations == nil {
		return
	}

	configuredByKey := tenantFabricAssociationsByFabricName(dg, configuredAssociations)
	if dg.HasError() {
		return
	}

	filtered := r.rscGetTenantFabricAssociations(dg, state.Name.ValueString(), configuredAssociations)
	if dg.HasError() {
		return
	}

	values := make([]FabricAssociationsValue, 0, len(filtered))
	for _, association := range filtered {
		association = normalizeTenantFabricAssociation(association)
		if configuredAssociation, ok := configuredByKey[association.FabricName]; ok {
			association.LocalName = configuredAssociation.LocalName
			association.TenantPrefix = configuredAssociation.TenantPrefix
			if configuredAssociation.AllowedVlans == nil {
				association.AllowedVlans = nil
			}
		}

		allowedVlanSet := types.SetNull(types.StringType)
		if association.AllowedVlans != nil {
			allowedVlans := make([]attr.Value, 0, len(association.AllowedVlans))
			for _, vlan := range association.AllowedVlans {
				allowedVlans = append(allowedVlans, types.StringValue(vlan))
			}

			var setDiags diag.Diagnostics
			allowedVlanSet, setDiags = types.SetValue(types.StringType, allowedVlans)
			dg.Append(setDiags...)
			if dg.HasError() {
				return
			}
		}

		localName := types.StringNull()
		if association.LocalName != "" {
			localName = types.StringValue(association.LocalName)
		}

		tenantPrefix := types.StringNull()
		if association.TenantPrefix != "" {
			tenantPrefix = types.StringValue(association.TenantPrefix)
		}

		values = append(values, FabricAssociationsValue{
			AllowedVlans: allowedVlanSet,
			FabricName:   types.StringValue(association.FabricName),
			LocalName:    localName,
			TenantPrefix: tenantPrefix,
			state:        attr.ValueStateKnown,
		})
	}

	associationSet, setDiags := types.SetValueFrom(ctx, FabricAssociationsValue{}.Type(ctx), values)
	dg.Append(setDiags...)
	if dg.HasError() {
		return
	}

	state.FabricAssociations = associationSet
}

func (r *tenantResource) rscSyncConfiguredTenantFabricAssociations(ctx context.Context, dg *diag.Diagnostics, oldState *TenantModel, tenantModel *TenantModel) {
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

	tenantName := tenantModel.Name.ValueString()
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
		oldAssociation = normalizeTenantFabricAssociation(oldAssociation)
		if !ok ||
			oldAssociation.FabricName != desiredAssociation.FabricName ||
			oldAssociation.LocalName != desiredAssociation.LocalName ||
			oldAssociation.TenantPrefix != desiredAssociation.TenantPrefix ||
			!slices.Equal(oldAssociation.AllowedVlans, desiredAssociation.AllowedVlans) {
			payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, desiredAssociation, true))
		}
	}

	r.rscPostTenantFabricAssociations(dg, payload)
}

func (r *tenantResource) rscDeleteTenantFabricAssociations(dg *diag.Diagnostics, state *TenantModel) {
	associations := r.rscGetTenantFabricAssociations(dg, state.Name.ValueString(), nil)
	if dg.HasError() {
		return
	}

	tenantName := state.Name.ValueString()
	payload := tenantFabricAssociationPayload{
		Items: make([]tenantFabricAssociationItem, 0, len(associations)),
	}
	for _, association := range associations {
		payload.Items = append(payload.Items, newTenantFabricAssociationItem(tenantName, association, false))
	}

	r.rscPostTenantFabricAssociations(dg, payload)
}
