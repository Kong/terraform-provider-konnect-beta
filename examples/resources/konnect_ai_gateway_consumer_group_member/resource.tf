resource "konnect_ai_gateway_consumer_group_member" "my_aigatewayconsumergroupmember" {
  provider = konnect-beta
  consumer_group_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  consumer_id       = "cf4c7e60-11db-49dd-b300-7c7e5f0f7e6b"
  gateway_id        = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
}