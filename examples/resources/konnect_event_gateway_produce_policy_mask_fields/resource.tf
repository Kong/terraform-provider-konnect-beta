resource "konnect_event_gateway_produce_policy_mask_fields" "my_eventgatewayproducepolicymaskfields" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    failure_mode = "passthrough"
    mask_fields = [
      {
        paths = [
          {
            match = "someObject.someArray[1].fieldName"
          }
        ]
        strategy = {
          replace = {
            replace = {
              phrase = "...my_phrase..."
            }
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
  parent_policy_id   = "871e4106-2200-408a-9305-98647c57aee2"
  virtual_cluster_id = "7c87e642-a8c0-47b0-bf5f-35f8045c2887"
}