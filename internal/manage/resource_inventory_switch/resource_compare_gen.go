// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_inventory_switch

import (
	"log"
	. "terraform-provider-nd/internal/common/types"
)

func (v NDFCSwitchesValue) DeepEqual(c NDFCSwitchesValue) int {
	cf := false
	if v.Hostname != c.Hostname {
		log.Printf("v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
		return RequiresUpdate
	}
	if v.IpAddress != c.IpAddress {
		log.Printf("v.IpAddress=%v, c.IpAddress=%v", v.IpAddress, c.IpAddress)
		return RequiresReplace
	}
	if v.Model != c.Model {
		log.Printf("v.Model=%v, c.Model=%v", v.Model, c.Model)
		return RequiresUpdate
	}
	if v.SoftwareVersion != c.SoftwareVersion {
		log.Printf("v.SoftwareVersion=%v, c.SoftwareVersion=%v", v.SoftwareVersion, c.SoftwareVersion)
		return RequiresUpdate
	}
	if v.SwitchRole != c.SwitchRole {
		log.Printf("v.SwitchRole=%v, c.SwitchRole=%v", v.SwitchRole, c.SwitchRole)
		return RequiresReplace
	}
	if v.GatewayIpMask != c.GatewayIpMask {
		log.Printf("v.GatewayIpMask=%v, c.GatewayIpMask=%v", v.GatewayIpMask, c.GatewayIpMask)
		return RequiresUpdate
	}
	if v.DiscoveryAuthProtocol != c.DiscoveryAuthProtocol {
		log.Printf("v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
		return RequiresUpdate
	}
	if v.PoapPassword != c.PoapPassword {
		log.Printf("v.PoapPassword=%v, c.PoapPassword=%v", v.PoapPassword, c.PoapPassword)
		return RequiresUpdate
	}
	if v.ImagePolicy != c.ImagePolicy {
		log.Printf("v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
		return RequiresUpdate
	}
	if v.DiscoveryUsername != c.DiscoveryUsername {
		log.Printf("v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
		return RequiresUpdate
	}
	if v.DiscoveryPassword != c.DiscoveryPassword {
		log.Printf("v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
		return RequiresUpdate
	}

	if len(v.ModuleModels) != len(c.ModuleModels) {
		log.Printf("len(v.ModuleModels)=%d, len(c.ModuleModels)=%d", len(v.ModuleModels), len(c.ModuleModels))
		return PortListUpdate
	}
	for i := range v.ModuleModels {
		if v.ModuleModels[i] != c.ModuleModels[i] {
			log.Printf("v.ModuleModels[%d]=%s, c.ModuleModels[%d]=%s", i, v.ModuleModels[i], i, c.ModuleModels[i])
			return PortListUpdate
		}
	}

	if v.VdcId != nil && c.VdcId != nil {
		if *v.VdcId != *c.VdcId {
			log.Printf("v.VdcId=%v, c.VdcId=%v", *v.VdcId, *c.VdcId)
			return RequiresUpdate
		}
	} else {
		if v.VdcId != nil {
			log.Printf("v.VdcId=%v", *v.VdcId)
			return RequiresUpdate
		} else if c.VdcId != nil {
			log.Printf("c.VdcId=%v", *c.VdcId)
			return RequiresUpdate
		}
	}

	if v.VdcMac != c.VdcMac {
		log.Printf("v.VdcMac=%v, c.VdcMac=%v", v.VdcMac, c.VdcMac)
		return RequiresUpdate
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCSwitchesValue) CreatePlan(c NDFCSwitchesValue, cf *bool) int {
	action := ActionNone

	if v.Hostname != c.Hostname {
		log.Printf("Hostname-Update: v.Hostname=%v, c.Hostname=%v", v.Hostname, c.Hostname)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.IpAddress != c.IpAddress {
		log.Printf("IpAddress-Update: v.IpAddress=%v, c.IpAddress=%v", v.IpAddress, c.IpAddress)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.Model != c.Model {
		log.Printf("Model-Update: v.Model=%v, c.Model=%v", v.Model, c.Model)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.SoftwareVersion != c.SoftwareVersion {
		log.Printf("SoftwareVersion-Update: v.SoftwareVersion=%v, c.SoftwareVersion=%v", v.SoftwareVersion, c.SoftwareVersion)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.SwitchRole != c.SwitchRole {
		log.Printf("SwitchRole-Update: v.SwitchRole=%v, c.SwitchRole=%v", v.SwitchRole, c.SwitchRole)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.GatewayIpMask != c.GatewayIpMask {
		log.Printf("GatewayIpMask-Update: v.GatewayIpMask=%v, c.GatewayIpMask=%v", v.GatewayIpMask, c.GatewayIpMask)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryAuthProtocol != c.DiscoveryAuthProtocol {
		log.Printf("DiscoveryAuthProtocol-Update: v.DiscoveryAuthProtocol=%v, c.DiscoveryAuthProtocol=%v", v.DiscoveryAuthProtocol, c.DiscoveryAuthProtocol)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.PoapPassword != c.PoapPassword {
		log.Printf("PoapPassword-Update: v.PoapPassword=%v, c.PoapPassword=%v", v.PoapPassword, c.PoapPassword)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ImagePolicy != c.ImagePolicy {
		log.Printf("ImagePolicy-Update: v.ImagePolicy=%v, c.ImagePolicy=%v", v.ImagePolicy, c.ImagePolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryUsername != c.DiscoveryUsername {
		log.Printf("DiscoveryUsername-Update: v.DiscoveryUsername=%v, c.DiscoveryUsername=%v", v.DiscoveryUsername, c.DiscoveryUsername)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.DiscoveryPassword != c.DiscoveryPassword {
		log.Printf("DiscoveryPassword-Update: v.DiscoveryPassword=%v, c.DiscoveryPassword=%v", v.DiscoveryPassword, c.DiscoveryPassword)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if len(v.ModuleModels) != len(c.ModuleModels) {
		log.Printf("Update: len(v.ModuleModels)=%d, len(c.ModuleModels)=%d", len(v.ModuleModels), len(c.ModuleModels))
		return RequiresUpdate
	}
	for i := range v.ModuleModels {
		if v.ModuleModels[i] != c.ModuleModels[i] {
			log.Printf("Update: v.ModuleModels[%d]=%s, c.ModuleModels[%d]=%s", i, v.ModuleModels[i], i, c.ModuleModels[i])
			return RequiresUpdate
		}
	}

	if v.VdcId != nil && c.VdcId != nil {
		if *v.VdcId != *c.VdcId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.VdcId=%v, c.VdcId=%v", *v.VdcId, *c.VdcId)
		}
	} else if v.VdcId != nil {
		log.Printf("Update: v.VdcId=%v, c.VdcId=nil", *v.VdcId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.VdcId != nil {
		v.VdcId = new(int64)
		log.Printf("Copy from state: v.VdcId=nil, c.VdcId=%v", *c.VdcId)
		*v.VdcId = *c.VdcId
	}
	if v.VdcMac != c.VdcMac {
		log.Printf("VdcMac-Update: v.VdcMac=%v, c.VdcMac=%v", v.VdcMac, c.VdcMac)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	return action
}
