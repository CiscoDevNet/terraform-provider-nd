package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNdVersion(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNdVersion,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.nd_version.test", "user"),
					resource.TestCheckResourceAttr("data.nd_version.test", "product_name", "Nexus Dashboard"),
					resource.TestCheckResourceAttr("data.nd_version.test", "product_id", "nd"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "build_host"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "build_time"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "commit_id"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "maintenance"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "major"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "minor"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "patch"),
					resource.TestCheckResourceAttrSet("data.nd_version.test", "release"),
				),
			},
		},
	})
}

const testConfigNdVersion = `
data "nd_version" "test" {
}
`
