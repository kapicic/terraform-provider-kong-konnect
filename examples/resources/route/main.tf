resource "kong_route" "example" {
    created_at = "created_at"

    headers = {
                key = "key"
}


    hosts = [
        "hosts"
    ]

    https_redirect_status_code = "https_redirect_status_code"

    id = "id"

    methods = [
        "methods"
    ]

    name = "name"

    path_handling = "path_handling"

    paths = [
        "paths"
    ]

    preserve_host = false

    protocols = [
        "protocols"
    ]

    regex_priority = "regex_priority"

    request_buffering = false

    response_buffering = false

    service = {
                id = "id"
}


    snis = [
        "snis"
    ]

    strip_path = false

    tags = [
        "tags"
    ]

    updated_at = "updated_at"

    runtime_group_id = "runtime_group_id"

    route_id = "route_id"

}
