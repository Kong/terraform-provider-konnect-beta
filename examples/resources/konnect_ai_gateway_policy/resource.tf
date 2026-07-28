resource "konnect_ai_gateway_policy" "my_aigatewaypolicy" {
  provider = konnect-beta
  config = {
    key = jsonencode("value")
  }
  display_name = "My Cool AI PII Sanitizer Policy"
  enabled      = true
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  global       = false
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "ai-pii-sanitizer-1234"
  type = "ai-sanitizer"
}