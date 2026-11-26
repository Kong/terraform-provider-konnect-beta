resource "konnect_identity_vault_provider" "my_identityvaultprovider" {
  provider = konnect-beta
  authorization_endpoint = "...my_authorization_endpoint..."
  client_id              = "...my_client_id..."
  client_secret          = "...my_client_secret..."
  credential_type        = "user"
  dcr_enabled            = true
  directory_id           = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  name                   = "...my_name..."
  pkce_enabled           = true
  refresh_buffer_seconds = 0
  registration_endpoint  = "...my_registration_endpoint..."
  scopes = [
    "..."
  ]
  template_name  = "github"
  token_endpoint = "...my_token_endpoint..."
}