// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_group_vxlan_aci

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = routeServerASNValidator{}
var _ validator.String = vniRangeValidator{}

// vniRangeValidator applies the existing identifier-range rules at schema level.
type vniRangeValidator struct{}

func (vniRangeValidator) Description(_ context.Context) string {
	return "must contain decimal IDs or start-end intervals separated by commas, with identifiers between 1 and 16777214, each start no greater than its end, and no overlapping intervals"
}

func (v vniRangeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (vniRangeValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if _, err := parseIdentifierRanges(req.ConfigValue.ValueString(), 1, 16777214); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid VNI Range", err.Error())
	}
}

// routeServerASNValidator keeps ASN validation at the attribute schema level.
// ParseUint enforces the numeric limits without a range-encoding regular expression.
type routeServerASNValidator struct{}

func (routeServerASNValidator) Description(_ context.Context) string {
	return "must be an ASN from 1 through 4294967295, or dotted notation with a first component from 1 through 65535 and a second component from 0 through 65535"
}

func (v routeServerASNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v routeServerASNValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	first, second, dotted := strings.Cut(req.ConfigValue.ValueString(), ".")
	valid := validASNComponent(first, 32, false)
	if dotted {
		valid = validASNComponent(first, 16, false) && validASNComponent(second, 16, true)
	}
	if !valid {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid Route Server ASN", v.Description(ctx))
	}
}

func validASNComponent(value string, bits int, allowZero bool) bool {
	// Preserve the API pattern's rejection of padded decimal components.
	if len(value) > 1 && value[0] == '0' {
		return false
	}
	number, err := strconv.ParseUint(value, 10, bits)
	return err == nil && (allowZero || number != 0)
}

type identifierRange struct {
	first uint64
	last  uint64
}

// validateFabricGroup handles range semantics and value-dependent relationships
// that the generated attribute schema cannot express.
func validateFabricGroup(config *FabricGroupVxlanAciModel) diag.Diagnostics {
	var diags diag.Diagnostics
	ranges := make(map[string][]identifierRange)
	for _, field := range []struct {
		name  string
		value types.String
		min   uint64
		max   uint64
	}{
		{"l2_vni_range", config.L2VniRange, 1, 16777214},
		{"l3_vni_range", config.L3VniRange, 1, 16777214},
		{"security_group_tag_id_range", config.SecurityGroupTagIdRange, 16, 65535},
	} {
		if field.value.IsNull() || field.value.IsUnknown() {
			continue
		}
		parsed, err := parseIdentifierRanges(field.value.ValueString(), field.min, field.max)
		if err != nil {
			diags.AddAttributeError(path.Root(field.name), "Invalid Identifier Range", fmt.Sprintf("%s: %s", field.name, err.Error()))
		} else {
			ranges[field.name] = parsed
		}
	}
	for _, l2 := range ranges["l2_vni_range"] {
		for _, l3 := range ranges["l3_vni_range"] {
			if l2.first <= l3.last && l3.first <= l2.last {
				diags.AddAttributeError(path.Root("l3_vni_range"), "Overlapping VNI Ranges", "l2_vni_range and l3_vni_range must not overlap.")
				break
			}
		}
	}
	mode := config.MultisiteOverlayInterConnectType
	method := mode.ValueString()
	if mode.IsNull() {
		method = "directPeering" // Generated schema default; not applied to raw config.
	}
	if !config.RouteServerCollection.IsNull() && !config.RouteServerCollection.IsUnknown() {
		servers := config.RouteServerCollection.Elements()
		for i, server := range servers {
			// The generator has no nested-object validator setting for list
			// elements. Child Required fields cannot validate a null object.
			if server.IsNull() {
				diags.AddAttributeError(path.Root("route_server_collection").AtListIndex(i), "Invalid Route Server", "Each route-server entry must be a non-null object.")
			}
		}
		if !mode.IsUnknown() && method != "routeServer" && len(servers) != 0 {
			diags.AddAttributeError(path.Root("route_server_collection"), "Invalid Route Server Configuration", "route_server_collection can contain entries only when multisite_overlay_inter_connect_type is routeServer.")
		}
	}
	if !mode.IsUnknown() {
		if method == "routeServer" {
			if config.RouteServerCollection.IsNull() || (!config.RouteServerCollection.IsUnknown() && len(config.RouteServerCollection.Elements()) == 0) {
				diags.AddAttributeError(path.Root("route_server_collection"), "Missing Route Servers", "Configure at least one route server when multisite_overlay_inter_connect_type is routeServer.")
			}
		} else if config.RouteServerRedistributeDirectRouteMap.ValueBool() {
			diags.AddAttributeError(path.Root("route_server_redistribute_direct_route_map"), "Invalid Route Server Setting", "route_server_redistribute_direct_route_map can be true only when multisite_overlay_inter_connect_type is routeServer.")
		}
	}
	return diags
}

// Identifier pools accept decimal IDs and start-end intervals separated by
// commas. Check every interval, including bounds, ordering, and duplicates.
func parseIdentifierRanges(value string, minimum, maximum uint64) ([]identifierRange, error) {
	var ranges []identifierRange
	for _, part := range strings.Split(value, ",") {
		bounds := strings.Split(strings.TrimSpace(part), "-")
		if len(bounds) > 2 {
			return nil, fmt.Errorf("use decimal IDs or start-end intervals separated by commas")
		}
		first, err := strconv.ParseUint(strings.TrimSpace(bounds[0]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("use decimal IDs or start-end intervals separated by commas")
		}
		last := first
		if len(bounds) == 2 {
			last, err = strconv.ParseUint(strings.TrimSpace(bounds[1]), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("each interval must have a decimal start and end")
			}
		}
		if first < minimum || last > maximum || first > last {
			return nil, fmt.Errorf("IDs must be between %d and %d, with each start no greater than its end", minimum, maximum)
		}
		for _, existing := range ranges {
			if first <= existing.last && existing.first <= last {
				return nil, fmt.Errorf("intervals within a range must not overlap")
			}
		}
		ranges = append(ranges, identifierRange{first: first, last: last})
	}
	return ranges, nil
}
