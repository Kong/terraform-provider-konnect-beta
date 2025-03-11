resource "konnect_mesh_control_plane" "my_meshcontrolplane" {
  provider = konnect-beta
  name        = "Mesh TF CP"
  description = "A control plane created using terraform"
  labels = {
    "terraform" = "true"
    "no_mesh" = "true"
  }
}

resource "konnect_mesh" "default" {
  provider = konnect-beta
  type  = "Mesh"
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id
  name  = "default"

  routing = {
    default_forbid_mesh_external_service_access = false
    # locality_aware_load_balancing               = true
    # zone_egress                                 = false
  }

  metrics = {
    backends = [
      {
        conf = {
          prometheus_metrics_backend_config = {
            port = "9090"
            path = "/metrics"
            tags = {
              "kuma.io/service" = "dataplane-metrics"
            }
          }
        }
        name = "prometheus_backend"
        type = "prometheus"
      }
    ]
    enabled_backend = "prometheus_backend"
  }

  logging = {
    backends = [
      {
        conf = {
          file_logging_backend_config = {
            path = "/var/log/mesh_access2.log"
          }
        }
        format = "{\"start_time\": \"%START_TIME%\", \"source\": \"%KUMA_SOURCE_SERVICE%\", \"destination\": \"%KUMA_DESTINATION_SERVICE%\", \"source_address\": \"%KUMA_SOURCE_ADDRESS_WITHOUT_PORT%\", \"destination_address\": \"%UPSTREAM_HOST%\", \"duration_millis\": \"%DURATION%\", \"bytes_received\": \"%BYTES_RECEIVED%\", \"bytes_sent\": \"%BYTES_SENT%\"}"
        name   = "file_backend"
        type   = "file"
      }
    ]
    default_backend = "file_backend"
  }

  constraints = {
    dataplane_proxy = {
      requirements = [
        {
          tags = {
            key = "a"
          }
        }
      ]
      restrictions = [
        {
          tags = {
            key = "*"
          }
        }
      ]
    }
  }

  labels = {
    environment = "production"
  }

  mesh_services = {
    mode = {
      str = "Exclusive"
    }
  }

  mtls = {
    "backends" = [
      {
        "conf" = {
          builtin_certificate_authority_config = {
            ca_cert = {
              expiration = "24h"
            }
          }
        }
        "dpCert" = {
          "rotation" : {
            "expiration" : "24h"
          }
        }
        "name" = "ca-1"
        "type" = "builtin"
      }
    ]
    "mode"           = "permissive"
    "enabledBackend" = "ca-1"
  }

  networking = {
    outbound = {
      passthrough = true
    }
  }

  tracing = {
    backends = [
      {
        conf = {
          datadog_tracing_backend_config = {
            address       = "datadog.datadog"
            port          = 4317
            split_service = true
          }
        }
        name     = "datadog_backend"
        sampling = 99
        type     = "datadog"
      },
      {
        conf = {
          zipkin_tracing_backend_config = {
            url          = "http://zipkin.example.com:9411/api/v2/spans"
            service_name = "mesh-trace-service"
          }
        }
        name     = "zipkin_backend"
        sampling = 100.0
        type     = "zipkin"
      }
    ]
    default_backend = "zipkin_backend"
  }
}
