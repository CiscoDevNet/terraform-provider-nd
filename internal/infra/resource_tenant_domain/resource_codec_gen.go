// Code generated;  DO NOT EDIT.

package resource_tenant_domain

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTenantDomainModel struct {
	Id          string   `json:"-"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	TenantNames []string `json:"tenantNames"`
}

func (v *TenantDomainModel) SetModelData(jsonData *NDFCTenantDomainModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

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

	if !v.Id.IsNull() && !v.Id.IsUnknown() {
		data.Id = v.Id.ValueString()
	} else {
		data.Id = ""
	}

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

	if !v.TenantNames.IsNull() && !v.TenantNames.IsUnknown() {
		listStringData := make([]string, len(v.TenantNames.Elements()))
		dg := v.TenantNames.ElementsAs(context.Background(), &listStringData, false)
		if dg.HasError() {
			panic(dg.Errors())
		}
		data.TenantNames = make([]string, len(listStringData))
		copy(data.TenantNames, listStringData)
	}

	return data
}
