resource "konnect-beta_auth_server" "my_authserver" {
  audience    = "...my_audience..."
  description = "...my_description..."
  labels = {
    key = "value"
  }
  name              = "...my_name..."
  signing_algorithm = "RS384"
}