# Import by fabric_name. Environment variables use its uppercase value with
# hyphens replaced by underscores; for example, apic1 uses APIC1.
# Credential variables are optional and, when set, are stored in imported state.
# export APIC1_USERNAME="admin"
# export APIC1_PASSWORD="<password>"
# export APIC1_LOGIN_DOMAIN="DefaultAuth"
#
# If username or password is omitted, configure it and run terraform apply
# before destroy. For forced destroy, set this Boolean and unset it afterward:
# export APIC1_FORCE="true"
terraform import nd_fabric_aci.test_resource_fabric_aci_1 apic1
