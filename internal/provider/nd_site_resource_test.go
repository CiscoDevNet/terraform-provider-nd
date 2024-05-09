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
					resource.TestCheckResourceAttr("nd_site.example_1", "site_name", "example_1"),
					resource.TestCheckResourceAttr("nd_site.example_1", "site_password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_1", "site_type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_1", "site_username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_1", "url", "10.195.219.154"),
				),
			},
		},
	})
}

// ND Site full configuration with import test
func TestAccResourceNdSiteWithImportTest(t *testing.T) {
	t.Setenv("ND_SITE_USERNAME", "admin")
	t.Setenv("ND_SITE_PASSWORD", "password")
	t.Setenv("ND_LOGIN_DOMAIN", "local")
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
					resource.TestCheckResourceAttr("nd_site.example_2", "site_name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
				),
			},
			// Import and verify values
			{
				ResourceName:      "nd_site.example_2",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_2", "inband_epg", "test_epg"),
					resource.TestCheckResourceAttr("nd_site.example_2", "latitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "login_domain", "local"),
					resource.TestCheckResourceAttr("nd_site.example_2", "longitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
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
					resource.TestCheckResourceAttr("nd_site.example_2", "site_name", "example_2"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_2", "site_username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_2", "url", "10.195.219.155"),
				),
			},
		},
	})
}

const testConfigNdSiteMinDependencyForDataSource = `
resource "nd_site" "example_0" {
  site_name     = "example_0"
  site_username = "admin"
  site_password = "password"
  url           = "10.195.219.154"
  site_type     = "aci"
}
`

const testConfigNdSiteMinDependency = `
resource "nd_site" "example_1" {
  site_name     = "example_1"
  site_username = "admin"
  site_password = "password"
  url           = "10.195.219.154"
  site_type     = "aci"
}
`

const testConfigNdSiteMinDependencyCreate = `
resource "nd_site" "example_2" {
  site_name     = "example_2"
  site_username = "admin"
  site_password = "password"
  url           = "10.195.219.155"
  site_type     = "aci"
  inband_epg    = "test_epg"
  latitude      = ""
  longitude     = ""
  login_domain  = "local"
}
`

const testConfigNdSiteAllDependencyUpdate = `
resource "nd_site" "example_2" {
  site_name     = "example_2"
  site_username = "admin"
  site_password = "password"
  url           = "10.195.219.155"
  site_type     = "aci"
  inband_epg    = "test_epg"
  latitude      = "19.36475238603211"
  longitude     = "-155.28865502961474"
  login_domain  = "local"
}
`
