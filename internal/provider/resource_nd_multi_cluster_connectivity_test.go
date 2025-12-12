package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Onboard ND
func TestAccResourceNdMultiClusterConnectivity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config:             testConfigResourceNdMultiClusterConnectivityCreate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "id", "nd1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "fabric_name", "nd1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_hostname", "198.18.133.203"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_type", "nd"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_password", "C1sco12345"),
				),
			},
			// Import and verify values
			{
				ResourceName:      "nd_multi_cluster_connectivity.onboard_nd",
				ImportState:       true,
				ImportStateVerify: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "id", "nd1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "fabric_name", "nd1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_hostname", "198.18.133.203"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_type", "nd"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_username", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "cluster_password", ""),
				),
			},
		},
	})
}

// Onboard APIC
func TestAccResourceApicMultiClusterConnectivity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config:             testConfigResourceApicMultiClusterConnectivityCreate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "id", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "fabric_name", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_password", "C1sco12345"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_type", "apic"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "license_tier", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "latitude", "0"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "longitude", "0"),
				),
			},
			// Import and verify values
			{
				ResourceName:      "nd_multi_cluster_connectivity.onboard_apic",
				ImportState:       true,
				ImportStateVerify: false,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "id", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "fabric_name", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_username", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_password", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_type", "apic"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "license_tier", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "latitude", "0"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "longitude", "0"),
				),
			},
			// Update
			{
				Config:             testConfigResourceApicMultiClusterConnectivityUpdate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "id", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "fabric_name", "apic1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_password", "C1sco12345"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "cluster_type", "apic"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "license_tier", "premier"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "latitude", "1.1"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "longitude", "1.2"),
				),
			},
		},
	})
}

// Validate Onboard Errors
func TestAccResourceMultiClusterConnectivityError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testConfigResourceNdMultiClusterConnectivityError,
				ExpectError: regexp.MustCompile("The telemetry_network is invalid attribute for cluster_type: nd"),
			},
			{
				Config:      testConfigResourceApicMultiClusterConnectivityError,
				ExpectError: regexp.MustCompile("The multi_cluster_login_domain is invalid attribute for cluster_type: apic"),
			},
		},
	})
}

const testConfigResourceNdMultiClusterConnectivityCreate = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name      = "nd1"
  cluster_username = "admin"
  cluster_password = "C1sco12345"
  cluster_hostname = "198.18.133.203"
  cluster_type     = "nd"
}
`
const testConfigResourceApicMultiClusterConnectivityCreate = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "C1sco12345"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
}
`

const testConfigResourceApicMultiClusterConnectivityUpdate = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "C1sco12345"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
  license_tier     = "premier"
  latitude         = 1.10
  longitude        = 1.20
}
`

const testConfigResourceNdMultiClusterConnectivityError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name      = "nd1"
  cluster_username = "admin"
  cluster_password = "C1sco12345"
  cluster_hostname = "198.18.133.203"
  cluster_type     = "nd"
  telemetry_network = "inband"
}
`

const testConfigResourceApicMultiClusterConnectivityError = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "C1sco12345"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
  license_tier     = "premier"
  latitude         = 1.10
  longitude        = 1.20
  multi_cluster_login_domain = "test"
}
`
