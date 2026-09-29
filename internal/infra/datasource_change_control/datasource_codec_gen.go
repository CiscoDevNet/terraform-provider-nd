// Code generated;  DO NOT EDIT.

package datasource_change_control

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCChangeControlModel struct {
	AdminStatus                  *bool  `json:"changeControlAdminStatus,omitempty"`
	Orchestration                *bool  `json:"changeControlOrchestration,omitempty"`
	NumberOfApprovers            *int64 `json:"numberOfApprovers,omitempty"`
	AllowSelfApproval            *bool  `json:"allowSelfApproval,omitempty"`
	NdManagedFabrics             *bool  `json:"changeControlNDManagedFabrics,omitempty"`
	BypassTelemetryChangeControl *bool  `json:"byPassTelemetryChangeControl,omitempty"`
	TicketNamePrefix             string `json:"ticketNamePrefix,omitempty"`
}

func (v *ChangeControlModel) SetModelData(jsonData *NDFCChangeControlModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.AdminStatus != nil {
		v.AdminStatus = types.BoolValue(*jsonData.AdminStatus)

	} else {
		v.AdminStatus = types.BoolNull()
	}

	if jsonData.Orchestration != nil {
		v.Orchestration = types.BoolValue(*jsonData.Orchestration)

	} else {
		v.Orchestration = types.BoolNull()
	}

	if jsonData.NumberOfApprovers != nil {
		v.NumberOfApprovers = types.Int64Value(*jsonData.NumberOfApprovers)
	} else {
		v.NumberOfApprovers = types.Int64Null()
	}

	if jsonData.AllowSelfApproval != nil {
		v.AllowSelfApproval = types.BoolValue(*jsonData.AllowSelfApproval)

	} else {
		v.AllowSelfApproval = types.BoolNull()
	}

	if jsonData.NdManagedFabrics != nil {
		v.NdManagedFabrics = types.BoolValue(*jsonData.NdManagedFabrics)

	} else {
		v.NdManagedFabrics = types.BoolNull()
	}

	if jsonData.BypassTelemetryChangeControl != nil {
		v.BypassTelemetryChangeControl = types.BoolValue(*jsonData.BypassTelemetryChangeControl)

	} else {
		v.BypassTelemetryChangeControl = types.BoolNull()
	}

	if jsonData.TicketNamePrefix != "" {
		v.TicketNamePrefix = types.StringValue(jsonData.TicketNamePrefix)
	} else {
		v.TicketNamePrefix = types.StringNull()
	}

	return err
}

func (v ChangeControlModel) GetModelData() *NDFCChangeControlModel {
	var data = new(NDFCChangeControlModel)

	//MARSHAL_BODY

	return data
}
