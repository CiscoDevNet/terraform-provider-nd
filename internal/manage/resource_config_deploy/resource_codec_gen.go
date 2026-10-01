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
	AlwaysExecute                 bool               `json:"-"`
	Behaviour                     NDFCBehaviourValue `json:"-"`
	Status                        string             `json:"status,omitempty"`
	TicketId                      string             `json:"-"`
}

type NDFCBehaviourValue struct {
	ExecuteOnCreate  *bool `json:"onCreate,omitempty"`
	ExecuteOnUpdate  *bool `json:"onUpdate,omitempty"`
	ExecuteOnDestroy *bool `json:"onDestroy,omitempty"`
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

	v.AlwaysExecute = types.BoolValue(jsonData.AlwaysExecute)
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

	if jsonData.ExecuteOnCreate != nil {
		v.ExecuteOnCreate = types.BoolValue(*jsonData.ExecuteOnCreate)

	} else {
		v.ExecuteOnCreate = types.BoolNull()
	}

	if jsonData.ExecuteOnUpdate != nil {
		v.ExecuteOnUpdate = types.BoolValue(*jsonData.ExecuteOnUpdate)

	} else {
		v.ExecuteOnUpdate = types.BoolNull()
	}

	if jsonData.ExecuteOnDestroy != nil {
		v.ExecuteOnDestroy = types.BoolValue(*jsonData.ExecuteOnDestroy)

	} else {
		v.ExecuteOnDestroy = types.BoolNull()
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

	if !v.AlwaysExecute.IsNull() && !v.AlwaysExecute.IsUnknown() {
		data.AlwaysExecute = v.AlwaysExecute.ValueBool()
	}

	if !v.TicketId.IsNull() && !v.TicketId.IsUnknown() {
		data.TicketId = v.TicketId.ValueString()
	} else {
		data.TicketId = ""
	}

	//MARSHAL_BODY

	// Nested types Behaviour # execute_on_create
	if !v.Behaviour.ExecuteOnCreate.IsNull() && !v.Behaviour.ExecuteOnCreate.IsUnknown() {
		data.Behaviour.ExecuteOnCreate = new(bool)
		*data.Behaviour.ExecuteOnCreate = v.Behaviour.ExecuteOnCreate.ValueBool()
	} else {
		data.Behaviour.ExecuteOnCreate = nil
	}

	// Nested types Behaviour # execute_on_update
	if !v.Behaviour.ExecuteOnUpdate.IsNull() && !v.Behaviour.ExecuteOnUpdate.IsUnknown() {
		data.Behaviour.ExecuteOnUpdate = new(bool)
		*data.Behaviour.ExecuteOnUpdate = v.Behaviour.ExecuteOnUpdate.ValueBool()
	} else {
		data.Behaviour.ExecuteOnUpdate = nil
	}

	// Nested types Behaviour # execute_on_destroy
	if !v.Behaviour.ExecuteOnDestroy.IsNull() && !v.Behaviour.ExecuteOnDestroy.IsUnknown() {
		data.Behaviour.ExecuteOnDestroy = new(bool)
		*data.Behaviour.ExecuteOnDestroy = v.Behaviour.ExecuteOnDestroy.ValueBool()
	} else {
		data.Behaviour.ExecuteOnDestroy = nil
	}

	return data
}
