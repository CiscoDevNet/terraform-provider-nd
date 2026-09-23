// Code generated;  DO NOT EDIT.

package datasource_tenant_domain

import (
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTenantDomainModel struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	TenantNames []string `json:"tenantNames,omitempty"`
}

func (v *TenantDomainModel) SetModelData(jsonData *NDFCTenantDomainModel) diag.Diagnostics {
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

	if len(jsonData.TenantNames) == 0 {
		log.Printf("v.TenantNames is empty")
		v.TenantNames = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.TenantNames))
		for i, item := range jsonData.TenantNames {
			listData[i] = types.StringValue(item)
		}
		v.TenantNames, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v TenantDomainModel) GetModelData() *NDFCTenantDomainModel {
	var data = new(NDFCTenantDomainModel)

	//MARSHAL_BODY

	if !v.Name.IsNull() && !v.Name.IsUnknown() {
		data.Name = v.Name.ValueString()
	} else {
		data.Name = ""
	}

	return data
}
