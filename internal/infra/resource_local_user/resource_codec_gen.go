// Code generated;  DO NOT EDIT.

package resource_local_user

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCLocalUserModel struct {
	LoginId                 string        `json:"loginID,omitempty"`
	UserPassword            string        `json:"password,omitempty"`
	Email                   string        `json:"email,omitempty"`
	FirstName               string        `json:"firstName,omitempty"`
	LastName                string        `json:"lastName,omitempty"`
	RemoteIdClaim           string        `json:"remoteIDClaim,omitempty"`
	RemoteUserAuthorization *bool         `json:"xLaunch,omitempty"`
	Rbac                    NDFCRbacValue `json:"rbac,omitempty"`
}

type NDFCRbacValue struct {
	TenantDomain    string                              `json:"tenantDomain,omitempty"`
	SecurityDomains map[string]NDFCSecurityDomainsValue `json:"domains,omitempty"`
}

type NDFCSecurityDomainsValue struct {
	Roles []string `json:"roles,omitempty"`
}

func (v *LocalUserModel) SetModelData(jsonData *NDFCLocalUserModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.LoginId != "" {
		v.LoginId = types.StringValue(jsonData.LoginId)
	} else {
		v.LoginId = types.StringNull()
	}

	if jsonData.UserPassword != "" {
		v.UserPassword = types.StringValue(jsonData.UserPassword)
	} else {
		v.UserPassword = types.StringNull()
	}

	if jsonData.Email != "" {
		v.Email = types.StringValue(jsonData.Email)
	} else {
		v.Email = types.StringNull()
	}

	if jsonData.FirstName != "" {
		v.FirstName = types.StringValue(jsonData.FirstName)
	} else {
		v.FirstName = types.StringNull()
	}

	if jsonData.LastName != "" {
		v.LastName = types.StringValue(jsonData.LastName)
	} else {
		v.LastName = types.StringNull()
	}

	if jsonData.RemoteIdClaim != "" {
		v.RemoteIdClaim = types.StringValue(jsonData.RemoteIdClaim)
	} else {
		v.RemoteIdClaim = types.StringNull()
	}

	if jsonData.RemoteUserAuthorization != nil {
		v.RemoteUserAuthorization = types.BoolValue(*jsonData.RemoteUserAuthorization)

	} else {
		v.RemoteUserAuthorization = types.BoolNull()
	}

	if jsonData.Rbac.TenantDomain != "" {
		v.TenantDomain = types.StringValue(jsonData.Rbac.TenantDomain)

	} else {
		v.TenantDomain = types.StringNull()
	}

	if len(jsonData.Rbac.SecurityDomains) == 0 {
		log.Printf("v.SecurityDomains is empty")
		v.SecurityDomains = types.MapNull(SecurityDomainsValue{}.Type(context.Background()))
	} else {
		mapData := make(map[string]SecurityDomainsValue)
		for key, item := range jsonData.Rbac.SecurityDomains {
			data := new(SecurityDomainsValue)
			err = data.SetValue(&item)
			if err != nil {
				log.Printf("Error in SecurityDomainsValue.SetValue")
				return err
			}
			data.state = attr.ValueStateKnown
			mapData[key] = *data
		}
		v.SecurityDomains, err = types.MapValueFrom(context.Background(), SecurityDomainsValue{}.Type(context.Background()), mapData)
		if err != nil {
			log.Printf("Error in converting map to  Map")
			return err
		}
	}

	return err
}

func (v *SecurityDomainsValue) SetValue(jsonData *NDFCSecurityDomainsValue) diag.Diagnostics {

	var err diag.Diagnostics
	err = nil

	if len(jsonData.Roles) == 0 {
		log.Printf("v.Roles is empty")
		v.Roles = types.SetNull(types.StringType)
		if err != nil {
			log.Printf("Error in converting []string to  List %v", err)
			return err
		}
	} else {
		listData := make([]attr.Value, len(jsonData.Roles))
		for i, item := range jsonData.Roles {
			listData[i] = types.StringValue(item)
		}
		v.Roles, err = types.SetValue(types.StringType, listData)
		if err != nil {
			log.Printf("Error in converting []string to  List")
			return err
		}
	}

	return err
}

func (v LocalUserModel) GetModelData() *NDFCLocalUserModel {
	var data = new(NDFCLocalUserModel)

	//MARSHAL_BODY

	if !v.LoginId.IsNull() && !v.LoginId.IsUnknown() {
		data.LoginId = v.LoginId.ValueString()
	} else {
		data.LoginId = ""
	}

	if !v.UserPassword.IsNull() && !v.UserPassword.IsUnknown() {
		data.UserPassword = v.UserPassword.ValueString()
	} else {
		data.UserPassword = ""
	}

	if !v.Email.IsNull() && !v.Email.IsUnknown() {
		data.Email = v.Email.ValueString()
	} else {
		data.Email = ""
	}

	if !v.FirstName.IsNull() && !v.FirstName.IsUnknown() {
		data.FirstName = v.FirstName.ValueString()
	} else {
		data.FirstName = ""
	}

	if !v.LastName.IsNull() && !v.LastName.IsUnknown() {
		data.LastName = v.LastName.ValueString()
	} else {
		data.LastName = ""
	}

	if !v.RemoteIdClaim.IsNull() && !v.RemoteIdClaim.IsUnknown() {
		data.RemoteIdClaim = v.RemoteIdClaim.ValueString()
	} else {
		data.RemoteIdClaim = ""
	}

	if !v.RemoteUserAuthorization.IsNull() && !v.RemoteUserAuthorization.IsUnknown() {
		data.RemoteUserAuthorization = new(bool)
		*data.RemoteUserAuthorization = v.RemoteUserAuthorization.ValueBool()
	} else {
		data.RemoteUserAuthorization = nil
	}

	if !v.TenantDomain.IsNull() && !v.TenantDomain.IsUnknown() {
		data.Rbac.TenantDomain = v.TenantDomain.ValueString()
	} else {
		data.Rbac.TenantDomain = ""
	}

	if !v.SecurityDomains.IsNull() && !v.SecurityDomains.IsUnknown() {
		elements1 := make(map[string]SecurityDomainsValue, len(v.SecurityDomains.Elements()))

		data.Rbac.SecurityDomains = make(map[string]NDFCSecurityDomainsValue)

		diag := v.SecurityDomains.ElementsAs(context.Background(), &elements1, false)
		if diag != nil {
			panic(diag)
		}
		for k1, ele1 := range elements1 {
			data1 := new(NDFCSecurityDomainsValue)

			// roles | List:String| []| false
			if !ele1.Roles.IsNull() && !ele1.Roles.IsUnknown() {

				listStringData := make([]string, len(ele1.Roles.Elements()))
				dg := ele1.Roles.ElementsAs(context.Background(), &listStringData, false)
				if dg.HasError() {
					panic(dg.Errors())
				}
				data1.Roles = make([]string, len(listStringData))
				copy(data1.Roles, listStringData)
			}

			data.Rbac.SecurityDomains[k1] = *data1

		}
	}

	return data
}
