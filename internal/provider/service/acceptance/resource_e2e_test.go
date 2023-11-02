//go:build acceptance
// +build acceptance

package acceptance

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/kong/internal/provider"
	"testing"
)

const (
	providerConfig = `
		provider "kong" {
        host = "host"
        auth_token = "auth_token"
}

	`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"kong": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func TestAcckongServiceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
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

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("kong_service.example", "ca_certificates.0", "ca_certificates"),
					resource.TestCheckResourceAttr("kong_service.example", "client_certificate.id", "id"),
					resource.TestCheckResourceAttr("kong_service.example", "connect_timeout", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "created_at", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "enabled", false),
					resource.TestCheckResourceAttr("kong_service.example", "host", "host"),
					resource.TestCheckResourceAttr("kong_service.example", "id", "id"),
					resource.TestCheckResourceAttr("kong_service.example", "name", "name"),
					resource.TestCheckResourceAttr("kong_service.example", "path", "path"),
					resource.TestCheckResourceAttr("kong_service.example", "port", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "protocol", "protocol"),
					resource.TestCheckResourceAttr("kong_service.example", "read_timeout", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "retries", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "tags.0", "tags"),
					resource.TestCheckResourceAttr("kong_service.example", "tls_verify", false),
					resource.TestCheckResourceAttr("kong_service.example", "tls_verify_depth", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "updated_at", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "url", "url"),
					resource.TestCheckResourceAttr("kong_service.example", "write_timeout", 1234),
					resource.TestCheckResourceAttr("kong_service.example", "runtime_group_id", "runtime_group_id"),
					resource.TestCheckResourceAttr("kong_service.example", "service_id", "service_id"),
				),
			},
		},
	})
}
