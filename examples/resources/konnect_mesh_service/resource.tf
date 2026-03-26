resource "konnect_mesh_service" "my_meshservice" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    identities = [
      {
        type  = "SpiffeID"
        value = "...my_value..."
      }
    ]
    ports = [
      {
        app_protocol = "tcp"
        name         = "...my_name..."
        port         = 8
        target_port = {
          integer = 7
        }
      }
    ]
    selector = {
      dataplane_labels = {
        match_labels = {
          key = "value"
        }
      }
      dataplane_ref = {
        name = "...my_name..."
      }
      dataplane_tags = {
        key = "value"
      }
    }
    state = "Unavailable"
  }
  type = "MeshService"
}