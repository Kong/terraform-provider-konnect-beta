terraform {
  required_providers {
    konnect = {
      source = "kong/konnect"
    }
  }
}


provider "konnect-beta" {
  server_url = "https://us.api.konghq.com"
}

provider "konnect" {
  server_url = "https://us.api.konghq.com"
}
