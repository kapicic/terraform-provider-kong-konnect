// +build acceptance

package acceptance

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"testing"
	"github.com/kong/internal/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func TestAcckongRouteResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
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

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("kong_route.example", "created_at", 1234),
					resource.TestCheckResourceAttr("kong_route.example", "headers.key", "key"),
					resource.TestCheckResourceAttr("kong_route.example", "hosts.0", "hosts"),
					resource.TestCheckResourceAttr("kong_route.example", "https_redirect_status_code", 1234),
					resource.TestCheckResourceAttr("kong_route.example", "id", "id"),
					resource.TestCheckResourceAttr("kong_route.example", "methods.0", "methods"),
					resource.TestCheckResourceAttr("kong_route.example", "name", "name"),
					resource.TestCheckResourceAttr("kong_route.example", "path_handling", "path_handling"),
					resource.TestCheckResourceAttr("kong_route.example", "paths.0", "paths"),
					resource.TestCheckResourceAttr("kong_route.example", "preserve_host", false),
					resource.TestCheckResourceAttr("kong_route.example", "protocols.0", "protocols"),
					resource.TestCheckResourceAttr("kong_route.example", "regex_priority", 1234),
					resource.TestCheckResourceAttr("kong_route.example", "request_buffering", false),
					resource.TestCheckResourceAttr("kong_route.example", "response_buffering", false),
					resource.TestCheckResourceAttr("kong_route.example", "service.id", "id"),
					resource.TestCheckResourceAttr("kong_route.example", "snis.0", "snis"),
					resource.TestCheckResourceAttr("kong_route.example", "strip_path", false),
					resource.TestCheckResourceAttr("kong_route.example", "tags.0", "tags"),
					resource.TestCheckResourceAttr("kong_route.example", "updated_at", 1234),
					resource.TestCheckResourceAttr("kong_route.example", "runtime_group_id", "runtime_group_id"),
					resource.TestCheckResourceAttr("kong_route.example", "route_id", "route_id"),
				),
			},
		},
	})
}
