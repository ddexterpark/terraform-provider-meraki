// DISCOVER ORGANIZATIONS //

// get list of orgs
data "meraki_organizations" "list" {
}

// Create human readable output variable
output "organizations" {
  value = data.meraki_organizations.list
}

// Find id of existing org to add to state:
// terraform output -json organizations | jq

// IMPORT EXISTING ORG  //

// Declare organization resource for state file
resource "meraki_organization" "terraform_example_org" {
  name = "terraform_example_org"

  lifecycle {
    prevent_destroy = true
  }

  depends_on = [data.meraki_organizations.list]
}

// Import from provider
//  terraform import meraki_organization.terraform_example_org "1234567890"

// Create human readable output variable
output "terraform_example_org" {
  value = meraki_organization.terraform_example_org
}

// Verify org is in state file
// terraform output -json terraform_example_org | jq


// Destroy an org after creation
// terraform state rm 'meraki_organization.terraform_example_org'