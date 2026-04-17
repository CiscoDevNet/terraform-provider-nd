// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package deployment

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// ConfigSave performs a config-save (recalculate) operation on a fabric
func ConfigSave(ctx context.Context, client *nd.Client, fabricName string, dg *diag.Diagnostics) (string, error) {
	configAPI := api.NewConfigAPI(nil, client)
	configAPI.FabricName = fabricName
	configAPI.SetOperation(api.OpConfigSave)

	resp, err := configAPI.Post(nil)
	if err != nil {
		return resp.String(), fmt.Errorf("config save failed for fabric %s: %w", fabricName, err)
	}

	tflog.Info(ctx, "Fabric config saved (recalculate)", map[string]interface{}{
		"fabric_name": fabricName,
	})
	return resp.String(), nil
}

// ConfigDeploy attempts a config-deploy and returns any error.
func ConfigDeploy(ctx context.Context, client *nd.Client, fabricName string) (string, error) {
	configAPI := api.NewConfigAPI(nil, client)
	configAPI.FabricName = fabricName
	configAPI.SetOperation(api.OpConfigDeploy)

	resp, err := configAPI.DeployPost(nil)
	if err != nil {
		return resp.String(), fmt.Errorf("config deploy failed for fabric %s: %w", fabricName, err)
	}

	tflog.Info(ctx, "Fabric config deployed", map[string]interface{}{
		"fabric_name": fabricName,
	})
	return resp.String(), nil
}

// ConfigSaveAndDeploy performs both config-save and config-deploy operations
func ConfigSaveAndDeploy(ctx context.Context, client *nd.Client, fabricName string, recalculate bool, deploy bool, dg *diag.Diagnostics) {
	if recalculate {
		respMsg, err := ConfigSave(ctx, client, fabricName, dg)
		if err != nil {
			dg.AddError("Error Saving Config", fmt.Sprintf("%v: %s", err, respMsg))
			return
		}
	}

	if deploy {
		respMsg, err := ConfigDeploy(ctx, client, fabricName)
		if err != nil {
			dg.AddError("Error Deploying Config", fmt.Sprintf("%v: %s", err, respMsg))
			return
		}
	}
}
