resource "konnect_portal_audit_log_webhook" "my_portalauditlogwebhook" {
  provider = konnect-beta
  audit_log_destination_id = "138339f6-7017-4c9e-ae91-4f542808e3e8"
  enabled                  = true
  include_principal_name   = true
  portal_id                = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
}