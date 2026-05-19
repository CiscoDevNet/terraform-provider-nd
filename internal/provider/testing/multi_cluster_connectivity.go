// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
)

// GenerateMultiClusterConnectivityObject builds a fresh
// NDFCMultiClusterConnectivityModel for use in acceptance tests.
//
// Required args mirror the TF schema's required attributes:
//   - hostname
//   - username
//   - password
//
// Optional values supplied via the overrides map (keys are the TF
// attribute names):
//   - "cluster_name"
//   - "login_domain"
//   - "multi_cluster_login_domain"
func GenerateMultiClusterConnectivityObject(
	obj **resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	hostname string,
	username string,
	password string,
	overrides map[string]interface{},
) {
	m := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)
	m.Spec.Hostname = hostname
	m.Spec.Credentials.Username = username
	m.Spec.Credentials.Password = password
	applyMultiClusterConnectivityValues(m, overrides)
	*obj = m
}

// ModifyMultiClusterConnectivityObject mutates an existing model between
// test steps. Keys mirror the TF attribute names; see
// GenerateMultiClusterConnectivityObject.
func ModifyMultiClusterConnectivityObject(
	obj **resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	values map[string]interface{},
) {
	m := *obj
	if m == nil {
		return
	}
	applyMultiClusterConnectivityValues(m, values)
	*obj = m
}

// applyMultiClusterConnectivityValues is the shared key->field map used by
// both Generate and Modify so the two stay in sync.
func applyMultiClusterConnectivityValues(
	m *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "cluster_name":
			m.Spec.ClusterName = val.(string)
		case "hostname":
			m.Spec.Hostname = val.(string)
		case "username":
			m.Spec.Credentials.Username = val.(string)
		case "password":
			m.Spec.Credentials.Password = val.(string)
		case "login_domain":
			m.Spec.Credentials.LoginDomain = val.(string)
		case "multi_cluster_login_domain":
			m.Spec.Nd.MultiClusterLoginDomain = val.(string)
		}
	}
}
