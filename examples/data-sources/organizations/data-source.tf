data "meraki_organizations" "list" {
}

output "organizations" {
  value = data.meraki_organizations.list
}

// Import an org from the list -
// terraform import meraki_organization.test_org 1234567890


