resource "konnect_event_gateway_schema_registry" "my_eventgatewayschemaregistry" {
  provider = konnect-beta
gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
schema_registry_confluent = {
config = {
authentication = {
basic = {
password = "${env['MY_SECRET']}"
username = "...my_username..."
}
}
endpoint = "https://upright-plumber.info"
schema_type = "json"
timeout_seconds = 3
}
description = "...my_description..."
labels = {
    key = "value"
}
name = "...my_name..."
}
}