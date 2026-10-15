resource "konnect_ai_gateway_ca_certificate" "my_aigatewaycacertificate" {
  provider = konnect-beta
  cert       = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "my-root-ca"
}