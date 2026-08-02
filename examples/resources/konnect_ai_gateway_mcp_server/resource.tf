resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver" {
  provider = konnect-beta
  conversion_listener = {
    access = {
      oauth_access_token = {
        access_token_claim_field = "...my_access_token_claim_field..."
        acl_attribute_type       = "oauth_access_token"
        acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
        default_tool_acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
      }
    }
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388608
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
      server = {
        forward_client_headers = true
        label                  = "...my_label..."
        session = {
          client = {
            secrets = [
              "..."
            ]
          }
          managed = true
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
              backlog   = 1275755412
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
          session_ttl = 86400
          strategy    = "redis"
        }
        timeout = 10000
      }
      url = "https://mcp.internal.kongair.com"
    }
    display_name = "Kong Air Flights"
    enabled      = true
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "kongair-flights"
    policies = [
      "..."
    ]
    tools = [
      {
        access = {
          acls = {
            allow = [
              "..."
            ]
            deny = [
              "..."
            ]
          }
        }
        annotations = {
          destructive_hint = true
          idempotent_hint  = false
          open_world_hint  = false
          read_only_hint   = true
          title            = "...my_title..."
        }
        description = "Search for available flights"
        headers     = "{ \"see\": \"documentation\" }"
        host        = "...my_host..."
        method      = "POST"
        name        = "...my_name..."
        parameters = [
          {
            description = "The origin airport code."
            in          = "query"
            name        = "origin"
            required    = true
            schema = {
              key = jsonencode("value")
            }
          }
        ]
        path         = "...my_path..."
        query        = "{ \"see\": \"documentation\" }"
        request_body = "{ \"see\": \"documentation\" }"
        responses    = "{ \"see\": \"documentation\" }"
        scheme       = "https"
      }
    ]
  }
  conversion_only = {
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388608
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
      url = "https://mcp.internal.kongair.com"
    }
    display_name = "Kong Air Flights"
    enabled      = true
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "kongair-flights"
    policies = [
      "..."
    ]
    tools = [
      {
        access = {
          acls = {
            allow = [
              "..."
            ]
            deny = [
              "..."
            ]
          }
        }
        annotations = {
          destructive_hint = false
          idempotent_hint  = true
          open_world_hint  = true
          read_only_hint   = false
          title            = "...my_title..."
        }
        description = "Search for available flights"
        headers     = "{ \"see\": \"documentation\" }"
        host        = "...my_host..."
        method      = "PATCH"
        name        = "...my_name..."
        parameters = [
          {
            description = "The origin airport code."
            in          = "query"
            name        = "origin"
            required    = true
            schema = {
              key = jsonencode("value")
            }
          }
        ]
        path         = "...my_path..."
        query        = "{ \"see\": \"documentation\" }"
        request_body = "{ \"see\": \"documentation\" }"
        responses    = "{ \"see\": \"documentation\" }"
        scheme       = "http"
      }
    ]
  }
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  listener = {
    access = {
      oauth_access_token = {
        access_token_claim_field = "...my_access_token_claim_field..."
        acl_attribute_type       = "oauth_access_token"
        acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
        default_tool_acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
      }
    }
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388608
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
      server = {
        forward_client_headers = true
        label                  = "...my_label..."
        session = {
          client = {
            secrets = [
              "..."
            ]
          }
          managed = true
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
              backlog   = 254844406
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
          session_ttl = 86400
          strategy    = "redis"
        }
        timeout = 10000
      }
    }
    display_name = "Kong Air Flights"
    enabled      = true
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "kongair-flights"
    policies = [
      "..."
    ]
    tools = [
      {
        access = {
          acls = {
            allow = [
              "..."
            ]
            deny = [
              "..."
            ]
          }
        }
        annotations = {
          destructive_hint = true
          idempotent_hint  = false
          open_world_hint  = false
          read_only_hint   = true
          title            = "...my_title..."
        }
        description = "Search for available flights"
        headers     = "{ \"see\": \"documentation\" }"
        host        = "...my_host..."
        method      = "PUT"
        name        = "...my_name..."
        parameters = [
          {
            description = "The origin airport code."
            in          = "query"
            name        = "origin"
            required    = true
            schema = {
              key = jsonencode("value")
            }
          }
        ]
        path         = "...my_path..."
        query        = "{ \"see\": \"documentation\" }"
        request_body = "{ \"see\": \"documentation\" }"
        responses    = "{ \"see\": \"documentation\" }"
        scheme       = "https"
      }
    ]
  }
  passthrough_listener = {
    access = {
      consumer = {
        acl_attribute_type = "consumer"
        acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
        default_tool_acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
      }
    }
    config = {
      logging = {
        audits   = false
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
          port = 43233
        }
        https_proxy = {
          host = "...my_host..."
          port = 31216
        }
        no_proxy     = "...my_no_proxy..."
        proxy_scheme = "http"
      }
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
      server = {
        forward_client_headers = true
        label                  = "...my_label..."
        session = {
          client = {
            secrets = [
              "..."
            ]
          }
          managed = true
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
              backlog   = 1750673053
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
          session_ttl = 86400
          strategy    = "redis"
        }
        timeout = 10000
      }
      url = "https://mcp.internal.kongair.com"
    }
    display_name = "Kong Air Flights"
    enabled      = true
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "kongair-flights"
    policies = [
      "..."
    ]
    tools = [
      {
        access = {
          acls = {
            allow = [
              "..."
            ]
            deny = [
              "..."
            ]
          }
        }
        annotations = {
          destructive_hint = true
          idempotent_hint  = false
          open_world_hint  = false
          read_only_hint   = true
          title            = "...my_title..."
        }
        description = "Search for available flights"
        headers     = "{ \"see\": \"documentation\" }"
        host        = "...my_host..."
        method      = "PUT"
        name        = "...my_name..."
        parameters = [
          {
            description = "The origin airport code."
            in          = "query"
            name        = "origin"
            required    = true
            schema = {
              key = jsonencode("value")
            }
          }
        ]
        path         = "...my_path..."
        query        = "{ \"see\": \"documentation\" }"
        request_body = "{ \"see\": \"documentation\" }"
        responses    = "{ \"see\": \"documentation\" }"
        scheme       = "http"
      }
    ]
  }
  upstream_server = {
    access = {
      consumer = {
        acl_attribute_type = "consumer"
        acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
        default_tool_acls = {
          allow = [
            "..."
          ]
          deny = [
            "..."
          ]
        }
      }
    }
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388608
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
      server = {
        forward_client_headers       = true
        label                        = "...my_label..."
        preserve_upstream_tool_names = false
        session = {
          client = {
            secrets = [
              "..."
            ]
          }
          managed = true
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
              backlog   = 176733398
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
          session_ttl = 86400
          strategy    = "redis"
        }
        timeout = 10000
        tools_list_auth = {
          jwt = {
            access_token_header = "Authorization"
            id_token_header     = "...my_id_token_header..."
            scope               = "...my_scope..."
          }
        }
      }
      tools_cache_ttl_seconds = 7
      url                     = "https://mcp.internal.kongair.com"
    }
    display_name = "Kong Air Flights"
    enabled      = true
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "kongair-flights"
    policies = [
      "..."
    ]
    tools = [
      {
        access = {
          acls = {
            allow = [
              "..."
            ]
            deny = [
              "..."
            ]
          }
        }
        annotations = {
          destructive_hint = true
          idempotent_hint  = true
          open_world_hint  = false
          read_only_hint   = true
          title            = "...my_title..."
        }
        description   = "Search for available flights"
        headers       = "{ \"see\": \"documentation\" }"
        host          = "...my_host..."
        input_schema  = "{ \"see\": \"documentation\" }"
        method        = "PUT"
        name          = "...my_name..."
        output_schema = "{ \"see\": \"documentation\" }"
        parameters = [
          {
            description = "The origin airport code."
            in          = "query"
            name        = "origin"
            required    = true
            schema = {
              key = jsonencode("value")
            }
          }
        ]
        path         = "...my_path..."
        query        = "{ \"see\": \"documentation\" }"
        request_body = "{ \"see\": \"documentation\" }"
        responses    = "{ \"see\": \"documentation\" }"
        scheme       = "http"
      }
    ]
  }
}