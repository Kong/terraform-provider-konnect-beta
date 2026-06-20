package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/kong/terraform-provider-konnect-beta/internal/provider"
	"github.com/stretchr/testify/require"
)

// providerFactoryBetaOnly contains only konnect-beta, used when ExternalProviders provides konnect
var providerFactoryBetaOnly = map[string]func() (tfprotov6.ProviderServer, error){
	"konnect-beta": providerserver.NewProtocol6WithError(provider.New("")()),
}

func TestIdentityDirectory(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Identity Directory CRUD", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.KonnectBeta,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.KonnectBeta
		controlPlane, err := hclbuilder.FromString(`
			resource "konnect_gateway_control_plane" "test_cp_for_directory" {
				name          = "tf-test-cp-for-identity-directory"
				cloud_gateway = false
			}
		`)
		require.NoError(t, err)

		// Create identity directory with allowed_control_planes referencing the control plane
		identityDirectory, err := hclbuilder.FromString(`
			resource "konnect_identity_directory" "test_directory_with_cp" {
				provider = konnect-beta

				name        = "tf-test-directory-with-allowed-cp"
				description = "Test directory with allowed control planes"

				allow_all_control_planes = true
				allowed_control_planes   = [konnect_gateway_control_plane.test_cp_for_directory.id]
			}
		`)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactoryBetaOnly,
			ExternalProviders: map[string]resource.ExternalProvider{
				"konnect": {Source: "kong/konnect"},
			},
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(controlPlane).Upsert(identityDirectory).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_identity_directory.test_directory_with_cp",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_identity_directory.test_directory_with_cp",
							"name",
							"tf-test-directory-with-allowed-cp",
						),
						resource.TestCheckResourceAttr(
							"konnect_identity_directory.test_directory_with_cp",
							"allow_all_control_planes",
							"true",
						),
						resource.TestCheckResourceAttr(
							"konnect_identity_directory.test_directory_with_cp",
							"allowed_control_planes.#",
							"1",
						),
						resource.TestCheckResourceAttrSet(
							"konnect_identity_directory.test_directory_with_cp",
							"allowed_control_planes.0",
						),
					),
				},
				{
					Config: builder.Upsert(controlPlane).Upsert(identityDirectory).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.Upsert(controlPlane).Upsert(
						identityDirectory.AddAttribute("description", `"Updated directory with control planes"`),
					).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_identity_directory.test_directory_with_cp",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_identity_directory.test_directory_with_cp",
							"description",
							"Updated directory with control planes",
						),
					),
				},
				{
					Config: builder.Upsert(controlPlane).Upsert(identityDirectory).Build(),
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
