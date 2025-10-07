resource "konnect_event_gateway_virtual_cluster_consume_policy_skip_record" "my_eventgatewayvirtualclusterconsumepolicyskiprecord" {
  provider = konnect-beta
  condition   = "context.topic.name.endsWith('my_suffix')"
  description = "...my_description..."
  enabled     = true
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "a788cc25-3732-4380-a1b7-5b1b257120ff"
  virtual_cluster_id = "375b7b37-bdda-47b7-b476-0d016b2f7992"
}