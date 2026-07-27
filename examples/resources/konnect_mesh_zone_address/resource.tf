resource "konnect_mesh_zone_address" "my_meshzoneaddress" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    address = "...my_address..."
    port    = 62290
  }
  type = "MeshZoneAddress"
}