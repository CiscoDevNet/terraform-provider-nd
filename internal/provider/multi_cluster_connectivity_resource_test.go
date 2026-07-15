// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/infra/api"
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/netascode/go-nd"
)

const (
	multiClusterConnectivityBackendAssignTimeout = 2 * time.Minute
	multiClusterConnectivityPollInterval         = 5 * time.Second
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
		t.Fatalf("failed to create ND client for multi-cluster connectivity acceptance helper: %s", err.Error())
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
		t.Fatalf("failed to list multi-cluster connectivity objects outside Terraform: %s", err.Error())
	}

	var clustersResp map[string][]resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel
	if err := json.Unmarshal(respData, &clustersResp); err != nil {
		t.Fatalf("failed to decode multi-cluster connectivity list response: %s", err.Error())
	}

	for _, cluster := range clustersResp["clusters"] {
		if cluster.Spec.Hostname == model.Spec.Hostname &&
			(cluster.Spec.ClusterType == "" || cluster.Spec.ClusterType == "ND") {
			return cluster.Spec.ClusterName
		}
	}

	return ""
}

func waitForMultiClusterConnectivityNameOutsideTerraform(t *testing.T, client *nd.Client, model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel, wantPresent bool) string {
	t.Helper()

	deadline := time.Now().Add(multiClusterConnectivityBackendAssignTimeout)
	for {
		clusterName := multiClusterConnectivityNameOutsideTerraform(t, client, model)
		if wantPresent && clusterName != "" {
			return clusterName
		}
		if !wantPresent && clusterName == "" {
			return ""
		}
		if time.Now().After(deadline) {
			if wantPresent {
				t.Fatalf("timed out waiting for multi-cluster connectivity object with hostname %q to be created", model.Spec.Hostname)
			}
			t.Fatalf("timed out waiting for multi-cluster connectivity object with hostname %q to be deleted", model.Spec.Hostname)
		}
		time.Sleep(multiClusterConnectivityPollInterval)
	}
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
		t.Fatalf("failed to delete multi-cluster connectivity %q outside Terraform: %s %s", clusterName, err.Error(), res.String())
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
		t.Fatalf("failed to marshal manual multi-cluster connectivity payload: %s", err.Error())
	}

	clusterAPI := api.NewClusterAPI(client)
	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		t.Fatalf("failed to create multi-cluster connectivity outside Terraform: %s %s", err.Error(), res.String())
	}
	return true
}

func multiClusterConnectivityImportIDFromState(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok || rs.Primary == nil {
			return "", fmt.Errorf("resource %q not found in state", resourceName)
		}

		clusterName := rs.Primary.Attributes["cluster_name"]
		if clusterName == "" {
			clusterName = rs.Primary.ID
		}
		if clusterName == "" {
			return "", fmt.Errorf("resource %q does not have a cluster_name or id in state", resourceName)
		}

		return clusterName, nil
	}
}

func missingMultiClusterConnectivityImportIDFromState(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		clusterName, err := multiClusterConnectivityImportIDFromState(resourceName)(state)
		if err != nil {
			return "", err
		}

		return clusterName + "_missing", nil
	}
}

