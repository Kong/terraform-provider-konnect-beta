resource "konnect-beta_mesh_timeout" "my_meshtimeout" {
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
          connection_timeout = "...my_connection_timeout..."
          http = {
            max_connection_duration = "...my_max_connection_duration..."
            max_stream_duration     = "...my_max_stream_duration..."
            request_headers_timeout = "...my_request_headers_timeout..."
            request_timeout         = "...my_request_timeout..."
            stream_idle_timeout     = "...my_stream_idle_timeout..."
          }
          idle_timeout = "...my_idle_timeout..."
        }
        matches = [
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
      }
    ]
    target_ref = {
      kind = "Mesh"
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
    to = [
      {
        default = {
          connection_timeout = "...my_connection_timeout..."
          http = {
            max_connection_duration = "...my_max_connection_duration..."
            max_stream_duration     = "...my_max_stream_duration..."
            request_headers_timeout = "...my_request_headers_timeout..."
            request_timeout         = "...my_request_timeout..."
            stream_idle_timeout     = "...my_stream_idle_timeout..."
          }
          idle_timeout = "...my_idle_timeout..."
        }
        target_ref = {
          kind = "MeshService"
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
    ]
  }
  type = "MeshTimeout"
}