resource "konnect_event_gateway_virtual_cluster" "my_eventgatewayvirtualcluster" {
  provider = konnect-beta
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
    type = "sasl_plain"
}
    }
    ]
    description = "...my_description..."
    destination = {
            id = "759b5471-3de4-485c-b7d3-6e8cb8929d81"
        name = "...my_name..."
    }
    dns_label = "vcluster-1"
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
    labels = {
        key = "value"
    }
    name = "...my_name..."
}