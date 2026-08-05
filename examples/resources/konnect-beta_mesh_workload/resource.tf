resource "konnect-beta_mesh_workload" "my_meshworkload" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    # ...
  }
  type = "Workload"
}