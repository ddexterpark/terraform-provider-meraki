variable "MERAKI_DASHBOARD_API_KEY" {
  default = "3b0779c586b85dca7aa810e48d8b3434d74f7e2f"
  type = string
}

variable "MERAKI_DASHBOARD_API_URL" {
  default = "https://api.meraki.com/api/"
  type = string
}


provider "meraki" {
  # example configuration here

}


terraform {
  required_providers {
    meraki = {
      source = "ddexterpark/meraki"
    }
  }
}

resource "scaffolding_example" "example" {
  configurable_attribute = "some-value"
}