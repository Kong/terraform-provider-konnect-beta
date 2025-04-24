terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.1.4"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}