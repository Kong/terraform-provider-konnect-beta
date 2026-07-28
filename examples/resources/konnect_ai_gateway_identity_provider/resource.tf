resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider" {
  provider = konnect-beta
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  key_auth = {
    config = {
      additional_properties = "{ \"see\": \"documentation\" }"
      hide_credentials      = true
      key_in_body           = false
      key_in_header         = true
      key_in_query          = true
      key_names = [
        "..."
      ]
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
    type = "key-auth"
  }
  openid_connect = {
    config = {
      additional_properties = "{ \"see\": \"documentation\" }"
      auth_methods = [
        "kong_oauth2"
      ]
      cache_tokens_salt = "...my_cache_tokens_salt..."
      client_id = [
        "..."
      ]
      client_secret = [
        "..."
      ]
      consumer_claims = [
        [
          # ...
        ]
      ]
      consumer_groups_claim = [
        "..."
      ]
      consumer_groups_optional = false
      consumer_optional        = false
      issuer                   = "https://dev-123456.okta.com"
      scopes = [
        "..."
      ]
      ssl_verify = true
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
    type = "openid-connect"
  }
}