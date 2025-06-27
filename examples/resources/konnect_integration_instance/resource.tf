resource "konnect_integration_instance" "my_integrationinstance" {
  provider = konnect-beta
  config = {
    key = jsonencode("value")
  }
  description      = "...my_description..."
  display_name     = "AWS (prod)"
  integration_name = "aws-lambda"
  name             = "aws-lambda-prod"
}