
resource "nd_fabric_external" "test_resource_fabric_external_1" {
  fabric_name                               = "my_ext_fabric"
  license_tier                              = "premier"
  telemetry_collection                      = false
  security_domain                           = "all"
  bgp_asn                                   = "65001"
  category                                  = "fabric"
  alert_suspend                             = "disabled"
  create_bgp_config                         = true
  advanced_ssh_option                       = false
  allow_same_loopback_ip_on_switches        = false
  allow_smart_switch_onboarding             = false
  connectivity_domain_name                  = "my_domain"
  day0_bootstrap                            = false
  cdp                                       = false
  copp_policy                               = "manual"
  local_dhcp_server                         = false
  extra_config_aaa                          = "radius-server host 10.1.1.1"
  extra_config_fabric                       = "no shutdown"
  inband_management                         = false
  real_time_interface_statistics_collection = false
  monitored_mode                            = false
  mpls_handoff                              = false
  netflow_enable                            = false
  nxapi                                     = false
  nxapi_https_port                          = 443
  nxapi_http                                = false
  nxapi_http_port                           = 80
  performance_monitoring                    = false
  power_redundancy_mode                     = "redundant"
  ptp                                       = false
  real_time_backup                          = true
  scheduled_backup                          = false
  snmp_trap                                 = true
  sub_interface_dot1q_range                 = "2-511"
}