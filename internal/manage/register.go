// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package manage

import (
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// GetResources returns all resources for the manage module.
// Resources auto-register via init() in their packages.
func GetResources() []func() resource.Resource {
	return registry.GetResources(ModuleKey)
}

// GetDataSources returns all data sources for the manage module.
// Data sources auto-register via init() in their packages.
func GetDataSources() []func() datasource.DataSource {
	return registry.GetDataSources(ModuleKey)
}
