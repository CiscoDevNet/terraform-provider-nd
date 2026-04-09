
resource "ndfc_inventory_switch" "test_resource_inventory_switch_1" {
  fabric_name             = "my_fabric"
  mode                    = "discovery"
  preserve_config         = true
  username                = "admin"
  password                = "mysecret"
  snmp_v3_auth_protocol   = "MD5"
  remote_credential_store = "local"
  max_hop                 = 2
  recalculate             = true
  deploy                  = true
}