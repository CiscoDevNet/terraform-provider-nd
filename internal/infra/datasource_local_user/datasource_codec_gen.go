// Code generated;  DO NOT EDIT.

package datasource_local_user

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCLocalUserModel struct {
	Id                      string        `json:"-"`
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

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

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

	return data
}
