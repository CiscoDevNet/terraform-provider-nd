package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// ND Site min configuration without import test
func TestAccResourceNdSiteTest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default ND values
			{
				Config:             testConfigNdSiteMinDependency,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_1", "inband_epg", ""),
					resource.TestCheckResourceAttr("nd_site.example_1", "latitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_1", "login_domain", ""),
					resource.TestCheckResourceAttr("nd_site.example_1", "longitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_1", "name", "example_1"),
					resource.TestCheckResourceAttr("nd_site.example_1", "password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_1", "type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_1", "username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_1", "url", "10.195.219.154"),
					resource.TestCheckResourceAttr("nd_site.example_1", "use_proxy", "false"),
				),
			},
		},
	})
}

// ND Site full configuration with import test
func TestAccResourceNdSiteWithImportTest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default ND values
			{
				Config:             testConfigNdSiteMinDependencyCreate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_2", "inband_epg", "test_epg"),
					resource.TestCheckResourceAttr("nd_site.example_2", "latitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "login_domain", "local"),
					resource.TestCheckResourceAttr("nd_site.example_2", "longitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_2", "type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
					resource.TestCheckResourceAttr("nd_site.example_2", "use_proxy", "true"),
				),
			},
			// Import and verify values
			{
				ResourceName:      "nd_site.example_2",
				ImportState:       true,
				ImportStateVerify: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_2", "inband_epg", "test_epg"),
					resource.TestCheckResourceAttr("nd_site.example_2", "latitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "login_domain", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "longitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "password", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "username", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
					resource.TestCheckResourceAttr("nd_site.example_2", "use_proxy", "true"),
				),
			},
			// Update with full config and verify default ND values
			{
				Config:             testConfigNdSiteAllDependencyUpdate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_2", "inband_epg", "test_epg"),
					resource.TestCheckResourceAttr("nd_site.example_2", "latitude", "19.36475238603211"),
					resource.TestCheckResourceAttr("nd_site.example_2", "login_domain", "local"),
					resource.TestCheckResourceAttr("nd_site.example_2", "longitude", "-155.28865502961474"),
					resource.TestCheckResourceAttr("nd_site.example_2", "name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_2", "type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
					resource.TestCheckResourceAttr("nd_site.example_2", "use_proxy", "false"),
				),
			},
		},
	})
}

const testConfigNdSiteMinDependencyForDataSource = `
resource "nd_site" "example_0" {
  name     = "example_0"
  username = "admin"
  password = "password"
  url      = "10.195.219.154"
  type     = "aci"
}
`

const testConfigNdSiteMinDependency = `
resource "nd_site" "example_1" {
  name     = "example_1"
  username = "admin"
  password = "password"
  url      = "10.195.219.154"
  type     = "aci"
}
`

const testConfigNdSiteMinDependencyCreate = `
resource "nd_site" "example_2" {
  name         = "example_2"
  username     = "admin"
  password     = "password"
  url          = "10.195.219.155"
  type         = "aci"
  inband_epg   = "test_epg"
  latitude     = ""
  longitude    = ""
  login_domain = "local"
  use_proxy    = true
}
`

const testConfigNdSiteAllDependencyUpdate = `
resource "nd_site" "example_2" {
  name         = "example_2"
  username     = "admin"
  password     = "password"
  url          = "10.195.219.155"
  type         = "aci"
  inband_epg   = "test_epg"
  latitude     = "19.36475238603211"
  longitude    = "-155.28865502961474"
  login_domain = "local"
  use_proxy    = false
}
`
