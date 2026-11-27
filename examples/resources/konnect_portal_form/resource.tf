resource "konnect_portal_form" "my_portalform" {
  provider = konnect-beta
  api_registration = {
    fields = [
      {
        textarea = {
          description = "Please provide as much detail as possible. Supports MDC markdown."
          label       = "Message"
          name        = "full-name"
          placeholder = "Tell us what you need..."
          required    = true
        }
      }
    ]
    name   = "flights-api-registration"
    status = "unpublished"
  }
  developer_registration = {
    fields = [
      {
        submit = {
          name  = "submit"
          value = "Create account"
        }
      }
    ]
    name   = "developer-registration"
    status = "unpublished"
  }
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
}