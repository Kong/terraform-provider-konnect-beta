resource "konnect_identity_group" "my_identitygroup" {
  provider = konnect-beta
  description  = "...my_description..."
  directory_id = "307cd9e3-ea12-451b-8bb2-b6c398132110"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  metadata = "{ \"see\": \"documentation\" }"
  name     = "...my_name..."
  type     = "group"
}