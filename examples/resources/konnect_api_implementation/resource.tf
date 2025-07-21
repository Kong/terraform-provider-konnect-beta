resource "konnect_api_implementation" "my_apiimplementation" {
  provider = konnect-beta
  api_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  api_implementation_control_plane_entity = {
    control_plane = {
      control_plane_id = "7710d5c4-d902-410b-992f-18b814155b53"
    }
  }
  api_implementation_gateway_service_entity = {
    service = {
      control_plane_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
      id               = "7710d5c4-d902-410b-992f-18b814155b53"
    }
  }
}