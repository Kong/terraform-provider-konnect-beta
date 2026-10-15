resource "konnect_mesh_access_role" "my_meshaccessrole" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  rules = [
    {
      access = [
        "GENERATE_ZONE_TOKEN"
      ]
      mesh = "...my_mesh..."
      names = [
        "..."
      ]
      types = [
        "..."
      ]
      when = [
        {
          dp_token = {
            tags = [
              {
                name  = "...my_name..."
                value = "...my_value..."
              }
            ]
          }
          from = {
            target_ref = {
              kind = "...my_kind..."
              labels = {
                key = "value"
              }
              name = "...my_name..."
            }
          }
          target_ref = {
            kind = "...my_kind..."
            labels = {
              key = "value"
            }
            name = "...my_name..."
          }
          to = {
            target_ref = {
              kind = "...my_kind..."
              labels = {
                key = "value"
              }
              name = "...my_name..."
            }
          }
        }
      ]
    }
  ]
  type = "...my_type..."
}