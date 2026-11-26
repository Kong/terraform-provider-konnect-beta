resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore" {
  provider = konnect-beta
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  redis_ce = {
    config = {
      cloud_authentication = {
        auth_provider = {
          str = "{vault://hcv/redis/auth_provider}"
        }
        aws_access_key_id        = "...my_aws_access_key_id..."
        aws_assume_role_arn      = "...my_aws_assume_role_arn..."
        aws_cache_name           = "...my_aws_cache_name..."
        aws_is_serverless        = true
        aws_region               = "...my_aws_region..."
        aws_role_session_name    = "...my_aws_role_session_name..."
        aws_secret_access_key    = "...my_aws_secret_access_key..."
        azure_client_id          = "...my_azure_client_id..."
        azure_client_secret      = "...my_azure_client_secret..."
        azure_tenant_id          = "...my_azure_tenant_id..."
        gcp_service_account_json = "...my_gcp_service_account_json..."
      }
      database    = 1
      host        = "127.0.0.1"
      password    = "...my_password..."
      port        = 60289
      server_name = "...my_server_name..."
      ssl         = true
      ssl_verify  = false
      timeout     = 1
      username    = "...my_username..."
    }
    description  = "Redis shared by rate limiting and semantic cache"
    display_name = "Shared Redis"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "shared-redis"
  }
  redis_ee = {
    config = {
      cloud_authentication = {
        auth_provider = {
          str = "{vault://hcv/redis/auth_provider}"
        }
        aws_access_key_id        = "...my_aws_access_key_id..."
        aws_assume_role_arn      = "...my_aws_assume_role_arn..."
        aws_cache_name           = "...my_aws_cache_name..."
        aws_is_serverless        = false
        aws_region               = "...my_aws_region..."
        aws_role_session_name    = "...my_aws_role_session_name..."
        aws_secret_access_key    = "...my_aws_secret_access_key..."
        azure_client_id          = "...my_azure_client_id..."
        azure_client_secret      = "...my_azure_client_secret..."
        azure_tenant_id          = "...my_azure_tenant_id..."
        gcp_service_account_json = "...my_gcp_service_account_json..."
      }
      cluster_max_redirections = 6
      cluster_nodes = [
        {
          ip   = "127.0.0.1"
          port = 6379
        }
      ]
      connect_timeout       = 9
      connection_is_proxied = true
      database              = 6
      host                  = "127.0.0.1"
      keepalive_backlog     = 5
      keepalive_pool_size   = 0
      password              = "...my_password..."
      port = {
        str = "{vault://hcv/redis/port}"
      }
      read_timeout    = 10
      send_timeout    = 2
      sentinel_master = "...my_sentinel_master..."
      sentinel_nodes = [
        {
          host = "127.0.0.1"
          port = 6379
        }
      ]
      sentinel_password = "...my_sentinel_password..."
      sentinel_role     = "master"
      sentinel_username = "...my_sentinel_username..."
      server_name       = "...my_server_name..."
      ssl               = false
      ssl_verify        = false
      username          = "...my_username..."
    }
    description  = "Redis shared by rate limiting and semantic cache"
    display_name = "Shared Redis"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "shared-redis"
  }
  token_vault = {
    config = {
      current_version = "v1"
      directory       = "my-directory"
      keys = [
        {
          key     = "{vault://hcv/vault-enc-key}"
          version = "v1"
        }
      ]
    }
    description  = "Redis shared by rate limiting and semantic cache"
    display_name = "Shared Redis"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "shared-redis"
  }
  vectordb = {
    config = {
      database     = "kong-pgvector"
      host         = "127.0.0.1"
      password     = "...my_password..."
      port         = 5432
      ssl          = false
      ssl_cert     = "...my_ssl_cert..."
      ssl_cert_key = "...my_ssl_cert_key..."
      ssl_required = false
      ssl_verify   = true
      ssl_version  = "tlsv1_2"
      timeout      = 5000
      user         = "postgres"
    }
    description  = "Redis shared by rate limiting and semantic cache"
    display_name = "Shared Redis"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "shared-redis"
  }
}