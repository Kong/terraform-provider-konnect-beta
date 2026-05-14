resource "konnect_api_publication" "my_apipublication" {
  provider = konnect-beta
  api_id = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  auth_strategy_ids = [
    "9c3bed4d-0322-4ea0-ba19-a4bd65d821f6"
  ]
  auto_approve_registrations = true
  form_id                    = "deafd39d-fc2f-4b87-9ee2-8e8feb2eae76"
  portal_id                  = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  visibility                 = "private"
}