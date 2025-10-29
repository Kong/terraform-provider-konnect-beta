resource "konnect_event_gateway_vault" "my_eventgatewayvault" {
  provider = konnect-beta
  env = {
    config = {
      prefix = "KONG_"
    }
    created_at  = "2022-11-04T20:10:06.927Z"
    description = "...my_description..."
    id          = "0e53c9e9-18d0-4412-b441-d5f3f499a4d3"
    labels = {
      key = "value"
    }
    name       = "...my_name..."
    updated_at = "2022-11-04T20:10:06.927Z"
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  konnect = {
    created_at  = "2022-11-04T20:10:06.927Z"
    description = "...my_description..."
    id          = "95dbb34a-94a6-4f3c-aede-d5eedd4cadaf"
    labels = {
      key = "value"
    }
    name       = "...my_name..."
    updated_at = "2022-11-04T20:10:06.927Z"
  }
}