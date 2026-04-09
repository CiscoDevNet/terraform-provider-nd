// Code generated;  DO NOT EDIT.

package resource_inventory_switch

import (
	"strconv"
	"terraform-provider-nd/internal/manage/resource_inventory_switch"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func InventorySwitchModelHelperStateCheck(RscName string, c resource_inventory_switch.NDFCInventorySwitchModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.Mode != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), c.Mode))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mode").String(), "discovery"))
	}

	if c.PreserveConfig != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), strconv.FormatBool(*c.PreserveConfig)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("preserve_config").String(), "false"))
	}
	for key, value := range c.Switches {
		attrNewPath := attrPath.AtName("switches").AtName(key)
		ret = append(ret, SwitchesValueHelperStateCheck(RscName, value, attrNewPath)...)
	}
	if c.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Username))
	}
	if c.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Password))
	}
	if c.PlatformType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("platform_type").String(), c.PlatformType))
	}
	if c.SnmpV3AuthProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_v3_auth_protocol").String(), c.SnmpV3AuthProtocol))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("snmp_v3_auth_protocol").String(), "MD5"))
	}
	if c.RemoteCredentialStore != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store").String(), c.RemoteCredentialStore))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store").String(), "local"))
	}
	if c.RemoteCredentialStoreKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("remote_credential_store_key").String(), c.RemoteCredentialStoreKey))
	}
	if c.MaxHop != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_hop").String(), strconv.Itoa(int(*c.MaxHop))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("max_hop").String(), "0"))
	}
	if c.Recalculate {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("recalculate").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("recalculate").String(), "false"))
	}
	if c.Deploy {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "true"))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("deploy").String(), "false"))
	}
	return ret
}

func SwitchesValueHelperStateCheck(RscName string, c resource_inventory_switch.NDFCSwitchesValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Hostname))
	}
	if c.SerialNumber != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("serial_number").String(), c.SerialNumber))
	}
	if c.IpAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip_address").String(), c.IpAddress))
	}
	if c.Model != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("model").String(), c.Model))
	}
	if c.SoftwareVersion != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("software_version").String(), c.SoftwareVersion))
	}
	if c.SwitchRole != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_role").String(), c.SwitchRole))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("switch_role").String(), "leaf"))
	}

	if c.Status != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("status").String(), c.Status))
	}
	if c.StatusReason != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("status_reason").String(), c.StatusReason))
	}
	if c.GatewayIpMask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("gateway_ip_mask").String(), c.GatewayIpMask))
	}
	if c.DiscoveryAuthProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("discovery_auth_protocol").String(), c.DiscoveryAuthProtocol))
	}
	if c.PoapPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("poap_password").String(), c.PoapPassword))
	}
	if c.VdcId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vdc_id").String(), strconv.Itoa(int(*c.VdcId))))
	}
	if c.VdcMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vdc_mac").String(), c.VdcMac))
	}
	return ret
}
