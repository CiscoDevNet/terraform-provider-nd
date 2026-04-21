// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity_nd

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCMultiClusterConnectivityNdModel struct {
	Id   string        `json:"-"`
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type NDFCSpecValue struct {
	ClusterType string                   `json:"clusterType,omitempty"`
	ClusterName string                   `json:"name,omitempty"`
	Hostname    string                   `json:"onboardUrl,omitempty"`
	Credentials NDFCSpecCredentialsValue `json:"credentials,omitempty"`
	Nd          NDFCSpecNdValue          `json:"nd,omitempty"`
}

type NDFCSpecCredentialsValue struct {
	Username    string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	LoginDomain string `json:"loginDomain,omitempty"`
}

type NDFCSpecNdValue struct {
	MultiClusterLoginDomain string `json:"multiClusterLoginDomainName,omitempty"`
}

func (v *MultiClusterConnectivityNdModel) SetModelData(jsonData *NDFCMultiClusterConnectivityNdModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.Spec.ClusterType != "" {
		v.ClusterType = types.StringValue(jsonData.Spec.ClusterType)

	} else {
		v.ClusterType = types.StringNull()
	}

	if jsonData.Spec.ClusterName != "" {
		v.ClusterName = types.StringValue(jsonData.Spec.ClusterName)

	} else {
		v.ClusterName = types.StringNull()
	}

	if jsonData.Spec.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Spec.Hostname)

	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.Spec.Credentials.Username != "" {
		v.Username = types.StringValue(jsonData.Spec.Credentials.Username)

	} else {
		v.Username = types.StringNull()
	}

	if jsonData.Spec.Credentials.Password != "" {
		v.Password = types.StringValue(jsonData.Spec.Credentials.Password)

	} else {
		v.Password = types.StringNull()
	}

	if jsonData.Spec.Credentials.LoginDomain != "" {
		v.LoginDomain = types.StringValue(jsonData.Spec.Credentials.LoginDomain)

	} else {
		v.LoginDomain = types.StringNull()
	}

	if jsonData.Spec.Nd.MultiClusterLoginDomain != "" {
		v.MultiClusterLoginDomain = types.StringValue(jsonData.Spec.Nd.MultiClusterLoginDomain)

	} else {
		v.MultiClusterLoginDomain = types.StringNull()
	}

	return err
}

func (v MultiClusterConnectivityNdModel) GetModelData() *NDFCMultiClusterConnectivityNdModel {
	var data = new(NDFCMultiClusterConnectivityNdModel)

	//MARSHAL_BODY

	if !v.ClusterType.IsNull() && !v.ClusterType.IsUnknown() {
		data.Spec.ClusterType = v.ClusterType.ValueString()
	} else {
		data.Spec.ClusterType = ""
	}

	if !v.ClusterName.IsNull() && !v.ClusterName.IsUnknown() {
		data.Spec.ClusterName = v.ClusterName.ValueString()
	} else {
		data.Spec.ClusterName = ""
	}

	if !v.Hostname.IsNull() && !v.Hostname.IsUnknown() {
		data.Spec.Hostname = v.Hostname.ValueString()
	} else {
		data.Spec.Hostname = ""
	}

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Spec.Credentials.Username = v.Username.ValueString()
	} else {
		data.Spec.Credentials.Username = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Spec.Credentials.Password = v.Password.ValueString()
	} else {
		data.Spec.Credentials.Password = ""
	}

	if !v.LoginDomain.IsNull() && !v.LoginDomain.IsUnknown() {
		data.Spec.Credentials.LoginDomain = v.LoginDomain.ValueString()
	} else {
		data.Spec.Credentials.LoginDomain = ""
	}

	if !v.MultiClusterLoginDomain.IsNull() && !v.MultiClusterLoginDomain.IsUnknown() {
		data.Spec.Nd.MultiClusterLoginDomain = v.MultiClusterLoginDomain.ValueString()
	} else {
		data.Spec.Nd.MultiClusterLoginDomain = ""
	}

	return data
}
