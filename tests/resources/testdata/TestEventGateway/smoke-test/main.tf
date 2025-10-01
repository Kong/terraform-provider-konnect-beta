
resource "konnect_event_gateway" "my_eventgateway" {
  provider = konnect-beta

  labels = {
    key = "values"
  }
  name = "myeventgateway"
}

resource "konnect_event_gateway_listener" "my_eventgatewaylistener" {
  provider = konnect-beta
  addresses = [
    "..."
  ]
  description = "description"
  gateway_id  = konnect_event_gateway.my_eventgateway.id
  labels = {
    key = "values"
  }
  name = "my-egw-listener"
  ports = [
    {
      str = 25401
    }
  ]
}

resource "konnect_event_gateway_backend_cluster" "my_eventgatewaybackendcluster" {
  provider = konnect-beta
    authentication = {
        sasl_scram = {
        algorithm = "sha256"
        password = "passs"
        type = "sasl_scram"
        username = "user"
    }
    }
    bootstrap_servers = [
        "10.0.0.1:8080"
    ]
    description = "description"
    gateway_id = konnect_event_gateway.my_eventgateway.id
    insecure_allow_anonymous_virtual_cluster_auth = false
    metadata_update_interval_seconds = 22808
    name = "backendname"
}

resource "konnect_event_gateway_virtual_cluster" "my_eventgatewayvirtualcluster" {
  provider = konnect-beta
    authentication = [
      {
        sasl_plain = {
            mediation = "passthrough"
            principals = [
              {
                    password = "secret"
                    username = "my_username"
              }
            ]
            type = "sasl_plain"
        }
      }
    ]
    description = "description"
    destination = {
      id = konnect_event_gateway_backend_cluster.my_eventgatewaybackendcluster.id
    }
    dns_label = "vcluster-label"
    gateway_id = konnect_event_gateway.my_eventgateway.id
    name = "virtualnamex"
}

resource "konnect_event_gateway_vault" "my_eventgatewayvault" {
  provider = konnect-beta
  env = {
    config = {
      prefix = "KONG_"
    }
    description = "description"
    name = "myvault"
    type = "env"
  }
  gateway_id = konnect_event_gateway.my_eventgateway.id
}

resource "konnect_event_gateway_schema_registry" "my_eventgatewayschemaregistry" {
  provider = konnect-beta
  confluent = {
    config = {
      authentication = {
        basic = {
            password = "upass"
            type = "basic"
            username = "uname"
        }
      }
      endpoint = "https://example.com"
      schema_type = "avro"
      timeout_seconds = 8
    }
    description = "description"
    labels = {
        key = "value"
    }
    name = "confname"
    type = "confluent"
  }
  gateway_id = konnect_event_gateway.my_eventgateway.id
}

resource "konnect_event_gateway_listener_policy" "my_eventgatewaylistenerpolicy" {
  provider = konnect-beta
  forward_to_virtual_cluster = {
    config = {
      port_mapping = {
        advertised_host = "host.example.com"
        bootstrap_port = "none"
        destination = {
          virtual_cluster_reference_by_id = {
              id = konnect_event_gateway_virtual_cluster.my_eventgatewayvirtualcluster.id
          }
        }
        min_broker_id = 5
        type = "port_mapping"
      }
    }
    description = "mydescription"
    enabled = true
    labels = {
        key = "value2"
    }
    name = "listenerpolicyname"
    type = "forward_to_virtual_cluster"
  }
  event_gateway_listener_id = konnect_event_gateway_listener.my_eventgatewaylistener.id
  gateway_id = konnect_event_gateway.my_eventgateway.id
}

resource "konnect_event_gateway_virtual_cluster_produce_policy" "my_eventgatewayvirtualclusterproducepolicy" {
  provider = konnect-beta
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.my_eventgatewayvirtualcluster.id
  gateway_id = konnect_event_gateway.my_eventgateway.id
  modify_headers = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      actions = [
        {
          remove = {
              key = "myKey"
              op = "remove"
          }
        }
      ]
    }
    description = "mydescription"
    enabled = true
    labels = {
        key = "value"
    }
    name = "myproducepolicy"
    type = "modify_headers"
  }
}

resource "konnect_event_gateway_virtual_cluster_consume_policy" "my_eventgatewayvirtualclusterconsumepolicy" {
  provider = konnect-beta
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.my_eventgatewayvirtualcluster.id
  gateway_id = konnect_event_gateway.my_eventgateway.id
  modify_headers = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {
      actions = [
        {
          remove = {
            key = "mykey"
            op  = "remove"
          }
        }
      ]
    }
    description = "mydescription"
    enabled     = true
    name = "myconsumepolicy"
    type = "modify_headers"
  }
}

resource "konnect_event_gateway_virtual_cluster_cluster_policy" "my_eventgatewayvirtualclusterclusterpolicy" {
  provider = konnect-beta
  virtual_cluster_id = konnect_event_gateway_virtual_cluster.my_eventgatewayvirtualcluster.id
  gateway_id = konnect_event_gateway.my_eventgateway.id
  id_prefix = {
    condition = "context.topic.name.endsWith('my_suffix')"
    config = {}
    description = "mydescription"
    enabled     = true
    name = "myclusterpolicy"
    type = "id_prefix"
  }
}