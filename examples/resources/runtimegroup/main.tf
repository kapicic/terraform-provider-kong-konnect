resource "kong_run_time_group" "example" {
    name = "name"

    description = "description"

    labels = {
                name = "name"
}


    cluster_type = "cluster_type"

    auth_type = "auth_type"

}
