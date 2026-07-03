
resource "nd_vpc_pair" "test_resource_vpc_pair_1" {
  switch_1_serial_number = "9FRGGCZT8LR"
  switch_2_serial_number = "96E0DGXYPV2"
  use_virtual_peerlink   = false
  deploy                 = true
}