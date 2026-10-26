
# Format of ID used for import:
# The login ID of the local user: local_user_123
# 
# Nexus Dashboard does not return user_password during import, so Terraform
# will show a plan difference for the configured user_password after import.
terraform import nd_local_user.test_resource_local_user local_user_123

