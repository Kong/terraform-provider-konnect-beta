package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMesh(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	t.Run("should not fail on creating a default mesh", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name     = "default"
  type     = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`)
		require.NoError(t, err)

		// Set cp_id and add dependency
		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		// if this grows move this to shared-speakeasy
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Add(cp).Add(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_mesh.default", plancheck.ResourceActionCreate),
						},
					},
				},
			},
		})
	})

	t.Run("create a mesh and modify fields on it", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name     = "default"
  type     = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`)
		require.NoError(t, err)

		// Set cp_id and add dependency
		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		builder.Add(cp)

		resource.ParallelTest(t, hclbuilder.CreateMeshAndModifyFields(providerFactory, builder, mesh))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name = "default"
  type = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`)
		require.NoError(t, err)

		mtp, err := hclbuilder.FromString(`
resource "konnect_mesh_traffic_permission" "allow_all" {
  provider = konnect-beta

  name = "allow-all"
  type = "MeshTrafficPermission"
}
`)
		require.NoError(t, err)

		// Set up dependencies
		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		mtp.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mtp.AddAttribute("mesh", mesh.ResourcePath()+".name")
		mtp.DependsOn(mesh)

		builder.Add(cp)

		resource.ParallelTest(t, hclbuilder.CreatePolicyAndModifyFields(providerFactory, builder, mesh, mtp))
	})

	t.Run("create mesh with oneOf", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "m1" {
  provider = konnect-beta

  name = "m1"
  type = "Mesh"

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
    enabledBackend = "mesh-a-acmpca"
  }
}
`)
		require.NoError(t, err)

		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Add(cp).Add(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_mesh.m1", plancheck.ResourceActionCreate),
						},
					},
				},
				{
					Config: builder.Add(cp).Add(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("create resource with status", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name = "default"
  type = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`)
		require.NoError(t, err)

		mes, err := hclbuilder.FromString(`
resource "konnect_mesh_external_service" "mes_1" {
  provider = konnect-beta

  name = "mes-1"
  type = "MeshExternalService"

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
}
`)
		require.NoError(t, err)

		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		mes.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mes.AddAttribute("mesh", mesh.ResourcePath()+".name")
		mes.DependsOn(mesh)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             builder.Add(cp).Add(mesh).Add(mes).Build(),
					ExpectNonEmptyPlan: true,
				},
			},
		})
	})

	t.Run("create a policy and remove arrays on it", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "e2e-test"
  description = "e2e test cp"
}
`)
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name = "default"
  type = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`)
		require.NoError(t, err)

		mtp1, err := hclbuilder.FromString(`
resource "konnect_mesh_traffic_permission" "allow_all" {
  provider = konnect-beta

  name = "allow-all"
  type = "MeshTrafficPermission"

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
  }
}
`)
		require.NoError(t, err)

		mtp2, err := hclbuilder.FromString(`
resource "konnect_mesh_traffic_permission" "allow_all" {
  provider = konnect-beta

  name = "allow-all"
  type = "MeshTrafficPermission"

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
  }
}
`)
		require.NoError(t, err)

		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		mtp1.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mtp1.AddAttribute("mesh", mesh.ResourcePath()+".name")
		mtp1.DependsOn(mesh)

		mtp2.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mtp2.AddAttribute("mesh", mesh.ResourcePath()+".name")
		mtp2.DependsOn(mesh)

		builder1 := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder1.ProviderProperty = hclbuilder.KonnectBeta
		config1 := builder1.Add(cp).Add(mesh).Add(mtp1).Build()

		builder2 := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder2.ProviderProperty = hclbuilder.KonnectBeta
		config2 := builder2.Add(cp).Add(mesh).Add(mtp2).Build()

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:             config1,
					ExpectNonEmptyPlan: true,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_mesh_traffic_permission.allow_all", plancheck.ResourceActionCreate),
						},
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue("konnect_mesh_traffic_permission.allow_all",
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
				{
					Config:             config1,
					ExpectNonEmptyPlan: true,
				},
				{
					Config:             config2,
					ExpectNonEmptyPlan: true,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_mesh_traffic_permission.allow_all", plancheck.ResourceActionUpdate),
						},
						PostApplyPostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue("konnect_mesh_traffic_permission.allow_all",
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
				{
					Config:             config2,
					ExpectNonEmptyPlan: true,
				},
			},
		})
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "m3"
		mtpName := "allow-all"
		cpName := fmt.Sprintf("e2e-test-%d", acctest.RandInt())

		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(fmt.Sprintf(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "%s"
  description = "e2e test cp"
}
`, cpName))
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(fmt.Sprintf(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name = "%s"
  type = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`, meshName))
		require.NoError(t, err)

		mtp, err := hclbuilder.FromString(fmt.Sprintf(`
resource "konnect_mesh_traffic_permission" "allow_all" {
  provider = konnect-beta

  name = "%s"
  type = "MeshTrafficPermission"
}
`, mtpName))
		require.NoError(t, err)

		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		mtp.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mtp.AddAttribute("mesh", mesh.ResourcePath()+".name")
		mtp.DependsOn(mesh)

		builder.Add(cp)

		resource.ParallelTest(t, hclbuilder.NotImportedResourceShouldError(providerFactory, builder, mesh, mtp, func() { createAnMTP(t, cpName, meshName, mtpName) }))
	})

	t.Run("should be able to store secrets", func(t *testing.T) {
		meshName := "m4"
		cpName := fmt.Sprintf("e2e-test-%d", acctest.RandInt())

		builder := hclbuilder.NewWithProvider(hclbuilder.KonnectBeta, fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.KonnectBeta

		cp, err := hclbuilder.FromString(fmt.Sprintf(`
resource "konnect_mesh_control_plane" "e2e-test" {
  provider = konnect-beta

  name = "%s"
  description = "e2e test cp"
}
`, cpName))
		require.NoError(t, err)

		mesh, err := hclbuilder.FromString(fmt.Sprintf(`
resource "konnect_mesh" "default" {
  provider = konnect-beta

  name = "%s"
  type = "Mesh"

  skip_creating_initial_policies = [ "*" ]
}
`, meshName))
		require.NoError(t, err)

		skey, err := hclbuilder.FromString(`
resource "konnect_mesh_secret" "skey" {
  provider = konnect-beta

  name = "skey"
  type = "Secret"
}
`)
		require.NoError(t, err)

		scert, err := hclbuilder.FromString(`
resource "konnect_mesh_secret" "scert" {
  provider = konnect-beta

  name = "scert"
  type = "Secret"
}
`)
		require.NoError(t, err)

		mesh.AddAttribute("cp_id", cp.ResourcePath()+".id")
		mesh.DependsOn(cp)

		skey.AddAttribute("cp_id", cp.ResourcePath()+".id")
		skey.AddAttribute("mesh", mesh.ResourcePath()+".name")
		skey.DependsOn(mesh)

		scert.AddAttribute("cp_id", cp.ResourcePath()+".id")
		scert.AddAttribute("mesh", mesh.ResourcePath()+".name")
		scert.DependsOn(mesh)

		builder.Add(cp).Add(mesh)

		resource.ParallelTest(t, hclbuilder.ShouldBeAbleToStoreSecrets(providerFactory, builder, mesh, scert, skey))
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
