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
	"regexp"
	"strings"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/netascode/go-nd"
)

const (
	multiClusterConnectivityHostname     = "198.18.133.100"
	multiClusterConnectivityClusterName  = "nexus-dashboard"
	multiClusterConnectivityCreateDomain = "local1"
	multiClusterConnectivityUpdateDomain = "local"
)

func newMultiClusterConnectivityTestClient(t *testing.T) *nd.Client {
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
		t.Fatalf("failed to create ND client for multi-cluster connectivity acceptance helper: %v", err)
	}

	return &client
}

func multiClusterConnectivityNameOutsideTerraform(t *testing.T, client *nd.Client, model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) string {
	t.Helper()

	if model.Spec.ClusterName != "" {
		return model.Spec.ClusterName
	}

	clusterAPI := api.NewClusterAPI(client)
	respData, err := clusterAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return ""
		}
		t.Fatalf("failed to list multi-cluster connectivity objects outside Terraform: %v", err)
	}

	var clustersResp map[string][]resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel
	if err := json.Unmarshal(respData, &clustersResp); err != nil {
		t.Fatalf("failed to decode multi-cluster connectivity list response: %v", err)
	}

	for _, cluster := range clustersResp["clusters"] {
		if cluster.Spec.Hostname == model.Spec.Hostname &&
			(cluster.Spec.ClusterType == "" || cluster.Spec.ClusterType == "ND") {
			return cluster.Spec.ClusterName
		}
	}

	return ""
}

func deleteMultiClusterConnectivityOutsideTerraform(t *testing.T, model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) {
	t.Helper()

	client := newMultiClusterConnectivityTestClient(t)
	clusterName := multiClusterConnectivityNameOutsideTerraform(t, client, model)
	if clusterName == "" {
		t.Fatalf("failed to resolve multi-cluster connectivity cluster name for hostname %q", model.Spec.Hostname)
	}

	clusterAPI := api.NewClusterAPI(client)
	clusterAPI.ClusterName = clusterName
	clusterAPI.Delete = true

	payload := []byte("{}")

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			t.Logf("multi-cluster connectivity %q already absent before Terraform delete", clusterName)
			return
		}
		t.Fatalf("failed to delete multi-cluster connectivity %q outside Terraform: %v %v", clusterName, err, res)
	}
}

func deleteMultiClusterConnectivityIfExistsOutsideTerraform(t *testing.T, model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) {
	t.Helper()

	client := newMultiClusterConnectivityTestClient(t)
	clusterName := multiClusterConnectivityNameOutsideTerraform(t, client, model)
	if clusterName == "" {
		return
	}

	deleteMultiClusterConnectivityOutsideTerraform(t, model)
}

func createMultiClusterConnectivityOutsideTerraform(t *testing.T, model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) bool {
	t.Helper()

	client := newMultiClusterConnectivityTestClient(t)
	if clusterName := multiClusterConnectivityNameOutsideTerraform(t, client, model); clusterName != "" {
		t.Logf("multi-cluster connectivity %q already exists before manual create", clusterName)
		return false
	}

	preexisting := *model
	preexisting.Spec.ClusterType = "ND"
	payload, err := json.Marshal(preexisting)
	if err != nil {
		t.Fatalf("failed to marshal manual multi-cluster connectivity payload: %v", err)
	}

	clusterAPI := api.NewClusterAPI(client)
	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to create multi-cluster connectivity outside Terraform: %v %v", err, res)
	}
	return true
}

func TestAccMultiClusterConnectivityResourcePreExistingObject(t *testing.T) {
	cfg := helper.GetConfig("global")

	testCases := []map[string]string{
		{
			"name":    "preexisting_remote_object_conflict",
			"purpose": "create the remote cluster outside Terraform first, then verify Terraform create reports the existing cluster conflict",
		},
	}
	for _, tc := range testCases {
		t.Logf("coverage: %s - %s", tc["name"], tc["purpose"])
	}

	x := &map[string]string{
		"RscType":  "nd_multi_cluster_connectivity",
		"RscName":  "rsc_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	var rsc *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel
	createdOutsideTerraform := false
	defer func() {
		if createdOutsideTerraform && rsc != nil {
			deleteMultiClusterConnectivityIfExistsOutsideTerraform(t, rsc)
		}
	}()

	s1 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create the object outside Terraform, then apply the
			// matching Terraform config. A normal resource create cannot adopt
			// the pre-existing connected cluster, so the API should return the
			// existing cluster-name conflict.
			{
				Config: func() string {
					*stepCount++
					s1.Index = *stepCount
					s1.Name = fmt.Sprintf("%s_%d_preexisting_remote_object_conflict", t.Name(), s1.Index)

					overrides := map[string]interface{}{
						"login_domain": multiClusterConnectivityCreateDomain,
					}
					helper.GenerateMultiClusterConnectivityObject(&rsc,
						multiClusterConnectivityHostname,
						cfg.ND.User,
						cfg.ND.Password,
						overrides,
					)

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)

					createdOutsideTerraform = createMultiClusterConnectivityOutsideTerraform(t, rsc)
				},
				ExpectError: regexp.MustCompile(
					`cluster name \([^)]+\) conflicts with an existing cluster name`,
				),
			},
		},
	})
}

