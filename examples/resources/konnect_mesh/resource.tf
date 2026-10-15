resource "konnect_mesh" "my_mesh" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh_services = {
    mode = "Everywhere"
  }
  name = "...my_name..."
  type = "...my_type..."
}