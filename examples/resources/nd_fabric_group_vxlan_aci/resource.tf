
resource "nd_fabric_group_vxlan_aci" "test_resource_fabric_group_vxlan_aci_1" {
  fabric_name                          = "tf-fabric-group-vxlan-aci"
  l2_vni_range                         = "30000-49000"
  l3_vni_range                         = "50000-59000"
  anycast_gateway_mac                  = "0022.bdf8.19ff"
  multisite_overlay_inter_connect_type = "directPeering"
}