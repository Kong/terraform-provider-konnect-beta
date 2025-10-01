resource "konnect_event_gateway_virtual_cluster_consume_policy" "my_eventgatewayvirtualclusterconsumepolicy" {
  provider = konnect-beta
  decrypt = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      decrypt = [
        {
          part_of_record = "key"
        }
      ]
      failure_mode = "error"
      key_sources = [
        {
          aws = {
            type = "aws"
          }
        }
      ]
    }
    description = "...my_description..."
    enabled     = false
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "decrypt"
  }
  gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  modify_headers = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      actions = [
        {
          remove = {
            key = "...my_key..."
            op  = "remove"
          }
        }
      ]
    }
    description = "...my_description..."
    enabled     = false
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "modify_headers"
  }
  parent_policy_id = "8df180ce-a71e-4f3c-8749-4f35ff3ab10d"
  schema_validation = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      key_validation_action = "skip"
      schema_registry = {
        schema_registry_reference_by_id = {
          id = "7671d8af-bc5a-4d37-9234-171178b1a48d"
        }
      }
      type                    = "json"
      value_validation_action = "skip"
    }
    description = "...my_description..."
    enabled     = true
    labels = {
      key = "value"
    }
    name = "...my_name..."
    type = "schema_validation"
  }
  virtual_cluster_id = "47a96a37-fce2-44e7-966b-b0c02a478a0f"
}