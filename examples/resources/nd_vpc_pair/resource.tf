
resource "nd_vpc_pair" "test_resource_vpc_pair_1" {
  switch_id_1          = "9B32CIGO4KA"
  switch_id_2          = "9CGCIOXRFUW"
  use_virtual_peerlink = false
  vpc_pair_details = {
    template_type                   = "default"
    admin_state                     = true
    allowed_vlans                   = "all"
    domain_id                       = 5
    enable_mirror_config            = false
    fabric_path_switch_id           = 100
    is_vpc_plus                     = false
    is_vteps                        = true
    keep_alive_hold_timeout         = 3
    keep_alive_vrf                  = "management"
    loopback_secondary_ip           = "10.3.0.2"
    nve_interface                   = 1
    peer_switch_keep_alive_local_ip = "192.168.10.102"
    peer_switch_member_interfaces   = ["e1/2"]
    peer_switch_native_vlan         = 3600
    peer_switch_po_description      = "vpc-peer-link leaf1--leaf2"
    peer_switch_po_id               = 500
    peer_switch_primary_ip          = "10.3.0.4"
    peer_switch_source_loopback     = 1
    po_mode                         = "active"
    switch_keep_alive_local_ip      = "192.168.10.101"
    switch_member_interfaces        = ["e1/2"]
    switch_native_vlan              = 3600
    switch_po_description           = "vpc-peer-link leaf1--leaf2"
    switch_po_id                    = 500
    switch_primary_ip               = "10.3.0.3"
    switch_source_loopback          = 1
  }
  deploy = true
}