// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_change_control

import (
	"context"
	"encoding/json"
	"fmt"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// rscGetChangeControl retrieves the singleton change control settings.
func (r *changeControlResource) rscGetChangeControl(ctx context.Context, dg *diag.Diagnostics, model *ChangeControlModel) {
	changeControlAPI := api.NewChangeControlAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	tflog.Debug(ctx, "Reading change control settings", map[string]interface{}{
		"resource": changeControlImportID,
	})

	respData, err := changeControlAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Reading Change Control",
			fmt.Sprintf("Could not read change control settings, unexpected error: %s", err.Error()),
		)
		return
	}

	var data NDFCChangeControlModel
	if err := json.Unmarshal(respData, &data); err != nil {
		dg.AddError(
			"Error Reading Change Control",
			fmt.Sprintf("Could not unmarshal change control response, unexpected error: %s", err.Error()),
		)
		return
	}

	dg.Append(model.SetModelData(&data)...)
	if !dg.HasError() {
		model.Id = types.StringValue(changeControlImportID)
	}
}

// rscPutChangeControl writes the singleton settings using the API's PUT
// operation. Create and Update apply the configured settings, while Delete
// uses the same operation to restore the defaults.
func (r *changeControlResource) rscPutChangeControl(ctx context.Context, dg *diag.Diagnostics, model *ChangeControlModel) {
	changeControlAPI := api.NewChangeControlAPI(r.infraClient.ApiClient, ndapi.DefaultFabric)

	payload, err := json.Marshal(model.GetModelData())
	if err != nil {
		dg.AddError(
			"Error Updating Change Control",
			fmt.Sprintf("Could not marshal change control payload, unexpected error: %s", err.Error()),
		)
		return
	}

	tflog.Debug(ctx, "Updating change control settings", map[string]interface{}{
		"resource": changeControlImportID,
	})

	res, err := changeControlAPI.Put(payload, nil)
	if err != nil {
		dg.AddError(
			"Error Updating Change Control",
			fmt.Sprintf("Could not update change control settings, unexpected error: %s; response: %s", err.Error(), res.String()),
		)
		return
	}

	r.rscGetChangeControl(ctx, dg, model)
}
