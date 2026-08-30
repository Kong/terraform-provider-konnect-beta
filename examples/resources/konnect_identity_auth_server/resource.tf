resource "konnect_identity_auth_server" "my_identityauthserver" {
  provider = konnect-beta
  audience = "...my_audience..."
  dcr = {
    default_access_token_duration = 300
    redirect_uri_allowlist = [
      "https://example.com/callback"
    ]
  }
  description   = "...my_description..."
  force_destroy = "false"
  labels = {
    key = "value"
  }
  name              = "...my_name..."
  open_dcr_enabled  = true
  signing_algorithm = "RS256"
  trusted_origins = [
    "https://example.com"
  ]
}