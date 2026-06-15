
resource "nd_tenant_domain" "test_resource_tenant_domain_1" {
  name         = "tenant_domain1"
  description  = "Tenant domain description"
  tenant_names = ["tenant1", "tenant2"]
}