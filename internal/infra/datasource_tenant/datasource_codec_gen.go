// Code generated;  DO NOT EDIT.

package datasource_tenant

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTenantModel struct {
	Name               string                                 `json:"name,omitempty"`
	Description        string                                 `json:"description,omitempty"`
	FabricAssociations map[string]NDFCFabricAssociationsValue `json:"-"`
}

type NDFCFabricAssociationsValue struct {
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
	LocalName    string   `json:"localName,omitempty"`
	AllowedVlans []string `json:"allowedVlans,string,omitempty"`
}

func (v *TenantModel) SetModelData(jsonData *NDFCTenantModel) diag.Diagnostics {
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

	if len(jsonData.FabricAssociations) == 0 {
		log.Printf("v.FabricAssociations is empty")
		v.FabricAssociations = types.MapNull(FabricAssociationsValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]FabricAssociationsValue)
		for key, item := range jsonData.FabricAssociations {
			data := new(FabricAssociationsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in FabricAssociationsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.FabricAssociations, err = types.MapValueFrom(context.Background(), FabricAssociationsValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map[string]FabricAssociationsValue to  Map")

		}
	}

	return err
}

func (v *FabricAssociationsValue) SetValue(jsonData *NDFCFabricAssociationsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.TenantPrefix != "" {
		v.TenantPrefix = types.StringValue(jsonData.TenantPrefix)
	} else {
		v.TenantPrefix = types.StringNull()
	}

	if jsonData.LocalName != "" {
		v.LocalName = types.StringValue(jsonData.LocalName)
	} else {
		v.LocalName = types.StringNull()
	}

	if len(jsonData.AllowedVlans) == 0 {
		log.Printf("v.AllowedVlans is empty")
		v.AllowedVlans = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.AllowedVlans))
		for i, item := range jsonData.AllowedVlans {
			listData[i] = types.StringValue(item)
		}
		v.AllowedVlans, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v TenantModel) GetModelData() *NDFCTenantModel {
	var data = new(NDFCTenantModel)

	//MARSHAL_BODY

	if !v.Name.IsNull() && !v.Name.IsUnknown() {
		data.Name = v.Name.ValueString()
	} else {
		data.Name = ""
	}

	return data
}
