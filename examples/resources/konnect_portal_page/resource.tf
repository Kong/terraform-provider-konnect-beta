resource "konnect_portal_page" "my_portalpage" {
  provider = konnect-beta
  content        = "# Welcome to My Page"
  description    = "A custom page about developer portals"
  parent_page_id = "6824a28d-9702-4426-9aed-403b50452182"
  portal_id      = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  slug           = "/my-page"
  status         = "published"
  title          = "My Page"
  visibility     = "public"
}