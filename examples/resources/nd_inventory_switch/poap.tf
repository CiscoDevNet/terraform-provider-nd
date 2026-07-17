# Example: Bootstrap (POAP) mode with 4 switches
#
# Switches must be in the POAP loop and reachable from NDFC.
# The provider queries the bootstrap list, merges API-sourced hardware
# details with user config, and issues importBootstrap.

resource "nd_inventory_switch" "poap_example" {
  fabric_name     = "my_fabric"
  mode            = "bootstrap"
  preserve_config = false

  # Common credentials used after bootstrap for switch discovery
  username = "admin"
  password = "mysecret"

  recalculate = true
  deploy      = true

  switches = {
    "FDO12345678" = {
      hostname                = "leaf-1"
      ip_address              = "10.1.1.11"
      switch_role             = "leaf"
      gateway_ip_mask         = "10.1.1.1/24"
      poap_password           = "bootstrap_secret"
      discovery_auth_protocol = "md5"
      image_policy            = "nxos_10_4"
      discovery_username      = "admin"
      discovery_password      = "disc_secret"
    }

    "FDO12345679" = {
      hostname                = "leaf-2"
      ip_address              = "10.1.1.12"
      switch_role             = "leaf"
      gateway_ip_mask         = "10.1.1.1/24"
      poap_password           = "bootstrap_secret"
      discovery_auth_protocol = "md5"
      image_policy            = "nxos_10_4"
      discovery_username      = "admin"
      discovery_password      = "disc_secret"
    }

    "FDO12345680" = {
      hostname                = "spine-1"
      ip_address              = "10.1.1.13"
      switch_role             = "spine"
      gateway_ip_mask         = "10.1.1.1/24"
      poap_password           = "bootstrap_secret"
      discovery_auth_protocol = "md5"
      image_policy            = "nxos_10_4"
      discovery_username      = "admin"
      discovery_password      = "disc_secret"
    }

    "FDO12345681" = {
      hostname                = "spine-2"
      ip_address              = "10.1.1.14"
      switch_role             = "spine"
      gateway_ip_mask         = "10.1.1.1/24"
      poap_password           = "bootstrap_secret"
      discovery_auth_protocol = "md5"
      image_policy            = "nxos_10_4"
      discovery_username      = "admin"
      discovery_password      = "disc_secret"
    }
  }
}
