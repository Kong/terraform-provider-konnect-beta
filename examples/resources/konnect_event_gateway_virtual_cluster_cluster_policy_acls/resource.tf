resource "konnect_event_gateway_virtual_cluster_cluster_policy_acls" "my_eventgatewayvirtualclusterclusterpolicyacls" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    rules = [
      {
        action = "deny"
        operations = [
          {
            name = "read"
          }
        ]
        resource_names = [
          {
            match = "...my_match..."
          }
        ]
        resource_type = "topic"
      }
    ]
  }
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "1fe3b7ae-4986-4b5b-b2ba-ea4dc3efe778"
  virtual_cluster_id = "85b98a90-14b8-4895-9187-f8bc07ab33e3"
}