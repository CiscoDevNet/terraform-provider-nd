// Code generated;  DO NOT EDIT.

package datasource_remote_storage_location

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCRemoteStorageLocationModel struct {
	Name               string           `json:"name,omitempty"`
	Description        string           `json:"description,omitempty"`
	Hostname           string           `json:"hostname,omitempty"`
	Path               string           `json:"path,omitempty"`
	Nfs                NDFCNfsValue     `json:"-"`
	ScpSftp            NDFCScpSftpValue `json:"-"`
	HealthState        string           `json:"-"`
	HealthStateMessage string           `json:"-"`
}

type NDFCNfsValue struct {
	Port           *int64 `json:"port,omitempty"`
	Limit          string `json:"limit,omitempty"`
	ReadWrite      *bool  `json:"readWrite,omitempty"`
	AlertThreshold *int64 `json:"alertThreshold,omitempty"`
}

type NDFCScpSftpValue struct {
	Protocol       string                  `json:"type,omitempty"`
	Port           *int64                  `json:"port,omitempty"`
	Authentication NDFCAuthenticationValue `json:"authentication,omitempty"`
}

type NDFCAuthenticationValue struct {
	Username                string `json:"username,omitempty"`
	IgnoreHostKeyValidation *bool  `json:"ignoreHostKeyValidation,omitempty"`
}

func (v *RemoteStorageLocationModel) SetModelData(jsonData *NDFCRemoteStorageLocationModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Name != "" {
		v.Name = types.StringValue(jsonData.Name)
	} else {
		v.Name = types.StringNull()
	}

	if jsonData.Description != "" {
		v.Description = types.StringValue(jsonData.Description)
	} else {
		v.Description = types.StringNull()
	}

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.Path != "" {
		v.Path = types.StringValue(jsonData.Path)
	} else {
		v.Path = types.StringNull()
	}

	v.Nfs.SetValue(&jsonData.Nfs)

	v.ScpSftp.SetValue(&jsonData.ScpSftp)

	if jsonData.HealthState != "" {
		v.HealthState = types.StringValue(jsonData.HealthState)
	} else {
		v.HealthState = types.StringNull()
	}

	if jsonData.HealthStateMessage != "" {
		v.HealthStateMessage = types.StringValue(jsonData.HealthStateMessage)
	} else {
		v.HealthStateMessage = types.StringNull()
	}

	return err
}

func (v *NfsValue) SetValue(jsonData *NDFCNfsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	valueStateKnown := false

	if jsonData.Port != nil {
		v.Port = types.Int64Value(*jsonData.Port)
		valueStateKnown = true
	} else {
		v.Port = types.Int64Null()
	}

	if jsonData.Limit != "" {
		v.Limit = types.StringValue(jsonData.Limit)
		valueStateKnown = true
	} else {
		v.Limit = types.StringNull()
	}

	if jsonData.ReadWrite != nil {
		v.ReadWrite = types.BoolValue(*jsonData.ReadWrite)
		valueStateKnown = true

	} else {
		v.ReadWrite = types.BoolNull()
	}

	if jsonData.AlertThreshold != nil {
		v.AlertThreshold = types.Int64Value(*jsonData.AlertThreshold)
		valueStateKnown = true
	} else {
		v.AlertThreshold = types.Int64Null()
	}

	if valueStateKnown {
		v.state = attr.ValueStateKnown
	}

	return err
}

func (v *ScpSftpValue) SetValue(jsonData *NDFCScpSftpValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	valueStateKnown := false
	if jsonData.Protocol != "" {
		v.Protocol = types.StringValue(jsonData.Protocol)
		valueStateKnown = true
	} else {
		v.Protocol = types.StringNull()
	}

	if jsonData.Port != nil {
		v.Port = types.Int64Value(*jsonData.Port)
		valueStateKnown = true
	} else {
		v.Port = types.Int64Null()
	}

	if jsonData.Authentication.Username != "" {
		v.Username = types.StringValue(jsonData.Authentication.Username)
		valueStateKnown = true
	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Authentication.IgnoreHostKeyValidation != nil {
		v.IgnoreHostKeyValidation = types.BoolValue(*jsonData.Authentication.IgnoreHostKeyValidation)
		valueStateKnown = true
	} else {
		v.IgnoreHostKeyValidation = types.BoolNull()
	}

	if valueStateKnown {
		v.state = attr.ValueStateKnown
	}

	return err
}

func (v RemoteStorageLocationModel) GetModelData() *NDFCRemoteStorageLocationModel {
	var data = new(NDFCRemoteStorageLocationModel)

	//MARSHAL_BODY

	if !v.Name.IsNull() && !v.Name.IsUnknown() {
		data.Name = v.Name.ValueString()
	} else {
		data.Name = ""
	}

	return data
}
