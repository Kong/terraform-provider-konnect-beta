resource "konnect_event_gateway_virtual_cluster_consume_policy_modify_headers" "my_eventgatewayvirtualclusterconsumepolicymodifyheaders" {
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
  parent_policy_id   = "54ea5e2d-b31c-4129-a8aa-a4a6d61eb00b"
  virtual_cluster_id = "44e783c8-0879-4a29-a504-6ab77bcc6b75"
}