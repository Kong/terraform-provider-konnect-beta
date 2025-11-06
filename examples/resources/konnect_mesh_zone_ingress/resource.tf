resource "konnect_mesh_zone_ingress" "my_meshzoneingress" {
  provider = konnect-beta
  available_services = [
    {
      external_service = false
      instances        = 10
      mesh             = "...my_mesh..."
      tags = {
        key = "value"
      }
    }
  ]
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 9
    }
    advertised_address = "...my_advertised_address..."
    advertised_port    = 2
    port               = 0
  }
  type = "...my_type..."
  zone = "...my_zone..."
}