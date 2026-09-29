// Code generated;  DO NOT EDIT.

package datasource_multi_cluster_connectivity

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCMultiClusterConnectivityModel struct {
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type NDFCSpecValue struct {
	ClusterName string `json:"name,omitempty"`
	Hostname    string `json:"onboardUrl,omitempty"`
}

func (v *MultiClusterConnectivityModel) SetModelData(jsonData *NDFCMultiClusterConnectivityModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

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

	return err
}

func (v MultiClusterConnectivityModel) GetModelData() *NDFCMultiClusterConnectivityModel {
	var data = new(NDFCMultiClusterConnectivityModel)

	//MARSHAL_BODY

	if !v.ClusterName.IsNull() && !v.ClusterName.IsUnknown() {
		data.Spec.ClusterName = v.ClusterName.ValueString()
	} else {
		data.Spec.ClusterName = ""
	}

	return data
}
