resource "konnect_api_package" "my_apipackage" {
  provider = konnect-beta
  description = "...my_description..."
  name        = "MyAPIPackage"
  rate_limiting_config = {
    duration = "hours"
    limit    = 9
  }
  slug    = "my-api-package-v1"
  version = "...my_version..."
}