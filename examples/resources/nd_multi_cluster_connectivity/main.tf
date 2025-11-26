resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name      = "nd1"
  cluster_username = "admin"
  cluster_password = "password"
  cluster_hostname = "198.18.133.203"
  cluster_type     = "nd"
}

resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "password"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
  license_tier     = "premier"
  latitude         = 1.10
  longitude        = 1.20
}

