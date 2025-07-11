---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_api_document Resource - terraform-provider-konnect-beta"
subcategory: ""
description: |-
  APIDocument Resource
---

# konnect_api_document (Resource)

APIDocument Resource

## Example Usage

```terraform
resource "konnect_api_document" "my_apidocument" {
  provider = konnect-beta
  api_id             = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  content            = "...my_content..."
  parent_document_id = "b689d9da-f357-4687-8303-ec1c14d44e37"
  slug               = "api-document"
  status             = "published"
  title              = "API Document"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_id` (String) The UUID API identifier
- `content` (String) Raw markdown content to display in your Portal

### Optional

- `parent_document_id` (String) API Documents may be rendered as a tree of files.

Specify the `id` of another API Document as the `parent_document_id` to add some heirarchy do your documents.
- `slug` (String) The `slug` is used in generated URLs to provide human readable paths.

Defaults to `slugify(title)`
- `status` (String) If `status=published` the document will be visible in your live portal. Default: "unpublished"; must be one of ["published", "unpublished"]
- `title` (String) The title of the document. Used to populate the `<title>` tag for the page

### Read-Only

- `created_at` (String) An ISO-8601 timestamp representation of entity creation date.
- `id` (String) The API document identifier.
- `updated_at` (String) An ISO-8601 timestamp representation of entity update date.

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_api_document.my_konnect_api_document "{ \"api_id\": \"9f5061ce-78f6-4452-9108-ad7c02821fd5\",  \"id\": \"de5c9818-be5c-42e6-b514-e3d4bc30ddeb\"}"
```
