resource "konnect-beta_auth_server_claims" "my_authserverclaims" {
  auth_server_id        = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  enabled               = false
  include_in_all_scopes = false
  include_in_scopes = [
    "ee406f81-4a0f-4902-bca4-091a176f4915"
  ]
  include_in_token = false
  name             = "...my_name..."
  value            = "...my_value..."
}