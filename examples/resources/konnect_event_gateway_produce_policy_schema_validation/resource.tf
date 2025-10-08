resource "konnect_event_gateway_produce_policy_schema_validation" "my_eventgatewayproducepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    confluent_schema_registry = {
      key_validation_action = "reject"
      schema_registry = {
        schema_registry_reference_by_name = {
          name = "...my_name..."
        }
      }
      value_validation_action = "reject"
    }
    json = {
      key_validation_action = "mark"
      schema_registry = {
        schema_registry_reference_by_name = {
          name = "...my_name..."
        }
      }
      value_validation_action = "reject"
    }
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "3fd88180-5e64-4151-9b2e-6733e74cddc6"
  virtual_cluster_id = "a3f4c612-4025-4392-861f-faa39b63e12d"
}