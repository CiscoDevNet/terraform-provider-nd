
resource "nd_fabric_aci" "test_resource_fabric_aci_1" {
  hostname                     = "198.18.133.101"
  latitude                     = 37.3
  longitude                    = -121.8
  username                     = "admin"
  password                     = "C1sco12345"
  login_domain                 = "DefaultAuth"
  fabric_name                  = "apic1"
  license_tier                 = "essentials"
  telemetry_status             = "enabled"
  telemetry_network            = "outband"
  telemetry_epg                = "uni/tn-mgmt/mgmtp-default/inb-nd-inb-mgmt"
  telemetry_streaming_protocol = "ipv4"
  orchestration_status         = "disabled"
  security_domain              = "all"
  verify_ca                    = false
}