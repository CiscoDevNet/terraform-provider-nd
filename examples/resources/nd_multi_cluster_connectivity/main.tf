resource "nd_multi_cluster_connectivity" "onboard_nd" {
  fabric_name = "nd1"
  username    = "admin"
  password    = "password"
  hostname    = "198.18.133.203"
  type        = "nd"
}

resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name  = "apic1"
  username     = "admin"
  password     = "password"
  hostname     = "198.18.133.101"
  type         = "apic"
  license_tier = "premier"
  latitude     = 1.10
  longitude    = 1.20
}