func TestAccMultiClusterConnectivityResourceCreateIgnoresRequestedClusterName(t *testing.T) {
	cfg := helper.GetConfig("global")
	mccCfg := cfg.ND.MultiClusterConnectivity

	testCases := []map[string]string{
		{
			"name":    "ignored_create_name",
			"purpose": "verify create payload clusterName is ignored and Nexus Dashboard returns its assigned cluster name",
		},
	}
	for _, tc := range testCases {
		t.Logf("coverage: %s - %s", tc["name"], tc["purpose"])
	}

	testAccPreCheck(t, "global")
	s1 := &helper.StepInfo{
		Index: 1,
		Name:  fmt.Sprintf("%s_%d_ignored_create_name", t.Name(), 1),
		Cfg:   "raw POST /infra/clusters payload with requested clusterName",
	}
	helper.LogStep(t, s1.Index, s1.Name, s1.Cfg)

	client := newMultiClusterConnectivityTestClient(t)

	loginPayload, err := json.Marshal(map[string]string{
		"userName":   client.Usr,
		"userPasswd": client.Pwd,
		"domain":     client.Domain,
	})
	if err != nil {
		t.Fatalf("failed to marshal Nexus Dashboard login payload: %s", err.Error())
	}
	loginReq, err := http.NewRequest(http.MethodPost, client.Url+"/login", bytes.NewReader(loginPayload))
	if err != nil {
		t.Fatalf("failed to build Nexus Dashboard login request: %s", err.Error())
	}
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp, err := client.HttpClient.Do(loginReq)
	if err != nil {
		t.Fatalf("failed to login to Nexus Dashboard: %s", err.Error())
	}
	defer loginResp.Body.Close()
	loginBody, err := io.ReadAll(loginResp.Body)
	if err != nil {
		t.Fatalf("failed to read Nexus Dashboard login response: %s", err.Error())
	}
	if loginResp.StatusCode < http.StatusOK || loginResp.StatusCode >= http.StatusMultipleChoices {
		t.Fatalf("Nexus Dashboard login failed: status=%s body=%s", loginResp.Status, string(loginBody))
	}
	var loginData struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(loginBody, &loginData); err != nil {
		t.Fatalf("failed to decode Nexus Dashboard login response: %s", err.Error())
	}
	if loginData.Token == "" {
		t.Fatalf("Nexus Dashboard login response did not include a token")
	}

	rawRequest := func(method, path string, payload []byte) []byte {
		req, err := http.NewRequest(method, client.Url+client.BasePath+path, bytes.NewReader(payload))
		if err != nil {
			t.Fatalf("failed to build Nexus Dashboard %s request for %s: %s", method, path, err.Error())
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+loginData.Token)

		resp, err := client.HttpClient.Do(req)
		if err != nil {
			t.Fatalf("failed Nexus Dashboard %s request for %s: %s", method, path, err.Error())
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read Nexus Dashboard %s response for %s: %s", method, path, err.Error())
		}
		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
			t.Fatalf("Nexus Dashboard %s request for %s failed: status=%s body=%s", method, path, resp.Status, string(respBody))
		}

		return respBody
	}

	findClusterName := func(model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) string {
		if model.Spec.ClusterName != "" {
			return model.Spec.ClusterName
		}

		respData := rawRequest(http.MethodGet, api.UrlCluster, nil)
		var clustersResp map[string][]resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel
		if err := json.Unmarshal(respData, &clustersResp); err != nil {
			t.Fatalf("failed to decode multi-cluster connectivity list response: %s", err.Error())
		}

		for _, cluster := range clustersResp["clusters"] {
			if cluster.Spec.Hostname == model.Spec.Hostname &&
				(cluster.Spec.ClusterType == "" || cluster.Spec.ClusterType == "ND") {
				return cluster.Spec.ClusterName
			}
		}

		return ""
	}

	waitForClusterName := func(model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel, wantPresent bool) string {
		deadline := time.Now().Add(multiClusterConnectivityBackendAssignTimeout)
		for {
			clusterName := findClusterName(model)
			if wantPresent && clusterName != "" {
				return clusterName
			}
			if !wantPresent && clusterName == "" {
				return ""
			}
			if time.Now().After(deadline) {
				if wantPresent {
					t.Fatalf("timed out waiting for multi-cluster connectivity object with hostname %q to be created", model.Spec.Hostname)
				}
				t.Fatalf("timed out waiting for multi-cluster connectivity object with hostname %q to be deleted", model.Spec.Hostname)
			}
			time.Sleep(multiClusterConnectivityPollInterval)
		}
	}

	deleteIfExists := func(model *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel) {
		clusterName := findClusterName(model)
		if clusterName == "" {
			return
		}

		removePath := fmt.Sprintf(api.UrlClusterRemoveByName, url.PathEscape(clusterName))
		rawRequest(http.MethodPost, removePath, []byte("{}"))
	}

	requestedClusterName := "cluster_name_test"

	rsc := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)
	helper.GenerateMultiClusterConnectivityObject(&rsc,
		mccCfg.Hostname,
		cfg.ND.User,
		cfg.ND.Password,
		map[string]interface{}{
			"login_domain": mccCfg.LoginDomain,
		},
	)
	rsc.Spec.ClusterName = requestedClusterName

	lookupModel := *rsc
	lookupModel.Spec.ClusterName = ""
	defer func() {
		deleteIfExists(&lookupModel)
		waitForClusterName(&lookupModel, false)
	}()
	deleteIfExists(&lookupModel)
	waitForClusterName(&lookupModel, false)

	createPayload := *rsc
	createPayload.Spec.ClusterType = "ND"
	payload, err := json.Marshal(createPayload)
	if err != nil {
		t.Fatalf("failed to marshal multi-cluster connectivity payload with requested clusterName: %s", err.Error())
	}

	rawRequest(http.MethodPost, api.UrlCluster, payload)

	assignedClusterName := waitForClusterName(&lookupModel, true)
	if assignedClusterName == requestedClusterName {
		t.Fatalf("expected backend to ignore requested clusterName %q during create, but read back the same cluster name", requestedClusterName)
	} else {
		// Expected result
		t.Logf("backend-assigned clusterName %q differs from requested create payload clusterName %q", assignedClusterName, requestedClusterName)
	}
}

