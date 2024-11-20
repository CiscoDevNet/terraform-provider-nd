terraform {
  required_providers {
    nd = {
      source = "ciscodevnet/nd"
    }
  }
}

provider "nd" {
  username = ""
  password = ""
  url      = ""
  insecure = true
}
