resource "konnect_ai_gateway_consumer" "my_aigatewayconsumer" {
  provider = konnect-beta
  custom_id    = "dev-users"
  display_name = "Greg's Dev Consumer"
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "gregs-dev-consumer"
  policies = [
    "..."
  ]
  type = "oauth"
}