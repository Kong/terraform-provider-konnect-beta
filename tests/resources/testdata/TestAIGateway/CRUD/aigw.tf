resource "konnect_ai_gateway" "my_aigateway" {
  provider = konnect-beta
  display_name          = "My Test AI Gateway"

  name = "my-test-ai-gateway"
}