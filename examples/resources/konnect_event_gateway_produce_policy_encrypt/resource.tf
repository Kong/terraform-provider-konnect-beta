resource "konnect_event_gateway_produce_policy_encrypt" "my_eventgatewayproducepolicyencrypt" {
  provider = konnect-beta
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            encrypt = [
        {
                        key_id = "static://static-key-named-in-source"
                part_of_record = "value"
        }
        ]
        failure_mode = "error"
        key_sources = [
        {
        static = {
    keys = [
    {
                id = "...my_id..."
            key = "${env['MY_SECRET']}"
    }
    ]
}
        }
        ]
    }
    description = "...my_description..."
    enabled = true
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
    labels = {
        key = "value"
    }
    name = "...my_name..."
    parent_policy_id = "d360a229-0d2f-4566-9b8e-dad95ffde3d0"
    virtual_cluster_id = "6ea3798e-38ca-4c28-a68e-1a577e478f2c"
}