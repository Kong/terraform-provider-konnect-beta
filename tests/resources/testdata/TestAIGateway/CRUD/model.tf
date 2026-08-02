resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  provider = konnect-beta
  api = {
    capabilities = [
      "files"
    ]
    config = {
      logging = {
        payloads = false
      }
      route = {
        headers = {
          key = jsonencode("value")
        }
        hosts = [
          "foo.example.com"
        ]
        https_redirect_status_code = 426
        model = {
          ai_gateway_model_alias_config_path = {
            path_aliases = [
              "@azure/claude-sonnet-5"
            ]
          }
        }
        preserve_host = false
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
      }
    }
    display_name = "My GPT 5 model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "my-gpt-5-model"
    targets = [
      {
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
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}