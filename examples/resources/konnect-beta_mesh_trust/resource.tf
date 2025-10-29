resource "konnect-beta_mesh_trust" "my_meshtrust" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    ca_bundles = [
      {
        pem = {
          value = "...my_value..."
        }
        type = "Pem"
      }
    ]
    origin = {
      kri = "...my_kri..."
    }
    trust_domain = "...my_trust_domain..."
  }
  type = "MeshTrust"
}