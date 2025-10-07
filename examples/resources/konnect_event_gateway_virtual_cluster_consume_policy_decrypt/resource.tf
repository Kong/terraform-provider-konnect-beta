resource "konnect_event_gateway_virtual_cluster_consume_policy_decrypt" "my_eventgatewayvirtualclusterconsumepolicydecrypt" {
  provider = konnect-beta
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
            decrypt = [
        {
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
    parent_policy_id = "a0deba9b-b5da-4259-8dd8-5ad0a02d313f"
    virtual_cluster_id = "49e1412d-8150-4edf-b053-cd5c67d5edf2"
}