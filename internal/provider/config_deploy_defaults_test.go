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
	"testing"

	"terraform-provider-nd/internal/manage/resource_config_deploy"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// TestConfigDeployBehaviourDefaultsWhenOmitted verifies, empirically, whether the
// terraform-plugin-framework applies the nested `behaviour.*` StaticBool defaults
// when the user omits the `behaviour` block entirely.
//
// It runs a real PlanResourceChange through the provider server with a config that
// sets only fabric_name + deploy, and inspects the planned `behaviour` value.
func TestConfigDeployBehaviourDefaultsWhenOmitted(t *testing.T) {
	ctx := context.Background()

	// Build the resource schema and its tftypes object type.
	sch := resource_config_deploy.ConfigDeployResourceSchema(ctx)
	objType, ok := sch.Type().TerraformType(ctx).(tftypes.Object)
	if !ok {
		t.Fatalf("expected object type, got %T", sch.Type().TerraformType(ctx))
	}
	attrTypes := objType.AttributeTypes

	// Config: only fabric_name + deploy set; behaviour omitted (null).
	configVal := tftypes.NewValue(objType, map[string]tftypes.Value{
		"id":                                tftypes.NewValue(attrTypes["id"], nil),
		"fabric_name":                       tftypes.NewValue(attrTypes["fabric_name"], "test_fabric"),
		"deploy":                            tftypes.NewValue(attrTypes["deploy"], true),
		"config_save":                       tftypes.NewValue(attrTypes["config_save"], nil),
		"switch_ids":                        tftypes.NewValue(attrTypes["switch_ids"], nil),
		"force_show_run":                    tftypes.NewValue(attrTypes["force_show_run"], nil),
		"include_all_fabric_group_switches": tftypes.NewValue(attrTypes["include_all_fabric_group_switches"], nil),
		"always_deploy":                     tftypes.NewValue(attrTypes["always_deploy"], nil),
		"behaviour":                         tftypes.NewValue(attrTypes["behaviour"], nil),
		"status":                            tftypes.NewValue(attrTypes["status"], nil),
		"ticket_id":                         tftypes.NewValue(attrTypes["ticket_id"], nil),
	})

	priorVal := tftypes.NewValue(objType, nil) // create: no prior state

	configDV, err := tfprotov6.NewDynamicValue(objType, configVal)
	if err != nil {
		t.Fatalf("encode config: %v", err)
	}
	proposedDV, err := tfprotov6.NewDynamicValue(objType, configVal)
	if err != nil {
		t.Fatalf("encode proposed: %v", err)
	}
	priorDV, err := tfprotov6.NewDynamicValue(objType, priorVal)
	if err != nil {
		t.Fatalf("encode prior: %v", err)
	}

	server, err := testAccProtoV6ProviderFactories["nd"]()
	if err != nil {
		t.Fatalf("create provider server: %v", err)
	}

	resp, err := server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{
		TypeName:         "nd_config_deploy",
		Config:           &configDV,
		PriorState:       &priorDV,
		ProposedNewState: &proposedDV,
	})
	if err != nil {
		t.Fatalf("PlanResourceChange: %v", err)
	}
	for _, d := range resp.Diagnostics {
		t.Logf("diagnostic: severity=%v summary=%q detail=%q", d.Severity, d.Summary, d.Detail)
	}
	if resp.PlannedState == nil {
		t.Fatal("planned state is nil")
	}

	plannedVal, err := resp.PlannedState.Unmarshal(objType)
	if err != nil {
		t.Fatalf("unmarshal planned state: %v", err)
	}

	var planned map[string]tftypes.Value
	if err := plannedVal.As(&planned); err != nil {
		t.Fatalf("planned.As: %v", err)
	}

	behaviour := planned["behaviour"]
	t.Logf("planned behaviour: null=%v known=%v value=%s",
		behaviour.IsNull(), behaviour.IsKnown(), behaviour.String())

	if behaviour.IsNull() || !behaviour.IsKnown() {
		t.Fatalf("behaviour should be defaulted to a known object when omitted, got null=%v known=%v",
			behaviour.IsNull(), behaviour.IsKnown())
	}

	// Block is present and known — verify nested defaults.
	var nested map[string]tftypes.Value
	if err := behaviour.As(&nested); err != nil {
		t.Fatalf("behaviour.As: %v", err)
	}

	assertBool := func(name string, want bool) {
		v := nested[name]
		if v.IsNull() || !v.IsKnown() {
			t.Fatalf("behaviour.%s should be known, got null=%v known=%v", name, v.IsNull(), v.IsKnown())
		}
		var got bool
		if err := v.As(&got); err != nil {
			t.Fatalf("behaviour.%s As: %v", name, err)
		}
		if got != want {
			t.Fatalf("behaviour.%s = %v, want %v", name, got, want)
		}
		t.Logf("behaviour.%s = %v (ok)", name, got)
	}

	assertBool("deploy_on_create", true)
	assertBool("deploy_on_update", true)
	assertBool("deploy_on_destroy", false)
	t.Log("RESULT: omitted behaviour block now materializes {create=true, update=true, destroy=false}.")
}
