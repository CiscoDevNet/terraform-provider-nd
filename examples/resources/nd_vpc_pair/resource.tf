terraform {
  required_providers {
    nd = {
      source = "CiscoDevNet/nd"
    }
  }
}


provider "nd" {
  username = "admin"
  password = "Bgl12@123"
  url      = "https://10.104.251.111"
  insecure = true
}

resource "nd_vpc_pair" "test_resource_vpc_pair_1" {
  fabric_name          = "new_nd4_fabric_vxlan"
  peer_switch_id       = "9FRGGCZT8LR"
  switch_id            = "96E0DGXYPV2"
  use_virtual_peerlink = false
  deploy               = true
}