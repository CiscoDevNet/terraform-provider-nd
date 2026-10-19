// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_interface_common

import (
	"log"
	. "terraform-provider-nd/internal/common/types"
)

func (v NDFCAssociationListValue) DeepEqual(c NDFCAssociationListValue) int {
	cf := false

	if v.PVlan != nil && c.PVlan != nil {
		if *v.PVlan != *c.PVlan {
			log.Printf("v.PVlan=%v, c.PVlan=%v", *v.PVlan, *c.PVlan)
			return RequiresUpdate
		}
	} else {
		if v.PVlan != nil {
			log.Printf("v.PVlan=%v", *v.PVlan)
			return RequiresUpdate
		} else if c.PVlan != nil {
			log.Printf("c.PVlan=%v", *c.PVlan)
			return RequiresUpdate
		}
	}

	if v.SVlan != nil && c.SVlan != nil {
		if *v.SVlan != *c.SVlan {
			log.Printf("v.SVlan=%v, c.SVlan=%v", *v.SVlan, *c.SVlan)
			return RequiresUpdate
		}
	} else {
		if v.SVlan != nil {
			log.Printf("v.SVlan=%v", *v.SVlan)
			return RequiresUpdate
		} else if c.SVlan != nil {
			log.Printf("c.SVlan=%v", *c.SVlan)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCMappingListValue) DeepEqual(c NDFCMappingListValue) int {
	cf := false

	if v.PVlan != nil && c.PVlan != nil {
		if *v.PVlan != *c.PVlan {
			log.Printf("v.PVlan=%v, c.PVlan=%v", *v.PVlan, *c.PVlan)
			return RequiresUpdate
		}
	} else {
		if v.PVlan != nil {
			log.Printf("v.PVlan=%v", *v.PVlan)
			return RequiresUpdate
		} else if c.PVlan != nil {
			log.Printf("c.PVlan=%v", *c.PVlan)
			return RequiresUpdate
		}
	}

	if v.SVlan != nil && c.SVlan != nil {
		if *v.SVlan != *c.SVlan {
			log.Printf("v.SVlan=%v, c.SVlan=%v", *v.SVlan, *c.SVlan)
			return RequiresUpdate
		}
	} else {
		if v.SVlan != nil {
			log.Printf("v.SVlan=%v", *v.SVlan)
			return RequiresUpdate
		} else if c.SVlan != nil {
			log.Printf("c.SVlan=%v", *c.SVlan)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCIgmpGroupRangesValue) DeepEqual(c NDFCIgmpGroupRangesValue) int {
	cf := false
	if v.GroupAddress != c.GroupAddress {
		log.Printf("v.GroupAddress=%v, c.GroupAddress=%v", v.GroupAddress, c.GroupAddress)
		return RequiresUpdate
	}

	if v.Prefix != nil && c.Prefix != nil {
		if *v.Prefix != *c.Prefix {
			log.Printf("v.Prefix=%v, c.Prefix=%v", *v.Prefix, *c.Prefix)
			return RequiresUpdate
		}
	} else {
		if v.Prefix != nil {
			log.Printf("v.Prefix=%v", *v.Prefix)
			return RequiresUpdate
		} else if c.Prefix != nil {
			log.Printf("c.Prefix=%v", *c.Prefix)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCMstPathCostListValue) DeepEqual(c NDFCMstPathCostListValue) int {
	cf := false

	if len(v.MstInstanceList) != len(c.MstInstanceList) {
		log.Printf("len(v.MstInstanceList)=%d, len(c.MstInstanceList)=%d", len(v.MstInstanceList), len(c.MstInstanceList))
		return PortListUpdate
	}
	for i := range v.MstInstanceList {
		if v.MstInstanceList[i] != c.MstInstanceList[i] {
			log.Printf("v.MstInstanceList[%d]=%s, c.MstInstanceList[%d]=%s", i, v.MstInstanceList[i], i, c.MstInstanceList[i])
			return PortListUpdate
		}
	}

	if v.PathCost != nil && c.PathCost != nil {
		if *v.PathCost != *c.PathCost {
			log.Printf("v.PathCost=%v, c.PathCost=%v", *v.PathCost, *c.PathCost)
			return RequiresUpdate
		}
	} else {
		if v.PathCost != nil {
			log.Printf("v.PathCost=%v", *v.PathCost)
			return RequiresUpdate
		} else if c.PathCost != nil {
			log.Printf("c.PathCost=%v", *c.PathCost)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCMstPortPrioListValue) DeepEqual(c NDFCMstPortPrioListValue) int {
	cf := false

	if len(v.MstInstanceList) != len(c.MstInstanceList) {
		log.Printf("len(v.MstInstanceList)=%d, len(c.MstInstanceList)=%d", len(v.MstInstanceList), len(c.MstInstanceList))
		return PortListUpdate
	}
	for i := range v.MstInstanceList {
		if v.MstInstanceList[i] != c.MstInstanceList[i] {
			log.Printf("v.MstInstanceList[%d]=%s, c.MstInstanceList[%d]=%s", i, v.MstInstanceList[i], i, c.MstInstanceList[i])
			return PortListUpdate
		}
	}

	if v.PortPriority != nil && c.PortPriority != nil {
		if *v.PortPriority != *c.PortPriority {
			log.Printf("v.PortPriority=%v, c.PortPriority=%v", *v.PortPriority, *c.PortPriority)
			return RequiresUpdate
		}
	} else {
		if v.PortPriority != nil {
			log.Printf("v.PortPriority=%v", *v.PortPriority)
			return RequiresUpdate
		} else if c.PortPriority != nil {
			log.Printf("c.PortPriority=%v", *c.PortPriority)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCVlanMappingEntriesValue) DeepEqual(c NDFCVlanMappingEntriesValue) int {
	cf := false

	if len(v.CustomerVlanId) != len(c.CustomerVlanId) {
		log.Printf("len(v.CustomerVlanId)=%d, len(c.CustomerVlanId)=%d", len(v.CustomerVlanId), len(c.CustomerVlanId))
		return PortListUpdate
	}
	for i := range v.CustomerVlanId {
		if v.CustomerVlanId[i] != c.CustomerVlanId[i] {
			log.Printf("v.CustomerVlanId[%d]=%s, c.CustomerVlanId[%d]=%s", i, v.CustomerVlanId[i], i, c.CustomerVlanId[i])
			return PortListUpdate
		}
	}
	if v.Dot1qTunnel != c.Dot1qTunnel {
		log.Printf("v.Dot1qTunnel=%v, c.Dot1qTunnel=%v", v.Dot1qTunnel, c.Dot1qTunnel)
		return RequiresUpdate
	}

	if v.CustomerInnerVlanId != nil && c.CustomerInnerVlanId != nil {
		if *v.CustomerInnerVlanId != *c.CustomerInnerVlanId {
			log.Printf("v.CustomerInnerVlanId=%v, c.CustomerInnerVlanId=%v", *v.CustomerInnerVlanId, *c.CustomerInnerVlanId)
			return RequiresUpdate
		}
	} else {
		if v.CustomerInnerVlanId != nil {
			log.Printf("v.CustomerInnerVlanId=%v", *v.CustomerInnerVlanId)
			return RequiresUpdate
		} else if c.CustomerInnerVlanId != nil {
			log.Printf("c.CustomerInnerVlanId=%v", *c.CustomerInnerVlanId)
			return RequiresUpdate
		}
	}

	if v.ProviderVlanId != nil && c.ProviderVlanId != nil {
		if *v.ProviderVlanId != *c.ProviderVlanId {
			log.Printf("v.ProviderVlanId=%v, c.ProviderVlanId=%v", *v.ProviderVlanId, *c.ProviderVlanId)
			return RequiresUpdate
		}
	} else {
		if v.ProviderVlanId != nil {
			log.Printf("v.ProviderVlanId=%v", *v.ProviderVlanId)
			return RequiresUpdate
		} else if c.ProviderVlanId != nil {
			log.Printf("c.ProviderVlanId=%v", *c.ProviderVlanId)
			return RequiresUpdate
		}
	}

	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v NDFCInterfacesValue) DeepEqual(c NDFCInterfacesValue) int {
	cf := false
	if v.SwitchId != c.SwitchId {
		log.Printf("v.SwitchId=%v, c.SwitchId=%v", v.SwitchId, c.SwitchId)
		return RequiresReplace
	}
	if v.ConfigData.Mode != c.ConfigData.Mode {
		log.Printf("v.ConfigData.Mode=%v, c.ConfigData.Mode=%v", v.ConfigData.Mode, c.ConfigData.Mode)
		return RequiresReplace
	}
	if v.ConfigData.NetworkOS.NetworkOsType != c.ConfigData.NetworkOS.NetworkOsType {
		log.Printf("v.ConfigData.NetworkOS.NetworkOsType=%v, c.ConfigData.NetworkOS.NetworkOsType=%v", v.ConfigData.NetworkOS.NetworkOsType, c.ConfigData.NetworkOS.NetworkOsType)
		return RequiresReplace
	}
	if v.ConfigData.NetworkOS.Policy.PolicyType != c.ConfigData.NetworkOS.Policy.PolicyType {
		log.Printf("v.ConfigData.NetworkOS.Policy.PolicyType=%v, c.ConfigData.NetworkOS.Policy.PolicyType=%v", v.ConfigData.NetworkOS.Policy.PolicyType, c.ConfigData.NetworkOS.Policy.PolicyType)
		return RequiresUpdate
	}
	if v.ConfigData.NetworkOS.Policy.TemplateName != "" {
		if v.ConfigData.NetworkOS.Policy.TemplateName != c.ConfigData.NetworkOS.Policy.TemplateName {
			log.Printf("v.ConfigData.NetworkOS.Policy.TemplateName=%v, c.ConfigData.NetworkOS.Policy.TemplateName=%v", v.ConfigData.NetworkOS.Policy.TemplateName, c.ConfigData.NetworkOS.Policy.TemplateName)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.TemplateName=%v, c.ConfigData.NetworkOS.Policy.TemplateName=%v", v.ConfigData.NetworkOS.Policy.TemplateName, c.ConfigData.NetworkOS.Policy.TemplateName)
	}

	if len(v.ConfigData.NetworkOS.Policy.TemplateConfig) != len(c.ConfigData.NetworkOS.Policy.TemplateConfig) {
		log.Printf("v.ConfigData.NetworkOS.Policy.TemplateConfig len=%d, c.ConfigData.NetworkOS.Policy.TemplateConfig len=%d", len(v.ConfigData.NetworkOS.Policy.TemplateConfig), len(c.ConfigData.NetworkOS.Policy.TemplateConfig))
		return RequiresUpdate
	}
	for kk, vv := range v.ConfigData.NetworkOS.Policy.TemplateConfig {
		cc, ok := c.ConfigData.NetworkOS.Policy.TemplateConfig[kk]
		if !ok || vv != cc {
			log.Printf("v.ConfigData.NetworkOS.Policy.TemplateConfig[%s]=%s, c=%s (ok=%v)", kk, vv, cc, ok)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.AdminState != nil && c.ConfigData.NetworkOS.Policy.AdminState != nil {
		if *v.ConfigData.NetworkOS.Policy.AdminState != *c.ConfigData.NetworkOS.Policy.AdminState {
			log.Printf("v.ConfigData.NetworkOS.Policy.AdminState=%v, c.ConfigData.NetworkOS.Policy.AdminState=%v", *v.ConfigData.NetworkOS.Policy.AdminState, *c.ConfigData.NetworkOS.Policy.AdminState)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.AdminState != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.AdminState=%v", *v.ConfigData.NetworkOS.Policy.AdminState)
		} else if c.ConfigData.NetworkOS.Policy.AdminState != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.AdminState=%v", *c.ConfigData.NetworkOS.Policy.AdminState)
		}
	}
	if v.ConfigData.NetworkOS.Policy.Description != "" {
		if v.ConfigData.NetworkOS.Policy.Description != c.ConfigData.NetworkOS.Policy.Description {
			log.Printf("v.ConfigData.NetworkOS.Policy.Description=%v, c.ConfigData.NetworkOS.Policy.Description=%v", v.ConfigData.NetworkOS.Policy.Description, c.ConfigData.NetworkOS.Policy.Description)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Description=%v, c.ConfigData.NetworkOS.Policy.Description=%v", v.ConfigData.NetworkOS.Policy.Description, c.ConfigData.NetworkOS.Policy.Description)
	}
	if v.ConfigData.NetworkOS.Policy.Mtu != "" {
		if v.ConfigData.NetworkOS.Policy.Mtu != c.ConfigData.NetworkOS.Policy.Mtu {
			log.Printf("v.ConfigData.NetworkOS.Policy.Mtu=%v, c.ConfigData.NetworkOS.Policy.Mtu=%v", v.ConfigData.NetworkOS.Policy.Mtu, c.ConfigData.NetworkOS.Policy.Mtu)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Mtu=%v, c.ConfigData.NetworkOS.Policy.Mtu=%v", v.ConfigData.NetworkOS.Policy.Mtu, c.ConfigData.NetworkOS.Policy.Mtu)
	}
	if v.ConfigData.NetworkOS.Policy.ExtraConfig != "" {
		if v.ConfigData.NetworkOS.Policy.ExtraConfig != c.ConfigData.NetworkOS.Policy.ExtraConfig {
			log.Printf("v.ConfigData.NetworkOS.Policy.ExtraConfig=%v, c.ConfigData.NetworkOS.Policy.ExtraConfig=%v", v.ConfigData.NetworkOS.Policy.ExtraConfig, c.ConfigData.NetworkOS.Policy.ExtraConfig)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.ExtraConfig=%v, c.ConfigData.NetworkOS.Policy.ExtraConfig=%v", v.ConfigData.NetworkOS.Policy.ExtraConfig, c.ConfigData.NetworkOS.Policy.ExtraConfig)
	}

	if v.ConfigData.NetworkOS.Policy.AccessVlan != nil && c.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.AccessVlan != *c.ConfigData.NetworkOS.Policy.AccessVlan {
			log.Printf("v.ConfigData.NetworkOS.Policy.AccessVlan=%v, c.ConfigData.NetworkOS.Policy.AccessVlan=%v", *v.ConfigData.NetworkOS.Policy.AccessVlan, *c.ConfigData.NetworkOS.Policy.AccessVlan)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.AccessVlan != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.AccessVlan=%v", *v.ConfigData.NetworkOS.Policy.AccessVlan)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.AccessVlan != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.AccessVlan=%v", *c.ConfigData.NetworkOS.Policy.AccessVlan)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.AllowedVlans != "" {
		if v.ConfigData.NetworkOS.Policy.AllowedVlans != c.ConfigData.NetworkOS.Policy.AllowedVlans {
			log.Printf("v.ConfigData.NetworkOS.Policy.AllowedVlans=%v, c.ConfigData.NetworkOS.Policy.AllowedVlans=%v", v.ConfigData.NetworkOS.Policy.AllowedVlans, c.ConfigData.NetworkOS.Policy.AllowedVlans)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.AllowedVlans=%v, c.ConfigData.NetworkOS.Policy.AllowedVlans=%v", v.ConfigData.NetworkOS.Policy.AllowedVlans, c.ConfigData.NetworkOS.Policy.AllowedVlans)
	}

	if v.ConfigData.NetworkOS.Policy.Bandwidth != nil && c.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		if *v.ConfigData.NetworkOS.Policy.Bandwidth != *c.ConfigData.NetworkOS.Policy.Bandwidth {
			log.Printf("v.ConfigData.NetworkOS.Policy.Bandwidth=%v, c.ConfigData.NetworkOS.Policy.Bandwidth=%v", *v.ConfigData.NetworkOS.Policy.Bandwidth, *c.ConfigData.NetworkOS.Policy.Bandwidth)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Bandwidth != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Bandwidth=%v", *v.ConfigData.NetworkOS.Policy.Bandwidth)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.Bandwidth != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Bandwidth=%v", *c.ConfigData.NetworkOS.Policy.Bandwidth)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil && c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		if *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent {
			log.Printf("v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v, c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v", *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent, *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v", *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v", *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.BorderRouter != nil && c.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		if *v.ConfigData.NetworkOS.Policy.BorderRouter != *c.ConfigData.NetworkOS.Policy.BorderRouter {
			log.Printf("v.ConfigData.NetworkOS.Policy.BorderRouter=%v, c.ConfigData.NetworkOS.Policy.BorderRouter=%v", *v.ConfigData.NetworkOS.Policy.BorderRouter, *c.ConfigData.NetworkOS.Policy.BorderRouter)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.BorderRouter != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.BorderRouter=%v", *v.ConfigData.NetworkOS.Policy.BorderRouter)
		} else if c.ConfigData.NetworkOS.Policy.BorderRouter != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.BorderRouter=%v", *c.ConfigData.NetworkOS.Policy.BorderRouter)
		}
	}
	if v.ConfigData.NetworkOS.Policy.BpduFilter != "" {
		if v.ConfigData.NetworkOS.Policy.BpduFilter != c.ConfigData.NetworkOS.Policy.BpduFilter {
			log.Printf("v.ConfigData.NetworkOS.Policy.BpduFilter=%v, c.ConfigData.NetworkOS.Policy.BpduFilter=%v", v.ConfigData.NetworkOS.Policy.BpduFilter, c.ConfigData.NetworkOS.Policy.BpduFilter)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.BpduFilter=%v, c.ConfigData.NetworkOS.Policy.BpduFilter=%v", v.ConfigData.NetworkOS.Policy.BpduFilter, c.ConfigData.NetworkOS.Policy.BpduFilter)
	}
	if v.ConfigData.NetworkOS.Policy.BpduGuard != "" {
		if v.ConfigData.NetworkOS.Policy.BpduGuard != c.ConfigData.NetworkOS.Policy.BpduGuard {
			log.Printf("v.ConfigData.NetworkOS.Policy.BpduGuard=%v, c.ConfigData.NetworkOS.Policy.BpduGuard=%v", v.ConfigData.NetworkOS.Policy.BpduGuard, c.ConfigData.NetworkOS.Policy.BpduGuard)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.BpduGuard=%v, c.ConfigData.NetworkOS.Policy.BpduGuard=%v", v.ConfigData.NetworkOS.Policy.BpduGuard, c.ConfigData.NetworkOS.Policy.BpduGuard)
	}

	if v.ConfigData.NetworkOS.Policy.Cdp != nil && c.ConfigData.NetworkOS.Policy.Cdp != nil {
		if *v.ConfigData.NetworkOS.Policy.Cdp != *c.ConfigData.NetworkOS.Policy.Cdp {
			log.Printf("v.ConfigData.NetworkOS.Policy.Cdp=%v, c.ConfigData.NetworkOS.Policy.Cdp=%v", *v.ConfigData.NetworkOS.Policy.Cdp, *c.ConfigData.NetworkOS.Policy.Cdp)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Cdp != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Cdp=%v", *v.ConfigData.NetworkOS.Policy.Cdp)
		} else if c.ConfigData.NetworkOS.Policy.Cdp != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Cdp=%v", *c.ConfigData.NetworkOS.Policy.Cdp)
		}
	}

	if v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil && c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		if *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer {
			log.Printf("v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer, *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v", *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.DebounceTimer != nil && c.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		if *v.ConfigData.NetworkOS.Policy.DebounceTimer != *c.ConfigData.NetworkOS.Policy.DebounceTimer {
			log.Printf("v.ConfigData.NetworkOS.Policy.DebounceTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceTimer, *c.ConfigData.NetworkOS.Policy.DebounceTimer)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.DebounceTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceTimer)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.DebounceTimer=%v", *c.ConfigData.NetworkOS.Policy.DebounceTimer)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.DuplexMode != "" {
		if v.ConfigData.NetworkOS.Policy.DuplexMode != c.ConfigData.NetworkOS.Policy.DuplexMode {
			log.Printf("v.ConfigData.NetworkOS.Policy.DuplexMode=%v, c.ConfigData.NetworkOS.Policy.DuplexMode=%v", v.ConfigData.NetworkOS.Policy.DuplexMode, c.ConfigData.NetworkOS.Policy.DuplexMode)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.DuplexMode=%v, c.ConfigData.NetworkOS.Policy.DuplexMode=%v", v.ConfigData.NetworkOS.Policy.DuplexMode, c.ConfigData.NetworkOS.Policy.DuplexMode)
	}

	if v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil && c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		if *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl {
			log.Printf("v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v, c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v", *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl, *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v", *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
		} else if c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v", *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
		}
	}
	if v.ConfigData.NetworkOS.Policy.Fec != "" {
		if v.ConfigData.NetworkOS.Policy.Fec != c.ConfigData.NetworkOS.Policy.Fec {
			log.Printf("v.ConfigData.NetworkOS.Policy.Fec=%v, c.ConfigData.NetworkOS.Policy.Fec=%v", v.ConfigData.NetworkOS.Policy.Fec, c.ConfigData.NetworkOS.Policy.Fec)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Fec=%v, c.ConfigData.NetworkOS.Policy.Fec=%v", v.ConfigData.NetworkOS.Policy.Fec, c.ConfigData.NetworkOS.Policy.Fec)
	}
	if v.ConfigData.NetworkOS.Policy.GuardMode != "" {
		if v.ConfigData.NetworkOS.Policy.GuardMode != c.ConfigData.NetworkOS.Policy.GuardMode {
			log.Printf("v.ConfigData.NetworkOS.Policy.GuardMode=%v, c.ConfigData.NetworkOS.Policy.GuardMode=%v", v.ConfigData.NetworkOS.Policy.GuardMode, c.ConfigData.NetworkOS.Policy.GuardMode)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.GuardMode=%v, c.ConfigData.NetworkOS.Policy.GuardMode=%v", v.ConfigData.NetworkOS.Policy.GuardMode, c.ConfigData.NetworkOS.Policy.GuardMode)
	}
	if v.ConfigData.NetworkOS.Policy.IfName != "" {
		if v.ConfigData.NetworkOS.Policy.IfName != c.ConfigData.NetworkOS.Policy.IfName {
			log.Printf("v.ConfigData.NetworkOS.Policy.IfName=%v, c.ConfigData.NetworkOS.Policy.IfName=%v", v.ConfigData.NetworkOS.Policy.IfName, c.ConfigData.NetworkOS.Policy.IfName)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.IfName=%v, c.ConfigData.NetworkOS.Policy.IfName=%v", v.ConfigData.NetworkOS.Policy.IfName, c.ConfigData.NetworkOS.Policy.IfName)
	}

	if v.ConfigData.NetworkOS.Policy.InheritBandwidth != nil && c.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		if *v.ConfigData.NetworkOS.Policy.InheritBandwidth != *c.ConfigData.NetworkOS.Policy.InheritBandwidth {
			log.Printf("v.ConfigData.NetworkOS.Policy.InheritBandwidth=%v, c.ConfigData.NetworkOS.Policy.InheritBandwidth=%v", *v.ConfigData.NetworkOS.Policy.InheritBandwidth, *c.ConfigData.NetworkOS.Policy.InheritBandwidth)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.InheritBandwidth=%v", *v.ConfigData.NetworkOS.Policy.InheritBandwidth)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.InheritBandwidth=%v", *c.ConfigData.NetworkOS.Policy.InheritBandwidth)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Ip != "" {
		if v.ConfigData.NetworkOS.Policy.Ip != c.ConfigData.NetworkOS.Policy.Ip {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ip=%v, c.ConfigData.NetworkOS.Policy.Ip=%v", v.ConfigData.NetworkOS.Policy.Ip, c.ConfigData.NetworkOS.Policy.Ip)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Ip=%v, c.ConfigData.NetworkOS.Policy.Ip=%v", v.ConfigData.NetworkOS.Policy.Ip, c.ConfigData.NetworkOS.Policy.Ip)
	}

	if v.ConfigData.NetworkOS.Policy.IpRedirects != nil && c.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		if *v.ConfigData.NetworkOS.Policy.IpRedirects != *c.ConfigData.NetworkOS.Policy.IpRedirects {
			log.Printf("v.ConfigData.NetworkOS.Policy.IpRedirects=%v, c.ConfigData.NetworkOS.Policy.IpRedirects=%v", *v.ConfigData.NetworkOS.Policy.IpRedirects, *c.ConfigData.NetworkOS.Policy.IpRedirects)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.IpRedirects != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.IpRedirects=%v", *v.ConfigData.NetworkOS.Policy.IpRedirects)
		} else if c.ConfigData.NetworkOS.Policy.IpRedirects != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.IpRedirects=%v", *c.ConfigData.NetworkOS.Policy.IpRedirects)
		}
	}
	if v.ConfigData.NetworkOS.Policy.Ipv6 != "" {
		if v.ConfigData.NetworkOS.Policy.Ipv6 != c.ConfigData.NetworkOS.Policy.Ipv6 {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ipv6=%v, c.ConfigData.NetworkOS.Policy.Ipv6=%v", v.ConfigData.NetworkOS.Policy.Ipv6, c.ConfigData.NetworkOS.Policy.Ipv6)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Ipv6=%v, c.ConfigData.NetworkOS.Policy.Ipv6=%v", v.ConfigData.NetworkOS.Policy.Ipv6, c.ConfigData.NetworkOS.Policy.Ipv6)
	}

	if v.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil && c.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		if *v.ConfigData.NetworkOS.Policy.Ipv6Prefix != *c.ConfigData.NetworkOS.Policy.Ipv6Prefix {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v, c.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v", *v.ConfigData.NetworkOS.Policy.Ipv6Prefix, *c.ConfigData.NetworkOS.Policy.Ipv6Prefix)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v", *v.ConfigData.NetworkOS.Policy.Ipv6Prefix)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v", *c.ConfigData.NetworkOS.Policy.Ipv6Prefix)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.LinkType != "" {
		if v.ConfigData.NetworkOS.Policy.LinkType != c.ConfigData.NetworkOS.Policy.LinkType {
			log.Printf("v.ConfigData.NetworkOS.Policy.LinkType=%v, c.ConfigData.NetworkOS.Policy.LinkType=%v", v.ConfigData.NetworkOS.Policy.LinkType, c.ConfigData.NetworkOS.Policy.LinkType)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.LinkType=%v, c.ConfigData.NetworkOS.Policy.LinkType=%v", v.ConfigData.NetworkOS.Policy.LinkType, c.ConfigData.NetworkOS.Policy.LinkType)
	}
	if v.ConfigData.NetworkOS.Policy.Mask != "" {
		if v.ConfigData.NetworkOS.Policy.Mask != c.ConfigData.NetworkOS.Policy.Mask {
			log.Printf("v.ConfigData.NetworkOS.Policy.Mask=%v, c.ConfigData.NetworkOS.Policy.Mask=%v", v.ConfigData.NetworkOS.Policy.Mask, c.ConfigData.NetworkOS.Policy.Mask)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Mask=%v, c.ConfigData.NetworkOS.Policy.Mask=%v", v.ConfigData.NetworkOS.Policy.Mask, c.ConfigData.NetworkOS.Policy.Mask)
	}

	if v.ConfigData.NetworkOS.Policy.Monitor != nil && c.ConfigData.NetworkOS.Policy.Monitor != nil {
		if *v.ConfigData.NetworkOS.Policy.Monitor != *c.ConfigData.NetworkOS.Policy.Monitor {
			log.Printf("v.ConfigData.NetworkOS.Policy.Monitor=%v, c.ConfigData.NetworkOS.Policy.Monitor=%v", *v.ConfigData.NetworkOS.Policy.Monitor, *c.ConfigData.NetworkOS.Policy.Monitor)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Monitor != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Monitor=%v", *v.ConfigData.NetworkOS.Policy.Monitor)
		} else if c.ConfigData.NetworkOS.Policy.Monitor != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Monitor=%v", *c.ConfigData.NetworkOS.Policy.Monitor)
		}
	}

	if v.ConfigData.NetworkOS.Policy.NativeVlan != nil && c.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.NativeVlan != *c.ConfigData.NetworkOS.Policy.NativeVlan {
			log.Printf("v.ConfigData.NetworkOS.Policy.NativeVlan=%v, c.ConfigData.NetworkOS.Policy.NativeVlan=%v", *v.ConfigData.NetworkOS.Policy.NativeVlan, *c.ConfigData.NetworkOS.Policy.NativeVlan)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.NativeVlan != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.NativeVlan=%v", *v.ConfigData.NetworkOS.Policy.NativeVlan)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.NativeVlan != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.NativeVlan=%v", *c.ConfigData.NetworkOS.Policy.NativeVlan)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.NegotiateAuto != nil && c.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		if *v.ConfigData.NetworkOS.Policy.NegotiateAuto != *c.ConfigData.NetworkOS.Policy.NegotiateAuto {
			log.Printf("v.ConfigData.NetworkOS.Policy.NegotiateAuto=%v, c.ConfigData.NetworkOS.Policy.NegotiateAuto=%v", *v.ConfigData.NetworkOS.Policy.NegotiateAuto, *c.ConfigData.NetworkOS.Policy.NegotiateAuto)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.NegotiateAuto=%v", *v.ConfigData.NetworkOS.Policy.NegotiateAuto)
		} else if c.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.NegotiateAuto=%v", *c.ConfigData.NetworkOS.Policy.NegotiateAuto)
		}
	}

	if v.ConfigData.NetworkOS.Policy.Netflow != nil && c.ConfigData.NetworkOS.Policy.Netflow != nil {
		if *v.ConfigData.NetworkOS.Policy.Netflow != *c.ConfigData.NetworkOS.Policy.Netflow {
			log.Printf("v.ConfigData.NetworkOS.Policy.Netflow=%v, c.ConfigData.NetworkOS.Policy.Netflow=%v", *v.ConfigData.NetworkOS.Policy.Netflow, *c.ConfigData.NetworkOS.Policy.Netflow)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Netflow != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Netflow=%v", *v.ConfigData.NetworkOS.Policy.Netflow)
		} else if c.ConfigData.NetworkOS.Policy.Netflow != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Netflow=%v", *c.ConfigData.NetworkOS.Policy.Netflow)
		}
	}
	if v.ConfigData.NetworkOS.Policy.NetflowMonitor != "" {
		if v.ConfigData.NetworkOS.Policy.NetflowMonitor != c.ConfigData.NetworkOS.Policy.NetflowMonitor {
			log.Printf("v.ConfigData.NetworkOS.Policy.NetflowMonitor=%v, c.ConfigData.NetworkOS.Policy.NetflowMonitor=%v", v.ConfigData.NetworkOS.Policy.NetflowMonitor, c.ConfigData.NetworkOS.Policy.NetflowMonitor)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.NetflowMonitor=%v, c.ConfigData.NetworkOS.Policy.NetflowMonitor=%v", v.ConfigData.NetworkOS.Policy.NetflowMonitor, c.ConfigData.NetworkOS.Policy.NetflowMonitor)
	}
	if v.ConfigData.NetworkOS.Policy.NetflowSampler != "" {
		if v.ConfigData.NetworkOS.Policy.NetflowSampler != c.ConfigData.NetworkOS.Policy.NetflowSampler {
			log.Printf("v.ConfigData.NetworkOS.Policy.NetflowSampler=%v, c.ConfigData.NetworkOS.Policy.NetflowSampler=%v", v.ConfigData.NetworkOS.Policy.NetflowSampler, c.ConfigData.NetworkOS.Policy.NetflowSampler)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.NetflowSampler=%v, c.ConfigData.NetworkOS.Policy.NetflowSampler=%v", v.ConfigData.NetworkOS.Policy.NetflowSampler, c.ConfigData.NetworkOS.Policy.NetflowSampler)
	}

	if v.ConfigData.NetworkOS.Policy.OrphanPort != nil && c.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		if *v.ConfigData.NetworkOS.Policy.OrphanPort != *c.ConfigData.NetworkOS.Policy.OrphanPort {
			log.Printf("v.ConfigData.NetworkOS.Policy.OrphanPort=%v, c.ConfigData.NetworkOS.Policy.OrphanPort=%v", *v.ConfigData.NetworkOS.Policy.OrphanPort, *c.ConfigData.NetworkOS.Policy.OrphanPort)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.OrphanPort != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.OrphanPort=%v", *v.ConfigData.NetworkOS.Policy.OrphanPort)
		} else if c.ConfigData.NetworkOS.Policy.OrphanPort != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.OrphanPort=%v", *c.ConfigData.NetworkOS.Policy.OrphanPort)
		}
	}

	if v.ConfigData.NetworkOS.Policy.PathCost != nil && c.ConfigData.NetworkOS.Policy.PathCost != nil {
		if *v.ConfigData.NetworkOS.Policy.PathCost != *c.ConfigData.NetworkOS.Policy.PathCost {
			log.Printf("v.ConfigData.NetworkOS.Policy.PathCost=%v, c.ConfigData.NetworkOS.Policy.PathCost=%v", *v.ConfigData.NetworkOS.Policy.PathCost, *c.ConfigData.NetworkOS.Policy.PathCost)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PathCost != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PathCost=%v", *v.ConfigData.NetworkOS.Policy.PathCost)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.PathCost != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PathCost=%v", *c.ConfigData.NetworkOS.Policy.PathCost)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Pfc != nil && c.ConfigData.NetworkOS.Policy.Pfc != nil {
		if *v.ConfigData.NetworkOS.Policy.Pfc != *c.ConfigData.NetworkOS.Policy.Pfc {
			log.Printf("v.ConfigData.NetworkOS.Policy.Pfc=%v, c.ConfigData.NetworkOS.Policy.Pfc=%v", *v.ConfigData.NetworkOS.Policy.Pfc, *c.ConfigData.NetworkOS.Policy.Pfc)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Pfc != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Pfc=%v", *v.ConfigData.NetworkOS.Policy.Pfc)
		} else if c.ConfigData.NetworkOS.Policy.Pfc != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Pfc=%v", *c.ConfigData.NetworkOS.Policy.Pfc)
		}
	}

	if v.ConfigData.NetworkOS.Policy.PimDrPriority != nil && c.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		if *v.ConfigData.NetworkOS.Policy.PimDrPriority != *c.ConfigData.NetworkOS.Policy.PimDrPriority {
			log.Printf("v.ConfigData.NetworkOS.Policy.PimDrPriority=%v, c.ConfigData.NetworkOS.Policy.PimDrPriority=%v", *v.ConfigData.NetworkOS.Policy.PimDrPriority, *c.ConfigData.NetworkOS.Policy.PimDrPriority)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PimDrPriority=%v", *v.ConfigData.NetworkOS.Policy.PimDrPriority)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PimDrPriority=%v", *c.ConfigData.NetworkOS.Policy.PimDrPriority)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PimSparse != nil && c.ConfigData.NetworkOS.Policy.PimSparse != nil {
		if *v.ConfigData.NetworkOS.Policy.PimSparse != *c.ConfigData.NetworkOS.Policy.PimSparse {
			log.Printf("v.ConfigData.NetworkOS.Policy.PimSparse=%v, c.ConfigData.NetworkOS.Policy.PimSparse=%v", *v.ConfigData.NetworkOS.Policy.PimSparse, *c.ConfigData.NetworkOS.Policy.PimSparse)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PimSparse != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PimSparse=%v", *v.ConfigData.NetworkOS.Policy.PimSparse)
		} else if c.ConfigData.NetworkOS.Policy.PimSparse != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PimSparse=%v", *c.ConfigData.NetworkOS.Policy.PimSparse)
		}
	}
	if v.ConfigData.NetworkOS.Policy.PortChannelId != "" {
		if v.ConfigData.NetworkOS.Policy.PortChannelId != c.ConfigData.NetworkOS.Policy.PortChannelId {
			log.Printf("v.ConfigData.NetworkOS.Policy.PortChannelId=%v, c.ConfigData.NetworkOS.Policy.PortChannelId=%v", v.ConfigData.NetworkOS.Policy.PortChannelId, c.ConfigData.NetworkOS.Policy.PortChannelId)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.PortChannelId=%v, c.ConfigData.NetworkOS.Policy.PortChannelId=%v", v.ConfigData.NetworkOS.Policy.PortChannelId, c.ConfigData.NetworkOS.Policy.PortChannelId)
	}
	if v.ConfigData.NetworkOS.Policy.PortChannelMode != "" {
		if v.ConfigData.NetworkOS.Policy.PortChannelMode != c.ConfigData.NetworkOS.Policy.PortChannelMode {
			log.Printf("v.ConfigData.NetworkOS.Policy.PortChannelMode=%v, c.ConfigData.NetworkOS.Policy.PortChannelMode=%v", v.ConfigData.NetworkOS.Policy.PortChannelMode, c.ConfigData.NetworkOS.Policy.PortChannelMode)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.PortChannelMode=%v, c.ConfigData.NetworkOS.Policy.PortChannelMode=%v", v.ConfigData.NetworkOS.Policy.PortChannelMode, c.ConfigData.NetworkOS.Policy.PortChannelMode)
	}

	if v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil && c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		if *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk {
			log.Printf("v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v, c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v", *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk, *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v", *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
		} else if c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v", *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
		}
	}
	if v.ConfigData.NetworkOS.Policy.Prefix != "" {
		if v.ConfigData.NetworkOS.Policy.Prefix != c.ConfigData.NetworkOS.Policy.Prefix {
			log.Printf("v.ConfigData.NetworkOS.Policy.Prefix=%v, c.ConfigData.NetworkOS.Policy.Prefix=%v", v.ConfigData.NetworkOS.Policy.Prefix, c.ConfigData.NetworkOS.Policy.Prefix)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Prefix=%v, c.ConfigData.NetworkOS.Policy.Prefix=%v", v.ConfigData.NetworkOS.Policy.Prefix, c.ConfigData.NetworkOS.Policy.Prefix)
	}

	if v.ConfigData.NetworkOS.Policy.Ptp != nil && c.ConfigData.NetworkOS.Policy.Ptp != nil {
		if *v.ConfigData.NetworkOS.Policy.Ptp != *c.ConfigData.NetworkOS.Policy.Ptp {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ptp=%v, c.ConfigData.NetworkOS.Policy.Ptp=%v", *v.ConfigData.NetworkOS.Policy.Ptp, *c.ConfigData.NetworkOS.Policy.Ptp)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Ptp != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Ptp=%v", *v.ConfigData.NetworkOS.Policy.Ptp)
		} else if c.ConfigData.NetworkOS.Policy.Ptp != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Ptp=%v", *c.ConfigData.NetworkOS.Policy.Ptp)
		}
	}
	if v.ConfigData.NetworkOS.Policy.PtpProfile != "" {
		if v.ConfigData.NetworkOS.Policy.PtpProfile != c.ConfigData.NetworkOS.Policy.PtpProfile {
			log.Printf("v.ConfigData.NetworkOS.Policy.PtpProfile=%v, c.ConfigData.NetworkOS.Policy.PtpProfile=%v", v.ConfigData.NetworkOS.Policy.PtpProfile, c.ConfigData.NetworkOS.Policy.PtpProfile)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.PtpProfile=%v, c.ConfigData.NetworkOS.Policy.PtpProfile=%v", v.ConfigData.NetworkOS.Policy.PtpProfile, c.ConfigData.NetworkOS.Policy.PtpProfile)
	}

	if v.ConfigData.NetworkOS.Policy.PtpRole != nil && c.ConfigData.NetworkOS.Policy.PtpRole != nil {
		if *v.ConfigData.NetworkOS.Policy.PtpRole != *c.ConfigData.NetworkOS.Policy.PtpRole {
			log.Printf("v.ConfigData.NetworkOS.Policy.PtpRole=%v, c.ConfigData.NetworkOS.Policy.PtpRole=%v", *v.ConfigData.NetworkOS.Policy.PtpRole, *c.ConfigData.NetworkOS.Policy.PtpRole)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PtpRole != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PtpRole=%v", *v.ConfigData.NetworkOS.Policy.PtpRole)
		} else if c.ConfigData.NetworkOS.Policy.PtpRole != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PtpRole=%v", *c.ConfigData.NetworkOS.Policy.PtpRole)
		}
	}

	if v.ConfigData.NetworkOS.Policy.PtpVlan != nil && c.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.PtpVlan != *c.ConfigData.NetworkOS.Policy.PtpVlan {
			log.Printf("v.ConfigData.NetworkOS.Policy.PtpVlan=%v, c.ConfigData.NetworkOS.Policy.PtpVlan=%v", *v.ConfigData.NetworkOS.Policy.PtpVlan, *c.ConfigData.NetworkOS.Policy.PtpVlan)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PtpVlan != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PtpVlan=%v", *v.ConfigData.NetworkOS.Policy.PtpVlan)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.PtpVlan != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PtpVlan=%v", *c.ConfigData.NetworkOS.Policy.PtpVlan)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans != "" {
		if v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans != c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans {
			log.Printf("v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v", v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v", v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans)
	}
	if v.ConfigData.NetworkOS.Policy.PvlanMode != "" {
		if v.ConfigData.NetworkOS.Policy.PvlanMode != c.ConfigData.NetworkOS.Policy.PvlanMode {
			log.Printf("v.ConfigData.NetworkOS.Policy.PvlanMode=%v, c.ConfigData.NetworkOS.Policy.PvlanMode=%v", v.ConfigData.NetworkOS.Policy.PvlanMode, c.ConfigData.NetworkOS.Policy.PvlanMode)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.PvlanMode=%v, c.ConfigData.NetworkOS.Policy.PvlanMode=%v", v.ConfigData.NetworkOS.Policy.PvlanMode, c.ConfigData.NetworkOS.Policy.PvlanMode)
	}

	if v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil && c.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan {
			log.Printf("v.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v, c.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v", *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan, *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v", *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v", *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Qos != nil && c.ConfigData.NetworkOS.Policy.Qos != nil {
		if *v.ConfigData.NetworkOS.Policy.Qos != *c.ConfigData.NetworkOS.Policy.Qos {
			log.Printf("v.ConfigData.NetworkOS.Policy.Qos=%v, c.ConfigData.NetworkOS.Policy.Qos=%v", *v.ConfigData.NetworkOS.Policy.Qos, *c.ConfigData.NetworkOS.Policy.Qos)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.Qos != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.Qos=%v", *v.ConfigData.NetworkOS.Policy.Qos)
		} else if c.ConfigData.NetworkOS.Policy.Qos != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.Qos=%v", *c.ConfigData.NetworkOS.Policy.Qos)
		}
	}
	if v.ConfigData.NetworkOS.Policy.QosPolicy != "" {
		if v.ConfigData.NetworkOS.Policy.QosPolicy != c.ConfigData.NetworkOS.Policy.QosPolicy {
			log.Printf("v.ConfigData.NetworkOS.Policy.QosPolicy=%v, c.ConfigData.NetworkOS.Policy.QosPolicy=%v", v.ConfigData.NetworkOS.Policy.QosPolicy, c.ConfigData.NetworkOS.Policy.QosPolicy)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.QosPolicy=%v, c.ConfigData.NetworkOS.Policy.QosPolicy=%v", v.ConfigData.NetworkOS.Policy.QosPolicy, c.ConfigData.NetworkOS.Policy.QosPolicy)
	}
	if v.ConfigData.NetworkOS.Policy.QueuingPolicy != "" {
		if v.ConfigData.NetworkOS.Policy.QueuingPolicy != c.ConfigData.NetworkOS.Policy.QueuingPolicy {
			log.Printf("v.ConfigData.NetworkOS.Policy.QueuingPolicy=%v, c.ConfigData.NetworkOS.Policy.QueuingPolicy=%v", v.ConfigData.NetworkOS.Policy.QueuingPolicy, c.ConfigData.NetworkOS.Policy.QueuingPolicy)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.QueuingPolicy=%v, c.ConfigData.NetworkOS.Policy.QueuingPolicy=%v", v.ConfigData.NetworkOS.Policy.QueuingPolicy, c.ConfigData.NetworkOS.Policy.QueuingPolicy)
	}
	if v.ConfigData.NetworkOS.Policy.RoutingTag != "" {
		if v.ConfigData.NetworkOS.Policy.RoutingTag != c.ConfigData.NetworkOS.Policy.RoutingTag {
			log.Printf("v.ConfigData.NetworkOS.Policy.RoutingTag=%v, c.ConfigData.NetworkOS.Policy.RoutingTag=%v", v.ConfigData.NetworkOS.Policy.RoutingTag, c.ConfigData.NetworkOS.Policy.RoutingTag)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.RoutingTag=%v, c.ConfigData.NetworkOS.Policy.RoutingTag=%v", v.ConfigData.NetworkOS.Policy.RoutingTag, c.ConfigData.NetworkOS.Policy.RoutingTag)
	}

	if v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil && c.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		if *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy {
			log.Printf("v.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v, c.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v", *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy, *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v", *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
		} else if c.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v", *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
		}
	}
	if v.ConfigData.NetworkOS.Policy.Speed != "" {
		if v.ConfigData.NetworkOS.Policy.Speed != c.ConfigData.NetworkOS.Policy.Speed {
			log.Printf("v.ConfigData.NetworkOS.Policy.Speed=%v, c.ConfigData.NetworkOS.Policy.Speed=%v", v.ConfigData.NetworkOS.Policy.Speed, c.ConfigData.NetworkOS.Policy.Speed)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.Speed=%v, c.ConfigData.NetworkOS.Policy.Speed=%v", v.ConfigData.NetworkOS.Policy.Speed, c.ConfigData.NetworkOS.Policy.Speed)
	}

	if v.ConfigData.NetworkOS.Policy.StormControl != nil && c.ConfigData.NetworkOS.Policy.StormControl != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControl != *c.ConfigData.NetworkOS.Policy.StormControl {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControl=%v, c.ConfigData.NetworkOS.Policy.StormControl=%v", *v.ConfigData.NetworkOS.Policy.StormControl, *c.ConfigData.NetworkOS.Policy.StormControl)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControl != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControl=%v", *v.ConfigData.NetworkOS.Policy.StormControl)
		} else if c.ConfigData.NetworkOS.Policy.StormControl != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControl=%v", *c.ConfigData.NetworkOS.Policy.StormControl)
		}
	}
	if v.ConfigData.NetworkOS.Policy.StormControlAction != "" {
		if v.ConfigData.NetworkOS.Policy.StormControlAction != c.ConfigData.NetworkOS.Policy.StormControlAction {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlAction=%v, c.ConfigData.NetworkOS.Policy.StormControlAction=%v", v.ConfigData.NetworkOS.Policy.StormControlAction, c.ConfigData.NetworkOS.Policy.StormControlAction)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.StormControlAction=%v, c.ConfigData.NetworkOS.Policy.StormControlAction=%v", v.ConfigData.NetworkOS.Policy.StormControlAction, c.ConfigData.NetworkOS.Policy.StormControlAction)
	}

	if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel, *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
		} else if c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel, *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
		} else if c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel, *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
		} else if c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil && c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		if *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent {
			log.Printf("v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v, c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v", *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent, *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v", *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v", *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.V6prefix != nil && c.ConfigData.NetworkOS.Policy.V6prefix != nil {
		if *v.ConfigData.NetworkOS.Policy.V6prefix != *c.ConfigData.NetworkOS.Policy.V6prefix {
			log.Printf("v.ConfigData.NetworkOS.Policy.V6prefix=%v, c.ConfigData.NetworkOS.Policy.V6prefix=%v", *v.ConfigData.NetworkOS.Policy.V6prefix, *c.ConfigData.NetworkOS.Policy.V6prefix)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.V6prefix != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.V6prefix=%v", *v.ConfigData.NetworkOS.Policy.V6prefix)
			return RequiresUpdate
		} else if c.ConfigData.NetworkOS.Policy.V6prefix != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.V6prefix=%v", *c.ConfigData.NetworkOS.Policy.V6prefix)
			return RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.VlanMapping != nil && c.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		if *v.ConfigData.NetworkOS.Policy.VlanMapping != *c.ConfigData.NetworkOS.Policy.VlanMapping {
			log.Printf("v.ConfigData.NetworkOS.Policy.VlanMapping=%v, c.ConfigData.NetworkOS.Policy.VlanMapping=%v", *v.ConfigData.NetworkOS.Policy.VlanMapping, *c.ConfigData.NetworkOS.Policy.VlanMapping)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.VlanMapping != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.VlanMapping=%v", *v.ConfigData.NetworkOS.Policy.VlanMapping)
		} else if c.ConfigData.NetworkOS.Policy.VlanMapping != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.VlanMapping=%v", *c.ConfigData.NetworkOS.Policy.VlanMapping)
		}
	}
	if v.ConfigData.NetworkOS.Policy.VrfInterface != "" {
		if v.ConfigData.NetworkOS.Policy.VrfInterface != c.ConfigData.NetworkOS.Policy.VrfInterface {
			log.Printf("v.ConfigData.NetworkOS.Policy.VrfInterface=%v, c.ConfigData.NetworkOS.Policy.VrfInterface=%v", v.ConfigData.NetworkOS.Policy.VrfInterface, c.ConfigData.NetworkOS.Policy.VrfInterface)
			return RequiresUpdate
		}
	} else {
		log.Printf("Skipping - v.ConfigData.NetworkOS.Policy.VrfInterface=%v, c.ConfigData.NetworkOS.Policy.VrfInterface=%v", v.ConfigData.NetworkOS.Policy.VrfInterface, c.ConfigData.NetworkOS.Policy.VrfInterface)
	}

	if v.ConfigData.NetworkOS.Policy.WanLink != nil && c.ConfigData.NetworkOS.Policy.WanLink != nil {
		if *v.ConfigData.NetworkOS.Policy.WanLink != *c.ConfigData.NetworkOS.Policy.WanLink {
			log.Printf("v.ConfigData.NetworkOS.Policy.WanLink=%v, c.ConfigData.NetworkOS.Policy.WanLink=%v", *v.ConfigData.NetworkOS.Policy.WanLink, *c.ConfigData.NetworkOS.Policy.WanLink)
			return RequiresUpdate
		}
	} else {
		if v.ConfigData.NetworkOS.Policy.WanLink != nil {
			log.Printf("v.ConfigData.NetworkOS.Policy.WanLink=%v", *v.ConfigData.NetworkOS.Policy.WanLink)
		} else if c.ConfigData.NetworkOS.Policy.WanLink != nil {
			log.Printf("c.ConfigData.NetworkOS.Policy.WanLink=%v", *c.ConfigData.NetworkOS.Policy.WanLink)
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.AssociationList) != len(c.ConfigData.NetworkOS.Policy.AssociationList) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.AssociationList)=%d, len(c.ConfigData.NetworkOS.Policy.AssociationList)=%d", len(v.ConfigData.NetworkOS.Policy.AssociationList), len(c.ConfigData.NetworkOS.Policy.AssociationList))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.AssociationList {
		retVal := v.ConfigData.NetworkOS.Policy.AssociationList[i].DeepEqual(c.ConfigData.NetworkOS.Policy.AssociationList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MappingList) != len(c.ConfigData.NetworkOS.Policy.MappingList) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.MappingList)=%d, len(c.ConfigData.NetworkOS.Policy.MappingList)=%d", len(v.ConfigData.NetworkOS.Policy.MappingList), len(c.ConfigData.NetworkOS.Policy.MappingList))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.MappingList {
		retVal := v.ConfigData.NetworkOS.Policy.MappingList[i].DeepEqual(c.ConfigData.NetworkOS.Policy.MappingList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges) != len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges)=%d, len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges)=%d", len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges), len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.IgmpGroupRanges {
		retVal := v.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i].DeepEqual(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MstPathCostList) != len(c.ConfigData.NetworkOS.Policy.MstPathCostList) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.MstPathCostList)=%d, len(c.ConfigData.NetworkOS.Policy.MstPathCostList)=%d", len(v.ConfigData.NetworkOS.Policy.MstPathCostList), len(c.ConfigData.NetworkOS.Policy.MstPathCostList))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.MstPathCostList {
		retVal := v.ConfigData.NetworkOS.Policy.MstPathCostList[i].DeepEqual(c.ConfigData.NetworkOS.Policy.MstPathCostList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MstPortPrioList) != len(c.ConfigData.NetworkOS.Policy.MstPortPrioList) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.MstPortPrioList)=%d, len(c.ConfigData.NetworkOS.Policy.MstPortPrioList)=%d", len(v.ConfigData.NetworkOS.Policy.MstPortPrioList), len(c.ConfigData.NetworkOS.Policy.MstPortPrioList))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.MstPortPrioList {
		retVal := v.ConfigData.NetworkOS.Policy.MstPortPrioList[i].DeepEqual(c.ConfigData.NetworkOS.Policy.MstPortPrioList[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries) != len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries) {
		log.Printf("len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries)=%d, len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries)=%d", len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries), len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries))
		return RequiresUpdate
	}
	for i := range v.ConfigData.NetworkOS.Policy.VlanMappingEntries {
		retVal := v.ConfigData.NetworkOS.Policy.VlanMappingEntries[i].DeepEqual(c.ConfigData.NetworkOS.Policy.VlanMappingEntries[i])
		if retVal != ValuesDeeplyEqual {
			return retVal
		}
	}
	if cf {
		return ControlFlagUpdate
	}
	return ValuesDeeplyEqual
}

func (v *NDFCAssociationListValue) CreatePlan(c NDFCAssociationListValue, cf *bool) int {
	action := ActionNone

	if v.PVlan != nil && c.PVlan != nil {
		if *v.PVlan != *c.PVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.PVlan=%v, c.PVlan=%v", *v.PVlan, *c.PVlan)
		}
	} else if v.PVlan != nil {
		log.Printf("Update: v.PVlan=%v, c.PVlan=nil", *v.PVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.PVlan != nil {
		v.PVlan = new(int64)
		log.Printf("Copy from state: v.PVlan=nil, c.PVlan=%v", *c.PVlan)
		*v.PVlan = *c.PVlan
	}
	if v.SVlan != nil && c.SVlan != nil {
		if *v.SVlan != *c.SVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.SVlan=%v, c.SVlan=%v", *v.SVlan, *c.SVlan)
		}
	} else if v.SVlan != nil {
		log.Printf("Update: v.SVlan=%v, c.SVlan=nil", *v.SVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.SVlan != nil {
		v.SVlan = new(int64)
		log.Printf("Copy from state: v.SVlan=nil, c.SVlan=%v", *c.SVlan)
		*v.SVlan = *c.SVlan
	}
	return action
}

func (v *NDFCMappingListValue) CreatePlan(c NDFCMappingListValue, cf *bool) int {
	action := ActionNone

	if v.PVlan != nil && c.PVlan != nil {
		if *v.PVlan != *c.PVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.PVlan=%v, c.PVlan=%v", *v.PVlan, *c.PVlan)
		}
	} else if v.PVlan != nil {
		log.Printf("Update: v.PVlan=%v, c.PVlan=nil", *v.PVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.PVlan != nil {
		v.PVlan = new(int64)
		log.Printf("Copy from state: v.PVlan=nil, c.PVlan=%v", *c.PVlan)
		*v.PVlan = *c.PVlan
	}
	if v.SVlan != nil && c.SVlan != nil {
		if *v.SVlan != *c.SVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.SVlan=%v, c.SVlan=%v", *v.SVlan, *c.SVlan)
		}
	} else if v.SVlan != nil {
		log.Printf("Update: v.SVlan=%v, c.SVlan=nil", *v.SVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.SVlan != nil {
		v.SVlan = new(int64)
		log.Printf("Copy from state: v.SVlan=nil, c.SVlan=%v", *c.SVlan)
		*v.SVlan = *c.SVlan
	}
	return action
}

func (v *NDFCIgmpGroupRangesValue) CreatePlan(c NDFCIgmpGroupRangesValue, cf *bool) int {
	action := ActionNone

	if v.GroupAddress != c.GroupAddress {
		log.Printf("GroupAddress-Update: v.GroupAddress=%v, c.GroupAddress=%v", v.GroupAddress, c.GroupAddress)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.Prefix != nil && c.Prefix != nil {
		if *v.Prefix != *c.Prefix {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.Prefix=%v, c.Prefix=%v", *v.Prefix, *c.Prefix)
		}
	} else if v.Prefix != nil {
		log.Printf("Update: v.Prefix=%v, c.Prefix=nil", *v.Prefix)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.Prefix != nil {
		v.Prefix = new(int64)
		log.Printf("Copy from state: v.Prefix=nil, c.Prefix=%v", *c.Prefix)
		*v.Prefix = *c.Prefix
	}
	return action
}

func (v *NDFCMstPathCostListValue) CreatePlan(c NDFCMstPathCostListValue, cf *bool) int {
	action := ActionNone

	if len(v.MstInstanceList) != len(c.MstInstanceList) {
		log.Printf("Update: len(v.MstInstanceList)=%d, len(c.MstInstanceList)=%d", len(v.MstInstanceList), len(c.MstInstanceList))
		return RequiresUpdate
	}
	for i := range v.MstInstanceList {
		if v.MstInstanceList[i] != c.MstInstanceList[i] {
			log.Printf("Update: v.MstInstanceList[%d]=%s, c.MstInstanceList[%d]=%s", i, v.MstInstanceList[i], i, c.MstInstanceList[i])
			return RequiresUpdate
		}
	}

	if v.PathCost != nil && c.PathCost != nil {
		if *v.PathCost != *c.PathCost {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.PathCost=%v, c.PathCost=%v", *v.PathCost, *c.PathCost)
		}
	} else if v.PathCost != nil {
		log.Printf("Update: v.PathCost=%v, c.PathCost=nil", *v.PathCost)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.PathCost != nil {
		v.PathCost = new(int64)
		log.Printf("Copy from state: v.PathCost=nil, c.PathCost=%v", *c.PathCost)
		*v.PathCost = *c.PathCost
	}
	return action
}

func (v *NDFCMstPortPrioListValue) CreatePlan(c NDFCMstPortPrioListValue, cf *bool) int {
	action := ActionNone

	if len(v.MstInstanceList) != len(c.MstInstanceList) {
		log.Printf("Update: len(v.MstInstanceList)=%d, len(c.MstInstanceList)=%d", len(v.MstInstanceList), len(c.MstInstanceList))
		return RequiresUpdate
	}
	for i := range v.MstInstanceList {
		if v.MstInstanceList[i] != c.MstInstanceList[i] {
			log.Printf("Update: v.MstInstanceList[%d]=%s, c.MstInstanceList[%d]=%s", i, v.MstInstanceList[i], i, c.MstInstanceList[i])
			return RequiresUpdate
		}
	}

	if v.PortPriority != nil && c.PortPriority != nil {
		if *v.PortPriority != *c.PortPriority {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.PortPriority=%v, c.PortPriority=%v", *v.PortPriority, *c.PortPriority)
		}
	} else if v.PortPriority != nil {
		log.Printf("Update: v.PortPriority=%v, c.PortPriority=nil", *v.PortPriority)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.PortPriority != nil {
		v.PortPriority = new(int64)
		log.Printf("Copy from state: v.PortPriority=nil, c.PortPriority=%v", *c.PortPriority)
		*v.PortPriority = *c.PortPriority
	}
	return action
}

func (v *NDFCVlanMappingEntriesValue) CreatePlan(c NDFCVlanMappingEntriesValue, cf *bool) int {
	action := ActionNone

	if len(v.CustomerVlanId) != len(c.CustomerVlanId) {
		log.Printf("Update: len(v.CustomerVlanId)=%d, len(c.CustomerVlanId)=%d", len(v.CustomerVlanId), len(c.CustomerVlanId))
		return RequiresUpdate
	}
	for i := range v.CustomerVlanId {
		if v.CustomerVlanId[i] != c.CustomerVlanId[i] {
			log.Printf("Update: v.CustomerVlanId[%d]=%s, c.CustomerVlanId[%d]=%s", i, v.CustomerVlanId[i], i, c.CustomerVlanId[i])
			return RequiresUpdate
		}
	}

	if v.Dot1qTunnel != c.Dot1qTunnel {
		log.Printf("Dot1qTunnel-Update: v.Dot1qTunnel=%v, c.Dot1qTunnel=%v", v.Dot1qTunnel, c.Dot1qTunnel)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.CustomerInnerVlanId != nil && c.CustomerInnerVlanId != nil {
		if *v.CustomerInnerVlanId != *c.CustomerInnerVlanId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.CustomerInnerVlanId=%v, c.CustomerInnerVlanId=%v", *v.CustomerInnerVlanId, *c.CustomerInnerVlanId)
		}
	} else if v.CustomerInnerVlanId != nil {
		log.Printf("Update: v.CustomerInnerVlanId=%v, c.CustomerInnerVlanId=nil", *v.CustomerInnerVlanId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.CustomerInnerVlanId != nil {
		v.CustomerInnerVlanId = new(int64)
		log.Printf("Copy from state: v.CustomerInnerVlanId=nil, c.CustomerInnerVlanId=%v", *c.CustomerInnerVlanId)
		*v.CustomerInnerVlanId = *c.CustomerInnerVlanId
	}
	if v.ProviderVlanId != nil && c.ProviderVlanId != nil {
		if *v.ProviderVlanId != *c.ProviderVlanId {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update:: v.ProviderVlanId=%v, c.ProviderVlanId=%v", *v.ProviderVlanId, *c.ProviderVlanId)
		}
	} else if v.ProviderVlanId != nil {
		log.Printf("Update: v.ProviderVlanId=%v, c.ProviderVlanId=nil", *v.ProviderVlanId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ProviderVlanId != nil {
		v.ProviderVlanId = new(int64)
		log.Printf("Copy from state: v.ProviderVlanId=nil, c.ProviderVlanId=%v", *c.ProviderVlanId)
		*v.ProviderVlanId = *c.ProviderVlanId
	}
	return action
}

func (v *NDFCInterfacesValue) CreatePlan(c NDFCInterfacesValue, cf *bool) int {
	action := ActionNone

	if v.SwitchId != c.SwitchId {
		log.Printf("SwitchId-Update: v.SwitchId=%v, c.SwitchId=%v", v.SwitchId, c.SwitchId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.ConfigData.Mode != c.ConfigData.Mode {
		log.Printf("Update: v.ConfigData.Mode=%v, c.ConfigData.Mode=%v", v.ConfigData.Mode, c.ConfigData.Mode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.ConfigData.NetworkOS.NetworkOsType != c.ConfigData.NetworkOS.NetworkOsType {
		log.Printf("Update: v.ConfigData.NetworkOS.NetworkOsType=%v, c.ConfigData.NetworkOS.NetworkOsType=%v", v.ConfigData.NetworkOS.NetworkOsType, c.ConfigData.NetworkOS.NetworkOsType)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresReplace
		}
	}

	if v.ConfigData.NetworkOS.Policy.PolicyType != c.ConfigData.NetworkOS.Policy.PolicyType {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PolicyType=%v, c.ConfigData.NetworkOS.Policy.PolicyType=%v", v.ConfigData.NetworkOS.Policy.PolicyType, c.ConfigData.NetworkOS.Policy.PolicyType)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.TemplateName != c.ConfigData.NetworkOS.Policy.TemplateName {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.TemplateName=%v, c.ConfigData.NetworkOS.Policy.TemplateName=%v", v.ConfigData.NetworkOS.Policy.TemplateName, c.ConfigData.NetworkOS.Policy.TemplateName)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.TemplateConfig) != len(c.ConfigData.NetworkOS.Policy.TemplateConfig) {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.TemplateConfig len=%d, c.ConfigData.NetworkOS.Policy.TemplateConfig len=%d", len(v.ConfigData.NetworkOS.Policy.TemplateConfig), len(c.ConfigData.NetworkOS.Policy.TemplateConfig))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else {
		for kk, vv := range v.ConfigData.NetworkOS.Policy.TemplateConfig {
			cc, ok := c.ConfigData.NetworkOS.Policy.TemplateConfig[kk]
			if !ok || vv != cc {
				log.Printf("Update: v.ConfigData.NetworkOS.Policy.TemplateConfig[%s]=%s, c=%s (ok=%v)", kk, vv, cc, ok)
				if action == ActionNone || action == RequiresUpdate {
					action = RequiresUpdate
				}
				break
			}
		}
	}
	if len(v.ConfigData.NetworkOS.Policy.TemplateConfig) == 0 && len(c.ConfigData.NetworkOS.Policy.TemplateConfig) > 0 {
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.TemplateConfig empty, c has %d entries", len(c.ConfigData.NetworkOS.Policy.TemplateConfig))
		v.ConfigData.NetworkOS.Policy.TemplateConfig = c.ConfigData.NetworkOS.Policy.TemplateConfig
	}

	if v.ConfigData.NetworkOS.Policy.AdminState != nil && c.ConfigData.NetworkOS.Policy.AdminState != nil {
		if *v.ConfigData.NetworkOS.Policy.AdminState != *c.ConfigData.NetworkOS.Policy.AdminState {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.AdminState=%v, c.ConfigData.NetworkOS.Policy.AdminState=%v", *v.ConfigData.NetworkOS.Policy.AdminState, *c.ConfigData.NetworkOS.Policy.AdminState)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.AdminState != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.AdminState=%v, c.nil", *v.ConfigData.NetworkOS.Policy.AdminState)
	} else if c.ConfigData.NetworkOS.Policy.AdminState != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.AdminState=%v", *c.ConfigData.NetworkOS.Policy.AdminState)
		v.ConfigData.NetworkOS.Policy.AdminState = new(bool)
		*v.ConfigData.NetworkOS.Policy.AdminState = *c.ConfigData.NetworkOS.Policy.AdminState
	}

	if v.ConfigData.NetworkOS.Policy.Description != c.ConfigData.NetworkOS.Policy.Description {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Description=%v, c.ConfigData.NetworkOS.Policy.Description=%v", v.ConfigData.NetworkOS.Policy.Description, c.ConfigData.NetworkOS.Policy.Description)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Mtu != c.ConfigData.NetworkOS.Policy.Mtu {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Mtu=%v, c.ConfigData.NetworkOS.Policy.Mtu=%v", v.ConfigData.NetworkOS.Policy.Mtu, c.ConfigData.NetworkOS.Policy.Mtu)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.ExtraConfig != c.ConfigData.NetworkOS.Policy.ExtraConfig {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.ExtraConfig=%v, c.ConfigData.NetworkOS.Policy.ExtraConfig=%v", v.ConfigData.NetworkOS.Policy.ExtraConfig, c.ConfigData.NetworkOS.Policy.ExtraConfig)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.AccessVlan != nil && c.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.AccessVlan != *c.ConfigData.NetworkOS.Policy.AccessVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.AccessVlan=%v, c.ConfigData.NetworkOS.Policy.AccessVlan=%v", *v.ConfigData.NetworkOS.Policy.AccessVlan, *c.ConfigData.NetworkOS.Policy.AccessVlan)
		}
	} else if v.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.AccessVlan=%v, c.ConfigData.NetworkOS.Policy.AccessVlan=nil", *v.ConfigData.NetworkOS.Policy.AccessVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.AccessVlan != nil {
		v.ConfigData.NetworkOS.Policy.AccessVlan = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.AccessVlan=nil, c.ConfigData.NetworkOS.Policy.AccessVlan=%v", *c.ConfigData.NetworkOS.Policy.AccessVlan)
		*v.ConfigData.NetworkOS.Policy.AccessVlan = *c.ConfigData.NetworkOS.Policy.AccessVlan
	}

	if v.ConfigData.NetworkOS.Policy.AllowedVlans != c.ConfigData.NetworkOS.Policy.AllowedVlans {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.AllowedVlans=%v, c.ConfigData.NetworkOS.Policy.AllowedVlans=%v", v.ConfigData.NetworkOS.Policy.AllowedVlans, c.ConfigData.NetworkOS.Policy.AllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Bandwidth != nil && c.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		if *v.ConfigData.NetworkOS.Policy.Bandwidth != *c.ConfigData.NetworkOS.Policy.Bandwidth {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Bandwidth=%v, c.ConfigData.NetworkOS.Policy.Bandwidth=%v", *v.ConfigData.NetworkOS.Policy.Bandwidth, *c.ConfigData.NetworkOS.Policy.Bandwidth)
		}
	} else if v.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Bandwidth=%v, c.ConfigData.NetworkOS.Policy.Bandwidth=nil", *v.ConfigData.NetworkOS.Policy.Bandwidth)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.Bandwidth != nil {
		v.ConfigData.NetworkOS.Policy.Bandwidth = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.Bandwidth=nil, c.ConfigData.NetworkOS.Policy.Bandwidth=%v", *c.ConfigData.NetworkOS.Policy.Bandwidth)
		*v.ConfigData.NetworkOS.Policy.Bandwidth = *c.ConfigData.NetworkOS.Policy.Bandwidth
	}

	if v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil && c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		if *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v, c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v", *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent, *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
		}
	} else if v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v, c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=nil", *v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent != nil {
		v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=nil, c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent=%v", *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent)
		*v.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent = *c.ConfigData.NetworkOS.Policy.BandwidthCapacityPercent
	}

	if v.ConfigData.NetworkOS.Policy.BorderRouter != nil && c.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		if *v.ConfigData.NetworkOS.Policy.BorderRouter != *c.ConfigData.NetworkOS.Policy.BorderRouter {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.BorderRouter=%v, c.ConfigData.NetworkOS.Policy.BorderRouter=%v", *v.ConfigData.NetworkOS.Policy.BorderRouter, *c.ConfigData.NetworkOS.Policy.BorderRouter)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.BorderRouter=%v, c.nil", *v.ConfigData.NetworkOS.Policy.BorderRouter)
	} else if c.ConfigData.NetworkOS.Policy.BorderRouter != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.BorderRouter=%v", *c.ConfigData.NetworkOS.Policy.BorderRouter)
		v.ConfigData.NetworkOS.Policy.BorderRouter = new(bool)
		*v.ConfigData.NetworkOS.Policy.BorderRouter = *c.ConfigData.NetworkOS.Policy.BorderRouter
	}

	if v.ConfigData.NetworkOS.Policy.BpduFilter != c.ConfigData.NetworkOS.Policy.BpduFilter {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.BpduFilter=%v, c.ConfigData.NetworkOS.Policy.BpduFilter=%v", v.ConfigData.NetworkOS.Policy.BpduFilter, c.ConfigData.NetworkOS.Policy.BpduFilter)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.BpduGuard != c.ConfigData.NetworkOS.Policy.BpduGuard {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.BpduGuard=%v, c.ConfigData.NetworkOS.Policy.BpduGuard=%v", v.ConfigData.NetworkOS.Policy.BpduGuard, c.ConfigData.NetworkOS.Policy.BpduGuard)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Cdp != nil && c.ConfigData.NetworkOS.Policy.Cdp != nil {
		if *v.ConfigData.NetworkOS.Policy.Cdp != *c.ConfigData.NetworkOS.Policy.Cdp {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Cdp=%v, c.ConfigData.NetworkOS.Policy.Cdp=%v", *v.ConfigData.NetworkOS.Policy.Cdp, *c.ConfigData.NetworkOS.Policy.Cdp)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Cdp != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Cdp=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Cdp)
	} else if c.ConfigData.NetworkOS.Policy.Cdp != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Cdp=%v", *c.ConfigData.NetworkOS.Policy.Cdp)
		v.ConfigData.NetworkOS.Policy.Cdp = new(bool)
		*v.ConfigData.NetworkOS.Policy.Cdp = *c.ConfigData.NetworkOS.Policy.Cdp
	}

	if v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil && c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		if *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer, *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
		}
	} else if v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=nil", *v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer != nil {
		v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=nil, c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer=%v", *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer)
		*v.ConfigData.NetworkOS.Policy.DebounceLinkupTimer = *c.ConfigData.NetworkOS.Policy.DebounceLinkupTimer
	}

	if v.ConfigData.NetworkOS.Policy.DebounceTimer != nil && c.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		if *v.ConfigData.NetworkOS.Policy.DebounceTimer != *c.ConfigData.NetworkOS.Policy.DebounceTimer {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.DebounceTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceTimer=%v", *v.ConfigData.NetworkOS.Policy.DebounceTimer, *c.ConfigData.NetworkOS.Policy.DebounceTimer)
		}
	} else if v.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.DebounceTimer=%v, c.ConfigData.NetworkOS.Policy.DebounceTimer=nil", *v.ConfigData.NetworkOS.Policy.DebounceTimer)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.DebounceTimer != nil {
		v.ConfigData.NetworkOS.Policy.DebounceTimer = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.DebounceTimer=nil, c.ConfigData.NetworkOS.Policy.DebounceTimer=%v", *c.ConfigData.NetworkOS.Policy.DebounceTimer)
		*v.ConfigData.NetworkOS.Policy.DebounceTimer = *c.ConfigData.NetworkOS.Policy.DebounceTimer
	}

	if v.ConfigData.NetworkOS.Policy.DuplexMode != c.ConfigData.NetworkOS.Policy.DuplexMode {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.DuplexMode=%v, c.ConfigData.NetworkOS.Policy.DuplexMode=%v", v.ConfigData.NetworkOS.Policy.DuplexMode, c.ConfigData.NetworkOS.Policy.DuplexMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil && c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		if *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v, c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v", *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl, *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v, c.nil", *v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
	} else if c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl=%v", *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl)
		v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl = new(bool)
		*v.ConfigData.NetworkOS.Policy.ErrorDetectionAcl = *c.ConfigData.NetworkOS.Policy.ErrorDetectionAcl
	}

	if v.ConfigData.NetworkOS.Policy.Fec != c.ConfigData.NetworkOS.Policy.Fec {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Fec=%v, c.ConfigData.NetworkOS.Policy.Fec=%v", v.ConfigData.NetworkOS.Policy.Fec, c.ConfigData.NetworkOS.Policy.Fec)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.GuardMode != c.ConfigData.NetworkOS.Policy.GuardMode {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.GuardMode=%v, c.ConfigData.NetworkOS.Policy.GuardMode=%v", v.ConfigData.NetworkOS.Policy.GuardMode, c.ConfigData.NetworkOS.Policy.GuardMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.IfName != c.ConfigData.NetworkOS.Policy.IfName {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.IfName=%v, c.ConfigData.NetworkOS.Policy.IfName=%v", v.ConfigData.NetworkOS.Policy.IfName, c.ConfigData.NetworkOS.Policy.IfName)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.InheritBandwidth != nil && c.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		if *v.ConfigData.NetworkOS.Policy.InheritBandwidth != *c.ConfigData.NetworkOS.Policy.InheritBandwidth {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.InheritBandwidth=%v, c.ConfigData.NetworkOS.Policy.InheritBandwidth=%v", *v.ConfigData.NetworkOS.Policy.InheritBandwidth, *c.ConfigData.NetworkOS.Policy.InheritBandwidth)
		}
	} else if v.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.InheritBandwidth=%v, c.ConfigData.NetworkOS.Policy.InheritBandwidth=nil", *v.ConfigData.NetworkOS.Policy.InheritBandwidth)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.InheritBandwidth != nil {
		v.ConfigData.NetworkOS.Policy.InheritBandwidth = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.InheritBandwidth=nil, c.ConfigData.NetworkOS.Policy.InheritBandwidth=%v", *c.ConfigData.NetworkOS.Policy.InheritBandwidth)
		*v.ConfigData.NetworkOS.Policy.InheritBandwidth = *c.ConfigData.NetworkOS.Policy.InheritBandwidth
	}

	if v.ConfigData.NetworkOS.Policy.Ip != c.ConfigData.NetworkOS.Policy.Ip {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ip=%v, c.ConfigData.NetworkOS.Policy.Ip=%v", v.ConfigData.NetworkOS.Policy.Ip, c.ConfigData.NetworkOS.Policy.Ip)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.IpRedirects != nil && c.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		if *v.ConfigData.NetworkOS.Policy.IpRedirects != *c.ConfigData.NetworkOS.Policy.IpRedirects {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.IpRedirects=%v, c.ConfigData.NetworkOS.Policy.IpRedirects=%v", *v.ConfigData.NetworkOS.Policy.IpRedirects, *c.ConfigData.NetworkOS.Policy.IpRedirects)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.IpRedirects=%v, c.nil", *v.ConfigData.NetworkOS.Policy.IpRedirects)
	} else if c.ConfigData.NetworkOS.Policy.IpRedirects != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.IpRedirects=%v", *c.ConfigData.NetworkOS.Policy.IpRedirects)
		v.ConfigData.NetworkOS.Policy.IpRedirects = new(bool)
		*v.ConfigData.NetworkOS.Policy.IpRedirects = *c.ConfigData.NetworkOS.Policy.IpRedirects
	}

	if v.ConfigData.NetworkOS.Policy.Ipv6 != c.ConfigData.NetworkOS.Policy.Ipv6 {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ipv6=%v, c.ConfigData.NetworkOS.Policy.Ipv6=%v", v.ConfigData.NetworkOS.Policy.Ipv6, c.ConfigData.NetworkOS.Policy.Ipv6)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil && c.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		if *v.ConfigData.NetworkOS.Policy.Ipv6Prefix != *c.ConfigData.NetworkOS.Policy.Ipv6Prefix {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v, c.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v", *v.ConfigData.NetworkOS.Policy.Ipv6Prefix, *c.ConfigData.NetworkOS.Policy.Ipv6Prefix)
		}
	} else if v.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v, c.ConfigData.NetworkOS.Policy.Ipv6Prefix=nil", *v.ConfigData.NetworkOS.Policy.Ipv6Prefix)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.Ipv6Prefix != nil {
		v.ConfigData.NetworkOS.Policy.Ipv6Prefix = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.Ipv6Prefix=nil, c.ConfigData.NetworkOS.Policy.Ipv6Prefix=%v", *c.ConfigData.NetworkOS.Policy.Ipv6Prefix)
		*v.ConfigData.NetworkOS.Policy.Ipv6Prefix = *c.ConfigData.NetworkOS.Policy.Ipv6Prefix
	}

	if v.ConfigData.NetworkOS.Policy.LinkType != c.ConfigData.NetworkOS.Policy.LinkType {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.LinkType=%v, c.ConfigData.NetworkOS.Policy.LinkType=%v", v.ConfigData.NetworkOS.Policy.LinkType, c.ConfigData.NetworkOS.Policy.LinkType)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Mask != c.ConfigData.NetworkOS.Policy.Mask {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Mask=%v, c.ConfigData.NetworkOS.Policy.Mask=%v", v.ConfigData.NetworkOS.Policy.Mask, c.ConfigData.NetworkOS.Policy.Mask)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Monitor != nil && c.ConfigData.NetworkOS.Policy.Monitor != nil {
		if *v.ConfigData.NetworkOS.Policy.Monitor != *c.ConfigData.NetworkOS.Policy.Monitor {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Monitor=%v, c.ConfigData.NetworkOS.Policy.Monitor=%v", *v.ConfigData.NetworkOS.Policy.Monitor, *c.ConfigData.NetworkOS.Policy.Monitor)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Monitor != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Monitor=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Monitor)
	} else if c.ConfigData.NetworkOS.Policy.Monitor != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Monitor=%v", *c.ConfigData.NetworkOS.Policy.Monitor)
		v.ConfigData.NetworkOS.Policy.Monitor = new(bool)
		*v.ConfigData.NetworkOS.Policy.Monitor = *c.ConfigData.NetworkOS.Policy.Monitor
	}

	if v.ConfigData.NetworkOS.Policy.NativeVlan != nil && c.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.NativeVlan != *c.ConfigData.NetworkOS.Policy.NativeVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.NativeVlan=%v, c.ConfigData.NetworkOS.Policy.NativeVlan=%v", *v.ConfigData.NetworkOS.Policy.NativeVlan, *c.ConfigData.NetworkOS.Policy.NativeVlan)
		}
	} else if v.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.NativeVlan=%v, c.ConfigData.NetworkOS.Policy.NativeVlan=nil", *v.ConfigData.NetworkOS.Policy.NativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.NativeVlan != nil {
		v.ConfigData.NetworkOS.Policy.NativeVlan = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.NativeVlan=nil, c.ConfigData.NetworkOS.Policy.NativeVlan=%v", *c.ConfigData.NetworkOS.Policy.NativeVlan)
		*v.ConfigData.NetworkOS.Policy.NativeVlan = *c.ConfigData.NetworkOS.Policy.NativeVlan
	}

	if v.ConfigData.NetworkOS.Policy.NegotiateAuto != nil && c.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		if *v.ConfigData.NetworkOS.Policy.NegotiateAuto != *c.ConfigData.NetworkOS.Policy.NegotiateAuto {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.NegotiateAuto=%v, c.ConfigData.NetworkOS.Policy.NegotiateAuto=%v", *v.ConfigData.NetworkOS.Policy.NegotiateAuto, *c.ConfigData.NetworkOS.Policy.NegotiateAuto)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.NegotiateAuto=%v, c.nil", *v.ConfigData.NetworkOS.Policy.NegotiateAuto)
	} else if c.ConfigData.NetworkOS.Policy.NegotiateAuto != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.NegotiateAuto=%v", *c.ConfigData.NetworkOS.Policy.NegotiateAuto)
		v.ConfigData.NetworkOS.Policy.NegotiateAuto = new(bool)
		*v.ConfigData.NetworkOS.Policy.NegotiateAuto = *c.ConfigData.NetworkOS.Policy.NegotiateAuto
	}

	if v.ConfigData.NetworkOS.Policy.Netflow != nil && c.ConfigData.NetworkOS.Policy.Netflow != nil {
		if *v.ConfigData.NetworkOS.Policy.Netflow != *c.ConfigData.NetworkOS.Policy.Netflow {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Netflow=%v, c.ConfigData.NetworkOS.Policy.Netflow=%v", *v.ConfigData.NetworkOS.Policy.Netflow, *c.ConfigData.NetworkOS.Policy.Netflow)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Netflow != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Netflow=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Netflow)
	} else if c.ConfigData.NetworkOS.Policy.Netflow != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Netflow=%v", *c.ConfigData.NetworkOS.Policy.Netflow)
		v.ConfigData.NetworkOS.Policy.Netflow = new(bool)
		*v.ConfigData.NetworkOS.Policy.Netflow = *c.ConfigData.NetworkOS.Policy.Netflow
	}

	if v.ConfigData.NetworkOS.Policy.NetflowMonitor != c.ConfigData.NetworkOS.Policy.NetflowMonitor {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.NetflowMonitor=%v, c.ConfigData.NetworkOS.Policy.NetflowMonitor=%v", v.ConfigData.NetworkOS.Policy.NetflowMonitor, c.ConfigData.NetworkOS.Policy.NetflowMonitor)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.NetflowSampler != c.ConfigData.NetworkOS.Policy.NetflowSampler {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.NetflowSampler=%v, c.ConfigData.NetworkOS.Policy.NetflowSampler=%v", v.ConfigData.NetworkOS.Policy.NetflowSampler, c.ConfigData.NetworkOS.Policy.NetflowSampler)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.OrphanPort != nil && c.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		if *v.ConfigData.NetworkOS.Policy.OrphanPort != *c.ConfigData.NetworkOS.Policy.OrphanPort {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.OrphanPort=%v, c.ConfigData.NetworkOS.Policy.OrphanPort=%v", *v.ConfigData.NetworkOS.Policy.OrphanPort, *c.ConfigData.NetworkOS.Policy.OrphanPort)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.OrphanPort=%v, c.nil", *v.ConfigData.NetworkOS.Policy.OrphanPort)
	} else if c.ConfigData.NetworkOS.Policy.OrphanPort != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.OrphanPort=%v", *c.ConfigData.NetworkOS.Policy.OrphanPort)
		v.ConfigData.NetworkOS.Policy.OrphanPort = new(bool)
		*v.ConfigData.NetworkOS.Policy.OrphanPort = *c.ConfigData.NetworkOS.Policy.OrphanPort
	}

	if v.ConfigData.NetworkOS.Policy.PathCost != nil && c.ConfigData.NetworkOS.Policy.PathCost != nil {
		if *v.ConfigData.NetworkOS.Policy.PathCost != *c.ConfigData.NetworkOS.Policy.PathCost {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PathCost=%v, c.ConfigData.NetworkOS.Policy.PathCost=%v", *v.ConfigData.NetworkOS.Policy.PathCost, *c.ConfigData.NetworkOS.Policy.PathCost)
		}
	} else if v.ConfigData.NetworkOS.Policy.PathCost != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PathCost=%v, c.ConfigData.NetworkOS.Policy.PathCost=nil", *v.ConfigData.NetworkOS.Policy.PathCost)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.PathCost != nil {
		v.ConfigData.NetworkOS.Policy.PathCost = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.PathCost=nil, c.ConfigData.NetworkOS.Policy.PathCost=%v", *c.ConfigData.NetworkOS.Policy.PathCost)
		*v.ConfigData.NetworkOS.Policy.PathCost = *c.ConfigData.NetworkOS.Policy.PathCost
	}

	if v.ConfigData.NetworkOS.Policy.Pfc != nil && c.ConfigData.NetworkOS.Policy.Pfc != nil {
		if *v.ConfigData.NetworkOS.Policy.Pfc != *c.ConfigData.NetworkOS.Policy.Pfc {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Pfc=%v, c.ConfigData.NetworkOS.Policy.Pfc=%v", *v.ConfigData.NetworkOS.Policy.Pfc, *c.ConfigData.NetworkOS.Policy.Pfc)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Pfc != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Pfc=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Pfc)
	} else if c.ConfigData.NetworkOS.Policy.Pfc != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Pfc=%v", *c.ConfigData.NetworkOS.Policy.Pfc)
		v.ConfigData.NetworkOS.Policy.Pfc = new(bool)
		*v.ConfigData.NetworkOS.Policy.Pfc = *c.ConfigData.NetworkOS.Policy.Pfc
	}

	if v.ConfigData.NetworkOS.Policy.PimDrPriority != nil && c.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		if *v.ConfigData.NetworkOS.Policy.PimDrPriority != *c.ConfigData.NetworkOS.Policy.PimDrPriority {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PimDrPriority=%v, c.ConfigData.NetworkOS.Policy.PimDrPriority=%v", *v.ConfigData.NetworkOS.Policy.PimDrPriority, *c.ConfigData.NetworkOS.Policy.PimDrPriority)
		}
	} else if v.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PimDrPriority=%v, c.ConfigData.NetworkOS.Policy.PimDrPriority=nil", *v.ConfigData.NetworkOS.Policy.PimDrPriority)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.PimDrPriority != nil {
		v.ConfigData.NetworkOS.Policy.PimDrPriority = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.PimDrPriority=nil, c.ConfigData.NetworkOS.Policy.PimDrPriority=%v", *c.ConfigData.NetworkOS.Policy.PimDrPriority)
		*v.ConfigData.NetworkOS.Policy.PimDrPriority = *c.ConfigData.NetworkOS.Policy.PimDrPriority
	}

	if v.ConfigData.NetworkOS.Policy.PimSparse != nil && c.ConfigData.NetworkOS.Policy.PimSparse != nil {
		if *v.ConfigData.NetworkOS.Policy.PimSparse != *c.ConfigData.NetworkOS.Policy.PimSparse {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PimSparse=%v, c.ConfigData.NetworkOS.Policy.PimSparse=%v", *v.ConfigData.NetworkOS.Policy.PimSparse, *c.ConfigData.NetworkOS.Policy.PimSparse)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.PimSparse != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PimSparse=%v, c.nil", *v.ConfigData.NetworkOS.Policy.PimSparse)
	} else if c.ConfigData.NetworkOS.Policy.PimSparse != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.PimSparse=%v", *c.ConfigData.NetworkOS.Policy.PimSparse)
		v.ConfigData.NetworkOS.Policy.PimSparse = new(bool)
		*v.ConfigData.NetworkOS.Policy.PimSparse = *c.ConfigData.NetworkOS.Policy.PimSparse
	}

	if v.ConfigData.NetworkOS.Policy.PortChannelId != c.ConfigData.NetworkOS.Policy.PortChannelId {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PortChannelId=%v, c.ConfigData.NetworkOS.Policy.PortChannelId=%v", v.ConfigData.NetworkOS.Policy.PortChannelId, c.ConfigData.NetworkOS.Policy.PortChannelId)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PortChannelMode != c.ConfigData.NetworkOS.Policy.PortChannelMode {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PortChannelMode=%v, c.ConfigData.NetworkOS.Policy.PortChannelMode=%v", v.ConfigData.NetworkOS.Policy.PortChannelMode, c.ConfigData.NetworkOS.Policy.PortChannelMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil && c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		if *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v, c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v", *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk, *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v, c.nil", *v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
	} else if c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk=%v", *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk)
		v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk = new(bool)
		*v.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk = *c.ConfigData.NetworkOS.Policy.PortTypeEdgeTrunk
	}

	if v.ConfigData.NetworkOS.Policy.Prefix != c.ConfigData.NetworkOS.Policy.Prefix {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Prefix=%v, c.ConfigData.NetworkOS.Policy.Prefix=%v", v.ConfigData.NetworkOS.Policy.Prefix, c.ConfigData.NetworkOS.Policy.Prefix)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.Ptp != nil && c.ConfigData.NetworkOS.Policy.Ptp != nil {
		if *v.ConfigData.NetworkOS.Policy.Ptp != *c.ConfigData.NetworkOS.Policy.Ptp {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ptp=%v, c.ConfigData.NetworkOS.Policy.Ptp=%v", *v.ConfigData.NetworkOS.Policy.Ptp, *c.ConfigData.NetworkOS.Policy.Ptp)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Ptp != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Ptp=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Ptp)
	} else if c.ConfigData.NetworkOS.Policy.Ptp != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Ptp=%v", *c.ConfigData.NetworkOS.Policy.Ptp)
		v.ConfigData.NetworkOS.Policy.Ptp = new(bool)
		*v.ConfigData.NetworkOS.Policy.Ptp = *c.ConfigData.NetworkOS.Policy.Ptp
	}

	if v.ConfigData.NetworkOS.Policy.PtpProfile != c.ConfigData.NetworkOS.Policy.PtpProfile {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PtpProfile=%v, c.ConfigData.NetworkOS.Policy.PtpProfile=%v", v.ConfigData.NetworkOS.Policy.PtpProfile, c.ConfigData.NetworkOS.Policy.PtpProfile)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PtpRole != nil && c.ConfigData.NetworkOS.Policy.PtpRole != nil {
		if *v.ConfigData.NetworkOS.Policy.PtpRole != *c.ConfigData.NetworkOS.Policy.PtpRole {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PtpRole=%v, c.ConfigData.NetworkOS.Policy.PtpRole=%v", *v.ConfigData.NetworkOS.Policy.PtpRole, *c.ConfigData.NetworkOS.Policy.PtpRole)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.PtpRole != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PtpRole=%v, c.nil", *v.ConfigData.NetworkOS.Policy.PtpRole)
	} else if c.ConfigData.NetworkOS.Policy.PtpRole != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.PtpRole=%v", *c.ConfigData.NetworkOS.Policy.PtpRole)
		v.ConfigData.NetworkOS.Policy.PtpRole = new(bool)
		*v.ConfigData.NetworkOS.Policy.PtpRole = *c.ConfigData.NetworkOS.Policy.PtpRole
	}

	if v.ConfigData.NetworkOS.Policy.PtpVlan != nil && c.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.PtpVlan != *c.ConfigData.NetworkOS.Policy.PtpVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PtpVlan=%v, c.ConfigData.NetworkOS.Policy.PtpVlan=%v", *v.ConfigData.NetworkOS.Policy.PtpVlan, *c.ConfigData.NetworkOS.Policy.PtpVlan)
		}
	} else if v.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PtpVlan=%v, c.ConfigData.NetworkOS.Policy.PtpVlan=nil", *v.ConfigData.NetworkOS.Policy.PtpVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.PtpVlan != nil {
		v.ConfigData.NetworkOS.Policy.PtpVlan = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.PtpVlan=nil, c.ConfigData.NetworkOS.Policy.PtpVlan=%v", *c.ConfigData.NetworkOS.Policy.PtpVlan)
		*v.ConfigData.NetworkOS.Policy.PtpVlan = *c.ConfigData.NetworkOS.Policy.PtpVlan
	}

	if v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans != c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans=%v", v.ConfigData.NetworkOS.Policy.PvlanAllowedVlans, c.ConfigData.NetworkOS.Policy.PvlanAllowedVlans)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PvlanMode != c.ConfigData.NetworkOS.Policy.PvlanMode {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PvlanMode=%v, c.ConfigData.NetworkOS.Policy.PvlanMode=%v", v.ConfigData.NetworkOS.Policy.PvlanMode, c.ConfigData.NetworkOS.Policy.PvlanMode)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil && c.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		if *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v, c.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v", *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan, *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
		}
	} else if v.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v, c.ConfigData.NetworkOS.Policy.PvlanNativeVlan=nil", *v.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.PvlanNativeVlan != nil {
		v.ConfigData.NetworkOS.Policy.PvlanNativeVlan = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.PvlanNativeVlan=nil, c.ConfigData.NetworkOS.Policy.PvlanNativeVlan=%v", *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan)
		*v.ConfigData.NetworkOS.Policy.PvlanNativeVlan = *c.ConfigData.NetworkOS.Policy.PvlanNativeVlan
	}

	if v.ConfigData.NetworkOS.Policy.Qos != nil && c.ConfigData.NetworkOS.Policy.Qos != nil {
		if *v.ConfigData.NetworkOS.Policy.Qos != *c.ConfigData.NetworkOS.Policy.Qos {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.Qos=%v, c.ConfigData.NetworkOS.Policy.Qos=%v", *v.ConfigData.NetworkOS.Policy.Qos, *c.ConfigData.NetworkOS.Policy.Qos)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.Qos != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Qos=%v, c.nil", *v.ConfigData.NetworkOS.Policy.Qos)
	} else if c.ConfigData.NetworkOS.Policy.Qos != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.Qos=%v", *c.ConfigData.NetworkOS.Policy.Qos)
		v.ConfigData.NetworkOS.Policy.Qos = new(bool)
		*v.ConfigData.NetworkOS.Policy.Qos = *c.ConfigData.NetworkOS.Policy.Qos
	}

	if v.ConfigData.NetworkOS.Policy.QosPolicy != c.ConfigData.NetworkOS.Policy.QosPolicy {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.QosPolicy=%v, c.ConfigData.NetworkOS.Policy.QosPolicy=%v", v.ConfigData.NetworkOS.Policy.QosPolicy, c.ConfigData.NetworkOS.Policy.QosPolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.QueuingPolicy != c.ConfigData.NetworkOS.Policy.QueuingPolicy {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.QueuingPolicy=%v, c.ConfigData.NetworkOS.Policy.QueuingPolicy=%v", v.ConfigData.NetworkOS.Policy.QueuingPolicy, c.ConfigData.NetworkOS.Policy.QueuingPolicy)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.RoutingTag != c.ConfigData.NetworkOS.Policy.RoutingTag {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.RoutingTag=%v, c.ConfigData.NetworkOS.Policy.RoutingTag=%v", v.ConfigData.NetworkOS.Policy.RoutingTag, c.ConfigData.NetworkOS.Policy.RoutingTag)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil && c.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		if *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v, c.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v", *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy, *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v, c.nil", *v.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
	} else if c.ConfigData.NetworkOS.Policy.SecurityGroupProxy != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.SecurityGroupProxy=%v", *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy)
		v.ConfigData.NetworkOS.Policy.SecurityGroupProxy = new(bool)
		*v.ConfigData.NetworkOS.Policy.SecurityGroupProxy = *c.ConfigData.NetworkOS.Policy.SecurityGroupProxy
	}

	if v.ConfigData.NetworkOS.Policy.Speed != c.ConfigData.NetworkOS.Policy.Speed {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.Speed=%v, c.ConfigData.NetworkOS.Policy.Speed=%v", v.ConfigData.NetworkOS.Policy.Speed, c.ConfigData.NetworkOS.Policy.Speed)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControl != nil && c.ConfigData.NetworkOS.Policy.StormControl != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControl != *c.ConfigData.NetworkOS.Policy.StormControl {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControl=%v, c.ConfigData.NetworkOS.Policy.StormControl=%v", *v.ConfigData.NetworkOS.Policy.StormControl, *c.ConfigData.NetworkOS.Policy.StormControl)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControl != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControl=%v, c.nil", *v.ConfigData.NetworkOS.Policy.StormControl)
	} else if c.ConfigData.NetworkOS.Policy.StormControl != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.StormControl=%v", *c.ConfigData.NetworkOS.Policy.StormControl)
		v.ConfigData.NetworkOS.Policy.StormControl = new(bool)
		*v.ConfigData.NetworkOS.Policy.StormControl = *c.ConfigData.NetworkOS.Policy.StormControl
	}

	if v.ConfigData.NetworkOS.Policy.StormControlAction != c.ConfigData.NetworkOS.Policy.StormControlAction {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlAction=%v, c.ConfigData.NetworkOS.Policy.StormControlAction=%v", v.ConfigData.NetworkOS.Policy.StormControlAction, c.ConfigData.NetworkOS.Policy.StormControlAction)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel, *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v, c.nil", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
	} else if c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel)
		v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel = new(float64)
		*v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel = *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevel
	}

	if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=nil", *v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps != nil {
		v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=nil, c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps)
		*v.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps = *c.ConfigData.NetworkOS.Policy.StormControlBroadcastLevelPps
	}

	if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel, *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v, c.nil", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
	} else if c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel)
		v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel = new(float64)
		*v.ConfigData.NetworkOS.Policy.StormControlMulticastLevel = *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevel
	}

	if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=nil", *v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps != nil {
		v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=nil, c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps)
		*v.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps = *c.ConfigData.NetworkOS.Policy.StormControlMulticastLevelPps
	}

	if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil && c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel, *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v, c.nil", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
	} else if c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel=%v", *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel)
		v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel = new(float64)
		*v.ConfigData.NetworkOS.Policy.StormControlUnicastLevel = *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevel
	}

	if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil && c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		if *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps, *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
		}
	} else if v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=nil", *v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps != nil {
		v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=nil, c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps=%v", *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps)
		*v.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps = *c.ConfigData.NetworkOS.Policy.StormControlUnicastLevelPps
	}

	if v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil && c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		if *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v, c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v", *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent, *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
		}
	} else if v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v, c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=nil", *v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent != nil {
		v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=nil, c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent=%v", *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent)
		*v.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent = *c.ConfigData.NetworkOS.Policy.UnicastBandwidthPercent
	}

	if v.ConfigData.NetworkOS.Policy.V6prefix != nil && c.ConfigData.NetworkOS.Policy.V6prefix != nil {
		if *v.ConfigData.NetworkOS.Policy.V6prefix != *c.ConfigData.NetworkOS.Policy.V6prefix {
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.V6prefix=%v, c.ConfigData.NetworkOS.Policy.V6prefix=%v", *v.ConfigData.NetworkOS.Policy.V6prefix, *c.ConfigData.NetworkOS.Policy.V6prefix)
		}
	} else if v.ConfigData.NetworkOS.Policy.V6prefix != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.V6prefix=%v, c.ConfigData.NetworkOS.Policy.V6prefix=nil", *v.ConfigData.NetworkOS.Policy.V6prefix)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	} else if c.ConfigData.NetworkOS.Policy.V6prefix != nil {
		v.ConfigData.NetworkOS.Policy.V6prefix = new(int64)
		log.Printf("Copy from state: v.ConfigData.NetworkOS.Policy.V6prefix=nil, c.ConfigData.NetworkOS.Policy.V6prefix=%v", *c.ConfigData.NetworkOS.Policy.V6prefix)
		*v.ConfigData.NetworkOS.Policy.V6prefix = *c.ConfigData.NetworkOS.Policy.V6prefix
	}

	if v.ConfigData.NetworkOS.Policy.VlanMapping != nil && c.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		if *v.ConfigData.NetworkOS.Policy.VlanMapping != *c.ConfigData.NetworkOS.Policy.VlanMapping {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.VlanMapping=%v, c.ConfigData.NetworkOS.Policy.VlanMapping=%v", *v.ConfigData.NetworkOS.Policy.VlanMapping, *c.ConfigData.NetworkOS.Policy.VlanMapping)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.VlanMapping=%v, c.nil", *v.ConfigData.NetworkOS.Policy.VlanMapping)
	} else if c.ConfigData.NetworkOS.Policy.VlanMapping != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.VlanMapping=%v", *c.ConfigData.NetworkOS.Policy.VlanMapping)
		v.ConfigData.NetworkOS.Policy.VlanMapping = new(bool)
		*v.ConfigData.NetworkOS.Policy.VlanMapping = *c.ConfigData.NetworkOS.Policy.VlanMapping
	}

	if v.ConfigData.NetworkOS.Policy.VrfInterface != c.ConfigData.NetworkOS.Policy.VrfInterface {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.VrfInterface=%v, c.ConfigData.NetworkOS.Policy.VrfInterface=%v", v.ConfigData.NetworkOS.Policy.VrfInterface, c.ConfigData.NetworkOS.Policy.VrfInterface)
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}

	if v.ConfigData.NetworkOS.Policy.WanLink != nil && c.ConfigData.NetworkOS.Policy.WanLink != nil {
		if *v.ConfigData.NetworkOS.Policy.WanLink != *c.ConfigData.NetworkOS.Policy.WanLink {
			log.Printf("Update: v.ConfigData.NetworkOS.Policy.WanLink=%v, c.ConfigData.NetworkOS.Policy.WanLink=%v", *v.ConfigData.NetworkOS.Policy.WanLink, *c.ConfigData.NetworkOS.Policy.WanLink)
			if action == ActionNone || action == RequiresUpdate {
				action = RequiresUpdate
			}
		}
	} else if v.ConfigData.NetworkOS.Policy.WanLink != nil {
		log.Printf("Update: v.ConfigData.NetworkOS.Policy.WanLink=%v, c.nil", *v.ConfigData.NetworkOS.Policy.WanLink)
	} else if c.ConfigData.NetworkOS.Policy.WanLink != nil {
		log.Printf("Copy from state: v.nil, c.ConfigData.NetworkOS.Policy.WanLink=%v", *c.ConfigData.NetworkOS.Policy.WanLink)
		v.ConfigData.NetworkOS.Policy.WanLink = new(bool)
		*v.ConfigData.NetworkOS.Policy.WanLink = *c.ConfigData.NetworkOS.Policy.WanLink
	}

	if len(v.ConfigData.NetworkOS.Policy.AssociationList) != len(c.ConfigData.NetworkOS.Policy.AssociationList) {
		log.Printf("AssociationList-Update: len(v.ConfigData.NetworkOS.Policy.AssociationList)=%d, len(c.ConfigData.NetworkOS.Policy.AssociationList)=%d", len(v.ConfigData.NetworkOS.Policy.AssociationList), len(c.ConfigData.NetworkOS.Policy.AssociationList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.AssociationList {
		retVal := v.ConfigData.NetworkOS.Policy.AssociationList[i].CreatePlan(c.ConfigData.NetworkOS.Policy.AssociationList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MappingList) != len(c.ConfigData.NetworkOS.Policy.MappingList) {
		log.Printf("MappingList-Update: len(v.ConfigData.NetworkOS.Policy.MappingList)=%d, len(c.ConfigData.NetworkOS.Policy.MappingList)=%d", len(v.ConfigData.NetworkOS.Policy.MappingList), len(c.ConfigData.NetworkOS.Policy.MappingList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.MappingList {
		retVal := v.ConfigData.NetworkOS.Policy.MappingList[i].CreatePlan(c.ConfigData.NetworkOS.Policy.MappingList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges) != len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges) {
		log.Printf("IgmpGroupRanges-Update: len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges)=%d, len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges)=%d", len(v.ConfigData.NetworkOS.Policy.IgmpGroupRanges), len(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.IgmpGroupRanges {
		retVal := v.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i].CreatePlan(c.ConfigData.NetworkOS.Policy.IgmpGroupRanges[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MstPathCostList) != len(c.ConfigData.NetworkOS.Policy.MstPathCostList) {
		log.Printf("MstPathCostList-Update: len(v.ConfigData.NetworkOS.Policy.MstPathCostList)=%d, len(c.ConfigData.NetworkOS.Policy.MstPathCostList)=%d", len(v.ConfigData.NetworkOS.Policy.MstPathCostList), len(c.ConfigData.NetworkOS.Policy.MstPathCostList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.MstPathCostList {
		retVal := v.ConfigData.NetworkOS.Policy.MstPathCostList[i].CreatePlan(c.ConfigData.NetworkOS.Policy.MstPathCostList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.MstPortPrioList) != len(c.ConfigData.NetworkOS.Policy.MstPortPrioList) {
		log.Printf("MstPortPrioList-Update: len(v.ConfigData.NetworkOS.Policy.MstPortPrioList)=%d, len(c.ConfigData.NetworkOS.Policy.MstPortPrioList)=%d", len(v.ConfigData.NetworkOS.Policy.MstPortPrioList), len(c.ConfigData.NetworkOS.Policy.MstPortPrioList))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.MstPortPrioList {
		retVal := v.ConfigData.NetworkOS.Policy.MstPortPrioList[i].CreatePlan(c.ConfigData.NetworkOS.Policy.MstPortPrioList[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}

	if len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries) != len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries) {
		log.Printf("VlanMappingEntries-Update: len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries)=%d, len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries)=%d", len(v.ConfigData.NetworkOS.Policy.VlanMappingEntries), len(c.ConfigData.NetworkOS.Policy.VlanMappingEntries))
		if action == ActionNone || action == RequiresUpdate {
			action = RequiresUpdate
		}
	}
	for i := range v.ConfigData.NetworkOS.Policy.VlanMappingEntries {
		retVal := v.ConfigData.NetworkOS.Policy.VlanMappingEntries[i].CreatePlan(c.ConfigData.NetworkOS.Policy.VlanMappingEntries[i], cf)
		if retVal != ActionNone {
			if action == ActionNone || action == RequiresUpdate {
				action = retVal
			}
		}
	}
	return action
}
