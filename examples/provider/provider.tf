terraform {
  required_providers {
    meraki = {
      source = "ddexterpark/meraki"
    }
  }
}

provider "meraki" {
  # example configuration here
  apikey = var.MERAKI_DASHBOARD_API_KEY
  path   = var.path
  host   = var.host
}

// Get List of Organizations
data "meraki_organizations" "list" {
}

// terraform output -json organizations | jq
output "organizations" {
  value = data.meraki_organizations.list
}

resource "meraki_organization" "terraform1" {
  name = "terraform1"

}

// terraform output -json terraform1 | jq
output "terraform1" {
  value = meraki_organization.terraform1
}

// Destroy an org after creation
// terraform state rm 'meraki_organization.terraform1'
