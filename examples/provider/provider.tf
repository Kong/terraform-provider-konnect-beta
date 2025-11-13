terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.12.1"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}