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
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "hostname", "198.18.133.203"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "type", "nd"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "password", "C1sco12345"),
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
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "hostname", "198.18.133.203"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "type", "nd"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "username", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_nd", "password", ""),
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
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "password", "C1sco12345"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "type", "apic"),
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
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "username", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "password", ""),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "type", "apic"),
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
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "username", "admin"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "password", "C1sco12345"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "hostname", "198.18.133.101"),
					resource.TestCheckResourceAttr("nd_multi_cluster_connectivity.onboard_apic", "type", "apic"),
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
				Config:      testConfigResourceTypeNdWithLicenseTierError,
				ExpectError: regexp.MustCompile("The 'license_tier' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithFeaturesError,
				ExpectError: regexp.MustCompile("The 'features' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithInbandEpgError,
				ExpectError: regexp.MustCompile("The 'inband_epg' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithSecurityDomainError,
				ExpectError: regexp.MustCompile("The 'security_domain' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithValidatePeerCertificateError,
				ExpectError: regexp.MustCompile("The 'validate_peer_certificate' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithTelemetryStreamingProtocolError,
				ExpectError: regexp.MustCompile("The 'telemetry_streaming_protocol' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceTypeNdWithTelemetryNetworkError,
				ExpectError: regexp.MustCompile("The 'telemetry_network' is invalid attribute for 'type': nd"),
			},
			{
				Config:      testConfigResourceApicWithClusterLoginDomainError,
				ExpectError: regexp.MustCompile("The 'login_domain' is invalid attribute for 'type': apic"),
			},
			{
				Config:      testConfigResourceApicWithMultiClusterLoginDomainError,
				ExpectError: regexp.MustCompile("The 'multi_cluster_login_domain' is invalid attribute for 'type': apic"),
			},
		},
	})
}

const testConfigResourceNdMultiClusterConnectivityCreate = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
  username    = "admin"
  password    = "C1sco12345"
  hostname    = "198.18.133.203"
  type        = "nd"
}
`

const testConfigResourceApicMultiClusterConnectivityCreate = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name = "apic1"
  username    = "admin"
  password    = "C1sco12345"
  hostname    = "198.18.133.101"
  type        = "apic"
}
`

const testConfigResourceApicMultiClusterConnectivityUpdate = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name  = "apic1"
  username     = "admin"
  password     = "C1sco12345"
  hostname     = "198.18.133.101"
  type         = "apic"
  license_tier = "premier"
  latitude     = 1.10
  longitude    = 1.20
}
`

const testConfigResourceTypeNdWithLicenseTierError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name  = "nd1"
  username     = "admin"
  password     = "C1sco12345"
  hostname     = "198.18.133.203"
  type         = "nd"
  license_tier = "premier"
}
`

const testConfigResourceTypeNdWithFeaturesError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
  username    = "admin"
  password    = "C1sco12345"
  hostname    = "198.18.133.203"
  type        = "nd"
  features    = ["telemetry", "orchestration"]
}
`

const testConfigResourceTypeNdWithInbandEpgError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
  username    = "admin"
  password    = "C1sco12345"
  hostname    = "198.18.133.203"
  type        = "nd"
  inband_epg  = "inband_epg"
}
`

const testConfigResourceTypeNdWithSecurityDomainError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name     = "nd1"
  username        = "admin"
  password        = "C1sco12345"
  hostname        = "198.18.133.203"
  type            = "nd"
  security_domain = "default"
}
`

const testConfigResourceTypeNdWithValidatePeerCertificateError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name               = "nd1"
  username                  = "admin"
  password                  = "C1sco12345"
  hostname                  = "198.18.133.203"
  type                      = "nd"
  validate_peer_certificate = true
}
`

const testConfigResourceTypeNdWithTelemetryStreamingProtocolError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name                  = "nd1"
  username                     = "admin"
  password                     = "C1sco12345"
  hostname                     = "198.18.133.203"
  type                         = "nd"
  telemetry_streaming_protocol = "ipv4"
}
`

const testConfigResourceTypeNdWithTelemetryNetworkError = `
resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name       = "nd1"
  username          = "admin"
  password          = "C1sco12345"
  hostname          = "198.18.133.203"
  type              = "nd"
  telemetry_network = "inband"
}
`

const testConfigResourceApicWithClusterLoginDomainError = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name  = "apic1"
  username     = "admin"
  password     = "C1sco12345"
  hostname     = "198.18.133.101"
  type         = "apic"
  license_tier = "premier"
  latitude     = 1.10
  longitude    = 1.20
  login_domain = "test"
}
`

const testConfigResourceApicWithMultiClusterLoginDomainError = `
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name                = "apic1"
  username                   = "admin"
  password                   = "C1sco12345"
  hostname                   = "198.18.133.101"
  type                       = "apic"
  license_tier               = "premier"
  latitude                   = 1.10
  longitude                  = 1.20
  multi_cluster_login_domain = "test"
}
`
