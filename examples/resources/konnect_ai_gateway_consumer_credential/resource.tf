resource "konnect_ai_gateway_consumer_credential" "my_aigatewayconsumercredential" {
  provider = konnect-beta
  api_key      = "sk-387788hd3xnej"
  consumer_id  = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  display_name = "Greg's Dev Key"
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "gregs-dev-key"
  ttl  = 86400
  type = "api-key"
}