resource "konnect_auth_server_scopes" "my_authserverscopes" {
  provider = konnect-beta
  auth_server_id      = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  default             = true
  description         = "...my_description..."
  enabled             = false
  include_in_metadata = false
  name                = "...my_name..."
}