resource "konnect_event_gateway_virtual_cluster_produce_policy_skip_record" "my_eventgatewayvirtualclusterproducepolicyskiprecord" {
  provider = konnect-beta
  condition   = "context.topic.name.endsWith('my_suffix')"
  description = "...my_description..."
  enabled     = false
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name               = "...my_name..."
  parent_policy_id   = "ef602745-6b1e-4f02-a3e9-c16032e5bd38"
  virtual_cluster_id = "540bae99-c32c-467f-ae67-44897861e0ae"
}