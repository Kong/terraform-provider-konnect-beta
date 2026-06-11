resource "konnect_event_gateway_produce_policy_schema_validation" "my_eventgatewayproducepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    json = {
      failure_mode          = "passthrough"
      key_validation_action = "mark"
      schema_registry = {
        id = "74577697-03b2-4d40-bfe2-929c891c4254"
      }
      validate_key            = true
      validate_value          = false
      value_validation_action = "reject"
    }
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "a3f4c612-4025-4392-861f-faa39b63e12d"
}