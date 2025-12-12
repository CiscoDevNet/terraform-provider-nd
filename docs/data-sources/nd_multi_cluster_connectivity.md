---
subcategory: "Multi-cluster connectivity"
layout: "nd"
page_title: "ND: nd_multi_cluster_connectivity"
sidebar_current: "docs-nd-data-source-nd_multi_cluster_connectivity"
description: |-
  Data source for Nexus Dashboard Multi-cluster connectivity
---

# nd_multi_cluster_connectivity #

Data source for Nexus Dashboard Multi-cluster connectivity

## API Information ##

* Multi-cluster connectivity Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/4-1-1/api-reference/)
* API Endpoint: `/api/v1/infra/clusters`

## GUI Information ##

* Location: `Admin -> System Settings -> Multi-cluster connectivity`

## Example Usage ##

```hcl
data "nd_multi_cluster_connectivity" "example" {
  fabric_name = "example"
}
```

## Schema ##

### Required ###

* `fabric_name` (name) - (String) The name of the cluster.

### Read-Only ###
* `id` (id) - (String) The ID of the cluster.
* `cluster_type` (clusterType) - (String) The type of the cluster.
* `cluster_hostname` (onboardUrl) - (String) The URL or Hostname of the cluster.
* `cluster_username` (user) - (String) The username of the cluster.
* `cluster_password` (password) - (String) The password of the cluster.
* `latitude` (latitude) - (Float) The latitude coordinate of the cluster.
* `longitude` (longitude) - (Float) The longitude coordinate of the cluster.
* `cluster_login_domain` (loginDomain) - (String) The login domain of the cluster.
* `multi_cluster_login_domain` (multiClusterLoginDomainName) - (String) The multi cluster login domain of the cluster.
* `license_tier` (licenseTier) - (String) The license tier of the cluster.
* `features` (orchestration,telemetry) - (List) The features of the cluster.
* `inband_epg` (epg) - (String) The Inband EPG name of the cluster.
* `security_domain` (securityDomain) - (String) The security domain of the cluster.
* `validate_peer_certificate` (verifyCA) - (Bool) The validate peer certificate flag of the cluster.
* `telemetry_streaming_protocol` (useProxy) - (String) The telemetry streaming protocol of the cluster.
* `telemetry_network` (network) - (String) The telemetry network type of the cluster.