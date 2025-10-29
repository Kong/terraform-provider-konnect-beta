package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	resourcesshared "github.com/kong/terraform-provider-kong-mesh/tests/resources/shared"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMesh(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	t.Run("should fail on creating a default mesh", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		mesh := tfbuilder.NewMeshBuilder("default", "default").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)

		builder.AddControlPlane(cp)

		// if this grows move this to shared-speakeasy
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddMesh(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionCreate),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
			},
		})
	})

	t.Run("create a mesh and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		mesh := tfbuilder.NewMeshBuilder("m1", "m1").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)

		builder.AddControlPlane(cp)

		resource.ParallelTest(t, tfbuilder.CreateMeshAndModifyFieldsOnIt(providerFactory, builder, mesh))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		builder.AddControlPlane(cp)
		mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", "allow-all", "MeshTrafficPermission").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)

		resource.ParallelTest(t, tfbuilder.CreatePolicyAndModifyFieldsOnIt(providerFactory, builder, mtp))
	})

	t.Run("create mesh with oneOf", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		mesh := tfbuilder.NewMeshBuilder("m1", "m1").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`
skip_creating_initial_policies = [ "*" ]
mtls = {
  backends = [
    {
      name = "mesh-a-acmpca"
      type = "acmpca"
      dp_cert = {
        rotation = {
          expiration = "24h"
        }
      }
      conf = {
        acm_certificate_authority_config = {
          arn = "arn:hello:world"
          ca_cert = {
            data_source_file = {
              file = "my-file"
            }
          }
          auth = {
            aws_credentials = {
              access_key = {
                data_source_inline_string = {
                  inline_string = "TestTestTest"
                }
              }
              access_key_secret = {
                data_source_inline_string = {
                  inline_string = "TestTestTest"
                }
              }
            }
          }
        }
      }
    }
  ]
  enabledBackend  = "mesh-a-acmpca"
}
`)

		builder.AddControlPlane(cp)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddMesh(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionCreate),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
			},
		})
	})

	t.Run("create resource with status", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		builder.AddControlPlane(cp)
		mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		builder.AddMesh(mesh)

		mes := tfbuilder.NewPolicyBuilder("mesh_external_service", "mes_1", "mes-1", "MeshExternalService").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName)).WithSpec(`
spec = {
  endpoints = [
    {
      address = "example.com"
      port    = 9478
    }
  ]
  match = {
    port     = 1444
    protocol = "tcp"
    type     = "HostnameGenerator"
  }
  tls = {
    allow_renegotiation = false
    enabled             = true
    verification = {
      ca_cert = {
        inline_string = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A...\n-----END CERTIFICATE-----"
      }
      client_cert = {
        inline_string = "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8B...\n-----END CERTIFICATE-----"
      }
      client_key = {
        inline_string = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAt...\n-----END RSA PRIVATE KEY-----"
      }
      mode        = "SkipAll"
      server_name = "my.server.name"
      subject_alt_names = [
        {
          type  = "Exact"
          value = "my.example.com"
        }
      ]
    }
    version = {
      max = "TLS12"
      min = "TLS10"
    }
  }
}
`)
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddPolicy(mes).Build(),
				},
			},
		})
	})

	t.Run("create a policy and remove arrays on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
		builder.AddControlPlane(cp)
		mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", "allow-all", "MeshTrafficPermission").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddPolicy(mtp.WithSpec(`
spec = {
  rules = [
    {
      default = {
        allow = [
          {
            spiffe_id = {
              type  = "Exact"
              value = "spiffe://hello/world"
            }
          }
        ]
      }
    }
  ]
}`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName), plancheck.ResourceActionCreate),
						},
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName),
								tfjsonpath.New("spec").AtMapKey("rules").AtSliceIndex(0).AtMapKey("default"),
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"allow": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"spiffe_id": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"type":  knownvalue.StringExact("Exact"),
												"value": knownvalue.StringExact("spiffe://hello/world"),
											}),
										}),
									}),
									"deny":                   knownvalue.ListExact([]knownvalue.Check{}),
									"allow_with_shadow_deny": knownvalue.ListExact([]knownvalue.Check{}),
								}),
							),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
				{
					Config: builder.AddPolicy(mtp.WithSpec(`
spec = {
  rules = [
    {
      default = {
        deny = [
          {
            spiffe_id = {
              type  = "Exact"
              value = "spiffe://hello/world"
            }
          }
		]
      }
    }
  ]
}`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName), plancheck.ResourceActionUpdate),
						},
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName),
								tfjsonpath.New("spec").AtMapKey("rules").AtSliceIndex(0).AtMapKey("default"),
								knownvalue.ObjectExact(map[string]knownvalue.Check{
									"deny": knownvalue.ListExact([]knownvalue.Check{
										knownvalue.ObjectExact(map[string]knownvalue.Check{
											"spiffe_id": knownvalue.ObjectExact(map[string]knownvalue.Check{
												"type":  knownvalue.StringExact("Exact"),
												"value": knownvalue.StringExact("spiffe://hello/world"),
											}),
										}),
									}),
									"allow_with_shadow_deny": knownvalue.ListExact([]knownvalue.Check{}),
									"allow":                  knownvalue.ListExact([]knownvalue.Check{}),
								}),
							),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
			},
		})
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "m3"
		mtpName := "allow-all"
		cpName := fmt.Sprintf("e2e-test-%d", acctest.RandInt())

		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", cpName, "e2e test cp")
		builder.AddControlPlane(cp)
		mesh := tfbuilder.NewMeshBuilder("default", meshName).
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", mtpName, "MeshTrafficPermission").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)

		resource.ParallelTest(t, tfbuilder.NotImportedResourceShouldErrorOutWithMeaningfulMessage(providerFactory, builder, mtp, func() { createAnMTP(t, cpName, meshName, mtpName) }))
	})

	t.Run("should be able to store secrets", func(t *testing.T) {
		meshName := "m4"
		secretName := "mysecret"
		cpName := fmt.Sprintf("e2e-test-%d", acctest.RandInt())

		builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
		cp := tfbuilder.NewControlPlane("e2e-test", cpName, "e2e test cp")
		builder.AddControlPlane(cp)
		mesh := tfbuilder.NewMeshBuilder("default", meshName).
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithSpec(`skip_creating_initial_policies = [ "*" ]`).
			WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName))
		secret := tfbuilder.NewPolicyBuilder("secret", "mysecret", secretName, "Secret").
			WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)
		resource.ParallelTest(t, resourcesshared.ShouldBeAbleToStoreSecrets(providerFactory, builder, secret, mesh))
	})
}

func createAnMTP(t *testing.T, cpName, meshName, mtpName string) {
	ctx := t.Context()
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	token := os.Getenv("KONNECT_TOKEN")
	if token == "" {
		token = os.Getenv("KONNECT_SPAT")
	}
	require.NotEmpty(t, token)
	client := sdk.New(
		sdk.WithServerURL(fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort)),
		sdk.WithSecurity(shared.Security{PersonalAccessToken: &token}),
	)
	cps, err := client.Mesh.ListMeshControlPlanes(ctx, operations.ListMeshControlPlanesRequest{})
	require.NoError(t, err)
	require.Equal(t, 200, cps.StatusCode)
	require.NotNil(t, cps.ListMeshControlPlanesResponse)

	var myCp shared.MeshControlPlane
	for _, cp := range cps.ListMeshControlPlanesResponse.Data {
		if cp.Name == cpName {
			myCp = cp
			break
		}
	}
	require.NotNil(t, myCp)

	action := shared.MeshTrafficPermissionItemActionAllow
	resp, err := client.MeshTrafficPermission.PutMeshTrafficPermission(ctx, operations.PutMeshTrafficPermissionRequest{
		Mesh: meshName,
		Name: mtpName,
		CpID: myCp.ID,
		MeshTrafficPermissionItem: shared.MeshTrafficPermissionItemInput{
			Mesh: &meshName,
			Name: mtpName,
			Type: shared.MeshTrafficPermissionItemTypeMeshTrafficPermission,
			Spec: shared.MeshTrafficPermissionItemSpec{
				From: []shared.MeshTrafficPermissionItemFrom{
					{
						TargetRef: shared.MeshTrafficPermissionItemSpecTargetRef{Kind: shared.MeshTrafficPermissionItemSpecKindMesh},
						Default:   &shared.MeshTrafficPermissionItemDefault{Action: &action},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
}
