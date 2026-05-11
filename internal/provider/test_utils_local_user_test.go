// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"strconv"

	"terraform-provider-nd/internal/infra/resource_local_user"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func LocalUserModelHelperStateCheck(RscName string, c resource_local_user.NDFCLocalUserModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.LoginId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("login_id").String(), c.LoginId))
	}
	if c.UserPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("user_password").String(), c.UserPassword))
	}
	if c.Email != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("email").String(), c.Email))
	}
	if c.FirstName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("first_name").String(), c.FirstName))
	}
	if c.LastName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("last_name").String(), c.LastName))
	}
	if c.RemoteIdClaim != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_id_claim").String(), c.RemoteIdClaim))
	}
	if c.RemoteUserAuthorization != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_user_authorization").String(), strconv.FormatBool(*c.RemoteUserAuthorization)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_user_authorization").String(), "false"))
	}
	if c.PasswordPolicy.ReuseLimitation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("reuse_limitation").String(), strconv.Itoa(int(*c.PasswordPolicy.ReuseLimitation))))
	}
	if c.PasswordPolicy.TimeIntervalLimitation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("time_interval_limitation").String(), strconv.Itoa(int(*c.PasswordPolicy.TimeIntervalLimitation))))
	}
	if c.Rbac.TenantDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_domain").String(), c.Rbac.TenantDomain))
	}
	for key, value := range c.Rbac.SecurityDomains {
		attrNewPath := attrPath.AtName("security_domains").AtName(key)
		ret = append(ret, SecurityDomainsValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	return ret
}

func SecurityDomainsValueHelperStateCheck(RscName string, c resource_local_user.NDFCSecurityDomainsValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	rolesPath := attrPath.AtName("roles")
	ret = append(ret, resource.TestCheckResourceAttr(RscName, rolesPath.String()+".#", strconv.Itoa(len(c.Roles))))
	for _, role := range c.Roles {
		ret = append(ret, resource.TestCheckTypeSetElemAttr(RscName, rolesPath.String()+".*", role))
	}
	return ret
}
