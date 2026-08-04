resource "konnect_ai_gateway_identity_provider" "my_aigatewayidentityprovider" {
  provider = konnect-beta
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  key_auth = {
    config = {
      anonymous        = "...my_anonymous..."
      hide_credentials = true
      identity_realms = [
        {
          id     = "...my_id..."
          region = "...my_region..."
          scope  = "realm"
        }
      ]
      key_in_body   = false
      key_in_header = true
      key_in_query  = true
      key_names = [
        "..."
      ]
      principals = {
        directory     = "...my_directory..."
        enabled       = true
        error_on_miss = true
      }
      realm            = "...my_realm..."
      run_on_preflight = false
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
  }
  openid_connect = {
    config = {
      anonymous = "...my_anonymous..."
      audience = [
        "..."
      ]
      audience_claim = [
        "..."
      ]
      audience_required = [
        "..."
      ]
      auth_methods = [
        "kong_oauth2"
      ]
      authenticated_groups_claim = [
        "..."
      ]
      authorization_cookie_domain    = "...my_authorization_cookie_domain..."
      authorization_cookie_http_only = true
      authorization_cookie_name      = "...my_authorization_cookie_name..."
      authorization_cookie_path      = "...my_authorization_cookie_path..."
      authorization_cookie_same_site = "Lax"
      authorization_cookie_secure    = false
      authorization_endpoint         = "...my_authorization_endpoint..."
      authorization_query_args_client = [
        "..."
      ]
      authorization_query_args_names = [
        "..."
      ]
      authorization_query_args_values = [
        "..."
      ]
      authorization_rolling_timeout = 4.95
      bearer_token_cookie_name      = "...my_bearer_token_cookie_name..."
      bearer_token_header_name      = "...my_bearer_token_header_name..."
      bearer_token_param_type = [
        "body"
      ]
      by_username_ignore_case = false
      cache_introspection     = true
      cache_token_exchange    = false
      cache_tokens            = false
      cache_tokens_salt       = "...my_cache_tokens_salt..."
      cache_ttl               = 0.89
      cache_ttl_max           = 7.18
      cache_ttl_min           = 0.73
      cache_ttl_neg           = 6.44
      cache_ttl_resurrect     = 6.44
      cache_user_info         = true
      claims_forbidden = [
        "..."
      ]
      client_alg = [
        "HS256"
      ]
      client_arg = "...my_client_arg..."
      client_auth = [
        "self_signed_tls_client_auth"
      ]
      client_credentials_param_type = [
        "header"
      ]
      client_id = [
        "..."
      ]
      client_jwk = [
        {
          alg    = "...my_alg..."
          crv    = "...my_crv..."
          d      = "...my_d..."
          dp     = "...my_dp..."
          dq     = "...my_dq..."
          e      = "...my_e..."
          issuer = "...my_issuer..."
          k      = "...my_k..."
          key_ops = [
            "..."
          ]
          kid = "...my_kid..."
          kty = "...my_kty..."
          n   = "...my_n..."
          oth = "...my_oth..."
          p   = "...my_p..."
          q   = "...my_q..."
          qi  = "...my_qi..."
          r   = "...my_r..."
          t   = "...my_t..."
          use = "...my_use..."
          x   = "...my_x..."
          x5c = [
            "..."
          ]
          x5t             = "...my_x5t..."
          x5t_number_s256 = "...my_x5t_number_s256..."
          x5u             = "...my_x5u..."
          y               = "...my_y..."
        }
      ]
      client_secret = [
        "..."
      ]
      cluster_cache_items = [
        "tokens"
      ]
      cluster_cache_redis = {
        cloud_authentication = {
          auth_provider            = "aws"
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
        cluster_max_redirections = 8
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 23504
          }
        ]
        connect_timeout       = 389588841
        connection_is_proxied = true
        database              = 5
        host                  = "...my_host..."
        keepalive_backlog     = 23143264
        keepalive_pool_size   = 522008638
        password              = "...my_password..."
        port                  = "...my_port..."
        read_timeout          = 2070681451
        send_timeout          = 903880050
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 32196
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "slave"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        ssl               = true
        ssl_verify        = false
        username          = "...my_username..."
      }
      cluster_cache_strategy = "redis"
      consumer_by = [
        "username"
      ]
      consumer_claims = [
        [
          # ...
        ]
      ]
      consumer_groups_claim = [
        "..."
      ]
      consumer_groups_optional = false
      consumer_optional        = true
      credential_claim = [
        "..."
      ]
      disable_session = [
        "refresh_token"
      ]
      discovery_headers_names = [
        "..."
      ]
      discovery_headers_values = [
        "..."
      ]
      display_errors = true
      domains = [
        "..."
      ]
      downstream_access_token_header     = "...my_downstream_access_token_header..."
      downstream_access_token_jwk_header = "...my_downstream_access_token_jwk_header..."
      downstream_headers = [
        {
          header = "...my_header..."
          path = [
            "..."
          ]
        }
      ]
      downstream_headers_claims = [
        "..."
      ]
      downstream_headers_names = [
        "..."
      ]
      downstream_id_token_header          = "...my_downstream_id_token_header..."
      downstream_id_token_jwk_header      = "...my_downstream_id_token_jwk_header..."
      downstream_introspection_header     = "...my_downstream_introspection_header..."
      downstream_introspection_jwt_header = "...my_downstream_introspection_jwt_header..."
      downstream_refresh_token_header     = "...my_downstream_refresh_token_header..."
      downstream_session_id_header        = "...my_downstream_session_id_header..."
      downstream_user_info_header         = "...my_downstream_user_info_header..."
      downstream_user_info_jwt_header     = "...my_downstream_user_info_jwt_header..."
      dpop_proof_lifetime                 = 8.5
      dpop_use_nonce                      = false
      enable_hs_signatures                = false
      end_session_endpoint                = "...my_end_session_endpoint..."
      expose_error_code                   = true
      extra_jwks_uris = [
        "..."
      ]
      forbidden_destroy_session = false
      forbidden_error_message   = "...my_forbidden_error_message..."
      forbidden_redirect_uri = [
        "..."
      ]
      groups_claim = [
        "..."
      ]
      groups_required = [
        "..."
      ]
      hide_credentials          = true
      http_proxy                = "...my_http_proxy..."
      http_proxy_authorization  = "...my_http_proxy_authorization..."
      http_version              = 1.21
      https_proxy               = "...my_https_proxy..."
      https_proxy_authorization = "...my_https_proxy_authorization..."
      id_token_param_name       = "...my_id_token_param_name..."
      id_token_param_type = [
        "query"
      ]
      ignore_signature = [
        "userinfo"
      ]
      introspect_jwt_tokens              = false
      introspection_accept               = "application/json"
      introspection_check_active         = true
      introspection_endpoint             = "...my_introspection_endpoint..."
      introspection_endpoint_auth_method = "tls_client_auth"
      introspection_headers_client = [
        "..."
      ]
      introspection_headers_names = [
        "..."
      ]
      introspection_headers_values = [
        "..."
      ]
      introspection_hint = "...my_introspection_hint..."
      introspection_post_args_client = [
        "..."
      ]
      introspection_post_args_client_headers = [
        "..."
      ]
      introspection_post_args_names = [
        "..."
      ]
      introspection_post_args_values = [
        "..."
      ]
      introspection_token_param_name = "...my_introspection_token_param_name..."
      issuer                         = "...my_issuer..."
      issuers_allowed = [
        "..."
      ]
      jwks_endpoint      = "...my_jwks_endpoint..."
      jwt_session_claim  = "...my_jwt_session_claim..."
      jwt_session_cookie = "...my_jwt_session_cookie..."
      keepalive          = false
      leeway             = 8.81
      login_action       = "response"
      login_methods = [
        "client_credentials"
      ]
      login_redirect_mode = "query"
      login_redirect_uri = [
        "..."
      ]
      login_tokens = [
        "refresh_token"
      ]
      logout_methods = [
        "GET"
      ]
      logout_post_arg  = "...my_logout_post_arg..."
      logout_query_arg = "...my_logout_query_arg..."
      logout_redirect_uri = [
        "..."
      ]
      logout_revoke               = true
      logout_revoke_access_token  = true
      logout_revoke_refresh_token = false
      logout_uri_suffix           = "...my_logout_uri_suffix..."
      max_age                     = 9.27
      mtls_introspection_endpoint = "...my_mtls_introspection_endpoint..."
      mtls_revocation_endpoint    = "...my_mtls_revocation_endpoint..."
      mtls_token_endpoint         = "...my_mtls_token_endpoint..."
      no_proxy                    = "...my_no_proxy..."
      password_param_type = [
        "query"
      ]
      preserve_query_args = false
      principals = {
        directory             = "...my_directory..."
        enabled               = true
        error_on_miss         = false
        match_consumer        = false
        match_consumer_groups = true
        principal_by          = "...my_principal_by..."
        principal_claim = [
          "..."
        ]
      }
      proof_of_possession_auth_methods_validation = true
      proof_of_possession_dpop                    = "off"
      proof_of_possession_mtls                    = "off"
      proof_of_possession_mtls_from_header = {
        allow_partial_chain = false
        ca_certificates = [
          "..."
        ]
        cert_cache_ttl            = 2.81
        certificate_header_format = "url_encoded"
        certificate_header_name   = "...my_certificate_header_name..."
        http_proxy_host           = "...my_http_proxy_host..."
        http_proxy_port           = 358
        http_timeout              = 3.21
        https_proxy_host          = "...my_https_proxy_host..."
        https_proxy_port          = 26146
        revocation_check_mode     = "IGNORE_CA_ERROR"
        secure_source             = true
        ssl_verify                = true
      }
      pushed_authorization_request_endpoint             = "...my_pushed_authorization_request_endpoint..."
      pushed_authorization_request_endpoint_auth_method = "client_secret_post"
      redirect_uri = [
        "..."
      ]
      redis = {
        cloud_authentication = {
          auth_provider            = "aws"
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
        cluster_max_redirections = 6
        cluster_nodes = [
          {
            ip   = "...my_ip..."
            port = 22648
          }
        ]
        connect_timeout       = 1281691808
        connection_is_proxied = false
        database              = 3
        host                  = "...my_host..."
        keepalive_backlog     = 1535934402
        keepalive_pool_size   = 281030404
        password              = "...my_password..."
        port                  = "...my_port..."
        prefix                = "...my_prefix..."
        read_timeout          = 1033518283
        send_timeout          = 1612129646
        sentinel_master       = "...my_sentinel_master..."
        sentinel_nodes = [
          {
            host = "...my_host..."
            port = 24150
          }
        ]
        sentinel_password = "...my_sentinel_password..."
        sentinel_role     = "any"
        sentinel_username = "...my_sentinel_username..."
        server_name       = "...my_server_name..."
        socket            = "...my_socket..."
        ssl               = false
        ssl_verify        = true
        username          = "...my_username..."
      }
      rediscovery_lifetime     = 0.31
      refresh_token_param_name = "...my_refresh_token_param_name..."
      refresh_token_param_type = [
        "header"
      ]
      refresh_tokens                        = true
      require_proof_key_for_code_exchange   = false
      require_pushed_authorization_requests = false
      require_signed_request_object         = false
      resolve_distributed_claims            = true
      response_mode                         = "fragment"
      response_type = [
        "..."
      ]
      reverify                        = true
      revocation_endpoint             = "...my_revocation_endpoint..."
      revocation_endpoint_auth_method = "private_key_jwt"
      revocation_token_param_name     = "...my_revocation_token_param_name..."
      roles_claim = [
        "..."
      ]
      roles_required = [
        "..."
      ]
      run_on_preflight = false
      scopes = [
        "..."
      ]
      scopes_claim = [
        "..."
      ]
      scopes_required = [
        "..."
      ]
      search_user_info         = false
      session_absolute_timeout = 9.21
      session_audience         = "...my_session_audience..."
      session_bind = [
        "user-agent"
      ]
      session_cookie_domain             = "...my_session_cookie_domain..."
      session_cookie_http_only          = true
      session_cookie_name               = "...my_session_cookie_name..."
      session_cookie_path               = "...my_session_cookie_path..."
      session_cookie_same_site          = "Default"
      session_cookie_secure             = true
      session_enforce_same_subject      = false
      session_hash_storage_key          = false
      session_hash_subject              = false
      session_idling_timeout            = 7.05
      session_memcached_host            = "...my_session_memcached_host..."
      session_memcached_port            = 1687
      session_memcached_prefix          = "...my_session_memcached_prefix..."
      session_memcached_socket          = "...my_session_memcached_socket..."
      session_memcached_ssl             = true
      session_memcached_ssl_verify      = false
      session_remember                  = true
      session_remember_absolute_timeout = 9.29
      session_remember_cookie_name      = "...my_session_remember_cookie_name..."
      session_remember_rolling_timeout  = 7.44
      session_request_headers = [
        "audience"
      ]
      session_response_headers = [
        "idling-timeout"
      ]
      session_rolling_timeout       = 3.66
      session_secret                = "...my_session_secret..."
      session_storage               = "memcached"
      session_store_metadata        = true
      ssl_verify                    = true
      timeout                       = 1.86
      tls_client_auth_cert_id       = "...my_tls_client_auth_cert_id..."
      tls_client_auth_ssl_verify    = false
      token_cache_key_include_scope = true
      token_endpoint                = "...my_token_endpoint..."
      token_endpoint_auth_method    = "private_key_jwt"
      token_exchange = {
        cache = {
          enabled = false
          ttl     = 1
        }
        request = {
          audience = [
            "..."
          ]
          empty_audience = true
          empty_scopes   = false
          scopes = [
            "..."
          ]
        }
        subject_token_issuers = [
          {
            conditions = {
              has_audience = [
                "..."
              ]
              has_scopes = [
                "..."
              ]
              missing_audience = [
                "..."
              ]
              missing_scopes = [
                "..."
              ]
            }
            issuer           = "...my_issuer..."
            jwks_uri         = "...my_jwks_uri..."
            verify_signature = true
          }
        ]
      }
      token_exchange_endpoint = "...my_token_exchange_endpoint..."
      token_headers_client = [
        "..."
      ]
      token_headers_grants = [
        "password"
      ]
      token_headers_names = [
        "..."
      ]
      token_headers_prefix = "...my_token_headers_prefix..."
      token_headers_replay = [
        "..."
      ]
      token_headers_values = [
        "..."
      ]
      token_post_args_client = [
        "..."
      ]
      token_post_args_names = [
        "..."
      ]
      token_post_args_values = [
        "..."
      ]
      unauthorized_destroy_session = false
      unauthorized_error_message   = "...my_unauthorized_error_message..."
      unauthorized_redirect_uri = [
        "..."
      ]
      unexpected_redirect_uri = [
        "..."
      ]
      upstream_access_token_header     = "...my_upstream_access_token_header..."
      upstream_access_token_jwk_header = "...my_upstream_access_token_jwk_header..."
      upstream_headers = [
        {
          header = "...my_header..."
          path = [
            "..."
          ]
        }
      ]
      upstream_headers_claims = [
        "..."
      ]
      upstream_headers_names = [
        "..."
      ]
      upstream_id_token_header          = "...my_upstream_id_token_header..."
      upstream_id_token_jwk_header      = "...my_upstream_id_token_jwk_header..."
      upstream_introspection_header     = "...my_upstream_introspection_header..."
      upstream_introspection_jwt_header = "...my_upstream_introspection_jwt_header..."
      upstream_refresh_token_header     = "...my_upstream_refresh_token_header..."
      upstream_session_id_header        = "...my_upstream_session_id_header..."
      upstream_user_info_header         = "...my_upstream_user_info_header..."
      upstream_user_info_jwt_header     = "...my_upstream_user_info_jwt_header..."
      userinfo_accept                   = "application/jwt"
      userinfo_endpoint                 = "...my_userinfo_endpoint..."
      userinfo_headers_client = [
        "..."
      ]
      userinfo_headers_names = [
        "..."
      ]
      userinfo_headers_values = [
        "..."
      ]
      userinfo_query_args_client = [
        "..."
      ]
      userinfo_query_args_names = [
        "..."
      ]
      userinfo_query_args_values = [
        "..."
      ]
      using_pseudo_issuer = true
      verify_claims       = true
      verify_nonce        = false
      verify_parameters   = true
      verify_signature    = false
    }
    display_name = "Okta AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "okta-ai-se"
  }
}