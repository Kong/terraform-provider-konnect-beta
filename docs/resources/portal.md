---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_portal Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  Portal Resource
---

# konnect_portal (Resource)

Portal Resource

## Example Usage

```terraform
resource "konnect_portal" "my_portal" {
  provider = konnect-beta
  authentication_enabled               = false
  auto_approve_applications            = false
  auto_approve_developers              = false
  default_api_visibility               = "public"
  default_application_auth_strategy_id = "e7d77a5f-c5f5-49db-9b2f-cabb4401add8"
  default_page_visibility              = "private"
  description                          = "...my_description..."
  display_name                         = "...my_display_name..."
  force_destroy                        = "false"
  labels = {
    key = "value"
  }
  name         = "...my_name..."
  rbac_enabled = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the portal, used to distinguish it from other portals. Name must be unique.

### Optional

- `authentication_enabled` (Boolean) Whether the portal supports developer authentication. If disabled, developers cannot register for accounts or create applications.
- `auto_approve_applications` (Boolean) Whether requests from applications to register for APIs will be automatically approved, or if they will be set to pending until approved by an admin.
- `auto_approve_developers` (Boolean) Whether developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.
- `default_api_visibility` (String) The default visibility of APIs in the portal. If set to `public`, newly published APIs are visible to unauthenticated developers. If set to `private`, newly published APIs are hidden from unauthenticated developers. must be one of ["public", "private"]
- `default_application_auth_strategy_id` (String) The default authentication strategy for APIs published to the portal. Newly published APIs will use this authentication strategy unless overridden during publication. If set to `null`, API publications will not use an authentication strategy unless set during publication.
- `default_page_visibility` (String) The default visibility of pages in the portal. If set to `public`, newly created pages are visible to unauthenticated developers. If set to `private`, newly created pages are hidden from unauthenticated developers. must be one of ["public", "private"]
- `description` (String) A description of the portal.
- `display_name` (String) The display name of the portal. This value will be the portal's `name` in Portal API.
- `force_destroy` (String) If set to "true", the portal and all child entities will be deleted when running `terraform destroy`.
If set to "false", the portal will not be deleted until all child entities are manually removed.
This will IRREVERSIBLY DELETE ALL REGISTERED DEVELOPERS AND THEIR CREDENTIALS. Only set to "true" if you want this behavior.
Default: "false"; must be one of ["true", "false"]
- `labels` (Map of String) Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. 

Labels are intended to store **INTERNAL** metadata.

Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
- `rbac_enabled` (Boolean) Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for APIs until unless assigned to teams with access to view and consume specific APIs. Authentication must be enabled to use RBAC.

### Read-Only

- `canonical_domain` (String) The canonical domain of the developer portal
- `created_at` (String) An ISO-8601 timestamp representation of entity creation date.
- `default_domain` (String) The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a `custom_domain``.
- `id` (String) Contains a unique identifier used for this resource.
- `updated_at` (String) An ISO-8601 timestamp representation of entity update date.

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_portal.my_konnect_portal "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
```
