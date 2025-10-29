resource "konnect-beta_event_gateway_produce_policy_skip_record" "my_eventgatewayproducepolicyskiprecord" {
  condition   = "context.topic.name.endsWith('my_suffix')"
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "d0f59208-a4ca-4178-ac9c-7da792d26d34"
  virtual_cluster_id = "f374230a-45cd-449a-b436-a516e43c3f9a"
}