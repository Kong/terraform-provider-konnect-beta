resource "konnect_event_gateway_vault" "my_eventgatewayvault" {
  provider = konnect-beta
  env = {
    config = {
      prefix = "KONG_"
    }
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "env"
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
}