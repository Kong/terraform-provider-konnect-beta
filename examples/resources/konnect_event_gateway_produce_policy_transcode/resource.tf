resource "konnect_event_gateway_produce_policy_transcode" "my_eventgatewayproducepolicytranscode" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    avro = {
      failure_mode     = "mark"
      on_unknown_field = "fail"
      schema_source = {
        reference = {
          schema_ref_destination = {
            none = {
              # ...
            }
          }
          schema_registry = {
            id = "de460e59-30fd-45da-af2b-292a5823961a"
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
  parent_policy_id   = "1c399fc8-e15c-4021-b66e-7a228d451dd7"
  virtual_cluster_id = "66575524-3680-4571-8e62-88377b472727"
}