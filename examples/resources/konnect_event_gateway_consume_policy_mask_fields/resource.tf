resource "konnect_event_gateway_consume_policy_mask_fields" "my_eventgatewayconsumepolicymaskfields" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    failure_mode = "mark"
    mask_fields = [
      {
        paths = [
          {
            match = "someObject.someArray[1].fieldName"
          }
        ]
        strategy = {
          keep_chars = {
            first  = 0
            last   = 0
            phrase = "...my_phrase..."
          }
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
  parent_policy_id   = "9f28729d-a908-496f-82e4-1630cf06a92b"
  virtual_cluster_id = "a2b1f942-a87e-4221-82e5-1ec1054fd0b0"
}