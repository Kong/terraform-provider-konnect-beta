resource "konnect_event_gateway_produce_policy_schema_validation" "my_eventgatewayproducepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    inline_schema = {
      failure_mode = "reject"
      key = {
        json = {
          schema = "...my_schema..."
        }
      }
      key_validation_action = "mark"
      schema_registry = {
        id = "588c80ab-6c75-4adb-8bab-6d8b7e5f5ae2"
      }
      validate_key   = true
      validate_value = true
      value = {
        avro = {
          schema = "...my_schema..."
        }
      }
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
  virtual_cluster_id = "a3f4c612-4025-4392-861f-faa39b63e12d"
}