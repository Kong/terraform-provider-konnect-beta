resource "konnect_cloud_gateway_addon" "my_cloudgatewayaddon" {
  provider = konnect-beta
  config = {
    managed_cache = {
      capacity_config = {
        tiered = {
          tier = "2xlarge"
        }
      }
    }
  }
  name = "my-add-on"
  owner = {
    control_plane_group = {
      control_plane_group_geo = "me"
      control_plane_group_id  = "123e4567-e89b-12d3-a456-426614174000"
    }
  }
}