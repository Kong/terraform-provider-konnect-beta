resource "konnect_event_gateway_produce_policy_encrypt" "my_eventgatewayproducepolicyencrypt" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith(\"my_suffix\") && record.headers[\"x-flag\"] == \"a-value\""
  config = {
    encryption_key = {
      static = {
        key = {
          id = "6e8740c2-fb76-4269-aeef-660d701c6ea1"
        }
      }
    }
    failure_mode = "error"
    part_of_record = [
      "key"
    ]
  }
  description = ""
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  virtual_cluster_id = "6ea3798e-38ca-4c28-a68e-1a577e478f2c"
}