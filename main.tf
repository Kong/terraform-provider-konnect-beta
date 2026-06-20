terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
    }

    konnect-beta = {
      source  = "kong/konnect-beta"
    }
  }
}

provider "konnect" {
  konnect_access_token = "kpat_C1A8rOGSxnKw3ew4epAfoaQBxh9l3lOKVBsfKEM9gBFQhhixZ"
  server_url = "https://us.api.konghq.com"
}

provider "konnect-beta" {
  konnect_access_token = "kpat_Mpvlpu3KROGvUdwgTkx4dnwjw0ZfvM6BpiWoZOxSMJ4MfoTbA"
  server_url = "https://global.api.konghq.com"
}


################################################
# Identity Directory (Beta Resource)
################################################

resource "konnect_identity_directory" "test_directory" {
  provider = konnect-beta

  name        = "tf-trial-directory"
  description = "Test directory for Terraform provider verification 123"

  allow_all_control_planes = true
  allowed_control_planes = ["ec0e8e32-cee8-4be2-a012-4085be259d8b"]

  labels = {
    environment = "test"
    managed_by  = "terraform"
  }

  managed_by = {
    owner = "terraform"
  }
}
