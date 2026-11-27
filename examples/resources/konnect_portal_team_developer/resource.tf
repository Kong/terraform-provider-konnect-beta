resource "konnect_portal_team_developer" "my_portalteamdeveloper" {
  provider = konnect-beta
  developer_id = "df120cb4-f60b-47bc-a2f8-6a28e6a3c63b"
  portal_id    = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  team_id      = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
}