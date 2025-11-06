resource "konnect_mesh_secret" "my_meshsecret" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  data  = "...my_data..."
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  type = "...my_type..."
}