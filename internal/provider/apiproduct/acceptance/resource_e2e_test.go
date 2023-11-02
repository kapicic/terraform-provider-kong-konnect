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

func TestAcckongApiProductResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "kong_api_product" "example" {
    name = "name"

    description = "description"

    labels = {
                name = "name"
}


}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("kong_api_product.example", "name", "name"),
					resource.TestCheckResourceAttr("kong_api_product.example", "description", "description"),
					resource.TestCheckResourceAttr("kong_api_product.example", "labels.name", "name"),
				),
			},
		},
	})
}
