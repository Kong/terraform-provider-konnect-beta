resource "konnect_event_gateway_listener_policy_tls_server" "my_eventgatewaylistenerpolicytlsserver" {
  provider = konnect-beta
config = {
allow_plaintext = false
certificates = [
{
certificate = "...my_certificate..."
key = "${vault.env['MY_ENV_VAR']}"
}
]
client_authentication = {
mode = "requested"
principal_mapping = "${context.certificate.subject['CN'] ? context.certificate.subject['CN'] : context.certificate.sans.uri[0]}"
tls_trust_bundles = [
{
id = "4207e6bd-68dd-4f60-bac1-adbb586553d5"
}
]
}
versions = {
max = "TLSv1.3"
min = "TLSv1.2"
}
}
description = ""
enabled = true
gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
labels = {
    key = "value"
}
listener_id = "f7d7b9be-5608-44c3-8f6a-46e055797c31"
name = "...my_name..."
}