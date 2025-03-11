resource "konnect_mesh_external_service" "my_meshexternalservice" {
  provider = konnect-beta
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id

  labels = {
    environment = "production"
    "kuma.io/mesh" = konnect_mesh.default.name
  }

  mesh = konnect_mesh.default.name
  name = "external-service-example"

  spec = {
    endpoints = [
      {
        address = "example.com"
        port    = 9478
      }
    ]
    match = {
      port     = 1444
      protocol = "tcp"
      type     = "HostnameGenerator"
    }
    tls = {
      allow_renegotiation = false
      enabled             = true
      verification = {
        ca_cert = {
          inline_string = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A...\n-----END CERTIFICATE-----"
          secret        = "my-ca-secret"
        }
        client_cert = {
          inline_string = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8B...\n-----END CERTIFICATE-----"
          secret        = "my-client-cert-secret"
        }
        client_key = {
          inline_string = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAt...\n-----END RSA PRIVATE KEY-----"
          secret        = "my-client-key-secret"
        }
        mode        = "SkipAll"
        server_name = "my.server.name"
        subject_alt_names = [
          {
            type  = "Exact"
            value = "my.example.com"
          }
        ]
      }
      version = {
        max = "TLS12"
        min = "TLS10"
      }
    }
  }

  type = "MeshExternalService"

  depends_on = [konnect_mesh.default]
}

resource "konnect_mesh_access_log" "my-access-log" {
  provider = konnect-beta
  type  = "MeshAccessLog"
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id
  depends_on = [konnect_mesh.default]
  name  = "my-access-log"
  mesh  = konnect_mesh.default.name

  labels = {
    "kuma.io/mesh" = konnect_mesh.default.name
    environment = "production"
    region      = "us-west"
    team        = "devops"
  }

  spec = {
    from = [
      {
        target_ref = {
          kind = "Mesh"
        }
        default = {
          backends = [
            {
              file = {
                format = {
                  json = [
                    {
                      key   = "start_time"
                      value = "%START_TIME%"
                    },
                    {
                      key   = "source"
                      value = "%KUMA_SOURCE_SERVICE%"
                    },
                    {
                      key   = "destination"
                      value = "%KUMA_DESTINATION_SERVICE%"
                    },
                    {
                      key   = "duration_millis"
                      value = "%DURATION%"
                    }
                  ]
                  type              = "Json"
                }
                path = "/var/log/mesh_access.log"
              }
              type = "File"
            },
            {
              open_telemetry = {
                attributes = [
                  {
                    key   = "mesh"
                    value = "%KUMA_MESH%"
                  },
                  {
                    key   = "service"
                    value = "%SERVICE_NAME%"
                  }
                ]
                body     = "{ \"values\": [{ \"key\": \"service\", \"value\": { \"stringValue\": \"%SERVICE_NAME%\" } }] }"
                endpoint = "otel-collector:4317"
              }
              type = "OpenTelemetry"
            },
            {
              tcp = {
                address = "127.0.0.1:5000"
                format = {
                  json = [
                    {
                      key   = "start_time"
                      value = "%START_TIME%"
                    },
                    {
                      key   = "bytes_sent"
                      value = "%BYTES_SENT%"
                    }
                  ]
                  omit_empty_values = false
                  type              = "Json"
                }
              }
              type = "Tcp"
            }
          ]
        }
      }
    ]
    to = [
      {
        target_ref = {
          kind = "Mesh"
        }
        default = {
          backends = [
            {
              file = {
                format = {
                  omit_empty_values = true
                  plain             = "[%START_TIME%] %KUMA_SOURCE_SERVICE% -> %KUMA_DESTINATION_SERVICE%"
                  type              = "Plain"
                }
                path = "/var/log/mesh_access.log"
              }
              type = "File"
            },
            {
              open_telemetry = {
                attributes = [
                  {
                    key   = "mesh"
                    value = "%KUMA_MESH%"
                  },
                  {
                    key   = "service"
                    value = "%SERVICE_NAME%"
                  }
                ]
                body     = "{ \"values\": [{ \"key\": \"service\", \"value\": { \"stringValue\": \"%SERVICE_NAME%\" } }] }"
                endpoint = "otel-collector:4317"
              }
              type = "OpenTelemetry"
            },
            {
              tcp = {
                address = "127.0.0.1:5000"
                format = {
                  plain = "[%START_TIME%] %KUMA_MESH% -> %UPSTREAM_HOST%"
                  type  = "Plain"
                }
              }
              type = "Tcp"
            }
          ]
        }
      }
    ]
  }
}

