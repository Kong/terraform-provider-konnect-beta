resource "konnect_mesh_zone" "my_meshzone" {
  provider = konnect-beta
  cp_id   = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  enabled = true
  labels = {
    key = "value"
  }
  name = "...my_name..."
  type = "...my_type..."
}