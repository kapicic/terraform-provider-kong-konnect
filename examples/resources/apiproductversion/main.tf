resource "kong_api_product_version" "example" {
  id = "id"

  name = "name"

  gateway_service = {
    id               = "id"
    control_plane_id = "control_plane_id"
  }


  publish_status = "publish_status"

  deprecated = true

  created_at = "created_at"

  updated_at = "updated_at"

  notify = true

  api_product_id = "api_product_id"

}
