
resource "nd_remote_storage_location" "test_resource_remote_storage_location_1" {
  name        = "scp-server"
  description = "Remote storage location description."
  hostname    = "192.168.100.100"
  path        = "/export/path/"
  scp_sftp = {
    protocol                   = "scp"
    port                       = 22
    username                   = "admin"
    password                   = "password"
    ignore_host_key_validation = false
    accept_host_key            = true
  }
}