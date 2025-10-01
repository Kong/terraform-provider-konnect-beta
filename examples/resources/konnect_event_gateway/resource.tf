resource "konnect_event_gateway" "my_eventgateway" {
  provider = konnect-beta
  labels = {
    key = "value"
  }
  name = "...my_name..."
}