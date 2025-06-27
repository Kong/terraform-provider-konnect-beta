resource "konnect_catalog_service" "my_catalogservice" {
  provider = konnect-beta
  custom_fields = "{ \"see\": \"documentation\" }"
  description   = "...my_description..."
  display_name  = "User Service"
  labels = {
    key = "value"
  }
  name = "user-svc"
}