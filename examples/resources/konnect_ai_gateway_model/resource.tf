resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  provider = konnect-beta
  api = {
    access = {
      acls = {
        allow = [
          "..."
        ]
        allow_when = [
          "..."
        ]
        deny = [
          "..."
        ]
        deny_when = [
          "..."
        ]
      }
      auth_strategies = [
        "okta-ai-se"
      ]
      identity_providers = [
        "okta-ai-se"
      ]
    }
    capabilities = [
      "files"
    ]
    config = {
      balancer = {
        priority = {
          connect_timeout = 60000
          fail_timeout    = 10000
          failover_criteria = [
            "http_503"
          ]
          max_fails    = 0
          read_timeout = 60000
          retries      = 5
          session_affinity = {
            enabled = true
            storage = {
              redis = {
                cloud_authentication = {
                  gcp = {
                    service_account_json = "...my_service_account_json..."
                  }
                }
                cluster = {
                  max_redirections = 5
                  nodes = [
                    {
                      ip   = "127.0.0.1"
                      port = 6379
                    }
                  ]
                }
                connect_timeout       = 2000
                connection_is_proxied = false
                database              = 0
                host                  = "127.0.0.1"
                keepalive = {
                  backlog   = 989403147
                  pool_size = 256
                }
                password     = "...my_password..."
                port         = 6379
                read_timeout = 2000
                send_timeout = 2000
                sentinel = {
                  master = "...my_master..."
                  nodes = [
                    {
                      host = "127.0.0.1"
                      port = 6379
                    }
                  ]
                  password = "...my_password..."
                  role     = "slave"
                  username = "...my_username..."
                }
                server_name = "...my_server_name..."
                ssl         = true
                ssl_verify  = true
                username    = "...my_username..."
              }
            }
            ttl = 21600
          }
          slots         = 10000
          write_timeout = 60000
        }
      }
      logging = {
        payloads = false
      }
      max_request_body_size = 3
      proxy = {
        auth = {
          password = "...my_password..."
          username = "...my_username..."
        }
        http_proxy = {
          host = "...my_host..."
          port = 29747
        }
        https_proxy = {
          host = "...my_host..."
          port = 12764
        }
        no_proxy     = "...my_no_proxy..."
        proxy_scheme = "http"
      }
      response_streaming = "allow"
      route = {
        headers = {
          key = jsonencode("value")
        }
        hosts = [
          "foo.example.com"
        ]
        https_redirect_status_code = 426
        methods = [
          "..."
        ]
        model = {
          body_param   = "model"
          header_param = "x-model"
          path_param   = "model_name"
          values = [
            "..."
          ]
        }
        paths = [
          "..."
        ]
        preserve_host = false
        protocols = [
          "..."
        ]
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
        tags = [
          "..."
        ]
      }
    }
    datastores = [
      {
        name    = "shared-redis"
        purpose = "semantic-balancer"
      }
    ]
    display_name = "My GPT 5 model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-gpt-5-model"
    policies = [
      "..."
    ]
    targets = [
      {
        allow_auth_override = false
        config = {
          ollama = {
            cache_read_cost = 6.36
            cache_read_cost_list = [
              {
                cost  = 3.01
                modal = "video"
              }
            ]
            cache_write_cost = 7.25
            cache_write_cost_list = [
              {
                cost = 0.01
                ttl  = "...my_ttl..."
              }
            ]
            context_window_factor = [
              {
                above         = "...my_above..."
                input_factor  = 3.5
                output_factor = 1.2
              }
            ]
            embeddings_dimensions = 1163998112
            input_cost            = 3.53
            input_cost_list = [
              {
                cost  = 2.6
                modal = "video"
              }
            ]
            max_tokens  = 2030002915
            output_cost = 7.78
            output_cost_list = [
              {
                cost  = 4.83
                modal = "image"
              }
            ]
            service_tier_factor = [
              {
                factor = 0.08
                tier   = "...my_tier..."
              }
            ]
            temperature       = 3.02
            top_k             = 164
            top_p             = 0.35
            upstream_base_url = "https://waterlogged-luck.com"
            upstream_url      = "https://brave-calculus.biz"
          }
        }
        name                 = "gpt-5-model"
        provider             = "azure-ai-se"
        semantic_description = "...my_semantic_description..."
        weight               = 100
      }
    ]
  }
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  model = {
    access = {
      acls = {
        allow = [
          "..."
        ]
        allow_when = [
          "..."
        ]
        deny = [
          "..."
        ]
        deny_when = [
          "..."
        ]
      }
      auth_strategies = [
        "okta-ai-se"
      ]
      identity_providers = [
        "okta-ai-se"
      ]
    }
    capabilities = [
      "audio/speech"
    ]
    config = {
      balancer = {
        priority = {
          connect_timeout = 60000
          fail_timeout    = 10000
          failover_criteria = [
            "http_408"
          ]
          max_fails    = 0
          read_timeout = 60000
          retries      = 5
          session_affinity = {
            enabled = true
            storage = {
              redis = {
                cloud_authentication = {
                  azure = {
                    client_id     = "...my_client_id..."
                    client_secret = "...my_client_secret..."
                    tenant_id     = "...my_tenant_id..."
                  }
                }
                cluster = {
                  max_redirections = 5
                  nodes = [
                    {
                      ip   = "127.0.0.1"
                      port = 6379
                    }
                  ]
                }
                connect_timeout       = 2000
                connection_is_proxied = false
                database              = 0
                host                  = "127.0.0.1"
                keepalive = {
                  backlog   = 1242809940
                  pool_size = 256
                }
                password     = "...my_password..."
                port         = 6379
                read_timeout = 2000
                send_timeout = 2000
                sentinel = {
                  master = "...my_master..."
                  nodes = [
                    {
                      host = "127.0.0.1"
                      port = 6379
                    }
                  ]
                  password = "...my_password..."
                  role     = "master"
                  username = "...my_username..."
                }
                server_name = "...my_server_name..."
                ssl         = true
                ssl_verify  = true
                username    = "...my_username..."
              }
            }
            ttl = 21600
          }
          slots         = 10000
          write_timeout = 60000
        }
      }
      logging = {
        payloads = false
      }
      max_request_body_size = 4
      model = {
        name_header = true
      }
      proxy = {
        auth = {
          password = "...my_password..."
          username = "...my_username..."
        }
        http_proxy = {
          host = "...my_host..."
          port = 30633
        }
        https_proxy = {
          host = "...my_host..."
          port = 29606
        }
        no_proxy     = "...my_no_proxy..."
        proxy_scheme = "http"
      }
      response_streaming = "allow"
      route = {
        headers = {
          key = jsonencode("value")
        }
        hosts = [
          "foo.example.com"
        ]
        https_redirect_status_code = 426
        methods = [
          "..."
        ]
        model = {
          body_param   = "model"
          header_param = "x-model"
          path_param   = "model_name"
          values = [
            "..."
          ]
        }
        paths = [
          "..."
        ]
        preserve_host = false
        protocols = [
          "..."
        ]
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
        tags = [
          "..."
        ]
      }
    }
    datastores = [
      {
        name    = "shared-redis"
        purpose = "semantic-balancer"
      }
    ]
    display_name = "My GPT 5 model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-gpt-5-model"
    policies = [
      "..."
    ]
    targets = [
      {
        allow_auth_override = false
        config = {
          typesafe = {
            input_cost        = 4.69
            output_cost       = 7.9
            upstream_base_url = "https://noteworthy-almighty.info"
            upstream_url      = "https://equatorial-pants.com/"
          }
        }
        name                 = "gpt-5-model"
        provider             = "azure-ai-se"
        semantic_description = "...my_semantic_description..."
        weight               = 100
      }
    ]
  }
}