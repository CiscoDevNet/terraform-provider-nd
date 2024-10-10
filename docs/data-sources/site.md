---
subcategory: "Sites"
layout: "nd"
page_title: "ND: nd_site"
sidebar_current: "docs-nd-data-source-nd_site"
description: |-
  Data source for Nexus Dashboard Sites
---

# nd_site #

Data source for Nexus Dashboard Sites

## API Information ##

* Site Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/3-1-1/api-reference/)
* API Endpoint: `nexus/api/sitemanagement/v4/sites`

## GUI Information ##

* Location: `Admin Console -> Manage -> Sites`

## Example Usage ##

```hcl
data "nd_site" "example" {
  name = "example"
}
```

## Schema ##

### Required ###

* `name` (name) - (String) The name of the site.

### Read-Only ###
* `id` (id) - (String) The ID of the site.
* `url` (host) - (String) The URL of the site.
* `username` (userName) - (String) The username of the site.
* `password` (password) - (String) The password of the site.
* `type` (siteType) - (String) The type of the site.
* `login_domain` (loginDomain) - (String) The login domain of the site.
* `inband_epg` (inband_epg) - (String) The In-Band Endpoint Group (EPG) used to connect ND to the site.
* `latitude` (latitude) - (String) The latitude location of the site.
* `longitude` (longitude) - (String) The longitude location of the site.
* `use_proxy` (useProxy) - (Bool) The use proxy of the site, used to route network traffic through a proxy server.
