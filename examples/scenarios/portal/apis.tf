resource "konnect_api" "my_api" {
  provider    = konnect-beta
  name        = "Cars"
  version     = "v1"
  description = "This is a description of a car API"
  attributes = jsonencode({
    color = ["red"],
    teams = ["team1", "team2"]
  })
  labels = {
    key = "value"
  }
}

resource "konnect_api_implementation" "my_apiimplementation" {
  provider = konnect-beta
  api_id   = konnect_api.my_api.id
  service = {
    control_plane_id = konnect_gateway_control_plane.demo_cp.id
    id               = konnect_gateway_service.httpbin.id
  }
}

resource "konnect_api_publication" "my_apipublication" {
  provider                   = konnect-beta
  api_id                     = konnect_api.my_api.id
  portal_id                  = konnect_portal.my_portal.id
  auto_approve_registrations = true
  visibility                 = "private"
  auth_strategy_ids = [
    konnect_application_auth_strategy.my_applicationauthstrategy.id
  ]
}

resource "konnect_api_specification" "my_apispecification" {
  provider = konnect-beta
  api_id   = konnect_api.my_api.id
  content  = file("specifications/cars.yaml")
  type     = "oas3"
}

resource "konnect_api_document" "my_apidocument" {
  provider = konnect-beta
  title    = "Sample Document"
  slug     = "saple-document"
  status   = "published"
  content  = file("documents/cars/sample.md")
  api_id   = konnect_api.my_api.id
}
