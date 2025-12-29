package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	testAuthServer = `
resource "konnect_auth_server" "my_authserver" {
  provider = konnect-beta

  name     = "tf-ci-testing-authserver-for-client"
  audience = "local-demo"
}
`
)

func TestAuthServerClients(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	t.Run("should do CRUD when ID is not given", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		authServer, err := hclbuilder.FromString(testAuthServer)
		require.NoError(t, err)

		authServerClient, err := hclbuilder.FromString(`resource "konnect_auth_server_clients" "my_client" {
			provider = konnect-beta

			name             = "my-client-without-id"
			allow_all_scopes = true
			grant_types = [
				"client_credentials",
				"authorization_code",
				"implicit"
			]

			response_types = [
				"id_token",
				"token",
				"code"
			]

			auth_server_id = konnect_auth_server.my_authserver.id
			}
			`)
		require.NoError(t, err)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(authServerClient).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_auth_server.my_authserver", plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction("konnect_auth_server_clients.my_client", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_auth_server_clients.my_client", "name", "my-client-without-id"),
						resource.TestCheckResourceAttrSet("konnect_auth_server_clients.my_client", "id"),
					),
				},
				// Update
				{
					Config: builder.Upsert(authServer).Upsert(authServerClient.AddAttribute("allow_all_scopes", "false")).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_auth_server_clients.my_client", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_auth_server_clients.my_client", "name", "my-client-without-id"),
						resource.TestCheckResourceAttrSet("konnect_auth_server_clients.my_client", "id"),
						resource.TestCheckResourceAttr("konnect_auth_server_clients.my_client", "allow_all_scopes", "false"),
					),
				},
			},
		})
	})

	// To be enabled once the provider supports ID and client secret in input for auth server clients
	t.Run("should do CRUD when ID and client secret are given in input", func(t *testing.T) {
		t.Skip()
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		authServer, err := hclbuilder.FromString(testAuthServer)
		require.NoError(t, err)

		authServerClient, err := hclbuilder.FromString(`resource "konnect_auth_server_clients" "my_client" {
			provider = konnect-beta

			name             = "my-client-without-id"
			allow_all_scopes = true
			grant_types = [
				"client_credentials",
				"authorization_code",
				"implicit"
			]
			client_secret  = "YAzsyUlNZ5gNGeKS9H3VAdxVPzhPo4ae"

			response_types = [
				"id_token",
				"token",
				"code"
			]

			auth_server_id = konnect_auth_server.my_authserver.id
			}
			`)
		require.NoError(t, err)

		uuidResource, err := hclbuilder.FromString(`resource "random_uuid" "example" {}`)
		require.NoError(t, err)

		authServerClient.AddAttribute("id", uuidResource.ResourcePath()+".id")

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(authServer).Upsert(authServerClient).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_auth_server.my_authserver", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_auth_server_clients.my_client", "name", "my-client-without-id"),
						resource.TestCheckResourceAttr("konnect_auth_server_clients.my_client", "id", uuidResource.ResourcePath()+".id"),
					),
				},
			},
		})
	})
}
