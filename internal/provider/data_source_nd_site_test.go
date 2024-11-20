package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNdSite(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNdSite,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_site.example_0", "inband_epg", ""),
					resource.TestCheckResourceAttr("nd_site.example_0", "latitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_0", "login_domain", ""),
					resource.TestCheckResourceAttr("nd_site.example_0", "longitude", ""),
					resource.TestCheckResourceAttr("nd_site.example_0", "name", "example_0"),
					resource.TestCheckResourceAttr("nd_site.example_0", "password", "password"),
					resource.TestCheckResourceAttr("nd_site.example_0", "type", "aci"),
					resource.TestCheckResourceAttr("nd_site.example_0", "username", "admin"),
					resource.TestCheckResourceAttr("nd_site.example_0", "url", "10.195.219.154"),
					resource.TestCheckResourceAttr("nd_site.example_0", "use_proxy", "false"),
				),
			},
			{
				Config:      testConfigNdSiteNonExisting,
				ExpectError: regexp.MustCompile("Failed to read nd_site data source"),
			},
		},
	})
}

const testConfigNdSite = testConfigNdSiteMinDependencyForDataSource + `
data "nd_site" "example_0" {
  name = "example_0"
  depends_on = [nd_site.example_0]
}
`

const testConfigNdSiteNonExisting = `
data "nd_site" "test" {
  name = "non_existing"
}
`
