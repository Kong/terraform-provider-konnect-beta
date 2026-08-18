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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
        }
      }
    }
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
  }
  vertex = {
    config = {
      auth = {
        vertex = {
          service_account_json = "...my_service_account_json..."
        }
      }
    }
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
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
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
  }
}