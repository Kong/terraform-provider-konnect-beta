# Mesh Example

This scenario creates a new Mesh Control Plane, Mesh and some demo policies.

In addition, if you set the `TF_VAR_k8s_cluster_config_path=~/.kube/config` environment variable, a mesh zone will be deployed to the Kubernetes cluster.

## Test it yourself

```bash
git clone https://github.com/Kong/terraform-provider-konnect-beta
cd examples/scenarios/mesh
terraform init
export TF_VAR_k8s_cluster_config_path="~/.kube/config"
terraform apply
```