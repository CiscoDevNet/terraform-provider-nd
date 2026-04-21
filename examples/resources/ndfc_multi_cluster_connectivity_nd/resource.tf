
resource "ndfc_multi_cluster_connectivity_nd" "test_resource_multi_cluster_connectivity_nd_1" {
  cluster_type               = "ND"
  cluster_name               = "nd4x"
  hostname                   = "10.15.1.111"
  username                   = "username"
  password                   = "password"
  login_domain               = "local"
  multi_cluster_login_domain = "multi_cluster_login_domain"
}