resource "konnect_organization_personal_access_token_settings" "my_organizationpersonalaccesstokensettings" {
  provider = konnect-beta
  max_expiration_period_days = 364
  organization_id            = "da10965a-36e6-4938-9e05-949d4e8c184e"
  pats_enabled               = true
}
