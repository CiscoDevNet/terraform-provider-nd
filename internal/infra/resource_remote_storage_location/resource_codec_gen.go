// Code generated;  DO NOT EDIT.

package resource_remote_storage_location

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCRemoteStorageLocationModel struct {
	Name                string                  `json:"name,omitempty"`
	Description         string                  `json:"description,omitempty"`
	StorageLocationType string                  `json:"type,omitempty"`
	ReadWrite           *bool                   `json:"readWrite,omitempty"`
	Hostname            string                  `json:"hostname,omitempty"`
	Port                *int64                  `json:"port,omitempty"`
	Path                string                  `json:"path,omitempty"`
	AlertThreshold      *int64                  `json:"alertThreshold,omitempty"`
	Limit               string                  `json:"limit,omitempty"`
	Authentication      NDFCAuthenticationValue `json:"authentication,omitempty"`
}

type NDFCAuthenticationValue struct {
	Username                string `json:"username,omitempty"`
	AuthenticationType      string `json:"type,omitempty"`
	Password                string `json:"password,omitempty"`
	SshKey                  string `json:"sshKey,omitempty"`
	Passphrase              string `json:"passphrase,omitempty"`
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

	if jsonData.StorageLocationType != "" {
		v.StorageLocationType = types.StringValue(jsonData.StorageLocationType)
	} else {
		v.StorageLocationType = types.StringNull()
	}

	if jsonData.ReadWrite != nil {
		v.ReadWrite = types.BoolValue(*jsonData.ReadWrite)

	} else {
		v.ReadWrite = types.BoolNull()
	}

	if jsonData.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Hostname)
	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.Port != nil {
		v.Port = types.Int64Value(*jsonData.Port)

	} else {
		v.Port = types.Int64Null()
	}

	if jsonData.Path != "" {
		v.Path = types.StringValue(jsonData.Path)
	} else {
		v.Path = types.StringNull()
	}

	if jsonData.AlertThreshold != nil {
		v.AlertThreshold = types.Int64Value(*jsonData.AlertThreshold)

	} else {
		v.AlertThreshold = types.Int64Null()
	}

	if jsonData.Limit != "" {
		v.Limit = types.StringValue(jsonData.Limit)
	} else {
		v.Limit = types.StringNull()
	}

	if jsonData.Authentication.Username != "" {
		v.Username = types.StringValue(jsonData.Authentication.Username)

	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Authentication.AuthenticationType != "" {
		v.AuthenticationType = types.StringValue(jsonData.Authentication.AuthenticationType)

	} else {
		v.AuthenticationType = types.StringNull()
	}

	if jsonData.Authentication.Password != "" {
		v.Password = types.StringValue(jsonData.Authentication.Password)

	} else {
		v.Password = types.StringNull()
	}

	if jsonData.Authentication.SshKey != "" {
		v.SshKey = types.StringValue(jsonData.Authentication.SshKey)

	} else {
		v.SshKey = types.StringNull()
	}

	if jsonData.Authentication.Passphrase != "" {
		v.Passphrase = types.StringValue(jsonData.Authentication.Passphrase)

	} else {
		v.Passphrase = types.StringNull()
	}

	if jsonData.Authentication.IgnoreHostKeyValidation != nil {
		v.IgnoreHostKeyValidation = types.BoolValue(*jsonData.Authentication.IgnoreHostKeyValidation)

	} else {
		v.IgnoreHostKeyValidation = types.BoolNull()
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

	if !v.Description.IsNull() && !v.Description.IsUnknown() {
		data.Description = v.Description.ValueString()
	} else {
		data.Description = ""
	}

	if !v.StorageLocationType.IsNull() && !v.StorageLocationType.IsUnknown() {
		data.StorageLocationType = v.StorageLocationType.ValueString()
	} else {
		data.StorageLocationType = ""
	}

	if !v.ReadWrite.IsNull() && !v.ReadWrite.IsUnknown() {
		data.ReadWrite = new(bool)
		*data.ReadWrite = v.ReadWrite.ValueBool()
	} else {
		data.ReadWrite = nil
	}

	if !v.Hostname.IsNull() && !v.Hostname.IsUnknown() {
		data.Hostname = v.Hostname.ValueString()
	} else {
		data.Hostname = ""
	}

	if !v.Port.IsNull() && !v.Port.IsUnknown() {
		data.Port = new(int64)
		*data.Port = v.Port.ValueInt64()

	} else {
		data.Port = nil
	}

	if !v.Path.IsNull() && !v.Path.IsUnknown() {
		data.Path = v.Path.ValueString()
	} else {
		data.Path = ""
	}

	if !v.AlertThreshold.IsNull() && !v.AlertThreshold.IsUnknown() {
		data.AlertThreshold = new(int64)
		*data.AlertThreshold = v.AlertThreshold.ValueInt64()

	} else {
		data.AlertThreshold = nil
	}

	if !v.Limit.IsNull() && !v.Limit.IsUnknown() {
		data.Limit = v.Limit.ValueString()
	} else {
		data.Limit = ""
	}

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Authentication.Username = v.Username.ValueString()
	} else {
		data.Authentication.Username = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Authentication.Password = v.Password.ValueString()
	} else {
		data.Authentication.Password = ""
	}

	if !v.SshKey.IsNull() && !v.SshKey.IsUnknown() {
		data.Authentication.SshKey = v.SshKey.ValueString()
	} else {
		data.Authentication.SshKey = ""
	}

	if !v.Passphrase.IsNull() && !v.Passphrase.IsUnknown() {
		data.Authentication.Passphrase = v.Passphrase.ValueString()
	} else {
		data.Authentication.Passphrase = ""
	}

	if !v.IgnoreHostKeyValidation.IsNull() && !v.IgnoreHostKeyValidation.IsUnknown() {
		data.Authentication.IgnoreHostKeyValidation = new(bool)
		*data.Authentication.IgnoreHostKeyValidation = v.IgnoreHostKeyValidation.ValueBool()
	} else {
		data.Authentication.IgnoreHostKeyValidation = nil
	}

	return data
}
