resource "konnect_identity_auth_server_client" "my_identityauthserverclient" {
  provider = konnect-beta
  access_token_duration = 300
  allow_all_scopes      = false
  allow_scopes = [
    "5ab7e333-0b71-4f32-881a-2cea4e043539"
  ]
  auth_server_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  client_secret  = "YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"
  grant_types = [
    "refresh_token"
  ]
  id                = "kYa9iQFU5xPDSIUH9z1z"
  id_token_duration = 300
  labels = {
    key = "value"
  }
  login_uri = "https://mushy-hose.info"
  name      = "...my_name..."
  redirect_uris = [
    "https://intrepid-flint.info"
  ]
  refresh_token_duration = 2592000
  response_types = [
    "none"
  ]
  token_endpoint_auth_method = "client_secret_post"
}