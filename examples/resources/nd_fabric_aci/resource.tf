
resource "nd_fabric_aci" "test_resource_fabric_aci_1" {
  hostname             = "1.1.1.1"
  latitude             = 37.3
  longitude            = -121.8
  username             = "admin"
  password             = "**********"
  login_domain         = "DefaultAuth"
  fabric_name          = "apic1"
  license_tier         = "premier"
  telemetry_status     = "disabled"
  orchestration_status = "disabled"
  security_domain      = "all"
  verify_ca            = false
}