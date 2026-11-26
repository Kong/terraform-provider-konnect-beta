resource "konnect_identity_vault_key" "my_identityvaultkey" {
  provider = konnect-beta
  directory_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  name         = "...my_name..."
  public_key   = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"
}