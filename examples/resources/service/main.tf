resource "kong_service" "example" {
    ca_certificates = [
        "ca_certificates"
    ]

    client_certificate = {
                id = "id"
}


    connect_timeout = "connect_timeout"

    created_at = "created_at"

    enabled = false

    host = "host"

    id = "id"

    name = "name"

    path = "path"

    port = "port"

    protocol = "protocol"

    read_timeout = "read_timeout"

    retries = "retries"

    tags = [
        "tags"
    ]

    tls_verify = false

    tls_verify_depth = "tls_verify_depth"

    updated_at = "updated_at"

    url = "url"

    write_timeout = "write_timeout"

    runtime_group_id = "runtime_group_id"

    service_id = "service_id"

}
