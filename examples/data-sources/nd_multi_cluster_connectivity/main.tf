data "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name = "apic1"
}

data "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
}
