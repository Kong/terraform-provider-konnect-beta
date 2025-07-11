terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.7.2"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}