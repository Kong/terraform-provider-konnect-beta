package tests

import (
    "github.com/Kong/shared-speakeasy/tfbuilder"
    "testing"

    "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMesh(t *testing.T) {
    serverHost, serverPort, serverScheme := providerConfigFromEnv()

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

        resource.ParallelTest(t, tfbuilder.CreatePolicyAndModifyFieldsOnIt(providerFactory, builder, mesh, mtp))
    })
}
