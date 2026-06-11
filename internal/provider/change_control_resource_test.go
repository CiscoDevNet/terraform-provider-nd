// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_change_control"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	nd "github.com/netascode/go-nd"
)

func newChangeControlAcceptanceClient(t *testing.T) *nd.Client {
	t.Helper()

	cfg := helper.GetConfig("global")
	client, err := nd.NewClient(
		cfg.ND.URL,
		"/api/v1",
		cfg.ND.User,
		cfg.ND.Password,
		"",
		cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("failed to create ND client for change control acceptance test: %s", err.Error())
	}

	return &client
}

func putChangeControlOutsideTerraform(
	t *testing.T,
	client *nd.Client,
	model *resource_change_control.NDFCChangeControlModel,
) {
	t.Helper()

	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("failed to marshal change control settings for out-of-band PUT: %s", err.Error())
	}

	changeControlAPI := api.NewChangeControlAPI(client, ndapi.DefaultFabric)
	res, err := changeControlAPI.Put(payload, nil)
	if err != nil {
		t.Fatalf("failed to PUT change control settings out of band: %s %s", err.Error(), res.String())
	}
}

func checkChangeControlBackend(
	client *nd.Client,
	expected *resource_change_control.NDFCChangeControlModel,
) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		changeControlAPI := api.NewChangeControlAPI(client, ndapi.DefaultFabric)
		respData, err := changeControlAPI.Get()
		if err != nil {
			return fmt.Errorf("failed to GET change control settings: %s", err.Error())
		}

		var actual resource_change_control.NDFCChangeControlModel
		if err := json.Unmarshal(respData, &actual); err != nil {
			return fmt.Errorf("failed to unmarshal change control settings: %s", err.Error())
		}

		if !reflect.DeepEqual(&actual, expected) {
			actualJSON, _ := json.Marshal(actual)
			expectedJSON, _ := json.Marshal(expected)
			return fmt.Errorf("unexpected change control settings: got %s, expected %s", actualJSON, expectedJSON)
		}

		return nil
	}
}

func changeControlModel(
	adminStatus bool,
	orchestration bool,
	ndManagedFabrics bool,
	bypassTelemetry bool,
	numberOfApprovers int64,
	allowSelfApproval bool,
	ticketNamePrefix string,
) *resource_change_control.NDFCChangeControlModel {
	return &resource_change_control.NDFCChangeControlModel{
		AdminStatus:                  &adminStatus,
		Orchestration:                &orchestration,
		NdManagedFabrics:             &ndManagedFabrics,
		BypassTelemetryChangeControl: &bypassTelemetry,
		NumberOfApprovers:            &numberOfApprovers,
		AllowSelfApproval:            &allowSelfApproval,
		TicketNamePrefix:             ticketNamePrefix,
	}
}

func changeControlStateChecks(
	resourceName string,
	model resource_change_control.NDFCChangeControlModel,
) []resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "id", "changeControl"),
	}

	return append(checks, ChangeControlModelHelperStateCheck(resourceName, model, path.Empty())...)
}

