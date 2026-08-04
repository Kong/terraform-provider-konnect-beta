resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider" {
  provider = konnect-beta
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  azure = {
    config = {
      auth = {
        azure = {
          client_id            = "test-client-id"
          client_secret        = "test-client-id"
          tenant_id            = "test-tenant-id"
          type                 = "azure"
          use_managed_identity = true
        }
      }
      instance = "kong-az-east"
    }
    display_name = "Test TF Azure AI SE"

    name = "tf-test-azure-ai-provider"
  }
}


resource "konnect_ai_gateway_model_provider" "my_aigatewaymodelprovider_anthropic" {
  provider = konnect-beta
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  anthropic = {
    config = {
      auth = {
        headers = []
        params = [
          {
            location = "body"
            name     = "param_name"
            value    = "...my_value..."
          }
        ]
      }
    }
    display_name = "TF Test Anthropic AI Provider"
    name = "tf-test-anthropic-provider"
  }
}