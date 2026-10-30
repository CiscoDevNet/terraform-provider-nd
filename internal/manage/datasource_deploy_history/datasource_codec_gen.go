// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package datasource_deploy_history

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCDeployHistoryModel struct {
	FabricName        string                      `json:"name,omitempty"`
	NoOfRecords       *int64                      `json:"no_of_records,omitempty"`
	DeploymentRecords NDFCDeploymentRecordsValues `json:"deploymentRecords,omitempty"`
}

type NDFCDeploymentRecordsValues []NDFCDeploymentRecordsValue

type NDFCDeploymentRecordsValue struct {
	CompleteTimestamp      string                           `json:"completeTimestamp,omitempty"`
	ConfigCommandResponses NDFCConfigCommandResponsesValues `json:"configCommandResponses,omitempty"`
	EntityName             string                           `json:"entityName,omitempty"`
	EntityType             string                           `json:"entityType,omitempty"`
	Hostname               string                           `json:"hostname,omitempty"`
	SecondaryEntityName    string                           `json:"secondaryEntityName,omitempty"`
	SecondaryEntityType    string                           `json:"secondaryEntityType,omitempty"`
	Source                 string                           `json:"source,omitempty"`
	StartTimestamp         string                           `json:"startTimestamp,omitempty"`
	Status                 string                           `json:"status,omitempty"`
	StatusDescription      string                           `json:"statusDescription,omitempty"`
	SwitchId               string                           `json:"switchId,omitempty"`
	SwitchIp               string                           `json:"switchIp,omitempty"`
	TicketId               string                           `json:"ticketId,omitempty"`
	Username               string                           `json:"username,omitempty"`
}

type NDFCConfigCommandResponsesValues []NDFCConfigCommandResponsesValue

type NDFCConfigCommandResponsesValue struct {
	CliResponse       string `json:"cliResponse,omitempty"`
	Command           string `json:"command,omitempty"`
	CompleteTimestamp string `json:"completeTimestamp,omitempty"`
	Status            string `json:"status,omitempty"`
}

func (v *DeployHistoryModel) SetModelData(jsonData *NDFCDeployHistoryModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.NoOfRecords != nil {
		v.NoOfRecords = types.Int64Value(*jsonData.NoOfRecords)
	} else {
		v.NoOfRecords = types.Int64Null()
	}

	if len(jsonData.DeploymentRecords) == 0 {
		log.Printf("v.DeploymentRecords is empty")
		v.DeploymentRecords = types.ListNull(DeploymentRecordsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.DeploymentRecords contains %d elements", len(jsonData.DeploymentRecords))
		listData := make([]DeploymentRecordsValue, 0)
		for _, item := range jsonData.DeploymentRecords {
			data := new(DeploymentRecordsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in DeploymentRecordsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.DeploymentRecords, err = types.ListValueFrom(context.Background(), DeploymentRecordsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []DeploymentRecordsValue to  List")
			return err
		}
	}

	return err
}

func (v *DeploymentRecordsValue) SetValue(jsonData *NDFCDeploymentRecordsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.CompleteTimestamp != "" {
		v.CompleteTimestamp = types.StringValue(jsonData.CompleteTimestamp)
	} else {
		v.CompleteTimestamp = types.StringNull()
	}

	if len(jsonData.ConfigCommandResponses) == 0 {
		log.Printf("v.ConfigCommandResponses is empty")
		v.ConfigCommandResponses = types.ListNull(ConfigCommandResponsesValue{}.Type(context.Background()))
	} else {
		log.Printf("v.ConfigCommandResponses contains %d elements", len(jsonData.ConfigCommandResponses))
		listData := make([]ConfigCommandResponsesValue, 0)
		for _, item := range jsonData.ConfigCommandResponses {
			data := new(ConfigCommandResponsesValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in ConfigCommandResponsesValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.ConfigCommandResponses, err = types.ListValueFrom(context.Background(), ConfigCommandResponsesValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []ConfigCommandResponsesValue to  List")
			return err
		}
	}
	if jsonData.EntityName != "" {
		v.EntityName = types.StringValue(jsonData.EntityName)
	} else {
		v.EntityName = types.StringNull()
	}

	if jsonData.EntityType != "" {
		v.EntityType = types.StringValue(jsonData.EntityType)
	} else {
		v.EntityType = types.StringNull()
	}

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.SecondaryEntityName != "" {
		v.SecondaryEntityName = types.StringValue(jsonData.SecondaryEntityName)
	} else {
		v.SecondaryEntityName = types.StringNull()
	}

	if jsonData.SecondaryEntityType != "" {
		v.SecondaryEntityType = types.StringValue(jsonData.SecondaryEntityType)
	} else {
		v.SecondaryEntityType = types.StringNull()
	}

	if jsonData.Source != "" {
		v.Source = types.StringValue(jsonData.Source)
	} else {
		v.Source = types.StringNull()
	}

	if jsonData.StartTimestamp != "" {
		v.StartTimestamp = types.StringValue(jsonData.StartTimestamp)
	} else {
		v.StartTimestamp = types.StringNull()
	}

	if jsonData.Status != "" {
		v.Status = types.StringValue(jsonData.Status)
	} else {
		v.Status = types.StringNull()
	}

	if jsonData.StatusDescription != "" {
		v.StatusDescription = types.StringValue(jsonData.StatusDescription)
	} else {
		v.StatusDescription = types.StringNull()
	}

	if jsonData.SwitchId != "" {
		v.SwitchId = types.StringValue(jsonData.SwitchId)
	} else {
		v.SwitchId = types.StringNull()
	}

	if jsonData.SwitchIp != "" {
		v.SwitchIp = types.StringValue(jsonData.SwitchIp)
	} else {
		v.SwitchIp = types.StringNull()
	}

	if jsonData.TicketId != "" {
		v.TicketId = types.StringValue(jsonData.TicketId)
	} else {
		v.TicketId = types.StringNull()
	}

	if jsonData.Username != "" {
		v.Username = types.StringValue(jsonData.Username)
	} else {
		v.Username = types.StringNull()
	}

	return err
}

func (v *ConfigCommandResponsesValue) SetValue(jsonData *NDFCConfigCommandResponsesValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.CliResponse != "" {
		v.CliResponse = types.StringValue(jsonData.CliResponse)
	} else {
		v.CliResponse = types.StringNull()
	}

	if jsonData.Command != "" {
		v.Command = types.StringValue(jsonData.Command)
	} else {
		v.Command = types.StringNull()
	}

	if jsonData.CompleteTimestamp != "" {
		v.CompleteTimestamp = types.StringValue(jsonData.CompleteTimestamp)
	} else {
		v.CompleteTimestamp = types.StringNull()
	}

	if jsonData.Status != "" {
		v.Status = types.StringValue(jsonData.Status)
	} else {
		v.Status = types.StringNull()
	}

	return err
}

func (v DeployHistoryModel) GetModelData() *NDFCDeployHistoryModel {
	var data = new(NDFCDeployHistoryModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.NoOfRecords.IsNull() && !v.NoOfRecords.IsUnknown() {
		data.NoOfRecords = new(int64)
		*data.NoOfRecords = v.NoOfRecords.ValueInt64()

	} else {
		data.NoOfRecords = nil
	}

	return data
}
