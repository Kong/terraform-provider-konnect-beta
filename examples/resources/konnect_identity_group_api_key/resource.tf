resource "konnect_identity_group_api_key" "my_identitygroupapikey" {
  provider = konnect-beta
  directory_id = "43362763-22ea-4f72-a11c-338abb6a15f8"
  expires_at   = "2022-11-04T20:10:06.927Z"
  group_id     = "6a68841d-e321-49ad-bef4-e849bcf5b944"
  managed_by = {
    key = "value"
  }
  metadata = "{ \"see\": \"documentation\" }"
  name     = "...my_name..."
  owner_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  secret   = "...my_secret..."
  type     = "v1"
}