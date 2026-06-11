resource "konnect_event_gateway_consume_policy_decrypt_fields" "my_eventgatewayconsumepolicydecryptfields" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    decrypt_fields = {
      paths = [
        {
          match = "someObject.someArray[1].fieldName"
        }
      ]
    }
    failure_mode = "mark"
    key_sources = [
      {
        static = {
          # ...
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
  parent_policy_id   = "1aaebe8e-3228-4062-baef-830b26256e35"
  virtual_cluster_id = "502b3102-936c-4278-b1c9-a3dc4c4dacc6"
}