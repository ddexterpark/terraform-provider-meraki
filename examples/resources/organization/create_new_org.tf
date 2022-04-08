// Create single organization
resource "meraki_organization" "terraform1" {
name = "terraform1"
}

// Create 2 new orgs at once  with count feature (value starts at 0)
resource "meraki_organization" "count_org" {
count = 2
name = "count-org-${count.index}"
}


// Create 2 new orgs with list
resource "meraki_organization" "testorg" {
for_each = toset( ["LAX", "SFO"] )
name = "list-org-${each.key}"
}


// Create 2 new orgs with hash map
resource "meraki_organization" "testorg" {
for_each = {
    a = "LAX"
    b = "SFO"
}
name = "map-org-${each.key}-${each.value}}"
}

// Destroy an org after creation
// terraform state rm 'meraki_organization.'