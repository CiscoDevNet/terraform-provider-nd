
resource "nd_policy" "test_resource_policy_1" {
  is_policy_group       = false
  deploy                = true
  entity_name           = "SWITCH"
  entity_type           = "switch"
  description           = "Enable LACP"
  template_name         = "feature_enable"
  source                = "UNDERLAY"
  priority              = 500
  serial_numbers        = ["FDO245206N5", "FDO245206N6"]
  policy_parameters     = { "featureName" : "lacp" }
  secondary_entity_name = "MyNetwork_30000"
  secondary_entity_type = "configProfile"
}