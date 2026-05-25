// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func getStringAttributeValue(ctx context.Context, attributeValue basetypes.StringValue, envKey string) string {
	if attributeValue.ValueString() == "" {
		return os.Getenv(envKey)
	}
	return attributeValue.ValueString()
}

// stepInfo carries the per-step name and rendered HCL from build time
// (when Config is evaluated) to step-execution time (when PreConfig fires).
// This lets us print the step header / snapshot path / TF config inline with
// the actual API call logs instead of all at once at test setup.
type stepInfo struct {
	name string
	cfg  string
}
