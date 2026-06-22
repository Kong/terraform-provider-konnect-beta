resource "konnect_identity_directory" "my_identitydirectory" {
  provider = konnect-beta
  allow_all_control_planes = false
  allowed_control_planes = [
    "ae809a12-d526-462e-bd1a-43f66dc49389"
  ]
  description   = "...my_description..."
  force_destroy = "false"
  labels = {
    key = "value"
  }
  managed_by = {
    key = "value"
  }
  name              = "...my_name..."
  negative_ttl_secs = 600
  ttl_secs          = 600
}