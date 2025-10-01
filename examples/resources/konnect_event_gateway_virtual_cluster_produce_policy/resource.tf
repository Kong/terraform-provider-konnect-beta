resource "konnect_event_gateway_virtual_cluster_produce_policy" "my_eventgatewayvirtualclusterproducepolicy" {
  provider = konnect-beta
encrypt = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            encrypt = [
        {
                        key_id = "static://static-key-named-in-source"
                part_of_record = "value"
        }
        ]
        failure_mode = "passthrough"
        key_sources = [
        {
        static = {
    keys = [
    {
                id = "...my_id..."
            key = "${env['MY_SECRET']}"
    }
    ]
    type = "static"
}
        }
        ]
    }
    description = "...my_description..."
    enabled = false
    labels = {
        key = "value"
    }
    name = "...my_name..."
    type = "encrypt"
}
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
modify_headers = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            actions = [
        {
        remove = {
    key = "...my_key..."
    op = "remove"
}
        }
        ]
    }
    description = "...my_description..."
    enabled = true
    labels = {
        key = "value"
    }
    name = "...my_name..."
    type = "modify_headers"
}
    parent_policy_id = "ce933254-b3f6-41d7-adfc-3db1ba8b90ca"
schema_validation = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
        # ...
    }
    description = "...my_description..."
    enabled = true
    labels = {
        key = "value"
    }
    name = "...my_name..."
    type = "schema_validation"
}
    virtual_cluster_id = "aacd895d-537a-41aa-baf6-3c8a82b63a74"
}