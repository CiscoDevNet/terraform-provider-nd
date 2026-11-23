// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_switch_config_diff

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCSwitchConfigDiffModel struct {
	FabricName  string                `json:"fabricName,omitempty"`
	SwitchId    string                `json:"switchId,omitempty"`
	DiffConfigs NDFCDiffConfigsValues `json:"diffConfigs,omitempty"`
}

type NDFCDiffConfigsValues []NDFCDiffConfigsValue

type NDFCDiffConfigsValue struct {
	DifferenceType string `json:"differenceType,omitempty"`
	ExpectedConfig string `json:"expectedConfig,omitempty"`
	RunningConfig  string `json:"runningConfig,omitempty"`
}

func (v *SwitchConfigDiffModel) SetModelData(jsonData *NDFCSwitchConfigDiffModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	if len(jsonData.DiffConfigs) == 0 {
		log.Printf("v.DiffConfigs is empty")
		v.DiffConfigs = types.ListNull(DiffConfigsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.DiffConfigs contains %d elements", len(jsonData.DiffConfigs))
		listData := make([]DiffConfigsValue, 0)
		for _, item := range jsonData.DiffConfigs {
			data := new(DiffConfigsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in DiffConfigsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.DiffConfigs, err = types.ListValueFrom(context.Background(), DiffConfigsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []DiffConfigsValue to  List")
			return err
		}
	}

	return err
}

func (v *DiffConfigsValue) SetValue(jsonData *NDFCDiffConfigsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.DifferenceType != "" {
		v.DifferenceType = types.StringValue(jsonData.DifferenceType)
	} else {
		v.DifferenceType = types.StringNull()
	}

	if jsonData.ExpectedConfig != "" {
		v.ExpectedConfig = types.StringValue(jsonData.ExpectedConfig)
	} else {
		v.ExpectedConfig = types.StringNull()
	}

	if jsonData.RunningConfig != "" {
		v.RunningConfig = types.StringValue(jsonData.RunningConfig)
	} else {
		v.RunningConfig = types.StringNull()
	}

	return err
}

func (v SwitchConfigDiffModel) GetModelData() *NDFCSwitchConfigDiffModel {
	var data = new(NDFCSwitchConfigDiffModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.SwitchId.IsNull() && !v.SwitchId.IsUnknown() {
		data.SwitchId = v.SwitchId.ValueString()
	} else {
		data.SwitchId = ""
	}

	return data
}
