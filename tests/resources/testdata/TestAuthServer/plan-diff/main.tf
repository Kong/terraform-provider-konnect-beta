resource "konnect_auth_server" "my_authserver" {
  provider = konnect-beta

  name     = "tf-ci-testing-authserver"
  audience = "local-demo"
}

resource "konnect_auth_server_scopes" "my_scope" {
  provider            = konnect-beta
  name                = "my-scope"
  description         = "My Scope"
  include_in_metadata = true
  enabled             = true

  auth_server_id = konnect_auth_server.my_authserver.id
}

resource "konnect_auth_server_claims" "my_authserverclaims" {
  provider = konnect-beta

  enabled               = true
  include_in_all_scopes = false
  include_in_scopes = [
    konnect_auth_server_scopes.my_scope.id
  ]
  include_in_token = true
  name             = "my-claim"
  value            = "some-value"

  auth_server_id = konnect_auth_server.my_authserver.id
}

resource "konnect_auth_server_clients" "my_client" {
  provider = konnect-beta

  name             = "my-client"
  allow_all_scopes = true
  grant_types = [
    "client_credentials",
    "authorization_code",
    "implicit"
  ]

  response_types = [
    "id_token",
    "token",
    "code"
  ]

  auth_server_id = konnect_auth_server.my_authserver.id
}

output "issuer" {
  value = konnect_auth_server.my_authserver.issuer
}

output "client_id" {
  value = konnect_auth_server_clients.my_client.id
}
output "client_secret" {
  value = konnect_auth_server_clients.my_client.client_secret
}
