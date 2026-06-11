// Code generated;  DO NOT EDIT.

package resource_change_control

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

	if !v.AdminStatus.IsNull() && !v.AdminStatus.IsUnknown() {
		data.AdminStatus = new(bool)
		*data.AdminStatus = v.AdminStatus.ValueBool()
	} else {
		data.AdminStatus = nil
	}

	if !v.Orchestration.IsNull() && !v.Orchestration.IsUnknown() {
		data.Orchestration = new(bool)
		*data.Orchestration = v.Orchestration.ValueBool()
	} else {
		data.Orchestration = nil
	}

	if !v.NumberOfApprovers.IsNull() && !v.NumberOfApprovers.IsUnknown() {
		data.NumberOfApprovers = new(int64)
		*data.NumberOfApprovers = v.NumberOfApprovers.ValueInt64()

	} else {
		data.NumberOfApprovers = nil
	}

	if !v.AllowSelfApproval.IsNull() && !v.AllowSelfApproval.IsUnknown() {
		data.AllowSelfApproval = new(bool)
		*data.AllowSelfApproval = v.AllowSelfApproval.ValueBool()
	} else {
		data.AllowSelfApproval = nil
	}

	if !v.NdManagedFabrics.IsNull() && !v.NdManagedFabrics.IsUnknown() {
		data.NdManagedFabrics = new(bool)
		*data.NdManagedFabrics = v.NdManagedFabrics.ValueBool()
	} else {
		data.NdManagedFabrics = nil
	}

	if !v.BypassTelemetryChangeControl.IsNull() && !v.BypassTelemetryChangeControl.IsUnknown() {
		data.BypassTelemetryChangeControl = new(bool)
		*data.BypassTelemetryChangeControl = v.BypassTelemetryChangeControl.ValueBool()
	} else {
		data.BypassTelemetryChangeControl = nil
	}

	if !v.TicketNamePrefix.IsNull() && !v.TicketNamePrefix.IsUnknown() {
		data.TicketNamePrefix = v.TicketNamePrefix.ValueString()
	} else {
		data.TicketNamePrefix = ""
	}

	return data
}
