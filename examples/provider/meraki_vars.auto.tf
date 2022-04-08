variable "MERAKI_DASHBOARD_API_KEY" {
  type      = string
  sensitive = true
}

variable "path" {
  type    = string
  default = "/api/v1"
  validation {
    condition     = can(regex("^/api/v", var.path))
    error_message = "Valid path must contain /api/v[0-1]."
  }
}

// terraform console: "host is set to:  %{if var.host == ""}Unset%{else}var.host%{endif}"
variable "host" {
  type    = string
  default = "api.meraki.com"
}


