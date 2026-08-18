resource "konnect_ai_gateway_agent" "my_aigatewayagent" {
  provider = konnect-beta
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
  config = {
    logging = {
      max_payload_size = 524288
      payloads         = false
    }
    max_request_body_size = 8388608
    proxy = {
      auth = {
        password = "...my_password..."
        username = "...my_username..."
      }
      http_proxy = {
        host = "...my_host..."
        port = 25961
      }
      https_proxy = {
        host = "...my_host..."
        port = 62168
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
    upstream = {
      auth = {
        aws = {
          access_key_id     = "...my_access_key_id..."
          assume_role_arn   = "...my_assume_role_arn..."
          region            = "...my_region..."
          role_session_name = "...my_role_session_name..."
          secret_access_key = "...my_secret_access_key..."
          session_token     = "...my_session_token..."
          sts_endpoint_url  = "...my_sts_endpoint_url..."
        }
      }
    }
    url = "https://booking-agent.internal.kongair.com"
  }
  display_name = "Kong Air Flight Booking Agent"
  enabled      = true
  gateway_id   = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name = "kongair-flight-booking-agent"
  policies = [
    "..."
  ]
  type = "a2a"
}