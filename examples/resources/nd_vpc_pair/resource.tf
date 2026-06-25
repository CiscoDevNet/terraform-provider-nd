
resource "nd_vpc_pair" "test_resource_vpc_pair_1" {
  fabric_name          = "my_fabric"
  switch_id_1          = "9FRGGCZT8LR"
  switch_id_2          = "96E0DGXYPV2"
  use_virtual_peerlink = false
  deploy               = true
}