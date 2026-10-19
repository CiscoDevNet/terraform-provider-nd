// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_ethernet

import (
	"terraform-provider-nd/internal/manage/resource_interface_common"
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewInterfaceEthernetResource)
	resource_interface_common.RegisterInterfaceHandler(interfaceTypeEthernet, ethernetHandler{})
}
