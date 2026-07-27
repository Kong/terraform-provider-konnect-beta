resource "konnect_mesh_open_telemetry_backend" "my_meshopentelemetrybackend" {
  provider = konnect-beta
  cp_id = "bf138ba2-c9b1-4229-b268-04d9d8a6410b"
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    endpoint = {
      address = "...my_address..."
      path    = "...my_path..."
      port    = 9
    }
    env = {
      allow_signal_overrides = true
      mode                   = "Optional"
      precedence             = "EnvFirst"
    }
    protocol = "grpc"
  }
  type = "MeshOpenTelemetryBackend"
}