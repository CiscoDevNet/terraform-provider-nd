package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNdMultiClusterConnectivity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNdMultiClusterConnectivity,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "id", "nd1"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "fabric_name", "nd1"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "cluster_hostname", "198.18.133.203"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "cluster_type", "nd"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "cluster_username", ""),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_nd", "cluster_password", ""),
				),
			},
		},
	})
}

const testConfigNdMultiClusterConnectivity = testConfigResourceNdMultiClusterConnectivityCreate + `
data "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
  depends_on  = [nd_multi_cluster_connectivity.onboard_nd]
}
`

func TestAccDataSourceApicMultiClusterConnectivity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigApicMultiClusterConnectivity,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "id", "apic1"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "fabric_name", "apic1"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "cluster_username", ""),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "cluster_password", ""),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "cluster_hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "cluster_type", "apic"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "license_tier", ""),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "latitude", "0"),
					resource.TestCheckResourceAttr("data.nd_multi_cluster_connectivity.onboard_apic", "longitude", "0"),
				),
			},
		},
	})
}

const testConfigApicMultiClusterConnectivity = testConfigResourceApicMultiClusterConnectivityCreate + `
data "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name = "apic1"
  depends_on  = [nd_multi_cluster_connectivity.onboard_apic]
}
`
