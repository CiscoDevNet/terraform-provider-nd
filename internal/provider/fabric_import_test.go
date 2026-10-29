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
	"sort"
	"strings"
	"testing"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/netascode/go-nd"
)

type fabricImportTarget struct {
	name   string
	bgpAsn string
}

func testAccImportExistingFabric(t *testing.T, resourceType string, compatibleFabricTypes ...string) {
	t.Helper()

	cfg := helper.GetConfig("global")
	target := findExistingFabric(t, cfg, compatibleFabricTypes...)
	resourceAddress := resourceType + ".existing"
	t.Logf("Selected existing %s fabric %q for import", resourceType, target.name)

	providerConfig := new(string)
	helper.GetTFConfigWithSingleResource(t.Name(), map[string]string{
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}, nil, &providerConfig)
	*providerConfig += fmt.Sprintf(`
resource %q "existing" {
  fabric_name = %q
  bgp_asn     = %q
}
`, resourceType, target.name, target.bgpAsn)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:       *providerConfig,
				ResourceName: resourceAddress,
				PreConfig: func() {
					t.Logf("Importing pre-existing fabric %q into %s", target.name, resourceAddress)
				},
				ImportState:   true,
				ImportStateId: target.name,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if len(states) != 1 {
						return fmt.Errorf("expected one imported fabric state, got %d", len(states))
					}
					if got := states[0].Attributes["fabric_name"]; got != target.name {
						return fmt.Errorf("expected imported fabric_name %q, got %q", target.name, got)
					}
					if got := states[0].Attributes["bgp_asn"]; got != target.bgpAsn {
						return fmt.Errorf("expected imported bgp_asn %q, got %q", target.bgpAsn, got)
					}
					return nil
				},
			},
		},
	})
}

func findExistingFabric(t *testing.T, cfg *helper.Config, compatibleFabricTypes ...string) fabricImportTarget {
	t.Helper()

	compatibleTypes := make(map[string]struct{}, len(compatibleFabricTypes))
	for _, fabricType := range compatibleFabricTypes {
		compatibleTypes[strings.ToLower(fabricType)] = struct{}{}
	}

	client, err := nd.NewClient(
		cfg.ND.URL, "/api/v1",
		cfg.ND.User, cfg.ND.Password,
		"", cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("Failed to create ND client to list fabrics: %v", err)
	}

	fabricsAPI := api.NewFabricAPI(&client, ndapi.DefaultFabric)
	listData, err := fabricsAPI.Get()
	if err != nil {
		t.Fatalf("Failed to list existing fabrics: %v", err)
	}

	fabricNames, err := parseFabricNames(listData)
	if err != nil {
		t.Fatalf("Failed to parse fabric list response: %v", err)
	}
	sort.Strings(fabricNames)

	for _, fabricName := range fabricNames {
		fabricAPI := api.NewFabricAPI(&client, ndapi.DefaultFabric)
		fabricAPI.FabricName = fabricName
		fabricData, err := fabricAPI.Get()
		if err != nil {
			t.Logf("Skipping fabric %q because its details could not be read: %v", fabricName, err)
			continue
		}

		var fabric resource_fabric_common.NDFCFabricCommonModel
		if err := json.Unmarshal(fabricData, &fabric); err != nil {
			t.Logf("Skipping fabric %q because its details could not be parsed: %v", fabricName, err)
			continue
		}

		if _, ok := compatibleTypes[strings.ToLower(fabric.Management.FabricType)]; ok {
			if fabric.Management.BgpAsn == "" {
				t.Logf("Skipping fabric %q because it has no BGP ASN required by the resource schema", fabricName)
				continue
			}
			return fabricImportTarget{name: fabricName, bgpAsn: fabric.Management.BgpAsn}
		}
	}

	t.Fatalf("No existing fabric of type %v with a BGP ASN found to import; discovered fabrics: %v", compatibleFabricTypes, fabricNames)
	return fabricImportTarget{}
}

func parseFabricNames(data []byte) ([]string, error) {
	var response interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	nameSet := make(map[string]struct{})
	var visit func(interface{})
	visit = func(value interface{}) {
		switch typed := value.(type) {
		case map[string]interface{}:
			for key, child := range typed {
				if key == "name" || key == "fabricName" {
					if name, ok := child.(string); ok && strings.TrimSpace(name) != "" {
						nameSet[strings.TrimSpace(name)] = struct{}{}
					}
				}
				visit(child)
			}
		case []interface{}:
			for _, child := range typed {
				if name, ok := child.(string); ok && strings.TrimSpace(name) != "" {
					nameSet[strings.TrimSpace(name)] = struct{}{}
					continue
				}
				visit(child)
			}
		}
	}
	visit(response)

	if len(nameSet) == 0 {
		return nil, fmt.Errorf("fabric list response contained no fabric names")
	}

	names := make([]string, 0, len(nameSet))
	for name := range nameSet {
		names = append(names, name)
	}
	return names, nil
}
