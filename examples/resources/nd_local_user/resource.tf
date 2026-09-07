
resource "nd_local_user" "test_resource_local_user_1" {
  login_id                  = "local_user_123"
  user_password             = "password"
  email                     = "local_user_123@mail.com"
  first_name                = "first_name"
  last_name                 = "last_name"
  remote_id_claim           = "tf_remote_id_claim"
  remote_user_authorization = true
  tenant_domain             = "all-tenants-domain"
  security_domains = {
    "all" = {
      roles = ["approver", "designer"]
    }
  }

}