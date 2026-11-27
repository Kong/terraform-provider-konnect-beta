resource "konnect_event_gateway_cluster_policy_request_rule_validator" "my_eventgatewayclusterpolicyrequestrulevalidator" {
  provider = konnect-beta
  condition = "context.auth.principal.name == \"this-user\""
  config = {
    requests = [
      {
        consumer_group_heartbeat = {
          rules = [
            {
              action      = "passthrough"
              description = "...my_description..."
              rule        = "context.group.id.startsWith('team-a.')"
            }
          ]
        }
      }
    ]
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "14fdb46c-b44c-45b7-974e-c5643ecb64db"
}