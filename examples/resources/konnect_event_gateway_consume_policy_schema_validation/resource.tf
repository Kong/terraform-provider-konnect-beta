resource "konnect_event_gateway_consume_policy_schema_validation" "my_eventgatewayconsumepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    key_validation_action = "mark"
    schema_registry = {
      id = "e1881384-290f-443c-a5bd-ed6f2e53d539"
    }
    type                    = "json"
    value_validation_action = "mark"
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "afccd415-a99c-4465-8754-9932a66f275f"
}