# Getting Started

This provider is available on the [Terraform registry](https://registry.terraform.io/providers/Kong/konnect-beta/latest).

## Sample manifest

Place the following content in `main.tf`, set your `personal_access_token` then run `terraform apply`.

```hcl
# Configure the provider to use your Kong Konnect account
terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
    }
  }
}

provider "konnect-beta" {
  personal_access_token = "kpat_YOUR_PAT"
  server_url            = "https://us.api.konghq.com"
}

# Create a new Control Plane
resource "konnect_gateway_control_plane" "tfdemo" {
  provider     = konnect-beta # This is important
  name         = "Terraform Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
  auth_type    = "pinned_client_certs"

  proxy_urls = [
    {
      host     = "example.com",
      port     = 443,
      protocol = "https"
    }
  ]
}
```