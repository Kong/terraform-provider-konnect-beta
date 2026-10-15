resource "konnect_organization_personal_access_token_settings" "my_organizationpersonalaccesstokensettings" {
  provider = konnect-beta
  max_expiration_period_days = 170
  organization_id            = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  pats_enabled               = true
}