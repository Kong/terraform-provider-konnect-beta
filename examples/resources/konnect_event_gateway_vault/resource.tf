resource "konnect_event_gateway_vault" "my_eventgatewayvault" {
  provider = konnect-beta
  event_gateway_env_vault = {
    config = {
      prefix = "KONG_"
    }
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
  }
  event_gateway_konnect_vault = {
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
}