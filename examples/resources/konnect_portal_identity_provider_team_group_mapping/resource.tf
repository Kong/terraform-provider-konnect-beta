resource "konnect_portal_identity_provider_team_group_mapping" "my_portalidentityproviderteamgroupmapping" {
  provider = konnect-beta
  group                = "API Engineers"
  identity_provider_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
  portal_id            = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  team_id              = "6801e673-cc10-498a-94cd-4271de07a0d3"
}