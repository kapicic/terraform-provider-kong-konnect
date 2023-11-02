resource "kong_api_product_version" "example" {
    name = "name"

    publish_status = "publish_status"

    deprecated = false

    gateway_service = {
                id = "id"
                control_plane_id = "control_plane_id"
}


    notify = false

    api_product_id = "api_product_id"

}
