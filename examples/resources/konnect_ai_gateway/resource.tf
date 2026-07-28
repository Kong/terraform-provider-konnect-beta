resource "konnect_ai_gateway" "my_aigateway" {
  provider = konnect-beta
  additional_properties = "{ \"see\": \"documentation\" }"
  description           = "An AI Gateway for my organization."
  display_name          = "My AI Gateway"
  labels = {
    key = "value"
  }
  name = "my-ai-gateway"
  proxy_urls = [
    {
      host     = "...my_host..."
      port     = 4
      protocol = "...my_protocol..."
    }
  ]
}