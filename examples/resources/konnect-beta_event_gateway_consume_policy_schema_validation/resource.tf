resource "konnect-beta_event_gateway_consume_policy_schema_validation" "my_eventgatewayconsumepolicyschemavalidation" {
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    json = {
      failure_mode          = "mark"
      key_validation_action = "skip"
      schema_registry = {
        id = "bd775494-cdf4-4d78-a6df-c21ccf6d9813"
      }
      validate_key            = false
      validate_value          = false
      value_validation_action = "skip"
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