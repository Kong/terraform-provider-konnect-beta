resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider" {
  provider = konnect-beta
  anthropic = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 16787121
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
      credential = {
        inbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  azure = {
    config = {
      auth = {
        azure = {
          client_id            = "...my_client_id..."
          client_secret        = "...my_client_secret..."
          tenant_id            = "...my_tenant_id..."
          use_managed_identity = true
        }
      }
      foundry = {
        domain   = "services.ai.azure.com"
        resource = "kong-foundry-east"
      }
      instance = "kong-az-east"
      service  = "azure-openai"
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 742839646
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  baseten = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 2017813820
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  bedrock = {
    config = {
      auth = {
        aws = {
          access_key_id     = "...my_access_key_id..."
          assume_role_arn   = "...my_assume_role_arn..."
          batch_role_arn    = "...my_batch_role_arn..."
          role_session_name = "...my_role_session_name..."
          secret_access_key = "...my_secret_access_key..."
          session_token     = "...my_session_token..."
          sts_endpoint_url  = "...my_sts_endpoint_url..."
        }
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
        redis = {
          cloud_authentication = {
            aws = {
              access_key_id     = "...my_access_key_id..."
              assume_role_arn   = "...my_assume_role_arn..."
              cache_name        = "...my_cache_name..."
              is_serverless     = true
              region            = "...my_region..."
              role_session_name = "...my_role_session_name..."
              secret_access_key = "...my_secret_access_key..."
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
            backlog   = 373020684
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  cerebras = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 999176054
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  cohere = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
        redis = {
          cloud_authentication = {
            aws = {
              access_key_id     = "...my_access_key_id..."
              assume_role_arn   = "...my_assume_role_arn..."
              cache_name        = "...my_cache_name..."
              is_serverless     = true
              region            = "...my_region..."
              role_session_name = "...my_role_session_name..."
              secret_access_key = "...my_secret_access_key..."
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
            backlog   = 1110664052
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  dashscope = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1909370302
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
            role     = "any"
            username = "...my_username..."
          }
          server_name = "...my_server_name..."
          ssl         = true
          ssl_verify  = true
          username    = "...my_username..."
        }
      }
      credential = {
        inbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  databricks = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 851042800
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
      credential = {
        inbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  deepseek = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1991302957
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
      credential = {
        inbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  gemini = {
    config = {
      auth = {
        gcp = {
          metadata_url            = "...my_metadata_url..."
          oauth_token_url         = "...my_oauth_token_url..."
          service_account_json    = "...my_service_account_json..."
          use_gcp_service_account = false
          workload_identity_federation = {
            aws_iam = {
              auth_json = "...my_auth_json..."
              aws = {
                access_key_id     = "...my_access_key_id..."
                assume_role_arn   = "...my_assume_role_arn..."
                region            = "...my_region..."
                role_session_name = "...my_role_session_name..."
                secret_access_key = "...my_secret_access_key..."
                session_token     = "...my_session_token..."
                sts_endpoint_url  = "...my_sts_endpoint_url..."
              }
              source = "aws_iam"
            }
          }
        }
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 385349533
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
            role     = "any"
            username = "...my_username..."
          }
          server_name = "...my_server_name..."
          ssl         = true
          ssl_verify  = true
          username    = "...my_username..."
        }
      }
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  huggingface = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1459411688
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  kimi = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1520118691
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  llama2 = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1238703384
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
            role     = "any"
            username = "...my_username..."
          }
          server_name = "...my_server_name..."
          ssl         = true
          ssl_verify  = true
          username    = "...my_username..."
        }
      }
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  mistral = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "body"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
        redis = {
          cloud_authentication = {
            aws = {
              access_key_id     = "...my_access_key_id..."
              assume_role_arn   = "...my_assume_role_arn..."
              cache_name        = "...my_cache_name..."
              is_serverless     = true
              region            = "...my_region..."
              role_session_name = "...my_role_session_name..."
              secret_access_key = "...my_secret_access_key..."
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
            backlog   = 1126717458
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  ollama = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1891587515
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  openai = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 1132104571
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  sagemaker = {
    config = {
      auth = {
        basic = {
          headers = [
            {
              name  = "...my_name..."
              value = "...my_value..."
            }
          ]
          params = [
            {
              location = "query"
              name     = "...my_name..."
              value    = "...my_value..."
            }
          ]
        }
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
        redis = {
          cloud_authentication = {
            aws = {
              access_key_id     = "...my_access_key_id..."
              assume_role_arn   = "...my_assume_role_arn..."
              cache_name        = "...my_cache_name..."
              is_serverless     = true
              region            = "...my_region..."
              role_session_name = "...my_role_session_name..."
              secret_access_key = "...my_secret_access_key..."
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
            backlog   = 1991958544
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
      credential = {
        inbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  typesafe = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 912262254
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  vercel = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 375826408
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
      credential = {
        inbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "query"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "access_token"
    }
  }
  vllm = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 2119930107
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
            role     = "any"
            username = "...my_username..."
          }
          server_name = "...my_server_name..."
          ssl         = true
          ssl_verify  = true
          username    = "...my_username..."
        }
      }
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "body"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
  xai = {
    config = {
      auth = {
        headers = [
          {
            name  = "...my_name..."
            value = "...my_value..."
          }
        ]
        params = [
          {
            location = "query"
            name     = "...my_name..."
            value    = "...my_value..."
          }
        ]
      }
    }
    datastores = [
      {
        name = "shared-redis"
      }
    ]
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    token_vault = {
      cache = {
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
            backlog   = 641476603
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
      credential = {
        inbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
        outbound = {
          location = "header"
          name     = "...my_name..."
          prefix   = "...my_prefix..."
        }
      }
      provider           = "...my_provider..."
      subject_token_type = "api_key"
    }
  }
}