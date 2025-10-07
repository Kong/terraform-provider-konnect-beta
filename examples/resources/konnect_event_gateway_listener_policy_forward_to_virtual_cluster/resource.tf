resource "konnect_event_gateway_listener_policy_forward_to_virtual_cluster" "my_eventgatewaylistenerpolicyforwardtovirtualcluster" {
  provider = konnect-beta
  condition = "context.topic.name.endsWith('my_suffix')"
  config = {
    sni = {
      advertised_port = 61579
      sni_suffix      = ".example.com"
    }
  }
  description               = "...my_description..."
  enabled                   = false
  event_gateway_listener_id = "6feda708-3b1b-4415-b1db-cf2694f34b09"
  gateway_id                = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name = "...my_name..."
}