resource "konnect_resource" "my_resource" {
  provider = konnect-beta
  data                    = "{ \"see\": \"documentation\" }"
  integration_instance_id = "3f51fa25-310a-421d-bd1a-007f859021a3"
  resource_id             = "AflTNLY0tTQhv2my"
  type                    = "gateway_svc"
}