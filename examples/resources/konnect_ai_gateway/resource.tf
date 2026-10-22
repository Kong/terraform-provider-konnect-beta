resource "konnect_ai_gateway" "my_aigateway" {
  provider = konnect-beta
  deployment_type = "hybrid"
  description     = "An AI Gateway for my organization."
  display_name    = "My AI Gateway"
  labels = {
    key = "value"
  }
  min_runtime_version = "2.0"
  name                = "my-ai-gateway"
  proxy_urls = [
    {
      host     = "...my_host..."
      port     = 4
      protocol = "...my_protocol..."
    }
  ]
  runtime_auto_upgrade = true
}