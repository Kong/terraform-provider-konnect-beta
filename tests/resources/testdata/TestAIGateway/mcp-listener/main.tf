resource "konnect_ai_gateway" "my_aigateway" {
  provider     = konnect-beta
  display_name = "TF Test AIGW - mcp-listener"
  name         = "tf-test-aigw-mcp-listener"
}

resource "konnect_ai_gateway_mcp_server" "my_aigatewaymcpserver_listener" {
  provider = konnect-beta
  listener = {
    access = {
      consumer = {
        allow = [
          "my-test-allow-consumer"
        ]
      }
    }
    config = {
      logging = {
        audits   = false
        payloads = false
      }
      max_request_body_size = 8388
      route = {
        hosts = [
          "foo.example.com"
        ]
        https_redirect_status_code = 426
        regex_priority             = 0
        request_buffering          = true
        response_buffering         = true
        strip_path                 = true
      }
    }
    display_name = "TF Test MCP Listener"
    enabled      = true
    name         = "tf-test-mcp-listener"
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}
