resource "konnect_ai_gateway_vault" "my_aigatewayvault" {
  provider = konnect-beta
  aws = {
    config = {
      assume_role_arn   = "...my_assume_role_arn..."
      base64_decode     = true
      endpoint_url      = "...my_endpoint_url..."
      neg_ttl           = 0
      region            = "us-east-1"
      resurrect_ttl     = 100000000
      role_session_name = "KongVault"
      sts_endpoint_url  = "...my_sts_endpoint_url..."
      ttl               = 0
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  azure = {
    config = {
      base64_decode      = true
      client_id          = "...my_client_id..."
      credentials_prefix = "AZURE"
      location           = "...my_location..."
      neg_ttl            = 0
      resurrect_ttl      = 100000000
      tenant_id          = "...my_tenant_id..."
      ttl                = 0
      type               = "secrets"
      vault_uri          = "...my_vault_uri..."
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  conjur = {
    config = {
      account       = "...my_account..."
      api_key       = "...my_api_key..."
      base64_decode = false
      endpoint_url  = "...my_endpoint_url..."
      login         = "...my_login..."
      neg_ttl       = 0
      resurrect_ttl = 100000000
      ttl           = 0
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  env = {
    config = {
      base64_decode = false
      prefix        = "MY_SECRET_"
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  gcp = {
    config = {
      base64_decode = true
      neg_ttl       = 0
      project_id    = "...my_project_id..."
      resurrect_ttl = 100000000
      ttl           = 0
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  hcv = {
    config = {
      jwt = {
        audiences      = "...my_audiences..."
        auth_method    = "jwt"
        base64_decode  = true
        client_id      = "...my_client_id..."
        client_secret  = "...my_client_secret..."
        host           = "...my_host..."
        kv             = "v1"
        mount          = "secret"
        namespace      = "...my_namespace..."
        neg_ttl        = 0
        port           = 5
        protocol       = "https"
        resurrect_ttl  = 100000000
        role           = "demo"
        ssl_verify     = true
        token_endpoint = "...my_token_endpoint..."
        ttl            = 0
      }
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
  konnect = {
    config = {
      config_store_id = "77426bee-2bca-4005-81af-284868fd3038"
    }
    description = "This vault is used to retrieve redis database access credentials"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name = "my-awesome-vault"
  }
}