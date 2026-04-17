// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_inventory_switch

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tfStringState returns a human-readable representation of a types.String value
func tfStringState(v types.String) string {
	if v.IsNull() {
		return "<null>"
	}
	if v.IsUnknown() {
		return "<unknown>"
	}
	return fmt.Sprintf("%q", v.ValueString())
}

// tfBoolState returns a human-readable representation of a types.Bool value
func tfBoolState(v types.Bool) string {
	if v.IsNull() {
		return "<null>"
	}
	if v.IsUnknown() {
		return "<unknown>"
	}
	return fmt.Sprintf("%v", v.ValueBool())
}

// tfInt64State returns a human-readable representation of a types.Int64 value
func tfInt64State(v types.Int64) string {
	if v.IsNull() {
		return "<null>"
	}
	if v.IsUnknown() {
		return "<unknown>"
	}
	return fmt.Sprintf("%d", v.ValueInt64())
}

// DumpInventorySwitchModel logs every field in the InventorySwitchModel,
// clearly indicating null/unknown/value for each attribute.
func DumpInventorySwitchModel(label string, m *InventorySwitchModel) {
	if m == nil {
		log.Printf("[DUMP] %s: InventorySwitchModel is nil", label)
		return
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("[DUMP] %s: InventorySwitchModel {\n", label))
	b.WriteString(fmt.Sprintf("  fabric_name:               %s\n", tfStringState(m.FabricName)))
	b.WriteString(fmt.Sprintf("  mode:                      %s\n", tfStringState(m.Mode)))
	b.WriteString(fmt.Sprintf("  username:                  %s\n", tfStringState(m.Username)))
	b.WriteString(fmt.Sprintf("  password:                  %s\n", tfStringState(m.Password)))
	b.WriteString(fmt.Sprintf("  platform_type:             %s\n", tfStringState(m.PlatformType)))
	b.WriteString(fmt.Sprintf("  preserve_config:           %s\n", tfBoolState(m.PreserveConfig)))
	b.WriteString(fmt.Sprintf("  recalculate:               %s\n", tfBoolState(m.Recalculate)))
	b.WriteString(fmt.Sprintf("  deploy:                    %s\n", tfBoolState(m.Deploy)))
	b.WriteString(fmt.Sprintf("  max_hop:                   %s\n", tfInt64State(m.MaxHop)))
	b.WriteString(fmt.Sprintf("  snmp_v3_auth_protocol:     %s\n", tfStringState(m.SnmpV3AuthProtocol)))
	b.WriteString(fmt.Sprintf("  remote_credential_store:   %s\n", tfStringState(m.RemoteCredentialStore)))
	b.WriteString(fmt.Sprintf("  remote_credential_store_key: %s\n", tfStringState(m.RemoteCredentialStoreKey)))

	// Switches map
	if m.Switches.IsNull() {
		b.WriteString("  switches: <null>\n")
	} else if m.Switches.IsUnknown() {
		b.WriteString("  switches: <unknown>\n")
	} else {
		elements := make(map[string]SwitchesValue, len(m.Switches.Elements()))
		diag := m.Switches.ElementsAs(context.Background(), &elements, false)
		if diag.HasError() {
			b.WriteString(fmt.Sprintf("  switches: <error reading elements: %v>\n", diag))
		} else {
			b.WriteString(fmt.Sprintf("  switches: (%d entries) {\n", len(elements)))
			for key, sw := range elements {
				b.WriteString(fmt.Sprintf("    [%q] {\n", key))
				b.WriteString(fmt.Sprintf("      ip_address:              %s\n", tfStringState(sw.IpAddress)))
				b.WriteString(fmt.Sprintf("      hostname:                %s\n", tfStringState(sw.Hostname)))
				b.WriteString(fmt.Sprintf("      serial_number:           %s\n", tfStringState(sw.SerialNumber)))
				b.WriteString(fmt.Sprintf("      model:                   %s\n", tfStringState(sw.Model)))
				b.WriteString(fmt.Sprintf("      software_version:        %s\n", tfStringState(sw.SoftwareVersion)))
				b.WriteString(fmt.Sprintf("      switch_role:             %s\n", tfStringState(sw.SwitchRole)))
				b.WriteString(fmt.Sprintf("      status:                  %s\n", tfStringState(sw.Status)))
				b.WriteString(fmt.Sprintf("      status_reason:           %s\n", tfStringState(sw.StatusReason)))
				b.WriteString(fmt.Sprintf("      gateway_ip_mask:         %s\n", tfStringState(sw.GatewayIpMask)))
				b.WriteString(fmt.Sprintf("      discovery_auth_protocol: %s\n", tfStringState(sw.DiscoveryAuthProtocol)))
				b.WriteString(fmt.Sprintf("      poap_password:           %s\n", tfStringState(sw.PoapPassword)))
				b.WriteString(fmt.Sprintf("      vdc_id:                  %s\n", tfInt64State(sw.VdcId)))
				b.WriteString(fmt.Sprintf("      vdc_mac:                 %s\n", tfStringState(sw.VdcMac)))
				b.WriteString("    }\n")
			}
			b.WriteString("  }\n")
		}
	}

	b.WriteString("}")
	log.Printf("%s", b.String())
}
