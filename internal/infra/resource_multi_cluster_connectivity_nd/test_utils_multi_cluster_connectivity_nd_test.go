// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity_nd

import (
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity_nd"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func MultiClusterConnectivityNdModelHelperStateCheck(RscName string, c resource_multi_cluster_connectivity_nd.NDFCMultiClusterConnectivityNdModel, attrPath path.Path) []resource.TestCheckFunc {
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
	if c.Spec.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Spec.Username))
	}
	if c.Spec.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Spec.Password))
	}
	if c.Spec.LoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("login_domain").String(), c.Spec.LoginDomain))
	}
	if c.Spec.MultiClusterLoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multi_cluster_login_domain").String(), c.Spec.MultiClusterLoginDomain))
	}
	return ret
}
