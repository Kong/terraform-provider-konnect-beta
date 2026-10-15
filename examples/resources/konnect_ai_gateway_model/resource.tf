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
        lowest_latency = {
          connect_timeout = 60000
          fail_timeout    = 10000
          failover_criteria = [
            "http_408"
          ]
          latency_strategy = "tpot"
          max_fails        = 0
          read_timeout     = 60000
          retries          = 5
          slots            = 10000
          write_timeout    = 60000
        }
      }
      logging = {
        payloads = false
      }
      max_request_body_size = 8388608
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
        name = "shared-redis"
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
          databricks = {
            cache_read_cost = 9.25
            cache_read_cost_list = [
              {
                cost  = 7.56
                modal = "text"
              }
            ]
            cache_write_cost = 8.79
            cache_write_cost_list = [
              {
                cost = 8.23
                ttl  = "...my_ttl..."
              }
            ]
            context_window_factor = [
              {
                above         = "...my_above..."
                input_factor  = 6.86
                output_factor = 3.98
              }
            ]
            embeddings_dimensions = 2017878719
            input_cost            = 3.88
            input_cost_list = [
              {
                cost  = 9.74
                modal = "video"
              }
            ]
            max_tokens  = 184807158
            output_cost = 3.66
            output_cost_list = [
              {
                cost  = 5.89
                modal = "audio"
              }
            ]
            service_tier_factor = [
              {
                factor = 4.06
                tier   = "...my_tier..."
              }
            ]
            temperature           = 4.92
            top_k                 = 299
            top_p                 = 0.68
            upstream_url          = "https://determined-skyscraper.name"
            workspace_instance_id = "...my_workspace_instance_id..."
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
      "audio/transcription"
    ]
    config = {
      balancer = {
        semantic = {
          connect_timeout = 60000
          embeddings = {
            allow_auth_override = false
            config = {
              ollama = {
                upstream_url = "...my_upstream_url..."
              }
            }
            name     = "...my_name..."
            provider = "azure-ai-se"
          }
          fail_timeout = 10000
          failover_criteria = [
            "http_429"
          ]
          max_fails    = 0
          read_timeout = 60000
          retries      = 5
          slots        = 10000
          vectordb = {
            pgvector = {
              database        = "kong-pgvector"
              dimensions      = 6
              distance_metric = "euclidean"
              host            = "127.0.0.1"
              password        = "...my_password..."
              port            = 5432
              ssl = {
                cert     = "...my_cert..."
                cert_key = "...my_cert_key..."
                enabled  = true
                required = true
                verify   = true
                version  = "tlsv1_2"
              }
              threshold = 3.66
              timeout   = 5000
              user      = "postgres"
            }
          }
          write_timeout = 60000
        }
      }
      logging = {
        payloads = false
      }
      max_request_body_size = 8388608
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
        name = "shared-redis"
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
          xai = {
            cache_read_cost = 2.26
            cache_read_cost_list = [
              {
                cost  = 1.04
                modal = "image"
              }
            ]
            cache_write_cost = 0.32
            cache_write_cost_list = [
              {
                cost = 3.9
                ttl  = "...my_ttl..."
              }
            ]
            context_window_factor = [
              {
                above         = "...my_above..."
                input_factor  = 2.92
                output_factor = 4.58
              }
            ]
            embeddings_dimensions = 748376912
            input_cost            = 3.63
            input_cost_list = [
              {
                cost  = 6.31
                modal = "audio"
              }
            ]
            max_tokens  = 1493087753
            output_cost = 1.58
            output_cost_list = [
              {
                cost  = 3.36
                modal = "audio"
              }
            ]
            service_tier_factor = [
              {
                factor = 2.95
                tier   = "...my_tier..."
              }
            ]
            temperature  = 0.46
            top_k        = 328
            top_p        = 0.12
            upstream_url = "https://carefree-hundred.org"
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