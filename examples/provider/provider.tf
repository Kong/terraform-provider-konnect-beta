terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.4.0"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}