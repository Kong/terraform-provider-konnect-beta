resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider" {
  provider = konnect-beta
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  key_auth = {
    config = {
      key = jsonencode("value")
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
  }
  openid_connect = {
    config = {
      key = jsonencode("value")
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
  }
}