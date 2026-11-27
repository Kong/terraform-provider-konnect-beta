resource "konnect_identity_principal_identity" "my_identityprincipalidentity" {
  provider = konnect-beta
  auth_server_client = {
    auth_server_id = "764ebd98-5bf8-4c1c-862c-0c115cc3fd21"
    client_id      = "...my_client_id..."
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
  }
  control_plane_consumer = {
    consumer_id      = "...my_consumer_id..."
    control_plane_id = "...my_control_plane_id..."
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    workspace_id = "...my_workspace_id..."
  }
  custom = {
    key = "...my_key..."
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
    value = "...my_value..."
  }
  directory_id = "b6e820c3-0070-4ea2-ad6f-3b44c07d409d"
  oidc = {
    claim = {
      name  = "...my_name..."
      value = "...my_value..."
    }
    issuer = "...my_issuer..."
    labels = {
      key = "value"
    }
    managed_by = {
      key = "value"
    }
  }
  principal_id = "b63e8652-e7d8-41ba-ba07-d093efe7352c"
}