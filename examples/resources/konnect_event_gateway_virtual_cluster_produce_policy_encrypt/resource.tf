resource "konnect_event_gateway_virtual_cluster_produce_policy_encrypt" "my_eventgatewayvirtualclusterproducepolicyencrypt" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    encrypt = [
      {
        key_id         = "static://static-key-named-in-source"
        part_of_record = "value"
      }
    ]
    failure_mode = "passthrough"
    key_sources = [
      {
        aws = {
          # ...
        }
      }
    ]
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "47ed5d4b-bb91-44d5-bf33-af3bff430488"
  virtual_cluster_id = "0d978e42-866e-4ee2-9e7d-5018d60a8a35"
}