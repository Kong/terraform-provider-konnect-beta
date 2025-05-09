---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_portal_team Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  PortalTeam Resource
---

# konnect_portal_team (Resource)

PortalTeam Resource

## Example Usage

```terraform
resource "konnect_portal_team" "my_portalteam" {
  provider = konnect-beta
  description = "The Identity Management (IDM) team."
  name        = "IDM - Developers"
  portal_id   = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `portal_id` (String) The Portal identifier

### Optional

- `description` (String)

### Read-Only

- `created_at` (String)
- `id` (String) The ID of this resource.
- `updated_at` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_portal_team.my_konnect_portal_team "{ \"portal_id\": \"f32d905a-ed33-46a3-a093-d8f536af9a8a\",  \"id\": \"d32d905a-ed33-46a3-a093-d8f536af9a8a\"}"
```
