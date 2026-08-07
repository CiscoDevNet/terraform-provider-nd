
resource "nd_fabric_vxlan_ibgp" "test_resource_fabric_vxlan_ibgp_1" {
  fabric_name                            = "my_fabric_ibgp"
  license_tier                           = "premier"
  telemetry_collection                   = false
  security_domain                        = "all"
  bgp_asn                                = "55000"
  target_subnet_mask                     = 30
  anycast_gateway_mac                    = "2020.0000.00aa"
  performance_monitoring                 = false
  replication_mode                       = "multicast"
  multicast_group_subnet                 = "239.1.1.0/25"
  underlay_multicast_group_address_limit = 128
  tenant_routed_multicast                = false
  rendezvous_point_count                 = 2
  category                               = "fabric"
  location = {
    latitude  = 42.0
    longitude = 42.0
  }
  alert_suspend                                = "disabled"
  rendezvous_point_loopback_id                 = 254
  vpc_peer_link_vlan                           = "3600"
  vpc_peer_link_enable_native_vlan             = false
  vpc_peer_keep_alive_option                   = "management"
  vpc_auto_recovery_timer                      = 360
  vpc_delay_restore_timer                      = 150
  vpc_peer_link_port_channel_id                = "500"
  vpc_ipv6_neighbor_discovery_sync             = true
  advertise_physical_ip                        = false
  vpc_domain_id_range                          = "1-1000"
  bgp_loopback_id                              = 0
  nve_loopback_id                              = 1
  vrf_template                                 = "Default_VRF_Universal"
  network_template                             = "Default_Network_Universal"
  vrf_extension_template                       = "Default_VRF_Extension_Universal"
  network_extension_template                   = "Default_Network_Extension_Universal"
  l3_vni_no_vlan_default_option                = false
  site_id                                      = "65001"
  fabric_mtu                                   = 9216
  l2_host_interface_mtu                        = 9216
  tenant_dhcp                                  = true
  nxapi                                        = false
  nxapi_https_port                             = 443
  nxapi_http                                   = true
  nxapi_http_port                              = 80
  snmp_trap                                    = true
  anycast_border_gateway_advertise_physical_ip = false
  greenfield_debug_flag                        = "disable"
  tcam_allocation                              = true
  real_time_interface_statistics_collection    = false
  bgp_loopback_ip_range                        = "10.2.0.0/22"
  nve_loopback_ip_range                        = "10.3.0.0/22"
  anycast_rendezvous_point_ip_range            = "10.254.254.0/24"
  intra_fabric_subnet_range                    = "10.4.0.0/16"
  l2_vni_range                                 = "30000-49000"
  l3_vni_range                                 = "50000-59000"
  network_vlan_range                           = "2300-2999"
  vrf_vlan_range                               = "2000-2299"
  sub_interface_dot1q_range                    = "2-511"
  vrf_lite_auto_config                         = "manual"
  vrf_lite_subnet_range                        = "10.33.0.0/16"
  vrf_lite_subnet_target_mask                  = 30
  vrf_lite_ipv6_subnet_range                   = "2001::10.33.0.0/16"
  vrf_lite_ipv6_subnet_target_mask             = 126
  auto_unique_vrf_lite_ip_prefix               = false
  per_vrf_loopback_auto_provision              = false
  per_vrf_loopback_auto_provision_ipv6         = false
  banner                                       = "@my_ibgp_banner@"
  day0_bootstrap                               = false
  local_dhcp_server                            = false
  bootstrap_subnet_collection = [
    {
      start_ip        = "10.6.0.2"
      end_ip          = "10.6.0.9"
      default_gateway = "10.6.0.1"
      subnet_prefix   = 24
    }
  ]
  real_time_backup                            = true
  scheduled_backup                            = true
  scheduled_backup_time                       = "22:00"
  underlay_ipv6                               = false
  tenant_routed_multicast_ipv6                = false
  vrf_route_import_id_reallocation            = false
  rendezvous_point_mode                       = "asm"
  auto_generate_multicast_group_address       = false
  advertise_physical_ip_on_border             = true
  fabric_vpc_domain_id                        = false
  vpc_layer3_peer_router                      = true
  fabric_vpc_qos                              = false
  fabric_vpc_qos_policy_name                  = "spine_qos_for_fabric_vpc_peering"
  enable_peer_switch                          = false
  bgp_authentication                          = false
  bgp_authentication_key_type                 = "3des"
  pim_hello_authentication                    = false
  bfd                                         = false
  bfd_ibgp                                    = false
  bfd_authentication                          = false
  macsec                                      = false
  overlay_mode                                = "cli"
  private_vlan                                = false
  power_redundancy_mode                       = "redundant"
  copp_policy                                 = "strict"
  nve_hold_down_timer                         = 180
  cdp                                         = false
  next_generation_oam                         = true
  ngoam_south_bound_loop_detect               = false
  strict_config_compliance_mode               = false
  advanced_ssh_option                         = false
  ptp                                         = false
  default_queuing_policy                      = false
  default_queuing_policy_cloudscale           = "queuing_policy_default_8q_cloudscale"
  default_queuing_policy_r_series             = "queuing_policy_default_r_series"
  default_queuing_policy_other                = "queuing_policy_default_other"
  aiml_qos                                    = false
  aiml_qos_policy                             = "400G"
  dlb                                         = false
  ai_load_sharing                             = false
  static_underlay_ip_allocation               = false
  extra_config_aaa                            = "radius-server host 10.1.1.1"
  aaa                                         = false
  ipv6_link_local                             = true
  fabric_interface_type                       = "p2p"
  ipv6_subnet_target_mask                     = 126
  link_state_routing_protocol                 = "ospf"
  route_reflector_count                       = 2
  vpc_tor_delay_restore_timer                 = 30
  leaf_tor_id_range                           = false
  link_state_routing_tag                      = "UNDERLAY"
  ospf_area_id                                = "0.0.0.0"
  ospf_authentication                         = false
  isis_level                                  = "level-2"
  isis_area_number                            = "0001"
  isis_authentication                         = false
  bfd_ospf                                    = false
  bfd_isis                                    = false
  bfd_pim                                     = false
  auto_bgp_neighbor_description               = true
  security_group_tag                          = false
  vrf_lite_macsec                             = false
  host_interface_admin_state                  = true
  brownfield_network_name_format              = "Auto_Net_VNI$$VNI$$_VLAN$$VLAN_ID$$"
  brownfield_skip_overlay_network_attachments = false
  policy_based_routing                        = false
  mpls_handoff                                = false
  mpls_isis_area_number                       = "0001"
  stp_root_option                             = "unmanaged"
  allow_vlan_on_leaf_tor_pairing              = "none"
  pre_interface_config_leaf                   = "speed 40000"
  pre_interface_config_spine                  = "speed 40000"
  pre_interface_config_tor                    = "speed 40000"
  extra_config_leaf                           = "no shutdown"
  extra_config_spine                          = "no shutdown"
  extra_config_tor                            = "no shutdown"
  extra_config_intra_fabric_links             = "no shutdown"
  auto_symmetric_vrf_lite                     = false
  auto_vrf_lite_default_vrf                   = false
  auto_symmetric_default_vrf                  = false
  ip_service_level_agreement_id_range         = "10000-19999"
  object_tracking_number_range                = "100-299"
  service_network_vlan_range                  = "3000-3199"
  route_map_sequence_number_range             = "1-65534"
  inband_management                           = false
  heartbeat_interval                          = 190
  allow_smart_switch_onboarding               = false
  netflow_enable                              = false
  netflow_exporter_collection = [
    {
      exporter_name         = "exporter1"
      exporter_ip           = "192.168.1.1"
      vrf                   = "default"
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
  netflow_sampler_collection = [
    {
      sampler_name  = "sampler1"
      num_samples   = 1
      sampling_rate = 1
    }
  ]
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