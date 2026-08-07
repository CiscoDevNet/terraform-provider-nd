
resource "nd_fabric_external" "test_resource_fabric_external_1" {
  fabric_name          = "my_ext_fabric"
  license_tier         = "premier"
  telemetry_collection = false
  security_domain      = "all"
  bgp_asn              = "65001"
  category             = "fabric"
  location = {
    latitude  = 37.77
    longitude = -122.41
  }
  alert_suspend                      = "disabled"
  create_bgp_config                  = true
  advanced_ssh_option                = false
  allow_same_loopback_ip_on_switches = false
  allow_smart_switch_onboarding      = false
  connectivity_domain_name           = "my_domain"
  day0_bootstrap                     = false
  bootstrap_subnet_collection = [
    {
      start_ip        = "192.168.1.10"
      end_ip          = "192.168.1.20"
      default_gateway = "192.168.1.1"
      subnet_prefix   = 24
    }
  ]
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
  netflow_exporter_collection = [
    {
      exporter_name         = "exporter1"
      exporter_ip           = "192.168.1.1"
      vrf                   = "management"
      source_interface_name = "loopback0"
      udp_port              = 1
    }
  ]
  netflow_record_collection = [
    {
      record_name     = "record1"
      record_template = "netflowIpv4Record"
      layer2_record   = false
    }
  ]
  netflow_monitor_collection = [
    {
      monitor_name        = "monitor1"
      monitor_record_name = "record1"
      exporter1_name      = "exporter1"
      exporter2_name      = "exporter2"
    }
  ]
  nxapi                     = false
  nxapi_https_port          = 443
  nxapi_http                = false
  nxapi_http_port           = 80
  performance_monitoring    = false
  power_redundancy_mode     = "redundant"
  ptp                       = false
  real_time_backup          = true
  scheduled_backup          = false
  snmp_trap                 = true
  sub_interface_dot1q_range = "2-511"
  vrf_flow_rules = [
    {
      vrf_flow_rule_name    = "rule1"
      vrf_flow_rule_uuid    = "67b140d39bbd5b6f18af0af8"
      vrf_flow_rule_tenant  = "tenant1"
      vrf_flow_rule_vrf     = "default"
      vrf_flow_rule_subnets = ["example_subnets_item"]
    }
  ]
  interface_flow_rules = [
    {
      interface_flow_rule_name    = "rule1"
      interface_flow_rule_uuid    = "67b140d39bbd5b6f18af0af8"
      interface_flow_rule_type    = "physical"
      interface_flow_rule_subnets = ["example_subnets_item"]
    }
  ]
  l3_out_flow_rules = [
    {
      l3_out_flow_rule_name    = "rule1"
      l3_out_flow_rule_uuid    = "67b140d39bbd5b6f18af0af8"
      l3_out_flow_rule_type    = "subInterface"
      l3_out_flow_rule_subnets = ["example_subnets_item"]
    }
  ]
  interface_rules = [
    {
      interface_rule_name                       = "TAInterfaceRule1"
      interface_rule_enabled                    = true
      interface_rule_enable_fabric_interconnect = true
      interface_rule_enable_l3_out              = true
      interface_rule_uuid                       = "67b140d39bbd5b6f18af0af8"
      interface_rule_subnets                    = ["10.0.0.0/24"]
    }
  ]
  cost = 1.2
  email = [
    {
      name                         = "rule1"
      receiver_email               = "example@example.com"
      format                       = "html"
      start_date                   = "2023-01-01T00:00:00Z"
      collection_frequency_in_days = 42
      collection_type              = "basic"
      anomalies                    = ["critical"]
      advisories                   = ["critical"]
      risk_and_conformance_reports = ["software"]
      only_include_active_alerts   = false
    }
  ]
  message_bus = [
    {
      server                              = "stream-server-1"
      collection_type                     = "alertsAndEvents"
      collection_settings_collection_type = "basic"
      anomalies                           = ["critical"]
      advisories                          = ["critical"]
      statistics                          = ["interfaces"]
      faults                              = ["critical"]
      audit_logs                          = ["creation"]
    }
  ]
}