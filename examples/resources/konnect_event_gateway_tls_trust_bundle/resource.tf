resource "konnect_event_gateway_tls_trust_bundle" "my_eventgatewaytlstrustbundle" {
  provider = konnect-beta
  config = {
    trusted_ca = "...my_trusted_ca..."
  }
  description = ""
  gateway_id  = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  labels = {
    key = "value"
  }
  name = "...my_name..."
}