resource "random_id" "hostname_suffix" {
  byte_length = 4
}

resource "konnect_portal" "my_portal" {
  provider                             = konnect-beta
  name                                 = "Hello World"
  authentication_enabled               = true
  auto_approve_applications            = true
  auto_approve_developers              = true
  default_api_visibility               = "public"
  default_application_auth_strategy_id = konnect_application_auth_strategy.my_applicationauthstrategy.id
  default_page_visibility              = "private"
  description                          = "This is a description"
  display_name                         = "Hello World"
  labels = {
    key = "value"
  }
  rbac_enabled = true
}

resource "konnect_portal_custom_domain" "my_portalcustomdomain" {
  provider  = konnect-beta
  enabled   = false
  # From the API: Custom hostnames ending in example.com, example.net, or example.org are prohibited
  hostname  = "tfautomatedtests-${random_id.hostname_suffix.hex}.mheap.xyz"
  portal_id = konnect_portal.my_portal.id
  ssl = {
    domain_verification_method = "http"
  }
}

resource "konnect_portal_page" "my_portalpage" {
  provider = konnect-beta

  content        = "Content for About page"
  description    = "A custom page about the business"
  parent_page_id = null
  portal_id      = konnect_portal.my_portal.id
  slug           = "/about"
  status         = "published"
  title          = "About"
  visibility     = "public"
}

resource "konnect_portal_snippet" "my_portalsnippet" {
  provider    = konnect-beta
  content     = "# Welcome to My Snippet"
  description = "A custom page about developer portals"
  name        = "my-snippet"
  portal_id   = konnect_portal.my_portal.id
  status      = "published"
  title       = "My Snippet"
  visibility  = "public"
}

locals {
  footer_sections = [
    for title in ["Page One", "Page Two", "Page Three", "Page Four"] : {
      external   = false
      path       = "/"
      title      = title
      visibility = "public"
    }
  ]
}

resource "konnect_portal_customization" "my_portalcustomization" {
  provider = konnect-beta

  menu = {
    main = [
      {
        external   = false
        path       = "/about"
        title      = "About"
        visibility = "public"
      },
      {
        external   = false
        path       = "/secret"
        title      = "Secret"
        visibility = "private"
      }
    ]

    footer_sections = [
      {
        title = "Section One"
        items = local.footer_sections
      },
      {
        title = "Section Two"
        items = local.footer_sections
      },
      {
        title = "Section Three"
        items = local.footer_sections
      },
      {
        title = "Section Four"
        items = local.footer_sections
      }
    ]

    footer_bottom = [
      {
        external   = false
        path       = "/about/company"
        title      = "My Page"
        visibility = "public"
      }
    ]
  }

  portal_id = konnect_portal.my_portal.id
  spec_renderer = {
    infinite_scroll = true
    show_schemas    = false
    try_it_insomnia = false
    try_it_ui       = true
  }
  theme = {
    mode = "system"
  }
}

resource "konnect_portal_logo" "my_portallogo" {
  provider  = konnect-beta
  data      = "data:image/jpeg;base64,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"
  portal_id = konnect_portal.my_portal.id
}

resource "konnect_portal_favicon" "my_portalfavicon" {
  provider  = konnect-beta
  data      = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAMAAABg3Am1AAAAIGNIUk0AAHomAACAhAAA+gAAAIDoAAB1MAAA6mAAADqYAAAXcJy6UTwAAAFNUExURQAAAAEAAAIAAAMAAAoAACsBAUYDAlYDAkMCAicBAQgAADgCAYgFA74HBNcIBd8JBeEJBd4IBdUIBbwHBIcFAzkCASgBAZ8GA9wIBdkIBdMIBdEIBdAIBdkJBd0JBaAGBCkCAVkDAtYIBd8IBdIIBdQIBV0EAmkEAuMJBeQJBXIEAlMDAuMIBVwDAiMBASkBAZwGA6IGAz8CAYoFA9AHBeAIBZEFA78HBNgIBcMHBA8AANMIBDECAUMCAdAIBEwDAU4CAtMHBeEIBVcDAk4DAtEHBUQCAdMHBNAHBE8DAioBAdQIBNsIBTcCAcgHBBQBAIkFA5cGA0gCAZ4GA90IBa0GA9wIBDYCAVoDAs8IBeUIBW4EAnEEAuUJBYQEAl8DAm8EAi0BAaYGA0ECAZQFA8cHBOIJBeIIBZYFA0YCARIAAFUDAmEDAhIBAP///3DEI5gAAAABYktHRG4iD1EXAAAAB3RJTUUH6QYeDScQDbMEtAAAAfhJREFUSMftVFlb1DAUvU1AlpHaosBMbtyahEUQQami41bcZQYFQVkEBAUX4P+/mqbtLB9phzd94DxlOufcnJPcXIBz/LdwCNUgzhnplJ5eFSCu29V9oaent68/3qtjeYDSxQH3kuf5g5evDJkPheZheKTsVRhyRHbVu3b9BhRF0fybgUAumYZikqNwR8HJV1AYG1dcKU1OwdnELSD5/MkplCpGJlCST9zOy0FgWg7KhMcakCq4k7MHgRnB2SlwfxYcu6FSGRv1M1c6D/p3raYo3PN4m5tUNxfet3lyoOvBvMx4LRIm3YePLKYIdLtNJ5pXbeauPLZsQeGJjy1GWhZPxTNLCArPQ0zcqCgTJAu+8MIqeBliUlGpKCZHUSp7lSN4LdAcY1q+GjGteKOzvA3fWQQE3gemusp6I+6pRRO6VrcI9LEOVKRqg5FXWbD0wXbXFD4ucNbCNaE1UCxbm4lCXbeGMimax6p/8XDF3q8OfBLcvIU2Ba6ugb37CPQFSjbyZhr5+Uveg6CwHj8g1jwlvd3G5lb+ICDwdZthQq2aS0e5uVPwqPU/u/M+Zi2lZ0d59VtOgFRBYG9fCFci51y6QqyVioZGkgO+/zioeYc/vdrBr9/6Q6cJGxP+1I+Oj4/qJ9Bp8KWSxrWSjuWzKNTgrOP+HP8AfwE9ZVapIb5/2wAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyNS0wNi0zMFQxMzozOTowOCswMDowMP+OT7EAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjUtMDYtMzBUMTM6Mzk6MDgrMDA6MDCO0/cNAAAAKHRFWHRkYXRlOnRpbWVzdGFtcAAyMDI1LTA2LTMwVDEzOjM5OjE2KzAwOjAwRVOtEQAAAABJRU5ErkJggg=="
  portal_id = konnect_portal.my_portal.id
}
