---
layout: "nd"
page_title: "Provider: ND"
sidebar_current: "docs-nd-index"
description: |-
  The Cisco ND provider is used to interact with the resources provided by Cisco Nexus Dashboard.
  The provider needs to be configured with the proper credentials before it can be used.
---

# Nexus Dashboard (ND)

Cisco Nexus Dashboard is a central management console for multiple data center sites and a common analytics solution for Cisco data center operations. Nexus Dashboard allows users to manage multiple data center sites and provide real time analytics, visibility, assurance for network policies and operations, as well as policy orchestration for the data center fabrics, such as Cisco ACI or Cisco NDFC or standalone Nexus 9000 switches.

# Cisco ND Provider

The Cisco ND terraform provider is used to interact with resources provided by Cisco Nexus Dashboard. The provider needs to be configured with proper credentials to authenticate with Cisco Nexus Dashboard.

## Authentication

Authentication with username and password.
 
Example:

```hcl
provider "nd" {
  username     = "admin"
  password     = "password"
  url          = "https://my-cisco-nd.com"
  login_domain = "DefaultAuth"
}
```

## Example Usage

```hcl
terraform {
  required_providers {
    nd = {
      source = "ciscodevnet/nd"
    }
  }
}

provider "nd" {
  username = "admin"
  password = "password"
  url      = "https://my-cisco-nd.com"
  insecure = false
}

resource "nd_site" "example" {
  name         = "example"
  username     = "admin"
  password     = "password"
  url          = "10.195.219.154"
  type         = "aci"
  inband_epg   = "test_epg"
  latitude     = "19.36475238603211"
  longitude    = "-155.28865502961474"
  login_domain = "local"
}
```

## Schema

## Required

- `username` (String) Username for the Nexus Dashboard Account.
  - Environment variable: `ND_USERNAME`
- `password` (String) Password for the Nexus Dashboard Account.
  - Environment variable: `ND_PASSWORD`
- `url` (String) URL of the Cisco Nexus Dashboard web interface.
  - Environment variable: `ND_URL`

## Optional

- `login_domain` (String) Login domain for the Nexus Dashboard Account.
  - Default: `DefaultAuth`
  - Environment variable: `ND_LOGIN_DOMAIN`
- `insecure` (Boolean) Allow insecure HTTPS client.
  - Default: `false`
  - Environment variable: `ND_INSECURE`
- `proxy_creds` (String) Proxy server credentials in the form of `username:password`.
  - Environment variable: `ND_PROXY_CREDS`
- `proxy_url` (String) Proxy Server URL with port number.
  - Environment variable: `ND_PROXY_URL`
- `retries` (Number) Number of retries for REST API calls.
  - Default: `2`
  - Environment variable: `ND_RETRIES`
