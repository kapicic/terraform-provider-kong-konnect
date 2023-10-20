resource "liblab_runtime_group" "example" {
  id = "id"

  name = "name"

  description = "description"

  labels = {
  }


  config = {
    control_plane_endpoint = "control_plane_endpoint"
    telemetry_endpoint     = "telemetry_endpoint"
  }


  created_at = "created_at"

  updated_at = "updated_at"

  cluster_type = "cluster_type"

  auth_type = "auth_type"

}
