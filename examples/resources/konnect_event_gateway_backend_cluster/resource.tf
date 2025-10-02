resource "konnect_event_gateway_backend_cluster" "my_eventgatewaybackendcluster" {
  provider = konnect-beta
    authentication = {
    sasl_scram = {
    algorithm = "sha256"
    password = "${env['MY_SECRET']}"
    username = "...my_username..."
}
    }
    bootstrap_servers = [
        "..."
    ]
    description = "...my_description..."
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
    insecure_allow_anonymous_virtual_cluster_auth = false
    labels = {
        key = "value"
    }
    metadata_update_interval_seconds = 22808
    name = "...my_name..."
    tls = {
            ca_bundle = "...my_ca_bundle..."
        insecure_skip_verify = false
        tls_versions = [
            "tls12"
        ]
    }
}