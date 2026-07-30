
resource "nd_remote_storage_location" "test_resource_remote_storage_location_1" {
  name                       = "scp-server"
  description                = "Remote storage location description."
  storage_location_type      = "scp"
  hostname                   = "192.168.100.100"
  port                       = 22
  path                       = "/export/path/"
  username                   = "root"
  password                   = "password"
  ignore_host_key_validation = true
}