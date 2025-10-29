resource "konnect_event_gateway_listener_policy_tls_server" "my_eventgatewaylistenerpolicytlsserver" {
  provider = konnect-beta
condition = "context.topic.name.endsWith('my_suffix')"
config = {
allow_plaintext = true
certificates = [
{
certificate = "...my_certificate..."
key = "${vault.env['MY_ENV_VAR']}"
}
]
versions = {
max = "TLSv1.3"
min = "TLSv1.3"
}
}
description = "...my_description..."
enabled = true
event_gateway_listener_id = "34102bf1-bf41-4c00-a62f-6fca747cb8f8"
gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
labels = {
    key = "value"
}
name = "...my_name..."
}