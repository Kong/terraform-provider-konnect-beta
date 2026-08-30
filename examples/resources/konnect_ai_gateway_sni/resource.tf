resource "konnect_ai_gateway_sni" "my_aigatewaysni" {
  provider = konnect-beta
  certificate  = "my-certificate"
  display_name = "My SNI"
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  hostname     = "example.org"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "my-sni"
}