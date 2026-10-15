resource "konnect_ai_gateway_data_plane_certificate" "my_aigatewaydataplanecertificate" {
  provider = konnect-beta
  cert        = "...my_cert..."
  description = "...my_description..."
  gateway_id  = "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"
  title       = "...my_title..."
}