resource "konnect_identity_trusted_idp" "my_identitytrustedidp" {
  provider = konnect-beta
  directory_id = "cc79837d-110f-4821-928a-6200135b77be"
  display_name = "...my_display_name..."
  issuer_url   = "https://acme.okta.com/oauth2/default"
  jit = {
    mapping = {
      display_name = "name"
      email        = "email"
      groups       = "groups"
      identifier   = "oid"
      metadata = {
        key = "value"
      }
    }
  }
  jit_provisioning_enabled = true
  jwks_uri                 = "https://acme.okta.com/oauth2/default/v1/keys"
  vault_access_enabled     = false
}