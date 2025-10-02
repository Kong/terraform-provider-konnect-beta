
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
    "25401"
  ]
}

resource "konnect_event_gateway_backend_cluster" "my_eventgatewaybackendcluster" {
  provider = konnect-beta
    authentication = {
        sasl_scram = {
        algorithm = "sha256"
        password = "passs"
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
  }
  gateway_id = konnect_event_gateway.my_eventgateway.id
}
