resource "konnect_event_gateway_virtual_cluster" "my_eventgatewayvirtualcluster" {
  provider = konnect-beta
    acl_mode = "enforce_on_gateway"
    authentication = [
    {
    sasl_plain = {
    mediation = "passthrough"
    principals = [
    {
                password = "${env['MY_SECRET']}"
            username = "...my_username..."
    }
    ]
}
    }
    ]
    description = "...my_description..."
    destination = {
    backend_cluster_reference_by_id = {
    id = "80206173-845d-4ab3-9e6a-f454c3cd4baf"
}
    }
    dns_label = "vcluster-1"
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
    labels = {
        key = "value"
    }
    name = "...my_name..."
    namespace = {
            additional = {
                    consumer_groups = [
            {
            glob = {
    glob = "...my_glob..."
}
            }
            ]
            topics = [
            {
            exact_list = {
    conflict = "warn"
    exact_list = [
    {
                backend = "...my_backend..."
    }
    ]
}
            }
            ]
        }
        mode = "hide_prefix"
        prefix = "...my_prefix..."
    }
}