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
  hostname  = "portal.mheap.xyz"
  portal_id = konnect_portal.my_portal.id
  ssl = {
    domain_verification_method = "http"
  }
}

locals {
  pages = [
    {
      title      = "Home"
      slug       = "/"
      visibility = "public"
      filename   = "home.md"
    },
    {
      title      = "About"
      slug       = "/about"
      visibility = "public"
      filename   = "about.md"
    },
    {
      title      = "Contact"
      slug       = "/contact"
      visibility = "public"
      filename   = "contact.md"
    },
    {
      title      = "Secrets"
      slug       = "/secret"
      visibility = "private"
      filename   = "secret.md"
    }
  ]
}

resource "konnect_portal_page" "my_portalpage" {
  provider = konnect-beta
  for_each = { for page in local.pages : page.slug => page }

  content        = file("pages/${each.value.filename}")
  description    = "A custom page about ${each.value.title}"
  parent_page_id = null
  portal_id      = konnect_portal.my_portal.id
  slug           = each.value.slug
  status         = "published"
  title          = each.value.title
  visibility     = each.value.visibility
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
