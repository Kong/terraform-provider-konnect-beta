# terraform-provider-konnect-beta

This repository contains a **BETA** Terraform provider for [Kong Konnect](https://konghq.com/products/kong-konnect/register?utm_source=github&utm_campaign=terraform&utm_content=terraform-provider-konnect-beta).

See [terraform-provider-konnect](https://github.com/kong/terraform-provider-konnect) for the official, supported provider

## Usage

The provider can be installed from the Terraform registry. Before using the provider, you must configure a `personal_access_token`. If you are running in a non-US region, you must also set the `server_url` configuration option.

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
```

You may also configure the `provider` block using environment variables:

- `personal_access_token` = `KONNECT_TOKEN`
- `system_account_access_token` = `KONNECT_SPAT`
- `server_url` = `KONNECT_SERVER_URL`

e.g. `KONNECT_TOKEN=kpat_YOUR_PAT terraform apply`

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Placeholder for Future Speakeasy SDK Sections -->
