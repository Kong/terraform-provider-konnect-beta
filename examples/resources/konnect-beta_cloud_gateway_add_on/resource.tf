resource "konnect-beta_cloud_gateway_add_on" "my_cloudgatewayaddon" {
  config = {
    create_managed_cache_add_on_config = {
      capacity_config = {
        tiered_capacity_config = {
          kind = "tiered"
          tier = "medium"
        }
      }
    }
  }
  name = "my-add-on"
  owner = {
    control_plane_group_add_on_owner = {
      control_plane_group_geo = "us"
      control_plane_group_id  = "123e4567-e89b-12d3-a456-426614174000"
    }
  }
}