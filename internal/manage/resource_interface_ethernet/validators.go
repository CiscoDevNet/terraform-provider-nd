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
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ethernetDiscriminatorMatrix encodes the (mode × network_os_type → legal
// policy_type set) matrix from `interface_ethernet.spec` §Discriminator
// matrix. The sentinel "userDefined" is universally allowed and handled
// outside the matrix.
var ethernetDiscriminatorMatrix = map[string]map[string][]string{
	"access": {
		"ios-xe": {"iosXeAccess"},
		"nx-os":  {"accessHost", "ipfmAccessHost"},
	},
	"dot1qTunnel": {
		"nx-os": {"dot1qTunnelHost"},
	},
	"managed": {
		"ios-xe": {"iosXeStackwiseDualActive", "iosXeStackwiseLinkMember"},
	},
	"pvlan": {
		"nx-os": {"pvlanHost"},
	},
	"routed": {
		"ios-xe": {
			"csrMultisiteIfcMember",
			"iosXeInternalL3PoMember",
			"iosXeNumbered",
			"iosXeRoutedHost",
		},
		"ios-xr": {"iosXrRoutedHost"},
		"nx-os": {
			"dataBrokerL3Host",
			"endPointLocator",
			"ipfmL3Port",
			"routedHost",
		},
	},
	"trunk": {
		"ios-xe": {"iosXeTrunkHost"},
		"nx-os": {
			"classicHost",
			"dataBrokerHost",
			"dataBrokerPoMember",
			"ipfmTrunkHost",
			"trunkHost",
		},
	},
	"unmanaged": {
		"ios-xe": {"iosXeMonitor"},
		"nx-os":  {"monitor"},
	},
}

// ethernetConfigValidator checks:
//   - switch_id mutual exclusion: resource-level and per-entry switch_id
//     cannot both be configured; exactly one level must supply the switch
//   - (mode, network_os_type) pair is supported for ethernet
//   - policy_type is legal for that pair (or == "userDefined")
//   - template_name is set iff policy_type == "userDefined"
type ethernetConfigValidator struct{}

func (v ethernetConfigValidator) Description(_ context.Context) string {
	return "Validate (mode, network_os_type, policy_type) tuples against the ethernet discriminator matrix and enforce the userDefined / template_name contract."
}

func (v ethernetConfigValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ethernetConfigValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var cfg InterfaceEthernetModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() || cfg.Interfaces.IsNull() || cfg.Interfaces.IsUnknown() {
		return
	}

	entries := make(map[string]InterfacesValue, len(cfg.Interfaces.Elements()))
	resp.Diagnostics.Append(cfg.Interfaces.ElementsAs(ctx, &entries, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ---- switch_id mutual exclusion ----
	// Resource-level and per-entry switch_id are mutually exclusive.
	// The user must supply switch identity at exactly one level.
	hasResourceSwitch := !cfg.SwitchId.IsNull() && !cfg.SwitchId.IsUnknown()
	var entriesWithSwitch []string
	for name, e := range entries {
		if !e.SwitchId.IsNull() && !e.SwitchId.IsUnknown() {
			entriesWithSwitch = append(entriesWithSwitch, name)
		}
	}
	sort.Strings(entriesWithSwitch) // deterministic error messages

	if hasResourceSwitch && len(entriesWithSwitch) > 0 {
		resp.Diagnostics.AddAttributeError(path.Root("switch_id"),
			"switch_id is mutually exclusive between resource level and interface level",
			fmt.Sprintf("Resource-level switch_id is set but the following interfaces also "+
				"configure switch_id: %s. Set switch_id at the resource level OR on each "+
				"interface entry, not both.", strings.Join(entriesWithSwitch, ", ")))
		return
	}
	if !hasResourceSwitch {
		var missing []string
		for name, e := range entries {
			if e.SwitchId.IsNull() || e.SwitchId.IsUnknown() {
				missing = append(missing, name)
			}
		}
		if len(missing) > 0 {
			sort.Strings(missing)
			resp.Diagnostics.AddAttributeError(path.Root("switch_id"),
				"switch_id is required",
				fmt.Sprintf("Resource-level switch_id is not set and the following interfaces "+
					"are missing switch_id: %s. Either set switch_id at the resource level (applies "+
					"to all interfaces) or set it on every interface entry.", strings.Join(missing, ", ")))
			return
		}
	}

	ifPath := path.Root("interfaces")
	for name, e := range entries {
		if e.Mode.IsNull() || e.Mode.IsUnknown() ||
			e.NetworkOsType.IsNull() || e.NetworkOsType.IsUnknown() ||
			e.PolicyType.IsNull() || e.PolicyType.IsUnknown() {
			continue
		}
		mode := e.Mode.ValueString()
		os := e.NetworkOsType.ValueString()
		policy := e.PolicyType.ValueString()
		entryPath := ifPath.AtMapKey(name)

		// userDefined requires template_name; non-userDefined forbids it.
		isUser := policy == "userDefined"
		hasTemplate := !e.TemplateName.IsNull() && !e.TemplateName.IsUnknown() && e.TemplateName.ValueString() != ""
		switch {
		case isUser && !hasTemplate:
			resp.Diagnostics.AddAttributeError(entryPath.AtName("template_name"),
				"template_name is required for policy_type=userDefined",
				fmt.Sprintf("Interface %q uses policy_type=userDefined but template_name is not set.", name))
		case !isUser && hasTemplate:
			resp.Diagnostics.AddAttributeError(entryPath.AtName("template_name"),
				"template_name is only valid for policy_type=userDefined",
				fmt.Sprintf("Interface %q has policy_type=%q but template_name is set; only userDefined accepts template_name.", name, policy))
		}

		// Per-policy attribute allow-list check: reject attributes that
		// don't belong to the selected policy type.
		if !isUser {
			if allowed, known := policyAllowedAttrs[policy]; known {
				configured := configuredPolicyAttrs(e)
				for attr := range configured {
					if !allowed[attr] {
						resp.Diagnostics.AddAttributeError(entryPath.AtName(attr),
							"Attribute not valid for this policy_type",
							fmt.Sprintf("Interface %q uses policy_type=%q which does not support attribute %q.", name, policy, attr))
					}
				}
			}
		}

		if isUser {
			// userDefined is always legal — nothing further to check.
			continue
		}

		byOS, modeOk := ethernetDiscriminatorMatrix[mode]
		if !modeOk {
			resp.Diagnostics.AddAttributeError(entryPath.AtName("mode"),
				"Unsupported ethernet mode",
				fmt.Sprintf("Interface %q uses mode=%q which is not supported for ethernet. Supported modes: %s.",
					name, mode, strings.Join(sortedKeys(ethernetDiscriminatorMatrix), ", ")))
			continue
		}
		policies, osOk := byOS[os]
		if !osOk {
			resp.Diagnostics.AddAttributeError(entryPath.AtName("network_os_type"),
				"Unsupported (mode, network_os_type) combination",
				fmt.Sprintf("Interface %q uses (mode=%q, network_os_type=%q) which is not a supported ethernet combination. Supported OSes for mode=%q: %s.",
					name, mode, os, mode, strings.Join(sortedKeys(byOS), ", ")))
			continue
		}
		if !contains(policies, policy) {
			resp.Diagnostics.AddAttributeError(entryPath.AtName("policy_type"),
				"policy_type is not valid for the chosen (mode, network_os_type)",
				fmt.Sprintf("Interface %q uses policy_type=%q which is not valid for (mode=%q, network_os_type=%q). Valid built-in policy types: %s, or \"userDefined\".",
					name, policy, mode, os, strings.Join(policies, ", ")))
		}
	}
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
