resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  provider = konnect-beta
  api = {
    access = {
      acls = {
        allow = [
          "..."
        ]
        deny = [
          "..."
        ]
      }
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
            "http_429"
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
          anthropic = {
            cache_read_cost  = 8.86
            cache_write_cost = 7.72
            cache_write_cost_list = [
              {
                cost = 1.85
                ttl  = "...my_ttl..."
              }
            ]
            context_window_factor = [
              {
                above         = "...my_above..."
                input_factor  = 7.89
                output_factor = 6.2
              }
            ]
            embeddings_dimensions = 3
            input_cost            = 9.85
            max_tokens            = 1
            output_cost           = 1.7
            service_tier_factor = [
              {
                factor = 7.15
                tier   = "...my_tier..."
              }
            ]
            temperature  = 6.58
            top_k        = 3
            top_p        = 4.84
            upstream_url = "https://ajar-summer.biz"
            version      = "2023-06-01"
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
        deny = [
          "..."
        ]
      }
      identity_providers = [
        "okta-ai-se"
      ]
    }
    capabilities = [
      "audio/speech"
    ]
    config = {
      balancer = {
        semantic = {
          connect_timeout = 60000
          embeddings = {
            allow_auth_override = false
            config = {
              huggingface = {
                upstream_url   = "...my_upstream_url..."
                use_cache      = false
                wait_for_model = false
              }
            }
            name     = "...my_name..."
            provider = "azure-ai-se"
          }
          fail_timeout = 10000
          failover_criteria = [
            "http_502"
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
          mistral = {
            cache_read_cost  = 9.91
            cache_write_cost = 1.14
            cache_write_cost_list = [
              {
                cost = 5.09
                ttl  = "...my_ttl..."
              }
            ]
            context_window_factor = [
              {
                above         = "...my_above..."
                input_factor  = 1.3
                output_factor = 9.49
              }
            ]
            embeddings_dimensions = 10
            format                = "ollama"
            input_cost            = 1.29
            max_tokens            = 5
            output_cost           = 5.4
            service_tier_factor = [
              {
                factor = 5.81
                tier   = "...my_tier..."
              }
            ]
            temperature  = 4.9
            top_k        = 6
            top_p        = 5.29
            upstream_url = "https://sad-thigh.net"
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