resource "konnect_event_gateway_consume_policy_transcode" "my_eventgatewayconsumepolicytranscode" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    avro = {
      failure_mode = "passthrough"
      schema_ref_destination = {
        none = {
          # ...
        }
      }
      schema_source = {
        reference = {
          schema_registry = {
            id = "8e2017b8-a61c-477c-98ba-09dd461a6484"
          }
          subject = "context.topic.name + '-value'"
          version = "'latest'"
        }
      }
    }
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "65e2e968-cbcf-4527-96dd-2760bd98c676"
  virtual_cluster_id = "cdbc3581-884d-4a5c-b9e2-7d1a0067abae"
}