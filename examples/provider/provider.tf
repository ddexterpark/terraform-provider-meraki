terraform {
  required_providers {
    meraki = {
      source = "ddexterpark/meraki"
    }
  }
}

provider "meraki" {
  # example configuration here
  apikey = ""
  path = "/api/v1"
  host     = "api.meraki.com"
}


resource "meraki_organization" "test" {
  id = ""
  name = "Dexter Park"
  url = "https://n168.meraki.com/o/iZWFOa/manage/organization/overview"
  cloud = "North America"
  licensing = {
    model = "per-device"
  }
  api = {
      enabled = true
    }
}

output "my_org" {
  value = meraki_organization.test
}

data "meraki_organizations" "example" {
}