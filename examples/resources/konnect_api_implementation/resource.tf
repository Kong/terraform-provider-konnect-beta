resource "konnect_api_implementation" "my_apiimplementation" {
  provider = konnect-beta
  api_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  control_plane_reference = {
    control_plane = {
      id = "7710d5c4-d902-410b-992f-18b814155b53"
    }
    environment_id = "2747d1e5-8246-4f65-a939-b392f1ee17f8"
  }
  service_reference = {
    environment_id = "2747d1e5-8246-4f65-a939-b392f1ee17f8"
    service = {
      control_plane_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
      id               = "7710d5c4-d902-410b-992f-18b814155b53"
    }
  }
}