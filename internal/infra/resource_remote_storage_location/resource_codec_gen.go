// Code generated;  DO NOT EDIT.

package resource_remote_storage_location

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
	AcceptHostKey  bool                    `json:"-"`
	Authentication NDFCAuthenticationValue `json:"authentication,omitempty"`
}

type NDFCAuthenticationValue struct {
	Username                string `json:"username,omitempty"`
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

	if jsonData.Authentication.Password != "" {
		v.Password = types.StringValue(jsonData.Authentication.Password)
		valueStateKnown = true
	} else {
		v.Password = types.StringNull()
	}

	if jsonData.Authentication.SshKey != "" {
		v.SshKey = types.StringValue(jsonData.Authentication.SshKey)
		valueStateKnown = true
	} else {
		v.SshKey = types.StringNull()
	}

	if jsonData.Authentication.Passphrase != "" {
		v.Passphrase = types.StringValue(jsonData.Authentication.Passphrase)
		valueStateKnown = true
	} else {
		v.Passphrase = types.StringNull()
	}

	if jsonData.Authentication.IgnoreHostKeyValidation != nil {
		v.IgnoreHostKeyValidation = types.BoolValue(*jsonData.Authentication.IgnoreHostKeyValidation)
		valueStateKnown = true
	} else {
		v.IgnoreHostKeyValidation = types.BoolNull()
	}

	v.AcceptHostKey = types.BoolValue(jsonData.AcceptHostKey)
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

	if !v.Description.IsNull() && !v.Description.IsUnknown() {
		data.Description = v.Description.ValueString()
	} else {
		data.Description = ""
	}

	if !v.Hostname.IsNull() && !v.Hostname.IsUnknown() {
		data.Hostname = v.Hostname.ValueString()
	} else {
		data.Hostname = ""
	}

	if !v.Path.IsNull() && !v.Path.IsUnknown() {
		data.Path = v.Path.ValueString()
	} else {
		data.Path = ""
	}

	//MARSHAL_BODY

	// Nested types Nfs # port
	if !v.Nfs.Port.IsNull() && !v.Nfs.Port.IsUnknown() {
		data.Nfs.Port = new(int64)
		*data.Nfs.Port = v.Nfs.Port.ValueInt64()

	} else {
		data.Nfs.Port = nil
	}

	// Nested types Nfs # limit
	if !v.Nfs.Limit.IsNull() && !v.Nfs.Limit.IsUnknown() {
		data.Nfs.Limit = v.Nfs.Limit.ValueString()
	} else {
		data.Nfs.Limit = ""
	}

	// Nested types Nfs # read_write
	if !v.Nfs.ReadWrite.IsNull() && !v.Nfs.ReadWrite.IsUnknown() {
		data.Nfs.ReadWrite = new(bool)
		*data.Nfs.ReadWrite = v.Nfs.ReadWrite.ValueBool()
	} else {
		data.Nfs.ReadWrite = nil
	}

	// Nested types Nfs # alert_threshold
	if !v.Nfs.AlertThreshold.IsNull() && !v.Nfs.AlertThreshold.IsUnknown() {
		data.Nfs.AlertThreshold = new(int64)
		*data.Nfs.AlertThreshold = v.Nfs.AlertThreshold.ValueInt64()

	} else {
		data.Nfs.AlertThreshold = nil
	}

	//MARSHAL_BODY

	// Nested types ScpSftp # protocol
	if !v.ScpSftp.Protocol.IsNull() && !v.ScpSftp.Protocol.IsUnknown() {
		data.ScpSftp.Protocol = v.ScpSftp.Protocol.ValueString()
	} else {
		data.ScpSftp.Protocol = ""
	}

	// Nested types ScpSftp # port
	if !v.ScpSftp.Port.IsNull() && !v.ScpSftp.Port.IsUnknown() {
		data.ScpSftp.Port = new(int64)
		*data.ScpSftp.Port = v.ScpSftp.Port.ValueInt64()

	} else {
		data.ScpSftp.Port = nil
	}

	// Nested types ScpSftp # username
	if !v.ScpSftp.Username.IsNull() && !v.ScpSftp.Username.IsUnknown() {
		data.ScpSftp.Authentication.Username = v.ScpSftp.Username.ValueString()
	} else {
		data.ScpSftp.Authentication.Username = ""
	}

	// Nested types ScpSftp # password
	if !v.ScpSftp.Password.IsNull() && !v.ScpSftp.Password.IsUnknown() {
		data.ScpSftp.Authentication.Password = v.ScpSftp.Password.ValueString()
	} else {
		data.ScpSftp.Authentication.Password = ""
	}

	// Nested types ScpSftp # ssh_key
	if !v.ScpSftp.SshKey.IsNull() && !v.ScpSftp.SshKey.IsUnknown() {
		data.ScpSftp.Authentication.SshKey = v.ScpSftp.SshKey.ValueString()
	} else {
		data.ScpSftp.Authentication.SshKey = ""
	}

	// Nested types ScpSftp # passphrase
	if !v.ScpSftp.Passphrase.IsNull() && !v.ScpSftp.Passphrase.IsUnknown() {
		data.ScpSftp.Authentication.Passphrase = v.ScpSftp.Passphrase.ValueString()
	} else {
		data.ScpSftp.Authentication.Passphrase = ""
	}

	// Nested types ScpSftp # ignore_host_key_validation
	if !v.ScpSftp.IgnoreHostKeyValidation.IsNull() && !v.ScpSftp.IgnoreHostKeyValidation.IsUnknown() {
		data.ScpSftp.Authentication.IgnoreHostKeyValidation = new(bool)
		*data.ScpSftp.Authentication.IgnoreHostKeyValidation = v.ScpSftp.IgnoreHostKeyValidation.ValueBool()
	} else {
		data.ScpSftp.Authentication.IgnoreHostKeyValidation = nil
	}

	// Nested types ScpSftp # accept_host_key
	if !v.ScpSftp.AcceptHostKey.IsNull() && !v.ScpSftp.AcceptHostKey.IsUnknown() {
		data.ScpSftp.AcceptHostKey = v.ScpSftp.AcceptHostKey.ValueBool()
	}

	return data
}
