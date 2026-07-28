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
        type = "basic"
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
    type = "anthropic"
  }
  azure = {
    config = {
      auth = {
        azure = {
          client_id            = "...my_client_id..."
          client_secret        = "...my_client_secret..."
          tenant_id            = "...my_tenant_id..."
          type                 = "azure"
          use_managed_identity = true
        }
      }
      instance = "kong-az-east"
    }
    display_name = "Azure AI SE"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "azure-ai-se"
    type = "azure"
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
          sts_endpoint_url  = "...my_sts_endpoint_url..."
          type              = "aws"
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
    type = "bedrock"
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
        type = "basic"
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
    type = "cerebras"
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
        type = "basic"
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
    type = "cohere"
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
        type = "basic"
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
    type = "dashscope"
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
        type = "basic"
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
    type = "databricks"
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
        type = "basic"
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
    type = "deepseek"
  }
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  gemini = {
    config = {
      auth = {
        gcp = {
          metadata_url            = "...my_metadata_url..."
          oauth_token_url         = "...my_oauth_token_url..."
          service_account_json    = "...my_service_account_json..."
          type                    = "gcp"
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
    type = "gemini"
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
        type = "basic"
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
    type = "huggingface"
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
        type = "basic"
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
    type = "kimi"
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
        type = "basic"
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
    type = "llama2"
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
        type = "basic"
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
    type = "mistral"
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
        type = "basic"
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
    type = "ollama"
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
        type = "basic"
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
    type = "openai"
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
        type = "basic"
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
    type = "vercel"
  }
  vertex = {
    config = {
      auth = {
        vertex = {
          service_account_json = "...my_service_account_json..."
          type                 = "vertex"
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
    type = "vertex"
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
        type = "basic"
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
    type = "vllm"
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
        type = "basic"
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
    type = "xai"
  }
}