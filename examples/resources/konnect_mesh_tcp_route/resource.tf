resource "konnect_mesh_tcp_route" "my_meshtcproute" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Mesh"
      labels = {
        key = "value"
      }
      section_name = "...my_section_name..."
    }
    to = [
      {
        rules = [
          {
            default = {
              backend_refs = [
                {
                  kind = "MeshExternalService"
                  labels = {
                    key = "value"
                  }
                  port         = 6
                  section_name = "...my_section_name..."
                  weight       = 4274592322
                }
              ]
            }
          }
        ]
        target_ref = {
          kind = "Mesh"
          labels = {
            key = "value"
          }
          section_name = "...my_section_name..."
        }
      }
    ]
  }
  type = "MeshTCPRoute"
}