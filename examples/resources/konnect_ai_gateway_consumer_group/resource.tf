resource "konnect_ai_gateway_consumer_group" "my_aigatewayconsumergroup" {
  provider = konnect-beta
  additional_properties = "{ \"see\": \"documentation\" }"
  display_name          = "Dev Users Group"
  gateway_id            = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "dev-users"
  policies = [
    "..."
  ]
}