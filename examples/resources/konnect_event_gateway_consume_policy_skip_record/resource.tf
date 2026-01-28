resource "konnect_event_gateway_consume_policy_skip_record" "my_eventgatewayconsumepolicyskiprecord" {
  provider = konnect-beta
  after       = "1fd29e98-9453-4db0-9ed3-3ce3c0e4b58b"
  before      = "75496408-24ab-49a9-af05-2c893ba297be"
  condition   = "record.value.content.foo.bar == \"a-value\""
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "ed46f132-b990-42ef-9430-a541cbd0a3f7"
  virtual_cluster_id = "d2c285d9-943d-4169-81ee-6317de3cc511"
}