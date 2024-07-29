---
subcategory: "Sites"
layout: "nd"
page_title: "ND: nd_site"
sidebar_current: "docs-nd-resource-nd_site"
description: |-
  Manages Sites for Nexus Dashboard
---

# nd_site #

Manages Sites for Nexus Dashboard

## API Information ##

* Site Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/3-1-1/api-reference/)
* API Endpoint: `nexus/api/sitemanagement/v4/sites`

## GUI Information ##

* Location: `Admin Console -> Manage -> Sites`
* [Guide](https://www.cisco.com/c/en/us/td/docs/dcn/nd/3x/articles-311/nexus-dashboard-sites-311.html#_adding_aci_sites)

## Example Usage ##

The configuration snippet below shows all possible attributes of the ND Site.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl
resource "nd_site" "example" {
  name         = "example"
  url          = "10.195.219.154"
  username     = "admin"
  password     = "password"
  type         = "aci"
  inband_epg   = "epg"
  login_domain = "local"
  latitude     = "19.36475238603211"
  longitude    = "-155.28865502961474"
}
```

All examples for the Site resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/examples/resources/nd_site) folder.

## Schema ##

### Required ###

* `name` (name) - (String) The name of the site.
* `url` (host) - (String) The URL of the site.
* `username` (userName) - (String) The username of the site.
* `password` (password) - (String) The password of the site.
* `type` (siteType) - (String) The type of the site.
  * Valid Values: `aci`, `dcnm`, `third_party`, `cloud_aci`, `dcnm_ng`, `ndfc`.

### Optional ###

* `login_domain` (loginDomain) - (String) The login domain of the site.
* `inband_epg` (inband_epg) - (String) The In-Band Endpoint Group (EPG) used to connect ND to the site.
* `latitude` (latitude) - (String) The latitude location of the site.
* `longitude` (longitude) - (String) The longitude location of the site.

### Read-Only ###

* `id` (id) - (String) The ID of the site.

## Importing

~> The environment variables `ND_SITE_USERNAME`, `ND_SITE_PASSWORD` and `ND_LOGIN_DOMAIN` must be set in order to import.

An existing Site can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its name (name), via the following command:

terraform import nd_site.example {name}

Starting in Terraform version 1.5, an existing Site can be imported using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  name = "{name}"
  to   = nd_site.example
}
```
