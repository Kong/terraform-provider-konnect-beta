resource "konnect_gateway_control_plane" "demo_cp" {
  name         = "New Portal Demo Control Plane"
  description  = "This is a sample description"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
  auth_type    = "pinned_client_certs"

  proxy_urls = []
}

# Configure a service and a route that we can use to test
resource "konnect_gateway_service" "httpbin" {
  name             = "HTTPBin"
  protocol         = "https"
  host             = "httpbin.org"
  port             = 443
  path             = "/"
  control_plane_id = konnect_gateway_control_plane.demo_cp.id
}

resource "konnect_gateway_route" "anything" {
  methods = ["GET"]
  name    = "Anything"
  paths   = ["/anything"]

  strip_path = false

  control_plane_id = konnect_gateway_control_plane.demo_cp.id
  service = {
    id = konnect_gateway_service.httpbin.id
  }
}
