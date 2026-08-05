resource "konnect-beta_portal_application" "my_portalapplication" {
  auth_strategy_id = "8a6c1dbf-8bf7-4797-8e1a-16660bb1a4a8"
  client_id        = "...my_client_id..."
  dcr_client_id    = "...my_dcr_client_id..."
  description      = "...my_description..."
  labels = {
    key = "value"
  }
  name = "...my_name..."
  owner = {
    id   = "137c3056-aa7f-420e-a3c3-78360e2918f6"
    type = "developer"
  }
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
}