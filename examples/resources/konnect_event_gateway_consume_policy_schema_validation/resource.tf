resource "konnect_event_gateway_consume_policy_schema_validation" "my_eventgatewayconsumepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    key_validation_action = "mark"
    schema_registry = {
      schema_registry_reference_by_name = {
        name = "...my_name..."
      }
    }
    type                    = "json"
    value_validation_action = "mark"
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "afccd415-a99c-4465-8754-9932a66f275f"
}