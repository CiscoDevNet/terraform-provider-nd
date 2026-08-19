
resource "nd_inventory_switch" "test_resource_inventory_switch_1" {
  fabric_name       = "my_fabric"
  mode              = "discovery"
  preserve_config   = true
  wait_for_discover = "10m"
  wait_for_ready    = "30m"
  switch_detail = {
    ip_address              = "10.23.244.81"
    model                   = "N9K-C93180YC-FX"
    software_version        = "10.3(3)"
    software_image          = "nxos64-cs.10.3.1.F.bin"
    switch_role             = "leaf"
    gateway_ip_mask         = "10.1.1.1/24"
    discovery_auth_protocol = "MD5"
  }
  discovery_username      = "admin"
  discovery_password      = "mysecret"
  snmp_v3_auth_protocol   = "md5"
  remote_credential_store = "local"
}