// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_policy"
)

// GeneratePolicyObject builds an nd_policy model for acceptance tests.
func GeneratePolicyObject(
	obj **resource_policy.NDFCPolicyModel,
	fabricName string,
	serialNumbers []string,
	templateName string,
	policyParameters map[string]string,
	isPolicyGroup bool,
	deploy bool,
) {
	policy := new(resource_policy.NDFCPolicyModel)
	policy.FabricName = fabricName
	policy.SerialNumbers = append([]string{}, serialNumbers...)
	policy.TemplateName = templateName
	policy.EntityType = "switch"
	policy.EntityName = "SWITCH"
	policy.Description = "Enable LACP"
	policy.IsPolicyGroup = isPolicyGroup
	policy.Deploy = deploy
	priority := int64(500)
	policy.Priority = &priority
	if policyParameters != nil {
		policy.PolicyParameters = copyStringMap(policyParameters)
	}

	*obj = policy
}

// ModifyPolicyObject updates selected fields on an existing policy model.
func ModifyPolicyObject(
	obj **resource_policy.NDFCPolicyModel,
	values map[string]interface{},
) {
	policy := *obj
	if policy == nil {
		return
	}

	for key, val := range values {
		switch key {
		case "fabric_name":
			policy.FabricName = val.(string)
		case "serial_numbers":
			policy.SerialNumbers = append([]string{}, val.([]string)...)
		case "template_name":
			policy.TemplateName = val.(string)
		case "entity_type":
			policy.EntityType = val.(string)
		case "entity_name":
			policy.EntityName = val.(string)
		case "description":
			policy.Description = val.(string)
		case "source":
			policy.Source = val.(string)
		case "priority":
			v := int64(val.(int))
			policy.Priority = &v
		case "is_policy_group":
			policy.IsPolicyGroup = val.(bool)
		case "deploy":
			policy.Deploy = val.(bool)
		case "policy_parameters":
			policy.PolicyParameters = copyStringMap(val.(map[string]string))
		case "secondary_entity_name":
			policy.SecondaryEntityName = val.(string)
		case "secondary_entity_type":
			policy.SecondaryEntityType = val.(string)
		}
	}

	*obj = policy
}

func copyStringMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
