resource "konnect_event_gateway_virtual_cluster_produce_policy_modify_headers" "my_eventgatewayvirtualclusterproducepolicymodifyheaders" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    actions = [
      {
        set = {
          key   = "...my_key..."
          value = "...my_value..."
        }
      }
    ]
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "496d65d8-7e33-4fc3-841b-1408d82c7581"
  virtual_cluster_id = "5a6056d4-54ac-459e-8c63-2541fa01a206"
}