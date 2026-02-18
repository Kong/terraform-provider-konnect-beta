resource "konnect-beta_cloud_gateway_add_on" "my_cloudgatewayaddon" {
  config = {
    managed_cache = {
      capacity_config = {
        tiered = {
          tier = "micro"
        }
      }
    }
  }
  name = "my-add-on"
  owner = {
    control_plane_group = {
      control_plane_group_geo = "sg"
      control_plane_group_id  = "123e4567-e89b-12d3-a456-426614174000"
    }
  }
}