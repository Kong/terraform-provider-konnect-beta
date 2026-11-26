resource "konnect_identity_principal" "my_identityprincipal" {
  provider = konnect-beta
  active       = true
  description  = "...my_description..."
  directory_id = "7ce7d7f7-55e9-4ce9-a436-233262cc9be9"
  display_name = "...my_display_name..."
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  metadata = "{ \"see\": \"documentation\" }"
  type     = "agent"
}