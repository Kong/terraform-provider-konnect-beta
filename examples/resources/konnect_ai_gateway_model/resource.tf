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
          algorithm       = "lowest-latency"
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
          ai_gateway_model_alias_config_path = {
            path_aliases = [
              "@azure/claude-sonnet-5"
            ]
          }
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
          xai = {
            embeddings_dimensions = 7
            input_cost            = 3.7
            max_tokens            = 6
            output_cost           = 6.56
            temperature           = 3.27
            top_k                 = 2
            top_p                 = 2.83
            type                  = "xai"
            upstream_url          = "https://baggy-trash.biz/"
          }
        }
        name                 = "gpt-5-model"
        provider             = "azure-ai-se"
        semantic_description = "...my_semantic_description..."
        weight               = 100
      }
    ]
    type = "api"
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
          algorithm       = "semantic"
          connect_timeout = 60000
          embeddings = {
            allow_auth_override = false
            config = {
              huggingface = {
                type           = "huggingface"
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
              type      = "pgvector"
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
          ai_gateway_model_alias_config_path = {
            path_aliases = [
              "@azure/claude-sonnet-5"
            ]
          }
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
          databricks = {
            embeddings_dimensions = 6
            input_cost            = 9.06
            max_tokens            = 8
            output_cost           = 7.78
            temperature           = 3.33
            top_k                 = 4
            top_p                 = 7.55
            type                  = "databricks"
            upstream_url          = "https://distant-antelope.com"
            workspace_instance_id = "...my_workspace_instance_id..."
          }
        }
        name                 = "gpt-5-model"
        provider             = "azure-ai-se"
        semantic_description = "...my_semantic_description..."
        weight               = 100
      }
    ]
    type = "model"
  }
}