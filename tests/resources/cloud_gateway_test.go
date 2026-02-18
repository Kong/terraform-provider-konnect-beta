package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider"
	"github.com/stretchr/testify/require"
)

var (
	providerFactoryBeta = map[string]func() (tfprotov6.ProviderServer, error){
		"konnect-beta": providerserver.NewProtocol6WithError(provider.New("")()),
	}
)

func TestCloudGatewayAddOn(t *testing.T) {
	t.Run("Cloud Gateways AddOns", func(t *testing.T) {
		builder := hclbuilder.New()
		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "test_cp" {
             name         = "tf-test-cp-us-external"
             cloud_gateway = true
          }
       `)
		require.NoError(t, err)
		addon, err := hclbuilder.FromString(`
         resource "konnect-beta_cloud_gateway_add_on" "my_addon" {
          provider = konnect-beta
          name     = "tf-test-add-on"

          config = {
            create_managed_cache_add_on_config = {
             kind = "managed-cache.v0"

             capacity_config = {
               tiered_capacity_config = {
                kind = "tiered"
                tier = "small"
               }
             }
            }
          }
          owner = {
            control_plane_add_on_owner = {
             kind             = "control-plane"
             control_plane_geo = "us"
             control_plane_id  = konnect_gateway_control_plane.test_cp.id
            }
          }
         }
       `)
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(cp).
			Upsert(addon).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactoryBeta,
			ExternalProviders: map[string]resource.ExternalProvider{
				"konnect": {Source: "kong/konnect"},
			},
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect-beta_cloud_gateway_add_on.my_addon",
								plancheck.ResourceActionCreate,
							),
						},
					},
				},
				{
					Config: fullConfig,
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
