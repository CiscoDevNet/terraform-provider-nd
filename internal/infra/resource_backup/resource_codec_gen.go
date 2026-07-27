// Code generated;  DO NOT EDIT.

package resource_backup

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCBackupModel struct {
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Destination   string            `json:"destination"`
	EncryptionKey string            `json:"encryptionKey,omitempty"`
	TelemetryData *bool             `json:"includeTelemetryOperationalData,omitempty"`
	Timeouts      NDFCTimeoutsValue `json:"-"`
}

type NDFCTimeoutsValue struct {
	Create string `json:"create,omitempty"`
	Read   string `json:"read,omitempty"`
}

func (v *BackupModel) SetModelData(jsonData *NDFCBackupModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Name != "" {
		v.Name = types.StringValue(jsonData.Name)
	} else {
		v.Name = types.StringNull()
	}

	if jsonData.Type != "" {
		v.Type = types.StringValue(jsonData.Type)
	} else {
		v.Type = types.StringNull()
	}

	if jsonData.Destination != "" {
		v.Destination = types.StringValue(jsonData.Destination)
	} else {
		v.Destination = types.StringNull()
	}

	if jsonData.EncryptionKey != "" {
		v.EncryptionKey = types.StringValue(jsonData.EncryptionKey)
	} else {
		v.EncryptionKey = types.StringNull()
	}

	if jsonData.TelemetryData != nil {
		v.TelemetryData = types.BoolValue(*jsonData.TelemetryData)

	} else {
		v.TelemetryData = types.BoolNull()
	}

	v.Timeouts.SetValue(&jsonData.Timeouts)
	v.Timeouts.state = attr.ValueStateKnown

	return err
}

func (v *TimeoutsValue) SetValue(jsonData *NDFCTimeoutsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.Create != "" {
		v.Create = types.StringValue(jsonData.Create)
	} else {
		v.Create = types.StringNull()
	}

	if jsonData.Read != "" {
		v.Read = types.StringValue(jsonData.Read)
	} else {
		v.Read = types.StringNull()
	}

	return err
}

func (v BackupModel) GetModelData() *NDFCBackupModel {
	var data = new(NDFCBackupModel)

	//MARSHAL_BODY

	if !v.Name.IsNull() && !v.Name.IsUnknown() {
		data.Name = v.Name.ValueString()
	} else {
		data.Name = ""
	}

	if !v.Type.IsNull() && !v.Type.IsUnknown() {
		data.Type = v.Type.ValueString()
	} else {
		data.Type = ""
	}

	if !v.Destination.IsNull() && !v.Destination.IsUnknown() {
		data.Destination = v.Destination.ValueString()
	} else {
		data.Destination = ""
	}

	if !v.EncryptionKey.IsNull() && !v.EncryptionKey.IsUnknown() {
		data.EncryptionKey = v.EncryptionKey.ValueString()
	} else {
		data.EncryptionKey = ""
	}

	if !v.TelemetryData.IsNull() && !v.TelemetryData.IsUnknown() {
		data.TelemetryData = new(bool)
		*data.TelemetryData = v.TelemetryData.ValueBool()
	} else {
		data.TelemetryData = nil
	}

	//MARSHAL_BODY

	// Nested types Timeouts # create
	if !v.Timeouts.Create.IsNull() && !v.Timeouts.Create.IsUnknown() {
		data.Timeouts.Create = v.Timeouts.Create.ValueString()
	} else {
		data.Timeouts.Create = ""
	}

	// Nested types Timeouts # read
	if !v.Timeouts.Read.IsNull() && !v.Timeouts.Read.IsUnknown() {
		data.Timeouts.Read = v.Timeouts.Read.ValueString()
	} else {
		data.Timeouts.Read = ""
	}

	return data
}
