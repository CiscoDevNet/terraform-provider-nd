// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_interface_common

import (
	"context"
	"strings"

	"terraform-provider-nd/internal/common/templates"
)

// resolveCollisionTypeFromTemplateDB queries the live ND TemplateDB for
// the declared parameterType of fieldName under the given policyType.
// Returns "number" or "string" when deterministic, or "" when the
// TemplateDB singleton is not yet initialised (e.g. unit tests) or the
// template/field is unknown.
func resolveCollisionTypeFromTemplateDB(policyType, fieldName string) string {
	// Guard: TemplateDB may not be initialised in unit tests.
	defer func() { recover() }() //nolint:errcheck

	tdb := templates.GetInstance()
	pt := tdb.GetParamTypeForPolicy(context.Background(), policyType, fieldName)
	if pt == "" {
		return ""
	}
	pt = strings.ToLower(pt)
	switch {
	case pt == "integer" || pt == "long" || pt == "float" || pt == "number":
		return "number"
	case pt == "string" || pt == "enum":
		return "string"
	default:
		return ""
	}
}
