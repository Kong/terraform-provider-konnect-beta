package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestEventGateway(t *testing.T) {
	t.Skip("Only runs in dev")
	t.Run("smoke-test", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			Steps: []resource.TestStep{
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						// Event Gateway
						resource.TestCheckResourceAttr("konnect_event_gateway.my_eventgateway", "name", "myeventgateway"),

						// Event Gateway Listener
						resource.TestCheckResourceAttr("konnect_event_gateway_listener.my_eventgatewaylistener", "name", "my-egw-listener"),

						// Event Gateway Backend Cluster
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.my_eventgatewaybackendcluster", "name", "backendname"),

						// Event Gateway Virtual Cluster
						resource.TestCheckResourceAttr("konnect_event_gateway_virtual_cluster.my_eventgatewayvirtualcluster", "name", "virtualnamex"),

						// Event Gateway Vault
						resource.TestCheckResourceAttr("konnect_event_gateway_vault.my_eventgatewayvault", "env.name", "myvault"),

						// Event Gateway Schema Registry
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.my_eventgatewayschemaregistry", "confluent.name", "confname"),

						// Event Gateway Listener Policy
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy.my_eventgatewaylistenerpolicy", "forward_to_virtual_cluster.name", "listenerpolicyname"),

						// Event Gateway Virtual Cluster Produce Policy
						resource.TestCheckResourceAttr("konnect_event_gateway_virtual_cluster_produce_policy.my_eventgatewayvirtualclusterproducepolicy", "modify_headers.name", "myproducepolicy"),

						// Event Gateway Virtual Cluster Consume Policy
						resource.TestCheckResourceAttr("konnect_event_gateway_virtual_cluster_consume_policy.my_eventgatewayvirtualclusterconsumepolicy", "modify_headers.name", "myconsumepolicy"),

						// Event Gateway Virtual Cluster Cluster Policy
						resource.TestCheckResourceAttr("konnect_event_gateway_virtual_cluster_cluster_policy.my_eventgatewayvirtualclusterclusterpolicy", "id_prefix.name", "myclusterpolicy"),
					),
				},
				{
					ProtoV6ProviderFactories: providerFactory,
					ConfigDirectory:          config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})
}
