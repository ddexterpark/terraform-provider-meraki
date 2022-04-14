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

// Get OrgId from above output and import any existing orgs:
// terraform import meraki_organization.terraform1 "1234567890"
// in order to modify resource you have to declare the configurable attributes and org id
// so after creation issue this command and copy the id, lic, and api fields into the resource below:
// terraform show
resource "meraki_organization" "terraform1" {
  name  = "terraform1"
  licensing = {
    model = "co-term"
  }
  api = {
    enabled = true
  }
}


// terraform output -json terraform1 | jq
output "terraform1" {
  value = meraki_organization.terraform1
}


// Destroy an org after creation
// terraform apply -destroy -target='meraki_organization.terraform1'

// remove an org from state (troubleshooting)
// terraform state rm 'meraki_organization.terraform1'