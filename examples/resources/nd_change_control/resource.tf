
resource "nd_change_control" "test_resource_change_control_1" {
  admin_status                    = true
  orchestration                   = true
  number_of_approvers             = 3
  allow_self_approval             = true
  nd_managed_fabrics              = true
  bypass_telemetry_change_control = true
  ticket_name_prefix              = "TICKET_"
}