# Import uses the cluster name of an already registered Nexus Dashboard cluster.
# Nexus Dashboard readback does not return username, password, login_domain, or
# multi_cluster_login_domain. These attributes will not be populated in
# Terraform state for the imported resource, and Terraform will emit a warning
# for this behavior during import.
#
# To update an imported cluster connection, include the username, password, and
# login_domain attributes in the Terraform configuration. These values are
# required for Terraform to re-register the imported cluster during update.
terraform import nd_multi_cluster_connectivity.test_resource_multi_cluster_connectivity_1 nd4x
