resource "ndfc_inventory_switch" "three_switches" {
  fabric_name           = "my_fabric"
  mode                  = "discovery"
  username              = "admin"
  password              = "mysecret"
  snmp_v3_auth_protocol = "MD5"
  max_hop               = 0
  recalculate           = true
  deploy                = true

  switches = {
    "SERIAL001" = {
      ip_address      = "10.122.84.63"
      switch_role     = "leaf"
      preserve_config = true
    }
    "SERIAL002" = {
      ip_address      = "10.122.84.71"
      switch_role     = "leaf"
      preserve_config = true
    }
    "SERIAL003" = {
      ip_address      = "10.122.84.57"
      switch_role     = "spine"
      preserve_config = true
    }
  }
}