resource "konnect_mesh_circuit_breaker" "my_meshcircuitbreaker" {
  provider = konnect-beta
  cp_id             = konnect_mesh_control_plane.my_meshcontrolplane.id
  name              = "web-to-backend-circuit-breaker"
  mesh              = konnect_mesh.default.name

  labels = {
    "kuma.io/mesh" = konnect_mesh.default.name
    environment = "production"
    team        = "devops"
  }

  spec = {
    # Inbound configuration (from)
    from = [
      {
        default = {
          connection_limits = {
            max_connection_pools = 9
            max_connections      = 1
            max_pending_requests = 3
            max_requests         = 7
            max_retries          = 3
          }
          outlier_detection = {
            base_ejection_time = "30s"
            detectors = {
              failure_percentage = {
                minimum_hosts  = 7
                request_volume = 5
                threshold      = 1
              }
              gateway_failures = {
                consecutive = 2
              }
              local_origin_failures = {
                consecutive = 10
              }
              success_rate = {
                minimum_hosts               = 9
                request_volume              = 6
                standard_deviation_factor = {
                  str = "1.9"
                }
              }
              total_failures = {
                consecutive = 7
              }
            }
            disabled                        = false
            interval                        = "5s"
            max_ejection_percent            = 20
            split_external_and_local_errors = true
          }
        }
        # Use a simple Mesh reference for 'from'
        target_ref = {
          kind = "Mesh"
        }
      }
    ]

    # Top-level targetRef: since kind is MeshSubset, only tags are allowed.
    target_ref = {
      kind = "MeshSubset"
      tags = {
        app = "web"
      }
    }

    # Outbound configuration (to)
    to = [
      {
        default = {
          connection_limits = {
            max_connection_pools = 4
            max_connections      = 8
            max_pending_requests = 3
            max_requests         = 2
            max_retries          = 3
          }
          outlier_detection = {
            base_ejection_time = "30s"
            detectors = {
              failure_percentage = {
                minimum_hosts  = 1
                request_volume = 10
                threshold      = 0
              }
              gateway_failures = {
                consecutive = 10
              }
              local_origin_failures = {
                consecutive = 5
              }
              success_rate = {
                minimum_hosts               = 9
                request_volume              = 3
                standard_deviation_factor = {
                  str = "1.9"
                }
              }
              total_failures = {
                consecutive = 2
              }
            }
            disabled                        = true
            interval                        = "5s"
            max_ejection_percent            = 4
            split_external_and_local_errors = true
          }
        }
        # Use MeshService for 'to' (as in the K8s example)
        target_ref = {
          kind         = "MeshService"
          name         = "backend"
          namespace    = "kuma-demo"
          section_name = "http"
        }
      }
    ]
  }

  type = "MeshCircuitBreaker"

  depends_on = [konnect_mesh.default]
}

