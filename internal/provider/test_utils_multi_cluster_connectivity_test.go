// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func MultiClusterConnectivityModelHelperStateCheck(RscName string, c resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.Spec.ClusterType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_type").String(), c.Spec.ClusterType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_type").String(), "ND"))
	}
	if c.Spec.ClusterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_name").String(), c.Spec.ClusterName))
	}
	if c.Spec.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Spec.Hostname))
	}
	if c.Spec.Credentials.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Spec.Credentials.Username))
	}
	if c.Spec.Credentials.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Spec.Credentials.Password))
	}
	if c.Spec.Credentials.LoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("login_domain").String(), c.Spec.Credentials.LoginDomain))
	}
	if c.Spec.Nd.MultiClusterLoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multi_cluster_login_domain").String(), c.Spec.Nd.MultiClusterLoginDomain))
	}
	return ret
}
