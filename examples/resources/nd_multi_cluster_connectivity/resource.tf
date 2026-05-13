
resource "nd_multi_cluster_connectivity" "test_resource_multi_cluster_connectivity_1" {
  cluster_type               = "ND"
  cluster_name               = "nd4x"
  hostname                   = "10.15.1.111"
  username                   = "username"
  password                   = "password"
  login_domain               = "local"
  multi_cluster_login_domain = "multi_cluster_login_domain"
}