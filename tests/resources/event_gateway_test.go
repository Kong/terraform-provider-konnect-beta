package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestEventGateway(t *testing.T) {
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
						resource.TestCheckResourceAttr("konnect_event_gateway_schema_registry.my_eventgatewayschemaregistry", "confluent.config.schema_type", "avro"),

						// Event Gateway Listener Policy Forward To Virtual Cluster
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.my_eventgatewaylistenerpolicy", "name", "listenerpolicyname"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.my_eventgatewaylistenerpolicy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.my_eventgatewaylistenerpolicy", "config.port_mapping.advertised_host", "host.example.com"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.my_eventgatewaylistenerpolicy", "config.port_mapping.bootstrap_port", "none"),
						resource.TestCheckResourceAttr("konnect_event_gateway_listener_policy_forward_to_virtual_cluster.my_eventgatewaylistenerpolicy", "config.port_mapping.min_broker_id", "5"),

						// Event Gateway Produce Policy Modify Headers
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.my_eventgatewayvirtualclusterproducepolicy", "name", "myproducename"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.my_eventgatewayvirtualclusterproducepolicy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.my_eventgatewayvirtualclusterproducepolicy", "condition", "context.topic.name.endsWith('my_suffix')"),
						resource.TestCheckResourceAttr("konnect_event_gateway_produce_policy_modify_headers.my_eventgatewayvirtualclusterproducepolicy", "config.actions.0.remove.key", "...my_key..."),

						// Event Gateway Consume Policy Modify Headers
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_modify_headers.my_eventgatewayvirtualclusterconsumepolicy", "name", "myconsumename"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_modify_headers.my_eventgatewayvirtualclusterconsumepolicy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_modify_headers.my_eventgatewayvirtualclusterconsumepolicy", "condition", "context.topic.name.endsWith('my_suffix')"),
						resource.TestCheckResourceAttr("konnect_event_gateway_consume_policy_modify_headers.my_eventgatewayvirtualclusterconsumepolicy", "config.actions.0.remove.key", "mykey"),

						// Event Gateway Cluster Policy ACLs
						resource.TestCheckResourceAttr("konnect_event_gateway_cluster_policy_acls.my_eventgatewayvirtualclusterclusterpolicy", "name", "mynamecluster"),
						resource.TestCheckResourceAttr("konnect_event_gateway_cluster_policy_acls.my_eventgatewayvirtualclusterclusterpolicy", "enabled", "true"),
						resource.TestCheckResourceAttr("konnect_event_gateway_cluster_policy_acls.my_eventgatewayvirtualclusterclusterpolicy", "config.rules.0.action", "deny"),
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
