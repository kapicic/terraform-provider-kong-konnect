resource "kong_api_product_version" "example" {
  name = "name"

  gateway_service = {
    id               = "id"
    control_plane_id = "control_plane_id"
  }


  publish_status = "publish_status"

  deprecated = false

  notify = false

  api_product_id = "api_product_id"

}
