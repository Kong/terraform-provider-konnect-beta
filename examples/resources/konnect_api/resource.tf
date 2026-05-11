resource "konnect_api" "my_api" {
  provider = konnect-beta
  attributes  = "{ \"see\": \"documentation\" }"
  description = "...my_description..."
  labels = {
    key = "value"
  }
  name = "MyAPI"
  slug = "my-api-v1"
  spec = {
    api_spec_content_payload = {
      content = "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"
    }
  }
  spec_content = "...my_spec_content..."
  version      = "...my_version..."
}