// TestAccMultiClusterConnectivityResourceCRUD exercises create/read/update
// against a remote ND cluster onboarded via nd_multi_cluster_connectivity.
//
// The resource input values mirror the recent create/update samples. The
// testbed config is still used for provider test setup and TLS behavior.
func TestAccMultiClusterConnectivityResourceCRUD(t *testing.T) {
	cfg := helper.GetConfig("global")

	testCases := []map[string]string{
		{
			"name":    "create_with_sample_payload",
			"purpose": "create with the recent sample hostname and local1 login domain",
		},
		{
			"name":    "update_sample_payload",
			"purpose": "update the nexus-dashboard cluster with the local login domain",
		},
		{
			"name":    "import_existing_cluster",
			"purpose": "verify import succeeds for an already registered cluster while ignoring attributes Nexus Dashboard does not return",
		},
		{
			"name":    "manual_delete_recreate",
			"purpose": "delete the remote cluster outside Terraform, then re-apply config to verify read removes stale state and recreates",
		},
		{
			"name":    "manual_delete_destroy",
			"purpose": "delete the remote cluster outside Terraform before destroy and verify provider delete treats 404 as success",
		},
	}
	for _, tc := range testCases {
		t.Logf("coverage: %s - %s", tc["name"], tc["purpose"])
	}

	x := &map[string]string{
		"RscType":  "nd_multi_cluster_connectivity",
		"RscName":  "rsc_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	rscAddr := "nd_multi_cluster_connectivity.rsc_test"
	tfConfig := new(string)
	stepCount := new(int)
	rsc := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with the recent sample POST payload values.
			{
				Config: func() string {
					*stepCount++
					s1.Index = *stepCount
					s1.Name = fmt.Sprintf("%s_%d_create_with_sample_payload", t.Name(), s1.Index)

					overrides := map[string]interface{}{
						"login_domain": multiClusterConnectivityCreateDomain,
					}
					helper.GenerateMultiClusterConnectivityObject(&rsc,
						multiClusterConnectivityHostname,
						cfg.ND.User,
						cfg.ND.Password,
						overrides,
					)

					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)
					deleteMultiClusterConnectivityIfExistsOutsideTerraform(t, rsc)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						MultiClusterConnectivityModelHelperStateCheck(
							rscAddr, *rsc, path.Empty(),
						),
						resource.TestCheckResourceAttrSet(rscAddr, "id"),
					)...,
				),
			},
			// Step 2: Update with the recent sample PUT payload values.
			{
				Config: func() string {
					*stepCount++
					s2.Index = *stepCount
					s2.Name = fmt.Sprintf("%s_%d_update_sample_payload", t.Name(), s2.Index)

					updates := map[string]interface{}{
						"cluster_name": multiClusterConnectivityClusterName,
						"login_domain": multiClusterConnectivityUpdateDomain,
					}
					helper.ModifyMultiClusterConnectivityObject(&rsc, updates)

					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{rsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					helper.LogStep(t, s2.Index, s2.Name, s2.Cfg)
				},
				Check: resource.ComposeTestCheckFunc(
					append(
						MultiClusterConnectivityModelHelperStateCheck(
							rscAddr, *rsc, path.Empty(),
						),
						resource.TestCheckResourceAttrSet(rscAddr, "id"),
					)...,
				),
			},
			// Step 3: Import the cluster by name. Nexus Dashboard readback does
			// not include credential/domain inputs, so import verification
			// ignores the attributes covered by the import warning.
			{
				PreConfig: func() {
					*stepCount++
					s3.Index = *stepCount
					s3.Name = fmt.Sprintf("%s_%d_import_existing_cluster", t.Name(), s3.Index)
					s3.Cfg = *tfConfig
					helper.LogStep(t, s3.Index, s3.Name, s3.Cfg)
				},
				ResourceName:                         rscAddr,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateId:                        multiClusterConnectivityClusterName,
				ImportStateVerifyIdentifierAttribute: "cluster_name",
				ImportStateVerifyIgnore: []string{
					"username",
					"password",
					"login_domain",
					"multi_cluster_login_domain",
				},
			},
			// Step 4: Delete the remote cluster outside Terraform before the
			// next apply. The provider Read should remove stale state on 404,
			// and Terraform should recreate the same configured resource.
			{
				PreConfig: func() {
					*stepCount++
					s4.Index = *stepCount
					s4.Name = fmt.Sprintf("%s_%d_manual_delete_recreate", t.Name(), s4.Index)
					s4.Cfg = *tfConfig
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
					deleteMultiClusterConnectivityOutsideTerraform(t, rsc)
				},
				Config: *tfConfig,
				Check: resource.ComposeTestCheckFunc(
					append(
						MultiClusterConnectivityModelHelperStateCheck(
							rscAddr, *rsc, path.Empty(),
						),
						resource.TestCheckResourceAttrSet(rscAddr, "id"),
					)...,
				),
			},
			// Step 5: Delete the remote cluster outside Terraform before
			// destroy. The provider Delete should treat the backend 404 as a
			// successful idempotent cleanup.
			{
				PreConfig: func() {
					*stepCount++
					s5.Index = *stepCount
					s5.Name = fmt.Sprintf("%s_%d_destroy_after_out_of_band_delete", t.Name(), s5.Index)
					s5.Cfg = *tfConfig
					helper.LogStep(t, s5.Index, s5.Name, s5.Cfg)
					deleteMultiClusterConnectivityOutsideTerraform(t, rsc)
				},
				Config:  *tfConfig,
				Destroy: true,
			},
		},
	})
}
