resource "kubernetes_namespace" "kong_mesh_system" {
  count = local.should_run ? 1 : 0
  metadata {
    name = "kong-mesh-system"
    labels = {
      "kuma.io/system-namespace" = "true"
    }
  }
}

resource "kubernetes_secret" "mesh_cp_token" {
  count = local.should_run ? 1 : 0
  metadata {
    name = "cp-token"
    namespace = "kong-mesh-system"
  }

  data = {
    token = konnect_system_account_access_token.zone_system_account_token[0].token
  }

  type = "opaque"

  depends_on = [kubernetes_namespace.kong_mesh_system]
}
