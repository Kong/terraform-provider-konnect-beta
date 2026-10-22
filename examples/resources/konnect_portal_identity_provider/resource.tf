resource "konnect_portal_identity_provider" "my_portalidentityprovider" {
  provider = konnect-beta
  config = {
    oidc_identity_provider_config = {
      back_channel_logout_enabled = false
      claim_mappings = {
        email  = "email"
        groups = "groups"
        name   = "name"
      }
      client_id     = "YOUR_CLIENT_ID"
      client_secret = "YOUR_CLIENT_SECRET"
      issuer_url    = "https://konghq.okta.com/oauth2/default"
      scopes = [
        "openid",
        "email",
        "profile",
      ]
    }
  }
  enabled   = true
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  type      = "oidc"
}