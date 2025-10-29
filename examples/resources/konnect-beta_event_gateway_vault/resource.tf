resource "konnect-beta_event_gateway_vault" "my_eventgatewayvault" {
  env = {
    config = {
      prefix = "KONG_"
    }
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  konnect = {
    description = "...my_description..."
    labels = {
      key = "value"
    }
    name = "...my_name..."
  }
}