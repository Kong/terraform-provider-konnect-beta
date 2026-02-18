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
<!-- Start Summary [summary] -->
## Summary

Konnect API (BETA): This is a BETA specification. Endpoints in this specification may change with zero notice
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [terraform-provider-konnect-beta](#terraform-provider-konnect-beta)
  * [Usage](#usage)
  * [Installation](#installation)
  * [Authentication](#authentication)
  * [Testing the provider locally](#testing-the-provider-locally)

<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    konnect-beta = {
      source  = "kong/konnect-beta"
      version = "0.15.0"
    }
  }
}

provider "konnect-beta" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Authentication [security] -->
## Authentication

This provider supports authentication configuration via environment variables and provider configuration.

The configuration precedence is:

- Provider configuration
- Environment variables

Available configuration:

| Provider Attribute | Description |
|---|---|
| `konnect_access_token` | The Konnect access token is meant to be used by the Konnect dashboard and the decK CLI authenticate with.
. |
| `personal_access_token` | The personal access token is meant to be used as an alternative to basic-auth when accessing Konnect via APIs.
You can generate a Personal Access Token (PAT) from the [personal access token page](https://cloud.konghq.com/global/account/tokens/) in the Konnect dashboard.
The PAT token must be passed in the header of a request, for example:
`curl -X GET 'https://global.api.konghq.com/v2/users/' --header 'Authorization: Bearer kpat_xgfT...'`
. Configurable via environment variable `KONNECT_TOKEN`. |
| `system_account_access_token` | The system account access token is meant for automations and integrations that are not directly associated with a human identity.
You can generate a system account Access Token by creating a system account and then obtaining a system account access token for that account.
The access token must be passed in the header of a request, for example:
`curl -X GET 'https://global.api.konghq.com/v2/users/' --header 'Authorization: Bearer spat_i2Ej...'`
. Configurable via environment variable `KONNECT_SPAT`. |
<!-- End Authentication [security] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-konnect-beta`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/kong/konnect-beta" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
