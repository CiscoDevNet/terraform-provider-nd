---
subcategory: "Sites"
layout: "nd"
page_title: "ND: nd_site"
sidebar_current: "docs-nd-data-source-nd_site"
description: |-
  Data source for the Nexus Dashboard Sites
---

# nd_site #

Data source for the Nexus Dashboard Sites

## API Information ##

* Site Management [API Information](https://developer.cisco.com/docs/nexus-dashboard/3-1-1/api-reference/)

## GUI Information ##

* Location: `Admin Console -> Manage -> Sites`
* GUI Configuration [Steps](https://www.cisco.com/c/en/us/td/docs/dcn/nd/3x/articles-311/nexus-dashboard-sites-311.html#_adding_aci_sites)

## Example Usage ##

```hcl
data "nd_site" "example" {
  site_name = "example"
}
```

## Schema ##

### Required ###

* `site_name` (name) - (String) The name of the site.

### Read-Only ###
* `id` (id) - (String) The ID of the site.
* `url` (host) - (String) The URL to reference the APICs.
* `site_username` (userName) - (String) The username for the APIC.
* `site_password` (password) - (String) The password for the APIC.
* `site_type` (siteType) - (String) The site type of the APICs.
* `login_domain` (loginDomain) - (String) The AAA login domain for the username of the APIC.
* `inband_epg` (inband_epg) - (String) The In-Band Endpoint Group (EPG) used to connect Nexus Dashboard to the fabric.
* `latitude` (latitude) - (String) The latitude of the location of the site.
* `longitude` (longitude) - (String) The longitude of the location of the site.
