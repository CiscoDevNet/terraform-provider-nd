
resource "nd_fabric_aci" "test_resource_fabric_aci_1" {
  hostname     = "1.1.1.1"
  latitude     = 37.3
  longitude    = -121.8
  username     = "admin"
  password     = "**********"
  login_domain = "DefaultAuth"
  fabric_name  = "apic1"
  license_tier = "premier"
  telemetry = {
    network            = "inband"
    epg                = "uni/tn-tf_tenant/ap-tf_ap/epg-tf_epg"
    streaming_protocol = "ipv4"
  }
  orchestration_status = "disabled"
  security_domain      = "all"
  verify_ca            = false
}