resource "konnect_portal_ip_allow_list" "my_portalipallowlist" {
  provider = konnect-beta
  allowed_ips = [
    "192.168.1.1",
    "192.168.1.0/22",
  ]
  portal_id = "f32d905a-ed33-46a3-a093-d8f536af9a8a"
}