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
	"testing"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFabricDataSource(t *testing.T) {
	cfg := helper.GetConfig("global")
	suffix := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	fabricName := fmt.Sprintf("%s_ds_%s", cfg.ND.Fabric, suffix)
	missingFabricName := fabricName + "_missing"

	resourceName := "nd_fabric_external.fabric_test"
	dataSourceName := "data.nd_fabric.fabric_test"

	providerConfig := map[string]string{
		"RscType":  "nd_fabric_external",
		"RscName":  "fabric_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	fabric := new(resource_fabric_common.NDFCFabricCommonModel)
	helper.GenerateFabricExternalObject(&fabric, fabricName, "65001", nil)

	resourceConfig := new(string)
	helper.GetTFConfigWithSingleResource(
		"fabric_datasource_create",
		providerConfig,
		[]interface{}{helper.ExternalResource(fabric)},
		&resourceConfig,
	)

	dataSourceConfig := new(string)
	helper.GetTFConfigWithSingleResource(
		"fabric_datasource_read",
		providerConfig,
		[]interface{}{helper.ExternalResource(fabric)},
		&dataSourceConfig,
	)
	*dataSourceConfig += fabricDataSourceBlock("fabric_test", fabricName, resourceName)

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "fabric_name", resourceName, "fabric_name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "license_tier", resourceName, "license_tier"),
		resource.TestCheckResourceAttrPair(dataSourceName, "category", resourceName, "category"),
		resource.TestCheckResourceAttrPair(dataSourceName, "security_domain", resourceName, "security_domain"),
		resource.TestCheckResourceAttrPair(dataSourceName, "bgp_asn", resourceName, "bgp_asn"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: *resourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "fabric_name", fabricName),
				),
			},
			{
				Config: *dataSourceConfig,
				Check:  resource.ComposeTestCheckFunc(matchingChecks...),
			},
		},
	})

	// Verify that the data source can read each VXLAN fabric variant as well as
	// the external fabric covered above.
	vxlanFabricCases := []struct {
		name          string
		resourceType  string
		fabricType    string
		generateModel func(**resource_fabric_common.NDFCFabricCommonModel, string)
		wrapResource  func(*resource_fabric_common.NDFCFabricCommonModel) *helper.FabricTestResource
	}{
		{
			name:         "vxlan",
			resourceType: "nd_fabric_vxlan",
			fabricType:   "vxlanIbgp",
			generateModel: func(model **resource_fabric_common.NDFCFabricCommonModel, name string) {
				helper.GenerateFabricVxlanObject(model, name, "55002", "vxlanIbgp", nil)
			},
			wrapResource: helper.VxlanResource,
		},
		{
			name:         "vxlan_ebgp",
			resourceType: "nd_fabric_vxlan_ebgp",
			fabricType:   "vxlanEbgp",
			generateModel: func(model **resource_fabric_common.NDFCFabricCommonModel, name string) {
				helper.GenerateFabricEbgpObject(model, name, "55003", nil)
			},
			wrapResource: helper.EbgpResource,
		},
		{
			name:         "vxlan_ibgp",
			resourceType: "nd_fabric_vxlan_ibgp",
			fabricType:   "vxlanIbgp",
			generateModel: func(model **resource_fabric_common.NDFCFabricCommonModel, name string) {
				helper.GenerateFabricIbgpObject(model, name, "55004", nil)
			},
			wrapResource: helper.IbgpResource,
		},
	}

	for _, fabricCase := range vxlanFabricCases {
		t.Run(fabricCase.name, func(t *testing.T) {
			caseFabricName := fmt.Sprintf("%s_%s_%s", cfg.ND.Fabric, fabricCase.name, acctest.RandStringFromCharSet(5, acctest.CharSetAlpha))
			caseResourceName := "nd_fabric_vxlan.fabric_test"
			if fabricCase.resourceType == "nd_fabric_vxlan_ebgp" {
				caseResourceName = "nd_fabric_vxlan_ebgp.fabric_test"
			} else if fabricCase.resourceType == "nd_fabric_vxlan_ibgp" {
				caseResourceName = "nd_fabric_vxlan_ibgp.fabric_test"
			}
			caseDataSourceName := "data.nd_fabric.fabric_test"

			caseProviderConfig := map[string]string{
				"RscType":  fabricCase.resourceType,
				"RscName":  "fabric_test",
				"User":     cfg.ND.User,
				"Password": cfg.ND.Password,
				"Host":     cfg.ND.URL,
				"Insecure": cfg.ND.Insecure,
			}

			caseFabric := new(resource_fabric_common.NDFCFabricCommonModel)
			fabricCase.generateModel(&caseFabric, caseFabricName)

			createConfig := new(string)
			helper.GetTFConfigWithSingleResource(
				"fabric_datasource_"+fabricCase.name+"_create",
				caseProviderConfig,
				[]interface{}{fabricCase.wrapResource(caseFabric)},
				&createConfig,
			)

			readConfig := new(string)
			helper.GetTFConfigWithSingleResource(
				"fabric_datasource_"+fabricCase.name+"_read",
				caseProviderConfig,
				[]interface{}{fabricCase.wrapResource(caseFabric)},
				&readConfig,
			)
			*readConfig += fabricDataSourceBlock("fabric_test", caseFabricName, caseResourceName)

			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t, "global") },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: *createConfig,
						Check:  resource.TestCheckResourceAttr(caseResourceName, "fabric_name", caseFabricName),
					},
					{
						Config: *readConfig,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrPair(caseDataSourceName, "fabric_name", caseResourceName, "fabric_name"),
							resource.TestCheckResourceAttrPair(caseDataSourceName, "license_tier", caseResourceName, "license_tier"),
							resource.TestCheckResourceAttrPair(caseDataSourceName, "category", caseResourceName, "category"),
							resource.TestCheckResourceAttrPair(caseDataSourceName, "security_domain", caseResourceName, "security_domain"),
							resource.TestCheckResourceAttrPair(caseDataSourceName, "bgp_asn", caseResourceName, "bgp_asn"),
							resource.TestCheckResourceAttr(caseDataSourceName, "fabric_type", fabricCase.fabricType),
						),
					},
				},
			})
		})
	}

	missingConfig := new(string)
	helper.GetTFConfigWithSingleResource(
		"fabric_datasource_missing",
		providerConfig,
		nil,
		&missingConfig,
	)
	*missingConfig += fabricDataSourceBlock("missing_fabric", missingFabricName, "")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      *missingConfig,
				ExpectError: regexp.MustCompile(fmt.Sprintf(`Could not read fabric with name\s+%q:\s+resource\s+not\s+found`, missingFabricName)),
			},
		},
	})
}

func fabricDataSourceBlock(dataSourceName, fabricName, dependsOn string) string {
	block := fmt.Sprintf(`
data "nd_fabric" %q {
  fabric_name = %q
`, dataSourceName, fabricName)
	if dependsOn != "" {
		block += fmt.Sprintf("  depends_on = [%s]\n", dependsOn)
	}
	return block + "}\n"
}
