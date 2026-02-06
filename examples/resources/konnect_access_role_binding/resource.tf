resource "konnect_access_role_binding" "my_accessrolebinding" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  name = "...my_name..."
  roles = [
    "..."
  ]
  subjects = [
    {
      name = "...my_name..."
      type = "...my_type..."
    }
  ]
  type = "...my_type..."
}