resource "konnect_auth_server_clients" "my_authserverclients" {
  provider              = konnect-beta
  access_token_duration = 17749669
  allow_all_scopes      = true
  allow_scopes = [
    "247c0ad0-5487-46a8-bedf-a569c07c2442"
  ]
  auth_server_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  grant_types = [
    "authorization_code"
  ]
  id_token_duration = 11692687
  labels = {
    key = "value"
  }
  login_uri = "https://inexperienced-meander.net"
  name      = "...my_name..."
  redirect_uris = [
    "https://flashy-sauerkraut.com/"
  ]
  response_types = [
    "none"
  ]
}
