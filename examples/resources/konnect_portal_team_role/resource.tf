resource "konnect_portal_team_role" "my_portalteamrole" {
  provider = konnect-beta
  entity_id        = "e67490ce-44dc-4cbd-b65e-b52c746fc26a"
  entity_region    = "eu"
  entity_type_name = "Services"
  portal_id        = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  role_name        = "API Viewer"
  team_id          = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
}