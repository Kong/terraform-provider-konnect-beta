resource "konnect_portal_customization" "my_portalcustomization" {
  provider = konnect-beta
  provider = konnect-beta
  provider = konnect-beta
  provider = konnect-beta
  css = "...my_css..."
  js = {
    custom = "...my_custom..."
    scripts = [
      "..."
    ]
  }
  layout = "...my_layout..."
  menu = {
    footer_bottom = [
      {
        external   = false
        path       = "/about/company"
        title      = "My Page"
        visibility = "public"
      }
    ]
    footer_sections = [
      {
        items = [
          {
            external   = true
            path       = "/about/company"
            title      = "My Page"
            visibility = "public"
          }
        ]
        title = "...my_title..."
      }
    ]
    main = [
      {
        external   = true
        path       = "/about/company"
        title      = "My Page"
        visibility = "public"
      }
    ]
  }
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
  robots    = "...my_robots..."
  spec_renderer = {
    infinite_scroll = false
    show_schemas    = false
    try_it_insomnia = false
    try_it_ui       = false
  }
  theme = {
    colors = {
      primary = "#000000"
    }
    mode = "system"
    name = "...my_name..."
  }
}