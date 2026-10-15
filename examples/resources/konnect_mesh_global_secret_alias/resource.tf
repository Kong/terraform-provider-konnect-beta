resource "konnect_mesh_global_secret_alias" "my_meshglobalsecretalias" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  data  = "...my_data..."
  labels = {
    key = "value"
  }
  name = "...my_name..."
  type = "...my_type..."
}