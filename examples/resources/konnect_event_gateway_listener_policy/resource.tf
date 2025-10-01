resource "konnect_event_gateway_listener_policy" "my_eventgatewaylistenerpolicy" {
  provider = konnect-beta
    event_gateway_listener_id = "48c9353f-7dda-464b-8fe4-e6972a3aeae2"
forward_to_virtual_cluster = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
    port_mapping = {
    advertised_host = "...my_advertised_host..."
    bootstrap_port = "none"
    destination = {
    virtual_cluster_reference_by_id = {
    id = "0cf1742b-423d-4928-be30-69ffe732f3fd"
}
    }
    min_broker_id = 5
    type = "port_mapping"
}
    }
    description = "...my_description..."
    enabled = true
    labels = {
        key = "value"
    }
    name = "...my_name..."
    type = "forward_to_virtual_cluster"
}
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
tls_server = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            allow_plaintext = false
        certificates = [
        {
                        certificate = "...my_certificate..."
                key = "${env['MY_SECRET']}"
        }
        ]
        versions = {
                    max = "TLSv1.2"
            min = "TLSv1.3"
        }
    }
    description = "...my_description..."
    enabled = true
    labels = {
        key = "value"
    }
    name = "...my_name..."
    type = "tls_server"
}
}