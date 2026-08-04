
resource "konnect_ai_gateway_model" "my_aigatewaymodel" {
  provider = konnect-beta
  api = {
    capabilities = [
      "batches"
    ]
    config = {
      route = {
        hosts = []
        paths = [
          "/my-test-path"
        ]
        https_redirect_status_code = 426
        preserve_host = false
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
      }
    }
    display_name = "My Test Claude 5 model"
    enabled      = true
    formats = [
      {
        type = "anthropic"
      }
    ]
    name = "tf-test-claude-5-model"
    targets = [
      {
        config = {
          anthropic = {
            embeddings_dimensions = 7
            input_cost            = 3.7
            max_tokens            = 6
            output_cost           = 6.56
            temperature           = 3.27
            top_k                 = 2
            top_p                 = 0.5
            upstream_url          = "https://baggy-trash.biz/"
            deployment_id = "claude-5-model"
          }
        }
        name                 = "claude-5-model"
        provider             = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider_anthropic.name
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}

resource "konnect_ai_gateway_model" "my_aigatewaymodel_model" {
  provider = konnect-beta
  model = {
    capabilities = [
      "generate"
    ]
    config = {
      route = {
        hosts = []
        model = {
            path_selector = {
              path_aliases = [
                "my model model"
              ]
            }
        }
        paths = [
          "/my-base-path"
        ]
        protocols = [
          "https"
        ]
        https_redirect_status_code = 426
        preserve_host = false
        regex_priority     = 0
        request_buffering  = true
        response_buffering = true
        strip_path         = true
      }
    }
    display_name = "My Test Azure model"
    enabled      = true
    formats = [
      {
        type = "openai"
      }
    ]
    name = "my-azure-model"
    targets = [
      {
        config = {
          azure = {
            api_version = "2023-05-15"
            deployment_id = "ahagshhh-1sjn-akjnda"
            type = "azure"
          }
        }
        name                 = "gpt-5.6-luna",
        provider             = konnect_ai_gateway_model_provider.my_aigatewaymodelprovider.name
      }
    ]
  }
  gateway_id = konnect_ai_gateway.my_aigateway.id
}