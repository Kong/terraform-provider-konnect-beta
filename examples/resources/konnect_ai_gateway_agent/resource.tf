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