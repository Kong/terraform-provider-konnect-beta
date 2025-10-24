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
            key = "${vault.env['MY_ENV_VAR']}"
    }
    ]
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
}
    gateway_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
modify_headers = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            actions = [
        {
        remove = {
    key = "...my_key..."
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
}
skip_record = {
    condition = "context.topic.name.endsWith('my_suffix')"
    description = "...my_description..."
    enabled = true
    labels = {
        key = "value"
    }
    name = "...my_name..."
}
    virtual_cluster_id = "aacd895d-537a-41aa-baf6-3c8a82b63a74"
}