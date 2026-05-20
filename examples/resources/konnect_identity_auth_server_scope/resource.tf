resource "konnect_identity_auth_server_scope" "my_identityauthserverscope" {
  provider = konnect-beta
  auth_server_id      = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  default             = false
  description         = "...my_description..."
  enabled             = true
  include_in_metadata = false
  name                = "...my_name..."
}