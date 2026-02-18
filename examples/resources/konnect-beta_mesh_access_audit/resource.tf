resource "konnect-beta_mesh_access_audit" "my_meshaccessaudit" {
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  rules = [
    {
      access = [
        {
          str = "...my_str..."
        }
      ]
      access_all = true
      mesh       = "...my_mesh..."
      types = [
        "..."
      ]
    }
  ]
  type = "...my_type..."
}