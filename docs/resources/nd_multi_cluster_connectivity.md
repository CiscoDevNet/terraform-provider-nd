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

* `latitude` (latitude) - (Float) The latitude coordinate of the cluster.
* `longitude` (longitude) - (Float) The longitude coordinate of the cluster.
* `cluster_login_domain` (loginDomain) - (String) The login domain of the cluster. This attribute is only applicable when `cluster_type` is set to `nd`.
* `multi_cluster_login_domain` (multiClusterLoginDomainName) - (String) The multi cluster login domain of the cluster. This attribute is only applicable when `cluster_type` is set to `nd`.
* `license_tier` (licenseTier) - (String) The license tier of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
  * Valid Values: `advantage`, or `essentials`, or `premier`.
* `features` (orchestration,telemetry) - (List) The features of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
  * Valid Values: `telemetry`, `orchestration`.
* `inband_epg` (epg) - (String) The Inband EPG name of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
* `security_domain` (securityDomain) - (String) The security domain of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
* `validate_peer_certificate` (verifyCA) - (Bool) The validate peer certificate flag of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
* `telemetry_streaming_protocol` (useProxy) - (String) The telemetry streaming protocol of the cluster. This attribute is only applicable when `cluster_type` is set to `apic`.
  * Valid Values: `ipv4`, or `ipv6`.
* `telemetry_network` (network) - (String) The telemetry network type of the cluster. Allowed values are `inband`, or `outband`. This attribute is only applicable when `cluster_type` is set to `apic`.

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

~> The values for `cluster_username`, `cluster_password`, `cluster_login_domain` and `multi_cluster_login_domain` attributes will not be imported when the `nd_multi_cluster_connectivity` resource imports an already registered cluster from Nexus Dashboard. To update/delete the imported cluster, use one of the following methods:

### Option 1: Using a Cluster Credentials File

Specify the path to a JSON file containing cluster credentials using the `CLUSTER_CREDENTIALS_FILE_LOCATION` environment variable. The fabric name of the cluster will be used to match the corresponding username and password within this file.

**Example `CLUSTER_CREDENTIALS_FILE_LOCATION` file content:**

```json
{
    "nd1": {"username": "user1", "password": "password1"},
    "apic1": {"username": "user2", "password": "password2"},
    "apic2": {"username": "user3", "password": "password3"}
}
```

### Option 2: Using Direct Environment Variables

Provide the cluster's username and password directly through the `CLUSTER_USERNAME` and `CLUSTER_PASSWORD` environment variables.