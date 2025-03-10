resource "konnect_portal_custom_domain" "my_portalcustomdomain" {
  provider = konnect-beta
  enabled   = false
  hostname  = "...my_hostname..."
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  ssl = {
    domain_verification_method = "http"
  }
}