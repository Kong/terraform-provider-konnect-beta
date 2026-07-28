resource "konnect_ai_gateway_config_store" "my_aigatewayconfigstore" {
  provider = konnect-beta
  display_name = "my-config-store"
  force        = false
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  name         = "my-config-store"
}