func TestAccChangeControlResource(t *testing.T) {
	cfg := helper.GetConfig("global")
	client := newChangeControlAcceptanceClient(t)
	cleanupEnabled := false
	defer func() {
		if cleanupEnabled {
			putChangeControlOutsideTerraform(t, client, changeControlModel(false, false, false, false, 1, true, "TICKET_"))
		}
	}()

	x := &map[string]string{
		"RscType":  "nd_change_control",
		"RscName":  "change_control_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	changeControl := new(resource_change_control.NDFCChangeControlModel)
	resourceName := "nd_change_control.change_control_test"

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t, "global")
			cleanupEnabled = true
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Manage the pre-existing change control singleton with schema defaults", t.Name())

					helper.GenerateChangeControlObject(&changeControl, map[string]interface{}{})
					helper.GetTFConfigWithSingleResource(s1.Name, *x, []interface{}{changeControl}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					putChangeControlOutsideTerraform(t, client, changeControlModel(false, false, false, false, 1, true, "TICKET_"))
				},
				Check: resource.ComposeTestCheckFunc(changeControlStateChecks(resourceName, *changeControl)...),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - Enable non-default change control settings", t.Name())

					helper.ModifyChangeControlObject(&changeControl, map[string]interface{}{
						"admin_status":                    true,
						"orchestration":                   true,
						"nd_managed_fabrics":              false,
						"bypass_telemetry_change_control": true,
						"number_of_approvers":             3,
						"allow_self_approval":             false,
						"ticket_name_prefix":              "TF_CC_",
					})
					helper.GetTFConfigWithSingleResource(s2.Name, *x, []interface{}{changeControl}, &tfConfig)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, s2.Cfg) },
				Check:     resource.ComposeTestCheckFunc(changeControlStateChecks(resourceName, *changeControl)...),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - Update all supported mutable change control settings", t.Name())

					helper.ModifyChangeControlObject(&changeControl, map[string]interface{}{
						"admin_status":                    true,
						"orchestration":                   true,
						"nd_managed_fabrics":              false,
						"bypass_telemetry_change_control": false,
						"number_of_approvers":             5,
						"allow_self_approval":             true,
						"ticket_name_prefix":              "UPDATED_CC_",
					})
					helper.GetTFConfigWithSingleResource(s3.Name, *x, []interface{}{changeControl}, &tfConfig)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, s3.Cfg) },
				Check:     resource.ComposeTestCheckFunc(changeControlStateChecks(resourceName, *changeControl)...),
			},
			{
				PreConfig: func() {
					s4.Index = 4
					s4.Name = fmt.Sprintf("%s - Import the singleton and verify canonical GET state", t.Name())
					s4.Cfg = *tfConfig
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
				},
				// The canonical GET returns every configurable schema attribute, and
				// the resource assigns the fixed singleton identifier after the GET.
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        "changeControl",
				ImportStateVerifyIdentifierAttribute: "id",
			},
			{
				Config: func() string {
					s5.Index = 5
					s5.Name = fmt.Sprintf("%s - Detect out-of-band mutable setting drift", t.Name())
					helper.GetTFConfigWithSingleResource(s5.Name, *x, []interface{}{changeControl}, &tfConfig)
					s5.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s5.Index, s5.Name, s5.Cfg)
					putChangeControlOutsideTerraform(t, client, changeControlModel(false, false, false, false, 1, true, "DRIFT_CC_"))
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: func() string {
					s6.Index = 6
					s6.Name = fmt.Sprintf("%s - Reconcile the out-of-band change control drift", t.Name())
					helper.GetTFConfigWithSingleResource(s6.Name, *x, []interface{}{changeControl}, &tfConfig)
					s6.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s6.Index, s6.Name, s6.Cfg) },
				Check:     resource.ComposeTestCheckFunc(changeControlStateChecks(resourceName, *changeControl)...),
			},
			{
				Config: func() string {
					s7.Index = 7
					s7.Name = fmt.Sprintf("%s - Reset the singleton to defaults and remove it from Terraform state", t.Name())
					helper.GetTFConfigWithSingleResource(s7.Name, *x, []interface{}{changeControl}, &tfConfig)
					s7.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s7.Index, s7.Name, s7.Cfg) },
				Destroy:   true,
				Check: checkChangeControlBackend(
					client,
					changeControlModel(false, false, false, false, 1, true, "TICKET_"),
				),
			},
		},
	})
}

func TestAccChangeControlResourceInvalidConfiguration(t *testing.T) {
	cfg := helper.GetConfig("global")
	x := &map[string]string{
		"RscType":  "nd_change_control",
		"RscName":  "invalid_change_control_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	changeControl := new(resource_change_control.NDFCChangeControlModel)
	s1 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - Reject admin status without an enabled change control target", t.Name())
					helper.GenerateChangeControlObject(&changeControl, map[string]interface{}{
						"admin_status":       true,
						"orchestration":      false,
						"nd_managed_fabrics": false,
					})
					helper.GetTFConfigWithSingleResource(s1.Name, *x, []interface{}{changeControl}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, s1.Cfg) },
				ExpectError: regexp.MustCompile(
					`admin_status can be enabled only when orchestration or nd_managed_fabrics is\s+enabled`,
				),
			},
		},
	})
}
