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

	"terraform-provider-nd/internal/manage/resource_interface_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// ethernetHandler overrides DefaultInterfaceHandler.Read to detect policy
// owner mismatches. When ND reports a policyType different from the one in
// state (e.g. `dataBrokerPoMember` when the user configured `trunkHost`),
// the interface is owned by a port-channel or vPC and should not be
// silently adopted — emit a clear error instead.
type ethernetHandler struct {
	resource_interface_common.DefaultInterfaceHandler
}

func (h ethernetHandler) Read(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model resource_interface_common.InterfaceModel) bool {
	// Capture prior state policyType per interface before the read overwrites it.
	priorData := model.GetModelData()
	priorPolicy := make(map[string]string, len(priorData.Interfaces))
	for name, entry := range priorData.Interfaces {
		priorPolicy[name] = entry.ConfigData.NetworkOS.Policy.PolicyType
	}

	// Delegate to the default read.
	found := h.DefaultInterfaceHandler.Read(ctx, client, dg, model)
	if !found || dg.HasError() {
		return found
	}

	// Compare returned policyType against the prior state for each entry.
	freshData := model.GetModelData()
	for name, freshEntry := range freshData.Interfaces {
		prior, ok := priorPolicy[name]
		if !ok || prior == "" {
			continue // newly discovered or empty prior — skip
		}
		got := freshEntry.ConfigData.NetworkOS.Policy.PolicyType
		if got != "" && got != prior {
			tflog.Warn(ctx, "interface policy owner mismatch", map[string]interface{}{
				"interface": name, "expected": prior, "got": got,
			})
			dg.AddError("Interface policy owner mismatch",
				fmt.Sprintf("Interface %q is now owned by policy %q (expected %q). "+
					"This typically means a port-channel or vPC membership has taken ownership. "+
					"Remove the interface from this resource or reconcile the ownership in ND.",
					name, got, prior))
		}
	}
	return found
}
