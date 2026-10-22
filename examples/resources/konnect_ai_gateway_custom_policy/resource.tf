resource "konnect_ai_gateway_custom_policy" "my_aigatewaycustompolicy" {
  provider = konnect-beta
  gateway_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  installed = {
    display_name = "Custom Policy - Installed plugin"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name   = "my-installed-custom-policy"
    schema = "<lua_schema>"
  }
  streaming = {
    display_name = "Custom Policy Streaming Plugin"
    handler      = "<lua_handler>"
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    name   = "my-streaming-custom-policy"
    schema = "<lua_schema>"
  }
}