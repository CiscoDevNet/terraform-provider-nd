// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-nd/internal/common/templates"
	"terraform-provider-nd/internal/manage/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// InterfaceType value fixed per spec §Discriminator matrix.
const interfaceTypeEthernet = "ethernet"

// GetFabricName returns the fabric name from the Terraform model.
func (v *InterfaceEthernetModel) GetFabricName() string {
	return v.FabricName.ValueString()
}

// GetInterfaceType returns the discriminator used for handler dispatch and
// payload interfaceType injection.
func (v *InterfaceEthernetModel) GetInterfaceType() string {
	return interfaceTypeEthernet
}

// PreMarshal injects per-entry fields that are either plumbed in from
// elsewhere in the model (switch_id defaulting from the resource-level
// field) or fixed per resource (interfaceType, interfaceName from the map
// key). The common CRUD backend calls this before any POST/PUT marshalling.
func (v *InterfaceEthernetModel) PreMarshal(_ context.Context, data *resource_interface_common.NDFCInterfaceCommonModel) {
	if data == nil {
		return
	}
	defaultSwitch := strings.TrimSpace(data.SwitchId)
	for name, entry := range data.Interfaces {
		entry.InterfaceName = name
		entry.InterfaceType = interfaceTypeEthernet
		if strings.TrimSpace(entry.SwitchId) == "" && defaultSwitch != "" {
			entry.SwitchId = defaultSwitch
		}
		data.Interfaces[name] = entry
	}
	// Build the resource id as `<fabric>/<switch_id>/<csv of switch:name pairs>`.
	// Always recomputed so it reflects the current interface set and per-entry
	// switch overrides after adds/removes.
	{
		type ifEntry struct {
			switchId string
			name     string
		}
		entries := make([]ifEntry, 0, len(data.Interfaces))
		for n, v := range data.Interfaces {
			entries = append(entries, ifEntry{switchId: v.SwitchId, name: n})
		}
		// Deterministic order: sort by name.
		for i := 1; i < len(entries); i++ {
			for j := i; j > 0 && entries[j-1].name > entries[j].name; j-- {
				entries[j-1], entries[j] = entries[j], entries[j-1]
			}
		}
		parts := make([]string, 0, len(entries))
		for _, e := range entries {
			if e.switchId != "" && e.switchId != defaultSwitch {
				parts = append(parts, e.switchId+":"+e.name)
			} else {
				parts = append(parts, e.name)
			}
		}
		data.Id = data.FabricName + "/" + defaultSwitch + "/" + strings.Join(parts, ",")
	}
}

// GetInterfacesMap extracts the typed InterfacesValue map from the TF
// model's Interfaces attribute. Returns nil if the map is null/unknown
// or conversion fails.
func (m *InterfaceEthernetModel) GetInterfacesMap(ctx context.Context) map[string]InterfacesValue {
	if m.Interfaces.IsNull() || m.Interfaces.IsUnknown() {
		return nil
	}
	entries := make(map[string]InterfacesValue, len(m.Interfaces.Elements()))
	diags := m.Interfaces.ElementsAs(ctx, &entries, false)
	if diags.HasError() {
		return nil
	}
	return entries
}

// validateAttrsAgainstTemplate queries the live ND Template Library for
// each unique policyType in the plan and warns about attributes that are
// not declared in the template's parameter list. This is an apply-time
// cross-check that catches drift between the hardcoded plan-time
// allow-list and the actual ND instance.
//
// It emits warnings (not errors) because ND is the final authority —
// the template may accept the field even if it was not in the cached
// parameter list.
func validateAttrsAgainstTemplate(ctx context.Context, dg *diag.Diagnostics, model *InterfaceEthernetModel) {
	tdb := templates.GetInstance()

	ifMap := model.GetInterfacesMap(ctx)
	for name, entry := range ifMap {
		pt := entry.PolicyType.ValueString()
		if pt == "" || pt == "userDefined" {
			continue
		}
		allowed := tdb.AllowedTFNamesForPolicy(ctx, pt)
		if allowed == nil {
			// Template not reachable or policyType unknown — skip silently;
			// the plan-time validator already checked the static allow-list.
			continue
		}
		configured := configuredPolicyAttrs(entry)
		for attr := range configured {
			if !allowed[attr] {
				dg.AddWarning(
					"Attribute may not be valid for this policy_type (live check)",
					fmt.Sprintf("Interface %q uses policy_type=%q whose ND template does not "+
						"declare attribute %q. ND may reject or ignore this value.", name, pt, attr),
				)
			}
		}
	}
}
