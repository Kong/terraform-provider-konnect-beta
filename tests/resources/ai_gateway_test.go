package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAIGateway(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_ai_gateway.my_aigateway", "name", "my-test-ai-gateway"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_data_plane_certificate.my_aigatewaydataplanecertificate", "title", "tf-test-dp-cert"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer.my_aigatewayconsumer", "name", "tf-test-consumer"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_group.my_aigatewayconsumergroup", "name", "tf-test-consumers"),
						resource.TestCheckResourceAttrSet("konnect_ai_gateway_consumer_group_member.my_aigatewayconsumergroupmember", "consumer_id"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_consumer_credential.my_aigatewayconsumercredential", "display_name", "TF Test Dev Key"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store.my_aigatewayconfigstore", "name", "tf-test-config-store"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_vault.my_aigatewayvault", "name", "tf-test-vault"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "key", "tf-test-secret-key"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_config_store_secret.my_aigatewayconfigstoresecret", "value", "tf-test-secret-value"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_identity_provider.my_aigatewayidentityprovider", "name", "tf-test-key-auth-identity-provider"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_identity_provider.my_aigatewayidentityprovider2", "name", "tf-test-openid-connect-identity-provider"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_policy.my_aigatewaypolicy", "name", "ai-pii-sanitizer-1234"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_agent.my_aigatewayagent", "name", "tf-test-flight-booking-agent"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_model_provider.my_aigatewaymodelprovider2", "name", "tf-test-azure-ai-se"),
						resource.TestCheckResourceAttr("konnect_ai_gateway_model.my_aigatewaymodel", "name", "tf-test-model"),
					),
				},
			},
		})
	})
}
