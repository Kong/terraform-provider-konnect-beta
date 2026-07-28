resource "konnect_ai_gateway_consumer" "my_aigatewayconsumer" {
  provider = konnect-beta
  display_name          = "TF Test Consumer"
  gateway_id            = konnect_ai_gateway.my_aigateway.id
  name = "tf-test-consumer"
  type = "api-key"
}

resource "konnect_ai_gateway_consumer_group" "my_aigatewayconsumergroup" {
  provider = konnect-beta

  display_name          = "TF Test Consumer Group"
  gateway_id            = konnect_ai_gateway.my_aigateway.id

  name = "tf-test-consumers"
}

resource "konnect_ai_gateway_consumer_group_member" "my_aigatewayconsumergroupmember" {
  provider = konnect-beta
  consumer_id            = konnect_ai_gateway_consumer.my_aigatewayconsumer.id
  consumer_group_id   = konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup.id

  gateway_id          = konnect_ai_gateway.my_aigateway.id
}

resource "konnect_ai_gateway_consumer_credential" "my_aigatewayconsumercredential" {
  provider = konnect-beta
  api_key      = "tf-test-dev-key"
  consumer_id  = konnect_ai_gateway_consumer.my_aigatewayconsumer.id
  display_name = "TF Test Dev Key"
  gateway_id   = konnect_ai_gateway.my_aigateway.id
  name = "tf-test-dev-key"
  ttl = 120
  type = "api-key"
}