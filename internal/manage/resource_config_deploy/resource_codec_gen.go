// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_config_deploy

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCConfigDeployModel struct {
	Id                            string             `json:"id,omitempty"`
	FabricName                    string             `json:"fabricName,omitempty"`
	Deploy                        bool               `json:"-"`
	ConfigSave                    bool               `json:"-"`
	SwitchIds                     []string           `json:"-"`
	ForceShowRun                  bool               `json:"-"`
	IncludeAllFabricGroupSwitches bool               `json:"-"`
	AlwaysDeploy                  bool               `json:"-"`
	Behaviour                     NDFCBehaviourValue `json:"-"`
	Status                        string             `json:"status,omitempty"`
	TicketId                      string             `json:"-"`
}

type NDFCBehaviourValue struct {
	DeployOnCreate  *bool `json:"onCreate,omitempty"`
	DeployOnUpdate  *bool `json:"onUpdate,omitempty"`
	DeployOnDestroy *bool `json:"onDestroy,omitempty"`
}

func (v *ConfigDeployModel) SetModelData(jsonData *NDFCConfigDeployModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	v.Deploy = types.BoolValue(jsonData.Deploy)

	v.ConfigSave = types.BoolValue(jsonData.ConfigSave)

	if len(jsonData.SwitchIds) == 0 {
		log.Printf("v.SwitchIds is empty")
		v.SwitchIds = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.SwitchIds))
		for i, item := range jsonData.SwitchIds {
			listData[i] = types.StringValue(item)
		}
		v.SwitchIds, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	v.ForceShowRun = types.BoolValue(jsonData.ForceShowRun)

	v.IncludeAllFabricGroupSwitches = types.BoolValue(jsonData.IncludeAllFabricGroupSwitches)

	v.AlwaysDeploy = types.BoolValue(jsonData.AlwaysDeploy)
	v.Behaviour.SetValue(&jsonData.Behaviour)
	v.Behaviour.state = attr.ValueStateKnown

	if jsonData.Status != "" {
		v.Status = types.StringValue(jsonData.Status)
	} else {
		v.Status = types.StringNull()
	}

	if jsonData.TicketId != "" {
		v.TicketId = types.StringValue(jsonData.TicketId)
	} else {
		v.TicketId = types.StringNull()
	}

	return err
}

func (v *BehaviourValue) SetValue(jsonData *NDFCBehaviourValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.DeployOnCreate != nil {
		v.DeployOnCreate = types.BoolValue(*jsonData.DeployOnCreate)

	} else {
		v.DeployOnCreate = types.BoolNull()
	}

	if jsonData.DeployOnUpdate != nil {
		v.DeployOnUpdate = types.BoolValue(*jsonData.DeployOnUpdate)

	} else {
		v.DeployOnUpdate = types.BoolNull()
	}

	if jsonData.DeployOnDestroy != nil {
		v.DeployOnDestroy = types.BoolValue(*jsonData.DeployOnDestroy)

	} else {
		v.DeployOnDestroy = types.BoolNull()
	}

	return err
}

func (v ConfigDeployModel) GetModelData() *NDFCConfigDeployModel {
	var data = new(NDFCConfigDeployModel)

	//MARSHAL_BODY

	if !v.Id.IsNull() && !v.Id.IsUnknown() {
		data.Id = v.Id.ValueString()
	} else {
		data.Id = ""
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.Deploy.IsNull() && !v.Deploy.IsUnknown() {
		data.Deploy = v.Deploy.ValueBool()
	}

	if !v.ConfigSave.IsNull() && !v.ConfigSave.IsUnknown() {
		data.ConfigSave = v.ConfigSave.ValueBool()
	}

	if !v.SwitchIds.IsNull() && !v.SwitchIds.IsUnknown() {
		listStringData := make([]string, len(v.SwitchIds.Elements()))
		dg := v.SwitchIds.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.SwitchIds = make([]string, len(listStringData))
		copy(data.SwitchIds, listStringData)
	}

	if !v.ForceShowRun.IsNull() && !v.ForceShowRun.IsUnknown() {
		data.ForceShowRun = v.ForceShowRun.ValueBool()
	}

	if !v.IncludeAllFabricGroupSwitches.IsNull() && !v.IncludeAllFabricGroupSwitches.IsUnknown() {
		data.IncludeAllFabricGroupSwitches = v.IncludeAllFabricGroupSwitches.ValueBool()
	}

	if !v.AlwaysDeploy.IsNull() && !v.AlwaysDeploy.IsUnknown() {
		data.AlwaysDeploy = v.AlwaysDeploy.ValueBool()
	}

	if !v.TicketId.IsNull() && !v.TicketId.IsUnknown() {
		data.TicketId = v.TicketId.ValueString()
	} else {
		data.TicketId = ""
	}

	//MARSHAL_BODY

	// Nested types Behaviour # deploy_on_create
	if !v.Behaviour.DeployOnCreate.IsNull() && !v.Behaviour.DeployOnCreate.IsUnknown() {
		data.Behaviour.DeployOnCreate = new(bool)
		*data.Behaviour.DeployOnCreate = v.Behaviour.DeployOnCreate.ValueBool()
	} else {
		data.Behaviour.DeployOnCreate = nil
	}

	// Nested types Behaviour # deploy_on_update
	if !v.Behaviour.DeployOnUpdate.IsNull() && !v.Behaviour.DeployOnUpdate.IsUnknown() {
		data.Behaviour.DeployOnUpdate = new(bool)
		*data.Behaviour.DeployOnUpdate = v.Behaviour.DeployOnUpdate.ValueBool()
	} else {
		data.Behaviour.DeployOnUpdate = nil
	}

	// Nested types Behaviour # deploy_on_destroy
	if !v.Behaviour.DeployOnDestroy.IsNull() && !v.Behaviour.DeployOnDestroy.IsUnknown() {
		data.Behaviour.DeployOnDestroy = new(bool)
		*data.Behaviour.DeployOnDestroy = v.Behaviour.DeployOnDestroy.ValueBool()
	} else {
		data.Behaviour.DeployOnDestroy = nil
	}

	return data
}
