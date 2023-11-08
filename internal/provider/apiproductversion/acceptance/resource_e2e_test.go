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

func TestAcckongApiProductVersionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "kong_api_product_version" "example" {
    name = "name"

    gateway_service = {
                id = "id"
                control_plane_id = "control_plane_id"
}


    publish_status = "publish_status"

    deprecated = false

    notify = false

    api_product_id = "api_product_id"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("kong_api_product_version.example", "name", "name"),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "gateway_service.id", "id"),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "gateway_service.control_plane_id", "control_plane_id"),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "publish_status", "publish_status"),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "deprecated", false),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "notify", false),
					resource.TestCheckResourceAttr("kong_api_product_version.example", "api_product_id", "api_product_id"),
				),
			},
		},
	})
}
