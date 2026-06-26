terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.20.1"
    }
  }
}

provider "konnect-beta" {
  server_url = "..." # Optional - can use KONNECT_SERVER_URL environment variable
}