resource "konnect_ai_gateway_certificate" "my_aigatewaycertificate" {
  provider = konnect-beta
  cert       = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
  cert_alt   = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  key        = "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n"
  key_alt    = "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "my-tls-cert"
}