resource "konnect-beta_portal_application_registration" "my_portalapplicationregistration" {
  additional_data = {
    key = {
      array_of_str = [
        "..."
      ]
      boolean = false
      number  = 7.82
      str     = "...my_str..."
    }
  }
  api_id         = "15a8f0a0-e185-4b99-8591-c6319de5704a"
  application_id = "f5e17ae7-79e1-4181-931f-143b2159d313"
  portal_id      = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  status         = "approved"
}