package tests

import (
    "github.com/Kong/shared-speakeasy/tfbuilder"
    "github.com/hashicorp/terraform-plugin-testing/plancheck"
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMesh(t *testing.T) {
    serverHost, serverPort, serverScheme := providerConfigFromEnv()

    t.Run("creates a mesh without skip_creating_initial_policies", func(t *testing.T) {
        t.Skip("Skipping test until https://github.com/Kong/shared-speakeasy/pull/4")
        builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
        cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
        mesh := tfbuilder.NewMeshBuilder("m1", "m1").
            WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id")
        builder.AddControlPlane(cp)
        builder.AddMesh(mesh)

        resource.Test(t, resource.TestCase{
            ProtoV6ProviderFactories: providerFactory,
            Steps: []resource.TestStep{
                {
                    Config: builder.Build(),
                    ConfigPlanChecks: resource.ConfigPlanChecks{
                        PreApply: []plancheck.PlanCheck{
                            plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionCreate),
                        },
                    },
                },
                {
                    // Re-apply the same config and ensure no changes occur
                    Config: builder.Build(),
                    ConfigPlanChecks: resource.ConfigPlanChecks{
                        PreApply: []plancheck.PlanCheck{
                            plancheck.ExpectEmptyPlan(),
                        },
                    },
                },
            },
        })
    })

    t.Run("creates a mesh with a policy", func(t *testing.T) {
        builder := tfbuilder.NewBuilder(tfbuilder.Konnect, serverScheme, serverHost, serverPort).WithProviderProperty(tfbuilder.KonnectBeta)
        cp := tfbuilder.NewControlPlane("e2e-test", "e2e-test", "e2e test cp")
        mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh").
            WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
            WithDependsOn(builder.ResourceAddress("mesh_control_plane", cp.ResourceName)).
            WithSpec(`skip_creating_initial_policies = [ "*" ]`)
        mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", "allow-all", "MeshTrafficPermission").
            WithCPID(builder.ResourceAddress("mesh_control_plane", cp.ResourceName) + ".id").
            WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
            WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName)).
            WithLabels(map[string]string{
                "kuma.io/mesh":   mesh.MeshName,
                "kuma.io/env":    "universal",
                "kuma.io/origin": "zone",
                "kuma.io/zone":   "default",
            }).
            WithSpec(tfbuilder.AllowAllTrafficPermissionSpec)
        builder.AddControlPlane(cp)
        builder.AddMesh(mesh)

        resource.Test(t, resource.TestCase{
            ProtoV6ProviderFactories: providerFactory,
            Steps: []resource.TestStep{
                {
                    Config: builder.Build(),
                    ConfigPlanChecks: resource.ConfigPlanChecks{
                        PreApply: []plancheck.PlanCheck{
                            plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionCreate),
                        },
                    },
                },
                {
                    // Re-apply the same config and ensure no changes occur
                    Config: builder.Build(),
                    ConfigPlanChecks: resource.ConfigPlanChecks{
                        PreApply: []plancheck.PlanCheck{
                            plancheck.ExpectEmptyPlan(),
                        },
                    },
                },
                {
                    Config: builder.AddPolicy(mtp).Build(),
                    ConfigPlanChecks: resource.ConfigPlanChecks{
                        PreApply: []plancheck.PlanCheck{
                            plancheck.ExpectResourceAction(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName), plancheck.ResourceActionCreate),
                        },
                    },
                },
                {
                    // Re-apply the same config and ensure no changes occur
                    Config: builder.Build(),
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
