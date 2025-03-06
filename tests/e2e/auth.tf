terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "999.99.9" # This ensures that we only test with the locally built provider
    }
  }
}

provider "konnect-beta" {
  server_url = "https://us.api.konghq.com"
}

