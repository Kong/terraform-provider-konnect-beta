terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.19.0"
    }
  }
}

provider "konnect-beta" {
  server_url = "..." # Optional - can use KONNECT_SERVER_URL environment variable
}