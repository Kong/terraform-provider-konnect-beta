resource "konnect_identity_auth_server" "my_identityauthserver" {
  provider = konnect-beta
  audience                          = "...my_audience..."
  dcr_default_access_token_duration = 300
  description                       = "...my_description..."
  force_destroy                     = "false"
  labels = {
    key = "value"
  }
  name              = "...my_name..."
  signing_algorithm = "RS256"
  trusted_origins = [
    "https://example.com"
  ]
}