resource "konnect_mesh_retry" "my_meshretry" {
  provider = konnect-beta
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id
  labels = {
    environment = "production"
  }
  mesh = konnect_mesh.default.name
  name = "web-to-backend-retry"

  spec = {
    # Top-level targetRef: using MeshSubset (virtual reference)
    target_ref = {
      kind = "MeshSubset"
      tags = {
        app = "web"
      }
    }

    to = [
      {
        default = {
          grpc = {
            back_off = {
              base_interval = "1s"
              max_interval  = "5s"
            }
            num_retries     = 2
            per_try_timeout = "2s"
            rate_limited_back_off = {
              max_interval = "5s"
              reset_headers = [
                {
                  format = "UnixTimestamp"
                  name   = "retry-after"
                }
              ]
            }
            retry_on = [
              "DeadlineExceeded"
            ]
          }
          http = {
            back_off = {
              base_interval = "1s"
              max_interval  = "5s"
            }
            host_selection = [
              {
                predicate        = "OmitPreviousHosts"
                tags             = { app = "web" }
                update_frequency = 6
              }
            ]
            host_selection_max_attempts = 1
            num_retries                 = 6
            per_try_timeout             = "2s"
            rate_limited_back_off = {
              max_interval = "5s"
              reset_headers = [
                {
                  format = "Seconds"
                  name   = "x-retry-reset"
                }
              ]
            }
            retriable_request_headers = [
              {
                name  = "x-retry"
                type  = "RegularExpression"
                value = "retry-value"
              }
            ]
            retriable_response_headers = [
              {
                name  = "x-retry-response"
                type  = "RegularExpression"
                value = "retry-response-value"
              }
            ]
            retry_on = [
              "503"
            ]
          }
          tcp = {
            max_connect_attempt = 1
          }
        }
        # In the "to" block, use a concrete target reference (here, a MeshService)
        target_ref = {
          kind         = "MeshService"
          name         = "backend"
          namespace    = "kuma-demo"
          section_name = "http"
        }
      }
    ]
  }

  type = "MeshRetry"

  depends_on = [konnect_mesh.default]
}

resource "konnect_mesh_fault_injection" "my_meshfaultinjection" {
  provider = konnect-beta
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id

  labels = {
    key = "value"
    "kuma.io/mesh" = "default"
  }

  mesh = konnect_mesh.default.name
  name = "fault-injection-example"

  spec = {
    # Top-level target_ref defines the overall scope of the policy.
    target_ref = {
      kind = "MeshSubset"
      tags = {
        key = "value"
      }
    }
    to = [
      {
        default = {
          http = [
            {
              abort = {
                http_status = 503
                percentage  = { integer = 10 }
              }
              delay = {
                percentage = { integer = 2 }
                value      = "200ms"
              }
              response_bandwidth = {
                limit      = "5kbps"
                percentage = { integer = 6 }
              }
            }
          ]
        }
        # Use a concrete destination target reference. "MeshService" is supported here.
        target_ref = {
          kind         = "Mesh"
        }
      }
    ]
  }

  type = "MeshFaultInjection"

  depends_on = [konnect_mesh.default]
}

resource "konnect_mesh_global_rate_limit" "my_meshglobalratelimit" {
  provider = konnect-beta
  cp_id = konnect_mesh_control_plane.my_meshcontrolplane.id

  labels = {
    environment = "production"
    "kuma.io/mesh" = "default"
  }

  mesh = konnect_mesh.default.name
  name = "global-rate-limit-example"

  spec = {
    from = [
      {
        default = {
          backend = {
            rate_limit_service = {
              limit_on_service_fail = true
              timeout               = "5s"
              url                   = "http://ratelimit.example.com"
            }
          }
          http = {
            disabled = false
            on_rate_limit = {
              headers = {
                add = [
                  {
                    name  = "x-ratelimit-add"
                    value = "add-value"
                  }
                ]
                set = [
                  {
                    name  = "x-ratelimit-set"
                    value = "set-value"
                  }
                ]
              }
              status = 429
            }
            request_rate = {
              interval = "1m0s"
              num      = 50
            }
          }
        }
        target_ref = {
          kind = "Mesh"
        }
      }
    ]
    target_ref = {
      kind = "MeshService"
      name = "client"
    }
  }

  type = "MeshGlobalRateLimit"

  depends_on = [konnect_mesh.default]
}