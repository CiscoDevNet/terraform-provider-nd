// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import "terraform-provider-nd/internal/infra/resource_change_control"

// GenerateChangeControlObject builds a change control model from explicitly
// supplied values. An omitted value remains unset so Terraform can exercise
// the schema default for that attribute.
func GenerateChangeControlObject(
	obj **resource_change_control.NDFCChangeControlModel,
	values map[string]interface{},
) {
	changeControl := new(resource_change_control.NDFCChangeControlModel)
	applyChangeControlValues(changeControl, values)
	*obj = changeControl
}

// ModifyChangeControlObject replaces the model values between Terraform steps.
// Resetting first allows a later configuration to omit an earlier attribute.
func ModifyChangeControlObject(
	obj **resource_change_control.NDFCChangeControlModel,
	values map[string]interface{},
) {
	changeControl := *obj
	if changeControl == nil {
		changeControl = new(resource_change_control.NDFCChangeControlModel)
	}
	*changeControl = resource_change_control.NDFCChangeControlModel{}
	applyChangeControlValues(changeControl, values)
	*obj = changeControl
}

func applyChangeControlValues(
	changeControl *resource_change_control.NDFCChangeControlModel,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "admin_status":
			v := val.(bool)
			changeControl.AdminStatus = &v
		case "orchestration":
			v := val.(bool)
			changeControl.Orchestration = &v
		case "nd_managed_fabrics":
			v := val.(bool)
			changeControl.NdManagedFabrics = &v
		case "bypass_telemetry_change_control":
			v := val.(bool)
			changeControl.BypassTelemetryChangeControl = &v
		case "number_of_approvers":
			v := int64(val.(int))
			changeControl.NumberOfApprovers = &v
		case "allow_self_approval":
			v := val.(bool)
			changeControl.AllowSelfApproval = &v
		case "ticket_name_prefix":
			changeControl.TicketNamePrefix = val.(string)
		}
	}
}
