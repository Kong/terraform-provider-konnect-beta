resource "konnect_identity_principal_basic_auth" "my_identityprincipalbasicauth" {
  provider = konnect-beta
  directory_id = "30f305d5-c16b-49e7-9c7c-7b672f7e6b73"
  managed_by = {
    key = "value"
  }
  principal_id = "dba80c6a-eafd-4513-95c6-b1e4e2c2095f"
  username     = "...my_username..."
}