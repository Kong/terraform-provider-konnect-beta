resource "konnect_portal_developer" "my_portaldeveloper" {
  provider = konnect-beta
  email                 = "Rhiannon.Johnson21@yahoo.com"
  full_name             = "...my_full_name..."
  portal_id             = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  send_invitation_email = false
  status                = "approved"
}