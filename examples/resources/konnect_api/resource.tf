resource "konnect_api" "my_api" {
  provider = konnect-beta
  attributes  = "{ \"see\": \"documentation\" }"
  description = "...my_description..."
  labels = {
    key = "value"
  }
  name         = "MyAPI"
  slug         = "my-api-v1"
  spec_content = "...my_spec_content..."
  version      = "...my_version..."
}