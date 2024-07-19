---
subcategory: "Version"
layout: "nd"
page_title: "ND: nd_version"
sidebar_current: "docs-nd-data-source-nd_version"
description: |-
  Data source for Nexus Dashboard Version
---

# nd_site #

Data source for Nexus Dashboard Version

## API Information ##

* Site Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/3-1-1/api-reference/)
* API Endpoint: `nexus/api/sitemanagement/v4/sites`

## GUI Information ##

* Location: `Help -> Welcome Screen`

## Example Usage ##

```hcl
data "nd_version" "example" {
}
```

## Schema ##

### Read-Only ###

* `build_host` (build_host) - (String) The build host of the ND Platform Version.
* `build_time` (build_time) - (String) The build time of the ND Platform Version.
* `commit_id` (commit_id) - (String) The commit id of the ND Platform Version.
* `maintenance` (maintenance) - (Number) The maintenance version number of the ND Platform Version.
* `major` (major) - (Number) The major version number of the ND Platform Version.
* `minor` (minor) - (Number) The minor version number of the ND Platform Version.
* `patch` (patch) - (String) The patch version letter of the ND Platform Version.
* `product_id` (product_id) - (String) The product id of the ND Platform Version.
* `product_name` (product_name) - (String) The product name of the ND Platform Version.
* `release` (release) - (Boolean) The release status of the ND Platform Version.
* `user` (user) - (String) The build user name of the ND Platform Version.