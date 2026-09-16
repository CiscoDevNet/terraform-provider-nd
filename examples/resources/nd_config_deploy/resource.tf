
resource "nd_config_deploy" "test_resource_config_deploy_1" {
  fabric_name                       = "my_fabric"
  deploy                            = true
  config_save                       = true
  switch_ids                        = ["ALL"]
  force_show_run                    = false
  include_all_fabric_group_switches = false
  always_deploy                     = false
  ticket_id                         = "Ticket_001"
}