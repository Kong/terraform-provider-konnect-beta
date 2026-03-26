resource "konnect_auth_server" "my_authserver" {
  provider = konnect-beta
  audience      = "...my_audience..."
  description   = "...my_description..."
  force_destroy = "false"
  labels = {
    key = "value"
  }
  name              = "...my_name..."
  signing_algorithm = "RS256"
  trusted_origins = [
    "https://example.com"
  ]
}