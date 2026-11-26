resource "konnect_identity_vault_provider_credential" "my_identityvaultprovidercredential" {
  provider = konnect-beta
  directory_id = "a0910333-5df8-4d68-aef7-6e231472328b"
  priority     = 1
  provider_id  = "4dc79705-5348-49ee-8649-386f87194874"
  secret       = "...my_secret..."
  selector     = "...my_selector..."
}