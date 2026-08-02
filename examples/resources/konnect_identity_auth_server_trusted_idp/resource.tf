resource "konnect_identity_auth_server_trusted_idp" "my_identityauthservertrustedidp" {
  provider = konnect-beta
  auth_server_id           = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  directory_id             = "1d6e153e-45ec-4b4b-836b-51b3da03a62f"
  id_jag_enabled           = false
  issuer_url               = "https://acme.okta.com/oauth2/default"
  jit_provisioning_enabled = false
  jwks_uri                 = "https://lone-discourse.net"
}