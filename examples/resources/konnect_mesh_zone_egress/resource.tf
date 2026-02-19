resource "konnect_mesh_zone_egress" "my_meshzoneegress" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 4
    }
    port = 2
  }
  type = "...my_type..."
  zone = "...my_zone..."
}