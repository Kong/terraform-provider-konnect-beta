resource "konnect_event_gateway_produce_policy_encrypt_fields" "my_eventgatewayproducepolicyencryptfields" {
  provider = konnect-beta
  condition = "record.value.content.foo.bar == \"a-value\""
  config = {
    encrypt_fields = [
      {
        encryption_key = {
          static = {
            key = {
              id = "5bf3a39d-29c8-4db0-a5cf-8b6709e7b297"
            }
          }
        }
        paths = [
          {
            match = "someObject.someArray[1].fieldName"
          }
        ]
        paths_expression = "$${context.auth.type == 'sasl_oauth_bearer' ? ['credentials.accessToken', 'credentials.refreshToken'] : ['credentials.password']}\n"
      }
    ]
    failure_mode = "mark"
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "53589770-fabe-4fbd-9218-75adfeed2fe4"
  virtual_cluster_id = "5a736330-9764-4c58-8d01-dbcb745874a1"
}