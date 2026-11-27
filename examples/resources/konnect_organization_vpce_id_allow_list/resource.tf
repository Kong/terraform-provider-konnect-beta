resource "konnect_organization_vpce_id_allow_list" "my_organizationvpceidallowlist" {
  provider = konnect-beta
  allowed_vpce_ids = [
    "vpce-0123456789abcdef0",
  ]
  enabled         = true
  organization_id = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
}