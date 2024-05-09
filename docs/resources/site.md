---
subcategory: "Sites"
layout: "nd"
page_title: "ND: nd_site"
sidebar_current: "docs-nd-resource-nd_site"
description: |-
  Manages Sites for the Nexus Dashboard
---

# nd_site #

Manages Sites for the Nexus Dashboard

## API Information ##

* Site Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/3-1-1/api-reference/)

## GUI Information ##

* Location: `Admin Console -> Manage -> Sites`
* GUI Configuration [Steps](https://www.cisco.com/c/en/us/td/docs/dcn/nd/3x/articles-311/nexus-dashboard-sites-311.html#_adding_aci_sites)

## Example Usage ##

The configuration snippet below shows all possible attributes of the ND Site.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl
resource "nd_site" "example" {
  site_name     = "example"
  url           = "10.195.219.154"
  site_username = "admin"
  site_password = "password"
  site_type     = "aci"
  inband_epg    = "example_epg"
  login_domain  = "local"
  latitude      = "19.36475238603211"
  longitude     = "-155.28865502961474"
}
```

All examples for the Site resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/examples/resources/nd_site) folder.

## Schema ##

### Required ###

* `site_name` (name) - (String) The name of the site.
* `url` (host) - (String) The URL to reference the APICs.
* `site_username` (userName) - (String) The username for the APIC.
* `site_password` (password) - (String) The password for the APIC.
* `site_type` (siteType) - (String) The site type of the APICs.
  * Valid Values: `aci`, `dcnm`, `third_party`, `cloud_aci`, `dcnm_ng`, `ndfc`.

### Optional ###

* `login_domain` (loginDomain) - (String) The AAA login domain for the username of the APIC.
* `inband_epg` (inband_epg) - (String) The In-Band Endpoint Group (EPG) used to connect Nexus Dashboard to the fabric.
* `latitude` (latitude) - (String) The latitude of the location of the site.
* `longitude` (longitude) - (String) The longitude of the location of the site.

### Read-Only ###

* `id` (id) - (String) The ID of the site.

## Importing

An existing Site can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its name (site_name), via the following command:

```
terraform import nd_site.example {site_name}
```

Starting in Terraform version 1.5, an existing Site can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  site_name = "{site_name}"
  to = nd_site.example
}
```

Note: `ND_SITE_USERNAME`, `ND_SITE_PASSWORD` and `ND_LOGIN_DOMAIN` must be set in order to import.