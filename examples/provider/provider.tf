terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.1.1"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}