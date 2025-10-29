resource "konnect-beta_mesh_hostname_generator" "my_meshhostnamegenerator" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  spec = {
    extension = {
      config = "{ \"see\": \"documentation\" }"
      type   = "...my_type..."
    }
    selector = {
      mesh_external_service = {
        match_labels = {
          key = "value"
        }
      }
      mesh_multi_zone_service = {
        match_labels = {
          key = "value"
        }
      }
      mesh_service = {
        match_labels = {
          key = "value"
        }
      }
    }
    template = "...my_template..."
  }
  type = "HostnameGenerator"
}