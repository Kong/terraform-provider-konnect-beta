resource "konnect_event_gateway_consume_policy_schema_validation" "my_eventgatewayconsumepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    confluent_schema_registry = {
      failure_mode          = "error"
      key_validation_action = "skip"
      schema_registry = {
        id = "7dae82e2-f401-4f9c-bf87-ce7f92d34aa1"
      }
      validate_key            = true
      validate_value          = false
      value_validation_action = "mark"
    }
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