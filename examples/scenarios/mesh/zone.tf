variable "zone_name" {
  type        = string
  description = "The name of the cluster."
  default     = "zone1"
}
resource "helm_release" "kong_mesh" {
  count = local.should_run ? 1 : 0
  name       = "kong-mesh"
  repository = "https://kong.github.io/kong-mesh-charts"
  chart      = "kong-mesh"
  version    = "2.9.3"

  namespace        = "kong-mesh-system"
  upgrade_install = true

  values = [templatefile("values.tftpl", {
    zone_name = var.zone_name,
    region    = var.region,
    cp_id     = konnect_mesh_control_plane.my_meshcontrolplane.id
  })]

  depends_on = [konnect_mesh_control_plane.my_meshcontrolplane, kubernetes_secret.mesh_cp_token]
}

resource "konnect_system_account" "zone_system_account" {
  count = local.should_run ? 1 : 0
  depends_on = [konnect_mesh_control_plane.my_meshcontrolplane]
  name            = "mesh_${konnect_mesh_control_plane.my_meshcontrolplane.id}_${var.zone_name}"
  description     = "Terraform generated system account for authentication zone ${var.zone_name} in ${konnect_mesh_control_plane.my_meshcontrolplane.id} control plane."
  konnect_managed = false
}

resource "konnect_system_account_role" "zone_system_account_role" {
  count = local.should_run ? 1 : 0
  depends_on = [konnect_system_account.zone_system_account]
  account_id       = konnect_system_account.zone_system_account[0].id
  entity_id        = konnect_mesh_control_plane.my_meshcontrolplane.id
  entity_region    = var.region
  entity_type_name = "Mesh Control Planes"
  role_name        = "Admin" # should be "Connector"
}

resource "time_offset" "one_year_from_now" {
  count = local.should_run ? 1 : 0
  offset_years = 1
}

resource "konnect_system_account_access_token" "zone_system_account_token" {
  count = local.should_run ? 1 : 0
  depends_on = [konnect_system_account.zone_system_account]
  account_id = konnect_system_account.zone_system_account[0].id
  expires_at = time_offset.one_year_from_now[0].rfc3339
  name       = konnect_system_account.zone_system_account[0].name
}
