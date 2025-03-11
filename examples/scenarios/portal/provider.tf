terraform {
  required_providers {
    konnect = {
      source = "kong/konnect"
    }
    konnect-beta = {
      source = "kong/konnect-beta"
    }
  }
}

variable "pat" {
  type = string
}

provider "konnect" {
  personal_access_token = var.pat
  server_url            = "https://us.api.konghq.com"
}
provider "konnect-beta" {
  personal_access_token = var.pat
  server_url            = "https://us.api.konghq.com"
}
