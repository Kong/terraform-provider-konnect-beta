resource "konnect-beta_mesh_traffic_permission" "my_meshtrafficpermission" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    rules = [
      {
        default = {
          allow = [
            {
              sni = {
                type  = "Exact"
                value = "...my_value..."
              }
              spiffe_id = {
                type  = "Exact"
                value = "...my_value..."
              }
            }
          ]
          allow_with_shadow_deny = [
            {
              sni = {
                type  = "Exact"
                value = "...my_value..."
              }
              spiffe_id = {
                type  = "Prefix"
                value = "...my_value..."
              }
            }
          ]
          deny = [
            {
              sni = {
                type  = "Exact"
                value = "...my_value..."
              }
              spiffe_id = {
                type  = "Prefix"
                value = "...my_value..."
              }
            }
          ]
        }
      }
    ]
    target_ref = {
      kind = "MeshHTTPRoute"
      labels = {
        key = "value"
      }
      mesh         = "...my_mesh..."
      name         = "...my_name..."
      namespace    = "...my_namespace..."
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
  }
  type = "MeshTrafficPermission"
}