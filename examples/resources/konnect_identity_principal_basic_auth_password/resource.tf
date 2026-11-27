resource "konnect_identity_principal_basic_auth_password" "my_identityprincipalbasicauthpassword" {
  provider = konnect-beta
  basic_auth_id = "d41f4f7a-b826-449f-833c-9e1c1604b8d3"
  directory_id  = "862d4d9d-af24-4ca2-ade8-ff5f18badf4d"
  expires_at    = "2022-11-04T20:10:06.927Z"
  managed_by = {
    key = "value"
  }
  principal_id = "ad5fb41f-04bc-44ce-8d4e-a29b3821b635"
  secret       = "...my_secret..."
}