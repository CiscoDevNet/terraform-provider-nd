// Code generated;  DO NOT EDIT.

package resource_tenant

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCTenantModel struct {
	Id                 string                        `json:"-"`
	Name               string                        `json:"name,omitempty"`
	Description        string                        `json:"description,omitempty"`
	FabricAssociations []NDFCFabricAssociationsValue `json:"-"`
}

type NDFCFabricAssociationsValue struct {
	FabricName   string   `json:"fabricName,omitempty"`
	LocalName    string   `json:"localName,omitempty"`
	TenantPrefix string   `json:"tenantPrefix,omitempty"`
	AllowedVlans []string `json:"allowedVlans,string,omitempty"`
}

func (v *TenantModel) SetModelData(jsonData *NDFCTenantModel) diag.Diagnostics {
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

	if len(jsonData.FabricAssociations) == 0 {
		log.Printf("v.FabricAssociations is empty")
		v.FabricAssociations = types.SetNull(FabricAssociationsValue{}.Type(context.Background()))
	} else {
		log.Printf("v.FabricAssociations contains %d elements", len(jsonData.FabricAssociations))
		listData := make([]FabricAssociationsValue, 0)
		for _, item := range jsonData.FabricAssociations {
			data := new(FabricAssociationsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in FabricAssociationsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			listData = append(listData, *data)
		}
		v.FabricAssociations, err = types.SetValueFrom(context.Background(), FabricAssociationsValue{}.Type(context.Background()), listData)
		if err != nil {
			log.Printf("Error in converting []FabricAssociationsValue to  List")
			return err
		}
	}

	return err
}

func (v *FabricAssociationsValue) SetValue(jsonData *NDFCFabricAssociationsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.LocalName != "" {
		v.LocalName = types.StringValue(jsonData.LocalName)
	} else {
		v.LocalName = types.StringNull()
	}

	if jsonData.TenantPrefix != "" {
		v.TenantPrefix = types.StringValue(jsonData.TenantPrefix)
	} else {
		v.TenantPrefix = types.StringNull()
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

	if !v.Description.IsNull() && !v.Description.IsUnknown() {
		data.Description = v.Description.ValueString()
	} else {
		data.Description = ""
	}

	//MARSHALL_LIST  BEGIN FabricAssociations[i1]

	if !v.FabricAssociations.IsNull() && !v.FabricAssociations.IsUnknown() {
		elements := make([]FabricAssociationsValue, len(v.FabricAssociations.Elements()))
		// Not augmenting

		// -- Set here 2 |.FabricAssociations|FabricAssociations[i1]| --
		data.FabricAssociations = make([]NDFCFabricAssociationsValue, len(v.FabricAssociations.Elements()))

		diag := v.FabricAssociations.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .FabricAssociations[i1] FabricAssociations[i1]
		for i1, ele1 := range elements {
			if !ele1.FabricName.IsNull() && !ele1.FabricName.IsUnknown() {

				data.FabricAssociations[i1].FabricName = ele1.FabricName.ValueString()
			} else {
				data.FabricAssociations[i1].FabricName = ""
			}

			if !ele1.LocalName.IsNull() && !ele1.LocalName.IsUnknown() {

				data.FabricAssociations[i1].LocalName = ele1.LocalName.ValueString()
			} else {
				data.FabricAssociations[i1].LocalName = ""
			}

			if !ele1.TenantPrefix.IsNull() && !ele1.TenantPrefix.IsUnknown() {

				data.FabricAssociations[i1].TenantPrefix = ele1.TenantPrefix.ValueString()
			} else {
				data.FabricAssociations[i1].TenantPrefix = ""
			}

			if !ele1.AllowedVlans.IsNull() && !ele1.AllowedVlans.IsUnknown() {

				// Nested List:String inside a list - which is not having NDFCNested |.FabricAssociations[i1]|[]|allowed_vlans|
				listStringData := make([]string, len(ele1.AllowedVlans.Elements()))
				dg := ele1.AllowedVlans.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data.FabricAssociations[i1].AllowedVlans = make([]string, len(listStringData))
				copy(data.FabricAssociations[i1].AllowedVlans, listStringData)
			}

		} /* for loop */
	} /* isNull if check */

	return data
}
