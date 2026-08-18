import {
  provider = konnect-beta
  to = konnect_ai_gateway_config_store_secret.my_konnect-beta_ai_gateway_config_store_secret
  id = jsonencode({
    config_store_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
    gateway_id      = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
    key             = "my-secret-key"
  })
}
