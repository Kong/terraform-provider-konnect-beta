resource "konnect_event_gateway_virtual_cluster_produce_policy_schema_validation" "my_eventgatewayvirtualclusterproducepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    # ...
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "78615144-4664-462d-8c78-4f4bfa6e6bce"
  virtual_cluster_id = "f74f1c09-da30-4395-b13c-b00196b2cbf0"
}