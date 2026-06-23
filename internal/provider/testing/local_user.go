// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/infra/resource_local_user"
)

// defaultLocalUserValues returns sensible defaults for a local user.
// Tests can override any of these via the overrides map passed to
// GenerateLocalUserObject.
func defaultLocalUserValues() map[string]interface{} {
	return map[string]interface{}{
		"email":                     "local_user_test@mail.com",
		"first_name":                "first_name",
		"last_name":                 "last_name",
		"remote_user_authorization": false,
		// Note: `remote_id_claim` and `tenant_domain` are intentionally
		// omitted from defaults. `remote_id_claim` must be unique per user
		// (API rejects duplicates with HTTP 500). `tenant_domain` is not
		// echoed back by the GET API today, so setting it produces a
		// "Provider produced inconsistent result after apply" diff.
		// Tests that need either field should pass it explicitly via
		// overrides.
	}
}

// GenerateLocalUserObject creates a local user model object for testing.
// loginID and userPassword are mandatory identifiers / credentials.
// securityDomains is a map of domain name -> list of roles
// (e.g. {"all": {"approver", "designer"}}). At least one entry is required.
// overrides lets each test supply unique values for any field listed in
// defaultLocalUserValues(). Any key not present in overrides gets the
// value from defaultLocalUserValues().
func GenerateLocalUserObject(
	obj **resource_local_user.NDFCLocalUserModel,
	loginID string,
	userPassword string,
	securityDomains map[string][]string,
	overrides map[string]interface{},
) {
	user := new(resource_local_user.NDFCLocalUserModel)

	user.Id = loginID
	user.LoginId = loginID
	user.UserPassword = userPassword

	// Merge defaults with caller overrides (overrides win)
	merged := defaultLocalUserValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyLocalUserValues(user, merged)

	user.Rbac.SecurityDomains = make(map[string]resource_local_user.NDFCSecurityDomainsValue)
	for name, roles := range securityDomains {
		rolesCopy := make([]string, len(roles))
		copy(rolesCopy, roles)
		user.Rbac.SecurityDomains[name] = resource_local_user.NDFCSecurityDomainsValue{
			Roles: rolesCopy,
		}
	}

	*obj = user
}

// ModifyLocalUserObject modifies fields on an existing local user model.
// Uses the same key set as GenerateLocalUserObject overrides.
func ModifyLocalUserObject(
	obj **resource_local_user.NDFCLocalUserModel,
	values map[string]interface{},
) {
	user := *obj
	if user == nil {
		return
	}

	applyLocalUserValues(user, values)

	*obj = user
}

// AddSecurityDomain adds (or replaces) a security domain on an existing
// local user model.
func AddSecurityDomain(
	obj **resource_local_user.NDFCLocalUserModel,
	name string,
	roles []string,
) {
	user := *obj
	if user == nil {
		return
	}
	if user.Rbac.SecurityDomains == nil {
		user.Rbac.SecurityDomains = make(map[string]resource_local_user.NDFCSecurityDomainsValue)
	}
	rolesCopy := make([]string, len(roles))
	copy(rolesCopy, roles)
	user.Rbac.SecurityDomains[name] = resource_local_user.NDFCSecurityDomainsValue{
		Roles: rolesCopy,
	}
	*obj = user
}

// DeleteSecurityDomain removes a security domain by name from an existing
// local user model.
func DeleteSecurityDomain(
	obj **resource_local_user.NDFCLocalUserModel,
	name string,
) {
	user := *obj
	if user == nil || user.Rbac.SecurityDomains == nil {
		return
	}
	delete(user.Rbac.SecurityDomains, name)
	*obj = user
}

// applyLocalUserValues is the shared engine that sets fields on a local user
// model from a key-value map. Used by both Generate and Modify.
func applyLocalUserValues(user *resource_local_user.NDFCLocalUserModel, values map[string]interface{}) {
	for key, val := range values {
		switch key {
		case "email":
			user.Email = val.(string)
		case "first_name":
			user.FirstName = val.(string)
		case "last_name":
			user.LastName = val.(string)
		case "remote_id_claim":
			user.RemoteIdClaim = val.(string)
		case "user_password":
			user.UserPassword = val.(string)
		case "remote_user_authorization":
			v := val.(bool)
			user.RemoteUserAuthorization = &v
		case "tenant_domain":
			user.Rbac.TenantDomain = val.(string)
		}
	}
}
