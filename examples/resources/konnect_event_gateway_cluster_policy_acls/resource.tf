resource "konnect_event_gateway_cluster_policy_acls" "my_eventgatewayclusterpolicyacls" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    rules = [
      {
        action = "deny"
        operations = [
          {
            name = "describe_configs"
          }
        ]
        resource_names = [
          {
            match = "...my_match..."
          }
        ]
        resource_type = "transactional_id"
      }
    ]
  }
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "528824ff-4d3e-47af-9e16-af5bb53cc0fa"
  virtual_cluster_id = "4a444990-e7d1-4dfb-b2bf-2d8e113d1b6e"
}