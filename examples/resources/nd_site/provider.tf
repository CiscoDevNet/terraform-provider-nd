terraform {
  required_providers {
    nd = {
      source = "hashicorp.com/edu/nd"
    }
  }
}

provider "nd" {
  username = ""
  password = ""
  url      = ""
  insecure = true
  platform = "nd"
}
