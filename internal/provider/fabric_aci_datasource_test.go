// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_fabric_aci"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccManageFabricAciDataSource(t *testing.T) {
	const (
		resourceName   = "nd_fabric_aci.fabric_test"
		dataSourceName = "data.nd_fabric_aci.fabric_test"
	)

	cfg := helper.GetConfig("global")
	hostname := cfg.ND.FabricAci.Hostname
	username := cfg.ND.FabricAci.Username
	password := cfg.ND.FabricAci.Password
	if hostname == "" || username == "" || password == "" {
		t.Fatal("fabric_aci hostname, username, and password must be set in the testbed configuration")
	}

	fabricName := "tf-" + acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	const (
		latitude            = 37.3
		longitude           = -121.8
		licenseTier         = "advantage"
		orchestrationStatus = "disabled"
		securityDomain      = "all"
	)
	providerConfig := map[string]string{
		"RscType":  "nd_fabric_aci",
		"RscName":  "fabric_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	fabric := new(resource_fabric_aci.NDFCFabricAciModel)
	fabricDataSource := &helper.FabricAciDataSourceTestData{
		RscName:    "fabric_test",
		FabricName: fabricName,
		DependsOn:  resourceName,
	}
	tfConfig := new(string)
	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}

	redactConfig := func(config string) string {
		config = strings.ReplaceAll(config, password, "<redacted>")
		if cfg.ND.Password != "" {
			config = strings.ReplaceAll(config, cfg.ND.Password, "<redacted>")
		}
		return config
	}

	// These fields are returned by the cluster GET API for both the resource
	// and datasource. AttrPair also checks that nullable fields are absent in
	// both states when the API omits them.
	attributes := []string{
		"fabric_name",
		"hostname",
		"last_update_message",
		"latitude",
		"license_tier",
		"longitude",
		"orchestration_status",
		"security_domain",
		"state",
	}
	matchingChecks := make([]resource.TestCheckFunc, 0, len(attributes))
	for _, attribute := range attributes {
		matchingChecks = append(matchingChecks,
			resource.TestCheckResourceAttrPair(dataSourceName, attribute, resourceName, attribute),
		)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Create the Fabric ACI cluster with populated optional attributes for datasource lookups")

					helper.GenerateFabricAciObject(
						&fabric,
						fabricName,
						hostname,
						username,
						password,
						map[string]interface{}{
							"latitude":             latitude,
							"longitude":            longitude,
							"license_tier":         licenseTier,
							"orchestration_status": orchestrationStatus,
							"security_domain":      securityDomain,
						},
					)
					helper.GetTFConfigWithSingleResource(s1.Name, providerConfig, []interface{}{fabric}, &tfConfig)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, redactConfig(s1.Cfg)) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "fabric_name", fabricName),
					resource.TestCheckResourceAttr(resourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(resourceName, "id", fabricName),
				),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf("%s - %s", t.Name(), "Read the existing Fabric ACI cluster and verify populated datasource attributes")

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						providerConfig,
						[]interface{}{fabric, fabricDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, redactConfig(s2.Cfg)) },
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "fabric_name", fabricName),
					resource.TestCheckResourceAttr(dataSourceName, "hostname", hostname),
					resource.TestCheckResourceAttr(dataSourceName, "latitude", fmt.Sprintf("%g", latitude)),
					resource.TestCheckResourceAttr(dataSourceName, "longitude", fmt.Sprintf("%g", longitude)),
					resource.TestCheckResourceAttr(dataSourceName, "license_tier", licenseTier),
					resource.TestCheckResourceAttr(dataSourceName, "orchestration_status", orchestrationStatus),
					resource.TestCheckResourceAttr(dataSourceName, "security_domain", securityDomain),
				),
			},
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf("%s - %s", t.Name(), "Match Fabric ACI datasource attributes with the configured cluster")

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						providerConfig,
						[]interface{}{fabric, fabricDataSource},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, redactConfig(s3.Cfg)) },
				Check: resource.ComposeTestCheckFunc(
					matchingChecks...,
				),
			},
		},
	})
}

func TestAccManageFabricAciDataSourceNotFound(t *testing.T) {
	cfg := helper.GetConfig("global")
	missingName := "tf-missing-" + acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	providerConfig := map[string]string{
		"RscType":  "nd_fabric_aci",
		"RscName":  "missing_fabric",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}
	missingDataSource := &helper.FabricAciDataSourceTestData{
		RscName:    "missing_fabric",
		FabricName: missingName,
	}

	tfConfig := new(string)
	s1 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf("%s - %s", t.Name(), "Reject datasource lookup for a missing Fabric ACI cluster")

					helper.GetTFConfigWithSingleResource(
						s1.Name,
						providerConfig,
						[]interface{}{missingDataSource},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					config := s1.Cfg
					if cfg.ND.Password != "" {
						config = strings.ReplaceAll(config, cfg.ND.Password, "<redacted>")
					}
					helper.LogStep(t, s1.Index, s1.Name, config)
				},
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`Could not read nd_fabric_aci with fabric_name\s+%q:\s+resource\s+not found`, missingName),
				),
			},
		},
	})
}
