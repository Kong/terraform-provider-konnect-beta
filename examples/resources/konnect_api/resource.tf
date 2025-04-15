resource "konnect_api" "my_api" {
  provider = konnect-beta
  provider = konnect-beta
  provider = konnect-beta
  provider = konnect-beta
  provider = konnect-beta
  deprecated  = false
  description = "...my_description..."
  labels = {
    key = "value"
  }
  name = "MyAPI"
  public_labels = {
    key = "value"
  }
  slug    = "my-api-v1"
  version = "...my_version..."
}