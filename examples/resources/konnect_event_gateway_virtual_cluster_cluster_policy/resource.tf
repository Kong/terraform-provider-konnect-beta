resource "konnect_event_gateway_virtual_cluster_cluster_policy" "my_eventgatewayvirtualclusterclusterpolicy" {
  provider = konnect-beta
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  id_prefix = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      # ...
    }
    description = "...my_description..."
    enabled     = true
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "id_prefix"
  }
  parent_policy_id = "7429f02c-38f4-4e79-ab4c-60de5093f673"
  topic_rewrite = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      expression = {
        backend_to_virtual = "context.topic.name.substring(0, context.topic.name.length - 7)"
        virtual_to_backend = "context.topic.name + '-suffix'"
      }
      prefix = {
        value = "...my_value..."
      }
      type = "expression"
    }
    description = "...my_description..."
    enabled     = true
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "topic_rewrite"
  }
  virtual_cluster_id = "bb8263d9-8ddb-4bc4-bee6-3372ef4b5cbc"
}