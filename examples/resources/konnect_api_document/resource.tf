resource "konnect_api_document" "my_apidocument" {
  provider = konnect-beta
  api_id             = "9f5061ce-78f6-4452-9108-ad7c02821fd5"
  content            = "...my_content..."
  parent_document_id = "0b689d9d-af35-4768-9730-3ec1c14d44e3"
  slug               = "api-document"
  status             = "unpublished"
  title              = "API Document"
}