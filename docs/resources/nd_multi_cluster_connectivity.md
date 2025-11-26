---
subcategory: "Multi-cluster connectivity"
layout: "nd"
page_title: "ND: nd_multi_cluster_connectivity"
sidebar_current: "docs-nd-resource-nd_multi_cluster_connectivity"
description: |-
  Manages Multi-cluster connectivity for Nexus Dashboard
---

# nd_multi_cluster_connectivity #

Manages Multi-cluster connectivity for Nexus Dashboard

## API Information ##

* Multi-cluster connectivity Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/4-1-1/api-reference/)
* API Endpoint: `/api/v1/infra/clusters`

## GUI Information ##

* Location: `Admin -> System Settings -> Multi-cluster connectivity`
* [Guide](https://www.cisco.com/c/en/us/td/docs/dcn/nd/4x/release-notes/cisco-nexus-dashboard-release-notes-411.html?dtid=osscdc000283&linkclickid=srch#OrchestrationNDO)

## Example Usage ##

The configuration snippet below shows all possible attributes of the ND clusters.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl
resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "password"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
  license_tier     = "premier"
  latitude         = 1.10
  longitude        = 1.20
}

resource "nd_multi_cluster_connectivity" "onboard_apic" {
  fabric_name      = "apic1"
  cluster_username = "admin"
  cluster_password = "password"
  cluster_hostname = "198.18.133.101"
  cluster_type     = "apic"
}
```

All examples for the Multi-cluster connectivity resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/examples/resources/nd_multi_cluster_connectivity) folder.

## Schema ##

### Required ###

* `fabric_name` (name) - (String) The name of the cluster.
* `cluster_type` (clusterType) - (String) The type of the cluster.
  * Valid Values: `nd`, or `apic`.
* `cluster_hostname` (onboardUrl) - (String) The URL or Hostname of the cluster.
* `cluster_username` (user) - (String) The username of the cluster.
* `cluster_password` (password) - (String) The password of the cluster.

### Optional ###

* `cluster_login_domain` (loginDomain) - (String) The login domain of the cluster.
* `multi_cluster_login_domain` (multiClusterLoginDomainName) - (String) The multi cluster login domain of the cluster.
* `license_tier` (licenseTier) - (String) The license tier of the cluster.
  * Valid Values: `advantage`, or `essentials`, or `premier`.
* `features` (orchestration,telemetry) - (List) The features of the cluster.
  * Valid Values: `telemetry`, `orchestration`.
* `inband_epg` (epg) - (String) The Inband EPG name of the cluster.
* `security_domain` (securityDomain) - (String) The security domain of the cluster.
* `validate_peer_certificate` (verifyCA) - (Bool) The validate peer certificate flag of the cluster.
* `latitude` (latitude) - (Float) The latitude location of the cluster.
* `longitude` (longitude) - (Float) The longitude location of the cluster.
* `telemetry_streaming_protocol` (useProxy) - (String) The telemetry streaming protocol of the cluster.
  * Valid Values: `ipv4`, or `ipv6`.

### Read-Only ###

* `id` (id) - (String) The ID of the cluster.

## Importing

An existing cluster can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its name (name), via the following command:

```
terraform import nd_multi_cluster_connectivity.example {name}
```

Starting in Terraform version 1.5, an existing cluster can be imported using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  name = "{name}"
  to   = nd_multi_cluster_connectivity.example
}
```

~> The values for `cluster_username`, `cluster_password`, `cluster_login_domain` and `multi_cluster_login_domain` attributes will not be imported when the `nd_multi_cluster_connectivity` resource imports an already registered cluster from Nexus Dashboard. Modifying the `fabric_name`, `cluster_type`, `cluster_username`, `cluster_password` and `cluster_login_domain` will not update the imported cluster configuration on Nexus Dashboard. Use the `-replace` option to force the cluster recreation and use the new provided `fabric_name`, `cluster_type`, `cluster_username`, `cluster_password` and `cluster_login_domain` attributes for the imported cluster.

```
terraform apply -replace="nd_multi_cluster_connectivity.example"
```
