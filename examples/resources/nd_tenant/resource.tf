
resource "nd_tenant" "test_resource_tenant_1" {
  name        = "tenant1"
  description = "Tenant description"
  fabric_associations = [
    {
      fabric_name   = "ansible_test"
      tenant_prefix = "tf_"
      local_name    = "tf_test_tenant1"
      allowed_vlans = ["7", "10-20", "30-40"]
    }
  ]
}