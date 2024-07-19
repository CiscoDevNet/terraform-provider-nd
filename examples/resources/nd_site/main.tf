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
