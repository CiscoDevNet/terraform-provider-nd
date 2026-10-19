// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"fmt"

	"terraform-provider-nd/internal/manage/resource_interface_common"
)

// GenerateInterfaceEthernetObject creates a new NDFCInterfaceCommonModel with the
// given interfaces. Each interface entry is built from the supplied parameters.
func GenerateInterfaceEthernetObject(
	obj **resource_interface_common.NDFCInterfaceCommonModel,
	fabricName string,
	switchId string,
	deploy bool,
	interfaces map[string]InterfaceEthernetEntry,
) {
	model := new(resource_interface_common.NDFCInterfaceCommonModel)
	model.FabricName = fabricName
	model.SwitchId = switchId
	model.Deploy = deploy
	model.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue, len(interfaces))

	for ifName, entry := range interfaces {
		model.Interfaces[ifName] = entry.toNDFCValue()
	}

	*obj = model
}

// InterfaceEthernetEntry is a convenience struct for building interface entries
// in tests. Only the fields that are set (non-zero) will be emitted in the
// generated HCL template.
type InterfaceEthernetEntry struct {
	Mode              string
	NetworkOsType     string
	PolicyType        string
	TemplateName      string
	AdminState        *bool
	Description       string
	Mtu               string
	AccessVlan        *int64
	AllowedVlans      string
	Speed             string
	Fec               string
	BpduGuard         string
	BpduFilter        string
	Cdp               *bool
	PortTypeEdgeTrunk *bool
	Ip                string
	Prefix            string
	RoutingTag        string
	VrfInterface      string
	NativeVlan        *int64
}

func (e InterfaceEthernetEntry) toNDFCValue() resource_interface_common.NDFCInterfacesValue {
	return resource_interface_common.NDFCInterfacesValue{
		ConfigData: resource_interface_common.NDFCConfigDataValue{
			Mode: e.Mode,
			NetworkOS: resource_interface_common.NDFCConfigDataNetworkOSValue{
				NetworkOsType: e.NetworkOsType,
				Policy: resource_interface_common.NDFCNetworkOSPolicyValue{
					PolicyType:        e.PolicyType,
					TemplateName:      e.TemplateName,
					AdminState:        e.AdminState,
					Description:       e.Description,
					Mtu:               e.Mtu,
					AccessVlan:        e.AccessVlan,
					AllowedVlans:      e.AllowedVlans,
					Speed:             e.Speed,
					Fec:               e.Fec,
					BpduGuard:         e.BpduGuard,
					BpduFilter:        e.BpduFilter,
					Cdp:               e.Cdp,
					PortTypeEdgeTrunk: e.PortTypeEdgeTrunk,
					Ip:                e.Ip,
					Prefix:            e.Prefix,
					RoutingTag:        e.RoutingTag,
					VrfInterface:      e.VrfInterface,
					NativeVlan:        e.NativeVlan,
				},
			},
		},
	}
}

// AddInterface adds a new interface entry to an existing model.
func AddInterface(
	obj **resource_interface_common.NDFCInterfaceCommonModel,
	ifName string,
	entry InterfaceEthernetEntry,
) {
	model := *obj
	if model.Interfaces == nil {
		model.Interfaces = make(map[string]resource_interface_common.NDFCInterfacesValue)
	}
	model.Interfaces[ifName] = entry.toNDFCValue()
}

// RemoveInterface removes an interface entry from an existing model.
func RemoveInterface(
	obj **resource_interface_common.NDFCInterfaceCommonModel,
	ifName string,
) {
	model := *obj
	delete(model.Interfaces, ifName)
}

// ModifyInterfaceEthernetObject modifies top-level and per-interface fields on an
// existing model. Supported top-level keys: "deploy". Per-interface keys are
// passed as "Ethernet1/21.access_vlan" (ifName.fieldName).
func ModifyInterfaceEthernetObject(
	obj **resource_interface_common.NDFCInterfaceCommonModel,
	values map[string]interface{},
) {
	model := *obj
	if model == nil {
		return
	}

	for key, val := range values {
		switch key {
		case "deploy":
			model.Deploy = val.(bool)
		case "fabric_name":
			model.FabricName = val.(string)
		case "switch_id":
			model.SwitchId = val.(string)
		default:
			// Expect "ifName.field" format for per-interface modifications.
			ifName, field := splitInterfaceKey(key)
			if ifName == "" {
				continue
			}
			entry, ok := model.Interfaces[ifName]
			if !ok {
				continue
			}
			modifyInterfaceEntry(&entry, field, val)
			model.Interfaces[ifName] = entry
		}
	}
	*obj = model
}

func splitInterfaceKey(key string) (ifName, field string) {
	for i := 0; i < len(key); i++ {
		if key[i] == '.' {
			return key[:i], key[i+1:]
		}
	}
	return "", ""
}

func modifyInterfaceEntry(entry *resource_interface_common.NDFCInterfacesValue, field string, val interface{}) {
	p := &entry.ConfigData.NetworkOS.Policy
	switch field {
	case "mode":
		entry.ConfigData.Mode = val.(string)
	case "network_os_type":
		entry.ConfigData.NetworkOS.NetworkOsType = val.(string)
	case "policy_type":
		p.PolicyType = val.(string)
	case "template_name":
		p.TemplateName = val.(string)
	case "admin_state":
		v := val.(bool)
		p.AdminState = &v
	case "description":
		p.Description = val.(string)
	case "mtu":
		p.Mtu = val.(string)
	case "access_vlan":
		v := int64(val.(int))
		p.AccessVlan = &v
	case "allowed_vlans":
		p.AllowedVlans = val.(string)
	case "speed":
		p.Speed = val.(string)
	case "fec":
		p.Fec = val.(string)
	case "bpdu_guard":
		p.BpduGuard = val.(string)
	case "bpdu_filter":
		p.BpduFilter = val.(string)
	case "cdp":
		v := val.(bool)
		p.Cdp = &v
	case "port_type_edge_trunk":
		v := val.(bool)
		p.PortTypeEdgeTrunk = &v
	case "ip":
		p.Ip = val.(string)
	case "prefix":
		p.Prefix = val.(string)
	case "routing_tag":
		p.RoutingTag = val.(string)
	case "vrf_interface":
		p.VrfInterface = val.(string)
	case "native_vlan":
		v := int64(val.(int))
		p.NativeVlan = &v
	default:
		panic(fmt.Sprintf("modifyInterfaceEntry: unknown field %q", field))
	}
}

// InterfaceEthernetConfigForStep is a convenience wrapper that increments the
// step counter and renders HCL via GetTFConfigWithSingleResource.
func InterfaceEthernetConfigForStep(
	testName string,
	stepCount *int,
	cfg map[string]string,
	model *resource_interface_common.NDFCInterfaceCommonModel,
) string {
	*stepCount++
	tfConfig := new(string)
	stepName := fmt.Sprintf("%s_%d", testName, *stepCount)
	GetTFConfigWithSingleResource(stepName, cfg, []interface{}{model}, &tfConfig)
	return *tfConfig
}