func TestAccMultiClusterConnectivityResourcePreExistingObject(t *testing.T) {
	cfg := helper.GetConfig("global")
	mccCfg := cfg.ND.MultiClusterConnectivity

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
						"login_domain": mccCfg.LoginDomain,
					}
					helper.GenerateMultiClusterConnectivityObject(&rsc,
						mccCfg.Hostname,
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
	mccCfg := cfg.ND.MultiClusterConnectivity

	testCases := []map[string]string{
		{
			"name":    "create_with_sample_payload",
			"purpose": "create with the configured sample hostname and login domain",
		},
		{
			"name":    "update_sample_payload",
			"purpose": "update the backend-assigned Nexus Dashboard cluster with the configured login domain",
		},
		{
			"name":    "import_existing_cluster",
			"purpose": "verify import succeeds for an already registered cluster while ignoring attributes Nexus Dashboard does not return",
		},
		{
			"name":    "import_missing_cluster",
			"purpose": "verify import returns a resource not found error for a missing cluster name",
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
	s6 := &helper.StepInfo{}

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
						"login_domain": mccCfg.LoginDomain,
					}
					helper.GenerateMultiClusterConnectivityObject(&rsc,
						mccCfg.Hostname,
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
						resource.TestCheckResourceAttrSet(rscAddr, "cluster_name"),
					)...,
				),
			},
			// Step 2: Update with the recent sample PUT payload values.
			{
				Config: func() string {
					*stepCount++
					s2.Index = *stepCount
					s2.Name = fmt.Sprintf("%s_%d_update_sample_payload", t.Name(), s2.Index)

					updates := map[string]interface{}{"login_domain": mccCfg.LoginDomain}
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
						resource.TestCheckResourceAttrSet(rscAddr, "cluster_name"),
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
				ImportStateIdFunc:                    multiClusterConnectivityImportIDFromState(rscAddr),
				ImportStateVerifyIdentifierAttribute: "cluster_name",
				ImportStateVerifyIgnore: []string{
					"username",
					"password",
					"login_domain",
					"multi_cluster_login_domain",
				},
			},
			// Step 4: Import a missing cluster name to verify the import path
			// reports resource not found instead of producing partial state.
			{
				PreConfig: func() {
					*stepCount++
					s4.Index = *stepCount
					s4.Name = fmt.Sprintf("%s_%d_import_missing_cluster", t.Name(), s4.Index)
					s4.Cfg = *tfConfig
					helper.LogStep(t, s4.Index, s4.Name, s4.Cfg)
				},
				ResourceName:      rscAddr,
				ImportState:       true,
				ImportStateIdFunc: missingMultiClusterConnectivityImportIDFromState(rscAddr),
				ExpectError: regexp.MustCompile(
					`Could not import nd multi cluster connectivity with id\s+"[^"]+_missing":\s+resource\s+not\s+found`,
				),
			},
			// Step 5: Delete the remote cluster outside Terraform before the
			// next apply. The provider Read should remove stale state on 404,
			// and Terraform should recreate the same configured resource.
			{
				PreConfig: func() {
					*stepCount++
					s5.Index = *stepCount
					s5.Name = fmt.Sprintf("%s_%d_manual_delete_recreate", t.Name(), s5.Index)
					s5.Cfg = *tfConfig
					helper.LogStep(t, s5.Index, s5.Name, s5.Cfg)
					deleteMultiClusterConnectivityOutsideTerraform(t, rsc)
				},
				Config: *tfConfig,
				Check: resource.ComposeTestCheckFunc(
					append(
						MultiClusterConnectivityModelHelperStateCheck(
							rscAddr, *rsc, path.Empty(),
						),
						resource.TestCheckResourceAttrSet(rscAddr, "id"),
						resource.TestCheckResourceAttrSet(rscAddr, "cluster_name"),
					)...,
				),
			},
			// Step 6: Delete the remote cluster outside Terraform before
			// destroy. The provider Delete should treat the backend 404 as a
			// successful idempotent cleanup.
			{
				PreConfig: func() {
					*stepCount++
					s6.Index = *stepCount
					s6.Name = fmt.Sprintf("%s_%d_destroy_after_out_of_band_delete", t.Name(), s6.Index)
					s6.Cfg = *tfConfig
					helper.LogStep(t, s6.Index, s6.Name, s6.Cfg)
					deleteMultiClusterConnectivityOutsideTerraform(t, rsc)
				},
				Config:  *tfConfig,
				Destroy: true,
			},
		},
	})
}
