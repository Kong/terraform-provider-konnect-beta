resource "konnect_event_gateway_virtual_cluster_consume_policy_schema_validation" "my_eventgatewayvirtualclusterconsumepolicyschemavalidation" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    key_validation_action = "skip"
    schema_registry = {
      schema_registry_reference_by_id = {
        id = "33d5eb35-0229-4eb7-b590-6a91e82ef8ff"
      }
    }
    type                    = "json"
    value_validation_action = "mark"
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "7c3f71d6-cac7-4f75-a3d1-4bd94b2fa3d3"
  virtual_cluster_id = "41536846-8e2c-4a95-bf7c-0f0a423eca01"
}