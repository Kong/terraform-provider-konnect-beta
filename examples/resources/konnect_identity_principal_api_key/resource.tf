resource "konnect_identity_principal_api_key" "my_identityprincipalapikey" {
  provider = konnect-beta
  directory_id = "99b4dbd8-6f80-482d-9d71-6253fa017915"
  expires_at   = "2022-11-04T20:10:06.927Z"
  managed_by = {
    key = "value"
  }
  metadata     = "{ \"see\": \"documentation\" }"
  name         = "...my_name..."
  owner_id     = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  principal_id = "a7bed842-b5bd-419a-87a7-bd011f1c692e"
  secret       = "...my_secret..."
  type         = "v1"
}