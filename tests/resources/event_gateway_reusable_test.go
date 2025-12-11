package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestEventGatewayReusable(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("EGW Control plane", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "my_event_gateway" {
  				name = "my_test_event_gateway"
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.my_event_gateway", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "name", "my_test_event_gateway"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway.my_event_gateway", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})

	})

	t.Run("EGW Backend Cluster", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		egwCp, err := hclbuilder.FromString(`
			resource "konnect_event_gateway" "tf_test_event_gateway" {
  				name = "my_test_event_gateway"
			}
		`)
		require.NoError(t, err)

		egwBackendCluster, err := hclbuilder.FromString(`
		resource "konnect_event_gateway_backend_cluster" "tf_test_egw_backend_cluster" {
			name = "tf_test_egw_backend_cluster"
			authentication = {
    			anonymous = {}
  			}
			bootstrap_servers = [
				"10.0.0.1:8080"
			]
			tls = {
				enabled = false
				insecure_skip_verify = true
			}
		}
		`)
		require.NoError(t, err)

		// Set up dependencies
		egwBackendCluster.AddAttribute("gateway_id", egwCp.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_event_gateway.tf_test_event_gateway", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "name", "tf_test_egw_backend_cluster"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(egwCp).Upsert(egwBackendCluster.AddAttribute("labels", `{key = "value"}`)).Build(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "labels.%", "1"),
						resource.TestCheckResourceAttr("konnect_event_gateway_backend_cluster.tf_test_egw_backend_cluster", "labels.key", "value"),
					),
				},
				{
					Config: builder.Upsert(egwCp).Build(),
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
