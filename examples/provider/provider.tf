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
// terraform import meraki_organization.terraform1 "762234236932456591"
// in order to modify resource you have to declare everything known about it in state
// terraform show
resource "meraki_organization" "terraform1" {
  cloud     = "North America"
  id        = "762234236932456591"
  name      = "terraform1"
  url       = "https://n354.meraki.com/o/mtGUhcIf/manage/organization/overview"
      // api = {
  //        enabled = true
  //      }
  // licensing = {
  //        model = "co-term"
  //      }

 }


// terraform output -json terraform1 | jq
output "terraform1" {
  value = meraki_organization.terraform1
}

// Destroy an org after creation
// terraform apply -destroy -target='meraki_organization.terraform1'
