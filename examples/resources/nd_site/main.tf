resource "nd_site" "example_site" {
  site_name     = "example_site"
  site_username = "admin"
  site_password = "password"
  url           = "https://example_site.com"
  site_type     = "aci"
  inband_epg    = "test_epg"
  latitude      = "19.36475238603211"
  longitude     = "-155.28865502961474"
  login_domain  = "local"
}
