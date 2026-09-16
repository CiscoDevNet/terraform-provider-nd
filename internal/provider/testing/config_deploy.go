// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/manage/resource_config_deploy"
)

// GenerateConfigDeployObject creates a config deploy model for testing.
func GenerateConfigDeployObject(
	obj **resource_config_deploy.NDFCConfigDeployModel,
	fabricName string,
	configSave bool,
	deploy bool,
	serialNumbers []string,
	behaviour *resource_config_deploy.NDFCBehaviourValue,
) {
	cd := new(resource_config_deploy.NDFCConfigDeployModel)

	cd.FabricName = fabricName
	cd.ConfigSave = configSave
	cd.Deploy = deploy
	cd.SwitchIds = serialNumbers

	if behaviour != nil {
		cd.Behaviour = *behaviour
	} else {
		t := true
		f := false
		cd.Behaviour = resource_config_deploy.NDFCBehaviourValue{
			DeployOnCreate:  &t,
			DeployOnUpdate:  &t,
			DeployOnDestroy: &f,
		}
	}

	*obj = cd
}

// ModifyConfigDeployObject modifies fields on an existing config deploy model.
func ModifyConfigDeployObject(
	obj **resource_config_deploy.NDFCConfigDeployModel,
	values map[string]interface{},
) {
	cd := *obj
	if cd == nil {
		return
	}

	for key, val := range values {
		switch key {
		case "deploy":
			cd.Deploy = val.(bool)
		case "config_save":
			cd.ConfigSave = val.(bool)
		case "switch_ids":
			cd.SwitchIds = val.([]string)
		case "always_deploy":
			cd.AlwaysDeploy = val.(bool)
		case "force_show_run":
			cd.ForceShowRun = val.(bool)
		case "include_all_fabric_group_switches":
			cd.IncludeAllFabricGroupSwitches = val.(bool)
		case "ticket_id":
			cd.TicketId = val.(string)
		case "behaviour":
			cd.Behaviour = val.(resource_config_deploy.NDFCBehaviourValue)
		}
	}

	*obj = cd
}

// BoolPtr returns a pointer to a bool value (helper for behaviour fields).
func BoolPtr(v bool) *bool {
	return &v
}
