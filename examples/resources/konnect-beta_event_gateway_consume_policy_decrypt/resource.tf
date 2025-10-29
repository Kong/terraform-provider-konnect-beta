resource "konnect-beta_event_gateway_consume_policy_decrypt" "my_eventgatewayconsumepolicydecrypt" {
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    decrypt = [
      {
        part_of_record = "key"
      }
    ]
    failure_mode = "passthrough"
    key_sources = [
      {
        aws = {
          # ...
        }
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
  parent_policy_id   = "969447b3-1e41-42d8-a020-1ebc4e88a916"
  virtual_cluster_id = "05c6c607-3c42-45e9-a9e8-3e6338120724"
}