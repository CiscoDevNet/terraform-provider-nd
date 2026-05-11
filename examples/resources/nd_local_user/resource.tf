
resource "nd_local_user" "test_resource_local_user_1" {
  login_id                  = "local_user_123"
  user_password             = "local_user_123"
  email                     = "local_user_123@mail.com"
  first_name                = "first_name"
  last_name                 = "last_name"
  remote_id_claim           = "tf_remote_id_claim"
  remote_user_authorization = false
  reuse_limitation          = 10
  time_interval_limitation  = 20
  tenant_domain             = "all-tenants-domain"
  security_domains = {
    "all" = {
      roles = ["approver", "designer"]
    }
  }

}